package zulip

// Whether [named topics] and the empty topic (i.e., ["general chat" topic] are enabled in this channel.
//   - TopicsPolicyInherit
//   - TopicsPolicyAllowEmptyTopic
//   - TopicsPolicyDisableEmptyTopic
//   - TopicsPolicyEmptyTopicOnly
//
// **Changes**: In Zulip 11.0 (feature level 404), the `"empty_topic_only"` option was added.  New in Zulip 11.0 (feature level 392).
//
// [named topics]: https://zulip.com/help/introduction-to-topics
// ["general chat" topic]: https://zulip.com/help/general-chat-topic
type TopicsPolicy string

const (
	// Messages can be sent to named topics in this channel, and the [organization-level `realm_topics_policy`] is used for whether messages can be sent to the empty topic in this channel.
	//
	// [organization-level `realm_topics_policy`]: https://zulip.com/help/require-topics#set-the-default-general-chat-topic-configuration
	TopicsPolicyInherit TopicsPolicy = "inherit"
	// Messages can be sent to both named topics and the empty topic in this channel.
	TopicsPolicyAllowEmptyTopic TopicsPolicy = "allow_empty_topic"
	// Messages can be sent to named topics in this channel, but the empty topic is disabled.
	TopicsPolicyDisableEmptyTopic TopicsPolicy = "disable_empty_topic"
	// Messages can be sent to the empty topic in this channel, but named topics are disabled.
	//
	// See ["general chat" channels]. The `"empty_topic_only"` policy can only be set if all existing messages in the channel are already in the empty topic.  When creating a new channel, if the `topics_policy` is not specified, the `"inherit"` option will be set.
	//
	// ["general chat" channels]: https://zulip.com/help/general-chat-channels
	TopicsPolicyEmptyTopicOnly TopicsPolicy = "empty_topic_only"
)

var allowedTopicsPolicyEnumValues = []TopicsPolicy{
	TopicsPolicyAllowEmptyTopic,
	TopicsPolicyDisableEmptyTopic,
	TopicsPolicyEmptyTopicOnly,
	TopicsPolicyInherit,
}

func NewTopicsPolicyFromValue(v string) (*TopicsPolicy, error) {
	ev := TopicsPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    allowedTopicsPolicyEnumValues,
			VarName: "TopicsPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TopicsPolicy) IsValid() bool {
	for _, existing := range allowedTopicsPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
