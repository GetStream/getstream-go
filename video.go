package getstream

type Video struct {
	client *Client
}

func NewVideo(client *Client) *Video {
	return &Video{
		client: client,
	}
}
