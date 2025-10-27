package events

import (
	"encoding/json"
	"time"

	"github.com/tum-zulip/go-zulip/zulip"
)

// UserTopicEvent Event sent to a user's clients when the user mutes/unmutes a topic, or otherwise modifies their personal per-topic configuration.
//
// **Changes**: New in Zulip 6.0 (feature level 134). Previously, clients were notified about changes in muted topic configuration via the `muted_topics` event type.
type UserTopicEvent struct {
	event
	UserTopic
}

// UserTopics Object describing the user's configuration for a given topic.
type UserTopic struct {
	// The Id of the channel to which the topic belongs.
	ChannelID int64 `json:"stream_id,omitempty"`
	// The name of the topic.  For clients that don't support the `empty_topic_name` [client capability], if the actual topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	TopicName string `json:"topic_name,omitempty"`
	// An integer UNIX timestamp representing when the user-topic relationship was last changed.
	LastUpdated time.Time `json:"last_updated,omitempty"`
	// An integer indicating the user's visibility preferences for the topic, such as whether the topic is muted.
	//   - VisibilityPolicyNone
	//   - VisibilityPolicyMuted
	//   - VisibilityPolicyUnmuted
	//   - VisibilityPolicyFollowed
	//
	// **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.  In Zulip 7.0 (feature level 170), added unmuted as a visibility policy option.
	VisibilityPolicy zulip.VisibilityPolicy `json:"visibility_policy,omitempty"`
}

type userTopicJSON struct {
	ChannelID        int64                  `json:"stream_id,omitempty"`
	TopicName        string                 `json:"topic_name,omitempty"`
	LastUpdated      int64                  `json:"last_updated,omitempty"`
	VisibilityPolicy zulip.VisibilityPolicy `json:"visibility_policy,omitempty"`
}

func (u *UserTopic) UnmarshalJSON(data []byte) error {
	var uj userTopicJSON
	if err := json.Unmarshal(data, &uj); err != nil {
		return err
	}

	u.ChannelID = uj.ChannelID
	u.TopicName = uj.TopicName
	u.LastUpdated = time.Unix(uj.LastUpdated, 0)
	u.VisibilityPolicy = uj.VisibilityPolicy

	return nil
}

func (u UserTopic) MarshalJSON() ([]byte, error) {
	uj := userTopicJSON{
		ChannelID:        u.ChannelID,
		TopicName:        u.TopicName,
		LastUpdated:      u.LastUpdated.Unix(),
		VisibilityPolicy: u.VisibilityPolicy,
	}

	return json.Marshal(uj)
}
