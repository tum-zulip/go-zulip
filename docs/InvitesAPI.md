# \InvitesAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInviteLink**](InvitesAPI.md#CreateInviteLink) | **Post** /invites/multiuse | Create a reusable invitation link
[**GetInvites**](InvitesAPI.md#GetInvites) | **Get** /invites | Get all invitations
[**ResendEmailInvite**](InvitesAPI.md#ResendEmailInvite) | **Post** /invites/{invite_id}/resend | Resend an email invitation
[**RevokeEmailInvite**](InvitesAPI.md#RevokeEmailInvite) | **Delete** /invites/{invite_id} | Revoke an email invitation
[**RevokeInviteLink**](InvitesAPI.md#RevokeInviteLink) | **Delete** /invites/multiuse/{invite_id} | Revoke a reusable invitation link
[**SendInvites**](InvitesAPI.md#SendInvites) | **Post** /invites | Send invitations



## CreateInviteLink

> CreateInviteLink200Response CreateInviteLink(ctx).InviteExpiresInMinutes(inviteExpiresInMinutes).InviteAs(inviteAs).StreamIds(streamIds).GroupIds(groupIds).IncludeRealmDefaultSubscriptions(includeRealmDefaultSubscriptions).WelcomeMessageCustomText(welcomeMessageCustomText).Execute()

Create a reusable invitation link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	inviteExpiresInMinutes := int32(56) // int32 | The number of minutes before the invitation will expire. If `null`, the invitation will never expire. If unspecified, the server will use a default value (based on the `INVITATION_LINK_VALIDITY_MINUTES` server setting, which defaults to 14400, i.e. 10 days) for when the invitation will expire.  **Changes**: New in Zulip 6.0 (feature level 126). Previously, there was an `invite_expires_in_days` parameter, which specified the duration in days instead of minutes.  (optional)
	inviteAs := openapiclient.InviteRoleParameter(100) // InviteRoleParameter |  (optional) (default to 400)
	streamIds := []int32{int32(123)} // []int32 | A list containing the [IDs of the channels](/api/get-stream-id) that the newly created user will be automatically subscribed to if the invitation is accepted, in addition to any default channels that the new user may be subscribed to based on the `include_realm_default_subscriptions` parameter.  Requested channels must either be default channels for the organization, or ones the acting user has permission to add subscribers to.  This list must be empty if the current user has the unlikely configuration of being able to create reusable invitation links while lacking permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: Prior to Zulip 10.0 (feature level 342), default channels that the acting user did not directly have permission to add subscribers to would be rejected.  [can-subscribe-others]: /help/configure-who-can-invite-to-channels  (optional)
	groupIds := []int32{int32(123)} // []int32 | A list containing the [IDs of the user groups](/api/get-user-groups) that the newly created user will be automatically added to if the invitation is accepted. If the list is empty, then the new user will not be added to any user groups. The acting user must have permission to add users to the groups listed in this request.  **Changes**: New in Zulip 10.0 (feature level 322).  (optional)
	includeRealmDefaultSubscriptions := true // bool | Boolean indicating whether the newly created user should be subscribed to the [default channels][default-channels] for the organization.  Note that this parameter can be `true` even if the current user does not generally have permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: New in Zulip 9.0 (feature level 261). Previous versions of Zulip behaved as though this parameter was always `false`; clients needed to include the organization's default channels in the `stream_ids` parameter for a newly created user to be automatically subscribed to them.  [default-channels]: /help/set-default-channels-for-new-users [can-subscribe-others]: /help/configure-who-can-invite-to-channels  (optional) (default to false)
	welcomeMessageCustomText := "welcomeMessageCustomText_example" // string | Custom message text, in Zulip Markdown format, to be sent by the Welcome Bot to new users that join the organization via this invitation.  Maximum length is 8000 characters.  Only organization administrators can use this feature; for other users, the value is always `null`.  - `null`: the organization's default `welcome_message_custom_text` is used. - Empty string: no Welcome Bot custom message is sent. - Otherwise, the provided string is the custom message.  **Changes**: New in Zulip 11.0 (feature level 416).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateInviteLink(context.Background()).InviteExpiresInMinutes(inviteExpiresInMinutes).InviteAs(inviteAs).StreamIds(streamIds).GroupIds(groupIds).IncludeRealmDefaultSubscriptions(includeRealmDefaultSubscriptions).WelcomeMessageCustomText(welcomeMessageCustomText).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.CreateInviteLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInviteLink`: CreateInviteLink200Response
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.CreateInviteLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteExpiresInMinutes** | **int32** | The number of minutes before the invitation will expire. If &#x60;null&#x60;, the invitation will never expire. If unspecified, the server will use a default value (based on the &#x60;INVITATION_LINK_VALIDITY_MINUTES&#x60; server setting, which defaults to 14400, i.e. 10 days) for when the invitation will expire.  **Changes**: New in Zulip 6.0 (feature level 126). Previously, there was an &#x60;invite_expires_in_days&#x60; parameter, which specified the duration in days instead of minutes.  | 
 **inviteAs** | [**InviteRoleParameter**](InviteRoleParameter.md) |  | [default to 400]
 **streamIds** | **[]int32** | A list containing the [IDs of the channels](/api/get-stream-id) that the newly created user will be automatically subscribed to if the invitation is accepted, in addition to any default channels that the new user may be subscribed to based on the &#x60;include_realm_default_subscriptions&#x60; parameter.  Requested channels must either be default channels for the organization, or ones the acting user has permission to add subscribers to.  This list must be empty if the current user has the unlikely configuration of being able to create reusable invitation links while lacking permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: Prior to Zulip 10.0 (feature level 342), default channels that the acting user did not directly have permission to add subscribers to would be rejected.  [can-subscribe-others]: /help/configure-who-can-invite-to-channels  | 
 **groupIds** | **[]int32** | A list containing the [IDs of the user groups](/api/get-user-groups) that the newly created user will be automatically added to if the invitation is accepted. If the list is empty, then the new user will not be added to any user groups. The acting user must have permission to add users to the groups listed in this request.  **Changes**: New in Zulip 10.0 (feature level 322).  | 
 **includeRealmDefaultSubscriptions** | **bool** | Boolean indicating whether the newly created user should be subscribed to the [default channels][default-channels] for the organization.  Note that this parameter can be &#x60;true&#x60; even if the current user does not generally have permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: New in Zulip 9.0 (feature level 261). Previous versions of Zulip behaved as though this parameter was always &#x60;false&#x60;; clients needed to include the organization&#39;s default channels in the &#x60;stream_ids&#x60; parameter for a newly created user to be automatically subscribed to them.  [default-channels]: /help/set-default-channels-for-new-users [can-subscribe-others]: /help/configure-who-can-invite-to-channels  | [default to false]
 **welcomeMessageCustomText** | **string** | Custom message text, in Zulip Markdown format, to be sent by the Welcome Bot to new users that join the organization via this invitation.  Maximum length is 8000 characters.  Only organization administrators can use this feature; for other users, the value is always &#x60;null&#x60;.  - &#x60;null&#x60;: the organization&#39;s default &#x60;welcome_message_custom_text&#x60; is used. - Empty string: no Welcome Bot custom message is sent. - Otherwise, the provided string is the custom message.  **Changes**: New in Zulip 11.0 (feature level 416).  | 

### Return type

[**CreateInviteLink200Response**](CreateInviteLink200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvites

> GetInvites200Response GetInvites(ctx).Execute()

Get all invitations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetInvites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.GetInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvites`: GetInvites200Response
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.GetInvites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitesRequest struct via the builder pattern


### Return type

[**GetInvites200Response**](GetInvites200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEmailInvite

> JsonSuccess ResendEmailInvite(ctx, inviteId).Execute()

Resend an email invitation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	inviteId := int32(1) // int32 | The ID of the email invitation to be resent. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResendEmailInvite(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.ResendEmailInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendEmailInvite`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.ResendEmailInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **int32** | The ID of the email invitation to be resent.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeEmailInvite

> JsonSuccess RevokeEmailInvite(ctx, inviteId).Execute()

Revoke an email invitation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	inviteId := int32(1) // int32 | The ID of the email invitation to be revoked. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevokeEmailInvite(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.RevokeEmailInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeEmailInvite`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.RevokeEmailInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **int32** | The ID of the email invitation to be revoked.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeEmailInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeInviteLink

> JsonSuccess RevokeInviteLink(ctx, inviteId).Execute()

Revoke a reusable invitation link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	inviteId := int32(1) // int32 | The ID of the reusable invitation link to be revoked. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevokeInviteLink(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.RevokeInviteLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeInviteLink`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.RevokeInviteLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **int32** | The ID of the reusable invitation link to be revoked.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInvites

> SendInvites200Response SendInvites(ctx).InviteeEmails(inviteeEmails).StreamIds(streamIds).InviteExpiresInMinutes(inviteExpiresInMinutes).InviteAs(inviteAs).GroupIds(groupIds).IncludeRealmDefaultSubscriptions(includeRealmDefaultSubscriptions).NotifyReferrerOnJoin(notifyReferrerOnJoin).WelcomeMessageCustomText(welcomeMessageCustomText).Execute()

Send invitations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	inviteeEmails := "inviteeEmails_example" // string | The string containing the email addresses, separated by commas or newlines, that will be sent an invitation. 
	streamIds := []int32{int32(123)} // []int32 | A list containing the [IDs of the channels](/api/get-stream-id) that the newly created user will be automatically subscribed to if the invitation is accepted, in addition to any default channels that the new user may be subscribed to based on the `include_realm_default_subscriptions` parameter.  Requested channels must either be default channels for the organization, or ones the acting user has permission to add subscribers to.  This list must be empty if the current user has the unlikely configuration of being able to send invitations while lacking permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: Prior to Zulip 10.0 (feature level 342), default channels that the acting user did not directly have permission to add subscribers to would be rejected.  Before Zulip 7.0 (feature level 180), specifying `stream_ids` as an empty list resulted in an error.  [can-subscribe-others]: /help/configure-who-can-invite-to-channels 
	inviteExpiresInMinutes := int32(56) // int32 | The number of minutes before the invitation will expire. If `null`, the invitation will never expire. If unspecified, the server will use a default value (based on the `INVITATION_LINK_VALIDITY_MINUTES` server setting, which defaults to 14400, i.e. 10 days) for when the invitation will expire.  **Changes**: New in Zulip 6.0 (feature level 126). Previously, there was an `invite_expires_in_days` parameter, which specified the duration in days instead of minutes.  (optional)
	inviteAs := openapiclient.InviteRoleParameter(100) // InviteRoleParameter |  (optional) (default to 400)
	groupIds := []int32{int32(123)} // []int32 | A list containing the [IDs of the user groups](/api/get-user-groups) that the newly created user will be automatically added to if the invitation is accepted. If the list is empty, then the new user will not be added to any user groups. The acting user must have permission to add users to the groups listed in this request.  **Changes**: New in Zulip 10.0 (feature level 322).  (optional)
	includeRealmDefaultSubscriptions := true // bool | Boolean indicating whether the newly created user should be subscribed to the [default channels][default-channels] for the organization.  Note that this parameter can be `true` even if the user creating the invitation does not generally have permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: New in Zulip 9.0 (feature level 261). Previous versions of Zulip behaved as though this parameter was always `false`; clients needed to include the organization's default channels in the `stream_ids` parameter for a newly created user to be automatically subscribed to them.  [default-channels]: /help/set-default-channels-for-new-users [can-subscribe-others]: /help/configure-who-can-invite-to-channels  (optional) (default to false)
	notifyReferrerOnJoin := true // bool | A boolean indicating whether the referrer would like to receive a direct message from [notification bot](/help/configure-automated-notices) when a user account is created using this invitation.  **Changes**: New in Zulip 9.0 (feature level 267). Previously, referrers always received such direct messages.  (optional) (default to true)
	welcomeMessageCustomText := "welcomeMessageCustomText_example" // string | Custom message text, in Zulip Markdown format, to be sent by the Welcome Bot to new users that join the organization via this invitation.  Maximum length is 8000 characters.  Only organization administrators can use this feature; for other users, the value is always `null`.  - `null`: the organization's default `welcome_message_custom_text` is used. - Empty string: no Welcome Bot custom message is sent. - Otherwise, the provided string is the custom message.  **Changes**: New in Zulip 11.0 (feature level 416).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendInvites(context.Background()).InviteeEmails(inviteeEmails).StreamIds(streamIds).InviteExpiresInMinutes(inviteExpiresInMinutes).InviteAs(inviteAs).GroupIds(groupIds).IncludeRealmDefaultSubscriptions(includeRealmDefaultSubscriptions).NotifyReferrerOnJoin(notifyReferrerOnJoin).WelcomeMessageCustomText(welcomeMessageCustomText).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.SendInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendInvites`: SendInvites200Response
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.SendInvites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteeEmails** | **string** | The string containing the email addresses, separated by commas or newlines, that will be sent an invitation.  | 
 **streamIds** | **[]int32** | A list containing the [IDs of the channels](/api/get-stream-id) that the newly created user will be automatically subscribed to if the invitation is accepted, in addition to any default channels that the new user may be subscribed to based on the &#x60;include_realm_default_subscriptions&#x60; parameter.  Requested channels must either be default channels for the organization, or ones the acting user has permission to add subscribers to.  This list must be empty if the current user has the unlikely configuration of being able to send invitations while lacking permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: Prior to Zulip 10.0 (feature level 342), default channels that the acting user did not directly have permission to add subscribers to would be rejected.  Before Zulip 7.0 (feature level 180), specifying &#x60;stream_ids&#x60; as an empty list resulted in an error.  [can-subscribe-others]: /help/configure-who-can-invite-to-channels  | 
 **inviteExpiresInMinutes** | **int32** | The number of minutes before the invitation will expire. If &#x60;null&#x60;, the invitation will never expire. If unspecified, the server will use a default value (based on the &#x60;INVITATION_LINK_VALIDITY_MINUTES&#x60; server setting, which defaults to 14400, i.e. 10 days) for when the invitation will expire.  **Changes**: New in Zulip 6.0 (feature level 126). Previously, there was an &#x60;invite_expires_in_days&#x60; parameter, which specified the duration in days instead of minutes.  | 
 **inviteAs** | [**InviteRoleParameter**](InviteRoleParameter.md) |  | [default to 400]
 **groupIds** | **[]int32** | A list containing the [IDs of the user groups](/api/get-user-groups) that the newly created user will be automatically added to if the invitation is accepted. If the list is empty, then the new user will not be added to any user groups. The acting user must have permission to add users to the groups listed in this request.  **Changes**: New in Zulip 10.0 (feature level 322).  | 
 **includeRealmDefaultSubscriptions** | **bool** | Boolean indicating whether the newly created user should be subscribed to the [default channels][default-channels] for the organization.  Note that this parameter can be &#x60;true&#x60; even if the user creating the invitation does not generally have permission to [subscribe other users to channels][can-subscribe-others].  **Changes**: New in Zulip 9.0 (feature level 261). Previous versions of Zulip behaved as though this parameter was always &#x60;false&#x60;; clients needed to include the organization&#39;s default channels in the &#x60;stream_ids&#x60; parameter for a newly created user to be automatically subscribed to them.  [default-channels]: /help/set-default-channels-for-new-users [can-subscribe-others]: /help/configure-who-can-invite-to-channels  | [default to false]
 **notifyReferrerOnJoin** | **bool** | A boolean indicating whether the referrer would like to receive a direct message from [notification bot](/help/configure-automated-notices) when a user account is created using this invitation.  **Changes**: New in Zulip 9.0 (feature level 267). Previously, referrers always received such direct messages.  | [default to true]
 **welcomeMessageCustomText** | **string** | Custom message text, in Zulip Markdown format, to be sent by the Welcome Bot to new users that join the organization via this invitation.  Maximum length is 8000 characters.  Only organization administrators can use this feature; for other users, the value is always &#x60;null&#x60;.  - &#x60;null&#x60;: the organization&#39;s default &#x60;welcome_message_custom_text&#x60; is used. - Empty string: no Welcome Bot custom message is sent. - Otherwise, the provided string is the custom message.  **Changes**: New in Zulip 11.0 (feature level 416).  | 

### Return type

[**SendInvites200Response**](SendInvites200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

