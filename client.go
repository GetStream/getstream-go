package getstream

import (
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
}

type ClientOption func(c *Client)

// WithBaseUrl sets the base URL for the client.
func WithBaseUrl(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

// NewClientFromEnvVars creates a new Client where the API key
// is retrieved from STREAM_KEY and the secret from STREAM_SECRET
// environmental variables.
func NewClientFromEnvVars() (*Client, error) {
	return NewClient(os.Getenv("STREAM_KEY"), os.Getenv("STREAM_SECRET"))
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

func (c *Client) CreateToken(userID string, expire time.Time, issuedAt ...time.Time) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is empty")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
	}
	if !expire.IsZero() {
		claims["exp"] = expire.Unix()
	}
	if len(issuedAt) > 0 && !issuedAt[0].IsZero() {
		claims["iat"] = issuedAt[0].Unix()
	}

	return c.createToken(claims)
}

func (c *Client) createToken(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(c.apiSecret)
}
