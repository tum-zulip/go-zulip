# DraftsAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Drafts** | Pointer to [**[]Draft**](Draft.md) | An array containing objects for the newly created drafts.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf69

`func NewGetEvents200ResponseAllOfEventsInnerOneOf69() *DraftsAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf69 instantiates a new DraftsAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf69WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf69WithDefaults() *DraftsAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf69WithDefaults instantiates a new DraftsAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DraftsAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DraftsAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DraftsAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DraftsAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DraftsAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DraftsAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DraftsAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DraftsAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *DraftsAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *DraftsAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *DraftsAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *DraftsAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetDrafts

`func (o *DraftsAddEvent) GetDrafts() []Draft`

GetDrafts returns the Drafts field if non-nil, zero value otherwise.

### GetDraftsOk

`func (o *DraftsAddEvent) GetDraftsOk() (*[]Draft, bool)`

GetDraftsOk returns a tuple with the Drafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrafts

`func (o *DraftsAddEvent) SetDrafts(v []Draft)`

SetDrafts sets Drafts field to given value.

### HasDrafts

`func (o *DraftsAddEvent) HasDrafts() bool`

HasDrafts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


