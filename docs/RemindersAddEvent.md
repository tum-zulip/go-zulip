# RemindersAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Reminders** | Pointer to [**[]Reminder**](Reminder.md) | An array of objects containing details of the newly created reminders.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf78

`func NewGetEvents200ResponseAllOfEventsInnerOneOf78() *RemindersAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf78 instantiates a new RemindersAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf78WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf78WithDefaults() *RemindersAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf78WithDefaults instantiates a new RemindersAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemindersAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemindersAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemindersAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RemindersAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RemindersAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RemindersAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RemindersAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RemindersAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RemindersAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RemindersAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RemindersAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RemindersAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetReminders

`func (o *RemindersAddEvent) GetReminders() []Reminder`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *RemindersAddEvent) GetRemindersOk() (*[]Reminder, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *RemindersAddEvent) SetReminders(v []Reminder)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *RemindersAddEvent) HasReminders() bool`

HasReminders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


