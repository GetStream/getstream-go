package getstream

import (
	"context"
)

type ChatClient struct {
	client *Client
}

func NewChatClient(client *Client) *ChatClient {
	return &ChatClient{
		client: client,
	}
}

// Query channels with filter query
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) QueryChannels(ctx context.Context, request *QueryChannelsRequest) (*QueryChannelsResponse, error) {
	var result QueryChannelsResponse
	err := MakeRequest[QueryChannelsRequest, QueryChannelsResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels", nil, request, &result, nil)
	return &result, err
}

// Allows to delete several channels at once asynchronously
//
// Sends events:
// - channel.deleted
//
// Required permissions:
// - DeleteChannel
func (c *ChatClient) DeleteChannels(ctx context.Context, request *DeleteChannelsRequest) (*DeleteChannelsResponse, error) {
	var result DeleteChannelsResponse
	err := MakeRequest[DeleteChannelsRequest, DeleteChannelsResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/delete", nil, request, &result, nil)
	return &result, err
}

// Marks channels as read up to the specific message. If no channels is given, mark all channel as read
//
// Sends events:
// - message.read
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) MarkChannelsRead(ctx context.Context, request *MarkChannelsReadRequest) (*MarkReadResponse, error) {
	var result MarkReadResponse
	err := MakeRequest[MarkChannelsReadRequest, MarkReadResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/read", nil, request, &result, nil)
	return &result, err
}

// This Method creates a channel or returns an existing one with matching attributes
//
// Sends events:
// - channel.created
// - member.added
// - member.removed
// - member.updated
// - user.watching.start
func (c *ChatClient) GetOrCreateDistinctChannel(ctx context.Context, _type string, request *ChannelGetOrCreateRequest) (*ChannelStateResponse, error) {
	var result ChannelStateResponse
	pathParams := map[string]string{
		"type": _type,
	}
	err := MakeRequest[ChannelGetOrCreateRequest, ChannelStateResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/query", nil, request, &result, pathParams)
	return &result, err
}

// Deletes channel
//
// Sends events:
// - channel.deleted
//
// Required permissions:
// - DeleteChannel
func (c *ChatClient) DeleteChannel(ctx context.Context, _type string, id string, hardDelete *bool) (*DeleteChannelResponse, error) {
	var result DeleteChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	queryParams := map[string]interface{}{
		"hard_delete": hardDelete,
	}
	err := MakeRequest[any, DeleteChannelResponse](c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Updates certain fields of the channel
//
// Sends events:
// - channel.updated
//
// Required permissions:
// - UpdateChannel
// - UpdateChannelCooldown
// - UpdateChannelFrozen
func (c *ChatClient) UpdateChannelPartial(ctx context.Context, _type string, id string, request *UpdateChannelPartialRequest) (*UpdateChannelPartialResponse, error) {
	var result UpdateChannelPartialResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[UpdateChannelPartialRequest, UpdateChannelPartialResponse, any](c.client, ctx, "PATCH", "/api/v2/chat/channels/{type}/{id}", nil, request, &result, pathParams)
	return &result, err
}

// Change channel data
//
// Sends events:
// - channel.updated
// - member.added
// - member.removed
// - member.updated
// - message.new
//
// Required permissions:
// - AddOwnChannelMembership
// - RemoveOwnChannelMembership
// - UpdateChannel
// - UpdateChannelCooldown
// - UpdateChannelFrozen
// - UpdateChannelMembers
func (c *ChatClient) UpdateChannel(ctx context.Context, _type string, id string, request *UpdateChannelRequest) (*UpdateChannelResponse, error) {
	var result UpdateChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[UpdateChannelRequest, UpdateChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}", nil, request, &result, pathParams)
	return &result, err
}

// Sends event to the channel
//
// Required permissions:
// - SendCustomEvent
func (c *ChatClient) SendEvent(ctx context.Context, _type string, id string, request *SendEventRequest) (*EventResponse, error) {
	var result EventResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[SendEventRequest, EventResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/event", nil, request, &result, pathParams)
	return &result, err
}

// Deletes previously uploaded file
//
// Required permissions:
// - DeleteAttachment
func (c *ChatClient) DeleteFile(ctx context.Context, _type string, id string, url *string) (*FileDeleteResponse, error) {
	var result FileDeleteResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	queryParams := map[string]interface{}{
		"url": url,
	}
	err := MakeRequest[any, FileDeleteResponse](c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}/file", queryParams, nil, &result, pathParams)
	return &result, err
}

// Uploads file
//
// Required permissions:
// - UploadAttachment
func (c *ChatClient) UploadFile(ctx context.Context, _type string, id string, request *FileUploadRequest) (*FileUploadResponse, error) {
	var result FileUploadResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[FileUploadRequest, FileUploadResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/file", nil, request, &result, pathParams)
	return &result, err
}

// Marks channel as hidden for current user
//
// Sends events:
// - channel.hidden
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) HideChannel(ctx context.Context, _type string, id string, request *HideChannelRequest) (*HideChannelResponse, error) {
	var result HideChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[HideChannelRequest, HideChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/hide", nil, request, &result, pathParams)
	return &result, err
}

// Deletes previously uploaded image
//
// Required permissions:
// - DeleteAttachment
func (c *ChatClient) DeleteImage(ctx context.Context, _type string, id string, url *string) (*FileDeleteResponse, error) {
	var result FileDeleteResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	queryParams := map[string]interface{}{
		"url": url,
	}
	err := MakeRequest[any, FileDeleteResponse](c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}/image", queryParams, nil, &result, pathParams)
	return &result, err
}

// Uploads image
//
// Required permissions:
// - UploadAttachment
func (c *ChatClient) UploadImage(ctx context.Context, _type string, id string, request *ImageUploadRequest) (*ImageUploadResponse, error) {
	var result ImageUploadResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[ImageUploadRequest, ImageUploadResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/image", nil, request, &result, pathParams)
	return &result, err
}

// Sends new message to the specified channel
//
// Sends events:
// - message.new
// - message.updated
//
// Required permissions:
// - AddLinks
// - CreateMessage
// - PinMessage
// - SkipChannelCooldown
// - SkipMessageModeration
// - UseFrozenChannel
func (c *ChatClient) SendMessage(ctx context.Context, _type string, id string, request *SendMessageRequest) (*SendMessageResponse, error) {
	var result SendMessageResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[SendMessageRequest, SendMessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/message", nil, request, &result, pathParams)
	return &result, err
}

// Returns list messages found by IDs
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetManyMessages(ctx context.Context, _type string, id string, ids []string) (*GetManyMessagesResponse, error) {
	var result GetManyMessagesResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	queryParams := map[string]interface{}{
		"ids": ids,
	}
	err := MakeRequest[any, GetManyMessagesResponse](c.client, ctx, "GET", "/api/v2/chat/channels/{type}/{id}/messages", queryParams, nil, &result, pathParams)
	return &result, err
}

// This Method creates a channel or returns an existing one with matching attributes
//
// Sends events:
// - channel.created
// - member.added
// - member.removed
// - member.updated
// - user.watching.start
func (c *ChatClient) GetOrCreateChannel(ctx context.Context, _type string, id string, request *ChannelGetOrCreateRequest) (*ChannelStateResponse, error) {
	var result ChannelStateResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[ChannelGetOrCreateRequest, ChannelStateResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/query", nil, request, &result, pathParams)
	return &result, err
}

// Marks channel as read up to the specific message
//
// Sends events:
// - message.read
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) MarkRead(ctx context.Context, _type string, id string, request *MarkReadRequest) (*MarkReadResponse, error) {
	var result MarkReadResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[MarkReadRequest, MarkReadResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/read", nil, request, &result, pathParams)
	return &result, err
}

// Shows previously hidden channel
//
// Sends events:
// - channel.visible
func (c *ChatClient) ShowChannel(ctx context.Context, _type string, id string, request *ShowChannelRequest) (*ShowChannelResponse, error) {
	var result ShowChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[ShowChannelRequest, ShowChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/show", nil, request, &result, pathParams)
	return &result, err
}

// Truncates channel
//
// Sends events:
// - channel.truncated
//
// Required permissions:
// - DeleteChannel
// - TruncateChannel
func (c *ChatClient) TruncateChannel(ctx context.Context, _type string, id string, request *TruncateChannelRequest) (*TruncateChannelResponse, error) {
	var result TruncateChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[TruncateChannelRequest, TruncateChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/truncate", nil, request, &result, pathParams)
	return &result, err
}

// Marks channel as unread from a specific message
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) MarkUnread(ctx context.Context, _type string, id string, request *MarkUnreadRequest) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[MarkUnreadRequest, Response, any](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/unread", nil, request, &result, pathParams)
	return &result, err
}

// Lists all available channel types
func (c *ChatClient) ListChannelTypes(ctx context.Context) (*ListChannelTypesResponse, error) {
	var result ListChannelTypesResponse
	err := MakeRequest[any, ListChannelTypesResponse, any](c.client, ctx, "GET", "/api/v2/chat/channeltypes", nil, nil, &result, nil)
	return &result, err
}

// Creates new channel type
func (c *ChatClient) CreateChannelType(ctx context.Context, request *CreateChannelTypeRequest) (*CreateChannelTypeResponse, error) {
	var result CreateChannelTypeResponse
	err := MakeRequest[CreateChannelTypeRequest, CreateChannelTypeResponse, any](c.client, ctx, "POST", "/api/v2/chat/channeltypes", nil, request, &result, nil)
	return &result, err
}

// Deletes channel type
func (c *ChatClient) DeleteChannelType(ctx context.Context, name string) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/chat/channeltypes/{name}", nil, nil, &result, pathParams)
	return &result, err
}

// Gets channel type
func (c *ChatClient) GetChannelType(ctx context.Context, name string) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "GET", "/api/v2/chat/channeltypes/{name}", nil, nil, &result, pathParams)
	return &result, err
}

// Updates channel type
func (c *ChatClient) UpdateChannelType(ctx context.Context, name string, request *UpdateChannelTypeRequest) (*UpdateChannelTypeResponse, error) {
	var result UpdateChannelTypeResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[UpdateChannelTypeRequest, UpdateChannelTypeResponse, any](c.client, ctx, "PUT", "/api/v2/chat/channeltypes/{name}", nil, request, &result, pathParams)
	return &result, err
}

// Returns all custom commands
func (c *ChatClient) ListCommands(ctx context.Context) (*ListCommandsResponse, error) {
	var result ListCommandsResponse
	err := MakeRequest[any, ListCommandsResponse, any](c.client, ctx, "GET", "/api/v2/chat/commands", nil, nil, &result, nil)
	return &result, err
}

// Creates custom chat command
func (c *ChatClient) CreateCommand(ctx context.Context, request *CreateCommandRequest) (*CreateCommandResponse, error) {
	var result CreateCommandResponse
	err := MakeRequest[CreateCommandRequest, CreateCommandResponse, any](c.client, ctx, "POST", "/api/v2/chat/commands", nil, request, &result, nil)
	return &result, err
}

// Deletes custom chat command
func (c *ChatClient) DeleteCommand(ctx context.Context, name string) (*DeleteCommandResponse, error) {
	var result DeleteCommandResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, DeleteCommandResponse, any](c.client, ctx, "DELETE", "/api/v2/chat/commands/{name}", nil, nil, &result, pathParams)
	return &result, err
}

// Returns custom command by its name
func (c *ChatClient) GetCommand(ctx context.Context, name string) (*GetCommandResponse, error) {
	var result GetCommandResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, GetCommandResponse, any](c.client, ctx, "GET", "/api/v2/chat/commands/{name}", nil, nil, &result, pathParams)
	return &result, err
}

// Updates custom chat command
func (c *ChatClient) UpdateCommand(ctx context.Context, name string, request *UpdateCommandRequest) (*UpdateCommandResponse, error) {
	var result UpdateCommandResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[UpdateCommandRequest, UpdateCommandResponse, any](c.client, ctx, "PUT", "/api/v2/chat/commands/{name}", nil, request, &result, pathParams)
	return &result, err
}

// Exports channel data to JSON file
func (c *ChatClient) ExportChannels(ctx context.Context, request *ExportChannelsRequest) (*ExportChannelsResponse, error) {
	var result ExportChannelsResponse
	err := MakeRequest[ExportChannelsRequest, ExportChannelsResponse, any](c.client, ctx, "POST", "/api/v2/chat/export_channels", nil, request, &result, nil)
	return &result, err
}

func (c *ChatClient) GetExportChannelsStatus(ctx context.Context, id string) (*GetExportChannelsStatusResponse, error) {
	var result GetExportChannelsStatusResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[any, GetExportChannelsStatusResponse, any](c.client, ctx, "GET", "/api/v2/chat/export_channels/{id}", nil, nil, &result, pathParams)
	return &result, err
}

// Find and filter channel members
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) QueryMembers(ctx context.Context, payload *QueryMembersRequest) (*MembersResponse, error) {
	var result MembersResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, MembersResponse](c.client, ctx, "GET", "/api/v2/chat/members", queryParams, nil, &result, nil)
	return &result, err
}

// Queries history for one message
func (c *ChatClient) QueryMessageHistory(ctx context.Context, request *QueryMessageHistoryRequest) (*QueryMessageHistoryResponse, error) {
	var result QueryMessageHistoryResponse
	err := MakeRequest[QueryMessageHistoryRequest, QueryMessageHistoryResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/history", nil, request, &result, nil)
	return &result, err
}

// Deletes message
//
// Sends events:
// - message.deleted
//
// Required permissions:
// - DeleteMessage
func (c *ChatClient) DeleteMessage(ctx context.Context, id string, hard *bool, deletedBy *string) (*DeleteMessageResponse, error) {
	var result DeleteMessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	queryParams := map[string]interface{}{
		"hard":       hard,
		"deleted_by": deletedBy,
	}
	err := MakeRequest[any, DeleteMessageResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Returns message by ID
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetMessage(ctx context.Context, id string, showDeletedMessage *bool) (*GetMessageResponse, error) {
	var result GetMessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	queryParams := map[string]interface{}{
		"show_deleted_message": showDeletedMessage,
	}
	err := MakeRequest[any, GetMessageResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Updates message with new data
//
// Sends events:
// - message.updated
//
// Required permissions:
// - AddLinks
// - PinMessage
// - SkipMessageModeration
// - UpdateMessage
func (c *ChatClient) UpdateMessage(ctx context.Context, id string, request *UpdateMessageRequest) (*UpdateMessageResponse, error) {
	var result UpdateMessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[UpdateMessageRequest, UpdateMessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}", nil, request, &result, pathParams)
	return &result, err
}

// Updates certain fields of the message
//
// Sends events:
// - message.updated
//
// Required permissions:
// - AddLinks
// - PinMessage
// - SkipMessageModeration
// - UpdateMessage
func (c *ChatClient) UpdateMessagePartial(ctx context.Context, id string, request *UpdateMessagePartialRequest) (*UpdateMessagePartialResponse, error) {
	var result UpdateMessagePartialResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[UpdateMessagePartialRequest, UpdateMessagePartialResponse, any](c.client, ctx, "PUT", "/api/v2/chat/messages/{id}", nil, request, &result, pathParams)
	return &result, err
}

// Executes message command action with given parameters
//
// Sends events:
// - message.new
//
// Required permissions:
// - RunMessageAction
func (c *ChatClient) RunMessageAction(ctx context.Context, id string, request *MessageActionRequest) (*MessageResponse, error) {
	var result MessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[MessageActionRequest, MessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/action", nil, request, &result, pathParams)
	return &result, err
}

// Commits a pending message, which will make it visible in the channel
//
// Sends events:
// - message.new
// - message.updated
func (c *ChatClient) CommitMessage(ctx context.Context, id string, request *CommitMessageRequest) (*MessageResponse, error) {
	var result MessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[CommitMessageRequest, MessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/commit", nil, request, &result, pathParams)
	return &result, err
}

// Sends reaction to specified message
//
// Sends events:
// - reaction.new
// - reaction.updated
//
// Required permissions:
// - CreateReaction
// - UseFrozenChannel
func (c *ChatClient) SendReaction(ctx context.Context, id string, request *SendReactionRequest) (*SendReactionResponse, error) {
	var result SendReactionResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[SendReactionRequest, SendReactionResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/reaction", nil, request, &result, pathParams)
	return &result, err
}

// Removes user reaction from the message
//
// Sends events:
// - reaction.deleted
//
// Required permissions:
// - DeleteReaction
func (c *ChatClient) DeleteReaction(ctx context.Context, id string, _type string, userId *string) (*ReactionRemovalResponse, error) {
	var result ReactionRemovalResponse
	pathParams := map[string]string{
		"id":   id,
		"type": _type,
	}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, ReactionRemovalResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{id}/reaction/{type}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Returns list of reactions of specific message
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetReactions(ctx context.Context, id string, limit *int, offset *int) (*GetReactionsResponse, error) {
	var result GetReactionsResponse
	pathParams := map[string]string{
		"id": id,
	}
	queryParams := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}
	err := MakeRequest[any, GetReactionsResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{id}/reactions", queryParams, nil, &result, pathParams)
	return &result, err
}

// Get reactions on a message
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) QueryReactions(ctx context.Context, id string, request *QueryReactionsRequest) (*QueryReactionsResponse, error) {
	var result QueryReactionsResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[QueryReactionsRequest, QueryReactionsResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/reactions", nil, request, &result, pathParams)
	return &result, err
}

// Translates message to a given language using automated translation software
//
// Sends events:
// - message.updated
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) TranslateMessage(ctx context.Context, id string, request *TranslateMessageRequest) (*MessageResponse, error) {
	var result MessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[TranslateMessageRequest, MessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/translate", nil, request, &result, pathParams)
	return &result, err
}

// Undelete a message that was previously soft-deleted
//
// Sends events:
// - message.undeleted
func (c *ChatClient) UndeleteMessage(ctx context.Context, id string, request *UpdateMessageRequest) (*UpdateMessageResponse, error) {
	var result UpdateMessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[UpdateMessageRequest, UpdateMessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/undelete", nil, request, &result, pathParams)
	return &result, err
}

// Cast a vote on a poll
//
// Sends events:
// - poll.vote_casted
//
// Required permissions:
// - CastVote
func (c *ChatClient) CastPollVote(ctx context.Context, messageId string, pollId string, request *CastPollVoteRequest) (*PollVoteResponse, error) {
	var result PollVoteResponse
	pathParams := map[string]string{
		"message_id": messageId,
		"poll_id":    pollId,
	}
	err := MakeRequest[CastPollVoteRequest, PollVoteResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{message_id}/polls/{poll_id}/vote", nil, request, &result, pathParams)
	return &result, err
}

// Delete a vote from a poll
//
// Sends events:
// - poll.vote_removed
//
// Required permissions:
// - CastVote
func (c *ChatClient) RemovePollVote(ctx context.Context, messageId string, pollId string, voteId string, userId *string) (*PollVoteResponse, error) {
	var result PollVoteResponse
	pathParams := map[string]string{
		"message_id": messageId,
		"poll_id":    pollId,
		"vote_id":    voteId,
	}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, PollVoteResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{message_id}/polls/{poll_id}/vote/{vote_id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Returns replies (thread) of the message
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetReplies(ctx context.Context, parentId string, limit *int, offset *int, idGte *string, idGt *string, idLte *string, idLt *string, createdAtAfterOrEqual *Timestamp, createdAtAfter *Timestamp, createdAtBeforeOrEqual *Timestamp, createdAtBefore *Timestamp, idAround *string, createdAtAround *Timestamp, sort *[]*SortParam) (*GetRepliesResponse, error) {
	var result GetRepliesResponse
	pathParams := map[string]string{
		"parent_id": parentId,
	}
	queryParams := map[string]interface{}{
		"limit":                      limit,
		"offset":                     offset,
		"id_gte":                     idGte,
		"id_gt":                      idGt,
		"id_lte":                     idLte,
		"id_lt":                      idLt,
		"created_at_after_or_equal":  createdAtAfterOrEqual,
		"created_at_after":           createdAtAfter,
		"created_at_before_or_equal": createdAtBeforeOrEqual,
		"created_at_before":          createdAtBefore,
		"id_around":                  idAround,
		"created_at_around":          createdAtAround,
		"sort":                       sort,
	}
	err := MakeRequest[any, GetRepliesResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{parent_id}/replies", queryParams, nil, &result, pathParams)
	return &result, err
}

// Find and filter message flags
//
// Required permissions:
// - ReadMessageFlags
func (c *ChatClient) QueryMessageFlags(ctx context.Context, payload *QueryMessageFlagsRequest) (*QueryMessageFlagsResponse, error) {
	var result QueryMessageFlagsResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, QueryMessageFlagsResponse](c.client, ctx, "GET", "/api/v2/chat/moderation/flags/message", queryParams, nil, &result, nil)
	return &result, err
}

// Mutes channel for user
//
// Sends events:
// - channel.muted
//
// Required permissions:
// - MuteChannel
func (c *ChatClient) MuteChannel(ctx context.Context, request *MuteChannelRequest) (*MuteChannelResponse, error) {
	var result MuteChannelResponse
	err := MakeRequest[MuteChannelRequest, MuteChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/moderation/mute/channel", nil, request, &result, nil)
	return &result, err
}

// Unmutes channel for user
//
// Sends events:
// - channel.unmuted
//
// Required permissions:
// - MuteChannel
func (c *ChatClient) UnmuteChannel(ctx context.Context, request *UnmuteChannelRequest) (*UnmuteResponse, error) {
	var result UnmuteResponse
	err := MakeRequest[UnmuteChannelRequest, UnmuteResponse, any](c.client, ctx, "POST", "/api/v2/chat/moderation/unmute/channel", nil, request, &result, nil)
	return &result, err
}

// Creates a new poll
//
// Sends events:
// - poll.created
//
// Required permissions:
// - CreatePoll
func (c *ChatClient) CreatePoll(ctx context.Context, request *CreatePollRequest) (*PollResponse, error) {
	var result PollResponse
	err := MakeRequest[CreatePollRequest, PollResponse, any](c.client, ctx, "POST", "/api/v2/chat/polls", nil, request, &result, nil)
	return &result, err
}

// Updates a poll
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - UpdatePoll
func (c *ChatClient) UpdatePoll(ctx context.Context, request *UpdatePollRequest) (*PollResponse, error) {
	var result PollResponse
	err := MakeRequest[UpdatePollRequest, PollResponse, any](c.client, ctx, "PUT", "/api/v2/chat/polls", nil, request, &result, nil)
	return &result, err
}

// Queries polls
func (c *ChatClient) QueryPolls(ctx context.Context, userId *string, request *QueryPollsRequest) (*QueryPollsResponse, error) {
	var result QueryPollsResponse
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[QueryPollsRequest, QueryPollsResponse](c.client, ctx, "POST", "/api/v2/chat/polls/query", queryParams, request, &result, nil)
	return &result, err
}

// Deletes a poll
//
// Sends events:
// - poll.deleted
//
// Required permissions:
// - DeletePoll
func (c *ChatClient) DeletePoll(ctx context.Context, pollId string, userId *string) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/polls/{poll_id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Retrieves a poll
func (c *ChatClient) GetPoll(ctx context.Context, pollId string, userId *string) (*PollResponse, error) {
	var result PollResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, PollResponse](c.client, ctx, "GET", "/api/v2/chat/polls/{poll_id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Updates a poll partially
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - UpdatePoll
func (c *ChatClient) UpdatePollPartial(ctx context.Context, pollId string, request *UpdatePollPartialRequest) (*PollResponse, error) {
	var result PollResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	err := MakeRequest[UpdatePollPartialRequest, PollResponse, any](c.client, ctx, "PATCH", "/api/v2/chat/polls/{poll_id}", nil, request, &result, pathParams)
	return &result, err
}

// Creates a poll option
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - CastVote
// - UpdatePoll
func (c *ChatClient) CreatePollOption(ctx context.Context, pollId string, request *CreatePollOptionRequest) (*PollOptionResponse, error) {
	var result PollOptionResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	err := MakeRequest[CreatePollOptionRequest, PollOptionResponse, any](c.client, ctx, "POST", "/api/v2/chat/polls/{poll_id}/options", nil, request, &result, pathParams)
	return &result, err
}

// Updates a poll option
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - UpdatePoll
func (c *ChatClient) UpdatePollOption(ctx context.Context, pollId string, request *UpdatePollOptionRequest) (*PollOptionResponse, error) {
	var result PollOptionResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	err := MakeRequest[UpdatePollOptionRequest, PollOptionResponse, any](c.client, ctx, "PUT", "/api/v2/chat/polls/{poll_id}/options", nil, request, &result, pathParams)
	return &result, err
}

// Deletes a poll option
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - UpdatePoll
func (c *ChatClient) DeletePollOption(ctx context.Context, pollId string, optionId string, userId *string) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"poll_id":   pollId,
		"option_id": optionId,
	}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/polls/{poll_id}/options/{option_id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Retrieves a poll option
func (c *ChatClient) GetPollOption(ctx context.Context, pollId string, optionId string, userId *string) (*PollOptionResponse, error) {
	var result PollOptionResponse
	pathParams := map[string]string{
		"poll_id":   pollId,
		"option_id": optionId,
	}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, PollOptionResponse](c.client, ctx, "GET", "/api/v2/chat/polls/{poll_id}/options/{option_id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Queries votes
func (c *ChatClient) QueryPollVotes(ctx context.Context, pollId string, userId *string, request *QueryPollVotesRequest) (*PollVotesResponse, error) {
	var result PollVotesResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[QueryPollVotesRequest, PollVotesResponse](c.client, ctx, "POST", "/api/v2/chat/polls/{poll_id}/votes", queryParams, request, &result, pathParams)
	return &result, err
}

// Find and filter channel scoped or global user bans
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) QueryBannedUsers(ctx context.Context, payload *QueryBannedUsersRequest) (*QueryBannedUsersResponse, error) {
	var result QueryBannedUsersResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, QueryBannedUsersResponse](c.client, ctx, "GET", "/api/v2/chat/query_banned_users", queryParams, nil, &result, nil)
	return &result, err
}

// Search messages across channels
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) Search(ctx context.Context, payload *SearchRequest) (*SearchResponse, error) {
	var result SearchResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, SearchResponse](c.client, ctx, "GET", "/api/v2/chat/search", queryParams, nil, &result, nil)
	return &result, err
}

// Returns the list of threads for specific user
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) QueryThreads(ctx context.Context, request *QueryThreadsRequest) (*QueryThreadsResponse, error) {
	var result QueryThreadsResponse
	err := MakeRequest[QueryThreadsRequest, QueryThreadsResponse, any](c.client, ctx, "POST", "/api/v2/chat/threads", nil, request, &result, nil)
	return &result, err
}

// Return a specific thread
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetThread(ctx context.Context, messageId string, connectionId *string, replyLimit *int, participantLimit *int, memberLimit *int) (*GetThreadResponse, error) {
	var result GetThreadResponse
	pathParams := map[string]string{
		"message_id": messageId,
	}
	queryParams := map[string]interface{}{
		"connection_id":     connectionId,
		"reply_limit":       replyLimit,
		"participant_limit": participantLimit,
		"member_limit":      memberLimit,
	}
	err := MakeRequest[any, GetThreadResponse](c.client, ctx, "GET", "/api/v2/chat/threads/{message_id}", queryParams, nil, &result, pathParams)
	return &result, err
}

// Updates certain fields of the thread
//
// Sends events:
// - thread.updated
//
// Required permissions:
// - ReadChannel
// - UpdateThread
func (c *ChatClient) UpdateThreadPartial(ctx context.Context, messageId string, request *UpdateThreadPartialRequest) (*UpdateThreadPartialResponse, error) {
	var result UpdateThreadPartialResponse
	pathParams := map[string]string{
		"message_id": messageId,
	}
	err := MakeRequest[UpdateThreadPartialRequest, UpdateThreadPartialResponse, any](c.client, ctx, "PATCH", "/api/v2/chat/threads/{message_id}", nil, request, &result, pathParams)
	return &result, err
}

// Fetch unread counts for a single user
func (c *ChatClient) UnreadCounts(ctx context.Context) (*WrappedUnreadCountsResponse, error) {
	var result WrappedUnreadCountsResponse
	err := MakeRequest[any, WrappedUnreadCountsResponse, any](c.client, ctx, "GET", "/api/v2/chat/unread", nil, nil, &result, nil)
	return &result, err
}

// Fetch unread counts in batch for multiple users in one call
func (c *ChatClient) UnreadCountsBatch(ctx context.Context, request *UnreadCountsBatchRequest) (*UnreadCountsBatchResponse, error) {
	var result UnreadCountsBatchResponse
	err := MakeRequest[UnreadCountsBatchRequest, UnreadCountsBatchResponse, any](c.client, ctx, "POST", "/api/v2/chat/unread_batch", nil, request, &result, nil)
	return &result, err
}

// Sends a custom event to a user
//
// Sends events:
// - *
func (c *ChatClient) SendUserCustomEvent(ctx context.Context, userId string, request *SendUserCustomEventRequest) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"user_id": userId,
	}
	err := MakeRequest[SendUserCustomEventRequest, Response, any](c.client, ctx, "POST", "/api/v2/chat/users/{user_id}/event", nil, request, &result, pathParams)
	return &result, err
}
