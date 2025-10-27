package events_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/events"
)

func Test_Unmarshal_HeartbeatEvent(t *testing.T) {

	data := []byte(`{"result": "success", "msg": "", "events": [{"type":"heartbeat","id":0}]}`)

	var resp real_time_events.GetEventsResponse
	err := json.Unmarshal(data, &resp)

	require.NoError(t, err)
	require.Len(t, resp.Events, 1)

	event := resp.Events[0]

	_, ok := event.(events.HeartbeatEvent)
	require.True(t, ok, "expected HeartbeatEvent, got %T", event)
	require.Equal(t, int64(0), event.GetID())
	require.Equal(t, events.EventTypeHeartbeat, event.GetType())

}
