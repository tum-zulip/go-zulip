# RegisterQueueResponsePresencesEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTimestamp** | Pointer to **int32** | The UNIX timestamp of the last time a client connected to Zulip reported that the user was actually present (e.g. via focusing a browser window or interacting with a computer running the desktop app).  Clients should display users with a current &#x60;active_timestamp&#x60; as fully present.  | [optional] 
**IdleTimestamp** | Pointer to **int32** | The UNIX timestamp of the last time the user had a client connected to Zulip, including idle clients where the user hasn&#39;t interacted with the system recently.  The Zulip server has no way of distinguishing whether an idle web app user is at their computer, but hasn&#39;t interacted with the Zulip tab recently, or simply left their desktop computer on when they left.  Thus, clients should display users with a current &#x60;idle_timestamp&#x60; but no current &#x60;active_timestamp&#x60; as potentially present.  | [optional] 

## Methods

### NewRegisterQueueResponsePresencesEntry

`func NewRegisterQueueResponsePresencesEntry() *RegisterQueueResponsePresencesEntry`

NewRegisterQueueResponsePresencesEntry instantiates a new RegisterQueueResponsePresencesEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponsePresencesEntryWithDefaults

`func NewRegisterQueueResponsePresencesEntryWithDefaults() *RegisterQueueResponsePresencesEntry`

NewRegisterQueueResponsePresencesEntryWithDefaults instantiates a new RegisterQueueResponsePresencesEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTimestamp

`func (o *RegisterQueueResponsePresencesEntry) GetActiveTimestamp() int32`

GetActiveTimestamp returns the ActiveTimestamp field if non-nil, zero value otherwise.

### GetActiveTimestampOk

`func (o *RegisterQueueResponsePresencesEntry) GetActiveTimestampOk() (*int32, bool)`

GetActiveTimestampOk returns a tuple with the ActiveTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTimestamp

`func (o *RegisterQueueResponsePresencesEntry) SetActiveTimestamp(v int32)`

SetActiveTimestamp sets ActiveTimestamp field to given value.

### HasActiveTimestamp

`func (o *RegisterQueueResponsePresencesEntry) HasActiveTimestamp() bool`

HasActiveTimestamp returns a boolean if a field has been set.

### GetIdleTimestamp

`func (o *RegisterQueueResponsePresencesEntry) GetIdleTimestamp() int32`

GetIdleTimestamp returns the IdleTimestamp field if non-nil, zero value otherwise.

### GetIdleTimestampOk

`func (o *RegisterQueueResponsePresencesEntry) GetIdleTimestampOk() (*int32, bool)`

GetIdleTimestampOk returns a tuple with the IdleTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimestamp

`func (o *RegisterQueueResponsePresencesEntry) SetIdleTimestamp(v int32)`

SetIdleTimestamp sets IdleTimestamp field to given value.

### HasIdleTimestamp

`func (o *RegisterQueueResponsePresencesEntry) HasIdleTimestamp() bool`

HasIdleTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


