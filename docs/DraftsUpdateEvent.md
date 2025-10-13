# DraftsUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Draft** | Pointer to [**Draft**](Draft.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf70

`func NewGetEvents200ResponseAllOfEventsInnerOneOf70() *DraftsUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf70 instantiates a new DraftsUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf70WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf70WithDefaults() *DraftsUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf70WithDefaults instantiates a new DraftsUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DraftsUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DraftsUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DraftsUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DraftsUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DraftsUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DraftsUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DraftsUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DraftsUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *DraftsUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *DraftsUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *DraftsUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *DraftsUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetDraft

`func (o *DraftsUpdateEvent) GetDraft() Draft`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *DraftsUpdateEvent) GetDraftOk() (*Draft, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *DraftsUpdateEvent) SetDraft(v Draft)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *DraftsUpdateEvent) HasDraft() bool`

HasDraft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


