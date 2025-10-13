# GroupSettingValueUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**New** | [**GroupSettingValue**](GroupSettingValue.md) | The new [group-setting value](zulip.com/api/group-setting-values for who would have this permission.  | 
**Old** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | The expected current [group-setting value](zulip.com/api/group-setting-values for who has this permission.  | [optional] 

## Methods

### NewGroupSettingValueUpdate

`func NewGroupSettingValueUpdate(new GroupSettingValue, ) *GroupSettingValueUpdate`

NewGroupSettingValueUpdate instantiates a new GroupSettingValueUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSettingValueUpdateWithDefaults

`func NewGroupSettingValueUpdateWithDefaults() *GroupSettingValueUpdate`

NewGroupSettingValueUpdateWithDefaults instantiates a new GroupSettingValueUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNew

`func (o *GroupSettingValueUpdate) GetNew() GroupSettingValue`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *GroupSettingValueUpdate) GetNewOk() (*GroupSettingValue, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *GroupSettingValueUpdate) SetNew(v GroupSettingValue)`

SetNew sets New field to given value.


### GetOld

`func (o *GroupSettingValueUpdate) GetOld() GroupSettingValue`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *GroupSettingValueUpdate) GetOldOk() (*GroupSettingValue, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *GroupSettingValueUpdate) SetOld(v GroupSettingValue)`

SetOld sets Old field to given value.

### HasOld

`func (o *GroupSettingValueUpdate) HasOld() bool`

HasOld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


