# RealmUpdateDictEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **string** | Always &#x60;\&quot;default\&quot;&#x60;. Present for backwards-compatibility with older clients that predate the &#x60;update_dict&#x60; event style.  **Deprecated** and will be removed in a future release.  | [optional] 
**Data** | Pointer to [**EventEnvelopeOneOf67Data**](EventEnvelopeOneOf67Data.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf67

`func NewEventEnvelopeOneOf67() *RealmUpdateDictEvent`

NewEventEnvelopeOneOf67 instantiates a new RealmUpdateDictEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf67WithDefaults

`func NewEventEnvelopeOneOf67WithDefaults() *RealmUpdateDictEvent`

NewEventEnvelopeOneOf67WithDefaults instantiates a new RealmUpdateDictEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmUpdateDictEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmUpdateDictEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmUpdateDictEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmUpdateDictEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmUpdateDictEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmUpdateDictEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmUpdateDictEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmUpdateDictEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmUpdateDictEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmUpdateDictEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmUpdateDictEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmUpdateDictEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetProperty

`func (o *RealmUpdateDictEvent) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RealmUpdateDictEvent) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RealmUpdateDictEvent) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RealmUpdateDictEvent) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetData

`func (o *RealmUpdateDictEvent) GetData() EventEnvelopeOneOf67Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RealmUpdateDictEvent) GetDataOk() (*EventEnvelopeOneOf67Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RealmUpdateDictEvent) SetData(v EventEnvelopeOneOf67Data)`

SetData sets Data field to given value.

### HasData

`func (o *RealmUpdateDictEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


