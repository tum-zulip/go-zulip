package events_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	realtimeevents "github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/events"
)

func Test_Unmarshal_HeartbeatEvent(t *testing.T) {
	data := []byte(`{"result": "success", "msg": "", "events": [{"type":"heartbeat","id":0}]}`)

	var resp realtimeevents.GetEventsResponse
	err := json.Unmarshal(data, &resp)

	require.NoError(t, err)
	require.Len(t, resp.Events, 1)

	event := resp.Events[0]

	_, ok := event.(events.HeartbeatEvent)
	require.True(t, ok, "expected HeartbeatEvent, got %T", event)
	require.Equal(t, int64(0), event.GetID())
	require.Equal(t, events.EventTypeHeartbeat, event.GetType())
}

func Test_Unmarshal_ChannelUpdateEvent_GroupSettingValueID(t *testing.T) {
	data := []byte(`{
		"result": "success",
		"msg": "",
		"events": [{
			"type": "stream",
			"op": "update",
			"id": 1,
			"stream_id": 2,
			"property": "can_subscribe_group",
			"value": 3
		}]
	}`)

	var resp realtimeevents.GetEventsResponse
	err := json.Unmarshal(data, &resp)

	require.NoError(t, err)
	require.Len(t, resp.Events, 1)

	event, ok := resp.Events[0].(events.ChannelUpdateEvent)
	require.True(t, ok, "expected ChannelUpdateEvent, got %T", resp.Events[0])
	require.NotNil(t, event.Value)
	require.NotNil(t, event.Value.GroupSettingValue)
	assert.Equal(t, int64(3), *event.Value.GroupSettingValue.GroupID)
}

func Test_Unmarshal_ChannelUpdateEvent_Int64Value(t *testing.T) {
	data := []byte(`{
		"result": "success",
		"msg": "",
		"events": [{
			"type": "stream",
			"op": "update",
			"id": 1,
			"stream_id": 2,
			"property": "folder_id",
			"value": 3
		}]
	}`)

	var resp realtimeevents.GetEventsResponse
	err := json.Unmarshal(data, &resp)

	require.NoError(t, err)
	require.Len(t, resp.Events, 1)

	event, ok := resp.Events[0].(events.ChannelUpdateEvent)
	require.True(t, ok, "expected ChannelUpdateEvent, got %T", resp.Events[0])
	require.NotNil(t, event.Value)
	require.NotNil(t, event.Value.Int64)
	assert.Equal(t, int64(3), *event.Value.Int64)
}
