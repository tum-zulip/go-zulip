# RealmBotUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Bot** | Pointer to [**BasicBot**](BasicBot.md) | Object containing details about the changed bot. It contains two properties: the user ID of the bot and the property to be changed. The changed property is one of the remaining properties listed below.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf60

`func NewGetEvents200ResponseAllOfEventsInnerOneOf60() *RealmBotUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf60 instantiates a new RealmBotUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf60WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf60WithDefaults() *RealmBotUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf60WithDefaults instantiates a new RealmBotUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmBotUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmBotUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmBotUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmBotUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmBotUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmBotUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmBotUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmBotUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmBotUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmBotUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmBotUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmBotUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetBot

`func (o *RealmBotUpdateEvent) GetBot() BasicBot`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *RealmBotUpdateEvent) GetBotOk() (*BasicBot, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *RealmBotUpdateEvent) SetBot(v BasicBot)`

SetBot sets Bot field to given value.

### HasBot

`func (o *RealmBotUpdateEvent) HasBot() bool`

HasBot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


