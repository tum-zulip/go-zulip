# SavedSnippetsAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**SavedSnippet** | Pointer to [**SavedSnippet**](SavedSnippet.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf75

`func NewEventEnvelopeOneOf75() *SavedSnippetsAddEvent`

NewEventEnvelopeOneOf75 instantiates a new SavedSnippetsAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf75WithDefaults

`func NewEventEnvelopeOneOf75WithDefaults() *SavedSnippetsAddEvent`

NewEventEnvelopeOneOf75WithDefaults instantiates a new SavedSnippetsAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SavedSnippetsAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedSnippetsAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedSnippetsAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SavedSnippetsAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SavedSnippetsAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SavedSnippetsAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SavedSnippetsAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SavedSnippetsAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *SavedSnippetsAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SavedSnippetsAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SavedSnippetsAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SavedSnippetsAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetSavedSnippet

`func (o *SavedSnippetsAddEvent) GetSavedSnippet() SavedSnippet`

GetSavedSnippet returns the SavedSnippet field if non-nil, zero value otherwise.

### GetSavedSnippetOk

`func (o *SavedSnippetsAddEvent) GetSavedSnippetOk() (*SavedSnippet, bool)`

GetSavedSnippetOk returns a tuple with the SavedSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippet

`func (o *SavedSnippetsAddEvent) SetSavedSnippet(v SavedSnippet)`

SetSavedSnippet sets SavedSnippet field to given value.

### HasSavedSnippet

`func (o *SavedSnippetsAddEvent) HasSavedSnippet() bool`

HasSavedSnippet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


