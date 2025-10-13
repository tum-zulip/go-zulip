# GetOwnUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**AvatarUrl** | Pointer to **string** | URL for the requesting user&#39;s avatar.  **Changes**: New in Zulip 2.1.0.  | [optional] 
**AvatarVersion** | Pointer to **int32** | Version for the requesting user&#39;s avatar. Used for cache-busting requests for the user&#39;s avatar. Clients generally shouldn&#39;t need to use this; most avatar URLs sent by Zulip will already end with &#x60;?v&#x3D;{avatar_version}&#x60;.  **Changes**: New in Zulip 3.0 (feature level 10).  | [optional] 
**Email** | Pointer to **string** | Zulip API email of the requesting user.  | [optional] 
**FullName** | Pointer to **string** | Full name of the requesting user.  | [optional] 
**IsAdmin** | Pointer to **bool** | A boolean indicating if the requesting user is an admin.  | [optional] 
**IsOwner** | Pointer to **bool** | A boolean indicating if the requesting user is an organization owner.  **Changes**: New in Zulip 3.0 (feature level 8).  | [optional] 
**Role** | Pointer to **int32** | [Organization-level role](/api/roles-and-permissions) of the requesting user. Possible values are:  - 100 &#x3D; Organization owner - 200 &#x3D; Organization administrator - 300 &#x3D; Organization moderator - 400 &#x3D; Member - 600 &#x3D; Guest  **Changes**: New in Zulip 4.0 (feature level 59).  | [optional] 
**IsGuest** | Pointer to **bool** | A boolean indicating if the requesting user is a guest.  **Changes**: New in Zulip 3.0 (feature level 10).  | [optional] 
**IsBot** | Pointer to **bool** | A boolean indicating if the requesting user is a bot.  | [optional] 
**IsActive** | Pointer to **bool** | A boolean specifying whether the requesting user account has been deactivated.  **Changes**: New in Zulip 3.0 (feature level 10).  | [optional] 
**Timezone** | Pointer to **string** | The IANA identifier of the requesting user&#39;s [profile time zone](/help/change-your-timezone), which is used primarily to display the user&#39;s local time to other users.  **Changes**: New in Zulip 3.0 (feature level 10).  | [optional] 
**DateJoined** | Pointer to **string** | The time the requesting user&#39;s account was created.  **Changes**: New in Zulip 3.0 (feature level 10).  | [optional] 
**MaxMessageId** | Pointer to **int32** | The integer ID of the last message received by the requesting user&#39;s account.  **Deprecated**. We plan to remove this in favor of recommending using &#x60;GET /messages&#x60; with &#x60;\&quot;anchor\&quot;: \&quot;newest\&quot;&#x60;.  | [optional] 
**UserId** | Pointer to **int32** | The user&#39;s ID.  | [optional] 
**DeliveryEmail** | Pointer to **string** | The requesting user&#39;s real email address.  **Changes**: Prior to Zulip 7.0 (feature level 163), this field was present only when &#x60;email_address_visibility&#x60; was restricted and the requesting user had permission to access realm users&#39; emails. As of this feature level, this field is always present.  | [optional] 
**ProfileData** | Pointer to [**map[string]ProfileDataValue**](ProfileDataValue.md) | Only present if &#x60;is_bot&#x60; is false; bots can&#39;t have custom profile fields.  A dictionary containing custom profile field data for the user. Each entry maps the integer ID of a custom profile field in the organization to a dictionary containing the user&#39;s data for that field. Generally the data includes just a single &#x60;value&#x60; key; for those custom profile fields supporting Markdown, a &#x60;rendered_value&#x60; key will also be present.  | [optional] 

## Methods

### NewGetOwnUser200Response

`func NewGetOwnUser200Response(result interface{}, msg interface{}, ) *GetOwnUser200Response`

NewGetOwnUser200Response instantiates a new GetOwnUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOwnUser200ResponseWithDefaults

`func NewGetOwnUser200ResponseWithDefaults() *GetOwnUser200Response`

NewGetOwnUser200ResponseWithDefaults instantiates a new GetOwnUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetOwnUser200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetOwnUser200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetOwnUser200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetOwnUser200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetOwnUser200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetOwnUser200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetOwnUser200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetOwnUser200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetOwnUser200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetOwnUser200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetOwnUser200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetOwnUser200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetOwnUser200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetOwnUser200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetOwnUser200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetOwnUser200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetAvatarUrl

`func (o *GetOwnUser200Response) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GetOwnUser200Response) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GetOwnUser200Response) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GetOwnUser200Response) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetAvatarVersion

`func (o *GetOwnUser200Response) GetAvatarVersion() int32`

GetAvatarVersion returns the AvatarVersion field if non-nil, zero value otherwise.

### GetAvatarVersionOk

`func (o *GetOwnUser200Response) GetAvatarVersionOk() (*int32, bool)`

GetAvatarVersionOk returns a tuple with the AvatarVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarVersion

`func (o *GetOwnUser200Response) SetAvatarVersion(v int32)`

SetAvatarVersion sets AvatarVersion field to given value.

### HasAvatarVersion

`func (o *GetOwnUser200Response) HasAvatarVersion() bool`

HasAvatarVersion returns a boolean if a field has been set.

### GetEmail

`func (o *GetOwnUser200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetOwnUser200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetOwnUser200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetOwnUser200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullName

`func (o *GetOwnUser200Response) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetOwnUser200Response) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetOwnUser200Response) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GetOwnUser200Response) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetIsAdmin

`func (o *GetOwnUser200Response) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *GetOwnUser200Response) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *GetOwnUser200Response) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *GetOwnUser200Response) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsOwner

`func (o *GetOwnUser200Response) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *GetOwnUser200Response) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *GetOwnUser200Response) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *GetOwnUser200Response) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetRole

`func (o *GetOwnUser200Response) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetOwnUser200Response) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetOwnUser200Response) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetOwnUser200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetIsGuest

`func (o *GetOwnUser200Response) GetIsGuest() bool`

GetIsGuest returns the IsGuest field if non-nil, zero value otherwise.

### GetIsGuestOk

`func (o *GetOwnUser200Response) GetIsGuestOk() (*bool, bool)`

GetIsGuestOk returns a tuple with the IsGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuest

`func (o *GetOwnUser200Response) SetIsGuest(v bool)`

SetIsGuest sets IsGuest field to given value.

### HasIsGuest

`func (o *GetOwnUser200Response) HasIsGuest() bool`

HasIsGuest returns a boolean if a field has been set.

### GetIsBot

`func (o *GetOwnUser200Response) GetIsBot() bool`

GetIsBot returns the IsBot field if non-nil, zero value otherwise.

### GetIsBotOk

`func (o *GetOwnUser200Response) GetIsBotOk() (*bool, bool)`

GetIsBotOk returns a tuple with the IsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBot

`func (o *GetOwnUser200Response) SetIsBot(v bool)`

SetIsBot sets IsBot field to given value.

### HasIsBot

`func (o *GetOwnUser200Response) HasIsBot() bool`

HasIsBot returns a boolean if a field has been set.

### GetIsActive

`func (o *GetOwnUser200Response) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GetOwnUser200Response) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GetOwnUser200Response) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *GetOwnUser200Response) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTimezone

`func (o *GetOwnUser200Response) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetOwnUser200Response) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetOwnUser200Response) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetOwnUser200Response) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDateJoined

`func (o *GetOwnUser200Response) GetDateJoined() string`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *GetOwnUser200Response) GetDateJoinedOk() (*string, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *GetOwnUser200Response) SetDateJoined(v string)`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *GetOwnUser200Response) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### GetMaxMessageId

`func (o *GetOwnUser200Response) GetMaxMessageId() int32`

GetMaxMessageId returns the MaxMessageId field if non-nil, zero value otherwise.

### GetMaxMessageIdOk

`func (o *GetOwnUser200Response) GetMaxMessageIdOk() (*int32, bool)`

GetMaxMessageIdOk returns a tuple with the MaxMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageId

`func (o *GetOwnUser200Response) SetMaxMessageId(v int32)`

SetMaxMessageId sets MaxMessageId field to given value.

### HasMaxMessageId

`func (o *GetOwnUser200Response) HasMaxMessageId() bool`

HasMaxMessageId returns a boolean if a field has been set.

### GetUserId

`func (o *GetOwnUser200Response) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetOwnUser200Response) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetOwnUser200Response) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetOwnUser200Response) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDeliveryEmail

`func (o *GetOwnUser200Response) GetDeliveryEmail() string`

GetDeliveryEmail returns the DeliveryEmail field if non-nil, zero value otherwise.

### GetDeliveryEmailOk

`func (o *GetOwnUser200Response) GetDeliveryEmailOk() (*string, bool)`

GetDeliveryEmailOk returns a tuple with the DeliveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEmail

`func (o *GetOwnUser200Response) SetDeliveryEmail(v string)`

SetDeliveryEmail sets DeliveryEmail field to given value.

### HasDeliveryEmail

`func (o *GetOwnUser200Response) HasDeliveryEmail() bool`

HasDeliveryEmail returns a boolean if a field has been set.

### GetProfileData

`func (o *GetOwnUser200Response) GetProfileData() map[string]ProfileDataValue`

GetProfileData returns the ProfileData field if non-nil, zero value otherwise.

### GetProfileDataOk

`func (o *GetOwnUser200Response) GetProfileDataOk() (*map[string]ProfileDataValue, bool)`

GetProfileDataOk returns a tuple with the ProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileData

`func (o *GetOwnUser200Response) SetProfileData(v map[string]ProfileDataValue)`

SetProfileData sets ProfileData field to given value.

### HasProfileData

`func (o *GetOwnUser200Response) HasProfileData() bool`

HasProfileData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


