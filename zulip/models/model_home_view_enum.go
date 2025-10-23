package models

// The [home view] used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.
//   - HomeViewRecentTopics = Recent conversations view
//   - HomeViewInbox = Inbox view
//   - HomeViewAllMessages = Combined feed view
//
// **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).
//
// [home view]: https://zulip.com/help/configure-home-view
type HomeView string

const (
	HomeViewRecentTopics HomeView = "recent_topics"
	HomeViewInbox        HomeView = "inbox"
	HomeViewAllMessages  HomeView = "all_messages"
)

var allowedHomeViewEnumValues = []HomeView{
	HomeViewRecentTopics,
	HomeViewInbox,
	HomeViewAllMessages,
}

func NewHomeViewFromValue(v string) (*HomeView, error) {
	ev := HomeView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    allowedHomeViewEnumValues,
			Value:   v,
			VarName: "HomeView",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HomeView) IsValid() bool {
	for _, existing := range allowedHomeViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
