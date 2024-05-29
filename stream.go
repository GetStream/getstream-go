package getstream

type Stream struct {
	client *Client
	chat   *Chat
	video  *Video
}

// New
func New(apiKey, apiSecret string) *Stream {
	client, err := NewClient(apiKey, apiSecret)
	if err != nil {
		return nil
	}
	return &Stream{
		client: client,
	}
}

// Chat
func (s *Stream) Chat() *Chat {
	if s.chat == nil {
		s.chat = NewChat(s.client)
	}
	return s.chat
}

// Video
func (s *Stream) Video() *Video {
	if s.video == nil {
		s.video = NewVideo(s.client)
	}
	return s.video
}
