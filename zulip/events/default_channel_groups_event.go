package events

// DefaultChannelGroupsEvent Event sent to all users in a Zulip organization when an organization administrator changes the organization's configured default channel groups.  Default channel groups are an **experimental** feature that is not yet stabilized.
type DefaultChannelGroupsEvent struct {
	event
	// An array of dictionaries where each dictionary contains details about a single default channel group.
	DefaultChannelGroups []DefaultChannelGroup `json:"default_stream_groups,omitempty"`
}

// DefaultChannelGroup Dictionary containing details of a default channel group.
type DefaultChannelGroup struct {
	// Name of the default channel group.
	Name string `json:"name,omitempty"`
	// Description of the default channel group.
	Description string `json:"description,omitempty"`
	// The ID of the default channel group.
	ID int64 `json:"id,omitempty"`
	// An array of IDs of all the channels in the default stream group.
	//
	// **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single stream in the default stream group.
	Channels []int64 `json:"streams,omitempty"`
}
