package getstream

type Chat struct {
	client *Client
}

// new
func NewChat(client *Client) *Chat {
	return &Chat{
		client: client,
	}
}
