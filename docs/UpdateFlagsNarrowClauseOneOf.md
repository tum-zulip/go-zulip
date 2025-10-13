# UpdateFlagsNarrowFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** |  | 
**Operand** | [**UpdateFlagsNarrowOperand**](UpdateFlagsNarrowOperand.md) |  | 
**Negated** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateFlagsNarrowFilter

`func NewUpdateFlagsNarrowFilter(operator string, operand UpdateFlagsNarrowOperand, ) *UpdateFlagsNarrowFilter`

NewUpdateFlagsNarrowFilter instantiates a new UpdateFlagsNarrowFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFlagsNarrowFilterWithDefaults

`func NewUpdateFlagsNarrowFilterWithDefaults() *UpdateFlagsNarrowFilter`

NewUpdateFlagsNarrowFilterWithDefaults instantiates a new UpdateFlagsNarrowFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *UpdateFlagsNarrowFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpdateFlagsNarrowFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpdateFlagsNarrowFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetOperand

`func (o *UpdateFlagsNarrowFilter) GetOperand() UpdateFlagsNarrowOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *UpdateFlagsNarrowFilter) GetOperandOk() (*UpdateFlagsNarrowOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *UpdateFlagsNarrowFilter) SetOperand(v UpdateFlagsNarrowOperand)`

SetOperand sets Operand field to given value.


### GetNegated

`func (o *UpdateFlagsNarrowFilter) GetNegated() bool`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *UpdateFlagsNarrowFilter) GetNegatedOk() (*bool, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *UpdateFlagsNarrowFilter) SetNegated(v bool)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *UpdateFlagsNarrowFilter) HasNegated() bool`

HasNegated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


