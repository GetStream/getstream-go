package getstream

import "context"

type Channels struct {
	channelType string
	channelD    string
	client      *ChatClient
}

func NewChannel(channelType string, channelD string, client *ChatClient) *Channels {
	return &Channels{
		channelType: channelType,
		channelD:    channelD,
		client:      client,
	}
}

func (c *Channels) Delete(ctx context.Context, hardDelete *bool) (*DeleteChannelResponse, error) {
	return c.client.DeleteChannel(ctx, c.channelType, c.channelD, hardDelete)
}

func (c *Channels) UpdateChannelPartial(ctx context.Context, request *UpdateChannelPartialRequest) (*UpdateChannelPartialResponse, error) {
	return c.client.UpdateChannelPartial(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) Update(ctx context.Context, request *UpdateChannelRequest) (*UpdateChannelResponse, error) {
	return c.client.UpdateChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) SendEvent(ctx context.Context, request *SendEventRequest) (*EventResponse, error) {
	return c.client.SendEvent(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) DeleteFile(ctx context.Context, url *string) (*FileDeleteResponse, error) {
	return c.client.DeleteFile(ctx, c.channelType, c.channelD, url)
}

func (c *Channels) UploadFile(ctx context.Context, request *FileUploadRequest) (*FileUploadResponse, error) {
	return c.client.UploadFile(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) Hide(ctx context.Context, request *HideChannelRequest) (*HideChannelResponse, error) {
	return c.client.HideChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) DeleteImage(ctx context.Context, url *string) (*FileDeleteResponse, error) {
	return c.client.DeleteImage(ctx, c.channelType, c.channelD, url)
}

func (c *Channels) UploadImage(ctx context.Context, request *ImageUploadRequest) (*ImageUploadResponse, error) {
	return c.client.UploadImage(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) SendMessage(ctx context.Context, request *SendMessageRequest) (*SendMessageResponse, error) {
	return c.client.SendMessage(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) GetManyMessages(ctx context.Context, ids []string) (*GetManyMessagesResponse, error) {
	return c.client.GetManyMessages(ctx, c.channelType, c.channelD, ids)
}

func (c *Channels) GetOrCreate(ctx context.Context, request *ChannelGetOrCreateRequest) (*ChannelStateResponse, error) {
	return c.client.GetOrCreateChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) MarkRead(ctx context.Context, request *MarkReadRequest) (*MarkReadResponse, error) {
	return c.client.MarkRead(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) Show(ctx context.Context, request *ShowChannelRequest) (*ShowChannelResponse, error) {
	return c.client.ShowChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) Truncate(ctx context.Context, request *TruncateChannelRequest) (*TruncateChannelResponse, error) {
	return c.client.TruncateChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) MarkUnread(ctx context.Context, request *MarkUnreadRequest) (*Response, error) {
	return c.client.MarkUnread(ctx, c.channelType, c.channelD, request)
}

func (c *ChatClient) Channel(channelType, channelD string) *Channels {
	return NewChannel(channelType, channelD, c)
}
