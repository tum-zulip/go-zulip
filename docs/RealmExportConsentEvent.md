# RealmExportConsentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** | The ID of the user whose setting was changed.  | [optional] 
**Consented** | Pointer to **bool** | Whether the user has consented for their private data export.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf58

`func NewGetEvents200ResponseAllOfEventsInnerOneOf58() *RealmExportConsentEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf58 instantiates a new RealmExportConsentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf58WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf58WithDefaults() *RealmExportConsentEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf58WithDefaults instantiates a new RealmExportConsentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmExportConsentEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmExportConsentEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmExportConsentEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmExportConsentEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmExportConsentEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmExportConsentEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmExportConsentEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmExportConsentEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *RealmExportConsentEvent) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RealmExportConsentEvent) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RealmExportConsentEvent) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RealmExportConsentEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetConsented

`func (o *RealmExportConsentEvent) GetConsented() bool`

GetConsented returns the Consented field if non-nil, zero value otherwise.

### GetConsentedOk

`func (o *RealmExportConsentEvent) GetConsentedOk() (*bool, bool)`

GetConsentedOk returns a tuple with the Consented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsented

`func (o *RealmExportConsentEvent) SetConsented(v bool)`

SetConsented sets Consented field to given value.

### HasConsented

`func (o *RealmExportConsentEvent) HasConsented() bool`

HasConsented returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


