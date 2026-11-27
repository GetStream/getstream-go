package getstream

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const tokenTestSecret = "token-test-secret"

func newTokenTestClient(t *testing.T) *Stream {
	t.Helper()
	client, err := NewClient("key", tokenTestSecret)
	require.NoError(t, err)
	return client
}

func decodeTestToken(t *testing.T, token string) jwt.MapClaims {
	t.Helper()
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(tokenTestSecret), nil
	})
	require.NoError(t, err)
	return claims
}

func TestCreateTokenRejectsReservedCustomClaims(t *testing.T) {
	client := newTokenTestClient(t)

	for _, key := range []string{"user_id", "server", "egress"} {
		t.Run(key, func(t *testing.T) {
			for _, value := range []interface{}{"", "other-user", true, false} {
				_, err := client.CreateToken("alice", WithClaims(Claims{
					CustomClaims: map[string]interface{}{key: value},
				}))
				require.Error(t, err, "value %v", value)
				assert.Contains(t, err.Error(), key)
			}
		})
	}
}

func TestCreateTokenRejectsEmptyUserID(t *testing.T) {
	client := newTokenTestClient(t)

	_, err := client.CreateToken("")
	require.Error(t, err)
}

func TestCreateTokenExpirationOptionWinsOverCustomExp(t *testing.T) {
	client := newTokenTestClient(t)
	farFuture := time.Now().Add(100 * 365 * 24 * time.Hour).Unix()

	before := time.Now().Unix()
	token, err := client.CreateToken("alice",
		WithExpiration(time.Hour),
		WithClaims(Claims{CustomClaims: map[string]interface{}{"exp": farFuture}}),
	)
	require.NoError(t, err)

	claims := decodeTestToken(t, token)
	exp, err := claims.GetExpirationTime()
	require.NoError(t, err)
	assert.InDelta(t, before+3600, exp.Unix(), 5)
}

func TestCreateTokenCustomExpHonoredWithoutExpirationOption(t *testing.T) {
	client := newTokenTestClient(t)
	exp := time.Now().Add(2 * time.Hour).Unix()

	token, err := client.CreateToken("alice",
		WithClaims(Claims{CustomClaims: map[string]interface{}{"exp": exp}}),
	)
	require.NoError(t, err)

	claims := decodeTestToken(t, token)
	assert.EqualValues(t, exp, claims["exp"])
}

func TestCreateTokenTypedClaimsWinOverCustomClaims(t *testing.T) {
	client := newTokenTestClient(t)

	token, err := client.CreateToken("alice", WithClaims(Claims{
		Role:        "user",
		CallCIDs:    []string{"default:call"},
		ChannelCIDs: []string{"messaging:general"},
		CustomClaims: map[string]interface{}{
			"role":         "admin",
			"call_cids":    []string{"default:other"},
			"channel_cids": []string{"messaging:other"},
			"tenant":       "acme",
		},
	}))
	require.NoError(t, err)

	claims := decodeTestToken(t, token)
	assert.Equal(t, "alice", claims["user_id"])
	assert.Equal(t, "user", claims["role"])
	assert.Equal(t, []interface{}{"default:call"}, claims["call_cids"])
	assert.Equal(t, []interface{}{"messaging:general"}, claims["channel_cids"])
	assert.Equal(t, "acme", claims["tenant"])
	assert.NotContains(t, claims, "server")
}

func TestCreateTokenScopedClaimsAllowedAsCustomClaims(t *testing.T) {
	client := newTokenTestClient(t)

	token, err := client.CreateToken("alice", WithClaims(Claims{
		CustomClaims: map[string]interface{}{
			"role":      "admin",
			"call_cids": []string{"default:call"},
		},
	}))
	require.NoError(t, err)

	claims := decodeTestToken(t, token)
	assert.Equal(t, "alice", claims["user_id"])
	assert.Equal(t, "admin", claims["role"])
	assert.Equal(t, []interface{}{"default:call"}, claims["call_cids"])
}

// backendTokenClaims has the shape of the claims struct the backend decodes a
// token into with encoding/json.
type backendTokenClaims struct {
	RawUserID     interface{} `json:"user_id"`
	ChannelCIDs   []string    `json:"channel_cids"`
	CallCIDs      []string    `json:"call_cids"`
	Role          *string     `json:"role"`
	Egress        bool        `json:"egress,omitempty"`
	CheckWrites   bool        `json:"stream_check_writes,omitempty"`
	ScopedCallCID string      `json:"scoped_call_cid,omitempty"`
	jwt.RegisteredClaims
}

func decodeBackendClaims(t *testing.T, token string) backendTokenClaims {
	t.Helper()
	payload, err := jwt.NewParser().DecodeSegment(strings.Split(token, ".")[1])
	require.NoError(t, err)
	var claims backendTokenClaims
	require.NoError(t, json.Unmarshal(payload, &claims))
	return claims
}

// Other spellings of claim names: case variants, the long s (U+017F, folds to
// s) and the Kelvin sign (U+212A, folds to k).
var foldedClaimNames = []string{
	"USER_ID", "User_Id", "u\u017Fer_id", "\u017Ferver", "SERVER",
	"Egress", "egre\u017F\u017F", "IAT", "Exp", "Role", "Call_CIDs",
	"stream_chec\u212A_writes", "\u017Fcoped_call_cid",
}

func TestBackendDecoderMatchesFoldedClaimNames(t *testing.T) {
	// The reason the SDK rejects these names: the backend's decoder reads each
	// as the canonical claim, and the last key wins.
	payload := `{"user_id":"alice","u\u017Fer_id":null,"Egre\u017F\u017F":true,"ROLE":"admin","stream_chec\u212A_writes":true}`
	var claims backendTokenClaims
	require.NoError(t, json.Unmarshal([]byte(payload), &claims))
	assert.Nil(t, claims.RawUserID)
	assert.True(t, claims.Egress)
	require.NotNil(t, claims.Role)
	assert.Equal(t, "admin", *claims.Role)
	assert.True(t, claims.CheckWrites)
}

func TestCreateTokenRejectsFoldedClaimNames(t *testing.T) {
	client := newTokenTestClient(t)

	for _, key := range foldedClaimNames {
		t.Run(key, func(t *testing.T) {
			for _, value := range []interface{}{nil, "", true} {
				_, err := client.CreateToken("alice", WithClaims(Claims{
					CustomClaims: map[string]interface{}{key: value},
				}))
				require.Error(t, err, "value %v", value)
				assert.Contains(t, err.Error(), "read by the backend")
			}
		})
	}
}

func TestCreateTokenValidatesCustomIat(t *testing.T) {
	client := newTokenTestClient(t)
	now := time.Now().Unix()

	for _, iat := range []interface{}{0, -1, 1.5, now + 3600, "123", nil, true} {
		_, err := client.CreateToken("alice", WithClaims(Claims{
			CustomClaims: map[string]interface{}{"iat": iat},
		}))
		require.Error(t, err, "iat %v", iat)
		assert.Contains(t, err.Error(), "iat")
	}

	past := now - 60
	for _, iat := range []interface{}{past, int(past), float64(past), json.Number("1700000000")} {
		_, err := client.CreateToken("alice", WithClaims(Claims{
			CustomClaims: map[string]interface{}{"iat": iat},
		}))
		require.NoError(t, err, "iat %v", iat)
	}
}

func TestCreateTokenIdentitySurvivesBackendDecoding(t *testing.T) {
	client := newTokenTestClient(t)

	token, err := client.CreateToken("alice", WithClaims(Claims{
		CallCIDs: []string{"default:call"},
		CustomClaims: map[string]interface{}{
			"scoped_call_cid": "default:call",
			"custom":          "value",
		},
	}))
	require.NoError(t, err)

	claims := decodeBackendClaims(t, token)
	assert.Equal(t, "alice", claims.RawUserID)
	assert.False(t, claims.Egress)
	assert.Equal(t, []string{"default:call"}, claims.CallCIDs)
	assert.Equal(t, "default:call", claims.ScopedCallCID)
	require.NotNil(t, claims.IssuedAt)
	assert.InDelta(t, time.Now().Unix(), claims.IssuedAt.Unix(), 5)
}
