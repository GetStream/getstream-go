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
func (c *ChatClient) QueryChannels(ctx context.Context, request *QueryChannelsRequest) (*StreamResponse[QueryChannelsResponse], error) {
	var result QueryChannelsResponse
	res, err := MakeRequest[QueryChannelsRequest, QueryChannelsResponse](c.client, ctx, "POST", "/api/v2/chat/channels", nil, request, &result, nil)
	return res, err
}

// Allows to delete several channels at once asynchronously
//
// Sends events:
// - channel.deleted
//
// Required permissions:
// - DeleteChannel
func (c *ChatClient) DeleteChannels(ctx context.Context, request *DeleteChannelsRequest) (*StreamResponse[DeleteChannelsResponse], error) {
	var result DeleteChannelsResponse
	res, err := MakeRequest[DeleteChannelsRequest, DeleteChannelsResponse](c.client, ctx, "POST", "/api/v2/chat/channels/delete", nil, request, &result, nil)
	return res, err
}

// Marks channels as read up to the specific message. If no channels is given, mark all channel as read
//
// Sends events:
// - message.read
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) MarkChannelsRead(ctx context.Context, request *MarkChannelsReadRequest) (*StreamResponse[MarkReadResponse], error) {
	var result MarkReadResponse
	res, err := MakeRequest[MarkChannelsReadRequest, MarkReadResponse](c.client, ctx, "POST", "/api/v2/chat/channels/read", nil, request, &result, nil)
	return res, err
}

// This Method creates a channel or returns an existing one with matching attributes
//
// Sends events:
// - channel.created
// - member.added
// - member.removed
// - member.updated
// - user.watching.start
func (c *ChatClient) GetOrCreateDistinctChannel(ctx context.Context, _type string, request *ChannelGetOrCreateRequest) (*StreamResponse[ChannelStateResponse], error) {
	var result ChannelStateResponse
	pathParams := map[string]string{
		"type": _type,
	}
	res, err := MakeRequest[ChannelGetOrCreateRequest, ChannelStateResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/query", nil, request, &result, pathParams)
	return res, err
}

// Deletes channel
//
// Sends events:
// - channel.deleted
//
// Required permissions:
// - DeleteChannel
func (c *ChatClient) DeleteChannel(ctx context.Context, _type string, id string, request *DeleteChannelRequest) (*StreamResponse[DeleteChannelResponse], error) {
	var result DeleteChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, DeleteChannelResponse](c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}", params, nil, &result, pathParams)
	return res, err
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
func (c *ChatClient) UpdateChannelPartial(ctx context.Context, _type string, id string, request *UpdateChannelPartialRequest) (*StreamResponse[UpdateChannelPartialResponse], error) {
	var result UpdateChannelPartialResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[UpdateChannelPartialRequest, UpdateChannelPartialResponse](c.client, ctx, "PATCH", "/api/v2/chat/channels/{type}/{id}", nil, request, &result, pathParams)
	return res, err
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
func (c *ChatClient) UpdateChannel(ctx context.Context, _type string, id string, request *UpdateChannelRequest) (*StreamResponse[UpdateChannelResponse], error) {
	var result UpdateChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[UpdateChannelRequest, UpdateChannelResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}", nil, request, &result, pathParams)
	return res, err
}

// Sends event to the channel
//
// Required permissions:
// - SendCustomEvent
func (c *ChatClient) SendEvent(ctx context.Context, _type string, id string, request *SendEventRequest) (*StreamResponse[EventResponse], error) {
	var result EventResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[SendEventRequest, EventResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/event", nil, request, &result, pathParams)
	return res, err
}

// Deletes previously uploaded file
//
// Required permissions:
// - DeleteAttachment
func (c *ChatClient) DeleteFile(ctx context.Context, _type string, id string, request *DeleteFileRequest) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}/file", params, nil, &result, pathParams)
	return res, err
}

// Uploads file
//
// Required permissions:
// - UploadAttachment
func (c *ChatClient) UploadFile(ctx context.Context, _type string, id string, request *FileUploadRequest) (*StreamResponse[FileUploadResponse], error) {
	var result FileUploadResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[FileUploadRequest, FileUploadResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/file", nil, request, &result, pathParams)
	return res, err
}

// Marks channel as hidden for current user
//
// Sends events:
// - channel.hidden
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) HideChannel(ctx context.Context, _type string, id string, request *HideChannelRequest) (*StreamResponse[HideChannelResponse], error) {
	var result HideChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[HideChannelRequest, HideChannelResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/hide", nil, request, &result, pathParams)
	return res, err
}

// Deletes previously uploaded image
//
// Required permissions:
// - DeleteAttachment
func (c *ChatClient) DeleteImage(ctx context.Context, _type string, id string, request *DeleteImageRequest) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}/image", params, nil, &result, pathParams)
	return res, err
}

// Uploads image
//
// Required permissions:
// - UploadAttachment
func (c *ChatClient) UploadImage(ctx context.Context, _type string, id string, request *ImageUploadRequest) (*StreamResponse[ImageUploadResponse], error) {
	var result ImageUploadResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[ImageUploadRequest, ImageUploadResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/image", nil, request, &result, pathParams)
	return res, err
}

func (c *ChatClient) UpdateMemberPartial(ctx context.Context, userId string, _type string, id string, request *UpdateMemberPartialRequest) (*StreamResponse[UpdateMemberPartialResponse], error) {
	var result UpdateMemberPartialResponse
	pathParams := map[string]string{
		"user_id": userId,
		"type":    _type,
		"id":      id,
	}
	res, err := MakeRequest[UpdateMemberPartialRequest, UpdateMemberPartialResponse](c.client, ctx, "PATCH", "/api/v2/chat/channels/{type}/{id}/member/{user_id}", nil, request, &result, pathParams)
	return res, err
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
func (c *ChatClient) SendMessage(ctx context.Context, _type string, id string, request *SendMessageRequest) (*StreamResponse[SendMessageResponse], error) {
	var result SendMessageResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[SendMessageRequest, SendMessageResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/message", nil, request, &result, pathParams)
	return res, err
}

// Returns list messages found by IDs
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetManyMessages(ctx context.Context, _type string, id string, request *GetManyMessagesRequest) (*StreamResponse[GetManyMessagesResponse], error) {
	var result GetManyMessagesResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetManyMessagesResponse](c.client, ctx, "GET", "/api/v2/chat/channels/{type}/{id}/messages", params, nil, &result, pathParams)
	return res, err
}

// This Method creates a channel or returns an existing one with matching attributes
//
// Sends events:
// - channel.created
// - member.added
// - member.removed
// - member.updated
// - user.watching.start
func (c *ChatClient) GetOrCreateChannel(ctx context.Context, _type string, id string, request *ChannelGetOrCreateRequest) (*StreamResponse[ChannelStateResponse], error) {
	var result ChannelStateResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[ChannelGetOrCreateRequest, ChannelStateResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/query", nil, request, &result, pathParams)
	return res, err
}

// Marks channel as read up to the specific message
//
// Sends events:
// - message.read
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) MarkRead(ctx context.Context, _type string, id string, request *MarkReadRequest) (*StreamResponse[MarkReadResponse], error) {
	var result MarkReadResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[MarkReadRequest, MarkReadResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/read", nil, request, &result, pathParams)
	return res, err
}

// Shows previously hidden channel
//
// Sends events:
// - channel.visible
func (c *ChatClient) ShowChannel(ctx context.Context, _type string, id string, request *ShowChannelRequest) (*StreamResponse[ShowChannelResponse], error) {
	var result ShowChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[ShowChannelRequest, ShowChannelResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/show", nil, request, &result, pathParams)
	return res, err
}

// Truncates channel
//
// Sends events:
// - channel.truncated
//
// Required permissions:
// - DeleteChannel
// - TruncateChannel
func (c *ChatClient) TruncateChannel(ctx context.Context, _type string, id string, request *TruncateChannelRequest) (*StreamResponse[TruncateChannelResponse], error) {
	var result TruncateChannelResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[TruncateChannelRequest, TruncateChannelResponse](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/truncate", nil, request, &result, pathParams)
	return res, err
}

// Marks channel as unread from a specific message
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) MarkUnread(ctx context.Context, _type string, id string, request *MarkUnreadRequest) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[MarkUnreadRequest, Response](c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/unread", nil, request, &result, pathParams)
	return res, err
}

// Lists all available channel types
func (c *ChatClient) ListChannelTypes(ctx context.Context) (*StreamResponse[ListChannelTypesResponse], error) {
	var result ListChannelTypesResponse
	res, err := MakeRequest[any, ListChannelTypesResponse](c.client, ctx, "GET", "/api/v2/chat/channeltypes", nil, nil, &result, nil)
	return res, err
}

// Creates new channel type
func (c *ChatClient) CreateChannelType(ctx context.Context, request *CreateChannelTypeRequest) (*StreamResponse[CreateChannelTypeResponse], error) {
	var result CreateChannelTypeResponse
	res, err := MakeRequest[CreateChannelTypeRequest, CreateChannelTypeResponse](c.client, ctx, "POST", "/api/v2/chat/channeltypes", nil, request, &result, nil)
	return res, err
}

// Deletes channel type
func (c *ChatClient) DeleteChannelType(ctx context.Context, name string) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/channeltypes/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Gets channel type
func (c *ChatClient) GetChannelType(ctx context.Context, name string) (*StreamResponse[GetChannelTypeResponse], error) {
	var result GetChannelTypeResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, GetChannelTypeResponse](c.client, ctx, "GET", "/api/v2/chat/channeltypes/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Updates channel type
func (c *ChatClient) UpdateChannelType(ctx context.Context, name string, request *UpdateChannelTypeRequest) (*StreamResponse[UpdateChannelTypeResponse], error) {
	var result UpdateChannelTypeResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[UpdateChannelTypeRequest, UpdateChannelTypeResponse](c.client, ctx, "PUT", "/api/v2/chat/channeltypes/{name}", nil, request, &result, pathParams)
	return res, err
}

// Returns all custom commands
func (c *ChatClient) ListCommands(ctx context.Context) (*StreamResponse[ListCommandsResponse], error) {
	var result ListCommandsResponse
	res, err := MakeRequest[any, ListCommandsResponse](c.client, ctx, "GET", "/api/v2/chat/commands", nil, nil, &result, nil)
	return res, err
}

// Creates custom chat command
func (c *ChatClient) CreateCommand(ctx context.Context, request *CreateCommandRequest) (*StreamResponse[CreateCommandResponse], error) {
	var result CreateCommandResponse
	res, err := MakeRequest[CreateCommandRequest, CreateCommandResponse](c.client, ctx, "POST", "/api/v2/chat/commands", nil, request, &result, nil)
	return res, err
}

// Deletes custom chat command
func (c *ChatClient) DeleteCommand(ctx context.Context, name string) (*StreamResponse[DeleteCommandResponse], error) {
	var result DeleteCommandResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, DeleteCommandResponse](c.client, ctx, "DELETE", "/api/v2/chat/commands/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Returns custom command by its name
func (c *ChatClient) GetCommand(ctx context.Context, name string) (*StreamResponse[GetCommandResponse], error) {
	var result GetCommandResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, GetCommandResponse](c.client, ctx, "GET", "/api/v2/chat/commands/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Updates custom chat command
func (c *ChatClient) UpdateCommand(ctx context.Context, name string, request *UpdateCommandRequest) (*StreamResponse[UpdateCommandResponse], error) {
	var result UpdateCommandResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[UpdateCommandRequest, UpdateCommandResponse](c.client, ctx, "PUT", "/api/v2/chat/commands/{name}", nil, request, &result, pathParams)
	return res, err
}

// Exports channel data to JSON file
func (c *ChatClient) ExportChannels(ctx context.Context, request *ExportChannelsRequest) (*StreamResponse[ExportChannelsResponse], error) {
	var result ExportChannelsResponse
	res, err := MakeRequest[ExportChannelsRequest, ExportChannelsResponse](c.client, ctx, "POST", "/api/v2/chat/export_channels", nil, request, &result, nil)
	return res, err
}

func (c *ChatClient) GetExportChannelsStatus(ctx context.Context, id string) (*StreamResponse[GetExportChannelsStatusResponse], error) {
	var result GetExportChannelsStatusResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[any, GetExportChannelsStatusResponse](c.client, ctx, "GET", "/api/v2/chat/export_channels/{id}", nil, nil, &result, pathParams)
	return res, err
}

// Find and filter channel members
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) QueryMembers(ctx context.Context, request *QueryMembersRequest) (*StreamResponse[MembersResponse], error) {
	var result MembersResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, MembersResponse](c.client, ctx, "GET", "/api/v2/chat/members", params, nil, &result, nil)
	return res, err
}

// Queries history for one message
func (c *ChatClient) QueryMessageHistory(ctx context.Context, request *QueryMessageHistoryRequest) (*StreamResponse[QueryMessageHistoryResponse], error) {
	var result QueryMessageHistoryResponse
	res, err := MakeRequest[QueryMessageHistoryRequest, QueryMessageHistoryResponse](c.client, ctx, "POST", "/api/v2/chat/messages/history", nil, request, &result, nil)
	return res, err
}

// Deletes message
//
// Sends events:
// - message.deleted
//
// Required permissions:
// - DeleteMessage
func (c *ChatClient) DeleteMessage(ctx context.Context, id string, request *DeleteMessageRequest) (*StreamResponse[DeleteMessageResponse], error) {
	var result DeleteMessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, DeleteMessageResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{id}", params, nil, &result, pathParams)
	return res, err
}

// Returns message by ID
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetMessage(ctx context.Context, id string, request *GetMessageRequest) (*StreamResponse[GetMessageResponse], error) {
	var result GetMessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetMessageResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{id}", params, nil, &result, pathParams)
	return res, err
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
func (c *ChatClient) UpdateMessage(ctx context.Context, id string, request *UpdateMessageRequest) (*StreamResponse[UpdateMessageResponse], error) {
	var result UpdateMessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[UpdateMessageRequest, UpdateMessageResponse](c.client, ctx, "POST", "/api/v2/chat/messages/{id}", nil, request, &result, pathParams)
	return res, err
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
func (c *ChatClient) UpdateMessagePartial(ctx context.Context, id string, request *UpdateMessagePartialRequest) (*StreamResponse[UpdateMessagePartialResponse], error) {
	var result UpdateMessagePartialResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[UpdateMessagePartialRequest, UpdateMessagePartialResponse](c.client, ctx, "PUT", "/api/v2/chat/messages/{id}", nil, request, &result, pathParams)
	return res, err
}

// Executes message command action with given parameters
//
// Sends events:
// - message.new
//
// Required permissions:
// - RunMessageAction
func (c *ChatClient) RunMessageAction(ctx context.Context, id string, request *MessageActionRequest) (*StreamResponse[MessageResponse], error) {
	var result MessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[MessageActionRequest, MessageResponse](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/action", nil, request, &result, pathParams)
	return res, err
}

// Commits a pending message, which will make it visible in the channel
//
// Sends events:
// - message.new
// - message.updated
func (c *ChatClient) CommitMessage(ctx context.Context, id string, request *CommitMessageRequest) (*StreamResponse[MessageResponse], error) {
	var result MessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[CommitMessageRequest, MessageResponse](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/commit", nil, request, &result, pathParams)
	return res, err
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
func (c *ChatClient) SendReaction(ctx context.Context, id string, request *SendReactionRequest) (*StreamResponse[SendReactionResponse], error) {
	var result SendReactionResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[SendReactionRequest, SendReactionResponse](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/reaction", nil, request, &result, pathParams)
	return res, err
}

// Removes user reaction from the message
//
// Sends events:
// - reaction.deleted
//
// Required permissions:
// - DeleteReaction
func (c *ChatClient) DeleteReaction(ctx context.Context, id string, _type string, request *DeleteReactionRequest) (*StreamResponse[ReactionRemovalResponse], error) {
	var result ReactionRemovalResponse
	pathParams := map[string]string{
		"id":   id,
		"type": _type,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, ReactionRemovalResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{id}/reaction/{type}", params, nil, &result, pathParams)
	return res, err
}

// Returns list of reactions of specific message
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetReactions(ctx context.Context, id string, request *GetReactionsRequest) (*StreamResponse[GetReactionsResponse], error) {
	var result GetReactionsResponse
	pathParams := map[string]string{
		"id": id,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetReactionsResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{id}/reactions", params, nil, &result, pathParams)
	return res, err
}

// Get reactions on a message
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) QueryReactions(ctx context.Context, id string, request *QueryReactionsRequest) (*StreamResponse[QueryReactionsResponse], error) {
	var result QueryReactionsResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[QueryReactionsRequest, QueryReactionsResponse](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/reactions", nil, request, &result, pathParams)
	return res, err
}

// Translates message to a given language using automated translation software
//
// Sends events:
// - message.updated
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) TranslateMessage(ctx context.Context, id string, request *TranslateMessageRequest) (*StreamResponse[MessageResponse], error) {
	var result MessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[TranslateMessageRequest, MessageResponse](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/translate", nil, request, &result, pathParams)
	return res, err
}

// Undelete a message that was previously soft-deleted
//
// Sends events:
// - message.undeleted
func (c *ChatClient) UndeleteMessage(ctx context.Context, id string, request *UpdateMessageRequest) (*StreamResponse[UpdateMessageResponse], error) {
	var result UpdateMessageResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[UpdateMessageRequest, UpdateMessageResponse](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/undelete", nil, request, &result, pathParams)
	return res, err
}

// Cast a vote on a poll
//
// Sends events:
// - poll.vote_casted
//
// Required permissions:
// - CastVote
func (c *ChatClient) CastPollVote(ctx context.Context, messageId string, pollId string, request *CastPollVoteRequest) (*StreamResponse[PollVoteResponse], error) {
	var result PollVoteResponse
	pathParams := map[string]string{
		"message_id": messageId,
		"poll_id":    pollId,
	}
	res, err := MakeRequest[CastPollVoteRequest, PollVoteResponse](c.client, ctx, "POST", "/api/v2/chat/messages/{message_id}/polls/{poll_id}/vote", nil, request, &result, pathParams)
	return res, err
}

// Delete a vote from a poll
//
// Sends events:
// - poll.vote_removed
//
// Required permissions:
// - CastVote
func (c *ChatClient) RemovePollVote(ctx context.Context, messageId string, pollId string, voteId string, request *RemovePollVoteRequest) (*StreamResponse[PollVoteResponse], error) {
	var result PollVoteResponse
	pathParams := map[string]string{
		"message_id": messageId,
		"poll_id":    pollId,
		"vote_id":    voteId,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, PollVoteResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{message_id}/polls/{poll_id}/vote/{vote_id}", params, nil, &result, pathParams)
	return res, err
}

// Returns replies (thread) of the message
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetReplies(ctx context.Context, parentId string, request *GetRepliesRequest) (*StreamResponse[GetRepliesResponse], error) {
	var result GetRepliesResponse
	pathParams := map[string]string{
		"parent_id": parentId,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetRepliesResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{parent_id}/replies", params, nil, &result, pathParams)
	return res, err
}

// Find and filter message flags
//
// Required permissions:
// - ReadMessageFlags
func (c *ChatClient) QueryMessageFlags(ctx context.Context, request *QueryMessageFlagsRequest) (*StreamResponse[QueryMessageFlagsResponse], error) {
	var result QueryMessageFlagsResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, QueryMessageFlagsResponse](c.client, ctx, "GET", "/api/v2/chat/moderation/flags/message", params, nil, &result, nil)
	return res, err
}

// Mutes channel for user
//
// Sends events:
// - channel.muted
//
// Required permissions:
// - MuteChannel
func (c *ChatClient) MuteChannel(ctx context.Context, request *MuteChannelRequest) (*StreamResponse[MuteChannelResponse], error) {
	var result MuteChannelResponse
	res, err := MakeRequest[MuteChannelRequest, MuteChannelResponse](c.client, ctx, "POST", "/api/v2/chat/moderation/mute/channel", nil, request, &result, nil)
	return res, err
}

// Unmutes channel for user
//
// Sends events:
// - channel.unmuted
//
// Required permissions:
// - MuteChannel
func (c *ChatClient) UnmuteChannel(ctx context.Context, request *UnmuteChannelRequest) (*StreamResponse[UnmuteResponse], error) {
	var result UnmuteResponse
	res, err := MakeRequest[UnmuteChannelRequest, UnmuteResponse](c.client, ctx, "POST", "/api/v2/chat/moderation/unmute/channel", nil, request, &result, nil)
	return res, err
}

// Creates a new poll
//
// Required permissions:
// - CreatePoll
func (c *ChatClient) CreatePoll(ctx context.Context, request *CreatePollRequest) (*StreamResponse[PollResponse], error) {
	var result PollResponse
	res, err := MakeRequest[CreatePollRequest, PollResponse](c.client, ctx, "POST", "/api/v2/chat/polls", nil, request, &result, nil)
	return res, err
}

// Updates a poll
//
// Sends events:
// - poll.closed
// - poll.updated
//
// Required permissions:
// - UpdatePoll
func (c *ChatClient) UpdatePoll(ctx context.Context, request *UpdatePollRequest) (*StreamResponse[PollResponse], error) {
	var result PollResponse
	res, err := MakeRequest[UpdatePollRequest, PollResponse](c.client, ctx, "PUT", "/api/v2/chat/polls", nil, request, &result, nil)
	return res, err
}

// Queries polls
func (c *ChatClient) QueryPolls(ctx context.Context, request *QueryPollsRequest) (*StreamResponse[QueryPollsResponse], error) {
	var result QueryPollsResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[QueryPollsRequest, QueryPollsResponse](c.client, ctx, "POST", "/api/v2/chat/polls/query", params, request, &result, nil)
	return res, err
}

// Deletes a poll
//
// Sends events:
// - poll.deleted
//
// Required permissions:
// - DeletePoll
func (c *ChatClient) DeletePoll(ctx context.Context, pollId string, request *DeletePollRequest) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/polls/{poll_id}", params, nil, &result, pathParams)
	return res, err
}

// Retrieves a poll
func (c *ChatClient) GetPoll(ctx context.Context, pollId string, request *GetPollRequest) (*StreamResponse[PollResponse], error) {
	var result PollResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, PollResponse](c.client, ctx, "GET", "/api/v2/chat/polls/{poll_id}", params, nil, &result, pathParams)
	return res, err
}

// Updates a poll partially
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - UpdatePoll
func (c *ChatClient) UpdatePollPartial(ctx context.Context, pollId string, request *UpdatePollPartialRequest) (*StreamResponse[PollResponse], error) {
	var result PollResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	res, err := MakeRequest[UpdatePollPartialRequest, PollResponse](c.client, ctx, "PATCH", "/api/v2/chat/polls/{poll_id}", nil, request, &result, pathParams)
	return res, err
}

// Creates a poll option
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - CastVote
// - UpdatePoll
func (c *ChatClient) CreatePollOption(ctx context.Context, pollId string, request *CreatePollOptionRequest) (*StreamResponse[PollOptionResponse], error) {
	var result PollOptionResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	res, err := MakeRequest[CreatePollOptionRequest, PollOptionResponse](c.client, ctx, "POST", "/api/v2/chat/polls/{poll_id}/options", nil, request, &result, pathParams)
	return res, err
}

// Updates a poll option
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - UpdatePoll
func (c *ChatClient) UpdatePollOption(ctx context.Context, pollId string, request *UpdatePollOptionRequest) (*StreamResponse[PollOptionResponse], error) {
	var result PollOptionResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	res, err := MakeRequest[UpdatePollOptionRequest, PollOptionResponse](c.client, ctx, "PUT", "/api/v2/chat/polls/{poll_id}/options", nil, request, &result, pathParams)
	return res, err
}

// Deletes a poll option
//
// Sends events:
// - poll.updated
//
// Required permissions:
// - UpdatePoll
func (c *ChatClient) DeletePollOption(ctx context.Context, pollId string, optionId string, request *DeletePollOptionRequest) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"poll_id":   pollId,
		"option_id": optionId,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/polls/{poll_id}/options/{option_id}", params, nil, &result, pathParams)
	return res, err
}

// Retrieves a poll option
func (c *ChatClient) GetPollOption(ctx context.Context, pollId string, optionId string, request *GetPollOptionRequest) (*StreamResponse[PollOptionResponse], error) {
	var result PollOptionResponse
	pathParams := map[string]string{
		"poll_id":   pollId,
		"option_id": optionId,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, PollOptionResponse](c.client, ctx, "GET", "/api/v2/chat/polls/{poll_id}/options/{option_id}", params, nil, &result, pathParams)
	return res, err
}

// Queries votes
func (c *ChatClient) QueryPollVotes(ctx context.Context, pollId string, request *QueryPollVotesRequest) (*StreamResponse[PollVotesResponse], error) {
	var result PollVotesResponse
	pathParams := map[string]string{
		"poll_id": pollId,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[QueryPollVotesRequest, PollVotesResponse](c.client, ctx, "POST", "/api/v2/chat/polls/{poll_id}/votes", params, request, &result, pathParams)
	return res, err
}

// Search messages across channels
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) Search(ctx context.Context, request *SearchRequest) (*StreamResponse[SearchResponse], error) {
	var result SearchResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, SearchResponse](c.client, ctx, "GET", "/api/v2/chat/search", params, nil, &result, nil)
	return res, err
}

// Returns the list of threads for specific user
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) QueryThreads(ctx context.Context, request *QueryThreadsRequest) (*StreamResponse[QueryThreadsResponse], error) {
	var result QueryThreadsResponse
	res, err := MakeRequest[QueryThreadsRequest, QueryThreadsResponse](c.client, ctx, "POST", "/api/v2/chat/threads", nil, request, &result, nil)
	return res, err
}

// Return a specific thread
//
// Required permissions:
// - ReadChannel
func (c *ChatClient) GetThread(ctx context.Context, messageId string, request *GetThreadRequest) (*StreamResponse[GetThreadResponse], error) {
	var result GetThreadResponse
	pathParams := map[string]string{
		"message_id": messageId,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetThreadResponse](c.client, ctx, "GET", "/api/v2/chat/threads/{message_id}", params, nil, &result, pathParams)
	return res, err
}

// Updates certain fields of the thread
//
// Sends events:
// - thread.updated
//
// Required permissions:
// - ReadChannel
// - UpdateThread
func (c *ChatClient) UpdateThreadPartial(ctx context.Context, messageId string, request *UpdateThreadPartialRequest) (*StreamResponse[UpdateThreadPartialResponse], error) {
	var result UpdateThreadPartialResponse
	pathParams := map[string]string{
		"message_id": messageId,
	}
	res, err := MakeRequest[UpdateThreadPartialRequest, UpdateThreadPartialResponse](c.client, ctx, "PATCH", "/api/v2/chat/threads/{message_id}", nil, request, &result, pathParams)
	return res, err
}

// Fetch unread counts for a single user
func (c *ChatClient) UnreadCounts(ctx context.Context) (*StreamResponse[WrappedUnreadCountsResponse], error) {
	var result WrappedUnreadCountsResponse
	res, err := MakeRequest[any, WrappedUnreadCountsResponse](c.client, ctx, "GET", "/api/v2/chat/unread", nil, nil, &result, nil)
	return res, err
}

// Fetch unread counts in batch for multiple users in one call
func (c *ChatClient) UnreadCountsBatch(ctx context.Context, request *UnreadCountsBatchRequest) (*StreamResponse[UnreadCountsBatchResponse], error) {
	var result UnreadCountsBatchResponse
	res, err := MakeRequest[UnreadCountsBatchRequest, UnreadCountsBatchResponse](c.client, ctx, "POST", "/api/v2/chat/unread_batch", nil, request, &result, nil)
	return res, err
}

// Get list of blocked Users
func (c *ChatClient) GetBlockedUsers(ctx context.Context, request *GetBlockedUsersRequest) (*StreamResponse[GetBlockedUsersResponse], error) {
	var result GetBlockedUsersResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetBlockedUsersResponse](c.client, ctx, "GET", "/api/v2/chat/users/block", params, nil, &result, nil)
	return res, err
}

// Block users
func (c *ChatClient) BlockUsers(ctx context.Context, request *BlockUsersRequest) (*StreamResponse[BlockUsersResponse], error) {
	var result BlockUsersResponse
	res, err := MakeRequest[BlockUsersRequest, BlockUsersResponse](c.client, ctx, "POST", "/api/v2/chat/users/block", nil, request, &result, nil)
	return res, err
}

// Unblock users
func (c *ChatClient) UnblockUsers(ctx context.Context, request *UnblockUsersRequest) (*StreamResponse[UnblockUsersResponse], error) {
	var result UnblockUsersResponse
	res, err := MakeRequest[UnblockUsersRequest, UnblockUsersResponse](c.client, ctx, "POST", "/api/v2/chat/users/unblock", nil, request, &result, nil)
	return res, err
}

// Sends a custom event to a user
//
// Sends events:
// - *
func (c *ChatClient) SendUserCustomEvent(ctx context.Context, userId string, request *SendUserCustomEventRequest) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"user_id": userId,
	}
	res, err := MakeRequest[SendUserCustomEventRequest, Response](c.client, ctx, "POST", "/api/v2/chat/users/{user_id}/event", nil, request, &result, pathParams)
	return res, err
}
