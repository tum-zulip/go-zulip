package zulip

type MessageRetentionDays string

const (
	MessageRetentionDaysRealmDefault MessageRetentionDays = "realm_default"
	MessageRetentionDaysUnlimited    MessageRetentionDays = "unlimited"
	MessageRetentionDaysForever      MessageRetentionDays = "forever" // Deprecated: use MessageRetentionDaysUnlimited
)

var allowedMessageRetentionDaysValues = []MessageRetentionDays{
	MessageRetentionDaysRealmDefault,
	MessageRetentionDaysUnlimited,
	MessageRetentionDaysForever,
}

func NewMessageRetationDaysFromValue(v string) (*MessageRetentionDays, error) {
	ev := MessageRetentionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    allowedMessageRetentionDaysValues,
			VarName: "Role",
		}
	}
}

func (v MessageRetentionDays) IsValid() bool {
	for _, m := range allowedMessageRetentionDaysValues {
		if m == v {
			return true
		}
	}
	return false
}
