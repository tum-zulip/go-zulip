# RealmBotDeleteEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Bot** | Pointer to [**EventEnvelopeOneOf62Bot**](EventEnvelopeOneOf62Bot.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf62

`func NewEventEnvelopeOneOf62() *RealmBotDeleteEvent`

NewEventEnvelopeOneOf62 instantiates a new RealmBotDeleteEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf62WithDefaults

`func NewEventEnvelopeOneOf62WithDefaults() *RealmBotDeleteEvent`

NewEventEnvelopeOneOf62WithDefaults instantiates a new RealmBotDeleteEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmBotDeleteEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmBotDeleteEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmBotDeleteEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmBotDeleteEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmBotDeleteEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmBotDeleteEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmBotDeleteEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmBotDeleteEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmBotDeleteEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmBotDeleteEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmBotDeleteEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmBotDeleteEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetBot

`func (o *RealmBotDeleteEvent) GetBot() EventEnvelopeOneOf62Bot`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *RealmBotDeleteEvent) GetBotOk() (*EventEnvelopeOneOf62Bot, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *RealmBotDeleteEvent) SetBot(v EventEnvelopeOneOf62Bot)`

SetBot sets Bot field to given value.

### HasBot

`func (o *RealmBotDeleteEvent) HasBot() bool`

HasBot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


