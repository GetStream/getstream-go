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

func (c *Client) createToken(userID string, claims *Claims, expiration int64) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is required")
	}

	now := time.Now().Unix()
	jwtClaims := jwt.MapClaims{
		"user_id": userID,
		"iat":     now,
	}

	if claims != nil {
		// Set expiration time if provided
		if expiration > 0 {
			jwtClaims["exp"] = now + expiration
		}

		// Set role if provided
		if claims.Role != "" {
			jwtClaims["role"] = claims.Role
		}

		// Set channel CIDs if provided
		if len(claims.ChannelCIDs) > 0 {
			jwtClaims["channel_cids"] = claims.ChannelCIDs
		}

		// Set call CIDs if provided
		if len(claims.CallCIDs) > 0 {
			jwtClaims["call_cids"] = claims.CallCIDs
		}

		// Set custom claims if provided
		if len(claims.CustomClaims) > 0 {
			for key, value := range claims.CustomClaims {
				jwtClaims[key] = value
			}
		}
	}

	return c.createTokenWithClaims(jwtClaims)
}

func (c *Client) createCallToken(userID string, claims *Claims, expiration int64) (string, error) {
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
