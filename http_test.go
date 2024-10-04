package getstream

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPath(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		pathParams map[string]string
		want       string
	}{
		{
			name:       "No parameters",
			path:       "/api/resource",
			pathParams: nil,
			want:       "/api/resource",
		},
		{
			name: "With parameters",
			path: "/api/{resource}/{id}",
			pathParams: map[string]string{
				"resource": "user",
				"id":       "123",
			},
			want: "/api/user/123",
		},
		{
			name: "Escaped characters",
			path: "/api/{query}",
			pathParams: map[string]string{
				"query": "special char/=&%",
			},
			want: "/api/special+char%2F%3D%26%25",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildPath(tt.path, tt.pathParams); got != tt.want {
				t.Errorf("buildPath(%q, %v) = %q, want %q", tt.path, tt.pathParams, got, tt.want)
			}
		})
	}
}

func TestExtractQueryParams(t *testing.T) {
	t.Run("Extract query params from GetCallRequest", func(t *testing.T) {
		request := &GetCallRequest{
			MembersLimit: PtrTo(10),
			Notify:       PtrTo(true),
			Ring:         PtrTo(false),
			Video:        PtrTo(true),
		}

		expected := url.Values{
			"members_limit": []string{"10"},
			"notify":        []string{"true"},
			"ring":          []string{"false"},
			"video":         []string{"true"},
		}

		result := extractQueryParams(request)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("extractQueryParams() = %v, want %v", result, expected)
		}
	})

	t.Run("Extract query params from nil request", func(t *testing.T) {
		result := extractQueryParams(nil)

		if len(result) != 0 {
			t.Errorf("extractQueryParams(nil) = %v, want empty url.Values", result)
		}
	})
}

func TestRequestURL(t *testing.T) {
	originalBaseURL := "https://api.example.com"
	mockLogger := DefaultLogger
	client := &Client{
		baseUrl: originalBaseURL,
		apiKey:  "testkey",
		logger:  mockLogger,
	}

	t.Run("Valid_URL_without_path_parameters", func(t *testing.T) {
		path := "/v1/resources"
		values := url.Values{
			"param1": {"value1"},
			"param2": {"value2"},
		}

		expectedURL := "https://api.example.com/v1/resources?api_key=testkey&param1=value1&param2=value2"

		got, err := client.requestURL(path, values, nil)
		if err != nil {
			t.Fatalf("requestURL returned error: %v", err)
		}

		if got != expectedURL {
			t.Errorf("requestURL() = %q, want %q", got, expectedURL)
		}
	})

	t.Run("Valid_URL_with_path_parameters", func(t *testing.T) {
		path := "/v1/resources/{id}"
		pathParams := map[string]string{
			"id": "123",
		}
		values := url.Values{}

		expectedURL := "https://api.example.com/v1/resources/123?api_key=testkey"

		got, err := client.requestURL(path, values, pathParams)
		if err != nil {
			t.Fatalf("requestURL returned error: %v", err)
		}

		if got != expectedURL {
			t.Errorf("requestURL() = %q, want %q", got, expectedURL)
		}
	})

	t.Run("Invalid_BaseURL", func(t *testing.T) {
		invalidBaseURL := "://invalid-url"
		client.baseUrl = invalidBaseURL

		_, err := client.requestURL("/path", nil, nil)
		if err == nil {
			t.Fatalf("Expected error due to invalid baseUrl, got nil")
		}
	})

	t.Run("URL_encoding_in_query_parameters", func(t *testing.T) {
		// Reset baseUrl to valid value before this subtest
		client.baseUrl = originalBaseURL

		path := "/v1/search"
		values := url.Values{
			"query": {"special chars &/?"},
		}

		expectedURL := "https://api.example.com/v1/search?api_key=testkey&query=special+chars+%26%2F%3F"

		got, err := client.requestURL(path, values, nil)
		if err != nil {
			t.Fatalf("requestURL returned error: %v", err)
		}

		if got != expectedURL {
			t.Errorf("requestURL() = %q, want %q", got, expectedURL)
		}
	})
}

func TestNewRequest(t *testing.T) {
	mockLogger := DefaultLogger
	client := &Client{
		baseUrl:   "https://api.example.com", // Set baseUrl
		apiKey:    "testkey",
		authToken: "Bearer testtoken",
		logger:    mockLogger,
	}

	ctx := context.Background()

	t.Run("GET request without body", func(t *testing.T) {
		method := http.MethodGet
		path := "/v1/resources"
		params := url.Values{
			"param": {"value"},
		}
		var data interface{}
		pathParams := map[string]string{}

		req, err := newRequest(client, ctx, method, path, params, data, pathParams)
		if err != nil {
			t.Fatalf("newRequest returned error: %v", err)
		}

		expectedURL := "https://api.example.com/v1/resources?api_key=testkey&param=value"
		if req.URL.String() != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, req.URL.String())
		}

		if req.Method != method {
			t.Errorf("Expected method %s, got %s", method, req.Method)
		}

		if req.Body != nil {
			t.Errorf("Expected no body for GET request, got %v", req.Body)
		}

		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got %s", req.Header.Get("Content-Type"))
		}

		if req.Header.Get("Authorization") != "Bearer testtoken" {
			t.Errorf("Expected Authorization 'Bearer testtoken', got %s", req.Header.Get("Authorization"))
		}
	})

	t.Run("POST request with JSON body", func(t *testing.T) {
		method := http.MethodPost
		path := "/v1/resources"
		params := url.Values{}
		data := map[string]interface{}{
			"field1": "value1",
			"field2": 2,
		}
		pathParams := map[string]string{}

		req, err := newRequest(client, ctx, method, path, params, data, pathParams)
		if err != nil {
			t.Fatalf("newRequest returned error: %v", err)
		}

		expectedURL := "https://api.example.com/v1/resources?api_key=testkey"
		if req.URL.String() != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, req.URL.String())
		}

		if req.Method != method {
			t.Errorf("Expected method %s, got %s", method, req.Method)
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		expectedBody := `{"field1":"value1","field2":2}`
		if string(body) != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, string(body))
		}

		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got %s", req.Header.Get("Content-Type"))
		}
	})

	t.Run("PUT request with io.Reader body", func(t *testing.T) {
		method := http.MethodPut
		path := "/v1/resources/{id}"
		params := url.Values{}
		bodyContent := "raw body data"
		data := strings.NewReader(bodyContent)
		pathParams := map[string]string{
			"id": "123",
		}

		req, err := newRequest(client, ctx, method, path, params, data, pathParams)
		if err != nil {
			t.Fatalf("newRequest returned error: %v", err)
		}

		expectedURL := "https://api.example.com/v1/resources/123?api_key=testkey"
		if req.URL.String() != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, req.URL.String())
		}

		if req.Method != method {
			t.Errorf("Expected method %s, got %s", method, req.Method)
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		if string(body) != bodyContent {
			t.Errorf("Expected body %s, got %s", bodyContent, string(body))
		}
	})

	t.Run("Unsupported data type", func(t *testing.T) {
		client := &Client{
			baseUrl: "https://api.stream-io-api.com",
			apiKey:  "key",
			logger:  DefaultLogger,
		}
		ctx := context.Background()
		unsupportedData := make(chan int)

		req, err := newRequest(client, ctx, http.MethodPost, "/example", nil, unsupportedData, nil)
		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Nil(t, req.Body) // The body should be nil for unsupported types
	})
}

func TestSetHeaders(t *testing.T) {
	client := &Client{
		authToken: "Bearer testtoken",
		logger:    DefaultLogger,
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.example.com", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	client.setHeaders(req)

	expectedHeaders := map[string]string{
		"Content-Type":     "application/json",
		"X-Stream-Client":  versionHeader(),
		"Authorization":    "Bearer testtoken",
		"Stream-Auth-Type": "jwt",
	}

	for key, expected := range expectedHeaders {
		got := req.Header.Get(key)
		if got != expected {
			t.Errorf("Header %s = %s, want %s", key, got, expected)
		}
	}
}
