# AttachmentAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Attachment** | Pointer to [**Attachment**](Attachment.md) |  | [optional] 
**UploadSpaceUsed** | Pointer to **int32** | The total size of all files uploaded by in the organization, in bytes.  | [optional] 

## Methods

### NewEventEnvelopeOneOf21

`func NewEventEnvelopeOneOf21() *AttachmentAddEvent`

NewEventEnvelopeOneOf21 instantiates a new AttachmentAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf21WithDefaults

`func NewEventEnvelopeOneOf21WithDefaults() *AttachmentAddEvent`

NewEventEnvelopeOneOf21WithDefaults instantiates a new AttachmentAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AttachmentAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AttachmentAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *AttachmentAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AttachmentAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AttachmentAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *AttachmentAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetAttachment

`func (o *AttachmentAddEvent) GetAttachment() Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *AttachmentAddEvent) GetAttachmentOk() (*Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *AttachmentAddEvent) SetAttachment(v Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *AttachmentAddEvent) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetUploadSpaceUsed

`func (o *AttachmentAddEvent) GetUploadSpaceUsed() int32`

GetUploadSpaceUsed returns the UploadSpaceUsed field if non-nil, zero value otherwise.

### GetUploadSpaceUsedOk

`func (o *AttachmentAddEvent) GetUploadSpaceUsedOk() (*int32, bool)`

GetUploadSpaceUsedOk returns a tuple with the UploadSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpaceUsed

`func (o *AttachmentAddEvent) SetUploadSpaceUsed(v int32)`

SetUploadSpaceUsed sets UploadSpaceUsed field to given value.

### HasUploadSpaceUsed

`func (o *AttachmentAddEvent) HasUploadSpaceUsed() bool`

HasUploadSpaceUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


