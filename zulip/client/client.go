package client

import (
	"log/slog"
	"net/http"

	"github.com/tum-zulip/go-zulip/zulip/api/authentication"
	"github.com/tum-zulip/go-zulip/zulip/api/channels"
	"github.com/tum-zulip/go-zulip/zulip/api/drafts"
	"github.com/tum-zulip/go-zulip/zulip/api/invites"
	"github.com/tum-zulip/go-zulip/zulip/api/messages"
	"github.com/tum-zulip/go-zulip/zulip/api/mobile"
	"github.com/tum-zulip/go-zulip/zulip/api/navigation_views"
	"github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/api/reminders"
	"github.com/tum-zulip/go-zulip/zulip/api/scheduled_messages"
	"github.com/tum-zulip/go-zulip/zulip/api/server_and_organizations"
	"github.com/tum-zulip/go-zulip/zulip/api/users"
	"github.com/tum-zulip/go-zulip/zulip/client/statistics"
	"github.com/tum-zulip/go-zulip/zulip/internal/clients"

	"github.com/tum-zulip/go-zulip/zulip/zuliprc"
)

type Client interface {
	authentication.APIAuthentication
	channels.APIChannels
	drafts.APIDrafts
	invites.APIInvites
	messages.APIMessages
	mobile.APIMobile
	navigation_views.APINavigationViews
	real_time_events.APIRealTimeEvents
	reminders.APIReminders
	scheduled_messages.APIScheduledMessages
	server_and_organizations.APIServerAndOrganizations
	users.APIUsers

	GetStatistics() map[string]statistics.Statistic
}

var _ Client = (*client)(nil)

type client struct {
	apiClient clients.RetryClient

	authentication.APIAuthentication
	channels.APIChannels
	drafts.APIDrafts
	invites.APIInvites
	messages.APIMessages
	mobile.APIMobile
	navigation_views.APINavigationViews
	real_time_events.APIRealTimeEvents
	reminders.APIReminders
	scheduled_messages.APIScheduledMessages
	server_and_organizations.APIServerAndOrganizations
	users.APIUsers
}

func NewClient(zuliprc *zuliprc.ZulipRC, opts ...clients.Option) (*client, error) {

	cfg, err := clients.NewConfig(zuliprc, opts...)
	if err != nil {
		return nil, err
	}

	c := &client{
		apiClient: clients.NewRetryClient(cfg),
	}
	c.initializeFromClient(&c.apiClient)

	return c, nil
}

func (c *client) GetStatistics() map[string]statistics.Statistic {
	return c.apiClient.Stats.GetStatistics()
}

func WithAPISuffix(suffix string) clients.Option {
	return func(cfg *clients.Config) {
		cfg.ApiSuffix = suffix
	}
}

func WithAPIVersion(version string) clients.Option {
	return func(cfg *clients.Config) {
		cfg.ApiVersion = version
	}
}

func WithLogger(logger *slog.Logger) clients.Option {
	return func(cfg *clients.Config) {
		if logger == nil {
			cfg.Logger = slog.Default()
			return
		}
		cfg.Logger = logger
	}
}

func WithHTTPClient(httpClient *http.Client) clients.Option {
	return func(cfg *clients.Config) {
		cfg.HttpClient = httpClient
	}
}

func WithClientName(name string) clients.Option {
	return func(cfg *clients.Config) {
		cfg.ClientName = name
	}
}

func SkipWarnOnInsecureTLS() clients.Option {
	return func(cfg *clients.Config) {
		cfg.InsecureWarning = false
	}
}

func EnableStatistics() clients.Option {
	return func(cfg *clients.Config) {
		cfg.GatherStats = true
	}
}

func WithMaxRetries(maxRetries int) clients.Option {
	return func(cfg *clients.Config) {
		cfg.MaxRetries = maxRetries
	}
}

func (c *client) initializeFromClient(apiClient clients.Client) {
	c.APIAuthentication = authentication.NewAuthenticationService(apiClient)
	c.APIChannels = channels.NewChannelsService(apiClient)
	c.APIDrafts = drafts.NewDraftsService(apiClient)
	c.APIInvites = invites.NewInvitesService(apiClient)
	c.APIMessages = messages.NewMessagesService(apiClient)
	c.APIMobile = mobile.NewMobileService(apiClient)
	c.APINavigationViews = navigation_views.NewNavigationViewsService(apiClient)
	c.APIRealTimeEvents = real_time_events.NewRealTimeEventsService(apiClient)
	c.APIReminders = reminders.NewRemindersService(apiClient)
	c.APIScheduledMessages = scheduled_messages.NewScheduledMessagesService(apiClient)
	c.APIServerAndOrganizations = server_and_organizations.NewServerAndOrganizationsService(apiClient)
	c.APIUsers = users.NewUsersService(apiClient)
}
