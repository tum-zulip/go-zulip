# GetUserPresence200ResponseAllOfPresenceValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int32** | When this update was received. If the timestamp is more than a few minutes in the past, the user is offline.  | [optional] 
**Status** | Pointer to **string** | Whether the user had recently interacted with Zulip at the time of the timestamp.  Will be either &#x60;\&quot;active\&quot;&#x60; or &#x60;\&quot;idle\&quot;&#x60;  | [optional] 

## Methods

### NewGetUserPresence200ResponseAllOfPresenceValue

`func NewGetUserPresence200ResponseAllOfPresenceValue() *GetUserPresence200ResponseAllOfPresenceValue`

NewGetUserPresence200ResponseAllOfPresenceValue instantiates a new GetUserPresence200ResponseAllOfPresenceValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserPresence200ResponseAllOfPresenceValueWithDefaults

`func NewGetUserPresence200ResponseAllOfPresenceValueWithDefaults() *GetUserPresence200ResponseAllOfPresenceValue`

NewGetUserPresence200ResponseAllOfPresenceValueWithDefaults instantiates a new GetUserPresence200ResponseAllOfPresenceValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GetUserPresence200ResponseAllOfPresenceValue) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetUserPresence200ResponseAllOfPresenceValue) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetUserPresence200ResponseAllOfPresenceValue) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetUserPresence200ResponseAllOfPresenceValue) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *GetUserPresence200ResponseAllOfPresenceValue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserPresence200ResponseAllOfPresenceValue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserPresence200ResponseAllOfPresenceValue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetUserPresence200ResponseAllOfPresenceValue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


