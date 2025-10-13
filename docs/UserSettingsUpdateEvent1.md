# MutedTopicsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**MutedTopics** | Pointer to [**[][]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner**]([]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner.md) | Array of tuples, where each tuple describes a muted topic. The first element of the tuple is the channel name in which the topic has to be muted, the second element is the topic name to be muted and the third element is an integer UNIX timestamp representing when the topic was muted.  **Changes**: Deprecated in Zulip 6.0 (feature level 134). Starting with this version, clients that explicitly requested the replacement &#x60;user_topic&#x60; event type when registering their event queue will not receive this legacy event type.  Before Zulip 3.0 (feature level 1), the &#x60;muted_topics&#x60; array objects were 2-item tuples and did not include the timestamp information for when the topic was muted.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf31

`func NewGetEvents200ResponseAllOfEventsInnerOneOf31() *MutedTopicsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf31 instantiates a new MutedTopicsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf31WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf31WithDefaults() *MutedTopicsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf31WithDefaults instantiates a new MutedTopicsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MutedTopicsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MutedTopicsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MutedTopicsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MutedTopicsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MutedTopicsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MutedTopicsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MutedTopicsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MutedTopicsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMutedTopics

`func (o *MutedTopicsEvent) GetMutedTopics() [][]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner`

GetMutedTopics returns the MutedTopics field if non-nil, zero value otherwise.

### GetMutedTopicsOk

`func (o *MutedTopicsEvent) GetMutedTopicsOk() (*[][]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner, bool)`

GetMutedTopicsOk returns a tuple with the MutedTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedTopics

`func (o *MutedTopicsEvent) SetMutedTopics(v [][]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner)`

SetMutedTopics sets MutedTopics field to given value.

### HasMutedTopics

`func (o *MutedTopicsEvent) HasMutedTopics() bool`

HasMutedTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


