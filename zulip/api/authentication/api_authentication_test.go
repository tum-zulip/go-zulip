package authentication_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_AuthenticationAPIService(t *testing.T) {
	t.Parallel()

	t.Run("DevFetchApiKey", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.DevFetchApiKey(ctx).Username(TestAdminUsername).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("FetchApiKey", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		// More complex authentication flow not trivial to test here
		t.Skip("Not implemented yet")
		ctx := context.Background()
		apiClient.FetchApiKey(ctx).Execute()
	}))
}
