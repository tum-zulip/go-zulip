package events

import (
	"encoding/json"
	"time"
)

// RestartEvent Event sent to all the users whenever the Zulip server restarts.  Specifically, this event is sent whenever the Tornado process for the user is restarted; in particular, this will always happen when the Zulip server is upgraded.  Clients should use this event to update their tracking of the server's capabilities, and to decide if they wish to get a new event queue after a server upgrade. Clients doing so must implement a random delay strategy to spread such restarts over 5 minutes or more to avoid creating a synchronized thundering herd effect.
//
// **Changes**: Removed the `immediate` flag, which was only used by web clients in development, in Zulip 9.0 (feature level 240).
type RestartEvent struct {
	event

	// The Zulip version number, in the format where this appears in the [server_settings] and [register] Responses.
	//
	// **Changes**: New in Zulip 4.0 (feature level 59).
	//
	// [server_settings]: https://zulip.com/api/get-server-settings
	// [register]: https://zulip.com/api/register-queue
	ZulipVersion string `json:"zulip_version,omitempty"`
	// The Zulip merge base number, in the format where this appears in the [server_settings] and [register] Responses.
	//
	// **Changes**: New in Zulip 5.0 (feature level 88).
	//
	// [server_settings]: https://zulip.com/api/get-server-settings
	// [register]: https://zulip.com/api/register-queue
	ZulipMergeBase string `json:"zulip_merge_base,omitempty"`
	// The [Zulip feature level] of the server after the restart.  Clients should use this to update their tracking of the server's capabilities, and may choose to refetch their state and create a new event queue when the API feature level has changed in a way that the client finds significant. Clients choosing to do so must implement a random delay strategy to spread such restarts over 5 or more minutes to avoid creating a synchronized thundering herd effect.
	//
	// **Changes**: New in Zulip 4.0 (feature level 59).
	//
	// [Zulip feature level]: https://zulip.com/api/changelog
	ZulipFeatureLevel int32 `json:"zulip_feature_level,omitempty"`
	// The timestamp at which the server started.
	ServerGeneration time.Time `json:"server_generation,omitempty"`
}

type restartEventJSON struct {
	event

	ZulipVersion      string  `json:"zulip_version,omitempty"`
	ZulipMergeBase    string  `json:"zulip_merge_base,omitempty"`
	ZulipFeatureLevel int32   `json:"zulip_feature_level,omitempty"`
	ServerGeneration  float64 `json:"server_generation,omitempty"`
}

// UnmarshalJSON Custom unmarshaler to handle conversion of `server_generation` from a float timestamp to a `time.Time` object.
func (e *RestartEvent) UnmarshalJSON(data []byte) error {
	var rej restartEventJSON
	if err := json.Unmarshal(data, &rej); err != nil {
		return err
	}

	e.event = rej.event
	e.ZulipVersion = rej.ZulipVersion
	e.ZulipMergeBase = rej.ZulipMergeBase
	e.ZulipFeatureLevel = rej.ZulipFeatureLevel
	e.ServerGeneration = time.Unix(int64(rej.ServerGeneration), 0)

	return nil
}

func (e *RestartEvent) MarshalJSON() ([]byte, error) {
	rej := restartEventJSON{
		event:             e.event,
		ZulipVersion:      e.ZulipVersion,
		ZulipMergeBase:    e.ZulipMergeBase,
		ZulipFeatureLevel: e.ZulipFeatureLevel,
		ServerGeneration:  float64(e.ServerGeneration.Unix()),
	}
	return json.Marshal(rej)
}
