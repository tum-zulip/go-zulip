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
		draft := createDraft(ctx, t, apiClient)
		assert.NotZero(t, draft.ID)
	})
}

func Test_CreateSavedSnippet(t *testing.T) {
	RequireFeatureLevel(t, 297)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		snippet := createSavedSnippet(ctx, t, apiClient)
		assert.NotZero(t, snippet.SavedSnippetID)
	})
}

func Test_DeleteDraft(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		draft := createDraft(ctx, t, apiClient)

		require.NotNil(t, draft.ID)

		resp := deleteDraft(ctx, t, apiClient, *draft.ID)
		assert.Equal(t, "success", resp.Result)

		draftsResp, _, err := apiClient.GetDrafts(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, draftsResp)

		for _, inner := range draftsResp.Drafts {
			require.NotNil(t, inner.ID)
			require.NotEqual(t, inner.ID, draft.ID, "Deleted draft still present")
		}
	})
}

func Test_DeleteSavedSnippet(t *testing.T) {
	RequireFeatureLevel(t, 297)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		snippet := createSavedSnippet(ctx, t, apiClient)

		resp := deleteSavedSnippet(ctx, t, apiClient, snippet.SavedSnippetID)
		assert.Equal(t, "success", resp.Result)

		snippetsResp, _, err := apiClient.GetSavedSnippets(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, snippetsResp)

		for _, s := range snippetsResp.SavedSnippets {
			require.NotEqual(t, snippet.SavedSnippetID, s.ID, "Deleted snippet still present")
		}
	})
}

func Test_EditDraft(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		draft := createDraft(ctx, t, apiClient)

		updatedDraft := *draft
		updatedDraft.ID = nil // ID is passed as path param, not in body
		updatedDraft.Topic = UniqueName("draft-topic")
		updatedDraft.Content = fmt.Sprintf("updated draft content %s", UniqueName("draft"))

		resp, httpResp, err := apiClient.EditDraft(ctx, *draft.ID).
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
			require.NotNil(t, d.ID)
			if *d.ID == *draft.ID {
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
		snippet := createSavedSnippet(ctx, t, apiClient)

		newTitle := UniqueName("snippet-title-updated")
		newContent := fmt.Sprintf("updated content %s", UniqueName("snippet"))

		resp, httpResp, err := apiClient.EditSavedSnippet(ctx, snippet.SavedSnippetID).
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
			if s.ID == snippet.SavedSnippetID {
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
		draft := createDraft(ctx, t, apiClient)

		resp, httpResp, err := apiClient.GetDrafts(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.GreaterOrEqual(t, len(resp.Drafts), 1)

		found := false
		for _, d := range resp.Drafts {
			require.NotNil(t, d.ID)
			if *d.ID == *draft.ID {
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
		snippet := createSavedSnippet(ctx, t, apiClient)

		resp, httpResp, err := apiClient.GetSavedSnippets(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.GreaterOrEqual(t, len(resp.SavedSnippets), 1)

		found := false
		for _, s := range resp.SavedSnippets {
			if s.ID == snippet.SavedSnippetID {
				found = true
				require.WithinDuration(t, time.Now(), s.DateCreated, 3*time.Minute)
			}
		}
		assert.True(t, found, "Created saved snippet not found in list")
	})
}

func createDraft(ctx context.Context, t *testing.T, apiClient client.Client) *z.Draft {
	t.Helper()

	userID := GetUserID(t, apiClient)
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
	require.NotEmpty(t, resp.IDs)

	draftID := resp.IDs[0]
	require.Positive(t, draftID)
	draft.ID = &draftID

	return draft
}

func deleteDraft(ctx context.Context, t *testing.T, apiClient client.Client, draftID int64) *z.Response {
	t.Helper()

	resp, httpResp, err := apiClient.DeleteDraft(ctx, draftID).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp
}

func createSavedSnippet(ctx context.Context, t *testing.T, apiClient client.Client) *drafts.CreateSavedSnippetResponse {
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

func deleteSavedSnippet(ctx context.Context, t *testing.T, apiClient client.Client, savedSnippetID int64) *z.Response {
	t.Helper()

	resp, httpResp, err := apiClient.DeleteSavedSnippet(ctx, savedSnippetID).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp
}
