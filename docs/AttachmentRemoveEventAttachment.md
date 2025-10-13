# AttachmentRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Attachment** | Pointer to [**EventEnvelopeOneOf23Attachment**](EventEnvelopeOneOf23Attachment.md) |  | [optional] 
**UploadSpaceUsed** | Pointer to **int32** | The total size of all files uploaded by in the organization, in bytes.  | [optional] 

## Methods

### NewEventEnvelopeOneOf23

`func NewEventEnvelopeOneOf23() *AttachmentRemoveEvent`

NewEventEnvelopeOneOf23 instantiates a new AttachmentRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf23WithDefaults

`func NewEventEnvelopeOneOf23WithDefaults() *AttachmentRemoveEvent`

NewEventEnvelopeOneOf23WithDefaults instantiates a new AttachmentRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AttachmentRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AttachmentRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *AttachmentRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AttachmentRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AttachmentRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *AttachmentRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetAttachment

`func (o *AttachmentRemoveEvent) GetAttachment() EventEnvelopeOneOf23Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *AttachmentRemoveEvent) GetAttachmentOk() (*EventEnvelopeOneOf23Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *AttachmentRemoveEvent) SetAttachment(v EventEnvelopeOneOf23Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *AttachmentRemoveEvent) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetUploadSpaceUsed

`func (o *AttachmentRemoveEvent) GetUploadSpaceUsed() int32`

GetUploadSpaceUsed returns the UploadSpaceUsed field if non-nil, zero value otherwise.

### GetUploadSpaceUsedOk

`func (o *AttachmentRemoveEvent) GetUploadSpaceUsedOk() (*int32, bool)`

GetUploadSpaceUsedOk returns a tuple with the UploadSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpaceUsed

`func (o *AttachmentRemoveEvent) SetUploadSpaceUsed(v int32)`

SetUploadSpaceUsed sets UploadSpaceUsed field to given value.

### HasUploadSpaceUsed

`func (o *AttachmentRemoveEvent) HasUploadSpaceUsed() bool`

HasUploadSpaceUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


