package zulip_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func Test_AuthenticationAPIService(t *testing.T) {
	t.Parallel()

	t.Run("DevFetchApiKey", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.DevFetchApiKey(ctx).Username(testAdminUsername).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("FetchApiKey", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		// More complex authentication flow not trivial to test here
		t.Skip("Not implemented yet")
		ctx := context.Background()
		apiClient.FetchApiKey(ctx).Execute()
	}))
}
