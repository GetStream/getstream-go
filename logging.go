package getstream

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"
)

var redactedQueryParams = map[string]struct{}{"api_key": {}, "api_secret": {}, "token": {}}

var redactedBodyKeys = []string{"api_secret", "token", "password"}

// redactQuery returns the encoded query string with known-secret param values
// replaced by <redacted>. The original values are never mutated.
func redactQuery(q url.Values) string {
	clone := url.Values{}
	for k, vs := range q {
		if _, secret := redactedQueryParams[strings.ToLower(k)]; secret {
			clone[k] = []string{"<redacted>"}
			continue
		}
		clone[k] = vs
	}
	return clone.Encode()
}

// redactJSONBody shallow-redacts known-secret top-level keys of a JSON object
// body. Non-JSON-object bodies pass through unchanged.
func redactJSONBody(b []byte) string {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return string(b)
	}
	changed := false
	for _, k := range redactedBodyKeys {
		if _, ok := m[k]; ok {
			m[k] = json.RawMessage(`"<redacted>"`)
			changed = true
		}
	}
	if !changed {
		return string(b)
	}
	out, err := json.Marshal(m)
	if err != nil {
		return string(b)
	}
	return string(out)
}

// logRequestSent logs the outgoing request. query must be the built request's
// actual query (e.g. r.URL.Query()), not the caller-supplied params — the
// caller's params never carry the api_key that requestURL injects.
func (c *Client) logRequestSent(method, path string, query url.Values, body []byte) {
	if c.logBodies && body != nil {
		c.logger.Debug("http.request.sent http.request.method=%s url.path=%s url.query=%s http.request.body=%s",
			method, path, redactQuery(query), redactJSONBody(body))
		return
	}
	c.logger.Debug("http.request.sent http.request.method=%s url.path=%s url.query=%s", method, path, redactQuery(query))
}

func (c *Client) logResponseReceived(method, path string, statusCode, bodySize int, d time.Duration, body []byte) {
	if c.logBodies {
		c.logger.Debug("http.response.received http.request.method=%s url.path=%s http.response.status_code=%d http.response.body.size=%d duration_ms=%d http.response.body=%s",
			method, path, statusCode, bodySize, d.Milliseconds(), redactJSONBody(body))
		return
	}
	c.logger.Debug("http.response.received http.request.method=%s url.path=%s http.response.status_code=%d http.response.body.size=%d duration_ms=%d",
		method, path, statusCode, bodySize, d.Milliseconds())
}

// safeErrorMessage returns a log-safe message for a transport-layer error.
// *url.Error.Error() embeds the full outgoing URL, including unredacted
// api_key/api_secret/token query values, so a *url.Error anywhere in the
// chain is unwrapped to its underlying cause instead of using Error()
// directly. err may be the raw transport error or a *StreamError wrapping
// it (via stackWrap) — errors.As/Unwrap see through either.
func safeErrorMessage(err error) string {
	var ue *url.Error
	if errors.As(err, &ue) {
		return ue.Err.Error()
	}
	for {
		u := errors.Unwrap(err)
		if u == nil {
			return err.Error()
		}
		err = u
	}
}

// logRequestFailed logs a final (non-retried) transport-layer failure at
// ERROR: retry disabled, attempts exhausted, or the failure was otherwise
// ineligible for retry. err may be the raw transport error or the
// *StreamError wrapping it; classification and message redaction work
// either way (see safeErrorMessage).
func (c *Client) logRequestFailed(method, path string, err error, d time.Duration) {
	c.logger.Error("http.request.failed http.request.method=%s url.path=%s error.type=%s error.message=%q duration_ms=%d",
		method, path, classifyTransportError(err), safeErrorMessage(err), d.Milliseconds())
}

// logRetryScheduled logs a retryable failure at DEBUG before the loop backs
// off and re-attempts. attempt is 1-indexed (the attempt number that just
// failed). error.type — the closed transport-only classifier enum — is
// included only when the failure is a transport error; a retried 429 carries
// no error.type since rate-limiting isn't a transport failure.
func (c *Client) logRetryScheduled(method, path string, err error, attempt int, delay time.Duration) {
	msg := safeErrorMessage(err)
	if !errors.Is(err, ErrTransport) {
		c.logger.Debug("http.request.failed http.request.method=%s url.path=%s error.message=%q retry.attempt=%d backoff_ms=%d",
			method, path, msg, attempt, delay.Milliseconds())
		return
	}
	var se *StreamError
	errors.As(err, &se)
	c.logger.Debug("http.request.failed http.request.method=%s url.path=%s error.type=%s error.message=%q retry.attempt=%d backoff_ms=%d",
		method, path, se.ErrorType, msg, attempt, delay.Milliseconds())
}
