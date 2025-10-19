package zulip

// Topic interaction types Zulip supports
//   - TopicInteractionyTopicsTheUserParticipatesIn = Topics the user participates in
//   - TopicInteractionyTopicsTheUserSendsAMessageTo = Topics the user sends a message to
//   - TopicInteractionyTopicsTheUserStarts = Topics the user starts
//   - TopicInteractionyNever = Never
//
// These can be used to set which topics to unmute or follow automatically.
type TopicInteraction int

const (
	TopicInteractionyTopicsTheUserParticipatesIn  TopicInteraction = 1
	TopicInteractionyTopicsTheUserSendsAMessageTo TopicInteraction = 2
	TopicInteractionyTopicsTheUserStarts          TopicInteraction = 3
	TopicInteractionyNever                        TopicInteraction = 4
)

var AllowedTopicInteractionEnumValues = []TopicInteraction{
	TopicInteractionyTopicsTheUserParticipatesIn,
	TopicInteractionyTopicsTheUserSendsAMessageTo,
	TopicInteractionyTopicsTheUserStarts,
	TopicInteractionyNever,
}

func NewTopicInteractionFromValue(v int) (*TopicInteraction, error) {
	ev := TopicInteraction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedTopicInteractionEnumValues,
			Value:   v,
			VarName: "TopicInteraction",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TopicInteraction) IsValid() bool {
	for _, existing := range AllowedTopicInteractionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
