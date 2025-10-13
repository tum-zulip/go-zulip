# RealmFiltersEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RealmFilters** | Pointer to [**[][]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner**]([]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner.md) | An array of tuples, where each tuple described a linkifier. The first element of the tuple was a string regex pattern which represented the pattern to be linkified on matching, for example &#x60;\&quot;#(?P&lt;id&gt;[123])\&quot;&#x60;. The second element was the URL format string that the pattern should be linkified with. A URL format string for the above example would be &#x60;\&quot;https://realm.com/my_realm_filter/%(id)s\&quot;&#x60;. And the third element was the ID of the realm filter.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf51

`func NewGetEvents200ResponseAllOfEventsInnerOneOf51() *RealmFiltersEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf51 instantiates a new RealmFiltersEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf51WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf51WithDefaults() *RealmFiltersEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf51WithDefaults instantiates a new RealmFiltersEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmFiltersEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmFiltersEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmFiltersEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmFiltersEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmFiltersEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmFiltersEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmFiltersEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmFiltersEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRealmFilters

`func (o *RealmFiltersEvent) GetRealmFilters() [][]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner`

GetRealmFilters returns the RealmFilters field if non-nil, zero value otherwise.

### GetRealmFiltersOk

`func (o *RealmFiltersEvent) GetRealmFiltersOk() (*[][]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner, bool)`

GetRealmFiltersOk returns a tuple with the RealmFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmFilters

`func (o *RealmFiltersEvent) SetRealmFilters(v [][]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner)`

SetRealmFilters sets RealmFilters field to given value.

### HasRealmFilters

`func (o *RealmFiltersEvent) HasRealmFilters() bool`

HasRealmFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


