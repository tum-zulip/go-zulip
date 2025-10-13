# RemindersRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**ReminderId** | Pointer to **int32** | The ID of the reminder that was deleted.  | [optional] 

## Methods

### NewEventEnvelopeOneOf79

`func NewEventEnvelopeOneOf79() *RemindersRemoveEvent`

NewEventEnvelopeOneOf79 instantiates a new RemindersRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf79WithDefaults

`func NewEventEnvelopeOneOf79WithDefaults() *RemindersRemoveEvent`

NewEventEnvelopeOneOf79WithDefaults instantiates a new RemindersRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemindersRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemindersRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemindersRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RemindersRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RemindersRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RemindersRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RemindersRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RemindersRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RemindersRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RemindersRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RemindersRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RemindersRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetReminderId

`func (o *RemindersRemoveEvent) GetReminderId() int32`

GetReminderId returns the ReminderId field if non-nil, zero value otherwise.

### GetReminderIdOk

`func (o *RemindersRemoveEvent) GetReminderIdOk() (*int32, bool)`

GetReminderIdOk returns a tuple with the ReminderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderId

`func (o *RemindersRemoveEvent) SetReminderId(v int32)`

SetReminderId sets ReminderId field to given value.

### HasReminderId

`func (o *RemindersRemoveEvent) HasReminderId() bool`

HasReminderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


