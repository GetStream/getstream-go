package getstream

import "context"

type Channels struct {
	channelType string
	channelD    string
	client      *ChatClient
}

func NewChannel(channelType string, channelD string, client *ChatClient) Channels {
	return Channels{
		channelType: channelType,
		channelD:    channelD,
		client:      client,
	}
}

func (c *Channels) Delete(ctx context.Context, hardDelete *bool) (*DeleteChannelResponse, error) {
	return c.client.DeleteChannel(ctx, c.channelType, c.channelD, hardDelete)
}

func (c *Channels) UpdateChannelPartial(ctx context.Context, updateChannelPartialRequest UpdateChannelPartialRequest) (*UpdateChannelPartialResponse, error) {
	return c.client.UpdateChannelPartial(ctx, c.channelType, c.channelD, updateChannelPartialRequest)
}

func (c *Channels) Update(ctx context.Context, updateChannelRequest UpdateChannelRequest) (*UpdateChannelResponse, error) {
	return c.client.UpdateChannel(ctx, c.channelType, c.channelD, updateChannelRequest)
}

func (c *Channels) SendEvent(ctx context.Context, sendEventRequest SendEventRequest) (*EventResponse, error) {
	return c.client.SendEvent(ctx, c.channelType, c.channelD, sendEventRequest)
}

func (c *Channels) DeleteFile(ctx context.Context, url *string) (*FileDeleteResponse, error) {
	return c.client.DeleteFile(ctx, c.channelType, c.channelD, url)
}

func (c *Channels) UploadFile(ctx context.Context, fileUploadRequest FileUploadRequest) (*FileUploadResponse, error) {
	return c.client.UploadFile(ctx, c.channelType, c.channelD, fileUploadRequest)
}

func (c *Channels) Hide(ctx context.Context, hideChannelRequest HideChannelRequest) (*HideChannelResponse, error) {
	return c.client.HideChannel(ctx, c.channelType, c.channelD, hideChannelRequest)
}

func (c *Channels) DeleteImage(ctx context.Context, url *string) (*FileDeleteResponse, error) {
	return c.client.DeleteImage(ctx, c.channelType, c.channelD, url)
}

func (c *Channels) UploadImage(ctx context.Context, imageUploadRequest ImageUploadRequest) (*ImageUploadResponse, error) {
	return c.client.UploadImage(ctx, c.channelType, c.channelD, imageUploadRequest)
}

func (c *Channels) SendMessage(ctx context.Context, sendMessageRequest SendMessageRequest) (*SendMessageResponse, error) {
	return c.client.SendMessage(ctx, c.channelType, c.channelD, sendMessageRequest)
}

func (c *Channels) GetManyMessages(ctx context.Context, ids []string) (*GetManyMessagesResponse, error) {
	return c.client.GetManyMessages(ctx, c.channelType, c.channelD, ids)
}

func (c *Channels) GetOrCreate(ctx context.Context, channelGetOrCreateRequest ChannelGetOrCreateRequest) (*ChannelStateResponse, error) {
	return c.client.GetOrCreateChannel(ctx, c.channelType, c.channelD, channelGetOrCreateRequest)
}

func (c *Channels) MarkRead(ctx context.Context, markReadRequest MarkReadRequest) (*MarkReadResponse, error) {
	return c.client.MarkRead(ctx, c.channelType, c.channelD, markReadRequest)
}

func (c *Channels) Show(ctx context.Context, showChannelRequest ShowChannelRequest) (*ShowChannelResponse, error) {
	return c.client.ShowChannel(ctx, c.channelType, c.channelD, showChannelRequest)
}

func (c *Channels) Truncate(ctx context.Context, truncateChannelRequest TruncateChannelRequest) (*TruncateChannelResponse, error) {
	return c.client.TruncateChannel(ctx, c.channelType, c.channelD, truncateChannelRequest)
}

func (c *Channels) MarkUnread(ctx context.Context, markUnreadRequest MarkUnreadRequest) (*Response, error) {
	return c.client.MarkUnread(ctx, c.channelType, c.channelD, markUnreadRequest)
}

func (c *ChatClient) Channel(ctx context.Context, channelType, channelD string) Channels {
	return NewChannel(channelType, channelD, c)
}
