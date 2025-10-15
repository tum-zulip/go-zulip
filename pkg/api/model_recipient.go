package api

import (
	"encoding/json"
	"fmt"

	"gopkg.in/validator.v2"
)

// Recipients - A message's tentative target audience.
// For channel messages, the integer Id of the channel.
// For direct messages, a list containing integer user Ids.
type Recipient struct {
	Users   *[]int64
	Channel *int64
}

// UsersAsRecipient is a convenience function that returns a list of UserIds wrapped in Recipient
func UsersAsRecipient(v []int64) Recipient {
	return Recipient{
		Users: &v,
	}
}

// UsersEmailsAsRecipient is a convenience function that returns a list of UserEmails wrapped in Recipient
func UsersEmailsAsRecipient(v []string) Recipient {
	return Recipient{
		Users:   nil,
		Channel: nil,
	}
}

// UserAsRecipient is a convenience function that returns UserId wrapped in Recipient
func UserAsRecipient(v int64) Recipient {
	return Recipient{
		Users: &[]int64{v},
	}
}

// ChannelAsReipient is a convenience function that returns ChannelId wrapped in Recipient
func ChannelAsReipient(v int64) Recipient {
	return Recipient{
		Channel: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Recipient) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UserIds
	err = newStrictDecoder(data).Decode(&dst.Users)
	if err == nil {
		jsonUserIds, _ := json.Marshal(dst.Users)
		if string(jsonUserIds) == "{}" { // empty struct
			dst.Users = nil
		} else {
			if err = validator.Validate(dst.Users); err != nil {
				dst.Users = nil
			} else {
				match++
			}
		}
	} else {
		dst.Users = nil
	}

	// try to unmarshal data into ChannelId
	err = newStrictDecoder(data).Decode(&dst.Channel)
	if err == nil {
		jsonChannelId, _ := json.Marshal(dst.Channel)
		if string(jsonChannelId) == "{}" { // empty struct
			dst.Channel = nil
		} else {
			if err = validator.Validate(dst.Channel); err != nil {
				dst.Channel = nil
			} else {
				match++
			}
		}
	} else {
		dst.Channel = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Users = nil
		dst.Channel = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Recipient)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Recipient)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Recipient) MarshalJSON() ([]byte, error) {
	if src.Users != nil {
		return json.Marshal(&src.Users)
	}

	if src.Channel != nil {
		return json.Marshal(&src.Channel)
	}

	return nil, nil // no data in oneOf schemas
}
