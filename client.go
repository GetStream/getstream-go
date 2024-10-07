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

type tokenOptions struct {
	claims     *Claims
	expiration *time.Duration
}

type TokenOption func(*tokenOptions)

func WithExpiration(d time.Duration) TokenOption {
	return func(t *tokenOptions) {
		t.expiration = &d
	}
}

func WithClaims(claims Claims) TokenOption {
	return func(t *tokenOptions) {
		t.claims = &claims
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
	tr.ExpectContinueTimeout = 2 * time.Second
	tr.IdleConnTimeout = 59 * time.Second

	client.httpClient = &http.Client{
		Timeout:   client.defaultTimeout,
		Transport: tr,
	}

	token, err := client.createTokenWithClaims(jwt.MapClaims{"server": true})
	if err != nil {
		return nil, err
	}

	client.authToken = token
	return client, nil
}

// Claims contains optional parameters for token creation.
type Claims struct {
	Role         string                 // Role assigned to the user
	ChannelCIDs  []string               // Channel IDs the user has access to
	CallCIDs     []string               // Call IDs the user has access to
	CustomClaims map[string]interface{} // Additional custom claims
}

func (c *Client) createToken(userID string, claims *Claims, expiration *time.Duration) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is required")
	}

	now := time.Now().Unix()
	jwtClaims := jwt.MapClaims{
		"user_id": userID,
		"iat":     now,
	}

	if expiration != nil && *expiration > 0 {
		jwtClaims["exp"] = now + int64(expiration.Seconds())
	}

	if claims != nil {
		for key, value := range claims.CustomClaims {
			jwtClaims[key] = value
		}

		if claims.Role != "" {
			jwtClaims["role"] = claims.Role
		}

		if len(claims.ChannelCIDs) > 0 {
			jwtClaims["channel_cids"] = claims.ChannelCIDs
		}

		if len(claims.CallCIDs) > 0 {
			jwtClaims["call_cids"] = claims.CallCIDs
		}
	}

	return c.createTokenWithClaims(jwtClaims)
}

func (c *Client) createCallToken(userID string, claims *Claims, expiration *time.Duration) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is required")
	}

	// Ensure that CallCIDs are included for call tokens
	if claims == nil {
		claims = &Claims{}
	}
	if len(claims.CallCIDs) == 0 {
		return "", errors.New("call_cids are required for call tokens")
	}

	return c.createToken(userID, claims, expiration)
}

// createToken signs the JWT with the provided claims.
//
// Parameters:
// - claims (jwt.MapClaims): The claims to include in the token.
//
// Returns:
// - (string): The signed JWT token.
// - (error): An error object if signing fails.
func (c *Client) createTokenWithClaims(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(c.apiSecret)
}

// VerifyWebhook validates if hmac signature is correct for message body.
func (c *Client) VerifyWebhook(body, signature []byte) (valid bool) {
	mac := hmac.New(crypto.SHA256.New, c.apiSecret)
	_, _ = mac.Write(body)

	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return bytes.Equal(signature, []byte(expectedMAC))
}
