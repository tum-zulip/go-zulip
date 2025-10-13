# UserUpdateEventEnvalop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user affected by this change.  | [optional] 
**FullName** | Pointer to **string** | The new full name for the user.  | [optional] 
**AvatarUrl** | Pointer to **string** | The URL of the new avatar for the user.  | [optional] 
**AvatarSource** | Pointer to **string** | The new avatar data source type for the user.  Value values are &#x60;G&#x60; (gravatar) and &#x60;U&#x60; (uploaded by user).  | [optional] 
**AvatarUrlMedium** | Pointer to **string** | The new medium-size avatar URL for user.  | [optional] 
**AvatarVersion** | Pointer to **int32** | The version number for the user&#39;s avatar. This is useful for cache-busting.  | [optional] 
**Email** | Pointer to **string** | The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the &#x60;user_id&#x60;.  | [optional] 
**Timezone** | Pointer to **string** | The IANA identifier of the new profile time zone for the user.  | [optional] 
**BotOwnerId** | Pointer to **int32** | The user ID of the new bot owner.  | [optional] 
**Role** | Pointer to **int32** | The new [role](/api/roles-and-permissions) of the user.  | [optional] 
**DeliveryEmail** | Pointer to **NullableString** | The new delivery email of the user.  This value can be &#x60;null&#x60; if the affected user changed their &#x60;email_address_visibility&#x60; setting such that you cannot access their real email.  **Changes**: Before Zulip 7.0 (feature level 163), &#x60;null&#x60; was not a possible value for this event as it was only sent to the affected user when their email address was changed.  | [optional] 
**CustomProfileField** | Pointer to [**UserUpdateEventCustomFieldDetails**](UserUpdateEventCustomFieldDetails.md) |  | [optional] 
**NewEmail** | Pointer to **string** | The new value of &#x60;email&#x60; for the user. The client should update any data structures associated with this user to use this new value as the user&#39;s Zulip API email address.  | [optional] 
**IsActive** | Pointer to **bool** | A boolean describing whether the user account has been deactivated.  | [optional] 

## Methods

### NewUserUpdateEventEnvalop

`func NewUserUpdateEventEnvalop() *UserUpdateEventEnvalop`

NewUserUpdateEventEnvalop instantiates a new UserUpdateEventEnvalop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateEventEnvalopWithDefaults

`func NewUserUpdateEventEnvalopWithDefaults() *UserUpdateEventEnvalop`

NewUserUpdateEventEnvalopWithDefaults instantiates a new UserUpdateEventEnvalop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserUpdateEventEnvalop) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUpdateEventEnvalop) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUpdateEventEnvalop) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserUpdateEventEnvalop) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetFullName

`func (o *UserUpdateEventEnvalop) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserUpdateEventEnvalop) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserUpdateEventEnvalop) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserUpdateEventEnvalop) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserUpdateEventEnvalop) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserUpdateEventEnvalop) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserUpdateEventEnvalop) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserUpdateEventEnvalop) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetAvatarSource

`func (o *UserUpdateEventEnvalop) GetAvatarSource() string`

GetAvatarSource returns the AvatarSource field if non-nil, zero value otherwise.

### GetAvatarSourceOk

`func (o *UserUpdateEventEnvalop) GetAvatarSourceOk() (*string, bool)`

GetAvatarSourceOk returns a tuple with the AvatarSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSource

`func (o *UserUpdateEventEnvalop) SetAvatarSource(v string)`

SetAvatarSource sets AvatarSource field to given value.

### HasAvatarSource

`func (o *UserUpdateEventEnvalop) HasAvatarSource() bool`

HasAvatarSource returns a boolean if a field has been set.

### GetAvatarUrlMedium

`func (o *UserUpdateEventEnvalop) GetAvatarUrlMedium() string`

GetAvatarUrlMedium returns the AvatarUrlMedium field if non-nil, zero value otherwise.

### GetAvatarUrlMediumOk

`func (o *UserUpdateEventEnvalop) GetAvatarUrlMediumOk() (*string, bool)`

GetAvatarUrlMediumOk returns a tuple with the AvatarUrlMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrlMedium

`func (o *UserUpdateEventEnvalop) SetAvatarUrlMedium(v string)`

SetAvatarUrlMedium sets AvatarUrlMedium field to given value.

### HasAvatarUrlMedium

`func (o *UserUpdateEventEnvalop) HasAvatarUrlMedium() bool`

HasAvatarUrlMedium returns a boolean if a field has been set.

### GetAvatarVersion

`func (o *UserUpdateEventEnvalop) GetAvatarVersion() int32`

GetAvatarVersion returns the AvatarVersion field if non-nil, zero value otherwise.

### GetAvatarVersionOk

`func (o *UserUpdateEventEnvalop) GetAvatarVersionOk() (*int32, bool)`

GetAvatarVersionOk returns a tuple with the AvatarVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarVersion

`func (o *UserUpdateEventEnvalop) SetAvatarVersion(v int32)`

SetAvatarVersion sets AvatarVersion field to given value.

### HasAvatarVersion

`func (o *UserUpdateEventEnvalop) HasAvatarVersion() bool`

HasAvatarVersion returns a boolean if a field has been set.

### GetEmail

`func (o *UserUpdateEventEnvalop) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUpdateEventEnvalop) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUpdateEventEnvalop) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserUpdateEventEnvalop) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTimezone

`func (o *UserUpdateEventEnvalop) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserUpdateEventEnvalop) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserUpdateEventEnvalop) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserUpdateEventEnvalop) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetBotOwnerId

`func (o *UserUpdateEventEnvalop) GetBotOwnerId() int32`

GetBotOwnerId returns the BotOwnerId field if non-nil, zero value otherwise.

### GetBotOwnerIdOk

`func (o *UserUpdateEventEnvalop) GetBotOwnerIdOk() (*int32, bool)`

GetBotOwnerIdOk returns a tuple with the BotOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotOwnerId

`func (o *UserUpdateEventEnvalop) SetBotOwnerId(v int32)`

SetBotOwnerId sets BotOwnerId field to given value.

### HasBotOwnerId

`func (o *UserUpdateEventEnvalop) HasBotOwnerId() bool`

HasBotOwnerId returns a boolean if a field has been set.

### GetRole

`func (o *UserUpdateEventEnvalop) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserUpdateEventEnvalop) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserUpdateEventEnvalop) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserUpdateEventEnvalop) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetDeliveryEmail

`func (o *UserUpdateEventEnvalop) GetDeliveryEmail() string`

GetDeliveryEmail returns the DeliveryEmail field if non-nil, zero value otherwise.

### GetDeliveryEmailOk

`func (o *UserUpdateEventEnvalop) GetDeliveryEmailOk() (*string, bool)`

GetDeliveryEmailOk returns a tuple with the DeliveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEmail

`func (o *UserUpdateEventEnvalop) SetDeliveryEmail(v string)`

SetDeliveryEmail sets DeliveryEmail field to given value.

### HasDeliveryEmail

`func (o *UserUpdateEventEnvalop) HasDeliveryEmail() bool`

HasDeliveryEmail returns a boolean if a field has been set.

### SetDeliveryEmailNil

`func (o *UserUpdateEventEnvalop) SetDeliveryEmailNil(b bool)`

 SetDeliveryEmailNil sets the value for DeliveryEmail to be an explicit nil

### UnsetDeliveryEmail
`func (o *UserUpdateEventEnvalop) UnsetDeliveryEmail()`

UnsetDeliveryEmail ensures that no value is present for DeliveryEmail, not even an explicit nil
### GetCustomProfileField

`func (o *UserUpdateEventEnvalop) GetCustomProfileField() UserUpdateEventCustomFieldDetails`

GetCustomProfileField returns the CustomProfileField field if non-nil, zero value otherwise.

### GetCustomProfileFieldOk

`func (o *UserUpdateEventEnvalop) GetCustomProfileFieldOk() (*UserUpdateEventCustomFieldDetails, bool)`

GetCustomProfileFieldOk returns a tuple with the CustomProfileField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProfileField

`func (o *UserUpdateEventEnvalop) SetCustomProfileField(v UserUpdateEventCustomFieldDetails)`

SetCustomProfileField sets CustomProfileField field to given value.

### HasCustomProfileField

`func (o *UserUpdateEventEnvalop) HasCustomProfileField() bool`

HasCustomProfileField returns a boolean if a field has been set.

### GetNewEmail

`func (o *UserUpdateEventEnvalop) GetNewEmail() string`

GetNewEmail returns the NewEmail field if non-nil, zero value otherwise.

### GetNewEmailOk

`func (o *UserUpdateEventEnvalop) GetNewEmailOk() (*string, bool)`

GetNewEmailOk returns a tuple with the NewEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEmail

`func (o *UserUpdateEventEnvalop) SetNewEmail(v string)`

SetNewEmail sets NewEmail field to given value.

### HasNewEmail

`func (o *UserUpdateEventEnvalop) HasNewEmail() bool`

HasNewEmail returns a boolean if a field has been set.

### GetIsActive

`func (o *UserUpdateEventEnvalop) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserUpdateEventEnvalop) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserUpdateEventEnvalop) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserUpdateEventEnvalop) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


