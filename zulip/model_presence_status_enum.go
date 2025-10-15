package zulip

type PresenceStatus string

const (
	PresenceStatusActive PresenceStatus = "active"
	PresenceStatusIdle   PresenceStatus = "idle"
)

var AllowedPresenceStatusEnumValues = []PresenceStatus{
	PresenceStatusActive,
	PresenceStatusIdle,
}

func NewPresenceStatusFromValue(v string) (*PresenceStatus, error) {
	ev := PresenceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    AllowedPresenceStatusEnumValues,
			VarName: "PresenceStatus",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PresenceStatus) IsValid() bool {
	for _, existing := range AllowedPresenceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
