package zulip_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func Test_EventEnvelope_UnmarshalJSON(t *testing.T) {
	t.Run("Decodes HeartbeatEvent", func(t *testing.T) {
		data := []byte(`{"result": "success", "msg": "", "events": [{"type":"heartbeat","id":0}]}`)

		var resp zulip.GetEventsResponse
		err := json.Unmarshal(data, &resp)

		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]

		_, ok := event.(zulip.HeartbeatEvent)
		require.True(t, ok, "expected HeartbeatEvent, got %T", event)
		require.Equal(t, int64(0), event.GetId())
		require.Equal(t, zulip.EventTypeHeartbeat, event.GetType())

	})

}
