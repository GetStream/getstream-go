package getstream

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// DefaultSRTTokenExpiration is how long the token in an SRT address works unless
// WithSRTTokenExpiration sets another duration.
const DefaultSRTTokenExpiration = 24 * time.Hour

// SRTCredentials holds the SRT address an encoder (OBS, ffmpeg, ...) publishes to.
//
// Address contains a user token, so keep it secret.
type SRTCredentials struct {
	Address string
	// ExpiresAt is when the token in Address expires.
	ExpiresAt time.Time
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

	return &SRTCredentials{
		Address:   address,
		ExpiresAt: exp.Time,
	}, nil
}
