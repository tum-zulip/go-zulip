package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	z "github.com/tum-zulip/go-zulip/zulip"
)

func TestAttachmentMarshalJSON_UsesUnixMillis(t *testing.T) {
	timestamp := time.UnixMilli(1700000000123)
	attachment := z.Attachment{
		CreateTime: timestamp,
	}

	data, err := json.Marshal(attachment)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["create_time"]
	require.True(t, ok, "create_time should be present in payload")
	require.IsType(t, float64(0), value)
	assert.InEpsilon(t, float64(timestamp.UnixMilli()), value, 0.001)
}

func TestAttachmentUnmarshalJSON_ParsesUnixMillis(t *testing.T) {
	raw := []byte(`{"create_time":1700000000123}`)

	var attachment z.Attachment
	require.NoError(t, json.Unmarshal(raw, &attachment))

	if assert.NotNil(t, attachment.CreateTime) {
		assert.Equal(t, int64(1700000000123), attachment.CreateTime.UnixMilli())
	}
}

func TestAttachmentMessagesMarshalJSON_UsesUnixMillis(t *testing.T) {
	timestamp := time.UnixMilli(1700000000456)
	message := z.AttachmentMessages{
		DateSent: timestamp,
	}

	data, err := json.Marshal(message)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["date_sent"]
	require.True(t, ok, "date_sent should be present in payload")
	require.IsType(t, float64(0), value)
	assert.InEpsilon(t, float64(timestamp.UnixMilli()), value, 0.001)
}

func TestAttachmentMessagesUnmarshalJSON_ParsesUnixMillis(t *testing.T) {
	raw := []byte(`{"date_sent":1700000000456}`)

	var message z.AttachmentMessages
	require.NoError(t, json.Unmarshal(raw, &message))

	if assert.NotNil(t, message.DateSent) {
		assert.Equal(t, int64(1700000000456), message.DateSent.UnixMilli())
	}
}
