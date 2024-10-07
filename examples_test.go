package getstream_test

import (
	"context"
	"testing"
	"time"

	"github.com/GetStream/getstream-go"
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
		//optionally you can create a new client with a different timeout
		client, err := getstream.NewClient(apiKey, apiSecret, getstream.WithTimeout(10_000*time.Millisecond))
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
				Teams:  []string{"red", "blue"},
				Name:   getstream.PtrTo("tommaso"),
				Role:   getstream.PtrTo("admin"),
				Custom: map[string]any{"country": "NL"},
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, response)

	// the token will be valid for 1 hour
	client.CreateToken("tommaso-id", nil, 3600)
}
