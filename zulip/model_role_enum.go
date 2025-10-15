package zulip

type Role int

const (
	RoleOwner     Role = 100
	RoleAdmin     Role = 200
	RoleModerator Role = 300
	RoleMember    Role = 400
	RoleGuest     Role = 600
)

// All allowed values of Role enum
var AllowedRoleEnumValues = []Role{
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
			Enum:    AllowedRoleEnumValues,
			VarName: "Role",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Role) IsValid() bool {
	for _, existing := range AllowedRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
