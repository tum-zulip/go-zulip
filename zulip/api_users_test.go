package zulip_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func Test_UsersAPIService_ActivateAndDeactivateUser(t *testing.T) {
	otherUserId := getOwnUserId(t, GetOtherNormalClient(t))

	t.Run("DeactivateUser", runForAdminAndOwnerClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()
		// ensure the user is active before deactivating
		apiClient.ReactivateUser(ctx, otherUserId).Execute()

		resp, httpRes, err := apiClient.DeactivateUser(ctx, otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	}))

	t.Run("ReactivateUser", runForAdminAndOwnerClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		// ensure the user is deactivated before reactivating
		apiClient.DeactivateUser(ctx, otherUserId).Execute()

		resp, httpRes, err := apiClient.ReactivateUser(ctx, otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	}))

	t.Run("DeactivateOwnUser", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		t.Skip("TODO: This test deactivates the user running the tests, so it should be the last test and the client should be recreated after this.")
		ctx := context.Background()

		resp, httpRes, err := apiClient.DeactivateOwnUser(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

}

func Test_UsersAPIService(t *testing.T) {
	t.Parallel()

	otherUserId := getOwnUserId(t, GetOtherNormalClient(t))

	t.Run("CreateUser", runForAdminAndOwnerClients(t, func(t *testing.T, apiClient zulip.Client) {
		t.Skip("TODO: not implemented")
		ctx := context.Background()

		resp, httpRes, err := apiClient.CreateUser(ctx).Email("test@example.com").Password("password").FullName("Test User").Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("AddAlertWords", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.AddAlertWords(ctx).AlertWords([]string{"word1", "word2"}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("CreateUserGroup", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.CreateUserGroup(ctx).Name(uniqueName("test-usergroup")).Description("Test User Group").Members([]int64{getOwnUserId(t, apiClient)}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("DeactivateOwnUser", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		t.Skip("TODO: This test deactivates the user running the tests, so it should be the last test and the client should be recreated after this.")
		ctx := context.Background()

		resp, httpRes, err := apiClient.DeactivateOwnUser(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("DeactivateUserGroup", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		userGroupId := createRandomUserGroup(t, apiClient, getOwnUserId(t, apiClient))
		resp, httpRes, err := apiClient.DeactivateUserGroup(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetAlertWords", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetAlertWords(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetAttachments", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetAttachments(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetIsUserGroupMember", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		userId := getOwnUserId(t, apiClient)
		userGroupId := createRandomUserGroup(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetIsUserGroupMember(ctx, userGroupId, userId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.True(t, resp.IsUserGroupMember)
	}))

	t.Run("GetOwnUser", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetOwnUser(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUser", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetUser(ctx, getOwnUserId(t, apiClient)).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserByEmail", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		var email string = getOwnUserEmail(t, apiClient)

		resp, httpRes, err := apiClient.GetUserByEmail(ctx, email).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserGroupMembers", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		userGroupId := createRandomUserGroup(t, apiClient, getOwnUserId(t, apiClient))

		resp, httpRes, err := apiClient.GetUserGroupMembers(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserGroupSubgroups", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		userGroupId := createRandomUserGroup(t, apiClient, getOwnUserId(t, apiClient))
		subGroupId := createRandomUserGroup(t, apiClient, getOwnUserId(t, apiClient))
		sgResp, sgHttpResp, err := apiClient.UpdateUserGroupSubgroups(ctx, userGroupId).Add([]int64{subGroupId}).Execute()
		require.NoError(t, err)
		require.NotNil(t, sgResp)
		assert.Equal(t, 200, sgHttpResp.StatusCode)

		resp, httpRes, err := apiClient.GetUserGroupSubgroups(ctx, userGroupId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserGroups", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetUserGroups(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserPresence", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		var userIdOrEmail string = fmt.Sprintf("%d", getOwnUserId(t, apiClient))

		resp, httpRes, err := apiClient.GetUserPresence(ctx, userIdOrEmail).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUserStatus", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		var userId int64 = getOwnUserId(t, apiClient)

		resp, httpRes, err := apiClient.GetUserStatus(ctx, userId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetUsers", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetUsers(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("MuteUser", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		// ensure the user is not muted before muting
		apiClient.UnmuteUser(ctx, otherUserId).Execute()

		resp, httpRes, err := apiClient.MuteUser(ctx, otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("RemoveAlertWords", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		alterWord := uniqueName("word")
		// ensure the alert word is added before removing
		apiClient.AddAlertWords(ctx).AlertWords([]string{alterWord}).Execute()

		resp, httpRes, err := apiClient.RemoveAlertWords(ctx).AlertWords([]string{alterWord}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("SetTypingStatus", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.SetTypingStatus(ctx).
			Op(zulip.TypingStatusOpStart).
			To(zulip.UserAsRecipient(getOwnUserId(t, apiClient))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("SetTypingStatusForMessageEdit", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		messageId := createDirectMessage(t, apiClient, otherUserId)

		resp, httpRes, err := apiClient.SetTypingStatusForMessageEdit(ctx, messageId).
			Op(zulip.TypingStatusOpStart).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UnmuteUser", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		// ensure the user is muted before unmuting
		apiClient.MuteUser(ctx, otherUserId).Execute()

		resp, httpRes, err := apiClient.UnmuteUser(ctx, otherUserId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdatePresence", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdatePresence(ctx).Status(zulip.PresenceStatusActive).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateSettings", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdateSettings(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateStatus", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdateStatus(ctx).StatusText(uniqueName("status")).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateStatusForUser", runForAdminAndOwnerClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdateStatusForUser(ctx, getOwnUserId(t, apiClient)).StatusText(uniqueName("status")).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUser", runForAdminAndOwnerClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.UpdateUser(ctx, getOwnUserId(t, apiClient)).ProfileData([]map[string]interface{}{{"id": 9, "value": uniqueName("they/them")}}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUserByEmail", runForAdminAndOwnerClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		var email string = getOwnUserEmail(t, apiClient)

		resp, httpRes, err := apiClient.UpdateUserByEmail(ctx, email).ProfileData([]map[string]interface{}{{"id": 9, "value": uniqueName("they/them")}}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUserGroup", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		userGroupId := createRandomUserGroup(t, apiClient, getOwnUserId(t, apiClient))

		resp, httpRes, err := apiClient.UpdateUserGroup(ctx, userGroupId).Description(uniqueName("test group")).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUserGroupMembers", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		userGroupId := createRandomUserGroup(t, apiClient, getOwnUserId(t, apiClient))
		resp, httpRes, err := apiClient.UpdateUserGroupMembers(ctx, userGroupId).Add([]int64{otherUserId}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("UpdateUserGroupSubgroups", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		userId := getOwnUserId(t, apiClient)
		userGroupId := createRandomUserGroup(t, apiClient, userId)
		subGroupId := createRandomUserGroup(t, apiClient, userId)

		resp, httpRes, err := apiClient.UpdateUserGroupSubgroups(ctx, userGroupId).Add([]int64{subGroupId}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("AddApnsToken", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		t.Skip("TODO: Not implemented")
		resp, httpRes, err := apiClient.AddApnsToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("AddFcmToken", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		t.Skip("TODO: Not implemented")
		resp, httpRes, err := apiClient.AddFcmToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("RemoveApnsToken", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		t.Skip("TODO: Not implemented")
		resp, httpRes, err := apiClient.RemoveApnsToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("RemoveAttachment", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		// Upload a file using the test helper
		uploadResp := uploadFileForTest(t, ctx, apiClient)
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

	t.Run("RemoveFcmToken", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()
		resp, httpRes, err := apiClient.RemoveFcmToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))
}

func getOwnUser(t *testing.T, apiClient zulip.Client) *zulip.GetOwnUserResponse {
	resp, httpRes, err := apiClient.GetOwnUser(context.Background()).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	return resp
}

func getOwnUserId(t *testing.T, apiClient zulip.Client) int64 {
	t.Helper()

	resp := getOwnUser(t, apiClient)

	return resp.UserId
}

func getOwnUserEmail(t *testing.T, apiClient zulip.Client) string {
	t.Helper()

	resp := getOwnUser(t, apiClient)

	return resp.Email
}

func createRandomUserGroup(t *testing.T, apiClient zulip.Client, members ...int64) int64 {
	t.Helper()

	groupId := rand.Intn(1000000)

	resp, httpRes, err := apiClient.CreateUserGroup(context.Background()).Name(fmt.Sprintf("test-group-%d", groupId)).Description("Test Group").Members(members).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)

	return resp.GroupId
}
