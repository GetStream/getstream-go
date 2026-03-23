package getstream_test

import (
	"context"
	"strings"
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

	t.Run("GetRetentionPolicy", func(t *testing.T) {
		policyName := "get-test-policy"
		maxAge := 720

		// Setup: create a retention policy so we can verify it appears in the list
		_, err := client.Chat().SetRetentionPolicy(ctx, &SetRetentionPolicyRequest{
			Policy:      PtrTo(policyName),
			MaxAgeHours: PtrTo(maxAge),
		})
		if err != nil {
			if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "Not Found") {
				t.Skip("Retention policy endpoints not available on this environment")
			}
			require.NoError(t, err)
		}
		defer func() {
			_, _ = client.Chat().DeleteRetentionPolicy(ctx, &DeleteRetentionPolicyRequest{
				Policy: PtrTo(policyName),
			})
		}()

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
	})

	t.Run("GetRetentionPolicyRuns", func(t *testing.T) {
		// List retention policy runs (may be empty, but should not error)
		runsResp, err := client.Chat().GetRetentionPolicyRuns(ctx, &GetRetentionPolicyRunsRequest{})
		if err != nil {
			if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "Not Found") {
				t.Skip("Retention policy endpoints not available on this environment")
			}
			require.NoError(t, err)
		}
		assert.NotNil(t, runsResp.Data.Runs)
	})

}
