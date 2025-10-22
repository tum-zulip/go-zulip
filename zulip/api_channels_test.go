package zulip_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func Test_ChannelsAPIService(t *testing.T) {
	t.Parallel()

	ownerClient := GetOwnerClient(t)
	ownerId := getOwnUserId(t, ownerClient)

	t.Run("AddDefaultChannel", func(t *testing.T) {
		ctx := context.Background()

		_, channelId := createRandomChannel(t, ownerClient, ownerId)

		resp, httpRes, err := ownerClient.AddDefaultChannel(ctx).ChannelId(channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	})

	t.Run("ArchiveChannel", func(t *testing.T) {
		ctx := context.Background()

		_, channelId := createRandomChannel(t, ownerClient, ownerId)

		resp, httpRes, err := ownerClient.ArchiveChannel(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	})

	t.Run("CreateChannelFolder", func(t *testing.T) {
		ctx := context.Background()

		resp, httpRes, err := ownerClient.CreateChannelFolder(ctx).
			Name(uniqueName("test-folder")).
			Description("Created during Channels API tests").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Greater(t, resp.ChannelFolderId, int64(0))
	})

	t.Run("DeleteTopic", func(t *testing.T) {
		ctx := context.Background()

		_, channelId := createRandomChannel(t, ownerClient, ownerId)
		topic := createTopicWithMessage(t, ownerClient, channelId)

		resp, httpRes, err := ownerClient.DeleteTopic(ctx, channelId).
			TopicName(topic).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.True(t, resp.Complete)
	})

	t.Run("PatchChannelFolders", func(t *testing.T) {
		ctx := context.Background()

		createChannelFolder(t, ownerClient, uniqueName("patch-folder"), "first test folder")
		createChannelFolder(t, ownerClient, uniqueName("patch-folder"), "second test folder")

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
		requireStatusOK(t, httpRes)
	})

	t.Run("RemoveDefaultChannel", func(t *testing.T) {
		ctx := context.Background()

		_, channelId := createRandomChannel(t, ownerClient, ownerId)
		reqResp, _, err := ownerClient.AddDefaultChannel(ctx).ChannelId(channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, reqResp)

		resp, httpRes, err := ownerClient.RemoveDefaultChannel(ctx).ChannelId(channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	})

	t.Run("UpdateChannelFolder", func(t *testing.T) {
		ctx := context.Background()

		folderId := createChannelFolder(t, ownerClient, uniqueName("update-folder"), "initial description")

		resp, httpRes, err := ownerClient.UpdateChannelFolder(ctx, folderId).
			Description("updated folder description").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	})

	t.Run("UpdateChannel", func(t *testing.T) {
		ctx := context.Background()

		_, channelId := createRandomChannel(t, ownerClient, ownerId)

		resp, httpRes, err := ownerClient.UpdateChannel(ctx, channelId).
			Description("updated by test").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	})

	t.Run("CreateBigBlueButtonVideoCall", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.CreateBigBlueButtonVideoCall(ctx).
			MeetingName(uniqueName("bbb-meeting")).
			Execute()
		if err != nil {
			skipIfBigBlueButtonUnavailable(t, err)
		}
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Url)
	}))

	t.Run("CreateChannel", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		userId := getOwnUserId(t, apiClient)

		createRandomChannel(t, apiClient, userId)
	}))

	t.Run("GetChannelFolders", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetChannelFolders(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("GetChannelById", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		_, channelId := createRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetChannelById(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Equal(t, channelId, resp.Channel.ChannelId)
	}))

	t.Run("GetChannelEmailAddress", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		_, channelId := createRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetChannelEmailAddress(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Email)
	}))

	t.Run("GetChannelId", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		channelName, channelId := createRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetChannelId(ctx).
			Channel(channelName).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Equal(t, channelId, resp.ChannelId)
	}))

	t.Run("GetChannelTopics", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		_, channelId := createRandomChannel(t, apiClient, userId)
		topic := createTopicWithMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.GetChannelTopics(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)

		found := false
		for _, entry := range resp.Topics {
			if strings.EqualFold(entry.Name, topic) {
				found = true
				break
			}
		}
		assert.True(t, found, "expected topic %q in list", topic)
	}))

	t.Run("GetChannels", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetChannels(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Channels)
	}))

	t.Run("GetSubscribers", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		_, channelId := createRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetSubscribers(ctx, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)

		subscribers := resp.Subscribers
		assert.Contains(t, subscribers, userId)
	}))

	t.Run("GetSubscriptionStatus", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		_, channelId := createRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.GetSubscriptionStatus(ctx, userId, channelId).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.True(t, resp.IsSubscribed)
	}))

	t.Run("GetSubscriptions", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.GetSubscriptions(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Subscriptions)
	}))

	t.Run("MuteTopic", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		_, channelId := createRandomChannel(t, apiClient, userId)
		topic := createTopicWithMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.MuteTopic(ctx).
			ChannelId(channelId).
			Topic(topic).
			Op("add").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("Subscribe", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		desc := "Subscribed by test"
		resp, httpRes, err := apiClient.Subscribe(ctx).
			Subscriptions([]z.SubscriptionRequest{{
				Name:        uniqueName("subscribe-channel"),
				Description: &desc,
			},
			}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Equal(t, "success", resp.Result)
	}))

	t.Run("Unsubscribe", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		channelName, _ := createRandomChannel(t, apiClient, userId)

		resp, httpRes, err := apiClient.Unsubscribe(ctx).
			Subscriptions([]string{channelName}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)

		if len(resp.Removed) > 0 {
			assert.Contains(t, resp.Removed, channelName)
		}
	}))

	t.Run("UpdateSubscriptionSettings", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		_, channelId := createRandomChannel(t, apiClient, userId)
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
		requireStatusOK(t, httpRes)
	}))

	t.Run("UpdateSubscriptions", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		channelName := uniqueName("update-subscriptions")
		desc := "Created in UpdateSubscriptions test"
		add := z.SubscriptionRequestWithColor{
			Name:        &channelName,
			Description: &desc,
		}

		resp, httpRes, err := apiClient.UpdateSubscriptions(ctx).
			Add([]z.SubscriptionRequestWithColor{add}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Equal(t, "success", resp.Result)
	}))

	t.Run("UpdateUserTopic", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()
		userId := getOwnUserId(t, apiClient)

		_, channelId := createRandomChannel(t, apiClient, userId)
		topic := createTopicWithMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.UpdateUserTopic(ctx).
			ChannelId(channelId).
			Topic(topic).
			VisibilityPolicy(1).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))
}

func createRandomChannel(t *testing.T, apiClient z.Client, subscribers ...int64) (string, int64) {
	t.Helper()

	subs := append([]int64(nil), subscribers...)
	if len(subs) == 0 {
		resp, httpRes, err := apiClient.GetOwnUser(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		subs = []int64{resp.UserId}
	}

	name := uniqueName("test-channel")
	resp, httpRes, err := apiClient.CreateChannel(context.Background()).
		Name(name).
		Description("Created by channel API tests").
		Subscribers(subs).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	requireStatusOK(t, httpRes)

	return name, resp.Id
}

func sendChannelMessage(t *testing.T, apiClient z.Client, channelId int64, topic, content string) int64 {
	t.Helper()

	resp, httpResp, err := apiClient.SendMessage(context.Background()).
		RecipientType(z.RecipientTypeChannel).
		To(z.ChannelAsRecipient(channelId)).
		Topic(topic).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	requireStatusOK(t, httpResp)

	return resp.Id
}

func createTopicWithMessage(t *testing.T, apiClient z.Client, channelId int64) string {
	t.Helper()

	topic := uniqueName("topic")
	messageId := sendChannelMessage(t, apiClient, channelId, topic, fmt.Sprintf("message for %s", topic))
	assert.Greater(t, messageId, int64(0))
	return topic
}

func createChannelFolder(t *testing.T, apiClient z.Client, name, description string) int64 {
	t.Helper()

	resp, httpRes, err := apiClient.CreateChannelFolder(context.Background()).
		Name(name).
		Description(description).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	requireStatusOK(t, httpRes)

	return resp.ChannelFolderId
}

func getChannelFolderIds(t *testing.T, apiClient z.Client) []int64 {
	t.Helper()

	resp, httpRes, err := apiClient.GetChannelFolders(context.Background()).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	requireStatusOK(t, httpRes)

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

func requireStatusOK(t *testing.T, httpRes *http.Response) {
	t.Helper()
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func createChannelWithAllClients(t *testing.T) (string, int64) {
	t.Helper()

	clientFactories := []func(*testing.T) z.Client{
		GetOwnerClient,
		GetAdminClient,
		GetModeratorClient,
		GetNormalClient,
		GetOtherNormalClient,
	}

	allClientIds := make([]int64, 0, len(clientFactories))
	for _, factory := range clientFactories {
		apiClient := factory(t)
		userId := getOwnUserId(t, apiClient)
		allClientIds = append(allClientIds, userId)
	}

	return createRandomChannel(t, GetAdminClient(t), allClientIds...)
}
