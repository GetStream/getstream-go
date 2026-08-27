package getstream

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// StreamError is the single concrete error type returned by the SDK.
//
// Category is signaled by the sentinel embedded via Is: callers branch with
// errors.Is(err, ErrApiResponse | ErrRateLimited | ErrTransport | ErrTaskFailed)
// and extract fields with errors.As(err, &streamErr).
type StreamError struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	ExceptionFields map[string]string `json:"exception_fields,omitempty"`
	StatusCode      int               `json:"StatusCode"`
	Duration        string            `json:"duration"`
	MoreInfo        string            `json:"more_info"`
	// Unrecoverable mirrors APIError.unrecoverable. When true, the request
	// that produced this error must not be retried.
	Unrecoverable bool `json:"unrecoverable,omitempty"`
	// Details carries the opaque APIError.details payload verbatim. nil if
	// the backend omitted the field.
	Details json.RawMessage `json:"details,omitempty"`
	// RawResponseBody is the unparsed response body. Always set on API-response
	// errors (including the unparseable-body case).
	RawResponseBody string `json:"-"`
	// RetryAfter is the parsed Retry-After header on HTTP 429. Zero otherwise.
	RetryAfter time.Duration `json:"-"`
	// ErrorType is populated only when the sentinel is ErrTransport. One of
	// ErrorTypeConnectionReset, ErrorTypeTimeout, ErrorTypeDNSFailure,
	// ErrorTypeTLSHandshake, ErrorTypeUnknown.
	ErrorType string `json:"-"`
	// Task carries the failed-task payload when the sentinel is ErrTaskFailed.
	Task *TaskErrorDetails `json:"-"`
	// RateLimit carries the rate-limit window info from response headers.
	RateLimit *RateLimitInfo `json:"-"`

	// sentinel selects the category surfaced via Is. cause is the wrapped
	// underlying error (typically a stack-bearing wrapper).
	sentinel error
	cause    error
}

func (e *StreamError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.Message
}

// Is reports whether target matches the StreamError's category sentinel.
// ErrRateLimited additionally matches ErrApiResponse.
func (e *StreamError) Is(target error) bool {
	if e == nil || target == nil {
		return false
	}
	if e.sentinel != nil && target == e.sentinel {
		return true
	}
	if e.sentinel == ErrRateLimited && target == ErrApiResponse {
		return true
	}
	return false
}

// Unwrap returns the underlying cause (typically a stack-bearing wrapper
// over the original transport error or JSON-parse error). Returns nil for
// API-response errors that have no upstream cause.
func (e *StreamError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

// Sentinel reports the category sentinel this error belongs to (e.g.
// ErrApiResponse). Returns nil if no category was set.
func (e *StreamError) Sentinel() error {
	if e == nil {
		return nil
	}
	return e.sentinel
}

// Response is the base response returned to the client
type StreamResponse[T any] struct {
	RateLimitInfo *RateLimitInfo `json:"ratelimit"`
	Data          T
}

// parseResponse parses the HTTP response into the provided result.
// On HTTP 4xx/5xx returns a *StreamError populated from the APIError
// envelope, or a sentinel-message StreamError when the body cannot be parsed.
func parseResponse[GResponse any](c *Client, resp *http.Response, body []byte, result *GResponse) (*StreamResponse[GResponse], error) {
	statusCode := resp.StatusCode
	if statusCode >= 399 {
		return nil, buildAPIError(resp, body)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return nil, stackWrap(err, "failed to unmarshal response body")
	}

	return addRateLimitInfo(resp.Header, result)
}

// buildAPIError constructs a *StreamError for an HTTP 4xx/5xx response.
// Returns *StreamError populated from the APIError envelope, or a
// sentinel-message StreamError when the body cannot be parsed.
func buildAPIError(resp *http.Response, body []byte) *StreamError {
	apiErr := &StreamError{
		StatusCode:      resp.StatusCode,
		RawResponseBody: string(body),
		ExceptionFields: map[string]string{},
		RateLimit:       NewRateLimitFromHeaders(resp.Header),
	}

	if len(body) > 0 {
		if err := json.Unmarshal(body, apiErr); err != nil {
			// A failed Unmarshal can still have assigned fields before it errored,
			// so restore the transport status rather than trust the partial parse.
			apiErr.StatusCode = resp.StatusCode
			apiErr.Code = 0
			apiErr.Message = fmt.Sprintf("failed to parse error response: unexpected server response code %d", resp.StatusCode)
			apiErr.ExceptionFields = map[string]string{}
			apiErr.Unrecoverable = false
			apiErr.cause = stackWrap(err, "parse api error response")
			apiErr.sentinel = ErrApiResponse
			return apiErr
		}
	} else {
		apiErr.Message = "empty response body"
	}

	if apiErr.StatusCode == 0 {
		apiErr.StatusCode = resp.StatusCode
	}
	if apiErr.ExceptionFields == nil {
		apiErr.ExceptionFields = map[string]string{}
	}

	if apiErr.StatusCode == http.StatusTooManyRequests {
		apiErr.sentinel = ErrRateLimited
		apiErr.RetryAfter = parseRetryAfter(resp.Header.Get("Retry-After"), time.Now())
	} else {
		apiErr.sentinel = ErrApiResponse
	}

	return apiErr
}

// requestURL constructs the full request URL
func (c *Client) requestURL(path string, values url.Values, pathParams map[string]string) (string, error) {
	path = buildPath(path, pathParams)

	u, err := url.Parse(c.baseUrl + path)
	if err != nil {
		return "", stackWrap(err, "url.Parse")
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

	// Do not set body if the method is GET
	if method == http.MethodGet {
		r.Body = nil
		return r, nil
	}

	switch t := any(data).(type) {
	case nil:
		r.Body = nil
	case io.ReadCloser:
		r.Body = t
	case io.Reader:
		r.Body = io.NopCloser(t)
	case *UploadFileRequest, *UploadImageRequest, *UploadChannelFileRequest, *UploadChannelImageRequest:
		return c.createMultipartRequest(r, t)
	default:
		b, err := json.Marshal(data)
		if err != nil {
			c.logger.Error("Error marshaling data: %+v, setting body to nil", err)
			r.Body = nil
		} else {
			setRetryableBody(r, b)
		}
	}

	return r, nil
}

// setRetryableBody sets an in-memory body with ContentLength and GetBody. GetBody
// lets the HTTP/2 transport retry the request on GOAWAY/REFUSED_STREAM, which it
// can't do without a rewindable body.
func setRetryableBody(r *http.Request, b []byte) {
	r.ContentLength = int64(len(b))
	r.Body = io.NopCloser(bytes.NewReader(b))
	r.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(b)), nil
	}
}

func getFileContent(fileName string, fileContent io.Reader) (io.Reader, error) {
	if fileContent != nil {
		return fileContent, nil
	}
	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, stackWrap(err, "failed to open file")
		}
		return file, nil
	}
	return nil, fmt.Errorf("either file name or file content must be provided")
}

// createMultipartRequest creates a multipart form request for file/image uploads
func (c *Client) createMultipartRequest(r *http.Request, data any) (*http.Request, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	var fileContent io.Reader
	var fileName string
	var err error
	// Handle both UploadFileRequest and UploadImageRequest
	switch req := data.(type) {
	case *UploadFileRequest:
		if req.File == nil {
			return nil, fmt.Errorf("file name must be provided")
		}
		fileName = *req.File
		fileContent, err = getFileContent(*req.File, nil)
		if err != nil {
			return nil, stackWrap(err, "failed to open file")
		}

		// Add user field if present
		if req.User != nil {
			userJSON, err := json.Marshal(req.User)
			if err != nil {
				return nil, stackWrap(err, "failed to marshal user")
			}
			err = writer.WriteField("user", string(userJSON))
			if err != nil {
				return nil, stackWrap(err, "failed to write user field")
			}
		}

	case *UploadImageRequest:
		if req.File == nil {
			return nil, fmt.Errorf("file name must be provided")
		}
		fileName = *req.File
		fileContent, err = getFileContent(*req.File, nil)
		if err != nil {
			return nil, stackWrap(err, "failed to open file")
		}

		// Add upload_sizes field if present
		if req.UploadSizes != nil && len(*req.UploadSizes) > 0 {
			uploadSizesJSON, err := json.Marshal(*req.UploadSizes)
			if err != nil {
				return nil, stackWrap(err, "failed to marshal upload_sizes")
			}
			err = writer.WriteField("upload_sizes", string(uploadSizesJSON))
			if err != nil {
				return nil, stackWrap(err, "failed to write upload_sizes field")
			}
		}

		// Add user field if present
		if req.User != nil {
			userJSON, err := json.Marshal(req.User)
			if err != nil {
				return nil, stackWrap(err, "failed to marshal user")
			}
			err = writer.WriteField("user", string(userJSON))
			if err != nil {
				return nil, stackWrap(err, "failed to write user field")
			}
		}

	case *UploadChannelFileRequest:
		if req.File == nil {
			return nil, fmt.Errorf("file name must be provided")
		}
		fileName = *req.File
		fileContent, err = getFileContent(*req.File, nil)
		if err != nil {
			return nil, stackWrap(err, "failed to open file")
		}

		// Add user field if present
		if req.User != nil {
			userJSON, err := json.Marshal(req.User)
			if err != nil {
				return nil, stackWrap(err, "failed to marshal user")
			}
			err = writer.WriteField("user", string(userJSON))
			if err != nil {
				return nil, stackWrap(err, "failed to write user field")
			}
		}

	case *UploadChannelImageRequest:
		if req.File == nil {
			return nil, fmt.Errorf("file name must be provided")
		}
		fileName = *req.File
		fileContent, err = getFileContent(*req.File, nil)
		if err != nil {
			return nil, stackWrap(err, "failed to open file")
		}

		// Add upload_sizes field if present
		if req.UploadSizes != nil && len(*req.UploadSizes) > 0 {
			uploadSizesJSON, err := json.Marshal(*req.UploadSizes)
			if err != nil {
				return nil, stackWrap(err, "failed to marshal upload_sizes")
			}
			err = writer.WriteField("upload_sizes", string(uploadSizesJSON))
			if err != nil {
				return nil, stackWrap(err, "failed to write upload_sizes field")
			}
		}

		// Add user field if present
		if req.User != nil {
			userJSON, err := json.Marshal(req.User)
			if err != nil {
				return nil, stackWrap(err, "failed to marshal user")
			}
			err = writer.WriteField("user", string(userJSON))
			if err != nil {
				return nil, stackWrap(err, "failed to write user field")
			}
		}

	default:
		return nil, fmt.Errorf("unsupported request type for multipart: %T", data)
	}

	// Add file field
	fileWriter, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, stackWrap(err, "failed to create form file")
	}

	_, err = io.Copy(fileWriter, fileContent)
	if err != nil {
		return nil, stackWrap(err, "failed to copy file content")
	}

	err = writer.Close()
	if err != nil {
		return nil, stackWrap(err, "failed to close multipart writer")
	}

	// buf isn't mutated after this, so sharing its slice is safe.
	setRetryableBody(r, buf.Bytes())
	r.Header.Set("Content-Type", writer.FormDataContentType())

	c.logger.Debug("Created multipart request with file: %s", fileName)
	return r, nil
}

// setHeaders sets necessary headers for the request
func (c *Client) setHeaders(r *http.Request) {
	if r.Header.Get("Content-Type") == "" {
		r.Header.Set("Content-Type", "application/json")
	}
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
		value := reflect.ValueOf(v)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			continue
		}
		values.Set(k, EncodeValueToQueryParam(v))
	}
	return values
}

// EncodeValueToQueryParam returns the string representation of a value ready to be used as a query param
func EncodeValueToQueryParam(value any) string {
	val := reflect.ValueOf(value)

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
	case reflect.Slice:
		// For query params, slices of primitives should be comma-separated (e.g. ids=a,b,c)
		parts := make([]string, val.Len())
		for i := 0; i < val.Len(); i++ {
			parts[i] = EncodeValueToQueryParam(val.Index(i).Interface())
		}
		return strings.Join(parts, ",")
	case reflect.Map, reflect.Struct:
		b, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		return string(b)
	default:
		return fmt.Sprintf("%v", val.Interface())
	}
}

// MakeRequest makes a generic HTTP request, auto-retrying per the client's
// opt-in RetryConfig (GET/HEAD on 429/transport errors only). Disabled by
// default: exactly one attempt, errors surface unchanged.
func MakeRequest[GRequest any, GResponse any](c *Client, ctx context.Context, method, path string, params url.Values, data *GRequest, response *GResponse, pathParams map[string]string) (*StreamResponse[GResponse], error) {
	for attempt := 0; ; attempt++ {
		start := time.Now()
		result, err := makeRequestOnce(c, ctx, method, path, params, data, response, pathParams)
		if err == nil {
			return result, nil
		}
		if !c.shouldRetry(err, method, attempt) {
			// Only transport failures ever reached http.request.failed before
			// retry existed (4xx/5xx, including 429, log via
			// http.response.received instead) — preserve that split here.
			if errors.Is(err, ErrTransport) {
				c.logRequestFailed(method, path, err, time.Since(start))
			}
			return result, err
		}
		delay := c.retryDelay(err, attempt)
		c.logRetryScheduled(method, path, err, attempt+1, delay)
		select {
		case <-ctx.Done():
			ctxErr := wrapTransportError(ctx.Err())
			c.logRequestFailed(method, path, ctxErr, time.Since(start))
			return nil, ctxErr
		case <-time.After(delay):
		}
	}
}

// makeRequestOnce performs a single HTTP attempt: builds the request, sends
// it, and parses the response. Callers (MakeRequest) own retry looping and
// the http.request.failed emission for transport failures.
func makeRequestOnce[GRequest any, GResponse any](c *Client, ctx context.Context, method, path string, params url.Values, data *GRequest, response *GResponse, pathParams map[string]string) (*StreamResponse[GResponse], error) {
	r, err := newRequest(c, ctx, method, path, params, data, pathParams)
	if err != nil {
		return nil, err
	}

	// Only re-marshal on opt-in (WithLogBodies) so the default path does zero
	// extra work. GET/HEAD never carry a body.
	var reqBody []byte
	if c.logBodies && data != nil && method != http.MethodGet && method != http.MethodHead {
		reqBody, _ = json.Marshal(data)
	}
	// r.URL.Query() is the actual built query (includes the api_key requestURL
	// injects), not the caller's params — params alone would log an empty
	// query for the ~246 of 316 call sites that pass nil.
	c.logRequestSent(method, path, r.URL.Query(), reqBody)

	start := time.Now()
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, wrapTransportError(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapTransportError(err)
	}

	duration := time.Since(start)
	c.logResponseReceived(method, path, resp.StatusCode, len(b), duration, b)

	return parseResponse(c, resp, b, response)
}

// addRateLimitInfo adds rate limit information to the result
func addRateLimitInfo[Gresponse any](headers http.Header, result *Gresponse) (*StreamResponse[Gresponse], error) {
	rateLimit := NewRateLimitFromHeaders(headers)
	return &StreamResponse[Gresponse]{RateLimitInfo: rateLimit, Data: *result}, nil
}
