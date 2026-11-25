package getstream

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// DefaultSRTTokenExpiration is how long the token in an SRT address works unless
// WithSRTTokenExpiration sets another duration.
const DefaultSRTTokenExpiration = 24 * time.Hour

// maxSRTStreamIDBytes is the SRT limit on the stream id; encoders refuse a longer one.
const maxSRTStreamIDBytes = 512

var srtStreamIDPattern = regexp.MustCompile(`[?&]streamid=([^&]*)`)

// SRTCredentials holds the SRT address an encoder (OBS, ffmpeg, ...) publishes to.
//
// Address contains a user token, so String and GoString leave it out: printing the
// credentials with fmt does not leak it. Read the field directly when you need it.
type SRTCredentials struct {
	Address string
	// ExpiresAt is when the token in Address expires.
	ExpiresAt time.Time
}

// String returns the credentials with the address redacted.
func (c SRTCredentials) String() string {
	return fmt.Sprintf("SRTCredentials{Address: [redacted], ExpiresAt: %s}", c.ExpiresAt.Format(time.RFC3339))
}

// GoString returns the credentials with the address redacted, for the %#v verb.
func (c SRTCredentials) GoString() string {
	return c.String()
}

type srtCredentialsOptions struct {
	expiration time.Duration
}

// SRTCredentialsOption configures CreateSRTCredentials.
type SRTCredentialsOption func(*srtCredentialsOptions)

// WithSRTTokenExpiration sets how long the SRT address works. It defaults to
// DefaultSRTTokenExpiration.
func WithSRTTokenExpiration(d time.Duration) SRTCredentialsOption {
	return func(o *srtCredentialsOptions) {
		o.expiration = d
	}
}

// CreateSRTCredentials creates the SRT address an encoder uses to publish into this call.
//
// The address embeds a user token that can only join this call and expires after 24 hours,
// or after the duration given with WithSRTTokenExpiration. The token is checked again
// whenever the stream reconnects, and at least every 6 hours during a stream, so it must
// stay valid until the stream ends. Treat the address as a secret.
func (c *Call) CreateSRTCredentials(userID string, opts ...SRTCredentialsOption) (*SRTCredentials, error) {
	if c.data == nil {
		return nil, errors.New("call is not initialized, make sure to call .Get or .GetOrCreate first")
	}

	options := srtCredentialsOptions{expiration: DefaultSRTTokenExpiration}
	for _, opt := range opts {
		opt(&options)
	}
	if options.expiration < time.Second {
		return nil, errors.New("SRT token expiration must be at least one second")
	}

	token, err := c.client.client.createToken(userID, &Claims{
		CustomClaims: map[string]interface{}{
			"scoped_call_cid": c.callType + ":" + c.callID,
		},
	}, &options.expiration)
	if err != nil {
		return nil, err
	}

	segments := strings.Split(token, ".")
	if len(segments) != 3 {
		return nil, errors.New("invalid token format")
	}
	claims := jwt.MapClaims{}
	if _, _, err := jwt.NewParser().ParseUnverified(token, claims); err != nil {
		return nil, err
	}
	exp, err := claims.GetExpirationTime()
	if err != nil || exp == nil {
		return nil, errors.New("SRT token has no expiration")
	}

	address := strings.ReplaceAll(c.data.Ingress.Srt.Address, "{passphrase}", segments[2])
	address = strings.ReplaceAll(address, "{token}", token)
	if m := srtStreamIDPattern.FindStringSubmatch(address); m != nil && len(m[1]) > maxSRTStreamIDBytes {
		return nil, fmt.Errorf("the SRT stream id is %d bytes, over the %d-byte SRT limit; use a shorter user ID or call ID", len(m[1]), maxSRTStreamIDBytes)
	}

	return &SRTCredentials{
		Address:   address,
		ExpiresAt: exp.Time,
	}, nil
}
