package authentication_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_DevFetchAPIKey(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, _, err := apiClient.DevFetchAPIKey(ctx).Username(TestAdminUsername).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_FetchAPIKey(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		// More complex authentication flow not trivial to test here
		t.Skip("Not implemented yet")
		ctx := context.Background()
		resp, _, err := apiClient.FetchAPIKey(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}
