package getstream

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"encoding/json"
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

	// defaultRequestTimeout is the default per-request timeout (was 6s prior to v4.2.0).
	defaultRequestTimeout = 30 * time.Second
	// defaultMaxConnsPerHost caps concurrent TCP connections per host and sizes the idle pool; extra requests wait for a free connection.
	defaultMaxConnsPerHost = 100
	// defaultIdleTimeout sits below the typical 60s LB idle timeout with a 5s safety margin.
	defaultIdleTimeout = 55 * time.Second
	// defaultConnectTimeout caps TCP + TLS handshake duration.
	defaultConnectTimeout = 10 * time.Second
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
	logBodies          bool // true iff WithLogBodies(true) was used; gates body fields on DEBUG events
	retry              RetryConfig
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

func (c *Client) MaxConnsPerHost() int {
	return c.maxConnsPerHost
}

func (c *Client) IdleTimeout() time.Duration {
	return c.idleTimeout
}

func (c *Client) ConnectTimeout() time.Duration {
	return c.connectTimeout
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

// WithMaxConnsPerHost caps concurrent TCP connections per host, which is also the number of idle connections kept for reuse. Default: 100.
// Ignored when WithHTTPClient is set.
func WithMaxConnsPerHost(n int) ClientOption {
	return func(c *Client) {
		c.maxConnsPerHost = n
	}
}

// WithIdleTimeout sets how long an idle connection lingers before being closed.
// Default: 55s (sits 5s below the typical 60s LB idle timeout). Ignored when WithHTTPClient is set.
func WithIdleTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		c.idleTimeout = d
	}
}

// WithConnectTimeout caps TCP+TLS handshake duration. Default: 10s. Ignored when WithHTTPClient is set.
func WithConnectTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		c.connectTimeout = d
	}
}

// WithRequestTimeout sets the default per-request timeout. Default: 30s.
// Callers can still override per-call via context.WithTimeout. Ignored when WithHTTPClient is set.
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

// WithLogBodies opts in to logging request/response bodies on the DEBUG
// events. Known-secret body keys are still redacted. Off by default.
func WithLogBodies(enabled bool) ClientOption {
	return func(c *Client) {
		c.logBodies = enabled
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

// buildDefaultHTTPClient constructs the SDK's default *http.Client with the spec-mandated transport tuning.
// It clones http.DefaultTransport so any runtime-provided ProxyFromEnvironment, ALPN, etc. defaults are preserved.
func buildDefaultHTTPClient(requestTimeout time.Duration, maxConnsPerHost int, idleTimeout, connectTimeout time.Duration) *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = maxConnsPerHost
	transport.MaxIdleConnsPerHost = maxConnsPerHost
	transport.MaxIdleConns = 0 // DefaultTransport's global 100 would otherwise cap the per-host idle pool
	transport.IdleConnTimeout = idleTimeout
	transport.DialContext = (&net.Dialer{
		Timeout:   connectTimeout,
		KeepAlive: 30 * time.Second, // OS-level TCP keep-alive; unrelated to HTTP keep-alive
	}).DialContext
	transport.DisableKeepAlives = false // keep-alive stays on; never emit Connection: close

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
		defaultTimeout:  defaultRequestTimeout,
		maxConnsPerHost: defaultMaxConnsPerHost,
		idleTimeout:     defaultIdleTimeout,
		connectTimeout:  defaultConnectTimeout,
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

	client.logger.Info(
		"client.initialized stream.sdk.name=getstream-go stream.sdk.version=%s stream.client.max_conns_per_host=%d stream.client.idle_timeout_seconds=%d stream.client.connect_timeout_seconds=%d stream.client.request_timeout_seconds=%d stream.client.gzip_enabled=%t stream.client.user_http_client=%t stream.client.log_bodies=%t",
		versionName, client.maxConnsPerHost, int(client.idleTimeout.Seconds()), int(client.connectTimeout.Seconds()), int(client.defaultTimeout.Seconds()), true, client.httpClientFromUser, client.logBodies,
	)
	if client.logBodies {
		client.logger.Warn("HTTP request/response bodies will be logged. Auth headers and known-secret fields are still redacted, but other sensitive data (messages, PII) may appear in logs. Disable for production.")
	}

	return client, nil
}

// Claims contains optional parameters for token creation.
type Claims struct {
	Role        string   // Role assigned to the user
	ChannelCIDs []string // Channel IDs the user has access to
	CallCIDs    []string // Call IDs the user has access to
	// CustomClaims are additional claims added to the token. CreateToken
	// returns an error if they set "user_id", "server" or "egress", if a key is
	// another spelling of a claim the backend reads (such as "User_ID"), or if
	// "iat" is not a positive whole Unix timestamp that is not in the future.
	// A custom "exp" is ignored when WithExpiration is used, and Role,
	// ChannelCIDs and CallCIDs take precedence over custom claims of the same name.
	CustomClaims map[string]interface{}
}

// reservedCustomClaims are claims a user token's custom claims may not set.
// The token's identity comes only from the userID argument (the backend treats
// a token without a user_id as a server token), and "egress" lets a token join
// calls without the app's role grants.
var reservedCustomClaims = []string{"user_id", "server", "egress"}

// backendClaims are the claim names the backend reads from a token. It decodes
// them with encoding/json, which matches names case-insensitively with Unicode
// simple folding (so "uſer_id" and "USER_ID" both set user_id) and lets the last
// matching key win, so a custom claim may use only the exact spelling.
var backendClaims = []string{
	"user_id", "call_cids", "channel_cids", "role", "egress", "stream_check_writes",
	"scoped_call_cid", "iss", "sub", "aud", "exp", "nbf", "iat", "jti", "server",
}

// checkCustomClaims returns an error if a custom claim may not be set on a user token.
func checkCustomClaims(customClaims map[string]interface{}, now int64) error {
	for key, value := range customClaims {
		for _, name := range backendClaims {
			// strings.EqualFold uses the same simple folding as encoding/json,
			// including the Kelvin sign (k) and the long s (s).
			if key != name && strings.EqualFold(key, name) {
				return fmt.Errorf("custom claim %q is read by the backend as %q; use %q", key, name, name)
			}
		}
		for _, reserved := range reservedCustomClaims {
			if key == reserved {
				return fmt.Errorf("custom claim %q is reserved and cannot be set on a user token", key)
			}
		}
		if key == "iat" {
			// The backend skips per-user token revocation for a token without
			// an iat, and a later iat outlives revocations made before it.
			if iat, ok := wholeNumber(value); !ok || iat <= 0 || iat > now {
				return errors.New(`custom claim "iat" must be a positive whole Unix timestamp that is not in the future`)
			}
		}
	}
	return nil
}

// wholeNumber returns v as an int64 if it is a whole number.
func wholeNumber(v interface{}) (int64, bool) {
	switch n := v.(type) {
	case int:
		return int64(n), true
	case int32:
		return int64(n), true
	case int64:
		return n, true
	case float64:
		if n != float64(int64(n)) {
			return 0, false
		}
		return int64(n), true
	case json.Number:
		i, err := n.Int64()
		return i, err == nil
	}
	return 0, false
}

func (c *Client) createToken(userID string, claims *Claims, expiration *time.Duration) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is required")
	}

	now := time.Now().Unix()
	jwtClaims := jwt.MapClaims{"iat": now}

	if claims != nil {
		if err := checkCustomClaims(claims.CustomClaims, now); err != nil {
			return "", err
		}
		for key, value := range claims.CustomClaims {
			jwtClaims[key] = value
		}
	}

	jwtClaims["user_id"] = userID
	if expiration != nil && *expiration > 0 {
		jwtClaims["exp"] = now + int64(expiration.Seconds())
	}

	if claims != nil {
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
