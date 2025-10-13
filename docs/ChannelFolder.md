# ChannelFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the channel folder.  | [optional] 
**DateCreated** | Pointer to **NullableInt32** | The UNIX timestamp for when the channel folder was created, in UTC seconds.  | [optional] 
**CreatorId** | Pointer to **NullableInt32** | The ID of the user who created this channel folder.  | [optional] 
**Description** | Pointer to **string** | The description of the channel folder.  See [Markdown message formatting](zulip.com/api/message-formatting for details on Zulip&#39;s HTML format.  | [optional] 
**RenderedDescription** | Pointer to **string** | The description of the channel folder rendered as HTML, intended to be used when displaying the channel folder description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.  | [optional] 
**Order** | Pointer to **int32** | This value determines in which order the channel folders will be displayed in the UI. The value is 0 indexed, and the value with the lower order will be displayed first.  **Changes**: New in Zulip 11.0 (feature level 414).  | [optional] 
**Id** | Pointer to **int32** | The ID of the channel folder.  | [optional] 
**IsArchived** | Pointer to **bool** | Whether the channel folder is archived or not.  | [optional] 

## Methods

### NewChannelFolder

`func NewChannelFolder() *ChannelFolder`

NewChannelFolder instantiates a new ChannelFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelFolderWithDefaults

`func NewChannelFolderWithDefaults() *ChannelFolder`

NewChannelFolderWithDefaults instantiates a new ChannelFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChannelFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDateCreated

`func (o *ChannelFolder) GetDateCreated() int32`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChannelFolder) GetDateCreatedOk() (*int32, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChannelFolder) SetDateCreated(v int32)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChannelFolder) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ChannelFolder) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ChannelFolder) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetCreatorId

`func (o *ChannelFolder) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ChannelFolder) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ChannelFolder) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ChannelFolder) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *ChannelFolder) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *ChannelFolder) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetDescription

`func (o *ChannelFolder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelFolder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelFolder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelFolder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRenderedDescription

`func (o *ChannelFolder) GetRenderedDescription() string`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *ChannelFolder) GetRenderedDescriptionOk() (*string, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *ChannelFolder) SetRenderedDescription(v string)`

SetRenderedDescription sets RenderedDescription field to given value.

### HasRenderedDescription

`func (o *ChannelFolder) HasRenderedDescription() bool`

HasRenderedDescription returns a boolean if a field has been set.

### GetOrder

`func (o *ChannelFolder) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ChannelFolder) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ChannelFolder) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ChannelFolder) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetId

`func (o *ChannelFolder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelFolder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelFolder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsArchived

`func (o *ChannelFolder) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *ChannelFolder) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *ChannelFolder) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *ChannelFolder) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


