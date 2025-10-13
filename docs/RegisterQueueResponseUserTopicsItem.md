# RegisterQueueResponseUserTopicsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamId** | Pointer to **int32** | The ID of the channel to which the topic belongs.  | [optional] 
**TopicName** | Pointer to **string** | The name of the topic.  Note that the empty string topic may have been rewritten by the server to the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue response depending on the value of the &#x60;empty_topic_name&#x60; [client capability][client-capabilities].  **Changes**: The &#x60;empty_topic_name&#x60; client capability is new in Zulip 10.0 (feature level 334).  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**LastUpdated** | Pointer to **int32** | An integer UNIX timestamp representing when the user-topic relationship was changed.  | [optional] 
**VisibilityPolicy** | Pointer to **int32** | An integer indicating the user&#39;s visibility configuration for the topic.  - 1 &#x3D; Muted. Used to record [muted topics](zulip.com/help/mute-a-topic. - 2 &#x3D; Unmuted. Used to record [unmuted topics](zulip.com/help/mute-a-topic. - 3 &#x3D; Followed. Used to record [followed topics](zulip.com/help/follow-a-topic.  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.  In Zulip 7.0 (feature level 170), added unmuted as a visibility policy option.  | [optional] 

## Methods

### NewRegisterQueueResponseUserTopicsItem

`func NewRegisterQueueResponseUserTopicsItem() *RegisterQueueResponseUserTopicsItem`

NewRegisterQueueResponseUserTopicsItem instantiates a new RegisterQueueResponseUserTopicsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseUserTopicsItemWithDefaults

`func NewRegisterQueueResponseUserTopicsItemWithDefaults() *RegisterQueueResponseUserTopicsItem`

NewRegisterQueueResponseUserTopicsItemWithDefaults instantiates a new RegisterQueueResponseUserTopicsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamId

`func (o *RegisterQueueResponseUserTopicsItem) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *RegisterQueueResponseUserTopicsItem) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *RegisterQueueResponseUserTopicsItem) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *RegisterQueueResponseUserTopicsItem) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTopicName

`func (o *RegisterQueueResponseUserTopicsItem) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *RegisterQueueResponseUserTopicsItem) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *RegisterQueueResponseUserTopicsItem) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *RegisterQueueResponseUserTopicsItem) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RegisterQueueResponseUserTopicsItem) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RegisterQueueResponseUserTopicsItem) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RegisterQueueResponseUserTopicsItem) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RegisterQueueResponseUserTopicsItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVisibilityPolicy

`func (o *RegisterQueueResponseUserTopicsItem) GetVisibilityPolicy() int32`

GetVisibilityPolicy returns the VisibilityPolicy field if non-nil, zero value otherwise.

### GetVisibilityPolicyOk

`func (o *RegisterQueueResponseUserTopicsItem) GetVisibilityPolicyOk() (*int32, bool)`

GetVisibilityPolicyOk returns a tuple with the VisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityPolicy

`func (o *RegisterQueueResponseUserTopicsItem) SetVisibilityPolicy(v int32)`

SetVisibilityPolicy sets VisibilityPolicy field to given value.

### HasVisibilityPolicy

`func (o *RegisterQueueResponseUserTopicsItem) HasVisibilityPolicy() bool`

HasVisibilityPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


