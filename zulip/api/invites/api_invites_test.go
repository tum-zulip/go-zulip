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

		baseline := inviteSnapshot(ctx, t, apiClient)

		resp, _, err := apiClient.CreateInviteLink(ctx).
			ChannelIDs([]int64{channelID}).
			IncludeRealmDefaultSubscriptions(true).
			InviteExpiresInMinutes(60).
			WelcomeMessageCustomText(fmt.Sprintf("Welcome via %s", channelName)).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		link := resp.InviteLink
		assert.NotEmpty(t, link)

		invites := loadInvites(ctx, t, apiClient)
		newInvite := findNewInvite(baseline, invites, func(inv z.Invite) bool {
			return inv.IsMultiuse && strings.EqualFold(inv.LinkURL, link)
		})
		require.NotNil(t, newInvite, "created invite link not present in GetInvites Response")
		require.NotZero(t, newInvite.ID)

		revokeResp, _, err := apiClient.RevokeInviteLink(ctx, newInvite.ID).Execute()
		require.NoError(t, err)
		require.NotNil(t, revokeResp)
		assert.Equal(t, "success", revokeResp.Result)

		remaining := loadInvites(ctx, t, apiClient)
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

		baseline := inviteSnapshot(ctx, t, apiClient)

		resp, _, err := apiClient.SendInvites(ctx).
			InviteeEmails(invitee).
			ChannelIDs([]int64{channelID}).
			InviteExpiresInMinutes(60).
			NotifyReferrerOnJoin(true).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, "success", resp.Result)

		invites := loadInvites(ctx, t, apiClient)
		emailInvite := findNewInvite(baseline, invites, func(inv z.Invite) bool {
			return !inv.IsMultiuse && strings.EqualFold(inv.Email, invitee)
		})
		require.NotNil(t, emailInvite, "email invitation not found in GetInvites Response")
		require.NotZero(t, emailInvite.ID)

		resendResp, _, err := apiClient.ResendEmailInvite(ctx, emailInvite.ID).Execute()
		require.NoError(t, err)
		require.NotNil(t, resendResp)
		assert.Equal(t, "success", resendResp.Result)

		revokeResp, _, err := apiClient.RevokeEmailInvite(ctx, emailInvite.ID).Execute()
		require.NoError(t, err)
		require.NotNil(t, revokeResp)
		assert.Equal(t, "success", revokeResp.Result)

		remaining := loadInvites(ctx, t, apiClient)
		assert.Nil(t, findNewInvite(baseline, remaining, func(inv z.Invite) bool {
			return !inv.IsMultiuse && inv.ID == emailInvite.ID
		}), "email invitation should be removed after revocation")
	})
}

func loadInvites(ctx context.Context, t *testing.T, apiClient client.Client) []z.Invite {
	t.Helper()

	resp, _, err := apiClient.GetInvites(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)

	return resp.Invites
}

func inviteSnapshot(ctx context.Context, t *testing.T, apiClient client.Client) map[string]struct{} {
	invites := loadInvites(ctx, t, apiClient)
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
			cp := inv
			return &cp
		}
	}
	return nil
}

func inviteKey(inv z.Invite) string {
	email := strings.ToLower(inv.Email)
	link := strings.ToLower(inv.LinkURL)
	return fmt.Sprintf("%d:%t:%s:%s", inv.ID, inv.IsMultiuse, email, link)
}
