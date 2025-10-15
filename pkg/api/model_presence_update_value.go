package api

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/validator.v2"
)

// PresenceUpdateValue - Will be one of these two formats (modern or legacy) for user presence data:
// ModernPresenceFormat `{user_id}`: Presence data (modern format) for the user with the specified Id.
type PresenceUpdateValue struct {
	ModernPresenceFormat *ModernPresenceFormat
	LegacyPresenceMap    map[string]LegacyPresenceFormat
}

type ModernPresenceFormat struct {
	// The UNIX timestamp of the last time a client connected to Zulip reported that the user was actually present (e.g. via focusing a browser window or interacting with a computer running the desktop app).  Clients should display users with a current `active_timestamp` as fully present.
	ActiveTimestamp time.Time `json:"active_timestamp,omitempty"`
	// The UNIX timestamp of the last time the user had a client connected to Zulip, including idle clients where the user hasn't interacted with the system recently.  The Zulip server has no way of distinguishing whether an idle web app user is at their computer, but hasn't interacted with the Zulip tab recently, or simply left their desktop computer on when they left.  Thus, clients should display users with a current `idle_timestamp` but no current `active_timestamp` as potentially present.
	IdleTimestamp time.Time `json:"idle_timestamp,omitempty"`
}

type modernPresenceFormatJSON struct {
	ActiveTimestamp int64 `json:"active_timestamp"`
	IdleTimestamp   int64 `json:"idle_timestamp"`
}

func (o ModernPresenceFormat) MarshalJSON() ([]byte, error) {
	aux := modernPresenceFormatJSON{
		ActiveTimestamp: o.ActiveTimestamp.Unix(),
		IdleTimestamp:   o.IdleTimestamp.Unix(),
	}
	return json.Marshal(aux)
}

func (o *ModernPresenceFormat) UnmarshalJSON(data []byte) error {
	var aux modernPresenceFormatJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.ActiveTimestamp = time.Unix(aux.ActiveTimestamp, 0).UTC()
	o.IdleTimestamp = time.Unix(aux.IdleTimestamp, 0).UTC()

	return nil
}

type LegacyPresenceFormat struct {
	// The client's platform name.  **Changes**: Starting with Zulip 7.0 (feature level 178), this will always be `\"website\"` as the server no longer stores which client submitted presence data.
	Client string `json:"client,omitempty"`
	// The status of the user on this client. Will be either `\"idle\"` or `\"active\"`.
	Status PresenceStatus `json:"status,omitempty"`
	// The UNIX timestamp of when this client sent the user's presence to the server with the precision of a second.
	Timestamp int32 `json:"timestamp,omitempty"`
	// Whether the client is capable of showing mobile/push notifications to the user.  Not present in objects with the `\"aggregated\"` key.  **Changes**: Starting with Zulip 7.0 (feature level 178), always `false` when present as the server no longer stores which client submitted presence data.
	Pushable bool `json:"pushable,omitempty"`
}

// ModernPresenceFormatAsPresenceUpdateValue is a convenience function that returns ModernPresenceFormat wrapped in PresenceUpdateValue
func ModernPresenceFormatAsPresenceUpdateValue(v ModernPresenceFormat) PresenceUpdateValue {
	return PresenceUpdateValue{
		ModernPresenceFormat: &v,
	}
}

// LegacyPresenceMapAsPresenceUpdateValue is a convenience function that returns map[string]LegacyPresenceFormat wrapped in PresenceUpdateValue
func LegacyPresenceMapAsPresenceUpdateValue(v map[string]LegacyPresenceFormat) PresenceUpdateValue {
	return PresenceUpdateValue{
		LegacyPresenceMap: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PresenceUpdateValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ModernPresenceFormat
	err = newStrictDecoder(data).Decode(&dst.ModernPresenceFormat)
	if err == nil {
		jsonModernPresenceFormat, _ := json.Marshal(dst.ModernPresenceFormat)
		if string(jsonModernPresenceFormat) == "{}" { // empty struct
			dst.ModernPresenceFormat = nil
		} else {
			if err = validator.Validate(dst.ModernPresenceFormat); err != nil {
				dst.ModernPresenceFormat = nil
			} else {
				match++
			}
		}
	} else {
		dst.ModernPresenceFormat = nil
	}

	// try to unmarshal data into LegacyPresenceMap
	err = newStrictDecoder(data).Decode(&dst.LegacyPresenceMap)
	if err == nil {
		jsonLegacyPresenceMap, _ := json.Marshal(dst.LegacyPresenceMap)
		if string(jsonLegacyPresenceMap) == "{}" { // empty struct
			dst.LegacyPresenceMap = nil
		} else {
			if err = validator.Validate(dst.LegacyPresenceMap); err != nil {
				dst.LegacyPresenceMap = nil
			} else {
				match++
			}
		}
	} else {
		dst.LegacyPresenceMap = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ModernPresenceFormat = nil
		dst.LegacyPresenceMap = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PresenceUpdateValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PresenceUpdateValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PresenceUpdateValue) MarshalJSON() ([]byte, error) {
	if src.ModernPresenceFormat != nil {
		return json.Marshal(&src.ModernPresenceFormat)
	}

	if src.LegacyPresenceMap != nil {
		return json.Marshal(&src.LegacyPresenceMap)
	}

	return nil, nil // no data in oneOf schemas
}
