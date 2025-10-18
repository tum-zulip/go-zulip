package zulip

// Which [topics to unmute automatically in muted channels](https://zulip.com/help/mute-a-topic).
// - 1 - Topics the user participates in
// - 2 - Topics the user sends a message to
// - 3 - Topics the user starts
// - 4 - Never
// **Changes**: New in Zulip 8.0 (feature level 214).
type AutomaticallyUnmuteTopicsInMutedChannelsPolicy int

const (
	AutomaticallyUnmuteTopicsInMutedChannelsPolicyTopicsTheUserParticipatesIn  AutomaticallyUnmuteTopicsInMutedChannelsPolicy = 1
	AutomaticallyUnmuteTopicsInMutedChannelsPolicyTopicsTheUserSendsAMessageTo AutomaticallyUnmuteTopicsInMutedChannelsPolicy = 2
	AutomaticallyUnmuteTopicsInMutedChannelsPolicyTopicsTheUserStarts          AutomaticallyUnmuteTopicsInMutedChannelsPolicy = 3
	AutomaticallyUnmuteTopicsInMutedChannelsPolicyNever                        AutomaticallyUnmuteTopicsInMutedChannelsPolicy = 4
)

var AllowedAutomaticallyUnmuteTopicsInMutedChannelsPolicyEnumValues = []AutomaticallyUnmuteTopicsInMutedChannelsPolicy{
	AutomaticallyUnmuteTopicsInMutedChannelsPolicyTopicsTheUserParticipatesIn,
	AutomaticallyUnmuteTopicsInMutedChannelsPolicyTopicsTheUserSendsAMessageTo,
	AutomaticallyUnmuteTopicsInMutedChannelsPolicyTopicsTheUserStarts,
	AutomaticallyUnmuteTopicsInMutedChannelsPolicyNever,
}

func NewAutomaticallyUnmuteTopicsInMutedChannelsPolicyFromValue(v int) (*AutomaticallyUnmuteTopicsInMutedChannelsPolicy, error) {
	ev := AutomaticallyUnmuteTopicsInMutedChannelsPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedAutomaticallyUnmuteTopicsInMutedChannelsPolicyEnumValues,
			Value:   v,
			VarName: "AutomaticallyUnmuteTopicsInMutedChannelsPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutomaticallyUnmuteTopicsInMutedChannelsPolicy) IsValid() bool {
	for _, existing := range AllowedAutomaticallyUnmuteTopicsInMutedChannelsPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
