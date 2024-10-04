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

func (s *Stream) CreateToken(userID string, claims *StreamJWTClaims) (string, error) {
	return s.CreateTokenWithClaims(userID, claims)
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
