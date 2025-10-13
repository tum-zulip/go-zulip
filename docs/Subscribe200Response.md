# Subscribe200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Subscribed** | Pointer to **map[string][]string** | A dictionary where the key is the ID of the user and the value is a list of the names of the channels that user was subscribed to as a result of the request.  **Changes**: Before Zulip 10.0 (feature level 289), the user keys were Zulip API email addresses, not user ID.  | [optional] 
**AlreadySubscribed** | Pointer to **map[string][]string** | A dictionary where the key is the ID of the user and the value is a list of the names of the channels that where the user was not added as a subscriber in this request, because they were already a subscriber.  **Changes**: Before Zulip 10.0 (feature level 289), the user keys were Zulip API email addresses, not user IDs.  | [optional] 
**Unauthorized** | Pointer to **[]string** | A list of names of channels that the requesting user/bot was not authorized to subscribe to. Only present if &#x60;\&quot;authorization_errors_fatal\&quot;: false&#x60;.  | [optional] 
**NewSubscriptionMessagesSent** | Pointer to **bool** | Only present if the parameter &#x60;send_new_subscription_messages&#x60; in the request was &#x60;true&#x60;.  Whether Notification Bot DMs in fact sent to the added subscribers as requested by the &#x60;send_new_subscription_messages&#x60; parameter. Clients may find this value useful to communicate with users about the effect of this request.  **Changes**: New in Zulip 11.0 (feature level 397).  | [optional] 

## Methods

### NewSubscribe200Response

`func NewSubscribe200Response(result interface{}, msg interface{}, ) *Subscribe200Response`

NewSubscribe200Response instantiates a new Subscribe200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribe200ResponseWithDefaults

`func NewSubscribe200ResponseWithDefaults() *Subscribe200Response`

NewSubscribe200ResponseWithDefaults instantiates a new Subscribe200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *Subscribe200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Subscribe200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Subscribe200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *Subscribe200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *Subscribe200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *Subscribe200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *Subscribe200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *Subscribe200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *Subscribe200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *Subscribe200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *Subscribe200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *Subscribe200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *Subscribe200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *Subscribe200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *Subscribe200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *Subscribe200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetSubscribed

`func (o *Subscribe200Response) GetSubscribed() map[string][]string`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *Subscribe200Response) GetSubscribedOk() (*map[string][]string, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *Subscribe200Response) SetSubscribed(v map[string][]string)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *Subscribe200Response) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.

### GetAlreadySubscribed

`func (o *Subscribe200Response) GetAlreadySubscribed() map[string][]string`

GetAlreadySubscribed returns the AlreadySubscribed field if non-nil, zero value otherwise.

### GetAlreadySubscribedOk

`func (o *Subscribe200Response) GetAlreadySubscribedOk() (*map[string][]string, bool)`

GetAlreadySubscribedOk returns a tuple with the AlreadySubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadySubscribed

`func (o *Subscribe200Response) SetAlreadySubscribed(v map[string][]string)`

SetAlreadySubscribed sets AlreadySubscribed field to given value.

### HasAlreadySubscribed

`func (o *Subscribe200Response) HasAlreadySubscribed() bool`

HasAlreadySubscribed returns a boolean if a field has been set.

### GetUnauthorized

`func (o *Subscribe200Response) GetUnauthorized() []string`

GetUnauthorized returns the Unauthorized field if non-nil, zero value otherwise.

### GetUnauthorizedOk

`func (o *Subscribe200Response) GetUnauthorizedOk() (*[]string, bool)`

GetUnauthorizedOk returns a tuple with the Unauthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthorized

`func (o *Subscribe200Response) SetUnauthorized(v []string)`

SetUnauthorized sets Unauthorized field to given value.

### HasUnauthorized

`func (o *Subscribe200Response) HasUnauthorized() bool`

HasUnauthorized returns a boolean if a field has been set.

### GetNewSubscriptionMessagesSent

`func (o *Subscribe200Response) GetNewSubscriptionMessagesSent() bool`

GetNewSubscriptionMessagesSent returns the NewSubscriptionMessagesSent field if non-nil, zero value otherwise.

### GetNewSubscriptionMessagesSentOk

`func (o *Subscribe200Response) GetNewSubscriptionMessagesSentOk() (*bool, bool)`

GetNewSubscriptionMessagesSentOk returns a tuple with the NewSubscriptionMessagesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSubscriptionMessagesSent

`func (o *Subscribe200Response) SetNewSubscriptionMessagesSent(v bool)`

SetNewSubscriptionMessagesSent sets NewSubscriptionMessagesSent field to given value.

### HasNewSubscriptionMessagesSent

`func (o *Subscribe200Response) HasNewSubscriptionMessagesSent() bool`

HasNewSubscriptionMessagesSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


