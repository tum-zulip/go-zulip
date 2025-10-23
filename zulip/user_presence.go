package zulip

import (
	"encoding/json"
	"time"
)

type UserPresence struct {
	// When this update was received. If the timestamp is more than a few minutes in the past, the user is offline.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Whether the user had recently interacted with Zulip at the time of the timestamp.  Will be either `"active"` or `"idle"`
	Status string `json:"status,omitempty"`
}

type userPresenceJSON struct {
	Timestamp int64  `json:"timestamp,omitempty"`
	Status    string `json:"status,omitempty"`
}

func (o UserPresence) MarshalJSON() ([]byte, error) {

	aux := userPresenceJSON{
		Status:    o.Status,
		Timestamp: o.Timestamp.Unix(),
	}

	return json.Marshal(aux)
}

func (o *UserPresence) UnmarshalJSON(data []byte) error {
	var aux userPresenceJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.Status = aux.Status
	o.Timestamp = time.Unix(aux.Timestamp, 0).UTC()

	return nil
}
