# Invite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the invitation.  Note that email invitations and reusable invitation links are stored in different database tables on the server, so each ID is guaranteed to be unique in combination with the boolean value of &#x60;is_multiuse&#x60;, e.g. there can only be one invitation with &#x60;id: 1&#x60; and &#x60;is_multiuse: true&#x60;.  | [optional] 
**InvitedByUserId** | Pointer to **int32** | The [user ID](/api/get-user) of the user who created the invitation.  **Changes**: New in Zulip 3.0 (feature level 22), replacing the &#x60;ref&#x60; field which contained the Zulip display email address of the user who created the invitation.  | [optional] 
**Invited** | Pointer to **int32** | The UNIX timestamp for when the invitation was created, in UTC seconds.  | [optional] 
**ExpiryDate** | Pointer to **NullableInt32** | The UNIX timestamp for when the invitation will expire, in UTC seconds. If &#x60;null&#x60;, the invitation never expires.  | [optional] 
**InvitedAs** | Pointer to **int32** | The [organization-level role](/api/roles-and-permissions) of the user that is created when the invitation is accepted. Possible values are:  - 100 &#x3D; Organization owner - 200 &#x3D; Organization administrator - 300 &#x3D; Organization moderator - 400 &#x3D; Member - 600 &#x3D; Guest  | [optional] 
**Email** | Pointer to **string** | The email address the invitation was sent to. This will not be present when &#x60;is_multiuse&#x60; is &#x60;true&#x60; (i.e. the invitation is a reusable invitation link).  | [optional] 
**NotifyReferrerOnJoin** | Pointer to **bool** | A boolean indicating whether the referrer has opted to receive a direct message from [notification bot](/help/configure-automated-notices) when a user account is created using this invitation.  **Changes**: New in Zulip 9.0 (feature level 267). Previously, referrers always received such direct messages.  | [optional] 
**LinkUrl** | Pointer to **string** | The URL of the reusable invitation link. This will not be present when &#x60;is_multiuse&#x60; is &#x60;false&#x60; (i.e. the invitation is an email invitation).  | [optional] 
**IsMultiuse** | Pointer to **bool** | A boolean specifying whether the [invitation](/help/invite-new-users) is a reusable invitation link or an email invitation.  | [optional] 

## Methods

### NewInvite

`func NewInvite() *Invite`

NewInvite instantiates a new Invite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteWithDefaults

`func NewInviteWithDefaults() *Invite`

NewInviteWithDefaults instantiates a new Invite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invite) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invite) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invite) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Invite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvitedByUserId

`func (o *Invite) GetInvitedByUserId() int32`

GetInvitedByUserId returns the InvitedByUserId field if non-nil, zero value otherwise.

### GetInvitedByUserIdOk

`func (o *Invite) GetInvitedByUserIdOk() (*int32, bool)`

GetInvitedByUserIdOk returns a tuple with the InvitedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedByUserId

`func (o *Invite) SetInvitedByUserId(v int32)`

SetInvitedByUserId sets InvitedByUserId field to given value.

### HasInvitedByUserId

`func (o *Invite) HasInvitedByUserId() bool`

HasInvitedByUserId returns a boolean if a field has been set.

### GetInvited

`func (o *Invite) GetInvited() int32`

GetInvited returns the Invited field if non-nil, zero value otherwise.

### GetInvitedOk

`func (o *Invite) GetInvitedOk() (*int32, bool)`

GetInvitedOk returns a tuple with the Invited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvited

`func (o *Invite) SetInvited(v int32)`

SetInvited sets Invited field to given value.

### HasInvited

`func (o *Invite) HasInvited() bool`

HasInvited returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Invite) GetExpiryDate() int32`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Invite) GetExpiryDateOk() (*int32, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Invite) SetExpiryDate(v int32)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Invite) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *Invite) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *Invite) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetInvitedAs

`func (o *Invite) GetInvitedAs() int32`

GetInvitedAs returns the InvitedAs field if non-nil, zero value otherwise.

### GetInvitedAsOk

`func (o *Invite) GetInvitedAsOk() (*int32, bool)`

GetInvitedAsOk returns a tuple with the InvitedAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedAs

`func (o *Invite) SetInvitedAs(v int32)`

SetInvitedAs sets InvitedAs field to given value.

### HasInvitedAs

`func (o *Invite) HasInvitedAs() bool`

HasInvitedAs returns a boolean if a field has been set.

### GetEmail

`func (o *Invite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Invite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Invite) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Invite) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNotifyReferrerOnJoin

`func (o *Invite) GetNotifyReferrerOnJoin() bool`

GetNotifyReferrerOnJoin returns the NotifyReferrerOnJoin field if non-nil, zero value otherwise.

### GetNotifyReferrerOnJoinOk

`func (o *Invite) GetNotifyReferrerOnJoinOk() (*bool, bool)`

GetNotifyReferrerOnJoinOk returns a tuple with the NotifyReferrerOnJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyReferrerOnJoin

`func (o *Invite) SetNotifyReferrerOnJoin(v bool)`

SetNotifyReferrerOnJoin sets NotifyReferrerOnJoin field to given value.

### HasNotifyReferrerOnJoin

`func (o *Invite) HasNotifyReferrerOnJoin() bool`

HasNotifyReferrerOnJoin returns a boolean if a field has been set.

### GetLinkUrl

`func (o *Invite) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *Invite) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *Invite) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *Invite) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetIsMultiuse

`func (o *Invite) GetIsMultiuse() bool`

GetIsMultiuse returns the IsMultiuse field if non-nil, zero value otherwise.

### GetIsMultiuseOk

`func (o *Invite) GetIsMultiuseOk() (*bool, bool)`

GetIsMultiuseOk returns a tuple with the IsMultiuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiuse

`func (o *Invite) SetIsMultiuse(v bool)`

SetIsMultiuse sets IsMultiuse field to given value.

### HasIsMultiuse

`func (o *Invite) HasIsMultiuse() bool`

HasIsMultiuse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


