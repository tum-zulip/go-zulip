# SavedSnippetsUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**SavedSnippet** | Pointer to [**SavedSnippet**](SavedSnippet.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf76

`func NewGetEvents200ResponseAllOfEventsInnerOneOf76() *SavedSnippetsUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf76 instantiates a new SavedSnippetsUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf76WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf76WithDefaults() *SavedSnippetsUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf76WithDefaults instantiates a new SavedSnippetsUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SavedSnippetsUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedSnippetsUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedSnippetsUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SavedSnippetsUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SavedSnippetsUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SavedSnippetsUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SavedSnippetsUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SavedSnippetsUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *SavedSnippetsUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SavedSnippetsUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SavedSnippetsUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SavedSnippetsUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetSavedSnippet

`func (o *SavedSnippetsUpdateEvent) GetSavedSnippet() SavedSnippet`

GetSavedSnippet returns the SavedSnippet field if non-nil, zero value otherwise.

### GetSavedSnippetOk

`func (o *SavedSnippetsUpdateEvent) GetSavedSnippetOk() (*SavedSnippet, bool)`

GetSavedSnippetOk returns a tuple with the SavedSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippet

`func (o *SavedSnippetsUpdateEvent) SetSavedSnippet(v SavedSnippet)`

SetSavedSnippet sets SavedSnippet field to given value.

### HasSavedSnippet

`func (o *SavedSnippetsUpdateEvent) HasSavedSnippet() bool`

HasSavedSnippet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


