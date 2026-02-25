package getstream_test

import (
	"context"
	"testing"

	"github.com/GetStream/getstream-go/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoIntegration(t *testing.T) {
	t.Parallel()
	client, err := getstream.NewClientFromEnvVars()
	require.NoError(t, err, "Failed to create client")

	t.Run("SRT URL", func(t *testing.T) {
		userID := uuid.NewString()
		call := client.Video().Call("default", "my-livestream")
		_, err := call.GetOrCreate(context.Background(), &getstream.GetOrCreateCallRequest{
			Data: &getstream.CallRequest{
				CreatedByID: &userID,
			},
		})
		require.NoError(t, err)
		credentials, err := call.CreateSRTCredentials("host-user-id")
		require.NoError(t, err)
		require.NotEmpty(t, credentials.Address)
	})
}
