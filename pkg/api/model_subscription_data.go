package api

import (
	"encoding/json"
	"fmt"

	"gopkg.in/validator.v2"
)

// SubscriptionData struct for SubscriptionData
type SubscriptionData struct {
	// The unique Id of a channel.
	StreamId int64 `json:"stream_id"`
	// One of the channel properties described below.
	// - `"color"`: The hex value of the user's display color for the channel.
	// - `"is_muted"`: Whether the channel is [muted](zulip.com/help/mute-a-channel).
	//   **Changes**: As of Zulip 6.0 (feature level 139), updating either `"is_muted"`
	//   or `"in_home_view"` generates two [subscription update
	//   events](zulip.com/api/get-events#subscription-update), one for each property.
	//   Prior to this feature level, updating either property only generated a
	//   subscription update event for `"in_home_view"`. Before Zulip 2.1.0, this
	//   feature was represented by the more confusingly named `"in_home_view"`
	//   (with the opposite value: `in_home_view=!is_muted`); for
	//   backwards-compatibility, modern Zulip still accepts that property.
	// - `"pin_to_top"`: Whether to pin the channel at the top of the channel list.
	// - `"desktop_notifications"`: Whether to show desktop notifications for all
	//   messages sent to the channel.
	// - `"audible_notifications"`: Whether to play a sound notification for all
	//   messages sent to the channel.
	// - `"push_notifications"`: Whether to trigger a mobile push notification for
	//   all messages sent to the channel.
	// - `"email_notifications"`: Whether to trigger an email notification for all
	//   messages sent to the channel.
	// - `"wildcard_mentions_notify"`: Whether wildcard mentions trigger
	//   notifications as though they were personal mentions in this channel.
	// - `"in_home_view"`: Legacy name for `"is_muted"`, provided for
	//   backwards-compatibility with older Zulip server versions.
	Property SubscriptionProperty `json:"property"`
	// SubscriptionDataValue - The new value of the property being modified.  If the property is `\"color\"`, then `value` is a string representing the hex value of the user's display color for the channel. For all other above properties, `value` is a boolean.
	Value SubscriptionDataValue `json:"value"`
}

type SubscriptionDataValue struct {
	Bool   *bool
	String *string
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionDataValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SubscriptionDataValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubscriptionDataValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionDataValue) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}
