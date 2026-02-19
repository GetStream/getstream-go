package getstream_test

import (
	"context"
	"testing"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatReactionIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	userIDs := createTestUsers(t, client, 2)
	userID := userIDs[0]
	userID2 := userIDs[1]

	t.Run("SendReaction", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})
		msgID := sendTestMessage(t, ch, userID, "React to this message")

		resp, err := client.Chat().SendReaction(ctx, msgID, &SendReactionRequest{
			Reaction: ReactionRequest{
				Type:   "like",
				UserID: PtrTo(userID2),
			},
		})
		require.NoError(t, err)
		assert.Equal(t, "like", resp.Data.Reaction.Type)
		assert.Equal(t, userID2, resp.Data.Reaction.UserID)
	})

	t.Run("GetReactions", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})
		msgID := sendTestMessage(t, ch, userID, "Message with reactions")

		// Send a couple of reactions
		_, err := client.Chat().SendReaction(ctx, msgID, &SendReactionRequest{
			Reaction: ReactionRequest{
				Type:   "like",
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		_, err = client.Chat().SendReaction(ctx, msgID, &SendReactionRequest{
			Reaction: ReactionRequest{
				Type:   "love",
				UserID: PtrTo(userID2),
			},
		})
		require.NoError(t, err)

		// Get reactions
		resp, err := client.Chat().GetReactions(ctx, msgID, &GetReactionsRequest{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Data.Reactions), 2)
	})

	t.Run("DeleteReaction", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Message for reaction deletion")

		// Send reaction
		_, err := client.Chat().SendReaction(ctx, msgID, &SendReactionRequest{
			Reaction: ReactionRequest{
				Type:   "like",
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		// Delete reaction
		_, err = client.Chat().DeleteReaction(ctx, msgID, "like", &DeleteReactionRequest{
			UserID: PtrTo(userID),
		})
		require.NoError(t, err)

		// Verify reaction is gone
		resp, err := client.Chat().GetReactions(ctx, msgID, &GetReactionsRequest{})
		require.NoError(t, err)

		for _, r := range resp.Data.Reactions {
			if r.UserID == userID {
				assert.NotEqual(t, "like", r.Type, "Like reaction should have been deleted")
			}
		}
	})

	t.Run("QueryReactions", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID, userID2})
		msgID := sendTestMessage(t, ch, userID, "Message for query reactions")

		// Send reactions
		_, err := client.Chat().SendReaction(ctx, msgID, &SendReactionRequest{
			Reaction: ReactionRequest{
				Type:   "like",
				UserID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		_, err = client.Chat().SendReaction(ctx, msgID, &SendReactionRequest{
			Reaction: ReactionRequest{
				Type:   "wow",
				UserID: PtrTo(userID2),
			},
		})
		require.NoError(t, err)

		// Query reactions
		resp, err := client.Chat().QueryReactions(ctx, msgID, &QueryReactionsRequest{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Data.Reactions), 2)
	})

	t.Run("EnforceUniqueReaction", func(t *testing.T) {
		ch, _ := createTestChannelWithMembers(t, client, userID, []string{userID})
		msgID := sendTestMessage(t, ch, userID, "Message for unique reaction test")

		// Send first like reaction
		_, err := client.Chat().SendReaction(ctx, msgID, &SendReactionRequest{
			Reaction: ReactionRequest{
				Type:   "like",
				UserID: PtrTo(userID),
			},
			EnforceUnique: PtrTo(true),
		})
		require.NoError(t, err)

		// Send second like reaction with EnforceUnique — should replace, not duplicate
		_, err = client.Chat().SendReaction(ctx, msgID, &SendReactionRequest{
			Reaction: ReactionRequest{
				Type:   "love",
				UserID: PtrTo(userID),
			},
			EnforceUnique: PtrTo(true),
		})
		require.NoError(t, err)

		// Query reactions — user should only have one
		resp, err := client.Chat().GetReactions(ctx, msgID, &GetReactionsRequest{})
		require.NoError(t, err)

		userReactions := 0
		for _, r := range resp.Data.Reactions {
			if r.UserID == userID {
				userReactions++
			}
		}
		assert.Equal(t, 1, userReactions, "EnforceUnique should keep only one reaction per user")
	})
}
