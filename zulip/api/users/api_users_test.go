package users_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_UsersAPIService_ActivateAndDeactivateUser(t *testing.T) {
	deactivateUserClient := GetTestClient(t, DeactivateTestUser)
	deactivateUserId := GetUserId(t, deactivateUserClient)

	t.Run("DeactivateUser", RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		// ensure the user is active before deactivating
		apiClient.ReactivateUser(ctx, deactivateUserId).Execute()

		resp, httpRes, err := apiClient.DeactivateUser(ctx, deactivateUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	}))

	t.Run("ReactivateUser", RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is deactivated before reactivating
		apiClient.DeactivateUser(ctx, deactivateUserId).Execute()

		resp, httpRes, err := apiClient.ReactivateUser(ctx, deactivateUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	}))

	t.Run("DeactivateOwnUser", func(t *testing.T) {
		t.Skip("TODO: This test deactivates the user running the tests, so it should be the last test and the client should be recreated after this.")
		ctx := context.Background()

		resp, httpRes, err := deactivateUserClient.DeactivateOwnUser(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})
}

func Test_UsersAPIService(t *testing.T) {
	t.Parallel()

	otherUserId := GetUserId(t, GetOtherNormalClient(t))

	t.Run("CreateUser", RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: not implemented")
		ctx := context.Background()

		resp, httpRes, err := apiClient.CreateUser(ctx).Email("test@example.com").Password("password").FullName("Test User").Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("AddAlertWords", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.AddAlertWords(ctx).AlertWords([]string{"word1", "word2"}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("CreateUserGroup", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.CreateUserGroup(ctx).Name(UniqueName("test-usergroup")).Description("Test User Group").Members([]int64{GetUserId(t, apiClient)}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("DeactivateOwnUser", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: This test deactivates the user running the tests, so it should be the last test and the client should be recreated after this.")
		ctx := context.Background()

		resp, httpRes, err := apiClient.DeactivateOwnUser(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("DeactivateUserGroup", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))
		resp, httpRes, err := apiClient.DeactivateUserGroup(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetAlertWords", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetAlertWords(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetAttachments", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetAttachments(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetIsUserGroupMember", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userId := GetUserId(t, apiClient)
		userGroupId := CreateRandomUserGroup(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetIsUserGroupMember(ctx, userGroupId, userId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.True(t, resp.IsUserGroupMember)
	}))

	t.Run("GetOwnUser", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetOwnUser(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUser", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetUser(ctx, GetUserId(t, apiClient)).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserByEmail", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var email string = GetUserEmail(t, apiClient)

		resp, httpRes, err := apiClient.GetUserByEmail(ctx, email).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserGroupMembers", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))

		resp, httpRes, err := apiClient.GetUserGroupMembers(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserGroupSubgroups", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))
		subGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))
		sgResp, sgHttpResp, err := apiClient.UpdateUserGroupSubgroups(ctx, userGroupId).Add([]int64{subGroupId}).Execute()
		require.NoError(t, err)
		require.NotNil(t, sgResp)
		assert.Equal(t, 200, sgHttpResp.StatusCode)

		resp, httpRes, err := apiClient.GetUserGroupSubgroups(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserGroups", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetUserGroups(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserPresence", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var userIdOrEmail string = fmt.Sprintf("%d", GetUserId(t, apiClient))

		resp, httpRes, err := apiClient.GetUserPresence(ctx, userIdOrEmail).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserStatus", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var userId int64 = GetUserId(t, apiClient)

		resp, httpRes, err := apiClient.GetUserStatus(ctx, userId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUsers", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetUsers(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("MuteUser", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is not muted before muting
		apiClient.UnmuteUser(ctx, otherUserId).Execute()

		resp, httpRes, err := apiClient.MuteUser(ctx, otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("RemoveAlertWords", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		alterWord := UniqueName("word")
		// ensure the alert word is added before removing
		apiClient.AddAlertWords(ctx).AlertWords([]string{alterWord}).Execute()

		resp, httpRes, err := apiClient.RemoveAlertWords(ctx).AlertWords([]string{alterWord}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("SetTypingStatus", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.SetTypingStatus(ctx).
			Op(z.TypingStatusOpStart).
			To(z.UserAsRecipient(GetUserId(t, apiClient))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("SetTypingStatusForMessageEdit", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		messageId := CreateDirectMessage(t, apiClient, otherUserId)

		resp, httpRes, err := apiClient.SetTypingStatusForMessageEdit(ctx, messageId).
			Op(z.TypingStatusOpStart).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UnmuteUser", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is muted before unmuting
		apiClient.MuteUser(ctx, otherUserId).Execute()

		resp, httpRes, err := apiClient.UnmuteUser(ctx, otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdatePresence", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdatePresence(ctx).Status(z.PresenceStatusActive).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateSettings", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdateSettings(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateStatus", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdateStatus(ctx).StatusText(UniqueName("status")).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateStatusForUser", RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdateStatusForUser(ctx, GetUserId(t, apiClient)).StatusText(UniqueName("status")).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUser", RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdateUser(ctx, GetUserId(t, apiClient)).ProfileData([]map[string]interface{}{{"id": 9, "value": UniqueName("they/them")}}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUserByEmail", RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var email string = GetUserEmail(t, apiClient)

		resp, httpRes, err := apiClient.UpdateUserByEmail(ctx, email).ProfileData([]map[string]interface{}{{"id": 9, "value": UniqueName("they/them")}}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUserGroup", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))

		resp, httpRes, err := apiClient.UpdateUserGroup(ctx, userGroupId).Description(UniqueName("test group")).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUserGroupMembers", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupId := CreateRandomUserGroup(t, apiClient, GetUserId(t, apiClient))
		resp, httpRes, err := apiClient.UpdateUserGroupMembers(ctx, userGroupId).Add([]int64{otherUserId}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUserGroupSubgroups", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userId := GetUserId(t, apiClient)
		userGroupId := CreateRandomUserGroup(t, apiClient, userId)
		subGroupId := CreateRandomUserGroup(t, apiClient, userId)

		resp, httpRes, err := apiClient.UpdateUserGroupSubgroups(ctx, userGroupId).Add([]int64{subGroupId}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("AddApnsToken", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()

		resp, httpRes, err := apiClient.AddApnsToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("AddFcmToken", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()

		resp, httpRes, err := apiClient.AddFcmToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("RemoveApnsToken", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()

		resp, httpRes, err := apiClient.RemoveApnsToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("RemoveAttachment", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
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

		resp, httpRes, err := apiClient.RemoveAttachment(ctx, attachmentId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("RemoveFcmToken", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()
		resp, httpRes, err := apiClient.RemoveFcmToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))
}
