package zulip_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func Test_EventEnvelope_UnmarshalJSON(t *testing.T) {
	t.Run("Decodes HeartbeatEvent", func(t *testing.T) {
		data := []byte(`{"type":"heartbeat","id":0}`)

		var envelope zulip.EventEnvelope
		err := json.Unmarshal(data, &envelope)

		require.NoError(t, err)
		_, ok := envelope.Event.(zulip.HeartbeatEvent)
		require.True(t, ok, "expected HeartbeatEvent, got %T", envelope.Event)
		require.Equal(t, int64(0), envelope.Event.GetId())
		require.Equal(t, zulip.EventTypeHeartbeat, envelope.Event.GetType())

	})

}
