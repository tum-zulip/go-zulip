package api

// Whether to [include organization name in subject of message notification emails](zulip.com/help/email-notifications#include-organization-name-in-subject-line.
// - 1 - Automatic
// - 2 - Always
// - 3 - Never
// **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.
type NameInEmailNotificationsPolicy int

const (
	NameInEmailNotificationsPolicyAutomatic NameInEmailNotificationsPolicy = 1
	NameInEmailNotificationsPolicyAlways    NameInEmailNotificationsPolicy = 2
	NameInEmailNotificationsPolicyNever     NameInEmailNotificationsPolicy = 3
)

var AllowedNameInEmailNotificationsPolicyEnumValues = []NameInEmailNotificationsPolicy{
	NameInEmailNotificationsPolicyAutomatic,
	NameInEmailNotificationsPolicyAlways,
	NameInEmailNotificationsPolicyNever,
}

func NewNameInEmailNotificationsPolicyFromValue(v int) (*NameInEmailNotificationsPolicy, error) {
	ev := NameInEmailNotificationsPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedNameInEmailNotificationsPolicyEnumValues,
			Value:   v,
			VarName: "NameInEmailNotificationsPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NameInEmailNotificationsPolicy) IsValid() bool {
	for _, existing := range AllowedNameInEmailNotificationsPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
