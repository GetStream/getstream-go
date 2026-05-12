package getstream_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	. "github.com/GetStream/getstream-go/v4"
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

		event, err := VerifyAndParseWebhook(validBody, sig, secret)
		require.NoError(t, err)
		require.NotNil(t, event)
		assert.Equal(t, "message.new", event.GetEventType())
	})

	t.Run("MissingSignature", func(t *testing.T) {
		_, err := VerifyAndParseWebhook(validBody, "", secret)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "signature")
	})

	t.Run("InvalidSignature", func(t *testing.T) {
		_, err := VerifyAndParseWebhook(validBody, "badsignature", secret)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "signature")
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		body := []byte(`not json at all`)
		sig := computeSignature(body, secret)

		_, err := VerifyAndParseWebhook(body, sig, secret)
		require.Error(t, err)
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
