# RegisterQueueResponseUnreadMsgsPmsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtherUserId** | Pointer to **int32** | The user ID of the other participant in this one-on-one direct message conversation. Will be the current user&#39;s ID for messages that they sent in a one-on-one direct message conversation with themself.  **Changes**: New in Zulip 5.0 (feature level 119), replacing the less clearly named &#x60;sender_id&#x60; field.  | [optional] 
**SenderId** | Pointer to **int32** | Old name for the &#x60;other_user_id&#x60; field. Clients should access this field in Zulip server versions that do not yet support &#x60;other_user_id&#x60;.  **Changes**: Deprecated in Zulip 5.0 (feature level 119). We expect to provide a next version of the full &#x60;unread_msgs&#x60; API before removing this legacy name.  | [optional] 
**UnreadMessageIds** | Pointer to **[]int32** | The message IDs of the recent unread direct messages sent by either user in this one-on-one direct message conversation, sorted in ascending order.  | [optional] 

## Methods

### NewRegisterQueueResponseUnreadMsgsPmsInner

`func NewRegisterQueueResponseUnreadMsgsPmsInner() *RegisterQueueResponseUnreadMsgsPmsInner`

NewRegisterQueueResponseUnreadMsgsPmsInner instantiates a new RegisterQueueResponseUnreadMsgsPmsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseUnreadMsgsPmsInnerWithDefaults

`func NewRegisterQueueResponseUnreadMsgsPmsInnerWithDefaults() *RegisterQueueResponseUnreadMsgsPmsInner`

NewRegisterQueueResponseUnreadMsgsPmsInnerWithDefaults instantiates a new RegisterQueueResponseUnreadMsgsPmsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtherUserId

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) GetOtherUserId() int32`

GetOtherUserId returns the OtherUserId field if non-nil, zero value otherwise.

### GetOtherUserIdOk

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) GetOtherUserIdOk() (*int32, bool)`

GetOtherUserIdOk returns a tuple with the OtherUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherUserId

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) SetOtherUserId(v int32)`

SetOtherUserId sets OtherUserId field to given value.

### HasOtherUserId

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) HasOtherUserId() bool`

HasOtherUserId returns a boolean if a field has been set.

### GetSenderId

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) GetSenderId() int32`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) GetSenderIdOk() (*int32, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) SetSenderId(v int32)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) GetUnreadMessageIds() []int32`

GetUnreadMessageIds returns the UnreadMessageIds field if non-nil, zero value otherwise.

### GetUnreadMessageIdsOk

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) GetUnreadMessageIdsOk() (*[]int32, bool)`

GetUnreadMessageIdsOk returns a tuple with the UnreadMessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) SetUnreadMessageIds(v []int32)`

SetUnreadMessageIds sets UnreadMessageIds field to given value.

### HasUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsPmsInner) HasUnreadMessageIds() bool`

HasUnreadMessageIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


