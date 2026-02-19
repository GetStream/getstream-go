package getstream_test

import (
	"context"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatChannelIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	// Create shared test users for all subtests
	userIDs := createTestUsers(t, client, 4)
	creatorID := userIDs[0]
	memberID1 := userIDs[1]
	memberID2 := userIDs[2]
	memberID3 := userIDs[3]

	t.Run("CreateChannelWithID", func(t *testing.T) {
		ch, channelID := createTestChannel(t, client, creatorID)
		_ = ch

		// Verify channel exists by querying
		resp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: map[string]any{
				"id": channelID,
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, resp.Data.Channels)
		assert.Equal(t, channelID, resp.Data.Channels[0].Channel.ID)
		assert.Equal(t, "messaging", resp.Data.Channels[0].Channel.Type)
	})

	t.Run("CreateChannelWithMembers", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1, memberID2})

		resp, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Data.Members), 3)
	})

	t.Run("CreateDistinctChannel", func(t *testing.T) {
		members := []ChannelMemberRequest{
			{UserID: creatorID},
			{UserID: memberID1},
		}

		resp, err := client.Chat().GetOrCreateDistinctChannel(ctx, "messaging", &GetOrCreateDistinctChannelRequest{
			Data: &ChannelInput{
				CreatedByID: PtrTo(creatorID),
				Members:     members,
			},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Channel)

		// Calling again with same members should return same channel
		resp2, err := client.Chat().GetOrCreateDistinctChannel(ctx, "messaging", &GetOrCreateDistinctChannelRequest{
			Data: &ChannelInput{
				CreatedByID: PtrTo(creatorID),
				Members:     members,
			},
		})
		require.NoError(t, err)
		assert.Equal(t, resp.Data.Channel.Cid, resp2.Data.Channel.Cid)

		// Cleanup
		t.Cleanup(func() {
			_, _ = client.Chat().DeleteChannel(context.Background(), "messaging", resp.Data.Channel.ID, &DeleteChannelRequest{})
		})
	})

	t.Run("QueryChannels", func(t *testing.T) {
		_, channelID := createTestChannel(t, client, creatorID)

		resp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: map[string]any{
				"type": "messaging",
				"id":   channelID,
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, resp.Data.Channels)
		assert.Equal(t, channelID, resp.Data.Channels[0].Channel.ID)
	})

	t.Run("UpdateChannel", func(t *testing.T) {
		ch, _ := createTestChannel(t, client, creatorID)

		resp, err := ch.Update(ctx, &UpdateChannelRequest{
			Data: &ChannelInputRequest{
				Custom: map[string]any{
					"color": "blue",
				},
			},
			Message: &MessageRequest{
				Text:   PtrTo("Channel updated!"),
				UserID: PtrTo(creatorID),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Channel)
		assert.Equal(t, "blue", resp.Data.Channel.Custom["color"])
	})

	t.Run("PartialUpdateChannel", func(t *testing.T) {
		ch, _ := createTestChannel(t, client, creatorID)

		// Set fields
		resp, err := ch.UpdateChannelPartial(ctx, &UpdateChannelPartialRequest{
			Set: map[string]any{
				"color":       "red",
				"description": "A test channel",
			},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Channel)
		assert.Equal(t, "red", resp.Data.Channel.Custom["color"])

		// Unset fields
		resp, err = ch.UpdateChannelPartial(ctx, &UpdateChannelPartialRequest{
			Unset: []string{"color"},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Channel)
		_, hasColor := resp.Data.Channel.Custom["color"]
		assert.False(t, hasColor, "color should be unset")
	})

	t.Run("DeleteChannel", func(t *testing.T) {
		channelID := "test-del-" + randomString(12)
		ch := client.Chat().Channel("messaging", channelID)

		_, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{
			Data: &ChannelInput{
				CreatedByID: PtrTo(creatorID),
			},
		})
		require.NoError(t, err)

		// Soft delete
		resp, err := ch.Delete(ctx, &DeleteChannelRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Channel)
	})

	t.Run("HardDeleteChannels", func(t *testing.T) {
		ch1, channelID1 := createTestChannel(t, client, creatorID)
		ch2, channelID2 := createTestChannel(t, client, creatorID)
		_ = ch1
		_ = ch2

		cid1 := "messaging:" + channelID1
		cid2 := "messaging:" + channelID2

		resp, err := client.Chat().DeleteChannels(ctx, &DeleteChannelsRequest{
			Cids:       []string{cid1, cid2},
			HardDelete: PtrTo(true),
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.TaskID)

		taskCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		taskResult, err := WaitForTask(taskCtx, client, *resp.Data.TaskID)
		require.NoError(t, err)
		assert.Equal(t, "completed", taskResult.Data.Status)
	})

	t.Run("AddRemoveMembers", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Add members
		_, err := ch.Update(ctx, &UpdateChannelRequest{
			AddMembers: []ChannelMemberRequest{
				{UserID: memberID2},
				{UserID: memberID3},
			},
		})
		require.NoError(t, err)

		// Verify members added
		resp, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Data.Members), 4)

		// Remove a member
		_, err = ch.Update(ctx, &UpdateChannelRequest{
			RemoveMembers: []string{memberID3},
		})
		require.NoError(t, err)

		// Verify member removed
		resp, err = ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
		require.NoError(t, err)
		memberFound := false
		for _, m := range resp.Data.Members {
			if m.UserID != nil && *m.UserID == memberID3 {
				memberFound = true
			}
		}
		assert.False(t, memberFound, "memberID3 should have been removed")
	})

	t.Run("QueryMembers", func(t *testing.T) {
		_, channelID := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1, memberID2})

		resp, err := client.Chat().QueryMembers(ctx, &QueryMembersRequest{
			Payload: &QueryMembersPayload{
				Type:             "messaging",
				ID:               PtrTo(channelID),
				FilterConditions: map[string]any{},
			},
		})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Data.Members), 3)
	})

	t.Run("InviteAcceptReject", func(t *testing.T) {
		channelID := "test-inv-" + randomString(12)
		ch := client.Chat().Channel("messaging", channelID)

		// Create channel with invite
		_, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{
			Data: &ChannelInput{
				CreatedByID: PtrTo(creatorID),
				Members: []ChannelMemberRequest{
					{UserID: creatorID},
				},
				Invites: []ChannelMemberRequest{
					{UserID: memberID1},
					{UserID: memberID2},
				},
			},
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{})
		})

		// Accept invite
		_, err = ch.Update(ctx, &UpdateChannelRequest{
			AcceptInvite: PtrTo(true),
			UserID:       PtrTo(memberID1),
		})
		require.NoError(t, err)

		// Reject invite
		_, err = ch.Update(ctx, &UpdateChannelRequest{
			RejectInvite: PtrTo(true),
			UserID:       PtrTo(memberID2),
		})
		require.NoError(t, err)
	})

	t.Run("HideShowChannel", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Hide
		_, err := ch.Hide(ctx, &HideChannelRequest{
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)

		// Show
		_, err = ch.Show(ctx, &ShowChannelRequest{
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)
	})

	t.Run("TruncateChannel", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Send some messages
		sendTestMessage(t, ch, creatorID, "Message 1")
		sendTestMessage(t, ch, creatorID, "Message 2")
		sendTestMessage(t, ch, creatorID, "Message 3")

		// Truncate
		_, err := ch.Truncate(ctx, &TruncateChannelRequest{})
		require.NoError(t, err)

		// Verify messages are gone
		resp, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
		require.NoError(t, err)
		assert.Empty(t, resp.Data.Messages, "Messages should be empty after truncation")
	})

	t.Run("FreezeUnfreezeChannel", func(t *testing.T) {
		ch, _ := createTestChannel(t, client, creatorID)

		// Freeze
		resp, err := ch.UpdateChannelPartial(ctx, &UpdateChannelPartialRequest{
			Set: map[string]any{
				"frozen": true,
			},
		})
		require.NoError(t, err)
		assert.True(t, resp.Data.Channel.Frozen)

		// Unfreeze
		resp, err = ch.UpdateChannelPartial(ctx, &UpdateChannelPartialRequest{
			Set: map[string]any{
				"frozen": false,
			},
		})
		require.NoError(t, err)
		assert.False(t, resp.Data.Channel.Frozen)
	})

	t.Run("MarkReadUnread", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Send a message
		msgID := sendTestMessage(t, ch, creatorID, "Message to mark read")

		// Mark read
		_, err := ch.MarkRead(ctx, &MarkReadRequest{
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)

		// Mark unread from this message
		_, err = ch.MarkUnread(ctx, &MarkUnreadRequest{
			UserID:    PtrTo(memberID1),
			MessageID: PtrTo(msgID),
		})
		require.NoError(t, err)
	})

	t.Run("MuteUnmuteChannel", func(t *testing.T) {
		_, channelID := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})
		cid := "messaging:" + channelID

		// Mute the channel for memberID1
		muteResp, err := client.Chat().MuteChannel(ctx, &MuteChannelRequest{
			ChannelCids: []string{cid},
			UserID:      PtrTo(memberID1),
		})
		require.NoError(t, err)
		require.NotNil(t, muteResp.Data)

		// Verify mute response details (like stream-chat-go does)
		require.NotNil(t, muteResp.Data.ChannelMute, "Mute response should contain ChannelMute")
		require.NotNil(t, muteResp.Data.ChannelMute.Channel, "ChannelMute should have Channel")
		require.NotNil(t, muteResp.Data.ChannelMute.User, "ChannelMute should have User")
		assert.Equal(t, cid, muteResp.Data.ChannelMute.Channel.Cid)
		assert.Equal(t, memberID1, muteResp.Data.ChannelMute.User.ID)

		// Verify via QueryChannels with muted=true and cid filter
		qResp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: map[string]any{
				"muted": true,
				"cid":   cid,
			},
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1, "Should find exactly 1 muted channel")
		assert.Equal(t, cid, qResp.Data.Channels[0].Channel.Cid)

		// Unmute the channel
		_, err = client.Chat().UnmuteChannel(ctx, &UnmuteChannelRequest{
			ChannelCids: []string{cid},
			UserID:      PtrTo(memberID1),
		})
		require.NoError(t, err)

		// Verify unmute via query with muted=false (like stream-chat-go does)
		qResp, err = client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: map[string]any{
				"muted": false,
				"cid":   cid,
			},
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1, "Unmuted channel should appear in muted=false query")
	})

	t.Run("MemberPartialUpdate", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Set custom fields on member
		resp, err := ch.UpdateMemberPartial(ctx, &UpdateMemberPartialRequest{
			UserID: PtrTo(memberID1),
			Set: map[string]any{
				"role_label": "moderator",
				"score":      42,
			},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.ChannelMember)
		assert.Equal(t, "moderator", resp.Data.ChannelMember.Custom["role_label"])

		// Unset a custom field
		resp, err = ch.UpdateMemberPartial(ctx, &UpdateMemberPartialRequest{
			UserID: PtrTo(memberID1),
			Unset:  []string{"score"},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.ChannelMember)
		_, hasScore := resp.Data.ChannelMember.Custom["score"]
		assert.False(t, hasScore, "score should be unset")
	})

	t.Run("AssignRoles", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Assign channel_moderator role to memberID1
		_, err := ch.Update(ctx, &UpdateChannelRequest{
			AssignRoles: []ChannelMemberRequest{
				{UserID: memberID1, ChannelRole: PtrTo("channel_moderator")},
			},
		})
		require.NoError(t, err)

		// Verify via QueryMembers that the role is set
		qResp, err := client.Chat().QueryMembers(ctx, &QueryMembersRequest{
			Payload: &QueryMembersPayload{
				Type: "messaging",
				ID:   PtrTo(channelID),
				FilterConditions: map[string]any{
					"id": memberID1,
				},
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, qResp.Data.Members)
		assert.Equal(t, "channel_moderator", qResp.Data.Members[0].ChannelRole)
	})

	t.Run("SendChannelEvent", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		_, err := ch.SendEvent(ctx, &SendEventRequest{
			Event: EventRequest{
				Type:   "typing.start",
				UserID: PtrTo(creatorID),
			},
		})
		require.NoError(t, err)
	})
}
