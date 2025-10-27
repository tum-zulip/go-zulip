package zulip

import (
	"encoding/json"
	"time"
)

// MutedUser Object containing the user ID and timestamp of a muted user.
type MutedUser struct {
	// The ID of the muted user.
	ID int64 `json:"id,omitempty"`
	// An integer UNIX timestamp representing when the user was muted.
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type mutedUserJSON struct {
	ID        int64 `json:"id,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}

func (m *MutedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&mutedUserJSON{
		ID:        m.ID,
		Timestamp: m.Timestamp.Unix(),
	})
}

func (m *MutedUser) UnmarshalJSON(data []byte) error {
	var r mutedUserJSON
	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	m.ID = r.ID
	m.Timestamp = time.Unix(r.Timestamp, 0)
	return nil
}
