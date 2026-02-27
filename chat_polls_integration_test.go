package getstream_test

import (
	"context"
	"strings"
	"testing"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatPollsIntegration(t *testing.T) {
	t.Parallel()
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 2)
	userID := userIDs[0]
	voterID := userIDs[1]

	t.Run("CreateGetUpdateDeletePoll", func(t *testing.T) {
		// Create a poll
		createResp, err := client.CreatePoll(ctx, &CreatePollRequest{
			Name:              "Favorite color?",
			Description:       PtrTo("Pick your favorite color"),
			EnforceUniqueVote: PtrTo(true),
			UserID:            PtrTo(userID),
			Options: []PollOptionInput{
				{Text: PtrTo("Red")},
				{Text: PtrTo("Blue")},
				{Text: PtrTo("Green")},
			},
		})
		require.NoError(t, err)
		pollID := createResp.Data.Poll.ID
		assert.NotEmpty(t, pollID)
		assert.Equal(t, "Favorite color?", createResp.Data.Poll.Name)
		assert.True(t, createResp.Data.Poll.EnforceUniqueVote)
		assert.Len(t, createResp.Data.Poll.Options, 3)

		t.Cleanup(func() {
			_, _ = client.DeletePoll(context.Background(), pollID, &DeletePollRequest{UserID: PtrTo(userID)})
		})

		// Get the poll
		getResp, err := client.GetPoll(ctx, pollID, &GetPollRequest{})
		require.NoError(t, err)
		assert.Equal(t, pollID, getResp.Data.Poll.ID)
		assert.Equal(t, "Favorite color?", getResp.Data.Poll.Name)

		// Update the poll
		updateResp, err := client.UpdatePoll(ctx, &UpdatePollRequest{
			ID:          pollID,
			Name:        "Updated: Favorite color?",
			Description: PtrTo("Updated description"),
			UserID:      PtrTo(userID),
		})
		require.NoError(t, err)
		assert.Equal(t, "Updated: Favorite color?", updateResp.Data.Poll.Name)

		// Delete the poll
		_, err = client.DeletePoll(ctx, pollID, &DeletePollRequest{UserID: PtrTo(userID)})
		require.NoError(t, err)
	})

	t.Run("QueryPolls", func(t *testing.T) {
		// Create a poll to query
		createResp, err := client.CreatePoll(ctx, &CreatePollRequest{
			Name:   "Query test poll " + randomString(8),
			UserID: PtrTo(userID),
			Options: []PollOptionInput{
				{Text: PtrTo("Option A")},
				{Text: PtrTo("Option B")},
			},
		})
		require.NoError(t, err)
		pollID := createResp.Data.Poll.ID

		t.Cleanup(func() {
			_, _ = client.DeletePoll(context.Background(), pollID, &DeletePollRequest{UserID: PtrTo(userID)})
		})

		// Query polls (server-side auth requires user_id)
		qResp, err := client.QueryPolls(ctx, &QueryPollsRequest{
			UserID: PtrTo(userID),
			Filter: map[string]any{
				"id": pollID,
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, qResp.Data.Polls)
		assert.Equal(t, pollID, qResp.Data.Polls[0].ID)
	})

	t.Run("CastPollVote", func(t *testing.T) {
		// Create a poll
		createResp, err := client.CreatePoll(ctx, &CreatePollRequest{
			Name:              "Vote test poll",
			EnforceUniqueVote: PtrTo(true),
			UserID:            PtrTo(userID),
			Options: []PollOptionInput{
				{Text: PtrTo("Yes")},
				{Text: PtrTo("No")},
			},
		})
		require.NoError(t, err)
		pollID := createResp.Data.Poll.ID
		optionID := createResp.Data.Poll.Options[0].ID

		t.Cleanup(func() {
			_, _ = client.DeletePoll(context.Background(), pollID, &DeletePollRequest{UserID: PtrTo(userID)})
		})

		// Send a message with the poll attached
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, voterID})
		sendResp, err := ch.SendMessage(ctx, &SendMessageRequest{
			Message: MessageRequest{
				Text:   PtrTo("Please vote!"),
				UserID: PtrTo(userID),
				PollID: PtrTo(pollID),
			},
		})
		if err != nil {
			if strings.Contains(err.Error(), "polls not enabled") {
				t.Skip("Polls not enabled for this channel")
			}
		}
		require.NoError(t, err)
		msgID := sendResp.Data.Message.ID

		// Cast a vote
		voteResp, err := client.Chat().CastPollVote(ctx, msgID, pollID, &CastPollVoteRequest{
			UserID: PtrTo(voterID),
			Vote: &VoteData{
				OptionID: PtrTo(optionID),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, voteResp.Data.Vote)
		assert.Equal(t, optionID, voteResp.Data.Vote.OptionID)

		// Verify poll has votes by getting it
		getResp, err := client.GetPoll(ctx, pollID, &GetPollRequest{})
		require.NoError(t, err)
		assert.Equal(t, 1, getResp.Data.Poll.VoteCount)
	})
}
