# RealmAuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean describing whether the authentication method (i.e. its key) is enabled in this organization.  | [optional] 
**Available** | Pointer to **bool** | Boolean describing whether the authentication method is available for use. If false, the organization is not eligible to enable the authentication method.  | [optional] 
**UnavailableReason** | Pointer to **string** | Reason why the authentication method is unavailable. This field is optional and is only present when &#39;available&#39; is false.  | [optional] 

## Methods

### NewRealmAuthenticationMethod

`func NewRealmAuthenticationMethod() *RealmAuthenticationMethod`

NewRealmAuthenticationMethod instantiates a new RealmAuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmAuthenticationMethodWithDefaults

`func NewRealmAuthenticationMethodWithDefaults() *RealmAuthenticationMethod`

NewRealmAuthenticationMethodWithDefaults instantiates a new RealmAuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RealmAuthenticationMethod) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RealmAuthenticationMethod) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RealmAuthenticationMethod) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RealmAuthenticationMethod) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAvailable

`func (o *RealmAuthenticationMethod) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *RealmAuthenticationMethod) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *RealmAuthenticationMethod) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *RealmAuthenticationMethod) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetUnavailableReason

`func (o *RealmAuthenticationMethod) GetUnavailableReason() string`

GetUnavailableReason returns the UnavailableReason field if non-nil, zero value otherwise.

### GetUnavailableReasonOk

`func (o *RealmAuthenticationMethod) GetUnavailableReasonOk() (*string, bool)`

GetUnavailableReasonOk returns a tuple with the UnavailableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableReason

`func (o *RealmAuthenticationMethod) SetUnavailableReason(v string)`

SetUnavailableReason sets UnavailableReason field to given value.

### HasUnavailableReason

`func (o *RealmAuthenticationMethod) HasUnavailableReason() bool`

HasUnavailableReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


