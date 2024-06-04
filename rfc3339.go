package getstream

import (
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
	rfc3339 bool
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData == "null" {
		t.Time = time.Time{}
		return nil
	}

	// Try parsing as RFC3339 first
	parsedTime, err := time.Parse(`"`+time.RFC3339Nano+`"`, strData)
	if err == nil {
		t.Time = parsedTime
		t.rfc3339 = true
		return nil
	}

	// Try parsing as Unix nanoseconds if RFC3339 fails
	ns, err := strconv.ParseInt(strData, 10, 64)
	if err == nil {
		t.Time = time.Unix(0, ns).UTC()
		return nil
	}

	// Handle high precision float values represented as strings
	floatNs, err := strconv.ParseFloat(strData, 64)
	if err != nil {
		return err
	}
	integerNs := int64(floatNs)
	t.Time = time.Unix(0, integerNs).UTC()
	return nil
}
