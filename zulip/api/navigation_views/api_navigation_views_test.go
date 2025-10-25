package navigation_views_test

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

func Test_AddNavigationView(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		view := createTestNavigationView(t, apiClient, true)

		ctx := context.Background()
		listResp, httpResp, err := apiClient.GetNavigationViews(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, listResp)
		RequireStatusOK(t, httpResp)

		found := navigationViewByFragment(listResp.NavigationViews, view.fragment)
		require.NotNil(t, found, "expected navigation view with fragment %s", view.fragment)
		require.NotNil(t, found.Name, "expected navigation view name to be set")
		assert.Equal(t, view.name, *found.Name)
		assert.Equal(t, view.isPinned, found.IsPinned)

	})
}

func Test_EditNavigationView(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		view := createTestNavigationView(t, apiClient, true)
		updatedName := fmt.Sprintf("Updated Navigation View %s", UniqueName("nav"))

		resp, httpResp, err := apiClient.EditNavigationView(ctx, view.fragment).
			IsPinned(false).
			Name(updatedName).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)

		listResp, listHTTPRes, err := apiClient.GetNavigationViews(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, listResp)
		RequireStatusOK(t, listHTTPRes)

		updated := navigationViewByFragment(listResp.NavigationViews, view.fragment)
		require.NotNil(t, updated, "expected navigation view with fragment %s", view.fragment)
		require.NotNil(t, updated.Name, "expected navigation view name to be set")
		assert.Equal(t, updatedName, *updated.Name)
		assert.False(t, updated.IsPinned)

	})
}

func Test_GetNavigationViews(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		view := createTestNavigationView(t, apiClient, false)

		resp, httpResp, err := apiClient.GetNavigationViews(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)

		assert.NotNil(t, navigationViewByFragment(resp.NavigationViews, view.fragment))

	})
}

func Test_RemoveNavigationView(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		view := createTestNavigationView(t, apiClient, true)

		resp, httpResp, err := apiClient.RemoveNavigationView(ctx, view.fragment).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)

		listResp, listHTTPRes, err := apiClient.GetNavigationViews(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, listResp)
		RequireStatusOK(t, listHTTPRes)

		assert.Nil(t, navigationViewByFragment(listResp.NavigationViews, view.fragment))

	})
}

type navigationViewState struct {
	fragment string
	name     string
	isPinned bool
}

func createTestNavigationView(t *testing.T, apiClient client.Client, isPinned bool) navigationViewState {
	t.Helper()

	ctx := context.Background()
	fragment := fmt.Sprintf("narrow/test/%s", UniqueName("nav"))
	name := fmt.Sprintf("Navigation View %s", UniqueName("nav"))

	resp, httpResp, err := apiClient.AddNavigationView(ctx).
		Fragment(fragment).
		IsPinned(isPinned).
		Name(name).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return navigationViewState{fragment: fragment, name: name, isPinned: isPinned}
}

func navigationViewByFragment(views []z.NavigationView, fragment string) *z.NavigationView {
	for i := range views {
		if views[i].Fragment == fragment {
			return &views[i]
		}
	}
	return nil
}
