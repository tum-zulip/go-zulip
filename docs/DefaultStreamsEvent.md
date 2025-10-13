# DefaultStreamsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DefaultStreams** | Pointer to **[]int32** | An array of IDs of all the [default channels](zulip.com/help/set-default-streams-for-new-users in the organization.  **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single default stream for the Zulip organization.  | [optional] 

## Methods

### NewEventEnvelopeOneOf29

`func NewEventEnvelopeOneOf29() *DefaultStreamsEvent`

NewEventEnvelopeOneOf29 instantiates a new DefaultStreamsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf29WithDefaults

`func NewEventEnvelopeOneOf29WithDefaults() *DefaultStreamsEvent`

NewEventEnvelopeOneOf29WithDefaults instantiates a new DefaultStreamsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DefaultStreamsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefaultStreamsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefaultStreamsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DefaultStreamsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DefaultStreamsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DefaultStreamsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DefaultStreamsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DefaultStreamsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultStreams

`func (o *DefaultStreamsEvent) GetDefaultStreams() []int32`

GetDefaultStreams returns the DefaultStreams field if non-nil, zero value otherwise.

### GetDefaultStreamsOk

`func (o *DefaultStreamsEvent) GetDefaultStreamsOk() (*[]int32, bool)`

GetDefaultStreamsOk returns a tuple with the DefaultStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStreams

`func (o *DefaultStreamsEvent) SetDefaultStreams(v []int32)`

SetDefaultStreams sets DefaultStreams field to given value.

### HasDefaultStreams

`func (o *DefaultStreamsEvent) HasDefaultStreams() bool`

HasDefaultStreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


