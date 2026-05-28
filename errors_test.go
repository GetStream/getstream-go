package getstream

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestStreamError_Is_Sentinels(t *testing.T) {
	tests := []struct {
		name     string
		sentinel error
		targets  []error
		want     map[error]bool
	}{
		{
			name:     "ErrApiResponse only",
			sentinel: ErrApiResponse,
			targets:  []error{ErrApiResponse, ErrRateLimited, ErrTransport, ErrTaskFailed},
			want: map[error]bool{
				ErrApiResponse: true,
				ErrRateLimited: false,
				ErrTransport:   false,
				ErrTaskFailed:  false,
			},
		},
		{
			name:     "ErrRateLimited also satisfies ErrApiResponse",
			sentinel: ErrRateLimited,
			targets:  []error{ErrApiResponse, ErrRateLimited, ErrTransport, ErrTaskFailed},
			want: map[error]bool{
				ErrApiResponse: true,
				ErrRateLimited: true,
				ErrTransport:   false,
				ErrTaskFailed:  false,
			},
		},
		{
			name:     "ErrTransport only",
			sentinel: ErrTransport,
			targets:  []error{ErrApiResponse, ErrRateLimited, ErrTransport, ErrTaskFailed},
			want: map[error]bool{
				ErrApiResponse: false,
				ErrRateLimited: false,
				ErrTransport:   true,
				ErrTaskFailed:  false,
			},
		},
		{
			name:     "ErrTaskFailed only",
			sentinel: ErrTaskFailed,
			targets:  []error{ErrApiResponse, ErrRateLimited, ErrTransport, ErrTaskFailed},
			want: map[error]bool{
				ErrApiResponse: false,
				ErrRateLimited: false,
				ErrTransport:   false,
				ErrTaskFailed:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &StreamError{sentinel: tt.sentinel, Message: "x"}
			var e error = err
			for _, target := range tt.targets {
				got := errors.Is(e, target)
				if got != tt.want[target] {
					t.Errorf("errors.Is(err, %v) = %v, want %v", target, got, tt.want[target])
				}
			}
		})
	}
}

func TestStreamError_As_ExtractsStruct(t *testing.T) {
	original := &StreamError{
		sentinel:        ErrApiResponse,
		StatusCode:      404,
		Code:            17,
		Message:         "not found",
		ExceptionFields: map[string]string{},
	}
	var err error = original

	var got *StreamError
	if !errors.As(err, &got) {
		t.Fatalf("errors.As(*StreamError) returned false")
	}
	if got.StatusCode != 404 || got.Code != 17 || got.Message != "not found" {
		t.Errorf("extracted struct mismatch: %+v", got)
	}
}

func TestStreamError_UnwrapReturnsCause(t *testing.T) {
	root := errors.New("network closed")
	err := &StreamError{
		sentinel: ErrTransport,
		cause:    stackWrap(root, "transport error"),
		Message:  "transport",
	}
	if !errors.Is(err, root) {
		t.Fatalf("errors.Is(err, root) returned false; cause chain broken")
	}
}

func TestParseRetryAfter_Integer(t *testing.T) {
	got := parseRetryAfter("30", time.Now())
	want := 30 * time.Second
	if got != want {
		t.Fatalf("parseRetryAfter(\"30\") = %v, want %v", got, want)
	}
}

func TestParseRetryAfter_HTTPDate(t *testing.T) {
	now := time.Date(2026, 5, 28, 12, 0, 0, 0, time.UTC)
	header := now.Add(45 * time.Second).UTC().Format(http.TimeFormat)
	got := parseRetryAfter(header, now)
	if got < 44*time.Second || got > 46*time.Second {
		t.Fatalf("parseRetryAfter(http-date) = %v, want ~45s", got)
	}
}

func TestParseRetryAfter_NegativeIntegerClamps(t *testing.T) {
	if got := parseRetryAfter("-5", time.Now()); got != 0 {
		t.Fatalf("parseRetryAfter(\"-5\") = %v, want 0", got)
	}
}

func TestParseRetryAfter_PastHTTPDateClamps(t *testing.T) {
	now := time.Date(2026, 5, 28, 12, 0, 0, 0, time.UTC)
	header := now.Add(-time.Hour).UTC().Format(http.TimeFormat)
	if got := parseRetryAfter(header, now); got != 0 {
		t.Fatalf("parseRetryAfter(past-date) = %v, want 0", got)
	}
}

func TestParseRetryAfter_EmptyAndJunk(t *testing.T) {
	now := time.Now()
	cases := []string{"", "  ", "soon", "1d"}
	for _, c := range cases {
		if got := parseRetryAfter(c, now); got != 0 {
			t.Errorf("parseRetryAfter(%q) = %v, want 0", c, got)
		}
	}
}

func TestClassifyTransportError(t *testing.T) {
	dnsErr := &net.DNSError{Err: "no such host", Name: "example.invalid"}
	timeoutErr := &net.OpError{Op: "read", Err: timeoutError{}}
	resetErr := errors.New("read tcp 127.0.0.1:1234: connection reset by peer")
	// `tls.CertificateVerificationError` is only available from Go 1.20+.
	// Use a generic error whose message the classifier matches via substring,
	// so the test works on Go 1.19 and newer.
	tlsErr := errors.New("tls: handshake failure: certificate verify failed")
	refusedErr := errors.New("dial tcp 127.0.0.1:1: connection refused")

	tests := []struct {
		name string
		err  error
		want string
	}{
		{"dns", dnsErr, ErrorTypeDNSFailure},
		{"timeout", timeoutErr, ErrorTypeTimeout},
		{"context deadline exceeded", context.DeadlineExceeded, ErrorTypeTimeout},
		{"connection reset", resetErr, ErrorTypeConnectionReset},
		{"connection refused", refusedErr, ErrorTypeConnectionReset},
		{"tls", tlsErr, ErrorTypeTLSHandshake},
		{"unknown", errors.New("something weird"), ErrorTypeUnknown},
		{"nil", nil, ErrorTypeUnknown},
		{
			name: "url.Error wrapping timeout",
			err:  &url.Error{Op: "Get", URL: "https://x", Err: timeoutError{}},
			want: ErrorTypeTimeout,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := classifyTransportError(tt.err); got != tt.want {
				t.Errorf("classifyTransportError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapTransportError_FieldsAndSentinel(t *testing.T) {
	root := errors.New("read tcp 1.2.3.4:5: connection reset by peer")
	err := wrapTransportError(root)

	if !errors.Is(err, ErrTransport) {
		t.Errorf("errors.Is(err, ErrTransport) = false")
	}
	if err.ErrorType != ErrorTypeConnectionReset {
		t.Errorf("ErrorType = %v, want %v", err.ErrorType, ErrorTypeConnectionReset)
	}
	if !errors.Is(err, root) {
		t.Errorf("errors.Is(err, root) = false; cause chain broken")
	}
	if !strings.Contains(err.Error(), "connection reset by peer") {
		t.Errorf("err.Error() = %q, want substring 'connection reset by peer'", err.Error())
	}
}

func TestBuildAPIError_429PopulatesRetryAfter(t *testing.T) {
	body := `{"code":9,"message":"rate limited","StatusCode":429}`
	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header: http.Header{
			"Retry-After": []string{"42"},
		},
	}
	got := buildAPIError(resp, []byte(body))

	if !errors.Is(got, ErrRateLimited) {
		t.Errorf("errors.Is(err, ErrRateLimited) = false")
	}
	if !errors.Is(got, ErrApiResponse) {
		t.Errorf("ErrRateLimited should also match ErrApiResponse via §4.2")
	}
	if got.RetryAfter != 42*time.Second {
		t.Errorf("RetryAfter = %v, want 42s", got.RetryAfter)
	}
	if got.StatusCode != 429 {
		t.Errorf("StatusCode = %d, want 429", got.StatusCode)
	}
	if got.Code != 9 {
		t.Errorf("Code = %d, want 9", got.Code)
	}
	if got.RawResponseBody != body {
		t.Errorf("RawResponseBody mismatch: got %q", got.RawResponseBody)
	}
}

func TestBuildAPIError_UnparseableBodyPath(t *testing.T) {
	body := `not json`
	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     http.Header{},
	}
	got := buildAPIError(resp, []byte(body))

	if !errors.Is(got, ErrApiResponse) {
		t.Errorf("unparseable body should map to ErrApiResponse, not transport")
	}
	if got.Code != 0 {
		t.Errorf("Code = %d, want 0 for unparseable", got.Code)
	}
	if got.Message != "failed to parse error response" {
		t.Errorf("Message = %q, want canonical unparseable text", got.Message)
	}
	if got.RawResponseBody != body {
		t.Errorf("RawResponseBody = %q, want raw body preserved", got.RawResponseBody)
	}
	// Cause chain must carry the JSON-parse error.
	var parseErr *json.SyntaxError
	if !errors.As(got, &parseErr) {
		t.Errorf("cause chain should carry the JSON parse error")
	}
}

func TestBuildAPIError_PopulatesUnrecoverableAndDetails(t *testing.T) {
	body := `{"code":4,"message":"bad request","StatusCode":400,"unrecoverable":true,"details":{"field":"value"},"more_info":"https://docs.example/4"}`
	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Header:     http.Header{},
	}
	got := buildAPIError(resp, []byte(body))

	if !got.Unrecoverable {
		t.Errorf("Unrecoverable = false, want true")
	}
	if got.MoreInfo != "https://docs.example/4" {
		t.Errorf("MoreInfo = %q, want canonical", got.MoreInfo)
	}
	if len(got.Details) == 0 {
		t.Errorf("Details empty, want preserved raw JSON")
	}
	var decoded map[string]string
	if err := json.Unmarshal(got.Details, &decoded); err != nil {
		t.Errorf("Details not valid JSON: %v", err)
	}
	if decoded["field"] != "value" {
		t.Errorf("Details decoded mismatch: %v", decoded)
	}
}

func TestBuildAPIError_EmptyExceptionFieldsAlwaysMap(t *testing.T) {
	body := `{"code":1,"message":"x","StatusCode":500}`
	resp := &http.Response{StatusCode: 500, Header: http.Header{}}
	got := buildAPIError(resp, []byte(body))
	if got.ExceptionFields == nil {
		t.Fatalf("ExceptionFields nil; spec §5.1 mandates empty map when absent")
	}
}

// timeoutError implements net.Error reporting Timeout=true.
type timeoutError struct{}

func (timeoutError) Error() string   { return "i/o timeout" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return true }
