package getstream

type Stream struct {
	*Client
	chat  *ChatClient
	video *VideoClient
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
