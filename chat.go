package getstream

import (
	"context"
	"net/url"
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

func (c *ChatClient) QueryChannels(ctx context.Context, QueryChannelsRequest QueryChannelsRequest) (QueryChannelsResponse, error) {
	var result QueryChannelsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels", nil, QueryChannelsRequest, &result)
	return result, err
}

func (c *ChatClient) DeleteChannels(ctx context.Context, DeleteChannelsRequest DeleteChannelsRequest) (DeleteChannelsResponse, error) {
	var result DeleteChannelsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/delete", nil, DeleteChannelsRequest, &result)
	return result, err
}

func (c *ChatClient) MarkChannelsRead(ctx context.Context, MarkChannelsReadRequest MarkChannelsReadRequest) (MarkReadResponse, error) {
	var result MarkReadResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/read", nil, MarkChannelsReadRequest, &result)
	return result, err
}

func (c *ChatClient) GetOrCreateDistinctChannel(ctx context.Context, Type string, ChannelGetOrCreateRequest ChannelGetOrCreateRequest) (ChannelStateResponse, error) {
	var result ChannelStateResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/query", nil, ChannelGetOrCreateRequest, &result, Type)
	return result, err
}

func (c *ChatClient) DeleteChannel(ctx context.Context, Type string, Id string, HardDelete *bool) (DeleteChannelResponse, error) {
	var result DeleteChannelResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}", url.Values{"HardDelete": []string{HardDelete}}, nil, &result, Type, Id)
	return result, err
}

func (c *ChatClient) UpdateChannelPartial(ctx context.Context, Type string, Id string, UpdateChannelPartialRequest UpdateChannelPartialRequest) (UpdateChannelPartialResponse, error) {
	var result UpdateChannelPartialResponse
	err := MakeRequest(c.client, ctx, "PATCH", "/api/v2/chat/channels/{type}/{id}", nil, UpdateChannelPartialRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) UpdateChannel(ctx context.Context, Type string, Id string, UpdateChannelRequest UpdateChannelRequest) (UpdateChannelResponse, error) {
	var result UpdateChannelResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}", nil, UpdateChannelRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) SendEvent(ctx context.Context, Type string, Id string, SendEventRequest SendEventRequest) (EventResponse, error) {
	var result EventResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/event", nil, SendEventRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) DeleteFile(ctx context.Context, Type string, Id string, Url *string) (FileDeleteResponse, error) {
	var result FileDeleteResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}/file", url.Values{"Url": []string{Url}}, nil, &result, Type, Id)
	return result, err
}

func (c *ChatClient) UploadFile(ctx context.Context, Type string, Id string, FileUploadRequest FileUploadRequest) (FileUploadResponse, error) {
	var result FileUploadResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/file", nil, FileUploadRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) HideChannel(ctx context.Context, Type string, Id string, HideChannelRequest HideChannelRequest) (HideChannelResponse, error) {
	var result HideChannelResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/hide", nil, HideChannelRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) DeleteImage(ctx context.Context, Type string, Id string, Url *string) (FileDeleteResponse, error) {
	var result FileDeleteResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/channels/{type}/{id}/image", url.Values{"Url": []string{Url}}, nil, &result, Type, Id)
	return result, err
}

func (c *ChatClient) UploadImage(ctx context.Context, Type string, Id string, ImageUploadRequest ImageUploadRequest) (ImageUploadResponse, error) {
	var result ImageUploadResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/image", nil, ImageUploadRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) SendMessage(ctx context.Context, Type string, Id string, SendMessageRequest SendMessageRequest) (SendMessageResponse, error) {
	var result SendMessageResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/message", nil, SendMessageRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) GetManyMessages(ctx context.Context, Type string, Id string, Ids []string) (GetManyMessagesResponse, error) {
	var result GetManyMessagesResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/channels/{type}/{id}/messages", url.Values{"Ids": []string{Ids}}, nil, &result, Type, Id)
	return result, err
}

func (c *ChatClient) GetOrCreateChannel(ctx context.Context, Type string, Id string, ChannelGetOrCreateRequest ChannelGetOrCreateRequest) (ChannelStateResponse, error) {
	var result ChannelStateResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/query", nil, ChannelGetOrCreateRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) MarkRead(ctx context.Context, Type string, Id string, MarkReadRequest MarkReadRequest) (MarkReadResponse, error) {
	var result MarkReadResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/read", nil, MarkReadRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) ShowChannel(ctx context.Context, Type string, Id string, ShowChannelRequest ShowChannelRequest) (ShowChannelResponse, error) {
	var result ShowChannelResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/show", nil, ShowChannelRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) TruncateChannel(ctx context.Context, Type string, Id string, TruncateChannelRequest TruncateChannelRequest) (TruncateChannelResponse, error) {
	var result TruncateChannelResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/truncate", nil, TruncateChannelRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) MarkUnread(ctx context.Context, Type string, Id string, MarkUnreadRequest MarkUnreadRequest) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channels/{type}/{id}/unread", nil, MarkUnreadRequest, &result, Type, Id)
	return result, err
}

func (c *ChatClient) ListChannelTypes(ctx context.Context) (ListChannelTypesResponse, error) {
	var result ListChannelTypesResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/channeltypes", nil, nil, &result)
	return result, err
}

func (c *ChatClient) CreateChannelType(ctx context.Context, CreateChannelTypeRequest CreateChannelTypeRequest) (CreateChannelTypeResponse, error) {
	var result CreateChannelTypeResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/channeltypes", nil, CreateChannelTypeRequest, &result)
	return result, err
}

func (c *ChatClient) DeleteChannelType(ctx context.Context, Name string) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/channeltypes/{name}", nil, nil, &result, Name)
	return result, err
}

func (c *ChatClient) GetChannelType(ctx context.Context, Name string) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/channeltypes/{name}", nil, nil, &result, Name)
	return result, err
}

func (c *ChatClient) UpdateChannelType(ctx context.Context, Name string, UpdateChannelTypeRequest UpdateChannelTypeRequest) (UpdateChannelTypeResponse, error) {
	var result UpdateChannelTypeResponse
	err := MakeRequest(c.client, ctx, "PUT", "/api/v2/chat/channeltypes/{name}", nil, UpdateChannelTypeRequest, &result, Name)
	return result, err
}

func (c *ChatClient) ListCommands(ctx context.Context) (ListCommandsResponse, error) {
	var result ListCommandsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/commands", nil, nil, &result)
	return result, err
}

func (c *ChatClient) CreateCommand(ctx context.Context, CreateCommandRequest CreateCommandRequest) (CreateCommandResponse, error) {
	var result CreateCommandResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/commands", nil, CreateCommandRequest, &result)
	return result, err
}

func (c *ChatClient) DeleteCommand(ctx context.Context, Name string) (DeleteCommandResponse, error) {
	var result DeleteCommandResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/commands/{name}", nil, nil, &result, Name)
	return result, err
}

func (c *ChatClient) GetCommand(ctx context.Context, Name string) (GetCommandResponse, error) {
	var result GetCommandResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/commands/{name}", nil, nil, &result, Name)
	return result, err
}

func (c *ChatClient) UpdateCommand(ctx context.Context, Name string, UpdateCommandRequest UpdateCommandRequest) (UpdateCommandResponse, error) {
	var result UpdateCommandResponse
	err := MakeRequest(c.client, ctx, "PUT", "/api/v2/chat/commands/{name}", nil, UpdateCommandRequest, &result, Name)
	return result, err
}

func (c *ChatClient) ExportChannels(ctx context.Context, ExportChannelsRequest ExportChannelsRequest) (ExportChannelsResponse, error) {
	var result ExportChannelsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/export_channels", nil, ExportChannelsRequest, &result)
	return result, err
}

func (c *ChatClient) GetExportChannelsStatus(ctx context.Context, Id string) (GetExportChannelsStatusResponse, error) {
	var result GetExportChannelsStatusResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/export_channels/{id}", nil, nil, &result, Id)
	return result, err
}

func (c *ChatClient) QueryMembers(ctx context.Context, Payload *QueryMembersRequest) (MembersResponse, error) {
	var result MembersResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/members", url.Values{"Payload": []string{Payload}}, nil, &result)
	return result, err
}

func (c *ChatClient) QueryMessageHistory(ctx context.Context, QueryMessageHistoryRequest QueryMessageHistoryRequest) (QueryMessageHistoryResponse, error) {
	var result QueryMessageHistoryResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/history", nil, QueryMessageHistoryRequest, &result)
	return result, err
}

func (c *ChatClient) DeleteMessage(ctx context.Context, Id string, Hard *bool, DeletedBy *string) (DeleteMessageResponse, error) {
	var result DeleteMessageResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/messages/{id}", url.Values{"Hard": []string{Hard}, "DeletedBy": []string{DeletedBy}}, nil, &result, Id)
	return result, err
}

func (c *ChatClient) GetMessage(ctx context.Context, Id string, ShowDeletedMessage *bool) (GetMessageResponse, error) {
	var result GetMessageResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/messages/{id}", url.Values{"ShowDeletedMessage": []string{ShowDeletedMessage}}, nil, &result, Id)
	return result, err
}

func (c *ChatClient) UpdateMessage(ctx context.Context, Id string, UpdateMessageRequest UpdateMessageRequest) (UpdateMessageResponse, error) {
	var result UpdateMessageResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/{id}", nil, UpdateMessageRequest, &result, Id)
	return result, err
}

func (c *ChatClient) UpdateMessagePartial(ctx context.Context, Id string, UpdateMessagePartialRequest UpdateMessagePartialRequest) (UpdateMessagePartialResponse, error) {
	var result UpdateMessagePartialResponse
	err := MakeRequest(c.client, ctx, "PUT", "/api/v2/chat/messages/{id}", nil, UpdateMessagePartialRequest, &result, Id)
	return result, err
}

func (c *ChatClient) RunMessageAction(ctx context.Context, Id string, MessageActionRequest MessageActionRequest) (MessageResponse, error) {
	var result MessageResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/{id}/action", nil, MessageActionRequest, &result, Id)
	return result, err
}

func (c *ChatClient) CommitMessage(ctx context.Context, Id string, CommitMessageRequest CommitMessageRequest) (MessageResponse, error) {
	var result MessageResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/{id}/commit", nil, CommitMessageRequest, &result, Id)
	return result, err
}

func (c *ChatClient) SendReaction(ctx context.Context, Id string, SendReactionRequest SendReactionRequest) (SendReactionResponse, error) {
	var result SendReactionResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/{id}/reaction", nil, SendReactionRequest, &result, Id)
	return result, err
}

func (c *ChatClient) DeleteReaction(ctx context.Context, Id string, Type string, UserId *string) (ReactionRemovalResponse, error) {
	var result ReactionRemovalResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/messages/{id}/reaction/{type}", url.Values{"UserId": []string{UserId}}, nil, &result, Id, Type)
	return result, err
}

func (c *ChatClient) GetReactions(ctx context.Context, Id string, Limit *int, Offset *int) (GetReactionsResponse, error) {
	var result GetReactionsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/messages/{id}/reactions", url.Values{"Limit": []string{Limit}, "Offset": []string{Offset}}, nil, &result, Id)
	return result, err
}

func (c *ChatClient) QueryReactions(ctx context.Context, Id string, QueryReactionsRequest QueryReactionsRequest) (QueryReactionsResponse, error) {
	var result QueryReactionsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/{id}/reactions", nil, QueryReactionsRequest, &result, Id)
	return result, err
}

func (c *ChatClient) TranslateMessage(ctx context.Context, Id string, TranslateMessageRequest TranslateMessageRequest) (MessageResponse, error) {
	var result MessageResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/{id}/translate", nil, TranslateMessageRequest, &result, Id)
	return result, err
}

func (c *ChatClient) UndeleteMessage(ctx context.Context, Id string, UpdateMessageRequest UpdateMessageRequest) (UpdateMessageResponse, error) {
	var result UpdateMessageResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/{id}/undelete", nil, UpdateMessageRequest, &result, Id)
	return result, err
}

func (c *ChatClient) CastPollVote(ctx context.Context, MessageId string, PollId string, CastPollVoteRequest CastPollVoteRequest) (PollVoteResponse, error) {
	var result PollVoteResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/messages/{message_id}/polls/{poll_id}/vote", nil, CastPollVoteRequest, &result, MessageId, PollId)
	return result, err
}

func (c *ChatClient) RemovePollVote(ctx context.Context, MessageId string, PollId string, VoteId string, UserId *string) (PollVoteResponse, error) {
	var result PollVoteResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/messages/{message_id}/polls/{poll_id}/vote/{vote_id}", url.Values{"UserId": []string{UserId}}, nil, &result, MessageId, PollId, VoteId)
	return result, err
}

func (c *ChatClient) GetReplies(ctx context.Context, ParentId string, Limit *int, Offset *int, IdGte *string, IdGt *string, IdLte *string, IdLt *string, CreatedAtAfterOrEqual *time.Time, CreatedAtAfter *time.Time, CreatedAtBeforeOrEqual *time.Time, CreatedAtBefore *time.Time, IdAround *string, CreatedAtAround *time.Time, Sort *[]*SortParam) (GetRepliesResponse, error) {
	var result GetRepliesResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/messages/{parent_id}/replies", url.Values{"Limit": []string{Limit}, "Offset": []string{Offset}, "IdGte": []string{IdGte}, "IdGt": []string{IdGt}, "IdLte": []string{IdLte}, "IdLt": []string{IdLt}, "CreatedAtAfterOrEqual": []string{CreatedAtAfterOrEqual}, "CreatedAtAfter": []string{CreatedAtAfter}, "CreatedAtBeforeOrEqual": []string{CreatedAtBeforeOrEqual}, "CreatedAtBefore": []string{CreatedAtBefore}, "IdAround": []string{IdAround}, "CreatedAtAround": []string{CreatedAtAround}, "Sort": []string{Sort}}, nil, &result, ParentId)
	return result, err
}

func (c *ChatClient) QueryMessageFlags(ctx context.Context, Payload *QueryMessageFlagsRequest) (QueryMessageFlagsResponse, error) {
	var result QueryMessageFlagsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/moderation/flags/message", url.Values{"Payload": []string{Payload}}, nil, &result)
	return result, err
}

func (c *ChatClient) MuteChannel(ctx context.Context, MuteChannelRequest MuteChannelRequest) (MuteChannelResponse, error) {
	var result MuteChannelResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/moderation/mute/channel", nil, MuteChannelRequest, &result)
	return result, err
}

func (c *ChatClient) UnmuteChannel(ctx context.Context, UnmuteChannelRequest UnmuteChannelRequest) (UnmuteResponse, error) {
	var result UnmuteResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/moderation/unmute/channel", nil, UnmuteChannelRequest, &result)
	return result, err
}

func (c *ChatClient) CreatePoll(ctx context.Context, CreatePollRequest CreatePollRequest) (PollResponse, error) {
	var result PollResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/polls", nil, CreatePollRequest, &result)
	return result, err
}

func (c *ChatClient) UpdatePoll(ctx context.Context, UpdatePollRequest UpdatePollRequest) (PollResponse, error) {
	var result PollResponse
	err := MakeRequest(c.client, ctx, "PUT", "/api/v2/chat/polls", nil, UpdatePollRequest, &result)
	return result, err
}

func (c *ChatClient) QueryPolls(ctx context.Context, UserId *string, QueryPollsRequest QueryPollsRequest) (QueryPollsResponse, error) {
	var result QueryPollsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/polls/query", url.Values{"UserId": []string{UserId}}, QueryPollsRequest, &result)
	return result, err
}

func (c *ChatClient) DeletePoll(ctx context.Context, PollId string, UserId *string) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/polls/{poll_id}", url.Values{"UserId": []string{UserId}}, nil, &result, PollId)
	return result, err
}

func (c *ChatClient) GetPoll(ctx context.Context, PollId string, UserId *string) (PollResponse, error) {
	var result PollResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/polls/{poll_id}", url.Values{"UserId": []string{UserId}}, nil, &result, PollId)
	return result, err
}

func (c *ChatClient) UpdatePollPartial(ctx context.Context, PollId string, UpdatePollPartialRequest UpdatePollPartialRequest) (PollResponse, error) {
	var result PollResponse
	err := MakeRequest(c.client, ctx, "PATCH", "/api/v2/chat/polls/{poll_id}", nil, UpdatePollPartialRequest, &result, PollId)
	return result, err
}

func (c *ChatClient) CreatePollOption(ctx context.Context, PollId string, CreatePollOptionRequest CreatePollOptionRequest) (PollOptionResponse, error) {
	var result PollOptionResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/polls/{poll_id}/options", nil, CreatePollOptionRequest, &result, PollId)
	return result, err
}

func (c *ChatClient) UpdatePollOption(ctx context.Context, PollId string, UpdatePollOptionRequest UpdatePollOptionRequest) (PollOptionResponse, error) {
	var result PollOptionResponse
	err := MakeRequest(c.client, ctx, "PUT", "/api/v2/chat/polls/{poll_id}/options", nil, UpdatePollOptionRequest, &result, PollId)
	return result, err
}

func (c *ChatClient) DeletePollOption(ctx context.Context, PollId string, OptionId string, UserId *string) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/chat/polls/{poll_id}/options/{option_id}", url.Values{"UserId": []string{UserId}}, nil, &result, PollId, OptionId)
	return result, err
}

func (c *ChatClient) GetPollOption(ctx context.Context, PollId string, OptionId string, UserId *string) (PollOptionResponse, error) {
	var result PollOptionResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/polls/{poll_id}/options/{option_id}", url.Values{"UserId": []string{UserId}}, nil, &result, PollId, OptionId)
	return result, err
}

func (c *ChatClient) QueryPollVotes(ctx context.Context, PollId string, UserId *string, QueryPollVotesRequest QueryPollVotesRequest) (PollVotesResponse, error) {
	var result PollVotesResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/polls/{poll_id}/votes", url.Values{"UserId": []string{UserId}}, QueryPollVotesRequest, &result, PollId)
	return result, err
}

func (c *ChatClient) QueryBannedUsers(ctx context.Context, Payload *QueryBannedUsersRequest) (QueryBannedUsersResponse, error) {
	var result QueryBannedUsersResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/query_banned_users", url.Values{"Payload": []string{Payload}}, nil, &result)
	return result, err
}

func (c *ChatClient) Search(ctx context.Context, Payload *SearchRequest) (SearchResponse, error) {
	var result SearchResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/search", url.Values{"Payload": []string{Payload}}, nil, &result)
	return result, err
}

func (c *ChatClient) QueryThreads(ctx context.Context, QueryThreadsRequest QueryThreadsRequest) (QueryThreadsResponse, error) {
	var result QueryThreadsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/threads", nil, QueryThreadsRequest, &result)
	return result, err
}

func (c *ChatClient) GetThread(ctx context.Context, MessageId string, ConnectionId *string, ReplyLimit *int, ParticipantLimit *int, MemberLimit *int) (GetThreadResponse, error) {
	var result GetThreadResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/threads/{message_id}", url.Values{"ConnectionId": []string{ConnectionId}, "ReplyLimit": []string{ReplyLimit}, "ParticipantLimit": []string{ParticipantLimit}, "MemberLimit": []string{MemberLimit}}, nil, &result, MessageId)
	return result, err
}

func (c *ChatClient) UpdateThreadPartial(ctx context.Context, MessageId string, UpdateThreadPartialRequest UpdateThreadPartialRequest) (UpdateThreadPartialResponse, error) {
	var result UpdateThreadPartialResponse
	err := MakeRequest(c.client, ctx, "PATCH", "/api/v2/chat/threads/{message_id}", nil, UpdateThreadPartialRequest, &result, MessageId)
	return result, err
}

func (c *ChatClient) UnreadCounts(ctx context.Context) (WrappedUnreadCountsResponse, error) {
	var result WrappedUnreadCountsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/chat/unread", nil, nil, &result)
	return result, err
}

func (c *ChatClient) UnreadCountsBatch(ctx context.Context, UnreadCountsBatchRequest UnreadCountsBatchRequest) (UnreadCountsBatchResponse, error) {
	var result UnreadCountsBatchResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/unread_batch", nil, UnreadCountsBatchRequest, &result)
	return result, err
}

func (c *ChatClient) SendUserCustomEvent(ctx context.Context, UserId string, SendUserCustomEventRequest SendUserCustomEventRequest) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/chat/users/{user_id}/event", nil, SendUserCustomEventRequest, &result, UserId)
	return result, err
}
