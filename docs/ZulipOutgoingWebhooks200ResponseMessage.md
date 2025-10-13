# ZulipOutgoingWebhooks200ResponseMessage

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
**RenderedContent** | Pointer to **string** | The content/body of the message rendered in HTML.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 

## Methods

### NewZulipOutgoingWebhooks200ResponseMessage

`func NewZulipOutgoingWebhooks200ResponseMessage() *ZulipOutgoingWebhooks200ResponseMessage`

NewZulipOutgoingWebhooks200ResponseMessage instantiates a new ZulipOutgoingWebhooks200ResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZulipOutgoingWebhooks200ResponseMessageWithDefaults

`func NewZulipOutgoingWebhooks200ResponseMessageWithDefaults() *ZulipOutgoingWebhooks200ResponseMessage`

NewZulipOutgoingWebhooks200ResponseMessageWithDefaults instantiates a new ZulipOutgoingWebhooks200ResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetAvatarUrl() interface{}`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetAvatarUrlOk() (*interface{}, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetAvatarUrl(v interface{})`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetClient

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetClient() interface{}`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetClientOk() (*interface{}, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetClient(v interface{})`

SetClient sets Client field to given value.

### HasClient

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetContent

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetContent() interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetContentOk() (*interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetContent(v interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentType

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetContentType() interface{}`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetContentTypeOk() (*interface{}, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetContentType(v interface{})`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDisplayRecipient

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetDisplayRecipient() interface{}`

GetDisplayRecipient returns the DisplayRecipient field if non-nil, zero value otherwise.

### GetDisplayRecipientOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetDisplayRecipientOk() (*interface{}, bool)`

GetDisplayRecipientOk returns a tuple with the DisplayRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayRecipient

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetDisplayRecipient(v interface{})`

SetDisplayRecipient sets DisplayRecipient field to given value.

### HasDisplayRecipient

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasDisplayRecipient() bool`

HasDisplayRecipient returns a boolean if a field has been set.

### SetDisplayRecipientNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetDisplayRecipientNil(b bool)`

 SetDisplayRecipientNil sets the value for DisplayRecipient to be an explicit nil

### UnsetDisplayRecipient
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetDisplayRecipient()`

UnsetDisplayRecipient ensures that no value is present for DisplayRecipient, not even an explicit nil
### GetEditHistory

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetEditHistory() interface{}`

GetEditHistory returns the EditHistory field if non-nil, zero value otherwise.

### GetEditHistoryOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetEditHistoryOk() (*interface{}, bool)`

GetEditHistoryOk returns a tuple with the EditHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditHistory

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetEditHistory(v interface{})`

SetEditHistory sets EditHistory field to given value.

### HasEditHistory

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasEditHistory() bool`

HasEditHistory returns a boolean if a field has been set.

### SetEditHistoryNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetEditHistoryNil(b bool)`

 SetEditHistoryNil sets the value for EditHistory to be an explicit nil

### UnsetEditHistory
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetEditHistory()`

UnsetEditHistory ensures that no value is present for EditHistory, not even an explicit nil
### GetId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsMeMessage

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetIsMeMessage() interface{}`

GetIsMeMessage returns the IsMeMessage field if non-nil, zero value otherwise.

### GetIsMeMessageOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetIsMeMessageOk() (*interface{}, bool)`

GetIsMeMessageOk returns a tuple with the IsMeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeMessage

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetIsMeMessage(v interface{})`

SetIsMeMessage sets IsMeMessage field to given value.

### HasIsMeMessage

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasIsMeMessage() bool`

HasIsMeMessage returns a boolean if a field has been set.

### SetIsMeMessageNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetIsMeMessageNil(b bool)`

 SetIsMeMessageNil sets the value for IsMeMessage to be an explicit nil

### UnsetIsMeMessage
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetIsMeMessage()`

UnsetIsMeMessage ensures that no value is present for IsMeMessage, not even an explicit nil
### GetLastEditTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetLastEditTimestamp() interface{}`

GetLastEditTimestamp returns the LastEditTimestamp field if non-nil, zero value otherwise.

### GetLastEditTimestampOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetLastEditTimestampOk() (*interface{}, bool)`

GetLastEditTimestampOk returns a tuple with the LastEditTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetLastEditTimestamp(v interface{})`

SetLastEditTimestamp sets LastEditTimestamp field to given value.

### HasLastEditTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasLastEditTimestamp() bool`

HasLastEditTimestamp returns a boolean if a field has been set.

### SetLastEditTimestampNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetLastEditTimestampNil(b bool)`

 SetLastEditTimestampNil sets the value for LastEditTimestamp to be an explicit nil

### UnsetLastEditTimestamp
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetLastEditTimestamp()`

UnsetLastEditTimestamp ensures that no value is present for LastEditTimestamp, not even an explicit nil
### GetLastMovedTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetLastMovedTimestamp() interface{}`

GetLastMovedTimestamp returns the LastMovedTimestamp field if non-nil, zero value otherwise.

### GetLastMovedTimestampOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetLastMovedTimestampOk() (*interface{}, bool)`

GetLastMovedTimestampOk returns a tuple with the LastMovedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMovedTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetLastMovedTimestamp(v interface{})`

SetLastMovedTimestamp sets LastMovedTimestamp field to given value.

### HasLastMovedTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasLastMovedTimestamp() bool`

HasLastMovedTimestamp returns a boolean if a field has been set.

### SetLastMovedTimestampNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetLastMovedTimestampNil(b bool)`

 SetLastMovedTimestampNil sets the value for LastMovedTimestamp to be an explicit nil

### UnsetLastMovedTimestamp
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetLastMovedTimestamp()`

UnsetLastMovedTimestamp ensures that no value is present for LastMovedTimestamp, not even an explicit nil
### GetReactions

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetReactions() interface{}`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetReactionsOk() (*interface{}, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetReactions(v interface{})`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### SetReactionsNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetReactionsNil(b bool)`

 SetReactionsNil sets the value for Reactions to be an explicit nil

### UnsetReactions
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetReactions()`

UnsetReactions ensures that no value is present for Reactions, not even an explicit nil
### GetRecipientId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetRecipientId() interface{}`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetRecipientIdOk() (*interface{}, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetRecipientId(v interface{})`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### SetRecipientIdNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetRecipientIdNil(b bool)`

 SetRecipientIdNil sets the value for RecipientId to be an explicit nil

### UnsetRecipientId
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetRecipientId()`

UnsetRecipientId ensures that no value is present for RecipientId, not even an explicit nil
### GetSenderEmail

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSenderEmail() interface{}`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSenderEmailOk() (*interface{}, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSenderEmail(v interface{})`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### SetSenderEmailNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSenderEmailNil(b bool)`

 SetSenderEmailNil sets the value for SenderEmail to be an explicit nil

### UnsetSenderEmail
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetSenderEmail()`

UnsetSenderEmail ensures that no value is present for SenderEmail, not even an explicit nil
### GetSenderFullName

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSenderFullName() interface{}`

GetSenderFullName returns the SenderFullName field if non-nil, zero value otherwise.

### GetSenderFullNameOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSenderFullNameOk() (*interface{}, bool)`

GetSenderFullNameOk returns a tuple with the SenderFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderFullName

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSenderFullName(v interface{})`

SetSenderFullName sets SenderFullName field to given value.

### HasSenderFullName

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasSenderFullName() bool`

HasSenderFullName returns a boolean if a field has been set.

### SetSenderFullNameNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSenderFullNameNil(b bool)`

 SetSenderFullNameNil sets the value for SenderFullName to be an explicit nil

### UnsetSenderFullName
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetSenderFullName()`

UnsetSenderFullName ensures that no value is present for SenderFullName, not even an explicit nil
### GetSenderId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSenderId() interface{}`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSenderIdOk() (*interface{}, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSenderId(v interface{})`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### SetSenderIdNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSenderIdNil(b bool)`

 SetSenderIdNil sets the value for SenderId to be an explicit nil

### UnsetSenderId
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetSenderId()`

UnsetSenderId ensures that no value is present for SenderId, not even an explicit nil
### GetSenderRealmStr

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSenderRealmStr() interface{}`

GetSenderRealmStr returns the SenderRealmStr field if non-nil, zero value otherwise.

### GetSenderRealmStrOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSenderRealmStrOk() (*interface{}, bool)`

GetSenderRealmStrOk returns a tuple with the SenderRealmStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderRealmStr

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSenderRealmStr(v interface{})`

SetSenderRealmStr sets SenderRealmStr field to given value.

### HasSenderRealmStr

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasSenderRealmStr() bool`

HasSenderRealmStr returns a boolean if a field has been set.

### SetSenderRealmStrNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSenderRealmStrNil(b bool)`

 SetSenderRealmStrNil sets the value for SenderRealmStr to be an explicit nil

### UnsetSenderRealmStr
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetSenderRealmStr()`

UnsetSenderRealmStr ensures that no value is present for SenderRealmStr, not even an explicit nil
### GetStreamId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetStreamId() interface{}`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetStreamIdOk() (*interface{}, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetStreamId(v interface{})`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### SetStreamIdNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetStreamIdNil(b bool)`

 SetStreamIdNil sets the value for StreamId to be an explicit nil

### UnsetStreamId
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetStreamId()`

UnsetStreamId ensures that no value is present for StreamId, not even an explicit nil
### GetSubject

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSubject() interface{}`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSubjectOk() (*interface{}, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSubject(v interface{})`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetSubmessages

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSubmessages() interface{}`

GetSubmessages returns the Submessages field if non-nil, zero value otherwise.

### GetSubmessagesOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetSubmessagesOk() (*interface{}, bool)`

GetSubmessagesOk returns a tuple with the Submessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmessages

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSubmessages(v interface{})`

SetSubmessages sets Submessages field to given value.

### HasSubmessages

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasSubmessages() bool`

HasSubmessages returns a boolean if a field has been set.

### SetSubmessagesNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetSubmessagesNil(b bool)`

 SetSubmessagesNil sets the value for Submessages to be an explicit nil

### UnsetSubmessages
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetSubmessages()`

UnsetSubmessages ensures that no value is present for Submessages, not even an explicit nil
### GetTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetTimestamp() interface{}`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetTimestampOk() (*interface{}, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetTimestamp(v interface{})`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTopicLinks

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetTopicLinks() interface{}`

GetTopicLinks returns the TopicLinks field if non-nil, zero value otherwise.

### GetTopicLinksOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetTopicLinksOk() (*interface{}, bool)`

GetTopicLinksOk returns a tuple with the TopicLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicLinks

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetTopicLinks(v interface{})`

SetTopicLinks sets TopicLinks field to given value.

### HasTopicLinks

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasTopicLinks() bool`

HasTopicLinks returns a boolean if a field has been set.

### SetTopicLinksNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetTopicLinksNil(b bool)`

 SetTopicLinksNil sets the value for TopicLinks to be an explicit nil

### UnsetTopicLinks
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetTopicLinks()`

UnsetTopicLinks ensures that no value is present for TopicLinks, not even an explicit nil
### GetType

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ZulipOutgoingWebhooks200ResponseMessage) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetRenderedContent

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetRenderedContent() string`

GetRenderedContent returns the RenderedContent field if non-nil, zero value otherwise.

### GetRenderedContentOk

`func (o *ZulipOutgoingWebhooks200ResponseMessage) GetRenderedContentOk() (*string, bool)`

GetRenderedContentOk returns a tuple with the RenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedContent

`func (o *ZulipOutgoingWebhooks200ResponseMessage) SetRenderedContent(v string)`

SetRenderedContent sets RenderedContent field to given value.

### HasRenderedContent

`func (o *ZulipOutgoingWebhooks200ResponseMessage) HasRenderedContent() bool`

HasRenderedContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


