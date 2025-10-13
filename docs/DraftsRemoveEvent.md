# DraftsRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**DraftId** | Pointer to **int32** | The ID of the draft that was just deleted.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf71

`func NewGetEvents200ResponseAllOfEventsInnerOneOf71() *DraftsRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf71 instantiates a new DraftsRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf71WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf71WithDefaults() *DraftsRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf71WithDefaults instantiates a new DraftsRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DraftsRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DraftsRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DraftsRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DraftsRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DraftsRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DraftsRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DraftsRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DraftsRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *DraftsRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *DraftsRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *DraftsRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *DraftsRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetDraftId

`func (o *DraftsRemoveEvent) GetDraftId() int32`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *DraftsRemoveEvent) GetDraftIdOk() (*int32, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *DraftsRemoveEvent) SetDraftId(v int32)`

SetDraftId sets DraftId field to given value.

### HasDraftId

`func (o *DraftsRemoveEvent) HasDraftId() bool`

HasDraftId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


