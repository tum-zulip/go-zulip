# BasicChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamId** | **interface{}** |  | 
**Name** | **interface{}** |  | 
**IsArchived** | **interface{}** |  | 
**Description** | **interface{}** |  | 
**DateCreated** | **interface{}** |  | 
**CreatorId** | **interface{}** |  | 
**InviteOnly** | **interface{}** |  | 
**RenderedDescription** | **interface{}** |  | 
**IsWebPublic** | **interface{}** |  | 
**StreamPostPolicy** | **interface{}** |  | 
**MessageRetentionDays** | **interface{}** |  | 
**HistoryPublicToSubscribers** | **interface{}** |  | 
**TopicsPolicy** | Pointer to **interface{}** |  | [optional] 
**FirstMessageId** | **interface{}** |  | 
**FolderId** | **interface{}** |  | 
**IsRecentlyActive** | **interface{}** |  | 
**IsAnnouncementOnly** | **interface{}** |  | 
**CanAddSubscribersGroup** | Pointer to **interface{}** |  | [optional] 
**CanRemoveSubscribersGroup** | **interface{}** |  | 
**CanAdministerChannelGroup** | Pointer to **interface{}** |  | [optional] 
**CanDeleteAnyMessageGroup** | Pointer to **interface{}** |  | [optional] 
**CanDeleteOwnMessageGroup** | Pointer to **interface{}** |  | [optional] 
**CanMoveMessagesOutOfChannelGroup** | Pointer to **interface{}** |  | [optional] 
**CanMoveMessagesWithinChannelGroup** | Pointer to **interface{}** |  | [optional] 
**CanSendMessageGroup** | Pointer to **interface{}** |  | [optional] 
**CanSubscribeGroup** | **interface{}** |  | 
**CanResolveTopicsGroup** | Pointer to **interface{}** |  | [optional] 
**SubscriberCount** | **interface{}** |  | 
**StreamWeeklyTraffic** | **NullableInt32** | The average number of messages sent to the channel per week, as estimated based on recent weeks, rounded to the nearest integer.  If &#x60;null&#x60;, no information is provided on the average traffic. This can be because the channel was recently created and there is insufficient data to make an estimate, or because the server wishes to omit this information for this client, this realm, or this endpoint or type of event.  **Changes**: New in Zulip 8.0 (feature level 199). Previously, this statistic was available only in subscription objects.  | 

## Methods

### NewBasicChannel

`func NewBasicChannel(streamId interface{}, name interface{}, isArchived interface{}, description interface{}, dateCreated interface{}, creatorId interface{}, inviteOnly interface{}, renderedDescription interface{}, isWebPublic interface{}, streamPostPolicy interface{}, messageRetentionDays interface{}, historyPublicToSubscribers interface{}, firstMessageId interface{}, folderId interface{}, isRecentlyActive interface{}, isAnnouncementOnly interface{}, canRemoveSubscribersGroup interface{}, canSubscribeGroup interface{}, subscriberCount interface{}, streamWeeklyTraffic NullableInt32, ) *BasicChannel`

NewBasicChannel instantiates a new BasicChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicChannelWithDefaults

`func NewBasicChannelWithDefaults() *BasicChannel`

NewBasicChannelWithDefaults instantiates a new BasicChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamId

`func (o *BasicChannel) GetStreamId() interface{}`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *BasicChannel) GetStreamIdOk() (*interface{}, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *BasicChannel) SetStreamId(v interface{})`

SetStreamId sets StreamId field to given value.


### SetStreamIdNil

`func (o *BasicChannel) SetStreamIdNil(b bool)`

 SetStreamIdNil sets the value for StreamId to be an explicit nil

### UnsetStreamId
`func (o *BasicChannel) UnsetStreamId()`

UnsetStreamId ensures that no value is present for StreamId, not even an explicit nil
### GetName

`func (o *BasicChannel) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicChannel) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicChannel) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *BasicChannel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BasicChannel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsArchived

`func (o *BasicChannel) GetIsArchived() interface{}`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *BasicChannel) GetIsArchivedOk() (*interface{}, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *BasicChannel) SetIsArchived(v interface{})`

SetIsArchived sets IsArchived field to given value.


### SetIsArchivedNil

`func (o *BasicChannel) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *BasicChannel) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil
### GetDescription

`func (o *BasicChannel) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BasicChannel) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BasicChannel) SetDescription(v interface{})`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *BasicChannel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BasicChannel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDateCreated

`func (o *BasicChannel) GetDateCreated() interface{}`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BasicChannel) GetDateCreatedOk() (*interface{}, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BasicChannel) SetDateCreated(v interface{})`

SetDateCreated sets DateCreated field to given value.


### SetDateCreatedNil

`func (o *BasicChannel) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *BasicChannel) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetCreatorId

`func (o *BasicChannel) GetCreatorId() interface{}`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *BasicChannel) GetCreatorIdOk() (*interface{}, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *BasicChannel) SetCreatorId(v interface{})`

SetCreatorId sets CreatorId field to given value.


### SetCreatorIdNil

`func (o *BasicChannel) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *BasicChannel) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetInviteOnly

`func (o *BasicChannel) GetInviteOnly() interface{}`

GetInviteOnly returns the InviteOnly field if non-nil, zero value otherwise.

### GetInviteOnlyOk

`func (o *BasicChannel) GetInviteOnlyOk() (*interface{}, bool)`

GetInviteOnlyOk returns a tuple with the InviteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteOnly

`func (o *BasicChannel) SetInviteOnly(v interface{})`

SetInviteOnly sets InviteOnly field to given value.


### SetInviteOnlyNil

`func (o *BasicChannel) SetInviteOnlyNil(b bool)`

 SetInviteOnlyNil sets the value for InviteOnly to be an explicit nil

### UnsetInviteOnly
`func (o *BasicChannel) UnsetInviteOnly()`

UnsetInviteOnly ensures that no value is present for InviteOnly, not even an explicit nil
### GetRenderedDescription

`func (o *BasicChannel) GetRenderedDescription() interface{}`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *BasicChannel) GetRenderedDescriptionOk() (*interface{}, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *BasicChannel) SetRenderedDescription(v interface{})`

SetRenderedDescription sets RenderedDescription field to given value.


### SetRenderedDescriptionNil

`func (o *BasicChannel) SetRenderedDescriptionNil(b bool)`

 SetRenderedDescriptionNil sets the value for RenderedDescription to be an explicit nil

### UnsetRenderedDescription
`func (o *BasicChannel) UnsetRenderedDescription()`

UnsetRenderedDescription ensures that no value is present for RenderedDescription, not even an explicit nil
### GetIsWebPublic

`func (o *BasicChannel) GetIsWebPublic() interface{}`

GetIsWebPublic returns the IsWebPublic field if non-nil, zero value otherwise.

### GetIsWebPublicOk

`func (o *BasicChannel) GetIsWebPublicOk() (*interface{}, bool)`

GetIsWebPublicOk returns a tuple with the IsWebPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebPublic

`func (o *BasicChannel) SetIsWebPublic(v interface{})`

SetIsWebPublic sets IsWebPublic field to given value.


### SetIsWebPublicNil

`func (o *BasicChannel) SetIsWebPublicNil(b bool)`

 SetIsWebPublicNil sets the value for IsWebPublic to be an explicit nil

### UnsetIsWebPublic
`func (o *BasicChannel) UnsetIsWebPublic()`

UnsetIsWebPublic ensures that no value is present for IsWebPublic, not even an explicit nil
### GetStreamPostPolicy

`func (o *BasicChannel) GetStreamPostPolicy() interface{}`

GetStreamPostPolicy returns the StreamPostPolicy field if non-nil, zero value otherwise.

### GetStreamPostPolicyOk

`func (o *BasicChannel) GetStreamPostPolicyOk() (*interface{}, bool)`

GetStreamPostPolicyOk returns a tuple with the StreamPostPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamPostPolicy

`func (o *BasicChannel) SetStreamPostPolicy(v interface{})`

SetStreamPostPolicy sets StreamPostPolicy field to given value.


### SetStreamPostPolicyNil

`func (o *BasicChannel) SetStreamPostPolicyNil(b bool)`

 SetStreamPostPolicyNil sets the value for StreamPostPolicy to be an explicit nil

### UnsetStreamPostPolicy
`func (o *BasicChannel) UnsetStreamPostPolicy()`

UnsetStreamPostPolicy ensures that no value is present for StreamPostPolicy, not even an explicit nil
### GetMessageRetentionDays

`func (o *BasicChannel) GetMessageRetentionDays() interface{}`

GetMessageRetentionDays returns the MessageRetentionDays field if non-nil, zero value otherwise.

### GetMessageRetentionDaysOk

`func (o *BasicChannel) GetMessageRetentionDaysOk() (*interface{}, bool)`

GetMessageRetentionDaysOk returns a tuple with the MessageRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRetentionDays

`func (o *BasicChannel) SetMessageRetentionDays(v interface{})`

SetMessageRetentionDays sets MessageRetentionDays field to given value.


### SetMessageRetentionDaysNil

`func (o *BasicChannel) SetMessageRetentionDaysNil(b bool)`

 SetMessageRetentionDaysNil sets the value for MessageRetentionDays to be an explicit nil

### UnsetMessageRetentionDays
`func (o *BasicChannel) UnsetMessageRetentionDays()`

UnsetMessageRetentionDays ensures that no value is present for MessageRetentionDays, not even an explicit nil
### GetHistoryPublicToSubscribers

`func (o *BasicChannel) GetHistoryPublicToSubscribers() interface{}`

GetHistoryPublicToSubscribers returns the HistoryPublicToSubscribers field if non-nil, zero value otherwise.

### GetHistoryPublicToSubscribersOk

`func (o *BasicChannel) GetHistoryPublicToSubscribersOk() (*interface{}, bool)`

GetHistoryPublicToSubscribersOk returns a tuple with the HistoryPublicToSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPublicToSubscribers

`func (o *BasicChannel) SetHistoryPublicToSubscribers(v interface{})`

SetHistoryPublicToSubscribers sets HistoryPublicToSubscribers field to given value.


### SetHistoryPublicToSubscribersNil

`func (o *BasicChannel) SetHistoryPublicToSubscribersNil(b bool)`

 SetHistoryPublicToSubscribersNil sets the value for HistoryPublicToSubscribers to be an explicit nil

### UnsetHistoryPublicToSubscribers
`func (o *BasicChannel) UnsetHistoryPublicToSubscribers()`

UnsetHistoryPublicToSubscribers ensures that no value is present for HistoryPublicToSubscribers, not even an explicit nil
### GetTopicsPolicy

`func (o *BasicChannel) GetTopicsPolicy() interface{}`

GetTopicsPolicy returns the TopicsPolicy field if non-nil, zero value otherwise.

### GetTopicsPolicyOk

`func (o *BasicChannel) GetTopicsPolicyOk() (*interface{}, bool)`

GetTopicsPolicyOk returns a tuple with the TopicsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsPolicy

`func (o *BasicChannel) SetTopicsPolicy(v interface{})`

SetTopicsPolicy sets TopicsPolicy field to given value.

### HasTopicsPolicy

`func (o *BasicChannel) HasTopicsPolicy() bool`

HasTopicsPolicy returns a boolean if a field has been set.

### SetTopicsPolicyNil

`func (o *BasicChannel) SetTopicsPolicyNil(b bool)`

 SetTopicsPolicyNil sets the value for TopicsPolicy to be an explicit nil

### UnsetTopicsPolicy
`func (o *BasicChannel) UnsetTopicsPolicy()`

UnsetTopicsPolicy ensures that no value is present for TopicsPolicy, not even an explicit nil
### GetFirstMessageId

`func (o *BasicChannel) GetFirstMessageId() interface{}`

GetFirstMessageId returns the FirstMessageId field if non-nil, zero value otherwise.

### GetFirstMessageIdOk

`func (o *BasicChannel) GetFirstMessageIdOk() (*interface{}, bool)`

GetFirstMessageIdOk returns a tuple with the FirstMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessageId

`func (o *BasicChannel) SetFirstMessageId(v interface{})`

SetFirstMessageId sets FirstMessageId field to given value.


### SetFirstMessageIdNil

`func (o *BasicChannel) SetFirstMessageIdNil(b bool)`

 SetFirstMessageIdNil sets the value for FirstMessageId to be an explicit nil

### UnsetFirstMessageId
`func (o *BasicChannel) UnsetFirstMessageId()`

UnsetFirstMessageId ensures that no value is present for FirstMessageId, not even an explicit nil
### GetFolderId

`func (o *BasicChannel) GetFolderId() interface{}`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BasicChannel) GetFolderIdOk() (*interface{}, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BasicChannel) SetFolderId(v interface{})`

SetFolderId sets FolderId field to given value.


### SetFolderIdNil

`func (o *BasicChannel) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *BasicChannel) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetIsRecentlyActive

`func (o *BasicChannel) GetIsRecentlyActive() interface{}`

GetIsRecentlyActive returns the IsRecentlyActive field if non-nil, zero value otherwise.

### GetIsRecentlyActiveOk

`func (o *BasicChannel) GetIsRecentlyActiveOk() (*interface{}, bool)`

GetIsRecentlyActiveOk returns a tuple with the IsRecentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecentlyActive

`func (o *BasicChannel) SetIsRecentlyActive(v interface{})`

SetIsRecentlyActive sets IsRecentlyActive field to given value.


### SetIsRecentlyActiveNil

`func (o *BasicChannel) SetIsRecentlyActiveNil(b bool)`

 SetIsRecentlyActiveNil sets the value for IsRecentlyActive to be an explicit nil

### UnsetIsRecentlyActive
`func (o *BasicChannel) UnsetIsRecentlyActive()`

UnsetIsRecentlyActive ensures that no value is present for IsRecentlyActive, not even an explicit nil
### GetIsAnnouncementOnly

`func (o *BasicChannel) GetIsAnnouncementOnly() interface{}`

GetIsAnnouncementOnly returns the IsAnnouncementOnly field if non-nil, zero value otherwise.

### GetIsAnnouncementOnlyOk

`func (o *BasicChannel) GetIsAnnouncementOnlyOk() (*interface{}, bool)`

GetIsAnnouncementOnlyOk returns a tuple with the IsAnnouncementOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnouncementOnly

`func (o *BasicChannel) SetIsAnnouncementOnly(v interface{})`

SetIsAnnouncementOnly sets IsAnnouncementOnly field to given value.


### SetIsAnnouncementOnlyNil

`func (o *BasicChannel) SetIsAnnouncementOnlyNil(b bool)`

 SetIsAnnouncementOnlyNil sets the value for IsAnnouncementOnly to be an explicit nil

### UnsetIsAnnouncementOnly
`func (o *BasicChannel) UnsetIsAnnouncementOnly()`

UnsetIsAnnouncementOnly ensures that no value is present for IsAnnouncementOnly, not even an explicit nil
### GetCanAddSubscribersGroup

`func (o *BasicChannel) GetCanAddSubscribersGroup() interface{}`

GetCanAddSubscribersGroup returns the CanAddSubscribersGroup field if non-nil, zero value otherwise.

### GetCanAddSubscribersGroupOk

`func (o *BasicChannel) GetCanAddSubscribersGroupOk() (*interface{}, bool)`

GetCanAddSubscribersGroupOk returns a tuple with the CanAddSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddSubscribersGroup

`func (o *BasicChannel) SetCanAddSubscribersGroup(v interface{})`

SetCanAddSubscribersGroup sets CanAddSubscribersGroup field to given value.

### HasCanAddSubscribersGroup

`func (o *BasicChannel) HasCanAddSubscribersGroup() bool`

HasCanAddSubscribersGroup returns a boolean if a field has been set.

### SetCanAddSubscribersGroupNil

`func (o *BasicChannel) SetCanAddSubscribersGroupNil(b bool)`

 SetCanAddSubscribersGroupNil sets the value for CanAddSubscribersGroup to be an explicit nil

### UnsetCanAddSubscribersGroup
`func (o *BasicChannel) UnsetCanAddSubscribersGroup()`

UnsetCanAddSubscribersGroup ensures that no value is present for CanAddSubscribersGroup, not even an explicit nil
### GetCanRemoveSubscribersGroup

`func (o *BasicChannel) GetCanRemoveSubscribersGroup() interface{}`

GetCanRemoveSubscribersGroup returns the CanRemoveSubscribersGroup field if non-nil, zero value otherwise.

### GetCanRemoveSubscribersGroupOk

`func (o *BasicChannel) GetCanRemoveSubscribersGroupOk() (*interface{}, bool)`

GetCanRemoveSubscribersGroupOk returns a tuple with the CanRemoveSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemoveSubscribersGroup

`func (o *BasicChannel) SetCanRemoveSubscribersGroup(v interface{})`

SetCanRemoveSubscribersGroup sets CanRemoveSubscribersGroup field to given value.


### SetCanRemoveSubscribersGroupNil

`func (o *BasicChannel) SetCanRemoveSubscribersGroupNil(b bool)`

 SetCanRemoveSubscribersGroupNil sets the value for CanRemoveSubscribersGroup to be an explicit nil

### UnsetCanRemoveSubscribersGroup
`func (o *BasicChannel) UnsetCanRemoveSubscribersGroup()`

UnsetCanRemoveSubscribersGroup ensures that no value is present for CanRemoveSubscribersGroup, not even an explicit nil
### GetCanAdministerChannelGroup

`func (o *BasicChannel) GetCanAdministerChannelGroup() interface{}`

GetCanAdministerChannelGroup returns the CanAdministerChannelGroup field if non-nil, zero value otherwise.

### GetCanAdministerChannelGroupOk

`func (o *BasicChannel) GetCanAdministerChannelGroupOk() (*interface{}, bool)`

GetCanAdministerChannelGroupOk returns a tuple with the CanAdministerChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdministerChannelGroup

`func (o *BasicChannel) SetCanAdministerChannelGroup(v interface{})`

SetCanAdministerChannelGroup sets CanAdministerChannelGroup field to given value.

### HasCanAdministerChannelGroup

`func (o *BasicChannel) HasCanAdministerChannelGroup() bool`

HasCanAdministerChannelGroup returns a boolean if a field has been set.

### SetCanAdministerChannelGroupNil

`func (o *BasicChannel) SetCanAdministerChannelGroupNil(b bool)`

 SetCanAdministerChannelGroupNil sets the value for CanAdministerChannelGroup to be an explicit nil

### UnsetCanAdministerChannelGroup
`func (o *BasicChannel) UnsetCanAdministerChannelGroup()`

UnsetCanAdministerChannelGroup ensures that no value is present for CanAdministerChannelGroup, not even an explicit nil
### GetCanDeleteAnyMessageGroup

`func (o *BasicChannel) GetCanDeleteAnyMessageGroup() interface{}`

GetCanDeleteAnyMessageGroup returns the CanDeleteAnyMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteAnyMessageGroupOk

`func (o *BasicChannel) GetCanDeleteAnyMessageGroupOk() (*interface{}, bool)`

GetCanDeleteAnyMessageGroupOk returns a tuple with the CanDeleteAnyMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteAnyMessageGroup

`func (o *BasicChannel) SetCanDeleteAnyMessageGroup(v interface{})`

SetCanDeleteAnyMessageGroup sets CanDeleteAnyMessageGroup field to given value.

### HasCanDeleteAnyMessageGroup

`func (o *BasicChannel) HasCanDeleteAnyMessageGroup() bool`

HasCanDeleteAnyMessageGroup returns a boolean if a field has been set.

### SetCanDeleteAnyMessageGroupNil

`func (o *BasicChannel) SetCanDeleteAnyMessageGroupNil(b bool)`

 SetCanDeleteAnyMessageGroupNil sets the value for CanDeleteAnyMessageGroup to be an explicit nil

### UnsetCanDeleteAnyMessageGroup
`func (o *BasicChannel) UnsetCanDeleteAnyMessageGroup()`

UnsetCanDeleteAnyMessageGroup ensures that no value is present for CanDeleteAnyMessageGroup, not even an explicit nil
### GetCanDeleteOwnMessageGroup

`func (o *BasicChannel) GetCanDeleteOwnMessageGroup() interface{}`

GetCanDeleteOwnMessageGroup returns the CanDeleteOwnMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteOwnMessageGroupOk

`func (o *BasicChannel) GetCanDeleteOwnMessageGroupOk() (*interface{}, bool)`

GetCanDeleteOwnMessageGroupOk returns a tuple with the CanDeleteOwnMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteOwnMessageGroup

`func (o *BasicChannel) SetCanDeleteOwnMessageGroup(v interface{})`

SetCanDeleteOwnMessageGroup sets CanDeleteOwnMessageGroup field to given value.

### HasCanDeleteOwnMessageGroup

`func (o *BasicChannel) HasCanDeleteOwnMessageGroup() bool`

HasCanDeleteOwnMessageGroup returns a boolean if a field has been set.

### SetCanDeleteOwnMessageGroupNil

`func (o *BasicChannel) SetCanDeleteOwnMessageGroupNil(b bool)`

 SetCanDeleteOwnMessageGroupNil sets the value for CanDeleteOwnMessageGroup to be an explicit nil

### UnsetCanDeleteOwnMessageGroup
`func (o *BasicChannel) UnsetCanDeleteOwnMessageGroup()`

UnsetCanDeleteOwnMessageGroup ensures that no value is present for CanDeleteOwnMessageGroup, not even an explicit nil
### GetCanMoveMessagesOutOfChannelGroup

`func (o *BasicChannel) GetCanMoveMessagesOutOfChannelGroup() interface{}`

GetCanMoveMessagesOutOfChannelGroup returns the CanMoveMessagesOutOfChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesOutOfChannelGroupOk

`func (o *BasicChannel) GetCanMoveMessagesOutOfChannelGroupOk() (*interface{}, bool)`

GetCanMoveMessagesOutOfChannelGroupOk returns a tuple with the CanMoveMessagesOutOfChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesOutOfChannelGroup

`func (o *BasicChannel) SetCanMoveMessagesOutOfChannelGroup(v interface{})`

SetCanMoveMessagesOutOfChannelGroup sets CanMoveMessagesOutOfChannelGroup field to given value.

### HasCanMoveMessagesOutOfChannelGroup

`func (o *BasicChannel) HasCanMoveMessagesOutOfChannelGroup() bool`

HasCanMoveMessagesOutOfChannelGroup returns a boolean if a field has been set.

### SetCanMoveMessagesOutOfChannelGroupNil

`func (o *BasicChannel) SetCanMoveMessagesOutOfChannelGroupNil(b bool)`

 SetCanMoveMessagesOutOfChannelGroupNil sets the value for CanMoveMessagesOutOfChannelGroup to be an explicit nil

### UnsetCanMoveMessagesOutOfChannelGroup
`func (o *BasicChannel) UnsetCanMoveMessagesOutOfChannelGroup()`

UnsetCanMoveMessagesOutOfChannelGroup ensures that no value is present for CanMoveMessagesOutOfChannelGroup, not even an explicit nil
### GetCanMoveMessagesWithinChannelGroup

`func (o *BasicChannel) GetCanMoveMessagesWithinChannelGroup() interface{}`

GetCanMoveMessagesWithinChannelGroup returns the CanMoveMessagesWithinChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesWithinChannelGroupOk

`func (o *BasicChannel) GetCanMoveMessagesWithinChannelGroupOk() (*interface{}, bool)`

GetCanMoveMessagesWithinChannelGroupOk returns a tuple with the CanMoveMessagesWithinChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesWithinChannelGroup

`func (o *BasicChannel) SetCanMoveMessagesWithinChannelGroup(v interface{})`

SetCanMoveMessagesWithinChannelGroup sets CanMoveMessagesWithinChannelGroup field to given value.

### HasCanMoveMessagesWithinChannelGroup

`func (o *BasicChannel) HasCanMoveMessagesWithinChannelGroup() bool`

HasCanMoveMessagesWithinChannelGroup returns a boolean if a field has been set.

### SetCanMoveMessagesWithinChannelGroupNil

`func (o *BasicChannel) SetCanMoveMessagesWithinChannelGroupNil(b bool)`

 SetCanMoveMessagesWithinChannelGroupNil sets the value for CanMoveMessagesWithinChannelGroup to be an explicit nil

### UnsetCanMoveMessagesWithinChannelGroup
`func (o *BasicChannel) UnsetCanMoveMessagesWithinChannelGroup()`

UnsetCanMoveMessagesWithinChannelGroup ensures that no value is present for CanMoveMessagesWithinChannelGroup, not even an explicit nil
### GetCanSendMessageGroup

`func (o *BasicChannel) GetCanSendMessageGroup() interface{}`

GetCanSendMessageGroup returns the CanSendMessageGroup field if non-nil, zero value otherwise.

### GetCanSendMessageGroupOk

`func (o *BasicChannel) GetCanSendMessageGroupOk() (*interface{}, bool)`

GetCanSendMessageGroupOk returns a tuple with the CanSendMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendMessageGroup

`func (o *BasicChannel) SetCanSendMessageGroup(v interface{})`

SetCanSendMessageGroup sets CanSendMessageGroup field to given value.

### HasCanSendMessageGroup

`func (o *BasicChannel) HasCanSendMessageGroup() bool`

HasCanSendMessageGroup returns a boolean if a field has been set.

### SetCanSendMessageGroupNil

`func (o *BasicChannel) SetCanSendMessageGroupNil(b bool)`

 SetCanSendMessageGroupNil sets the value for CanSendMessageGroup to be an explicit nil

### UnsetCanSendMessageGroup
`func (o *BasicChannel) UnsetCanSendMessageGroup()`

UnsetCanSendMessageGroup ensures that no value is present for CanSendMessageGroup, not even an explicit nil
### GetCanSubscribeGroup

`func (o *BasicChannel) GetCanSubscribeGroup() interface{}`

GetCanSubscribeGroup returns the CanSubscribeGroup field if non-nil, zero value otherwise.

### GetCanSubscribeGroupOk

`func (o *BasicChannel) GetCanSubscribeGroupOk() (*interface{}, bool)`

GetCanSubscribeGroupOk returns a tuple with the CanSubscribeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSubscribeGroup

`func (o *BasicChannel) SetCanSubscribeGroup(v interface{})`

SetCanSubscribeGroup sets CanSubscribeGroup field to given value.


### SetCanSubscribeGroupNil

`func (o *BasicChannel) SetCanSubscribeGroupNil(b bool)`

 SetCanSubscribeGroupNil sets the value for CanSubscribeGroup to be an explicit nil

### UnsetCanSubscribeGroup
`func (o *BasicChannel) UnsetCanSubscribeGroup()`

UnsetCanSubscribeGroup ensures that no value is present for CanSubscribeGroup, not even an explicit nil
### GetCanResolveTopicsGroup

`func (o *BasicChannel) GetCanResolveTopicsGroup() interface{}`

GetCanResolveTopicsGroup returns the CanResolveTopicsGroup field if non-nil, zero value otherwise.

### GetCanResolveTopicsGroupOk

`func (o *BasicChannel) GetCanResolveTopicsGroupOk() (*interface{}, bool)`

GetCanResolveTopicsGroupOk returns a tuple with the CanResolveTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResolveTopicsGroup

`func (o *BasicChannel) SetCanResolveTopicsGroup(v interface{})`

SetCanResolveTopicsGroup sets CanResolveTopicsGroup field to given value.

### HasCanResolveTopicsGroup

`func (o *BasicChannel) HasCanResolveTopicsGroup() bool`

HasCanResolveTopicsGroup returns a boolean if a field has been set.

### SetCanResolveTopicsGroupNil

`func (o *BasicChannel) SetCanResolveTopicsGroupNil(b bool)`

 SetCanResolveTopicsGroupNil sets the value for CanResolveTopicsGroup to be an explicit nil

### UnsetCanResolveTopicsGroup
`func (o *BasicChannel) UnsetCanResolveTopicsGroup()`

UnsetCanResolveTopicsGroup ensures that no value is present for CanResolveTopicsGroup, not even an explicit nil
### GetSubscriberCount

`func (o *BasicChannel) GetSubscriberCount() interface{}`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *BasicChannel) GetSubscriberCountOk() (*interface{}, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *BasicChannel) SetSubscriberCount(v interface{})`

SetSubscriberCount sets SubscriberCount field to given value.


### SetSubscriberCountNil

`func (o *BasicChannel) SetSubscriberCountNil(b bool)`

 SetSubscriberCountNil sets the value for SubscriberCount to be an explicit nil

### UnsetSubscriberCount
`func (o *BasicChannel) UnsetSubscriberCount()`

UnsetSubscriberCount ensures that no value is present for SubscriberCount, not even an explicit nil
### GetStreamWeeklyTraffic

`func (o *BasicChannel) GetStreamWeeklyTraffic() int32`

GetStreamWeeklyTraffic returns the StreamWeeklyTraffic field if non-nil, zero value otherwise.

### GetStreamWeeklyTrafficOk

`func (o *BasicChannel) GetStreamWeeklyTrafficOk() (*int32, bool)`

GetStreamWeeklyTrafficOk returns a tuple with the StreamWeeklyTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamWeeklyTraffic

`func (o *BasicChannel) SetStreamWeeklyTraffic(v int32)`

SetStreamWeeklyTraffic sets StreamWeeklyTraffic field to given value.


### SetStreamWeeklyTrafficNil

`func (o *BasicChannel) SetStreamWeeklyTrafficNil(b bool)`

 SetStreamWeeklyTrafficNil sets the value for StreamWeeklyTraffic to be an explicit nil

### UnsetStreamWeeklyTraffic
`func (o *BasicChannel) UnsetStreamWeeklyTraffic()`

UnsetStreamWeeklyTraffic ensures that no value is present for StreamWeeklyTraffic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


