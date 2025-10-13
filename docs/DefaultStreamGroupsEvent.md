# DefaultStreamGroupsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DefaultStreamGroups** | Pointer to [**[]DefaultChannelGroup**](DefaultChannelGroup.md) | An array of dictionaries where each dictionary contains details about a single default channel group.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf28

`func NewGetEvents200ResponseAllOfEventsInnerOneOf28() *DefaultStreamGroupsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf28 instantiates a new DefaultStreamGroupsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf28WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf28WithDefaults() *DefaultStreamGroupsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf28WithDefaults instantiates a new DefaultStreamGroupsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DefaultStreamGroupsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefaultStreamGroupsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefaultStreamGroupsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DefaultStreamGroupsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DefaultStreamGroupsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DefaultStreamGroupsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DefaultStreamGroupsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DefaultStreamGroupsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultStreamGroups

`func (o *DefaultStreamGroupsEvent) GetDefaultStreamGroups() []DefaultChannelGroup`

GetDefaultStreamGroups returns the DefaultStreamGroups field if non-nil, zero value otherwise.

### GetDefaultStreamGroupsOk

`func (o *DefaultStreamGroupsEvent) GetDefaultStreamGroupsOk() (*[]DefaultChannelGroup, bool)`

GetDefaultStreamGroupsOk returns a tuple with the DefaultStreamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStreamGroups

`func (o *DefaultStreamGroupsEvent) SetDefaultStreamGroups(v []DefaultChannelGroup)`

SetDefaultStreamGroups sets DefaultStreamGroups field to given value.

### HasDefaultStreamGroups

`func (o *DefaultStreamGroupsEvent) HasDefaultStreamGroups() bool`

HasDefaultStreamGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


