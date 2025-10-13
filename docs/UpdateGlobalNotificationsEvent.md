# UpdateGlobalNotificationsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** | The Zulip API email of the user.  | [optional] 
**NotificationName** | Pointer to **string** | Name of the changed notification setting.  | [optional] 
**Setting** | Pointer to [**EventEnvelopeOneOf1Setting**](EventEnvelopeOneOf1Setting.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf2

`func NewEventEnvelopeOneOf2() *UpdateGlobalNotificationsEvent`

NewEventEnvelopeOneOf2 instantiates a new UpdateGlobalNotificationsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf2WithDefaults

`func NewEventEnvelopeOneOf2WithDefaults() *UpdateGlobalNotificationsEvent`

NewEventEnvelopeOneOf2WithDefaults instantiates a new UpdateGlobalNotificationsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateGlobalNotificationsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateGlobalNotificationsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateGlobalNotificationsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateGlobalNotificationsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateGlobalNotificationsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateGlobalNotificationsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateGlobalNotificationsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateGlobalNotificationsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *UpdateGlobalNotificationsEvent) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateGlobalNotificationsEvent) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateGlobalNotificationsEvent) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateGlobalNotificationsEvent) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetNotificationName

`func (o *UpdateGlobalNotificationsEvent) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *UpdateGlobalNotificationsEvent) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *UpdateGlobalNotificationsEvent) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *UpdateGlobalNotificationsEvent) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetSetting

`func (o *UpdateGlobalNotificationsEvent) GetSetting() EventEnvelopeOneOf1Setting`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *UpdateGlobalNotificationsEvent) GetSettingOk() (*EventEnvelopeOneOf1Setting, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *UpdateGlobalNotificationsEvent) SetSetting(v EventEnvelopeOneOf1Setting)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *UpdateGlobalNotificationsEvent) HasSetting() bool`

HasSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


