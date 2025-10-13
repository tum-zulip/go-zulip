# RealmLinkifiersEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RealmLinkifiers** | Pointer to [**[]EventEnvelopeOneOf50RealmLinkifiersInner**](EventEnvelopeOneOf50RealmLinkifiersInner.md) | An ordered array of dictionaries where each dictionary contains details about a single linkifier.  Clients should always process linkifiers in the order given; this is important if the realm has linkifiers with overlapping patterns. The order can be modified using [&#x60;PATCH /realm/linkifiers&#x60;](zulip.com/api/reorder-linkifiers.  | [optional] 

## Methods

### NewEventEnvelopeOneOf50

`func NewEventEnvelopeOneOf50() *RealmLinkifiersEvent`

NewEventEnvelopeOneOf50 instantiates a new RealmLinkifiersEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf50WithDefaults

`func NewEventEnvelopeOneOf50WithDefaults() *RealmLinkifiersEvent`

NewEventEnvelopeOneOf50WithDefaults instantiates a new RealmLinkifiersEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmLinkifiersEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmLinkifiersEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmLinkifiersEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmLinkifiersEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmLinkifiersEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmLinkifiersEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmLinkifiersEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmLinkifiersEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRealmLinkifiers

`func (o *RealmLinkifiersEvent) GetRealmLinkifiers() []EventEnvelopeOneOf50RealmLinkifiersInner`

GetRealmLinkifiers returns the RealmLinkifiers field if non-nil, zero value otherwise.

### GetRealmLinkifiersOk

`func (o *RealmLinkifiersEvent) GetRealmLinkifiersOk() (*[]EventEnvelopeOneOf50RealmLinkifiersInner, bool)`

GetRealmLinkifiersOk returns a tuple with the RealmLinkifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmLinkifiers

`func (o *RealmLinkifiersEvent) SetRealmLinkifiers(v []EventEnvelopeOneOf50RealmLinkifiersInner)`

SetRealmLinkifiers sets RealmLinkifiers field to given value.

### HasRealmLinkifiers

`func (o *RealmLinkifiersEvent) HasRealmLinkifiers() bool`

HasRealmLinkifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


