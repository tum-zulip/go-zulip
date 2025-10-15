package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func TestUserPresenceMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	t.Parallel()

	ts := time.Unix(1700000000, 123000000)
	presence := zulip.UserPresence{
		Timestamp: ts,
		Status:    "active",
	}

	data, err := json.Marshal(presence)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), value)
	assert.Equal(t, float64(ts.Unix()), value)
	assert.Equal(t, "active", payload["status"])
}

func TestUserPresenceUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {
	t.Parallel()

	raw := []byte(`{"timestamp":1700000000,"status":"idle"}`)

	var presence zulip.UserPresence
	require.NoError(t, json.Unmarshal(raw, &presence))

	assert.Equal(t, "idle", presence.Status)
	assert.False(t, presence.Timestamp.IsZero())
	assert.Equal(t, int64(1700000000), presence.Timestamp.Unix())
	assert.Equal(t, time.UTC, presence.Timestamp.Location())
}
