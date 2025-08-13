package getstream_test

import (
	"net/http"
	"testing"

	"github.com/GetStream/getstream-go/v3"
	"github.com/stretchr/testify/assert"
)

func TestNewRateLimitFromHeaders(t *testing.T) {
	headers := http.Header{}
	headers.Set(getstream.HeaderRateLimit, "100")
	headers.Set(getstream.HeaderRateRemaining, "50")
	headers.Set(getstream.HeaderRateReset, "1609459200")

	rateLimitInfo := getstream.NewRateLimitFromHeaders(headers)

	assert.Equal(t, int64(100), rateLimitInfo.Limit)
	assert.Equal(t, int64(50), rateLimitInfo.Remaining)
	assert.Equal(t, int64(1609459200), rateLimitInfo.Reset)
}
