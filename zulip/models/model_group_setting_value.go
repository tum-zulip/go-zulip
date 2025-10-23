package models

// GroupSettingValue - struct for GroupSettingValue
type GroupSettingValue struct {
	ComplexGroupSettingValue *ComplexGroupSettingValue
	GroupId                  *int64
}

// GroupSettingValueUpdate struct for GroupSettingValueUpdate
type GroupSettingValueUpdate struct {
	New GroupSettingValue  `json:"new"`
	Old *GroupSettingValue `json:"old,omitempty"`
}

// GroupSettingValueOneOf An object with these fields:
type ComplexGroupSettingValue struct {
	// The list of Ids of individual users in the collection of users with this permission.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 303), this list would include deactivated users who had the permission before being deactivated.
	DirectMembers []int64 `json:"direct_members"`
	// The list of Ids of the groups in the collection of users with this permission.
	DirectSubgroups []int64 `json:"direct_subgroups"`
}

// special json marshaler and unmarshaler for union GroupSettingValue
func (o GroupSettingValue) MarshalJSON() ([]byte, error) {
	return marshalUnionType(o)
}

// Unmarshal JSON data into one of the pointers in the struct
func (o *GroupSettingValue) UnmarshalJSON(data []byte) error {
	return unmarshalUnionType(data, o)
}
