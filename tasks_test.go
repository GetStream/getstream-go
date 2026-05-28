package getstream_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v4"
)

// scriptedHTTPClient returns a queued sequence of HTTP responses. Each
// invocation pops one entry from the queue; if the queue is exhausted the
// final entry is reused (useful for "infinite polling" scenarios).
type scriptedHTTPClient struct {
	bodies []string
	calls  int32
}

func (c *scriptedHTTPClient) Do(req *http.Request) (*http.Response, error) {
	n := atomic.AddInt32(&c.calls, 1) - 1
	idx := int(n)
	if idx >= len(c.bodies) {
		idx = len(c.bodies) - 1
	}
	body := c.bodies[idx]
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     http.Header{},
		Request:    req,
	}, nil
}

func newTaskClient(t *testing.T, bodies []string) (*Stream, *scriptedHTTPClient) {
	t.Helper()
	mock := &scriptedHTTPClient{bodies: bodies}
	client, err := NewClient(
		"key", "secret",
		WithBaseUrl("https://api.example.invalid"),
		WithAuthToken("Bearer t"),
		WithHTTPClient(mock),
	)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return client, mock
}

func mustMarshal(t *testing.T, v any) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return string(b)
}

func TestWaitForTask_Completed(t *testing.T) {
	t.Parallel()
	completedBody := mustMarshal(t, map[string]any{
		"task_id": "task-1",
		"status":  "completed",
		"result":  map[string]any{"count": 42},
	})
	client, mock := newTaskClient(t, []string{completedBody})

	res, err := WaitForTask(
		context.Background(), client, "task-1",
		WithWaitForTaskPollInterval(10*time.Millisecond),
		WithWaitForTaskTimeout(2*time.Second),
	)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if res == nil || res.Data.Status != "completed" {
		t.Fatalf("expected completed status, got %+v", res)
	}
	if atomic.LoadInt32(&mock.calls) != 1 {
		t.Errorf("expected 1 call, got %d", mock.calls)
	}
}

func TestWaitForTask_PollsUntilTerminal(t *testing.T) {
	t.Parallel()
	runningBody := mustMarshal(t, map[string]any{
		"task_id": "task-1",
		"status":  "running",
	})
	completedBody := mustMarshal(t, map[string]any{
		"task_id": "task-1",
		"status":  "completed",
	})
	client, mock := newTaskClient(t, []string{runningBody, runningBody, completedBody})

	res, err := WaitForTask(
		context.Background(), client, "task-1",
		WithWaitForTaskPollInterval(5*time.Millisecond),
		WithWaitForTaskTimeout(2*time.Second),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Data.Status != "completed" {
		t.Fatalf("status = %q, want completed", res.Data.Status)
	}
	if mock.calls < 3 {
		t.Errorf("expected at least 3 polls, got %d", mock.calls)
	}
}

func TestWaitForTask_Failed_ReturnsErrTaskFailed(t *testing.T) {
	t.Parallel()
	failedBody := mustMarshal(t, map[string]any{
		"task_id": "task-1",
		"status":  "failed",
		"error": map[string]any{
			"type":        "ValidationError",
			"description": "channel does not exist",
			"stacktrace":  "...",
			"version":     "v1",
		},
	})
	client, _ := newTaskClient(t, []string{failedBody})

	_, err := WaitForTask(
		context.Background(), client, "task-1",
		WithWaitForTaskPollInterval(5*time.Millisecond),
		WithWaitForTaskTimeout(2*time.Second),
	)
	if err == nil {
		t.Fatalf("expected error on failed task, got nil")
	}
	if !errors.Is(err, ErrTaskFailed) {
		t.Errorf("errors.Is(err, ErrTaskFailed) = false; got %v", err)
	}

	var se *StreamError
	if !errors.As(err, &se) {
		t.Fatalf("errors.As(*StreamError) = false")
	}
	if se.Task == nil {
		t.Fatalf("StreamError.Task is nil")
	}
	if se.Task.TaskID != "task-1" || se.Task.ErrorType != "ValidationError" ||
		se.Task.Description != "channel does not exist" || se.Task.StackTrace != "..." ||
		se.Task.Version != "v1" {
		t.Errorf("Task fields mismatch: %+v", se.Task)
	}
}

func TestWaitForTask_Timeout_ReturnsTransportErrorTimeout(t *testing.T) {
	t.Parallel()
	runningBody := mustMarshal(t, map[string]any{
		"task_id": "task-1",
		"status":  "running",
	})
	client, _ := newTaskClient(t, []string{runningBody})

	_, err := WaitForTask(
		context.Background(), client, "task-1",
		WithWaitForTaskPollInterval(10*time.Millisecond),
		WithWaitForTaskTimeout(40*time.Millisecond),
	)
	if err == nil {
		t.Fatalf("expected timeout error, got nil")
	}
	if !errors.Is(err, ErrTransport) {
		t.Errorf("errors.Is(err, ErrTransport) = false; got %v", err)
	}

	var se *StreamError
	if !errors.As(err, &se) {
		t.Fatalf("errors.As(*StreamError) = false")
	}
	if se.ErrorType != ErrorTypeTimeout {
		t.Errorf("ErrorType = %q, want %q", se.ErrorType, ErrorTypeTimeout)
	}
	if !strings.Contains(se.Message, "task-1") {
		t.Errorf("Message should reference task ID; got %q", se.Message)
	}
}

func TestWaitForTask_ContextCancel(t *testing.T) {
	t.Parallel()
	runningBody := mustMarshal(t, map[string]any{
		"task_id": "task-1",
		"status":  "running",
	})
	client, _ := newTaskClient(t, []string{runningBody})

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(30 * time.Millisecond)
		cancel()
	}()

	_, err := WaitForTask(
		ctx, client, "task-1",
		WithWaitForTaskPollInterval(10*time.Millisecond),
		WithWaitForTaskTimeout(5*time.Second),
	)
	if err == nil {
		t.Fatalf("expected error from ctx cancel, got nil")
	}
	// ctx cancel propagates as either an ErrTransport timeout (if our wait
	// loop fired first) or as the SDK transport-error wrap of the cancelled
	// HTTP request. Either way, ErrTransport must match.
	if !errors.Is(err, ErrTransport) {
		t.Errorf("errors.Is(err, ErrTransport) = false; got %v", err)
	}
}

func TestWaitForTask_PropagatesAPIError(t *testing.T) {
	t.Parallel()
	// Mock returns a 200 with the running body. We override one entry to
	// simulate a 500 by using an httpClient that returns a non-2xx for the
	// first call.
	httpClient := &errorOnceClient{
		failNext:    true,
		successBody: mustMarshal(t, map[string]any{"task_id": "task-1", "status": "completed"}),
	}
	client, err := NewClient(
		"k", "s",
		WithBaseUrl("https://x.invalid"),
		WithAuthToken("Bearer t"),
		WithHTTPClient(httpClient),
	)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	_, err = WaitForTask(
		context.Background(), client, "task-1",
		WithWaitForTaskPollInterval(5*time.Millisecond),
		WithWaitForTaskTimeout(2*time.Second),
	)
	if err == nil {
		t.Fatalf("expected APIError to propagate, got nil")
	}
	if !errors.Is(err, ErrApiResponse) {
		t.Errorf("errors.Is(err, ErrApiResponse) = false; got %v", err)
	}
}

type errorOnceClient struct {
	failNext    bool
	successBody string
	mu          sync.Mutex
}

func (c *errorOnceClient) Do(req *http.Request) (*http.Response, error) {
	c.mu.Lock()
	fail := c.failNext
	c.failNext = false
	c.mu.Unlock()
	if fail {
		body := `{"code":1,"message":"server error","StatusCode":500}`
		return &http.Response{
			StatusCode: 500,
			Status:     "500 Internal Server Error",
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     http.Header{},
			Request:    req,
		}, nil
	}
	return &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Body:       io.NopCloser(strings.NewReader(c.successBody)),
		Header:     http.Header{},
		Request:    req,
	}, nil
}

func TestStreamError_TransportRoundTrip(t *testing.T) {
	t.Parallel()
	httpClient := &alwaysErrorClient{}
	client, err := NewClient(
		"k", "s",
		WithBaseUrl("https://x.invalid"),
		WithAuthToken("Bearer t"),
		WithHTTPClient(httpClient),
	)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	_, err = client.GetTask(context.Background(), "task-1", &GetTaskRequest{})
	if err == nil {
		t.Fatalf("expected transport error, got nil")
	}
	if !errors.Is(err, ErrTransport) {
		t.Fatalf("errors.Is(err, ErrTransport) = false; got %v", err)
	}
	var se *StreamError
	if !errors.As(err, &se) {
		t.Fatalf("errors.As(*StreamError) = false")
	}
	if se.ErrorType == "" {
		t.Errorf("ErrorType empty; expected a classification")
	}
}

type alwaysErrorClient struct{}

func (alwaysErrorClient) Do(_ *http.Request) (*http.Response, error) {
	return nil, errors.New("dial tcp 1.2.3.4:443: connection refused")
}
