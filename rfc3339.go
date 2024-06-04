package getstream

import (
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
	rfc3339 bool
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t.rfc3339 {
		return t.Time.MarshalJSON()
	}
	return t.formatUnix()
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	err := t.Time.UnmarshalJSON(data)
	if err != nil {
		return t.parseUnix(data)
	}
	t.rfc3339 = true
	return nil
}

func (t Timestamp) formatUnix() ([]byte, error) {
	sec := float64(t.Time.UnixNano()) * float64(time.Nanosecond) / float64(time.Second)
	return strconv.AppendFloat(nil, sec, 'f', -1, 64), nil
}

func (t *Timestamp) parseUnix(data []byte) error {
	f, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(0, int64(f*float64(time.Second/time.Nanosecond)))
	return nil
}
