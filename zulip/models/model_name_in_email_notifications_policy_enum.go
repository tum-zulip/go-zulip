package models

// Whether to [include organization name in subject of message notification emails].
//   - NameInEmailNotificationsPolicyAutomatic
//   - NameInEmailNotificationsPolicyAlways
//   - NameInEmailNotificationsPolicyNever
//
// **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.
//
// [include organization name in subject of message notification emails]: https://zulip.com/help/email-notifications#include-organization-name-in-subject-line
type NameInEmailNotificationsPolicy int

const (
	NameInEmailNotificationsPolicyAutomatic NameInEmailNotificationsPolicy = 1
	NameInEmailNotificationsPolicyAlways    NameInEmailNotificationsPolicy = 2
	NameInEmailNotificationsPolicyNever     NameInEmailNotificationsPolicy = 3
)

var allowedNameInEmailNotificationsPolicyEnumValues = []NameInEmailNotificationsPolicy{
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
			Enum:    allowedNameInEmailNotificationsPolicyEnumValues,
			Value:   v,
			VarName: "NameInEmailNotificationsPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NameInEmailNotificationsPolicy) IsValid() bool {
	for _, existing := range allowedNameInEmailNotificationsPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
