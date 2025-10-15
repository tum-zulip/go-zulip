package zulip

// Unread count badge (appears in desktop sidebar and browser tab)
// - 1 - All unread messages
// - 2 - DMs, mentions, and followed topics
// - 3 - DMs and mentions
// - 4 - None
// **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.
type DesktopIconCountDisplay int

const (
	DesktopIconCountDisplayAllUnreadMessages            DesktopIconCountDisplay = 1
	DesktopIconCountDisplayDMsMentionsAndFollowedTopics DesktopIconCountDisplay = 2
	DesktopIconCountDisplayDMsAndMentions               DesktopIconCountDisplay = 3
	DesktopIconCountDisplayNone                         DesktopIconCountDisplay = 4
)

var AllowedDesktopIconCountDisplayEnumValues = []DesktopIconCountDisplay{
	DesktopIconCountDisplayAllUnreadMessages,
	DesktopIconCountDisplayDMsMentionsAndFollowedTopics,
	DesktopIconCountDisplayDMsAndMentions,
	DesktopIconCountDisplayNone,
}

func NewDesktopIconCountDisplayFromValue(v int) (*DesktopIconCountDisplay, error) {
	ev := DesktopIconCountDisplay(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedDesktopIconCountDisplayEnumValues,
			Value:   v,
			VarName: "DesktopIconCountDisplay",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DesktopIconCountDisplay) IsValid() bool {
	for _, existing := range AllowedDesktopIconCountDisplayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
