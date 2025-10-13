# DeactivateOwnUser400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**Entity** | **string** | An internationalized string that notes if the current user is the only organization owner or the only user in the organization.  | 
**IsLastOwner** | **bool** | Whether the current user is the only organization owner (meaning there are other active users in the organization) or the only active user in the organization.  | 

## Methods

### NewDeactivateOwnUser400Response

`func NewDeactivateOwnUser400Response(result interface{}, msg interface{}, code interface{}, entity string, isLastOwner bool, ) *DeactivateOwnUser400Response`

NewDeactivateOwnUser400Response instantiates a new DeactivateOwnUser400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeactivateOwnUser400ResponseWithDefaults

`func NewDeactivateOwnUser400ResponseWithDefaults() *DeactivateOwnUser400Response`

NewDeactivateOwnUser400ResponseWithDefaults instantiates a new DeactivateOwnUser400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeactivateOwnUser400Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeactivateOwnUser400Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeactivateOwnUser400Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *DeactivateOwnUser400Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *DeactivateOwnUser400Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *DeactivateOwnUser400Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DeactivateOwnUser400Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DeactivateOwnUser400Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *DeactivateOwnUser400Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *DeactivateOwnUser400Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *DeactivateOwnUser400Response) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeactivateOwnUser400Response) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeactivateOwnUser400Response) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *DeactivateOwnUser400Response) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *DeactivateOwnUser400Response) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetEntity

`func (o *DeactivateOwnUser400Response) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *DeactivateOwnUser400Response) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *DeactivateOwnUser400Response) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetIsLastOwner

`func (o *DeactivateOwnUser400Response) GetIsLastOwner() bool`

GetIsLastOwner returns the IsLastOwner field if non-nil, zero value otherwise.

### GetIsLastOwnerOk

`func (o *DeactivateOwnUser400Response) GetIsLastOwnerOk() (*bool, bool)`

GetIsLastOwnerOk returns a tuple with the IsLastOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastOwner

`func (o *DeactivateOwnUser400Response) SetIsLastOwner(v bool)`

SetIsLastOwner sets IsLastOwner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


