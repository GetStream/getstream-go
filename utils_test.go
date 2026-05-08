package getstream_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type StubHTTPClient struct{}

func (c *StubHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Status: "200 OK",
		Body:   io.NopCloser(bytes.NewBufferString("{}")),
	}, nil
}

// WaitForTask polls the task until it reaches a terminal status
// ("completed" or "failed") and returns the final observation.
//
// Polling cadence ramps from 1s to a 5s cap. The 60-attempt budget
// (~5 minutes) is sized to outlast async-task retry cycles on the
// backend (e.g. rate-limited bulk hard-deletes that retry every
// 10–15s for several minutes before clearing).
//
// If the budget is exhausted before a terminal status is observed,
// the last observation is returned alongside a timeout error so
// the caller can log the in-flight status for debugging.
func WaitForTask(ctx context.Context, client *Stream, taskID string) (*StreamResponse[GetTaskResponse], error) {
	const maxAttempts = 60
	var lastResult *StreamResponse[GetTaskResponse]
	for i := 0; i < maxAttempts; i++ {
		taskResult, err := client.GetTask(context.Background(), taskID, &GetTaskRequest{})
		if err != nil {
			return nil, fmt.Errorf("failed to get task result: %w", err)
		}
		lastResult = taskResult
		switch taskResult.Data.Status {
		case "completed", "failed":
			return taskResult, nil
		}

		interval := time.Duration(i+1) * time.Second
		if interval > 5*time.Second {
			interval = 5 * time.Second
		}

		select {
		case <-ctx.Done():
			return lastResult, fmt.Errorf("context expired waiting for task %s: %w", taskID, ctx.Err())
		case <-time.After(interval):
		}
	}
	return lastResult, fmt.Errorf("task %s did not reach terminal status after %d attempts", taskID, maxAttempts)
}

// ResourceManager manages resource cleanup for tests.
type ResourceManager struct {
	t *testing.T
}

// NewResourceManager initializes a new ResourceManager.
func NewResourceManager(t *testing.T) *ResourceManager {
	return &ResourceManager{t: t}
}

// RegisterCleanup registers a cleanup function to be called when the test finishes.
func (rm *ResourceManager) RegisterCleanup(cleanup func()) {
	rm.t.Cleanup(cleanup)
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

func getUser(t *testing.T, client *Stream, name *string, image *string, custom map[string]any) (*FullUserResponse, error) {
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

	res, err := client.UpdateUsers(ctx, &UpdateUsersRequest{Users: usersMap})
	if err != nil {
		return nil, err
	}
	user := res.Data.Users[userID]
	return &user, nil
}
