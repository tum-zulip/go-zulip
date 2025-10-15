package zulip

// Controls visibility policy
// 0 = None. Removes the visibility policy previously set for the topic.
// 1 = Muted. Mutes the topic in a channel.
// 2 = Unmuted. Unmutes the topic in a muted channel.
// 3 = Followed. Follows the topic.
type VisibilityPolicy int

const (
	VisibilityPolicyNone     VisibilityPolicy = 0
	VisibilityPolicyMuted    VisibilityPolicy = 1
	VisibilityPolicyUnmuted  VisibilityPolicy = 2
	VisibilityPolicyFollowed VisibilityPolicy = 3
)

var AllowedVisibilityPolicyEnumValues = []VisibilityPolicy{
	VisibilityPolicyNone,
	VisibilityPolicyMuted,
	VisibilityPolicyUnmuted,
	VisibilityPolicyFollowed,
}

func NewVisibilityPolicyFromValue(v int) (*VisibilityPolicy, error) {
	ev := VisibilityPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    AllowedVisibilityPolicyEnumValues,
			VarName: "VisibilityPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VisibilityPolicy) IsValid() bool {
	for _, existing := range AllowedVisibilityPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
