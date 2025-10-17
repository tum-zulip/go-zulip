package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func TestSnapshotMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	t.Parallel()

	ts := time.Unix(1700000000, 250*int64(time.Millisecond)).UTC()
	channel := int64(5)
	prevChannel := int64(4)
	prevTopic := "old"
	snapshot := zulip.Snapshot{
		Topic:           "new",
		Content:         "content",
		RenderedContent: "<p>content</p>",
		ContentHtmlDiff: "<ins>new</ins>",
	}
	snapshot.Timestamp = ts
	snapshot.PrevTopic = &prevTopic
	snapshot.Channel = &channel
	snapshot.PrevChannel = &prevChannel

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
	t.Parallel()

	raw := []byte(`{"topic":"new","timestamp":1700000000}`)

	var snapshot zulip.Snapshot
	require.NoError(t, json.Unmarshal(raw, &snapshot))

	assert.Equal(t, "new", snapshot.Topic)
	assert.Equal(t, int64(1700000000), snapshot.Timestamp.Unix())
	assert.Equal(t, time.UTC, snapshot.Timestamp.Location())
}
