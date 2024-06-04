package getstream

type Stream struct {
	client *Client
	chat   *ChatClient
	video  *VideoClient
	common *CommonClient
}

func New(apiKey, apiSecret string, options ...ClientOption) *Stream {
	client, err := NewClient(apiKey, apiSecret, options...)
	if err != nil {
		return nil
	}
	return &Stream{
		client: client,
	}
}

// Chat
func (s *Stream) Chat() *ChatClient {
	if s.chat == nil {
		s.chat = NewChatClient(s.client)
	}
	return s.chat
}

// Video
func (s *Stream) Video() *VideoClient {
	if s.video == nil {
		s.video = NewVideoClient(s.client)
	}
	return s.video
}

// common
func (s *Stream) Common() *CommonClient {
	if s.common == nil {
		s.common = NewCommonClient(s.client)
	}
	return s.common
}
