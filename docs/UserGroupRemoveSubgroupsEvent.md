# UserGroupRemoveSubgroupsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **int32** | The ID of the user group whose details have changed.  | [optional] 
**DirectSubgroupIds** | Pointer to **[]int32** | Array containing the IDs of the subgroups that have been removed from the user group.  **Changes**: New in Zulip 6.0 (feature level 131). Previously, this was called &#x60;subgroup_ids&#x60;, but clients can ignore older events as this feature level predates subgroups being fully implemented.  | [optional] 

## Methods

### NewEventEnvelopeOneOf48

`func NewEventEnvelopeOneOf48() *UserGroupRemoveSubgroupsEvent`

NewEventEnvelopeOneOf48 instantiates a new UserGroupRemoveSubgroupsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf48WithDefaults

`func NewEventEnvelopeOneOf48WithDefaults() *UserGroupRemoveSubgroupsEvent`

NewEventEnvelopeOneOf48WithDefaults instantiates a new UserGroupRemoveSubgroupsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupRemoveSubgroupsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupRemoveSubgroupsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupRemoveSubgroupsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupRemoveSubgroupsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserGroupRemoveSubgroupsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroupRemoveSubgroupsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroupRemoveSubgroupsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroupRemoveSubgroupsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *UserGroupRemoveSubgroupsEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UserGroupRemoveSubgroupsEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UserGroupRemoveSubgroupsEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UserGroupRemoveSubgroupsEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetGroupId

`func (o *UserGroupRemoveSubgroupsEvent) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UserGroupRemoveSubgroupsEvent) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UserGroupRemoveSubgroupsEvent) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UserGroupRemoveSubgroupsEvent) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDirectSubgroupIds

`func (o *UserGroupRemoveSubgroupsEvent) GetDirectSubgroupIds() []int32`

GetDirectSubgroupIds returns the DirectSubgroupIds field if non-nil, zero value otherwise.

### GetDirectSubgroupIdsOk

`func (o *UserGroupRemoveSubgroupsEvent) GetDirectSubgroupIdsOk() (*[]int32, bool)`

GetDirectSubgroupIdsOk returns a tuple with the DirectSubgroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSubgroupIds

`func (o *UserGroupRemoveSubgroupsEvent) SetDirectSubgroupIds(v []int32)`

SetDirectSubgroupIds sets DirectSubgroupIds field to given value.

### HasDirectSubgroupIds

`func (o *UserGroupRemoveSubgroupsEvent) HasDirectSubgroupIds() bool`

HasDirectSubgroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


