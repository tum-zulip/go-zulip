# StreamDeleteEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Streams** | Pointer to [**[]EventEnvelopeOneOf17StreamsInner**](EventEnvelopeOneOf17StreamsInner.md) | Array of objects, each containing ID of the channel that was deleted.  **Changes**: **Deprecated** in Zulip 10.0 (feature level 343) and will be removed in a future release. Previously, these objects additionally contained all the standard fields for a channel object.  | [optional] 
**StreamIds** | Pointer to **[]int32** | Array containing the IDs of the channels that were deleted.  **Changes**: New in Zulip 10.0 (feature level 343). Previously, these IDs were available only via the legacy &#x60;streams&#x60; array.  | [optional] 

## Methods

### NewEventEnvelopeOneOf17

`func NewEventEnvelopeOneOf17() *StreamDeleteEvent`

NewEventEnvelopeOneOf17 instantiates a new StreamDeleteEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf17WithDefaults

`func NewEventEnvelopeOneOf17WithDefaults() *StreamDeleteEvent`

NewEventEnvelopeOneOf17WithDefaults instantiates a new StreamDeleteEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamDeleteEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamDeleteEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamDeleteEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StreamDeleteEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *StreamDeleteEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamDeleteEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamDeleteEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamDeleteEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *StreamDeleteEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *StreamDeleteEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *StreamDeleteEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *StreamDeleteEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetStreams

`func (o *StreamDeleteEvent) GetStreams() []EventEnvelopeOneOf17StreamsInner`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *StreamDeleteEvent) GetStreamsOk() (*[]EventEnvelopeOneOf17StreamsInner, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *StreamDeleteEvent) SetStreams(v []EventEnvelopeOneOf17StreamsInner)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *StreamDeleteEvent) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetStreamIds

`func (o *StreamDeleteEvent) GetStreamIds() []int32`

GetStreamIds returns the StreamIds field if non-nil, zero value otherwise.

### GetStreamIdsOk

`func (o *StreamDeleteEvent) GetStreamIdsOk() (*[]int32, bool)`

GetStreamIdsOk returns a tuple with the StreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamIds

`func (o *StreamDeleteEvent) SetStreamIds(v []int32)`

SetStreamIds sets StreamIds field to given value.

### HasStreamIds

`func (o *StreamDeleteEvent) HasStreamIds() bool`

HasStreamIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


