# RegisterQueueResponseServerSupportedPermissionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Realm** | Pointer to [**map[string]GroupPermissionSetting**](GroupPermissionSetting.md) | Configuration for realm level group permission settings.  | [optional] 
**Stream** | Pointer to [**map[string]GroupPermissionSetting**](GroupPermissionSetting.md) | Configuration for channel level group permission settings.  | [optional] 
**Group** | Pointer to [**map[string]GroupPermissionSetting**](GroupPermissionSetting.md) | Configuration for group level group permission settings.  | [optional] 

## Methods

### NewRegisterQueueResponseServerSupportedPermissionSettings

`func NewRegisterQueueResponseServerSupportedPermissionSettings() *RegisterQueueResponseServerSupportedPermissionSettings`

NewRegisterQueueResponseServerSupportedPermissionSettings instantiates a new RegisterQueueResponseServerSupportedPermissionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseServerSupportedPermissionSettingsWithDefaults

`func NewRegisterQueueResponseServerSupportedPermissionSettingsWithDefaults() *RegisterQueueResponseServerSupportedPermissionSettings`

NewRegisterQueueResponseServerSupportedPermissionSettingsWithDefaults instantiates a new RegisterQueueResponseServerSupportedPermissionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealm

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) GetRealm() map[string]GroupPermissionSetting`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) GetRealmOk() (*map[string]GroupPermissionSetting, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) SetRealm(v map[string]GroupPermissionSetting)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetStream

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) GetStream() map[string]GroupPermissionSetting`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) GetStreamOk() (*map[string]GroupPermissionSetting, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) SetStream(v map[string]GroupPermissionSetting)`

SetStream sets Stream field to given value.

### HasStream

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetGroup

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) GetGroup() map[string]GroupPermissionSetting`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) GetGroupOk() (*map[string]GroupPermissionSetting, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) SetGroup(v map[string]GroupPermissionSetting)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *RegisterQueueResponseServerSupportedPermissionSettings) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


