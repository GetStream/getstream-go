package getstream_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func computeSignature(body []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	return hex.EncodeToString(h.Sum(nil))
}

// Timestamps in this SDK use nanosecond Unix epochs (not ISO strings).
// 1704067200000000000 = 2024-01-01T00:00:00Z in nanoseconds.
const testTimestampNs = "1704067200000000000"

func TestVerifyAndParseWebhook(t *testing.T) {
	secret := "test-webhook-secret-123"

	validBody := []byte(`{"type":"message.new","message_id":"msg123","created_at":` + testTimestampNs + `,"watcher_count":0,"custom":{},"message":{"id":"msg123","text":"hello","type":"regular","created_at":` + testTimestampNs + `,"updated_at":` + testTimestampNs + `,"cid":"messaging:test","html":""}}`)

	t.Run("ValidRequest", func(t *testing.T) {
		sig := computeSignature(validBody, secret)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(validBody))
		req.Header.Set("X-Signature", sig)

		event, err := VerifyAndParseWebhook(req, secret)
		require.NoError(t, err)
		require.NotNil(t, event)
		assert.Equal(t, "message.new", event.GetEventType())
	})

	t.Run("MissingSignatureHeader", func(t *testing.T) {
		body := []byte(`{"type":"message.new"}`)
		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))

		_, err := VerifyAndParseWebhook(req, secret)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "X-Signature")
	})

	t.Run("InvalidSignature", func(t *testing.T) {
		body := []byte(`{"type":"message.new"}`)
		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
		req.Header.Set("X-Signature", "badsignature")

		_, err := VerifyAndParseWebhook(req, secret)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid signature")
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		body := []byte(`not json at all`)
		sig := computeSignature(body, secret)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
		req.Header.Set("X-Signature", sig)

		_, err := VerifyAndParseWebhook(req, secret)
		require.Error(t, err)
	})

	t.Run("BodyRestoredAfterParsing", func(t *testing.T) {
		sig := computeSignature(validBody, secret)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(validBody))
		req.Header.Set("X-Signature", sig)

		_, err := VerifyAndParseWebhook(req, secret)
		require.NoError(t, err)

		// Body should still be readable
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(req.Body)
		require.NoError(t, err)
		assert.NotEmpty(t, buf.Bytes(), "Body should be restored after VerifyAndParseWebhook")
	})
}

func TestWebhookMiddleware(t *testing.T) {
	secret := "middleware-test-secret"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		event := r.Context().Value(WebhookEventKey)
		if event == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	middleware := WebhookMiddleware(secret)
	wrapped := middleware(handler)

	t.Run("RejectsGET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/webhook", nil)
		rr := httptest.NewRecorder()
		wrapped.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	})

	t.Run("RejectsBadSignature", func(t *testing.T) {
		body := []byte(`{"type":"message.new"}`)
		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
		req.Header.Set("X-Signature", "invalidsig")
		rr := httptest.NewRecorder()
		wrapped.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("PassesValidRequest", func(t *testing.T) {
		body := []byte(`{"type":"message.new","message_id":"msg1","created_at":` + testTimestampNs + `,"watcher_count":0,"custom":{},"message":{"id":"msg1","text":"hi","type":"regular","created_at":` + testTimestampNs + `,"updated_at":` + testTimestampNs + `,"cid":"messaging:test","html":""}}`)
		sig := computeSignature(body, secret)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
		req.Header.Set("X-Signature", sig)
		rr := httptest.NewRecorder()
		wrapped.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestParseWebhookEventWithPayload(t *testing.T) {
	t.Run("MessageNewEvent", func(t *testing.T) {
		payload := []byte(`{
			"type": "message.new",
			"message_id": "msg-abc-123",
			"created_at": ` + testTimestampNs + `,
			"watcher_count": 5,
			"custom": {},
			"message": {
				"id": "msg-abc-123",
				"text": "Hello world!",
				"type": "regular",
				"created_at": ` + testTimestampNs + `,
				"updated_at": ` + testTimestampNs + `,
				"cid": "messaging:general",
				"html": ""
			}
		}`)
		event, err := ParseWebhookEvent(payload)
		require.NoError(t, err)
		assert.Equal(t, "message.new", event.GetEventType())

		msgEvent, ok := event.(*MessageNewEvent)
		require.True(t, ok, "Should be *MessageNewEvent")
		assert.Equal(t, "msg-abc-123", msgEvent.MessageID)
		assert.Equal(t, "msg-abc-123", msgEvent.Message.ID)
		assert.Equal(t, "Hello world!", msgEvent.Message.Text)
		assert.Equal(t, 5, msgEvent.WatcherCount)
	})

	t.Run("ChannelCreatedEvent", func(t *testing.T) {
		payload := []byte(`{
			"type": "channel.created",
			"created_at": ` + testTimestampNs + `,
			"custom": {},
			"channel": {
				"cid": "messaging:new-channel",
				"id": "new-channel",
				"type": "messaging",
				"created_at": ` + testTimestampNs + `,
				"updated_at": ` + testTimestampNs + `,
				"frozen": false,
				"disabled": false,
				"member_count": 2,
				"config": {
					"automod": "disabled",
					"automod_behavior": "flag",
					"commands": [],
					"connect_events": true,
					"created_at": ` + testTimestampNs + `,
					"max_message_length": 5000,
					"message_retention": "infinite",
					"mutes": true,
					"name": "messaging",
					"reactions": true,
					"read_events": true,
					"replies": true,
					"search": true,
					"typing_events": true,
					"updated_at": ` + testTimestampNs + `,
					"uploads": true,
					"url_enrichment": true
				}
			}
		}`)
		event, err := ParseWebhookEvent(payload)
		require.NoError(t, err)
		assert.Equal(t, "channel.created", event.GetEventType())

		chEvent, ok := event.(*ChannelCreatedEvent)
		require.True(t, ok, "Should be *ChannelCreatedEvent")
		assert.Equal(t, "messaging:new-channel", chEvent.Channel.Cid)
		assert.Equal(t, "new-channel", chEvent.Channel.ID)
		assert.Equal(t, "messaging", chEvent.Channel.Type)
		require.NotNil(t, chEvent.Channel.MemberCount)
		assert.Equal(t, 2, *chEvent.Channel.MemberCount)
	})

	t.Run("UnknownEventType", func(t *testing.T) {
		payload := []byte(`{"type":"completely.unknown.event"}`)
		_, err := ParseWebhookEvent(payload)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unknown webhook event type")
	})
}

func TestWebhookEventKey(t *testing.T) {
	secret := "event-key-test-secret"

	var capturedEvent WebhookEvent

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(WebhookEventKey)
		if val != nil {
			capturedEvent = val.(WebhookEvent)
		}
		w.WriteHeader(http.StatusOK)
	})

	middleware := WebhookMiddleware(secret)
	wrapped := middleware(handler)

	body := []byte(`{"type":"channel.created","created_at":` + testTimestampNs + `,"custom":{},"channel":{"cid":"messaging:test","id":"test","type":"messaging","created_at":` + testTimestampNs + `,"updated_at":` + testTimestampNs + `,"frozen":false,"disabled":false,"member_count":0,"config":{"automod":"disabled","automod_behavior":"flag","commands":[],"connect_events":true,"created_at":` + testTimestampNs + `,"max_message_length":5000,"message_retention":"infinite","mutes":true,"name":"messaging","reactions":true,"read_events":true,"replies":true,"search":true,"typing_events":true,"updated_at":` + testTimestampNs + `,"uploads":true,"url_enrichment":true}}}`)
	sig := computeSignature(body, secret)

	req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
	req.Header.Set("X-Signature", sig)
	rr := httptest.NewRecorder()
	wrapped.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.NotNil(t, capturedEvent, "Event should be captured from context")
	assert.Equal(t, "channel.created", capturedEvent.GetEventType())

	chEvent, ok := capturedEvent.(*ChannelCreatedEvent)
	require.True(t, ok)
	assert.Equal(t, "messaging:test", chEvent.Channel.Cid)
}
