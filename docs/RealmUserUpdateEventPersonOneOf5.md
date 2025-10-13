# GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user affected by this change.  | [optional] 
**DeliveryEmail** | Pointer to **NullableString** | The new delivery email of the user.  This value can be &#x60;null&#x60; if the affected user changed their &#x60;email_address_visibility&#x60; setting such that you cannot access their real email.  **Changes**: Before Zulip 7.0 (feature level 163), &#x60;null&#x60; was not a possible value for this event as it was only sent to the affected user when their email address was changed.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5

`func NewGetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5() *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5`

NewGetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5 instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5WithDefaults() *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5`

NewGetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5WithDefaults instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDeliveryEmail

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) GetDeliveryEmail() string`

GetDeliveryEmail returns the DeliveryEmail field if non-nil, zero value otherwise.

### GetDeliveryEmailOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) GetDeliveryEmailOk() (*string, bool)`

GetDeliveryEmailOk returns a tuple with the DeliveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEmail

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) SetDeliveryEmail(v string)`

SetDeliveryEmail sets DeliveryEmail field to given value.

### HasDeliveryEmail

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) HasDeliveryEmail() bool`

HasDeliveryEmail returns a boolean if a field has been set.

### SetDeliveryEmailNil

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) SetDeliveryEmailNil(b bool)`

 SetDeliveryEmailNil sets the value for DeliveryEmail to be an explicit nil

### UnsetDeliveryEmail
`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf5) UnsetDeliveryEmail()`

UnsetDeliveryEmail ensures that no value is present for DeliveryEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


