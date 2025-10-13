# SubscriptionRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to [**[]EventEnvelopeOneOf6SubscriptionsInner**](EventEnvelopeOneOf6SubscriptionsInner.md) | A list of dictionaries, where each dictionary contains information about one of the newly unsubscribed channels.  | [optional] 

## Methods

### NewEventEnvelopeOneOf6

`func NewEventEnvelopeOneOf6() *SubscriptionRemoveEvent`

NewEventEnvelopeOneOf6 instantiates a new SubscriptionRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf6WithDefaults

`func NewEventEnvelopeOneOf6WithDefaults() *SubscriptionRemoveEvent`

NewEventEnvelopeOneOf6WithDefaults instantiates a new SubscriptionRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *SubscriptionRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SubscriptionRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SubscriptionRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SubscriptionRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetSubscriptions

`func (o *SubscriptionRemoveEvent) GetSubscriptions() []EventEnvelopeOneOf6SubscriptionsInner`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *SubscriptionRemoveEvent) GetSubscriptionsOk() (*[]EventEnvelopeOneOf6SubscriptionsInner, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *SubscriptionRemoveEvent) SetSubscriptions(v []EventEnvelopeOneOf6SubscriptionsInner)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *SubscriptionRemoveEvent) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


