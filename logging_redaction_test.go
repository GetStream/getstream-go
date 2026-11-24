package getstream

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedactJSONBodyRecursive(t *testing.T) {
	raw := []byte(`{
		"access_token": "tok-123",
		"keep": "visible",
		"storage": {
			"s3_secret": "aws-secret-value",
			"firebase_credentials": "{\"private_key\":\"pem\"}",
			"bucket": "my-bucket"
		},
		"items": [{"apn_auth_key": "p8-key", "name": "ios"}]
	}`)

	out := redactJSONBody(raw)
	require.NotContains(t, out, "tok-123")
	require.NotContains(t, out, "aws-secret-value")
	require.NotContains(t, out, "private_key")
	require.NotContains(t, out, "p8-key")

	var parsed map[string]any
	require.NoError(t, json.Unmarshal([]byte(out), &parsed))
	require.Equal(t, "<redacted>", parsed["access_token"])
	require.Equal(t, "visible", parsed["keep"])

	storage := parsed["storage"].(map[string]any)
	require.Equal(t, "<redacted>", storage["s3_secret"])
	require.Equal(t, "<redacted>", storage["firebase_credentials"])
	require.Equal(t, "my-bucket", storage["bucket"])

	items := parsed["items"].([]any)
	first := items[0].(map[string]any)
	require.Equal(t, "<redacted>", first["apn_auth_key"])
	require.Equal(t, "ios", first["name"])
}
