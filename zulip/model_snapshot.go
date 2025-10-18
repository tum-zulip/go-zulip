package zulip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Snapshot struct for Snapshot
type SnapshotCommon struct {
	// Only present if message's topic was edited.  The topic of the message immediately prior to this edit event.
	PrevTopic *string `json:"prev_topic,omitempty"`
	// Only present if message's channel was edited.  The Id of the channel containing the message immediately after this edit event.  **Changes**: New in Zulip 5.0 (feature level 118).
	Channel *int64 `json:"stream,omitempty"`
	// Only present if message's channel was edited.  The Id of the channel containing the message immediately prior to this edit event.  **Changes**: New in Zulip 3.0 (feature level 1).
	PrevChannel *int64 `json:"prev_stream,omitempty"`
	// The raw [Zulip-flavored Markdown](https://zulip.com/help/format-your-message-using-markdown) content of the message immediately after this edit event.
	Content string `json:"content,omitempty"`
	// The rendered HTML representation of `content`.  See [Markdown message formatting](https://zulip.com/api/message-formatting) for details on Zulip's HTML format.
	RenderedContent string `json:"rendered_content,omitempty"`
	// Only present if message's content was edited.  The raw [Zulip-flavored Markdown](https://zulip.com/help/format-your-message-using-markdown) content of the message immediately prior to this edit event.
	PrevContent *string `json:"prev_content,omitempty"`
	// Only present if message's content was edited.  The rendered HTML representation of `prev_content`.  See [Markdown message formatting](https://zulip.com/api/message-formatting) for details on Zulip's HTML format.
	PrevRenderedContent *string `json:"prev_rendered_content,omitempty"`
	// The Id of the user that made the edit.  Will be `null` only for edit history events predating March 2017.  Clients can display edit history events where this is `null` as modified by either the sender (for content edits) or an unknown user (for topic edits).
	UserId *int64 `json:"user_id,omitempty"`
	// Only present if message's content was edited.  An HTML diff between this version of the message and the previous one.
	ContentHtmlDiff string `json:"content_html_diff,omitempty"`
	// The UNIX timestamp for this edit.
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// EditHistory struct for EditHistory
type EditHistory struct {
	SnapshotCommon
	// Only present if message's topic was edited.  The topic of the message immediately after this edit event.  **Changes**: New in Zulip 5.0 (feature level 118).
	Topic *string `json:"topic,omitempty"`
}

func (o EditHistory) MarshalJSON() ([]byte, error) {
	var snapshotJSON snapshotJSON
	snapshotJSON.fromSnapshotCommon(o.SnapshotCommon)
	snapshotJSON.Topic = o.Topic
	return json.Marshal(&snapshotJSON)
}

func (o *EditHistory) UnmarshalJSON(data []byte) error {
	var snapshotJSON snapshotJSON
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&snapshotJSON); err != nil {
		return err
	}
	o.SnapshotCommon.fromSnapshotJSON(snapshotJSON)
	o.Topic = snapshotJSON.Topic
	return nil
}

// Snapshot struct for Snapshot
type Snapshot struct {
	SnapshotCommon

	// The topic of the message immediately after this edit event.
	Topic string `json:"topic,omitempty"`

	// The raw [Zulip-flavored Markdown](https://zulip.com/help/format-your-message-using-markdown) content of the message immediately after this edit event.
	Content string `json:"content,omitempty"`
	// The rendered HTML representation of `content`.  See [Markdown message formatting](https://zulip.com/api/message-formatting) for details on Zulip's HTML format.
	RenderedContent string `json:"rendered_content,omitempty"`

	// Only present if message's content was edited.  An HTML diff between this version of the message and the previous one.
	ContentHtmlDiff string `json:"content_html_diff,omitempty"`
}

func (o Snapshot) MarshalJSON() ([]byte, error) {
	var snapshotJSON snapshotJSON
	snapshotJSON.fromSnapshotCommon(o.SnapshotCommon)
	snapshotJSON.Topic = &o.Topic
	snapshotJSON.Content = o.Content
	snapshotJSON.RenderedContent = o.RenderedContent
	snapshotJSON.ContentHtmlDiff = o.ContentHtmlDiff
	return json.Marshal(&snapshotJSON)
}

func (o *Snapshot) UnmarshalJSON(data []byte) error {
	var snapshotJSON snapshotJSON
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&snapshotJSON); err != nil {
		return err
	}
	o.fromSnapshotJSON(snapshotJSON)
	if snapshotJSON.Topic == nil {
		return fmt.Errorf("required field topic missing")
	}
	o.Topic = *snapshotJSON.Topic
	o.Content = snapshotJSON.Content
	o.RenderedContent = snapshotJSON.RenderedContent
	o.ContentHtmlDiff = snapshotJSON.ContentHtmlDiff
	return nil
}

type snapshotJSON struct {
	Topic               *string `json:"topic,omitempty"`
	PrevTopic           *string `json:"prev_topic,omitempty"`
	Channel             *int64  `json:"stream,omitempty"`
	PrevChannel         *int64  `json:"prev_stream,omitempty"`
	Content             string  `json:"content,omitempty"`
	RenderedContent     string  `json:"rendered_content,omitempty"`
	PrevContent         *string `json:"prev_content,omitempty"`
	PrevRenderedContent *string `json:"prev_rendered_content,omitempty"`
	UserId              *int64  `json:"user_id,omitempty"`
	ContentHtmlDiff     string  `json:"content_html_diff,omitempty"`
	Timestamp           int64   `json:"timestamp,omitempty"`
}

func (o *snapshotJSON) fromSnapshotCommon(base SnapshotCommon) {
	o.PrevTopic = base.PrevTopic
	o.Channel = base.Channel
	o.PrevChannel = base.PrevChannel
	o.Content = base.Content
	o.RenderedContent = base.RenderedContent
	o.PrevContent = base.PrevContent
	o.PrevRenderedContent = base.PrevRenderedContent
	o.UserId = base.UserId
	o.ContentHtmlDiff = base.ContentHtmlDiff
	o.Timestamp = base.Timestamp.Unix()
}

func (o *SnapshotCommon) fromSnapshotJSON(data snapshotJSON) {
	o.PrevTopic = data.PrevTopic
	o.Channel = data.Channel
	o.PrevChannel = data.PrevChannel
	o.Content = data.Content
	o.RenderedContent = data.RenderedContent
	o.PrevContent = data.PrevContent
	o.PrevRenderedContent = data.PrevRenderedContent
	o.UserId = data.UserId
	o.ContentHtmlDiff = data.ContentHtmlDiff
	if data.Timestamp != 0 {
		o.Timestamp = time.Unix(data.Timestamp, 0).UTC()
	}
}
