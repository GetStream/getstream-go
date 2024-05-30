package getstream

import (
	"context"
	"time"
)

type ChatClient struct {
	client *Client
}

func NewChatClient(client *Client) *ChatClient {
	return &ChatClient{
		client: client,
	}
}
func (c *ChatClient) QueryChannels(ctx context.Context, queryChannelsRequest QueryChannelsRequest) (QueryChannelsResponse, error) {
	var result QueryChannelsResponse
	err := MakeRequest[QueryChannelsRequest, QueryChannelsResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels", nil, &queryChannelsRequest, &result)
	return result, err
}
func (c *ChatClient) DeleteChannels(ctx context.Context, deleteChannelsRequest DeleteChannelsRequest) (DeleteChannelsResponse, error) {
	var result DeleteChannelsResponse
	err := MakeRequest[DeleteChannelsRequest, DeleteChannelsResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/delete", nil, &deleteChannelsRequest, &result)
	return result, err
}
func (c *ChatClient) MarkChannelsRead(ctx context.Context, markChannelsReadRequest MarkChannelsReadRequest) (MarkReadResponse, error) {
	var result MarkReadResponse
	err := MakeRequest[MarkChannelsReadRequest, MarkReadResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/read", nil, &markChannelsReadRequest, &result)
	return result, err
}
func (c *ChatClient) GetOrCreateDistinctChannel(ctx context.Context, _type string, channelGetOrCreateRequest ChannelGetOrCreateRequest) (ChannelStateResponse, error) {
	var result ChannelStateResponse
	pathParams := []string{_type}
	err := MakeRequest[ChannelGetOrCreateRequest, ChannelStateResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/query", nil, &channelGetOrCreateRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) DeleteChannel(ctx context.Context, _type string, id string, hardDelete *bool) (DeleteChannelResponse, error) {
	var result DeleteChannelResponse
	pathParams := []string{_type, id}
	queryParams := map[string]interface{}{
		"hard_delete": hardDelete,
	}
	err := MakeRequest[any, DeleteChannelResponse](c.client, ctx, "DELETE", "/api/v2/chat/channels/{_type}/{id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdateChannelPartial(ctx context.Context, _type string, id string, updateChannelPartialRequest UpdateChannelPartialRequest) (UpdateChannelPartialResponse, error) {
	var result UpdateChannelPartialResponse
	pathParams := []string{_type, id}
	err := MakeRequest[UpdateChannelPartialRequest, UpdateChannelPartialResponse, any](c.client, ctx, "PATCH", "/api/v2/chat/channels/{_type}/{id}", nil, &updateChannelPartialRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdateChannel(ctx context.Context, _type string, id string, updateChannelRequest UpdateChannelRequest) (UpdateChannelResponse, error) {
	var result UpdateChannelResponse
	pathParams := []string{_type, id}
	err := MakeRequest[UpdateChannelRequest, UpdateChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}", nil, &updateChannelRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) SendEvent(ctx context.Context, _type string, id string, sendEventRequest SendEventRequest) (EventResponse, error) {
	var result EventResponse
	pathParams := []string{_type, id}
	err := MakeRequest[SendEventRequest, EventResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/event", nil, &sendEventRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) DeleteFile(ctx context.Context, _type string, id string, url *string) (FileDeleteResponse, error) {
	var result FileDeleteResponse
	pathParams := []string{_type, id}
	queryParams := map[string]interface{}{
		"url": url,
	}
	err := MakeRequest[any, FileDeleteResponse](c.client, ctx, "DELETE", "/api/v2/chat/channels/{_type}/{id}/file", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UploadFile(ctx context.Context, _type string, id string, fileUploadRequest FileUploadRequest) (FileUploadResponse, error) {
	var result FileUploadResponse
	pathParams := []string{_type, id}
	err := MakeRequest[FileUploadRequest, FileUploadResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/file", nil, &fileUploadRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) HideChannel(ctx context.Context, _type string, id string, hideChannelRequest HideChannelRequest) (HideChannelResponse, error) {
	var result HideChannelResponse
	pathParams := []string{_type, id}
	err := MakeRequest[HideChannelRequest, HideChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/hide", nil, &hideChannelRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) DeleteImage(ctx context.Context, _type string, id string, url *string) (FileDeleteResponse, error) {
	var result FileDeleteResponse
	pathParams := []string{_type, id}
	queryParams := map[string]interface{}{
		"url": url,
	}
	err := MakeRequest[any, FileDeleteResponse](c.client, ctx, "DELETE", "/api/v2/chat/channels/{_type}/{id}/image", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UploadImage(ctx context.Context, _type string, id string, imageUploadRequest ImageUploadRequest) (ImageUploadResponse, error) {
	var result ImageUploadResponse
	pathParams := []string{_type, id}
	err := MakeRequest[ImageUploadRequest, ImageUploadResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/image", nil, &imageUploadRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) SendMessage(ctx context.Context, _type string, id string, sendMessageRequest SendMessageRequest) (SendMessageResponse, error) {
	var result SendMessageResponse
	pathParams := []string{_type, id}
	err := MakeRequest[SendMessageRequest, SendMessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/message", nil, &sendMessageRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetManyMessages(ctx context.Context, _type string, id string, ids []string) (GetManyMessagesResponse, error) {
	var result GetManyMessagesResponse
	pathParams := []string{_type, id}
	queryParams := map[string]interface{}{
		"ids": ids,
	}
	err := MakeRequest[any, GetManyMessagesResponse](c.client, ctx, "GET", "/api/v2/chat/channels/{_type}/{id}/messages", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetOrCreateChannel(ctx context.Context, _type string, id string, channelGetOrCreateRequest ChannelGetOrCreateRequest) (ChannelStateResponse, error) {
	var result ChannelStateResponse
	pathParams := []string{_type, id}
	err := MakeRequest[ChannelGetOrCreateRequest, ChannelStateResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/query", nil, &channelGetOrCreateRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) MarkRead(ctx context.Context, _type string, id string, markReadRequest MarkReadRequest) (MarkReadResponse, error) {
	var result MarkReadResponse
	pathParams := []string{_type, id}
	err := MakeRequest[MarkReadRequest, MarkReadResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/read", nil, &markReadRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) ShowChannel(ctx context.Context, _type string, id string, showChannelRequest ShowChannelRequest) (ShowChannelResponse, error) {
	var result ShowChannelResponse
	pathParams := []string{_type, id}
	err := MakeRequest[ShowChannelRequest, ShowChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/show", nil, &showChannelRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) TruncateChannel(ctx context.Context, _type string, id string, truncateChannelRequest TruncateChannelRequest) (TruncateChannelResponse, error) {
	var result TruncateChannelResponse
	pathParams := []string{_type, id}
	err := MakeRequest[TruncateChannelRequest, TruncateChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/truncate", nil, &truncateChannelRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) MarkUnread(ctx context.Context, _type string, id string, markUnreadRequest MarkUnreadRequest) (Response, error) {
	var result Response
	pathParams := []string{_type, id}
	err := MakeRequest[MarkUnreadRequest, Response, any](c.client, ctx, "POST", "/api/v2/chat/channels/{_type}/{id}/unread", nil, &markUnreadRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) ListChannelTypes(ctx context.Context) (ListChannelTypesResponse, error) {
	var result ListChannelTypesResponse
	err := MakeRequest[any, ListChannelTypesResponse, any](c.client, ctx, "GET", "/api/v2/chat/channeltypes", nil, nil, &result)
	return result, err
}
func (c *ChatClient) CreateChannelType(ctx context.Context, createChannelTypeRequest CreateChannelTypeRequest) (CreateChannelTypeResponse, error) {
	var result CreateChannelTypeResponse
	err := MakeRequest[CreateChannelTypeRequest, CreateChannelTypeResponse, any](c.client, ctx, "POST", "/api/v2/chat/channeltypes", nil, &createChannelTypeRequest, &result)
	return result, err
}
func (c *ChatClient) DeleteChannelType(ctx context.Context, name string) (Response, error) {
	var result Response
	pathParams := []string{name}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/chat/channeltypes/{name}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetChannelType(ctx context.Context, name string) (Response, error) {
	var result Response
	pathParams := []string{name}
	err := MakeRequest[any, Response, any](c.client, ctx, "GET", "/api/v2/chat/channeltypes/{name}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdateChannelType(ctx context.Context, name string, updateChannelTypeRequest UpdateChannelTypeRequest) (UpdateChannelTypeResponse, error) {
	var result UpdateChannelTypeResponse
	pathParams := []string{name}
	err := MakeRequest[UpdateChannelTypeRequest, UpdateChannelTypeResponse, any](c.client, ctx, "PUT", "/api/v2/chat/channeltypes/{name}", nil, &updateChannelTypeRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) ListCommands(ctx context.Context) (ListCommandsResponse, error) {
	var result ListCommandsResponse
	err := MakeRequest[any, ListCommandsResponse, any](c.client, ctx, "GET", "/api/v2/chat/commands", nil, nil, &result)
	return result, err
}
func (c *ChatClient) CreateCommand(ctx context.Context, createCommandRequest CreateCommandRequest) (CreateCommandResponse, error) {
	var result CreateCommandResponse
	err := MakeRequest[CreateCommandRequest, CreateCommandResponse, any](c.client, ctx, "POST", "/api/v2/chat/commands", nil, &createCommandRequest, &result)
	return result, err
}
func (c *ChatClient) DeleteCommand(ctx context.Context, name string) (DeleteCommandResponse, error) {
	var result DeleteCommandResponse
	pathParams := []string{name}
	err := MakeRequest[any, DeleteCommandResponse, any](c.client, ctx, "DELETE", "/api/v2/chat/commands/{name}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetCommand(ctx context.Context, name string) (GetCommandResponse, error) {
	var result GetCommandResponse
	pathParams := []string{name}
	err := MakeRequest[any, GetCommandResponse, any](c.client, ctx, "GET", "/api/v2/chat/commands/{name}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdateCommand(ctx context.Context, name string, updateCommandRequest UpdateCommandRequest) (UpdateCommandResponse, error) {
	var result UpdateCommandResponse
	pathParams := []string{name}
	err := MakeRequest[UpdateCommandRequest, UpdateCommandResponse, any](c.client, ctx, "PUT", "/api/v2/chat/commands/{name}", nil, &updateCommandRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) ExportChannels(ctx context.Context, exportChannelsRequest ExportChannelsRequest) (ExportChannelsResponse, error) {
	var result ExportChannelsResponse
	err := MakeRequest[ExportChannelsRequest, ExportChannelsResponse, any](c.client, ctx, "POST", "/api/v2/chat/export_channels", nil, &exportChannelsRequest, &result)
	return result, err
}
func (c *ChatClient) GetExportChannelsStatus(ctx context.Context, id string) (GetExportChannelsStatusResponse, error) {
	var result GetExportChannelsStatusResponse
	pathParams := []string{id}
	err := MakeRequest[any, GetExportChannelsStatusResponse, any](c.client, ctx, "GET", "/api/v2/chat/export_channels/{id}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) QueryMembers(ctx context.Context, payload *QueryMembersRequest) (MembersResponse, error) {
	var result MembersResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, MembersResponse](c.client, ctx, "GET", "/api/v2/chat/members", queryParams, nil, &result)
	return result, err
}
func (c *ChatClient) QueryMessageHistory(ctx context.Context, queryMessageHistoryRequest QueryMessageHistoryRequest) (QueryMessageHistoryResponse, error) {
	var result QueryMessageHistoryResponse
	err := MakeRequest[QueryMessageHistoryRequest, QueryMessageHistoryResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/history", nil, &queryMessageHistoryRequest, &result)
	return result, err
}
func (c *ChatClient) DeleteMessage(ctx context.Context, id string, hard *bool, deletedBy *string) (DeleteMessageResponse, error) {
	var result DeleteMessageResponse
	pathParams := []string{id}
	queryParams := map[string]interface{}{
		"hard":       hard,
		"deleted_by": deletedBy,
	}
	err := MakeRequest[any, DeleteMessageResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetMessage(ctx context.Context, id string, showDeletedMessage *bool) (GetMessageResponse, error) {
	var result GetMessageResponse
	pathParams := []string{id}
	queryParams := map[string]interface{}{
		"show_deleted_message": showDeletedMessage,
	}
	err := MakeRequest[any, GetMessageResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdateMessage(ctx context.Context, id string, updateMessageRequest UpdateMessageRequest) (UpdateMessageResponse, error) {
	var result UpdateMessageResponse
	pathParams := []string{id}
	err := MakeRequest[UpdateMessageRequest, UpdateMessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}", nil, &updateMessageRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdateMessagePartial(ctx context.Context, id string, updateMessagePartialRequest UpdateMessagePartialRequest) (UpdateMessagePartialResponse, error) {
	var result UpdateMessagePartialResponse
	pathParams := []string{id}
	err := MakeRequest[UpdateMessagePartialRequest, UpdateMessagePartialResponse, any](c.client, ctx, "PUT", "/api/v2/chat/messages/{id}", nil, &updateMessagePartialRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) RunMessageAction(ctx context.Context, id string, messageActionRequest MessageActionRequest) (MessageResponse, error) {
	var result MessageResponse
	pathParams := []string{id}
	err := MakeRequest[MessageActionRequest, MessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/action", nil, &messageActionRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) CommitMessage(ctx context.Context, id string, commitMessageRequest CommitMessageRequest) (MessageResponse, error) {
	var result MessageResponse
	pathParams := []string{id}
	err := MakeRequest[CommitMessageRequest, MessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/commit", nil, &commitMessageRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) SendReaction(ctx context.Context, id string, sendReactionRequest SendReactionRequest) (SendReactionResponse, error) {
	var result SendReactionResponse
	pathParams := []string{id}
	err := MakeRequest[SendReactionRequest, SendReactionResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/reaction", nil, &sendReactionRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) DeleteReaction(ctx context.Context, id string, _type string, userId *string) (ReactionRemovalResponse, error) {
	var result ReactionRemovalResponse
	pathParams := []string{id, _type}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, ReactionRemovalResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{id}/reaction/{_type}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetReactions(ctx context.Context, id string, limit *int, offset *int) (GetReactionsResponse, error) {
	var result GetReactionsResponse
	pathParams := []string{id}
	queryParams := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}
	err := MakeRequest[any, GetReactionsResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{id}/reactions", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) QueryReactions(ctx context.Context, id string, queryReactionsRequest QueryReactionsRequest) (QueryReactionsResponse, error) {
	var result QueryReactionsResponse
	pathParams := []string{id}
	err := MakeRequest[QueryReactionsRequest, QueryReactionsResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/reactions", nil, &queryReactionsRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) TranslateMessage(ctx context.Context, id string, translateMessageRequest TranslateMessageRequest) (MessageResponse, error) {
	var result MessageResponse
	pathParams := []string{id}
	err := MakeRequest[TranslateMessageRequest, MessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/translate", nil, &translateMessageRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UndeleteMessage(ctx context.Context, id string, updateMessageRequest UpdateMessageRequest) (UpdateMessageResponse, error) {
	var result UpdateMessageResponse
	pathParams := []string{id}
	err := MakeRequest[UpdateMessageRequest, UpdateMessageResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{id}/undelete", nil, &updateMessageRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) CastPollVote(ctx context.Context, messageId string, pollId string, castPollVoteRequest CastPollVoteRequest) (PollVoteResponse, error) {
	var result PollVoteResponse
	pathParams := []string{messageId, pollId}
	err := MakeRequest[CastPollVoteRequest, PollVoteResponse, any](c.client, ctx, "POST", "/api/v2/chat/messages/{message_id}/polls/{poll_id}/vote", nil, &castPollVoteRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) RemovePollVote(ctx context.Context, messageId string, pollId string, voteId string, userId *string) (PollVoteResponse, error) {
	var result PollVoteResponse
	pathParams := []string{messageId, pollId, voteId}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, PollVoteResponse](c.client, ctx, "DELETE", "/api/v2/chat/messages/{message_id}/polls/{poll_id}/vote/{vote_id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetReplies(ctx context.Context, parentId string, limit *int, offset *int, idGte *string, idGt *string, idLte *string, idLt *string, createdAtAfterOrEqual *time.Time, createdAtAfter *time.Time, createdAtBeforeOrEqual *time.Time, createdAtBefore *time.Time, idAround *string, createdAtAround *time.Time, sort *[]*SortParam) (GetRepliesResponse, error) {
	var result GetRepliesResponse
	pathParams := []string{parentId}
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
	err := MakeRequest[any, GetRepliesResponse](c.client, ctx, "GET", "/api/v2/chat/messages/{parent_id}/replies", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) QueryMessageFlags(ctx context.Context, payload *QueryMessageFlagsRequest) (QueryMessageFlagsResponse, error) {
	var result QueryMessageFlagsResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, QueryMessageFlagsResponse](c.client, ctx, "GET", "/api/v2/chat/moderation/flags/message", queryParams, nil, &result)
	return result, err
}
func (c *ChatClient) MuteChannel(ctx context.Context, muteChannelRequest MuteChannelRequest) (MuteChannelResponse, error) {
	var result MuteChannelResponse
	err := MakeRequest[MuteChannelRequest, MuteChannelResponse, any](c.client, ctx, "POST", "/api/v2/chat/moderation/mute/channel", nil, &muteChannelRequest, &result)
	return result, err
}
func (c *ChatClient) UnmuteChannel(ctx context.Context, unmuteChannelRequest UnmuteChannelRequest) (UnmuteResponse, error) {
	var result UnmuteResponse
	err := MakeRequest[UnmuteChannelRequest, UnmuteResponse, any](c.client, ctx, "POST", "/api/v2/chat/moderation/unmute/channel", nil, &unmuteChannelRequest, &result)
	return result, err
}
func (c *ChatClient) CreatePoll(ctx context.Context, createPollRequest CreatePollRequest) (PollResponse, error) {
	var result PollResponse
	err := MakeRequest[CreatePollRequest, PollResponse, any](c.client, ctx, "POST", "/api/v2/chat/polls", nil, &createPollRequest, &result)
	return result, err
}
func (c *ChatClient) UpdatePoll(ctx context.Context, updatePollRequest UpdatePollRequest) (PollResponse, error) {
	var result PollResponse
	err := MakeRequest[UpdatePollRequest, PollResponse, any](c.client, ctx, "PUT", "/api/v2/chat/polls", nil, &updatePollRequest, &result)
	return result, err
}
func (c *ChatClient) QueryPolls(ctx context.Context, userId *string, queryPollsRequest QueryPollsRequest) (QueryPollsResponse, error) {
	var result QueryPollsResponse
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[QueryPollsRequest, QueryPollsResponse](c.client, ctx, "POST", "/api/v2/chat/polls/query", queryParams, &queryPollsRequest, &result)
	return result, err
}
func (c *ChatClient) DeletePoll(ctx context.Context, pollId string, userId *string) (Response, error) {
	var result Response
	pathParams := []string{pollId}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/polls/{poll_id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetPoll(ctx context.Context, pollId string, userId *string) (PollResponse, error) {
	var result PollResponse
	pathParams := []string{pollId}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, PollResponse](c.client, ctx, "GET", "/api/v2/chat/polls/{poll_id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdatePollPartial(ctx context.Context, pollId string, updatePollPartialRequest UpdatePollPartialRequest) (PollResponse, error) {
	var result PollResponse
	pathParams := []string{pollId}
	err := MakeRequest[UpdatePollPartialRequest, PollResponse, any](c.client, ctx, "PATCH", "/api/v2/chat/polls/{poll_id}", nil, &updatePollPartialRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) CreatePollOption(ctx context.Context, pollId string, createPollOptionRequest CreatePollOptionRequest) (PollOptionResponse, error) {
	var result PollOptionResponse
	pathParams := []string{pollId}
	err := MakeRequest[CreatePollOptionRequest, PollOptionResponse, any](c.client, ctx, "POST", "/api/v2/chat/polls/{poll_id}/options", nil, &createPollOptionRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdatePollOption(ctx context.Context, pollId string, updatePollOptionRequest UpdatePollOptionRequest) (PollOptionResponse, error) {
	var result PollOptionResponse
	pathParams := []string{pollId}
	err := MakeRequest[UpdatePollOptionRequest, PollOptionResponse, any](c.client, ctx, "PUT", "/api/v2/chat/polls/{poll_id}/options", nil, &updatePollOptionRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) DeletePollOption(ctx context.Context, pollId string, optionId string, userId *string) (Response, error) {
	var result Response
	pathParams := []string{pollId, optionId}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/chat/polls/{poll_id}/options/{option_id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) GetPollOption(ctx context.Context, pollId string, optionId string, userId *string) (PollOptionResponse, error) {
	var result PollOptionResponse
	pathParams := []string{pollId, optionId}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, PollOptionResponse](c.client, ctx, "GET", "/api/v2/chat/polls/{poll_id}/options/{option_id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) QueryPollVotes(ctx context.Context, pollId string, userId *string, queryPollVotesRequest QueryPollVotesRequest) (PollVotesResponse, error) {
	var result PollVotesResponse
	pathParams := []string{pollId}
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[QueryPollVotesRequest, PollVotesResponse](c.client, ctx, "POST", "/api/v2/chat/polls/{poll_id}/votes", queryParams, &queryPollVotesRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) QueryBannedUsers(ctx context.Context, payload *QueryBannedUsersRequest) (QueryBannedUsersResponse, error) {
	var result QueryBannedUsersResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, QueryBannedUsersResponse](c.client, ctx, "GET", "/api/v2/chat/query_banned_users", queryParams, nil, &result)
	return result, err
}
func (c *ChatClient) Search(ctx context.Context, payload *SearchRequest) (SearchResponse, error) {
	var result SearchResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, SearchResponse](c.client, ctx, "GET", "/api/v2/chat/search", queryParams, nil, &result)
	return result, err
}
func (c *ChatClient) QueryThreads(ctx context.Context, queryThreadsRequest QueryThreadsRequest) (QueryThreadsResponse, error) {
	var result QueryThreadsResponse
	err := MakeRequest[QueryThreadsRequest, QueryThreadsResponse, any](c.client, ctx, "POST", "/api/v2/chat/threads", nil, &queryThreadsRequest, &result)
	return result, err
}
func (c *ChatClient) GetThread(ctx context.Context, messageId string, connectionId *string, replyLimit *int, participantLimit *int, memberLimit *int) (GetThreadResponse, error) {
	var result GetThreadResponse
	pathParams := []string{messageId}
	queryParams := map[string]interface{}{
		"connection_id":     connectionId,
		"reply_limit":       replyLimit,
		"participant_limit": participantLimit,
		"member_limit":      memberLimit,
	}
	err := MakeRequest[any, GetThreadResponse](c.client, ctx, "GET", "/api/v2/chat/threads/{message_id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UpdateThreadPartial(ctx context.Context, messageId string, updateThreadPartialRequest UpdateThreadPartialRequest) (UpdateThreadPartialResponse, error) {
	var result UpdateThreadPartialResponse
	pathParams := []string{messageId}
	err := MakeRequest[UpdateThreadPartialRequest, UpdateThreadPartialResponse, any](c.client, ctx, "PATCH", "/api/v2/chat/threads/{message_id}", nil, &updateThreadPartialRequest, &result, pathParams...)
	return result, err
}
func (c *ChatClient) UnreadCounts(ctx context.Context) (WrappedUnreadCountsResponse, error) {
	var result WrappedUnreadCountsResponse
	err := MakeRequest[any, WrappedUnreadCountsResponse, any](c.client, ctx, "GET", "/api/v2/chat/unread", nil, nil, &result)
	return result, err
}
func (c *ChatClient) UnreadCountsBatch(ctx context.Context, unreadCountsBatchRequest UnreadCountsBatchRequest) (UnreadCountsBatchResponse, error) {
	var result UnreadCountsBatchResponse
	err := MakeRequest[UnreadCountsBatchRequest, UnreadCountsBatchResponse, any](c.client, ctx, "POST", "/api/v2/chat/unread_batch", nil, &unreadCountsBatchRequest, &result)
	return result, err
}
func (c *ChatClient) SendUserCustomEvent(ctx context.Context, userId string, sendUserCustomEventRequest SendUserCustomEventRequest) (Response, error) {
	var result Response
	pathParams := []string{userId}
	err := MakeRequest[SendUserCustomEventRequest, Response, any](c.client, ctx, "POST", "/api/v2/chat/users/{user_id}/event", nil, &sendUserCustomEventRequest, &result, pathParams...)
	return result, err
}
