package getstream

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "Special date",
			time:     time.Date(2018, 10, 5, 4, 20, 0, 0, time.UTC),
			expected: `"2018-10-05T04:20:00Z"`,
		},
		{
			name:     "RFC3339 time",
			time:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: `"2020-01-01T00:00:00Z"`,
		},
		{
			name:     "Future date",
			time:     time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: `"2030-01-01T00:00:00Z"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := Timestamp{&tc.time}
			data, err := json.Marshal(ts)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, string(data))
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		expected *Timestamp
	}{
		{
			name:     "empty",
			data:     []byte("null"),
			expected: &Timestamp{},
		},
		{
			name: "special date",
			data: []byte("1538713200000000000"),
			expected: &Timestamp{
				func() *time.Time {
					t := time.Date(2018, 10, 5, 4, 20, 0, 0, time.UTC)
					return &t
				}(),
			},
		},
		{
			name: "future date",
			data: []byte("2233023600000000000"),
			expected: &Timestamp{
				func() *time.Time {
					t := time.Date(2040, 10, 5, 4, 20, 0, 0, time.UTC)
					return &t
				}(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := &Timestamp{}
			err := json.Unmarshal(tc.data, ts)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, ts)
		})
	}
}
