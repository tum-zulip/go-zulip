package zulip

// Unread count badge (appears in desktop sidebar and browser tab)
//   - BadgeCountAllUnreadMessages = All unread messages
//   - BadgeCountDMsMentionsAndFollowedTopics = DMs, mentions, and followed topics
//   - BadgeCountDMsAndMentions = DMs and mentions
//   - BadgeCountNone = None
//
// **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.
type BadgeCount int

const (
	BadgeCountAllUnreadMessages            BadgeCount = 1
	BadgeCountDMsMentionsAndFollowedTopics BadgeCount = 2
	BadgeCountDMsAndMentions               BadgeCount = 3
	BadgeCountNone                         BadgeCount = 4
)

var allowedBadgeCountEnumValues = []BadgeCount{
	BadgeCountAllUnreadMessages,
	BadgeCountDMsMentionsAndFollowedTopics,
	BadgeCountDMsAndMentions,
	BadgeCountNone,
}

func NewDesktopIconCountDisplayFromValue(v int) (*BadgeCount, error) {
	ev := BadgeCount(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    allowedBadgeCountEnumValues,
			Value:   v,
			VarName: "BadgeCount",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BadgeCount) IsValid() bool {
	for _, existing := range allowedBadgeCountEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
