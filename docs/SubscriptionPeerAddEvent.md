# SubscriptionPeerAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**StreamIds** | Pointer to **[]int32** | The IDs of channels that have new or updated subscriber data.  **Changes**: New in Zulip 4.0 (feature level 35), replacing the &#x60;stream_id&#x60; integer.  | [optional] 
**UserIds** | Pointer to **[]int32** | The IDs of the users who are newly visible as subscribed to the specified channels.  **Changes**: New in Zulip 4.0 (feature level 35), replacing the &#x60;user_id&#x60; integer.  | [optional] 

## Methods

### NewEventEnvelopeOneOf8

`func NewEventEnvelopeOneOf8() *SubscriptionPeerAddEvent`

NewEventEnvelopeOneOf8 instantiates a new SubscriptionPeerAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf8WithDefaults

`func NewEventEnvelopeOneOf8WithDefaults() *SubscriptionPeerAddEvent`

NewEventEnvelopeOneOf8WithDefaults instantiates a new SubscriptionPeerAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionPeerAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPeerAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPeerAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionPeerAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionPeerAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionPeerAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionPeerAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionPeerAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *SubscriptionPeerAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SubscriptionPeerAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SubscriptionPeerAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SubscriptionPeerAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetStreamIds

`func (o *SubscriptionPeerAddEvent) GetStreamIds() []int32`

GetStreamIds returns the StreamIds field if non-nil, zero value otherwise.

### GetStreamIdsOk

`func (o *SubscriptionPeerAddEvent) GetStreamIdsOk() (*[]int32, bool)`

GetStreamIdsOk returns a tuple with the StreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamIds

`func (o *SubscriptionPeerAddEvent) SetStreamIds(v []int32)`

SetStreamIds sets StreamIds field to given value.

### HasStreamIds

`func (o *SubscriptionPeerAddEvent) HasStreamIds() bool`

HasStreamIds returns a boolean if a field has been set.

### GetUserIds

`func (o *SubscriptionPeerAddEvent) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *SubscriptionPeerAddEvent) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *SubscriptionPeerAddEvent) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *SubscriptionPeerAddEvent) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


