# UpdateSubscriptionSettingsRequestSubscriptionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamId** | **int32** | The unique ID of a channel.  | 
**Property** | **string** | One of the channel properties described below:  - &#x60;\&quot;color\&quot;&#x60;: The hex value of the user&#39;s display color for the channel.  - &#x60;\&quot;is_muted\&quot;&#x60;: Whether the channel is [muted](/help/mute-a-channel).&lt;br&gt;   **Changes**: As of Zulip 6.0 (feature level 139), updating either   &#x60;\&quot;is_muted\&quot;&#x60; or &#x60;\&quot;in_home_view\&quot;&#x60; generates two [subscription update   events](/api/get-events#subscription-update), one for each property,   that are sent to clients. Prior to this feature level, updating either   property only generated a subscription update event for   &#x60;\&quot;in_home_view\&quot;&#x60;. &lt;br&gt;   Prior to Zulip 2.1.0, this feature was represented   by the more confusingly named &#x60;\&quot;in_home_view\&quot;&#x60; (with the   opposite value: &#x60;in_home_view&#x3D;!is_muted&#x60;); for   backwards-compatibility, modern Zulip still accepts that property.  - &#x60;\&quot;pin_to_top\&quot;&#x60;: Whether to pin the channel at the top of the channel list.  - &#x60;\&quot;desktop_notifications\&quot;&#x60;: Whether to show desktop notifications   for all messages sent to the channel.  - &#x60;\&quot;audible_notifications\&quot;&#x60;: Whether to play a sound   notification for all messages sent to the channel.  - &#x60;\&quot;push_notifications\&quot;&#x60;: Whether to trigger a mobile push   notification for all messages sent to the channel.  - &#x60;\&quot;email_notifications\&quot;&#x60;: Whether to trigger an email   notification for all messages sent to the channel.  - &#x60;\&quot;wildcard_mentions_notify\&quot;&#x60;: Whether wildcard mentions trigger   notifications as though they were personal mentions in this channel.  | 
**Value** | [**UpdateSubscriptionSettingsRequestSubscriptionDataInnerValue**](UpdateSubscriptionSettingsRequestSubscriptionDataInnerValue.md) |  | 

## Methods

### NewUpdateSubscriptionSettingsRequestSubscriptionDataInner

`func NewUpdateSubscriptionSettingsRequestSubscriptionDataInner(streamId int32, property string, value UpdateSubscriptionSettingsRequestSubscriptionDataInnerValue, ) *UpdateSubscriptionSettingsRequestSubscriptionDataInner`

NewUpdateSubscriptionSettingsRequestSubscriptionDataInner instantiates a new UpdateSubscriptionSettingsRequestSubscriptionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionSettingsRequestSubscriptionDataInnerWithDefaults

`func NewUpdateSubscriptionSettingsRequestSubscriptionDataInnerWithDefaults() *UpdateSubscriptionSettingsRequestSubscriptionDataInner`

NewUpdateSubscriptionSettingsRequestSubscriptionDataInnerWithDefaults instantiates a new UpdateSubscriptionSettingsRequestSubscriptionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamId

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.


### GetProperty

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetValue

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) GetValue() UpdateSubscriptionSettingsRequestSubscriptionDataInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) GetValueOk() (*UpdateSubscriptionSettingsRequestSubscriptionDataInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSubscriptionSettingsRequestSubscriptionDataInner) SetValue(v UpdateSubscriptionSettingsRequestSubscriptionDataInnerValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


