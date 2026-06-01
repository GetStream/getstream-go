package getstream

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Sentinel errors expose the SDK error categories. All concrete errors
// returned from the SDK are *StreamError; callers branch on category
// with errors.Is(err, sentinel) and extract structured fields with
// errors.As(err, &streamErr).
var (
	// ErrApiResponse fires when the backend returned an HTTP 4xx/5xx
	// (auth, validation, server error, or any other API failure with the
	// APIError envelope). Also satisfied by ErrRateLimited via the
	// Unwrap chain.
	ErrApiResponse = errors.New("stream: api error")

	// ErrRateLimited fires when the backend returned HTTP 429.
	// StreamError.RetryAfter carries the parsed Retry-After header.
	// errors.Is(err, ErrApiResponse) also returns true.
	ErrRateLimited = errors.New("stream: rate limited")

	// ErrTransport fires when a network-layer failure prevented an HTTP
	// response from being received (connection reset, timeout, TLS,
	// DNS). StreamError.ErrorType identifies the subtype; the original
	// error is preserved via errors.Unwrap.
	ErrTransport = errors.New("stream: transport error")

	// ErrTaskFailed fires when WaitForTask observes status=="failed".
	// StreamError.Task carries the task's ErrorResult.
	ErrTaskFailed = errors.New("stream: task failed")
)

// Transport-error subtype values populated on StreamError.ErrorType when the
// sentinel is ErrTransport.
const (
	ErrorTypeConnectionReset = "connection_reset"
	ErrorTypeTimeout         = "timeout"
	ErrorTypeDNSFailure      = "dns_failure"
	ErrorTypeTLSHandshake    = "tls_handshake_failed"
	ErrorTypeUnknown         = "unknown"
)

// TaskErrorDetails carries the failed-task payload exposed on
// StreamError.Task when the sentinel is ErrTaskFailed.
type TaskErrorDetails struct {
	TaskID      string
	ErrorType   string
	Description string
	StackTrace  string
	Version     string
}

// classifyTransportError maps a transport-layer error to an errorType
// enum. The mapping inspects the error chain via errors.Is / errors.As
// so it survives library wrapping.
func classifyTransportError(err error) string {
	if err == nil {
		return ErrorTypeUnknown
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return ErrorTypeTimeout
	}

	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return ErrorTypeTimeout
	}

	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return ErrorTypeDNSFailure
	}

	var opErr *net.OpError
	if errors.As(err, &opErr) {
		if opErr.Op == "dial" {
			if strings.Contains(opErr.Err.Error(), "no such host") {
				return ErrorTypeDNSFailure
			}
		}
	}

	msg := err.Error()
	switch {
	case strings.Contains(msg, "tls:"),
		strings.Contains(msg, "x509:"),
		strings.Contains(msg, "TLS handshake"):
		return ErrorTypeTLSHandshake
	case strings.Contains(msg, "connection reset"),
		strings.Contains(msg, "connection refused"),
		strings.Contains(msg, "broken pipe"),
		strings.Contains(msg, "EOF"):
		return ErrorTypeConnectionReset
	case strings.Contains(msg, "no such host"):
		return ErrorTypeDNSFailure
	}

	// url.Error often hides the actual cause; recurse into it.
	var urlErr *url.Error
	if errors.As(err, &urlErr) && urlErr.Err != nil && urlErr.Err != err {
		if sub := classifyTransportError(urlErr.Err); sub != ErrorTypeUnknown {
			return sub
		}
	}

	return ErrorTypeUnknown
}

// wrapTransportError converts a raw transport-layer error from the HTTP
// client into a *StreamError with the ErrTransport sentinel, populated
// ErrorType, and the original error preserved via stack-bearing wrap.
func wrapTransportError(err error) *StreamError {
	return &StreamError{
		sentinel:  ErrTransport,
		ErrorType: classifyTransportError(err),
		Message:   "stream transport error: " + err.Error(),
		cause:     stackWrap(err, "transport error"),
	}
}

// parseRetryAfter parses an HTTP Retry-After header value per RFC 7231 §7.1.3.
// Accepts either a non-negative integer in seconds or an HTTP-date. Returns 0
// if the header is missing or unparseable.
//
// now is injectable for deterministic tests; pass time.Now in production.
func parseRetryAfter(value string, now time.Time) time.Duration {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0
	}

	if seconds, err := strconv.Atoi(value); err == nil {
		if seconds < 0 {
			return 0
		}
		return time.Duration(seconds) * time.Second
	}

	if t, err := http.ParseTime(value); err == nil {
		d := t.Sub(now)
		if d < 0 {
			return 0
		}
		return d
	}

	return 0
}
