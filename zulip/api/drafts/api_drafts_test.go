package drafts_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/drafts"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_CreateDrafts(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		draft := createDraft(t, ctx, apiClient)
		assert.NotZero(t, draft.Id)
	})
}

func Test_CreateSavedSnippet(t *testing.T) {
	RequireFeatureLevel(t, 297)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		snippet := createSavedSnippet(t, ctx, apiClient)
		assert.NotZero(t, snippet.SavedSnippetId)
	})
}

func Test_DeleteDraft(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		draft := createDraft(t, ctx, apiClient)

		require.NotNil(t, draft.Id)

		resp := deleteDraft(t, ctx, apiClient, *draft.Id)
		assert.Equal(t, "success", resp.Result)

		draftsResp, _, err := apiClient.GetDrafts(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, draftsResp)

		for _, inner := range draftsResp.Drafts {
			require.NotNil(t, inner.Id)
			require.NotEqual(t, inner.Id, draft.Id, "Deleted draft still present")
		}
	})
}

func Test_DeleteSavedSnippet(t *testing.T) {
	RequireFeatureLevel(t, 297)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		snippet := createSavedSnippet(t, ctx, apiClient)

		resp := deleteSavedSnippet(t, ctx, apiClient, snippet.SavedSnippetId)
		assert.Equal(t, "success", resp.Result)

		snippetsResp, _, err := apiClient.GetSavedSnippets(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, snippetsResp)

		for _, s := range snippetsResp.SavedSnippets {
			require.NotEqual(t, snippet.SavedSnippetId, s.Id, "Deleted snippet still present")
		}
	})
}

func Test_EditDraft(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		draft := createDraft(t, ctx, apiClient)

		updatedDraft := *draft
		updatedDraft.Id = nil // ID is passed as path param, not in body
		updatedDraft.Topic = UniqueName("draft-topic")
		updatedDraft.Content = fmt.Sprintf("updated draft content %s", UniqueName("draft"))

		resp, httpResp, err := apiClient.EditDraft(ctx, *draft.Id).
			Draft(updatedDraft).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, "success", resp.Result)

		draftsResp, _, err := apiClient.GetDrafts(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, draftsResp)

		found := false
		for _, d := range draftsResp.Drafts {
			require.NotNil(t, d.Id)
			if *d.Id == *draft.Id {
				found = true
				assert.Equal(t, d.Topic, updatedDraft.Topic)
				assert.Equal(t, d.Content, updatedDraft.Content)
			}
		}
		assert.True(t, found, "Updated draft not found in draft list")
	})
}

func Test_EditSavedSnippet(t *testing.T) {
	RequireFeatureLevel(t, 368)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		snippet := createSavedSnippet(t, ctx, apiClient)

		newTitle := UniqueName("snippet-title-updated")
		newContent := fmt.Sprintf("updated content %s", UniqueName("snippet"))

		resp, httpResp, err := apiClient.EditSavedSnippet(ctx, snippet.SavedSnippetId).
			Title(newTitle).
			Content(newContent).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, "success", resp.Result)

		snippetsResp, _, err := apiClient.GetSavedSnippets(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, snippetsResp)

		found := false
		for _, s := range snippetsResp.SavedSnippets {
			if s.Id == snippet.SavedSnippetId {
				found = true
				assert.Equal(t, newTitle, s.Title)
				assert.Equal(t, newContent, s.Content)
			}
		}
		assert.True(t, found, "Updated saved snippet not found in list")
	})
}

func Test_GetDrafts(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		draft := createDraft(t, ctx, apiClient)

		resp, httpResp, err := apiClient.GetDrafts(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.GreaterOrEqual(t, len(resp.Drafts), 1)

		found := false
		for _, d := range resp.Drafts {
			require.NotNil(t, d.Id)
			if *d.Id == *draft.Id {
				found = true
				assert.Equal(t, d.Topic, draft.Topic)
				assert.Equal(t, d.Content, draft.Content)
				require.WithinDuration(t, time.Now(), d.Timestamp, 3*time.Minute)

			}
		}
		assert.True(t, found, "Created draft not found in list of drafts")
	})
}

func Test_GetSavedSnippets(t *testing.T) {
	RequireFeatureLevel(t, 297)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		snippet := createSavedSnippet(t, ctx, apiClient)

		resp, httpResp, err := apiClient.GetSavedSnippets(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.GreaterOrEqual(t, len(resp.SavedSnippets), 1)

		found := false
		for _, s := range resp.SavedSnippets {
			if s.Id == snippet.SavedSnippetId {
				found = true
				require.WithinDuration(t, time.Now(), s.DateCreated, 3*time.Minute)
			}
		}
		assert.True(t, found, "Created saved snippet not found in list")
	})
}

func createDraft(t *testing.T, ctx context.Context, apiClient client.Client) *z.Draft {
	t.Helper()

	userID := GetUserId(t, apiClient)
	_, channelID := CreateRandomChannel(t, apiClient, userID)

	draft := &z.Draft{
		Type:    z.RecipientTypeChannel,
		To:      z.ChannelAsRecipient(channelID),
		Topic:   UniqueName("draft-topic"),
		Content: fmt.Sprintf("draft content %s", UniqueName("draft")),
		// in the past to avoid any clock skew issues
		Timestamp: time.Now().Add(-1 * time.Minute),
	}

	resp, httpResp, err := apiClient.CreateDrafts(ctx).
		Drafts([]z.Draft{*draft}).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
	require.NotEmpty(t, resp.Ids)

	draftId := resp.Ids[0]
	require.Greater(t, draftId, int64(0))
	draft.Id = &draftId

	return draft
}

func deleteDraft(t *testing.T, ctx context.Context, apiClient client.Client, draftId int64) *z.Response {
	t.Helper()

	resp, httpResp, err := apiClient.DeleteDraft(ctx, draftId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp
}

func createSavedSnippet(t *testing.T, ctx context.Context, apiClient client.Client) *drafts.CreateSavedSnippetResponse {
	t.Helper()

	title := UniqueName("snippet-title")
	content := fmt.Sprintf("snippet content %s", UniqueName("snippet"))

	resp, httpResp, err := apiClient.CreateSavedSnippet(ctx).
		Title(title).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp
}

func deleteSavedSnippet(t *testing.T, ctx context.Context, apiClient client.Client, savedSnippetId int64) *z.Response {
	t.Helper()

	resp, httpResp, err := apiClient.DeleteSavedSnippet(ctx, savedSnippetId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp
}
