package zulip

// One of the channel properties described below:
//   - SubscriptionPropertyColor
//   - SubscriptionPropertyIsMuted
//   - SubscriptionPropertyPinToTop
//   - SubscriptionPropertyDesktopNotifications
//   - SubscriptionPropertyAudibleNotifications
//   - SubscriptionPropertyPushNotifications
//   - SubscriptionPropertyEmailNotifications
//   - SubscriptionPropertyWildcardMentionsNotify
//   - SubscriptionPropertyInHomeView
//
// [muted]: https://zulip.com/help/mute-a-channel
// [subscription update events]: https://zulip.com/api/get-events#subscription-update
type SubscriptionProperty string

const (
	// Whether the channel is [muted].
	//
	// **Changes**: As of Zulip 6.0 (feature level 139), updating either "is_muted" or "in_home_view" generates two [subscription update events], one for each property, that are sent to clients. Prior to this feature level, updating either property only generated a subscription update event for "in_home_view". Prior to Zulip 2.1.0, this feature was represented by the more confusingly named "in_home_view" (with the opposite value: `in_home_view=!is_muted`); for backwards-compatibility, modern Zulip still accepts that property.
	//
	// [muted]: https://zulip.com/help/mute-a-channel
	// [subscription update events]: https://zulip.com/api/get-events#subscription-update
	SubscriptionPropertyIsMuted                SubscriptionProperty = "is_muted"
	SubscriptionPropertyColor                  SubscriptionProperty = "color"                    // The hex value of the user's display color for the channel.
	SubscriptionPropertyPinToTop               SubscriptionProperty = "pin_to_top"               // Whether to pin the channel at the top of the channel list.
	SubscriptionPropertyDesktopNotifications   SubscriptionProperty = "desktop_notifications"    // Whether to show desktop notifications for all messages sent to the channel.
	SubscriptionPropertyAudibleNotifications   SubscriptionProperty = "audible_notifications"    // Whether to play a sound notification for all messages sent to the channel.
	SubscriptionPropertyPushNotifications      SubscriptionProperty = "push_notifications"       // Whether to trigger a mobile push notification for all messages sent to the channel.
	SubscriptionPropertyEmailNotifications     SubscriptionProperty = "email_notifications"      // Whether to trigger an email notification for all messages sent to the channel.
	SubscriptionPropertyWildcardMentionsNotify SubscriptionProperty = "wildcard_mentions_notify" // Whether wildcard mentions trigger notifications as though they were personal mentions in this channel.
	SubscriptionPropertyInHomeView             SubscriptionProperty = "in_home_view"             // Legacy name for `is_muted`, provided for backwards-compatibility with older Zulip server versions.
)

var AllowedSubscriptionPropertyEnumValues = []SubscriptionProperty{
	SubscriptionPropertyColor,
	SubscriptionPropertyIsMuted,
	SubscriptionPropertyPinToTop,
	SubscriptionPropertyDesktopNotifications,
	SubscriptionPropertyAudibleNotifications,
	SubscriptionPropertyPushNotifications,
	SubscriptionPropertyEmailNotifications,
	SubscriptionPropertyWildcardMentionsNotify,
	SubscriptionPropertyInHomeView,
}

func NewSubscriptionPropertyFromValue(v string) (*SubscriptionProperty, error) {
	ev := SubscriptionProperty(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, &ErrInvalidEnumValue{
		Value:   v,
		Enum:    AllowedSubscriptionPropertyEnumValues,
		VarName: "SubscriptionProperty",
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionProperty) IsValid() bool {
	for _, existing := range AllowedSubscriptionPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
