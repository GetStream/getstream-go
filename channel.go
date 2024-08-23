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

func (c *Channels) Delete(ctx context.Context, request *DeleteChannelRequest) (*StreamResponse[DeleteChannelResponse], error) {
	return c.client.DeleteChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) UpdateChannelPartial(ctx context.Context, request *UpdateChannelPartialRequest) (*StreamResponse[UpdateChannelPartialResponse], error) {
	return c.client.UpdateChannelPartial(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) Update(ctx context.Context, request *UpdateChannelRequest) (*StreamResponse[UpdateChannelResponse], error) {
	return c.client.UpdateChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) SendEvent(ctx context.Context, request *SendEventRequest) (*StreamResponse[EventResponse], error) {
	return c.client.SendEvent(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) DeleteFile(ctx context.Context, request *DeleteFileRequest) (*StreamResponse[Response], error) {
	return c.client.DeleteFile(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) UploadFile(ctx context.Context, request *FileUploadRequest) (*StreamResponse[FileUploadResponse], error) {
	return c.client.UploadFile(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) Hide(ctx context.Context, request *HideChannelRequest) (*StreamResponse[HideChannelResponse], error) {
	return c.client.HideChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) DeleteImage(ctx context.Context, request *DeleteImageRequest) (*StreamResponse[Response], error) {
	return c.client.DeleteImage(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) UploadImage(ctx context.Context, request *ImageUploadRequest) (*StreamResponse[ImageUploadResponse], error) {
	return c.client.UploadImage(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) SendMessage(ctx context.Context, request *SendMessageRequest) (*StreamResponse[SendMessageResponse], error) {
	return c.client.SendMessage(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) GetManyMessages(ctx context.Context, request *GetManyMessagesRequest) (*StreamResponse[GetManyMessagesResponse], error) {
	return c.client.GetManyMessages(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) GetOrCreate(ctx context.Context, request *ChannelGetOrCreateRequest) (*StreamResponse[ChannelStateResponse], error) {
	return c.client.GetOrCreateChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) MarkRead(ctx context.Context, request *MarkReadRequest) (*StreamResponse[MarkReadResponse], error) {
	return c.client.MarkRead(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) Show(ctx context.Context, request *ShowChannelRequest) (*StreamResponse[ShowChannelResponse], error) {
	return c.client.ShowChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) Truncate(ctx context.Context, request *TruncateChannelRequest) (*StreamResponse[TruncateChannelResponse], error) {
	return c.client.TruncateChannel(ctx, c.channelType, c.channelD, request)
}

func (c *Channels) MarkUnread(ctx context.Context, request *MarkUnreadRequest) (*StreamResponse[Response], error) {
	return c.client.MarkUnread(ctx, c.channelType, c.channelD, request)
}

func (c *ChatClient) Channel(channelType, channelD string) *Channels {
	return NewChannel(channelType, channelD, c)
}
