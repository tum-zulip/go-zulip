# StreamUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**StreamId** | Pointer to **int32** | The ID of the channel whose details have changed.  | [optional] 
**Name** | Pointer to **string** | The name of the channel whose details have changed.  | [optional] 
**Property** | Pointer to **string** | The property of the channel which has changed. See [GET /streams](/api/get-streams#response) response for details on the various properties of a channel.  Clients should handle an \&quot;unknown\&quot; property received here without crashing, since that can happen when connecting to a server running a newer version of Zulip with new features.  | [optional] 
**Value** | Pointer to [**EventEnvelopeOneOf18Value**](EventEnvelopeOneOf18Value.md) |  | [optional] 
**RenderedDescription** | Pointer to **string** | Note: Only present if the changed property was &#x60;description&#x60;.  The short description of the channel rendered as HTML, intended to be used when displaying the channel description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**HistoryPublicToSubscribers** | Pointer to **bool** | Note: Only present if the changed property was &#x60;invite_only&#x60;.  Whether the history of the channel is public to its subscribers.  Currently always true for public channels (i.e. &#x60;\&quot;invite_only\&quot;: false&#x60; implies &#x60;\&quot;history_public_to_subscribers\&quot;: true&#x60;), but clients should not make that assumption, as we may change that behavior in the future.  | [optional] 
**IsWebPublic** | Pointer to **bool** | Note: Only present if the changed property was &#x60;invite_only&#x60;.  Whether the channel&#39;s history is now readable by web-public spectators.  **Changes**: New in Zulip 5.0 (feature level 71).  | [optional] 

## Methods

### NewEventEnvelopeOneOf18

`func NewEventEnvelopeOneOf18() *StreamUpdateEvent`

NewEventEnvelopeOneOf18 instantiates a new StreamUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf18WithDefaults

`func NewEventEnvelopeOneOf18WithDefaults() *StreamUpdateEvent`

NewEventEnvelopeOneOf18WithDefaults instantiates a new StreamUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StreamUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *StreamUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *StreamUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *StreamUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *StreamUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *StreamUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetStreamId

`func (o *StreamUpdateEvent) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *StreamUpdateEvent) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *StreamUpdateEvent) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *StreamUpdateEvent) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetName

`func (o *StreamUpdateEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamUpdateEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamUpdateEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamUpdateEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperty

`func (o *StreamUpdateEvent) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamUpdateEvent) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamUpdateEvent) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *StreamUpdateEvent) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *StreamUpdateEvent) GetValue() EventEnvelopeOneOf18Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StreamUpdateEvent) GetValueOk() (*EventEnvelopeOneOf18Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StreamUpdateEvent) SetValue(v EventEnvelopeOneOf18Value)`

SetValue sets Value field to given value.

### HasValue

`func (o *StreamUpdateEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRenderedDescription

`func (o *StreamUpdateEvent) GetRenderedDescription() string`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *StreamUpdateEvent) GetRenderedDescriptionOk() (*string, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *StreamUpdateEvent) SetRenderedDescription(v string)`

SetRenderedDescription sets RenderedDescription field to given value.

### HasRenderedDescription

`func (o *StreamUpdateEvent) HasRenderedDescription() bool`

HasRenderedDescription returns a boolean if a field has been set.

### GetHistoryPublicToSubscribers

`func (o *StreamUpdateEvent) GetHistoryPublicToSubscribers() bool`

GetHistoryPublicToSubscribers returns the HistoryPublicToSubscribers field if non-nil, zero value otherwise.

### GetHistoryPublicToSubscribersOk

`func (o *StreamUpdateEvent) GetHistoryPublicToSubscribersOk() (*bool, bool)`

GetHistoryPublicToSubscribersOk returns a tuple with the HistoryPublicToSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPublicToSubscribers

`func (o *StreamUpdateEvent) SetHistoryPublicToSubscribers(v bool)`

SetHistoryPublicToSubscribers sets HistoryPublicToSubscribers field to given value.

### HasHistoryPublicToSubscribers

`func (o *StreamUpdateEvent) HasHistoryPublicToSubscribers() bool`

HasHistoryPublicToSubscribers returns a boolean if a field has been set.

### GetIsWebPublic

`func (o *StreamUpdateEvent) GetIsWebPublic() bool`

GetIsWebPublic returns the IsWebPublic field if non-nil, zero value otherwise.

### GetIsWebPublicOk

`func (o *StreamUpdateEvent) GetIsWebPublicOk() (*bool, bool)`

GetIsWebPublicOk returns a tuple with the IsWebPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebPublic

`func (o *StreamUpdateEvent) SetIsWebPublic(v bool)`

SetIsWebPublic sets IsWebPublic field to given value.

### HasIsWebPublic

`func (o *StreamUpdateEvent) HasIsWebPublic() bool`

HasIsWebPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


