# SubscriptionPeerRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**StreamIds** | Pointer to **[]int32** | The IDs of the channels from which the users have been unsubscribed from.  When a user is deactivated, the server will send this event removing the user&#39;s subscriptions before the &#x60;realm_user&#x60; event for the user&#39;s deactivation.  **Changes**: Before Zulip 10.0 (feature level 377), this event was not sent on user deactivation. Clients supporting older server versions and maintaining peer subscriber data need to remove all channel subscriptions for a user when processing the &#x60;realm_user&#x60; event with &#x60;op&#x3D;\&quot;remove\&quot;&#x60;.  **Changes**: New in Zulip 4.0 (feature level 35), replacing the &#x60;stream_id&#x60; integer.  | [optional] 
**UserIds** | Pointer to **[]int32** | The IDs of the users who have been unsubscribed.  **Changes**: New in Zulip 4.0 (feature level 35), replacing the &#x60;user_id&#x60; integer.  | [optional] 

## Methods

### NewEventEnvelopeOneOf9

`func NewEventEnvelopeOneOf9() *SubscriptionPeerRemoveEvent`

NewEventEnvelopeOneOf9 instantiates a new SubscriptionPeerRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf9WithDefaults

`func NewEventEnvelopeOneOf9WithDefaults() *SubscriptionPeerRemoveEvent`

NewEventEnvelopeOneOf9WithDefaults instantiates a new SubscriptionPeerRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionPeerRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPeerRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPeerRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionPeerRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionPeerRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionPeerRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionPeerRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionPeerRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *SubscriptionPeerRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SubscriptionPeerRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SubscriptionPeerRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SubscriptionPeerRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetStreamIds

`func (o *SubscriptionPeerRemoveEvent) GetStreamIds() []int32`

GetStreamIds returns the StreamIds field if non-nil, zero value otherwise.

### GetStreamIdsOk

`func (o *SubscriptionPeerRemoveEvent) GetStreamIdsOk() (*[]int32, bool)`

GetStreamIdsOk returns a tuple with the StreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamIds

`func (o *SubscriptionPeerRemoveEvent) SetStreamIds(v []int32)`

SetStreamIds sets StreamIds field to given value.

### HasStreamIds

`func (o *SubscriptionPeerRemoveEvent) HasStreamIds() bool`

HasStreamIds returns a boolean if a field has been set.

### GetUserIds

`func (o *SubscriptionPeerRemoveEvent) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *SubscriptionPeerRemoveEvent) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *SubscriptionPeerRemoveEvent) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *SubscriptionPeerRemoveEvent) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


