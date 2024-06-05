package getstream

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initVideoClient(t *testing.T) *VideoClient {
	t.Helper()

	c, err := NewClientFromEnvVars()
	require.NoError(t, err, "Failed to create client from env vars")
	video := NewVideoClient(c)

	return video
}

var (
	once         sync.Once
	call         *Call
	callTypeName string
	client       *VideoClient
	setupErr     error
)

// setup initializes the call object etc once for all tests
func setup(t *testing.T) {
	once.Do(func() {
		callType := "default"
		callID := randomString(10)
		callTypeName = "calltype" + randomString(10)
		var err error

		client = initVideoClient(t)
		call = client.Call(callType, callID)
		setupErr = err
	})
	if setupErr != nil {
		panic("Failed to setup call object for testing: " + setupErr.Error())
	}
	t.Cleanup(resetSharedResource)
}
func resetSharedResource() {
	once = sync.Once{}
	client = nil
	call = nil
	callTypeName = ""
	setupErr = nil
}

func TestCRUDCallOperations(t *testing.T) {
	setup(t)

	t.Run("Create", func(t *testing.T) {
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

		c, err := call.GetOrCreate(ctx, callRequest)
		assert.NoError(t, err)
		assert.Equal(t, "john", c.Call.CreatedBy.Id)
		assert.False(t, c.Call.Settings.Screensharing.Enabled)
	})

	t.Run("Update", func(t *testing.T) {

		ctx := context.Background()
		callRequest := UpdateCallRequest{
			SettingsOverride: &CallSettingsRequest{
				Audio: &AudioSettingsRequest{
					MicDefaultOn:  PtrTo(true),
					DefaultDevice: "speaker",
				},
			},
		}
		c, err := call.Update(ctx, callRequest)
		assert.NoError(t, err)
		assert.True(t, c.Call.Settings.Audio.MicDefaultOn)
	})
}

func TestCRUDCallTypeOperations(t *testing.T) {
	setup(t)

	t.Run("Create", func(t *testing.T) {
		ctx := context.Background()

		callSettings := &CallSettingsRequest{
			Audio: &AudioSettingsRequest{
				DefaultDevice: "speaker",
				MicDefaultOn:  PtrTo(true),
			},
			Screensharing: &ScreensharingSettingsRequest{
				AccessRequestEnabled: PtrTo(false),
				Enabled:              PtrTo(true),
			},
		}

		notificationSettings := &NotificationSettings{
			Enabled: true,
			CallNotification: EventNotificationSettings{
				Apns: APNS{
					Title: "{{ user.display_name }} invites you to a call",
					Body:  "",
				},
				Enabled: true,
			},
			SessionStarted: EventNotificationSettings{
				Apns: APNS{
					Body:  "",
					Title: "{{ user.display_name }} invites you to a call",
				},
				Enabled: false,
			},
			CallLiveStarted: EventNotificationSettings{
				Apns: APNS{
					Body:  "",
					Title: "{{ user.display_name }} invites you to a call",
				},
				Enabled: false,
			},
			CallRing: EventNotificationSettings{
				Apns: APNS{
					Body:  "",
					Title: "{{ user.display_name }} invites you to a call",
				},
				Enabled: false,
			},
		}

		grants := map[string][]string{
			"admin": {
				SEND_AUDIO.String(),
				SEND_VIDEO.String(),
				MUTE_USERS.String(),
			},
			"user": {
				SEND_AUDIO.String(),
				SEND_VIDEO.String(),
			},
		}

		response, err := client.CreateCallType(ctx, CreateCallTypeRequest{Grants: &grants, Name: callTypeName, Settings: callSettings, NotificationSettings: notificationSettings})
		assert.NoError(t, err)
		assert.Equal(t, callTypeName, response.Name)
		assert.True(t, response.Settings.Audio.MicDefaultOn)
		assert.Equal(t, "speaker", response.Settings.Audio.DefaultDevice)
		assert.False(t, response.Settings.Screensharing.AccessRequestEnabled)
		assert.True(t, response.Settings.Screensharing.Enabled)
		assert.True(t, response.NotificationSettings.Enabled)
		assert.False(t, response.NotificationSettings.SessionStarted.Enabled)
		assert.True(t, response.NotificationSettings.CallNotification.Enabled)
		assert.Equal(t, "{{ user.display_name }} invites you to a call", response.NotificationSettings.CallNotification.Apns.Title)
	})

	t.Run("Update", func(t *testing.T) {
		ctx := context.Background()
		grants := map[string][]string{
			"host": {JOIN_BACKSTAGE.String()},
		}
		response, err := client.UpdateCallType(ctx, callTypeName, UpdateCallTypeRequest{Settings: &CallSettingsRequest{
			Audio: &AudioSettingsRequest{
				DefaultDevice: "earpiece",
				MicDefaultOn:  PtrTo(false),
			},
			Recording: &RecordSettingsRequest{
				Mode: "disabled",
			},
			Backstage: &BackstageSettingsRequest{
				Enabled: PtrTo(true),
			},
		}, Grants: &grants})

		assert.NoError(t, err)
		assert.False(t, response.Settings.Audio.MicDefaultOn)
		assert.Equal(t, "earpiece", response.Settings.Audio.DefaultDevice)
		assert.Equal(t, "disabled", response.Settings.Recording.Mode)
		assert.True(t, response.Settings.Backstage.Enabled)
		assert.Equal(t, []string{JOIN_BACKSTAGE.String()}, response.Grants["host"])
	})
	t.Run("Read", func(t *testing.T) {
		ctx := context.Background()

		response, err := client.GetCallType(ctx, callTypeName)
		assert.NoError(t, err)
		assert.Equal(t, callTypeName, response.Name)
	})

	t.Run("Update", func(t *testing.T) {
		ctx := context.Background()

		layoutOptions := map[string]string{
			"logo.image_url":                             "https://theme.zdassets.com/theme_assets/9442057/efc3820e436f9150bc8cf34267fff4df052a1f9c.png",
			"logo.horizontal_position":                   "center",
			"title.text":                                 "Building Stream Video Q&A",
			"title.horizontal_position":                  "center",
			"title.color":                                "black",
			"participant_label.border_radius":            "0px",
			"participant.border_radius":                  "0px",
			"layout.spotlight.participants_bar_position": "top",
			"layout.background_color":                    "#f2f2f2",
			"participant.placeholder_background_color":   "#1f1f1f",
			"layout.single-participant.padding_inline":   "20%",
			"participant_label.background_color":         "transparent",
		}

		_, err := client.UpdateCallType(ctx, callTypeName, UpdateCallTypeRequest{Settings: &CallSettingsRequest{
			Recording: &RecordSettingsRequest{
				Mode:      "available",
				AudioOnly: PtrTo(false),
				Quality:   PtrTo("1080p"),
				Layout: &LayoutSettingsRequest{
					Name:    "spotlight",
					Options: &layoutOptions,
				},
			},
		}})
		assert.NoError(t, err)
	})

	t.Run("Update Custom Recording Style css", func(t *testing.T) {
		ctx := context.Background()

		_, err := client.UpdateCallType(ctx, callTypeName, UpdateCallTypeRequest{Settings: &CallSettingsRequest{
			Recording: &RecordSettingsRequest{
				Mode:      "available",
				AudioOnly: PtrTo(false),
				Quality:   PtrTo("1080p"),
				Layout: &LayoutSettingsRequest{
					Name:           "spotlight",
					ExternalCssUrl: PtrTo("https://path/to/custom.css"),
				},
			},
		}})
		assert.NoError(t, err)
	})

	t.Run("Update Custom Recording Website", func(t *testing.T) {
		
		ctx := context.Background()

		_, err := client.UpdateCallType(ctx, callTypeName, UpdateCallTypeRequest{Settings: &CallSettingsRequest{
			Recording: &RecordSettingsRequest{
				Mode:      "available",
				AudioOnly: PtrTo(false),
				Quality:   PtrTo("1080p"),
				Layout: &LayoutSettingsRequest{
					Name:           "custom",
					ExternalAppUrl: PtrTo("https://path/to/layout/app"),
				},
			},
		}})
		assert.NoError(t, err)
	})
}

// func TestDeleteCallType(t *testing.T) {
// 	setup(t)
// 	ctx := context.Background()
// 	_, err := client.DeleteCallType(ctx, callTypeName)

// 	assert.NoError(t, err)
// }
