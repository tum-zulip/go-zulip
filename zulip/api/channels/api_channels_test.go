package channels_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/channels"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_AddDefaultChannel(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerId := GetUserId(t, ownerClient)

	_, channelId := CreateRandomChannel(t, ownerClient, ownerId)

	resp, httpRes, err := ownerClient.AddDefaultChannel(ctx).ChannelId(channelId).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)
}

func Test_ArchiveChannel(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerId := GetUserId(t, ownerClient)

	_, channelId := CreateRandomChannel(t, ownerClient, ownerId)

	resp, httpRes, err := ownerClient.ArchiveChannel(ctx, channelId).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)
}

func Test_CreateChannelFolder(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)

	resp, httpRes, err := ownerClient.CreateChannelFolder(ctx).
		Name(UniqueName("test-folder")).
		Description("Created during Channels API tests").
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)
	assert.Greater(t, resp.ChannelFolderId, int64(0))
}

func Test_DeleteTopic(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerId := GetUserId(t, ownerClient)

	_, channelId := CreateRandomChannel(t, ownerClient, ownerId)
	topic := createTopicWithMessage(t, ownerClient, channelId)

	resp, httpRes, err := ownerClient.DeleteTopic(ctx, channelId).
		TopicName(topic).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)
	assert.True(t, resp.Complete)
}

func Test_PatchChannelFolders(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)

	createChannelFolder(t, ownerClient, UniqueName("patch-folder"), "first test folder")
	createChannelFolder(t, ownerClient, UniqueName("patch-folder"), "second test folder")

	originalOrder := getChannelFolderIds(t, ownerClient)
	if len(originalOrder) < 2 {
		t.Skip("need at least two channel folders to reorder")
	}

	reordered := append([]int64(nil), originalOrder...)
	reverseSlice(reordered)

	resp, httpRes, err := ownerClient.PatchChannelFolders(ctx).
		Order(reordered).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)
}

func Test_RemoveDefaultChannel(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerId := GetUserId(t, ownerClient)

	_, channelId := CreateRandomChannel(t, ownerClient, ownerId)
	reqResp, _, err := ownerClient.AddDefaultChannel(ctx).ChannelId(channelId).Execute()
	require.NoError(t, err)
	require.NotNil(t, reqResp)

	resp, httpRes, err := ownerClient.RemoveDefaultChannel(ctx).ChannelId(channelId).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)
}

func Test_UpdateChannelFolder(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)

	folderId := createChannelFolder(t, ownerClient, UniqueName("update-folder"), "initial description")

	resp, httpRes, err := ownerClient.UpdateChannelFolder(ctx, folderId).
		Description("updated folder description").
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)
}

func Test_UpdateChannel(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerId := GetUserId(t, ownerClient)

	_, channelId := CreateRandomChannel(t, ownerClient, ownerId)

	resp, httpRes, err := ownerClient.UpdateChannel(ctx, channelId).
		Description("updated by test").
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)
}

func Test_CreateBigBlueButtonVideoCall(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.CreateBigBlueButtonVideoCall(ctx).
			MeetingName(UniqueName("bbb-meeting")).
			Execute()
		if err != nil {
			skipIfBigBlueButtonUnavailable(t, err)
		}
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Url)
	})
}

func Test_CreateChannel(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		userId := GetUserId(t, apiClient)

		CreateRandomChannel(t, apiClient, userId)
	})
}

func Test_GetChannelFolders(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetChannelFolders(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_GetChannelById(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		_, channelId := CreateRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetChannelById(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.Equal(t, channelId, resp.Channel.ChannelId)
	})
}

func Test_GetChannelEmailAddress(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		_, channelId := CreateRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetChannelEmailAddress(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Email)
	})
}

func Test_GetChannelId(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		channelName, channelId := CreateRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetChannelId(ctx).
			Channel(channelName).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.Equal(t, channelId, resp.ChannelId)
	})
}

func Test_GetChannelTopics(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		_, channelId := CreateRandomChannel(t, apiClient, userId)
		topic := createTopicWithMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.GetChannelTopics(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)

		found := false
		for _, entry := range resp.Topics {
			if strings.EqualFold(entry.Name, topic) {
				found = true
				break
			}
		}
		assert.True(t, found, "expected topic %q in list", topic)
	})
}

func Test_GetChannels(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetChannels(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Channels)
	})
}

func Test_GetSubscribers(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		_, channelId := CreateRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetSubscribers(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)

		subscribers := resp.Subscribers
		assert.Contains(t, subscribers, userId)
	})
}

func Test_GetSubscriptionStatus(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		_, channelId := CreateRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetSubscriptionStatus(ctx, userId, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.True(t, resp.IsSubscribed)
	})
}

func Test_GetSubscriptions(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetSubscriptions(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Subscriptions)
	})
}

func Test_MuteTopic(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		_, channelId := CreateRandomChannel(t, apiClient, userId)
		topic := createTopicWithMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.MuteTopic(ctx).
			ChannelId(channelId).
			Topic(topic).
			Op("add").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_Subscribe(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		desc := "Subscribed by test"
		resp, httpRes, err := apiClient.Subscribe(ctx).
			Subscriptions([]channels.SubscriptionRequest{{
				Name:        UniqueName("subscribe-channel"),
				Description: &desc,
			},
			}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.Equal(t, "success", resp.Result)
	})
}

func Test_Unsubscribe(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		channelName, _ := CreateRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.Unsubscribe(ctx).
			Subscriptions([]string{channelName}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)

		if len(resp.Removed) > 0 {
			assert.Contains(t, resp.Removed, channelName)
		}
	})
}

func Test_UpdateSubscriptionSettings(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		_, channelId := CreateRandomChannel(t, apiClient, userId)
		mute := true

		resp, httpRes, err := apiClient.UpdateSubscriptionSettings(ctx).
			SubscriptionData([]z.SubscriptionData{{
				ChannelId: channelId,
				Property:  z.SubscriptionPropertyIsMuted,
				Value:     z.SubscriptionDataValue{Bool: &mute},
			}}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_UpdateSubscriptions(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		channelName := UniqueName("update-subscriptions")
		desc := "Created in UpdateSubscriptions test"
		add := channels.SubscriptionRequestWithColor{
			Name:        &channelName,
			Description: &desc,
		}

		resp, httpRes, err := apiClient.UpdateSubscriptions(ctx).
			Add([]channels.SubscriptionRequestWithColor{add}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.Equal(t, "success", resp.Result)
	})
}

func Test_UpdateUserTopic(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userId := GetUserId(t, apiClient)

		_, channelId := CreateRandomChannel(t, apiClient, userId)
		topic := createTopicWithMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.UpdateUserTopic(ctx).
			ChannelId(channelId).
			Topic(topic).
			VisibilityPolicy(1).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func createTopicWithMessage(t *testing.T, apiClient client.Client, channelId int64) string {
	t.Helper()

	topic := UniqueName("topic")
	messageId := SendChannelMessage(t, apiClient, channelId, topic, fmt.Sprintf("message for %s", topic))
	assert.Greater(t, messageId, int64(0))
	return topic
}

func createChannelFolder(t *testing.T, apiClient client.Client, name, description string) int64 {
	t.Helper()

	resp, httpRes, err := apiClient.CreateChannelFolder(context.Background()).
		Name(name).
		Description(description).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)

	return resp.ChannelFolderId
}

func getChannelFolderIds(t *testing.T, apiClient client.Client) []int64 {
	t.Helper()

	resp, httpRes, err := apiClient.GetChannelFolders(context.Background()).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpRes)

	var ids []int64
	for _, folder := range resp.ChannelFolders {
		ids = append(ids, folder.Id)
	}
	return ids
}

func reverseSlice(values []int64) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}

func skipIfBigBlueButtonUnavailable(t *testing.T, err error) {
	t.Helper()

	var apiErr *z.APIError
	if !errors.As(err, &apiErr) {
		return
	}

	message := strings.ToLower(fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body())))
	if strings.Contains(message, "bigbluebutton") ||
		strings.Contains(message, "not configured") ||
		strings.Contains(message, "not implemented") {
		t.Skipf("BigBlueButton not available: %s", strings.TrimSpace(string(apiErr.Body())))
	}
}
