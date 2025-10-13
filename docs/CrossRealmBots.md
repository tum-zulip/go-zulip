# CrossRealmBots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **interface{}** |  | [optional] 
**DeliveryEmail** | Pointer to **interface{}** |  | [optional] 
**Email** | Pointer to **interface{}** |  | [optional] 
**FullName** | Pointer to **interface{}** |  | [optional] 
**DateJoined** | Pointer to **interface{}** |  | [optional] 
**IsActive** | Pointer to **interface{}** |  | [optional] 
**IsOwner** | Pointer to **interface{}** |  | [optional] 
**IsAdmin** | Pointer to **interface{}** |  | [optional] 
**IsGuest** | Pointer to **interface{}** |  | [optional] 
**IsBot** | Pointer to **interface{}** |  | [optional] 
**BotType** | Pointer to **interface{}** |  | [optional] 
**BotOwnerId** | Pointer to **interface{}** |  | [optional] 
**Role** | Pointer to **interface{}** |  | [optional] 
**Timezone** | Pointer to **interface{}** |  | [optional] 
**AvatarUrl** | Pointer to **interface{}** |  | [optional] 
**AvatarVersion** | Pointer to **interface{}** |  | [optional] 
**ProfileData** | Pointer to **interface{}** |  | [optional] 
**IsSystemBot** | Pointer to **bool** | Whether the user is a system bot. System bots are special bot user accounts that are managed by the system, rather than the organization&#39;s administrators.  **Changes**: This field was called &#x60;is_cross_realm_bot&#x60; before Zulip 5.0 (feature level 83).  | [optional] 

## Methods

### NewRegisterQueueResponseCrossRealmBotsItem

`func NewRegisterQueueResponseCrossRealmBotsItem() *RegisterQueueResponseCrossRealmBotsItem`

NewRegisterQueueResponseCrossRealmBotsItem instantiates a new RegisterQueueResponseCrossRealmBotsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseCrossRealmBotsItemWithDefaults

`func NewRegisterQueueResponseCrossRealmBotsItemWithDefaults() *RegisterQueueResponseCrossRealmBotsItem`

NewRegisterQueueResponseCrossRealmBotsItemWithDefaults instantiates a new RegisterQueueResponseCrossRealmBotsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetUserId() interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetUserIdOk() (*interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetUserId(v interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDeliveryEmail

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetDeliveryEmail() interface{}`

GetDeliveryEmail returns the DeliveryEmail field if non-nil, zero value otherwise.

### GetDeliveryEmailOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetDeliveryEmailOk() (*interface{}, bool)`

GetDeliveryEmailOk returns a tuple with the DeliveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEmail

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetDeliveryEmail(v interface{})`

SetDeliveryEmail sets DeliveryEmail field to given value.

### HasDeliveryEmail

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasDeliveryEmail() bool`

HasDeliveryEmail returns a boolean if a field has been set.

### SetDeliveryEmailNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetDeliveryEmailNil(b bool)`

 SetDeliveryEmailNil sets the value for DeliveryEmail to be an explicit nil

### UnsetDeliveryEmail
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetDeliveryEmail()`

UnsetDeliveryEmail ensures that no value is present for DeliveryEmail, not even an explicit nil
### GetEmail

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullName

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetFullName() interface{}`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetFullNameOk() (*interface{}, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetFullName(v interface{})`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDateJoined

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetDateJoined() interface{}`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetDateJoinedOk() (*interface{}, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetDateJoined(v interface{})`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### SetDateJoinedNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetDateJoinedNil(b bool)`

 SetDateJoinedNil sets the value for DateJoined to be an explicit nil

### UnsetDateJoined
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetDateJoined()`

UnsetDateJoined ensures that no value is present for DateJoined, not even an explicit nil
### GetIsActive

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsActive() interface{}`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsActiveOk() (*interface{}, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsActive(v interface{})`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsOwner

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsOwner() interface{}`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsOwnerOk() (*interface{}, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsOwner(v interface{})`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### SetIsOwnerNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsOwnerNil(b bool)`

 SetIsOwnerNil sets the value for IsOwner to be an explicit nil

### UnsetIsOwner
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetIsOwner()`

UnsetIsOwner ensures that no value is present for IsOwner, not even an explicit nil
### GetIsAdmin

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsAdmin() interface{}`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsAdminOk() (*interface{}, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsAdmin(v interface{})`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### SetIsAdminNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsAdminNil(b bool)`

 SetIsAdminNil sets the value for IsAdmin to be an explicit nil

### UnsetIsAdmin
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetIsAdmin()`

UnsetIsAdmin ensures that no value is present for IsAdmin, not even an explicit nil
### GetIsGuest

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsGuest() interface{}`

GetIsGuest returns the IsGuest field if non-nil, zero value otherwise.

### GetIsGuestOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsGuestOk() (*interface{}, bool)`

GetIsGuestOk returns a tuple with the IsGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuest

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsGuest(v interface{})`

SetIsGuest sets IsGuest field to given value.

### HasIsGuest

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasIsGuest() bool`

HasIsGuest returns a boolean if a field has been set.

### SetIsGuestNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsGuestNil(b bool)`

 SetIsGuestNil sets the value for IsGuest to be an explicit nil

### UnsetIsGuest
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetIsGuest()`

UnsetIsGuest ensures that no value is present for IsGuest, not even an explicit nil
### GetIsBot

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsBot() interface{}`

GetIsBot returns the IsBot field if non-nil, zero value otherwise.

### GetIsBotOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsBotOk() (*interface{}, bool)`

GetIsBotOk returns a tuple with the IsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBot

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsBot(v interface{})`

SetIsBot sets IsBot field to given value.

### HasIsBot

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasIsBot() bool`

HasIsBot returns a boolean if a field has been set.

### SetIsBotNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsBotNil(b bool)`

 SetIsBotNil sets the value for IsBot to be an explicit nil

### UnsetIsBot
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetIsBot()`

UnsetIsBot ensures that no value is present for IsBot, not even an explicit nil
### GetBotType

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetBotType() interface{}`

GetBotType returns the BotType field if non-nil, zero value otherwise.

### GetBotTypeOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetBotTypeOk() (*interface{}, bool)`

GetBotTypeOk returns a tuple with the BotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotType

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetBotType(v interface{})`

SetBotType sets BotType field to given value.

### HasBotType

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasBotType() bool`

HasBotType returns a boolean if a field has been set.

### SetBotTypeNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetBotTypeNil(b bool)`

 SetBotTypeNil sets the value for BotType to be an explicit nil

### UnsetBotType
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetBotType()`

UnsetBotType ensures that no value is present for BotType, not even an explicit nil
### GetBotOwnerId

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetBotOwnerId() interface{}`

GetBotOwnerId returns the BotOwnerId field if non-nil, zero value otherwise.

### GetBotOwnerIdOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetBotOwnerIdOk() (*interface{}, bool)`

GetBotOwnerIdOk returns a tuple with the BotOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotOwnerId

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetBotOwnerId(v interface{})`

SetBotOwnerId sets BotOwnerId field to given value.

### HasBotOwnerId

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasBotOwnerId() bool`

HasBotOwnerId returns a boolean if a field has been set.

### SetBotOwnerIdNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetBotOwnerIdNil(b bool)`

 SetBotOwnerIdNil sets the value for BotOwnerId to be an explicit nil

### UnsetBotOwnerId
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetBotOwnerId()`

UnsetBotOwnerId ensures that no value is present for BotOwnerId, not even an explicit nil
### GetRole

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetRole() interface{}`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetRoleOk() (*interface{}, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetRole(v interface{})`

SetRole sets Role field to given value.

### HasRole

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTimezone

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetTimezone() interface{}`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetTimezoneOk() (*interface{}, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetTimezone(v interface{})`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetAvatarUrl

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetAvatarUrl() interface{}`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetAvatarUrlOk() (*interface{}, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetAvatarUrl(v interface{})`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetAvatarVersion

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetAvatarVersion() interface{}`

GetAvatarVersion returns the AvatarVersion field if non-nil, zero value otherwise.

### GetAvatarVersionOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetAvatarVersionOk() (*interface{}, bool)`

GetAvatarVersionOk returns a tuple with the AvatarVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarVersion

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetAvatarVersion(v interface{})`

SetAvatarVersion sets AvatarVersion field to given value.

### HasAvatarVersion

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasAvatarVersion() bool`

HasAvatarVersion returns a boolean if a field has been set.

### SetAvatarVersionNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetAvatarVersionNil(b bool)`

 SetAvatarVersionNil sets the value for AvatarVersion to be an explicit nil

### UnsetAvatarVersion
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetAvatarVersion()`

UnsetAvatarVersion ensures that no value is present for AvatarVersion, not even an explicit nil
### GetProfileData

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetProfileData() interface{}`

GetProfileData returns the ProfileData field if non-nil, zero value otherwise.

### GetProfileDataOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetProfileDataOk() (*interface{}, bool)`

GetProfileDataOk returns a tuple with the ProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileData

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetProfileData(v interface{})`

SetProfileData sets ProfileData field to given value.

### HasProfileData

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasProfileData() bool`

HasProfileData returns a boolean if a field has been set.

### SetProfileDataNil

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetProfileDataNil(b bool)`

 SetProfileDataNil sets the value for ProfileData to be an explicit nil

### UnsetProfileData
`func (o *RegisterQueueResponseCrossRealmBotsItem) UnsetProfileData()`

UnsetProfileData ensures that no value is present for ProfileData, not even an explicit nil
### GetIsSystemBot

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsSystemBot() bool`

GetIsSystemBot returns the IsSystemBot field if non-nil, zero value otherwise.

### GetIsSystemBotOk

`func (o *RegisterQueueResponseCrossRealmBotsItem) GetIsSystemBotOk() (*bool, bool)`

GetIsSystemBotOk returns a tuple with the IsSystemBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemBot

`func (o *RegisterQueueResponseCrossRealmBotsItem) SetIsSystemBot(v bool)`

SetIsSystemBot sets IsSystemBot field to given value.

### HasIsSystemBot

`func (o *RegisterQueueResponseCrossRealmBotsItem) HasIsSystemBot() bool`

HasIsSystemBot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


