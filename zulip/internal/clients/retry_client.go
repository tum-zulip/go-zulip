package clients

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/tum-zulip/go-zulip/zulip"
)

// Will not be returned if max-retires is set to zero
var ErrMaxRetriesReached = errors.New("hit max retires")

type RetryClient struct {
	SimpleClient
}

var _ Client = (*RetryClient)(nil)

func NewRetryClient(cfg Config) RetryClient {
	return RetryClient{
		SimpleClient: SimpleClient{Config: cfg},
	}
}

func (c *RetryClient) CallAPI(ctx context.Context, endpoint string, req *http.Request, model ResponseModel) (httpResp *http.Response, err error) {
	retryForever := c.MaxRetries == retryIndefinitely

	for i := 0; retryForever || i <= c.MaxRetries; i++ {

		httpResp, err = c.SimpleClient.CallAPI(ctx, endpoint, req, model)

		if retryAfter, ok := getRateLimit(httpResp, err); ok {
			c.Logger.WarnContext(ctx, "hit API rate-limit", "retry-after", retryAfter, "attempt", i+1)
			time.Sleep(retryAfter)

			if c.GatherStats {
				c.Stats.Retry(endpoint)
				c.Stats.Duration("rate-limit-timeout", retryAfter)
			}
			continue
		}

		return
	}
	c.Logger.DebugContext(ctx, "hit max retires", "max-retires", c.MaxRetries)
	return httpResp, ErrMaxRetriesReached
}

func getRateLimit(resp *http.Response, err error) (time.Duration, bool) {
	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		if apiErr, ok := err.(*zulip.APIError); ok {
			if rateLimitedError, ok := apiErr.Model().(zulip.RateLimitedError); ok {
				return rateLimitedError.RetryAfter, true
			}
		}
	}
	return 0, false
}
