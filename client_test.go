package getstream

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func cleanupAllExternalStorages(t *testing.T, client *Stream) {
//     t.Helper()
//     ctx := context.Background()

//     // List all external storages
//     response, err := client.ListExternalStorage(ctx)
//     if err != nil {
//         t.Logf("Warning: Failed to list external storages: %v", err)
//         return
//     }

//     // Delete each external storage
//     for _, storage := range response.Data.ExternalStorages {
//         _, err := client.DeleteExternalStorage(ctx, storage.Name)
//         if err != nil {
//             t.Logf("Warning: Failed to delete external storage %s: %v", storage.Name, err)
//         } else {
//             t.Logf("Successfully deleted external storage: %s", storage.Name)
//         }
//     }
// }

// func TestCleanupAllExternalStorages(t *testing.T) {
//     client := initClient(t)
//     cleanupAllExternalStorages(t, client)
// }

// initClient initializes the Stream client using environment variables.
func initClient(t *testing.T) *Stream {
	t.Helper()

	stream, err := NewStreamFromEnvVars(WithLogLevel(LogLevelError))
	require.NoError(t, err, "Failed to create client from env vars")

	return stream
}

// setup initializes the client, call object, and optionally creates a call type for each test.
// If createCallType is true, it creates a unique call type and registers a cleanup function.
func setup(t *testing.T, rm *ResourceManager, createCallType bool) (*Stream, *Call, string) {
	t.Helper()

	client := initClient(t)
	callType := "default"
	callID := randomString(10)
	callTypeName := "calltype_" + fmt.Sprintf("%s_%s", t.Name(), randomString(10))

	if createCallType {
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

		response, err := client.Video().CreateCallType(ctx, &CreateCallTypeRequest{
			Grants:               &grants,
			Name:                 callTypeName,
			Settings:             callSettings,
			NotificationSettings: notificationSettings,
		})
		require.NoError(t, err, "Failed to create call type")
		assert.Equal(t, callTypeName, response.Data.Name)
		assert.True(t, response.Data.Settings.Audio.MicDefaultOn)
		assert.Equal(t, "speaker", response.Data.Settings.Audio.DefaultDevice)
		assert.False(t, response.Data.Settings.Screensharing.AccessRequestEnabled)
		assert.True(t, response.Data.Settings.Screensharing.Enabled)
		assert.True(t, response.Data.NotificationSettings.Enabled)
		assert.False(t, response.Data.NotificationSettings.SessionStarted.Enabled)
		assert.True(t, response.Data.NotificationSettings.CallNotification.Enabled)
		assert.Equal(t, "{{ user.display_name }} invites you to a call", response.Data.NotificationSettings.CallNotification.Apns.Title)

		// Register cleanup for the created call type
		rm.RegisterCleanup(func() {
			resetSharedResource(t, client, callTypeName)
		})
	}

	call := client.Video().Call(callType, callID)

	// Register cleanup for the created call
	rm.RegisterCleanup(func() {
		resetCallResource(t, call)
	})

	return client, call, callTypeName
}

// resetSharedResource deletes the specified call type.
func resetSharedResource(t *testing.T, client *Stream, callTypeName string) {
	ctx := context.Background()
	_, err := client.Video().DeleteCallType(ctx, callTypeName)
	if err != nil {
		t.Logf("Warning: Failed to delete call type %s: %v", callTypeName, err)
	}
}

// resetCallResource deletes the specified call.
func resetCallResource(t *testing.T, call *Call) {
	ctx := context.Background()
	_, err := call.Delete(ctx, &DeleteCallRequest{})
	if err != nil {
		// Check if the error is due to the call not being found
		if !strings.Contains(err.Error(), "Can't find call with id") {
			t.Logf("Warning: Failed to delete call: %v", err)
		}
	}
}

// createExternalStorageWithCleanup creates an external storage and registers its cleanup.
func createExternalStorageWithCleanup(t *testing.T, rm *ResourceManager, client *Stream, storageName string) {
	ctx := context.Background()
	path := "test-directory/"
	s3apiKey := "test-access-key"
	s3secret := "test-secret"

	_, err := client.CreateExternalStorage(ctx, &CreateExternalStorageRequest{
		Bucket:      "test-bucket",
		Name:        storageName,
		StorageType: "s3",
		Path:        &path,
		AwsS3: &S3Request{
			S3Region: "us-east-1",
			S3ApiKey: &s3apiKey,
			S3Secret: &s3secret,
		},
	})
	require.NoError(t, err, "Failed to create external storage")

	// Register cleanup for the created external storage
	rm.RegisterCleanup(func() {
		resetExternalStorage(t, client, storageName)
	})
}

// resetExternalStorage deletes the specified external storage.
func resetExternalStorage(t *testing.T, client *Stream, storageName string) {
	ctx := context.Background()
	_, err := client.DeleteExternalStorage(ctx, storageName)
	if err != nil {
		t.Logf("Warning: Failed to delete external storage %s: %v", storageName, err)
	}
}

// TestCRUDCallTypeOperations tests Create, Read, Update, and Delete operations for call types.
func TestCRUDCallTypeOperations(t *testing.T) {
	rm := NewResourceManager(t)
	client, call, callTypeName := setup(t, rm, true)

	t.Run("Update Call Type Settings", func(t *testing.T) {
		ctx := context.Background()
		grants := map[string][]string{
			"host": {JOIN_BACKSTAGE.String()},
		}
		response, err := client.Video().UpdateCallType(ctx, callTypeName, &UpdateCallTypeRequest{
			Settings: &CallSettingsRequest{
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
			},
			Grants: &grants,
		})

		assert.NoError(t, err)
		assert.False(t, response.Data.Settings.Audio.MicDefaultOn)
		assert.Equal(t, "earpiece", response.Data.Settings.Audio.DefaultDevice)
		assert.Equal(t, "disabled", response.Data.Settings.Recording.Mode)
		assert.True(t, response.Data.Settings.Backstage.Enabled)
		assert.Equal(t, []string{JOIN_BACKSTAGE.String()}, response.Data.Grants["host"])
	})

	t.Run("Update Layout Options", func(t *testing.T) {
		ctx := context.Background()

		layoutOptions := map[string]any{
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

		_, err := client.Video().UpdateCallType(ctx, callTypeName, &UpdateCallTypeRequest{
			Settings: &CallSettingsRequest{
				Recording: &RecordSettingsRequest{
					Mode:      "available",
					AudioOnly: PtrTo(false),
					Quality:   PtrTo("1080p"),
					Layout: &LayoutSettingsRequest{
						Name:    "spotlight",
						Options: &layoutOptions,
					},
				},
			},
		})
		assert.NoError(t, err)
	})

	t.Run("Update Custom Recording Style CSS", func(t *testing.T) {
		ctx := context.Background()

		_, err := client.Video().UpdateCallType(ctx, callTypeName, &UpdateCallTypeRequest{
			Settings: &CallSettingsRequest{
				Recording: &RecordSettingsRequest{
					Mode:      "available",
					AudioOnly: PtrTo(false),
					Quality:   PtrTo("1080p"),
					Layout: &LayoutSettingsRequest{
						Name:           "spotlight",
						ExternalCssUrl: PtrTo("https://path/to/custom.css"),
					},
				},
			},
		})
		assert.NoError(t, err)
	})

	t.Run("Update Custom Recording Website", func(t *testing.T) {
		ctx := context.Background()

		_, err := client.Video().UpdateCallType(ctx, callTypeName, &UpdateCallTypeRequest{
			Settings: &CallSettingsRequest{
				Recording: &RecordSettingsRequest{
					Mode:      "available",
					AudioOnly: PtrTo(false),
					Quality:   PtrTo("1080p"),
					Layout: &LayoutSettingsRequest{
						Name:           "custom",
						ExternalAppUrl: PtrTo("https://path/to/layout/app"),
					},
				},
			},
		})
		assert.NoError(t, err)
	})

	t.Run("Read Call Type", func(t *testing.T) {
		ctx := context.Background()

		response, err := client.Video().GetCallType(ctx, callTypeName)
		assert.NoError(t, err)
		assert.Equal(t, callTypeName, response.Data.Name)
	})

	t.Run("CreatingStorageWithReservedNameShouldFail", func(t *testing.T) {
		ctx := context.Background()
		path := "directory_name/"
		s3apiKey := "my-access-key"
		s3secret := "my-secret"
		_, err := client.CreateExternalStorage(ctx, &CreateExternalStorageRequest{
			Bucket:      "my-bucket",
			Name:        "stream-s3",
			StorageType: "s3",
			Path:        &path,
			AwsS3: &S3Request{
				S3Region: "us-east-1",
				S3ApiKey: &s3apiKey,
				S3Secret: &s3secret,
			},
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "stream-s3 name reserved for internal use")
	})

	t.Run("ShouldBeAbleToListExternalStorage", func(t *testing.T) {
		ctx := context.Background()
		_, err := client.ListExternalStorage(ctx)
		require.NoError(t, err)
	})

	t.Run("Create Call", func(t *testing.T) {
		ctx := context.Background()
		callRequest := GetOrCreateCallRequest{
			Data: &CallRequest{
				CreatedByID: PtrTo("john"),
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

		c, err := call.GetOrCreate(ctx, &callRequest)
		assert.NoError(t, err)
		assert.Equal(t, "john", c.Data.Call.CreatedBy.ID)
		assert.False(t, c.Data.Call.Settings.Screensharing.Enabled)
	})

	t.Run("Update Call", func(t *testing.T) {
		ctx := context.Background()
		callRequest := UpdateCallRequest{
			SettingsOverride: &CallSettingsRequest{
				Audio: &AudioSettingsRequest{
					MicDefaultOn:  PtrTo(true),
					DefaultDevice: "speaker",
				},
			},
		}
		c, err := call.Update(ctx, &callRequest)
		assert.NoError(t, err)
		assert.True(t, c.Data.Call.Settings.Audio.MicDefaultOn)
	})
}

// TestVideoExamples tests various video-related functionalities without creating call types.
func TestVideoExamples(t *testing.T) {
	rm := NewResourceManager(t)
	client, _, _ := setup(t, rm, false)

	t.Run("Create User", func(t *testing.T) {
		ctx := context.Background()

		thierrySet := map[string]any{
			"role":    "admin",
			"country": "NL",
		}
		tommasoSet := map[string]any{
			"role":    "admin",
			"country": "US",
		}
		users := []UpdateUserPartialRequest{
			{
				ID:    "thierry-id",
				Set:   &thierrySet,
				Unset: &[]string{"custom"},
			},
			{
				ID:    "tommaso-id",
				Set:   &tommasoSet,
				Unset: &[]string{"custom"},
			},
		}
		_, err := client.UpdateUsersPartial(ctx, &UpdateUsersPartialRequest{Users: users})
		assert.NoError(t, err)

		token, err := client.CreateToken("tommaso-id", nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Create Call With Members", func(t *testing.T) {
		ctx := context.Background()
		call := newCall(t, client)
		members := []MemberRequest{
			{UserID: "thierry-id"},
			{UserID: "tommaso-id"},
		}
		callRequest := GetOrCreateCallRequest{
			Data: &CallRequest{
				CreatedByID: PtrTo("tommaso-id"),
				Members:     &members,
			},
		}
		_, err := call.GetOrCreate(ctx, &callRequest)
		assert.NoError(t, err)
	})

	t.Run("Block Unblock User From Calls", func(t *testing.T) {
		ctx := context.Background()

		call := newCall(t, client)

		badUser, err := getUser(t, client, nil, nil, nil)
		require.NoError(t, err)

		_, err = call.BlockUser(ctx, &BlockUserRequest{UserID: badUser.ID})
		assert.NoError(t, err)

		response, err := call.Get(ctx, nil)
		assert.NoError(t, err)
		assert.Contains(t, response.Data.Call.BlockedUserIDs, badUser.ID)

		_, err = call.UnblockUser(ctx, &UnblockUserRequest{UserID: badUser.ID})
		assert.NoError(t, err)

		response, err = call.Get(ctx, nil)
		assert.NoError(t, err)
		assert.NotContains(t, response.Data.Call.BlockedUserIDs, badUser.ID)
	})
}

// TestSendCustomEvent tests sending custom events within a call.
func TestSendCustomEvent(t *testing.T) {
	rm := NewResourceManager(t)
	client, call, _ := setup(t, rm, false)

	ctx := context.Background()
	user, err := getUser(t, client, nil, nil, nil)
	require.NoError(t, err)

	callRequest := GetOrCreateCallRequest{
		Data: &CallRequest{
			CreatedByID: PtrTo("tommaso-id"),
		},
	}
	_, err = call.GetOrCreate(ctx, &callRequest)
	require.NoError(t, err)

	customEvent := map[string]interface{}{
		"bananas": "good",
	}
	sendEventRequest := SendCallEventRequest{
		UserID: &user.ID,
		Custom: &customEvent,
	}
	_, err = call.SendCallEvent(ctx, &sendEventRequest)
	assert.NoError(t, err)
}

// TestMuteAll tests muting all users in a call.
func TestMuteAll(t *testing.T) {
	rm := NewResourceManager(t)
	_, call, _ := setup(t, rm, false)

	ctx := context.Background()
	userID := randomString(10)

	callRequest := GetOrCreateCallRequest{
		Data: &CallRequest{
			CreatedByID: PtrTo(userID),
		},
	}
	_, err := call.GetOrCreate(ctx, &callRequest)
	require.NoError(t, err)

	muteRequest := MuteUsersRequest{
		MutedByID:    &userID,
		MuteAllUsers: PtrTo(true),
		Audio:        PtrTo(true),
	}
	_, err = call.MuteUsers(ctx, &muteRequest)
	assert.NoError(t, err)
}

// TestVideoExamplesAdditional tests additional video-related functionalities.
func TestVideoExamplesAdditional(t *testing.T) {
	rm := NewResourceManager(t)
	client, call, _ := setup(t, rm, false)

	t.Run("MuteSomeUsers", func(t *testing.T) {
		ctx := context.Background()
		alice, err := getUser(t, client, nil, nil, nil)
		require.NoError(t, err)
		bob, err := getUser(t, client, nil, nil, nil)
		require.NoError(t, err)

		userID := randomString(10)
		_, err = call.GetOrCreate(ctx, &GetOrCreateCallRequest{
			Data: &CallRequest{
				CreatedByID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		_, err = call.MuteUsers(ctx, &MuteUsersRequest{
			MutedByID:        &userID,
			UserIDs:          &[]string{alice.ID, bob.ID},
			Audio:            PtrTo(true),
			Video:            PtrTo(true),
			Screenshare:      PtrTo(true),
			ScreenshareAudio: PtrTo(true),
		})
		assert.NoError(t, err)
	})

	t.Run("UpdateUserPermissions", func(t *testing.T) {
		ctx := context.Background()
		userID := randomString(10)
		_, err := call.GetOrCreate(ctx, &GetOrCreateCallRequest{
			Data: &CallRequest{
				CreatedByID: PtrTo(userID),
			},
		})
		require.NoError(t, err)

		alice, err := getUser(t, client, nil, nil, nil)
		require.NoError(t, err)

		_, err = call.UpdateUserPermissions(ctx, &UpdateUserPermissionsRequest{
			UserID:            alice.ID,
			RevokePermissions: &[]string{SEND_AUDIO.String()},
		})
		assert.NoError(t, err)

		_, err = call.UpdateUserPermissions(ctx, &UpdateUserPermissionsRequest{
			UserID:           alice.ID,
			GrantPermissions: &[]string{SEND_AUDIO.String()},
		})
		assert.NoError(t, err)
	})

	t.Run("DeactivateUser", func(t *testing.T) {
		ctx := context.Background()
		alice, err := getUser(t, client, nil, nil, nil)
		require.NoError(t, err)
		bob, err := getUser(t, client, nil, nil, nil)
		require.NoError(t, err)

		_, err = client.DeactivateUser(ctx, alice.ID, &DeactivateUserRequest{})
		assert.NoError(t, err)

		_, err = client.ReactivateUser(ctx, alice.ID, &ReactivateUserRequest{})
		assert.NoError(t, err)

		response, err := client.DeactivateUsers(ctx, &DeactivateUsersRequest{UserIDs: []string{alice.ID, bob.ID}})
		assert.NoError(t, err)
		taskID := response.Data.TaskID

		// Note: In a real scenario, you might want to implement a retry mechanism or use a channel to wait for the task to complete
		time.Sleep(2 * time.Second)

		taskStatus, err := client.GetTask(ctx, taskID)
		assert.NoError(t, err)

		if taskStatus.Data.Status == "completed" {
			t.Logf("Task result: %v", taskStatus.Data.Result)
		}
	})

	t.Run("CreateCallWithSessionTimer", func(t *testing.T) {
		ctx := context.Background()
		userID := randomString(10)
		response, err := call.GetOrCreate(ctx, &GetOrCreateCallRequest{
			Data: &CallRequest{
				CreatedByID: PtrTo(userID),
				SettingsOverride: &CallSettingsRequest{
					Limits: &LimitsSettingsRequest{
						MaxDurationSeconds: PtrTo(3600),
					},
				},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, PtrTo(3600), response.Data.Call.Settings.Limits.MaxDurationSeconds)

		res, err := call.Update(ctx, &UpdateCallRequest{
			SettingsOverride: &CallSettingsRequest{
				Limits: &LimitsSettingsRequest{
					MaxDurationSeconds: PtrTo(7200),
				},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, PtrTo(7200), res.Data.Call.Settings.Limits.MaxDurationSeconds)

		res2, err := call.Update(ctx, &UpdateCallRequest{
			SettingsOverride: &CallSettingsRequest{
				Limits: &LimitsSettingsRequest{
					MaxDurationSeconds: PtrTo(0),
				},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, PtrTo(0), res2.Data.Call.Settings.Limits.MaxDurationSeconds)
	})

	t.Run("UserBlocking", func(t *testing.T) {
		t.Skip()
		ctx := context.Background()
		alice, err := getUser(t, client, nil, nil, nil)
		require.NoError(t, err)
		bob, err := getUser(t, client, nil, nil, nil)
		require.NoError(t, err)

		_, err = client.BlockUsers(ctx, &BlockUsersRequest{BlockedUserID: bob.ID, UserID: &alice.ID})

		response, err := client.GetBlockedUsers(ctx, &GetBlockedUsersRequest{UserID: &alice.ID})
		assert.NoError(t, err)
		assert.Len(t, response.Data.Blocks, 1)
		assert.Equal(t, alice.ID, response.Data.Blocks[0].UserID)
		assert.Equal(t, bob.ID, response.Data.Blocks[0].BlockedUserID)

		_, err = client.UnblockUsers(ctx, &UnblockUsersRequest{BlockedUserID: bob.ID, UserID: &alice.ID})

		response, err = client.GetBlockedUsers(ctx, &GetBlockedUsersRequest{UserID: &alice.ID})
		assert.NoError(t, err)
		assert.Empty(t, response.Data.Blocks)
	})

	t.Run("CreateCallWithBackstageAndJoinAheadSet", func(t *testing.T) {
		ctx := context.Background()
		userID := randomString(10)
		startsAt := time.Now().Add(30 * time.Minute)
		ts := NewTimestamp(startsAt, true)

		response, err := call.GetOrCreate(ctx, &GetOrCreateCallRequest{
			Data: &CallRequest{
				StartsAt:    &ts,
				CreatedByID: PtrTo(userID),
				SettingsOverride: &CallSettingsRequest{
					Backstage: &BackstageSettingsRequest{
						Enabled:              PtrTo(true),
						JoinAheadTimeSeconds: PtrTo(300),
					},
				},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, PtrTo(300), response.Data.Call.JoinAheadTimeSeconds)

		res, err := call.Update(ctx, &UpdateCallRequest{
			SettingsOverride: &CallSettingsRequest{
				Backstage: &BackstageSettingsRequest{
					JoinAheadTimeSeconds: PtrTo(600),
				},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, PtrTo(600), res.Data.Call.JoinAheadTimeSeconds)

		res2, err := call.Update(ctx, &UpdateCallRequest{
			SettingsOverride: &CallSettingsRequest{
				Backstage: &BackstageSettingsRequest{
					JoinAheadTimeSeconds: PtrTo(0),
				},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, PtrTo(0), res2.Data.Call.JoinAheadTimeSeconds)
	})
}

// TestDeleteCall tests the soft deletion of a call.
func TestDeleteCall(t *testing.T) {
	client := initClient(t)
	ctx := context.Background()
	call := client.Video().Call("default", randomString(10))

	t.Run("SoftDelete", func(t *testing.T) {
		_, err := call.GetOrCreate(ctx, &GetOrCreateCallRequest{
			Data: &CallRequest{
				CreatedByID: PtrTo("john"),
			},
		})
		require.NoError(t, err)

		response, err := call.Delete(ctx, &DeleteCallRequest{})
		require.NoError(t, err)
		assert.NotNil(t, response.Data.Call)
		assert.Nil(t, response.Data.TaskID)

		_, err = call.Get(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Can't find call with id")
	})
}

// TestTeams tests functionalities related to teams.
func TestTeams(t *testing.T) {
	client := initClient(t)
	ctx := context.Background()

	callID := randomString(10)
	userID := randomString(10)

	_, err := client.UpdateUsers(ctx, &UpdateUsersRequest{
		Users: map[string]UserRequest{
			userID: {
				ID:    userID,
				Teams: &[]string{"red", "blue"},
			},
		},
	})
	require.NoError(t, err)

	call := client.Video().Call("default", callID)
	response, err := call.GetOrCreate(ctx, &GetOrCreateCallRequest{
		Data: &CallRequest{
			CreatedByID: PtrTo(userID),
			Team:        PtrTo("blue"),
		},
	})
	require.NoError(t, err)
	assert.Equal(t, PtrTo("blue"), response.Data.Call.Team)

	usersResponse, err := client.QueryUsers(ctx, &QueryUsersRequest{
		Payload: &QueryUsersPayload{FilterConditions: map[string]interface{}{
			"id":    userID,
			"teams": map[string]interface{}{"$in": []string{"red", "blue"}},
		}},
	})
	require.NoError(t, err)
	assert.Greater(t, len(usersResponse.Data.Users), 0)
	userIDs := make([]string, 0, len(usersResponse.Data.Users))
	for _, user := range usersResponse.Data.Users {
		userIDs = append(userIDs, user.ID)
	}
	assert.Contains(t, userIDs, userID)

	usersResponse, err = client.QueryUsers(ctx, &QueryUsersRequest{
		Payload: &QueryUsersPayload{FilterConditions: map[string]interface{}{
			"teams": nil,
		}},
	})
	require.NoError(t, err)
	for _, user := range usersResponse.Data.Users {
		assert.Empty(t, user.Teams)
	}

	callsResponse, err := client.Video().QueryCalls(ctx, &QueryCallsRequest{
		FilterConditions: &map[string]interface{}{
			"id":   callID,
			"team": map[string]interface{}{"$eq": "blue"},
		},
	})
	require.NoError(t, err)
	assert.Greater(t, len(callsResponse.Data.Calls), 0)
}

func TestExternalStorageOperations(t *testing.T) {
	rm := NewResourceManager(t)
	client, _, _ := setup(t, rm, false)
	ctx := context.Background()
	// Create a unique name for the test storage
	uniqueName := "test-storage-" + randomString(10)

	defer func() {
		resetExternalStorage(t, client, uniqueName)
	}()
	t.Run("CreateExternalStorage", func(t *testing.T) {
		path := "test-directory/"
		s3apiKey := "test-access-key"
		s3secret := "test-secret"

		_, err := client.CreateExternalStorage(ctx, &CreateExternalStorageRequest{
			Bucket:      "test-bucket",
			Name:        uniqueName,
			StorageType: "s3",
			Path:        &path,
			AwsS3: &S3Request{
				S3Region: "us-east-1",
				S3ApiKey: &s3apiKey,
				S3Secret: &s3secret,
			},
		})
		require.NoError(t, err)
	})

	t.Run("ListExternalStorage", func(t *testing.T) {
		response, err := client.ListExternalStorage(ctx)
		require.NoError(t, err)
		assert.NotEmpty(t, response.Data.ExternalStorages, "External storages list should not be empty")

		// Check if our created storage is in the list
		found := false
		for _, storage := range response.Data.ExternalStorages {
			if storage.Name == uniqueName {
				found = true
				break
			}
		}
		assert.True(t, found, "Created external storage should be in the list")
	})
}

func TestEnableCallRecordingAndBackstageMode(t *testing.T) {
	rm := NewResourceManager(t)
	_, call, _ := setup(t, rm, false)
	ctx := context.Background()
	// Create the call first
	_, err := call.GetOrCreate(ctx, &GetOrCreateCallRequest{
		Data: &CallRequest{
			CreatedByID: PtrTo("test-user"),
		},
	})
	require.NoError(t, err)
	t.Run("EnableCallRecording", func(t *testing.T) {
		response, err := call.Update(ctx, &UpdateCallRequest{
			SettingsOverride: &CallSettingsRequest{
				Recording: &RecordSettingsRequest{
					Mode: "available",
				},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, "available", response.Data.Call.Settings.Recording.Mode)
	})

	t.Run("EnableBackstageMode", func(t *testing.T) {
		response, err := call.Update(ctx, &UpdateCallRequest{
			SettingsOverride: &CallSettingsRequest{
				Backstage: &BackstageSettingsRequest{
					Enabled: PtrTo(true),
				},
			},
		})
		require.NoError(t, err)
		assert.True(t, response.Data.Call.Settings.Backstage.Enabled)
	})
}

func TestDeleteRecordingsAndTranscriptions(t *testing.T) {
	rm := NewResourceManager(t)
	_, call, _ := setup(t, rm, false)
	ctx := context.Background()

	t.Run("DeleteNonExistentRecording", func(t *testing.T) {
		_, err := call.DeleteRecording(ctx, "non-existent-session", "non-existent-filename")
		require.Error(t, err)
	})

	t.Run("DeleteNonExistentTranscription", func(t *testing.T) {
		_, err := call.DeleteTranscription(ctx, "non-existent-session", "non-existent-filename")
		require.Error(t, err)
	})
}

func TestHardDeleteCall(t *testing.T) {
	rm := NewResourceManager(t)
	client, call, _ := setup(t, rm, false)
	ctx := context.Background()

	t.Run("HardDelete", func(t *testing.T) {
		_, err := call.GetOrCreate(ctx, &GetOrCreateCallRequest{
			Data: &CallRequest{
				CreatedByID: PtrTo("test-user"),
			},
		})
		require.NoError(t, err)

		response, err := call.Delete(ctx, &DeleteCallRequest{
			Hard: PtrTo(true),
		})
		require.NoError(t, err)
		assert.NotNil(t, response.Data.Call)
		assert.NotNil(t, response.Data.TaskID)
		time.Sleep(2 * time.Second)
		// Wait for the task to complete
		taskStatus, err := client.GetTask(ctx, *response.Data.TaskID)
		require.NoError(t, err)

		assert.Equal(t, "completed", taskStatus.Data.Status)

		// Verify the call is deleted
		_, err = call.Get(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Can't find call with id")
	})
}
