# RealmUserSettingsDefaultsUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **string** | The name of the property that was changed.  | [optional] 
**Value** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf68Value**](GetEvents200ResponseAllOfEventsInnerOneOf68Value.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf68

`func NewGetEvents200ResponseAllOfEventsInnerOneOf68() *RealmUserSettingsDefaultsUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf68 instantiates a new RealmUserSettingsDefaultsUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf68WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf68WithDefaults() *RealmUserSettingsDefaultsUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf68WithDefaults instantiates a new RealmUserSettingsDefaultsUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmUserSettingsDefaultsUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmUserSettingsDefaultsUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmUserSettingsDefaultsUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmUserSettingsDefaultsUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmUserSettingsDefaultsUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmUserSettingsDefaultsUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetProperty

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RealmUserSettingsDefaultsUpdateEvent) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RealmUserSettingsDefaultsUpdateEvent) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetValue() GetEvents200ResponseAllOfEventsInnerOneOf68Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RealmUserSettingsDefaultsUpdateEvent) GetValueOk() (*GetEvents200ResponseAllOfEventsInnerOneOf68Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RealmUserSettingsDefaultsUpdateEvent) SetValue(v GetEvents200ResponseAllOfEventsInnerOneOf68Value)`

SetValue sets Value field to given value.

### HasValue

`func (o *RealmUserSettingsDefaultsUpdateEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


