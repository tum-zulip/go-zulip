# InvitesChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf12

`func NewGetEvents200ResponseAllOfEventsInnerOneOf12() *InvitesChangedEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf12 instantiates a new InvitesChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf12WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf12WithDefaults() *InvitesChangedEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf12WithDefaults instantiates a new InvitesChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvitesChangedEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitesChangedEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitesChangedEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvitesChangedEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InvitesChangedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvitesChangedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvitesChangedEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvitesChangedEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


