package getstream

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"testing"
)

type recordingLogger struct {
	debug, info, warn, errs []string
}

func (l *recordingLogger) Debug(f string, v ...interface{}) {
	l.debug = append(l.debug, fmt.Sprintf(f, v...))
}
func (l *recordingLogger) Info(f string, v ...interface{}) {
	l.info = append(l.info, fmt.Sprintf(f, v...))
}
func (l *recordingLogger) Warn(f string, v ...interface{}) {
	l.warn = append(l.warn, fmt.Sprintf(f, v...))
}
func (l *recordingLogger) Error(f string, v ...interface{}) {
	l.errs = append(l.errs, fmt.Sprintf(f, v...))
}

type oneShotClient struct {
	status int
	body   string
	err    error
	calls  int
}

func (s *oneShotClient) Do(r *http.Request) (*http.Response, error) {
	s.calls++
	if s.err != nil {
		return nil, s.err
	}
	h := http.Header{}
	h.Set("Content-Type", "application/json")
	return &http.Response{StatusCode: s.status, Header: h, Body: io.NopCloser(strings.NewReader(s.body))}, nil
}

func loggedGET(t *testing.T, fake HttpClient, opts ...ClientOption) (*recordingLogger, error) {
	t.Helper()
	rec := &recordingLogger{}
	c, err := NewClient("key", "secret", append([]ClientOption{WithHTTPClient(fake), WithLogger(rec)}, opts...)...)
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]any
	_, err = MakeRequest[map[string]any, map[string]any](c.Client, context.Background(), http.MethodGet, "/api/v2/app", url.Values{"api_key": {"key"}}, nil, &out, nil)
	return rec, err
}

func has(entries []string, substr string) bool {
	for _, e := range entries {
		if strings.Contains(e, substr) {
			return true
		}
	}
	return false
}

func TestClientInitializedEmittedOnce(t *testing.T) {
	rec, _ := loggedGET(t, &oneShotClient{status: 200, body: `{}`})
	count := 0
	for _, e := range rec.info {
		if strings.Contains(e, "client.initialized") {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("want exactly 1 client.initialized, got %d in %v", count, rec.info)
	}
	if !has(rec.info, "stream.sdk.name=getstream-go") || !has(rec.info, "max_conns_per_host") {
		t.Fatalf("client.initialized missing schema fields: %v", rec.info)
	}
}

func TestRequestAndResponseEventsOnSuccess(t *testing.T) {
	rec, err := loggedGET(t, &oneShotClient{status: 200, body: `{}`})
	if err != nil {
		t.Fatal(err)
	}
	if !has(rec.debug, "http.request.sent") || !has(rec.debug, "http.response.received") {
		t.Fatalf("missing request/response events: %v", rec.debug)
	}
	if !has(rec.debug, "http.response.status_code=200") {
		t.Fatalf("missing status code: %v", rec.debug)
	}
}

func TestErrorStatusFlowsThroughResponseReceived(t *testing.T) {
	rec, err := loggedGET(t, &oneShotClient{status: 500, body: `{"code":1,"message":"boom"}`})
	if err == nil {
		t.Fatal("want error")
	}
	if !has(rec.debug, "http.response.status_code=500") {
		t.Fatalf("want 500 via http.response.received: %v", rec.debug)
	}
	if has(rec.errs, "http.request.failed") {
		t.Fatalf("4xx/5xx must not emit http.request.failed: %v", rec.errs)
	}
}

func TestTransportFailureEmitsRequestFailed(t *testing.T) {
	rec, err := loggedGET(t, &oneShotClient{err: syscall.ECONNRESET})
	if err == nil {
		t.Fatal("want error")
	}
	if !has(rec.errs, "http.request.failed") || !has(rec.errs, "error.type=connection_reset") {
		t.Fatalf("want http.request.failed with error.type: %v", rec.errs)
	}
}

func TestQueryRedaction(t *testing.T) {
	rec, _ := loggedGET(t, &oneShotClient{status: 200, body: `{}`})
	for _, e := range rec.debug {
		if strings.Contains(e, "api_key=key") {
			t.Fatalf("api_key leaked: %q", e)
		}
	}
	if !has(rec.debug, "api_key=%3Credacted%3E") && !has(rec.debug, "api_key=<redacted>") {
		t.Fatalf("want redacted api_key in url.query: %v", rec.debug)
	}
}

func TestNoHeadersInDebugOutput(t *testing.T) {
	rec, _ := loggedGET(t, &oneShotClient{status: 200, body: `{}`})
	for _, e := range append(rec.debug, rec.info...) {
		if strings.Contains(e, "Authorization") {
			t.Fatalf("Authorization header leaked: %q", e)
		}
	}
}

func TestLogBodiesOptInAndWarn(t *testing.T) {
	recOff, _ := loggedGET(t, &oneShotClient{status: 200, body: `{"secret":"x"}`})
	if has(recOff.debug, "http.response.body=") {
		t.Fatalf("bodies logged without opt-in: %v", recOff.debug)
	}
	recOn, _ := loggedGET(t, &oneShotClient{status: 200, body: `{"token":"tok","ok":true}`}, WithLogBodies(true))
	if len(recOn.warn) != 1 || !strings.Contains(recOn.warn[0], "bodies will be logged") {
		t.Fatalf("want exactly one body-logging WARN, got %v", recOn.warn)
	}
	if !has(recOn.debug, "http.response.body=") {
		t.Fatalf("want body field with opt-in: %v", recOn.debug)
	}
	if has(recOn.debug, `"token":"tok"`) {
		t.Fatalf("token body key leaked: %v", recOn.debug)
	}
}

func TestRedactJSONBody(t *testing.T) {
	out := redactJSONBody([]byte(`{"token":"t","password":"p","api_secret":"s","keep":"v"}`))
	for _, leak := range []string{`"t"`, `"p"`, `"s"`} {
		if strings.Contains(out, leak) {
			t.Fatalf("leak %s in %s", leak, out)
		}
	}
	if !strings.Contains(out, `"keep":"v"`) {
		t.Fatalf("non-secret key dropped: %s", out)
	}
	if got := redactJSONBody([]byte(`not json`)); got != "not json" {
		t.Fatalf("non-JSON body must pass through, got %q", got)
	}
}
