package getstream

type Stream struct {
	*Client
	chat  *ChatClient
	video *VideoClient
}

func NewStreamFromEnvVars(options ...ClientOption) (*Stream, error) {
	client, err := NewClientFromEnvVars(options...)
	if err != nil {
		return nil, err
	}
	return &Stream{
		Client: client,
	}, nil
}

func New(apiKey, apiSecret string, options ...ClientOption) *Stream {
	client, err := NewClient(apiKey, apiSecret, options...)
	if err != nil {
		return nil
	}
	return &Stream{
		Client: client,
	}
}

func (s *Stream) CreateToken(userID string, claims *StreamJWTClaims) (string, error) {
	return s.CreateToken(userID, claims)
}

// Chat
func (s *Stream) Chat() *ChatClient {
	if s.chat == nil {
		s.chat = NewChatClient(s.Client)
	}
	return s.chat
}

// Video
func (s *Stream) Video() *VideoClient {
	if s.video == nil {
		s.video = NewVideoClient(s.Client)
	}
	return s.video
}
