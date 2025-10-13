# StreamCreateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Streams** | Pointer to [**[]BasicChannel**](BasicChannel.md) | Array of objects, each containing details about the newly added channel(s).  | [optional] 

## Methods

### NewEventEnvelopeOneOf16

`func NewEventEnvelopeOneOf16() *StreamCreateEvent`

NewEventEnvelopeOneOf16 instantiates a new StreamCreateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf16WithDefaults

`func NewEventEnvelopeOneOf16WithDefaults() *StreamCreateEvent`

NewEventEnvelopeOneOf16WithDefaults instantiates a new StreamCreateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamCreateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamCreateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamCreateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StreamCreateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *StreamCreateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamCreateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamCreateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamCreateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *StreamCreateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *StreamCreateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *StreamCreateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *StreamCreateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetStreams

`func (o *StreamCreateEvent) GetStreams() []BasicChannel`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *StreamCreateEvent) GetStreamsOk() (*[]BasicChannel, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *StreamCreateEvent) SetStreams(v []BasicChannel)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *StreamCreateEvent) HasStreams() bool`

HasStreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


