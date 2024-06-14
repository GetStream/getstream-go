package getstream

import "context"

type Call struct {
	callType string
	callID   string
	client   *VideoClient
}

func NewCall(callType string, callID string, client *VideoClient) *Call {
	return &Call{
		callType: callType,
		callID:   callID,
		client:   client,
	}
}

func (c *Call) Get(ctx context.Context, params *GetCallParams) (*GetCallResponse, error) {
	return c.client.GetCall(ctx, c.callType, c.callID, params)
}

func (c *Call) Update(ctx context.Context, request *UpdateCallRequest) (*UpdateCallResponse, error) {
	return c.client.UpdateCall(ctx, c.callType, c.callID, request)
}

func (c *Call) GetOrCreate(ctx context.Context, request *GetOrCreateCallRequest) (*GetOrCreateCallResponse, error) {
	return c.client.GetOrCreateCall(ctx, c.callType, c.callID, request)
}

func (c *Call) BlockUser(ctx context.Context, request *BlockUserRequest) (*BlockUserResponse, error) {
	return c.client.BlockUser(ctx, c.callType, c.callID, request)
}

func (c *Call) SendCallEvent(ctx context.Context, request *SendCallEventRequest) (*SendCallEventResponse, error) {
	return c.client.SendCallEvent(ctx, c.callType, c.callID, request)
}

func (c *Call) CollectUserFeedback(ctx context.Context, session string, request *CollectUserFeedbackRequest) (*CollectUserFeedbackResponse, error) {
	return c.client.CollectUserFeedback(ctx, c.callType, c.callID, session, request)
}

func (c *Call) GoLive(ctx context.Context, request *GoLiveRequest) (*GoLiveResponse, error) {
	return c.client.GoLive(ctx, c.callType, c.callID, request)
}

func (c *Call) End(ctx context.Context) (*EndCallResponse, error) {
	return c.client.EndCall(ctx, c.callType, c.callID)
}

func (c *Call) UpdateCallMembers(ctx context.Context, request *UpdateCallMembersRequest) (*UpdateCallMembersResponse, error) {
	return c.client.UpdateCallMembers(ctx, c.callType, c.callID, request)
}

func (c *Call) MuteUsers(ctx context.Context, request *MuteUsersRequest) (*MuteUsersResponse, error) {
	return c.client.MuteUsers(ctx, c.callType, c.callID, request)
}

func (c *Call) VideoPin(ctx context.Context, request *PinRequest) (*PinResponse, error) {
	return c.client.VideoPin(ctx, c.callType, c.callID, request)
}

func (c *Call) ListRecordings(ctx context.Context) (*ListRecordingsResponse, error) {
	return c.client.ListRecordings(ctx, c.callType, c.callID)
}

func (c *Call) StartHLSBroadcasting(ctx context.Context) (*StartHLSBroadcastingResponse, error) {
	return c.client.StartHLSBroadcasting(ctx, c.callType, c.callID)
}

func (c *Call) StartRecording(ctx context.Context, request *StartRecordingRequest) (*StartRecordingResponse, error) {
	return c.client.StartRecording(ctx, c.callType, c.callID, request)
}

func (c *Call) StartTranscription(ctx context.Context, request *StartTranscriptionRequest) (*StartTranscriptionResponse, error) {
	return c.client.StartTranscription(ctx, c.callType, c.callID, request)
}

func (c *Call) GetCallStats(ctx context.Context, session string) (*GetCallStatsResponse, error) {
	return c.client.GetCallStats(ctx, c.callType, c.callID, session)
}

func (c *Call) StopHLSBroadcasting(ctx context.Context) (*StopHLSBroadcastingResponse, error) {
	return c.client.StopHLSBroadcasting(ctx, c.callType, c.callID)
}

func (c *Call) StopLive(ctx context.Context) (*StopLiveResponse, error) {
	return c.client.StopLive(ctx, c.callType, c.callID)
}

func (c *Call) StopRecording(ctx context.Context) (*StopRecordingResponse, error) {
	return c.client.StopRecording(ctx, c.callType, c.callID)
}

func (c *Call) StopTranscription(ctx context.Context) (*StopTranscriptionResponse, error) {
	return c.client.StopTranscription(ctx, c.callType, c.callID)
}

func (c *Call) ListTranscriptions(ctx context.Context) (*ListTranscriptionsResponse, error) {
	return c.client.ListTranscriptions(ctx, c.callType, c.callID)
}

func (c *Call) UnblockUser(ctx context.Context, request *UnblockUserRequest) (*UnblockUserResponse, error) {
	return c.client.UnblockUser(ctx, c.callType, c.callID, request)
}

func (c *Call) VideoUnpin(ctx context.Context, request *UnpinRequest) (*UnpinResponse, error) {
	return c.client.VideoUnpin(ctx, c.callType, c.callID, request)
}

func (c *Call) UpdateUserPermissions(ctx context.Context, request *UpdateUserPermissionsRequest) (*UpdateUserPermissionsResponse, error) {
	return c.client.UpdateUserPermissions(ctx, c.callType, c.callID, request)
}

func (c *Call) DeleteRecording(ctx context.Context, session string, filename string) (*DeleteRecordingResponse, error) {
	return c.client.DeleteRecording(ctx, c.callType, c.callID, session, filename)
}

func (c *Call) DeleteTranscription(ctx context.Context, session string, filename string) (*DeleteTranscriptionResponse, error) {
	return c.client.DeleteTranscription(ctx, c.callType, c.callID, session, filename)
}

func (c *VideoClient) Call(callType, callID string) *Call {
	return NewCall(callType, callID, c)
}
