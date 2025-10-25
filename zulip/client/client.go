package client

import (
	"log/slog"
	"net/http"
	"time"

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

	GetStatistics() map[string]time.Duration
}

var _ Client = (*client)(nil)

type client struct {
	apiClient *api_client.APIClient

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

type Option = api_client.Option

func NewClient(zuliprc *zuliprc.ZulipRC, options ...Option) (*client, error) {
	apiClient, err := api_client.NewAPIClient(zuliprc, options...)

	if err != nil {
		return nil, err
	}

	c := &client{apiClient: apiClient}

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

	return c, nil
}

func (c *client) GetStatistics() map[string]time.Duration {
	return c.apiClient.Stats.GetStatistics()
}

func WithAPISuffix(suffix string) Option {
	return api_client.WithAPISuffix(suffix)
}

func WithAPIVersion(version string) Option {
	return api_client.WithAPIVersion(version)

}

func WithLogger(logger *slog.Logger) Option {
	return api_client.WithLogger(logger)
}

func WithMaxRetries(maxRetries int) Option {
	return api_client.WithMaxRetries(maxRetries)
}

func WithHTTPClient(httpClient *http.Client) Option {
	return api_client.WithHTTPClient(httpClient)

}

func WithClientName(name string) Option {
	return api_client.WithClientName(name)

}

func SkipWarnOnInsecureTLS() Option {
	return api_client.SkipWarnOnInsecureTLS()
}

func GatherStatistics() Option {
	return api_client.GatherStatistics()
}
