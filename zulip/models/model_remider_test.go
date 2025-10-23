package models_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func TestReminderMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	t.Parallel()

	ts := time.Unix(1700000000, 789000000).UTC()

	reminder := z.Reminder{
		ReminderId:                 123,
		Type:                       z.RecipientTypePrivate,
		To:                         []int64{1, 2, 3},
		Content:                    "Don't forget",
		RenderedContent:            "<p>Don't forget</p>",
		ScheduledDeliveryTimestamp: ts,
		Failed:                     true,
		ReminderTargetMessageId:    456,
	}

	data, err := json.Marshal(reminder)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["scheduled_delivery_timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), value)
	assert.Equal(t, float64(ts.Unix()), value)
}

func TestReminderUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {
	t.Parallel()

	raw := []byte(`{"reminder_id":123,"type":"private","to":[1,2,3],"content":"Don't forget","rendered_content":"<p>Don't forget</p>","scheduled_delivery_timestamp":1700000000,"failed":true,"reminder_target_message_id":456}`)

	var reminder z.Reminder
	require.NoError(t, json.Unmarshal(raw, &reminder))

	assert.Equal(t, int64(123), reminder.ReminderId)
	assert.Equal(t, z.RecipientTypeDirect, reminder.Type)
	assert.Equal(t, []int64{1, 2, 3}, reminder.To)
	assert.Equal(t, "Don't forget", reminder.Content)
	assert.Equal(t, "<p>Don't forget</p>", reminder.RenderedContent)
	assert.Equal(t, int64(1700000000), reminder.ScheduledDeliveryTimestamp.Unix())
	assert.Equal(t, time.UTC, reminder.ScheduledDeliveryTimestamp.Location())
	assert.True(t, reminder.Failed)
	assert.Equal(t, int64(456), reminder.ReminderTargetMessageId)
}
