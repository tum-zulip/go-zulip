# SubscriptionUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**StreamId** | Pointer to **int32** | The ID of the channel whose subscription details have changed.  | [optional] 
**Property** | Pointer to **string** | The property of the subscription which has changed. For details on the various subscription properties that a user can change, see [POST /users/me/subscriptions/properties](zulip.com/api/update-subscription-settings.  Clients should generally handle an unknown property received here without crashing, since that will naturally happen when connecting to a Zulip server running a new version that adds a new subscription property.  **Changes**: As of Zulip 6.0 (feature level 139), updates to the &#x60;is_muted&#x60; property or the deprecated &#x60;in_home_view&#x60; property will send two &#x60;subscription&#x60; update events, one for each property, to support clients fully migrating to use the &#x60;is_muted&#x60; property. Prior to this feature level, updates to either property only sent one event with the deprecated &#x60;in_home_view&#x60; property.  | [optional] 
**Value** | Pointer to [**EventEnvelopeOneOf7Value**](EventEnvelopeOneOf7Value.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf7

`func NewEventEnvelopeOneOf7() *SubscriptionUpdateEvent`

NewEventEnvelopeOneOf7 instantiates a new SubscriptionUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf7WithDefaults

`func NewEventEnvelopeOneOf7WithDefaults() *SubscriptionUpdateEvent`

NewEventEnvelopeOneOf7WithDefaults instantiates a new SubscriptionUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *SubscriptionUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SubscriptionUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SubscriptionUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SubscriptionUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetStreamId

`func (o *SubscriptionUpdateEvent) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *SubscriptionUpdateEvent) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *SubscriptionUpdateEvent) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *SubscriptionUpdateEvent) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetProperty

`func (o *SubscriptionUpdateEvent) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SubscriptionUpdateEvent) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SubscriptionUpdateEvent) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SubscriptionUpdateEvent) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *SubscriptionUpdateEvent) GetValue() EventEnvelopeOneOf7Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SubscriptionUpdateEvent) GetValueOk() (*EventEnvelopeOneOf7Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SubscriptionUpdateEvent) SetValue(v EventEnvelopeOneOf7Value)`

SetValue sets Value field to given value.

### HasValue

`func (o *SubscriptionUpdateEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


