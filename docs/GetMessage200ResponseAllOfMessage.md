# GetMessage200ResponseAllOfMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **interface{}** |  | [optional] 
**Client** | Pointer to **interface{}** |  | [optional] 
**Content** | Pointer to **interface{}** |  | [optional] 
**ContentType** | Pointer to **interface{}** |  | [optional] 
**DisplayRecipient** | Pointer to **interface{}** |  | [optional] 
**EditHistory** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**IsMeMessage** | Pointer to **interface{}** |  | [optional] 
**LastEditTimestamp** | Pointer to **interface{}** |  | [optional] 
**LastMovedTimestamp** | Pointer to **interface{}** |  | [optional] 
**Reactions** | Pointer to **interface{}** |  | [optional] 
**RecipientId** | Pointer to **interface{}** |  | [optional] 
**SenderEmail** | Pointer to **interface{}** |  | [optional] 
**SenderFullName** | Pointer to **interface{}** |  | [optional] 
**SenderId** | Pointer to **interface{}** |  | [optional] 
**SenderRealmStr** | Pointer to **interface{}** |  | [optional] 
**StreamId** | Pointer to **interface{}** |  | [optional] 
**Subject** | Pointer to **interface{}** |  | [optional] 
**Submessages** | Pointer to **interface{}** |  | [optional] 
**Timestamp** | Pointer to **interface{}** |  | [optional] 
**TopicLinks** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 
**Flags** | Pointer to **[]string** | The user&#39;s [message flags][message-flags] for the message.  **Changes**: In Zulip 8.0 (feature level 224), the &#x60;wildcard_mentioned&#x60; flag was deprecated in favor of the &#x60;stream_wildcard_mentioned&#x60; and &#x60;topic_wildcard_mentioned&#x60; flags. The &#x60;wildcard_mentioned&#x60; flag exists for backwards compatibility with older clients and equals &#x60;stream_wildcard_mentioned || topic_wildcard_mentioned&#x60;. Clients supporting older server versions should treat this field as a previous name for the &#x60;stream_wildcard_mentioned&#x60; flag as topic wildcard mentions were not available prior to this feature level.  [message-flags]: /api/update-message-flags#available-flags  | [optional] 

## Methods

### NewGetMessage200ResponseAllOfMessage

`func NewGetMessage200ResponseAllOfMessage() *GetMessage200ResponseAllOfMessage`

NewGetMessage200ResponseAllOfMessage instantiates a new GetMessage200ResponseAllOfMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessage200ResponseAllOfMessageWithDefaults

`func NewGetMessage200ResponseAllOfMessageWithDefaults() *GetMessage200ResponseAllOfMessage`

NewGetMessage200ResponseAllOfMessageWithDefaults instantiates a new GetMessage200ResponseAllOfMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *GetMessage200ResponseAllOfMessage) GetAvatarUrl() interface{}`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GetMessage200ResponseAllOfMessage) GetAvatarUrlOk() (*interface{}, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GetMessage200ResponseAllOfMessage) SetAvatarUrl(v interface{})`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GetMessage200ResponseAllOfMessage) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *GetMessage200ResponseAllOfMessage) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *GetMessage200ResponseAllOfMessage) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetClient

`func (o *GetMessage200ResponseAllOfMessage) GetClient() interface{}`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *GetMessage200ResponseAllOfMessage) GetClientOk() (*interface{}, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *GetMessage200ResponseAllOfMessage) SetClient(v interface{})`

SetClient sets Client field to given value.

### HasClient

`func (o *GetMessage200ResponseAllOfMessage) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *GetMessage200ResponseAllOfMessage) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *GetMessage200ResponseAllOfMessage) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetContent

`func (o *GetMessage200ResponseAllOfMessage) GetContent() interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetMessage200ResponseAllOfMessage) GetContentOk() (*interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetMessage200ResponseAllOfMessage) SetContent(v interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *GetMessage200ResponseAllOfMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *GetMessage200ResponseAllOfMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *GetMessage200ResponseAllOfMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentType

`func (o *GetMessage200ResponseAllOfMessage) GetContentType() interface{}`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetMessage200ResponseAllOfMessage) GetContentTypeOk() (*interface{}, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetMessage200ResponseAllOfMessage) SetContentType(v interface{})`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GetMessage200ResponseAllOfMessage) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *GetMessage200ResponseAllOfMessage) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *GetMessage200ResponseAllOfMessage) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDisplayRecipient

`func (o *GetMessage200ResponseAllOfMessage) GetDisplayRecipient() interface{}`

GetDisplayRecipient returns the DisplayRecipient field if non-nil, zero value otherwise.

### GetDisplayRecipientOk

`func (o *GetMessage200ResponseAllOfMessage) GetDisplayRecipientOk() (*interface{}, bool)`

GetDisplayRecipientOk returns a tuple with the DisplayRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayRecipient

`func (o *GetMessage200ResponseAllOfMessage) SetDisplayRecipient(v interface{})`

SetDisplayRecipient sets DisplayRecipient field to given value.

### HasDisplayRecipient

`func (o *GetMessage200ResponseAllOfMessage) HasDisplayRecipient() bool`

HasDisplayRecipient returns a boolean if a field has been set.

### SetDisplayRecipientNil

`func (o *GetMessage200ResponseAllOfMessage) SetDisplayRecipientNil(b bool)`

 SetDisplayRecipientNil sets the value for DisplayRecipient to be an explicit nil

### UnsetDisplayRecipient
`func (o *GetMessage200ResponseAllOfMessage) UnsetDisplayRecipient()`

UnsetDisplayRecipient ensures that no value is present for DisplayRecipient, not even an explicit nil
### GetEditHistory

`func (o *GetMessage200ResponseAllOfMessage) GetEditHistory() interface{}`

GetEditHistory returns the EditHistory field if non-nil, zero value otherwise.

### GetEditHistoryOk

`func (o *GetMessage200ResponseAllOfMessage) GetEditHistoryOk() (*interface{}, bool)`

GetEditHistoryOk returns a tuple with the EditHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditHistory

`func (o *GetMessage200ResponseAllOfMessage) SetEditHistory(v interface{})`

SetEditHistory sets EditHistory field to given value.

### HasEditHistory

`func (o *GetMessage200ResponseAllOfMessage) HasEditHistory() bool`

HasEditHistory returns a boolean if a field has been set.

### SetEditHistoryNil

`func (o *GetMessage200ResponseAllOfMessage) SetEditHistoryNil(b bool)`

 SetEditHistoryNil sets the value for EditHistory to be an explicit nil

### UnsetEditHistory
`func (o *GetMessage200ResponseAllOfMessage) UnsetEditHistory()`

UnsetEditHistory ensures that no value is present for EditHistory, not even an explicit nil
### GetId

`func (o *GetMessage200ResponseAllOfMessage) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMessage200ResponseAllOfMessage) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMessage200ResponseAllOfMessage) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *GetMessage200ResponseAllOfMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GetMessage200ResponseAllOfMessage) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GetMessage200ResponseAllOfMessage) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsMeMessage

`func (o *GetMessage200ResponseAllOfMessage) GetIsMeMessage() interface{}`

GetIsMeMessage returns the IsMeMessage field if non-nil, zero value otherwise.

### GetIsMeMessageOk

`func (o *GetMessage200ResponseAllOfMessage) GetIsMeMessageOk() (*interface{}, bool)`

GetIsMeMessageOk returns a tuple with the IsMeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeMessage

`func (o *GetMessage200ResponseAllOfMessage) SetIsMeMessage(v interface{})`

SetIsMeMessage sets IsMeMessage field to given value.

### HasIsMeMessage

`func (o *GetMessage200ResponseAllOfMessage) HasIsMeMessage() bool`

HasIsMeMessage returns a boolean if a field has been set.

### SetIsMeMessageNil

`func (o *GetMessage200ResponseAllOfMessage) SetIsMeMessageNil(b bool)`

 SetIsMeMessageNil sets the value for IsMeMessage to be an explicit nil

### UnsetIsMeMessage
`func (o *GetMessage200ResponseAllOfMessage) UnsetIsMeMessage()`

UnsetIsMeMessage ensures that no value is present for IsMeMessage, not even an explicit nil
### GetLastEditTimestamp

`func (o *GetMessage200ResponseAllOfMessage) GetLastEditTimestamp() interface{}`

GetLastEditTimestamp returns the LastEditTimestamp field if non-nil, zero value otherwise.

### GetLastEditTimestampOk

`func (o *GetMessage200ResponseAllOfMessage) GetLastEditTimestampOk() (*interface{}, bool)`

GetLastEditTimestampOk returns a tuple with the LastEditTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditTimestamp

`func (o *GetMessage200ResponseAllOfMessage) SetLastEditTimestamp(v interface{})`

SetLastEditTimestamp sets LastEditTimestamp field to given value.

### HasLastEditTimestamp

`func (o *GetMessage200ResponseAllOfMessage) HasLastEditTimestamp() bool`

HasLastEditTimestamp returns a boolean if a field has been set.

### SetLastEditTimestampNil

`func (o *GetMessage200ResponseAllOfMessage) SetLastEditTimestampNil(b bool)`

 SetLastEditTimestampNil sets the value for LastEditTimestamp to be an explicit nil

### UnsetLastEditTimestamp
`func (o *GetMessage200ResponseAllOfMessage) UnsetLastEditTimestamp()`

UnsetLastEditTimestamp ensures that no value is present for LastEditTimestamp, not even an explicit nil
### GetLastMovedTimestamp

`func (o *GetMessage200ResponseAllOfMessage) GetLastMovedTimestamp() interface{}`

GetLastMovedTimestamp returns the LastMovedTimestamp field if non-nil, zero value otherwise.

### GetLastMovedTimestampOk

`func (o *GetMessage200ResponseAllOfMessage) GetLastMovedTimestampOk() (*interface{}, bool)`

GetLastMovedTimestampOk returns a tuple with the LastMovedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMovedTimestamp

`func (o *GetMessage200ResponseAllOfMessage) SetLastMovedTimestamp(v interface{})`

SetLastMovedTimestamp sets LastMovedTimestamp field to given value.

### HasLastMovedTimestamp

`func (o *GetMessage200ResponseAllOfMessage) HasLastMovedTimestamp() bool`

HasLastMovedTimestamp returns a boolean if a field has been set.

### SetLastMovedTimestampNil

`func (o *GetMessage200ResponseAllOfMessage) SetLastMovedTimestampNil(b bool)`

 SetLastMovedTimestampNil sets the value for LastMovedTimestamp to be an explicit nil

### UnsetLastMovedTimestamp
`func (o *GetMessage200ResponseAllOfMessage) UnsetLastMovedTimestamp()`

UnsetLastMovedTimestamp ensures that no value is present for LastMovedTimestamp, not even an explicit nil
### GetReactions

`func (o *GetMessage200ResponseAllOfMessage) GetReactions() interface{}`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *GetMessage200ResponseAllOfMessage) GetReactionsOk() (*interface{}, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *GetMessage200ResponseAllOfMessage) SetReactions(v interface{})`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *GetMessage200ResponseAllOfMessage) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### SetReactionsNil

`func (o *GetMessage200ResponseAllOfMessage) SetReactionsNil(b bool)`

 SetReactionsNil sets the value for Reactions to be an explicit nil

### UnsetReactions
`func (o *GetMessage200ResponseAllOfMessage) UnsetReactions()`

UnsetReactions ensures that no value is present for Reactions, not even an explicit nil
### GetRecipientId

`func (o *GetMessage200ResponseAllOfMessage) GetRecipientId() interface{}`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *GetMessage200ResponseAllOfMessage) GetRecipientIdOk() (*interface{}, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *GetMessage200ResponseAllOfMessage) SetRecipientId(v interface{})`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *GetMessage200ResponseAllOfMessage) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### SetRecipientIdNil

`func (o *GetMessage200ResponseAllOfMessage) SetRecipientIdNil(b bool)`

 SetRecipientIdNil sets the value for RecipientId to be an explicit nil

### UnsetRecipientId
`func (o *GetMessage200ResponseAllOfMessage) UnsetRecipientId()`

UnsetRecipientId ensures that no value is present for RecipientId, not even an explicit nil
### GetSenderEmail

`func (o *GetMessage200ResponseAllOfMessage) GetSenderEmail() interface{}`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *GetMessage200ResponseAllOfMessage) GetSenderEmailOk() (*interface{}, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *GetMessage200ResponseAllOfMessage) SetSenderEmail(v interface{})`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *GetMessage200ResponseAllOfMessage) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### SetSenderEmailNil

`func (o *GetMessage200ResponseAllOfMessage) SetSenderEmailNil(b bool)`

 SetSenderEmailNil sets the value for SenderEmail to be an explicit nil

### UnsetSenderEmail
`func (o *GetMessage200ResponseAllOfMessage) UnsetSenderEmail()`

UnsetSenderEmail ensures that no value is present for SenderEmail, not even an explicit nil
### GetSenderFullName

`func (o *GetMessage200ResponseAllOfMessage) GetSenderFullName() interface{}`

GetSenderFullName returns the SenderFullName field if non-nil, zero value otherwise.

### GetSenderFullNameOk

`func (o *GetMessage200ResponseAllOfMessage) GetSenderFullNameOk() (*interface{}, bool)`

GetSenderFullNameOk returns a tuple with the SenderFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderFullName

`func (o *GetMessage200ResponseAllOfMessage) SetSenderFullName(v interface{})`

SetSenderFullName sets SenderFullName field to given value.

### HasSenderFullName

`func (o *GetMessage200ResponseAllOfMessage) HasSenderFullName() bool`

HasSenderFullName returns a boolean if a field has been set.

### SetSenderFullNameNil

`func (o *GetMessage200ResponseAllOfMessage) SetSenderFullNameNil(b bool)`

 SetSenderFullNameNil sets the value for SenderFullName to be an explicit nil

### UnsetSenderFullName
`func (o *GetMessage200ResponseAllOfMessage) UnsetSenderFullName()`

UnsetSenderFullName ensures that no value is present for SenderFullName, not even an explicit nil
### GetSenderId

`func (o *GetMessage200ResponseAllOfMessage) GetSenderId() interface{}`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *GetMessage200ResponseAllOfMessage) GetSenderIdOk() (*interface{}, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *GetMessage200ResponseAllOfMessage) SetSenderId(v interface{})`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *GetMessage200ResponseAllOfMessage) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### SetSenderIdNil

`func (o *GetMessage200ResponseAllOfMessage) SetSenderIdNil(b bool)`

 SetSenderIdNil sets the value for SenderId to be an explicit nil

### UnsetSenderId
`func (o *GetMessage200ResponseAllOfMessage) UnsetSenderId()`

UnsetSenderId ensures that no value is present for SenderId, not even an explicit nil
### GetSenderRealmStr

`func (o *GetMessage200ResponseAllOfMessage) GetSenderRealmStr() interface{}`

GetSenderRealmStr returns the SenderRealmStr field if non-nil, zero value otherwise.

### GetSenderRealmStrOk

`func (o *GetMessage200ResponseAllOfMessage) GetSenderRealmStrOk() (*interface{}, bool)`

GetSenderRealmStrOk returns a tuple with the SenderRealmStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderRealmStr

`func (o *GetMessage200ResponseAllOfMessage) SetSenderRealmStr(v interface{})`

SetSenderRealmStr sets SenderRealmStr field to given value.

### HasSenderRealmStr

`func (o *GetMessage200ResponseAllOfMessage) HasSenderRealmStr() bool`

HasSenderRealmStr returns a boolean if a field has been set.

### SetSenderRealmStrNil

`func (o *GetMessage200ResponseAllOfMessage) SetSenderRealmStrNil(b bool)`

 SetSenderRealmStrNil sets the value for SenderRealmStr to be an explicit nil

### UnsetSenderRealmStr
`func (o *GetMessage200ResponseAllOfMessage) UnsetSenderRealmStr()`

UnsetSenderRealmStr ensures that no value is present for SenderRealmStr, not even an explicit nil
### GetStreamId

`func (o *GetMessage200ResponseAllOfMessage) GetStreamId() interface{}`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *GetMessage200ResponseAllOfMessage) GetStreamIdOk() (*interface{}, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *GetMessage200ResponseAllOfMessage) SetStreamId(v interface{})`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *GetMessage200ResponseAllOfMessage) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### SetStreamIdNil

`func (o *GetMessage200ResponseAllOfMessage) SetStreamIdNil(b bool)`

 SetStreamIdNil sets the value for StreamId to be an explicit nil

### UnsetStreamId
`func (o *GetMessage200ResponseAllOfMessage) UnsetStreamId()`

UnsetStreamId ensures that no value is present for StreamId, not even an explicit nil
### GetSubject

`func (o *GetMessage200ResponseAllOfMessage) GetSubject() interface{}`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetMessage200ResponseAllOfMessage) GetSubjectOk() (*interface{}, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetMessage200ResponseAllOfMessage) SetSubject(v interface{})`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetMessage200ResponseAllOfMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *GetMessage200ResponseAllOfMessage) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *GetMessage200ResponseAllOfMessage) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetSubmessages

`func (o *GetMessage200ResponseAllOfMessage) GetSubmessages() interface{}`

GetSubmessages returns the Submessages field if non-nil, zero value otherwise.

### GetSubmessagesOk

`func (o *GetMessage200ResponseAllOfMessage) GetSubmessagesOk() (*interface{}, bool)`

GetSubmessagesOk returns a tuple with the Submessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmessages

`func (o *GetMessage200ResponseAllOfMessage) SetSubmessages(v interface{})`

SetSubmessages sets Submessages field to given value.

### HasSubmessages

`func (o *GetMessage200ResponseAllOfMessage) HasSubmessages() bool`

HasSubmessages returns a boolean if a field has been set.

### SetSubmessagesNil

`func (o *GetMessage200ResponseAllOfMessage) SetSubmessagesNil(b bool)`

 SetSubmessagesNil sets the value for Submessages to be an explicit nil

### UnsetSubmessages
`func (o *GetMessage200ResponseAllOfMessage) UnsetSubmessages()`

UnsetSubmessages ensures that no value is present for Submessages, not even an explicit nil
### GetTimestamp

`func (o *GetMessage200ResponseAllOfMessage) GetTimestamp() interface{}`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetMessage200ResponseAllOfMessage) GetTimestampOk() (*interface{}, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetMessage200ResponseAllOfMessage) SetTimestamp(v interface{})`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetMessage200ResponseAllOfMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *GetMessage200ResponseAllOfMessage) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *GetMessage200ResponseAllOfMessage) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTopicLinks

`func (o *GetMessage200ResponseAllOfMessage) GetTopicLinks() interface{}`

GetTopicLinks returns the TopicLinks field if non-nil, zero value otherwise.

### GetTopicLinksOk

`func (o *GetMessage200ResponseAllOfMessage) GetTopicLinksOk() (*interface{}, bool)`

GetTopicLinksOk returns a tuple with the TopicLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicLinks

`func (o *GetMessage200ResponseAllOfMessage) SetTopicLinks(v interface{})`

SetTopicLinks sets TopicLinks field to given value.

### HasTopicLinks

`func (o *GetMessage200ResponseAllOfMessage) HasTopicLinks() bool`

HasTopicLinks returns a boolean if a field has been set.

### SetTopicLinksNil

`func (o *GetMessage200ResponseAllOfMessage) SetTopicLinksNil(b bool)`

 SetTopicLinksNil sets the value for TopicLinks to be an explicit nil

### UnsetTopicLinks
`func (o *GetMessage200ResponseAllOfMessage) UnsetTopicLinks()`

UnsetTopicLinks ensures that no value is present for TopicLinks, not even an explicit nil
### GetType

`func (o *GetMessage200ResponseAllOfMessage) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMessage200ResponseAllOfMessage) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMessage200ResponseAllOfMessage) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GetMessage200ResponseAllOfMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GetMessage200ResponseAllOfMessage) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GetMessage200ResponseAllOfMessage) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFlags

`func (o *GetMessage200ResponseAllOfMessage) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetMessage200ResponseAllOfMessage) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetMessage200ResponseAllOfMessage) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *GetMessage200ResponseAllOfMessage) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


