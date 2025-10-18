package zulip

// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.
//   - 1 = All channels
//   - 2 = Unmuted channels and topics
//   - 3 = No channels
//
// **Changes**: New in Zulip 8.0 (feature level 210).
type WebChannelUnreadsCountDisplayPolicy int

const (
	WebChannelUnreadsCountDisplayPolicyAllChannels              WebChannelUnreadsCountDisplayPolicy = 1
	WebChannelUnreadsCountDisplayPolicyUnmutedChannelsAndTopics WebChannelUnreadsCountDisplayPolicy = 2
	WebChannelUnreadsCountDisplayPolicyNoChannels               WebChannelUnreadsCountDisplayPolicy = 3
)

var AllowedWebChannelUnreadsCountDisplayPolicyEnumValues = []WebChannelUnreadsCountDisplayPolicy{
	WebChannelUnreadsCountDisplayPolicyAllChannels,
	WebChannelUnreadsCountDisplayPolicyUnmutedChannelsAndTopics,
	WebChannelUnreadsCountDisplayPolicyNoChannels,
}

func NewWebChannelUnreadsCountDisplayPolicyFromValue(v int) (*WebChannelUnreadsCountDisplayPolicy, error) {
	ev := WebChannelUnreadsCountDisplayPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedWebChannelUnreadsCountDisplayPolicyEnumValues,
			Value:   v,
			VarName: " WebChannelUnreadsCountDisplayPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebChannelUnreadsCountDisplayPolicy) IsValid() bool {
	for _, existing := range AllowedWebChannelUnreadsCountDisplayPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
