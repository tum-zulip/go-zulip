package zulip

import (
	"github.com/tum-zulip/go-zulip/zulip/internal/union"
)

// GroupSettingValue - struct for GroupSettingValue.
type GroupSettingValue struct {
	ComplexGroupSettingValue *ComplexGroupSettingValue
	GroupID                  *int64
}

// GroupSettingValueUpdate struct for GroupSettingValueUpdate.
type GroupSettingValueUpdate struct {
	New GroupSettingValue  `json:"new"`
	Old *GroupSettingValue `json:"old,omitempty"`
}

// ComplexGroupSettingValue struct for ComplexGroupSettingValue.
type ComplexGroupSettingValue struct {
	// The list of IDs of individual users in the collection of users with this permission.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 303), this list would include deactivated users who had the permission before being deactivated.
	DirectMembers []int64 `json:"direct_members"`
	// The list of IDs of the groups in the collection of users with this permission.
	DirectSubgroups []int64 `json:"direct_subgroups"`
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (o GroupSettingValue) MarshalJSON() ([]byte, error) {
	return union.Marshal(o)
}

// Unmarshal JSON data into one of the pointers in the struct.
func (o *GroupSettingValue) UnmarshalJSON(data []byte) error {
	return union.Unmarshal(data, o)
}
