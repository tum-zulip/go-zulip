# MessagesBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **NullableString** | The URL of the message sender&#39;s avatar. Can be &#x60;null&#x60; only if the current user has access to the sender&#39;s real email address and &#x60;client_gravatar&#x60; was &#x60;true&#x60;.  If &#x60;null&#x60;, then the sender has not uploaded an avatar in Zulip, and the client can compute the gravatar URL by hashing the sender&#39;s email address, which corresponds in this case to their real email address.  **Changes**: Before Zulip 7.0 (feature level 163), access to a user&#39;s real email address was a realm-level setting. As of this feature level, &#x60;email_address_visibility&#x60; is a user setting.  | [optional] 
**Client** | Pointer to **string** | A Zulip \&quot;client\&quot; string, describing what Zulip client sent the message.  | [optional] 
**Content** | Pointer to **string** | The content/body of the message. When &#x60;apply_markdown&#x60; is set, it will be in HTML format.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**ContentType** | Pointer to **string** | The HTTP &#x60;content_type&#x60; for the message content. This will be &#x60;text/html&#x60; or &#x60;text/x-markdown&#x60;, depending on whether &#x60;apply_markdown&#x60; was set.  See the help center article on [message formatting](/help/format-your-message-using-markdown) for details on Zulip-flavored Markdown.  | [optional] 
**DisplayRecipient** | Pointer to [**MessagesBaseDisplayRecipient**](MessagesBaseDisplayRecipient.md) |  | [optional] 
**EditHistory** | Pointer to [**[]MessagesBaseEditHistoryInner**](MessagesBaseEditHistoryInner.md) | An array of objects, with each object documenting the changes in a previous edit made to the message, ordered chronologically from most recent to least recent edit.  Not present if the message has never been edited or moved, or if [viewing message edit history][edit-history-access] is not allowed in the organization.  Every object will contain &#x60;user_id&#x60; and &#x60;timestamp&#x60;.  The other fields are optional, and will be present or not depending on whether the channel, topic, and/or message content were modified in the edit event. For example, if only the topic was edited, only &#x60;prev_topic&#x60; and &#x60;topic&#x60; will be present in addition to &#x60;user_id&#x60; and &#x60;timestamp&#x60;.  [edit-history-access]: /help/restrict-message-edit-history-access  **Changes**: In Zulip 10.0 (feature level 284), removed the &#x60;prev_rendered_content_version&#x60; field as it is an internal server implementation detail not used by any client.  | [optional] 
**Id** | Pointer to **int32** | The unique message ID. Messages should always be displayed sorted by ID.  | [optional] 
**IsMeMessage** | Pointer to **bool** | Whether the message is a [/me status message][status-messages]  [status-messages]: /help/format-your-message-using-markdown#status-messages  | [optional] 
**LastEditTimestamp** | Pointer to **int32** | The UNIX timestamp for when the message&#39;s content was last edited, in UTC seconds.  Not present if the message&#39;s content has never been edited.  Clients should use this field, rather than parsing the &#x60;edit_history&#x60; array, to display an indicator that the message has been edited.  **Changes**: Prior to Zulip 10.0 (feature level 365), this was the time when the message was last edited or moved.  | [optional] 
**LastMovedTimestamp** | Pointer to **int32** | The UNIX timestamp for when the message was last moved to a different channel or topic, in UTC seconds.  Not present if the message has never been moved, or if the only topic moves for the message are [resolving or unresolving](/help/resolve-a-topic) the message&#39;s topic.  Clients should use this field, rather than parsing the &#x60;edit_history&#x60; array, to display an indicator that the message has been moved.  **Changes**: New in Zulip 10.0 (feature level 365). Previously, parsing the &#x60;edit_history&#x60; array was required in order to correctly display moved message indicators.  | [optional] 
**Reactions** | Pointer to [**[]EmojiReaction**](EmojiReaction.md) | Data on any reactions to the message.  | [optional] 
**RecipientId** | Pointer to **int32** | A unique ID for the set of users receiving the message (either a channel or group of users). Useful primarily for hashing.  **Changes**: Before Zulip 10.0 (feature level 327), &#x60;recipient_id&#x60; was the same across all incoming 1:1 direct messages. Now, each incoming message uniquely shares a &#x60;recipient_id&#x60; with outgoing messages in the same conversation.  | [optional] 
**SenderEmail** | Pointer to **string** | The Zulip API email address of the message&#39;s sender.  | [optional] 
**SenderFullName** | Pointer to **string** | The full name of the message&#39;s sender.  | [optional] 
**SenderId** | Pointer to **int32** | The user ID of the message&#39;s sender.  | [optional] 
**SenderRealmStr** | Pointer to **string** | A string identifier for the realm the sender is in. Unique only within the context of a given Zulip server.  E.g. on &#x60;example.zulip.com&#x60;, this will be &#x60;example&#x60;.  | [optional] 
**StreamId** | Pointer to **int32** | Only present for channel messages; the ID of the channel.  | [optional] 
**Subject** | Pointer to **string** | The &#x60;topic&#x60; of the message. Currently always &#x60;\&quot;\&quot;&#x60; for direct messages, though this could change if Zulip adds support for topics in direct message conversations.  The field name is a legacy holdover from when topics were called \&quot;subjects\&quot; and will eventually change.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], the empty string value is replaced with the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](/api/register-queue) response, for channel messages.  **Changes**: Before Zulip 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**Submessages** | Pointer to [**[]MessagesBaseSubmessagesInner**](MessagesBaseSubmessagesInner.md) | Data used for certain experimental Zulip integrations.  | [optional] 
**Timestamp** | Pointer to **int32** | The UNIX timestamp for when the message was sent, in UTC seconds.  | [optional] 
**TopicLinks** | Pointer to [**[]GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner**](GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner.md) | Data on any links to be included in the &#x60;topic&#x60; line (these are generated by [custom linkification filters](/help/add-a-custom-linkifier) that match content in the message&#39;s topic.)  **Changes**: This field contained a list of urls before Zulip 4.0 (feature level 46).  New in Zulip 3.0 (feature level 1). Previously, this field was called &#x60;subject_links&#x60;; clients are recommended to rename &#x60;subject_links&#x60; to &#x60;topic_links&#x60; if present for compatibility with older Zulip servers.  | [optional] 
**Type** | Pointer to **string** | The type of the message: &#x60;\&quot;stream\&quot;&#x60; or &#x60;\&quot;private\&quot;&#x60;.  | [optional] 

## Methods

### NewMessagesBase

`func NewMessagesBase() *MessagesBase`

NewMessagesBase instantiates a new MessagesBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesBaseWithDefaults

`func NewMessagesBaseWithDefaults() *MessagesBase`

NewMessagesBaseWithDefaults instantiates a new MessagesBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *MessagesBase) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *MessagesBase) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *MessagesBase) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *MessagesBase) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *MessagesBase) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *MessagesBase) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetClient

`func (o *MessagesBase) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *MessagesBase) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *MessagesBase) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *MessagesBase) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetContent

`func (o *MessagesBase) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessagesBase) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessagesBase) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MessagesBase) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentType

`func (o *MessagesBase) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MessagesBase) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MessagesBase) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MessagesBase) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDisplayRecipient

`func (o *MessagesBase) GetDisplayRecipient() MessagesBaseDisplayRecipient`

GetDisplayRecipient returns the DisplayRecipient field if non-nil, zero value otherwise.

### GetDisplayRecipientOk

`func (o *MessagesBase) GetDisplayRecipientOk() (*MessagesBaseDisplayRecipient, bool)`

GetDisplayRecipientOk returns a tuple with the DisplayRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayRecipient

`func (o *MessagesBase) SetDisplayRecipient(v MessagesBaseDisplayRecipient)`

SetDisplayRecipient sets DisplayRecipient field to given value.

### HasDisplayRecipient

`func (o *MessagesBase) HasDisplayRecipient() bool`

HasDisplayRecipient returns a boolean if a field has been set.

### GetEditHistory

`func (o *MessagesBase) GetEditHistory() []MessagesBaseEditHistoryInner`

GetEditHistory returns the EditHistory field if non-nil, zero value otherwise.

### GetEditHistoryOk

`func (o *MessagesBase) GetEditHistoryOk() (*[]MessagesBaseEditHistoryInner, bool)`

GetEditHistoryOk returns a tuple with the EditHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditHistory

`func (o *MessagesBase) SetEditHistory(v []MessagesBaseEditHistoryInner)`

SetEditHistory sets EditHistory field to given value.

### HasEditHistory

`func (o *MessagesBase) HasEditHistory() bool`

HasEditHistory returns a boolean if a field has been set.

### GetId

`func (o *MessagesBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessagesBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessagesBase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MessagesBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMeMessage

`func (o *MessagesBase) GetIsMeMessage() bool`

GetIsMeMessage returns the IsMeMessage field if non-nil, zero value otherwise.

### GetIsMeMessageOk

`func (o *MessagesBase) GetIsMeMessageOk() (*bool, bool)`

GetIsMeMessageOk returns a tuple with the IsMeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeMessage

`func (o *MessagesBase) SetIsMeMessage(v bool)`

SetIsMeMessage sets IsMeMessage field to given value.

### HasIsMeMessage

`func (o *MessagesBase) HasIsMeMessage() bool`

HasIsMeMessage returns a boolean if a field has been set.

### GetLastEditTimestamp

`func (o *MessagesBase) GetLastEditTimestamp() int32`

GetLastEditTimestamp returns the LastEditTimestamp field if non-nil, zero value otherwise.

### GetLastEditTimestampOk

`func (o *MessagesBase) GetLastEditTimestampOk() (*int32, bool)`

GetLastEditTimestampOk returns a tuple with the LastEditTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditTimestamp

`func (o *MessagesBase) SetLastEditTimestamp(v int32)`

SetLastEditTimestamp sets LastEditTimestamp field to given value.

### HasLastEditTimestamp

`func (o *MessagesBase) HasLastEditTimestamp() bool`

HasLastEditTimestamp returns a boolean if a field has been set.

### GetLastMovedTimestamp

`func (o *MessagesBase) GetLastMovedTimestamp() int32`

GetLastMovedTimestamp returns the LastMovedTimestamp field if non-nil, zero value otherwise.

### GetLastMovedTimestampOk

`func (o *MessagesBase) GetLastMovedTimestampOk() (*int32, bool)`

GetLastMovedTimestampOk returns a tuple with the LastMovedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMovedTimestamp

`func (o *MessagesBase) SetLastMovedTimestamp(v int32)`

SetLastMovedTimestamp sets LastMovedTimestamp field to given value.

### HasLastMovedTimestamp

`func (o *MessagesBase) HasLastMovedTimestamp() bool`

HasLastMovedTimestamp returns a boolean if a field has been set.

### GetReactions

`func (o *MessagesBase) GetReactions() []EmojiReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *MessagesBase) GetReactionsOk() (*[]EmojiReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *MessagesBase) SetReactions(v []EmojiReaction)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *MessagesBase) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetRecipientId

`func (o *MessagesBase) GetRecipientId() int32`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *MessagesBase) GetRecipientIdOk() (*int32, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *MessagesBase) SetRecipientId(v int32)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *MessagesBase) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetSenderEmail

`func (o *MessagesBase) GetSenderEmail() string`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *MessagesBase) GetSenderEmailOk() (*string, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *MessagesBase) SetSenderEmail(v string)`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *MessagesBase) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### GetSenderFullName

`func (o *MessagesBase) GetSenderFullName() string`

GetSenderFullName returns the SenderFullName field if non-nil, zero value otherwise.

### GetSenderFullNameOk

`func (o *MessagesBase) GetSenderFullNameOk() (*string, bool)`

GetSenderFullNameOk returns a tuple with the SenderFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderFullName

`func (o *MessagesBase) SetSenderFullName(v string)`

SetSenderFullName sets SenderFullName field to given value.

### HasSenderFullName

`func (o *MessagesBase) HasSenderFullName() bool`

HasSenderFullName returns a boolean if a field has been set.

### GetSenderId

`func (o *MessagesBase) GetSenderId() int32`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *MessagesBase) GetSenderIdOk() (*int32, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *MessagesBase) SetSenderId(v int32)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *MessagesBase) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetSenderRealmStr

`func (o *MessagesBase) GetSenderRealmStr() string`

GetSenderRealmStr returns the SenderRealmStr field if non-nil, zero value otherwise.

### GetSenderRealmStrOk

`func (o *MessagesBase) GetSenderRealmStrOk() (*string, bool)`

GetSenderRealmStrOk returns a tuple with the SenderRealmStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderRealmStr

`func (o *MessagesBase) SetSenderRealmStr(v string)`

SetSenderRealmStr sets SenderRealmStr field to given value.

### HasSenderRealmStr

`func (o *MessagesBase) HasSenderRealmStr() bool`

HasSenderRealmStr returns a boolean if a field has been set.

### GetStreamId

`func (o *MessagesBase) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *MessagesBase) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *MessagesBase) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *MessagesBase) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetSubject

`func (o *MessagesBase) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MessagesBase) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MessagesBase) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MessagesBase) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubmessages

`func (o *MessagesBase) GetSubmessages() []MessagesBaseSubmessagesInner`

GetSubmessages returns the Submessages field if non-nil, zero value otherwise.

### GetSubmessagesOk

`func (o *MessagesBase) GetSubmessagesOk() (*[]MessagesBaseSubmessagesInner, bool)`

GetSubmessagesOk returns a tuple with the Submessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmessages

`func (o *MessagesBase) SetSubmessages(v []MessagesBaseSubmessagesInner)`

SetSubmessages sets Submessages field to given value.

### HasSubmessages

`func (o *MessagesBase) HasSubmessages() bool`

HasSubmessages returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessagesBase) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessagesBase) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessagesBase) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessagesBase) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTopicLinks

`func (o *MessagesBase) GetTopicLinks() []GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner`

GetTopicLinks returns the TopicLinks field if non-nil, zero value otherwise.

### GetTopicLinksOk

`func (o *MessagesBase) GetTopicLinksOk() (*[]GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner, bool)`

GetTopicLinksOk returns a tuple with the TopicLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicLinks

`func (o *MessagesBase) SetTopicLinks(v []GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner)`

SetTopicLinks sets TopicLinks field to given value.

### HasTopicLinks

`func (o *MessagesBase) HasTopicLinks() bool`

HasTopicLinks returns a boolean if a field has been set.

### GetType

`func (o *MessagesBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessagesBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessagesBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MessagesBase) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


