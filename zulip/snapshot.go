package zulip

import (
	"encoding/json"
	"time"

	"github.com/tum-zulip/go-zulip/zulip/internal/utils"
)

// Snapshot struct for Snapshot
type Snapshot struct {
	// Only present if message's topic was edited.  The topic of the message immediately after this edit event.
	//
	// **Changes**: New in Zulip 5.0 (feature level 118).
	Topic *string `json:"topic,omitempty"`
	// Only present if message's topic was edited.  The topic of the message immediately prior to this edit event.
	PrevTopic *string `json:"prev_topic,omitempty"`
	// Only present if message's channel was edited.  The Id of the channel containing the message immediately after this edit event.
	//
	// **Changes**: New in Zulip 5.0 (feature level 118).
	Channel *int64 `json:"stream,omitempty"`
	// Only present if message's channel was edited.  The Id of the channel containing the message immediately prior to this edit event.
	//
	// **Changes**: New in Zulip 3.0 (feature level 1).
	PrevChannel *int64 `json:"prev_stream,omitempty"`
	// The raw [Zulip-flavored Markdown] content of the message immediately after this edit event.
	//
	// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
	Content string `json:"content,omitempty"`
	// The rendered HTML representation of `content`.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	RenderedContent string `json:"rendered_content,omitempty"`
	// Only present if message's content was edited.  The raw [Zulip-flavored Markdown] content of the message immediately prior to this edit event.
	//
	// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
	PrevContent *string `json:"prev_content,omitempty"`
	// Only present if message's content was edited.  The rendered HTML representation of `prev_content`.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	PrevRenderedContent *string `json:"prev_rendered_content,omitempty"`
	// The Id of the user that made the edit.  Will be `null` only for edit history events predating March 2017.  Clients can display edit history events where this is `null` as modified by either the sender (for content edits) or an unknown user (for topic edits).
	UserID *int64 `json:"user_id,omitempty"`
	// Only present if message's content was edited.  An HTML diff between this version of the message and the previous one.
	ContentHtmlDiff *string `json:"content_html_diff,omitempty"`
	// The UNIX timestamp for this edit.
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (o Snapshot) MarshalJSON() ([]byte, error) {
	snapshotJSON := snapshotJSON{
		Topic:               o.Topic,
		PrevTopic:           o.PrevTopic,
		Channel:             o.Channel,
		PrevChannel:         o.PrevChannel,
		Content:             o.Content,
		RenderedContent:     o.RenderedContent,
		PrevContent:         o.PrevContent,
		PrevRenderedContent: o.PrevRenderedContent,
		UserID:              o.UserID,
		ContentHtmlDiff:     o.ContentHtmlDiff,
		Timestamp:           o.Timestamp.Unix(),
	}
	return json.Marshal(&snapshotJSON)
}

func (o *Snapshot) UnmarshalJSON(data []byte) error {
	var j snapshotJSON
	dec := utils.NewStrictDecoder(data)
	if err := dec.Decode(&j); err != nil {
		return err
	}

	o.Topic = j.Topic
	o.PrevTopic = j.PrevTopic
	o.Channel = j.Channel
	o.PrevChannel = j.PrevChannel
	o.Content = j.Content
	o.RenderedContent = j.RenderedContent
	o.PrevContent = j.PrevContent
	o.PrevRenderedContent = j.PrevRenderedContent
	o.UserID = j.UserID
	o.ContentHtmlDiff = j.ContentHtmlDiff
	o.Timestamp = time.Unix(j.Timestamp, int64(0))
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
	UserID              *int64  `json:"user_id,omitempty"`
	ContentHtmlDiff     *string `json:"content_html_diff,omitempty"`
	Timestamp           int64   `json:"timestamp,omitempty"`
}
