package getstream_test

import (
	"context"
	"testing"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// skipIfShort skips integration tests when running with -short flag.
func skipIfShort(t *testing.T) {
	t.Helper()
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
}

// createTestUsers creates n test users and returns their IDs.
func createTestUsers(t *testing.T, client *Stream, n int) []string {
	t.Helper()
	ctx := context.Background()
	users := make(map[string]UserRequest, n)
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		id := "test-user-" + uuid.New().String()
		ids[i] = id
		users[id] = UserRequest{
			ID:   id,
			Name: PtrTo("Test User " + id[:8]),
			Role: PtrTo("user"),
		}
	}
	_, err := client.UpdateUsers(ctx, &UpdateUsersRequest{Users: users})
	require.NoError(t, err, "Failed to create test users")
	return ids
}

// createTestChannel creates a messaging channel and registers cleanup to delete it.
func createTestChannel(t *testing.T, client *Stream, creatorID string) (*Channels, string) {
	t.Helper()
	ctx := context.Background()
	channelID := "test-ch-" + randomString(12)
	ch := client.Chat().Channel("messaging", channelID)

	_, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{
		Data: &ChannelInput{
			CreatedByID: PtrTo(creatorID),
		},
	})
	require.NoError(t, err, "Failed to create test channel")

	t.Cleanup(func() {
		_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{})
	})

	return ch, channelID
}

// createTestChannelWithMembers creates a messaging channel with members and registers cleanup.
func createTestChannelWithMembers(t *testing.T, client *Stream, creatorID string, memberIDs []string) (*Channels, string) {
	t.Helper()
	ctx := context.Background()
	channelID := "test-ch-" + randomString(12)
	ch := client.Chat().Channel("messaging", channelID)

	members := make([]ChannelMemberRequest, len(memberIDs))
	for i, id := range memberIDs {
		members[i] = ChannelMemberRequest{UserID: id}
	}

	_, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{
		Data: &ChannelInput{
			CreatedByID: PtrTo(creatorID),
			Members:     members,
		},
	})
	require.NoError(t, err, "Failed to create test channel with members")

	t.Cleanup(func() {
		_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{})
	})

	return ch, channelID
}

// sendTestMessage sends a message to a channel and returns the message ID.
func sendTestMessage(t *testing.T, ch *Channels, userID string, text string) string {
	t.Helper()
	ctx := context.Background()
	resp, err := ch.SendMessage(ctx, &SendMessageRequest{
		Message: MessageRequest{
			Text:   PtrTo(text),
			UserID: PtrTo(userID),
		},
	})
	require.NoError(t, err, "Failed to send test message")
	require.NotEmpty(t, resp.Data.Message.ID, "Message ID should not be empty")
	return resp.Data.Message.ID
}
