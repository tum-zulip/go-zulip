package api

import (
	"encoding/json"
	"fmt"
)

// GroupSettingValue - struct for GroupSettingValue
type GroupSettingValue struct {
	ComplexGroupSettingValue *ComplexGroupSettingValue
	GroupId                  *int64
}

// GroupSettingValueOneOf An object with these fields:
type ComplexGroupSettingValue struct {
	// The list of Ids of individual users in the collection of users with this permission.  **Changes**: Prior to Zulip 10.0 (feature level 303), this list would include deactivated users who had the permission before being deactivated.
	DirectMembers []int64 `json:"direct_members,omitempty"`
	// The list of Ids of the groups in the collection of users with this permission.
	DirectSubgroups []int64 `json:"direct_subgroups,omitempty"`
}

// special json marshaller and unmarshaller for union GroupSettingValue
func (o GroupSettingValue) MarshalJSON() ([]byte, error) {
	if o.ComplexGroupSettingValue != nil {
		return json.Marshal(&o.ComplexGroupSettingValue)
	}

	if o.GroupId != nil {
		return json.Marshal(&o.GroupId)
	}

	return nil, nil // no data in union
}

// Unmarshal JSON data into one of the pointers in the struct
func (o *GroupSettingValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ComplexGroupSettingValue
	err = newStrictDecoder(data).Decode(&o.ComplexGroupSettingValue)
	if err == nil {
		match++
	} else {
		o.ComplexGroupSettingValue = nil
	}

	// try to unmarshal data into GroupId
	err = newStrictDecoder(data).Decode(&o.GroupId)
	if err == nil {
		match++
	} else {
		o.GroupId = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		o.ComplexGroupSettingValue = nil
		o.GroupId = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GroupSettingValue)")
	} else if match == 0 { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GroupSettingValue)")
	}

	return nil
}
