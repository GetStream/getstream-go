package getstream

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initVideoClient(t *testing.T) *VideoClient {
	t.Helper()

	c, err := NewClientFromEnvVars()
	require.NoError(t, err, "new client")
	video := NewVideoClient(c)

	return video
}

func TestCreateCall(t *testing.T) {
	client := initVideoClient(t)
	ctx := context.Background()
	callRequest := GetOrCreateCallRequest{
		Data: &CallRequest{
			CreatedById: PtrTo("john"),
			SettingsOverride: &CallSettingsRequest{
				Geofencing: &GeofenceSettingsRequest{
					Names: PtrTo([]string{"canada"}),
				},
				Screensharing: &ScreensharingSettingsRequest{
					Enabled: PtrTo(false),
				},
			},
		},
	}

	callType := "default"
	callID := randomString(10)
	call := client.Call(ctx, callType, callID)
	c, err := call.GetOrCreate(ctx, callRequest)
	assert.NoError(t, err)
	assert.Equal(t, "john", c.Call.CreatedBy.Id)
	assert.False(t, c.Call.Settings.Screensharing.Enabled)
}
