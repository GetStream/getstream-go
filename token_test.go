package getstream

import (
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

	for _, key := range []string{"user_id", "server"} {
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
