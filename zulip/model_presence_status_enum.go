package zulip

// PresenceStatus types Zulip supports
//   - PresenceStatusActive = active
//   - PresenceStatusIdle = idle
//
// These can be used to represent a user's presence status.
type PresenceStatus string

const (
	PresenceStatusActive PresenceStatus = "active"
	PresenceStatusIdle   PresenceStatus = "idle"
)

var allowedPresenceStatusEnumValues = []PresenceStatus{
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
			Enum:    allowedPresenceStatusEnumValues,
			VarName: "PresenceStatus",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PresenceStatus) IsValid() bool {
	for _, existing := range allowedPresenceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
