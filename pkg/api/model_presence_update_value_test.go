package api_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tum-zulip/go-zulip/pkg/api"
)

func TestPresenceUpdateValue_MarshalUnmarshal_Modern(t *testing.T) {
	// create PresenceUpdateValue with modern presence format
	p := api.PresenceUpdateValue{
		ModernPresenceFormat: &api.ModernPresenceFormat{
			ActiveTimestamp: time.Unix(1760534744, 0).UTC(),
			IdleTimestamp:   time.Unix(1760534744, 0).UTC(),
		},
	}

	b, err := json.Marshal(p)
	require.NoError(t, err)

	// exact JSON string expected from the envelope marshal
	assert.Equal(t, `{"active_timestamp":1760534744,"idle_timestamp":1760534744}`, string(b))

	// round-trip through the envelope
	var got api.PresenceUpdateValue
	require.NoError(t, json.Unmarshal(b, &got))
	assert.Equal(t, p, got)
}

func TestPresenceUpdateValue_MarshalUnmarshal_Legacy(t *testing.T) {
	// expected legacy map
	expected := api.PresenceUpdateValue{
		LegacyPresenceMap: map[string]api.LegacyPresenceFormat{
			"aggregated": {
				Client:    "website",
				Status:    "active",
				Timestamp: 1760534744,
			},
			"website": {
				Client:    "website",
				Status:    "active",
				Timestamp: 1760534744,
				Pushable:  false,
			},
		},
	}

	// build envelope from map and marshal
	b, err := json.Marshal(expected)
	require.NoError(t, err)

	assert.Equal(t, `{"aggregated":{"client":"website","status":"active","timestamp":1760534744},"website":{"client":"website","status":"active","timestamp":1760534744}}`, string(b))

	// unmarshal produced bytes back into a map and compare
	var got api.PresenceUpdateValue
	require.NoError(t, json.Unmarshal(b, &got))
	assert.Equal(t, expected, got)
}
