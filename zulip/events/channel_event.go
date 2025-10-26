package events

import (
	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/internal/utils"
)

// ChannelCreateEvent Event sent when a new channel is created to users who can see the new channel exists (for private channels, only subscribers and organization administrators will receive this event).  This event is also sent when a user gains access to a channel they previously [could not access], such as when their [role] changes, a private channel is made public, or a guest user is subscribed to a public (or private) channel.  This event is also sent when a channel is unarchived but only to clients that did not declare the `archived_channels` [client capability].  Note that organization administrators who are not subscribed will not be able to see content on the channel; just that it exists.
//
// **Changes**: Prior to Zulip 11.0 (feature level 378), this event was sent to all the users who could see the channel when it was unarchived.  Prior to Zulip 8.0 (feature level 220), this event was incorrectly not sent to guest users a web-public channel was created.  Prior to Zulip 8.0 (feature level 205), this event was not sent when a user gained access to a channel due to their role changing.  Prior to Zulip 8.0 (feature level 192), this event was not sent when guest users gained access to a public channel by being subscribed.  Prior to Zulip 6.0 (feature level 134), this event was not sent when a private channel was made public.
//
// [could not access]: https://zulip.com/help/channel-permissions
// [role]: https://zulip.com/help/user-roles
type ChannelCreateEvent struct {
	event
	// Array of objects, each containing details about the newly added channel(s).
	Channels []zulip.Channel `json:"streams,omitempty"`
}

// ChannelDeleteEvent Event sent when a user loses access to a channel they previously [could access] because they are unsubscribed from a private channel or their [role] has changed.  This event is also sent when a channel is archived but only to clients that did not declare the `archived_channels` [client capability].
//
// **Changes**: Prior to Zulip 11.0 (feature level 378), this event was sent to all the users who could see the channel when it was archived.  Prior to Zulip 8.0 (feature level 205), this event was not sent when a user lost access to a channel due to their role changing.
//
// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
// [could access]: https://zulip.com/help/channel-permissions
// [role]: https://zulip.com/help/user-roles
type ChannelDeleteEvent struct {
	event

	// Array of objects, each containing Id of the channel that was deleted.
	//
	// **Changes**:**Deprecated** in Zulip 10.0 (feature level 343) and will be removed in a future release. Previously, these objects additionally contained all the standard fields for a channel object.
	// Deprecated
	Channels []interface{} `json:"streams,omitempty"`

	// Array containing the Ids of the channels that were deleted.
	//
	// **Changes**: New in Zulip 10.0 (feature level 343). Previously, these Ids were available only via the legacy `streams` array.
	ChannelIds []int64 `json:"stream_ids,omitempty"`
}

// ChannelUpdateEvent Event sent to all users who can see that a channel exists when a property of that channel changes. See [GET /streams] response for details on the various properties of a channel.  This event is also sent when archiving or unarchiving a channel to all the users who can see that channel exists but only to the clients that declared the `archived_channels` [client capability].
//
// **Changes**: Prior to Zulip 11.0 (feature level 378), this event was never sent when archiving or unarchiving a channel.  Before Zulip 9.0 (feature level 256), this event was never sent when the `first_message_id` property of a channel was updated because the oldest message that had been sent to it changed.
//
// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
// [GET /streams]: https://zulip.com/api/get-streams#response
type ChannelUpdateEvent struct {
	event
	// The Id of the channel whose details have changed.
	ChannelId int64 `json:"stream_id,omitempty"`
	// The name of the channel whose details have changed.
	Name string `json:"name,omitempty"`
	// The property of the channel which has changed. See [GET /streams] response for details on the various properties of a channel.  Clients should handle an "unknown" property received here without crashing, since that can happen when connecting to a server running a newer version of Zulip with new features.
	//
	// [GET /streams]: https://zulip.com/api/get-streams#response
	Property string `json:"property,omitempty"`
	// ChannelEventUpdateValue - The new value of the changed property.
	//
	// **Changes**: Starting with Zulip 11.0 (feature level 389), this value can be `null` when a channel is removed from the folder.  Starting with Zulip 10.0 (feature level 320), this field can be an object for `can_remove_subscribers_group` property, which is a [group-setting value], when the setting is set to a combination of users and groups.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	Value *ChannelEventUpdateValue `json:"value,omitempty"`
	// Note: Only present if the changed property was `description`.  The short description of the channel rendered as HTML, intended to be used when displaying the channel description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	RenderedDescription *string `json:"rendered_description,omitempty"`
	// Note: Only present if the changed property was `invite_only`.  Whether the history of the channel is public to its subscribers.  Currently always true for public channels (i.e. `"invite_only": false` implies `"history_public_to_subscribers": true`), but clients should not make that assumption, as we may change that behavior in the future.
	HistoryPublicToSubscribers *bool `json:"history_public_to_subscribers,omitempty"`
	// Note: Only present if the changed property was `invite_only`.  Whether the channel's history is now readable by web-public spectators.
	//
	// **Changes**: New in Zulip 5.0 (feature level 71).
	IsWebPublic *bool `json:"is_web_public,omitempty"`
}

type ChannelEventUpdateValue struct {
	GroupSettingValue *zulip.GroupSettingValue
	Bool              *bool
	Int64             *int64
	String            *string
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChannelEventUpdateValue) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalUnionType(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChannelEventUpdateValue) MarshalJSON() ([]byte, error) {
	return utils.MarshalUnionType(src)
}
