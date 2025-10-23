package zulip

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
	"github.com/tum-zulip/go-zulip/zulip/internal/api_client"
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
}

var _ Client = (*client)(nil)

type client struct {
	apiClient api_client.APIClient

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

func NewClient(zuliprc *zuliprc.ZulipRC, options ...Option) *client {
	c := &client{}

	c.APIAuthentication = authentication.authenticationService{apiClient: &c.apiClient}
	c.APIChannels = channels.channelsService{apiClient: &c.apiClient}
	c.APIDrafts = drafts.draftsService{apiClient: &c.apiClient}
	c.APIInvites = invites.invitesService{apiClient: &c.apiClient}
	c.APIMessages = messages.messagesService{apiClient: &c.apiClient}
	c.APIMobile = mobile.mobileService{apiClient: &c.apiClient}
	c.APINavigationViews = navigation_views.navigationViewsService{apiClient: &c.apiClient}
	c.APIRealTimeEvents = real_time_events.realTimeEventsService{apiClient: &c.apiClient}
	c.APIReminders = reminders.remindersService{apiClient: &c.apiClient}
	c.APIScheduledMessages = scheduled_messages.scheduledMessagesService{apiClient: &c.apiClient}
	c.APIServerAndOrganizations = server_and_organizations.serverAndOrganizationsService{apiClient: &c.apiClient}
	c.APIUsers = users.usersService{apiClient: &c.apiClient}

	return c
}

type Option func(*client)

func WithAPISuffix(suffix string) Option {
	return func(c *client) {
		c.apiClient.ApiSuffix = suffix
	}
}

func WithAPIVersion(version string) Option {
	return func(c *client) {
		c.apiClient.ApiVersion = version
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(c *client) {
		if logger == nil {
			c.apiClient.Logger = slog.Default()
			return
		}
		c.apiClient.Logger = logger
	}
}

func WithMaxRetries(maxRetries int) Option {
	return func(c *client) {
		c.apiClient.MaxRetries = maxRetries
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *client) {
		c.apiClient.HttpClient = httpClient
	}
}

func WithClientName(name string) Option {
	return func(c *client) {
		c.apiClient.ClientName = name
	}
}

func SkipWarnOnInsecureTLS() Option {
	return func(c *client) {
		c.apiClient.InsecureWarning = false
	}
}
