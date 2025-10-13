# RegisterQueueResponseRecentPrivateConversationsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMessageId** | Pointer to **int32** | The highest message ID of the conversation, intended to support sorting the conversations by recency.  | [optional] 
**UserIds** | Pointer to **[]int32** | The list of users other than the current user in the direct message conversation. This will be an empty list for direct messages sent to oneself.  | [optional] 

## Methods

### NewRegisterQueueResponseRecentPrivateConversationsItem

`func NewRegisterQueueResponseRecentPrivateConversationsItem() *RegisterQueueResponseRecentPrivateConversationsItem`

NewRegisterQueueResponseRecentPrivateConversationsItem instantiates a new RegisterQueueResponseRecentPrivateConversationsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseRecentPrivateConversationsItemWithDefaults

`func NewRegisterQueueResponseRecentPrivateConversationsItemWithDefaults() *RegisterQueueResponseRecentPrivateConversationsItem`

NewRegisterQueueResponseRecentPrivateConversationsItemWithDefaults instantiates a new RegisterQueueResponseRecentPrivateConversationsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMessageId

`func (o *RegisterQueueResponseRecentPrivateConversationsItem) GetMaxMessageId() int32`

GetMaxMessageId returns the MaxMessageId field if non-nil, zero value otherwise.

### GetMaxMessageIdOk

`func (o *RegisterQueueResponseRecentPrivateConversationsItem) GetMaxMessageIdOk() (*int32, bool)`

GetMaxMessageIdOk returns a tuple with the MaxMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageId

`func (o *RegisterQueueResponseRecentPrivateConversationsItem) SetMaxMessageId(v int32)`

SetMaxMessageId sets MaxMessageId field to given value.

### HasMaxMessageId

`func (o *RegisterQueueResponseRecentPrivateConversationsItem) HasMaxMessageId() bool`

HasMaxMessageId returns a boolean if a field has been set.

### GetUserIds

`func (o *RegisterQueueResponseRecentPrivateConversationsItem) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *RegisterQueueResponseRecentPrivateConversationsItem) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *RegisterQueueResponseRecentPrivateConversationsItem) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *RegisterQueueResponseRecentPrivateConversationsItem) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


