package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func TestDraftMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	t.Parallel()

	ts := time.Unix(1700000000, 123000000).UTC()
	id := int64(42)
	draft := zulip.Draft{
		Id:        &id,
		Type:      "stream",
		To:        []int64{1, 2},
		Topic:     "topic",
		Content:   "hello",
		Timestamp: ts,
	}

	data, err := json.Marshal(draft)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), value)
	assert.Equal(t, float64(ts.Unix()), value)
}

func TestDraftUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {
	t.Parallel()

	raw := []byte(`{"id":42,"type":"channel","to":[1,2],"topic":"topic","content":"hello","timestamp":1700000000}`)

	var draft zulip.Draft
	require.NoError(t, json.Unmarshal(raw, &draft))

	assert.Equal(t, int64(42), *draft.Id)
	assert.Equal(t, zulip.RecipientTypeChannel, draft.Type)
	assert.Equal(t, []int64{1, 2}, draft.To)
	assert.Equal(t, "topic", draft.Topic)
	assert.Equal(t, "hello", draft.Content)
	assert.Equal(t, int64(1700000000), draft.Timestamp.Unix())
	assert.Equal(t, time.UTC, draft.Timestamp.Location())
}
