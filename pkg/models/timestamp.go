package models

import (
	"encoding/json"
	"time"
)

type Timestamp int64

func (t Timestamp) ToInt64() int64 {
	return int64(t)
}

func (t *Timestamp) FromInt64(v int64) {
	*t = Timestamp(v)
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToInt64())
}

func (t *Timestamp) ToTime() time.Time {
	return time.Unix(int64(*t), 0)
}

func (t *Timestamp) FromTime(v time.Time) {
	*t = Timestamp(v.Unix())
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	t.FromInt64(v)
	return nil
}

func (t Timestamp) String() string {
	return t.ToTime().String()
}

// printers for log and slog
func (t Timestamp) GoString() string {
	return t.String()
}

func (t Timestamp) Format(s string) string {
	return t.ToTime().Format(s)
}

func (t Timestamp) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *Timestamp) UnmarshalText(data []byte) error {
	tt, err := time.Parse(time.RFC3339, string(data))
	if err != nil {
		return err
	}
	t.FromTime(tt)
	return nil
}
