package clients

import (
	"context"
	"net/http"
	"time"
)

type TestClient struct {
	RetryClient
}

var _ Client = (*TestClient)(nil)

func NewTestClient(cfg Config) *TestClient {
	return &TestClient{
		RetryClient: *NewRetryClient(cfg),
	}
}

func (c *TestClient) CallAPI(
	ctx context.Context,
	endpoint string,
	req *http.Request,
	model ResponseModel,
) (*http.Response, error) {
	const maxRetries = 5
	const waitDuration = 100 * time.Millisecond

	var httpResp *http.Response
	var err error
	for range maxRetries {
		httpResp, err = c.RetryClient.CallAPI(ctx, endpoint, req, model)
		if err == nil && httpResp != nil && httpResp.StatusCode < http.StatusOK {
			return httpResp, nil
		}

		time.Sleep(waitDuration)
		c.RetryClient.SimpleClient.Stats.Retry(endpoint)
		c.RetryClient.SimpleClient.Stats.Duration("test-retry-timeout", waitDuration)
	}
	return httpResp, err
}
