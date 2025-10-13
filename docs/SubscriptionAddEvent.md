# SubscriptionAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to [**[]Subscription**](Subscription.md) | A list of dictionaries where each dictionary contains information about one of the subscribed channels.  **Changes**: Removed &#x60;email_address&#x60; field from the dictionary in Zulip 8.0 (feature level 226).  Removed &#x60;role&#x60; field from the dictionary in Zulip 6.0 (feature level 133).  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf5

`func NewGetEvents200ResponseAllOfEventsInnerOneOf5() *SubscriptionAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf5 instantiates a new SubscriptionAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf5WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf5WithDefaults() *SubscriptionAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf5WithDefaults instantiates a new SubscriptionAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *SubscriptionAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SubscriptionAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SubscriptionAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SubscriptionAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetSubscriptions

`func (o *SubscriptionAddEvent) GetSubscriptions() []Subscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *SubscriptionAddEvent) GetSubscriptionsOk() (*[]Subscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *SubscriptionAddEvent) SetSubscriptions(v []Subscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *SubscriptionAddEvent) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


