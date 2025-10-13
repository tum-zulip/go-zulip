# AttachmentMessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateSent** | Pointer to **int32** | Time when the message was sent as a UNIX timestamp multiplied by 1000 (matching the format of getTime() in JavaScript).  **Changes**: Changed in Zulip 3.0 (feature level 22). This field was previously strangely called &#x60;name&#x60; and was a floating point number.  | [optional] 
**Id** | Pointer to **int32** | The unique message ID. Messages should always be displayed sorted by ID.  | [optional] 

## Methods

### NewAttachmentMessagesInner

`func NewAttachmentMessagesInner() *AttachmentMessagesInner`

NewAttachmentMessagesInner instantiates a new AttachmentMessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentMessagesInnerWithDefaults

`func NewAttachmentMessagesInnerWithDefaults() *AttachmentMessagesInner`

NewAttachmentMessagesInnerWithDefaults instantiates a new AttachmentMessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateSent

`func (o *AttachmentMessagesInner) GetDateSent() int32`

GetDateSent returns the DateSent field if non-nil, zero value otherwise.

### GetDateSentOk

`func (o *AttachmentMessagesInner) GetDateSentOk() (*int32, bool)`

GetDateSentOk returns a tuple with the DateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSent

`func (o *AttachmentMessagesInner) SetDateSent(v int32)`

SetDateSent sets DateSent field to given value.

### HasDateSent

`func (o *AttachmentMessagesInner) HasDateSent() bool`

HasDateSent returns a boolean if a field has been set.

### GetId

`func (o *AttachmentMessagesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentMessagesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentMessagesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentMessagesInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


