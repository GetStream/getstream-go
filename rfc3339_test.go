package getstream

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		time     time.Time
		rfc3339  bool
		expected string
	}{
		{
			name:     "Zero time",
			time:     time.Time{},
			rfc3339:  false,
			expected: "null",
		},
		{
			name:     "RFC3339 time",
			time:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			rfc3339:  true,
			expected: `"2020-01-01T00:00:00Z"`,
		},
		// time in the future
		{
			name:     "RFC3339 time in the future",
			time:     time.Date(4102, 1, 1, 0, 0, 0, 0, time.UTC),
			rfc3339:  true,
			expected: `"4102-01-01T00:00:00Z"`,
		},
		{
			name:     "Unix nanoseconds",
			time:     time.Date(2020, 1, 1, 0, 0, 0, 123456789, time.UTC),
			rfc3339:  false,
			expected: "1577836800123456789",
		},
		{
			name:     "Negative Unix nanoseconds",
			time:     time.Date(1969, 12, 31, 23, 59, 59, 0, time.UTC),
			rfc3339:  false,
			expected: "-1000000000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := NewTimestamp(tc.time, tc.rfc3339)
			data, err := json.Marshal(ts)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if string(data) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, string(data))
			}

			// Test round-trip
			var unmarshalled Timestamp
			err = json.Unmarshal(data, &unmarshalled)
			if err != nil {
				t.Fatalf("Expected no error during unmarshal, got %v", err)
			}
			if !unmarshalled.Equal(ts.Time) {
				t.Errorf("Round-trip failed. Expected %v, got %v", ts.Time, unmarshalled.Time)
			}
		})
	}
}

func TestUnmarshalJSON_Null(t *testing.T) {
	data := "null"
	var ts Timestamp
	err := json.Unmarshal([]byte(data), &ts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !ts.Time.IsZero() {
		t.Errorf("Expected zero time, got %v", ts.Time)
	}
}

func TestUnmarshalJSON_NegativeUnix(t *testing.T) {
	data := "-1000000000"
	var ts Timestamp
	err := json.Unmarshal([]byte(data), &ts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := time.Date(1969, 12, 31, 23, 59, 59, 0, time.UTC)
	if !ts.Time.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, ts.Time)
	}
}

func TestUnmarshalJSON_Precision(t *testing.T) {
	data := "1577836800123456789"
	var ts Timestamp
	err := json.Unmarshal([]byte(data), &ts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// Expected precision is to the nearest microsecond because Go time.Time supports up to nanoseconds
	expected := time.Date(2020, 1, 1, 0, 0, 0, 123456789, time.UTC)
	if !ts.Time.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, ts.Time)
	}
}

func TestUnmarshalJSON_FutureTimestamp(t *testing.T) {
	data := "4102444800000000000"
	var ts Timestamp
	err := json.Unmarshal([]byte(data), &ts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)
	if !ts.Time.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, ts.Time)
	}
}
