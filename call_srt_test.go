package getstream

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

const (
	srtTestAPIKey = "srt-test-key"
	srtTestSecret = "srt-test-secret-that-is-at-least-32-bytes"
)

var srtTestTokenPattern = regexp.MustCompile(`streamid=[^/]+/(.+)$`)

func newSRTTestCall(t *testing.T, callType, callID string) *Call {
	t.Helper()
	client, err := NewClient(srtTestAPIKey, srtTestSecret)
	require.NoError(t, err)
	call := client.Video().Call(callType, callID)
	call.data = &CallResponse{Ingress: CallIngressResponse{Srt: SRTIngress{
		Address: fmt.Sprintf("srt://ingress.stream-io-video.com:10080?mode=caller&passphrase={passphrase}&streamid=%s.%s.%s/{token}",
			srtTestAPIKey, callType, callID),
	}}}
	return call
}

func parseSRTToken(t *testing.T, address string) (string, jwt.MapClaims) {
	t.Helper()
	m := srtTestTokenPattern.FindStringSubmatch(address)
	require.NotNil(t, m)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(m[1], claims, func(*jwt.Token) (interface{}, error) {
		return []byte(srtTestSecret), nil
	})
	require.NoError(t, err)
	return m[1], claims
}

func TestCreateSRTCredentials(t *testing.T) {
	t.Run("requires an initialized call", func(t *testing.T) {
		client, err := NewClient(srtTestAPIKey, srtTestSecret)
		require.NoError(t, err)

		_, err = client.Video().Call("livestream", "my-call").CreateSRTCredentials("john")
		require.ErrorContains(t, err, "not initialized")
	})

	t.Run("call-scoped token that expires after 24 hours by default", func(t *testing.T) {
		creds, err := newSRTTestCall(t, "livestream", "my-call").CreateSRTCredentials("john")
		require.NoError(t, err)

		token, claims := parseSRTToken(t, creds.Address)
		require.Equal(t, "john", claims["user_id"])
		require.Equal(t, "livestream:my-call", claims["scoped_call_cid"])
		require.NotContains(t, claims, "call_cids")
		require.EqualValues(t, 24*60*60, claims["exp"].(float64)-claims["iat"].(float64))
		require.Equal(t, int64(claims["exp"].(float64)), creds.ExpiresAt.Unix())
		require.Contains(t, creds.Address, "passphrase="+strings.Split(token, ".")[2]+"&")
	})

	t.Run("custom expiration", func(t *testing.T) {
		creds, err := newSRTTestCall(t, "livestream", "my-call").
			CreateSRTCredentials("john", WithSRTTokenExpiration(3*time.Hour))
		require.NoError(t, err)

		_, claims := parseSRTToken(t, creds.Address)
		require.EqualValues(t, 3*60*60, claims["exp"].(float64)-claims["iat"].(float64))
	})

	t.Run("rejects an expiration under one second", func(t *testing.T) {
		for _, d := range []time.Duration{0, -time.Hour, 500 * time.Millisecond} {
			_, err := newSRTTestCall(t, "livestream", "my-call").
				CreateSRTCredentials("john", WithSRTTokenExpiration(d))
			require.ErrorContains(t, err, "at least one second", d)
		}
	})

	t.Run("rejects a stream id over the SRT limit", func(t *testing.T) {
		_, err := newSRTTestCall(t, "livestream", strings.Repeat("c", 200)).
			CreateSRTCredentials(strings.Repeat("u", 100))
		require.ErrorContains(t, err, "512-byte SRT limit")
	})

	t.Run("fmt does not print the address", func(t *testing.T) {
		creds, err := newSRTTestCall(t, "livestream", "my-call").CreateSRTCredentials("john")
		require.NoError(t, err)
		token, _ := parseSRTToken(t, creds.Address)

		for _, verb := range []string{"%v", "%+v", "%#v", "%s"} {
			for _, v := range []interface{}{creds, *creds} {
				printed := fmt.Sprintf(verb, v)
				require.NotContains(t, printed, token, verb)
				require.Contains(t, printed, "[redacted]", verb)
			}
		}
	})
}
