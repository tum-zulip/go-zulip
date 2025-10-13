# RealmEmoji

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID for this emoji, same as the object&#39;s key.  | [optional] 
**Name** | Pointer to **string** | The user-friendly name for this emoji. Users in the organization can use this emoji by writing this name between colons (&#x60;:name :&#x60;).  | [optional] 
**SourceUrl** | Pointer to **string** | The path relative to the organization&#39;s URL where the emoji&#39;s image can be found.  | [optional] 
**StillUrl** | Pointer to **NullableString** | Only non-null when the emoji&#39;s image is animated.  The path relative to the organization&#39;s URL where a still (not animated) version of the emoji can be found. (This is currently always the first frame of the animation).  This is useful for clients to display the emoji in contexts where continuously animating it would be a bad user experience (E.g. because it would be distracting).  **Changes**: New in Zulip 5.0 (added as optional field in feature level 97 and then made mandatory, but nullable, in feature level 113).  | [optional] 
**Deactivated** | Pointer to **bool** | Whether the emoji has been deactivated or not.  | [optional] 
**AuthorId** | Pointer to **NullableInt32** | The user ID of the user who uploaded the custom emoji. Will be &#x60;null&#x60; if the uploader is unknown.  **Changes**: New in Zulip 3.0 (feature level 7). Previously was accessible via an &#x60;author&#x60; object with an &#x60;id&#x60; field.  | [optional] 

## Methods

### NewRealmEmoji

`func NewRealmEmoji() *RealmEmoji`

NewRealmEmoji instantiates a new RealmEmoji object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmEmojiWithDefaults

`func NewRealmEmojiWithDefaults() *RealmEmoji`

NewRealmEmojiWithDefaults instantiates a new RealmEmoji object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmEmoji) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmEmoji) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmEmoji) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RealmEmoji) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RealmEmoji) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealmEmoji) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealmEmoji) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RealmEmoji) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceUrl

`func (o *RealmEmoji) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RealmEmoji) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RealmEmoji) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RealmEmoji) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetStillUrl

`func (o *RealmEmoji) GetStillUrl() string`

GetStillUrl returns the StillUrl field if non-nil, zero value otherwise.

### GetStillUrlOk

`func (o *RealmEmoji) GetStillUrlOk() (*string, bool)`

GetStillUrlOk returns a tuple with the StillUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStillUrl

`func (o *RealmEmoji) SetStillUrl(v string)`

SetStillUrl sets StillUrl field to given value.

### HasStillUrl

`func (o *RealmEmoji) HasStillUrl() bool`

HasStillUrl returns a boolean if a field has been set.

### SetStillUrlNil

`func (o *RealmEmoji) SetStillUrlNil(b bool)`

 SetStillUrlNil sets the value for StillUrl to be an explicit nil

### UnsetStillUrl
`func (o *RealmEmoji) UnsetStillUrl()`

UnsetStillUrl ensures that no value is present for StillUrl, not even an explicit nil
### GetDeactivated

`func (o *RealmEmoji) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *RealmEmoji) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *RealmEmoji) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *RealmEmoji) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.

### GetAuthorId

`func (o *RealmEmoji) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *RealmEmoji) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *RealmEmoji) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *RealmEmoji) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### SetAuthorIdNil

`func (o *RealmEmoji) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *RealmEmoji) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


