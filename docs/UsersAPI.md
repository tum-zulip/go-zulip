# \UsersAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAlertWords**](UsersAPI.md#AddAlertWords) | **Post** /users/me/alert_words | Add alert words
[**AddApnsToken**](UsersAPI.md#AddApnsToken) | **Post** /users/me/apns_device_token | Add an APNs device token
[**AddFcmToken**](UsersAPI.md#AddFcmToken) | **Post** /users/me/android_gcm_reg_id | Add an FCM registration token
[**CreateUser**](UsersAPI.md#CreateUser) | **Post** /users | Create a user
[**CreateUserGroup**](UsersAPI.md#CreateUserGroup) | **Post** /user_groups/create | Create a user group
[**DeactivateOwnUser**](UsersAPI.md#DeactivateOwnUser) | **Delete** /users/me | Deactivate own user
[**DeactivateUser**](UsersAPI.md#DeactivateUser) | **Delete** /users/{user_id} | Deactivate a user
[**DeactivateUserGroup**](UsersAPI.md#DeactivateUserGroup) | **Post** /user_groups/{user_group_id}/deactivate | Deactivate a user group
[**GetAlertWords**](UsersAPI.md#GetAlertWords) | **Get** /users/me/alert_words | Get all alert words
[**GetAttachments**](UsersAPI.md#GetAttachments) | **Get** /attachments | Get attachments
[**GetIsUserGroupMember**](UsersAPI.md#GetIsUserGroupMember) | **Get** /user_groups/{user_group_id}/members/{user_id} | Get user group membership status
[**GetOwnUser**](UsersAPI.md#GetOwnUser) | **Get** /users/me | Get own user
[**GetUser**](UsersAPI.md#GetUser) | **Get** /users/{user_id} | Get a user
[**GetUserByEmail**](UsersAPI.md#GetUserByEmail) | **Get** /users/{email} | Get a user by email
[**GetUserGroupMembers**](UsersAPI.md#GetUserGroupMembers) | **Get** /user_groups/{user_group_id}/members | Get user group members
[**GetUserGroupSubgroups**](UsersAPI.md#GetUserGroupSubgroups) | **Get** /user_groups/{user_group_id}/subgroups | Get subgroups of a user group
[**GetUserGroups**](UsersAPI.md#GetUserGroups) | **Get** /user_groups | Get user groups
[**GetUserPresence**](UsersAPI.md#GetUserPresence) | **Get** /users/{user_id_or_email}/presence | Get a user&#39;s presence
[**GetUserStatus**](UsersAPI.md#GetUserStatus) | **Get** /users/{user_id}/status | Get a user&#39;s status
[**GetUsers**](UsersAPI.md#GetUsers) | **Get** /users | Get users
[**MuteUser**](UsersAPI.md#MuteUser) | **Post** /users/me/muted_users/{muted_user_id} | Mute a user
[**ReactivateUser**](UsersAPI.md#ReactivateUser) | **Post** /users/{user_id}/reactivate | Reactivate a user
[**RemoveAlertWords**](UsersAPI.md#RemoveAlertWords) | **Delete** /users/me/alert_words | Remove alert words
[**RemoveApnsToken**](UsersAPI.md#RemoveApnsToken) | **Delete** /users/me/apns_device_token | Remove an APNs device token
[**RemoveAttachment**](UsersAPI.md#RemoveAttachment) | **Delete** /attachments/{attachment_id} | Delete an attachment
[**RemoveFcmToken**](UsersAPI.md#RemoveFcmToken) | **Delete** /users/me/android_gcm_reg_id | Remove an FCM registration token
[**SetTypingStatus**](UsersAPI.md#SetTypingStatus) | **Post** /typing | Set \&quot;typing\&quot; status
[**SetTypingStatusForMessageEdit**](UsersAPI.md#SetTypingStatusForMessageEdit) | **Post** /messages/{message_id}/typing | Set \&quot;typing\&quot; status for message editing
[**UnmuteUser**](UsersAPI.md#UnmuteUser) | **Delete** /users/me/muted_users/{muted_user_id} | Unmute a user
[**UpdatePresence**](UsersAPI.md#UpdatePresence) | **Post** /users/me/presence | Update your presence
[**UpdateSettings**](UsersAPI.md#UpdateSettings) | **Patch** /settings | Update settings
[**UpdateStatus**](UsersAPI.md#UpdateStatus) | **Post** /users/me/status | Update your status
[**UpdateStatusForUser**](UsersAPI.md#UpdateStatusForUser) | **Post** /users/{user_id}/status | Update user status
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Patch** /users/{user_id} | Update a user
[**UpdateUserByEmail**](UsersAPI.md#UpdateUserByEmail) | **Patch** /users/{email} | Update a user by email
[**UpdateUserGroup**](UsersAPI.md#UpdateUserGroup) | **Patch** /user_groups/{user_group_id} | Update a user group
[**UpdateUserGroupMembers**](UsersAPI.md#UpdateUserGroupMembers) | **Post** /user_groups/{user_group_id}/members | Update user group members
[**UpdateUserGroupSubgroups**](UsersAPI.md#UpdateUserGroupSubgroups) | **Post** /user_groups/{user_group_id}/subgroups | Update subgroups of a user group



## AddAlertWords

> AddAlertWords200Response AddAlertWords(ctx).AlertWords(alertWords).Execute()

Add alert words



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
	alertWords := []string{"Inner_example"} // []string | An array of strings to be added to the user's set of configured alert words. Strings already present in the user's set of alert words already are ignored.  Alert words are case insensitive. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddAlertWords(context.Background()).AlertWords(alertWords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddAlertWords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAlertWords`: AddAlertWords200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddAlertWords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAlertWordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertWords** | **[]string** | An array of strings to be added to the user&#39;s set of configured alert words. Strings already present in the user&#39;s set of alert words already are ignored.  Alert words are case insensitive.  | 

### Return type

[**AddAlertWords200Response**](AddAlertWords200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddApnsToken

> JsonSuccess AddApnsToken(ctx).Token(token).Appid(appid).Execute()

Add an APNs device token



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
	token := "token_example" // string | The token provided by the device. 
	appid := "appid_example" // string | The ID of the Zulip app that is making the request.  **Changes**: In Zulip 8.0 (feature level 223), this parameter was made required. Previously, if it was unspecified, the server would use a default value (based on the `ZULIP_IOS_APP_ID` server setting, which defaulted to `\\\"org.zulip.Zulip\\\"`). 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddApnsToken(context.Background()).Token(token).Appid(appid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddApnsToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddApnsToken`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddApnsToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddApnsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token provided by the device.  | 
 **appid** | **string** | The ID of the Zulip app that is making the request.  **Changes**: In Zulip 8.0 (feature level 223), this parameter was made required. Previously, if it was unspecified, the server would use a default value (based on the &#x60;ZULIP_IOS_APP_ID&#x60; server setting, which defaulted to &#x60;\\\&quot;org.zulip.Zulip\\\&quot;&#x60;).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFcmToken

> JsonSuccess AddFcmToken(ctx).Token(token).Execute()

Add an FCM registration token



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
	token := "token_example" // string | The token provided by the device. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddFcmToken(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddFcmToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFcmToken`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddFcmToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFcmTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token provided by the device.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> CreateUser200Response CreateUser(ctx).Email(email).Password(password).FullName(fullName).Execute()

Create a user



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
	email := "email_example" // string | The email address of the new user. 
	password := "password_example" // string | The password of the new user. 
	fullName := "fullName_example" // string | The full name of the new user. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateUser(context.Background()).Email(email).Password(password).FullName(fullName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: CreateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | The email address of the new user.  | 
 **password** | **string** | The password of the new user.  | 
 **fullName** | **string** | The full name of the new user.  | 

### Return type

[**CreateUser200Response**](CreateUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserGroup

> CreateUserGroup200Response CreateUserGroup(ctx).Name(name).Description(description).Members(members).Subgroups(subgroups).CanAddMembersGroup(canAddMembersGroup).CanJoinGroup(canJoinGroup).CanLeaveGroup(canLeaveGroup).CanManageGroup(canManageGroup).CanMentionGroup(canMentionGroup).CanRemoveMembersGroup(canRemoveMembersGroup).Execute()

Create a user group



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
	name := "name_example" // string | The name of the user group. 
	description := "description_example" // string | The description of the user group. 
	members := []int32{int32(123)} // []int32 | An array containing the user IDs of the initial members for the new user group. 
	subgroups := []int32{int32(123)} // []int32 | An array containing the IDs of the initial subgroups for the new user group.  User can add subgroups to the new group irrespective of other permissions for the new group.  **Changes**: New in Zulip 10.0 (feature level 311).  (optional)
	canAddMembersGroup := openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()} // GroupSettingValue | A [group-setting value][setting-values] defining the set of users who have permission to add members to this user group.  **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the `can_manage_group` setting.  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups  (optional)
	canJoinGroup := openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()} // GroupSettingValue | A [group-setting value][setting-values] defining the set of users who have permission to join this user group.  **Changes**: New in Zulip 10.0 (feature level 301).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups  (optional)
	canLeaveGroup := openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()} // GroupSettingValue | A [group-setting value][setting-values] defining the set of users who have permission to leave this user group.  **Changes**: New in Zulip 10.0 (feature level 308).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups  (optional)
	canManageGroup := openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()} // GroupSettingValue | A [group-setting value][setting-values] defining the set of users who have permission to [manage this user group][manage-user-groups].  This setting cannot be set to `\\\"role:internet\\\"` and `\\\"role:everyone\\\"` [system groups][system-groups].  **Changes**: New in Zulip 10.0 (feature level 283).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups [manage-user-groups]: /help/manage-user-groups  (optional)
	canMentionGroup := openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()} // GroupSettingValue | A [group-setting value][setting-values] defining the set of users who have permission to [mention this user group][mentions].  This setting cannot be set to `\\\"role:internet\\\"` and `\\\"role:owners\\\"` [system groups][system-groups].  Before Zulip 9.0 (feature level 258), this parameter could only be the integer form of a [group-setting value][setting-values].  Before Zulip 8.0 (feature level 198), this parameter was named `can_mention_group_id`.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups][system-groups].  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups [mentions]: /help/mention-a-user-or-group  (optional)
	canRemoveMembersGroup := openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()} // GroupSettingValue | A [group-setting value][setting-values] defining the set of users who have permission to remove members from this user group.  **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the `can_manage_group` setting.  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateUserGroup(context.Background()).Name(name).Description(description).Members(members).Subgroups(subgroups).CanAddMembersGroup(canAddMembersGroup).CanJoinGroup(canJoinGroup).CanLeaveGroup(canLeaveGroup).CanManageGroup(canManageGroup).CanMentionGroup(canMentionGroup).CanRemoveMembersGroup(canRemoveMembersGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserGroup`: CreateUserGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the user group.  | 
 **description** | **string** | The description of the user group.  | 
 **members** | **[]int32** | An array containing the user IDs of the initial members for the new user group.  | 
 **subgroups** | **[]int32** | An array containing the IDs of the initial subgroups for the new user group.  User can add subgroups to the new group irrespective of other permissions for the new group.  **Changes**: New in Zulip 10.0 (feature level 311).  | 
 **canAddMembersGroup** | [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to add members to this user group.  **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups  | 
 **canJoinGroup** | [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to join this user group.  **Changes**: New in Zulip 10.0 (feature level 301).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups  | 
 **canLeaveGroup** | [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to leave this user group.  **Changes**: New in Zulip 10.0 (feature level 308).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups  | 
 **canManageGroup** | [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to [manage this user group][manage-user-groups].  This setting cannot be set to &#x60;\\\&quot;role:internet\\\&quot;&#x60; and &#x60;\\\&quot;role:everyone\\\&quot;&#x60; [system groups][system-groups].  **Changes**: New in Zulip 10.0 (feature level 283).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups [manage-user-groups]: /help/manage-user-groups  | 
 **canMentionGroup** | [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to [mention this user group][mentions].  This setting cannot be set to &#x60;\\\&quot;role:internet\\\&quot;&#x60; and &#x60;\\\&quot;role:owners\\\&quot;&#x60; [system groups][system-groups].  Before Zulip 9.0 (feature level 258), this parameter could only be the integer form of a [group-setting value][setting-values].  Before Zulip 8.0 (feature level 198), this parameter was named &#x60;can_mention_group_id&#x60;.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups][system-groups].  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups [mentions]: /help/mention-a-user-or-group  | 
 **canRemoveMembersGroup** | [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to remove members from this user group.  **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups  | 

### Return type

[**CreateUserGroup200Response**](CreateUserGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateOwnUser

> JsonSuccess DeactivateOwnUser(ctx).Execute()

Deactivate own user



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
	resp, r, err := apiClient.DeactivateOwnUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeactivateOwnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateOwnUser`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeactivateOwnUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateOwnUserRequest struct via the builder pattern


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


## DeactivateUser

> JsonSuccess DeactivateUser(ctx, userId).DeactivationNotificationComment(deactivationNotificationComment).Execute()

Deactivate a user



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
	userId := int32(12) // int32 | The target user's ID. 
	deactivationNotificationComment := "deactivationNotificationComment_example" // string | If not `null`, requests that the deactivated user receive a notification email about their account deactivation.  If not `\\\"\\\"`, encodes custom text written by the administrator to be included in the notification email.  **Changes**: New in Zulip 5.0 (feature level 135).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeactivateUser(context.Background(), userId).DeactivationNotificationComment(deactivationNotificationComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeactivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateUser`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeactivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The target user&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deactivationNotificationComment** | **string** | If not &#x60;null&#x60;, requests that the deactivated user receive a notification email about their account deactivation.  If not &#x60;\\\&quot;\\\&quot;&#x60;, encodes custom text written by the administrator to be included in the notification email.  **Changes**: New in Zulip 5.0 (feature level 135).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateUserGroup

> JsonSuccess DeactivateUserGroup(ctx, userGroupId).Execute()

Deactivate a user group



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
	userGroupId := int32(38) // int32 | The ID of the target user group. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeactivateUserGroup(context.Background(), userGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeactivateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateUserGroup`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeactivateUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** | The ID of the target user group.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateUserGroupRequest struct via the builder pattern


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


## GetAlertWords

> GetAlertWords200Response GetAlertWords(ctx).Execute()

Get all alert words



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
	resp, r, err := apiClient.GetAlertWords(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetAlertWords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertWords`: GetAlertWords200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetAlertWords`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertWordsRequest struct via the builder pattern


### Return type

[**GetAlertWords200Response**](GetAlertWords200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachments

> GetAttachments200Response GetAttachments(ctx).Execute()

Get attachments



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
	resp, r, err := apiClient.GetAttachments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachments`: GetAttachments200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetAttachments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentsRequest struct via the builder pattern


### Return type

[**GetAttachments200Response**](GetAttachments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsUserGroupMember

> GetIsUserGroupMember200Response GetIsUserGroupMember(ctx, userGroupId, userId).DirectMemberOnly(directMemberOnly).Execute()

Get user group membership status



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
	userGroupId := int32(38) // int32 | The ID of the target user group. 
	userId := int32(12) // int32 | The target user's ID. 
	directMemberOnly := false // bool | Whether to consider only the direct members of user group and not members of its subgroups. Default is `false`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetIsUserGroupMember(context.Background(), userGroupId, userId).DirectMemberOnly(directMemberOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetIsUserGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsUserGroupMember`: GetIsUserGroupMember200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetIsUserGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** | The ID of the target user group.  | 
**userId** | **int32** | The target user&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIsUserGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **directMemberOnly** | **bool** | Whether to consider only the direct members of user group and not members of its subgroups. Default is &#x60;false&#x60;.  | 

### Return type

[**GetIsUserGroupMember200Response**](GetIsUserGroupMember200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnUser

> GetOwnUser200Response GetOwnUser(ctx).Execute()

Get own user



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
	resp, r, err := apiClient.GetOwnUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetOwnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnUser`: GetOwnUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetOwnUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnUserRequest struct via the builder pattern


### Return type

[**GetOwnUser200Response**](GetOwnUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> GetUserByEmail200Response GetUser(ctx, userId).ClientGravatar(clientGravatar).IncludeCustomProfileFields(includeCustomProfileFields).Execute()

Get a user



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
	userId := int32(12) // int32 | The target user's ID. 
	clientGravatar := false // bool | Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.  **Changes**: The default value of this parameter was `false` prior to Zulip 5.0 (feature level 92).  (optional) (default to true)
	includeCustomProfileFields := true // bool | Whether the client wants [custom profile field](/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetUser(context.Background(), userId).ClientGravatar(clientGravatar).IncludeCustomProfileFields(includeCustomProfileFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: GetUserByEmail200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The target user&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientGravatar** | **bool** | Whether the client supports computing gravatars URLs. If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  **Changes**: The default value of this parameter was &#x60;false&#x60; prior to Zulip 5.0 (feature level 92).  | [default to true]
 **includeCustomProfileFields** | **bool** | Whether the client wants [custom profile field](/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.  | [default to false]

### Return type

[**GetUserByEmail200Response**](GetUserByEmail200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByEmail

> GetUserByEmail200Response GetUserByEmail(ctx, email).ClientGravatar(clientGravatar).IncludeCustomProfileFields(includeCustomProfileFields).Execute()

Get a user by email



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
	email := "iago@zulip.com" // string | The email address of the user to fetch. Two forms are supported:  - The real email address of the user (`delivery_email`). The lookup will   succeed if and only if the user exists and their email address visibility   setting permits the client to see the email address.  - The dummy Zulip API email address of the form `user{user_id}@{realm_host}`. This   is identical to simply [getting user by ID](/api/get-user). If the server or   realm change domains, the dummy email address used has to be adjustment to   match the new realm domain. This is legacy behavior for   backwards-compatibility, and will be removed in a future release.  **Changes**: Starting with Zulip 10.0 (feature level 302), lookups by real email address match the semantics of the target's email visibility setting and dummy email addresses work for all users, independently of their email visibility setting.  Previously, lookups were done only using the Zulip API email addresses. 
	clientGravatar := false // bool | Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.  **Changes**: The default value of this parameter was `false` prior to Zulip 5.0 (feature level 92).  (optional) (default to true)
	includeCustomProfileFields := true // bool | Whether the client wants [custom profile field](/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetUserByEmail(context.Background(), email).ClientGravatar(clientGravatar).IncludeCustomProfileFields(includeCustomProfileFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserByEmail`: GetUserByEmail200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | The email address of the user to fetch. Two forms are supported:  - The real email address of the user (&#x60;delivery_email&#x60;). The lookup will   succeed if and only if the user exists and their email address visibility   setting permits the client to see the email address.  - The dummy Zulip API email address of the form &#x60;user{user_id}@{realm_host}&#x60;. This   is identical to simply [getting user by ID](/api/get-user). If the server or   realm change domains, the dummy email address used has to be adjustment to   match the new realm domain. This is legacy behavior for   backwards-compatibility, and will be removed in a future release.  **Changes**: Starting with Zulip 10.0 (feature level 302), lookups by real email address match the semantics of the target&#39;s email visibility setting and dummy email addresses work for all users, independently of their email visibility setting.  Previously, lookups were done only using the Zulip API email addresses.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientGravatar** | **bool** | Whether the client supports computing gravatars URLs. If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  **Changes**: The default value of this parameter was &#x60;false&#x60; prior to Zulip 5.0 (feature level 92).  | [default to true]
 **includeCustomProfileFields** | **bool** | Whether the client wants [custom profile field](/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.  | [default to false]

### Return type

[**GetUserByEmail200Response**](GetUserByEmail200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupMembers

> GetUserGroupMembers200Response GetUserGroupMembers(ctx, userGroupId).DirectMemberOnly(directMemberOnly).Execute()

Get user group members



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
	userGroupId := int32(38) // int32 | The ID of the target user group. 
	directMemberOnly := false // bool | Whether to consider only the direct members of user group and not members of its subgroups. Default is `false`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetUserGroupMembers(context.Background(), userGroupId).DirectMemberOnly(directMemberOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupMembers`: GetUserGroupMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** | The ID of the target user group.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **directMemberOnly** | **bool** | Whether to consider only the direct members of user group and not members of its subgroups. Default is &#x60;false&#x60;.  | 

### Return type

[**GetUserGroupMembers200Response**](GetUserGroupMembers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupSubgroups

> GetUserGroupSubgroups200Response GetUserGroupSubgroups(ctx, userGroupId).DirectSubgroupOnly(directSubgroupOnly).Execute()

Get subgroups of a user group



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
	userGroupId := int32(38) // int32 | The ID of the target user group. 
	directSubgroupOnly := true // bool | Whether to consider only direct subgroups of the user group or subgroups of subgroups also.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetUserGroupSubgroups(context.Background(), userGroupId).DirectSubgroupOnly(directSubgroupOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroupSubgroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupSubgroups`: GetUserGroupSubgroups200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroupSubgroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** | The ID of the target user group.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupSubgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **directSubgroupOnly** | **bool** | Whether to consider only direct subgroups of the user group or subgroups of subgroups also.  | [default to false]

### Return type

[**GetUserGroupSubgroups200Response**](GetUserGroupSubgroups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroups

> GetUserGroups200Response GetUserGroups(ctx).IncludeDeactivatedGroups(includeDeactivatedGroups).Execute()

Get user groups



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
	includeDeactivatedGroups := true // bool | Whether to include deactivated user groups in the response.  **Changes**: In Zulip 10.0 (feature level 294), renamed `allow_deactivated` to `include_deactivated_groups`.  New in Zulip 10.0 (feature level 290). Previously, deactivated user groups did not exist and thus would never be included in the response.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetUserGroups(context.Background()).IncludeDeactivatedGroups(includeDeactivatedGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroups`: GetUserGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDeactivatedGroups** | **bool** | Whether to include deactivated user groups in the response.  **Changes**: In Zulip 10.0 (feature level 294), renamed &#x60;allow_deactivated&#x60; to &#x60;include_deactivated_groups&#x60;.  New in Zulip 10.0 (feature level 290). Previously, deactivated user groups did not exist and thus would never be included in the response.  | [default to false]

### Return type

[**GetUserGroups200Response**](GetUserGroups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPresence

> GetUserPresence200Response GetUserPresence(ctx, userIdOrEmail).Execute()

Get a user's presence



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
	userIdOrEmail := "iago@zulip.com" // string | The ID or Zulip API email address of the user whose presence you want to fetch.  **Changes**: New in Zulip 4.0 (feature level 43). Previous versions only supported identifying the user by Zulip API email. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetUserPresence(context.Background(), userIdOrEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserPresence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPresence`: GetUserPresence200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserPresence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrEmail** | **string** | The ID or Zulip API email address of the user whose presence you want to fetch.  **Changes**: New in Zulip 4.0 (feature level 43). Previous versions only supported identifying the user by Zulip API email.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserPresence200Response**](GetUserPresence200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserStatus

> GetUserStatus200Response GetUserStatus(ctx, userId).Execute()

Get a user's status



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
	userId := int32(12) // int32 | The target user's ID. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetUserStatus(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserStatus`: GetUserStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The target user&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserStatus200Response**](GetUserStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> GetUsers200Response GetUsers(ctx).ClientGravatar(clientGravatar).IncludeCustomProfileFields(includeCustomProfileFields).UserIds(userIds).Execute()

Get users



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
	clientGravatar := false // bool | Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.  **Changes**: The default value of this parameter was `false` prior to Zulip 5.0 (feature level 92).  (optional) (default to true)
	includeCustomProfileFields := true // bool | Whether the client wants [custom profile field](/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.  (optional) (default to false)
	userIds := []int32{int32(123)} // []int32 | Limits the results to the specified user IDs. If not provided, the server will return all accessible users in the organization.  **Changes**: New in Zulip 11.0 (feature level 384).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetUsers(context.Background()).ClientGravatar(clientGravatar).IncludeCustomProfileFields(includeCustomProfileFields).UserIds(userIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: GetUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGravatar** | **bool** | Whether the client supports computing gravatars URLs. If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  **Changes**: The default value of this parameter was &#x60;false&#x60; prior to Zulip 5.0 (feature level 92).  | [default to true]
 **includeCustomProfileFields** | **bool** | Whether the client wants [custom profile field](/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.  | [default to false]
 **userIds** | **[]int32** | Limits the results to the specified user IDs. If not provided, the server will return all accessible users in the organization.  **Changes**: New in Zulip 11.0 (feature level 384).  | 

### Return type

[**GetUsers200Response**](GetUsers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteUser

> JsonSuccess MuteUser(ctx, mutedUserId).Execute()

Mute a user



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
	mutedUserId := int32(10) // int32 | The ID of the user to mute/unmute.  **Changes**: Before Zulip 8.0 (feature level 188), bot users could not be muted/unmuted, and specifying a bot user's ID returned an error response. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MuteUser(context.Background(), mutedUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.MuteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MuteUser`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.MuteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mutedUserId** | **int32** | The ID of the user to mute/unmute.  **Changes**: Before Zulip 8.0 (feature level 188), bot users could not be muted/unmuted, and specifying a bot user&#39;s ID returned an error response.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMuteUserRequest struct via the builder pattern


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


## ReactivateUser

> JsonSuccess ReactivateUser(ctx, userId).Execute()

Reactivate a user



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
	userId := int32(12) // int32 | The target user's ID. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactivateUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ReactivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactivateUser`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ReactivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The target user&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateUserRequest struct via the builder pattern


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


## RemoveAlertWords

> RemoveAlertWords200Response RemoveAlertWords(ctx).AlertWords(alertWords).Execute()

Remove alert words



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
	alertWords := []string{"Inner_example"} // []string | An array of strings to be removed from the user's set of configured alert words. Strings that are not in the user's set of alert words are ignored. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveAlertWords(context.Background()).AlertWords(alertWords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveAlertWords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAlertWords`: RemoveAlertWords200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RemoveAlertWords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAlertWordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertWords** | **[]string** | An array of strings to be removed from the user&#39;s set of configured alert words. Strings that are not in the user&#39;s set of alert words are ignored.  | 

### Return type

[**RemoveAlertWords200Response**](RemoveAlertWords200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveApnsToken

> JsonSuccess RemoveApnsToken(ctx).Token(token).Execute()

Remove an APNs device token



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
	token := "token_example" // string | The token provided by the device. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveApnsToken(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveApnsToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveApnsToken`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RemoveApnsToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveApnsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token provided by the device.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAttachment

> JsonSuccess RemoveAttachment(ctx, attachmentId).Execute()

Delete an attachment



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
	attachmentId := int32(1) // int32 | The ID of the attachment to be deleted. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveAttachment(context.Background(), attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAttachment`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RemoveAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **int32** | The ID of the attachment to be deleted.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAttachmentRequest struct via the builder pattern


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


## RemoveFcmToken

> JsonSuccess RemoveFcmToken(ctx).Token(token).Execute()

Remove an FCM registration token



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
	token := "token_example" // string | The token provided by the device. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveFcmToken(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveFcmToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFcmToken`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RemoveFcmToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFcmTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token provided by the device.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTypingStatus

> JsonSuccess SetTypingStatus(ctx).Op(op).Type_(type_).To(to).StreamId(streamId).Topic(topic).Execute()

Set \"typing\" status



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
	op := "op_example" // string | Whether the user has started (`\\\"start\\\"`) or stopped (`\\\"stop\\\"`) typing. 
	type_ := "type__example" // string | Type of the message being composed.  **Changes**: In Zulip 9.0 (feature level 248), `\\\"channel\\\"` was added as an additional value for this parameter to indicate a channel message is being composed.  In Zulip 8.0 (feature level 215), stopped supporting `\\\"private\\\"` as a valid value for this parameter.  In Zulip 7.0 (feature level 174), `\\\"direct\\\"` was added as the preferred way to indicate a direct message is being composed, becoming the default value for this parameter and deprecating the original `\\\"private\\\"`.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  (optional) (default to "direct")
	to := []int32{int32(123)} // []int32 | User IDs of the recipients of the message being typed. Required for the `\\\"direct\\\"` type. Ignored in the case of `\\\"stream\\\"` or `\\\"channel\\\"` type.  Clients should send a JSON-encoded list of user IDs, even if there is only one recipient.  **Changes**: In Zulip 8.0 (feature level 215), stopped using this parameter for the `\\\"stream\\\"` type. Previously, in the case of the `\\\"stream\\\"` type, it accepted a single-element list containing the ID of the channel. A new parameter, `stream_id`, is now used for this. Note that the `\\\"channel\\\"` type did not exist at this feature level.  Support for typing notifications for channel' messages is new in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  Before Zulip 2.0.0, this parameter accepted only a JSON-encoded list of email addresses. Support for the email address-based format was removed in Zulip 3.0 (feature level 11).  (optional)
	streamId := int32(56) // int32 | ID of the channel in which the message is being typed. Required for the `\\\"stream\\\"` or `\\\"channel\\\"` type. Ignored in the case of `\\\"direct\\\"` type.  **Changes**: New in Zulip 8.0 (feature level 215). Previously, a single-element list containing the ID of the channel was passed in `to` parameter.  (optional)
	topic := "topic_example" // string | Topic to which message is being typed. Required for the `\\\"stream\\\"` or `\\\"channel\\\"` type. Ignored in the case of `\\\"direct\\\"` type.  Note: When `\\\"(no topic)\\\"` or the value of `realm_empty_topic_display_name` found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 372), `\\\"(no topic)\\\"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetTypingStatus(context.Background()).Op(op).Type_(type_).To(to).StreamId(streamId).Topic(topic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SetTypingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTypingStatus`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SetTypingStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTypingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **op** | **string** | Whether the user has started (&#x60;\\\&quot;start\\\&quot;&#x60;) or stopped (&#x60;\\\&quot;stop\\\&quot;&#x60;) typing.  | 
 **type_** | **string** | Type of the message being composed.  **Changes**: In Zulip 9.0 (feature level 248), &#x60;\\\&quot;channel\\\&quot;&#x60; was added as an additional value for this parameter to indicate a channel message is being composed.  In Zulip 8.0 (feature level 215), stopped supporting &#x60;\\\&quot;private\\\&quot;&#x60; as a valid value for this parameter.  In Zulip 7.0 (feature level 174), &#x60;\\\&quot;direct\\\&quot;&#x60; was added as the preferred way to indicate a direct message is being composed, becoming the default value for this parameter and deprecating the original &#x60;\\\&quot;private\\\&quot;&#x60;.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  | [default to &quot;direct&quot;]
 **to** | **[]int32** | User IDs of the recipients of the message being typed. Required for the &#x60;\\\&quot;direct\\\&quot;&#x60; type. Ignored in the case of &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; type.  Clients should send a JSON-encoded list of user IDs, even if there is only one recipient.  **Changes**: In Zulip 8.0 (feature level 215), stopped using this parameter for the &#x60;\\\&quot;stream\\\&quot;&#x60; type. Previously, in the case of the &#x60;\\\&quot;stream\\\&quot;&#x60; type, it accepted a single-element list containing the ID of the channel. A new parameter, &#x60;stream_id&#x60;, is now used for this. Note that the &#x60;\\\&quot;channel\\\&quot;&#x60; type did not exist at this feature level.  Support for typing notifications for channel&#39; messages is new in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  Before Zulip 2.0.0, this parameter accepted only a JSON-encoded list of email addresses. Support for the email address-based format was removed in Zulip 3.0 (feature level 11).  | 
 **streamId** | **int32** | ID of the channel in which the message is being typed. Required for the &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; type. Ignored in the case of &#x60;\\\&quot;direct\\\&quot;&#x60; type.  **Changes**: New in Zulip 8.0 (feature level 215). Previously, a single-element list containing the ID of the channel was passed in &#x60;to&#x60; parameter.  | 
 **topic** | **string** | Topic to which message is being typed. Required for the &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; type. Ignored in the case of &#x60;\\\&quot;direct\\\&quot;&#x60; type.  Note: When &#x60;\\\&quot;(no topic)\\\&quot;&#x60; or the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 372), &#x60;\\\&quot;(no topic)\\\&quot;&#x60; was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTypingStatusForMessageEdit

> JsonSuccess SetTypingStatusForMessageEdit(ctx, messageId).Op(op).Execute()

Set \"typing\" status for message editing



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
	messageId := int32(47) // int32 | The target message's ID. 
	op := "op_example" // string | Whether the user has started (`\\\"start\\\"`) or stopped (`\\\"stop\\\"`) editing. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetTypingStatusForMessageEdit(context.Background(), messageId).Op(op).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SetTypingStatusForMessageEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTypingStatusForMessageEdit`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SetTypingStatusForMessageEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTypingStatusForMessageEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **op** | **string** | Whether the user has started (&#x60;\\\&quot;start\\\&quot;&#x60;) or stopped (&#x60;\\\&quot;stop\\\&quot;&#x60;) editing.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmuteUser

> JsonSuccess UnmuteUser(ctx, mutedUserId).Execute()

Unmute a user



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
	mutedUserId := int32(10) // int32 | The ID of the user to mute/unmute.  **Changes**: Before Zulip 8.0 (feature level 188), bot users could not be muted/unmuted, and specifying a bot user's ID returned an error response. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnmuteUser(context.Background(), mutedUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UnmuteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnmuteUser`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UnmuteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mutedUserId** | **int32** | The ID of the user to mute/unmute.  **Changes**: Before Zulip 8.0 (feature level 188), bot users could not be muted/unmuted, and specifying a bot user&#39;s ID returned an error response.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteUserRequest struct via the builder pattern


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


## UpdatePresence

> UpdatePresence200Response UpdatePresence(ctx).Status(status).LastUpdateId(lastUpdateId).HistoryLimitDays(historyLimitDays).NewUserInput(newUserInput).PingOnly(pingOnly).SlimPresence(slimPresence).Execute()

Update your presence



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
	status := "status_example" // string | The status of the user on this client.  Clients should report the user as `\\\"active\\\"` on this device if the client knows that the user is presently using the device (and thus would potentially see a notification immediately), even if the user has not directly interacted with the Zulip client.  Otherwise, it should report the user as `\\\"idle\\\"`.  See the related [`new_user_input`](#parameter-new_user_input) parameter for how a client should report whether the user is actively using the Zulip client. 
	lastUpdateId := int32(56) // int32 | The identifier that specifies what presence data the client already has received, which allows the server to only return more recent user presence data.  This should be set to `-1` during initialization of the client in order to fetch all user presence data, unless the client is obtaining initial user presence metadata from the [`POST /register`](/api/register-queue) endpoint.  In subsequent queries to this endpoint, this value should be set to the most recent value of `presence_last_update_id` returned by the server in this endpoint's response, which implements incremental fetching of user presence data.  When this parameter is passed, the user presence data in the response will always be in the modern format.  **Changes**: New in Zulip 9.0 (feature level 263). Previously, the server sent user presence data for all users who had been active in the last two weeks unconditionally.  (optional)
	historyLimitDays := int32(56) // int32 | Limits how far back in time to fetch user presence data. If not specified, defaults to 14 days. A value of N means that the oldest presence data fetched will be from at most N days ago.  Note that this is only useful during the initial user presence data fetch, as subsequent fetches should use the `last_update_id` parameter, which will act as the limit on how much presence data is returned. `history_limit_days` is ignored if `last_update_id` is passed with a value greater than `0`, indicating that the client already has some presence data.  **Changes**: New in Zulip 10.0 (feature level 288).  (optional)
	newUserInput := true // bool | Whether the user has interacted with the client (e.g. moved the mouse, used the keyboard, etc.) since the previous presence request from this client.  The server uses data from this parameter to implement certain [usage statistics](/help/analytics).  User interface clients that might run in the background, without the user ever interacting with them, should be careful to only pass `true` if the user has actually interacted with the client in order to avoid corrupting usage statistics graphs.  (optional) (default to false)
	pingOnly := true // bool | Whether the client is sending a ping-only request, meaning it only wants to update the user's presence `status` on the server.  Otherwise, also requests the server return user presence data for all users in the organization, which is further specified by the [`last_update_id`](#parameter-last_update_id) parameter.  (optional) (default to false)
	slimPresence := true // bool | Legacy parameter for configuring the format (modern or legacy) in which the server will return user presence data for the organization.  Modern clients should use [`last_update_id`](#parameter-last_update_id), which guarantees that user presence data will be returned in the modern format, and should not pass this parameter as `true` unless interacting with an older server.  Legacy clients that do not yet support `last_update_id` may use the value of `true` to request the modern format for user presence data.  **Note**: The legacy format for user presence data will be removed entirely in a future release.  **Changes**: **Deprecated** in Zulip 9.0 (feature level 263). Using the modern `last_update_id` parameter is the recommended way to request the modern format for user presence data.  New in Zulip 3.0 (no feature level as it was an unstable API at that point).  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdatePresence(context.Background()).Status(status).LastUpdateId(lastUpdateId).HistoryLimitDays(historyLimitDays).NewUserInput(newUserInput).PingOnly(pingOnly).SlimPresence(slimPresence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdatePresence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePresence`: UpdatePresence200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdatePresence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | The status of the user on this client.  Clients should report the user as &#x60;\\\&quot;active\\\&quot;&#x60; on this device if the client knows that the user is presently using the device (and thus would potentially see a notification immediately), even if the user has not directly interacted with the Zulip client.  Otherwise, it should report the user as &#x60;\\\&quot;idle\\\&quot;&#x60;.  See the related [&#x60;new_user_input&#x60;](#parameter-new_user_input) parameter for how a client should report whether the user is actively using the Zulip client.  | 
 **lastUpdateId** | **int32** | The identifier that specifies what presence data the client already has received, which allows the server to only return more recent user presence data.  This should be set to &#x60;-1&#x60; during initialization of the client in order to fetch all user presence data, unless the client is obtaining initial user presence metadata from the [&#x60;POST /register&#x60;](/api/register-queue) endpoint.  In subsequent queries to this endpoint, this value should be set to the most recent value of &#x60;presence_last_update_id&#x60; returned by the server in this endpoint&#39;s response, which implements incremental fetching of user presence data.  When this parameter is passed, the user presence data in the response will always be in the modern format.  **Changes**: New in Zulip 9.0 (feature level 263). Previously, the server sent user presence data for all users who had been active in the last two weeks unconditionally.  | 
 **historyLimitDays** | **int32** | Limits how far back in time to fetch user presence data. If not specified, defaults to 14 days. A value of N means that the oldest presence data fetched will be from at most N days ago.  Note that this is only useful during the initial user presence data fetch, as subsequent fetches should use the &#x60;last_update_id&#x60; parameter, which will act as the limit on how much presence data is returned. &#x60;history_limit_days&#x60; is ignored if &#x60;last_update_id&#x60; is passed with a value greater than &#x60;0&#x60;, indicating that the client already has some presence data.  **Changes**: New in Zulip 10.0 (feature level 288).  | 
 **newUserInput** | **bool** | Whether the user has interacted with the client (e.g. moved the mouse, used the keyboard, etc.) since the previous presence request from this client.  The server uses data from this parameter to implement certain [usage statistics](/help/analytics).  User interface clients that might run in the background, without the user ever interacting with them, should be careful to only pass &#x60;true&#x60; if the user has actually interacted with the client in order to avoid corrupting usage statistics graphs.  | [default to false]
 **pingOnly** | **bool** | Whether the client is sending a ping-only request, meaning it only wants to update the user&#39;s presence &#x60;status&#x60; on the server.  Otherwise, also requests the server return user presence data for all users in the organization, which is further specified by the [&#x60;last_update_id&#x60;](#parameter-last_update_id) parameter.  | [default to false]
 **slimPresence** | **bool** | Legacy parameter for configuring the format (modern or legacy) in which the server will return user presence data for the organization.  Modern clients should use [&#x60;last_update_id&#x60;](#parameter-last_update_id), which guarantees that user presence data will be returned in the modern format, and should not pass this parameter as &#x60;true&#x60; unless interacting with an older server.  Legacy clients that do not yet support &#x60;last_update_id&#x60; may use the value of &#x60;true&#x60; to request the modern format for user presence data.  **Note**: The legacy format for user presence data will be removed entirely in a future release.  **Changes**: **Deprecated** in Zulip 9.0 (feature level 263). Using the modern &#x60;last_update_id&#x60; parameter is the recommended way to request the modern format for user presence data.  New in Zulip 3.0 (no feature level as it was an unstable API at that point).  | [default to false]

### Return type

[**UpdatePresence200Response**](UpdatePresence200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettings

> IgnoredParametersSuccess UpdateSettings(ctx).FullName(fullName).Email(email).OldPassword(oldPassword).NewPassword(newPassword).TwentyFourHourTime(twentyFourHourTime).WebMarkReadOnScrollPolicy(webMarkReadOnScrollPolicy).WebChannelDefaultView(webChannelDefaultView).StarredMessageCounts(starredMessageCounts).ReceivesTypingNotifications(receivesTypingNotifications).WebSuggestUpdateTimezone(webSuggestUpdateTimezone).FluidLayoutWidth(fluidLayoutWidth).HighContrastMode(highContrastMode).WebFontSizePx(webFontSizePx).WebLineHeightPercent(webLineHeightPercent).ColorScheme(colorScheme).EnableDraftsSynchronization(enableDraftsSynchronization).TranslateEmoticons(translateEmoticons).DisplayEmojiReactionUsers(displayEmojiReactionUsers).DefaultLanguage(defaultLanguage).WebHomeView(webHomeView).WebEscapeNavigatesToHomeView(webEscapeNavigatesToHomeView).LeftSideUserlist(leftSideUserlist).Emojiset(emojiset).DemoteInactiveStreams(demoteInactiveStreams).UserListStyle(userListStyle).WebAnimateImagePreviews(webAnimateImagePreviews).WebStreamUnreadsCountDisplayPolicy(webStreamUnreadsCountDisplayPolicy).HideAiFeatures(hideAiFeatures).WebLeftSidebarShowChannelFolders(webLeftSidebarShowChannelFolders).WebLeftSidebarUnreadsCountSummary(webLeftSidebarUnreadsCountSummary).Timezone(timezone).EnableStreamDesktopNotifications(enableStreamDesktopNotifications).EnableStreamEmailNotifications(enableStreamEmailNotifications).EnableStreamPushNotifications(enableStreamPushNotifications).EnableStreamAudibleNotifications(enableStreamAudibleNotifications).NotificationSound(notificationSound).EnableDesktopNotifications(enableDesktopNotifications).EnableSounds(enableSounds).EmailNotificationsBatchingPeriodSeconds(emailNotificationsBatchingPeriodSeconds).EnableOfflineEmailNotifications(enableOfflineEmailNotifications).EnableOfflinePushNotifications(enableOfflinePushNotifications).EnableOnlinePushNotifications(enableOnlinePushNotifications).EnableFollowedTopicDesktopNotifications(enableFollowedTopicDesktopNotifications).EnableFollowedTopicEmailNotifications(enableFollowedTopicEmailNotifications).EnableFollowedTopicPushNotifications(enableFollowedTopicPushNotifications).EnableFollowedTopicAudibleNotifications(enableFollowedTopicAudibleNotifications).EnableDigestEmails(enableDigestEmails).EnableMarketingEmails(enableMarketingEmails).EnableLoginEmails(enableLoginEmails).MessageContentInEmailNotifications(messageContentInEmailNotifications).PmContentInDesktopNotifications(pmContentInDesktopNotifications).WildcardMentionsNotify(wildcardMentionsNotify).EnableFollowedTopicWildcardMentionsNotify(enableFollowedTopicWildcardMentionsNotify).DesktopIconCountDisplay(desktopIconCountDisplay).RealmNameInEmailNotificationsPolicy(realmNameInEmailNotificationsPolicy).AutomaticallyFollowTopicsPolicy(automaticallyFollowTopicsPolicy).AutomaticallyUnmuteTopicsInMutedStreamsPolicy(automaticallyUnmuteTopicsInMutedStreamsPolicy).AutomaticallyFollowTopicsWhereMentioned(automaticallyFollowTopicsWhereMentioned).ResolvedTopicNoticeAutoReadPolicy(resolvedTopicNoticeAutoReadPolicy).PresenceEnabled(presenceEnabled).EnterSends(enterSends).SendPrivateTypingNotifications(sendPrivateTypingNotifications).SendStreamTypingNotifications(sendStreamTypingNotifications).SendReadReceipts(sendReadReceipts).AllowPrivateDataExport(allowPrivateDataExport).EmailAddressVisibility(emailAddressVisibility).WebNavigateToSentMessage(webNavigateToSentMessage).Execute()

Update settings



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
	fullName := "fullName_example" // string | A new display name for the user.  (optional)
	email := "email_example" // string | Asks the server to initiate a confirmation sequence to change the user's email address to the indicated value. The user will need to demonstrate control of the new email address by clicking a confirmation link sent to that address.  (optional)
	oldPassword := "oldPassword_example" // string | The user's old Zulip password (or LDAP password, if LDAP authentication is in use).  Required only when sending the `new_password` parameter.  (optional)
	newPassword := "newPassword_example" // string | The user's new Zulip password (or LDAP password, if LDAP authentication is in use).  The `old_password` parameter must be included in the request.  (optional)
	twentyFourHourTime := true // bool | Whether time should be [displayed in 24-hour notation](/help/change-the-time-format).  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  (optional)
	webMarkReadOnScrollPolicy := int32(56) // int32 | Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \\\"Always\\\" setting when marking messages as read.  (optional)
	webChannelDefaultView := int32(56) // int32 | Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \\\"Top unread topic in channel\\\" is new in Zulip 11.0 (feature level 401).  The \\\"List of topics\\\" option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \\\"Channel feed\\\" behavior.  (optional)
	starredMessageCounts := true // bool | Whether clients should display the [number of starred messages](/help/star-a-message#display-the-number-of-starred-messages).  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  (optional)
	receivesTypingNotifications := true // bool | Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  By default, this is set to true, enabling user to receive typing notifications from other users.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.  (optional)
	webSuggestUpdateTimezone := true // bool | Whether the user should be shown an alert, offering to update their [profile time zone](/help/change-your-timezone), when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.  **Changes**: New in Zulip 10.0 (feature level 329).  (optional)
	fluidLayoutWidth := true // bool | Whether to use the [maximum available screen width](/help/enable-full-width-display) for the web app's center panel (message feed, recent conversations) on wide screens.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  (optional)
	highContrastMode := true // bool | This setting is reserved for use to control variations in Zulip's design to help visually impaired users.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  (optional)
	webFontSizePx := int32(56) // int32 | User-configured primary `font-size` for the web application, in pixels.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.  (optional)
	webLineHeightPercent := int32(56) // int32 | User-configured primary `line-height` for the web application, in percent, so a value of 120 represents a `line-height` of 1.2.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.  (optional)
	colorScheme := int32(56) // int32 | Controls which [color theme](/help/dark-theme) to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard `prefers-color-scheme` media query.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  (optional)
	enableDraftsSynchronization := true // bool | A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.  **Changes**: New in Zulip 5.0 (feature level 87).  (optional)
	translateEmoticons := true // bool | Whether to [translate emoticons to emoji](/help/configure-emoticon-translations) in messages the user sends.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  (optional)
	displayEmojiReactionUsers := true // bool | Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.  **Changes**: New in Zulip 6.0 (feature level 125).  (optional)
	defaultLanguage := "defaultLanguage_example" // string | What [default language](/help/change-your-language) to use for the account.  This controls both the Zulip UI as well as email notifications sent to the user.  The value needs to be a standard language code that the Zulip server has translation data for; for example, `\\\"en\\\"` for English or `\\\"de\\\"` for German.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 63).  (optional)
	webHomeView := "webHomeView_example" // string | The [home view](/help/configure-home-view) used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.  - \\\"recent_topics\\\" - Recent conversations view - \\\"inbox\\\" - Inbox view - \\\"all_messages\\\" - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).  Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).  (optional)
	webEscapeNavigatesToHomeView := true // bool | Whether the escape key navigates to the [configured home view](/help/configure-home-view).  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `escape_navigates_to_default_view`, which was new in Zulip 5.0 (feature level 107).  (optional)
	leftSideUserlist := true // bool | Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  (optional)
	emojiset := "emojiset_example" // string | The user's configured [emoji set](/help/emoji-and-emoticons#use-emoticons), used to display emoji to the user everywhere they appear in the UI.  - \\\"google\\\" - Google modern - \\\"google-blob\\\" - Google classic - \\\"twitter\\\" - Twitter - \\\"text\\\" - Plain text  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).  (optional)
	demoteInactiveStreams := int32(56) // int32 | Whether to [hide inactive channels](/help/manage-inactive-channels) in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  (optional)
	userListStyle := int32(56) // int32 | The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).  (optional)
	webAnimateImagePreviews := "webAnimateImagePreviews_example" // string | Controls how animated images should be played in the message feed in the web/desktop application.  - \\\"always\\\" - Always play the animated images in the message feed. - \\\"on_hover\\\" - Play the animated images on hover over them in the message feed. - \\\"never\\\" - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275).  (optional)
	webStreamUnreadsCountDisplayPolicy := int32(56) // int32 | Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).  (optional)
	hideAiFeatures := true // bool | Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).  (optional)
	webLeftSidebarShowChannelFolders := true // bool | Determines whether the web/desktop application's left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).  (optional)
	webLeftSidebarUnreadsCountSummary := true // bool | Determines whether the web/desktop application's left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).  (optional)
	timezone := "timezone_example" // string | The IANA identifier of the user's [profile time zone](/help/change-your-timezone), which is used primarily to display the user's local time to other users.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).  (optional)
	enableStreamDesktopNotifications := true // bool | Enable visual desktop notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableStreamEmailNotifications := true // bool | Enable email notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableStreamPushNotifications := true // bool | Enable mobile notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableStreamAudibleNotifications := true // bool | Enable audible desktop notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	notificationSound := "notificationSound_example" // string | Notification sound name.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 63).  (optional)
	enableDesktopNotifications := true // bool | Enable visual desktop notifications for direct messages and @-mentions.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableSounds := true // bool | Enable audible desktop notifications for direct messages and @-mentions.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	emailNotificationsBatchingPeriodSeconds := int32(56) // int32 | The duration (in seconds) for which the server should wait to batch email notifications before sending them.  **Changes**: New in Zulip 5.0 (feature level 82)  (optional)
	enableOfflineEmailNotifications := true // bool | Enable email notifications for direct messages and @-mentions received when the user is offline.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableOfflinePushNotifications := true // bool | Enable mobile notification for direct messages and @-mentions received when the user is offline.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableOnlinePushNotifications := true // bool | Enable mobile notification for direct messages and @-mentions received when the user is online.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableFollowedTopicDesktopNotifications := true // bool | Enable visual desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	enableFollowedTopicEmailNotifications := true // bool | Enable email notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	enableFollowedTopicPushNotifications := true // bool | Enable push notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	enableFollowedTopicAudibleNotifications := true // bool | Enable audible desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	enableDigestEmails := true // bool | Enable digest emails when the user is away.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableMarketingEmails := true // bool | Enable marketing emails. Has no function outside Zulip Cloud.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableLoginEmails := true // bool | Enable email notifications for new logins to account.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	messageContentInEmailNotifications := true // bool | Include the message's content in email notifications for new messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	pmContentInDesktopNotifications := true // bool | Include content of direct messages in desktop notifications.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	wildcardMentionsNotify := true // bool | Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enableFollowedTopicWildcardMentionsNotify := true // bool | Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	desktopIconCountDisplay := int32(56) // int32 | Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.  Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	realmNameInEmailNotificationsPolicy := int32(56) // int32 | Whether to [include organization name in subject of message notification emails](/help/email-notifications#include-organization-name-in-subject-line).  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.  Before Zulip 5.0 (feature level 80), the previous `realm_name_in_notifications` setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	automaticallyFollowTopicsPolicy := int32(56) // int32 | Which [topics to follow automatically](/help/mute-a-topic).  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  (optional)
	automaticallyUnmuteTopicsInMutedStreamsPolicy := int32(56) // int32 | Which [topics to unmute automatically in muted channels](/help/mute-a-topic).  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  (optional)
	automaticallyFollowTopicsWhereMentioned := true // bool | Whether the server will automatically mark the user as following topics where the user is mentioned.  **Changes**: New in Zulip 8.0 (feature level 235).  (optional)
	resolvedTopicNoticeAutoReadPolicy := "resolvedTopicNoticeAutoReadPolicy_example" // string | Controls whether the resolved-topic notices are marked as read.  - \\\"always\\\" - Always mark resolved-topic notices as read. - \\\"except_followed\\\" - Mark resolved-topic notices as read in topics not followed by the user. - \\\"never\\\" - Never mark resolved-topic notices as read.  **Changes**: New in Zulip 11.0 (feature level 385).  (optional)
	presenceEnabled := true // bool | Display the presence status to other users when online.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  (optional)
	enterSends := true // bool | Whether pressing Enter in the compose box sends a message (or saves a message edit).  **Changes**: Before Zulip 5.0 (feature level 81), this setting was managed by the `POST /users/me/enter-sends` endpoint, with the same parameter format.  (optional)
	sendPrivateTypingNotifications := true // bool | Whether [typing notifications](/help/typing-notifications) be sent when composing direct messages.  **Changes**: New in Zulip 5.0 (feature level 105).  (optional)
	sendStreamTypingNotifications := true // bool | Whether [typing notifications](/help/typing-notifications) be sent when composing channel messages.  **Changes**: New in Zulip 5.0 (feature level 105).  (optional)
	sendReadReceipts := true // bool | Whether other users are allowed to see whether you've read messages.  **Changes**: New in Zulip 5.0 (feature level 105).  (optional)
	allowPrivateDataExport := true // bool | Whether organization administrators are allowed to export your private data.  **Changes**: New in Zulip 10.0 (feature level 293).  (optional)
	emailAddressVisibility := int32(56) // int32 | The [policy][permission-level] this user has selected for [which other users][help-email-visibility] in this organization can see their real email address.  - 1 = Everyone - 2 = Members only - 3 = Administrators only - 4 = Nobody - 5 = Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.  [permission-level]: /api/roles-and-permissions#permission-levels [help-email-visibility]: /help/configure-email-visibility  (optional)
	webNavigateToSentMessage := true // bool | Web/desktop app setting for whether the user's view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateSettings(context.Background()).FullName(fullName).Email(email).OldPassword(oldPassword).NewPassword(newPassword).TwentyFourHourTime(twentyFourHourTime).WebMarkReadOnScrollPolicy(webMarkReadOnScrollPolicy).WebChannelDefaultView(webChannelDefaultView).StarredMessageCounts(starredMessageCounts).ReceivesTypingNotifications(receivesTypingNotifications).WebSuggestUpdateTimezone(webSuggestUpdateTimezone).FluidLayoutWidth(fluidLayoutWidth).HighContrastMode(highContrastMode).WebFontSizePx(webFontSizePx).WebLineHeightPercent(webLineHeightPercent).ColorScheme(colorScheme).EnableDraftsSynchronization(enableDraftsSynchronization).TranslateEmoticons(translateEmoticons).DisplayEmojiReactionUsers(displayEmojiReactionUsers).DefaultLanguage(defaultLanguage).WebHomeView(webHomeView).WebEscapeNavigatesToHomeView(webEscapeNavigatesToHomeView).LeftSideUserlist(leftSideUserlist).Emojiset(emojiset).DemoteInactiveStreams(demoteInactiveStreams).UserListStyle(userListStyle).WebAnimateImagePreviews(webAnimateImagePreviews).WebStreamUnreadsCountDisplayPolicy(webStreamUnreadsCountDisplayPolicy).HideAiFeatures(hideAiFeatures).WebLeftSidebarShowChannelFolders(webLeftSidebarShowChannelFolders).WebLeftSidebarUnreadsCountSummary(webLeftSidebarUnreadsCountSummary).Timezone(timezone).EnableStreamDesktopNotifications(enableStreamDesktopNotifications).EnableStreamEmailNotifications(enableStreamEmailNotifications).EnableStreamPushNotifications(enableStreamPushNotifications).EnableStreamAudibleNotifications(enableStreamAudibleNotifications).NotificationSound(notificationSound).EnableDesktopNotifications(enableDesktopNotifications).EnableSounds(enableSounds).EmailNotificationsBatchingPeriodSeconds(emailNotificationsBatchingPeriodSeconds).EnableOfflineEmailNotifications(enableOfflineEmailNotifications).EnableOfflinePushNotifications(enableOfflinePushNotifications).EnableOnlinePushNotifications(enableOnlinePushNotifications).EnableFollowedTopicDesktopNotifications(enableFollowedTopicDesktopNotifications).EnableFollowedTopicEmailNotifications(enableFollowedTopicEmailNotifications).EnableFollowedTopicPushNotifications(enableFollowedTopicPushNotifications).EnableFollowedTopicAudibleNotifications(enableFollowedTopicAudibleNotifications).EnableDigestEmails(enableDigestEmails).EnableMarketingEmails(enableMarketingEmails).EnableLoginEmails(enableLoginEmails).MessageContentInEmailNotifications(messageContentInEmailNotifications).PmContentInDesktopNotifications(pmContentInDesktopNotifications).WildcardMentionsNotify(wildcardMentionsNotify).EnableFollowedTopicWildcardMentionsNotify(enableFollowedTopicWildcardMentionsNotify).DesktopIconCountDisplay(desktopIconCountDisplay).RealmNameInEmailNotificationsPolicy(realmNameInEmailNotificationsPolicy).AutomaticallyFollowTopicsPolicy(automaticallyFollowTopicsPolicy).AutomaticallyUnmuteTopicsInMutedStreamsPolicy(automaticallyUnmuteTopicsInMutedStreamsPolicy).AutomaticallyFollowTopicsWhereMentioned(automaticallyFollowTopicsWhereMentioned).ResolvedTopicNoticeAutoReadPolicy(resolvedTopicNoticeAutoReadPolicy).PresenceEnabled(presenceEnabled).EnterSends(enterSends).SendPrivateTypingNotifications(sendPrivateTypingNotifications).SendStreamTypingNotifications(sendStreamTypingNotifications).SendReadReceipts(sendReadReceipts).AllowPrivateDataExport(allowPrivateDataExport).EmailAddressVisibility(emailAddressVisibility).WebNavigateToSentMessage(webNavigateToSentMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSettings`: IgnoredParametersSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullName** | **string** | A new display name for the user.  | 
 **email** | **string** | Asks the server to initiate a confirmation sequence to change the user&#39;s email address to the indicated value. The user will need to demonstrate control of the new email address by clicking a confirmation link sent to that address.  | 
 **oldPassword** | **string** | The user&#39;s old Zulip password (or LDAP password, if LDAP authentication is in use).  Required only when sending the &#x60;new_password&#x60; parameter.  | 
 **newPassword** | **string** | The user&#39;s new Zulip password (or LDAP password, if LDAP authentication is in use).  The &#x60;old_password&#x60; parameter must be included in the request.  | 
 **twentyFourHourTime** | **bool** | Whether time should be [displayed in 24-hour notation](/help/change-the-time-format).  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  | 
 **webMarkReadOnScrollPolicy** | **int32** | Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \\\&quot;Always\\\&quot; setting when marking messages as read.  | 
 **webChannelDefaultView** | **int32** | Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \\\&quot;Top unread topic in channel\\\&quot; is new in Zulip 11.0 (feature level 401).  The \\\&quot;List of topics\\\&quot; option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \\\&quot;Channel feed\\\&quot; behavior.  | 
 **starredMessageCounts** | **bool** | Whether clients should display the [number of starred messages](/help/star-a-message#display-the-number-of-starred-messages).  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  | 
 **receivesTypingNotifications** | **bool** | Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  By default, this is set to true, enabling user to receive typing notifications from other users.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.  | 
 **webSuggestUpdateTimezone** | **bool** | Whether the user should be shown an alert, offering to update their [profile time zone](/help/change-your-timezone), when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.  **Changes**: New in Zulip 10.0 (feature level 329).  | 
 **fluidLayoutWidth** | **bool** | Whether to use the [maximum available screen width](/help/enable-full-width-display) for the web app&#39;s center panel (message feed, recent conversations) on wide screens.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  | 
 **highContrastMode** | **bool** | This setting is reserved for use to control variations in Zulip&#39;s design to help visually impaired users.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  | 
 **webFontSizePx** | **int32** | User-configured primary &#x60;font-size&#x60; for the web application, in pixels.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.  | 
 **webLineHeightPercent** | **int32** | User-configured primary &#x60;line-height&#x60; for the web application, in percent, so a value of 120 represents a &#x60;line-height&#x60; of 1.2.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.  | 
 **colorScheme** | **int32** | Controls which [color theme](/help/dark-theme) to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard &#x60;prefers-color-scheme&#x60; media query.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  | 
 **enableDraftsSynchronization** | **bool** | A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.  **Changes**: New in Zulip 5.0 (feature level 87).  | 
 **translateEmoticons** | **bool** | Whether to [translate emoticons to emoji](/help/configure-emoticon-translations) in messages the user sends.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  | 
 **displayEmojiReactionUsers** | **bool** | Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.  **Changes**: New in Zulip 6.0 (feature level 125).  | 
 **defaultLanguage** | **string** | What [default language](/help/change-your-language) to use for the account.  This controls both the Zulip UI as well as email notifications sent to the user.  The value needs to be a standard language code that the Zulip server has translation data for; for example, &#x60;\\\&quot;en\\\&quot;&#x60; for English or &#x60;\\\&quot;de\\\&quot;&#x60; for German.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 63).  | 
 **webHomeView** | **string** | The [home view](/help/configure-home-view) used when opening a new Zulip web app window or hitting the &#x60;Esc&#x60; keyboard shortcut repeatedly.  - \\\&quot;recent_topics\\\&quot; - Recent conversations view - \\\&quot;inbox\\\&quot; - Inbox view - \\\&quot;all_messages\\\&quot; - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;default_view&#x60;, which was new in Zulip 4.0 (feature level 42).  Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).  | 
 **webEscapeNavigatesToHomeView** | **bool** | Whether the escape key navigates to the [configured home view](/help/configure-home-view).  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;escape_navigates_to_default_view&#x60;, which was new in Zulip 5.0 (feature level 107).  | 
 **leftSideUserlist** | **bool** | Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  | 
 **emojiset** | **string** | The user&#39;s configured [emoji set](/help/emoji-and-emoticons#use-emoticons), used to display emoji to the user everywhere they appear in the UI.  - \\\&quot;google\\\&quot; - Google modern - \\\&quot;google-blob\\\&quot; - Google classic - \\\&quot;twitter\\\&quot; - Twitter - \\\&quot;text\\\&quot; - Plain text  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).  | 
 **demoteInactiveStreams** | **int32** | Whether to [hide inactive channels](/help/manage-inactive-channels) in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  | 
 **userListStyle** | **int32** | The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).  | 
 **webAnimateImagePreviews** | **string** | Controls how animated images should be played in the message feed in the web/desktop application.  - \\\&quot;always\\\&quot; - Always play the animated images in the message feed. - \\\&quot;on_hover\\\&quot; - Play the animated images on hover over them in the message feed. - \\\&quot;never\\\&quot; - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275).  | 
 **webStreamUnreadsCountDisplayPolicy** | **int32** | Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).  | 
 **hideAiFeatures** | **bool** | Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).  | 
 **webLeftSidebarShowChannelFolders** | **bool** | Determines whether the web/desktop application&#39;s left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).  | 
 **webLeftSidebarUnreadsCountSummary** | **bool** | Determines whether the web/desktop application&#39;s left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).  | 
 **timezone** | **string** | The IANA identifier of the user&#39;s [profile time zone](/help/change-your-timezone), which is used primarily to display the user&#39;s local time to other users.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).  | 
 **enableStreamDesktopNotifications** | **bool** | Enable visual desktop notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableStreamEmailNotifications** | **bool** | Enable email notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableStreamPushNotifications** | **bool** | Enable mobile notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableStreamAudibleNotifications** | **bool** | Enable audible desktop notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **notificationSound** | **string** | Notification sound name.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 63).  | 
 **enableDesktopNotifications** | **bool** | Enable visual desktop notifications for direct messages and @-mentions.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableSounds** | **bool** | Enable audible desktop notifications for direct messages and @-mentions.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **emailNotificationsBatchingPeriodSeconds** | **int32** | The duration (in seconds) for which the server should wait to batch email notifications before sending them.  **Changes**: New in Zulip 5.0 (feature level 82)  | 
 **enableOfflineEmailNotifications** | **bool** | Enable email notifications for direct messages and @-mentions received when the user is offline.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableOfflinePushNotifications** | **bool** | Enable mobile notification for direct messages and @-mentions received when the user is offline.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableOnlinePushNotifications** | **bool** | Enable mobile notification for direct messages and @-mentions received when the user is online.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableFollowedTopicDesktopNotifications** | **bool** | Enable visual desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **enableFollowedTopicEmailNotifications** | **bool** | Enable email notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **enableFollowedTopicPushNotifications** | **bool** | Enable push notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **enableFollowedTopicAudibleNotifications** | **bool** | Enable audible desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **enableDigestEmails** | **bool** | Enable digest emails when the user is away.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableMarketingEmails** | **bool** | Enable marketing emails. Has no function outside Zulip Cloud.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableLoginEmails** | **bool** | Enable email notifications for new logins to account.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **messageContentInEmailNotifications** | **bool** | Include the message&#39;s content in email notifications for new messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **pmContentInDesktopNotifications** | **bool** | Include content of direct messages in desktop notifications.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **wildcardMentionsNotify** | **bool** | Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enableFollowedTopicWildcardMentionsNotify** | **bool** | Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **desktopIconCountDisplay** | **int32** | Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added &#x60;DMs, mentions, and followed topics&#x60; option, renumbering the options to insert it in order.  Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **realmNameInEmailNotificationsPolicy** | **int32** | Whether to [include organization name in subject of message notification emails](/help/email-notifications#include-organization-name-in-subject-line).  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous &#x60;realm_name_in_notifications&#x60; boolean; &#x60;true&#x60; corresponded to &#x60;Always&#x60;, and &#x60;false&#x60; to &#x60;Never&#x60;.  Before Zulip 5.0 (feature level 80), the previous &#x60;realm_name_in_notifications&#x60; setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **automaticallyFollowTopicsPolicy** | **int32** | Which [topics to follow automatically](/help/mute-a-topic).  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  | 
 **automaticallyUnmuteTopicsInMutedStreamsPolicy** | **int32** | Which [topics to unmute automatically in muted channels](/help/mute-a-topic).  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  | 
 **automaticallyFollowTopicsWhereMentioned** | **bool** | Whether the server will automatically mark the user as following topics where the user is mentioned.  **Changes**: New in Zulip 8.0 (feature level 235).  | 
 **resolvedTopicNoticeAutoReadPolicy** | **string** | Controls whether the resolved-topic notices are marked as read.  - \\\&quot;always\\\&quot; - Always mark resolved-topic notices as read. - \\\&quot;except_followed\\\&quot; - Mark resolved-topic notices as read in topics not followed by the user. - \\\&quot;never\\\&quot; - Never mark resolved-topic notices as read.  **Changes**: New in Zulip 11.0 (feature level 385).  | 
 **presenceEnabled** | **bool** | Display the presence status to other users when online.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  | 
 **enterSends** | **bool** | Whether pressing Enter in the compose box sends a message (or saves a message edit).  **Changes**: Before Zulip 5.0 (feature level 81), this setting was managed by the &#x60;POST /users/me/enter-sends&#x60; endpoint, with the same parameter format.  | 
 **sendPrivateTypingNotifications** | **bool** | Whether [typing notifications](/help/typing-notifications) be sent when composing direct messages.  **Changes**: New in Zulip 5.0 (feature level 105).  | 
 **sendStreamTypingNotifications** | **bool** | Whether [typing notifications](/help/typing-notifications) be sent when composing channel messages.  **Changes**: New in Zulip 5.0 (feature level 105).  | 
 **sendReadReceipts** | **bool** | Whether other users are allowed to see whether you&#39;ve read messages.  **Changes**: New in Zulip 5.0 (feature level 105).  | 
 **allowPrivateDataExport** | **bool** | Whether organization administrators are allowed to export your private data.  **Changes**: New in Zulip 10.0 (feature level 293).  | 
 **emailAddressVisibility** | **int32** | The [policy][permission-level] this user has selected for [which other users][help-email-visibility] in this organization can see their real email address.  - 1 &#x3D; Everyone - 2 &#x3D; Members only - 3 &#x3D; Administrators only - 4 &#x3D; Nobody - 5 &#x3D; Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.  [permission-level]: /api/roles-and-permissions#permission-levels [help-email-visibility]: /help/configure-email-visibility  | 
 **webNavigateToSentMessage** | **bool** | Web/desktop app setting for whether the user&#39;s view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.  | 

### Return type

[**IgnoredParametersSuccess**](IgnoredParametersSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStatus

> JsonSuccess UpdateStatus(ctx).StatusText(statusText).Away(away).EmojiName(emojiName).EmojiCode(emojiCode).ReactionType(reactionType).Execute()

Update your status



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
	statusText := "statusText_example" // string | The text content of the status message. Sending the empty string will clear the user's status.  **Note**: The limit on the size of the message is 60 characters.  (optional)
	away := true // bool | Whether the user should be marked as \\\"away\\\".  **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, `away` is a legacy way to access the user's `presence_enabled` setting, with `away = !presence_enabled`. To be removed in a future release.  (optional)
	emojiName := "emojiName_example" // string | The name for the emoji to associate with this status.  **Changes**: New in Zulip 5.0 (feature level 86).  (optional)
	emojiCode := "emojiCode_example" // string | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.  **Changes**: New in Zulip 5.0 (feature level 86).  (optional)
	reactionType := "reactionType_example" // string | A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:  - `unicode_emoji` : In this namespace, `emoji_code` will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - `realm_emoji` : In this namespace, `emoji_code` will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - `zulip_extra_emoji` : These are special emoji included with Zulip.   In this namespace, `emoji_code` will be the name of the emoji (e.g.   \\\"zulip\\\").  **Changes**: New in Zulip 5.0 (feature level 86).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateStatus(context.Background()).StatusText(statusText).Away(away).EmojiName(emojiName).EmojiCode(emojiCode).ReactionType(reactionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStatus`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusText** | **string** | The text content of the status message. Sending the empty string will clear the user&#39;s status.  **Note**: The limit on the size of the message is 60 characters.  | 
 **away** | **bool** | Whether the user should be marked as \\\&quot;away\\\&quot;.  **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, &#x60;away&#x60; is a legacy way to access the user&#39;s &#x60;presence_enabled&#x60; setting, with &#x60;away &#x3D; !presence_enabled&#x60;. To be removed in a future release.  | 
 **emojiName** | **string** | The name for the emoji to associate with this status.  **Changes**: New in Zulip 5.0 (feature level 86).  | 
 **emojiCode** | **string** | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the &#x60;reaction_type&#x60;.  **Changes**: New in Zulip 5.0 (feature level 86).  | 
 **reactionType** | **string** | A string indicating the type of emoji. Each emoji &#x60;reaction_type&#x60; has an independent namespace for values of &#x60;emoji_code&#x60;.  Must be one of the following values:  - &#x60;unicode_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - &#x60;realm_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - &#x60;zulip_extra_emoji&#x60; : These are special emoji included with Zulip.   In this namespace, &#x60;emoji_code&#x60; will be the name of the emoji (e.g.   \\\&quot;zulip\\\&quot;).  **Changes**: New in Zulip 5.0 (feature level 86).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStatusForUser

> JsonSuccess UpdateStatusForUser(ctx, userId).StatusText(statusText).EmojiName(emojiName).EmojiCode(emojiCode).ReactionType(reactionType).Execute()

Update user status



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
	userId := int32(12) // int32 | The target user's ID. 
	statusText := "statusText_example" // string | The text content of the status message. Sending the empty string will clear the user's status.  **Note**: The limit on the size of the message is 60 characters.  (optional)
	emojiName := "emojiName_example" // string | The name for the emoji to associate with this status.  **Changes**: New in Zulip 5.0 (feature level 86).  (optional)
	emojiCode := "emojiCode_example" // string | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.  **Changes**: New in Zulip 5.0 (feature level 86).  (optional)
	reactionType := "reactionType_example" // string | A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:  - `unicode_emoji` : In this namespace, `emoji_code` will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - `realm_emoji` : In this namespace, `emoji_code` will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - `zulip_extra_emoji` : These are special emoji included with Zulip.   In this namespace, `emoji_code` will be the name of the emoji (e.g.   \\\"zulip\\\").  **Changes**: New in Zulip 5.0 (feature level 86).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateStatusForUser(context.Background(), userId).StatusText(statusText).EmojiName(emojiName).EmojiCode(emojiCode).ReactionType(reactionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateStatusForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStatusForUser`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateStatusForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The target user&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatusForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statusText** | **string** | The text content of the status message. Sending the empty string will clear the user&#39;s status.  **Note**: The limit on the size of the message is 60 characters.  | 
 **emojiName** | **string** | The name for the emoji to associate with this status.  **Changes**: New in Zulip 5.0 (feature level 86).  | 
 **emojiCode** | **string** | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the &#x60;reaction_type&#x60;.  **Changes**: New in Zulip 5.0 (feature level 86).  | 
 **reactionType** | **string** | A string indicating the type of emoji. Each emoji &#x60;reaction_type&#x60; has an independent namespace for values of &#x60;emoji_code&#x60;.  Must be one of the following values:  - &#x60;unicode_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - &#x60;realm_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - &#x60;zulip_extra_emoji&#x60; : These are special emoji included with Zulip.   In this namespace, &#x60;emoji_code&#x60; will be the name of the emoji (e.g.   \\\&quot;zulip\\\&quot;).  **Changes**: New in Zulip 5.0 (feature level 86).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> JsonSuccess UpdateUser(ctx, userId).FullName(fullName).Role(role).ProfileData(profileData).NewEmail(newEmail).Execute()

Update a user



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
	userId := int32(12) // int32 | The target user's ID. 
	fullName := "fullName_example" // string | The user's full name.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 5.0 (feature level 106).  (optional)
	role := int32(56) // int32 | New [role](/api/roles-and-permissions) for the user. Roles are encoded as:  - Organization owner: 100 - Organization administrator: 200 - Organization moderator: 300 - Member: 400 - Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.  **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of `is_admin` and `is_guest` boolean parameters. Organization moderator role added in Zulip 4.0 (feature level 60).  (optional)
	profileData := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | A dictionary containing the updated custom profile field data for the user.  (optional)
	newEmail := "newEmail_example" // string | New email address for the user. Requires the user making the request to be an organization owner and additionally have the `.can_change_user_emails` special permission.  **Changes**: New in Zulip 10.0 (feature level 285).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateUser(context.Background(), userId).FullName(fullName).Role(role).ProfileData(profileData).NewEmail(newEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The target user&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullName** | **string** | The user&#39;s full name.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 5.0 (feature level 106).  | 
 **role** | **int32** | New [role](/api/roles-and-permissions) for the user. Roles are encoded as:  - Organization owner: 100 - Organization administrator: 200 - Organization moderator: 300 - Member: 400 - Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.  **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of &#x60;is_admin&#x60; and &#x60;is_guest&#x60; boolean parameters. Organization moderator role added in Zulip 4.0 (feature level 60).  | 
 **profileData** | **[]map[string]interface{}** | A dictionary containing the updated custom profile field data for the user.  | 
 **newEmail** | **string** | New email address for the user. Requires the user making the request to be an organization owner and additionally have the &#x60;.can_change_user_emails&#x60; special permission.  **Changes**: New in Zulip 10.0 (feature level 285).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserByEmail

> JsonSuccess UpdateUserByEmail(ctx, email).FullName(fullName).Role(role).ProfileData(profileData).NewEmail(newEmail).Execute()

Update a user by email



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
	email := "hamlet@zulip.com" // string | The email address of the user, specified following the same rules as [`GET /users/{email}`](/api/get-user-by-email). 
	fullName := "fullName_example" // string | The user's full name.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 5.0 (feature level 106).  (optional)
	role := int32(56) // int32 | New [role](/api/roles-and-permissions) for the user. Roles are encoded as:  - Organization owner: 100 - Organization administrator: 200 - Organization moderator: 300 - Member: 400 - Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.  **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of `is_admin` and `is_guest` boolean parameters. Organization moderator role added in Zulip 4.0 (feature level 60).  (optional)
	profileData := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | A dictionary containing the updated custom profile field data for the user.  (optional)
	newEmail := "newEmail_example" // string | New email address for the user. Requires the user making the request to be an organization owner and additionally have the `.can_change_user_emails` special permission.  **Changes**: New in Zulip 10.0 (feature level 285).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateUserByEmail(context.Background(), email).FullName(fullName).Role(role).ProfileData(profileData).NewEmail(newEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserByEmail`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | The email address of the user, specified following the same rules as [&#x60;GET /users/{email}&#x60;](/api/get-user-by-email).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullName** | **string** | The user&#39;s full name.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 5.0 (feature level 106).  | 
 **role** | **int32** | New [role](/api/roles-and-permissions) for the user. Roles are encoded as:  - Organization owner: 100 - Organization administrator: 200 - Organization moderator: 300 - Member: 400 - Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.  **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of &#x60;is_admin&#x60; and &#x60;is_guest&#x60; boolean parameters. Organization moderator role added in Zulip 4.0 (feature level 60).  | 
 **profileData** | **[]map[string]interface{}** | A dictionary containing the updated custom profile field data for the user.  | 
 **newEmail** | **string** | New email address for the user. Requires the user making the request to be an organization owner and additionally have the &#x60;.can_change_user_emails&#x60; special permission.  **Changes**: New in Zulip 10.0 (feature level 285).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserGroup

> JsonSuccess UpdateUserGroup(ctx, userGroupId).Name(name).Description(description).CanAddMembersGroup(canAddMembersGroup).CanJoinGroup(canJoinGroup).CanLeaveGroup(canLeaveGroup).CanManageGroup(canManageGroup).CanMentionGroup(canMentionGroup).CanRemoveMembersGroup(canRemoveMembersGroup).Deactivated(deactivated).Execute()

Update a user group



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
	userGroupId := int32(38) // int32 | The ID of the target user group. 
	name := "name_example" // string | The new name of the group.  **Changes**: Before Zulip 7.0 (feature level 165), this was a required field.  (optional)
	description := "description_example" // string | The new description of the group.  **Changes**: Before Zulip 7.0 (feature level 165), this was a required field.  (optional)
	canAddMembersGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to add members to this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the `can_manage_group` setting.  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups  (optional)
	canJoinGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to join this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 301).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups  (optional)
	canLeaveGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to leave this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 308).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups  (optional)
	canManageGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to [manage this user group][manage-user-groups] expressed as an [update to a group-setting value][update-group-setting].  This setting cannot be set to `\\\"role:internet\\\"` and `\\\"role:everyone\\\"` [system groups][system-groups].  **Changes**: New in Zulip 10.0 (feature level 283).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups [manage-user-groups]: /help/manage-user-groups  (optional)
	canMentionGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to [mention this group][mentions], expressed as an [update to a group-setting value][update-group-setting].  This setting cannot be set to `\\\"role:internet\\\"` and `\\\"role:owners\\\"` [system groups][system-groups].  **Changes**: In Zulip 9.0 (feature level 260), this parameter was updated to only accept an object with the `old` and `new` fields described below. Prior to this feature level, this parameter could be either of the two forms of a [group-setting value][setting-values].  Before Zulip 9.0 (feature level 258), this parameter could only be the integer form of a [group-setting value][setting-values].  Before Zulip 8.0 (feature level 198), this parameter was named `can_mention_group_id`.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups][system-groups].  [mentions]: /help/mention-a-user-or-group [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups [setting-values]: /api/group-setting-values  (optional)
	canRemoveMembersGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to remove members from this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the `can_manage_group` setting.  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups  (optional)
	deactivated := true // bool | A deactivated user group can be reactivated by passing this parameter as `false`.  Passing `true` does nothing as user group is deactivated using [`POST /user_groups/{user_group_id}/deactivate`](deactivate-user-group) endpoint.  **Changes**: New in Zulip 11.0 (feature level 386).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateUserGroup(context.Background(), userGroupId).Name(name).Description(description).CanAddMembersGroup(canAddMembersGroup).CanJoinGroup(canJoinGroup).CanLeaveGroup(canLeaveGroup).CanManageGroup(canManageGroup).CanMentionGroup(canMentionGroup).CanRemoveMembersGroup(canRemoveMembersGroup).Deactivated(deactivated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserGroup`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** | The ID of the target user group.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The new name of the group.  **Changes**: Before Zulip 7.0 (feature level 165), this was a required field.  | 
 **description** | **string** | The new description of the group.  **Changes**: Before Zulip 7.0 (feature level 165), this was a required field.  | 
 **canAddMembersGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to add members to this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups  | 
 **canJoinGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to join this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 301).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups  | 
 **canLeaveGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to leave this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 308).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups  | 
 **canManageGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to [manage this user group][manage-user-groups] expressed as an [update to a group-setting value][update-group-setting].  This setting cannot be set to &#x60;\\\&quot;role:internet\\\&quot;&#x60; and &#x60;\\\&quot;role:everyone\\\&quot;&#x60; [system groups][system-groups].  **Changes**: New in Zulip 10.0 (feature level 283).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups [manage-user-groups]: /help/manage-user-groups  | 
 **canMentionGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to [mention this group][mentions], expressed as an [update to a group-setting value][update-group-setting].  This setting cannot be set to &#x60;\\\&quot;role:internet\\\&quot;&#x60; and &#x60;\\\&quot;role:owners\\\&quot;&#x60; [system groups][system-groups].  **Changes**: In Zulip 9.0 (feature level 260), this parameter was updated to only accept an object with the &#x60;old&#x60; and &#x60;new&#x60; fields described below. Prior to this feature level, this parameter could be either of the two forms of a [group-setting value][setting-values].  Before Zulip 9.0 (feature level 258), this parameter could only be the integer form of a [group-setting value][setting-values].  Before Zulip 8.0 (feature level 198), this parameter was named &#x60;can_mention_group_id&#x60;.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups][system-groups].  [mentions]: /help/mention-a-user-or-group [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups [setting-values]: /api/group-setting-values  | 
 **canRemoveMembersGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to remove members from this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups  | 
 **deactivated** | **bool** | A deactivated user group can be reactivated by passing this parameter as &#x60;false&#x60;.  Passing &#x60;true&#x60; does nothing as user group is deactivated using [&#x60;POST /user_groups/{user_group_id}/deactivate&#x60;](deactivate-user-group) endpoint.  **Changes**: New in Zulip 11.0 (feature level 386).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserGroupMembers

> JsonSuccess UpdateUserGroupMembers(ctx, userGroupId).Delete(delete).Add(add).DeleteSubgroups(deleteSubgroups).AddSubgroups(addSubgroups).Execute()

Update user group members



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
	userGroupId := int32(38) // int32 | The ID of the target user group. 
	delete := []int32{int32(123)} // []int32 | The list of user IDs to be removed from the user group.  (optional)
	add := []int32{int32(123)} // []int32 | The list of user IDs to be added to the user group.  (optional)
	deleteSubgroups := []int32{int32(123)} // []int32 | The list of user group IDs to be removed from the user group.  **Changes**: New in Zulip 10.0 (feature level 311).  (optional)
	addSubgroups := []int32{int32(123)} // []int32 | The list of user group IDs to be added to the user group.  **Changes**: New in Zulip 10.0 (feature level 311).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateUserGroupMembers(context.Background(), userGroupId).Delete(delete).Add(add).DeleteSubgroups(deleteSubgroups).AddSubgroups(addSubgroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserGroupMembers`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** | The ID of the target user group.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delete** | **[]int32** | The list of user IDs to be removed from the user group.  | 
 **add** | **[]int32** | The list of user IDs to be added to the user group.  | 
 **deleteSubgroups** | **[]int32** | The list of user group IDs to be removed from the user group.  **Changes**: New in Zulip 10.0 (feature level 311).  | 
 **addSubgroups** | **[]int32** | The list of user group IDs to be added to the user group.  **Changes**: New in Zulip 10.0 (feature level 311).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserGroupSubgroups

> JsonSuccess UpdateUserGroupSubgroups(ctx, userGroupId).Delete(delete).Add(add).Execute()

Update subgroups of a user group



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
	userGroupId := int32(38) // int32 | The ID of the target user group. 
	delete := []int32{int32(123)} // []int32 | The list of user group IDs to be removed from the user group.  (optional)
	add := []int32{int32(123)} // []int32 | The list of user group IDs to be added to the user group.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateUserGroupSubgroups(context.Background(), userGroupId).Delete(delete).Add(add).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserGroupSubgroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserGroupSubgroups`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserGroupSubgroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** | The ID of the target user group.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupSubgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delete** | **[]int32** | The list of user group IDs to be removed from the user group.  | 
 **add** | **[]int32** | The list of user group IDs to be added to the user group.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

