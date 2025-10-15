package api_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tum-zulip/go-zulip/pkg/api"
)

func TestModernPresenceFormatMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	t.Parallel()

	active := time.Unix(1700000000, 500*int64(time.Millisecond))
	idle := time.Unix(1700000100, 250*int64(time.Millisecond))
	presence := api.ModernPresenceFormat{
		ActiveTimestamp: active,
		IdleTimestamp:   idle,
	}

	data, err := json.Marshal(presence)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	activeVal, ok := payload["active_timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), activeVal)
	assert.Equal(t, float64(active.Unix()), activeVal)

	idleVal, ok := payload["idle_timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), idleVal)
	assert.Equal(t, float64(idle.Unix()), idleVal)
}

func TestModernPresenceFormatUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {
	t.Parallel()

	raw := []byte(`{"active_timestamp":1700000000,"idle_timestamp":1700000100}`)

	var presence api.ModernPresenceFormat
	require.NoError(t, json.Unmarshal(raw, &presence))

	assert.Equal(t, int64(1700000000), presence.ActiveTimestamp.Unix())
	assert.Equal(t, time.UTC, presence.ActiveTimestamp.Location())
	assert.Equal(t, int64(1700000100), presence.IdleTimestamp.Unix())
	assert.Equal(t, time.UTC, presence.IdleTimestamp.Location())
}
