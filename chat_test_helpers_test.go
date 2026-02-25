package getstream_test

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// rateLimitClient wraps an HttpClient and automatically retries on 429 responses
// with exponential backoff and jitter to avoid thundering herd.
type rateLimitClient struct {
	inner HttpClient
}

func (c *rateLimitClient) Do(r *http.Request) (*http.Response, error) {
	// Buffer the body so we can replay it on retry
	var bodyBytes []byte
	if r.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	const maxRetries = 8
	for i := 0; i < maxRetries; i++ {
		if i > 0 && bodyBytes != nil {
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
		resp, err := c.inner.Do(r)
		if err != nil {
			return resp, err
		}
		if resp.StatusCode != http.StatusTooManyRequests {
			return resp, nil
		}
		// Drain and close the 429 response body before retrying
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()

		// Use the reset header if available, otherwise exponential backoff + jitter
		backoff := rateLimitBackoff(resp.Header, i)
		time.Sleep(backoff)
	}
	// Exhausted retries — replay one last time and return whatever we get
	if bodyBytes != nil {
		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}
	return c.inner.Do(r)
}

// rateLimitBackoff calculates a backoff duration from the rate limit reset header.
// Falls back to exponential backoff with jitter if the header is missing.
func rateLimitBackoff(headers http.Header, attempt int) time.Duration {
	if resetStr := headers.Get("X-Ratelimit-Reset"); resetStr != "" {
		if resetUnix, err := strconv.ParseInt(resetStr, 10, 64); err == nil && resetUnix > 0 {
			wait := time.Until(time.Unix(resetUnix, 0))
			if wait > 0 && wait < 90*time.Second {
				// Add small jitter to desynchronize concurrent retries
				jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
				return wait + jitter
			}
		}
	}
	// Exponential backoff: 2s, 4s, 8s, 16s, capped at 30s — plus jitter
	base := time.Duration(1<<uint(attempt+1)) * time.Second
	if base > 30*time.Second {
		base = 30 * time.Second
	}
	jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
	return base + jitter
}

// newRateLimitClient wraps an http.Client with automatic 429 retry.
// The inner client timeout applies per-attempt, not for the whole retry chain.
func newRateLimitClient() *rateLimitClient {
	return &rateLimitClient{inner: &http.Client{Timeout: 60 * time.Second}}
}

// requireNoErrorOrSkipRateLimit asserts no error, but skips the test if the
// error is a rate limit ("Too many requests"). Use this for API calls that are
// heavily contended in parallel test runs.
func requireNoErrorOrSkipRateLimit(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		if strings.Contains(err.Error(), "Too many requests") {
			t.Skipf("Skipping: rate limited after retries (%s)", t.Name())
		}
	}
	require.NoError(t, err)
}

// deleteUsersWithRetry calls DeleteUsers with retry logic to handle rate limiting.
// Used in t.Cleanup to avoid "Too many requests" failures.
// The rateLimitClient handles 429 retries at the HTTP level, so this just
// does a single attempt — cleanup failures are acceptable.
func deleteUsersWithRetry(client *Stream, userIDs []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	_, _ = client.DeleteUsers(ctx, &DeleteUsersRequest{
		UserIds:       userIDs,
		User:          PtrTo("hard"),
		Messages:      PtrTo("hard"),
		Conversations: PtrTo("hard"),
	})
}

// skipIfShort skips integration tests when running with -short flag.
func skipIfShort(t *testing.T) {
	t.Helper()
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
}

// createTestUsers creates n test users and returns their IDs.
// Registers t.Cleanup to hard-delete users (matching stream-chat-go's randomUser pattern).
func createTestUsers(t *testing.T, client *Stream, n int) []string {
	t.Helper()
	ctx := context.Background()
	users := make(map[string]UserRequest, n)
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		id := "test-user-" + uuid.New().String()
		ids[i] = id
		users[id] = UserRequest{
			ID:   id,
			Name: PtrTo("Test User " + id[:8]),
			Role: PtrTo("user"),
		}
	}
	_, err := client.UpdateUsers(ctx, &UpdateUsersRequest{Users: users})
	require.NoError(t, err, "Failed to create test users")

	t.Cleanup(func() {
		deleteUsersWithRetry(client, ids)
	})

	return ids
}

// createTestChannel creates a messaging channel and registers cleanup to hard-delete it.
func createTestChannel(t *testing.T, client *Stream, creatorID string) (*Channels, string) {
	t.Helper()
	ctx := context.Background()
	channelID := "test-ch-" + randomString(12)
	ch := client.Chat().Channel("messaging", channelID)

	_, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{
		Data: &ChannelInput{
			CreatedByID: PtrTo(creatorID),
		},
	})
	require.NoError(t, err, "Failed to create test channel")

	t.Cleanup(func() {
		_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{
			HardDelete: PtrTo(true),
		})
	})

	return ch, channelID
}

// createTestChannelWithMembers creates a messaging channel with members and registers cleanup to hard-delete it.
func createTestChannelWithMembers(t *testing.T, client *Stream, creatorID string, memberIDs []string) (*Channels, string) {
	t.Helper()
	ctx := context.Background()
	channelID := "test-ch-" + randomString(12)
	ch := client.Chat().Channel("messaging", channelID)

	members := make([]ChannelMemberRequest, len(memberIDs))
	for i, id := range memberIDs {
		members[i] = ChannelMemberRequest{UserID: id}
	}

	_, err := ch.GetOrCreate(ctx, &GetOrCreateChannelRequest{
		Data: &ChannelInput{
			CreatedByID: PtrTo(creatorID),
			Members:     members,
		},
	})
	require.NoError(t, err, "Failed to create test channel with members")

	t.Cleanup(func() {
		_, _ = ch.Delete(context.Background(), &DeleteChannelRequest{
			HardDelete: PtrTo(true),
		})
	})

	return ch, channelID
}

// sendTestMessage sends a message to a channel and returns the message ID.
func sendTestMessage(t *testing.T, ch *Channels, userID string, text string) string {
	t.Helper()
	ctx := context.Background()
	resp, err := ch.SendMessage(ctx, &SendMessageRequest{
		Message: MessageRequest{
			Text:   PtrTo(text),
			UserID: PtrTo(userID),
		},
	})
	require.NoError(t, err, "Failed to send test message")
	require.NotEmpty(t, resp.Data.Message.ID, "Message ID should not be empty")
	return resp.Data.Message.ID
}
