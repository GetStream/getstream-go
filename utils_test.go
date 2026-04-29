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

func WaitForTask(ctx context.Context, client *Stream, taskID string) (*StreamResponse[GetTaskResponse], error) {
	// Poll with progressive intervals: start at 1s, increase by 1s each
	// attempt up to 5s, for a total ceiling of ~190s. This handles slow
	// task completion under heavy parallel load while still finishing
	// quickly when the server is responsive.
	//
	// "failed" is treated as non-terminal until it persists for at least
	// failedConfirmDuration. The chat backend writes Status="failed" before
	// asynq retries rate-limited or transient-internal failures; retries are
	// scheduled 10-15s out (rateLimitAwareRetryDelay), and the next attempt
	// overwrites the result with "running" then "completed". Any non-failed
	// observation resets the deadline so chains of retries don't time out.
	const maxAttempts = 40
	const failedConfirmDuration = 30 * time.Second
	var lastResult *StreamResponse[GetTaskResponse]
	var firstFailedAt time.Time
	for i := 0; i < maxAttempts; i++ {
		taskResult, err := client.GetTask(context.Background(), taskID, &GetTaskRequest{})
		if err != nil {
			return nil, fmt.Errorf("failed to get task result: %w", err)
		}
		lastResult = taskResult
		switch taskResult.Data.Status {
		case "completed":
			return taskResult, nil
		case "failed":
			if firstFailedAt.IsZero() {
				firstFailedAt = time.Now()
			}
			if time.Since(firstFailedAt) >= failedConfirmDuration {
				return taskResult, nil
			}
		default:
			firstFailedAt = time.Time{}
		}

		interval := time.Duration(i+1) * time.Second
		if interval > 5*time.Second {
			interval = 5 * time.Second
		}

		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context expired waiting for task %s: %w", taskID, ctx.Err())
		case <-time.After(interval):
		}
	}
	if lastResult != nil {
		return lastResult, fmt.Errorf("task %s did not complete after %d attempts (last status %q)", taskID, maxAttempts, lastResult.Data.Status)
	}
	return nil, fmt.Errorf("task %s did not complete after %d attempts", taskID, maxAttempts)
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
