package getstream

import "net/http"

type Stream struct {
	*Client
	chat       *ChatClient
	video      *VideoClient
	feeds      *FeedsClient
	moderation *ModerationClient
}

func NewClientFromEnvVars(options ...ClientOption) (*Stream, error) {
	client, err := newClientFromEnvVars(options...)
	if err != nil {
		return nil, err
	}
	return &Stream{
		Client: client,
	}, nil
}

func NewClient(apiKey, apiSecret string, options ...ClientOption) (*Stream, error) {
	client, err := newClient(apiKey, apiSecret, options...)
	if err != nil {
		return nil, err
	}
	return &Stream{
		Client: client,
	}, nil
}

// CreateToken generates a token for a given user ID, with optional claims.
//
// Parameters:
// - userID (string): The unique identifier of the user for whom the token is being created.
// - claims (*Claims): A pointer to a Claims struct containing optional parameters.
//
// Returns:
// - (string): The generated JWT token.
// - (error): An error object if token creation fails.
//
// token, err := client.CreateToken("userID", getstream.WithExpiration(time.Hour))
func (s *Stream) CreateToken(userID string, opts ...TokenOption) (string, error) {
	o := tokenOptions{}
	for _, opt := range opts {
		opt(&o)
	}
	return s.createToken(userID, o.claims, o.expiration)
}

func (s *Stream) Chat() *ChatClient {
	if s.chat == nil {
		s.chat = NewChatClient(s.Client)
	}
	return s.chat
}

func (s *Stream) Video() *VideoClient {
	if s.video == nil {
		s.video = NewVideoClient(s.Client)
	}
	return s.video
}

// Feeds client
func (s *Stream) Feeds() *FeedsClient {
	if s.feeds == nil {
		s.feeds = NewFeedsClient(s.Client)
	}
	return s.feeds
}

// Moderation client
func (s *Stream) Moderation() *ModerationClient {
	if s.moderation == nil {
		s.moderation = NewModerationClient(s.Client)
	}
	return s.moderation
}

// VerifyWebhookSignature verifies the HMAC-SHA256 signature of a webhook body
// using this client's API secret. Convenience wrapper around the package-level
// VerifyWebhookSignature function — drops the secret parameter in favor of
// the secret stored on the client.
func (s *Stream) VerifyWebhookSignature(body []byte, signature string) bool {
	return VerifyWebhookSignature(body, signature, string(s.apiSecret))
}

// VerifyAndParseWebhook verifies and parses a webhook payload from an
// *http.Request using this client's API secret. The request body is restored
// so downstream handlers can read it again. Convenience wrapper around the
// package-level VerifyAndParseWebhook — drops the secret parameter.
func (s *Stream) VerifyAndParseWebhook(r *http.Request) (WebhookEvent, error) {
	return VerifyAndParseWebhook(r, string(s.apiSecret))
}

// VerifyAndParseWebhookBytes verifies and parses a webhook payload (raw bytes)
// using this client's API secret. Convenience wrapper around the package-level
// VerifyAndParseWebhookBytes — drops the secret parameter.
func (s *Stream) VerifyAndParseWebhookBytes(body []byte, signature string) (WebhookEvent, error) {
	return VerifyAndParseWebhookBytes(body, signature, string(s.apiSecret))
}

// ParseSqs is a convenience wrapper that calls the package-level ParseSqs.
// No signature is required; SQS deliveries are authenticated via AWS IAM.
func (s *Stream) ParseSqs(messageBody string) (WebhookEvent, error) {
	return ParseSqs(messageBody)
}

// ParseSns is a convenience wrapper that calls the package-level ParseSns.
// No signature is required; SNS deliveries are authenticated via AWS IAM.
func (s *Stream) ParseSns(notificationBody string) (WebhookEvent, error) {
	return ParseSns(notificationBody)
}
