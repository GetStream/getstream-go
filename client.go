package getstream

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"errors"
	"fmt"
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

func PtrTo[T any](v T) *T {
	return &v
}

type Client struct {
	apiKey         string
	apiSecret      []byte
	authToken      string
	baseUrl        string
	defaultTimeout time.Duration
	httpClient     *http.Client
	logger         *Logger
}

func (c *Client) HttpClient() *http.Client {
	return c.httpClient
}

func (c *Client) Logger() *Logger {
	return c.logger
}

func (c *Client) ApiKey() string {
	return c.apiKey
}

func (c *Client) BaseUrl() string {
	return c.baseUrl
}

func (c *Client) DefaultTimeout() time.Duration {
	return c.defaultTimeout
}

type ClientOption func(c *Client)

// WithTimeout sets a custom timeout for all API requests
func WithTimeout(t time.Duration) ClientOption {
	return func(c *Client) {
		c.defaultTimeout = t
	}
}

// WithBaseUrl sets the base URL for the client.
func WithBaseUrl(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseUrl = baseURL
	}
}

const (
	EnvStreamApiKey      = "STREAM_API_KEY"
	EnvStreamApiSecret   = "STREAM_API_SECRET"
	EnvStreamBaseUrl     = "STREAM_BASE_URL"
	EnvStreamHttpTimeout = "STREAM_HTTP_TIMEOUT"
)

func newClientFromEnvVars(options ...ClientOption) (*Client, error) {
	apiKey := os.Getenv(EnvStreamApiKey)
	if apiKey == "" {
		return nil, errors.New(EnvStreamApiKey + " is empty")
	}
	apiSecret := os.Getenv(EnvStreamApiSecret)
	if apiSecret == "" {
		return nil, errors.New(EnvStreamApiSecret + " is empty")
	}
	return newClient(apiKey, apiSecret, options...)
}

func newClient(apiKey, apiSecret string, options ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("API key is empty")
	}

	if apiSecret == "" {
		return nil, errors.New("API secret is empty")
	}

	baseURL := DefaultBaseURL
	if baseURLEnv := os.Getenv(EnvStreamBaseUrl); strings.HasPrefix(baseURLEnv, "http") {
		baseURL = baseURLEnv
	}

	client := &Client{
		apiKey:         apiKey,
		apiSecret:      []byte(apiSecret),
		baseUrl:        baseURL,
		logger:         DefaultLogger,
		defaultTimeout: defaultTimeout,
	}

	if timeoutEnv := os.Getenv(EnvStreamHttpTimeout); timeoutEnv != "" {
		i, err := strconv.Atoi(timeoutEnv)
		if err != nil {
			return nil, fmt.Errorf("cannot convert "+EnvStreamHttpTimeout+" into a valid timeout %w", err)
		}
		client.defaultTimeout = time.Duration(i) * time.Second
	}

	for _, fn := range options {
		fn(client)
	}

	tr := http.DefaultTransport.(*http.Transport).Clone() //nolint:forcetypeassert
	tr.MaxIdleConnsPerHost = 5
	tr.IdleConnTimeout = 59 * time.Second
	tr.ExpectContinueTimeout = 2 * time.Second

	client.httpClient = &http.Client{
		Timeout:   client.defaultTimeout,
		Transport: tr,
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
