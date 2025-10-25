package users_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_DeactivateUser(t *testing.T) {
	t.Parallel()

	deactivateUserClient := GetTestClient(t, DeactivateTestUser)
	deactivateUserId := GetUserId(t, deactivateUserClient)

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		// ensure the user is active before deactivating
		apiClient.ReactivateUser(ctx, deactivateUserId).Execute()

		resp, httpResp, err := apiClient.DeactivateUser(ctx, deactivateUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_ReactivateUser(t *testing.T) {
	t.Parallel()

	deactivateUserClient := GetTestClient(t, DeactivateTestUser)
	deactivateUserId := GetUserId(t, deactivateUserClient)

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is deactivated before reactivating
		apiClient.DeactivateUser(ctx, deactivateUserId).Execute()

		resp, httpResp, err := apiClient.ReactivateUser(ctx, deactivateUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_DeactivateOwnUser(t *testing.T) {
	t.Parallel()

	t.Skip("TODO: This test deactivates the user running the tests, so it should be the last test and the client should be recreated after this.")
	ctx := context.Background()

	deactivateUserClient := GetTestClient(t, DeactivateTestUser)

	resp, httpResp, err := deactivateUserClient.DeactivateOwnUser(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpResp.StatusCode)

}

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: not implemented")
		ctx := context.Background()

		resp, httpResp, err := apiClient.CreateUser(ctx).Email("test@example.com").Password("password").FullName("Test User").Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_AddAlertWords(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.AddAlertWords(ctx).AlertWords([]string{"word1", "word2"}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_CreateUserGroup(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.CreateUserGroup(ctx).Name(UniqueName("test-usergroup")).Description("Test User Group").Members([]int64{GetUserId(t, apiClient)}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_DeactivateUserGroup(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))
		resp, httpResp, err := apiClient.DeactivateUserGroup(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetAlertWords(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetAlertWords(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetAttachments(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetAttachments(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetIsUserGroupMember(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userId := GetUserId(t, apiClient)
		userGroupId := CreateRandomUserGroup(t, apiClient, userId)

		resp, httpResp, err := apiClient.GetIsUserGroupMember(ctx, userGroupId, userId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
		assert.True(t, resp.IsUserGroupMember)
	})
}

func Test_GetOwnUser(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetOwnUser(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetUser(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetUser(ctx, GetUserId(t, apiClient)).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetUserByEmail(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var email string = GetUserEmail(t, apiClient)

		resp, httpResp, err := apiClient.GetUserByEmail(ctx, email).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetUserGroupMembers(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))

		resp, httpResp, err := apiClient.GetUserGroupMembers(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetUserGroupSubgroups(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))
		subGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))
		sgResp, sgHttpResp, err := apiClient.UpdateUserGroupSubgroups(ctx, userGroupId).Add([]int64{subGroupId}).Execute()
		require.NoError(t, err)
		require.NotNil(t, sgResp)
		assert.Equal(t, 200, sgHttpResp.StatusCode)

		resp, httpResp, err := apiClient.GetUserGroupSubgroups(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetUserGroups(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))

		resp, httpResp, err := apiClient.GetUserGroups(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

		found := false
		for _, g := range resp.UserGroups {
			if g.Id == userGroupId {
				found = true
				require.NotNil(t, g.DateCreated)
				require.WithinDuration(t, time.Now(), *g.DateCreated, 3*time.Minute)
			}
		}
		require.True(t, found)

	})
}

func Test_GetUserPresence(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var userIdOrEmail string = fmt.Sprintf("%d", GetUserId(t, apiClient))

		resp, httpResp, err := apiClient.GetUserPresence(ctx, userIdOrEmail).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetUserStatus(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var userId int64 = GetUserId(t, apiClient)

		resp, httpResp, err := apiClient.GetUserStatus(ctx, userId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetUsers(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetUsers(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_MuteUser(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)
	otherUserId := GetUserId(t, otherClient)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is not muted before muting
		apiClient.UnmuteUser(ctx, otherUserId).Execute()

		resp, httpResp, err := apiClient.MuteUser(ctx, otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_RemoveAlertWords(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		alterWord := UniqueName("word")
		// ensure the alert word is added before removing
		apiClient.AddAlertWords(ctx).AlertWords([]string{alterWord}).Execute()

		resp, httpResp, err := apiClient.RemoveAlertWords(ctx).AlertWords([]string{alterWord}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_SetTypingStatus(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.SetTypingStatus(ctx).
			Op(z.TypingStatusOpStart).
			To(z.UserAsRecipient(GetUserId(t, apiClient))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_SetTypingStatusForMessageEdit(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)
	otherUserId := GetUserId(t, otherClient)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		messageId := CreateDirectMessage(t, apiClient, otherUserId)

		resp, httpResp, err := apiClient.SetTypingStatusForMessageEdit(ctx, messageId).
			Op(z.TypingStatusOpStart).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UnmuteUser(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)
	otherUserId := GetUserId(t, otherClient)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is muted before unmuting
		apiClient.MuteUser(ctx, otherUserId).Execute()

		resp, httpResp, err := apiClient.UnmuteUser(ctx, otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UpdatePresence(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdatePresence(ctx).Status(z.PresenceStatusActive).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
		{
			userIdOrEmail := fmt.Sprintf("%d", GetUserId(t, apiClient))

			resp, httpResp, err := apiClient.GetUserPresence(ctx, userIdOrEmail).Execute()

			require.NoError(t, err)
			require.NotNil(t, resp)
			RequireStatusOK(t, httpResp)
			require.NotNil(t, resp.Presence)
			p := resp.Presence["aggregated"]
			require.WithinDuration(t, time.Now(), p.Timestamp, 3*time.Minute)
		}

	})
}

func Test_UpdateSettings(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateSettings(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UpdateStatus(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateStatus(ctx).StatusText(UniqueName("status")).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UpdateStatusForUser(t *testing.T) {
	t.Parallel()

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateStatusForUser(ctx, GetUserId(t, apiClient)).StatusText(UniqueName("status")).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UpdateUser(t *testing.T) {
	t.Parallel()

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateUser(ctx, GetUserId(t, apiClient)).ProfileData([]map[string]interface{}{{"id": 9, "value": UniqueName("they/them")}}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UpdateUserByEmail(t *testing.T) {
	t.Parallel()

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var email string = GetUserEmail(t, apiClient)

		resp, httpResp, err := apiClient.UpdateUserByEmail(ctx, email).ProfileData([]map[string]interface{}{{"id": 9, "value": UniqueName("they/them")}}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UpdateUserGroup(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))

		resp, httpResp, err := apiClient.UpdateUserGroup(ctx, userGroupId).Description(UniqueName("test group")).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UpdateUserGroupMembers(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)
	otherUserId := GetUserId(t, otherClient)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))
		resp, httpResp, err := apiClient.UpdateUserGroupMembers(ctx, userGroupId).Add(otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_UpdateUserGroupSubgroups(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userId := GetUserId(t, apiClient)
		userGroupId := CreateRandomUserGroup(t, apiClient, userId)
		subGroupId := CreateRandomUserGroup(t, apiClient, userId)

		resp, httpResp, err := apiClient.UpdateUserGroupSubgroups(ctx, userGroupId).Add([]int64{subGroupId}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_AddApnsToken(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()

		resp, httpResp, err := apiClient.AddApnsToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_AddFcmToken(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()

		resp, httpResp, err := apiClient.AddFcmToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_RemoveApnsToken(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()

		resp, httpResp, err := apiClient.RemoveApnsToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_RemoveAttachment(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Upload a file using the test helper
		uploadResp := UploadFileForTest(t, ctx, apiClient)
		fileName := uploadResp.Filename

		// Get the attachmentId from GetAttachments
		attachmentsResp, attachmentsHttpRes, err := apiClient.GetAttachments(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, attachmentsResp)
		assert.Equal(t, 200, attachmentsHttpRes.StatusCode)
		var attachmentId int64 = -1
		for _, att := range attachmentsResp.Attachments {
			if att.Name == fileName {
				attachmentId = att.Id
				break
			}
		}
		require.NotEqual(t, int64(-1), attachmentId, "Uploaded attachment not found")

		resp, httpResp, err := apiClient.RemoveAttachment(ctx, attachmentId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_RemoveFcmToken(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()
		resp, httpResp, err := apiClient.RemoveFcmToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}
