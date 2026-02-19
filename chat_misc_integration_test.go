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

		// Need a small delay before deleting a channel type
		time.Sleep(2 * time.Second)

		_, err = client.Chat().DeleteChannelType(ctx, delName, &DeleteChannelTypeRequest{})
		require.NoError(t, err)
	})
}

func TestChatThreadIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 2)
	userID := userIDs[0]
	userID2 := userIDs[1]

	ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})

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
		resp, err := client.Chat().QueryThreads(ctx, &QueryThreadsRequest{
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)
		require.NotEmpty(t, resp.Data.Threads, "Should have at least one thread")

		found := false
		for _, thread := range resp.Data.Threads {
			if thread.ParentMessageID == parentID {
				found = true
				assert.Equal(t, userID, thread.CreatedByUserID)
			}
		}
		assert.True(t, found, "Created thread should appear in query results")
	})

	t.Run("GetThread", func(t *testing.T) {
		resp, err := client.Chat().GetThread(ctx, parentID, &GetThreadRequest{})
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
		// Get current settings first
		getResp, err := client.GetApp(ctx, &GetAppRequest{})
		require.NoError(t, err)

		originalWebhookURL := getResp.Data.App.WebhookUrl

		// Update a setting
		testURL := "https://example.com/webhook-integration-test"
		_, err = client.UpdateApp(ctx, &UpdateAppRequest{
			WebhookUrl: PtrTo(testURL),
		})
		require.NoError(t, err)

		// Verify update
		getResp, err = client.GetApp(ctx, &GetAppRequest{})
		require.NoError(t, err)
		assert.Equal(t, testURL, getResp.Data.App.WebhookUrl)

		// Restore original
		_, err = client.UpdateApp(ctx, &UpdateAppRequest{
			WebhookUrl: PtrTo(originalWebhookURL),
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
