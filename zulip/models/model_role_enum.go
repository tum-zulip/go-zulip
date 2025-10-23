package models

// [Organization-level role] of the user. Possible values are:
//   - RoleOwner = Organization owner
//   - RoleAdmin = Organization administrator
//   - RoleModerator = Organization moderator
//   - RoleMember = Member
//   - RoleGuest = Guest
//
// **Changes**: New in Zulip 4.0 (feature level 59).
//
// [Organization-level role]: https://zulip.com/api/roles-and-permissions
type Role int

// [Organization-level role] of the user. Possible values are:
//   - RoleOwner = Organization owner
//   - RoleAdmin = Organization administrator
//   - RoleModerator = Organization moderator
//   - RoleMember = Member
//   - RoleGuest = Guest
//
// **Changes**: New in Zulip 4.0 (feature level 59).
//
// [Organization-level role]: https://zulip.com/api/roles-and-permissions
const (
	RoleOwner     Role = 100
	RoleAdmin     Role = 200
	RoleModerator Role = 300
	RoleMember    Role = 400
	RoleGuest     Role = 600
)

// All allowed values of Role enum
var allowedRoleEnumValues = []Role{
	RoleMember,
	RoleGuest,
	RoleModerator,
	RoleAdmin,
	RoleOwner,
}

// NewInviteRoleParameterFromValue returns a pointer to a valid InviteRoleParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleFromValue(v int) (*Role, error) {
	ev := Role(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    allowedRoleEnumValues,
			VarName: "Role",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Role) IsValid() bool {
	for _, existing := range allowedRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
