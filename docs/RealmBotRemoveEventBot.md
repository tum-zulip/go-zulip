# RealmBotRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Bot** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf61Bot**](GetEvents200ResponseAllOfEventsInnerOneOf61Bot.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf61

`func NewGetEvents200ResponseAllOfEventsInnerOneOf61() *RealmBotRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf61 instantiates a new RealmBotRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf61WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf61WithDefaults() *RealmBotRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf61WithDefaults instantiates a new RealmBotRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmBotRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmBotRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmBotRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmBotRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmBotRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmBotRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmBotRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmBotRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmBotRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmBotRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmBotRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmBotRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetBot

`func (o *RealmBotRemoveEvent) GetBot() GetEvents200ResponseAllOfEventsInnerOneOf61Bot`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *RealmBotRemoveEvent) GetBotOk() (*GetEvents200ResponseAllOfEventsInnerOneOf61Bot, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *RealmBotRemoveEvent) SetBot(v GetEvents200ResponseAllOfEventsInnerOneOf61Bot)`

SetBot sets Bot field to given value.

### HasBot

`func (o *RealmBotRemoveEvent) HasBot() bool`

HasBot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


