# UpdateFlagsNarrowClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** |  | 
**Operand** | [**UpdateFlagsNarrowOperand**](UpdateFlagsNarrowOperand.md) |  | 
**Negated** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateMessageFlagsForNarrowRequestNarrowInner

`func NewUpdateMessageFlagsForNarrowRequestNarrowInner(operator string, operand UpdateFlagsNarrowOperand, ) *UpdateFlagsNarrowClause`

NewUpdateMessageFlagsForNarrowRequestNarrowInner instantiates a new UpdateFlagsNarrowClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessageFlagsForNarrowRequestNarrowInnerWithDefaults

`func NewUpdateMessageFlagsForNarrowRequestNarrowInnerWithDefaults() *UpdateFlagsNarrowClause`

NewUpdateMessageFlagsForNarrowRequestNarrowInnerWithDefaults instantiates a new UpdateFlagsNarrowClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *UpdateFlagsNarrowClause) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpdateFlagsNarrowClause) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpdateFlagsNarrowClause) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetOperand

`func (o *UpdateFlagsNarrowClause) GetOperand() UpdateFlagsNarrowOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *UpdateFlagsNarrowClause) GetOperandOk() (*UpdateFlagsNarrowOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *UpdateFlagsNarrowClause) SetOperand(v UpdateFlagsNarrowOperand)`

SetOperand sets Operand field to given value.


### GetNegated

`func (o *UpdateFlagsNarrowClause) GetNegated() bool`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *UpdateFlagsNarrowClause) GetNegatedOk() (*bool, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *UpdateFlagsNarrowClause) SetNegated(v bool)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *UpdateFlagsNarrowClause) HasNegated() bool`

HasNegated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


