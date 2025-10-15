package api

import (
	"encoding/json"
	"fmt"

	"gopkg.in/validator.v2"
)

// Principals - A list of user Ids (preferred) or Zulip API email addresses of the users to be subscribed to or unsubscribed from the channels specified in the `subscriptions` parameter. If not provided, then the requesting user/bot is subscribed.  **Changes**: The integer format is new in Zulip 3.0 (feature level 9).
type Principals struct {
	UserIds    *[]int64
	UserEmails *[]string
}

// []UserIdsAsPrincipals is a convenience function that returns []int64 wrapped in Principals
func UserIdsAsPrincipals(v []int64) Principals {
	return Principals{
		UserIds: &v,
	}
}

// []UserEmailsAsPrincipals is a convenience function that returns []string wrapped in Principals
func UserEmailsAsPrincipals(v []string) Principals {
	return Principals{
		UserEmails: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Principals) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UserIds
	err = newStrictDecoder(data).Decode(&dst.UserIds)
	if err == nil {
		jsonUserIds, _ := json.Marshal(dst.UserIds)
		if string(jsonUserIds) == "{}" { // empty struct
			dst.UserIds = nil
		} else {
			if err = validator.Validate(dst.UserIds); err != nil {
				dst.UserIds = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserIds = nil
	}

	// try to unmarshal data into UserEmails
	err = newStrictDecoder(data).Decode(&dst.UserEmails)
	if err == nil {
		jsonUserEmails, _ := json.Marshal(dst.UserEmails)
		if string(jsonUserEmails) == "{}" { // empty struct
			dst.UserEmails = nil
		} else {
			if err = validator.Validate(dst.UserEmails); err != nil {
				dst.UserEmails = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserEmails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UserIds = nil
		dst.UserEmails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Principals)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Principals)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Principals) MarshalJSON() ([]byte, error) {
	if src.UserIds != nil {
		return json.Marshal(&src.UserIds)
	}

	if src.UserEmails != nil {
		return json.Marshal(&src.UserEmails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Principals) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UserIds != nil {
		return obj.UserIds
	}

	if obj.UserEmails != nil {
		return obj.UserEmails
	}

	// all schemas are nil
	return nil
}
