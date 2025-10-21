package zulip

// Whether or not to mark messages as read when the user scrolls through their feed.
//   - MarkReadOnScrollPolicyAlways = Always
//   - MarkReadOnScrollPolicyOnlyInConversationViews = Only in conversation views
//   - MarkReadOnScrollPolicyNever = Never
//
// **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the "Always" setting when marking messages as read.
type MarkReadOnScrollPolicy int

const (
	MarkReadOnScrollPolicyAlways                  MarkReadOnScrollPolicy = 1
	MarkReadOnScrollPolicyOnlyInConversationViews MarkReadOnScrollPolicy = 2
	MarkReadOnScrollPolicyNever                   MarkReadOnScrollPolicy = 3
)

var allowedMarkReadOnScrollPolicyEnumValues = []MarkReadOnScrollPolicy{
	MarkReadOnScrollPolicyAlways,
	MarkReadOnScrollPolicyOnlyInConversationViews,
	MarkReadOnScrollPolicyNever,
}

func NewMarkReadOnScrollPolicyFromValue(v int) (*MarkReadOnScrollPolicy, error) {
	ev := MarkReadOnScrollPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    allowedMarkReadOnScrollPolicyEnumValues,
			Value:   v,
			VarName: "MarkReadOnScrollPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarkReadOnScrollPolicy) IsValid() bool {
	for _, existing := range allowedMarkReadOnScrollPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
