package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func TestChannelMarshalJSON_EncodesUnixSeconds(t *testing.T) {

	creatorId := int64(7)
	date := time.Unix(1700000000, 500*int64(time.Millisecond)).UTC()
	channel := z.Channel{
		ChannelId:            123,
		Name:                 "general",
		DateCreated:          date,
		CreatorId:            &creatorId,
		ChannelPostPolicy:    2,
		ChannelWeeklyTraffic: ptrTo[int](10),
	}

	data, err := json.Marshal(channel)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["date_created"]
	require.True(t, ok)
	require.IsType(t, float64(0), value)
	assert.Equal(t, float64(date.Unix()), value)
}

func TestChannelUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {

	raw := []byte(`{"stream_id":123,"name":"general","date_created":1700000000}`)

	var channel z.Channel
	require.NoError(t, json.Unmarshal(raw, &channel))

	assert.Equal(t, int64(123), channel.ChannelId)
	assert.Equal(t, "general", channel.Name)
	assert.Equal(t, int64(1700000000), channel.DateCreated.Unix())
}

func ptrTo[T any](v T) *T {
	return &v
}
