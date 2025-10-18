package zulip

// The [policy] for [which other users] in this organization can see the user's real email address.
//   - 1 = Everyone
//   - 2 = Members only
//   - 3 = Administrators only
//   - 4 = Nobody
//   - 5 = Moderators only
//
// **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.
//
// [policy]: https://zulip.com/api/roles-and-permissions#permission-levels
// [which other users]: https://zulip.com/help/configure-email-visibility
type EmailAddressVisibility int

const (
	EmailAddressVisibilityEveryone           EmailAddressVisibility = 1
	EmailAddressVisibilityMembersOnly        EmailAddressVisibility = 2
	EmailAddressVisibilityAdministratorsOnly EmailAddressVisibility = 3
	EmailAddressVisibilityNobody             EmailAddressVisibility = 4
	EmailAddressVisibilityModeratorsOnly     EmailAddressVisibility = 5
)

var AllowedEmailAddressVisibilityEnumValues = []EmailAddressVisibility{
	EmailAddressVisibilityEveryone,
	EmailAddressVisibilityMembersOnly,
	EmailAddressVisibilityAdministratorsOnly,
	EmailAddressVisibilityNobody,
	EmailAddressVisibilityModeratorsOnly,
}

func NewEmailAddressVisibilityFromValue(v int) (*EmailAddressVisibility, error) {
	ev := EmailAddressVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedEmailAddressVisibilityEnumValues,
			Value:   v,
			VarName: "EmailAddressVisibility",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailAddressVisibility) IsValid() bool {
	for _, existing := range AllowedEmailAddressVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
