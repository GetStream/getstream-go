package getstream

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	Time *time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	strData := string(data)
	if strData == "null" {
		return nil
	}

	ns, err := strconv.ParseInt(strData, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}

	t.Time = PtrTo(time.Unix(0, ns).UTC())
	return nil
}
