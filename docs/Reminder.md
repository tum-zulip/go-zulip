# Reminder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReminderId** | **int32** | The unique ID of the reminder, which can be used to delete the reminder.  This is different from the unique ID that the message would have after being sent.  | 
**Type** | **string** | The type of the reminder. Always set to &#x60;\&quot;private\&quot;&#x60;.  | 
**To** | **[]int32** | Contains the ID of the user who scheduled the reminder, and to which the reminder will be sent.  | 
**Content** | **string** | The content/body of the reminder, in [Zulip-flavored Markdown](/help/format-your-message-using-markdown) format.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | 
**RenderedContent** | **string** | The content/body of the reminder rendered in HTML.  | 
**ScheduledDeliveryTimestamp** | **int32** | The UNIX timestamp for when the message will be sent by the server, in UTC seconds.  | 
**Failed** | **bool** | Whether the server has tried to send the reminder and it failed to successfully send.  Clients that support unscheduling reminders should display scheduled messages with &#x60;\&quot;failed\&quot;: true&#x60; with an indicator that the server failed to send the message at the scheduled time, so that the user is aware of the failure and can get the content of the scheduled message.  | 
**ReminderTargetMessageId** | **int32** | The ID of the message that the reminder is created for.  | 

## Methods

### NewReminder

`func NewReminder(reminderId int32, type_ string, to []int32, content string, renderedContent string, scheduledDeliveryTimestamp int32, failed bool, reminderTargetMessageId int32, ) *Reminder`

NewReminder instantiates a new Reminder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReminderWithDefaults

`func NewReminderWithDefaults() *Reminder`

NewReminderWithDefaults instantiates a new Reminder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReminderId

`func (o *Reminder) GetReminderId() int32`

GetReminderId returns the ReminderId field if non-nil, zero value otherwise.

### GetReminderIdOk

`func (o *Reminder) GetReminderIdOk() (*int32, bool)`

GetReminderIdOk returns a tuple with the ReminderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderId

`func (o *Reminder) SetReminderId(v int32)`

SetReminderId sets ReminderId field to given value.


### GetType

`func (o *Reminder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Reminder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Reminder) SetType(v string)`

SetType sets Type field to given value.


### GetTo

`func (o *Reminder) GetTo() []int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Reminder) GetToOk() (*[]int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Reminder) SetTo(v []int32)`

SetTo sets To field to given value.


### GetContent

`func (o *Reminder) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Reminder) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Reminder) SetContent(v string)`

SetContent sets Content field to given value.


### GetRenderedContent

`func (o *Reminder) GetRenderedContent() string`

GetRenderedContent returns the RenderedContent field if non-nil, zero value otherwise.

### GetRenderedContentOk

`func (o *Reminder) GetRenderedContentOk() (*string, bool)`

GetRenderedContentOk returns a tuple with the RenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedContent

`func (o *Reminder) SetRenderedContent(v string)`

SetRenderedContent sets RenderedContent field to given value.


### GetScheduledDeliveryTimestamp

`func (o *Reminder) GetScheduledDeliveryTimestamp() int32`

GetScheduledDeliveryTimestamp returns the ScheduledDeliveryTimestamp field if non-nil, zero value otherwise.

### GetScheduledDeliveryTimestampOk

`func (o *Reminder) GetScheduledDeliveryTimestampOk() (*int32, bool)`

GetScheduledDeliveryTimestampOk returns a tuple with the ScheduledDeliveryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryTimestamp

`func (o *Reminder) SetScheduledDeliveryTimestamp(v int32)`

SetScheduledDeliveryTimestamp sets ScheduledDeliveryTimestamp field to given value.


### GetFailed

`func (o *Reminder) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *Reminder) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *Reminder) SetFailed(v bool)`

SetFailed sets Failed field to given value.


### GetReminderTargetMessageId

`func (o *Reminder) GetReminderTargetMessageId() int32`

GetReminderTargetMessageId returns the ReminderTargetMessageId field if non-nil, zero value otherwise.

### GetReminderTargetMessageIdOk

`func (o *Reminder) GetReminderTargetMessageIdOk() (*int32, bool)`

GetReminderTargetMessageIdOk returns a tuple with the ReminderTargetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderTargetMessageId

`func (o *Reminder) SetReminderTargetMessageId(v int32)`

SetReminderTargetMessageId sets ReminderTargetMessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


