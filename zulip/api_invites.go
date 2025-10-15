package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type InvitesAPI interface {

	/*
			CreateInviteLink Create a reusable invitation link

			Create a [reusable invitation link](zulip.com/help/invite-new-users#create-a-reusable-invitation-link
		which can be used to invite new users to the organization.

		**Changes**: In Zulip 8.0 (feature level 209), added support for non-admin
		users [with permission](zulip.com/help/restrict-account-creation#change-who-can-send-invitations
		to use this endpoint. Previously, it was restricted to administrators only.

		In Zulip 6.0 (feature level 126), the `invite_expires_in_days`
		parameter was removed and replaced by `invite_expires_in_minutes`.

		In Zulip 5.0 (feature level 117), added support for passing `null` as
		the `invite_expires_in_days` parameter to request an invitation that never
		expires.

		In Zulip 5.0 (feature level 96), the `invite_expires_in_days` parameter was
		added which specified the number of days before the invitation would expire.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return CreateInviteLinkRequest
	*/
	CreateInviteLink(ctx context.Context) CreateInviteLinkRequest

	// CreateInviteLinkExecute executes the request
	//  @return CreateInviteLinkResponse
	CreateInviteLinkExecute(r CreateInviteLinkRequest) (*CreateInviteLinkResponse, *http.Response, error)

	/*
			GetInvites Get all invitations

			Fetch all unexpired [invitations](zulip.com/help/invite-new-users) (i.e. email
		invitations and reusable invitation links) that can be managed by the user.

		Note that administrators can manage invitations that were created by other users.

		**Changes**: Prior to Zulip 8.0 (feature level 209), non-admin users could
		only create email invitations, and therefore the response would never include
		reusable invitation links for these users.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GetInvitesRequest
	*/
	GetInvites(ctx context.Context) GetInvitesRequest

	// GetInvitesExecute executes the request
	//  @return GetInvitesResponse
	GetInvitesExecute(r GetInvitesRequest) (*GetInvitesResponse, *http.Response, error)

	/*
			ResendEmailInvite Resend an email invitation

			Resend an [email invitation](zulip.com/help/invite-new-users#send-email-invitations.

		A user can only resend [invitations that they can
		manage](zulip.com/help/invite-new-users#manage-pending-invitations.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param inviteId The Id of the email invitation to be resent.
			@return ResendEmailInviteRequest
	*/
	ResendEmailInvite(ctx context.Context, inviteId int64) ResendEmailInviteRequest

	// ResendEmailInviteExecute executes the request
	//  @return Response
	ResendEmailInviteExecute(r ResendEmailInviteRequest) (*Response, *http.Response, error)

	/*
			RevokeEmailInvite Revoke an email invitation

			Revoke an [email invitation](zulip.com/help/invite-new-users#send-email-invitations.

		A user can only revoke [invitations that they can
		manage](zulip.com/help/invite-new-users#manage-pending-invitations.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param inviteId The Id of the email invitation to be revoked.
			@return RevokeEmailInviteRequest
	*/
	RevokeEmailInvite(ctx context.Context, inviteId int64) RevokeEmailInviteRequest

	// RevokeEmailInviteExecute executes the request
	//  @return Response
	RevokeEmailInviteExecute(r RevokeEmailInviteRequest) (*Response, *http.Response, error)

	/*
			RevokeInviteLink Revoke a reusable invitation link

			Revoke a [reusable invitation link](zulip.com/help/invite-new-users#create-a-reusable-invitation-link.

		A user can only revoke [invitations that they can
		manage](zulip.com/help/invite-new-users#manage-pending-invitations.

		**Changes**: Prior to Zulip 8.0 (feature level 209), only organization
		administrators were able to create and revoke reusable invitation links.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param inviteId The Id of the reusable invitation link to be revoked.
			@return RevokeInviteLinkRequest
	*/
	RevokeInviteLink(ctx context.Context, inviteId int64) RevokeInviteLinkRequest

	// RevokeInviteLinkExecute executes the request
	//  @return Response
	RevokeInviteLinkExecute(r RevokeInviteLinkRequest) (*Response, *http.Response, error)

	/*
			SendInvites Send invitations

			Send [invitations](zulip.com/help/invite-new-users) to specified email addresses.

		**Changes**: In Zulip 6.0 (feature level 126), the `invite_expires_in_days`
		parameter was removed and replaced by `invite_expires_in_minutes`.

		In Zulip 5.0 (feature level 117), added support for passing `null` as
		the `invite_expires_in_days` parameter to request an invitation that never
		expires.

		In Zulip 5.0 (feature level 96), the `invite_expires_in_days` parameter was
		added which specified the number of days before the invitation would expire.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return SendInvitesRequest
	*/
	SendInvites(ctx context.Context) SendInvitesRequest

	// SendInvitesExecute executes the request
	//  @return Response
	SendInvitesExecute(r SendInvitesRequest) (*Response, *http.Response, error)
}

type CreateInviteLinkRequest struct {
	ctx                              context.Context
	ApiService                       InvitesAPI
	inviteExpiresInMinutes           *int32
	inviteAs                         *Role
	streamIds                        *[]int64
	groupIds                         *[]int64
	includeRealmDefaultSubscriptions *bool
	welcomeMessageCustomText         *string
}

// The number of minutes before the invitation will expire. If &#x60;null&#x60;, the invitation will never expire. If unspecified, the server will use a default value (based on the &#x60;INVITATION_LINK_VALIdITY_MINUTES&#x60; server setting, which defaults to 14400, i.e. 10 days) for when the invitation will expire.  **Changes**: New in Zulip 6.0 (feature level 126). Previously, there was an &#x60;invite_expires_in_days&#x60; parameter, which specified the duration in days instead of minutes.
func (r CreateInviteLinkRequest) InviteExpiresInMinutes(inviteExpiresInMinutes int32) CreateInviteLinkRequest {
	r.inviteExpiresInMinutes = &inviteExpiresInMinutes
	return r
}

// The [organization-level role](zulip.com/api/roles-and-permissions) of the user that is created when the invitation is accepted. Possible values are:  - 100 = Organization owner - 200 = Organization administrator - 300 = Organization moderator - 400 = Member - 600 = Guest  Users can only create invitation links for [roles with equal or stricter restrictions](zulip.com/api/roles-and-permissions#permission-levels as their own. For example, a moderator cannot invite someone to be an owner or administrator, but they can invite them to be a moderator or member.  **Changes**: In Zulip 4.0 (feature level 61), added support for inviting users as moderators.
func (r CreateInviteLinkRequest) InviteAs(inviteAs Role) CreateInviteLinkRequest {
	r.inviteAs = &inviteAs
	return r
}

// A list containing the [Ids of the channels](zulip.com/api/get-stream-id) that the newly created user will be automatically subscribed to if the invitation is accepted, in addition to any default channels that the new user may be subscribed to based on the &#x60;include_realm_default_subscriptions&#x60; parameter.  Requested channels must either be default channels for the organization, or ones the acting user has permission to add subscribers to.  This list must be empty if the current user has the unlikely configuration of being able to create reusable invitation links while lacking permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: Prior to Zulip 10.0 (feature level 342), default channels that the acting user did not directly have permission to add subscribers to would be rejected.  [can-subscribe-others]: /help/configure-who-can-invite-to-channels
func (r CreateInviteLinkRequest) StreamIds(streamIds []int64) CreateInviteLinkRequest {
	r.streamIds = &streamIds
	return r
}

// A list containing the [Ids of the user groups](zulip.com/api/get-user-groups) that the newly created user will be automatically added to if the invitation is accepted. If the list is empty, then the new user will not be added to any user groups. The acting user must have permission to add users to the groups listed in this request.  **Changes**: New in Zulip 10.0 (feature level 322).
func (r CreateInviteLinkRequest) GroupIds(groupIds []int64) CreateInviteLinkRequest {
	r.groupIds = &groupIds
	return r
}

// Boolean indicating whether the newly created user should be subscribed to the [default channels][default-channels] for the organization.  Note that this parameter can be &#x60;true&#x60; even if the current user does not generally have permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: New in Zulip 9.0 (feature level 261). Previous versions of Zulip behaved as though this parameter was always &#x60;false&#x60;; clients needed to include the organization&#39;s default channels in the &#x60;stream_ids&#x60; parameter for a newly created user to be automatically subscribed to them.  [default-channels]: /help/set-default-channels-for-new-users [can-subscribe-others]: /help/configure-who-can-invite-to-channels
func (r CreateInviteLinkRequest) IncludeRealmDefaultSubscriptions(includeRealmDefaultSubscriptions bool) CreateInviteLinkRequest {
	r.includeRealmDefaultSubscriptions = &includeRealmDefaultSubscriptions
	return r
}

// Custom message text, in Zulip Markdown format, to be sent by the Welcome Bot to new users that join the organization via this invitation.  Maximum length is 8000 characters.  Only organization administrators can use this feature; for other users, the value is always &#x60;null&#x60;.  - &#x60;null&#x60;: the organization&#39;s default &#x60;welcome_message_custom_text&#x60; is used. - Empty string: no Welcome Bot custom message is sent. - Otherwise, the provided string is the custom message.  **Changes**: New in Zulip 11.0 (feature level 416).
func (r CreateInviteLinkRequest) WelcomeMessageCustomText(welcomeMessageCustomText string) CreateInviteLinkRequest {
	r.welcomeMessageCustomText = &welcomeMessageCustomText
	return r
}

func (r CreateInviteLinkRequest) Execute() (*CreateInviteLinkResponse, *http.Response, error) {
	return r.ApiService.CreateInviteLinkExecute(r)
}

/*
CreateInviteLink Create a reusable invitation link

Create a [reusable invitation link](zulip.com/help/invite-new-users#create-a-reusable-invitation-link
which can be used to invite new users to the organization.

**Changes**: In Zulip 8.0 (feature level 209), added support for non-admin
users [with permission](zulip.com/help/restrict-account-creation#change-who-can-send-invitations
to use this endpoint. Previously, it was restricted to administrators only.

In Zulip 6.0 (feature level 126), the `invite_expires_in_days`
parameter was removed and replaced by `invite_expires_in_minutes`.

In Zulip 5.0 (feature level 117), added support for passing `null` as
the `invite_expires_in_days` parameter to request an invitation that never
expires.

In Zulip 5.0 (feature level 96), the `invite_expires_in_days` parameter was
added which specified the number of days before the invitation would expire.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateInviteLinkRequest
*/
func (c *Client) CreateInviteLink(ctx context.Context) CreateInviteLinkRequest {
	return CreateInviteLinkRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateInviteLinkResponse
func (c *Client) CreateInviteLinkExecute(r CreateInviteLinkRequest) (*CreateInviteLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateInviteLinkResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invites/multiuse"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.inviteExpiresInMinutes != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "invite_expires_in_minutes", r.inviteExpiresInMinutes, "form", "")
	}
	if r.inviteAs != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "invite_as", r.inviteAs, "", "")
	}
	if r.streamIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "stream_ids", r.streamIds, "form", "multi")
	}
	if r.groupIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_ids", r.groupIds, "form", "multi")
	}
	if r.includeRealmDefaultSubscriptions != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "include_realm_default_subscriptions", r.includeRealmDefaultSubscriptions, "", "")
	}
	if r.welcomeMessageCustomText != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "welcome_message_custom_text", r.welcomeMessageCustomText, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetInvitesRequest struct {
	ctx        context.Context
	ApiService InvitesAPI
}

func (r GetInvitesRequest) Execute() (*GetInvitesResponse, *http.Response, error) {
	return r.ApiService.GetInvitesExecute(r)
}

/*
GetInvites Get all invitations

Fetch all unexpired [invitations](zulip.com/help/invite-new-users) (i.e. email
invitations and reusable invitation links) that can be managed by the user.

Note that administrators can manage invitations that were created by other users.

**Changes**: Prior to Zulip 8.0 (feature level 209), non-admin users could
only create email invitations, and therefore the response would never include
reusable invitation links for these users.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetInvitesRequest
*/
func (c *Client) GetInvites(ctx context.Context) GetInvitesRequest {
	return GetInvitesRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetInvitesResponse
func (c *Client) GetInvitesExecute(r GetInvitesRequest) (*GetInvitesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetInvitesResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ResendEmailInviteRequest struct {
	ctx        context.Context
	ApiService InvitesAPI
	inviteId   int64
}

func (r ResendEmailInviteRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.ResendEmailInviteExecute(r)
}

/*
ResendEmailInvite Resend an email invitation

Resend an [email invitation](zulip.com/help/invite-new-users#send-email-invitations.

A user can only resend [invitations that they can
manage](zulip.com/help/invite-new-users#manage-pending-invitations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inviteId The Id of the email invitation to be resent.
	@return ResendEmailInviteRequest
*/
func (c *Client) ResendEmailInvite(ctx context.Context, inviteId int64) ResendEmailInviteRequest {
	return ResendEmailInviteRequest{
		ApiService: c,
		ctx:        ctx,
		inviteId:   inviteId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *Client) ResendEmailInviteExecute(r ResendEmailInviteRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invites/{invite_id}/resend"
	localVarPath = strings.Replace(localVarPath, "{"+"invite_id"+"}", url.PathEscape(parameterValueToString(r.inviteId, "inviteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RevokeEmailInviteRequest struct {
	ctx        context.Context
	ApiService InvitesAPI
	inviteId   int64
}

func (r RevokeEmailInviteRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RevokeEmailInviteExecute(r)
}

/*
RevokeEmailInvite Revoke an email invitation

Revoke an [email invitation](zulip.com/help/invite-new-users#send-email-invitations.

A user can only revoke [invitations that they can
manage](zulip.com/help/invite-new-users#manage-pending-invitations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inviteId The Id of the email invitation to be revoked.
	@return RevokeEmailInviteRequest
*/
func (c *Client) RevokeEmailInvite(ctx context.Context, inviteId int64) RevokeEmailInviteRequest {
	return RevokeEmailInviteRequest{
		ApiService: c,
		ctx:        ctx,
		inviteId:   inviteId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *Client) RevokeEmailInviteExecute(r RevokeEmailInviteRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invites/{invite_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"invite_id"+"}", url.PathEscape(parameterValueToString(r.inviteId, "inviteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RevokeInviteLinkRequest struct {
	ctx        context.Context
	ApiService InvitesAPI
	inviteId   int64
}

func (r RevokeInviteLinkRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RevokeInviteLinkExecute(r)
}

/*
RevokeInviteLink Revoke a reusable invitation link

Revoke a [reusable invitation link](zulip.com/help/invite-new-users#create-a-reusable-invitation-link.

A user can only revoke [invitations that they can
manage](zulip.com/help/invite-new-users#manage-pending-invitations.

**Changes**: Prior to Zulip 8.0 (feature level 209), only organization
administrators were able to create and revoke reusable invitation links.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inviteId The Id of the reusable invitation link to be revoked.
	@return RevokeInviteLinkRequest
*/
func (c *Client) RevokeInviteLink(ctx context.Context, inviteId int64) RevokeInviteLinkRequest {
	return RevokeInviteLinkRequest{
		ApiService: c,
		ctx:        ctx,
		inviteId:   inviteId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *Client) RevokeInviteLinkExecute(r RevokeInviteLinkRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invites/multiuse/{invite_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"invite_id"+"}", url.PathEscape(parameterValueToString(r.inviteId, "inviteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SendInvitesRequest struct {
	ctx                              context.Context
	ApiService                       InvitesAPI
	inviteeEmails                    *string
	streamIds                        *[]int64
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

// A list containing the [Ids of the channels](zulip.com/api/get-stream-id) that the newly created user will be automatically subscribed to if the invitation is accepted, in addition to any default channels that the new user may be subscribed to based on the &#x60;include_realm_default_subscriptions&#x60; parameter.  Requested channels must either be default channels for the organization, or ones the acting user has permission to add subscribers to.  This list must be empty if the current user has the unlikely configuration of being able to send invitations while lacking permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: Prior to Zulip 10.0 (feature level 342), default channels that the acting user did not directly have permission to add subscribers to would be rejected.  Before Zulip 7.0 (feature level 180), specifying &#x60;stream_ids&#x60; as an empty list resulted in an error.  [can-subscribe-others]: /help/configure-who-can-invite-to-channels
func (r SendInvitesRequest) StreamIds(streamIds []int64) SendInvitesRequest {
	r.streamIds = &streamIds
	return r
}

// The number of minutes before the invitation will expire. If &#x60;null&#x60;, the invitation will never expire. If unspecified, the server will use a default value (based on the &#x60;INVITATION_LINK_VALIdITY_MINUTES&#x60; server setting, which defaults to 14400, i.e. 10 days) for when the invitation will expire.  **Changes**: New in Zulip 6.0 (feature level 126). Previously, there was an &#x60;invite_expires_in_days&#x60; parameter, which specified the duration in days instead of minutes.
func (r SendInvitesRequest) InviteExpiresInMinutes(inviteExpiresInMinutes int32) SendInvitesRequest {
	r.inviteExpiresInMinutes = &inviteExpiresInMinutes
	return r
}

// The [organization-level role](zulip.com/api/roles-and-permissions) of the user that is created when the invitation is accepted. Possible values are:  - 100 = Organization owner - 200 = Organization administrator - 300 = Organization moderator - 400 = Member - 600 = Guest  Users can only create invitation links for [roles with equal or stricter restrictions](zulip.com/api/roles-and-permissions#permission-levels as their own. For example, a moderator cannot invite someone to be an owner or administrator, but they can invite them to be a moderator or member.  **Changes**: In Zulip 4.0 (feature level 61), added support for inviting users as moderators.
func (r SendInvitesRequest) InviteAs(inviteAs Role) SendInvitesRequest {
	r.inviteAs = &inviteAs
	return r
}

// A list containing the [Ids of the user groups](zulip.com/api/get-user-groups) that the newly created user will be automatically added to if the invitation is accepted. If the list is empty, then the new user will not be added to any user groups. The acting user must have permission to add users to the groups listed in this request.  **Changes**: New in Zulip 10.0 (feature level 322).
func (r SendInvitesRequest) GroupIds(groupIds []int64) SendInvitesRequest {
	r.groupIds = &groupIds
	return r
}

// Boolean indicating whether the newly created user should be subscribed to the [default channels][default-channels] for the organization.  Note that this parameter can be &#x60;true&#x60; even if the user creating the invitation does not generally have permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: New in Zulip 9.0 (feature level 261). Previous versions of Zulip behaved as though this parameter was always &#x60;false&#x60;; clients needed to include the organization&#39;s default channels in the &#x60;stream_ids&#x60; parameter for a newly created user to be automatically subscribed to them.  [default-channels]: /help/set-default-channels-for-new-users [can-subscribe-others]: /help/configure-who-can-invite-to-channels
func (r SendInvitesRequest) IncludeRealmDefaultSubscriptions(includeRealmDefaultSubscriptions bool) SendInvitesRequest {
	r.includeRealmDefaultSubscriptions = &includeRealmDefaultSubscriptions
	return r
}

// A boolean indicating whether the referrer would like to receive a direct message from [notification bot](zulip.com/help/configure-automated-notices) when a user account is created using this invitation.  **Changes**: New in Zulip 9.0 (feature level 267). Previously, referrers always received such direct messages.
func (r SendInvitesRequest) NotifyReferrerOnJoin(notifyReferrerOnJoin bool) SendInvitesRequest {
	r.notifyReferrerOnJoin = &notifyReferrerOnJoin
	return r
}

// Custom message text, in Zulip Markdown format, to be sent by the Welcome Bot to new users that join the organization via this invitation.  Maximum length is 8000 characters.  Only organization administrators can use this feature; for other users, the value is always &#x60;null&#x60;.  - &#x60;null&#x60;: the organization&#39;s default &#x60;welcome_message_custom_text&#x60; is used. - Empty string: no Welcome Bot custom message is sent. - Otherwise, the provided string is the custom message.  **Changes**: New in Zulip 11.0 (feature level 416).
func (r SendInvitesRequest) WelcomeMessageCustomText(welcomeMessageCustomText string) SendInvitesRequest {
	r.welcomeMessageCustomText = &welcomeMessageCustomText
	return r
}

func (r SendInvitesRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.SendInvitesExecute(r)
}

/*
SendInvites Send invitations

Send [invitations](zulip.com/help/invite-new-users) to specified email addresses.

**Changes**: In Zulip 6.0 (feature level 126), the `invite_expires_in_days`
parameter was removed and replaced by `invite_expires_in_minutes`.

In Zulip 5.0 (feature level 117), added support for passing `null` as
the `invite_expires_in_days` parameter to request an invitation that never
expires.

In Zulip 5.0 (feature level 96), the `invite_expires_in_days` parameter was
added which specified the number of days before the invitation would expire.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SendInvitesRequest
*/
func (c *Client) SendInvites(ctx context.Context) SendInvitesRequest {
	return SendInvitesRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *Client) SendInvitesExecute(r SendInvitesRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inviteeEmails == nil {
		return localVarReturnValue, nil, reportError("inviteeEmails is required and must be specified")
	}
	if r.streamIds == nil {
		return localVarReturnValue, nil, reportError("streamIds is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "invitee_emails", r.inviteeEmails, "", "")
	if r.inviteExpiresInMinutes != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "invite_expires_in_minutes", r.inviteExpiresInMinutes, "form", "")
	}
	if r.inviteAs != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "invite_as", r.inviteAs, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "stream_ids", r.streamIds, "form", "multi")
	if r.groupIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_ids", r.groupIds, "form", "multi")
	}
	if r.includeRealmDefaultSubscriptions != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "include_realm_default_subscriptions", r.includeRealmDefaultSubscriptions, "", "")
	}
	if r.notifyReferrerOnJoin != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "notify_referrer_on_join", r.notifyReferrerOnJoin, "", "")
	}
	if r.welcomeMessageCustomText != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "welcome_message_custom_text", r.welcomeMessageCustomText, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v SendInvites400Response
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
