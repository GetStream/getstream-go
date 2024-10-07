package getstream_test

import (
	"context"
	"testing"
	"time"

	"github.com/GetStream/getstream-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	client *getstream.Client
	ctx    context.Context
)

func init() {
	ctx = context.Background()
}

func TestCreatingClient(t *testing.T) {
	apiKey := "key"
	apiSecret := "secret"
	client, err := getstream.NewClient(apiKey, apiSecret)
	require.NoError(t, err)
	require.NotNil(t, client)

	{
		// optionally you can create a new client with a different timeout
		client, err := getstream.NewClient(
			apiKey, apiSecret,
			getstream.WithTimeout(10_000*time.Millisecond),
		)
		require.NoError(t, err)
		require.NotNil(t, client)
	}
}

func TestCreateUserAndToken(t *testing.T) {
	client := initClient(t)

	// optional values are passed as pointers, you can use `getstream.PtrTo` to get pointers from literals of any type
	response, err := client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
		Users: map[string]getstream.UserRequest{
			"user-id": {
				ID:     "user-id",
				Name:   getstream.PtrTo("tommaso"),
				Role:   getstream.PtrTo("admin"),
				Custom: map[string]any{"country": "NL"},
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, response)

	token, err := client.CreateToken("tommaso-id", getstream.WithExpiration(24*time.Hour))
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestCreateCall(t *testing.T) {
	client := initClient(t)

	call := client.Video().Call("default", uuid.NewString())

	members := []getstream.MemberRequest{
		{UserID: "john", Role: getstream.PtrTo("admin")},
		{UserID: "jack"},
	}

	callRequest := getstream.GetOrCreateCallRequest{
		Data: &getstream.CallRequest{
			CreatedByID: getstream.PtrTo("sacha"),
			Members:     members,
			Custom: map[string]any{
				"color": "blue",
			},
		},
	}

	response, err := call.GetOrCreate(ctx, &callRequest)
	require.NoError(t, err)
	require.NotEmpty(t, response)

	// pre-create the users
	{
		response, err := client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
			Users: map[string]getstream.UserRequest{
				"sara": {
					ID: "sara",
				},
				"emily": {
					ID: "emily",
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, response)
	}

	// call members example:
	{
		// Call members need to be existing users (use `client.UpdateUsers` for that)
		// You can also update the role of existing members
		response, err := call.UpdateCallMembers(ctx, &getstream.UpdateCallMembersRequest{
			UpdateMembers: []getstream.MemberRequest{
				{UserID: "sara"},
				{UserID: "emily", Role: getstream.PtrTo("admin")},
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, response)
	}

	// call members remove example:
	{
		response, err := call.UpdateCallMembers(ctx, &getstream.UpdateCallMembersRequest{
			RemoveMembers: []string{
				"sara",
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, response)
	}

	// updating a call
	{
		// update some custom data for this call
		response, err := call.Update(ctx, &getstream.UpdateCallRequest{
			Custom: map[string]any{"color": "red"},
		})

		// update settings for this call
		response, err = call.Update(ctx, &getstream.UpdateCallRequest{
			SettingsOverride: &getstream.CallSettingsRequest{
				Screensharing: &getstream.ScreensharingSettingsRequest{
					Enabled:              getstream.PtrTo(true),
					AccessRequestEnabled: getstream.PtrTo(true),
				},
			},
		})

		require.NoError(t, err)
		require.NotEmpty(t, response)
		assert.Equal(t, map[string]any{"color": "red"}, response.Data.Call.Custom)
		assert.True(t, response.Data.Call.Settings.Screensharing.Enabled)
		assert.True(t, response.Data.Call.Settings.Screensharing.AccessRequestEnabled)
	}
}

func TestUsers(t *testing.T) {
	client := initClient(t)

	response, err := client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
		Users: map[string]getstream.UserRequest{
			"user_id": {
				ID:   "user_id",
				Role: getstream.PtrTo("admin"),
				Custom: map[string]interface{}{
					"color": "red",
				},
				Name:  getstream.PtrTo("This is a test user"),
				Image: getstream.PtrTo("link/to/profile/image"),
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, response)

	{
		response, err := client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
			Users: map[string]getstream.UserRequest{
				"user_id": {
					ID:   "user_id",
					Role: getstream.PtrTo("user"),
					Custom: map[string]interface{}{
						"color": "blue",
					},
					Name:  getstream.PtrTo("This is a test user"),
					Image: getstream.PtrTo("link/to/profile/image"),
				},
			},
		})

		// or
		response, err = client.UpdateUsersPartial(ctx, &getstream.UpdateUsersPartialRequest{
			Users: []getstream.UpdateUserPartialRequest{
				{
					ID: "user_id",
					Set: map[string]interface{}{
						"new-field": "value",
					},
					Unset: []string{"name"},
				},
			},
		})

		require.NoError(t, err)
		require.NotEmpty(t, response)
		require.Empty(t, response.Data.Users["user_id"].Name)
		require.Equal(t, map[string]interface{}{
			"color":     "blue",
			"new-field": "value",
		}, response.Data.Users["user_id"].Custom)
	}

	aliceID := uuid.New().String()
	bobID := uuid.New().String()

	// pre-create the users
	{
		response, err := client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
			Users: map[string]getstream.UserRequest{
				aliceID: {
					ID: aliceID,
				},
				bobID: {
					ID: bobID,
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, response)
	}

	{
		// deactivate one user
		response, err := client.DeactivateUser(ctx, aliceID, &getstream.DeactivateUserRequest{})
		require.NoError(t, err)
		require.NotEmpty(t, response)

		// reactivates the user
		_, err = client.ReactivateUser(ctx, aliceID, &getstream.ReactivateUserRequest{})
		require.NoError(t, err)

		// deactivates users in bulk, this is an async operation
		_, err = client.DeactivateUsers(ctx, &getstream.DeactivateUsersRequest{
			UserIds: []string{aliceID, bobID},
		})
		require.NoError(t, err)
	}

	// deleting users
	{
		response, err := client.DeleteUsers(ctx, &getstream.DeleteUsersRequest{UserIds: []string{"<id>"}})

		// restore a soft-deleted user
		_, err = client.RestoreUsers(ctx, &getstream.RestoreUsersRequest{UserIds: []string{"<id>"}})

		require.NoError(t, err)
		require.NotEmpty(t, response)
	}
}

func TestCallToken(t *testing.T) {
	client := initClient(t)

	// the list of call IDs this token applies to
	tokenClaims := getstream.Claims{CallCIDs: []string{"default:call1", "livestream:call2"}}

	token, err := client.CreateToken("john",
		getstream.WithClaims(tokenClaims),
		getstream.WithExpiration(24*time.Hour),
	)
	require.NoError(t, err)
	require.NotEmpty(t, token)
}
