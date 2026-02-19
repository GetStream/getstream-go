package getstream_test

import (
	"context"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatDeviceIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 1)
	userID := userIDs[0]

	t.Run("CreateListDeleteDevice", func(t *testing.T) {
		deviceID := "integration-test-device-" + randomString(12)

		// Create device
		_, err := client.CreateDevice(ctx, &CreateDeviceRequest{
			ID:           deviceID,
			PushProvider: "firebase",
			UserID:       PtrTo(userID),
		})
		require.NoError(t, err)

		// List devices
		listResp, err := client.ListDevices(ctx, &ListDevicesRequest{
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)

		found := false
		for _, d := range listResp.Data.Devices {
			if d.ID == deviceID {
				found = true
				assert.Equal(t, "firebase", d.PushProvider)
				assert.Equal(t, userID, d.UserID)
			}
		}
		assert.True(t, found, "Created device should appear in list")

		// Delete device
		_, err = client.DeleteDevice(ctx, &DeleteDeviceRequest{
			ID:     deviceID,
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)

		// Verify deleted
		listResp, err = client.ListDevices(ctx, &ListDevicesRequest{
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)

		for _, d := range listResp.Data.Devices {
			assert.NotEqual(t, deviceID, d.ID, "Device should be deleted")
		}
	})
}

func TestChatBlocklistIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	blocklistName := "test-blocklist-" + randomString(8)

	t.Cleanup(func() {
		_, _ = client.DeleteBlockList(context.Background(), blocklistName, &DeleteBlockListRequest{})
	})

	t.Run("CreateBlockList", func(t *testing.T) {
		resp, err := client.CreateBlockList(ctx, &CreateBlockListRequest{
			Name:  blocklistName,
			Words: []string{"badword1", "badword2", "badword3"},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data)
	})

	t.Run("GetBlockList", func(t *testing.T) {
		resp, err := client.GetBlockList(ctx, blocklistName, &GetBlockListRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Blocklist)
		assert.Equal(t, blocklistName, resp.Data.Blocklist.Name)
		assert.Len(t, resp.Data.Blocklist.Words, 3)
	})

	t.Run("UpdateBlockList", func(t *testing.T) {
		_, err := client.UpdateBlockList(ctx, blocklistName, &UpdateBlockListRequest{
			Words: []string{"badword1", "badword2", "badword3", "badword4"},
		})
		require.NoError(t, err)

		// Verify update
		resp, err := client.GetBlockList(ctx, blocklistName, &GetBlockListRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Blocklist)
		assert.Len(t, resp.Data.Blocklist.Words, 4)
	})

	t.Run("ListBlockLists", func(t *testing.T) {
		resp, err := client.ListBlockLists(ctx, &ListBlockListsRequest{})
		require.NoError(t, err)

		found := false
		for _, bl := range resp.Data.Blocklists {
			if bl.Name == blocklistName {
				found = true
			}
		}
		assert.True(t, found, "Created blocklist should appear in list")
	})

	t.Run("DeleteBlockList", func(t *testing.T) {
		// Create a separate one to delete
		deleteName := "test-del-bl-" + randomString(8)
		_, err := client.CreateBlockList(ctx, &CreateBlockListRequest{
			Name:  deleteName,
			Words: []string{"word1"},
		})
		require.NoError(t, err)

		_, err = client.DeleteBlockList(ctx, deleteName, &DeleteBlockListRequest{})
		require.NoError(t, err)
	})
}

func TestChatCommandIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	cmdName := "testcmd" + randomString(6)

	t.Cleanup(func() {
		_, _ = client.Chat().DeleteCommand(context.Background(), cmdName, &DeleteCommandRequest{})
	})

	t.Run("CreateCommand", func(t *testing.T) {
		resp, err := client.Chat().CreateCommand(ctx, &CreateCommandRequest{
			Name:        cmdName,
			Description: "A test command",
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Command)
		assert.Equal(t, cmdName, resp.Data.Command.Name)
		assert.Equal(t, "A test command", resp.Data.Command.Description)
	})

	t.Run("GetCommand", func(t *testing.T) {
		resp, err := client.Chat().GetCommand(ctx, cmdName, &GetCommandRequest{})
		require.NoError(t, err)
		assert.Equal(t, cmdName, resp.Data.Name)
		assert.Equal(t, "A test command", resp.Data.Description)
	})

	t.Run("UpdateCommand", func(t *testing.T) {
		_, err := client.Chat().UpdateCommand(ctx, cmdName, &UpdateCommandRequest{
			Description: "Updated test command",
		})
		require.NoError(t, err)

		// Verify update
		resp, err := client.Chat().GetCommand(ctx, cmdName, &GetCommandRequest{})
		require.NoError(t, err)
		assert.Equal(t, "Updated test command", resp.Data.Description)
	})

	t.Run("ListCommands", func(t *testing.T) {
		resp, err := client.Chat().ListCommands(ctx, &ListCommandsRequest{})
		require.NoError(t, err)

		found := false
		for _, cmd := range resp.Data.Commands {
			if cmd.Name == cmdName {
				found = true
			}
		}
		assert.True(t, found, "Created command should appear in list")
	})

	t.Run("DeleteCommand", func(t *testing.T) {
		delName := "testdelcmd" + randomString(6)
		_, err := client.Chat().CreateCommand(ctx, &CreateCommandRequest{
			Name:        delName,
			Description: "Command to delete",
		})
		require.NoError(t, err)

		resp, err := client.Chat().DeleteCommand(ctx, delName, &DeleteCommandRequest{})
		require.NoError(t, err)
		assert.Equal(t, delName, resp.Data.Name)
	})
}

func TestChatChannelTypeIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	typeName := "testtype" + randomString(6)

	t.Cleanup(func() {
		_, _ = client.Chat().DeleteChannelType(context.Background(), typeName, &DeleteChannelTypeRequest{})
	})

	t.Run("CreateChannelType", func(t *testing.T) {
		resp, err := client.Chat().CreateChannelType(ctx, &CreateChannelTypeRequest{
			Name:            typeName,
			Automod:         "disabled",
			AutomodBehavior: "flag",
			MaxMessageLength: 5000,
		})
		require.NoError(t, err)
		assert.Equal(t, typeName, resp.Data.Name)
		assert.Equal(t, 5000, resp.Data.MaxMessageLength)
	})

	t.Run("GetChannelType", func(t *testing.T) {
		resp, err := client.Chat().GetChannelType(ctx, typeName, &GetChannelTypeRequest{})
		require.NoError(t, err)
		assert.Equal(t, typeName, resp.Data.Name)
	})

	t.Run("UpdateChannelType", func(t *testing.T) {
		resp, err := client.Chat().UpdateChannelType(ctx, typeName, &UpdateChannelTypeRequest{
			Automod:          "disabled",
			AutomodBehavior:  "flag",
			MaxMessageLength: 10000,
			TypingEvents:     PtrTo(false),
		})
		require.NoError(t, err)
		assert.Equal(t, 10000, resp.Data.MaxMessageLength)
		assert.False(t, resp.Data.TypingEvents)
	})

	t.Run("ListChannelTypes", func(t *testing.T) {
		resp, err := client.Chat().ListChannelTypes(ctx, &ListChannelTypesRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.ChannelTypes)

		_, found := resp.Data.ChannelTypes[typeName]
		assert.True(t, found, "Created channel type should appear in list")
	})

	t.Run("DeleteChannelType", func(t *testing.T) {
		delName := "testdeltype" + randomString(6)
		_, err := client.Chat().CreateChannelType(ctx, &CreateChannelTypeRequest{
			Name:             delName,
			Automod:          "disabled",
			AutomodBehavior:  "flag",
			MaxMessageLength: 5000,
		})
		require.NoError(t, err)

		// stream-chat-go sleeps 6s after create and retries delete up to 5 times
		time.Sleep(6 * time.Second)

		var deleteErr error
		for i := 0; i < 5; i++ {
			_, deleteErr = client.Chat().DeleteChannelType(ctx, delName, &DeleteChannelTypeRequest{})
			if deleteErr == nil {
				break
			}
			time.Sleep(time.Second)
		}
		require.NoError(t, deleteErr)
	})
}

func TestChatThreadIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 2)
	userID := userIDs[0]
	userID2 := userIDs[1]

	ch, channelID := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})
	channelCID := "messaging:" + channelID

	// Create a thread by sending a parent + reply
	parentID := sendTestMessage(t, ch, userID, "Thread parent message")

	_, err := ch.SendMessage(ctx, &SendMessageRequest{
		Message: MessageRequest{
			Text:     PtrTo("First reply in thread"),
			UserID:   PtrTo(userID2),
			ParentID: PtrTo(parentID),
		},
	})
	require.NoError(t, err)

	_, err = ch.SendMessage(ctx, &SendMessageRequest{
		Message: MessageRequest{
			Text:     PtrTo("Second reply in thread"),
			UserID:   PtrTo(userID),
			ParentID: PtrTo(parentID),
		},
	})
	require.NoError(t, err)

	t.Run("QueryThreads", func(t *testing.T) {
		// Filter by channel_cid to only get threads from our test channel
		// (same approach as stream-chat-go)
		resp, err := client.Chat().QueryThreads(ctx, &QueryThreadsRequest{
			UserID: PtrTo(userID),
			Filter: map[string]any{
				"channel_cid": map[string]any{
					"$eq": channelCID,
				},
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, resp.Data.Threads, "Should have at least one thread")

		found := false
		for _, thread := range resp.Data.Threads {
			if thread.ParentMessageID == parentID {
				found = true
				// CreatedByUserID is the first reply sender, not the parent sender
				assert.Equal(t, userID2, thread.CreatedByUserID)
			}
		}
		assert.True(t, found, "Thread should appear in query results for channel %s", channelCID)
	})

	t.Run("GetThread", func(t *testing.T) {
		resp, err := client.Chat().GetThread(ctx, parentID, &GetThreadRequest{
			ReplyLimit: PtrTo(10),
		})
		require.NoError(t, err)
		assert.Equal(t, parentID, resp.Data.Thread.ParentMessageID)
		assert.GreaterOrEqual(t, len(resp.Data.Thread.LatestReplies), 2)
	})
}

func TestChatAppSettingsIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	t.Run("GetApp", func(t *testing.T) {
		resp, err := client.GetApp(ctx, &GetAppRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp.Data)
		// App name should not be empty
		assert.NotEmpty(t, resp.Data.App.Name)
	})

	t.Run("UpdateAndVerifyApp", func(t *testing.T) {
		// Get current settings to restore later
		getResp, err := client.GetApp(ctx, &GetAppRequest{})
		require.NoError(t, err)
		originalValue := getResp.Data.App.EnforceUniqueUsernames

		// Toggle enforce_unique_usernames — safe to change on any app
		newValue := "no"
		if originalValue == "no" {
			newValue = "app"
		}
		_, err = client.UpdateApp(ctx, &UpdateAppRequest{
			EnforceUniqueUsernames: PtrTo(newValue),
		})
		require.NoError(t, err)

		// Verify update
		getResp, err = client.GetApp(ctx, &GetAppRequest{})
		require.NoError(t, err)
		assert.Equal(t, newValue, getResp.Data.App.EnforceUniqueUsernames)

		// Restore original
		_, err = client.UpdateApp(ctx, &UpdateAppRequest{
			EnforceUniqueUsernames: PtrTo(originalValue),
		})
		require.NoError(t, err)
	})
}

func TestChatUnreadCountsIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 2)
	userID := userIDs[0]
	userID2 := userIDs[1]

	// Create a channel and send a message so there's something to count
	ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})
	sendTestMessage(t, ch, userID, "Message for unread counts test")

	t.Run("UnreadCounts", func(t *testing.T) {
		resp, err := client.Chat().UnreadCounts(ctx, &UnreadCountsRequest{
			UserID: PtrTo(userID2),
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data)
		// The response should have fields — total unread count might be 0 or more
		assert.GreaterOrEqual(t, resp.Data.TotalUnreadCount, 0)
	})

	t.Run("UnreadCountsBatch", func(t *testing.T) {
		resp, err := client.Chat().UnreadCountsBatch(ctx, &UnreadCountsBatchRequest{
			UserIds: []string{userID, userID2},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.CountsByUser)
		assert.Contains(t, resp.Data.CountsByUser, userID)
		assert.Contains(t, resp.Data.CountsByUser, userID2)
	})
}

func TestChatBanIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 3)
	adminID := userIDs[0]
	targetID := userIDs[1]
	targetID2 := userIDs[2]

	t.Run("BanAndQueryBannedUsers", func(t *testing.T) {
		// Ban a user with reason and timeout
		_, err := client.Moderation().Ban(ctx, &BanRequest{
			TargetUserID: targetID,
			BannedByID:   PtrTo(adminID),
			Reason:       PtrTo("test ban reason"),
			Timeout:      PtrTo(60), // 60 minutes
		})
		require.NoError(t, err)

		// Query banned users (use $eq operator like stream-chat-go)
		qResp, err := client.Chat().QueryBannedUsers(ctx, &QueryBannedUsersRequest{
			Payload: &QueryBannedUsersPayload{
				FilterConditions: map[string]any{
					"user_id": map[string]string{"$eq": targetID},
				},
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, qResp.Data.Bans, "Should find the banned user")

		ban := qResp.Data.Bans[0]
		require.NotNil(t, ban.Reason)
		assert.Equal(t, "test ban reason", *ban.Reason)
		// When timeout is set, Expires should be populated
		assert.NotNil(t, ban.Expires, "Ban with timeout should have Expires set")

		// Unban the user
		_, err = client.Moderation().Unban(ctx, &UnbanRequest{
			TargetUserID: targetID,
		})
		require.NoError(t, err)

		// Verify ban is gone after unban
		qResp, err = client.Chat().QueryBannedUsers(ctx, &QueryBannedUsersRequest{
			Payload: &QueryBannedUsersPayload{
				FilterConditions: map[string]any{
					"user_id": map[string]string{"$eq": targetID},
				},
			},
		})
		require.NoError(t, err)
		assert.Empty(t, qResp.Data.Bans, "Bans should be empty after unban")
	})

	t.Run("ChannelBan", func(t *testing.T) {
		_, channelID := createTestChannelWithMembers(t, client, adminID, []string{adminID, targetID2})
		cid := "messaging:" + channelID

		// Ban user in the specific channel
		_, err := client.Moderation().Ban(ctx, &BanRequest{
			TargetUserID: targetID2,
			BannedByID:   PtrTo(adminID),
			ChannelCid:   PtrTo(cid),
			Reason:       PtrTo("channel-specific ban"),
		})
		require.NoError(t, err)

		// Unban from channel
		_, err = client.Moderation().Unban(ctx, &UnbanRequest{
			TargetUserID: targetID2,
			ChannelCid:   PtrTo(cid),
		})
		require.NoError(t, err)

		// Verify ban is gone after unban (same pattern as stream-chat-go)
		qResp, err := client.Chat().QueryBannedUsers(ctx, &QueryBannedUsersRequest{
			Payload: &QueryBannedUsersPayload{
				FilterConditions: map[string]any{
					"channel_cid": map[string]string{"$eq": cid},
				},
			},
		})
		require.NoError(t, err)
		assert.Empty(t, qResp.Data.Bans, "Channel bans should be empty after unban")
	})
}

func TestChatMuteIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 4)
	muterID := userIDs[0]
	targetID := userIDs[1]
	targetID2 := userIDs[2]
	targetID3 := userIDs[3]

	t.Run("MuteUnmuteUser", func(t *testing.T) {
		// Mute a user (without timeout)
		muteResp, err := client.Moderation().Mute(ctx, &MuteRequest{
			TargetIds: []string{targetID},
			UserID:    PtrTo(muterID),
		})
		require.NoError(t, err)
		require.NotEmpty(t, muteResp.Data.Mutes, "Mute response should contain mutes")

		mute := muteResp.Data.Mutes[0]
		assert.NotNil(t, mute.User, "Mute should have a User")
		assert.NotNil(t, mute.Target, "Mute should have a Target")
		assert.Nil(t, mute.Expires, "Mute without timeout should have no Expires")

		// Verify mute appears in QueryUsers (like stream-chat-go does)
		qResp, err := client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": map[string]string{"$eq": muterID},
				},
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, qResp.Data.Users)
		require.NotEmpty(t, qResp.Data.Users[0].Mutes, "User should have Mutes after muting")

		// Unmute the user
		_, err = client.Moderation().Unmute(ctx, &UnmuteRequest{
			TargetIds: []string{targetID},
			UserID:    PtrTo(muterID),
		})
		require.NoError(t, err)
	})

	t.Run("MuteWithTimeout", func(t *testing.T) {
		// Mute two users with timeout — expiration should be set
		muteResp, err := client.Moderation().Mute(ctx, &MuteRequest{
			TargetIds: []string{targetID2, targetID3},
			UserID:    PtrTo(muterID),
			Timeout:   PtrTo(60),
		})
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(muteResp.Data.Mutes), 2, "Should have at least 2 mute entries")

		// When timeout is set, Expires should be populated
		for _, m := range muteResp.Data.Mutes {
			assert.NotNil(t, m.Expires, "Mute with timeout should have Expires")
			assert.NotNil(t, m.User, "Mute should have User")
			assert.NotNil(t, m.Target, "Mute should have Target")
		}

		// Cleanup: unmute both
		_, err = client.Moderation().Unmute(ctx, &UnmuteRequest{
			TargetIds: []string{targetID2, targetID3},
			UserID:    PtrTo(muterID),
		})
		require.NoError(t, err)
	})
}

func TestChatFlagIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 2)
	userID := userIDs[0]
	flaggerID := userIDs[1]

	t.Run("FlagMessageAndQuery", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, userID, []string{userID, flaggerID})
		msgID := sendTestMessage(t, ch, userID, "Message to be flagged")

		// Flag the message using the moderation v2 API (entity_id + entity_type format).
		// Note: stream-chat-go uses the v1 FlagMessage (target_message_id format) which
		// populates the v1 chat flags store. Our SDK exposes the v2 moderation API,
		// so QueryMessageFlags (v1) may not see flags created via Moderation().Flag() (v2).
		flagResp, err := client.Moderation().Flag(ctx, &FlagRequest{
			EntityID:        msgID,
			EntityType:      "stream:chat:v1:message",
			EntityCreatorID: PtrTo(userID),
			UserID:          PtrTo(flaggerID),
			Reason:          PtrTo("inappropriate content"),
		})
		require.NoError(t, err)
		assert.NotEmpty(t, flagResp.Data.ItemID, "Flag should return an item ID")

		// Verify QueryMessageFlags endpoint works with channel_cid filter
		cid := "messaging:" + channelID
		qResp, err := client.Chat().QueryMessageFlags(ctx, &QueryMessageFlagsRequest{
			Payload: &QueryMessageFlagsPayload{
				FilterConditions: map[string]any{
					"channel_cid": cid,
				},
			},
		})
		require.NoError(t, err)
		_ = qResp // flags may be empty since v2 flags don't populate v1 store

		// Also verify QueryMessageFlags works with user_id filter
		qResp2, err := client.Chat().QueryMessageFlags(ctx, &QueryMessageFlagsRequest{
			Payload: &QueryMessageFlagsPayload{
				FilterConditions: map[string]any{
					"user_id": flaggerID,
				},
			},
		})
		require.NoError(t, err)
		_ = qResp2
	})
}

func TestChatPermissionsIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	t.Run("CreateListDeleteRole", func(t *testing.T) {
		roleName := "testrole" + randomString(6)

		// Create role
		_, err := client.CreateRole(ctx, &CreateRoleRequest{
			Name: roleName,
		})
		require.NoError(t, err)

		// List roles and verify
		listResp, err := client.ListRoles(ctx, &ListRolesRequest{})
		require.NoError(t, err)

		found := false
		for _, role := range listResp.Data.Roles {
			if role.Name == roleName {
				found = true
				assert.True(t, role.Custom, "Created role should be custom")
			}
		}
		assert.True(t, found, "Created role should appear in list")

		// Delete role (may need retry due to propagation delay)
		time.Sleep(2 * time.Second)
		var deleteErr error
		for i := 0; i < 5; i++ {
			_, deleteErr = client.DeleteRole(ctx, roleName, &DeleteRoleRequest{})
			if deleteErr == nil {
				break
			}
			time.Sleep(time.Second)
		}
		require.NoError(t, deleteErr)
	})

	t.Run("ListAndGetPermission", func(t *testing.T) {
		// List all permissions
		listResp, err := client.ListPermissions(ctx, &ListPermissionsRequest{})
		require.NoError(t, err)
		assert.NotEmpty(t, listResp.Data.Permissions, "Should have at least one permission")

		// Get a specific well-known permission
		getResp, err := client.GetPermission(ctx, "create-channel", &GetPermissionRequest{})
		require.NoError(t, err)
		assert.Equal(t, "create-channel", getResp.Data.Permission.ID)
		assert.NotEmpty(t, getResp.Data.Permission.Action)
	})
}

func TestChatExportChannelsIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 1)
	userID := userIDs[0]

	ch, channelID := createTestChannelWithMembers(t, client, userID, []string{userID})
	sendTestMessage(t, ch, userID, "Message for export test")

	cid := "messaging:" + channelID

	// Export channels
	exportResp, err := client.Chat().ExportChannels(ctx, &ExportChannelsRequest{
		Channels: []ChannelExport{
			{Cid: PtrTo(cid)},
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, exportResp.Data.TaskID)

	// Wait for the export task to complete
	taskCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	taskResult, err := WaitForTask(taskCtx, client, exportResp.Data.TaskID)
	require.NoError(t, err)
	assert.Equal(t, "completed", taskResult.Data.Status)
}

func TestChatCustomEventIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 1)
	userID := userIDs[0]

	// Send a custom event to a user (dots not allowed in event type)
	_, err := client.Chat().SendUserCustomEvent(ctx, userID, &SendUserCustomEventRequest{
		Event: UserCustomEventRequest{
			Type: "friendship_request",
			Custom: map[string]any{
				"message": "Let's be friends!",
			},
		},
	})
	require.NoError(t, err)
}

func TestChatRestoreUsersIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	// Create a dedicated user for this test
	userIDs := createTestUsers(t, client, 1)
	userID := userIDs[0]

	// Delete the user (soft delete)
	delResp, err := client.DeleteUsers(ctx, &DeleteUsersRequest{
		UserIds: []string{userID},
		User:    PtrTo("soft"),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, delResp.Data.TaskID)

	// Wait for deletion to complete
	taskCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	taskResult, err := WaitForTask(taskCtx, client, delResp.Data.TaskID)
	require.NoError(t, err)
	assert.Equal(t, "completed", taskResult.Data.Status)

	// Restore the user
	_, err = client.RestoreUsers(ctx, &RestoreUsersRequest{
		UserIds: []string{userID},
	})
	require.NoError(t, err)

	// Verify user exists after restore
	qResp, err := client.QueryUsers(ctx, &QueryUsersRequest{
		Payload: &QueryUsersPayload{
			FilterConditions: map[string]any{
				"id": userID,
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, qResp.Data.Users, "Restored user should be queryable")
	assert.Equal(t, userID, qResp.Data.Users[0].ID)
}
