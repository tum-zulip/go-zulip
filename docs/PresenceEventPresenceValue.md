# PresenceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Presences** | Pointer to [**map[string]ModernPresenceFormat**](ModernPresenceFormat.md) | Only present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  A dictionary mapping user IDs to the presence data (modern format) for the modified user(s). Clients should support updating multiple users in a single event.  **Changes**: New in Zulip 11.0 (feature level 419).  | [optional] 
**UserId** | Pointer to **int32** | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  The ID of the modified user.  | [optional] 
**Email** | Pointer to **string** | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the &#x60;user_id&#x60;.  | [optional] 
**ServerTimestamp** | Pointer to **float32** | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  The timestamp of when the Zulip server received the user&#39;s presence as a UNIX timestamp.  | [optional] 
**Presence** | Pointer to [**map[string]GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue**](GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue.md) | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  Object containing the presence data (legacy format) of of the modified user.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf15

`func NewGetEvents200ResponseAllOfEventsInnerOneOf15() *PresenceEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf15 instantiates a new PresenceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf15WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf15WithDefaults() *PresenceEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf15WithDefaults instantiates a new PresenceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PresenceEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PresenceEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PresenceEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PresenceEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PresenceEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PresenceEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PresenceEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PresenceEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPresences

`func (o *PresenceEvent) GetPresences() map[string]ModernPresenceFormat`

GetPresences returns the Presences field if non-nil, zero value otherwise.

### GetPresencesOk

`func (o *PresenceEvent) GetPresencesOk() (*map[string]ModernPresenceFormat, bool)`

GetPresencesOk returns a tuple with the Presences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresences

`func (o *PresenceEvent) SetPresences(v map[string]ModernPresenceFormat)`

SetPresences sets Presences field to given value.

### HasPresences

`func (o *PresenceEvent) HasPresences() bool`

HasPresences returns a boolean if a field has been set.

### GetUserId

`func (o *PresenceEvent) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PresenceEvent) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PresenceEvent) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PresenceEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmail

`func (o *PresenceEvent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PresenceEvent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PresenceEvent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PresenceEvent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetServerTimestamp

`func (o *PresenceEvent) GetServerTimestamp() float32`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *PresenceEvent) GetServerTimestampOk() (*float32, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *PresenceEvent) SetServerTimestamp(v float32)`

SetServerTimestamp sets ServerTimestamp field to given value.

### HasServerTimestamp

`func (o *PresenceEvent) HasServerTimestamp() bool`

HasServerTimestamp returns a boolean if a field has been set.

### GetPresence

`func (o *PresenceEvent) GetPresence() map[string]GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *PresenceEvent) GetPresenceOk() (*map[string]GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *PresenceEvent) SetPresence(v map[string]GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *PresenceEvent) HasPresence() bool`

HasPresence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


