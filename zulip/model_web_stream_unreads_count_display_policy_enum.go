package zulip

// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.
// - 1 - All channels
// - 2 - Unmuted channels and topics
// - 3 - No channels
// **Changes**: New in Zulip 8.0 (feature level 210).
type WebStreamUnreadsCountDisplayPolicy int

const (
	WebStreamUnreadsCountDisplayPolicyAllChannels              WebStreamUnreadsCountDisplayPolicy = 1
	WebStreamUnreadsCountDisplayPolicyUnmutedChannelsAndTopics WebStreamUnreadsCountDisplayPolicy = 2
	WebStreamUnreadsCountDisplayPolicyNoChannels               WebStreamUnreadsCountDisplayPolicy = 3
)

var AllowedWebStreamUnreadsCountDisplayPolicyEnumValues = []WebStreamUnreadsCountDisplayPolicy{
	WebStreamUnreadsCountDisplayPolicyAllChannels,
	WebStreamUnreadsCountDisplayPolicyUnmutedChannelsAndTopics,
	WebStreamUnreadsCountDisplayPolicyNoChannels,
}

func NewWebStreamUnreadsCountDisplayPolicyFromValue(v int) (*WebStreamUnreadsCountDisplayPolicy, error) {
	ev := WebStreamUnreadsCountDisplayPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedWebStreamUnreadsCountDisplayPolicyEnumValues,
			Value:   v,
			VarName: "WebStreamUnreadsCountDisplayPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebStreamUnreadsCountDisplayPolicy) IsValid() bool {
	for _, existing := range AllowedWebStreamUnreadsCountDisplayPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
