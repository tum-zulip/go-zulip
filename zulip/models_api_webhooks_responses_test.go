package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func TestZulipOutgoingWebhooksResponse_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	raw := []byte(`{
        "bot_email": "outgoing-bot@localhost",
        "bot_full_name": "Outgoing webhook test",
        "data": "@**Outgoing webhook test** Zulip is the world’s most productive group chat!",
        "trigger": "mention",
        "token": "xvOzfurIutdRRVLzpXrIIHXJvNfaJLJ0",
        "message": {
            "subject": "Verona2",
            "sender_email": "iago@zulip.com",
            "timestamp": 1527876931,
            "client": "website",
            "submessages": [],
            "recipient_id": 20,
            "topic_links": [],
            "sender_full_name": "Iago",
            "avatar_url": "https://secure.gravatar.com/avatar/1f4f1575bf002ae562fea8fc4b861b09?d=identicon&version=1",
            "rendered_content": "<p><span class=\"user-mention\" data-user-id=\"25\">@Outgoing webhook test</span> Zulip is the world’s most productive group chat!</p>",
            "sender_id": 5,
            "stream_id": 5,
            "content": "@**Outgoing webhook test** Zulip is the world’s most productive group chat!",
            "display_recipient": "Verona",
            "type": "stream",
            "id": 112,
            "is_me_message": false,
            "reactions": [],
            "sender_realm_str": "zulip"
        }
    }`)

	var resp zulip.ZulipOutgoingWebhooksResponse
	require.NoError(t, json.Unmarshal(raw, &resp))

	assert.Equal(t, "outgoing-bot@localhost", resp.BotEmail)
	assert.Equal(t, "Outgoing webhook test", resp.BotFullName)
	assert.Equal(t, "@**Outgoing webhook test** Zulip is the world’s most productive group chat!", resp.Data)
	assert.Equal(t, zulip.WebhookTriggerMention, resp.Trigger)
	assert.Equal(t, "xvOzfurIutdRRVLzpXrIIHXJvNfaJLJ0", resp.Token)

	if assert.NotNil(t, resp.Message) {
		msg := resp.Message
		assert.Equal(t, int64(112), msg.Id)
		assert.Equal(t, "Verona2", msg.Subject)
		assert.Equal(t, "iago@zulip.com", msg.SenderEmail)
		assert.Equal(t, "Iago", msg.SenderFullName)
		assert.Equal(t, int64(5), msg.SenderId)
		assert.Equal(t, "zulip", msg.SenderRealmStr)
		assert.Equal(t, zulip.RecipientTypeStream, msg.Type)
		assert.Equal(t, "website", msg.Client)
		assert.Equal(t, "@**Outgoing webhook test** Zulip is the world’s most productive group chat!", msg.Content)
		require.NotNil(t, msg.RenderedContent)
		assert.Contains(t, *msg.RenderedContent, "@Outgoing webhook test")
		require.NotNil(t, msg.AvatarUrl)
		assert.Contains(t, *msg.AvatarUrl, "gravatar.com")
		require.NotNil(t, msg.ChannelId)
		assert.Equal(t, int64(5), *msg.ChannelId)
		assert.Equal(t, int64(1527876931), msg.Timestamp.Unix())
		assert.False(t, msg.IsMeMessage)
	}
}

func TestZulipOutgoingWebhooksResponse_MarshalJSON(t *testing.T) {
	t.Parallel()

	streamID := int64(7)
	channelName := "integrations"
	rendered := "<p><strong>Hello</strong> world!</p>"
	timestamp := time.Unix(1_700_000_100, 0).UTC()

	resp := zulip.ZulipOutgoingWebhooksResponse{
		BotEmail:    "hook@example.com",
		BotFullName: "Webhook Tester",
		Data:        "**Hello** world!",
		Trigger:     zulip.WebhookTriggerMention,
		Token:       "totally-secret-token",
		Message: &zulip.Message{
			Id:               99,
			Type:             zulip.RecipientTypeStream,
			Client:           "website",
			Content:          "**Hello** world!",
			RenderedContent:  &rendered,
			SenderEmail:      "iago@zulip.com",
			SenderFullName:   "Iago",
			SenderId:         5,
			SenderRealmStr:   "zulip",
			RecipientId:      42,
			ChannelId:        &streamID,
			Subject:          "Greetings",
			DisplayRecipient: zulip.ChannelNameAsDisplayRecipient(&channelName),
			Timestamp:        timestamp,
		},
	}

	data, err := json.Marshal(resp)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	assert.Equal(t, "hook@example.com", payload["bot_email"])

	messageValue, ok := payload["message"]
	require.True(t, ok)
	messageMap, ok := messageValue.(map[string]any)
	require.True(t, ok)

	timestampValue, ok := messageMap["timestamp"]
	require.True(t, ok)
	require.IsType(t, float64(0), timestampValue)
	assert.Equal(t, float64(timestamp.Unix()), timestampValue)

	assert.Equal(t, "integrations", messageMap["display_recipient"])
	assert.Equal(t, float64(streamID), messageMap["stream_id"])
	assert.Equal(t, rendered, messageMap["rendered_content"])
}
