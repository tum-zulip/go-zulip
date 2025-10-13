# SubscriptionPeerAddEvent4Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name of the channel folder. Only present if the channel folder&#39;s name changed.  | [optional] 
**Description** | Pointer to **string** | The new description of the channel folder. Only present if the description changed.  See [Markdown message formatting](zulip.com/api/message-formatting for details on Zulip&#39;s HTML format.  | [optional] 
**RenderedDescription** | Pointer to **string** | The new rendered description of the channel folder. Only present if the description changed.  | [optional] 
**IsArchived** | Pointer to **bool** | Whether the channel folder is archived or not. Only present if the channel folder is archived or unarchived.  | [optional] 

## Methods

### NewEventEnvelopeOneOf84Data

`func NewEventEnvelopeOneOf84Data() *EventEnvelopeOneOf84Data`

NewEventEnvelopeOneOf84Data instantiates a new EventEnvelopeOneOf84Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf84DataWithDefaults

`func NewEventEnvelopeOneOf84DataWithDefaults() *EventEnvelopeOneOf84Data`

NewEventEnvelopeOneOf84DataWithDefaults instantiates a new EventEnvelopeOneOf84Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventEnvelopeOneOf84Data) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventEnvelopeOneOf84Data) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventEnvelopeOneOf84Data) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventEnvelopeOneOf84Data) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EventEnvelopeOneOf84Data) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventEnvelopeOneOf84Data) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventEnvelopeOneOf84Data) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventEnvelopeOneOf84Data) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRenderedDescription

`func (o *EventEnvelopeOneOf84Data) GetRenderedDescription() string`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *EventEnvelopeOneOf84Data) GetRenderedDescriptionOk() (*string, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *EventEnvelopeOneOf84Data) SetRenderedDescription(v string)`

SetRenderedDescription sets RenderedDescription field to given value.

### HasRenderedDescription

`func (o *EventEnvelopeOneOf84Data) HasRenderedDescription() bool`

HasRenderedDescription returns a boolean if a field has been set.

### GetIsArchived

`func (o *EventEnvelopeOneOf84Data) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *EventEnvelopeOneOf84Data) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *EventEnvelopeOneOf84Data) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *EventEnvelopeOneOf84Data) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


