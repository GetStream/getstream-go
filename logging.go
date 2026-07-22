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

// logRequestFailed logs a transport-layer failure. Classification runs on the
// original err (needed to see through *url.Error to the real cause), but the
// logged message unwraps *url.Error to its underlying cause: url.Error.Error()
// embeds the full request URL, including unredacted api_key/api_secret/token
// query values, so logging it verbatim would leak secrets on every real
// *http.Client transport failure.
func (c *Client) logRequestFailed(method, path string, err error, d time.Duration) {
	msg := err.Error()
	var ue *url.Error
	if errors.As(err, &ue) {
		msg = ue.Err.Error()
	}
	c.logger.Error("http.request.failed http.request.method=%s url.path=%s error.type=%s error.message=%q duration_ms=%d",
		method, path, classifyTransportError(err), msg, d.Milliseconds())
}
