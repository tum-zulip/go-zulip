package api

import (
	"encoding/json"
	"time"
)

// UserSettingsUpdateEvent3MutedUsersInner Object containing the user Id and timestamp of a muted user.
type MutedUser struct {
	// The Id of the muted user.
	Id int64 `json:"id,omitempty"`
	// An integer UNIX timestamp representing when the user was muted.
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type mutedUserJSON struct {
	Id        int64 `json:"id,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}

func (m *MutedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&mutedUserJSON{
		Id:        m.Id,
		Timestamp: m.Timestamp.Unix(),
	})
}

func (m *MutedUser) UnmarshalJSON(data []byte) error {
	var r mutedUserJSON
	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	m.Id = r.Id
	m.Timestamp = time.Unix(r.Timestamp, 0)
	return nil
}
