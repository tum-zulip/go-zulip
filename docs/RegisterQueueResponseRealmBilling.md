# RegisterQueueResponseRealmBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasPendingSponsorshipRequest** | Pointer to **bool** | Whether there is a pending sponsorship request for the organization. Note that this field will always be &#x60;false&#x60; if the user is not in &#x60;can_manage_billing_group&#x60;.  **Changes**: New in Zulip 10.0 (feature level 363).  | [optional] 

## Methods

### NewRegisterQueueResponseRealmBilling

`func NewRegisterQueueResponseRealmBilling() *RegisterQueueResponseRealmBilling`

NewRegisterQueueResponseRealmBilling instantiates a new RegisterQueueResponseRealmBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseRealmBillingWithDefaults

`func NewRegisterQueueResponseRealmBillingWithDefaults() *RegisterQueueResponseRealmBilling`

NewRegisterQueueResponseRealmBillingWithDefaults instantiates a new RegisterQueueResponseRealmBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasPendingSponsorshipRequest

`func (o *RegisterQueueResponseRealmBilling) GetHasPendingSponsorshipRequest() bool`

GetHasPendingSponsorshipRequest returns the HasPendingSponsorshipRequest field if non-nil, zero value otherwise.

### GetHasPendingSponsorshipRequestOk

`func (o *RegisterQueueResponseRealmBilling) GetHasPendingSponsorshipRequestOk() (*bool, bool)`

GetHasPendingSponsorshipRequestOk returns a tuple with the HasPendingSponsorshipRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingSponsorshipRequest

`func (o *RegisterQueueResponseRealmBilling) SetHasPendingSponsorshipRequest(v bool)`

SetHasPendingSponsorshipRequest sets HasPendingSponsorshipRequest field to given value.

### HasHasPendingSponsorshipRequest

`func (o *RegisterQueueResponseRealmBilling) HasHasPendingSponsorshipRequest() bool`

HasHasPendingSponsorshipRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


