package getstream

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"testing"
	"time"
)

// scriptedRetryClient replays a fixed sequence of responses/errors, one per
// call. Panics (via index out of range) if called more times than scripted —
// that's a test bug, not a case to handle gracefully.
type scriptedRetryClient struct {
	calls     int
	responses []func() (*http.Response, error)
}

func (s *scriptedRetryClient) Do(_ *http.Request) (*http.Response, error) {
	step := s.responses[s.calls]
	s.calls++
	return step()
}

func canned(status int, body string, headers map[string]string) func() (*http.Response, error) {
	return func() (*http.Response, error) {
		h := http.Header{}
		h.Set("Content-Type", "application/json")
		for k, v := range headers {
			h.Set(k, v)
		}
		return &http.Response{StatusCode: status, Header: h, Body: io.NopCloser(strings.NewReader(body))}, nil
	}
}

func newRetryTestClient(t *testing.T, script *scriptedRetryClient, cfg *RetryConfig) *Client {
	t.Helper()
	opts := []ClientOption{WithHTTPClient(script)}
	if cfg != nil {
		opts = append(opts, WithRetry(*cfg))
	}
	c, err := newClient("key", "secret", opts...)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func doGET(c *Client) error {
	var out map[string]any
	_, err := MakeRequest[map[string]any, map[string]any](c, context.Background(), http.MethodGet, "/api/v2/x", url.Values{}, nil, &out, nil)
	return err
}

func TestRetryDisabledByDefault(t *testing.T) {
	script := &scriptedRetryClient{responses: []func() (*http.Response, error){
		canned(429, `{}`, map[string]string{"Retry-After": "1"}),
	}}
	err := doGET(newRetryTestClient(t, script, nil))
	if !errors.Is(err, ErrRateLimited) {
		t.Fatalf("want ErrRateLimited, got %v", err)
	}
	if script.calls != 1 {
		t.Fatalf("want 1 call, got %d", script.calls)
	}
}

func TestRetryEnabledGet429ThenSuccess(t *testing.T) {
	script := &scriptedRetryClient{responses: []func() (*http.Response, error){
		canned(429, `{}`, nil),
		canned(200, `{}`, nil),
	}}
	c := newRetryTestClient(t, script, &RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: time.Millisecond})
	if err := doGET(c); err != nil {
		t.Fatalf("want success, got %v", err)
	}
	if script.calls != 2 {
		t.Fatalf("want 2 calls, got %d", script.calls)
	}
}

func TestRetryNeverRetriesPost(t *testing.T) {
	script := &scriptedRetryClient{responses: []func() (*http.Response, error){
		canned(429, `{}`, nil),
	}}
	c := newRetryTestClient(t, script, &RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: time.Millisecond})
	var out map[string]any
	_, err := MakeRequest[map[string]any, map[string]any](c, context.Background(), http.MethodPost, "/api/v2/x", url.Values{}, nil, &out, nil)
	if !errors.Is(err, ErrRateLimited) || script.calls != 1 {
		t.Fatalf("want single rate-limited call, got calls=%d err=%v", script.calls, err)
	}
}

func TestRetryHonorsUnrecoverable(t *testing.T) {
	script := &scriptedRetryClient{responses: []func() (*http.Response, error){
		canned(429, `{"code":9,"message":"nope","unrecoverable":true}`, nil),
	}}
	c := newRetryTestClient(t, script, &RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: time.Millisecond})
	err := doGET(c)
	if !errors.Is(err, ErrRateLimited) || script.calls != 1 {
		t.Fatalf("want single unrecoverable call, got calls=%d err=%v", script.calls, err)
	}
}

func TestRetryTransportErrorThenSuccess(t *testing.T) {
	script := &scriptedRetryClient{responses: []func() (*http.Response, error){
		func() (*http.Response, error) { return nil, syscall.ECONNRESET },
		canned(200, `{}`, nil),
	}}
	c := newRetryTestClient(t, script, &RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: time.Millisecond})
	if err := doGET(c); err != nil {
		t.Fatalf("want success, got %v", err)
	}
	if script.calls != 2 {
		t.Fatalf("want 2 calls, got %d", script.calls)
	}
}

func TestRetryExhaustionSurfacesLastError(t *testing.T) {
	script := &scriptedRetryClient{responses: []func() (*http.Response, error){
		canned(429, `{}`, nil), canned(429, `{}`, nil), canned(429, `{}`, nil),
	}}
	c := newRetryTestClient(t, script, &RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: time.Millisecond})
	err := doGET(c)
	if !errors.Is(err, ErrRateLimited) || script.calls != 3 {
		t.Fatalf("want 3 calls ending rate-limited, got calls=%d err=%v", script.calls, err)
	}
}

func TestRetryDelayClampAndJitterBounds(t *testing.T) {
	c := &Client{retry: RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: 30 * time.Second}}
	rateLimited := &StreamError{sentinel: ErrRateLimited, RetryAfter: 600 * time.Second}
	if d := c.retryDelay(rateLimited, 0); d != 30*time.Second {
		t.Fatalf("want clamp to 30s, got %s", d)
	}
	transport := &StreamError{sentinel: ErrTransport}
	for attempt := 0; attempt < 3; attempt++ {
		ceil := time.Second << uint(attempt)
		if ceil > c.retry.MaxBackoff {
			ceil = c.retry.MaxBackoff
		}
		for i := 0; i < 50; i++ {
			if d := c.retryDelay(transport, attempt); d < 0 || d > ceil {
				t.Fatalf("attempt %d: delay %s out of [0,%s]", attempt, d, ceil)
			}
		}
	}
}

// Cross-SDK schema rule: error.type is a closed transport-only enum
// (connection_reset/timeout/dns_failure/tls_handshake_failed/unknown), so a
// retried 429 must not carry it, while a retried transport error must.
func TestRetryLogSchemaErrorTypePresenceMatchesFailureKind(t *testing.T) {
	rec := &recordingLogger{}
	script := &scriptedRetryClient{responses: []func() (*http.Response, error){
		canned(429, `{}`, nil),
		canned(200, `{}`, nil),
	}}
	c, err := newClient("key", "secret",
		WithHTTPClient(script), WithLogger(rec),
		WithRetry(RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: time.Millisecond}))
	if err != nil {
		t.Fatal(err)
	}
	if err := doGET(c); err != nil {
		t.Fatalf("want success, got %v", err)
	}
	var line429 string
	for _, e := range rec.debug {
		if strings.Contains(e, "http.request.failed") {
			line429 = e
		}
	}
	if line429 == "" {
		t.Fatalf("want a retry http.request.failed DEBUG log: %v", rec.debug)
	}
	if !strings.Contains(line429, "retry.attempt=1") {
		t.Fatalf("want retry.attempt=1, got %q", line429)
	}
	if strings.Contains(line429, "error.type=") {
		t.Fatalf("429 retry log must not carry error.type: %q", line429)
	}

	rec2 := &recordingLogger{}
	script2 := &scriptedRetryClient{responses: []func() (*http.Response, error){
		func() (*http.Response, error) { return nil, syscall.ECONNRESET },
		canned(200, `{}`, nil),
	}}
	c2, err := newClient("key", "secret",
		WithHTTPClient(script2), WithLogger(rec2),
		WithRetry(RetryConfig{Enabled: true, MaxAttempts: 3, MaxBackoff: time.Millisecond}))
	if err != nil {
		t.Fatal(err)
	}
	if err := doGET(c2); err != nil {
		t.Fatalf("want success, got %v", err)
	}
	var lineTransport string
	for _, e := range rec2.debug {
		if strings.Contains(e, "http.request.failed") {
			lineTransport = e
		}
	}
	if lineTransport == "" {
		t.Fatalf("want a retry http.request.failed DEBUG log: %v", rec2.debug)
	}
	if !strings.Contains(lineTransport, "error.type=connection_reset") {
		t.Fatalf("transport retry log must carry error.type=connection_reset: %q", lineTransport)
	}
}

// Retry disabled must reproduce today's logging exactly: one ERROR
// http.request.failed on transport failure, none on a non-retried 429.
func TestRetryDisabledLoggingMatchesPreRetryBehavior(t *testing.T) {
	rec := &recordingLogger{}
	script := &scriptedRetryClient{responses: []func() (*http.Response, error){
		func() (*http.Response, error) { return nil, syscall.ECONNRESET },
	}}
	c, err := newClient("key", "secret", WithHTTPClient(script), WithLogger(rec))
	if err != nil {
		t.Fatal(err)
	}
	if err := doGET(c); !errors.Is(err, ErrTransport) {
		t.Fatalf("want ErrTransport, got %v", err)
	}
	count := 0
	for _, e := range rec.errs {
		if strings.Contains(e, "http.request.failed") {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("want exactly 1 ERROR http.request.failed, got %d: %v", count, rec.errs)
	}

	rec2 := &recordingLogger{}
	script2 := &scriptedRetryClient{responses: []func() (*http.Response, error){
		canned(429, `{}`, nil),
	}}
	c2, err := newClient("key", "secret", WithHTTPClient(script2), WithLogger(rec2))
	if err != nil {
		t.Fatal(err)
	}
	if err := doGET(c2); !errors.Is(err, ErrRateLimited) {
		t.Fatalf("want ErrRateLimited, got %v", err)
	}
	for _, e := range rec2.errs {
		if strings.Contains(e, "http.request.failed") {
			t.Fatalf("non-retried 429 must not emit http.request.failed: %v", rec2.errs)
		}
	}
}
