// Package client provides the main Zulip API client interface and implementation,
// aggregating all API domain services for the Zulip platform.
package client

import (
	"log/slog"
	"net/http"

	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/authentication"
	"github.com/tum-zulip/go-zulip/zulip/api/channels"
	"github.com/tum-zulip/go-zulip/zulip/api/drafts"
	"github.com/tum-zulip/go-zulip/zulip/api/invites"
	"github.com/tum-zulip/go-zulip/zulip/api/messages"
	"github.com/tum-zulip/go-zulip/zulip/api/mobile"
	navigationviews "github.com/tum-zulip/go-zulip/zulip/api/navigation_views"
	realtimeevents "github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/api/reminders"
	scheduledmessages "github.com/tum-zulip/go-zulip/zulip/api/scheduled_messages"
	serverandorganizations "github.com/tum-zulip/go-zulip/zulip/api/server_and_organizations"
	"github.com/tum-zulip/go-zulip/zulip/api/users"
	"github.com/tum-zulip/go-zulip/zulip/client/statistics"
	"github.com/tum-zulip/go-zulip/zulip/internal/clients"
)

type Client interface {
	authentication.APIAuthentication
	channels.APIChannels
	drafts.APIDrafts
	invites.APIInvites
	messages.APIMessages
	mobile.APIMobile
	navigationviews.APINavigationViews
	realtimeevents.APIRealTimeEvents
	reminders.APIReminders
	scheduledmessages.APIScheduledMessages
	serverandorganizations.APIServerAndOrganizations
	users.APIUsers

	GetStatistics() statistics.Statistics
}

var _ Client = (*client)(nil)

type client struct {
	authentication.APIAuthentication
	channels.APIChannels
	drafts.APIDrafts
	invites.APIInvites
	messages.APIMessages
	mobile.APIMobile
	navigationviews.APINavigationViews
	realtimeevents.APIRealTimeEvents
	reminders.APIReminders
	scheduledmessages.APIScheduledMessages
	serverandorganizations.APIServerAndOrganizations
	users.APIUsers

	apiClient clients.RetryClient
}

func NewClient(zuliprc *zulip.RC, opts ...clients.Option) (Client, error) {
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

func (c *client) GetStatistics() statistics.Statistics {
	return c.apiClient.Stats.GetStatistics()
}

func WithAPISuffix(suffix string) clients.Option {
	return func(cfg *clients.Config) {
		cfg.APISuffix = suffix
	}
}

func WithAPIVersion(version string) clients.Option {
	return func(cfg *clients.Config) {
		cfg.APIVersion = version
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
		cfg.HTTPClient = httpClient
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
	c.APINavigationViews = navigationviews.NewNavigationViewsService(apiClient)
	c.APIRealTimeEvents = realtimeevents.NewRealTimeEventsService(apiClient)
	c.APIReminders = reminders.NewRemindersService(apiClient)
	c.APIScheduledMessages = scheduledmessages.NewScheduledMessagesService(apiClient)
	c.APIServerAndOrganizations = serverandorganizations.NewServerAndOrganizationsService(apiClient)
	c.APIUsers = users.NewUsersService(apiClient)
}
