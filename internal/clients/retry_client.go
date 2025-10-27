package clients

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/tum-zulip/go-zulip/zulip"
)

// ErrMaxRetriesReached is returned when the maximum number of retries is reached.
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

func (c *RetryClient) CallAPI(
	ctx context.Context,
	endpoint string,
	req *http.Request,
	model ResponseModel,
) (*http.Response, error) {
	retryForever := c.MaxRetries == retryIndefinitely

	var httpResp *http.Response
	var err error
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

		return httpResp, err
	}
	c.Logger.DebugContext(ctx, "hit max retires", "max-retires", c.MaxRetries)
	return httpResp, ErrMaxRetriesReached
}

func getRateLimit(resp *http.Response, err error) (time.Duration, bool) {
	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		var rateLimitedError zulip.RateLimitedError
		if errors.As(err, &rateLimitedError) {
			return rateLimitedError.RetryAfter, true
		}
	}
	return 0, false
}
