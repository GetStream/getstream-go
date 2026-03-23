package getstream_test

import (
	"context"
	"testing"

	. "github.com/GetStream/getstream-go/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatRetentionPolicyIntegration(t *testing.T) {
	t.Parallel()
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	t.Run("SetGetDeleteRetentionPolicy", func(t *testing.T) {
		policyName := "old-messages"
		maxAge := 720

		// Set a retention policy
		setResp, err := client.Chat().SetRetentionPolicy(ctx, &SetRetentionPolicyRequest{
			Policy:      PtrTo(policyName),
			MaxAgeHours: PtrTo(maxAge),
		})
		require.NoError(t, err)
		assert.Equal(t, policyName, setResp.Data.Policy.Policy)
		assert.Equal(t, maxAge, setResp.Data.Policy.Config.MaxAgeHours)

		// Get retention policies and verify ours is present
		getResp, err := client.Chat().GetRetentionPolicy(ctx, &GetRetentionPolicyRequest{})
		require.NoError(t, err)

		found := false
		for _, p := range getResp.Data.Policies {
			if p.Policy == policyName {
				found = true
				assert.Equal(t, maxAge, p.Config.MaxAgeHours)
			}
		}
		assert.True(t, found, "Created retention policy should appear in list")

		// Update the retention policy with a new max age
		updatedMaxAge := 360
		updateResp, err := client.Chat().SetRetentionPolicy(ctx, &SetRetentionPolicyRequest{
			Policy:      PtrTo(policyName),
			MaxAgeHours: PtrTo(updatedMaxAge),
		})
		require.NoError(t, err)
		assert.Equal(t, policyName, updateResp.Data.Policy.Policy)
		assert.Equal(t, updatedMaxAge, updateResp.Data.Policy.Config.MaxAgeHours)

		// Delete the retention policy
		_, err = client.Chat().DeleteRetentionPolicy(ctx, &DeleteRetentionPolicyRequest{
			Policy: PtrTo(policyName),
		})
		require.NoError(t, err)

		// Verify it was deleted
		getResp, err = client.Chat().GetRetentionPolicy(ctx, &GetRetentionPolicyRequest{})
		require.NoError(t, err)

		for _, p := range getResp.Data.Policies {
			assert.NotEqual(t, policyName, p.Policy, "Deleted retention policy should not appear in list")
		}
	})

	t.Run("GetRetentionPolicyRuns", func(t *testing.T) {
		// List retention policy runs (may be empty, but should not error)
		runsResp, err := client.Chat().GetRetentionPolicyRuns(ctx, &GetRetentionPolicyRunsRequest{})
		require.NoError(t, err)
		assert.NotNil(t, runsResp.Data.Runs)
	})

	t.Run("GetRetentionPolicyRunsWithPagination", func(t *testing.T) {
		runsResp, err := client.Chat().GetRetentionPolicyRuns(ctx, &GetRetentionPolicyRunsRequest{
			Limit:  PtrTo(5),
			Offset: PtrTo(0),
		})
		require.NoError(t, err)
		assert.NotNil(t, runsResp.Data.Runs)
	})
}
