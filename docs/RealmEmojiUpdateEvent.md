# RealmEmojiUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**RealmEmoji** | Pointer to [**map[string]RealmEmoji**](RealmEmoji.md) | An object in which each key describes a realm emoji.  | [optional] 

## Methods

### NewEventEnvelopeOneOf53

`func NewEventEnvelopeOneOf53() *RealmEmojiUpdateEvent`

NewEventEnvelopeOneOf53 instantiates a new RealmEmojiUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf53WithDefaults

`func NewEventEnvelopeOneOf53WithDefaults() *RealmEmojiUpdateEvent`

NewEventEnvelopeOneOf53WithDefaults instantiates a new RealmEmojiUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmEmojiUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmEmojiUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmEmojiUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmEmojiUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmEmojiUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmEmojiUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmEmojiUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmEmojiUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmEmojiUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmEmojiUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmEmojiUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmEmojiUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetRealmEmoji

`func (o *RealmEmojiUpdateEvent) GetRealmEmoji() map[string]RealmEmoji`

GetRealmEmoji returns the RealmEmoji field if non-nil, zero value otherwise.

### GetRealmEmojiOk

`func (o *RealmEmojiUpdateEvent) GetRealmEmojiOk() (*map[string]RealmEmoji, bool)`

GetRealmEmojiOk returns a tuple with the RealmEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmoji

`func (o *RealmEmojiUpdateEvent) SetRealmEmoji(v map[string]RealmEmoji)`

SetRealmEmoji sets RealmEmoji field to given value.

### HasRealmEmoji

`func (o *RealmEmojiUpdateEvent) HasRealmEmoji() bool`

HasRealmEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


