package events

// DefaultChannelsEvent Event sent to all users in a Zulip organization when the default channels in the organization are changed by an organization administrator.
type DefaultChannelsEvent struct {
	event
	// An array of IDs of all the [default channels] in the organization.
	//
	// **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single default stream for the Zulip organization.
	//
	// [default channels]: https://zulip.com/help/set-default-streams-for-new-users
	DefaultChannels []int64 `json:"default_streams,omitempty"`
}
