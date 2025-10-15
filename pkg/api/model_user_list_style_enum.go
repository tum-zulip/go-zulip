package api

// The style selected by the user for the right sidebar user list.
// - 1 - Compact
// - 2 - With status
// - 3 - With avatar and status
// **Changes**: New in Zulip 6.0 (feature level 141).
type UserListStyle int

const (
	UserListStyleCompact             UserListStyle = 1
	UserListStyleWithStatus          UserListStyle = 2
	UserListStyleWithAvatarAndStatus UserListStyle = 3
)

var AllowedUserListStyleEnumValues = []UserListStyle{
	UserListStyleCompact,
	UserListStyleWithStatus,
	UserListStyleWithAvatarAndStatus,
}

func NewUserListStyleFromValue(v int) (*UserListStyle, error) {
	ev := UserListStyle(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedUserListStyleEnumValues,
			Value:   v,
			VarName: "UserListStyle",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserListStyle) IsValid() bool {
	for _, existing := range AllowedUserListStyleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
