# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID for the attachment.  | [optional] 
**Name** | Pointer to **string** | Name of the uploaded file.  | [optional] 
**PathId** | Pointer to **string** | A representation of the path of the file within the repository of user-uploaded files. If the &#x60;path_id&#x60; of a file is &#x60;{realm_id}/ab/cdef/temp_file.py&#x60;, its URL will be: &#x60;{server_url}/user_uploads/{realm_id}/ab/cdef/temp_file.py&#x60;.  | [optional] 
**Size** | Pointer to **int32** | Size of the file in bytes.  | [optional] 
**CreateTime** | Pointer to **int32** | Time when the attachment was uploaded as a UNIX timestamp multiplied by 1000 (matching the format of getTime() in JavaScript).  **Changes**: Changed in Zulip 3.0 (feature level 22). This field was previously a floating point number.  | [optional] 
**Messages** | Pointer to [**[]AttachmentMessagesInner**](AttachmentMessagesInner.md) | Contains basic details on any Zulip messages that have been sent referencing this [uploaded file](/api/upload-file). This includes messages sent by any user in the Zulip organization who sent a message containing a link to the uploaded file.  | [optional] 

## Methods

### NewAttachment

`func NewAttachment() *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attachment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attachment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attachment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Attachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Attachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Attachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Attachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Attachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPathId

`func (o *Attachment) GetPathId() string`

GetPathId returns the PathId field if non-nil, zero value otherwise.

### GetPathIdOk

`func (o *Attachment) GetPathIdOk() (*string, bool)`

GetPathIdOk returns a tuple with the PathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathId

`func (o *Attachment) SetPathId(v string)`

SetPathId sets PathId field to given value.

### HasPathId

`func (o *Attachment) HasPathId() bool`

HasPathId returns a boolean if a field has been set.

### GetSize

`func (o *Attachment) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Attachment) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Attachment) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Attachment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreateTime

`func (o *Attachment) GetCreateTime() int32`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Attachment) GetCreateTimeOk() (*int32, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Attachment) SetCreateTime(v int32)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Attachment) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetMessages

`func (o *Attachment) GetMessages() []AttachmentMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Attachment) GetMessagesOk() (*[]AttachmentMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Attachment) SetMessages(v []AttachmentMessagesInner)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Attachment) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


