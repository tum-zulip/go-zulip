package api

// Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.
// - 1 - Top topic in the channel
// - 2 - Channel feed
// - 3 - List of topics
// - 4 - Top unread topic in channel
// **Changes**: The \"Top unread topic in channel\" is new in Zulip 11.0 (feature level 401).  The \"List of topics\" option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \"Channel feed\" behavior.
type ChannelDefaultView int

const (
	ChannelDefaultViewTopTopicInChannel       ChannelDefaultView = 1
	ChannelDefaultViewChannelFeed             ChannelDefaultView = 2
	ChannelDefaultViewListOfTopics            ChannelDefaultView = 3
	ChannelDefaultViewTopUnreadTopicInChannel ChannelDefaultView = 4
)

var AllowedChannelDefaultViewEnumValues = []ChannelDefaultView{
	ChannelDefaultViewTopTopicInChannel,
	ChannelDefaultViewChannelFeed,
	ChannelDefaultViewListOfTopics,
	ChannelDefaultViewTopUnreadTopicInChannel,
}

func NewChannelDefaultViewFromValue(v int) (*ChannelDefaultView, error) {
	ev := ChannelDefaultView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedChannelDefaultViewEnumValues,
			Value:   v,
			VarName: "ChannelDefaultView",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelDefaultView) IsValid() bool {
	for _, existing := range AllowedChannelDefaultViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
