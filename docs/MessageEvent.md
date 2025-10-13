# MessageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to [**MessagesEvent**](MessagesEvent.md) |  | [optional] 
**Flags** | Pointer to **[]string** | The user&#39;s [message flags][message-flags] for the message.  Clients should inspect the flags field rather than assuming that new messages are unread; [muted users](zulip.com/api/mute-user, messages sent by the current user, and more subtle scenarios can result in a new message that the server has already marked as read for the user.  **Changes**: In Zulip 8.0 (feature level 224), the &#x60;wildcard_mentioned&#x60; flag was deprecated in favor of the &#x60;stream_wildcard_mentioned&#x60; and &#x60;topic_wildcard_mentioned&#x60; flags. The &#x60;wildcard_mentioned&#x60; flag exists for backwards compatibility with older clients and equals &#x60;stream_wildcard_mentioned || topic_wildcard_mentioned&#x60;. Clients supporting older server versions should treat this field as a previous name for the &#x60;stream_wildcard_mentioned&#x60; flag as topic wildcard mentions were not available prior to this feature level.  [message-flags]: /api/update-message-flags#available-flags  | [optional] 

## Methods

### NewEventEnvelopeOneOf10

`func NewEventEnvelopeOneOf10() *MessageEvent`

NewEventEnvelopeOneOf10 instantiates a new MessageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf10WithDefaults

`func NewEventEnvelopeOneOf10WithDefaults() *MessageEvent`

NewEventEnvelopeOneOf10WithDefaults instantiates a new MessageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MessageEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MessageEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MessageEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *MessageEvent) GetMessage() MessagesEvent`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageEvent) GetMessageOk() (*MessagesEvent, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageEvent) SetMessage(v MessagesEvent)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MessageEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFlags

`func (o *MessageEvent) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageEvent) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageEvent) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MessageEvent) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


