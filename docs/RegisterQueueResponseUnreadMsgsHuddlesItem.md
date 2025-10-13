# RegisterQueueResponseUnreadMsgsHuddlesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIdsString** | Pointer to **string** | A string containing the IDs of all users in the group direct message conversation, including the current user, separated by commas and sorted numerically; for example: &#x60;\&quot;1,2,3\&quot;&#x60;.  | [optional] 
**UnreadMessageIds** | Pointer to **[]int32** | The message IDs of the recent unread messages which have been sent in this group direct message conversation, sorted in ascending order.  | [optional] 

## Methods

### NewRegisterQueueResponseUnreadMsgsHuddlesInner

`func NewRegisterQueueResponseUnreadMsgsHuddlesInner() *RegisterQueueResponseUnreadMsgsHuddlesInner`

NewRegisterQueueResponseUnreadMsgsHuddlesInner instantiates a new RegisterQueueResponseUnreadMsgsHuddlesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseUnreadMsgsHuddlesInnerWithDefaults

`func NewRegisterQueueResponseUnreadMsgsHuddlesInnerWithDefaults() *RegisterQueueResponseUnreadMsgsHuddlesInner`

NewRegisterQueueResponseUnreadMsgsHuddlesInnerWithDefaults instantiates a new RegisterQueueResponseUnreadMsgsHuddlesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIdsString

`func (o *RegisterQueueResponseUnreadMsgsHuddlesInner) GetUserIdsString() string`

GetUserIdsString returns the UserIdsString field if non-nil, zero value otherwise.

### GetUserIdsStringOk

`func (o *RegisterQueueResponseUnreadMsgsHuddlesInner) GetUserIdsStringOk() (*string, bool)`

GetUserIdsStringOk returns a tuple with the UserIdsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdsString

`func (o *RegisterQueueResponseUnreadMsgsHuddlesInner) SetUserIdsString(v string)`

SetUserIdsString sets UserIdsString field to given value.

### HasUserIdsString

`func (o *RegisterQueueResponseUnreadMsgsHuddlesInner) HasUserIdsString() bool`

HasUserIdsString returns a boolean if a field has been set.

### GetUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsHuddlesInner) GetUnreadMessageIds() []int32`

GetUnreadMessageIds returns the UnreadMessageIds field if non-nil, zero value otherwise.

### GetUnreadMessageIdsOk

`func (o *RegisterQueueResponseUnreadMsgsHuddlesInner) GetUnreadMessageIdsOk() (*[]int32, bool)`

GetUnreadMessageIdsOk returns a tuple with the UnreadMessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsHuddlesInner) SetUnreadMessageIds(v []int32)`

SetUnreadMessageIds sets UnreadMessageIds field to given value.

### HasUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsHuddlesInner) HasUnreadMessageIds() bool`

HasUnreadMessageIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


