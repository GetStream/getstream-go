package getstream

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Error represents an API error
type Error struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	ExceptionFields map[string]string `json:"exception_fields,omitempty"`
	StatusCode      int               `json:"StatusCode"`
	Duration        string            `json:"duration"`
	MoreInfo        string            `json:"more_info"`
	RateLimit       *RateLimitInfo    `json:"-"`
}

func (e Error) Error() string {
	return e.Message
}

// Response is the base response returned to the client
type Response struct {
	RateLimitInfo *RateLimitInfo `json:"ratelimit"`
}

// BuildQueryParam constructs a map of query parameters from various data types.
func BuildQueryParam[T any](params map[string]T) url.Values {
	values := url.Values{}
	for key, value := range params {
		switch v := any(value).(type) {
		case string:
			values.Add(key, v)
		case int:
			values.Add(key, strconv.Itoa(v))
		case int32, int64:
			if converted, ok := any(v).(int64); ok {
				values.Add(key, strconv.FormatInt(converted, 10))
			}
		case uint, uint32, uint64:
			if converted, ok := any(v).(uint64); ok {
				values.Add(key, strconv.FormatUint(converted, 10))
			}
		case float32:
			values.Add(key, strconv.FormatFloat(float64(v), 'f', -1, 32))
		case float64:
			values.Add(key, strconv.FormatFloat(v, 'f', -1, 64))
		default:
			// Attempt to marshal as JSON for any other types, including structs
			jsonData, err := json.Marshal(v)
			if err == nil {
				values.Add(key, string(jsonData))
			}
		}
	}
	return values
}

// parseResponse parses the HTTP response into the provided result
func parseResponse[U any](resp *http.Response, result *U) error {
	if resp.Body == nil {
		return errors.New("http body is nil")
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read HTTP response: %w", err)
	}

	if resp.StatusCode >= 399 {
		var apiErr Error
		err := json.Unmarshal(b, &apiErr)
		if err != nil {
			apiErr.Message = string(b)
			apiErr.StatusCode = resp.StatusCode
			return apiErr
		}
		apiErr.RateLimit = NewRateLimitFromHeaders(resp.Header)
		return apiErr
	}

	// unmarshal result
	err = json.Unmarshal(b, result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal HTTP response: %w", err)
	}
	return nil

	// return addRateLimitInfo(resp.Header, result)
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
	u.RawQuery = values.Encode()

	return u.String(), nil
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

	switch t := any(data).(type) {
	case nil:
		r.Body = nil
	case io.ReadCloser:
		r.Body = t
	case io.Reader:
		r.Body = io.NopCloser(t)
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		r.Body = io.NopCloser(bytes.NewReader(b))
	}

	return r, nil
}

// setHeaders sets necessary headers for the request
func (c *Client) setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Stream-Client", versionHeader())
	r.Header.Set("Authorization", c.authToken)
	r.Header.Set("Stream-Auth-Type", "jwt")
}

func ToMap(item interface{}) (map[string]interface{}, error) {
	bytes, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// makeRequest makes a generic HTTP request
func MakeRequest[GRequest any, GResponse any, GParams any](c *Client, ctx context.Context, method, path string, params map[string]GParams, data *GRequest, response *GResponse, pathParams map[string]string) error {
	queryParams := BuildQueryParam(params)
	r, err := newRequest(c, ctx, method, path, queryParams, data, pathParams)
	if err != nil {
		return err
	}

	resp, err := c.HTTP.Do(r)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return err
	}

	return parseResponse(resp, response)
}

// TODO: revisit this
// addRateLimitInfo adds rate limit information to the result
func addRateLimitInfo[U any](headers http.Header, result *U) error {
	rl := map[string]interface{}{
		"ratelimit": NewRateLimitFromHeaders(headers),
	}

	b, err := json.Marshal(rl)
	if err != nil {
		return fmt.Errorf("cannot marshal rate limit info: %w", err)
	}

	err = json.Unmarshal(b, result)
	if err != nil {
		return fmt.Errorf("cannot unmarshal rate limit info: %w", err)
	}
	return nil
}

// versionHeader returns the version header (implementation omitted for brevity)
func (c *Client) version() string {
	return versionHeader()
}
