package channels_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/channels"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_AddDefaultChannel(t *testing.T) {
	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerID := GetUserID(t, ownerClient)

	_, channelID := CreateRandomChannel(t, ownerClient, ownerID)

	resp, httpResp, err := ownerClient.AddDefaultChannel(ctx).ChannelID(channelID).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
}

func Test_ArchiveChannel(t *testing.T) {
	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerID := GetUserID(t, ownerClient)

	_, channelID := CreateRandomChannel(t, ownerClient, ownerID)

	resp, httpResp, err := ownerClient.ArchiveChannel(ctx, channelID).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
}

func Test_CreateChannelFolder(t *testing.T) {
	RequireFeatureLevel(t, 389)

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)

	resp, httpResp, err := ownerClient.CreateChannelFolder(ctx).
		Name(UniqueName("test-folder")).
		Description("Created during Channels API tests").
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
	assert.Positive(t, resp.ChannelFolderID)
}

func Test_DeleteTopic(t *testing.T) {
	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerID := GetUserID(t, ownerClient)

	_, channelID := CreateRandomChannel(t, ownerClient, ownerID)
	topic := createTopicWithMessage(t, ownerClient, channelID)

	resp, httpResp, err := ownerClient.DeleteTopic(ctx, channelID).
		TopicName(topic).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
	assert.True(t, resp.Complete)
}

func Test_PatchChannelFolders(t *testing.T) {
	RequireFeatureLevel(t, 414)

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)

	createChannelFolder(t, ownerClient, UniqueName("patch-folder"), "first test folder")
	createChannelFolder(t, ownerClient, UniqueName("patch-folder"), "second test folder")

	originalOrder := getChannelFolderIDs(t, ownerClient)
	if len(originalOrder) < 2 {
		t.Skip("need at least two channel folders to reorder")
	}

	reordered := append([]int64(nil), originalOrder...)
	reverseSlice(reordered)

	resp, httpResp, err := ownerClient.PatchChannelFolders(ctx).
		Order(reordered).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
}

func Test_RemoveDefaultChannel(t *testing.T) {
	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerID := GetUserID(t, ownerClient)

	_, channelID := CreateRandomChannel(t, ownerClient, ownerID)
	reqResp, _, err := ownerClient.AddDefaultChannel(ctx).ChannelID(channelID).Execute()
	require.NoError(t, err)
	require.NotNil(t, reqResp)

	resp, httpResp, err := ownerClient.RemoveDefaultChannel(ctx).ChannelID(channelID).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
}

func Test_UpdateChannelFolder(t *testing.T) {
	RequireFeatureLevel(t, 389)

	ctx := context.Background()
	ownerClient := GetOwnerClient(t)

	folderID := createChannelFolder(t, ownerClient, UniqueName("update-folder"), "initial description")

	resp, httpResp, err := ownerClient.UpdateChannelFolder(ctx, folderID).
		Description("updated folder description").
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
}

func Test_UpdateChannel(t *testing.T) {
	ctx := context.Background()
	ownerClient := GetOwnerClient(t)
	ownerID := GetUserID(t, ownerClient)

	_, channelID := CreateRandomChannel(t, ownerClient, ownerID)

	resp, httpResp, err := ownerClient.UpdateChannel(ctx, channelID).
		Description("updated by test").
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
}

func Test_CreateBigBlueButtonVideoCall(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.CreateBigBlueButtonVideoCall(ctx).
			MeetingName(UniqueName("bbb-meeting")).
			Execute()
		if err != nil {
			skipIfBigBlueButtonUnavailable(t, err)
		}
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.NotEmpty(t, resp.Url)
	})
}

func Test_CreateChannel(t *testing.T) {
	RequireFeatureLevel(t, 417)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		userID := GetUserID(t, apiClient)

		resp, httpResp, err := apiClient.CreateChannel(context.Background()).
			Name(UniqueName("test-channel")).
			Description("Created by channel API tests").
			Subscribers([]int64{userID}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
	})
}

func Test_GetChannelFolders(t *testing.T) {
	RequireFeatureLevel(t, 389)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetChannelFolders(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
	})
}

func Test_GetChannelByID(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		_, channelID := CreateRandomChannel(t, apiClient, userID)

		resp, httpResp, err := apiClient.GetChannelByID(ctx, channelID).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, channelID, resp.Channel.ChannelID)
		require.WithinDuration(t, time.Now(), resp.Channel.DateCreated, 3*time.Minute)
	})
}

func Test_GetChannelEmailAddress(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		_, channelID := CreateRandomChannel(t, apiClient, userID)

		resp, httpResp, err := apiClient.GetChannelEmailAddress(ctx, channelID).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.NotEmpty(t, resp.Email)
	})
}

func Test_GetChannelID(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		channelName, channelID := CreateRandomChannel(t, apiClient, userID)

		resp, httpResp, err := apiClient.GetChannelID(ctx).
			Channel(channelName).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, channelID, resp.ChannelID)
	})
}

func Test_GetChannelTopics(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		_, channelID := CreateRandomChannel(t, apiClient, userID)
		topic := createTopicWithMessage(t, apiClient, channelID)

		resp, httpResp, err := apiClient.GetChannelTopics(ctx, channelID).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)

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
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetChannels(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.NotEmpty(t, resp.Channels)
	})
}

func Test_GetSubscribers(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		_, channelID := CreateRandomChannel(t, apiClient, userID)

		resp, httpResp, err := apiClient.GetSubscribers(ctx, channelID).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)

		subscribers := resp.Subscribers
		assert.Contains(t, subscribers, userID)
	})
}

func Test_GetSubscriptionStatus(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		_, channelID := CreateRandomChannel(t, apiClient, userID)

		resp, httpResp, err := apiClient.GetSubscriptionStatus(ctx, userID, channelID).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.True(t, resp.IsSubscribed)
	})
}

func Test_GetSubscriptions(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetSubscriptions(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.NotEmpty(t, resp.Subscriptions)
	})
}

func Test_MuteTopic(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		_, channelID := CreateRandomChannel(t, apiClient, userID)
		topic := createTopicWithMessage(t, apiClient, channelID)

		resp, httpResp, err := apiClient.MuteTopic(ctx).
			ChannelID(channelID).
			Topic(topic).
			Op("add").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
	})
}

func Test_Subscribe(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		desc := "Subscribed by test"
		resp, httpResp, err := apiClient.Subscribe(ctx).
			Subscriptions([]channels.SubscriptionRequest{{
				Name:        UniqueName("subscribe-channel"),
				Description: &desc,
			}}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, "success", resp.Result)
	})
}

func Test_Unsubscribe(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		channelName, _ := CreateRandomChannel(t, apiClient, userID)

		resp, httpResp, err := apiClient.Unsubscribe(ctx).
			Subscriptions([]string{channelName}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)

		if len(resp.Removed) > 0 {
			assert.Contains(t, resp.Removed, channelName)
		}
	})
}

func Test_UpdateSubscriptionSettings(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		_, channelID := CreateRandomChannel(t, apiClient, userID)
		mute := true

		resp, httpResp, err := apiClient.UpdateSubscriptionSettings(ctx).
			SubscriptionData([]z.SubscriptionData{{
				ChannelID: channelID,
				Property:  z.SubscriptionPropertyIsMuted,
				Value:     z.SubscriptionDataValue{Bool: &mute},
			}}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
	})
}

func Test_UpdateSubscriptions(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		channelName := UniqueName("update-subscriptions")
		desc := "Created in UpdateSubscriptions test"
		add := channels.SubscriptionRequestWithColor{
			Name:        &channelName,
			Description: &desc,
		}

		resp, httpResp, err := apiClient.UpdateSubscriptions(ctx).
			Add([]channels.SubscriptionRequestWithColor{add}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, "success", resp.Result)
	})
}

func Test_UpdateUserTopic(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		userID := GetUserID(t, apiClient)

		_, channelID := CreateRandomChannel(t, apiClient, userID)
		topic := createTopicWithMessage(t, apiClient, channelID)

		resp, httpResp, err := apiClient.UpdateUserTopic(ctx).
			ChannelID(channelID).
			Topic(topic).
			VisibilityPolicy(1).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
	})
}

func createTopicWithMessage(t *testing.T, apiClient client.Client, channelID int64) string {
	t.Helper()

	topic := UniqueName("topic")
	messageID := SendChannelMessage(t, apiClient, channelID, topic, fmt.Sprintf("message for %s", topic))
	assert.Positive(t, messageID)
	return topic
}

func createChannelFolder(t *testing.T, apiClient client.Client, name, description string) int64 {
	t.Helper()

	resp, httpResp, err := apiClient.CreateChannelFolder(context.Background()).
		Name(name).
		Description(description).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp.ChannelFolderID
}

func getChannelFolderIDs(t *testing.T, apiClient client.Client) []int64 {
	t.Helper()

	resp, httpResp, err := apiClient.GetChannelFolders(context.Background()).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	var ids []int64
	for _, folder := range resp.ChannelFolders {
		ids = append(ids, folder.ID)
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
