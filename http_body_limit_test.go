package getstream

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// bodyLimitServer serves body on every request and counts the TCP
// connections it accepted. chunked drops Content-Length so the client can
// only find out the size by reading.
func bodyLimitServer(t *testing.T, body string, chunked bool) (*httptest.Server, *atomic.Int32, *atomic.Int32) {
	t.Helper()
	var conns, requests atomic.Int32
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if chunked {
			w.(http.Flusher).Flush()
		}
		_, _ = w.Write([]byte(body))
	}))
	srv.Config.ConnState = func(_ net.Conn, state http.ConnState) {
		if state == http.StateNew {
			conns.Add(1)
		}
	}
	srv.Start()
	t.Cleanup(srv.Close)
	return srv, &conns, &requests
}

func bodyLimitGET(t *testing.T, c *Client) error {
	t.Helper()
	var out map[string]any
	_, err := MakeRequest[map[string]any, map[string]any](c, context.Background(), http.MethodGet, "/api/v2/app", nil, nil, &out, nil)
	return err
}

func newBodyLimitClient(t *testing.T, baseURL string, opts ...ClientOption) *Client {
	t.Helper()
	c, err := newClient("key", "secret", append([]ClientOption{WithBaseUrl(baseURL)}, opts...)...)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

// jsonOfSize returns a valid JSON object of exactly n bytes.
func jsonOfSize(n int) string {
	return `{"a":"` + strings.Repeat("x", n-8) + `"}`
}

func TestResponseBodyLimitDefault(t *testing.T) {
	c := newBodyLimitClient(t, "http://unused")
	if got := c.MaxResponseBodyBytes(); got != DefaultMaxResponseBodyBytes {
		t.Fatalf("default limit = %d, want %d", got, DefaultMaxResponseBodyBytes)
	}
	c = newBodyLimitClient(t, "http://unused", WithMaxResponseBodyBytes(0))
	if got := c.MaxResponseBodyBytes(); got != DefaultMaxResponseBodyBytes {
		t.Fatalf("WithMaxResponseBodyBytes(0) limit = %d, want default %d", got, DefaultMaxResponseBodyBytes)
	}
	if got := (&Client{}).MaxResponseBodyBytes(); got != DefaultMaxResponseBodyBytes {
		t.Fatalf("zero Client limit = %d, want default %d", got, DefaultMaxResponseBodyBytes)
	}
}

func TestResponseBodyAtLimitIsRead(t *testing.T) {
	for _, chunked := range []bool{false, true} {
		srv, _, _ := bodyLimitServer(t, jsonOfSize(64), chunked)
		c := newBodyLimitClient(t, srv.URL, WithMaxResponseBodyBytes(64))
		if err := bodyLimitGET(t, c); err != nil {
			t.Fatalf("chunked=%v: body at the limit must succeed, got %v", chunked, err)
		}
	}
}

func TestResponseBodyOverLimitFails(t *testing.T) {
	for _, chunked := range []bool{false, true} {
		srv, _, _ := bodyLimitServer(t, jsonOfSize(65), chunked)
		c := newBodyLimitClient(t, srv.URL, WithMaxResponseBodyBytes(64))
		err := bodyLimitGET(t, c)
		if !errors.Is(err, ErrResponseTooLarge) {
			t.Fatalf("chunked=%v: want ErrResponseTooLarge, got %v", chunked, err)
		}
		var streamErr *StreamError
		if !errors.As(err, &streamErr) || streamErr.StatusCode != http.StatusOK {
			t.Fatalf("chunked=%v: want *StreamError with StatusCode 200, got %#v", chunked, err)
		}
		if errors.Is(err, ErrTransport) || errors.Is(err, ErrApiResponse) {
			t.Fatalf("chunked=%v: oversized body must not look like a transport/API error: %v", chunked, err)
		}
		if !strings.Contains(err.Error(), "64-byte limit") {
			t.Fatalf("chunked=%v: error should name the limit: %v", chunked, err)
		}
	}
}

func TestResponseBodyOverLimitNotRetried(t *testing.T) {
	srv, _, requests := bodyLimitServer(t, jsonOfSize(65), false)
	c := newBodyLimitClient(t, srv.URL, WithMaxResponseBodyBytes(64),
		WithRetry(RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: time.Millisecond}))
	if err := bodyLimitGET(t, c); !errors.Is(err, ErrResponseTooLarge) {
		t.Fatalf("want ErrResponseTooLarge, got %v", err)
	}
	if got := requests.Load(); got != 1 {
		t.Fatalf("oversized response must not be retried, server saw %d requests", got)
	}
}

// A body within the limit is read to EOF, so keep-alive reuses the connection.
func TestResponseBodyWithinLimitReusesConnection(t *testing.T) {
	srv, conns, _ := bodyLimitServer(t, `{"ok":true}`, true)
	c := newBodyLimitClient(t, srv.URL)
	for i := 0; i < 3; i++ {
		if err := bodyLimitGET(t, c); err != nil {
			t.Fatal(err)
		}
	}
	if got := conns.Load(); got != 1 {
		t.Fatalf("want 1 reused connection, got %d", got)
	}
}
