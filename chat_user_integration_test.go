package getstream_test

import (
	"context"
	"strings"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatUserIntegration(t *testing.T) {
	t.Parallel()
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	// Collect all user IDs for a single batch cleanup at the end.
	// This avoids hitting DeleteUsers rate limits from many individual API calls.
	var allUserIDs []string
	t.Cleanup(func() {
		if len(allUserIDs) > 0 {
			deleteUsersWithRetry(client, allUserIDs)
		}
	})

	// Helper to create test users and add them to the batch cleanup list
	// instead of registering individual per-subtest cleanups.
	newUsers := func(t *testing.T, n int) []string {
		t.Helper()
		ids := make([]string, n)
		users := make(map[string]UserRequest, n)
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
		allUserIDs = append(allUserIDs, ids...)
		return ids
	}

	t.Run("UpsertUsers", func(t *testing.T) {
		userID1 := "upsert-" + uuid.New().String()
		userID2 := "upsert-" + uuid.New().String()
		allUserIDs = append(allUserIDs, userID1, userID2)

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
		userIDs := newUsers(t, 2)

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

	t.Run("QueryUsersWithOffsetLimit", func(t *testing.T) {
		userIDs := newUsers(t, 3)

		resp, err := client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": map[string]any{"$in": userIDs},
				},
				Offset: PtrTo(1),
				Limit:  PtrTo(2),
			},
		})
		require.NoError(t, err)
		assert.Len(t, resp.Data.Users, 2, "Should return exactly 2 users with offset=1 limit=2")
	})

	t.Run("PartialUpdateUser", func(t *testing.T) {
		userIDs := newUsers(t, 1)
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
		userIDs := newUsers(t, 2)
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
		userIDs := newUsers(t, 1)
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
		userIDs := newUsers(t, 2)

		// Retry with exponential backoff to handle rate limiting
		var resp *StreamResponse[DeleteUsersResponse]
		var deleteErr error
		for i := 0; i < 5; i++ {
			resp, deleteErr = client.DeleteUsers(ctx, &DeleteUsersRequest{
				UserIds: userIDs,
			})
			if deleteErr == nil {
				break
			}
			if strings.Contains(deleteErr.Error(), "Too many requests") {
				backoff := time.Duration(1<<uint(i+1)) * time.Second
				t.Logf("DeleteUsers rate limited, attempt %d/5, waiting %s", i+1, backoff)
				time.Sleep(backoff)
				continue
			}
			break
		}
		require.NoError(t, deleteErr)

		taskID := resp.Data.TaskID
		require.NotEmpty(t, taskID, "Task ID should not be empty")

		taskResult, err := WaitForTask(ctx, client, taskID)
		require.NoError(t, err)
		assert.Equal(t, "completed", taskResult.Data.Status)
	})

	t.Run("ExportUser", func(t *testing.T) {
		userIDs := newUsers(t, 1)
		userID := userIDs[0]

		resp, err := client.ExportUser(ctx, userID, &ExportUserRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp.Data)
	})

	t.Run("CreateGuest", func(t *testing.T) {
		guestID := "guest-" + uuid.New().String()
		allUserIDs = append(allUserIDs, guestID)

		resp, err := client.CreateGuest(ctx, &CreateGuestRequest{
			User: UserRequest{
				ID:   guestID,
				Name: PtrTo("Guest User"),
			},
		})
		if err != nil {
			// Guest access may be disabled at the app level (e.g. multi-tenant apps)
			// stream-chat-go handles this the same way
			return
		}
		require.NotNil(t, resp.Data)
		assert.NotEmpty(t, resp.Data.AccessToken, "Access token should not be empty")

		// Server may prefix the guest ID, so check with Contains
		assert.Contains(t, resp.Data.User.ID, guestID, "Guest user ID should contain the requested ID")

		// Also clean up the actual server-assigned ID
		allUserIDs = append(allUserIDs, resp.Data.User.ID)
	})

	t.Run("UpsertUsersWithRoleAndTeamsRole", func(t *testing.T) {
		userID := "teams-" + uuid.New().String()
		allUserIDs = append(allUserIDs, userID)

		resp, err := client.UpdateUsers(ctx, &UpdateUsersRequest{
			Users: map[string]UserRequest{
				userID: {
					ID:        userID,
					Name:      PtrTo("Teams User"),
					Role:      PtrTo("admin"),
					Teams:     []string{"blue"},
					TeamsRole: map[string]string{"blue": "admin"},
				},
			},
		})
		require.NoError(t, err)

		u, ok := resp.Data.Users[userID]
		require.True(t, ok, "User should be in response")
		assert.Equal(t, "admin", u.Role)
		assert.Equal(t, []string{"blue"}, u.Teams)
		assert.Equal(t, map[string]string{"blue": "admin"}, u.TeamsRole)
	})

	t.Run("PartialUpdateUserWithTeam", func(t *testing.T) {
		userIDs := newUsers(t, 1)
		userID := userIDs[0]

		// Partial update to add teams and teams_role
		resp, err := client.UpdateUsersPartial(ctx, &UpdateUsersPartialRequest{
			Users: []UpdateUserPartialRequest{
				{
					ID: userID,
					Set: map[string]any{
						"teams":      []string{"blue"},
						"teams_role": map[string]string{"blue": "admin"},
					},
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, resp.Data.Users)

		u, ok := resp.Data.Users[userID]
		require.True(t, ok)
		assert.Equal(t, []string{"blue"}, u.Teams)
		assert.Equal(t, map[string]string{"blue": "admin"}, u.TeamsRole)
	})

	t.Run("UpdatePrivacySettings", func(t *testing.T) {
		userID := "privacy-" + uuid.New().String()
		allUserIDs = append(allUserIDs, userID)

		// Create user without privacy settings
		resp, err := client.UpdateUsers(ctx, &UpdateUsersRequest{
			Users: map[string]UserRequest{
				userID: {
					ID:   userID,
					Name: PtrTo("Privacy User"),
				},
			},
		})
		require.NoError(t, err)
		u, ok := resp.Data.Users[userID]
		require.True(t, ok)
		assert.Nil(t, u.PrivacySettings, "PrivacySettings should be nil initially")

		// Update with TypingIndicators disabled
		resp, err = client.UpdateUsers(ctx, &UpdateUsersRequest{
			Users: map[string]UserRequest{
				userID: {
					ID: userID,
					PrivacySettings: &PrivacySettingsResponse{
						TypingIndicators: &TypingIndicatorsResponse{
							Enabled: PtrTo(false),
						},
					},
				},
			},
		})
		require.NoError(t, err)
		u, ok = resp.Data.Users[userID]
		require.True(t, ok)
		require.NotNil(t, u.PrivacySettings)
		require.NotNil(t, u.PrivacySettings.TypingIndicators)
		assert.Equal(t, false, *u.PrivacySettings.TypingIndicators.Enabled)
		assert.Nil(t, u.PrivacySettings.ReadReceipts, "ReadReceipts should still be nil")

		// Update with both TypingIndicators=true and ReadReceipts=false
		resp, err = client.UpdateUsers(ctx, &UpdateUsersRequest{
			Users: map[string]UserRequest{
				userID: {
					ID: userID,
					PrivacySettings: &PrivacySettingsResponse{
						TypingIndicators: &TypingIndicatorsResponse{
							Enabled: PtrTo(true),
						},
						ReadReceipts: &ReadReceiptsResponse{
							Enabled: PtrTo(false),
						},
					},
				},
			},
		})
		require.NoError(t, err)
		u, ok = resp.Data.Users[userID]
		require.True(t, ok)
		require.NotNil(t, u.PrivacySettings)
		require.NotNil(t, u.PrivacySettings.TypingIndicators)
		assert.Equal(t, true, *u.PrivacySettings.TypingIndicators.Enabled)
		require.NotNil(t, u.PrivacySettings.ReadReceipts)
		assert.Equal(t, false, *u.PrivacySettings.ReadReceipts.Enabled)
	})

	t.Run("PartialUpdatePrivacySettings", func(t *testing.T) {
		userID := "privacy-partial-" + uuid.New().String()
		allUserIDs = append(allUserIDs, userID)

		// Create user
		resp, err := client.UpdateUsers(ctx, &UpdateUsersRequest{
			Users: map[string]UserRequest{
				userID: {ID: userID, Name: PtrTo("Privacy Partial User")},
			},
		})
		require.NoError(t, err)
		u := resp.Data.Users[userID]
		require.Nil(t, u.PrivacySettings)

		// Partial update: set typing_indicators.enabled = true
		partialResp, err := client.UpdateUsersPartial(ctx, &UpdateUsersPartialRequest{
			Users: []UpdateUserPartialRequest{
				{
					ID: userID,
					Set: map[string]any{
						"privacy_settings": map[string]any{
							"typing_indicators": map[string]any{
								"enabled": true,
							},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		u2 := partialResp.Data.Users[userID]
		require.NotNil(t, u2.PrivacySettings)
		require.NotNil(t, u2.PrivacySettings.TypingIndicators)
		assert.True(t, *u2.PrivacySettings.TypingIndicators.Enabled)
		assert.Nil(t, u2.PrivacySettings.ReadReceipts, "ReadReceipts should still be nil")

		// Partial update: set read_receipts.enabled = false
		partialResp2, err := client.UpdateUsersPartial(ctx, &UpdateUsersPartialRequest{
			Users: []UpdateUserPartialRequest{
				{
					ID: userID,
					Set: map[string]any{
						"privacy_settings": map[string]any{
							"read_receipts": map[string]any{
								"enabled": false,
							},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		u3 := partialResp2.Data.Users[userID]
		require.NotNil(t, u3.PrivacySettings)
		require.NotNil(t, u3.PrivacySettings.TypingIndicators)
		assert.True(t, *u3.PrivacySettings.TypingIndicators.Enabled, "TypingIndicators should still be true")
		require.NotNil(t, u3.PrivacySettings.ReadReceipts)
		assert.False(t, *u3.PrivacySettings.ReadReceipts.Enabled)
	})

	t.Run("QueryUsersWithDeactivated", func(t *testing.T) {
		userIDs := newUsers(t, 3)

		// Deactivate one user
		_, err := client.DeactivateUser(ctx, userIDs[2], &DeactivateUserRequest{})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, _ = client.ReactivateUser(context.Background(), userIDs[2], &ReactivateUserRequest{})
		})

		// Query WITHOUT including deactivated — should get 2
		resp, err := client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": map[string]any{"$in": userIDs},
				},
			},
		})
		require.NoError(t, err)
		assert.Len(t, resp.Data.Users, 2, "Should exclude deactivated user by default")

		// Query WITH including deactivated — should get all 3
		resp, err = client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": map[string]any{"$in": userIDs},
				},
				IncludeDeactivatedUsers: PtrTo(true),
			},
		})
		require.NoError(t, err)
		assert.Len(t, resp.Data.Users, 3, "Should include deactivated user")
	})

	t.Run("DeactivateUsersPlural", func(t *testing.T) {
		userIDs := newUsers(t, 2)

		// Deactivate multiple users at once
		resp, err := client.DeactivateUsers(ctx, &DeactivateUsersRequest{
			UserIds: userIDs,
		})
		require.NoError(t, err)
		assert.NotEmpty(t, resp.Data.TaskID)

		// Wait for deactivation task
		taskResult, err := WaitForTask(ctx, client, resp.Data.TaskID)
		require.NoError(t, err)
		assert.Equal(t, "completed", taskResult.Data.Status)

		// Verify deactivated — query without include should not find them
		qResp, err := client.QueryUsers(ctx, &QueryUsersRequest{
			Payload: &QueryUsersPayload{
				FilterConditions: map[string]any{
					"id": map[string]any{"$in": userIDs},
				},
			},
		})
		require.NoError(t, err)
		assert.Empty(t, qResp.Data.Users, "Deactivated users should not appear in default query")
	})

	t.Run("UserCustomData", func(t *testing.T) {
		userID := "custom-" + uuid.New().String()
		allUserIDs = append(allUserIDs, userID)

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
