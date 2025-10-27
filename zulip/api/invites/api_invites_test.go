package invites_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_InviteLinkLifecycle(t *testing.T) {
	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		channelName, channelID := CreateRandomChannel(t, apiClient, GetUserID(t, apiClient))

		baseline := inviteSnapshot(t, ctx, apiClient)

		resp, httpResp, err := apiClient.CreateInviteLink(ctx).
			ChannelIDs([]int64{channelID}).
			IncludeRealmDefaultSubscriptions(true).
			InviteExpiresInMinutes(60).
			WelcomeMessageCustomText(fmt.Sprintf("Welcome via %s", channelName)).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		link := resp.InviteLink
		assert.NotEmpty(t, link)

		invites := loadInvites(t, ctx, apiClient)
		newInvite := findNewInvite(baseline, invites, func(inv z.Invite) bool {
			return inv.IsMultiuse && strings.EqualFold(inv.LinkUrl, link)
		})
		require.NotNil(t, newInvite, "created invite link not present in GetInvites response")
		require.NotZero(t, newInvite.ID)

		revokeResp, revokeHTTP, err := apiClient.RevokeInviteLink(ctx, newInvite.ID).Execute()
		require.NoError(t, err)
		require.NotNil(t, revokeResp)
		RequireStatusOK(t, revokeHTTP)
		assert.Equal(t, "success", revokeResp.Result)

		remaining := loadInvites(t, ctx, apiClient)
		assert.Nil(t, findNewInvite(baseline, remaining, func(inv z.Invite) bool {
			return inv.IsMultiuse && inv.ID == newInvite.ID
		}), "multiuse invite should be removed after revocation")
	})
}

func Test_EmailInviteLifecycle(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		_, channelID := CreateRandomChannel(t, apiClient, GetUserID(t, apiClient))
		invitee := fmt.Sprintf("%s@zulip.com", strings.ToLower(UniqueName("invitee")))

		baseline := inviteSnapshot(t, ctx, apiClient)

		resp, httpResp, err := apiClient.SendInvites(ctx).
			InviteeEmails(invitee).
			ChannelIDs([]int64{channelID}).
			InviteExpiresInMinutes(60).
			NotifyReferrerOnJoin(true).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, "success", resp.Result)

		invites := loadInvites(t, ctx, apiClient)
		emailInvite := findNewInvite(baseline, invites, func(inv z.Invite) bool {
			return !inv.IsMultiuse && strings.EqualFold(inv.Email, invitee)
		})
		require.NotNil(t, emailInvite, "email invitation not found in GetInvites response")
		require.NotZero(t, emailInvite.ID)

		resendResp, resendHTTP, err := apiClient.ResendEmailInvite(ctx, emailInvite.ID).Execute()
		require.NoError(t, err)
		require.NotNil(t, resendResp)
		RequireStatusOK(t, resendHTTP)
		assert.Equal(t, "success", resendResp.Result)

		revokeResp, revokeHTTP, err := apiClient.RevokeEmailInvite(ctx, emailInvite.ID).Execute()
		require.NoError(t, err)
		require.NotNil(t, revokeResp)
		RequireStatusOK(t, revokeHTTP)
		assert.Equal(t, "success", revokeResp.Result)

		remaining := loadInvites(t, ctx, apiClient)
		assert.Nil(t, findNewInvite(baseline, remaining, func(inv z.Invite) bool {
			return !inv.IsMultiuse && inv.ID == emailInvite.ID
		}), "email invitation should be removed after revocation")
	})
}

func loadInvites(t *testing.T, ctx context.Context, apiClient client.Client) []z.Invite {
	t.Helper()

	resp, httpResp, err := apiClient.GetInvites(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp.Invites
}

func inviteSnapshot(t *testing.T, ctx context.Context, apiClient client.Client) map[string]struct{} {
	invites := loadInvites(t, ctx, apiClient)
	snapshot := make(map[string]struct{}, len(invites))
	for _, inv := range invites {
		snapshot[inviteKey(inv)] = struct{}{}
	}
	return snapshot
}

func findNewInvite(snapshot map[string]struct{}, invites []z.Invite, match func(z.Invite) bool) *z.Invite {
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

func inviteKey(inv z.Invite) string {
	email := strings.ToLower(inv.Email)
	link := strings.ToLower(inv.LinkUrl)
	return fmt.Sprintf("%d:%t:%s:%s", inv.ID, inv.IsMultiuse, email, link)
}
