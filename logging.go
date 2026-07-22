package getstream

import (
	"encoding/json"
	"net/url"
	"time"
)

var redactedQueryParams = map[string]struct{}{"api_key": {}, "api_secret": {}, "token": {}}

var redactedBodyKeys = []string{"api_secret", "token", "password"}

// redactQuery returns the encoded query string with known-secret param values
// replaced by <redacted>. The original values are never mutated.
func redactQuery(q url.Values) string {
	clone := url.Values{}
	for k, vs := range q {
		if _, secret := redactedQueryParams[toLowerASCII(k)]; secret {
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

func toLowerASCII(s string) string {
	b := []byte(s)
	for i := range b {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 'a' - 'A'
		}
	}
	return string(b)
}

func (c *Client) logRequestSent(method, path string, params url.Values, body []byte) {
	if c.logBodies && body != nil {
		c.logger.Debug("http.request.sent http.request.method=%s url.path=%s url.query=%s http.request.body=%s",
			method, path, redactQuery(params), redactJSONBody(body))
		return
	}
	c.logger.Debug("http.request.sent http.request.method=%s url.path=%s url.query=%s", method, path, redactQuery(params))
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

func (c *Client) logRequestFailed(method, path string, err error, d time.Duration) {
	c.logger.Error("http.request.failed http.request.method=%s url.path=%s error.type=%s error.message=%q duration_ms=%d",
		method, path, classifyTransportError(err), err.Error(), d.Milliseconds())
}
