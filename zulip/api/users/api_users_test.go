package users_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_DeactivateUser(t *testing.T) {
	_, deactivateUserClient := GetTestClient(t, DeactivateTestUser)
	deactivateUserID := GetUserID(t, deactivateUserClient)

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is active before deactivating
		//nolint: errcheck, gosec // returns an error if the user is already active, which we can ignore
		apiClient.ReactivateUser(ctx, deactivateUserID).Execute()

		resp, httpResp, err := apiClient.DeactivateUser(ctx, deactivateUserID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_ReactivateUser(t *testing.T) {
	_, deactivateUserClient := GetTestClient(t, DeactivateTestUser)
	deactivateUserID := GetUserID(t, deactivateUserClient)

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		{
			// ensure the user is deactivated before reactivating
			_, _, _ = apiClient.DeactivateUser(ctx, deactivateUserID).Execute()
		}

		resp, httpResp, err := apiClient.ReactivateUser(ctx, deactivateUserID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_DeactivateOwnUser(t *testing.T) {
	t.Skip(
		"TODO: This test deactivates the user running the tests, so it should be the last test and the client should be recreated after this.",
	)
	ctx := context.Background()

	_, deactivateUserClient := GetTestClient(t, DeactivateTestUser)

	resp, httpResp, err := deactivateUserClient.DeactivateOwnUser(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpResp.StatusCode)
}

func Test_CreateUser(t *testing.T) {
	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: not implemented")
		ctx := context.Background()

		resp, httpResp, err := apiClient.CreateUser(ctx).
			Email("test@example.com").
			Password("password").
			FullName("Test User").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_AddAlertWords(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.AddAlertWords(ctx).AlertWords([]string{"word1", "word2"}).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_CreateUserGroup(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.CreateUserGroup(ctx).
			Name(UniqueName("test-usergroup")).
			Description("Test User Group").
			Members([]int64{GetUserID(t, apiClient)}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_DeactivateUserGroup(t *testing.T) {
	RequireFeatureLevel(t, 290)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupID := CreateRandomUserGroup(t, apiClient, GetUserID(t, apiClient))
		resp, httpResp, err := apiClient.DeactivateUserGroup(ctx, userGroupID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetAlertWords(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetAlertWords(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetAttachments(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetAttachments(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetIsUserGroupMember(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userID := GetUserID(t, apiClient)
		userGroupID := CreateRandomUserGroup(t, apiClient, userID)

		resp, httpResp, err := apiClient.GetIsUserGroupMember(ctx, userGroupID, userID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
		assert.True(t, resp.IsUserGroupMember)
	})
}

func Test_GetOwnUser(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetOwnUser(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetUser(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetUser(ctx, GetUserID(t, apiClient)).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetUserByEmail(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		email := GetUserEmail(t, apiClient)

		resp, httpResp, err := apiClient.GetUserByEmail(ctx, email).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetUserGroupMembers(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupID := CreateRandomUserGroup(t, apiClient, GetUserID(t, apiClient))

		resp, httpResp, err := apiClient.GetUserGroupMembers(ctx, userGroupID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetUserGroupSubgroups(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupID := CreateRandomUserGroup(t, apiClient, GetUserID(t, apiClient))
		subGroupID := CreateRandomUserGroup(t, apiClient, GetUserID(t, apiClient))

		sgResp, sgHTTPResp, err := apiClient.UpdateUserGroupSubgroups(ctx, userGroupID).
			Add([]int64{subGroupID}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, sgResp)
		assert.Equal(t, 200, sgHTTPResp.StatusCode)

		resp, httpResp, err := apiClient.GetUserGroupSubgroups(ctx, userGroupID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetUserGroups(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupID := CreateRandomUserGroup(t, apiClient, GetUserID(t, apiClient))

		resp, httpResp, err := apiClient.GetUserGroups(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

		found := false
		for _, g := range resp.UserGroups {
			if g.ID == userGroupID {
				found = true

				if GetFeatureLevel(t) >= 292 {
					require.NotNil(t, g.DateCreated)
					require.WithinDuration(t, time.Now(), *g.DateCreated, 3*time.Minute)
				}
			}
		}
		require.True(t, found, "Created user group not found in list")
	})
}

func Test_GetUserPresence(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userIDOrEmail := strconv.FormatInt(GetUserID(t, apiClient), 10)

		resp, httpResp, err := apiClient.GetUserPresence(ctx, userIDOrEmail).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetUserStatus(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userID := GetUserID(t, apiClient)

		resp, httpResp, err := apiClient.GetUserStatus(ctx, userID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_GetUsers(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetUsers(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_MuteUser(t *testing.T) {
	otherClient := GetOtherNormalClient(t)
	otherUserID := GetUserID(t, otherClient)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is not muted before muting
		//nolint: errcheck, gosec // returns an error if the user is not muted, which we can ignore
		apiClient.UnmuteUser(ctx, otherUserID).Execute()

		resp, httpResp, err := apiClient.MuteUser(ctx, otherUserID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_RemoveAlertWords(t *testing.T) {
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
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.SetTypingStatus(ctx).
			Op(z.TypingStatusOpStart).
			To(z.UserAsRecipient(GetUserID(t, apiClient))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_SetTypingStatusForMessageEdit(t *testing.T) {
	RequireFeatureLevel(t, 361)

	otherClient := GetOtherNormalClient(t)
	otherUserID := GetUserID(t, otherClient)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		messageID := CreateDirectMessage(t, apiClient, otherUserID)

		resp, httpResp, err := apiClient.SetTypingStatusForMessageEdit(ctx, messageID).
			Op(z.TypingStatusOpStart).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UnmuteUser(t *testing.T) {
	otherClient := GetOtherNormalClient(t)
	otherUserID := GetUserID(t, otherClient)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// ensure the user is muted before unmuting
		//nolint: errcheck,gosec // returns an error if the user is not muted, which we can ignore
		apiClient.MuteUser(ctx, otherUserID).Execute()

		resp, httpResp, err := apiClient.UnmuteUser(ctx, otherUserID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UpdatePresence(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdatePresence(ctx).Status(z.PresenceStatusActive).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
		{
			userIDOrEmail := strconv.FormatInt(GetUserID(t, apiClient), 10)

			presenceResp, _, presenceErr := apiClient.GetUserPresence(ctx, userIDOrEmail).Execute()

			require.NoError(t, presenceErr)
			require.NotNil(t, presenceResp)
			require.NotNil(t, presenceResp.Presence)
			p := presenceResp.Presence["aggregated"]
			require.WithinDuration(t, time.Now(), p.Timestamp, 3*time.Minute)
		}
	})
}

func Test_UpdateSettings(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateSettings(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UpdateStatus(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateStatus(ctx).StatusText(UniqueName("status")).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UpdateStatusForUser(t *testing.T) {
	RequireFeatureLevel(t, 407)

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateStatusForUser(ctx, GetUserID(t, apiClient)).
			StatusText(UniqueName("status")).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UpdateUser(t *testing.T) {
	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateUser(ctx, GetUserID(t, apiClient)).
			ProfileData([]map[string]interface{}{{"id": 9, "value": UniqueName("they/them")}}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UpdateUserByEmail(t *testing.T) {
	RequireFeatureLevel(t, 313)

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		var email string = GetUserEmail(t, apiClient)

		resp, httpResp, err := apiClient.UpdateUserByEmail(ctx, email).
			ProfileData([]map[string]interface{}{{"id": 9, "value": UniqueName("they/them")}}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UpdateUserGroup(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupID := CreateRandomUserGroup(t, apiClient, GetUserID(t, apiClient))

		resp, httpResp, err := apiClient.UpdateUserGroup(ctx, userGroupID).
			Description(UniqueName("test group")).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UpdateUserGroupMembers(t *testing.T) {
	otherClient := GetOtherNormalClient(t)
	otherUserID := GetUserID(t, otherClient)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userGroupID := CreateRandomUserGroup(t, apiClient, GetUserID(t, apiClient))
		resp, httpResp, err := apiClient.UpdateUserGroupMembers(ctx, userGroupID).
			Add([]int64{otherUserID}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_UpdateUserGroupSubgroups(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		userID := GetUserID(t, apiClient)
		userGroupID := CreateRandomUserGroup(t, apiClient, userID)
		subGroupID := CreateRandomUserGroup(t, apiClient, userID)

		resp, httpResp, err := apiClient.UpdateUserGroupSubgroups(ctx, userGroupID).
			Add([]int64{subGroupID}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_AddApnsToken(t *testing.T) {
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
	RequireFeatureLevel(t, 285)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Upload a file using the test helper
		uploadResp := UploadFileForTest(t, ctx, apiClient)
		fileName := uploadResp.Filename

		// Get the attachmentID from GetAttachments
		attachmentsResp, attachmentsHttpRes, err := apiClient.GetAttachments(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, attachmentsResp)
		assert.Equal(t, 200, attachmentsHttpRes.StatusCode)
		var attachmentID int64 = -1
		for _, att := range attachmentsResp.Attachments {
			if att.Name == fileName {
				attachmentID = att.ID
				break
			}
		}
		require.NotEqual(t, int64(-1), attachmentID, "Uploaded attachment not found")

		resp, httpResp, err := apiClient.RemoveAttachment(ctx, attachmentID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

func Test_RemoveFcmToken(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		t.Skip("TODO: Not implemented")

		ctx := context.Background()
		resp, httpResp, err := apiClient.RemoveFcmToken(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}
