# SavedSnippetsRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**SavedSnippetId** | Pointer to **int32** | The ID of the saved snippet that was just deleted.  **Changes**: New in Zulip 10.0 (feature level 297).  | [optional] 

## Methods

### NewEventEnvelopeOneOf77

`func NewEventEnvelopeOneOf77() *SavedSnippetsRemoveEvent`

NewEventEnvelopeOneOf77 instantiates a new SavedSnippetsRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf77WithDefaults

`func NewEventEnvelopeOneOf77WithDefaults() *SavedSnippetsRemoveEvent`

NewEventEnvelopeOneOf77WithDefaults instantiates a new SavedSnippetsRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SavedSnippetsRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedSnippetsRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedSnippetsRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SavedSnippetsRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SavedSnippetsRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SavedSnippetsRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SavedSnippetsRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SavedSnippetsRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *SavedSnippetsRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SavedSnippetsRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SavedSnippetsRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SavedSnippetsRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetSavedSnippetId

`func (o *SavedSnippetsRemoveEvent) GetSavedSnippetId() int32`

GetSavedSnippetId returns the SavedSnippetId field if non-nil, zero value otherwise.

### GetSavedSnippetIdOk

`func (o *SavedSnippetsRemoveEvent) GetSavedSnippetIdOk() (*int32, bool)`

GetSavedSnippetIdOk returns a tuple with the SavedSnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippetId

`func (o *SavedSnippetsRemoveEvent) SetSavedSnippetId(v int32)`

SetSavedSnippetId sets SavedSnippetId field to given value.

### HasSavedSnippetId

`func (o *SavedSnippetsRemoveEvent) HasSavedSnippetId() bool`

HasSavedSnippetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


