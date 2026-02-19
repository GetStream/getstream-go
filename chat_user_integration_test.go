package getstream_test

import (
	"context"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatUserIntegration(t *testing.T) {
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	t.Run("UpsertUsers", func(t *testing.T) {
		userID1 := "upsert-" + uuid.New().String()
		userID2 := "upsert-" + uuid.New().String()

		resp, err := client.UpdateUsers(ctx, &UpdateUsersRequest{
			Users: map[string]UserRequest{
				userID1: {
					ID:   userID1,
					Name: PtrTo("User One"),
					Role: PtrTo("user"),
				},
				userID2: {
					ID:   userID2,
					Name: PtrTo("User Two"),
					Role: PtrTo("user"),
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Users)

		u1, ok := resp.Data.Users[userID1]
		require.True(t, ok, "User 1 should be in response")
		assert.Equal(t, userID1, u1.ID)
		assert.Equal(t, "User One", *u1.Name)

		u2, ok := resp.Data.Users[userID2]
		require.True(t, ok, "User 2 should be in response")
		assert.Equal(t, userID2, u2.ID)
		assert.Equal(t, "User Two", *u2.Name)
	})

	t.Run("QueryUsers", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 2)

		resp, err := client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": map[string]any{"$in": userIDs},
				},
			},
		})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Data.Users), 2)

		foundIDs := make(map[string]bool)
		for _, u := range resp.Data.Users {
			foundIDs[u.ID] = true
		}
		for _, id := range userIDs {
			assert.True(t, foundIDs[id], "User %s should be found in query results", id)
		}
	})

	t.Run("PartialUpdateUser", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 1)
		userID := userIDs[0]

		_, err := client.UpdateUsersPartial(ctx, &UpdateUsersPartialRequest{
			Users: []UpdateUserPartialRequest{
				{
					ID: userID,
					Set: map[string]any{
						"country": "NL",
						"role":    "admin",
					},
				},
			},
		})
		require.NoError(t, err)

		// Verify the update
		resp, err := client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": userID,
				},
			},
		})
		require.NoError(t, err)
		require.Len(t, resp.Data.Users, 1)

		// Now unset
		_, err = client.UpdateUsersPartial(ctx, &UpdateUsersPartialRequest{
			Users: []UpdateUserPartialRequest{
				{
					ID:    userID,
					Unset: []string{"country"},
				},
			},
		})
		require.NoError(t, err)
	})

	t.Run("BlockUnblockUser", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 2)
		alice := userIDs[0]
		bob := userIDs[1]

		// Block
		_, err := client.BlockUsers(ctx, &BlockUsersRequest{
			BlockedUserID: bob,
			UserID:        &alice,
		})
		require.NoError(t, err)

		// Verify blocked
		resp, err := client.GetBlockedUsers(ctx, &GetBlockedUsersRequest{
			UserID: &alice,
		})
		require.NoError(t, err)
		require.NotEmpty(t, resp.Data.Blocks, "Should have at least one block")

		found := false
		for _, b := range resp.Data.Blocks {
			if b.BlockedUserID == bob {
				found = true
				break
			}
		}
		assert.True(t, found, "Bob should be in Alice's blocked list")

		// Unblock
		_, err = client.UnblockUsers(ctx, &UnblockUsersRequest{
			BlockedUserID: bob,
			UserID:        &alice,
		})
		require.NoError(t, err)

		// Verify unblocked
		resp, err = client.GetBlockedUsers(ctx, &GetBlockedUsersRequest{
			UserID: &alice,
		})
		require.NoError(t, err)
		for _, b := range resp.Data.Blocks {
			assert.NotEqual(t, bob, b.BlockedUserID, "Bob should no longer be blocked")
		}
	})

	t.Run("DeactivateReactivateUser", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 1)
		userID := userIDs[0]

		// Deactivate
		_, err := client.DeactivateUser(ctx, userID, &DeactivateUserRequest{})
		require.NoError(t, err)

		// Reactivate
		_, err = client.ReactivateUser(ctx, userID, &ReactivateUserRequest{})
		require.NoError(t, err)

		// Verify user is active again by querying
		resp, err := client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": userID,
				},
			},
		})
		require.NoError(t, err)
		require.Len(t, resp.Data.Users, 1)
		assert.Equal(t, userID, resp.Data.Users[0].ID)
	})

	t.Run("DeleteUsers", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 2)

		resp, err := client.DeleteUsers(ctx, &DeleteUsersRequest{
			UserIds: userIDs,
		})
		require.NoError(t, err)

		taskID := resp.Data.TaskID
		require.NotEmpty(t, taskID, "Task ID should not be empty")

		taskCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		taskResult, err := WaitForTask(taskCtx, client, taskID)
		require.NoError(t, err)
		assert.Equal(t, "completed", taskResult.Data.Status)
	})

	t.Run("ExportUser", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 1)
		userID := userIDs[0]

		resp, err := client.ExportUser(ctx, userID, &ExportUserRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp.Data)
	})

	t.Run("CreateGuest", func(t *testing.T) {
		guestID := "guest-" + uuid.New().String()

		resp, err := client.CreateGuest(ctx, &CreateGuestRequest{
			User: UserRequest{
				ID:   guestID,
				Name: PtrTo("Guest User"),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data)
		assert.NotEmpty(t, resp.Data.AccessToken, "Access token should not be empty")
		assert.Equal(t, guestID, resp.Data.User.ID)
	})

	t.Run("UserCustomData", func(t *testing.T) {
		userID := "custom-" + uuid.New().String()

		custom := map[string]any{
			"favorite_color": "blue",
			"age":            float64(30),
			"tags":           []any{"vip", "early_adopter"},
		}

		resp, err := client.UpdateUsers(ctx, &UpdateUsersRequest{
			Users: map[string]UserRequest{
				userID: {
					ID:     userID,
					Name:   PtrTo("Custom User"),
					Custom: custom,
				},
			},
		})
		require.NoError(t, err)

		u, ok := resp.Data.Users[userID]
		require.True(t, ok)
		assert.Equal(t, "blue", u.Custom["favorite_color"])
		assert.Equal(t, float64(30), u.Custom["age"])

		// Query back to verify persistence
		queryResp, err := client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": userID,
				},
			},
		})
		require.NoError(t, err)
		require.Len(t, queryResp.Data.Users, 1)
		assert.Equal(t, "blue", queryResp.Data.Users[0].Custom["favorite_color"])
	})
}
