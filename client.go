package getstream

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// DefaultBaseURL is the default base URL for the stream chat api.
	// It works like CDN style and connects you to the closest production server.
	// By default, there is no real reason to change it. Use it only if you know what you are doing.
	DefaultBaseURL = "https://chat.stream-io-api.com"
	defaultTimeout = 6 * time.Second
)

type Client struct {
	BaseURL string
	HTTP    *http.Client `json:"-"`

	apiKey    string
	apiSecret []byte
	authToken string
	logger    *Logger
}

type ClientOption func(c *Client)

// WithLogger sets a custom logger for the client
func WithLogger(l *Logger) ClientOption {
	return func(c *Client) {
		c.logger = l
	}
}

// WithLogLevel sets the log level for the default logger
func WithLogLevel(level LogLevel) ClientOption {
	return func(c *Client) {
		if c.logger == nil {
			c.logger = DefaultLogger
		}
		c.logger.SetLevel(level)
	}
}

// WithBaseUrl sets the base URL for the client.
func WithBaseUrl(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

// NewClientFromEnvVars creates a new Client where the API key
// is retrieved from STREAM_KEY and the secret from STREAM_SECRET
// environmental variables.
func NewClientFromEnvVars(options ...ClientOption) (*Client, error) {
	return NewClient(os.Getenv("STREAM_KEY"), os.Getenv("STREAM_SECRET"), options...)
}

// NewClient creates new stream chat api client.
func NewClient(apiKey, apiSecret string, options ...ClientOption) (*Client, error) {
	switch {
	case apiKey == "":
		return nil, errors.New("API key is empty")
	case apiSecret == "":
		return nil, errors.New("API secret is empty")
	}

	baseURL := DefaultBaseURL
	if baseURLEnv := os.Getenv("STREAM_CHAT_URL"); strings.HasPrefix(baseURLEnv, "http") {
		baseURL = baseURLEnv
	}

	timeout := defaultTimeout
	if timeoutEnv := os.Getenv("STREAM_CHAT_TIMEOUT"); timeoutEnv != "" {
		i, err := strconv.Atoi(timeoutEnv)
		if err != nil {
			return nil, err
		}
		timeout = time.Duration(i) * time.Second
	}

	tr := http.DefaultTransport.(*http.Transport).Clone() //nolint:forcetypeassert
	tr.MaxIdleConnsPerHost = 5
	tr.IdleConnTimeout = 59 * time.Second // load balancer's idle timeout is 60 sec
	tr.ExpectContinueTimeout = 2 * time.Second

	client := &Client{
		apiKey:    apiKey,
		apiSecret: []byte(apiSecret),
		BaseURL:   baseURL,
		HTTP: &http.Client{
			Timeout:   timeout,
			Transport: tr,
		},
		logger: DefaultLogger,
	}

	for _, fn := range options {
		fn(client)
	}

	token, err := client.createToken(jwt.MapClaims{"server": true})
	if err != nil {
		return nil, err
	}

	client.authToken = token

	return client, nil
}

type StreamJWTClaims struct {
	Expire      *time.Time
	IssuedAt    *time.Time
	ChannelCIDs []string
	CallCIDs    []string
	Role        string
}

func (c *Client) CreateTokenWithClaims(userID string, claims *StreamJWTClaims) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is empty")
	}

	jwtClaims := jwt.MapClaims{
		"user_id": userID,
	}

	// Set issued at time; use the provided time or the current time if not provided.
	if claims != nil {
		if claims.IssuedAt != nil && !claims.IssuedAt.IsZero() {
			jwtClaims["iat"] = claims.IssuedAt.Unix()
		} else {
			now := time.Now()
			jwtClaims["iat"] = now.Unix()
		}

		// Set expiration time if provided.
		if claims.Expire != nil && !claims.Expire.IsZero() {
			jwtClaims["exp"] = claims.Expire.Unix()
		}

		// Add channel IDs if provided.
		if len(claims.ChannelCIDs) > 0 {
			jwtClaims["channel_cids"] = claims.ChannelCIDs
		}

		// Add call IDs if provided.
		if len(claims.CallCIDs) > 0 {
			jwtClaims["call_cids"] = claims.CallCIDs
		}

		// Add role if provided.
		if claims.Role != "" {
			jwtClaims["role"] = claims.Role
		}
	}

	return c.createToken(jwtClaims)
}

func (c *Client) createToken(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(c.apiSecret)
}

// VerifyWebhook validates if hmac signature is correct for message body.
func (c *Client) VerifyWebhook(body, signature []byte) (valid bool) {
	mac := hmac.New(crypto.SHA256.New, c.apiSecret)
	_, _ = mac.Write(body)

	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return bytes.Equal(signature, []byte(expectedMAC))
}
