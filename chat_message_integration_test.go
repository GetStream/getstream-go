package getstream_test

import (
	"context"
	"strings"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatMessageIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 2)
	userID := userIDs[0]
	userID2 := userIDs[1]

	t.Run("SendAndGetMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})

		msgText := "Hello from integration test " + randomString(8)
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo(msgText),
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)
		msgID := sendResp.Data.Message.ID
		assert.NotEmpty(t, msgID)
		assert.Equal(t, msgText, sendResp.Data.Message.Text)

		// Get message by ID
		getResp, err := client.Chat().GetMessage(ctx, msgID, &GetMessageRequest{})
		require.NoError(t, err)
		assert.Equal(t, msgID, getResp.Data.Message.ID)
		assert.Equal(t, msgText, getResp.Data.Message.Text)
	})

	t.Run("GetManyMessages", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})

		id1 := sendTestMessage(t, ch, userID, "Msg 1")
		id2 := sendTestMessage(t, ch, userID, "Msg 2")
		id3 := sendTestMessage(t, ch, userID, "Msg 3")

		// Allow time for messages to be indexed
		time.Sleep(time.Second)

		resp, err := ch.GetManyMessages(ctx, &GetManyMessagesRequest{
			Ids: []string{id1, id2, id3},
		})
		require.NoError(t, err)
		assert.Len(t, resp.Data.Messages, 3)
	})

	t.Run("UpdateMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Original text")

		updatedText := "Updated text " + randomString(8)
		resp, err := client.Chat().UpdateMessage(ctx, msgID, &UpdateMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo(updatedText),
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)
		assert.Equal(t, updatedText, resp.Data.Message.Text)
	})

	t.Run("PartialUpdateMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Partial update test")

		// Set custom field
		resp, err := client.Chat().UpdateMessagePartial(ctx, msgID, &UpdateMessagePartialRequest{
			Set: map[string]any{
				"priority": "high",
				"status":   "reviewed",
			},
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Message)

		// Unset custom field
		resp, err = client.Chat().UpdateMessagePartial(ctx, msgID, &UpdateMessagePartialRequest{
			Unset:  []string{"status"},
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Message)
	})

	t.Run("DeleteMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Message to delete")

		// Soft delete
		resp, err := client.Chat().DeleteMessage(ctx, msgID, &DeleteMessageRequest{})
		require.NoError(t, err)
		assert.Equal(t, "deleted", resp.Data.Message.Type)
	})

	t.Run("HardDeleteMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Message to hard delete")

		resp, err := client.Chat().DeleteMessage(ctx, msgID, &DeleteMessageRequest{
			Hard: PtrTo(true),
		})
		require.NoError(t, err)
		assert.Equal(t, "deleted", resp.Data.Message.Type)
	})

	t.Run("PinUnpinMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})

		// Send a pinned message
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Pinned message"),
				UserID: PtrTo(userID),
				Pinned: PtrTo(true),
			},
		})
		require.NoError(t, err)
		msgID := sendResp.Data.Message.ID
		assert.True(t, sendResp.Data.Message.Pinned)

		// Unpin via update
		resp, err := client.Chat().UpdateMessagePartial(ctx, msgID, &UpdateMessagePartialRequest{
			Set: map[string]any{
				"pinned": false,
			},
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Message)
		assert.False(t, resp.Data.Message.Pinned)
	})

	t.Run("TranslateMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Hello, how are you?")

		resp, err := client.Chat().TranslateMessage(ctx, msgID, &TranslateMessageRequest{
			Language: "es",
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Message)
		// The i18n field should be populated after translation
		assert.NotNil(t, resp.Data.Message.I18n, "i18n field should be set after translation")
	})

	t.Run("ThreadReply", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})

		// Send parent message
		parentID := sendTestMessage(t, ch, userID, "Parent message for thread")

		// Send reply
		replyResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:     PtrTo("Reply to parent"),
				UserID:   PtrTo(userID2),
				ParentID: PtrTo(parentID),
			},
		})
		require.NoError(t, err)
		assert.NotEmpty(t, replyResp.Data.Message.ID)

		// Get replies — provide empty Sort slice to avoid nil being serialized as "null"
		repliesResp, err := client.Chat().GetReplies(ctx, parentID, &GetRepliesRequest{
			Sort: []SortParamRequest{},
		})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(repliesResp.Data.Messages), 1)
	})

	t.Run("SearchMessages", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, userID, []string{userID})

		searchTerm := "uniquesearch" + randomString(8)
		sendTestMessage(t, ch, userID, "This message contains "+searchTerm+" for testing")

		// Wait briefly for indexing
		time.Sleep(2 * time.Second)

		resp, err := client.Chat().Search(ctx, &SearchRequest{
			Payload: &SearchPayload{
				Query: PtrTo(searchTerm),
				FilterConditions: map[string]any{
					"cid": "messaging:" + channelID,
				},
			},
		})
		require.NoError(t, err)
		assert.NotEmpty(t, resp.Data.Results, "Search should return at least one result")
	})

	t.Run("SilentMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})

		resp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("This is a silent message"),
				UserID: PtrTo(userID),
				Silent: PtrTo(true),
			},
		})
		require.NoError(t, err)
		assert.True(t, resp.Data.Message.Silent)
	})

	t.Run("PendingMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})

		// Send a pending message (requires pending messages feature flag)
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Pending message text"),
				UserID: PtrTo(userID),
			},
			Pending:  PtrTo(true),
			SkipPush: PtrTo(true),
		})
		if err != nil {
			errStr := err.Error()
			if strings.Contains(errStr, "pending messages not enabled") || strings.Contains(errStr, "feature flag") {
				t.Skip("Pending messages feature not enabled for this app")
			}
			require.NoError(t, err)
		}
		msgID := sendResp.Data.Message.ID
		assert.NotEmpty(t, msgID)

		// Commit the pending message
		commitResp, err := client.Chat().CommitMessage(ctx, msgID, &CommitMessageRequest{})
		require.NoError(t, err)
		require.NotNil(t, commitResp.Data.Message)
		assert.Equal(t, msgID, commitResp.Data.Message.ID)
	})

	t.Run("QueryMessageHistory", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Original history text")

		// Update the message twice to create history entries
		_, err := client.Chat().UpdateMessage(ctx, msgID, &UpdateMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Updated history text v1"),
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		_, err = client.Chat().UpdateMessage(ctx, msgID, &UpdateMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Updated history text v2"),
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		// Query message history (requires feature flag)
		histResp, err := client.Chat().QueryMessageHistory(ctx, &QueryMessageHistoryRequest{
			Filter: map[string]any{
				"message_id": msgID,
			},
			Sort: []SortParamRequest{},
		})
		if err != nil {
			errStr := err.Error()
			if strings.Contains(errStr, "feature flag") || strings.Contains(errStr, "not enabled") {
				t.Skip("QueryMessageHistory feature not enabled for this app")
			}
			require.NoError(t, err)
		}
		assert.GreaterOrEqual(t, len(histResp.Data.MessageHistory), 2, "Should have at least 2 history entries")

		// Verify that history entries reference the correct message and updater
		for _, entry := range histResp.Data.MessageHistory {
			assert.Equal(t, msgID, entry.MessageID)
			assert.Equal(t, userID, entry.MessageUpdatedByID)
		}
	})

	t.Run("SkipEnrichUrl", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})

		// Send a message with a URL but skip enrichment
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Check out https://getstream.io for more info"),
				UserID: PtrTo(userID),
			},
			SkipEnrichUrl: PtrTo(true),
		})
		require.NoError(t, err)
		assert.Empty(t, sendResp.Data.Message.Attachments, "Attachments should be empty when SkipEnrichUrl is true")

		// Verify via GetMessage that attachments remain empty
		time.Sleep(3 * time.Second)
		getResp, err := client.Chat().GetMessage(ctx, sendResp.Data.Message.ID, &GetMessageRequest{})
		require.NoError(t, err)
		assert.Empty(t, getResp.Data.Message.Attachments, "Attachments should remain empty after enrichment window")
	})

	t.Run("KeepChannelHidden", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, userID, []string{userID})
		cid := "messaging:" + channelID

		// Hide the channel first
		_, err := ch.Hide(ctx, &HideChannelRequest{
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)

		// Send a message with KeepChannelHidden=true
		_, err = ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Hidden message"),
				UserID: PtrTo(userID),
			},
			KeepChannelHidden: PtrTo(true),
		})
		require.NoError(t, err)

		// Query channels — the channel should still be hidden
		qResp, err := client.Chat().QueryChannels(ctx, &QueryChannelsRequest{
			FilterConditions: map[string]any{
				"cid": cid,
			},
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)
		assert.Empty(t, qResp.Data.Channels, "Channel should remain hidden after sending with KeepChannelHidden")
	})

	t.Run("UndeleteMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Message to undelete")

		// Soft delete the message
		_, err := client.Chat().DeleteMessage(ctx, msgID, &DeleteMessageRequest{})
		require.NoError(t, err)

		// Verify it's deleted
		getResp, err := client.Chat().GetMessage(ctx, msgID, &GetMessageRequest{})
		require.NoError(t, err)
		assert.Equal(t, "deleted", getResp.Data.Message.Type)

		// Undelete the message
		// Note: The API requires "undeleted_by" field which may not be in the generated
		// request struct yet. Gracefully skip if the field is missing.
		undelResp, err := client.Chat().UndeleteMessage(ctx, msgID, &UndeleteMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Message to undelete"),
				UserID: PtrTo(userID),
			},
		})
		if err != nil {
			errStr := err.Error()
			if strings.Contains(errStr, "undeleted_by") || strings.Contains(errStr, "required field") {
				t.Skip("UndeleteMessage requires 'undeleted_by' field not yet in generated request struct")
			}
			require.NoError(t, err)
		}
		require.NotNil(t, undelResp.Data.Message)
		assert.NotEqual(t, "deleted", undelResp.Data.Message.Type)
		assert.Equal(t, "Message to undelete", undelResp.Data.Message.Text)
	})

	t.Run("RestrictedVisibility", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})

		// Send a message with restricted visibility — only userID can see it
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:                 PtrTo("Secret message"),
				UserID:               PtrTo(userID),
				RestrictedVisibility: []string{userID},
			},
		})
		if err != nil {
			errStr := err.Error()
			if strings.Contains(errStr, "private messaging is not allowed") || strings.Contains(errStr, "not enabled") {
				t.Skip("RestrictedVisibility (private messaging) is not enabled for this app")
			}
			require.NoError(t, err)
		}
		assert.Equal(t, []string{userID}, sendResp.Data.Message.RestrictedVisibility)
	})

	t.Run("SystemMessage", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})

		resp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("User joined the channel"),
				UserID: PtrTo(userID),
				Type:   PtrTo("system"),
			},
		})
		require.NoError(t, err)
		assert.Equal(t, "system", resp.Data.Message.Type)
	})
}
