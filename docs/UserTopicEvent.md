# UserTopicEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**StreamId** | Pointer to **int32** | The ID of the channel to which the topic belongs.  | [optional] 
**TopicName** | Pointer to **string** | The name of the topic.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**LastUpdated** | Pointer to **int32** | An integer UNIX timestamp representing when the user-topic relationship was last changed.  | [optional] 
**VisibilityPolicy** | Pointer to **int32** | An integer indicating the user&#39;s visibility preferences for the topic, such as whether the topic is muted.  - 0 &#x3D; None. Used to indicate that the user no   longer has a special visibility policy for this topic. - 1 &#x3D; Muted. Used to record [muted topics](zulip.com/help/mute-a-topic. - 2 &#x3D; Unmuted. Used to record unmuted topics. - 3 &#x3D; Followed. Used to record [followed topics](zulip.com/help/follow-a-topic.  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.  In Zulip 7.0 (feature level 170), added unmuted as a visibility policy option.  | [optional] 

## Methods

### NewEventEnvelopeOneOf32

`func NewEventEnvelopeOneOf32() *UserTopicEvent`

NewEventEnvelopeOneOf32 instantiates a new UserTopicEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf32WithDefaults

`func NewEventEnvelopeOneOf32WithDefaults() *UserTopicEvent`

NewEventEnvelopeOneOf32WithDefaults instantiates a new UserTopicEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserTopicEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserTopicEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserTopicEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserTopicEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserTopicEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserTopicEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserTopicEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserTopicEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStreamId

`func (o *UserTopicEvent) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *UserTopicEvent) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *UserTopicEvent) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *UserTopicEvent) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTopicName

`func (o *UserTopicEvent) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *UserTopicEvent) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *UserTopicEvent) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *UserTopicEvent) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UserTopicEvent) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserTopicEvent) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserTopicEvent) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserTopicEvent) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVisibilityPolicy

`func (o *UserTopicEvent) GetVisibilityPolicy() int32`

GetVisibilityPolicy returns the VisibilityPolicy field if non-nil, zero value otherwise.

### GetVisibilityPolicyOk

`func (o *UserTopicEvent) GetVisibilityPolicyOk() (*int32, bool)`

GetVisibilityPolicyOk returns a tuple with the VisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityPolicy

`func (o *UserTopicEvent) SetVisibilityPolicy(v int32)`

SetVisibilityPolicy sets VisibilityPolicy field to given value.

### HasVisibilityPolicy

`func (o *UserTopicEvent) HasVisibilityPolicy() bool`

HasVisibilityPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


