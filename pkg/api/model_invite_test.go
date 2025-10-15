package api_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tum-zulip/go-zulip/pkg/api"
)

func TestInviteMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	t.Parallel()

	invited := time.Unix(1700000000, 123000000).UTC()
	expiry := time.Unix(1700003600, 0).UTC()
	invite := api.Invite{
		Id:                   10,
		InvitedByUserId:      3,
		Invited:              invited,
		ExpiryDate:           &expiry,
		InvitedAs:            400,
		Email:                "user@example.com",
		NotifyReferrerOnJoin: true,
		LinkUrl:              "https://example.com/invite",
		IsMultiuse:           false,
	}

	data, err := json.Marshal(invite)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	invitedValue, ok := payload["invited"]
	require.True(t, ok)
	require.IsType(t, float64(0), invitedValue)
	assert.Equal(t, float64(invited.Unix()), invitedValue)

	expiryValue, ok := payload["expiry_date"]
	require.True(t, ok)
	require.IsType(t, float64(0), expiryValue)
	assert.Equal(t, float64(expiry.Unix()), expiryValue)
}

func TestInviteUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {
	t.Parallel()

	raw := []byte(`{"id":10,"invited_by_user_id":3,"invited":1700000000,"expiry_date":1700003600,"invited_as":400,"email":"user@example.com","notify_referrer_on_join":true,"link_url":"https://example.com/invite","is_multiuse":false}`)

	var invite api.Invite
	require.NoError(t, json.Unmarshal(raw, &invite))

	assert.Equal(t, int64(10), invite.Id)
	assert.Equal(t, int64(3), invite.InvitedByUserId)
	assert.Equal(t, int64(1700000000), invite.Invited.Unix())
	assert.Equal(t, time.UTC, invite.Invited.Location())

	if assert.NotNil(t, invite.ExpiryDate) {
		assert.Equal(t, int64(1700003600), invite.ExpiryDate.Unix())
		assert.Equal(t, time.UTC, invite.ExpiryDate.Location())
	}

	assert.Equal(t, int32(400), invite.InvitedAs)
	assert.Equal(t, "user@example.com", invite.Email)
	assert.True(t, invite.NotifyReferrerOnJoin)
	assert.Equal(t, "https://example.com/invite", invite.LinkUrl)
	assert.False(t, invite.IsMultiuse)
}
