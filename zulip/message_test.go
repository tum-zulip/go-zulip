package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func TestMessageMarshalJSON_EncodesUnixSeconds(t *testing.T) {

	timestamp := time.Unix(1700000000, 0).UTC()
	lastEdit := time.Unix(1700000100, 0).UTC()
	lastMoved := time.Unix(1700000200, 0).UTC()

	message := z.Message{
		ID:                 42,
		Client:             "web",
		Content:            "hello",
		ContentType:        "text/html",
		LastEditTimestamp:  lastEdit,
		LastMovedTimestamp: lastMoved,
		Timestamp:          timestamp,
		Flags:              []string{"read"},
	}

	data, err := json.Marshal(message)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	ts, ok := payload["timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), ts)
	assert.Equal(t, float64(timestamp.Unix()), ts)

	lastEditVal, ok := payload["last_edit_timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), lastEditVal)
	assert.Equal(t, float64(lastEdit.Unix()), lastEditVal)

	lastMovedVal, ok := payload["last_moved_timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), lastMovedVal)
	assert.Equal(t, float64(lastMoved.Unix()), lastMovedVal)
}

func TestMessageUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {

	raw := []byte(`{
        "id": 42,
        "client": "web",
        "content": "hello",
        "content_type": "text/html",
        "display_recipient": [{"id": 1, "email": "alice@example.com", "full_name": "Alice"}],
        "last_edit_timestamp": 1700000100,
        "last_moved_timestamp": 1700000200,
        "timestamp": 1700000000,
        "flags": ["read"]
    }`)

	var message z.Message
	require.NoError(t, json.Unmarshal(raw, &message))

	assert.Equal(t, int64(42), message.ID)
	assert.Equal(t, "web", message.Client)
	assert.Equal(t, "hello", message.Content)
	assert.Equal(t, "text/html", message.ContentType)
	assert.Equal(t, int64(1700000000), message.Timestamp.Unix())
	assert.Equal(t, time.UTC, message.Timestamp.Location())
	assert.Equal(t, int64(1700000100), message.LastEditTimestamp.Unix())
	assert.Equal(t, time.UTC, message.LastEditTimestamp.Location())
	assert.Equal(t, int64(1700000200), message.LastMovedTimestamp.Unix())
	assert.Equal(t, time.UTC, message.LastMovedTimestamp.Location())
	assert.Equal(t, []string{"read"}, message.Flags)
}
