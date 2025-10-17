package zulip_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func Test_NavigationViewsAPIService(t *testing.T) {
	t.Parallel()

	t.Run("AddNavigationView", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		view := createTestNavigationView(t, apiClient, true)

		ctx := context.Background()
		listResp, httpRes, err := apiClient.GetNavigationViews(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, listResp)
		requireStatusOK(t, httpRes)

		found := navigationViewByFragment(listResp.NavigationViews, view.fragment)
		require.NotNil(t, found, "expected navigation view with fragment %s", view.fragment)
		require.NotNil(t, found.Name, "expected navigation view name to be set")
		assert.Equal(t, view.name, *found.Name)
		assert.Equal(t, view.isPinned, found.IsPinned)

	}))

	t.Run("EditNavigationView", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		view := createTestNavigationView(t, apiClient, true)
		updatedName := fmt.Sprintf("Updated Navigation View %s", uniqueName("nav"))

		resp, httpRes, err := apiClient.EditNavigationView(ctx, view.fragment).
			IsPinned(false).
			Name(updatedName).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)

		listResp, listHTTPRes, err := apiClient.GetNavigationViews(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, listResp)
		requireStatusOK(t, listHTTPRes)

		updated := navigationViewByFragment(listResp.NavigationViews, view.fragment)
		require.NotNil(t, updated, "expected navigation view with fragment %s", view.fragment)
		require.NotNil(t, updated.Name, "expected navigation view name to be set")
		assert.Equal(t, updatedName, *updated.Name)
		assert.False(t, updated.IsPinned)

	}))

	t.Run("GetNavigationViews", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		view := createTestNavigationView(t, apiClient, false)

		resp, httpRes, err := apiClient.GetNavigationViews(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)

		assert.NotNil(t, navigationViewByFragment(resp.NavigationViews, view.fragment))

	}))

	t.Run("RemoveNavigationView", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		view := createTestNavigationView(t, apiClient, true)

		resp, httpRes, err := apiClient.RemoveNavigationView(ctx, view.fragment).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)

		listResp, listHTTPRes, err := apiClient.GetNavigationViews(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, listResp)
		requireStatusOK(t, listHTTPRes)

		assert.Nil(t, navigationViewByFragment(listResp.NavigationViews, view.fragment))

	}))
}

type navigationViewState struct {
	fragment string
	name     string
	isPinned bool
}

func createTestNavigationView(t *testing.T, apiClient zulip.Client, isPinned bool) navigationViewState {
	t.Helper()

	ctx := context.Background()
	fragment := fmt.Sprintf("narrow/test/%s", uniqueName("nav"))
	name := fmt.Sprintf("Navigation View %s", uniqueName("nav"))

	resp, httpRes, err := apiClient.AddNavigationView(ctx).
		Fragment(fragment).
		IsPinned(isPinned).
		Name(name).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	requireStatusOK(t, httpRes)

	return navigationViewState{fragment: fragment, name: name, isPinned: isPinned}
}

func navigationViewByFragment(views []zulip.NavigationView, fragment string) *zulip.NavigationView {
	for i := range views {
		if views[i].Fragment == fragment {
			return &views[i]
		}
	}
	return nil
}
