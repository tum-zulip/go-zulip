package models

// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.
//   - UnreadsCountDisplayAllChannels = All channels
//   - UnreadsCountDisplayUnmutedChannelsAndTopics = Unmuted channels and topics
//   - UnreadsCountDisplayNoChannels = No channels
//
// **Changes**: New in Zulip 8.0 (feature level 210).
type UnreadsCountDisplay int

const (
	UnreadsCountDisplayAllChannels              UnreadsCountDisplay = 1
	UnreadsCountDisplayUnmutedChannelsAndTopics UnreadsCountDisplay = 2
	UnreadsCountDisplayNoChannels               UnreadsCountDisplay = 3
)

var allowedUnreadsCountDisplayEnumValues = []UnreadsCountDisplay{
	UnreadsCountDisplayAllChannels,
	UnreadsCountDisplayUnmutedChannelsAndTopics,
	UnreadsCountDisplayNoChannels,
}

func NewUnreadsCountDisplayFromValue(v int) (*UnreadsCountDisplay, error) {
	ev := UnreadsCountDisplay(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    allowedUnreadsCountDisplayEnumValues,
			Value:   v,
			VarName: "UnreadsCountDisplay",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnreadsCountDisplay) IsValid() bool {
	for _, existing := range allowedUnreadsCountDisplayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
