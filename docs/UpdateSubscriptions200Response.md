# UpdateSubscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Subscribed** | **map[string][]string** | A dictionary where the key is the Zulip API email address of the user/bot and the value is a list of the names of the channels that were subscribed to as a result of the query.  | 
**AlreadySubscribed** | **map[string][]string** | A dictionary where the key is the Zulip API email address of the user/bot and the value is a list of the names of the channels that the user/bot is already subscribed to.  | 
**NotRemoved** | Pointer to **[]string** | A list of the names of channels that the user is already unsubscribed from, and hence doesn&#39;t need to be unsubscribed.  | [optional] 
**Removed** | **[]string** | A list of the names of channels which were unsubscribed from as a result of the query.  | 
**NewSubscriptionMessagesSent** | Pointer to **bool** | Only present if the parameter &#x60;send_new_subscription_messages&#x60; in the request was &#x60;true&#x60;.  Whether Notification Bot DMs in fact sent to the added subscribers as requested by the &#x60;send_new_subscription_messages&#x60; parameter. Clients may find this value useful to communicate with users about the effect of this request.  **Changes**: New in Zulip 11.0 (feature level 397).  | [optional] 

## Methods

### NewUpdateSubscriptions200Response

`func NewUpdateSubscriptions200Response(result interface{}, msg interface{}, subscribed map[string][]string, alreadySubscribed map[string][]string, removed []string, ) *UpdateSubscriptions200Response`

NewUpdateSubscriptions200Response instantiates a new UpdateSubscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptions200ResponseWithDefaults

`func NewUpdateSubscriptions200ResponseWithDefaults() *UpdateSubscriptions200Response`

NewUpdateSubscriptions200ResponseWithDefaults instantiates a new UpdateSubscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateSubscriptions200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateSubscriptions200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateSubscriptions200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *UpdateSubscriptions200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *UpdateSubscriptions200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *UpdateSubscriptions200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateSubscriptions200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateSubscriptions200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *UpdateSubscriptions200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateSubscriptions200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *UpdateSubscriptions200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *UpdateSubscriptions200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *UpdateSubscriptions200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *UpdateSubscriptions200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *UpdateSubscriptions200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *UpdateSubscriptions200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetSubscribed

`func (o *UpdateSubscriptions200Response) GetSubscribed() map[string][]string`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *UpdateSubscriptions200Response) GetSubscribedOk() (*map[string][]string, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *UpdateSubscriptions200Response) SetSubscribed(v map[string][]string)`

SetSubscribed sets Subscribed field to given value.


### GetAlreadySubscribed

`func (o *UpdateSubscriptions200Response) GetAlreadySubscribed() map[string][]string`

GetAlreadySubscribed returns the AlreadySubscribed field if non-nil, zero value otherwise.

### GetAlreadySubscribedOk

`func (o *UpdateSubscriptions200Response) GetAlreadySubscribedOk() (*map[string][]string, bool)`

GetAlreadySubscribedOk returns a tuple with the AlreadySubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadySubscribed

`func (o *UpdateSubscriptions200Response) SetAlreadySubscribed(v map[string][]string)`

SetAlreadySubscribed sets AlreadySubscribed field to given value.


### GetNotRemoved

`func (o *UpdateSubscriptions200Response) GetNotRemoved() []string`

GetNotRemoved returns the NotRemoved field if non-nil, zero value otherwise.

### GetNotRemovedOk

`func (o *UpdateSubscriptions200Response) GetNotRemovedOk() (*[]string, bool)`

GetNotRemovedOk returns a tuple with the NotRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRemoved

`func (o *UpdateSubscriptions200Response) SetNotRemoved(v []string)`

SetNotRemoved sets NotRemoved field to given value.

### HasNotRemoved

`func (o *UpdateSubscriptions200Response) HasNotRemoved() bool`

HasNotRemoved returns a boolean if a field has been set.

### GetRemoved

`func (o *UpdateSubscriptions200Response) GetRemoved() []string`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *UpdateSubscriptions200Response) GetRemovedOk() (*[]string, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *UpdateSubscriptions200Response) SetRemoved(v []string)`

SetRemoved sets Removed field to given value.


### GetNewSubscriptionMessagesSent

`func (o *UpdateSubscriptions200Response) GetNewSubscriptionMessagesSent() bool`

GetNewSubscriptionMessagesSent returns the NewSubscriptionMessagesSent field if non-nil, zero value otherwise.

### GetNewSubscriptionMessagesSentOk

`func (o *UpdateSubscriptions200Response) GetNewSubscriptionMessagesSentOk() (*bool, bool)`

GetNewSubscriptionMessagesSentOk returns a tuple with the NewSubscriptionMessagesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSubscriptionMessagesSent

`func (o *UpdateSubscriptions200Response) SetNewSubscriptionMessagesSent(v bool)`

SetNewSubscriptionMessagesSent sets NewSubscriptionMessagesSent field to given value.

### HasNewSubscriptionMessagesSent

`func (o *UpdateSubscriptions200Response) HasNewSubscriptionMessagesSent() bool`

HasNewSubscriptionMessagesSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


