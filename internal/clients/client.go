// Package clients provides HTTP client abstractions and configurations for the Zulip API.
// It defines the core Client interface for making API calls, Response models, and configuration
// options for customizing client behavior including retry policies, logging, and user agents.
package clients

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client/statistics"
)

const retryIndefinitely = -1

// ResponseModel defines the interface that API Response models must implement.
// It provides information about parameters that were ignored during unmarshaling.
type ResponseModel interface {
	GetIgnoredParametersUnsupported() []string
}

// Client defines the interface for making authenticated API calls to the Zulip server.
// Implementations handle request execution, retry logic, and HTTP communication.
type Client interface {
	CallAPI(
		ctx context.Context,
		endpoint string,
		req *http.Request,
		model ResponseModel,
	) (httpResp *http.Response, err error)
	ServerURL() string
	GetUserAgent() string

	GetStatistics() statistics.Statistics
}

// Option is a functional option for configuring a Client.
type Option func(*Config)

// Config holds the configuration for a Zulip API client.
// It includes server credentials, HTTP client settings, logging configuration, and API version information.
type Config struct {
	RC *zulip.RC

	UserAgent       string
	HTTPClient      *http.Client
	Logger          *slog.Logger
	InsecureWarning bool

	ClientName string

	GatherStats bool

	MaxRetries int

	APISuffix    string
	APIVersion   string
	FeatureLevel int
}

// NewConfig creates and returns a new Config with default values and applied options.
// It builds the underlying HTTP client and sets the user agent. Returns an error if
// HTTP client initialization fails.
func NewConfig(rc *zulip.RC, opts ...Option) (Config, error) {
	if rc == nil {
		return Config{}, errors.New("invalid configuration: zulip.RC is nil")
	}

	cfg := Config{
		RC:              rc,
		ClientName:      defaultClientName,
		MaxRetries:      retryIndefinitely,
		InsecureWarning: true,
		GatherStats:     false,
		APISuffix:       defaultAPISuffix,
		APIVersion:      defaultAPIVersion,
		Logger:          slog.Default(),
	}

	for _, option := range opts {
		option(&cfg)
	}

	httpClient, userAgent, err := buildHTTPClient(
		cfg.HTTPClient,
		rc,
		cfg.ClientName,
		cfg.Logger,
		cfg.InsecureWarning,
	)
	if err != nil {
		return cfg, err
	}

	cfg.HTTPClient = httpClient
	cfg.UserAgent = userAgent
	return cfg, nil
}
