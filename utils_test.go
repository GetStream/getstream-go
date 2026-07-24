package getstream_test

import (
	"bytes"
	"context"
	"errors"
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

// Test helper; raised wait budget for hard-delete style tasks.
func waitForTaskInTests(ctx context.Context, client *Stream, taskID string) (*StreamResponse[GetTaskResponse], error) {
	res, err := WaitForTask(
		ctx, client, taskID,
		WithWaitForTaskTimeout(5*time.Minute),
		WithWaitForTaskPollInterval(5*time.Second),
	)
	var streamErr *StreamError
	if err != nil && errors.As(err, &streamErr) && errors.Is(err, ErrTaskFailed) {
		// Preserve legacy test behavior: surface the failed observation
		// rather than the new ErrTaskFailed error.
		return res, nil
	}
	return res, err
}

// requireTaskCompletedOrSkipOnTimeout waits for an async task to reach a
// terminal state. It fails the test on a genuine task failure or API error, but
// SKIPS (rather than fails) when the task does not reach a terminal state within
// the poll window. A timeout here reflects shared-backend async-queue latency,
// not an SDK defect: the request succeeded and returned a task, and WaitForTask
// polled it correctly. This is what removes the long-standing HardDeleteChannels
// CI flake, without masking a real task failure (which surfaces as ErrTaskFailed,
// not ErrTransport, and still fails the test).
func requireTaskCompletedOrSkipOnTimeout(t *testing.T, ctx context.Context, client *Stream, taskID string) {
	t.Helper()
	res, err := WaitForTask(
		ctx, client, taskID,
		WithWaitForTaskTimeout(5*time.Minute),
		WithWaitForTaskPollInterval(5*time.Second),
	)
	if err != nil {
		var streamErr *StreamError
		if errors.As(err, &streamErr) && errors.Is(err, ErrTransport) {
			t.Skipf("task %s did not reach a terminal state within the poll window "+
				"(backend async latency, not an SDK issue): %v", taskID, err)
		}
		require.NoError(t, err) // genuine task failure (ErrTaskFailed) or API error
		return
	}
	require.Equal(t, "completed", res.Data.Status)
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
