// Code generated by GetStream internal OpenAPI code generator. DO NOT EDIT.
package getstream_test

import (
	"context"
	"testing"

	"github.com/GetStream/getstream-go"
	"github.com/stretchr/testify/require"
)

func TestVideoQueryCallMembers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().QueryCallMembers(context.Background(), &getstream.QueryCallMembersRequest{})
	require.NoError(t, err)
}
func TestVideoQueryCallStats(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().QueryCallStats(context.Background(), &getstream.QueryCallStatsRequest{})
	require.NoError(t, err)
}
func TestVideoGetCall(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().GetCall(context.Background(), "", "", &getstream.GetCallRequest{})
	require.NoError(t, err)
}
func TestVideoUpdateCall(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().UpdateCall(context.Background(), "", "", &getstream.UpdateCallRequest{})
	require.NoError(t, err)
}
func TestVideoGetOrCreateCall(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().GetOrCreateCall(context.Background(), "", "", &getstream.GetOrCreateCallRequest{})
	require.NoError(t, err)
}
func TestVideoBlockUser(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().BlockUser(context.Background(), "", "", &getstream.BlockUserRequest{})
	require.NoError(t, err)
}
func TestVideoDeleteCall(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().DeleteCall(context.Background(), "", "", &getstream.DeleteCallRequest{})
	require.NoError(t, err)
}
func TestVideoSendCallEvent(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().SendCallEvent(context.Background(), "", "", &getstream.SendCallEventRequest{})
	require.NoError(t, err)
}
func TestVideoCollectUserFeedback(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().CollectUserFeedback(context.Background(), "", "", "", &getstream.CollectUserFeedbackRequest{})
	require.NoError(t, err)
}
func TestVideoGoLive(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().GoLive(context.Background(), "", "", &getstream.GoLiveRequest{})
	require.NoError(t, err)
}
func TestVideoEndCall(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().EndCall(context.Background(), "", "", &getstream.EndCallRequest{})
	require.NoError(t, err)
}
func TestVideoUpdateCallMembers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().UpdateCallMembers(context.Background(), "", "", &getstream.UpdateCallMembersRequest{})
	require.NoError(t, err)
}
func TestVideoMuteUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().MuteUsers(context.Background(), "", "", &getstream.MuteUsersRequest{})
	require.NoError(t, err)
}
func TestVideoVideoPin(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().VideoPin(context.Background(), "", "", &getstream.VideoPinRequest{})
	require.NoError(t, err)
}
func TestVideoListRecordings(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().ListRecordings(context.Background(), "", "", &getstream.ListRecordingsRequest{})
	require.NoError(t, err)
}
func TestVideoStartRTMPBroadcasts(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StartRTMPBroadcasts(context.Background(), "", "", &getstream.StartRTMPBroadcastsRequest{})
	require.NoError(t, err)
}
func TestVideoStopAllRTMPBroadcasts(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StopAllRTMPBroadcasts(context.Background(), "", "", &getstream.StopAllRTMPBroadcastsRequest{})
	require.NoError(t, err)
}
func TestVideoStopRTMPBroadcast(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StopRTMPBroadcast(context.Background(), "", "", "", &getstream.StopRTMPBroadcastRequest{})
	require.NoError(t, err)
}
func TestVideoStartHLSBroadcasting(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StartHLSBroadcasting(context.Background(), "", "", &getstream.StartHLSBroadcastingRequest{})
	require.NoError(t, err)
}
func TestVideoStartClosedCaptions(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StartClosedCaptions(context.Background(), "", "", &getstream.StartClosedCaptionsRequest{})
	require.NoError(t, err)
}
func TestVideoStartRecording(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StartRecording(context.Background(), "", "", &getstream.StartRecordingRequest{})
	require.NoError(t, err)
}
func TestVideoStartTranscription(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StartTranscription(context.Background(), "", "", &getstream.StartTranscriptionRequest{})
	require.NoError(t, err)
}
func TestVideoGetCallStats(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().GetCallStats(context.Background(), "", "", "", &getstream.GetCallStatsRequest{})
	require.NoError(t, err)
}
func TestVideoStopHLSBroadcasting(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StopHLSBroadcasting(context.Background(), "", "", &getstream.StopHLSBroadcastingRequest{})
	require.NoError(t, err)
}
func TestVideoStopClosedCaptions(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StopClosedCaptions(context.Background(), "", "", &getstream.StopClosedCaptionsRequest{})
	require.NoError(t, err)
}
func TestVideoStopLive(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StopLive(context.Background(), "", "", &getstream.StopLiveRequest{})
	require.NoError(t, err)
}
func TestVideoStopRecording(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StopRecording(context.Background(), "", "", &getstream.StopRecordingRequest{})
	require.NoError(t, err)
}
func TestVideoStopTranscription(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().StopTranscription(context.Background(), "", "", &getstream.StopTranscriptionRequest{})
	require.NoError(t, err)
}
func TestVideoListTranscriptions(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().ListTranscriptions(context.Background(), "", "", &getstream.ListTranscriptionsRequest{})
	require.NoError(t, err)
}
func TestVideoUnblockUser(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().UnblockUser(context.Background(), "", "", &getstream.UnblockUserRequest{})
	require.NoError(t, err)
}
func TestVideoVideoUnpin(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().VideoUnpin(context.Background(), "", "", &getstream.VideoUnpinRequest{})
	require.NoError(t, err)
}
func TestVideoUpdateUserPermissions(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().UpdateUserPermissions(context.Background(), "", "", &getstream.UpdateUserPermissionsRequest{})
	require.NoError(t, err)
}
func TestVideoDeleteRecording(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().DeleteRecording(context.Background(), "", "", "", "", &getstream.DeleteRecordingRequest{})
	require.NoError(t, err)
}
func TestVideoDeleteTranscription(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().DeleteTranscription(context.Background(), "", "", "", "", &getstream.DeleteTranscriptionRequest{})
	require.NoError(t, err)
}
func TestVideoQueryCalls(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().QueryCalls(context.Background(), &getstream.QueryCallsRequest{})
	require.NoError(t, err)
}
func TestVideoListCallTypes(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().ListCallTypes(context.Background(), &getstream.ListCallTypesRequest{})
	require.NoError(t, err)
}
func TestVideoCreateCallType(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().CreateCallType(context.Background(), &getstream.CreateCallTypeRequest{})
	require.NoError(t, err)
}
func TestVideoDeleteCallType(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().DeleteCallType(context.Background(), "", &getstream.DeleteCallTypeRequest{})
	require.NoError(t, err)
}
func TestVideoGetCallType(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().GetCallType(context.Background(), "", &getstream.GetCallTypeRequest{})
	require.NoError(t, err)
}
func TestVideoUpdateCallType(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().UpdateCallType(context.Background(), "", &getstream.UpdateCallTypeRequest{})
	require.NoError(t, err)
}
func TestVideoGetEdges(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().GetEdges(context.Background(), &getstream.GetEdgesRequest{})
	require.NoError(t, err)
}
func TestVideoQueryAggregateCallStats(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.Video().QueryAggregateCallStats(context.Background(), &getstream.QueryAggregateCallStatsRequest{})
	require.NoError(t, err)
}