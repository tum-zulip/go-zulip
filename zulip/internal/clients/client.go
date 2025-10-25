package clients

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/tum-zulip/go-zulip/zulip/internal/utils"
	"github.com/tum-zulip/go-zulip/zulip/zuliprc"
)

const retryIndefinitely = -1

type ResponseModel interface {
	GetIgnoredParametersUnsupported() []string
}

type Client interface {
	CallAPI(ctx context.Context, endpoint string, req *http.Request, model ResponseModel) (httpResp *http.Response, err error)
	ServerURL() (string, error)
	GetUserAgent() string
}

type Option func(*Config)

type Config struct {
	RC *zuliprc.ZulipRC

	UserAgent       string
	HttpClient      *http.Client
	Logger          *slog.Logger
	InsecureWarning bool

	ClientName string

	GatherStats bool

	MaxRetries int

	ApiSuffix  string
	ApiVersion string
}

func NewConfig(rc *zuliprc.ZulipRC, opts ...Option) (Config, error) {
	cfg := Config{
		RC:              rc,
		ClientName:      defaultClientName,
		MaxRetries:      retryIndefinitely,
		InsecureWarning: true,
		GatherStats:     false,
		ApiSuffix:       defaultAPISuffix,
		ApiVersion:      defaultAPIVersion,
		Logger:          slog.Default(),
	}

	for _, option := range opts {
		option(&cfg)
	}

	httpClient, userAgent, err := utils.BuildHTTPClient(rc, cfg.ClientName, cfg.Logger, cfg.InsecureWarning)
	if err != nil {
		return cfg, err
	}

	cfg.HttpClient = httpClient
	cfg.UserAgent = userAgent
	return cfg, nil
}
