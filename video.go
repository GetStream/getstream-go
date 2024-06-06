package getstream

import (
	"context"
)

type VideoClient struct {
	client *Client
}

func NewVideoClient(client *Client) *VideoClient {
	return &VideoClient{
		client: client,
	}
}

func (c *VideoClient) QueryCallMembers(ctx context.Context, request *QueryCallMembersRequest) (*QueryCallMembersResponse, error) {
	var result QueryCallMembersResponse
	err := MakeRequest[QueryCallMembersRequest, QueryCallMembersResponse, any](c.client, ctx, "POST", "/api/v2/video/call/members", nil, request, &result, nil)
	return &result, err
}

func (c *VideoClient) QueryCallStats(ctx context.Context, request *QueryCallStatsRequest) (*QueryCallStatsResponse, error) {
	var result QueryCallStatsResponse
	err := MakeRequest[QueryCallStatsRequest, QueryCallStatsResponse, any](c.client, ctx, "POST", "/api/v2/video/call/stats", nil, request, &result, nil)
	return &result, err
}

func (c *VideoClient) GetCall(ctx context.Context, _type string, id string, membersLimit *int, ring *bool, notify *bool) (*GetCallResponse, error) {
	var result GetCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	queryParams := map[string]interface{}{
		"members_limit": membersLimit,
		"ring":          ring,
		"notify":        notify,
	}
	err := MakeRequest[any, GetCallResponse](c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}", queryParams, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) UpdateCall(ctx context.Context, _type string, id string, request *UpdateCallRequest) (*UpdateCallResponse, error) {
	var result UpdateCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[UpdateCallRequest, UpdateCallResponse, any](c.client, ctx, "PATCH", "/api/v2/video/call/{type}/{id}", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) GetOrCreateCall(ctx context.Context, _type string, id string, request *GetOrCreateCallRequest) (*GetOrCreateCallResponse, error) {
	var result GetOrCreateCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[GetOrCreateCallRequest, GetOrCreateCallResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) BlockUser(ctx context.Context, _type string, id string, request *BlockUserRequest) (*BlockUserResponse, error) {
	var result BlockUserResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[BlockUserRequest, BlockUserResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/block", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) SendCallEvent(ctx context.Context, _type string, id string, request *SendCallEventRequest) (*SendCallEventResponse, error) {
	var result SendCallEventResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[SendCallEventRequest, SendCallEventResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/event", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) CollectUserFeedback(ctx context.Context, _type string, id string, session string, request *CollectUserFeedbackRequest) (*CollectUserFeedbackResponse, error) {
	var result CollectUserFeedbackResponse
	pathParams := map[string]string{
		"type":    _type,
		"id":      id,
		"session": session,
	}
	err := MakeRequest[CollectUserFeedbackRequest, CollectUserFeedbackResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/feedback/{session}", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) GoLive(ctx context.Context, _type string, id string, request *GoLiveRequest) (*GoLiveResponse, error) {
	var result GoLiveResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[GoLiveRequest, GoLiveResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/go_live", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) EndCall(ctx context.Context, _type string, id string) (*EndCallResponse, error) {
	var result EndCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[any, EndCallResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/mark_ended", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) UpdateCallMembers(ctx context.Context, _type string, id string, request *UpdateCallMembersRequest) (*UpdateCallMembersResponse, error) {
	var result UpdateCallMembersResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[UpdateCallMembersRequest, UpdateCallMembersResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/members", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) MuteUsers(ctx context.Context, _type string, id string, request *MuteUsersRequest) (*MuteUsersResponse, error) {
	var result MuteUsersResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[MuteUsersRequest, MuteUsersResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/mute_users", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) VideoPin(ctx context.Context, _type string, id string, request *PinRequest) (*PinResponse, error) {
	var result PinResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[PinRequest, PinResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/pin", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) ListRecordings(ctx context.Context, _type string, id string) (*ListRecordingsResponse, error) {
	var result ListRecordingsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[any, ListRecordingsResponse, any](c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/recordings", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) StartHLSBroadcasting(ctx context.Context, _type string, id string) (*StartHLSBroadcastingResponse, error) {
	var result StartHLSBroadcastingResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[any, StartHLSBroadcastingResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_broadcasting", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) StartRecording(ctx context.Context, _type string, id string, request *StartRecordingRequest) (*StartRecordingResponse, error) {
	var result StartRecordingResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[StartRecordingRequest, StartRecordingResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_recording", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) StartTranscription(ctx context.Context, _type string, id string, request *StartTranscriptionRequest) (*StartTranscriptionResponse, error) {
	var result StartTranscriptionResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[StartTranscriptionRequest, StartTranscriptionResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_transcription", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) GetCallStats(ctx context.Context, _type string, id string, session string) (*GetCallStatsResponse, error) {
	var result GetCallStatsResponse
	pathParams := map[string]string{
		"type":    _type,
		"id":      id,
		"session": session,
	}
	err := MakeRequest[any, GetCallStatsResponse, any](c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/stats/{session}", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) StopHLSBroadcasting(ctx context.Context, _type string, id string) (*StopHLSBroadcastingResponse, error) {
	var result StopHLSBroadcastingResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[any, StopHLSBroadcastingResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_broadcasting", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) StopLive(ctx context.Context, _type string, id string) (*StopLiveResponse, error) {
	var result StopLiveResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[any, StopLiveResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_live", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) StopRecording(ctx context.Context, _type string, id string) (*StopRecordingResponse, error) {
	var result StopRecordingResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[any, StopRecordingResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_recording", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) StopTranscription(ctx context.Context, _type string, id string) (*StopTranscriptionResponse, error) {
	var result StopTranscriptionResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[any, StopTranscriptionResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_transcription", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) ListTranscriptions(ctx context.Context, _type string, id string) (*ListTranscriptionsResponse, error) {
	var result ListTranscriptionsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[any, ListTranscriptionsResponse, any](c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/transcriptions", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) UnblockUser(ctx context.Context, _type string, id string, request *UnblockUserRequest) (*UnblockUserResponse, error) {
	var result UnblockUserResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[UnblockUserRequest, UnblockUserResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/unblock", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) VideoUnpin(ctx context.Context, _type string, id string, request *UnpinRequest) (*UnpinResponse, error) {
	var result UnpinResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[UnpinRequest, UnpinResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/unpin", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) UpdateUserPermissions(ctx context.Context, _type string, id string, request *UpdateUserPermissionsRequest) (*UpdateUserPermissionsResponse, error) {
	var result UpdateUserPermissionsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	err := MakeRequest[UpdateUserPermissionsRequest, UpdateUserPermissionsResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/user_permissions", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) DeleteRecording(ctx context.Context, _type string, id string, session string, filename string) (*DeleteRecordingResponse, error) {
	var result DeleteRecordingResponse
	pathParams := map[string]string{
		"type":     _type,
		"id":       id,
		"session":  session,
		"filename": filename,
	}
	err := MakeRequest[any, DeleteRecordingResponse, any](c.client, ctx, "DELETE", "/api/v2/video/call/{type}/{id}/{session}/recordings/{filename}", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) DeleteTranscription(ctx context.Context, _type string, id string, session string, filename string) (*DeleteTranscriptionResponse, error) {
	var result DeleteTranscriptionResponse
	pathParams := map[string]string{
		"type":     _type,
		"id":       id,
		"session":  session,
		"filename": filename,
	}
	err := MakeRequest[any, DeleteTranscriptionResponse, any](c.client, ctx, "DELETE", "/api/v2/video/call/{type}/{id}/{session}/transcriptions/{filename}", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) QueryCalls(ctx context.Context, request *QueryCallsRequest) (*QueryCallsResponse, error) {
	var result QueryCallsResponse
	err := MakeRequest[QueryCallsRequest, QueryCallsResponse, any](c.client, ctx, "POST", "/api/v2/video/calls", nil, request, &result, nil)
	return &result, err
}

func (c *VideoClient) ListCallTypes(ctx context.Context) (*ListCallTypeResponse, error) {
	var result ListCallTypeResponse
	err := MakeRequest[any, ListCallTypeResponse, any](c.client, ctx, "GET", "/api/v2/video/calltypes", nil, nil, &result, nil)
	return &result, err
}

func (c *VideoClient) CreateCallType(ctx context.Context, request *CreateCallTypeRequest) (*CreateCallTypeResponse, error) {
	var result CreateCallTypeResponse
	err := MakeRequest[CreateCallTypeRequest, CreateCallTypeResponse, any](c.client, ctx, "POST", "/api/v2/video/calltypes", nil, request, &result, nil)
	return &result, err
}

func (c *VideoClient) DeleteCallType(ctx context.Context, name string) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/video/calltypes/{name}", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) GetCallType(ctx context.Context, name string) (*GetCallTypeResponse, error) {
	var result GetCallTypeResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, GetCallTypeResponse, any](c.client, ctx, "GET", "/api/v2/video/calltypes/{name}", nil, nil, &result, pathParams)
	return &result, err
}

func (c *VideoClient) UpdateCallType(ctx context.Context, name string, request *UpdateCallTypeRequest) (*UpdateCallTypeResponse, error) {
	var result UpdateCallTypeResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[UpdateCallTypeRequest, UpdateCallTypeResponse, any](c.client, ctx, "PUT", "/api/v2/video/calltypes/{name}", nil, request, &result, pathParams)
	return &result, err
}

func (c *VideoClient) GetEdges(ctx context.Context) (*GetEdgesResponse, error) {
	var result GetEdgesResponse
	err := MakeRequest[any, GetEdgesResponse, any](c.client, ctx, "GET", "/api/v2/video/edges", nil, nil, &result, nil)
	return &result, err
}
