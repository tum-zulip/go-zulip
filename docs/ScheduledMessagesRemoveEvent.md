# ScheduledMessagesRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**ScheduledMessageId** | Pointer to **int32** | The ID of the scheduled message that was deleted.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf82

`func NewGetEvents200ResponseAllOfEventsInnerOneOf82() *ScheduledMessagesRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf82 instantiates a new ScheduledMessagesRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf82WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf82WithDefaults() *ScheduledMessagesRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf82WithDefaults instantiates a new ScheduledMessagesRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledMessagesRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledMessagesRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledMessagesRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduledMessagesRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ScheduledMessagesRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduledMessagesRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduledMessagesRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScheduledMessagesRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *ScheduledMessagesRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ScheduledMessagesRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ScheduledMessagesRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ScheduledMessagesRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetScheduledMessageId

`func (o *ScheduledMessagesRemoveEvent) GetScheduledMessageId() int32`

GetScheduledMessageId returns the ScheduledMessageId field if non-nil, zero value otherwise.

### GetScheduledMessageIdOk

`func (o *ScheduledMessagesRemoveEvent) GetScheduledMessageIdOk() (*int32, bool)`

GetScheduledMessageIdOk returns a tuple with the ScheduledMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessageId

`func (o *ScheduledMessagesRemoveEvent) SetScheduledMessageId(v int32)`

SetScheduledMessageId sets ScheduledMessageId field to given value.

### HasScheduledMessageId

`func (o *ScheduledMessagesRemoveEvent) HasScheduledMessageId() bool`

HasScheduledMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


