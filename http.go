package getstream

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// logRequest logs the details of an HTTP request
func (c *Client) logRequest(req *http.Request) {
	c.logger.Debug("---> %s %s", req.Method, req.URL.String())
	c.logger.Debug("Host: %s", req.Host)
	for key, values := range req.Header {
		c.logger.Debug("%s: %s", key, strings.Join(values, ", "))
	}
	if req.Body != nil {
		var buf bytes.Buffer
		body, _ := io.ReadAll(req.Body)
		buf.Write(body)
		req.Body = io.NopCloser(&buf)
		c.logger.Debug("\n%s", string(body))
	}
}

// logResponse logs the details of an HTTP response
func (c *Client) logResponse(resp *http.Response, body []byte, duration time.Duration) {
	c.logger.Debug("<--- %d %s (%s)", resp.StatusCode, http.StatusText(resp.StatusCode), duration)
	for key, values := range resp.Header {
		c.logger.Debug("%s: %s", key, strings.Join(values, ", "))
	}
	c.logger.Debug("\n%s", string(body))
}

// Error represents an API error
type StreamError struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	ExceptionFields map[string]string `json:"exception_fields,omitempty"`
	StatusCode      int               `json:"StatusCode"`
	Duration        string            `json:"duration"`
	MoreInfo        string            `json:"more_info"`
	RateLimit       *RateLimitInfo    `json:"-"`
}

func (e StreamError) Error() string {
	return e.Message
}

// Response is the base response returned to the client
type StreamResponse[T any] struct {
	RateLimitInfo *RateLimitInfo `json:"ratelimit"`
	Data          T
}

// parseResponse parses the HTTP response into the provided result
func parseResponse[GResponse any](c *Client, resp *http.Response, body []byte, result *GResponse) (*StreamResponse[GResponse], error) {
	statusCode := resp.StatusCode
	c.logger.Debug("Status Code: %d", statusCode)
	// If status code indicates an error
	if statusCode >= 399 {
		var apiErr StreamError
		err := json.Unmarshal(body, &apiErr)
		if err != nil {
			apiErr.Message = string(body)
			apiErr.StatusCode = resp.StatusCode
			return nil, apiErr
		}
		apiErr.RateLimit = NewRateLimitFromHeaders(resp.Header)
		return nil, apiErr
	}

	// Attempt to unmarshal the response into the result
	err := json.Unmarshal(body, result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// Add rate limit info to the result
	return addRateLimitInfo(resp.Header, result)
}

// requestURL constructs the full request URL
func (c *Client) requestURL(path string, values url.Values, pathParams map[string]string) (string, error) {
	path = buildPath(path, pathParams)

	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return "", fmt.Errorf("url.Parse: %w", err)
	}

	if values == nil {
		values = make(url.Values)
	}

	values.Add("api_key", c.apiKey)
	c.logger.Debug("Query parameters: %v", values)
	u.RawQuery = values.Encode()
	url := u.String()
	c.logger.Debug("Full URL: %s", url)
	return url, nil
}

// buildPath constructs a URL path with parameters, escaping them appropriately.
func buildPath(path string, pathParams map[string]string) string {
	if pathParams == nil {
		return path
	}
	for k, v := range pathParams {
		pathParams[k] = url.QueryEscape(v)
	}
	return replaceParams(path, pathParams)
}

// replaceParams replaces placeholders in the path with the corresponding values from pathParams.
func replaceParams(path string, pathParams map[string]string) string {
	for k, v := range pathParams {
		path = strings.ReplaceAll(path, "{"+k+"}", v)
	}
	return path
}

// newRequest creates a new HTTP request
func newRequest[T any](c *Client, ctx context.Context, method, path string, params url.Values, data T, pathParams map[string]string) (*http.Request, error) {
	u, err := c.requestURL(path, params, pathParams)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, method, u, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(r)

	// Do not set body if the method is GET
	if method == http.MethodGet {
		r.Body = nil
		c.logger.Debug("GET request: No body set")
		return r, nil
	}

	// Handle other methods with body
	c.logger.Debug("Method: %s, Data: %#v (Type: %T)", method, data, data)
	switch t := any(data).(type) {
	case nil:
		c.logger.Debug("Data is nil")
		r.Body = nil
	case io.ReadCloser:
		c.logger.Debug("Data is io.ReadCloser")
		r.Body = t
	case io.Reader:
		c.logger.Debug("Data is io.Reader")
		r.Body = io.NopCloser(t)
	default:
		c.logger.Debug("Data is of type %T, attempting to marshal to JSON", t)
		b, err := json.Marshal(data)
		if err != nil {
			c.logger.Error("Error marshaling data: %v", err)
			c.logger.Warn("Unable to marshal data, setting body to nil")
			r.Body = nil
		} else {
			r.Body = io.NopCloser(bytes.NewReader(b))
			c.logger.Debug("Request body set with JSON: %s", string(b))
		}
	}

	return r, nil
}

// isNil checks if a generic value is nil using reflection
func isNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}

// setHeaders sets necessary headers for the request
func (c *Client) setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Stream-Client", versionHeader())
	r.Header.Set("Authorization", c.authToken)
	r.Header.Set("Stream-Auth-Type", "jwt")
}

func StructToMapWithTags(input any, tagName string) (map[string]any, error) {
	result := make(map[string]any)
	err := extractFields(reflect.ValueOf(input), tagName, result)
	return result, err
}

// Recursive function to extract fields
func extractFields(val reflect.Value, tagName string, result map[string]any) error {
	// Check if the input is a pointer and get the actual value
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Ensure the provided input is a struct
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("input must be a struct or a pointer to a struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)

		// Check if the field is an embedded struct and extract its fields
		if structField.Anonymous && field.Kind() == reflect.Struct {

			err := extractFields(field, tagName, result)
			if err != nil {
				return err
			}
			continue
		}

		// Check for the 'path' tag
		if tag, ok := structField.Tag.Lookup(tagName); ok {
			result[tag] = field.Interface()
		}
	}
	return nil
}

func extractQueryParams(v any) url.Values {
	if v == nil || reflect.ValueOf(v).IsNil() {
		return url.Values{}
	}
	m, err := StructToMapWithTags(v, "query")
	if err != nil {
		panic(err)
	}
	values := url.Values{}
	for k, v := range m {
		values.Set(k, EncodeValueToQueryParam(v))
	}
	return values
}

// EncodeValueToQueryParam returns the string representation of a value ready to be used as a query param
func EncodeValueToQueryParam(value any) string {
	val := reflect.ValueOf(value)

	if val.Kind() == reflect.Ptr && val.IsNil() {
		return ""
	}

	switch val.Kind() {
	case reflect.Ptr:
		return EncodeValueToQueryParam(val.Elem().Interface())
	case reflect.String:
		return val.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())
	case reflect.Map, reflect.Struct, reflect.Slice:
		b, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		return string(b)
	default:
		return fmt.Sprintf("%v", val.Interface())
	}
}

// makeRequest makes a generic HTTP request
func MakeRequest[GRequest any, GResponse any](c *Client, ctx context.Context, method, path string, params url.Values, data *GRequest, response *GResponse, pathParams map[string]string) (*StreamResponse[GResponse], error) {
	r, err := newRequest(c, ctx, method, path, params, data, pathParams)
	if err != nil {
		return nil, err
	}

	c.logRequest(r)

	start := time.Now()
	resp, err := c.HTTP.Do(r)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTTP response: %w", err)
	}

	duration := time.Since(start)
	c.logResponse(resp, b, duration)

	return parseResponse(c, resp, b, response)
}

// TODO: revisit this
// addRateLimitInfo adds rate limit information to the result
func addRateLimitInfo[Gresponse any](headers http.Header, result *Gresponse) (*StreamResponse[Gresponse], error) {
	rateLimit := NewRateLimitFromHeaders(headers)
	return &StreamResponse[Gresponse]{RateLimitInfo: rateLimit, Data: *result}, nil
}

// versionHeader returns the version header (implementation omitted for brevity)
func (c *Client) version() string {
	return versionHeader()
}
