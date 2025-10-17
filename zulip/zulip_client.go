package zulip

import (
	"regexp"
	"strings"
)

const (
	defaultClientName = "go-zulip/1.0.0"
	defaultAPISuffix  = "api"
	defaultAPIVersion = "v1"
)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer("%5B", "[", "%5D", "]")
)

type Client interface {
	AuthenticationAPI
	ChannelsAPI
	DraftsAPI
	InvitesAPI
	MessagesAPI
	MobileAPI
	NavigationViewsAPI
	RealTimeEventsAPI
	RemindersAPI
	ScheduledMessagesAPI
	ServerAndOrganizationsAPI
	UsersAPI
	WebhooksAPI
}
