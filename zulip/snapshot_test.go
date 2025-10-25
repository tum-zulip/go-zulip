package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func TestSnapshotMarshalJSON_EncodesUnixSeconds(t *testing.T) {

	ts := time.Unix(1700000000, 250*int64(time.Millisecond)).UTC()
	channel := int64(5)
	prevChannel := int64(4)
	prevTopic := "old"
	topic := "new"
	diff := "<ins>new</ins>"
	snapshot := z.Snapshot{
		Topic:           &topic,
		Timestamp:       ts,
		PrevTopic:       &prevTopic,
		Channel:         &channel,
		PrevChannel:     &prevChannel,
		Content:         "content",
		RenderedContent: "<p>content</p>",
		ContentHtmlDiff: &diff,
	}

	data, err := json.Marshal(snapshot)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), value)
	assert.Equal(t, float64(ts.Unix()), value)
}

func TestSnapshotUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {

	raw := []byte(`{"topic":"new","timestamp":1700000000}`)

	var snapshot z.Snapshot
	require.NoError(t, json.Unmarshal(raw, &snapshot))

	require.NotNil(t, snapshot.Topic)
	assert.Equal(t, "new", *snapshot.Topic)
	assert.Equal(t, int64(1700000000), snapshot.Timestamp.Unix())
}
