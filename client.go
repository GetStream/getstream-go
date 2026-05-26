package getstream

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
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

	// DefaultRequestTimeout is the default per-request timeout (was 6s prior to v4.2.0).
	DefaultRequestTimeout = 30 * time.Second
	// DefaultMaxConnsPerHost caps concurrent TCP connections per host.
	DefaultMaxConnsPerHost = 5
	// DefaultIdleTimeout sits below the typical 60s LB idle timeout with a 5s safety margin.
	DefaultIdleTimeout = 55 * time.Second
	// DefaultConnectTimeout caps TCP + TLS handshake duration.
	DefaultConnectTimeout = 10 * time.Second

	// defaultTimeout is preserved for backwards compatibility of internal references.
	defaultTimeout = DefaultRequestTimeout
)

func PtrTo[T any](v T) *T {
	return &v
}

type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

type Client struct {
	apiKey             string
	apiSecret          []byte
	authToken          string
	baseUrl            string
	defaultTimeout     time.Duration
	maxConnsPerHost    int
	idleTimeout        time.Duration
	connectTimeout     time.Duration
	httpClient         HttpClient
	httpClientFromUser bool // true iff WithHTTPClient was used; gates transport build
	logger             Logger
}

func (c *Client) HttpClient() HttpClient {
	return c.httpClient
}

func (c *Client) Logger() Logger {
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

func WithHTTPClient(httpClient HttpClient) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
		c.httpClientFromUser = true
	}
}

// WithTimeout sets a custom timeout for all API requests
func WithTimeout(t time.Duration) ClientOption {
	return func(c *Client) {
		c.defaultTimeout = t
	}
}

// WithMaxConnsPerHost caps concurrent TCP connections per host. Default: 5.
// Ignored when WithHTTPClient is set.
func WithMaxConnsPerHost(n int) ClientOption {
	return func(c *Client) {
		c.maxConnsPerHost = n
	}
}

// WithIdleTimeout sets how long an idle connection lingers before being closed.
// Default: 55s (sits 5s below the typical 60s LB idle timeout). Ignored when
// WithHTTPClient is set.
func WithIdleTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		c.idleTimeout = d
	}
}

// WithConnectTimeout caps TCP+TLS handshake duration. Default: 10s. Ignored
// when WithHTTPClient is set.
func WithConnectTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		c.connectTimeout = d
	}
}

// WithRequestTimeout sets the default per-request timeout. Default: 30s.
// Callers can still override per-call via context.WithTimeout. Ignored when
// WithHTTPClient is set.
func WithRequestTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		c.defaultTimeout = d
	}
}

// WithBaseUrl sets the base URL for the client.
func WithBaseUrl(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseUrl = baseURL
	}
}

// WithLogger sets a custom logger for the client.
func WithLogger(logger Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithAuthToken sets the auth token for the client.
func WithAuthToken(authToken string) ClientOption {
	return func(c *Client) {
		c.authToken = authToken
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

// buildDefaultHTTPClient constructs the SDK's default *http.Client with the
// spec-mandated transport tuning. It clones http.DefaultTransport so any
// runtime-provided ProxyFromEnvironment, ALPN, etc. defaults are preserved.
func buildDefaultHTTPClient(requestTimeout time.Duration, maxConnsPerHost int, idleTimeout, connectTimeout time.Duration) *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = maxConnsPerHost
	transport.MaxIdleConnsPerHost = maxConnsPerHost
	transport.IdleConnTimeout = idleTimeout
	transport.DialContext = (&net.Dialer{
		Timeout:   connectTimeout,
		KeepAlive: 30 * time.Second, // OS-level TCP keep-alive; unrelated to HTTP keep-alive
	}).DialContext
	transport.DisableKeepAlives = false // §5 invariant 4

	return &http.Client{
		Timeout:   requestTimeout,
		Transport: transport,
	}
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
		apiKey:          apiKey,
		apiSecret:       []byte(apiSecret),
		baseUrl:         baseURL,
		defaultTimeout:  DefaultRequestTimeout,
		maxConnsPerHost: DefaultMaxConnsPerHost,
		idleTimeout:     DefaultIdleTimeout,
		connectTimeout:  DefaultConnectTimeout,
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

	// Set default logger if not provided
	if client.logger == nil {
		client.logger = DefaultLoggerInstance
	}

	if client.httpClient == nil {
		client.httpClient = buildDefaultHTTPClient(
			client.defaultTimeout,
			client.maxConnsPerHost,
			client.idleTimeout,
			client.connectTimeout,
		)
	}

	if client.authToken == "" {
		token, err := client.createTokenWithClaims(jwt.MapClaims{"server": true})
		if err != nil {
			return nil, err
		}
		client.authToken = token
	}
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
