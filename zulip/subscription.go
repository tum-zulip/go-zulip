package zulip

import "encoding/json"

// Subscription struct for Subscription.
type Subscription struct {
	Channel

	// A list of user IDs of users who are also subscribed to a given channel. Included only if `include_subscribers` is `true`.
	Subscribers []int64 `json:"subscribers,omitempty"`
	// If [`include_subscribers="partial"`].
	//
	// [`include_subscribers="partial"`]: https://zulip.com/api/get-subscriptions#parameter-include_subscribers was requested, the server may, at its discretion, send a `partial_subscribers` list rather than a `subscribers` list for channels with a large number of subscribers.  The `partial_subscribers` list contains an arbitrary subset of the channel's subscribers that is guaranteed to include all bot user subscribers as well as all users who have been active in the last 14 days, but otherwise can be chosen arbitrarily by the server.
	//
	// **Changes**: New in Zulip 11.0 (feature level 412
	PartialSubscribers []int64 `json:"partial_subscribers,omitempty"`
	// A boolean specifying whether desktop notifications are enabled for the given channel.  A `null` value means the value of this setting should be inherited from the user-level default setting, `enable_stream_desktop_notifications`, for this channel.
	DesktopNotifications *bool `json:"desktop_notifications,omitempty"`
	// A boolean specifying whether email notifications are enabled for the given channel.  A `null` value means the value of this setting should be inherited from the user-level default setting, `enable_stream_email_notifications`, for this channel.
	EmailNotifications *bool `json:"email_notifications,omitempty"`
	// A boolean specifying whether wildcard mentions trigger notifications as though they were personal mentions in this channel.  A `null` value means the value of this setting should be inherited from the user-level default setting, wildcard_mentions_notify, for this channel.
	WildcardMentionsNotify *bool `json:"wildcard_mentions_notify,omitempty"`
	// A boolean specifying whether push notifications are enabled for the given channel.  A `null` value means the value of this setting should be inherited from the user-level default setting, `enable_stream_push_notifications`, for this channel.
	PushNotifications *bool `json:"push_notifications,omitempty"`
	// A boolean specifying whether audible notifications are enabled for the given channel.  A `null` value means the value of this setting should be inherited from the user-level default setting, `enable_stream_audible_notifications`, for this channel.
	AudibleNotifications *bool `json:"audible_notifications,omitempty"`
	// A boolean specifying whether the given channel has been pinned to the top.
	PinToTop bool `json:"pin_to_top,omitempty"`
	// Whether the user has muted the channel. Muted channels do not count towards your total unread count and do not show up in the `Combined feed` view (previously known as `All messages`).
	//
	// **Changes**: Prior to Zulip 2.1.0, this feature was represented by the more confusingly named `in_home_view` (with the opposite value, `in_home_view=!is_muted`).
	IsMuted bool `json:"is_muted,omitempty"`
	// Legacy property for if the given channel is muted, with inverted meaning.
	//
	// **Changes**: Deprecated in Zulip 2.1.0. Clients should use `is_muted` where available.
	// Deprecated
	InHomeView bool `json:"in_home_view,omitempty"`
	// The user's personal color for the channel.
	Color string `json:"color,omitempty"`
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	var obj channelJSON
	obj.fromChannel(o.Channel)
	obj.Subscribers = o.Subscribers
	obj.PartialSubscribers = o.PartialSubscribers
	obj.DesktopNotifications = o.DesktopNotifications
	obj.EmailNotifications = o.EmailNotifications
	obj.WildcardMentionsNotify = o.WildcardMentionsNotify
	obj.PushNotifications = o.PushNotifications
	obj.AudibleNotifications = o.AudibleNotifications
	obj.PinToTop = o.PinToTop
	obj.IsMuted = o.IsMuted
	obj.InHomeView = o.InHomeView
	obj.Color = o.Color

	return json.Marshal(obj)
}

func (o *Subscription) UnmarshalJSON(data []byte) error {
	var obj channelJSON
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	o.Channel.fromChannelJSON(obj)

	o.Subscribers = obj.Subscribers
	o.PartialSubscribers = obj.PartialSubscribers
	o.DesktopNotifications = obj.DesktopNotifications
	o.EmailNotifications = obj.EmailNotifications
	o.WildcardMentionsNotify = obj.WildcardMentionsNotify
	o.PushNotifications = obj.PushNotifications
	o.AudibleNotifications = obj.AudibleNotifications
	o.PinToTop = obj.PinToTop
	o.IsMuted = obj.IsMuted
	o.InHomeView = obj.InHomeView
	o.Color = obj.Color

	return nil
}
