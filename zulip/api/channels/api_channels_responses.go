package channels

import "github.com/tum-zulip/go-zulip/zulip"

type CreateChannelResponse struct {
	zulip.Response

	// The ID of the newly created channel.
	ID int64 `json:"id,omitempty"`
}

type CreateChannelFolderResponse struct {
	zulip.Response

	// The unique ID of the created channel folder.
	ChannelFolderID int64 `json:"channel_folder_id"`
}

type MarkAllAsReadResponse struct {
	zulip.Response

	// Whether all unread messages were marked as read.  Will be `false` if the request successfully marked some, but not all, messages as read.
	Complete bool `json:"complete"`
}

type GetChannelFoldersResponse struct {
	zulip.Response

	// A list of channel folder objects.
	ChannelFolders []zulip.ChannelFolder `json:"channel_folders,omitempty"`
}

type GetChannelResponse struct {
	zulip.Response

	Channel zulip.Channel `json:"stream,omitempty"`
}

// GetChannelsResponse struct for GetChannelsResponse.
type GetChannelsResponse struct {
	zulip.Response

	// A list of channel objects with details on the requested channels.
	Channels []zulip.Channel `json:"streams,omitempty"`
}

type GetChannelEmailAddressResponse struct {
	zulip.Response

	// Email address of the channel.
	Email string `json:"email,omitempty"`
}

// GetChannelIDResponse struct for GetChannelIDResponse.
type GetChannelIDResponse struct {
	zulip.Response

	// The ID of the given channel.
	ChannelID int64 `json:"stream_id,omitempty"`
}

// GetChannelTopicsResponse struct for GetChannelTopicsResponse.
type GetChannelTopicsResponse struct {
	zulip.Response

	// An array of objects with information about user-accessible topics in the specified channel, sorted by recency (i.e., the topic with the most recent message is ordered first).
	Topics []Topic `json:"topics,omitempty"`
}

// Topic struct for Topic.
type Topic struct {
	// The message ID of the last message sent to this topic.
	MaxID int64 `json:"max_id,omitempty"`
	// The name of the topic.
	Name string `json:"name,omitempty"`
}

// GetSubscribersResponse struct for GetSubscribersResponse.
type GetSubscribersResponse struct {
	zulip.Response

	// A list containing the IDs of all active users who are subscribed to the channel.
	Subscribers []int64 `json:"subscribers,omitempty"`
}

// GetSubscriptionStatusResponse struct for GetSubscriptionStatusResponse.
type GetSubscriptionStatusResponse struct {
	zulip.Response

	// Whether the user is subscribed to the channel.
	IsSubscribed bool `json:"is_subscribed,omitempty"`
}

// GetSubscriptionsResponse struct for GetSubscriptionsResponse.
type GetSubscriptionsResponse struct {
	zulip.Response

	// A list of dictionaries where each dictionary contains information about one of the subscribed channels.
	//
	// **Changes**: Removed `email_address` field from the dictionary in Zulip 8.0 (feature level 226).  Removed `role` field from the dictionary in Zulip 6.0 (feature level 133).
	Subscriptions []zulip.Subscription `json:"subscriptions"`
}

// SubscribeResponse struct for SubscribeResponse.
type SubscribeResponse struct {
	zulip.Response

	// A dictionary where the key is the ID of the user and the value is a list of the names of the channels that user was subscribed to as a result of the request.
	//
	// **Changes**: Before Zulip 10.0 (feature level 289), the user keys were Zulip API email addresses, not user ID.
	Subscribed map[string][]string `json:"subscribed,omitempty"`
	// A dictionary where the key is the ID of the user and the value is a list of the names of the channels that where the user was not added as a subscriber in this request, because they were already a subscriber.
	//
	// **Changes**: Before Zulip 10.0 (feature level 289), the user keys were Zulip API email addresses, not user IDs.
	AlreadySubscribed map[string][]string `json:"already_subscribed,omitempty"`
	// A list of names of channels that the requesting user/bot was not authorized to subscribe to. Only present if `"authorization_errors_fatal": false`.
	Unauthorized []string `json:"unauthorized,omitempty"`
	// Only present if the parameter `send_new_subscription_messages` in the request was `true`.  Whether Notification Bot DMs in fact sent to the added subscribers as requested by the `send_new_subscription_messages` parameter. Clients may find this value useful to communicate with users about the effect of this request.
	//
	// **Changes**: New in Zulip 11.0 (feature level 397).
	NewSubscriptionMessagesSent bool `json:"new_subscription_messages_sent,omitempty"`
}

// CreateBigBlueButtonVideoCallResponse struct for CreateBigBlueButtonVideoCallResponse.
type CreateBigBlueButtonVideoCallResponse struct {
	zulip.Response

	// The URL for the BigBlueButton video call.
	URL string `json:"url,omitempty"`
}

// UnsubscribeResponse struct for UnsubscribeResponse.
type UnsubscribeResponse struct {
	zulip.Response

	// A list of the names of channels that the user is already unsubscribed from, and hence doesn't need to be unsubscribed.
	NotRemoved []string `json:"not_removed,omitempty"`
	// A list of the names of channels which were unsubscribed from as a result of the query.
	Removed []string `json:"removed,omitempty"`
}

// UpdateSubscriptionsResponse struct for UpdateSubscriptionsResponse.
type UpdateSubscriptionsResponse struct {
	zulip.Response

	// A dictionary where the key is the Zulip API email address of the user/bot and the value is a list of the names of the channels that were subscribed to as a result of the query.
	Subscribed map[string][]string `json:"subscribed"`
	// A dictionary where the key is the Zulip API email address of the user/bot and the value is a list of the names of the channels that the user/bot is already subscribed to.
	AlreadySubscribed map[string][]string `json:"already_subscribed"`
	// A list of the names of channels that the user is already unsubscribed from, and hence doesn't need to be unsubscribed.
	NotRemoved []string `json:"not_removed,omitempty"`
	// A list of the names of channels which were unsubscribed from as a result of the query.
	Removed []string `json:"removed"`
	// Only present if the parameter `send_new_subscription_messages` in the request was `true`.  Whether Notification Bot DMs in fact sent to the added subscribers as requested by the `send_new_subscription_messages` parameter. Clients may find this value useful to communicate with users about the effect of this request.
	//
	// **Changes**: New in Zulip 11.0 (feature level 397).
	NewSubscriptionMessagesSent *bool `json:"new_subscription_messages_sent,omitempty"`
}
