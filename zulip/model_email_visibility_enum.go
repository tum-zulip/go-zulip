package zulip

// The [policy] for [which other users] in this organization can see the user's real email address.
//   - EmailVisibilityEveryone = Everyone
//   - EmailVisibilityMembersOnly = Members only
//   - EmailVisibilityAdministratorsOnly = Administrators only
//   - EmailVisibilityNobody = Nobody
//   - EmailVisibilityModeratorsOnly = Moderators only
//
// **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.
//
// [policy]: https://zulip.com/api/roles-and-permissions#permission-levels
// [which other users]: https://zulip.com/help/configure-email-visibility
type EmailVisibility int

const (
	EmailVisibilityEveryone           EmailVisibility = 1
	EmailVisibilityMembersOnly        EmailVisibility = 2
	EmailVisibilityAdministratorsOnly EmailVisibility = 3
	EmailVisibilityNobody             EmailVisibility = 4
	EmailVisibilityModeratorsOnly     EmailVisibility = 5
)

var allowedEmailVisibilityEnumValues = []EmailVisibility{
	EmailVisibilityEveryone,
	EmailVisibilityMembersOnly,
	EmailVisibilityAdministratorsOnly,
	EmailVisibilityNobody,
	EmailVisibilityModeratorsOnly,
}

func NewEmailAddressVisibilityFromValue(v int) (*EmailVisibility, error) {
	ev := EmailVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    allowedEmailVisibilityEnumValues,
			Value:   v,
			VarName: "EmailVisibility",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailVisibility) IsValid() bool {
	for _, existing := range allowedEmailVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
