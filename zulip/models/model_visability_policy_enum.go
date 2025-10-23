package models

// Controls visibility policy
//   - VisibilityPolicyNone = None. Removes the visibility policy previously set for the topic.
//   - VisibilityPolicyMuted = Muted. [Mutes the topic] in a channel.
//   - VisibilityPolicyUnmuted = Unmuted. [Unmutes the topic] in a muted channel.
//   - VisibilityPolicyFollowed = Followed. [Follows the topic].
//
// [Mutes the topic]: https://zulip.com/help/mute-a-topic
// [Follows the topic]: https://zulip.com/help/follow-a-topic
type VisibilityPolicy int

const (
	VisibilityPolicyNone     VisibilityPolicy = 0
	VisibilityPolicyMuted    VisibilityPolicy = 1
	VisibilityPolicyUnmuted  VisibilityPolicy = 2
	VisibilityPolicyFollowed VisibilityPolicy = 3
)

var allowedVisibilityPolicyEnumValues = []VisibilityPolicy{
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
			Enum:    allowedVisibilityPolicyEnumValues,
			VarName: "VisibilityPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VisibilityPolicy) IsValid() bool {
	for _, existing := range allowedVisibilityPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
