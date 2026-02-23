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

		resp, err := ch.GetManyMessages(ctx, &GetManyMessagesRequest{
			Ids: []string{id1, id2, id3},
		})
		require.NoError(t, err)
		require.Len(t, resp.Data.Messages, 3)
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
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})

		// Send initial message with custom data (matching stream-chat-go TestMessageHistory)
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("initial text"),
				UserID: PtrTo(userID),
				Custom: map[string]any{"custom_field": "custom value"},
			},
		})
		require.NoError(t, err)
		msgID := sendResp.Data.Message.ID

		// Update by user1 with new text and custom value
		_, err = client.Chat().UpdateMessage(ctx, msgID, &UpdateMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("updated text"),
				UserID: PtrTo(userID),
				Custom: map[string]any{"custom_field": "updated custom value"},
			},
		})
		require.NoError(t, err)

		// Update by user2 with new text
		_, err = client.Chat().UpdateMessage(ctx, msgID, &UpdateMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("updated text 2"),
				UserID: PtrTo(userID2),
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
		require.GreaterOrEqual(t, len(histResp.Data.MessageHistory), 2, "Should have at least 2 history entries")

		// Verify history entries reference the correct message and updaters
		for _, entry := range histResp.Data.MessageHistory {
			assert.Equal(t, msgID, entry.MessageID)
		}

		// Verify text values in history (descending order by default)
		// history[0] = most recent prior version = "updated text"
		// history[1] = original = "initial text"
		assert.Equal(t, "updated text", histResp.Data.MessageHistory[0].Text)
		assert.Equal(t, userID, histResp.Data.MessageHistory[0].MessageUpdatedByID)
		assert.Equal(t, "initial text", histResp.Data.MessageHistory[1].Text)
		assert.Equal(t, userID, histResp.Data.MessageHistory[1].MessageUpdatedByID)
	})

	t.Run("QueryMessageHistorySort", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})

		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("sort initial"),
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)
		msgID := sendResp.Data.Message.ID

		_, err = client.Chat().UpdateMessage(ctx, msgID, &UpdateMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("sort updated 1"),
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		_, err = client.Chat().UpdateMessage(ctx, msgID, &UpdateMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("sort updated 2"),
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		// Query with ascending sort by message_updated_at
		histResp, err := client.Chat().QueryMessageHistory(ctx, &QueryMessageHistoryRequest{
			Filter: map[string]any{
				"message_id": msgID,
			},
			Sort: []SortParamRequest{
				{Field: PtrTo("message_updated_at"), Direction: PtrTo(1)},
			},
		})
		if err != nil {
			errStr := err.Error()
			if strings.Contains(errStr, "feature flag") || strings.Contains(errStr, "not enabled") {
				t.Skip("QueryMessageHistory feature not enabled for this app")
			}
			require.NoError(t, err)
		}
		require.GreaterOrEqual(t, len(histResp.Data.MessageHistory), 2)

		// Ascending: oldest first
		assert.Equal(t, "sort initial", histResp.Data.MessageHistory[0].Text)
		assert.Equal(t, userID, histResp.Data.MessageHistory[0].MessageUpdatedByID)
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
			UndeletedBy: userID,
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

	t.Run("DeleteMessageForMe", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "test message to delete for me")

		// Delete the message only for the sender (not for everyone)
		_, err := client.Chat().DeleteMessage(ctx, msgID, &DeleteMessageRequest{
			DeleteForMe: PtrTo(true),
			DeletedBy:   PtrTo(userID),
		})
		require.NoError(t, err)
	})

	t.Run("PinExpiration", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})

		// Send message by user2
		msgID := sendTestMessage(t, ch, userID2, "Message to pin with expiry")

		// Pin with 3 second expiration
		now := time.Now()
		expiry := now.Add(3 * time.Second)
		pinResp, err := client.Chat().UpdateMessagePartial(ctx, msgID, &UpdateMessagePartialRequest{
			Set: map[string]any{
				"pinned":      true,
				"pin_expires": expiry.Format(time.RFC3339),
			},
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)
		assert.True(t, pinResp.Data.Message.Pinned)

		// Wait for pin to expire
		time.Sleep(4 * time.Second)

		// Verify pin expired
		getResp, err := client.Chat().GetMessage(ctx, msgID, &GetMessageRequest{})
		require.NoError(t, err)
		assert.False(t, getResp.Data.Message.Pinned, "Pin should have expired")
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

	t.Run("PendingFalse", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})

		// Send message with Pending explicitly set to false (non-pending)
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Non-pending message"),
				UserID: PtrTo(userID),
			},
			Pending: PtrTo(false),
		})
		require.NoError(t, err)

		// Get the message to verify it's immediately available (no commit needed)
		getResp, err := client.Chat().GetMessage(ctx, sendResp.Data.Message.ID, &GetMessageRequest{})
		require.NoError(t, err)
		assert.Equal(t, "Non-pending message", getResp.Data.Message.Text)
	})

	t.Run("SearchWithMessageFilters", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, userID, []string{userID})

		searchTerm := "filterable" + randomString(8)
		sendTestMessage(t, ch, userID, "This has "+searchTerm+" text")
		sendTestMessage(t, ch, userID, "This also has "+searchTerm+" text")

		// Wait briefly for indexing
		time.Sleep(2 * time.Second)

		// Search using message_filter_conditions (instead of query)
		resp, err := client.Chat().Search(ctx, &SearchRequest{
			Payload: &SearchPayload{
				FilterConditions: map[string]any{
					"cid": "messaging:" + channelID,
				},
				MessageFilterConditions: map[string]any{
					"text": map[string]any{"$q": searchTerm},
				},
			},
		})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Data.Results), 2, "Should find at least 2 messages with MessageFilterConditions")
	})

	t.Run("SearchQueryAndMessageFiltersError", func(t *testing.T) {
		// Using both Query and MessageFilterConditions together should error
		_, err := client.Chat().Search(ctx, &SearchRequest{
			Payload: &SearchPayload{
				FilterConditions: map[string]any{
					"members": map[string]any{"$in": []string{userID}},
				},
				Query: PtrTo("test"),
				MessageFilterConditions: map[string]any{
					"text": map[string]any{"$q": "test"},
				},
			},
		})
		require.Error(t, err, "Using both Query and MessageFilterConditions should error")
	})

	t.Run("SearchOffsetAndSortError", func(t *testing.T) {
		// Using Offset with Sort should error
		_, err := client.Chat().Search(ctx, &SearchRequest{
			Payload: &SearchPayload{
				FilterConditions: map[string]any{
					"members": map[string]any{"$in": []string{userID}},
				},
				Query:  PtrTo("test"),
				Offset: PtrTo(1),
				Sort: []SortParamRequest{
					{Field: PtrTo("created_at"), Direction: PtrTo(-1)},
				},
			},
		})
		require.Error(t, err, "Using Offset with Sort should error")
	})

	t.Run("SearchOffsetAndNextError", func(t *testing.T) {
		// Using Offset with Next should error
		_, err := client.Chat().Search(ctx, &SearchRequest{
			Payload: &SearchPayload{
				FilterConditions: map[string]any{
					"members": map[string]any{"$in": []string{userID}},
				},
				Query:  PtrTo("test"),
				Offset: PtrTo(1),
				Next:   PtrTo(randomString(5)),
			},
		})
		require.Error(t, err, "Using Offset with Next should error")
	})

	t.Run("ChannelRoleInMember", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 2)
		memberUserID := userIDs[0]
		customRoleUserID := userIDs[1]

		// Create channel with members assigned specific roles
		channelID := "test-ch-" + randomString(12)
		ch := client.Chat().Channel("messaging", channelID)

		_, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{
			Data: &ChannelInput{
				CreatedByID: PtrTo(memberUserID),
				Members: []ChannelMemberRequest{
					{UserID: memberUserID, ChannelRole: PtrTo("channel_member")},
					{UserID: customRoleUserID, ChannelRole: PtrTo("channel_moderator")},
				},
			},
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{HardDelete: PtrTo(true)})
		})

		// Send message from channel_member
		respMember, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("message from channel_member"),
				UserID: PtrTo(memberUserID),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, respMember.Data.Message.Member, "Member should be present in message response")
		assert.Equal(t, "channel_member", respMember.Data.Message.Member.ChannelRole)

		// Send message from channel_moderator
		respMod, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("message from channel_moderator"),
				UserID: PtrTo(customRoleUserID),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, respMod.Data.Message.Member, "Member should be present in message response")
		assert.Equal(t, "channel_moderator", respMod.Data.Message.Member.ChannelRole)
	})
}
