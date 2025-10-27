package clients

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client/statistics"
)

const (
	defaultClientName = "go-zulip/1.0.0"
	defaultAPISuffix  = "api"
	defaultAPIVersion = "v1"
)

// Client manages communication with the Zulip REST API API v1.0.0
// In most cases there should be only one, shared, Client.
type SimpleClient struct {
	Config

	Stats stats
}

var _ Client = (*SimpleClient)(nil)

// NewSimpleClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewSimpleClient(cfg Config) *SimpleClient {
	if cfg.MaxRetries != 0 {
		cfg.Logger.Warn("simple client does not support retries, but option was set")
	}
	return &SimpleClient{
		Config: cfg,
	}
}

func (c *SimpleClient) GetStatistics() statistics.Statistics {
	return c.Stats.GetStatistics()
}

func (c *SimpleClient) ServerURL() string {
	return fmt.Sprintf("%s/%s/%s", c.RC.Site, c.APISuffix, c.APIVersion)
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior.
func (c *SimpleClient) GetZulipRC() *zulip.RC {
	return c.RC
}

func (c *SimpleClient) CallAPI(
	ctx context.Context,
	endpoint string,
	req *http.Request,
	model ResponseModel,
) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if model == nil {
		return nil, errors.New("response model cannot be nil")
	}

	httpResp, err := c.doHTTPCall(ctx, endpoint, req)
	if err != nil || httpResp == nil {
		if c.GatherStats {
			c.Stats.Error(endpoint)
		}

		return httpResp, err
	}

	body, err := io.ReadAll(httpResp.Body)
	closeErr := httpResp.Body.Close()
	if closeErr != nil {
		c.Logger.ErrorContext(ctx, "failed to close Response body", "error", closeErr)
		return httpResp, closeErr
	}
	httpResp.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		if c.GatherStats {
			c.Stats.Error(endpoint)
		}

		return httpResp, err
	}

	if httpResp.StatusCode >= http.StatusMultipleChoices { // 300+
		if c.GatherStats {
			c.Stats.Error(endpoint)
		}

		return httpResp, unmarshallAPIError(ctx, c.Logger, httpResp.StatusCode, body)
	}

	err = decode(model, body, httpResp.Header.Get("Content-Type"))
	if err != nil {
		if c.GatherStats {
			c.Stats.Error(endpoint)
		}
		return httpResp, zulip.NewAPIError(body, err)
	}

	c.handleUnsupportedParameters(ctx, model.GetIgnoredParametersUnsupported())
	return httpResp, nil
}

func (c *SimpleClient) GetUserAgent() string {
	return c.UserAgent
}

// callAPI do the request.
func (c *SimpleClient) doHTTPCall(ctx context.Context, endpoint string, request *http.Request) (*http.Response, error) {
	debug := c.Logger.Enabled(ctx, slog.LevelDebug)

	if debug {
		dump, dumpErr := httputil.DumpRequestOut(request, true)
		if dumpErr != nil {
			c.Logger.ErrorContext(ctx, "failed to dump http Request to string", "error", dumpErr)
		}
		c.Logger.DebugContext(ctx, "HTTP Request", "dump", string(dump))
	}

	begin := time.Now()
	resp, err := c.HTTPClient.Do(request)
	end := time.Now()

	if c.GatherStats {
		duration := end.Sub(begin)
		c.Stats.Duration(endpoint, duration)
		c.Stats.Count(endpoint)
	}

	if debug && resp != nil {
		dump, dumpErr := httputil.DumpResponse(resp, true)
		if dumpErr != nil {
			c.Logger.ErrorContext(ctx, "failed to dump http Response to string", "error", dumpErr)
		}
		c.Logger.DebugContext(ctx, "HTTP Response", "dump", string(dump))
	}

	return resp, err
}

func (c *SimpleClient) handleUnsupportedParameters(ctx context.Context, parameters []string) {
	if len(parameters) > 0 {
		c.Logger.WarnContext(ctx, "Unsupported parameters were ignored", "parameters", parameters)
	}
}
