# WebReloadClientEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Immediate** | Pointer to **bool** | Whether the client should fetch a new event queue immediately, rather than using a backoff strategy to avoid thundering herds. A Zulip development server uses this parameter to reload clients immediately.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf66

`func NewGetEvents200ResponseAllOfEventsInnerOneOf66() *WebReloadClientEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf66 instantiates a new WebReloadClientEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf66WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf66WithDefaults() *WebReloadClientEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf66WithDefaults instantiates a new WebReloadClientEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebReloadClientEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebReloadClientEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebReloadClientEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WebReloadClientEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *WebReloadClientEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebReloadClientEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebReloadClientEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebReloadClientEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImmediate

`func (o *WebReloadClientEvent) GetImmediate() bool`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *WebReloadClientEvent) GetImmediateOk() (*bool, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *WebReloadClientEvent) SetImmediate(v bool)`

SetImmediate sets Immediate field to given value.

### HasImmediate

`func (o *WebReloadClientEvent) HasImmediate() bool`

HasImmediate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


