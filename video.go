package getstream

import (
	"context"
	"net/http"
	"net/url"

	"github.com/getstream/go-client/models"
	"github.com/getstream/go-client"
)


type VideoClient struct {
	client *Client
}

func NewVideoClient(client *Client) *VideoClient {
	return &VideoClient{
		client: NewClient(client),
	}
}


func (c *VideoClient) QueryCallMembers(ctx context.Context,
	QueryCallMembersRequest QueryCallMembersRequest
) (QueryCallMembersResponse, error) {
	var result QueryCallMembersResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/members", nil,QueryCallMembersRequest, &result,)
	return result, err
}


func (c *VideoClient) QueryCallStats(ctx context.Context,
	QueryCallStatsRequest QueryCallStatsRequest
) (QueryCallStatsResponse, error) {
	var result QueryCallStatsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/stats", nil,QueryCallStatsRequest, &result,)
	return result, err
}


func (c *VideoClient) GetCall(ctx context.Context,
	Type string
	, Id string
	, MembersLimit *int
	, Ring *bool
	, Notify *bool
) (GetCallResponse, error) {
	var result GetCallResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}",url.Values{"MembersLimit": []string { MembersLimit }, "Ring": []string { Ring }, "Notify": []string { Notify } }, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) UpdateCall(ctx context.Context,
	Type string
	, Id string
	, UpdateCallRequest UpdateCallRequest
) (UpdateCallResponse, error) {
	var result UpdateCallResponse
	err := MakeRequest(c.client, ctx, "PATCH", "/api/v2/video/call/{type}/{id}", nil,UpdateCallRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) GetOrCreateCall(ctx context.Context,
	Type string
	, Id string
	, GetOrCreateCallRequest GetOrCreateCallRequest
) (GetOrCreateCallResponse, error) {
	var result GetOrCreateCallResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}", nil,GetOrCreateCallRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) BlockUser(ctx context.Context,
	Type string
	, Id string
	, BlockUserRequest BlockUserRequest
) (BlockUserResponse, error) {
	var result BlockUserResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/block", nil,BlockUserRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) SendCallEvent(ctx context.Context,
	Type string
	, Id string
	, SendCallEventRequest SendCallEventRequest
) (SendCallEventResponse, error) {
	var result SendCallEventResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/event", nil,SendCallEventRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) CollectUserFeedback(ctx context.Context,
	Type string
	, Id string
	, Session string
	, CollectUserFeedbackRequest CollectUserFeedbackRequest
) (CollectUserFeedbackResponse, error) {
	var result CollectUserFeedbackResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/feedback/{session}", nil,CollectUserFeedbackRequest, &result, Type, Id, Session,)
	return result, err
}


func (c *VideoClient) GoLive(ctx context.Context,
	Type string
	, Id string
	, GoLiveRequest GoLiveRequest
) (GoLiveResponse, error) {
	var result GoLiveResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/go_live", nil,GoLiveRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) EndCall(ctx context.Context,
	Type string
	, Id string
) (EndCallResponse, error) {
	var result EndCallResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/mark_ended", nil, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) UpdateCallMembers(ctx context.Context,
	Type string
	, Id string
	, UpdateCallMembersRequest UpdateCallMembersRequest
) (UpdateCallMembersResponse, error) {
	var result UpdateCallMembersResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/members", nil,UpdateCallMembersRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) MuteUsers(ctx context.Context,
	Type string
	, Id string
	, MuteUsersRequest MuteUsersRequest
) (MuteUsersResponse, error) {
	var result MuteUsersResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/mute_users", nil,MuteUsersRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) VideoPin(ctx context.Context,
	Type string
	, Id string
	, PinRequest PinRequest
) (PinResponse, error) {
	var result PinResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/pin", nil,PinRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) ListRecordings(ctx context.Context,
	Type string
	, Id string
) (ListRecordingsResponse, error) {
	var result ListRecordingsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/recordings", nil, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) StartHLSBroadcasting(ctx context.Context,
	Type string
	, Id string
) (StartHLSBroadcastingResponse, error) {
	var result StartHLSBroadcastingResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_broadcasting", nil, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) StartRecording(ctx context.Context,
	Type string
	, Id string
	, StartRecordingRequest StartRecordingRequest
) (StartRecordingResponse, error) {
	var result StartRecordingResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_recording", nil,StartRecordingRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) StartTranscription(ctx context.Context,
	Type string
	, Id string
	, StartTranscriptionRequest StartTranscriptionRequest
) (StartTranscriptionResponse, error) {
	var result StartTranscriptionResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_transcription", nil,StartTranscriptionRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) GetCallStats(ctx context.Context,
	Type string
	, Id string
	, Session string
) (GetCallStatsResponse, error) {
	var result GetCallStatsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/stats/{session}", nil, nil, &result, Type, Id, Session,)
	return result, err
}


func (c *VideoClient) StopHLSBroadcasting(ctx context.Context,
	Type string
	, Id string
) (StopHLSBroadcastingResponse, error) {
	var result StopHLSBroadcastingResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_broadcasting", nil, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) StopLive(ctx context.Context,
	Type string
	, Id string
) (StopLiveResponse, error) {
	var result StopLiveResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_live", nil, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) StopRecording(ctx context.Context,
	Type string
	, Id string
) (StopRecordingResponse, error) {
	var result StopRecordingResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_recording", nil, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) StopTranscription(ctx context.Context,
	Type string
	, Id string
) (StopTranscriptionResponse, error) {
	var result StopTranscriptionResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_transcription", nil, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) ListTranscriptions(ctx context.Context,
	Type string
	, Id string
) (ListTranscriptionsResponse, error) {
	var result ListTranscriptionsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/transcriptions", nil, nil, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) UnblockUser(ctx context.Context,
	Type string
	, Id string
	, UnblockUserRequest UnblockUserRequest
) (UnblockUserResponse, error) {
	var result UnblockUserResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/unblock", nil,UnblockUserRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) VideoUnpin(ctx context.Context,
	Type string
	, Id string
	, UnpinRequest UnpinRequest
) (UnpinResponse, error) {
	var result UnpinResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/unpin", nil,UnpinRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) UpdateUserPermissions(ctx context.Context,
	Type string
	, Id string
	, UpdateUserPermissionsRequest UpdateUserPermissionsRequest
) (UpdateUserPermissionsResponse, error) {
	var result UpdateUserPermissionsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/user_permissions", nil,UpdateUserPermissionsRequest, &result, Type, Id,)
	return result, err
}


func (c *VideoClient) DeleteRecording(ctx context.Context,
	Type string
	, Id string
	, Session string
	, Filename string
) (DeleteRecordingResponse, error) {
	var result DeleteRecordingResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/video/call/{type}/{id}/{session}/recordings/{filename}", nil, nil, &result, Type, Id, Session, Filename,)
	return result, err
}


func (c *VideoClient) DeleteTranscription(ctx context.Context,
	Type string
	, Id string
	, Session string
	, Filename string
) (DeleteTranscriptionResponse, error) {
	var result DeleteTranscriptionResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/video/call/{type}/{id}/{session}/transcriptions/{filename}", nil, nil, &result, Type, Id, Session, Filename,)
	return result, err
}


func (c *VideoClient) QueryCalls(ctx context.Context,
	QueryCallsRequest QueryCallsRequest
) (QueryCallsResponse, error) {
	var result QueryCallsResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/calls", nil,QueryCallsRequest, &result,)
	return result, err
}


func (c *VideoClient) ListCallTypes(ctx context.Context) (ListCallTypeResponse, error) {
	var result ListCallTypeResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/video/calltypes", nil, nil, &result,)
	return result, err
}


func (c *VideoClient) CreateCallType(ctx context.Context,
	CreateCallTypeRequest CreateCallTypeRequest
) (CreateCallTypeResponse, error) {
	var result CreateCallTypeResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/video/calltypes", nil,CreateCallTypeRequest, &result,)
	return result, err
}


func (c *VideoClient) DeleteCallType(ctx context.Context,
	Name string
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/video/calltypes/{name}", nil, nil, &result, Name,)
	return result, err
}


func (c *VideoClient) GetCallType(ctx context.Context,
	Name string
) (GetCallTypeResponse, error) {
	var result GetCallTypeResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/video/calltypes/{name}", nil, nil, &result, Name,)
	return result, err
}


func (c *VideoClient) UpdateCallType(ctx context.Context,
	Name string
	, UpdateCallTypeRequest UpdateCallTypeRequest
) (UpdateCallTypeResponse, error) {
	var result UpdateCallTypeResponse
	err := MakeRequest(c.client, ctx, "PUT", "/api/v2/video/calltypes/{name}", nil,UpdateCallTypeRequest, &result, Name,)
	return result, err
}


func (c *VideoClient) GetEdges(ctx context.Context) (GetEdgesResponse, error) {
	var result GetEdgesResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/video/edges", nil, nil, &result,)
	return result, err
}










