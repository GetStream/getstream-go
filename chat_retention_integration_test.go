package getstream_test

import (
	"context"
	"strings"
	"testing"

	. "github.com/GetStream/getstream-go/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatRetentionPolicyRunsIntegration(t *testing.T) {
	t.Parallel()
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	// List retention policy runs (may be empty, but should not error)
	runsResp, err := client.Chat().GetRetentionPolicyRuns(ctx, &GetRetentionPolicyRunsRequest{
		Limit: PtrTo(10),
	})
	if err != nil {
		if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "Not Found") {
			t.Skip("Retention policy endpoints not available on this environment")
		}
		require.NoError(t, err)
	}
	assert.NotNil(t, runsResp.Data.Runs)
}
