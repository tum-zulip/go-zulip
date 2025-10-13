# RegisterQueueResponseUnreadMsgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of unread messages to display. This includes one-on-one and group direct messages, as well as channel messages that are not [muted](zulip.com/help/mute-a-topic.  **Changes**: Before Zulip 8.0 (feature level 213), the unmute and follow topic features were not handled correctly in calculating this field.  | [optional] 
**Pms** | Pointer to [**[]RegisterQueueResponseUnreadMsgsPmsInner**](RegisterQueueResponseUnreadMsgsPmsInner.md) | An array of objects where each object contains details of unread one-on-one direct messages with a specific user.  Note that it is possible for a message that the current user sent to the specified user to be marked as unread and thus appear here.  | [optional] 
**Streams** | Pointer to [**[]RegisterQueueResponseUnreadMsgsStreamsInner**](RegisterQueueResponseUnreadMsgsStreamsInner.md) | An array of dictionaries where each dictionary contains details of all unread messages of a single subscribed channel. This includes muted channels and muted topics, even though those messages are excluded from &#x60;count&#x60;.  **Changes**: Prior to Zulip 5.0 (feature level 90), these objects included a &#x60;sender_ids&#x60; property, which listed the set of IDs of users who had sent the unread messages.  | [optional] 
**Huddles** | Pointer to [**[]RegisterQueueResponseUnreadMsgsHuddlesInner**](RegisterQueueResponseUnreadMsgsHuddlesInner.md) | An array of objects where each object contains details of unread group direct messages with a specific group of users.  | [optional] 
**Mentions** | Pointer to **[]int32** | Array containing the IDs of all unread messages in which the user was mentioned directly, and unread [non-muted](zulip.com/help/mute-a-topic messages in which the user was mentioned through a wildcard.  **Changes**: Before Zulip 8.0 (feature level 213), the unmute and follow topic features were not handled correctly in calculating this field.  | [optional] 
**OldUnreadsMissing** | Pointer to **bool** | Whether this data set was truncated because the user has too many unread messages. When truncation occurs, only the most recent &#x60;MAX_UNREAD_MESSAGES&#x60; (currently 50000) messages will be considered when forming this response. When &#x60;true&#x60;, we recommend that clients display a warning, as they are likely to produce erroneous results until reloaded with the user having fewer than &#x60;MAX_UNREAD_MESSAGES&#x60; unread messages.  **Changes**: New in Zulip 4.0 (feature level 44).  | [optional] 

## Methods

### NewRegisterQueueResponseUnreadMsgs

`func NewRegisterQueueResponseUnreadMsgs() *RegisterQueueResponseUnreadMsgs`

NewRegisterQueueResponseUnreadMsgs instantiates a new RegisterQueueResponseUnreadMsgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseUnreadMsgsWithDefaults

`func NewRegisterQueueResponseUnreadMsgsWithDefaults() *RegisterQueueResponseUnreadMsgs`

NewRegisterQueueResponseUnreadMsgsWithDefaults instantiates a new RegisterQueueResponseUnreadMsgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RegisterQueueResponseUnreadMsgs) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RegisterQueueResponseUnreadMsgs) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RegisterQueueResponseUnreadMsgs) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RegisterQueueResponseUnreadMsgs) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPms

`func (o *RegisterQueueResponseUnreadMsgs) GetPms() []RegisterQueueResponseUnreadMsgsPmsInner`

GetPms returns the Pms field if non-nil, zero value otherwise.

### GetPmsOk

`func (o *RegisterQueueResponseUnreadMsgs) GetPmsOk() (*[]RegisterQueueResponseUnreadMsgsPmsInner, bool)`

GetPmsOk returns a tuple with the Pms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPms

`func (o *RegisterQueueResponseUnreadMsgs) SetPms(v []RegisterQueueResponseUnreadMsgsPmsInner)`

SetPms sets Pms field to given value.

### HasPms

`func (o *RegisterQueueResponseUnreadMsgs) HasPms() bool`

HasPms returns a boolean if a field has been set.

### GetStreams

`func (o *RegisterQueueResponseUnreadMsgs) GetStreams() []RegisterQueueResponseUnreadMsgsStreamsInner`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *RegisterQueueResponseUnreadMsgs) GetStreamsOk() (*[]RegisterQueueResponseUnreadMsgsStreamsInner, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *RegisterQueueResponseUnreadMsgs) SetStreams(v []RegisterQueueResponseUnreadMsgsStreamsInner)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *RegisterQueueResponseUnreadMsgs) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetHuddles

`func (o *RegisterQueueResponseUnreadMsgs) GetHuddles() []RegisterQueueResponseUnreadMsgsHuddlesInner`

GetHuddles returns the Huddles field if non-nil, zero value otherwise.

### GetHuddlesOk

`func (o *RegisterQueueResponseUnreadMsgs) GetHuddlesOk() (*[]RegisterQueueResponseUnreadMsgsHuddlesInner, bool)`

GetHuddlesOk returns a tuple with the Huddles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuddles

`func (o *RegisterQueueResponseUnreadMsgs) SetHuddles(v []RegisterQueueResponseUnreadMsgsHuddlesInner)`

SetHuddles sets Huddles field to given value.

### HasHuddles

`func (o *RegisterQueueResponseUnreadMsgs) HasHuddles() bool`

HasHuddles returns a boolean if a field has been set.

### GetMentions

`func (o *RegisterQueueResponseUnreadMsgs) GetMentions() []int32`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *RegisterQueueResponseUnreadMsgs) GetMentionsOk() (*[]int32, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *RegisterQueueResponseUnreadMsgs) SetMentions(v []int32)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *RegisterQueueResponseUnreadMsgs) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetOldUnreadsMissing

`func (o *RegisterQueueResponseUnreadMsgs) GetOldUnreadsMissing() bool`

GetOldUnreadsMissing returns the OldUnreadsMissing field if non-nil, zero value otherwise.

### GetOldUnreadsMissingOk

`func (o *RegisterQueueResponseUnreadMsgs) GetOldUnreadsMissingOk() (*bool, bool)`

GetOldUnreadsMissingOk returns a tuple with the OldUnreadsMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldUnreadsMissing

`func (o *RegisterQueueResponseUnreadMsgs) SetOldUnreadsMissing(v bool)`

SetOldUnreadsMissing sets OldUnreadsMissing field to given value.

### HasOldUnreadsMissing

`func (o *RegisterQueueResponseUnreadMsgs) HasOldUnreadsMissing() bool`

HasOldUnreadsMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


