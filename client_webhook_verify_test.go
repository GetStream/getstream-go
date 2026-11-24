package getstream

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestClientVerifyWebhook_RejectsEmptySecretAndSignature(t *testing.T) {
	body := []byte(`{"type":"message.new"}`)

	mac := hmac.New(sha256.New, []byte(""))
	mac.Write(body)
	forgedEmptySig := []byte(hex.EncodeToString(mac.Sum(nil)))

	emptySecretClient := &Client{apiSecret: nil}
	if emptySecretClient.VerifyWebhook(body, forgedEmptySig) {
		t.Fatal("expected VerifyWebhook to reject empty secret")
	}

	validClient := &Client{apiSecret: []byte("real-webhook-secret")}
	if validClient.VerifyWebhook(body, nil) {
		t.Fatal("expected VerifyWebhook to reject empty signature")
	}

	validMac := hmac.New(sha256.New, []byte("real-webhook-secret"))
	validMac.Write(body)
	validSig := []byte(hex.EncodeToString(validMac.Sum(nil)))
	if !validClient.VerifyWebhook(body, validSig) {
		t.Fatal("expected VerifyWebhook to accept valid signature")
	}
}
