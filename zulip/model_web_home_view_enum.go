package zulip

// The [home view] used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.
//   - "recent_topics" - Recent conversations view
//   - "inbox" - Inbox view
//   - "all_messages" - Combined feed view
//
// **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).
//
// [home view]: https://zulip.com/help/configure-home-view
type WebHomeView string

const (
	WebHomeViewRecentTopics WebHomeView = "recent_topics"
	WebHomeViewInbox        WebHomeView = "inbox"
	WebHomeViewAllMessages  WebHomeView = "all_messages"
)

var AllowedWebHomeViewEnumValues = []WebHomeView{
	WebHomeViewRecentTopics,
	WebHomeViewInbox,
	WebHomeViewAllMessages,
}

func NewWebHomeViewFromValue(v string) (*WebHomeView, error) {
	ev := WebHomeView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedWebHomeViewEnumValues,
			Value:   v,
			VarName: "WebHomeView",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebHomeView) IsValid() bool {
	for _, existing := range AllowedWebHomeViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
