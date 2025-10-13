# GetEvents200ResponseAllOfEventsInnerOneOf39Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of message being composed. Must be &#x60;\&quot;channel\&quot;&#x60; or &#x60;\&quot;direct\&quot;&#x60;.  | [optional] 
**ChannelId** | Pointer to **int32** | Only present if &#x60;type&#x60; is &#x60;\&quot;channel\&quot;&#x60;.  The unique ID of the channel to which message is being edited.  | [optional] 
**Topic** | Pointer to **string** | Only present if &#x60;type&#x60; is &#x60;\&quot;channel\&quot;&#x60;.  Topic within the channel where the message is being edited.  | [optional] 
**UserIds** | Pointer to **[]int32** | Present only if &#x60;type&#x60; is &#x60;direct&#x60;.  The user IDs of every recipient of this direct message.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf39Recipient

`func NewGetEvents200ResponseAllOfEventsInnerOneOf39Recipient() *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient`

NewGetEvents200ResponseAllOfEventsInnerOneOf39Recipient instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf39Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf39RecipientWithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf39RecipientWithDefaults() *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient`

NewGetEvents200ResponseAllOfEventsInnerOneOf39RecipientWithDefaults instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf39Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChannelId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetTopic

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUserIds

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf39Recipient) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


