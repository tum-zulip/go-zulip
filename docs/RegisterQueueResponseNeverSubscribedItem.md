# RegisterQueueResponseNeverSubscribedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamId** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **interface{}** |  | [optional] 
**IsArchived** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **interface{}** |  | [optional] 
**DateCreated** | Pointer to **interface{}** |  | [optional] 
**CreatorId** | Pointer to **interface{}** |  | [optional] 
**InviteOnly** | Pointer to **interface{}** |  | [optional] 
**RenderedDescription** | Pointer to **interface{}** |  | [optional] 
**IsWebPublic** | Pointer to **interface{}** |  | [optional] 
**StreamPostPolicy** | Pointer to **interface{}** |  | [optional] 
**MessageRetentionDays** | Pointer to **interface{}** |  | [optional] 
**HistoryPublicToSubscribers** | Pointer to **interface{}** |  | [optional] 
**TopicsPolicy** | Pointer to **interface{}** |  | [optional] 
**FirstMessageId** | Pointer to **interface{}** |  | [optional] 
**FolderId** | Pointer to **interface{}** |  | [optional] 
**IsRecentlyActive** | Pointer to **interface{}** |  | [optional] 
**IsAnnouncementOnly** | Pointer to **interface{}** |  | [optional] 
**CanAddSubscribersGroup** | Pointer to **interface{}** |  | [optional] 
**CanRemoveSubscribersGroup** | Pointer to **interface{}** |  | [optional] 
**CanAdministerChannelGroup** | Pointer to **interface{}** |  | [optional] 
**CanDeleteAnyMessageGroup** | Pointer to **interface{}** |  | [optional] 
**CanDeleteOwnMessageGroup** | Pointer to **interface{}** |  | [optional] 
**CanMoveMessagesOutOfChannelGroup** | Pointer to **interface{}** |  | [optional] 
**CanMoveMessagesWithinChannelGroup** | Pointer to **interface{}** |  | [optional] 
**CanSendMessageGroup** | Pointer to **interface{}** |  | [optional] 
**CanSubscribeGroup** | Pointer to **interface{}** |  | [optional] 
**CanResolveTopicsGroup** | Pointer to **interface{}** |  | [optional] 
**SubscriberCount** | Pointer to **interface{}** |  | [optional] 
**StreamWeeklyTraffic** | Pointer to **NullableInt32** | The average number of messages sent to the channel per week, as estimated based on recent weeks, rounded to the nearest integer.  If &#x60;null&#x60;, the channel was recently created and there is insufficient data to estimate the average traffic.  | [optional] 
**Subscribers** | Pointer to **[]int32** | A list of user IDs of users who are subscribed to the channel. Included only if &#x60;include_subscribers&#x60; is &#x60;true&#x60;.  If a user is not allowed to know the subscribers for a channel, we will send an empty array. API authors should use other data to determine whether users like guest users are forbidden to know the subscribers.  | [optional] 
**PartialSubscribers** | Pointer to **[]int32** | If [&#x60;include_subscribers&#x3D;\&quot;partial\&quot;&#x60;](/api/get-subscriptions#parameter-include_subscribers) was requested, the server may, at its discretion, send a &#x60;partial_subscribers&#x60; list rather than a &#x60;subscribers&#x60; list for channels with a large number of subscribers.  The &#x60;partial_subscribers&#x60; list contains an arbitrary subset of the channel&#39;s subscribers that is guaranteed to include all bot user subscribers as well as all users who have been active in the last 14 days, but otherwise can be chosen arbitrarily by the server.  If a user is not allowed to know the subscribers for a channel, we will send an empty array. API authors should use other data to determine whether users like guest users are forbidden to know the subscribers.  **Changes**: New in Zulip 11.0 (feature level 412).  | [optional] 

## Methods

### NewRegisterQueueResponseNeverSubscribedItem

`func NewRegisterQueueResponseNeverSubscribedItem() *RegisterQueueResponseNeverSubscribedItem`

NewRegisterQueueResponseNeverSubscribedItem instantiates a new RegisterQueueResponseNeverSubscribedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseNeverSubscribedItemWithDefaults

`func NewRegisterQueueResponseNeverSubscribedItemWithDefaults() *RegisterQueueResponseNeverSubscribedItem`

NewRegisterQueueResponseNeverSubscribedItemWithDefaults instantiates a new RegisterQueueResponseNeverSubscribedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamId

`func (o *RegisterQueueResponseNeverSubscribedItem) GetStreamId() interface{}`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetStreamIdOk() (*interface{}, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *RegisterQueueResponseNeverSubscribedItem) SetStreamId(v interface{})`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *RegisterQueueResponseNeverSubscribedItem) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### SetStreamIdNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetStreamIdNil(b bool)`

 SetStreamIdNil sets the value for StreamId to be an explicit nil

### UnsetStreamId
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetStreamId()`

UnsetStreamId ensures that no value is present for StreamId, not even an explicit nil
### GetName

`func (o *RegisterQueueResponseNeverSubscribedItem) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterQueueResponseNeverSubscribedItem) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *RegisterQueueResponseNeverSubscribedItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsArchived

`func (o *RegisterQueueResponseNeverSubscribedItem) GetIsArchived() interface{}`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetIsArchivedOk() (*interface{}, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *RegisterQueueResponseNeverSubscribedItem) SetIsArchived(v interface{})`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *RegisterQueueResponseNeverSubscribedItem) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchivedNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil
### GetDescription

`func (o *RegisterQueueResponseNeverSubscribedItem) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegisterQueueResponseNeverSubscribedItem) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegisterQueueResponseNeverSubscribedItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDateCreated

`func (o *RegisterQueueResponseNeverSubscribedItem) GetDateCreated() interface{}`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetDateCreatedOk() (*interface{}, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *RegisterQueueResponseNeverSubscribedItem) SetDateCreated(v interface{})`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *RegisterQueueResponseNeverSubscribedItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetCreatorId

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCreatorId() interface{}`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCreatorIdOk() (*interface{}, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCreatorId(v interface{})`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetInviteOnly

`func (o *RegisterQueueResponseNeverSubscribedItem) GetInviteOnly() interface{}`

GetInviteOnly returns the InviteOnly field if non-nil, zero value otherwise.

### GetInviteOnlyOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetInviteOnlyOk() (*interface{}, bool)`

GetInviteOnlyOk returns a tuple with the InviteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteOnly

`func (o *RegisterQueueResponseNeverSubscribedItem) SetInviteOnly(v interface{})`

SetInviteOnly sets InviteOnly field to given value.

### HasInviteOnly

`func (o *RegisterQueueResponseNeverSubscribedItem) HasInviteOnly() bool`

HasInviteOnly returns a boolean if a field has been set.

### SetInviteOnlyNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetInviteOnlyNil(b bool)`

 SetInviteOnlyNil sets the value for InviteOnly to be an explicit nil

### UnsetInviteOnly
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetInviteOnly()`

UnsetInviteOnly ensures that no value is present for InviteOnly, not even an explicit nil
### GetRenderedDescription

`func (o *RegisterQueueResponseNeverSubscribedItem) GetRenderedDescription() interface{}`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetRenderedDescriptionOk() (*interface{}, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *RegisterQueueResponseNeverSubscribedItem) SetRenderedDescription(v interface{})`

SetRenderedDescription sets RenderedDescription field to given value.

### HasRenderedDescription

`func (o *RegisterQueueResponseNeverSubscribedItem) HasRenderedDescription() bool`

HasRenderedDescription returns a boolean if a field has been set.

### SetRenderedDescriptionNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetRenderedDescriptionNil(b bool)`

 SetRenderedDescriptionNil sets the value for RenderedDescription to be an explicit nil

### UnsetRenderedDescription
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetRenderedDescription()`

UnsetRenderedDescription ensures that no value is present for RenderedDescription, not even an explicit nil
### GetIsWebPublic

`func (o *RegisterQueueResponseNeverSubscribedItem) GetIsWebPublic() interface{}`

GetIsWebPublic returns the IsWebPublic field if non-nil, zero value otherwise.

### GetIsWebPublicOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetIsWebPublicOk() (*interface{}, bool)`

GetIsWebPublicOk returns a tuple with the IsWebPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebPublic

`func (o *RegisterQueueResponseNeverSubscribedItem) SetIsWebPublic(v interface{})`

SetIsWebPublic sets IsWebPublic field to given value.

### HasIsWebPublic

`func (o *RegisterQueueResponseNeverSubscribedItem) HasIsWebPublic() bool`

HasIsWebPublic returns a boolean if a field has been set.

### SetIsWebPublicNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetIsWebPublicNil(b bool)`

 SetIsWebPublicNil sets the value for IsWebPublic to be an explicit nil

### UnsetIsWebPublic
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetIsWebPublic()`

UnsetIsWebPublic ensures that no value is present for IsWebPublic, not even an explicit nil
### GetStreamPostPolicy

`func (o *RegisterQueueResponseNeverSubscribedItem) GetStreamPostPolicy() interface{}`

GetStreamPostPolicy returns the StreamPostPolicy field if non-nil, zero value otherwise.

### GetStreamPostPolicyOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetStreamPostPolicyOk() (*interface{}, bool)`

GetStreamPostPolicyOk returns a tuple with the StreamPostPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamPostPolicy

`func (o *RegisterQueueResponseNeverSubscribedItem) SetStreamPostPolicy(v interface{})`

SetStreamPostPolicy sets StreamPostPolicy field to given value.

### HasStreamPostPolicy

`func (o *RegisterQueueResponseNeverSubscribedItem) HasStreamPostPolicy() bool`

HasStreamPostPolicy returns a boolean if a field has been set.

### SetStreamPostPolicyNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetStreamPostPolicyNil(b bool)`

 SetStreamPostPolicyNil sets the value for StreamPostPolicy to be an explicit nil

### UnsetStreamPostPolicy
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetStreamPostPolicy()`

UnsetStreamPostPolicy ensures that no value is present for StreamPostPolicy, not even an explicit nil
### GetMessageRetentionDays

`func (o *RegisterQueueResponseNeverSubscribedItem) GetMessageRetentionDays() interface{}`

GetMessageRetentionDays returns the MessageRetentionDays field if non-nil, zero value otherwise.

### GetMessageRetentionDaysOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetMessageRetentionDaysOk() (*interface{}, bool)`

GetMessageRetentionDaysOk returns a tuple with the MessageRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRetentionDays

`func (o *RegisterQueueResponseNeverSubscribedItem) SetMessageRetentionDays(v interface{})`

SetMessageRetentionDays sets MessageRetentionDays field to given value.

### HasMessageRetentionDays

`func (o *RegisterQueueResponseNeverSubscribedItem) HasMessageRetentionDays() bool`

HasMessageRetentionDays returns a boolean if a field has been set.

### SetMessageRetentionDaysNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetMessageRetentionDaysNil(b bool)`

 SetMessageRetentionDaysNil sets the value for MessageRetentionDays to be an explicit nil

### UnsetMessageRetentionDays
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetMessageRetentionDays()`

UnsetMessageRetentionDays ensures that no value is present for MessageRetentionDays, not even an explicit nil
### GetHistoryPublicToSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) GetHistoryPublicToSubscribers() interface{}`

GetHistoryPublicToSubscribers returns the HistoryPublicToSubscribers field if non-nil, zero value otherwise.

### GetHistoryPublicToSubscribersOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetHistoryPublicToSubscribersOk() (*interface{}, bool)`

GetHistoryPublicToSubscribersOk returns a tuple with the HistoryPublicToSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPublicToSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) SetHistoryPublicToSubscribers(v interface{})`

SetHistoryPublicToSubscribers sets HistoryPublicToSubscribers field to given value.

### HasHistoryPublicToSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) HasHistoryPublicToSubscribers() bool`

HasHistoryPublicToSubscribers returns a boolean if a field has been set.

### SetHistoryPublicToSubscribersNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetHistoryPublicToSubscribersNil(b bool)`

 SetHistoryPublicToSubscribersNil sets the value for HistoryPublicToSubscribers to be an explicit nil

### UnsetHistoryPublicToSubscribers
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetHistoryPublicToSubscribers()`

UnsetHistoryPublicToSubscribers ensures that no value is present for HistoryPublicToSubscribers, not even an explicit nil
### GetTopicsPolicy

`func (o *RegisterQueueResponseNeverSubscribedItem) GetTopicsPolicy() interface{}`

GetTopicsPolicy returns the TopicsPolicy field if non-nil, zero value otherwise.

### GetTopicsPolicyOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetTopicsPolicyOk() (*interface{}, bool)`

GetTopicsPolicyOk returns a tuple with the TopicsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsPolicy

`func (o *RegisterQueueResponseNeverSubscribedItem) SetTopicsPolicy(v interface{})`

SetTopicsPolicy sets TopicsPolicy field to given value.

### HasTopicsPolicy

`func (o *RegisterQueueResponseNeverSubscribedItem) HasTopicsPolicy() bool`

HasTopicsPolicy returns a boolean if a field has been set.

### SetTopicsPolicyNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetTopicsPolicyNil(b bool)`

 SetTopicsPolicyNil sets the value for TopicsPolicy to be an explicit nil

### UnsetTopicsPolicy
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetTopicsPolicy()`

UnsetTopicsPolicy ensures that no value is present for TopicsPolicy, not even an explicit nil
### GetFirstMessageId

`func (o *RegisterQueueResponseNeverSubscribedItem) GetFirstMessageId() interface{}`

GetFirstMessageId returns the FirstMessageId field if non-nil, zero value otherwise.

### GetFirstMessageIdOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetFirstMessageIdOk() (*interface{}, bool)`

GetFirstMessageIdOk returns a tuple with the FirstMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessageId

`func (o *RegisterQueueResponseNeverSubscribedItem) SetFirstMessageId(v interface{})`

SetFirstMessageId sets FirstMessageId field to given value.

### HasFirstMessageId

`func (o *RegisterQueueResponseNeverSubscribedItem) HasFirstMessageId() bool`

HasFirstMessageId returns a boolean if a field has been set.

### SetFirstMessageIdNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetFirstMessageIdNil(b bool)`

 SetFirstMessageIdNil sets the value for FirstMessageId to be an explicit nil

### UnsetFirstMessageId
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetFirstMessageId()`

UnsetFirstMessageId ensures that no value is present for FirstMessageId, not even an explicit nil
### GetFolderId

`func (o *RegisterQueueResponseNeverSubscribedItem) GetFolderId() interface{}`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetFolderIdOk() (*interface{}, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RegisterQueueResponseNeverSubscribedItem) SetFolderId(v interface{})`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RegisterQueueResponseNeverSubscribedItem) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetIsRecentlyActive

`func (o *RegisterQueueResponseNeverSubscribedItem) GetIsRecentlyActive() interface{}`

GetIsRecentlyActive returns the IsRecentlyActive field if non-nil, zero value otherwise.

### GetIsRecentlyActiveOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetIsRecentlyActiveOk() (*interface{}, bool)`

GetIsRecentlyActiveOk returns a tuple with the IsRecentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecentlyActive

`func (o *RegisterQueueResponseNeverSubscribedItem) SetIsRecentlyActive(v interface{})`

SetIsRecentlyActive sets IsRecentlyActive field to given value.

### HasIsRecentlyActive

`func (o *RegisterQueueResponseNeverSubscribedItem) HasIsRecentlyActive() bool`

HasIsRecentlyActive returns a boolean if a field has been set.

### SetIsRecentlyActiveNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetIsRecentlyActiveNil(b bool)`

 SetIsRecentlyActiveNil sets the value for IsRecentlyActive to be an explicit nil

### UnsetIsRecentlyActive
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetIsRecentlyActive()`

UnsetIsRecentlyActive ensures that no value is present for IsRecentlyActive, not even an explicit nil
### GetIsAnnouncementOnly

`func (o *RegisterQueueResponseNeverSubscribedItem) GetIsAnnouncementOnly() interface{}`

GetIsAnnouncementOnly returns the IsAnnouncementOnly field if non-nil, zero value otherwise.

### GetIsAnnouncementOnlyOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetIsAnnouncementOnlyOk() (*interface{}, bool)`

GetIsAnnouncementOnlyOk returns a tuple with the IsAnnouncementOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnouncementOnly

`func (o *RegisterQueueResponseNeverSubscribedItem) SetIsAnnouncementOnly(v interface{})`

SetIsAnnouncementOnly sets IsAnnouncementOnly field to given value.

### HasIsAnnouncementOnly

`func (o *RegisterQueueResponseNeverSubscribedItem) HasIsAnnouncementOnly() bool`

HasIsAnnouncementOnly returns a boolean if a field has been set.

### SetIsAnnouncementOnlyNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetIsAnnouncementOnlyNil(b bool)`

 SetIsAnnouncementOnlyNil sets the value for IsAnnouncementOnly to be an explicit nil

### UnsetIsAnnouncementOnly
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetIsAnnouncementOnly()`

UnsetIsAnnouncementOnly ensures that no value is present for IsAnnouncementOnly, not even an explicit nil
### GetCanAddSubscribersGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanAddSubscribersGroup() interface{}`

GetCanAddSubscribersGroup returns the CanAddSubscribersGroup field if non-nil, zero value otherwise.

### GetCanAddSubscribersGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanAddSubscribersGroupOk() (*interface{}, bool)`

GetCanAddSubscribersGroupOk returns a tuple with the CanAddSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddSubscribersGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanAddSubscribersGroup(v interface{})`

SetCanAddSubscribersGroup sets CanAddSubscribersGroup field to given value.

### HasCanAddSubscribersGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanAddSubscribersGroup() bool`

HasCanAddSubscribersGroup returns a boolean if a field has been set.

### SetCanAddSubscribersGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanAddSubscribersGroupNil(b bool)`

 SetCanAddSubscribersGroupNil sets the value for CanAddSubscribersGroup to be an explicit nil

### UnsetCanAddSubscribersGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanAddSubscribersGroup()`

UnsetCanAddSubscribersGroup ensures that no value is present for CanAddSubscribersGroup, not even an explicit nil
### GetCanRemoveSubscribersGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanRemoveSubscribersGroup() interface{}`

GetCanRemoveSubscribersGroup returns the CanRemoveSubscribersGroup field if non-nil, zero value otherwise.

### GetCanRemoveSubscribersGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanRemoveSubscribersGroupOk() (*interface{}, bool)`

GetCanRemoveSubscribersGroupOk returns a tuple with the CanRemoveSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemoveSubscribersGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanRemoveSubscribersGroup(v interface{})`

SetCanRemoveSubscribersGroup sets CanRemoveSubscribersGroup field to given value.

### HasCanRemoveSubscribersGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanRemoveSubscribersGroup() bool`

HasCanRemoveSubscribersGroup returns a boolean if a field has been set.

### SetCanRemoveSubscribersGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanRemoveSubscribersGroupNil(b bool)`

 SetCanRemoveSubscribersGroupNil sets the value for CanRemoveSubscribersGroup to be an explicit nil

### UnsetCanRemoveSubscribersGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanRemoveSubscribersGroup()`

UnsetCanRemoveSubscribersGroup ensures that no value is present for CanRemoveSubscribersGroup, not even an explicit nil
### GetCanAdministerChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanAdministerChannelGroup() interface{}`

GetCanAdministerChannelGroup returns the CanAdministerChannelGroup field if non-nil, zero value otherwise.

### GetCanAdministerChannelGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanAdministerChannelGroupOk() (*interface{}, bool)`

GetCanAdministerChannelGroupOk returns a tuple with the CanAdministerChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdministerChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanAdministerChannelGroup(v interface{})`

SetCanAdministerChannelGroup sets CanAdministerChannelGroup field to given value.

### HasCanAdministerChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanAdministerChannelGroup() bool`

HasCanAdministerChannelGroup returns a boolean if a field has been set.

### SetCanAdministerChannelGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanAdministerChannelGroupNil(b bool)`

 SetCanAdministerChannelGroupNil sets the value for CanAdministerChannelGroup to be an explicit nil

### UnsetCanAdministerChannelGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanAdministerChannelGroup()`

UnsetCanAdministerChannelGroup ensures that no value is present for CanAdministerChannelGroup, not even an explicit nil
### GetCanDeleteAnyMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanDeleteAnyMessageGroup() interface{}`

GetCanDeleteAnyMessageGroup returns the CanDeleteAnyMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteAnyMessageGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanDeleteAnyMessageGroupOk() (*interface{}, bool)`

GetCanDeleteAnyMessageGroupOk returns a tuple with the CanDeleteAnyMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteAnyMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanDeleteAnyMessageGroup(v interface{})`

SetCanDeleteAnyMessageGroup sets CanDeleteAnyMessageGroup field to given value.

### HasCanDeleteAnyMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanDeleteAnyMessageGroup() bool`

HasCanDeleteAnyMessageGroup returns a boolean if a field has been set.

### SetCanDeleteAnyMessageGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanDeleteAnyMessageGroupNil(b bool)`

 SetCanDeleteAnyMessageGroupNil sets the value for CanDeleteAnyMessageGroup to be an explicit nil

### UnsetCanDeleteAnyMessageGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanDeleteAnyMessageGroup()`

UnsetCanDeleteAnyMessageGroup ensures that no value is present for CanDeleteAnyMessageGroup, not even an explicit nil
### GetCanDeleteOwnMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanDeleteOwnMessageGroup() interface{}`

GetCanDeleteOwnMessageGroup returns the CanDeleteOwnMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteOwnMessageGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanDeleteOwnMessageGroupOk() (*interface{}, bool)`

GetCanDeleteOwnMessageGroupOk returns a tuple with the CanDeleteOwnMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteOwnMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanDeleteOwnMessageGroup(v interface{})`

SetCanDeleteOwnMessageGroup sets CanDeleteOwnMessageGroup field to given value.

### HasCanDeleteOwnMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanDeleteOwnMessageGroup() bool`

HasCanDeleteOwnMessageGroup returns a boolean if a field has been set.

### SetCanDeleteOwnMessageGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanDeleteOwnMessageGroupNil(b bool)`

 SetCanDeleteOwnMessageGroupNil sets the value for CanDeleteOwnMessageGroup to be an explicit nil

### UnsetCanDeleteOwnMessageGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanDeleteOwnMessageGroup()`

UnsetCanDeleteOwnMessageGroup ensures that no value is present for CanDeleteOwnMessageGroup, not even an explicit nil
### GetCanMoveMessagesOutOfChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanMoveMessagesOutOfChannelGroup() interface{}`

GetCanMoveMessagesOutOfChannelGroup returns the CanMoveMessagesOutOfChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesOutOfChannelGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanMoveMessagesOutOfChannelGroupOk() (*interface{}, bool)`

GetCanMoveMessagesOutOfChannelGroupOk returns a tuple with the CanMoveMessagesOutOfChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesOutOfChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanMoveMessagesOutOfChannelGroup(v interface{})`

SetCanMoveMessagesOutOfChannelGroup sets CanMoveMessagesOutOfChannelGroup field to given value.

### HasCanMoveMessagesOutOfChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanMoveMessagesOutOfChannelGroup() bool`

HasCanMoveMessagesOutOfChannelGroup returns a boolean if a field has been set.

### SetCanMoveMessagesOutOfChannelGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanMoveMessagesOutOfChannelGroupNil(b bool)`

 SetCanMoveMessagesOutOfChannelGroupNil sets the value for CanMoveMessagesOutOfChannelGroup to be an explicit nil

### UnsetCanMoveMessagesOutOfChannelGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanMoveMessagesOutOfChannelGroup()`

UnsetCanMoveMessagesOutOfChannelGroup ensures that no value is present for CanMoveMessagesOutOfChannelGroup, not even an explicit nil
### GetCanMoveMessagesWithinChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanMoveMessagesWithinChannelGroup() interface{}`

GetCanMoveMessagesWithinChannelGroup returns the CanMoveMessagesWithinChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesWithinChannelGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanMoveMessagesWithinChannelGroupOk() (*interface{}, bool)`

GetCanMoveMessagesWithinChannelGroupOk returns a tuple with the CanMoveMessagesWithinChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesWithinChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanMoveMessagesWithinChannelGroup(v interface{})`

SetCanMoveMessagesWithinChannelGroup sets CanMoveMessagesWithinChannelGroup field to given value.

### HasCanMoveMessagesWithinChannelGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanMoveMessagesWithinChannelGroup() bool`

HasCanMoveMessagesWithinChannelGroup returns a boolean if a field has been set.

### SetCanMoveMessagesWithinChannelGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanMoveMessagesWithinChannelGroupNil(b bool)`

 SetCanMoveMessagesWithinChannelGroupNil sets the value for CanMoveMessagesWithinChannelGroup to be an explicit nil

### UnsetCanMoveMessagesWithinChannelGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanMoveMessagesWithinChannelGroup()`

UnsetCanMoveMessagesWithinChannelGroup ensures that no value is present for CanMoveMessagesWithinChannelGroup, not even an explicit nil
### GetCanSendMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanSendMessageGroup() interface{}`

GetCanSendMessageGroup returns the CanSendMessageGroup field if non-nil, zero value otherwise.

### GetCanSendMessageGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanSendMessageGroupOk() (*interface{}, bool)`

GetCanSendMessageGroupOk returns a tuple with the CanSendMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanSendMessageGroup(v interface{})`

SetCanSendMessageGroup sets CanSendMessageGroup field to given value.

### HasCanSendMessageGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanSendMessageGroup() bool`

HasCanSendMessageGroup returns a boolean if a field has been set.

### SetCanSendMessageGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanSendMessageGroupNil(b bool)`

 SetCanSendMessageGroupNil sets the value for CanSendMessageGroup to be an explicit nil

### UnsetCanSendMessageGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanSendMessageGroup()`

UnsetCanSendMessageGroup ensures that no value is present for CanSendMessageGroup, not even an explicit nil
### GetCanSubscribeGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanSubscribeGroup() interface{}`

GetCanSubscribeGroup returns the CanSubscribeGroup field if non-nil, zero value otherwise.

### GetCanSubscribeGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanSubscribeGroupOk() (*interface{}, bool)`

GetCanSubscribeGroupOk returns a tuple with the CanSubscribeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSubscribeGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanSubscribeGroup(v interface{})`

SetCanSubscribeGroup sets CanSubscribeGroup field to given value.

### HasCanSubscribeGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanSubscribeGroup() bool`

HasCanSubscribeGroup returns a boolean if a field has been set.

### SetCanSubscribeGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanSubscribeGroupNil(b bool)`

 SetCanSubscribeGroupNil sets the value for CanSubscribeGroup to be an explicit nil

### UnsetCanSubscribeGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanSubscribeGroup()`

UnsetCanSubscribeGroup ensures that no value is present for CanSubscribeGroup, not even an explicit nil
### GetCanResolveTopicsGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanResolveTopicsGroup() interface{}`

GetCanResolveTopicsGroup returns the CanResolveTopicsGroup field if non-nil, zero value otherwise.

### GetCanResolveTopicsGroupOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetCanResolveTopicsGroupOk() (*interface{}, bool)`

GetCanResolveTopicsGroupOk returns a tuple with the CanResolveTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResolveTopicsGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanResolveTopicsGroup(v interface{})`

SetCanResolveTopicsGroup sets CanResolveTopicsGroup field to given value.

### HasCanResolveTopicsGroup

`func (o *RegisterQueueResponseNeverSubscribedItem) HasCanResolveTopicsGroup() bool`

HasCanResolveTopicsGroup returns a boolean if a field has been set.

### SetCanResolveTopicsGroupNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetCanResolveTopicsGroupNil(b bool)`

 SetCanResolveTopicsGroupNil sets the value for CanResolveTopicsGroup to be an explicit nil

### UnsetCanResolveTopicsGroup
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetCanResolveTopicsGroup()`

UnsetCanResolveTopicsGroup ensures that no value is present for CanResolveTopicsGroup, not even an explicit nil
### GetSubscriberCount

`func (o *RegisterQueueResponseNeverSubscribedItem) GetSubscriberCount() interface{}`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetSubscriberCountOk() (*interface{}, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *RegisterQueueResponseNeverSubscribedItem) SetSubscriberCount(v interface{})`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *RegisterQueueResponseNeverSubscribedItem) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.

### SetSubscriberCountNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetSubscriberCountNil(b bool)`

 SetSubscriberCountNil sets the value for SubscriberCount to be an explicit nil

### UnsetSubscriberCount
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetSubscriberCount()`

UnsetSubscriberCount ensures that no value is present for SubscriberCount, not even an explicit nil
### GetStreamWeeklyTraffic

`func (o *RegisterQueueResponseNeverSubscribedItem) GetStreamWeeklyTraffic() int32`

GetStreamWeeklyTraffic returns the StreamWeeklyTraffic field if non-nil, zero value otherwise.

### GetStreamWeeklyTrafficOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetStreamWeeklyTrafficOk() (*int32, bool)`

GetStreamWeeklyTrafficOk returns a tuple with the StreamWeeklyTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamWeeklyTraffic

`func (o *RegisterQueueResponseNeverSubscribedItem) SetStreamWeeklyTraffic(v int32)`

SetStreamWeeklyTraffic sets StreamWeeklyTraffic field to given value.

### HasStreamWeeklyTraffic

`func (o *RegisterQueueResponseNeverSubscribedItem) HasStreamWeeklyTraffic() bool`

HasStreamWeeklyTraffic returns a boolean if a field has been set.

### SetStreamWeeklyTrafficNil

`func (o *RegisterQueueResponseNeverSubscribedItem) SetStreamWeeklyTrafficNil(b bool)`

 SetStreamWeeklyTrafficNil sets the value for StreamWeeklyTraffic to be an explicit nil

### UnsetStreamWeeklyTraffic
`func (o *RegisterQueueResponseNeverSubscribedItem) UnsetStreamWeeklyTraffic()`

UnsetStreamWeeklyTraffic ensures that no value is present for StreamWeeklyTraffic, not even an explicit nil
### GetSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) GetSubscribers() []int32`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetSubscribersOk() (*[]int32, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) SetSubscribers(v []int32)`

SetSubscribers sets Subscribers field to given value.

### HasSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) HasSubscribers() bool`

HasSubscribers returns a boolean if a field has been set.

### GetPartialSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) GetPartialSubscribers() []int32`

GetPartialSubscribers returns the PartialSubscribers field if non-nil, zero value otherwise.

### GetPartialSubscribersOk

`func (o *RegisterQueueResponseNeverSubscribedItem) GetPartialSubscribersOk() (*[]int32, bool)`

GetPartialSubscribersOk returns a tuple with the PartialSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) SetPartialSubscribers(v []int32)`

SetPartialSubscribers sets PartialSubscribers field to given value.

### HasPartialSubscribers

`func (o *RegisterQueueResponseNeverSubscribedItem) HasPartialSubscribers() bool`

HasPartialSubscribers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


