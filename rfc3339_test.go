package getstream

import (
	"encoding/json"
	"testing"
	"time"
)

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
