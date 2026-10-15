package getstream_test

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// errCodeScheduledMessagesDisabled is the API error code returned when
// scheduled messages are not enabled for the app.
const errCodeScheduledMessagesDisabled = 114

// scheduleTestMessage schedules a message an hour from now and returns the
// scheduled message. It skips the test when the app cannot schedule: the
// feature is opt-in per app, and backends without it ignore scheduled_at and
// send the message right away.
func scheduleTestMessage(t *testing.T, ch *Channels, userID, text string) ScheduledMessageResponse {
	t.Helper()
	resp, err := ch.SendMessage(context.Background(), &SendMessageRequest{
		Message: MessageRequest{
			Text:   PtrTo(text),
			UserID: PtrTo(userID),
		},
		ScheduledAt: &Timestamp{Time: PtrTo(time.Now().Add(time.Hour))},
	})
	var streamErr *StreamError
	if errors.As(err, &streamErr) && streamErr.Code == errCodeScheduledMessagesDisabled {
		t.Skip("Scheduled messages are not enabled for this app")
	}
	require.NoError(t, err)
	if resp.Data.ScheduledMessage == nil {
		t.Skip("Backend does not support scheduled messages: scheduled_at was ignored")
	}
	return *resp.Data.ScheduledMessage
}

func TestChatScheduledMessageIntegration(t *testing.T) {
	t.Parallel()
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 1)
	userID := userIDs[0]

	t.Run("ScheduleQueryUpdateAndSendNow", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, userID, []string{userID})
		cid := "messaging:" + channelID

		scheduled := scheduleTestMessage(t, ch, userID, "Scheduled from Go SDK")
		assert.NotEmpty(t, scheduled.ID)
		assert.Equal(t, "scheduled", scheduled.Status)
		assert.Equal(t, cid, scheduled.Cid)
		assert.Equal(t, "Scheduled from Go SDK", scheduled.Message.Text)

		queryResp, err := client.Chat().QueryScheduledMessages(ctx, &QueryScheduledMessagesRequest{
			UserID: PtrTo(userID),
			Filter: &ScheduledMessagesFilter{Cid: PtrTo(cid)},
		})
		require.NoError(t, err)
		require.Len(t, queryResp.Data.ScheduledMessages, 1)
		assert.Equal(t, scheduled.ID, queryResp.Data.ScheduledMessages[0].ID)

		newTime := time.Now().Add(2 * time.Hour).Truncate(time.Millisecond)
		updateResp, err := client.Chat().UpdateScheduledMessage(ctx, scheduled.ID, &UpdateScheduledMessageRequest{
			ScheduledAt: &Timestamp{Time: &newTime},
			Message: &MessageRequest{
				Text: PtrTo("Rescheduled from Go SDK"),
			},
		})
		require.NoError(t, err)
		assert.Equal(t, "Rescheduled from Go SDK", updateResp.Data.ScheduledMessage.Message.Text)
		require.NotNil(t, updateResp.Data.ScheduledMessage.ScheduledAt.Time)
		assert.True(t, newTime.Equal(*updateResp.Data.ScheduledMessage.ScheduledAt.Time))

		sendResp, err := client.Chat().SendScheduledMessageNow(ctx, scheduled.ID, &SendScheduledMessageNowRequest{})
		require.NoError(t, err)
		assert.Equal(t, scheduled.ID, sendResp.Data.ScheduledMessage.ID)
	})

	t.Run("DeleteScheduledMessage", func(t *testing.T) {
		ch, channelID := createTestChannelWithMembers(t, client, userID, []string{userID})
		cid := "messaging:" + channelID

		scheduled := scheduleTestMessage(t, ch, userID, "Deleted before sending")

		deleteResp, err := client.Chat().DeleteScheduledMessage(ctx, scheduled.ID, &DeleteScheduledMessageRequest{})
		require.NoError(t, err)
		assert.Equal(t, scheduled.ID, deleteResp.Data.ScheduledMessage.ID)
		assert.Equal(t, "canceled", deleteResp.Data.ScheduledMessage.Status)

		queryResp, err := client.Chat().QueryScheduledMessages(ctx, &QueryScheduledMessagesRequest{
			UserID: PtrTo(userID),
			Filter: &ScheduledMessagesFilter{Cid: PtrTo(cid)},
		})
		require.NoError(t, err)
		assert.Empty(t, queryResp.Data.ScheduledMessages)
	})
}
