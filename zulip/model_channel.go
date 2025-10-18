package zulip

import (
	"encoding/json"
	"time"
)

type Channel struct {
	// The unique Id of the channel.
	ChannelId int64 `json:"stream_id"`
	// The name of the channel.
	Name string `json:"name"`
	// A boolean indicating whether the channel is [archived](https://zulip.com/help/archive-a-channel).  **Changes**: New in Zulip 10.0 (feature level 315). Previously, this endpoint never returned archived channels.
	IsArchived bool `json:"is_archived"`
	// The short description of the channel in [Zulip-flavored Markdown](https://zulip.com/help/format-your-message-using-markdown) format, intended to be used to prepopulate UI for editing a channel's description.  See [Markdown message formatting](https://zulip.com/api/message-formatting) for details on Zulip's HTML format.
	Description string `json:"description"`
	// The UNIX timestamp for when the channel was created, in UTC seconds.  **Changes**: New in Zulip 4.0 (feature level 30).
	DateCreated time.Time `json:"date_created"`
	CreatorId   *int64    `json:"creator_id"`
	// Specifies whether the channel is private or not. Only people who have been invited can access a private channel.
	InviteOnly bool `json:"invite_only"`
	// The short description of the channel rendered as HTML, intended to be used when displaying the channel description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.
	RenderedDescription string `json:"rendered_description"`
	// Whether the channel has been configured to allow unauthenticated access to its message history from the web.  **Changes**: New in Zulip 2.1.0.
	IsWebPublic bool `json:"is_web_public"`
	// A deprecated representation of a superset of the users who have permission to post messages to the channel available for backwards-compatibility. Clients should use `can_send_message_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:
	//   - 1 = Any user can post.
	//   - 2 = Only administrators can post.
	//   - 3 = Only [full members] can post.
	//   - 4 = Only moderators can post.
	//
	// **Changes**: Deprecated in Zulip 10.0 (feature level 333) and replaced by `can_send_message_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  New in Zulip 3.0 (feature level 1), replacing the previous `is_announcement_only` boolean.
	//
	// [full members]: https://zulip.com/api/roles-and-permissions#determining-if-a-user-is-a-full-member
	// Deprecated
	ChannelPostPolicy    int32 `json:"stream_post_policy"`
	MessageRetentionDays *int  `json:"message_retention_days"`
	// Whether the history of the channel is public to its subscribers.  Currently always true for public channels (i.e. `\"invite_only\": false` implies `\"history_public_to_subscribers\": true`), but clients should not make that assumption, as we may change that behavior in the future.
	HistoryPublicToSubscribers bool         `json:"history_public_to_subscribers"`
	TopicsPolicy               TopicsPolicy `json:"topics_policy,omitempty"`
	// The Id of the first message in the channel.  Intended to help clients determine whether they need to display UI like the \"show all topics\" widget that would suggest the channel has older history that can be accessed.  Is `null` for channels with no message history.
	FirstMessageId *int64 `json:"first_message_id"`
	// The Id of the folder to which the channel belongs.  Is `null` if channel does not belong to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).
	FolderId *int64 `json:"folder_id"`
	// Whether the channel has recent message activity. Clients should use this to implement [hide inactive channels](https://zulip.com/help/manage-inactive-channels) if `demote_inactive_streams` is enabled.  **Changes**: New in Zulip 10.0 (feature level 323). Previously, clients implemented the demote_inactive_streams from local message history, resulting in a choppy loading experience.
	IsRecentlyActive bool `json:"is_recently_active"`
	// Whether the given channel is announcement only or not.  **Changes**: Deprecated in Zulip 3.0 (feature level 1). Clients should use `stream_post_policy` instead.
	// Deprecated
	IsAnnouncementOnly                bool              `json:"is_announcement_only"`
	CanAddSubscribersGroup            GroupSettingValue `json:"can_add_subscribers_group,omitempty"`
	CanRemoveSubscribersGroup         GroupSettingValue `json:"can_remove_subscribers_group"`
	CanAdministerChannelGroup         GroupSettingValue `json:"can_administer_channel_group,omitempty"`
	CanDeleteAnyMessageGroup          GroupSettingValue `json:"can_delete_any_message_group,omitempty"`
	CanDeleteOwnMessageGroup          GroupSettingValue `json:"can_delete_own_message_group,omitempty"`
	CanMoveMessagesOutOfChannelGroup  GroupSettingValue `json:"can_move_messages_out_of_channel_group,omitempty"`
	CanMoveMessagesWithinChannelGroup GroupSettingValue `json:"can_move_messages_within_channel_group,omitempty"`
	CanSendMessageGroup               GroupSettingValue `json:"can_send_message_group,omitempty"`
	CanSubscribeGroup                 GroupSettingValue `json:"can_subscribe_group"`
	CanResolveTopicsGroup             GroupSettingValue `json:"can_resolve_topics_group,omitempty"`
	// The total number of non-deactivated users (including bots) who are subscribed to the channel. Clients are responsible for updating this value using `peer_add` and `peer_remove` events.  The server's internals cannot guarantee this value is correctly synced with `peer_add` and `peer_remove` events for the channel. As a result, if a (rare) race occurs between a change in the channel's subscribers and fetching this value, it is possible for a client that is correctly following the events protocol to end up with a permanently off-by-one error in the channel's subscriber count.  Clients are recommended to fetch full subscriber data for a channel in contexts where it is important to avoid this risk. The official web application, for example, uses this field primarily while waiting to fetch a given channel's full subscriber list from the server.  **Changes**: New in Zulip 11.0 (feature level 394).
	SubscriberCount float32 `json:"subscriber_count"`
	// The average number of messages sent to the channel per week, as estimated based on recent weeks, rounded to the nearest integer.  If `null`, no information is provided on the average traffic. This can be because the channel was recently created and there is insufficient data to make an estimate, or because the server wishes to omit this information for this client, this realm, or this endpoint or type of event.  **Changes**: New in Zulip 8.0 (feature level 199). Previously, this statistic was available only in subscription objects.
	ChannelWeeklyTraffic *int `json:"stream_weekly_traffic"`
}

type ChannelWithIsDefault struct {
	Channel
	// Only present when [`include_default`] parameter is `true`.  Whether the given channel is a [default channel](https://zulip.com/help/set-default-channels-for-new-users).  [`include_default`]: https://zulip.com/api/get-streams#parameter-include_default
	IsDefault bool `json:"is_default,omitempty"`
}

// ChannelWithSubscribers struct for ChannelWithSubscribers
type ChannelWithSubscribers struct {
	Channel

	// A list of user Ids of users who are subscribed to the channel. Included only if `include_subscribers` is `true`.  If a user is not allowed to know the subscribers for a channel, we will send an empty array. API authors should use other data to determine whether users like guest users are forbidden to know the subscribers.
	Subscribers []int64 `json:"subscribers,omitempty"`
	// If [`include_subscribers=\"partial\"`](https://zulip.com/api/get-subscriptions#parameter-include_subscribers was requested, the server may, at its discretion, send a `partial_subscribers` list rather than a `subscribers` list for channels with a large number of subscribers.  The `partial_subscribers` list contains an arbitrary subset of the channel's subscribers that is guaranteed to include all bot user subscribers as well as all users who have been active in the last 14 days, but otherwise can be chosen arbitrarily by the server.  If a user is not allowed to know the subscribers for a channel, we will send an empty array. API authors should use other data to determine whether users like guest users are forbidden to know the subscribers.  **Changes**: New in Zulip 11.0 (feature level 412).
	PartialSubscribers []int64 `json:"partial_subscribers,omitempty"`
}

func (o Channel) MarshalJSON() ([]byte, error) {
	var channelJSON channelJSON
	channelJSON.fromChannel(o)
	return json.Marshal(&channelJSON)
}

func (o *Channel) UnmarshalJSON(data []byte) error {
	var err error
	var channelJSON channelJSON
	err = newStrictDecoder(data).Decode(&channelJSON)
	if err != nil {
		return err
	}
	o.fromChannelJSON(channelJSON)
	return nil
}

func (o ChannelWithIsDefault) MarshalJSON() ([]byte, error) {
	var channelJSON channelJSON
	channelJSON.fromChannel(o.Channel)
	channelJSON.IsDefault = &o.IsDefault
	return json.Marshal(&channelJSON)
}

func (o *ChannelWithIsDefault) UnmarshalJSON(data []byte) error {
	var err error
	var channelJSON channelJSON
	err = newStrictDecoder(data).Decode(&channelJSON)
	if err != nil {
		return err
	}
	o.fromChannelJSON(channelJSON)
	if channelJSON.IsDefault != nil {
		o.IsDefault = *channelJSON.IsDefault
	} else {
		o.IsDefault = false
	}
	return nil
}

func (o *channelJSON) fromChannel(cb Channel) {
	o.ChannelId = cb.ChannelId
	o.Name = cb.Name
	o.IsArchived = cb.IsArchived
	o.Description = cb.Description
	o.DateCreated = cb.DateCreated.Unix()
	o.CreatorId = cb.CreatorId
	o.InviteOnly = cb.InviteOnly
	o.RenderedDescription = cb.RenderedDescription
	o.IsWebPublic = cb.IsWebPublic
	o.ChannelPostPolicy = cb.ChannelPostPolicy
	o.MessageRetentionDays = cb.MessageRetentionDays
	o.HistoryPublicToSubscribers = cb.HistoryPublicToSubscribers
	o.TopicsPolicy = cb.TopicsPolicy
	o.FirstMessageId = cb.FirstMessageId
	o.FolderId = cb.FolderId
	o.IsRecentlyActive = cb.IsRecentlyActive
	o.IsAnnouncementOnly = cb.IsAnnouncementOnly
	o.CanAddSubscribersGroup = cb.CanAddSubscribersGroup
	o.CanRemoveSubscribersGroup = cb.CanRemoveSubscribersGroup
	o.CanAdministerChannelGroup = cb.CanAdministerChannelGroup
	o.CanDeleteAnyMessageGroup = cb.CanDeleteAnyMessageGroup
	o.CanDeleteOwnMessageGroup = cb.CanDeleteOwnMessageGroup
	o.CanMoveMessagesOutOfChannelGroup = cb.CanMoveMessagesOutOfChannelGroup
	o.CanMoveMessagesWithinChannelGroup = cb.CanMoveMessagesWithinChannelGroup
	o.CanSendMessageGroup = cb.CanSendMessageGroup
	o.CanSubscribeGroup = cb.CanSubscribeGroup
	o.CanResolveTopicsGroup = cb.CanResolveTopicsGroup
	o.SubscriberCount = cb.SubscriberCount
	o.ChannelWeeklyTraffic = cb.ChannelWeeklyTraffic
}

func (o *Channel) fromChannelJSON(cb channelJSON) {
	o.ChannelId = cb.ChannelId
	o.Name = cb.Name
	o.IsArchived = cb.IsArchived
	o.Description = cb.Description
	o.DateCreated = time.Unix(cb.DateCreated, 0)
	o.CreatorId = cb.CreatorId
	o.InviteOnly = cb.InviteOnly
	o.RenderedDescription = cb.RenderedDescription
	o.IsWebPublic = cb.IsWebPublic
	o.ChannelPostPolicy = cb.ChannelPostPolicy
	o.MessageRetentionDays = cb.MessageRetentionDays
	o.HistoryPublicToSubscribers = cb.HistoryPublicToSubscribers
	o.TopicsPolicy = cb.TopicsPolicy
	o.FirstMessageId = cb.FirstMessageId
	o.FolderId = cb.FolderId
	o.IsRecentlyActive = cb.IsRecentlyActive
	o.IsAnnouncementOnly = cb.IsAnnouncementOnly
	o.CanAddSubscribersGroup = cb.CanAddSubscribersGroup
	o.CanRemoveSubscribersGroup = cb.CanRemoveSubscribersGroup
	o.CanAdministerChannelGroup = cb.CanAdministerChannelGroup
	o.CanDeleteAnyMessageGroup = cb.CanDeleteAnyMessageGroup
	o.CanDeleteOwnMessageGroup = cb.CanDeleteOwnMessageGroup
	o.CanMoveMessagesOutOfChannelGroup = cb.CanMoveMessagesOutOfChannelGroup
	o.CanMoveMessagesWithinChannelGroup = cb.CanMoveMessagesWithinChannelGroup
	o.CanSendMessageGroup = cb.CanSendMessageGroup
	o.CanSubscribeGroup = cb.CanSubscribeGroup
	o.CanResolveTopicsGroup = cb.CanResolveTopicsGroup
	o.SubscriberCount = cb.SubscriberCount
	o.ChannelWeeklyTraffic = cb.ChannelWeeklyTraffic
}

type channelJSON struct {
	ChannelId                         int64             `json:"stream_id"`
	Name                              string            `json:"name"`
	IsArchived                        bool              `json:"is_archived"`
	Description                       string            `json:"description"`
	DateCreated                       int64             `json:"date_created"`
	CreatorId                         *int64            `json:"creator_id"`
	InviteOnly                        bool              `json:"invite_only"`
	RenderedDescription               string            `json:"rendered_description"`
	IsWebPublic                       bool              `json:"is_web_public"`
	ChannelPostPolicy                 int32             `json:"stream_post_policy"`
	MessageRetentionDays              *int              `json:"message_retention_days"`
	HistoryPublicToSubscribers        bool              `json:"history_public_to_subscribers"`
	TopicsPolicy                      TopicsPolicy      `json:"topics_policy,omitempty"`
	FirstMessageId                    *int64            `json:"first_message_id"`
	FolderId                          *int64            `json:"folder_id"`
	IsRecentlyActive                  bool              `json:"is_recently_active"`
	IsAnnouncementOnly                bool              `json:"is_announcement_only"`
	CanAddSubscribersGroup            GroupSettingValue `json:"can_add_subscribers_group,omitempty"`
	CanRemoveSubscribersGroup         GroupSettingValue `json:"can_remove_subscribers_group"`
	CanAdministerChannelGroup         GroupSettingValue `json:"can_administer_channel_group,omitempty"`
	CanDeleteAnyMessageGroup          GroupSettingValue `json:"can_delete_any_message_group,omitempty"`
	CanDeleteOwnMessageGroup          GroupSettingValue `json:"can_delete_own_message_group,omitempty"`
	CanMoveMessagesOutOfChannelGroup  GroupSettingValue `json:"can_move_messages_out_of_channel_group,omitempty"`
	CanMoveMessagesWithinChannelGroup GroupSettingValue `json:"can_move_messages_within_channel_group,omitempty"`
	CanSendMessageGroup               GroupSettingValue `json:"can_send_message_group,omitempty"`
	CanSubscribeGroup                 GroupSettingValue `json:"can_subscribe_group"`
	CanResolveTopicsGroup             GroupSettingValue `json:"can_resolve_topics_group,omitempty"`
	SubscriberCount                   float32           `json:"subscriber_count"`
	ChannelWeeklyTraffic              *int              `json:"stream_weekly_traffic"`
	IsDefault                         *bool             `json:"is_default,omitempty"`

	// members for subscription
	Subscribers            []int64 `json:"subscribers,omitempty"`
	PartialSubscribers     []int64 `json:"partial_subscribers,omitempty"`
	DesktopNotifications   *bool   `json:"desktop_notifications,omitempty"`
	EmailNotifications     *bool   `json:"email_notifications,omitempty"`
	WildcardMentionsNotify *bool   `json:"wildcard_mentions_notify,omitempty"`
	PushNotifications      *bool   `json:"push_notifications,omitempty"`
	AudibleNotifications   *bool   `json:"audible_notifications,omitempty"`
	PinToTop               bool    `json:"pin_to_top,omitempty"`
	IsMuted                bool    `json:"is_muted,omitempty"`
	InHomeView             bool    `json:"in_home_view,omitempty"`
	Color                  string  `json:"color,omitempty"`
}
