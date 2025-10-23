package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func TestScheduledMessageMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	t.Parallel()

	ts := time.Unix(1700000000, 456000000).UTC()
	topic := "project"
	recipients := []int64{1, 2, 3}

	msg := z.ScheduledMessage{
		ScheduledMessageId:         99,
		Type:                       z.RecipientTypePrivate,
		To:                         z.Recipient{Users: recipients},
		Topic:                      &topic,
		Content:                    "Reminder",
		RenderedContent:            "<p>Reminder</p>",
		ScheduledDeliveryTimestamp: ts,
		Failed:                     true,
	}

	data, err := json.Marshal(msg)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["scheduled_delivery_timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), value)
	assert.Equal(t, float64(ts.Unix()), value)
}

func TestScheduledMessageUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {
	t.Parallel()

	raw := []byte(`{"scheduled_message_id":99,"type":"channel","to":[1,2,3],"topic":"project","content":"Reminder","rendered_content":"<p>Reminder</p>","scheduled_delivery_timestamp":1700000000,"failed":false}`)

	var msg z.ScheduledMessage
	require.NoError(t, json.Unmarshal(raw, &msg))

	assert.Equal(t, int64(99), msg.ScheduledMessageId)
	assert.Equal(t, z.RecipientTypeChannel, msg.Type)
	if assert.NotNil(t, msg.To.Users) {
		assert.Equal(t, []int64{1, 2, 3}, msg.To.Users)
	}
	require.NotNil(t, msg.Topic)
	assert.Equal(t, "project", *msg.Topic)
	assert.Equal(t, "Reminder", msg.Content)
	assert.Equal(t, "<p>Reminder</p>", msg.RenderedContent)
	assert.Equal(t, int64(1700000000), msg.ScheduledDeliveryTimestamp.Unix())
	assert.Equal(t, time.UTC, msg.ScheduledDeliveryTimestamp.Location())
	assert.False(t, msg.Failed)
}
