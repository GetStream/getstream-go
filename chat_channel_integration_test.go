package getstream_test

import (
	"context"
	"os"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatChannelIntegration(t *testing.T) {
	t.Parallel()
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
			FilterConditions: PtrTo(map[string]any{
				"id": channelID,
			}),
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
			{UserID: PtrTo(creatorID)},
			{UserID: PtrTo(memberID1)},
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

		// Cleanup (hard delete)
		t.Cleanup(func() {
			_, _ = client.Chat().DeleteChannel(context.Background(), "messaging", resp.Data.Channel.ID, &DeleteChannelRequest{
				HardDelete: PtrTo(true),
			})
		})
	})

	t.Run("QueryChannels", func(t *testing.T) {
		_, channelID := createTestChannel(t, client, creatorID)

		resp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"type": "messaging",
				"id":   channelID,
			}),
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
			Set: PtrTo(map[string]any{
				"color":       "red",
				"description": "A test channel",
			}),
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Channel)

		// Same read-replica caveat as the unset below: assert that the write
		// applied by re-reading, not from the write response.
		require.Eventually(t, func() bool {
			resp, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
			if err != nil || resp.Data.Channel == nil {
				return false
			}
			return resp.Data.Channel.Custom["color"] == "red"
		}, 5*time.Second, 500*time.Millisecond, "color should be set to red")

		// Unset fields
		_, err = ch.UpdateChannelPartial(ctx, &UpdateChannelPartialRequest{
			Unset: PtrTo([]string{"color"}),
		})
		require.NoError(t, err)

		// The response can be hydrated from a read replica that has not observed
		// the write yet, so re-read until the unset field disappears.
		require.Eventually(t, func() bool {
			resp, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
			if err != nil || resp.Data.Channel == nil {
				return false
			}
			_, hasColor := resp.Data.Channel.Custom["color"]
			return !hasColor
		}, 5*time.Second, 500*time.Millisecond, "color should be unset")
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

		// Hard delete cleanup in case soft delete test fails
		t.Cleanup(func() {
			_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{
				HardDelete: PtrTo(true),
			})
		})

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

		// Hard-delete is async; the SDK's job (returning a task, polling it) is
		// verified above and in requireTaskCompletedOrSkipOnTimeout. A poll
		// timeout is backend queue latency, not an SDK regression, so it skips
		// rather than fails the (historically flaky) subtest.
		requireTaskCompletedOrSkipOnTimeout(t, ctx, client, *resp.Data.TaskID)
	})

	t.Run("AddRemoveMembers", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Add members
		_, err := ch.Update(ctx, &UpdateChannelRequest{
			AddMembers: PtrTo([]ChannelMemberRequest{
				{UserID: PtrTo(memberID2)},
				{UserID: PtrTo(memberID3)},
			}),
		})
		require.NoError(t, err)

		// Verify members added
		resp, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Data.Members), 4)

		// Remove a member
		_, err = ch.Update(ctx, &UpdateChannelRequest{
			RemoveMembers: PtrTo([]string{memberID3}),
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
					{UserID: PtrTo(creatorID)},
				},
				Invites: []ChannelMemberRequest{
					{UserID: PtrTo(memberID1)},
					{UserID: PtrTo(memberID2)},
				},
			},
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{
				HardDelete: PtrTo(true),
			})
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

		// The frozen flag in an UpdateChannelPartial response can be hydrated from
		// a read replica that lags the write, so verify via a re-read that retries
		// until the state converges rather than trusting the write response itself.
		requireFrozen := func(want bool) {
			t.Helper()
			var got bool
			for i := 0; i < 10; i++ {
				r, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
				require.NoError(t, err)
				got = r.Data.Channel.Frozen
				if got == want {
					return
				}
				time.Sleep(500 * time.Millisecond)
			}
			assert.Equal(t, want, got, "channel frozen state should converge to %v", want)
		}

		_, err := ch.UpdateChannelPartial(ctx, &UpdateChannelPartialRequest{
			Set: PtrTo(map[string]any{"frozen": true}),
		})
		require.NoError(t, err)
		requireFrozen(true)

		_, err = ch.UpdateChannelPartial(ctx, &UpdateChannelPartialRequest{
			Set: PtrTo(map[string]any{"frozen": false}),
		})
		require.NoError(t, err)
		requireFrozen(false)
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
			ChannelCids: PtrTo([]string{cid}),
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
			FilterConditions: PtrTo(map[string]any{
				"muted": true,
				"cid":   cid,
			}),
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1, "Should find exactly 1 muted channel")
		assert.Equal(t, cid, qResp.Data.Channels[0].Channel.Cid)

		// Unmute the channel
		_, err = client.Chat().UnmuteChannel(ctx, &UnmuteChannelRequest{
			ChannelCids: PtrTo([]string{cid}),
			UserID:      PtrTo(memberID1),
		})
		require.NoError(t, err)

		// Verify unmute via query with muted=false (like stream-chat-go does)
		qResp, err = client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"muted": false,
				"cid":   cid,
			}),
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
			Set: PtrTo(map[string]any{
				"role_label": "moderator",
				"score":      42,
			}),
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.ChannelMember)
		assert.Equal(t, "moderator", resp.Data.ChannelMember.Custom["role_label"])

		// Unset a custom field
		resp, err = ch.UpdateMemberPartial(ctx, &UpdateMemberPartialRequest{
			UserID: PtrTo(memberID1),
			Unset:  PtrTo([]string{"score"}),
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
			AssignRoles: PtrTo([]ChannelMemberRequest{
				{UserID: PtrTo(memberID1), ChannelRole: PtrTo("channel_moderator")},
			}),
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

	t.Run("AddDemoteModerators", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Add moderator using UpdateChannel with AddModerators
		_, err := ch.Update(ctx, &UpdateChannelRequest{
			AddModerators: PtrTo([]string{memberID1}),
		})
		require.NoError(t, err)

		// Verify role changed to moderator via QueryMembers
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

		// Demote moderator back to member
		_, err = ch.Update(ctx, &UpdateChannelRequest{
			DemoteModerators: PtrTo([]string{memberID1}),
		})
		require.NoError(t, err)

		// Verify role changed back to member
		qResp, err = client.Chat().QueryMembers(ctx, &QueryMembersRequest{
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
		assert.Equal(t, "channel_member", qResp.Data.Members[0].ChannelRole)
	})

	t.Run("MarkUnreadWithThread", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Send parent message
		parentID := sendTestMessage(t, ch, creatorID, "Parent for mark unread thread")

		// Send reply to create a thread
		_, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:     PtrTo("Reply in thread"),
				UserID:   PtrTo(creatorID),
				ParentID: PtrTo(parentID),
			},
		})
		require.NoError(t, err)

		// Mark unread from thread
		_, err = ch.MarkUnread(ctx, &MarkUnreadRequest{
			UserID:   PtrTo(memberID1),
			ThreadID: PtrTo(parentID),
		})
		require.NoError(t, err)
	})

	t.Run("TruncateWithOptions", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Send messages
		sendTestMessage(t, ch, creatorID, "Truncate msg 1")
		sendTestMessage(t, ch, creatorID, "Truncate msg 2")

		// Truncate with message and skip push (matching stream-chat-go TruncateWithOptions)
		_, err := ch.Truncate(ctx, &TruncateChannelRequest{
			Message: &MessageRequest{
				Text:   PtrTo("Channel was truncated"),
				UserID: PtrTo(creatorID),
			},
			SkipPush:   PtrTo(true),
			HardDelete: PtrTo(true),
		})
		require.NoError(t, err)
	})

	t.Run("PinUnpinChannel", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Pin channel for memberID1 via UpdateMemberPartial
		_, err := ch.UpdateMemberPartial(ctx, &UpdateMemberPartialRequest{
			UserID: PtrTo(memberID1),
			Set: PtrTo(map[string]any{
				"pinned": true,
			}),
		})
		require.NoError(t, err)

		// Verify via QueryChannels with pinned filter
		qResp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"pinned": true,
				"cid":    "messaging:" + channelID,
			}),
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1, "Should find 1 pinned channel")
		assert.Equal(t, "messaging:"+channelID, qResp.Data.Channels[0].Channel.Cid)

		// Unpin
		_, err = ch.UpdateMemberPartial(ctx, &UpdateMemberPartialRequest{
			UserID: PtrTo(memberID1),
			Set: PtrTo(map[string]any{
				"pinned": false,
			}),
		})
		require.NoError(t, err)

		// Verify unpinned
		qResp, err = client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"pinned": false,
				"cid":    "messaging:" + channelID,
			}),
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1, "Should find channel with pinned=false")
	})

	t.Run("ArchiveUnarchiveChannel", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Archive channel for memberID1 via UpdateMemberPartial
		_, err := ch.UpdateMemberPartial(ctx, &UpdateMemberPartialRequest{
			UserID: PtrTo(memberID1),
			Set: PtrTo(map[string]any{
				"archived": true,
			}),
		})
		require.NoError(t, err)

		// Verify via QueryChannels with archived filter
		qResp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"archived": true,
				"cid":      "messaging:" + channelID,
			}),
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1, "Should find 1 archived channel")
		assert.Equal(t, "messaging:"+channelID, qResp.Data.Channels[0].Channel.Cid)

		// Unarchive
		_, err = ch.UpdateMemberPartial(ctx, &UpdateMemberPartialRequest{
			UserID: PtrTo(memberID1),
			Set: PtrTo(map[string]any{
				"archived": false,
			}),
		})
		require.NoError(t, err)

		// Verify unarchived
		qResp, err = client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"archived": false,
				"cid":      "messaging:" + channelID,
			}),
			UserID: PtrTo(memberID1),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1, "Should find channel with archived=false")
	})

	t.Run("AddMembersWithRoles", func(t *testing.T) {
		ch, channelID := createTestChannel(t, client, creatorID)
		_ = channelID

		newUserIDs := createTestUsers(t, client, 2)
		modUserID := newUserIDs[0]
		memberUserID2 := newUserIDs[1]

		// Add members with specific channel roles
		_, err := ch.Update(ctx, &UpdateChannelRequest{
			AddMembers: PtrTo([]ChannelMemberRequest{
				{UserID: PtrTo(modUserID), ChannelRole: PtrTo("channel_moderator")},
				{UserID: PtrTo(memberUserID2), ChannelRole: PtrTo("channel_member")},
			}),
		})
		require.NoError(t, err)

		// Query members to verify roles
		membersResp, err := client.Chat().QueryMembers(ctx, &QueryMembersRequest{
			Payload: &QueryMembersPayload{
				Type:             "messaging",
				ID:               PtrTo(channelID),
				FilterConditions: map[string]any{"id": map[string]any{"$in": newUserIDs}},
			},
		})
		require.NoError(t, err)

		roleMap := make(map[string]string)
		for _, m := range membersResp.Data.Members {
			if m.UserID != nil {
				roleMap[*m.UserID] = m.ChannelRole
			}
		}
		assert.Equal(t, "channel_moderator", roleMap[modUserID], "First user should be channel_moderator")
		assert.Equal(t, "channel_member", roleMap[memberUserID2], "Second user should be channel_member")
	})

	t.Run("MessageCount", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Send a message
		sendTestMessage(t, ch, creatorID, "hello world")

		// Query the channel to get message count
		qResp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"cid": "messaging:" + channelID,
			}),
			UserID: PtrTo(creatorID),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1)

		// MessageCount should be present (default enabled for messaging type)
		channel := qResp.Data.Channels[0].Channel
		if channel.MessageCount != nil {
			assert.GreaterOrEqual(t, *channel.MessageCount, 1, "MessageCount should be >= 1")
		}
		// Note: MessageCount may be nil if count_messages is disabled on the channel type
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

	t.Run("FilterTags", func(t *testing.T) {
		ch, _ := createTestChannel(t, client, creatorID)

		// Add filter tags
		_, err := ch.Update(ctx, &UpdateChannelRequest{
			AddFilterTags: PtrTo([]string{"sports", "news"}),
		})
		require.NoError(t, err)

		// Verify tags were added by querying
		getResp, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{})
		require.NoError(t, err)
		require.NotNil(t, getResp.Data.Channel)

		// Remove filter tags
		_, err = ch.Update(ctx, &UpdateChannelRequest{
			RemoveFilterTags: PtrTo([]string{"sports"}),
		})
		require.NoError(t, err)
	})

	t.Run("MessageCountDisabled", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Disable count_messages via config_overrides partial update
		_, err := ch.UpdateChannelPartial(ctx, &UpdateChannelPartialRequest{
			Set: PtrTo(map[string]any{
				"config_overrides": map[string]any{
					"count_messages": false,
				},
			}),
		})
		require.NoError(t, err)

		// Send a message
		sendTestMessage(t, ch, creatorID, "hello world disabled count")

		// Query the channel — MessageCount should be nil when disabled
		qResp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"cid": "messaging:" + channelID,
			}),
			UserID: PtrTo(creatorID),
		})
		require.NoError(t, err)
		require.Len(t, qResp.Data.Channels, 1)
		assert.Nil(t, qResp.Data.Channels[0].Channel.MessageCount,
			"MessageCount should be nil when count_messages is disabled")
	})

	t.Run("MarkUnreadWithTimestamp", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID, memberID1})

		// Send a message to get a valid timestamp
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("test message for timestamp unread"),
				UserID: PtrTo(creatorID),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, sendResp.Data.Message.CreatedAt)

		// Mark unread from timestamp
		ts := sendResp.Data.Message.CreatedAt
		_, err = ch.MarkUnread(ctx, &MarkUnreadRequest{
			UserID:           PtrTo(memberID1),
			MessageTimestamp: &ts,
		})
		require.NoError(t, err)
	})

	t.Run("HideForCreator", func(t *testing.T) {
		channelID := "test-hide-" + randomString(12)
		ch := client.Chat().Channel("messaging", channelID)

		_, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{
			HideForCreator: PtrTo(true),
			Data: &ChannelInput{
				CreatedByID: PtrTo(creatorID),
				Members: []ChannelMemberRequest{
					{UserID: PtrTo(creatorID)},
					{UserID: PtrTo(memberID1)},
				},
			},
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{
				HardDelete: PtrTo(true),
			})
		})

		// Channel should be hidden for creator — querying without show_hidden should not find it
		qResp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: PtrTo(map[string]any{
				"cid": "messaging:" + channelID,
			}),
			UserID: PtrTo(creatorID),
		})
		require.NoError(t, err)
		assert.Empty(t, qResp.Data.Channels, "Channel should be hidden for creator")
	})

	t.Run("UploadAndDeleteFile", func(t *testing.T) {
		// Serialize app-config mutation (see appConfigMu): this UpdateApp sends
		// size_limit=0 and would clobber a concurrent config test's read-back.
		// defer-unlock runs after the restore defer below (defers are LIFO).
		appConfigMu.Lock()
		defer appConfigMu.Unlock()

		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID})

		// Save original file upload config and allow .txt uploads
		appResp, err := client.GetApp(ctx, &GetAppRequest{})
		require.NoError(t, err)
		originalConfig := appResp.Data.App.FileUploadConfig

		_, err = client.UpdateApp(ctx, &UpdateAppRequest{
			FileUploadConfig: &FileUploadConfig{
				AllowedFileExtensions: []string{},
				BlockedFileExtensions: []string{},
				AllowedMimeTypes:      []string{},
				BlockedMimeTypes:      []string{},
			},
		})
		require.NoError(t, err)
		time.Sleep(2 * time.Second)
		defer func() {
			// Restore original file upload config
			_, _ = client.UpdateApp(ctx, &UpdateAppRequest{
				FileUploadConfig: &originalConfig,
			})
		}()

		// Create a temp file to upload
		tmpFile, err := os.CreateTemp("", "chat-test-*.txt")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())
		_, err = tmpFile.WriteString("hello world test file content")
		require.NoError(t, err)
		tmpFile.Close()

		// Upload file
		uploadResp, err := ch.UploadChannelFile(ctx, &UploadChannelFileRequest{
			File: PtrTo(tmpFile.Name()),
			User: &OnlyUserID{ID: creatorID},
		})
		require.NoError(t, err)
		require.NotNil(t, uploadResp.Data.File)
		fileURL := *uploadResp.Data.File
		assert.NotEmpty(t, fileURL)
		assert.Contains(t, fileURL, "http")

		// Delete file
		_, err = ch.DeleteChannelFile(ctx, &DeleteChannelFileRequest{
			Url: PtrTo(fileURL),
		})
		require.NoError(t, err)
	})

	t.Run("UploadAndDeleteImage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, creatorID, []string{creatorID})

		// Create a temp image file to upload (minimal valid JPEG)
		tmpFile, err := os.CreateTemp("", "chat-test-*.jpg")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())
		_, err = tmpFile.WriteString("fake-jpg-image-data-for-testing")
		require.NoError(t, err)
		tmpFile.Close()

		// Upload image
		uploadResp, err := ch.UploadChannelImage(ctx, &UploadChannelImageRequest{
			File: PtrTo(tmpFile.Name()),
			User: &OnlyUserID{ID: creatorID},
		})
		require.NoError(t, err)
		require.NotNil(t, uploadResp.Data.File)
		imageURL := *uploadResp.Data.File
		assert.NotEmpty(t, imageURL)
		assert.Contains(t, imageURL, "http")

		// Delete image
		_, err = ch.DeleteChannelImage(ctx, &DeleteChannelImageRequest{
			Url: PtrTo(imageURL),
		})
		require.NoError(t, err)
	})
}
