# EventEnvelopeOneOf15PresenceValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to **string** | The client&#39;s platform name.  **Changes**: Starting with Zulip 7.0 (feature level 178), this will always be &#x60;\&quot;website\&quot;&#x60; as the server no longer stores which client submitted presence updates.  | [optional] 
**Status** | Pointer to **string** | The status of the user on this client. Will be either &#x60;idle&#x60; or &#x60;active&#x60;.  | [optional] 
**Timestamp** | Pointer to **int32** | The UNIX timestamp of when this client sent the user&#39;s presence to the server with the precision of a second.  | [optional] 
**Pushable** | Pointer to **bool** | Whether the client is capable of showing mobile/push notifications to the user.  **Changes**: Starting with Zulip 7.0 (feature level 178), this will always be &#x60;false&#x60; as the server no longer stores which client submitted presence updates.  | [optional] 

## Methods

### NewEventEnvelopeOneOf15PresenceValue

`func NewEventEnvelopeOneOf15PresenceValue() *EventEnvelopeOneOf15PresenceValue`

NewEventEnvelopeOneOf15PresenceValue instantiates a new EventEnvelopeOneOf15PresenceValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf15PresenceValueWithDefaults

`func NewEventEnvelopeOneOf15PresenceValueWithDefaults() *EventEnvelopeOneOf15PresenceValue`

NewEventEnvelopeOneOf15PresenceValueWithDefaults instantiates a new EventEnvelopeOneOf15PresenceValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *EventEnvelopeOneOf15PresenceValue) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *EventEnvelopeOneOf15PresenceValue) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *EventEnvelopeOneOf15PresenceValue) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *EventEnvelopeOneOf15PresenceValue) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetStatus

`func (o *EventEnvelopeOneOf15PresenceValue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventEnvelopeOneOf15PresenceValue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventEnvelopeOneOf15PresenceValue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventEnvelopeOneOf15PresenceValue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventEnvelopeOneOf15PresenceValue) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventEnvelopeOneOf15PresenceValue) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventEnvelopeOneOf15PresenceValue) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EventEnvelopeOneOf15PresenceValue) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPushable

`func (o *EventEnvelopeOneOf15PresenceValue) GetPushable() bool`

GetPushable returns the Pushable field if non-nil, zero value otherwise.

### GetPushableOk

`func (o *EventEnvelopeOneOf15PresenceValue) GetPushableOk() (*bool, bool)`

GetPushableOk returns a tuple with the Pushable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushable

`func (o *EventEnvelopeOneOf15PresenceValue) SetPushable(v bool)`

SetPushable sets Pushable field to given value.

### HasPushable

`func (o *EventEnvelopeOneOf15PresenceValue) HasPushable() bool`

HasPushable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


