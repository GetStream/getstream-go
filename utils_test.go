package getstream

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func PtrTo[T any](v T) *T {
	return &v
}

func randomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = byte(65 + r.Intn(26)) // should be 26 to include 'Z'
	}
	return string(bytes)
}

func newCall(t *testing.T, client *Stream) *Call {
	t.Helper()
	ctx := context.Background()
	callID := uuid.New().String()
	call := client.Video().Call("default", callID)
	callRequest := GetOrCreateCallRequest{
		Data: &CallRequest{
			CreatedByID: PtrTo("tommaso-id"),
		},
	}
	_, err := call.GetOrCreate(ctx, &callRequest)
	require.NoError(t, err, "Error creating call")

	return call
}

func getUser(t *testing.T, client *Stream, name *string, image *string, custom *map[string]any) (*FullUserResponse, error) {
	t.Helper()
	ctx := context.Background()
	userID := uuid.New().String()
	users := []UserRequest{
		{
			ID:     userID,
			Name:   name,
			Image:  image,
			Custom: custom,
		},
	}
	usersMap := make(map[string]UserRequest)
	for _, user := range users {
		usersMap[user.ID] = user
	}

	res, err := client.Common().UpdateUsers(ctx, &UpdateUsersRequest{Users: usersMap})
	if err != nil {
		return nil, err
	}
	user := res.Data.Users[userID]
	return &user, nil
}
