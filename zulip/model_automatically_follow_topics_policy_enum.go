package zulip

// Which [topics to follow automatically](https://zulip.com/help/mute-a-topic).
//   - 1 = Topics the user participates in
//   - 2 = Topics the user sends a message to
//   - 3 = Topics the user starts
//   - 4 = Never
//
// **Changes**: New in Zulip 8.0 (feature level 214).
type AutomaticallyFollowTopicsPolicy int

const (
	AutomaticallyFollowTopicsPolicyTopicsTheUserParticipatesIn  AutomaticallyFollowTopicsPolicy = 1
	AutomaticallyFollowTopicsPolicyTopicsTheUserSendsAMessageTo AutomaticallyFollowTopicsPolicy = 2
	AutomaticallyFollowTopicsPolicyTopicsTheUserStarts          AutomaticallyFollowTopicsPolicy = 3
	AutomaticallyFollowTopicsPolicyNever                        AutomaticallyFollowTopicsPolicy = 4
)

var AllowedAutomaticallyFollowTopicsPolicyEnumValues = []AutomaticallyFollowTopicsPolicy{
	AutomaticallyFollowTopicsPolicyTopicsTheUserParticipatesIn,
	AutomaticallyFollowTopicsPolicyTopicsTheUserSendsAMessageTo,
	AutomaticallyFollowTopicsPolicyTopicsTheUserStarts,
	AutomaticallyFollowTopicsPolicyNever,
}

func NewAutomaticallyFollowTopicsPolicyFromValue(v int) (*AutomaticallyFollowTopicsPolicy, error) {
	ev := AutomaticallyFollowTopicsPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedAutomaticallyFollowTopicsPolicyEnumValues,
			Value:   v,
			VarName: "AutomaticallyFollowTopicsPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutomaticallyFollowTopicsPolicy) IsValid() bool {
	for _, existing := range AllowedAutomaticallyFollowTopicsPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
