package zulip_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func Test_WebhooksAPIService(t *testing.T) {
	t.Parallel()

	t.Run("ZulipOutgoingWebhooks", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.ZulipOutgoingWebhooks(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, httpRes)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

	}))
}
