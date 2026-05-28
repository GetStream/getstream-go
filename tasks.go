package getstream

import (
	"context"
	"errors"
	"time"
)

// Default polling parameters for WaitForTask, per Server-Side SDK Error
// Handling Spec §8.
const (
	defaultTaskPollInterval = 1 * time.Second
	defaultTaskWaitTimeout  = 60 * time.Second
)

// WaitForTaskOption configures WaitForTask.
type WaitForTaskOption func(*waitForTaskConfig)

type waitForTaskConfig struct {
	pollInterval time.Duration
	timeout      time.Duration
}

// WithWaitForTaskPollInterval sets how often the task-status endpoint is
// polled. Default 1s. Values <= 0 are ignored.
func WithWaitForTaskPollInterval(d time.Duration) WaitForTaskOption {
	return func(c *waitForTaskConfig) {
		if d > 0 {
			c.pollInterval = d
		}
	}
}

// WithWaitForTaskTimeout sets the maximum wait before returning a timeout
// error. Default 60s. Values <= 0 disable the timeout (waits until ctx ends
// or terminal status is reached).
func WithWaitForTaskTimeout(d time.Duration) WaitForTaskOption {
	return func(c *waitForTaskConfig) {
		c.timeout = d
	}
}

// WaitForTask polls the task-status endpoint until the task reaches a
// terminal state, the configured timeout elapses, or ctx is cancelled.
//
//   - On status "completed": returns the final StreamResponse.
//   - On status "failed": returns a *StreamError with sentinel ErrTaskFailed
//     and the populated Task field.
//   - On timeout or ctx cancellation: returns a *StreamError with sentinel
//     ErrTransport and ErrorType "timeout"; the original cause (a
//     context.DeadlineExceeded or ctx.Err()) is reachable via errors.Unwrap.
//
// Behavior matches Server-Side SDK Error Handling Spec §8. Defaults:
// pollInterval = 1s, timeout = 60s; override with WaitForTaskOption.
func WaitForTask(ctx context.Context, client *Stream, taskID string, opts ...WaitForTaskOption) (*StreamResponse[GetTaskResponse], error) {
	cfg := &waitForTaskConfig{
		pollInterval: defaultTaskPollInterval,
		timeout:      defaultTaskWaitTimeout,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	deadline := time.Time{}
	if cfg.timeout > 0 {
		deadline = time.Now().Add(cfg.timeout)
	}

	var lastResult *StreamResponse[GetTaskResponse]
	for {
		taskResult, err := client.GetTask(ctx, taskID, &GetTaskRequest{})
		if err != nil {
			return lastResult, err
		}
		lastResult = taskResult

		switch taskResult.Data.Status {
		case "completed":
			return taskResult, nil
		case "failed":
			return lastResult, taskFailureError(taskID, taskResult.Data.Error)
		}

		// Honor deadline before sleeping.
		if !deadline.IsZero() && !time.Now().Before(deadline) {
			return lastResult, taskTimeoutError(taskID, context.DeadlineExceeded)
		}

		select {
		case <-ctx.Done():
			return lastResult, taskTimeoutError(taskID, ctx.Err())
		case <-time.After(cfg.pollInterval):
		}
	}
}

// taskFailureError constructs the *StreamError surfaced when a task reaches
// status "failed".
func taskFailureError(taskID string, errRes *ErrorResult) *StreamError {
	details := &TaskErrorDetails{TaskID: taskID}
	if errRes != nil {
		details.ErrorType = errRes.Type
		details.Description = errRes.Description
		if errRes.Stacktrace != nil {
			details.StackTrace = *errRes.Stacktrace
		}
		if errRes.Version != nil {
			details.Version = *errRes.Version
		}
	}

	msg := "stream task failed"
	if details.Description != "" {
		msg = "stream task failed: " + details.Description
	}

	return &StreamError{
		sentinel: ErrTaskFailed,
		Message:  msg,
		Task:     details,
		cause:    stackWrap(errors.New(msg), "task failed"),
	}
}

// taskTimeoutError constructs the *StreamError surfaced when WaitForTask
// exhausts its deadline or ctx is cancelled.
func taskTimeoutError(taskID string, cause error) *StreamError {
	if cause == nil {
		cause = context.DeadlineExceeded
	}
	return &StreamError{
		sentinel:  ErrTransport,
		ErrorType: ErrorTypeTimeout,
		Message:   "stream wait for task timed out: " + taskID,
		cause:     stackWrap(cause, "wait for task "+taskID),
	}
}
