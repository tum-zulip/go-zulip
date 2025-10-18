package zulip

import "encoding/json"

// Subscription struct for Subscription
type Subscription struct {
	Channel

	// A list of user Ids of users who are also subscribed to a given channel. Included only if `include_subscribers` is `true`.
	Subscribers []int64 `json:"subscribers,omitempty"`
	// If [`include_subscribers=\"partial\"`](https://zulip.com/api/get-subscriptions#parameter-include_subscribers was requested, the server may, at its discretion, send a `partial_subscribers` list rather than a `subscribers` list for channels with a large number of subscribers.  The `partial_subscribers` list contains an arbitrary subset of the channel's subscribers that is guaranteed to include all bot user subscribers as well as all users who have been active in the last 14 days, but otherwise can be chosen arbitrarily by the server.  **Changes**: New in Zulip 11.0 (feature level 412).
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
	// Whether the user has muted the channel. Muted channels do not count towards your total unread count and do not show up in the `Combined feed` view (previously known as `All messages`).  **Changes**: Prior to Zulip 2.1.0, this feature was represented by the more confusingly named `in_home_view` (with the opposite value, `in_home_view=!is_muted`).
	IsMuted bool `json:"is_muted,omitempty"`
	// Legacy property for if the given channel is muted, with inverted meaning.  **Changes**: Deprecated in Zulip 2.1.0. Clients should use `is_muted` where available.
	// Deprecated
	InHomeView bool `json:"in_home_view,omitempty"`
	// The user's personal color for the channel.
	Color string `json:"color,omitempty"`
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	var channelJSON channelJSON
	channelJSON.fromChannel(o.Channel)
	channelJSON.Subscribers = o.Subscribers
	channelJSON.PartialSubscribers = o.PartialSubscribers
	channelJSON.DesktopNotifications = o.DesktopNotifications
	channelJSON.EmailNotifications = o.EmailNotifications
	channelJSON.WildcardMentionsNotify = o.WildcardMentionsNotify
	channelJSON.PushNotifications = o.PushNotifications
	channelJSON.AudibleNotifications = o.AudibleNotifications
	channelJSON.PinToTop = o.PinToTop
	channelJSON.IsMuted = o.IsMuted
	channelJSON.InHomeView = o.InHomeView
	channelJSON.Color = o.Color

	return json.Marshal(channelJSON)
}

func (o *Subscription) UnmarshalJSON(data []byte) error {
	var channelJSON channelJSON
	if err := json.Unmarshal(data, &channelJSON); err != nil {
		return err
	}
	o.Channel.fromChannelJSON(channelJSON)

	o.Subscribers = channelJSON.Subscribers
	o.PartialSubscribers = channelJSON.PartialSubscribers
	o.DesktopNotifications = channelJSON.DesktopNotifications
	o.EmailNotifications = channelJSON.EmailNotifications
	o.WildcardMentionsNotify = channelJSON.WildcardMentionsNotify
	o.PushNotifications = channelJSON.PushNotifications
	o.AudibleNotifications = channelJSON.AudibleNotifications
	o.PinToTop = channelJSON.PinToTop
	o.IsMuted = channelJSON.IsMuted
	o.InHomeView = channelJSON.InHomeView
	o.Color = channelJSON.Color

	return nil
}
