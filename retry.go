package getstream

import (
	"errors"
	"math/rand"
	"net/http"
	"time"
)

const (
	defaultRetryMaxAttempts = 3
	defaultRetryMaxBackoff  = 30 * time.Second
	retryBackoffBase        = time.Second
)

// RetryConfig is the opt-in auto-retry policy. Disabled by default: the
// client performs exactly one attempt and surfaces errors unchanged. When
// enabled, only GET/HEAD requests failing with HTTP 429 or a transport error
// are retried, and never when the backend marked the error unrecoverable.
type RetryConfig struct {
	// Enabled turns retries on. Default false.
	Enabled bool
	// MaxAttempts is the total attempt budget including the initial request.
	// Default 3 (1 initial + 2 retries).
	MaxAttempts int
	// MaxBackoff caps every wait between attempts, including Retry-After
	// hints from the server. Default 30s.
	MaxBackoff time.Duration
}

// WithRetry enables the opt-in auto-retry policy. Zero values for
// MaxAttempts/MaxBackoff fall back to the documented defaults (3 attempts,
// 30s cap).
func WithRetry(cfg RetryConfig) ClientOption {
	return func(c *Client) {
		if cfg.MaxAttempts <= 0 {
			cfg.MaxAttempts = defaultRetryMaxAttempts
		}
		if cfg.MaxBackoff <= 0 {
			cfg.MaxBackoff = defaultRetryMaxBackoff
		}
		c.retry = cfg
	}
}

// shouldRetry reports whether a failed attempt may be retried. attempt is
// 0-indexed and counts completed attempts.
func (c *Client) shouldRetry(err error, method string, attempt int) bool {
	if !c.retry.Enabled || err == nil {
		return false
	}
	if method != http.MethodGet && method != http.MethodHead {
		return false
	}
	if attempt+1 >= c.retry.MaxAttempts {
		return false
	}
	var streamErr *StreamError
	if !errors.As(err, &streamErr) {
		return false
	}
	if streamErr.Unrecoverable {
		return false
	}
	return errors.Is(err, ErrRateLimited) || errors.Is(err, ErrTransport)
}

// retryDelay returns the wait before the next attempt: a positive Retry-After
// hint clamped to MaxBackoff, otherwise exponential backoff with full jitter.
func (c *Client) retryDelay(err error, attempt int) time.Duration {
	var streamErr *StreamError
	if errors.As(err, &streamErr) && streamErr.RetryAfter > 0 {
		if streamErr.RetryAfter > c.retry.MaxBackoff {
			return c.retry.MaxBackoff
		}
		return streamErr.RetryAfter
	}
	ceil := retryBackoffBase << uint(attempt)
	if ceil <= 0 || ceil > c.retry.MaxBackoff {
		ceil = c.retry.MaxBackoff
	}
	if ceil <= 0 {
		return 0
	}
	return time.Duration(rand.Int63n(int64(ceil) + 1))
}
