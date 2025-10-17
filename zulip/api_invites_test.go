package zulip_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func Test_InvitesAPIService(t *testing.T) {
	t.Parallel()

	t.Run("InviteLinkLifecycle", runForAdminAndOwnerClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()
		channelName, channelId := createRandomChannel(t, apiClient, getOwnUserId(t, apiClient))

		baseline := inviteSnapshot(t, ctx, apiClient)

		resp, httpRes, err := apiClient.CreateInviteLink(ctx).
			ChannelIds([]int64{channelId}).
			IncludeRealmDefaultSubscriptions(true).
			InviteExpiresInMinutes(60).
			WelcomeMessageCustomText(fmt.Sprintf("Welcome via %s", channelName)).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		link := resp.InviteLink
		assert.NotEmpty(t, link)

		invites := loadInvites(t, ctx, apiClient)
		newInvite := findNewInvite(baseline, invites, func(inv zulip.Invite) bool {
			return inv.IsMultiuse && strings.EqualFold(inv.LinkUrl, link)
		})
		require.NotNil(t, newInvite, "created invite link not present in GetInvites response")
		require.NotZero(t, newInvite.Id)

		revokeResp, revokeHTTP, err := apiClient.RevokeInviteLink(ctx, newInvite.Id).Execute()
		require.NoError(t, err)
		require.NotNil(t, revokeResp)
		requireStatusOK(t, revokeHTTP)
		assert.Equal(t, "success", revokeResp.Result)

		remaining := loadInvites(t, ctx, apiClient)
		assert.Nil(t, findNewInvite(baseline, remaining, func(inv zulip.Invite) bool {
			return inv.IsMultiuse && inv.Id == newInvite.Id
		}), "multiuse invite should be removed after revocation")
	}))

	t.Run("EmailInviteLifecycle", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()
		_, channelId := createRandomChannel(t, apiClient, getOwnUserId(t, apiClient))
		invitee := fmt.Sprintf("%s@example.com", strings.ToLower(uniqueName("invitee")))

		baseline := inviteSnapshot(t, ctx, apiClient)

		resp, httpRes, err := apiClient.SendInvites(ctx).
			InviteeEmails(invitee).
			ChannelIds([]int64{channelId}).
			InviteExpiresInMinutes(60).
			NotifyReferrerOnJoin(true).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Equal(t, "success", resp.Result)

		invites := loadInvites(t, ctx, apiClient)
		emailInvite := findNewInvite(baseline, invites, func(inv zulip.Invite) bool {
			return !inv.IsMultiuse && strings.EqualFold(inv.Email, invitee)
		})
		require.NotNil(t, emailInvite, "email invitation not found in GetInvites response")
		require.NotZero(t, emailInvite.Id)

		resendResp, resendHTTP, err := apiClient.ResendEmailInvite(ctx, emailInvite.Id).Execute()
		require.NoError(t, err)
		require.NotNil(t, resendResp)
		requireStatusOK(t, resendHTTP)
		assert.Equal(t, "success", resendResp.Result)

		revokeResp, revokeHTTP, err := apiClient.RevokeEmailInvite(ctx, emailInvite.Id).Execute()
		require.NoError(t, err)
		require.NotNil(t, revokeResp)
		requireStatusOK(t, revokeHTTP)
		assert.Equal(t, "success", revokeResp.Result)

		remaining := loadInvites(t, ctx, apiClient)
		assert.Nil(t, findNewInvite(baseline, remaining, func(inv zulip.Invite) bool {
			return !inv.IsMultiuse && inv.Id == emailInvite.Id
		}), "email invitation should be removed after revocation")
	}))
}

func loadInvites(t *testing.T, ctx context.Context, apiClient zulip.Client) []zulip.Invite {
	t.Helper()

	resp, httpRes, err := apiClient.GetInvites(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	requireStatusOK(t, httpRes)

	return resp.Invites
}

func inviteSnapshot(t *testing.T, ctx context.Context, apiClient zulip.Client) map[string]struct{} {
	invites := loadInvites(t, ctx, apiClient)
	snapshot := make(map[string]struct{}, len(invites))
	for _, inv := range invites {
		snapshot[inviteKey(inv)] = struct{}{}
	}
	return snapshot
}

func findNewInvite(snapshot map[string]struct{}, invites []zulip.Invite, match func(zulip.Invite) bool) *zulip.Invite {
	for _, inv := range invites {
		key := inviteKey(inv)
		if _, seen := snapshot[key]; seen {
			continue
		}
		if match(inv) {
			copy := inv
			return &copy
		}
	}
	return nil
}

func inviteKey(inv zulip.Invite) string {
	email := strings.ToLower(inv.Email)
	link := strings.ToLower(inv.LinkUrl)
	return fmt.Sprintf("%d:%t:%s:%s", inv.Id, inv.IsMultiuse, email, link)
}
