package events

// MutedTopicsEvent Event sent to a user's clients when that user's set of configured muted topics have changed.
type MutedTopicsEvent struct {
	event
	// Array of tuples, where each tuple describes a muted topic. The first element of the tuple is the channel name in which the topic has to be muted, the second element is the topic name to be muted and the third element is an integer UNIX timestamp representing when the topic was muted.
	//
	// **Changes**: Deprecated in Zulip 6.0 (feature level 134). Starting with this version, clients that explicitly requested the replacement `user_topic` event type when registering their event queue will not receive this legacy event type.  Before Zulip 3.0 (feature level 1), the `muted_topics` array objects were 2-item tuples and did not include the timestamp information for when the topic was muted.
	// Deprecated
	MutedTopics [][]interface{} `json:"muted_topics,omitempty"`
}
