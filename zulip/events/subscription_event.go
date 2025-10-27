package events

import "github.com/tum-zulip/go-zulip/zulip"

// SubscriptionAddEvent Event sent to a user's clients when that user's channel subscriptions have changed (either the set of subscriptions or their properties).
type SubscriptionAddEvent struct {
	event

	// A list of dictionaries where each dictionary contains information about one of the subscribed channels.
	//
	// **Changes**: Removed `email_address` field from the dictionary in Zulip 8.0 (feature level 226).  Removed `role` field from the dictionary in Zulip 6.0 (feature level 133).
	Subscriptions []zulip.Subscription `json:"subscriptions,omitempty"`
}

// SubscriptionUpdateEvent Event sent to a user's clients when a property of the user's subscription to a channel has been updated. This event is used only for personal properties like `is_muted` or `pin_to_top`. See the [`stream op: update` event] for updates to global properties of a channel.
//
// [`stream op: update` event]: https://zulip.com/api/get-events#stream-update
type SubscriptionUpdateEvent struct {
	event

	ChannelID int64 `json:"stream_id,omitempty"`
	// The property of the subscription which has changed. For details on the various subscription properties that a user can change, see [POST /users/me/subscriptions/properties].  Clients should generally handle an unknown property received here without crashing, since that will naturally happen when connecting to a Zulip server running a new version that adds a new subscription property.
	//
	// **Changes**: As of Zulip 6.0 (feature level 139), updates to the `is_muted` property or the deprecated `in_home_view` property will send two `subscription` update events, one for each property, to support clients fully migrating to use the `is_muted` property. Prior to this feature level, updates to either property only sent one event with the deprecated `in_home_view` property.
	//
	// [POST /users/me/subscriptions/properties]: https://zulip.com/api/update-subscription-settings
	Property string      `json:"property,omitempty"`
	Value    UpdateValue `json:"value,omitempty"`
}

// SubscriptionRemoveEvent Event sent to a user's clients when that user has been unsubscribed from one or more channels.
type SubscriptionRemoveEvent struct {
	event

	// A list of dictionaries, where each dictionary contains information about one of the newly unsubscribed channels.
	Subscriptions []SubscriptionRemoveData `json:"subscriptions,omitempty"`
}

// SubscriptionPeerAddEvent Event sent when another user subscribes to a channel, or their subscription is newly visible to the current user.  When a user subscribes to a channel, the current user will receive this event only if they [have permission to see the channel's subscriber list]. When the current user gains permission to see a given channel's subscriber list, they will receive this event for the existing subscriptions to the channel.
//
// **Changes**: Prior to Zulip 8.0 (feature level 220), this event was incorrectly not sent to guest users when subscribers to web-public channels and subscribed public channels changed.  Prior to Zulip 8.0 (feature level 205), this event was not sent when a user gained access to a channel due to their [role changing].  Prior to Zulip 6.0 (feature level 134), this event was not sent when a private channel was made public.  In Zulip 4.0 (feature level 35), the singular `user_id` and `stream_id` integers included in this event were replaced with plural `user_ids` and `stream_ids` integer arrays.  In Zulip 3.0 (feature level 19), the `stream_id` field was added to identify the channel the user subscribed to, replacing the `name` field.
//
// [have permission to see the channel's subscriber list]: https://zulip.com/help/channel-permissions
// [role changing]: https://zulip.com/help/user-roles
type SubscriptionPeerAddEvent struct {
	event

	// The Ids of channels that have new or updated subscriber data.
	//
	// **Changes**: New in Zulip 4.0 (feature level 35), replacing the `stream_id` integer.
	ChannelIds []int64 `json:"stream_ids,omitempty"`
	// The Ids of the users who are newly visible as subscribed to the specified channels.
	//
	// **Changes**: New in Zulip 4.0 (feature level 35), replacing the `user_id` integer.
	UserIds []int64 `json:"user_ids,omitempty"`
}

// SubscriptionPeerRemoveEvent Event sent to other users when users have been unsubscribed from channels. Sent to all users if the channel is public or to only the existing subscribers if the channel is private.
//
// **Changes**: Prior to Zulip 8.0 (feature level 220), this event was incorrectly not sent to guest users when subscribers to web-public channels and subscribed public channels changed.  In Zulip 4.0 (feature level 35), the singular `user_id` and `stream_id` integers included in this event were replaced with plural `user_ids` and `stream_ids` integer arrays.  In Zulip 3.0 (feature level 19), the `stream_id` field was added to identify the channel the user unsubscribed from, replacing the `name` field.
type SubscriptionPeerRemoveEvent struct {
	event

	// The Ids of the channels from which the users have been unsubscribed from.  When a user is deactivated, the server will send this event removing the user's subscriptions before the `realm_user` event for the user's deactivation.
	//
	// **Changes**: Before Zulip 10.0 (feature level 377), this event was not sent on user deactivation. Clients supporting older server versions and maintaining peer subscriber data need to remove all channel subscriptions for a user when processing the `realm_user` event with `op="remove"`.
	//
	// **Changes**: New in Zulip 4.0 (feature level 35), replacing the `stream_id` integer.
	ChannelIds []int64 `json:"stream_ids,omitempty"`
	// The Ids of the users who have been unsubscribed.
	//
	// **Changes**: New in Zulip 4.0 (feature level 35), replacing the `user_id` integer.
	UserIds []int64 `json:"user_ids,omitempty"`
}

// SubscriptionRemoveData Dictionary containing details about the unsubscribed channel.
type SubscriptionRemoveData struct {
	// The Id of the channel.
	ChannelID int64 `json:"stream_id,omitempty"`
	// The name of the channel.
	Name string `json:"name,omitempty"`
}
