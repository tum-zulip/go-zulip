package api

// Which [topics to unmute automatically in muted channels](zulip.com/help/mute-a-topic.
// - 1 - Topics the user participates in
// - 2 - Topics the user sends a message to
// - 3 - Topics the user starts
// - 4 - Never
// **Changes**: New in Zulip 8.0 (feature level 214).
type AutomaticallyUnmuteTopicsInMutedStreamsPolicy int

const (
	AutomaticallyUnmuteTopicsInMutedStreamsPolicyTopicsTheUserParticipatesIn  AutomaticallyUnmuteTopicsInMutedStreamsPolicy = 1
	AutomaticallyUnmuteTopicsInMutedStreamsPolicyTopicsTheUserSendsAMessageTo AutomaticallyUnmuteTopicsInMutedStreamsPolicy = 2
	AutomaticallyUnmuteTopicsInMutedStreamsPolicyTopicsTheUserStarts          AutomaticallyUnmuteTopicsInMutedStreamsPolicy = 3
	AutomaticallyUnmuteTopicsInMutedStreamsPolicyNever                        AutomaticallyUnmuteTopicsInMutedStreamsPolicy = 4
)

var AllowedAutomaticallyUnmuteTopicsInMutedStreamsPolicyEnumValues = []AutomaticallyUnmuteTopicsInMutedStreamsPolicy{
	AutomaticallyUnmuteTopicsInMutedStreamsPolicyTopicsTheUserParticipatesIn,
	AutomaticallyUnmuteTopicsInMutedStreamsPolicyTopicsTheUserSendsAMessageTo,
	AutomaticallyUnmuteTopicsInMutedStreamsPolicyTopicsTheUserStarts,
	AutomaticallyUnmuteTopicsInMutedStreamsPolicyNever,
}

func NewAutomaticallyUnmuteTopicsInMutedStreamsPolicyFromValue(v int) (*AutomaticallyUnmuteTopicsInMutedStreamsPolicy, error) {
	ev := AutomaticallyUnmuteTopicsInMutedStreamsPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedAutomaticallyUnmuteTopicsInMutedStreamsPolicyEnumValues,
			Value:   v,
			VarName: "AutomaticallyUnmuteTopicsInMutedStreamsPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutomaticallyUnmuteTopicsInMutedStreamsPolicy) IsValid() bool {
	for _, existing := range AllowedAutomaticallyUnmuteTopicsInMutedStreamsPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
