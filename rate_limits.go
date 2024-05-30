package getstream

import (
	"net/http"
	"strconv"
)

const (
	HeaderRateLimit     = "X-Ratelimit-Limit"
	HeaderRateRemaining = "X-Ratelimit-Remaining"
	HeaderRateReset     = "X-Ratelimit-Reset"
)

// RateLimitInfo represents the quota and usage for a single endpoint.
type RateLimitInfo struct {
	// Limit is the maximum number of API calls for a single time window (1 minute).
	Limit int64 `json:"limit"`
	// Remaining is the number of API calls remaining in the current time window (1 minute).
	Remaining int64 `json:"remaining"`
	// Reset is the Unix timestamp of the expiration of the current rate limit time window.
	Reset int64 `json:"reset"`
}

func NewRateLimitFromHeaders(headers http.Header) *RateLimitInfo {
	var rl RateLimitInfo

	limit, err := strconv.ParseInt(headers.Get(HeaderRateLimit), 10, 64)
	if err == nil {
		rl.Limit = limit
	}
	remaining, err := strconv.ParseInt(headers.Get(HeaderRateRemaining), 10, 64)
	if err == nil {
		rl.Remaining = remaining
	}
	reset, err := strconv.ParseInt(headers.Get(HeaderRateReset), 10, 64)
	if err == nil && reset > 0 {
		rl.Reset = reset
	}

	return &rl
}
