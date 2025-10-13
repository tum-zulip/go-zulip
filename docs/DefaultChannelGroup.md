# DefaultChannelGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the default channel group.  | [optional] 
**Description** | Pointer to **string** | Description of the default channel group.  | [optional] 
**Id** | Pointer to **int32** | The ID of the default channel group.  | [optional] 
**Streams** | Pointer to **[]int32** | An array of IDs of all the channels in the default stream group.  **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single stream in the default stream group.  | [optional] 

## Methods

### NewDefaultChannelGroup

`func NewDefaultChannelGroup() *DefaultChannelGroup`

NewDefaultChannelGroup instantiates a new DefaultChannelGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultChannelGroupWithDefaults

`func NewDefaultChannelGroupWithDefaults() *DefaultChannelGroup`

NewDefaultChannelGroupWithDefaults instantiates a new DefaultChannelGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DefaultChannelGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefaultChannelGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefaultChannelGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefaultChannelGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DefaultChannelGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DefaultChannelGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DefaultChannelGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DefaultChannelGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DefaultChannelGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefaultChannelGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefaultChannelGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DefaultChannelGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStreams

`func (o *DefaultChannelGroup) GetStreams() []int32`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *DefaultChannelGroup) GetStreamsOk() (*[]int32, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *DefaultChannelGroup) SetStreams(v []int32)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *DefaultChannelGroup) HasStreams() bool`

HasStreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


