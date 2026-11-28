package getstream

import (
	"math"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const expirationTestSecret = "expiration-test-secret"

func mintExpirationTestToken(t *testing.T, opts ...TokenOption) (jwt.MapClaims, error) {
	t.Helper()
	client, err := NewClient("key", expirationTestSecret)
	require.NoError(t, err)

	token, err := client.CreateToken("alice", opts...)
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(expirationTestSecret), nil
	})
	require.NoError(t, err)
	return claims, nil
}

func lifetimeOf(t *testing.T, claims jwt.MapClaims) int64 {
	t.Helper()
	iat, err := claims.GetIssuedAt()
	require.NoError(t, err)
	exp, err := claims.GetExpirationTime()
	require.NoError(t, err)
	require.NotNil(t, exp, "token must carry an exp claim")
	return exp.Unix() - iat.Unix()
}

func TestCreateTokenRejectsNonPositiveExpiration(t *testing.T) {
	for _, d := range []time.Duration{0, -time.Nanosecond, -time.Hour, math.MinInt64} {
		t.Run(d.String(), func(t *testing.T) {
			_, err := mintExpirationTestToken(t, WithExpiration(d))
			require.Error(t, err, "an explicit non-positive expiration must not mint a non-expiring token")
			assert.Contains(t, err.Error(), "expiration")
		})
	}
}

func TestCreateTokenRoundsSubSecondExpirationUpToOneSecond(t *testing.T) {
	for _, d := range []time.Duration{time.Nanosecond, time.Millisecond, 999 * time.Millisecond} {
		t.Run(d.String(), func(t *testing.T) {
			claims, err := mintExpirationTestToken(t, WithExpiration(d))
			require.NoError(t, err)
			assert.EqualValues(t, 1, lifetimeOf(t, claims))
		})
	}
}

func TestCreateTokenTruncatesWholeSecondExpirations(t *testing.T) {
	cases := map[time.Duration]int64{
		time.Second:             1,
		1999 * time.Millisecond: 1,
		time.Hour:               3600,
	}
	for d, want := range cases {
		t.Run(d.String(), func(t *testing.T) {
			claims, err := mintExpirationTestToken(t, WithExpiration(d))
			require.NoError(t, err)
			assert.EqualValues(t, want, lifetimeOf(t, claims))
		})
	}
}

func TestCreateTokenAcceptsLargestExpiration(t *testing.T) {
	claims, err := mintExpirationTestToken(t, WithExpiration(time.Duration(math.MaxInt64)))
	require.NoError(t, err)
	assert.EqualValues(t, int64(time.Duration(math.MaxInt64)/time.Second), lifetimeOf(t, claims))
}

func TestCreateTokenWithoutExpirationHasNoExp(t *testing.T) {
	claims, err := mintExpirationTestToken(t)
	require.NoError(t, err)
	assert.NotContains(t, claims, "exp")
}
