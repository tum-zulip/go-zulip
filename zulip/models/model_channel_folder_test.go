package models_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func TestChannelFolderMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	t.Parallel()

	ts := time.Unix(1700000000, 0).UTC()
	creator := int64(42)
	folder := z.ChannelFolder{
		Name:        "general",
		DateCreated: &ts,
		CreatorId:   &creator,
		Order:       3,
		Id:          99,
		IsArchived:  true,
	}

	data, err := json.Marshal(folder)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["date_created"]
	require.True(t, ok)
	require.IsType(t, float64(0), value)
	assert.Equal(t, float64(ts.Unix()), value)
	assert.Equal(t, "general", payload["name"])
}

func TestChannelFolderUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {
	t.Parallel()

	raw := []byte(`{"name":"general","date_created":1700000000}`)

	var folder z.ChannelFolder
	require.NoError(t, json.Unmarshal(raw, &folder))

	if assert.NotNil(t, folder.DateCreated) {
		assert.Equal(t, int64(1700000000), folder.DateCreated.Unix())
		assert.Equal(t, time.UTC, folder.DateCreated.Location())
	}
}
