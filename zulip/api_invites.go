package zulip

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type InvitesAPI interface {

	// CreateInviteLink Create a reusable invitation link
	//
	// Create a [reusable invitation link]
	// which can be used to invite new users to the organization.
	//
	// *Changes**: In Zulip 8.0 (feature level 209), added support for non-admin
	// users [with permission]
	// to use this endpoint. Previously, it was restricted to administrators only.
	//
	// In Zulip 6.0 (feature level 126), the `invite_expires_in_days`
	// parameter was removed and replaced by `invite_expires_in_minutes`.
	//
	// In Zulip 5.0 (feature level 117), added support for passing `null` as
	// the `invite_expires_in_days` parameter to request an invitation that never
	// expires.
	//
	// In Zulip 5.0 (feature level 96), the `invite_expires_in_days` parameter was
	// added which specified the number of days before the invitation would expire.
	//
	// [reusable invitation link]: https://zulip.com/help/invite-new-users#create-a-reusable-invitation-link
	// [with permission]: https://zulip.com/help/restrict-account-creation#change-who-can-send-invitations
	CreateInviteLink(ctx context.Context) CreateInviteLinkRequest

	// CreateInviteLinkExecute executes the request
	CreateInviteLinkExecute(r CreateInviteLinkRequest) (*CreateInviteLinkResponse, *http.Response, error)

	// GetInvites Get all invitations
	//
	// Fetch all unexpired [invitations] (i.e. email
	// invitations and reusable invitation links) that can be managed by the user.
	//
	// Note that administrators can manage invitations that were created by other users.
	//
	// *Changes**: Prior to Zulip 8.0 (feature level 209), non-admin users could
	// only create email invitations, and therefore the response would never include
	// reusable invitation links for these users.
	//
	// [invitations]: https://zulip.com/help/invite-new-users
	GetInvites(ctx context.Context) GetInvitesRequest

	// GetInvitesExecute executes the request
	GetInvitesExecute(r GetInvitesRequest) (*GetInvitesResponse, *http.Response, error)

	// ResendEmailInvite Resend an email invitation
	//
	// Resend an [email invitation].
	//
	// A user can only resend [invitations that they can manage].
	//
	// [email invitation]: https://zulip.com/help/invite-new-users#send-email-invitations
	// [invitations that they can manage]: https://zulip.com/help/invite-new-users#manage-pending-invitations
	ResendEmailInvite(ctx context.Context, inviteId int64) ResendEmailInviteRequest

	// ResendEmailInviteExecute executes the request
	ResendEmailInviteExecute(r ResendEmailInviteRequest) (*Response, *http.Response, error)

	// RevokeEmailInvite Revoke an email invitation
	//
	// Revoke an [email invitation].
	//
	// A user can only revoke [invitations that they can manage].
	//
	// [email invitation]: https://zulip.com/help/invite-new-users#send-email-invitations
	// [invitations that they can manage]: https://zulip.com/help/invite-new-users#manage-pending-invitations
	RevokeEmailInvite(ctx context.Context, inviteId int64) RevokeEmailInviteRequest

	// RevokeEmailInviteExecute executes the request
	RevokeEmailInviteExecute(r RevokeEmailInviteRequest) (*Response, *http.Response, error)

	// RevokeInviteLink Revoke a reusable invitation link
	//
	// Revoke a [reusable invitation link].
	//
	// A user can only revoke [invitations that they can manage].
	//
	// *Changes**: Prior to Zulip 8.0 (feature level 209), only organization
	// administrators were able to create and revoke reusable invitation links.
	//
	// [reusable invitation link]: https://zulip.com/help/invite-new-users#create-a-reusable-invitation-link
	// [invitations that they can manage]: https://zulip.com/help/invite-new-users#manage-pending-invitations
	RevokeInviteLink(ctx context.Context, inviteId int64) RevokeInviteLinkRequest

	// RevokeInviteLinkExecute executes the request
	RevokeInviteLinkExecute(r RevokeInviteLinkRequest) (*Response, *http.Response, error)

	// SendInvites Send invitations
	//
	// Send [invitations] to specified email addresses.
	//
	// *Changes**: In Zulip 6.0 (feature level 126), the `invite_expires_in_days`
	// parameter was removed and replaced by `invite_expires_in_minutes`.
	//
	// In Zulip 5.0 (feature level 117), added support for passing `null` as
	// the `invite_expires_in_days` parameter to request an invitation that never
	// expires.
	//
	// In Zulip 5.0 (feature level 96), the `invite_expires_in_days` parameter was
	// added which specified the number of days before the invitation would expire.
	//
	// [invitations]: https://zulip.com/help/invite-new-users
	SendInvites(ctx context.Context) SendInvitesRequest

	// SendInvitesExecute executes the request
	SendInvitesExecute(r SendInvitesRequest) (*Response, *http.Response, error)
}

type CreateInviteLinkRequest struct {
	ctx                              context.Context
	apiService                       InvitesAPI
	inviteExpiresInMinutes           *int32
	inviteAs                         *Role
	channelIds                       *[]int64
	groupIds                         *[]int64
	includeRealmDefaultSubscriptions *bool
	welcomeMessageCustomText         *string
}

// The number of minutes before the invitation will expire. If `null`, the invitation will never expire. If unspecified, the server will use a default value (based on the `INVITATION_LINK_VALIdITY_MINUTES` server setting, which defaults to 14400, i.e. 10 days) for when the invitation will expire.
//
// **Changes**: New in Zulip 6.0 (feature level 126). Previously, there was an `invite_expires_in_days` parameter, which specified the duration in days instead of minutes.
func (r CreateInviteLinkRequest) InviteExpiresInMinutes(inviteExpiresInMinutes int32) CreateInviteLinkRequest {
	r.inviteExpiresInMinutes = &inviteExpiresInMinutes
	return r
}

// The [organization-level role] of the user that is created when the invitation is accepted. Possible values are:
//   - RoleOwner
//   - RoleAdmin
//   - RoleModerator
//   - RoleMember
//   - RoleGuest
//
// Users can only create invitation links for [roles with equal or stricter restrictions] as their own. For example, a moderator cannot invite someone to be an owner or administrator, but they can invite them to be a moderator or member.
//
// **Changes**: In Zulip 4.0 (feature level 61), added support for inviting users as moderators.
//
// [organization-level role]: https://zulip.com/api/roles-and-permissions
// [roles with equal or stricter restrictions]: https://zulip.com/api/roles-and-permissions#permission-levels
func (r CreateInviteLinkRequest) InviteAs(inviteAs Role) CreateInviteLinkRequest {
	r.inviteAs = &inviteAs
	return r
}

// A list containing the [Ids of the channels] that the newly created user will be automatically subscribed to if the invitation is accepted, in addition to any default channels that the new user may be subscribed to based on the `include_realm_default_subscriptions` parameter.  Requested channels must either be default channels for the organization, or ones the acting user has permission to add subscribers to.  This list must be empty if the current user has the unlikely configuration of being able to create reusable invitation links while lacking permission to [subscribe other users to channels].
//
// **Changes**: Prior to Zulip 10.0 (feature level 342), default channels that the acting user did not directly have permission to add subscribers to would be rejected.
//
// [subscribe other users to channels]: https://zulip.com/help/configure-who-can-invite-to-channels
// [Ids of the channels]: https://zulip.com/api/get-stream-id
func (r CreateInviteLinkRequest) ChannelIds(channelIds []int64) CreateInviteLinkRequest {
	r.channelIds = &channelIds
	return r
}

// A list containing the [Ids of the user groups] that the newly created user will be automatically added to if the invitation is accepted. If the list is empty, then the new user will not be added to any user groups. The acting user must have permission to add users to the groups listed in this request.
//
// **Changes**: New in Zulip 10.0 (feature level 322).
//
// [Ids of the user groups]: https://zulip.com/api/get-user-groups
func (r CreateInviteLinkRequest) GroupIds(groupIds []int64) CreateInviteLinkRequest {
	r.groupIds = &groupIds
	return r
}

// Boolean indicating whether the newly created user should be subscribed to the [default channels] for the organization.  Note that this parameter can be `true` even if the current user does not generally have permission to [subscribe other users to channels].
//
// **Changes**: New in Zulip 9.0 (feature level 261). Previous versions of Zulip behaved as though this parameter was always `false`; clients needed to include the organization's default channels in the `stream_ids` parameter for a newly created user to be automatically subscribed to them.
//
// [default channels]: https://zulip.com/help/set-default-channels-for-new-users
// [subscribe other users to channels]: https://zulip.com/help/configure-who-can-invite-to-channels
func (r CreateInviteLinkRequest) IncludeRealmDefaultSubscriptions(includeRealmDefaultSubscriptions bool) CreateInviteLinkRequest {
	r.includeRealmDefaultSubscriptions = &includeRealmDefaultSubscriptions
	return r
}

// Custom message text, in Zulip Markdown format, to be sent by the Welcome Bot to new users that join the organization via this invitation.  Maximum length is 8000 characters.  Only organization administrators can use this feature; for other users, the value is always `null`.  - `null`: the organization's default `welcome_message_custom_text` is used. - Empty string: no Welcome Bot custom message is sent. - Otherwise, the provided string is the custom message.
//
// **Changes**: New in Zulip 11.0 (feature level 416).
func (r CreateInviteLinkRequest) WelcomeMessageCustomText(welcomeMessageCustomText string) CreateInviteLinkRequest {
	r.welcomeMessageCustomText = &welcomeMessageCustomText
	return r
}

func (r CreateInviteLinkRequest) Execute() (*CreateInviteLinkResponse, *http.Response, error) {
	return r.apiService.CreateInviteLinkExecute(r)
}

// CreateInviteLink Create a reusable invitation link
//
// Create a [reusable invitation link]
// which can be used to invite new users to the organization.
//
// *Changes**: In Zulip 8.0 (feature level 209), added support for non-admin
// users [with permission]
// to use this endpoint. Previously, it was restricted to administrators only.
//
// In Zulip 6.0 (feature level 126), the `invite_expires_in_days`
// parameter was removed and replaced by `invite_expires_in_minutes`.
//
// In Zulip 5.0 (feature level 117), added support for passing `null` as
// the `invite_expires_in_days` parameter to request an invitation that never
// expires.
//
// In Zulip 5.0 (feature level 96), the `invite_expires_in_days` parameter was
// added which specified the number of days before the invitation would expire.
//
// [reusable invitation link]: https://zulip.com/help/invite-new-users#create-a-reusable-invitation-link
// [with permission]: https://zulip.com/help/restrict-account-creation#change-who-can-send-invitations
func (c *simpleClient) CreateInviteLink(ctx context.Context) CreateInviteLinkRequest {
	return CreateInviteLinkRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateInviteLinkExecute(r CreateInviteLinkRequest) (*CreateInviteLinkResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateInviteLinkResponse{}
		endpoint = "/invites/multiuse"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addOptionalParam(form, "invite_expires_in_minutes", r.inviteExpiresInMinutes)
	addOptionalParam(form, "invite_as", r.inviteAs)
	addOptionalParam(form, "stream_ids", r.channelIds)
	addOptionalParam(form, "group_ids", r.groupIds)
	addOptionalParam(form, "include_realm_default_subscriptions", r.includeRealmDefaultSubscriptions)
	addOptionalParam(form, "welcome_message_custom_text", r.welcomeMessageCustomText)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetInvitesRequest struct {
	ctx        context.Context
	apiService InvitesAPI
}

func (r GetInvitesRequest) Execute() (*GetInvitesResponse, *http.Response, error) {
	return r.apiService.GetInvitesExecute(r)
}

// GetInvites Get all invitations
//
// Fetch all unexpired [invitations] (i.e. email
// invitations and reusable invitation links) that can be managed by the user.
//
// Note that administrators can manage invitations that were created by other users.
//
// *Changes**: Prior to Zulip 8.0 (feature level 209), non-admin users could
// only create email invitations, and therefore the response would never include
// reusable invitation links for these users.
//
// [invitations]: https://zulip.com/help/invite-new-users
func (c *simpleClient) GetInvites(ctx context.Context) GetInvitesRequest {
	return GetInvitesRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetInvitesExecute(r GetInvitesRequest) (*GetInvitesResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetInvitesResponse{}
		endpoint = "/invites"
	)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type ResendEmailInviteRequest struct {
	ctx        context.Context
	apiService InvitesAPI
	inviteId   int64
}

func (r ResendEmailInviteRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.ResendEmailInviteExecute(r)
}

// ResendEmailInvite Resend an email invitation
//
// Resend an [email invitation].
//
// A user can only resend [invitations that they can manage].
//
// [email invitation]: https://zulip.com/help/invite-new-users#send-email-invitations
// [invitations that they can manage]: https://zulip.com/help/invite-new-users#manage-pending-invitations
func (c *simpleClient) ResendEmailInvite(ctx context.Context, inviteId int64) ResendEmailInviteRequest {
	return ResendEmailInviteRequest{
		apiService: c,
		ctx:        ctx,
		inviteId:   inviteId,
	}
}

// Execute executes the request
func (c *simpleClient) ResendEmailInviteExecute(r ResendEmailInviteRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/invites/{invite_id}/resend"
	)

	endpoint = strings.Replace(endpoint, "{invite_id}", idToString(r.inviteId), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RevokeEmailInviteRequest struct {
	ctx        context.Context
	apiService InvitesAPI
	inviteId   int64
}

func (r RevokeEmailInviteRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.RevokeEmailInviteExecute(r)
}

// RevokeEmailInvite Revoke an email invitation
//
// Revoke an [email invitation].
//
// A user can only revoke [invitations that they can manage].
//
// [email invitation]: https://zulip.com/help/invite-new-users#send-email-invitations
// [invitations that they can manage]: https://zulip.com/help/invite-new-users#manage-pending-invitations
func (c *simpleClient) RevokeEmailInvite(ctx context.Context, inviteId int64) RevokeEmailInviteRequest {
	return RevokeEmailInviteRequest{
		apiService: c,
		ctx:        ctx,
		inviteId:   inviteId,
	}
}

// Execute executes the request
func (c *simpleClient) RevokeEmailInviteExecute(r RevokeEmailInviteRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/invites/{invite_id}"
	)

	endpoint = strings.Replace(endpoint, "{invite_id}", idToString(r.inviteId), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RevokeInviteLinkRequest struct {
	ctx        context.Context
	apiService InvitesAPI
	inviteId   int64
}

func (r RevokeInviteLinkRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.RevokeInviteLinkExecute(r)
}

// RevokeInviteLink Revoke a reusable invitation link
//
// Revoke a [reusable invitation link].
//
// A user can only revoke [invitations that they can manage].
//
// *Changes**: Prior to Zulip 8.0 (feature level 209), only organization
// administrators were able to create and revoke reusable invitation links.
//
// [reusable invitation link]: https://zulip.com/help/invite-new-users#create-a-reusable-invitation-link
// [invitations that they can manage]: https://zulip.com/help/invite-new-users#manage-pending-invitations
func (c *simpleClient) RevokeInviteLink(ctx context.Context, inviteId int64) RevokeInviteLinkRequest {
	return RevokeInviteLinkRequest{
		apiService: c,
		ctx:        ctx,
		inviteId:   inviteId,
	}
}

// Execute executes the request
func (c *simpleClient) RevokeInviteLinkExecute(r RevokeInviteLinkRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/invites/multiuse/{invite_id}"
	)

	endpoint = strings.Replace(endpoint, "{invite_id}", idToString(r.inviteId), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type SendInvitesRequest struct {
	ctx                              context.Context
	apiService                       InvitesAPI
	inviteeEmails                    *string
	channelIds                       *[]int64
	inviteExpiresInMinutes           *int32
	inviteAs                         *Role
	groupIds                         *[]int64
	includeRealmDefaultSubscriptions *bool
	notifyReferrerOnJoin             *bool
	welcomeMessageCustomText         *string
}

// The string containing the email addresses, separated by commas or newlines, that will be sent an invitation.
func (r SendInvitesRequest) InviteeEmails(inviteeEmails string) SendInvitesRequest {
	r.inviteeEmails = &inviteeEmails
	return r
}

// A list containing the [Ids of the channels] that the newly created user will be automatically subscribed to if the invitation is accepted, in addition to any default channels that the new user may be subscribed to based on the `include_realm_default_subscriptions` parameter.  Requested channels must either be default channels for the organization, or ones the acting user has permission to add subscribers to.  This list must be empty if the current user has the unlikely configuration of being able to send invitations while lacking permission to [subscribe other users to channels].
//
// **Changes**: Prior to Zulip 10.0 (feature level 342), default channels that the acting user did not directly have permission to add subscribers to would be rejected.  Before Zulip 7.0 (feature level 180), specifying `stream_ids` as an empty list resulted in an error.
//
// [subscribe other users to channels]: https://zulip.com/help/configure-who-can-invite-to-channels
// [Ids of the channels]: https://zulip.com/api/get-stream-id
func (r SendInvitesRequest) ChannelIds(channelIds []int64) SendInvitesRequest {
	r.channelIds = &channelIds
	return r
}

// The number of minutes before the invitation will expire. If `null`, the invitation will never expire. If unspecified, the server will use a default value (based on the `INVITATION_LINK_VALIdITY_MINUTES` server setting, which defaults to 14400, i.e. 10 days) for when the invitation will expire.
//
// **Changes**: New in Zulip 6.0 (feature level 126). Previously, there was an `invite_expires_in_days` parameter, which specified the duration in days instead of minutes.
func (r SendInvitesRequest) InviteExpiresInMinutes(inviteExpiresInMinutes int32) SendInvitesRequest {
	r.inviteExpiresInMinutes = &inviteExpiresInMinutes
	return r
}

// The [organization-level role] of the user that is created when the invitation is accepted. Possible values are:
//   - RoleOwner
//   - RoleAdmin
//   - RoleModerator
//   - RoleMember
//   - RoleGuest
//
// Users can only create invitation links for [roles with equal or stricter restrictions] as their own. For example, a moderator cannot invite someone to be an owner or administrator, but they can invite them to be a moderator or member.
//
// **Changes**: In Zulip 4.0 (feature level 61), added support for inviting users as moderators.
//
// [organization-level role]: https://zulip.com/api/roles-and-permissions
// [roles with equal or stricter restrictions]: https://zulip.com/api/roles-and-permissions#permission-levels
func (r SendInvitesRequest) InviteAs(inviteAs Role) SendInvitesRequest {
	r.inviteAs = &inviteAs
	return r
}

// A list containing the [Ids of the user groups] that the newly created user will be automatically added to if the invitation is accepted. If the list is empty, then the new user will not be added to any user groups. The acting user must have permission to add users to the groups listed in this request.
//
// **Changes**: New in Zulip 10.0 (feature level 322).
//
// [Ids of the user groups]: https://zulip.com/api/get-user-groups
func (r SendInvitesRequest) GroupIds(groupIds []int64) SendInvitesRequest {
	r.groupIds = &groupIds
	return r
}

// Boolean indicating whether the newly created user should be subscribed to the [default channels] for the organization.  Note that this parameter can be `true` even if the user creating the invitation does not generally have permission to [subscribe other users to channels].
//
// **Changes**: New in Zulip 9.0 (feature level 261). Previous versions of Zulip behaved as though this parameter was always `false`; clients needed to include the organization's default channels in the `stream_ids` parameter for a newly created user to be automatically subscribed to them.
//
// [default channels]: https://zulip.com/help/set-default-channels-for-new-users
// [subscribe other users to channels]: https://zulip.com/help/configure-who-can-invite-to-channels
func (r SendInvitesRequest) IncludeRealmDefaultSubscriptions(includeRealmDefaultSubscriptions bool) SendInvitesRequest {
	r.includeRealmDefaultSubscriptions = &includeRealmDefaultSubscriptions
	return r
}

// A boolean indicating whether the referrer would like to receive a direct message from [notification bot] when a user account is created using this invitation.
//
// **Changes**: New in Zulip 9.0 (feature level 267). Previously, referrers always received such direct messages.
//
// [notification bot]: https://zulip.com/help/configure-automated-notices
func (r SendInvitesRequest) NotifyReferrerOnJoin(notifyReferrerOnJoin bool) SendInvitesRequest {
	r.notifyReferrerOnJoin = &notifyReferrerOnJoin
	return r
}

// Custom message text, in Zulip Markdown format, to be sent by the Welcome Bot to new users that join the organization via this invitation.  Maximum length is 8000 characters.  Only organization administrators can use this feature; for other users, the value is always `null`.  - `null`: the organization's default `welcome_message_custom_text` is used. - Empty string: no Welcome Bot custom message is sent. - Otherwise, the provided string is the custom message.
//
// **Changes**: New in Zulip 11.0 (feature level 416).
func (r SendInvitesRequest) WelcomeMessageCustomText(welcomeMessageCustomText string) SendInvitesRequest {
	r.welcomeMessageCustomText = &welcomeMessageCustomText
	return r
}

func (r SendInvitesRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.SendInvitesExecute(r)
}

// SendInvites Send invitations
//
// Send [invitations] to specified email addresses.
//
// *Changes**: In Zulip 6.0 (feature level 126), the `invite_expires_in_days`
// parameter was removed and replaced by `invite_expires_in_minutes`.
//
// In Zulip 5.0 (feature level 117), added support for passing `null` as
// the `invite_expires_in_days` parameter to request an invitation that never
// expires.
//
// In Zulip 5.0 (feature level 96), the `invite_expires_in_days` parameter was
// added which specified the number of days before the invitation would expire.
//
// [invitations]: https://zulip.com/help/invite-new-users
func (c *simpleClient) SendInvites(ctx context.Context) SendInvitesRequest {
	return SendInvitesRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) SendInvitesExecute(r SendInvitesRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/invites"
	)
	if r.inviteeEmails == nil {
		return nil, nil, reportError("inviteeEmails is required and must be specified")
	}
	if r.channelIds == nil {
		return nil, nil, reportError("channelIds is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "invitee_emails", r.inviteeEmails)
	addOptionalParam(form, "invite_expires_in_minutes", r.inviteExpiresInMinutes)
	addOptionalParam(form, "invite_as", r.inviteAs)
	addParam(form, "stream_ids", r.channelIds)
	addOptionalParam(form, "group_ids", r.groupIds)
	addOptionalParam(form, "include_realm_default_subscriptions", r.includeRealmDefaultSubscriptions)
	addOptionalParam(form, "notify_referrer_on_join", r.notifyReferrerOnJoin)
	addOptionalParam(form, "welcome_message_custom_text", r.welcomeMessageCustomText)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}
