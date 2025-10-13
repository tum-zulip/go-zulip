# GetStreams200ResponseAllOfStreamsInner

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
**IsDefault** | Pointer to **bool** | Only present when [&#x60;include_default&#x60;][include_default] parameter is &#x60;true&#x60;.  Whether the given channel is a [default channel](zulip.com/help/set-default-channels-for-new-users.  [include_default]: /api/get-streams#parameter-include_default  | [optional] 

## Methods

### NewGetStreams200ResponseAllOfStreamsInner

`func NewGetStreams200ResponseAllOfStreamsInner(streamId interface{}, name interface{}, isArchived interface{}, description interface{}, dateCreated interface{}, creatorId interface{}, inviteOnly interface{}, renderedDescription interface{}, isWebPublic interface{}, streamPostPolicy interface{}, messageRetentionDays interface{}, historyPublicToSubscribers interface{}, firstMessageId interface{}, folderId interface{}, isRecentlyActive interface{}, isAnnouncementOnly interface{}, canRemoveSubscribersGroup interface{}, canSubscribeGroup interface{}, subscriberCount interface{}, streamWeeklyTraffic NullableInt32, ) *GetStreams200ResponseAllOfStreamsInner`

NewGetStreams200ResponseAllOfStreamsInner instantiates a new GetStreams200ResponseAllOfStreamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStreams200ResponseAllOfStreamsInnerWithDefaults

`func NewGetStreams200ResponseAllOfStreamsInnerWithDefaults() *GetStreams200ResponseAllOfStreamsInner`

NewGetStreams200ResponseAllOfStreamsInnerWithDefaults instantiates a new GetStreams200ResponseAllOfStreamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamId

`func (o *GetStreams200ResponseAllOfStreamsInner) GetStreamId() interface{}`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetStreamIdOk() (*interface{}, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *GetStreams200ResponseAllOfStreamsInner) SetStreamId(v interface{})`

SetStreamId sets StreamId field to given value.


### SetStreamIdNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetStreamIdNil(b bool)`

 SetStreamIdNil sets the value for StreamId to be an explicit nil

### UnsetStreamId
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetStreamId()`

UnsetStreamId ensures that no value is present for StreamId, not even an explicit nil
### GetName

`func (o *GetStreams200ResponseAllOfStreamsInner) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetStreams200ResponseAllOfStreamsInner) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsArchived

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsArchived() interface{}`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsArchivedOk() (*interface{}, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsArchived(v interface{})`

SetIsArchived sets IsArchived field to given value.


### SetIsArchivedNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil
### GetDescription

`func (o *GetStreams200ResponseAllOfStreamsInner) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetStreams200ResponseAllOfStreamsInner) SetDescription(v interface{})`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDateCreated

`func (o *GetStreams200ResponseAllOfStreamsInner) GetDateCreated() interface{}`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetDateCreatedOk() (*interface{}, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetStreams200ResponseAllOfStreamsInner) SetDateCreated(v interface{})`

SetDateCreated sets DateCreated field to given value.


### SetDateCreatedNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetCreatorId

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCreatorId() interface{}`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCreatorIdOk() (*interface{}, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCreatorId(v interface{})`

SetCreatorId sets CreatorId field to given value.


### SetCreatorIdNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetInviteOnly

`func (o *GetStreams200ResponseAllOfStreamsInner) GetInviteOnly() interface{}`

GetInviteOnly returns the InviteOnly field if non-nil, zero value otherwise.

### GetInviteOnlyOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetInviteOnlyOk() (*interface{}, bool)`

GetInviteOnlyOk returns a tuple with the InviteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteOnly

`func (o *GetStreams200ResponseAllOfStreamsInner) SetInviteOnly(v interface{})`

SetInviteOnly sets InviteOnly field to given value.


### SetInviteOnlyNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetInviteOnlyNil(b bool)`

 SetInviteOnlyNil sets the value for InviteOnly to be an explicit nil

### UnsetInviteOnly
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetInviteOnly()`

UnsetInviteOnly ensures that no value is present for InviteOnly, not even an explicit nil
### GetRenderedDescription

`func (o *GetStreams200ResponseAllOfStreamsInner) GetRenderedDescription() interface{}`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetRenderedDescriptionOk() (*interface{}, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *GetStreams200ResponseAllOfStreamsInner) SetRenderedDescription(v interface{})`

SetRenderedDescription sets RenderedDescription field to given value.


### SetRenderedDescriptionNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetRenderedDescriptionNil(b bool)`

 SetRenderedDescriptionNil sets the value for RenderedDescription to be an explicit nil

### UnsetRenderedDescription
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetRenderedDescription()`

UnsetRenderedDescription ensures that no value is present for RenderedDescription, not even an explicit nil
### GetIsWebPublic

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsWebPublic() interface{}`

GetIsWebPublic returns the IsWebPublic field if non-nil, zero value otherwise.

### GetIsWebPublicOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsWebPublicOk() (*interface{}, bool)`

GetIsWebPublicOk returns a tuple with the IsWebPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebPublic

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsWebPublic(v interface{})`

SetIsWebPublic sets IsWebPublic field to given value.


### SetIsWebPublicNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsWebPublicNil(b bool)`

 SetIsWebPublicNil sets the value for IsWebPublic to be an explicit nil

### UnsetIsWebPublic
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetIsWebPublic()`

UnsetIsWebPublic ensures that no value is present for IsWebPublic, not even an explicit nil
### GetStreamPostPolicy

`func (o *GetStreams200ResponseAllOfStreamsInner) GetStreamPostPolicy() interface{}`

GetStreamPostPolicy returns the StreamPostPolicy field if non-nil, zero value otherwise.

### GetStreamPostPolicyOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetStreamPostPolicyOk() (*interface{}, bool)`

GetStreamPostPolicyOk returns a tuple with the StreamPostPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamPostPolicy

`func (o *GetStreams200ResponseAllOfStreamsInner) SetStreamPostPolicy(v interface{})`

SetStreamPostPolicy sets StreamPostPolicy field to given value.


### SetStreamPostPolicyNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetStreamPostPolicyNil(b bool)`

 SetStreamPostPolicyNil sets the value for StreamPostPolicy to be an explicit nil

### UnsetStreamPostPolicy
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetStreamPostPolicy()`

UnsetStreamPostPolicy ensures that no value is present for StreamPostPolicy, not even an explicit nil
### GetMessageRetentionDays

`func (o *GetStreams200ResponseAllOfStreamsInner) GetMessageRetentionDays() interface{}`

GetMessageRetentionDays returns the MessageRetentionDays field if non-nil, zero value otherwise.

### GetMessageRetentionDaysOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetMessageRetentionDaysOk() (*interface{}, bool)`

GetMessageRetentionDaysOk returns a tuple with the MessageRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRetentionDays

`func (o *GetStreams200ResponseAllOfStreamsInner) SetMessageRetentionDays(v interface{})`

SetMessageRetentionDays sets MessageRetentionDays field to given value.


### SetMessageRetentionDaysNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetMessageRetentionDaysNil(b bool)`

 SetMessageRetentionDaysNil sets the value for MessageRetentionDays to be an explicit nil

### UnsetMessageRetentionDays
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetMessageRetentionDays()`

UnsetMessageRetentionDays ensures that no value is present for MessageRetentionDays, not even an explicit nil
### GetHistoryPublicToSubscribers

`func (o *GetStreams200ResponseAllOfStreamsInner) GetHistoryPublicToSubscribers() interface{}`

GetHistoryPublicToSubscribers returns the HistoryPublicToSubscribers field if non-nil, zero value otherwise.

### GetHistoryPublicToSubscribersOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetHistoryPublicToSubscribersOk() (*interface{}, bool)`

GetHistoryPublicToSubscribersOk returns a tuple with the HistoryPublicToSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPublicToSubscribers

`func (o *GetStreams200ResponseAllOfStreamsInner) SetHistoryPublicToSubscribers(v interface{})`

SetHistoryPublicToSubscribers sets HistoryPublicToSubscribers field to given value.


### SetHistoryPublicToSubscribersNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetHistoryPublicToSubscribersNil(b bool)`

 SetHistoryPublicToSubscribersNil sets the value for HistoryPublicToSubscribers to be an explicit nil

### UnsetHistoryPublicToSubscribers
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetHistoryPublicToSubscribers()`

UnsetHistoryPublicToSubscribers ensures that no value is present for HistoryPublicToSubscribers, not even an explicit nil
### GetTopicsPolicy

`func (o *GetStreams200ResponseAllOfStreamsInner) GetTopicsPolicy() interface{}`

GetTopicsPolicy returns the TopicsPolicy field if non-nil, zero value otherwise.

### GetTopicsPolicyOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetTopicsPolicyOk() (*interface{}, bool)`

GetTopicsPolicyOk returns a tuple with the TopicsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsPolicy

`func (o *GetStreams200ResponseAllOfStreamsInner) SetTopicsPolicy(v interface{})`

SetTopicsPolicy sets TopicsPolicy field to given value.

### HasTopicsPolicy

`func (o *GetStreams200ResponseAllOfStreamsInner) HasTopicsPolicy() bool`

HasTopicsPolicy returns a boolean if a field has been set.

### SetTopicsPolicyNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetTopicsPolicyNil(b bool)`

 SetTopicsPolicyNil sets the value for TopicsPolicy to be an explicit nil

### UnsetTopicsPolicy
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetTopicsPolicy()`

UnsetTopicsPolicy ensures that no value is present for TopicsPolicy, not even an explicit nil
### GetFirstMessageId

`func (o *GetStreams200ResponseAllOfStreamsInner) GetFirstMessageId() interface{}`

GetFirstMessageId returns the FirstMessageId field if non-nil, zero value otherwise.

### GetFirstMessageIdOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetFirstMessageIdOk() (*interface{}, bool)`

GetFirstMessageIdOk returns a tuple with the FirstMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessageId

`func (o *GetStreams200ResponseAllOfStreamsInner) SetFirstMessageId(v interface{})`

SetFirstMessageId sets FirstMessageId field to given value.


### SetFirstMessageIdNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetFirstMessageIdNil(b bool)`

 SetFirstMessageIdNil sets the value for FirstMessageId to be an explicit nil

### UnsetFirstMessageId
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetFirstMessageId()`

UnsetFirstMessageId ensures that no value is present for FirstMessageId, not even an explicit nil
### GetFolderId

`func (o *GetStreams200ResponseAllOfStreamsInner) GetFolderId() interface{}`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetFolderIdOk() (*interface{}, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GetStreams200ResponseAllOfStreamsInner) SetFolderId(v interface{})`

SetFolderId sets FolderId field to given value.


### SetFolderIdNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetIsRecentlyActive

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsRecentlyActive() interface{}`

GetIsRecentlyActive returns the IsRecentlyActive field if non-nil, zero value otherwise.

### GetIsRecentlyActiveOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsRecentlyActiveOk() (*interface{}, bool)`

GetIsRecentlyActiveOk returns a tuple with the IsRecentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecentlyActive

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsRecentlyActive(v interface{})`

SetIsRecentlyActive sets IsRecentlyActive field to given value.


### SetIsRecentlyActiveNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsRecentlyActiveNil(b bool)`

 SetIsRecentlyActiveNil sets the value for IsRecentlyActive to be an explicit nil

### UnsetIsRecentlyActive
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetIsRecentlyActive()`

UnsetIsRecentlyActive ensures that no value is present for IsRecentlyActive, not even an explicit nil
### GetIsAnnouncementOnly

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsAnnouncementOnly() interface{}`

GetIsAnnouncementOnly returns the IsAnnouncementOnly field if non-nil, zero value otherwise.

### GetIsAnnouncementOnlyOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsAnnouncementOnlyOk() (*interface{}, bool)`

GetIsAnnouncementOnlyOk returns a tuple with the IsAnnouncementOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnouncementOnly

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsAnnouncementOnly(v interface{})`

SetIsAnnouncementOnly sets IsAnnouncementOnly field to given value.


### SetIsAnnouncementOnlyNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsAnnouncementOnlyNil(b bool)`

 SetIsAnnouncementOnlyNil sets the value for IsAnnouncementOnly to be an explicit nil

### UnsetIsAnnouncementOnly
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetIsAnnouncementOnly()`

UnsetIsAnnouncementOnly ensures that no value is present for IsAnnouncementOnly, not even an explicit nil
### GetCanAddSubscribersGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanAddSubscribersGroup() interface{}`

GetCanAddSubscribersGroup returns the CanAddSubscribersGroup field if non-nil, zero value otherwise.

### GetCanAddSubscribersGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanAddSubscribersGroupOk() (*interface{}, bool)`

GetCanAddSubscribersGroupOk returns a tuple with the CanAddSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddSubscribersGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanAddSubscribersGroup(v interface{})`

SetCanAddSubscribersGroup sets CanAddSubscribersGroup field to given value.

### HasCanAddSubscribersGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) HasCanAddSubscribersGroup() bool`

HasCanAddSubscribersGroup returns a boolean if a field has been set.

### SetCanAddSubscribersGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanAddSubscribersGroupNil(b bool)`

 SetCanAddSubscribersGroupNil sets the value for CanAddSubscribersGroup to be an explicit nil

### UnsetCanAddSubscribersGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanAddSubscribersGroup()`

UnsetCanAddSubscribersGroup ensures that no value is present for CanAddSubscribersGroup, not even an explicit nil
### GetCanRemoveSubscribersGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanRemoveSubscribersGroup() interface{}`

GetCanRemoveSubscribersGroup returns the CanRemoveSubscribersGroup field if non-nil, zero value otherwise.

### GetCanRemoveSubscribersGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanRemoveSubscribersGroupOk() (*interface{}, bool)`

GetCanRemoveSubscribersGroupOk returns a tuple with the CanRemoveSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemoveSubscribersGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanRemoveSubscribersGroup(v interface{})`

SetCanRemoveSubscribersGroup sets CanRemoveSubscribersGroup field to given value.


### SetCanRemoveSubscribersGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanRemoveSubscribersGroupNil(b bool)`

 SetCanRemoveSubscribersGroupNil sets the value for CanRemoveSubscribersGroup to be an explicit nil

### UnsetCanRemoveSubscribersGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanRemoveSubscribersGroup()`

UnsetCanRemoveSubscribersGroup ensures that no value is present for CanRemoveSubscribersGroup, not even an explicit nil
### GetCanAdministerChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanAdministerChannelGroup() interface{}`

GetCanAdministerChannelGroup returns the CanAdministerChannelGroup field if non-nil, zero value otherwise.

### GetCanAdministerChannelGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanAdministerChannelGroupOk() (*interface{}, bool)`

GetCanAdministerChannelGroupOk returns a tuple with the CanAdministerChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdministerChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanAdministerChannelGroup(v interface{})`

SetCanAdministerChannelGroup sets CanAdministerChannelGroup field to given value.

### HasCanAdministerChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) HasCanAdministerChannelGroup() bool`

HasCanAdministerChannelGroup returns a boolean if a field has been set.

### SetCanAdministerChannelGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanAdministerChannelGroupNil(b bool)`

 SetCanAdministerChannelGroupNil sets the value for CanAdministerChannelGroup to be an explicit nil

### UnsetCanAdministerChannelGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanAdministerChannelGroup()`

UnsetCanAdministerChannelGroup ensures that no value is present for CanAdministerChannelGroup, not even an explicit nil
### GetCanDeleteAnyMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanDeleteAnyMessageGroup() interface{}`

GetCanDeleteAnyMessageGroup returns the CanDeleteAnyMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteAnyMessageGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanDeleteAnyMessageGroupOk() (*interface{}, bool)`

GetCanDeleteAnyMessageGroupOk returns a tuple with the CanDeleteAnyMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteAnyMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanDeleteAnyMessageGroup(v interface{})`

SetCanDeleteAnyMessageGroup sets CanDeleteAnyMessageGroup field to given value.

### HasCanDeleteAnyMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) HasCanDeleteAnyMessageGroup() bool`

HasCanDeleteAnyMessageGroup returns a boolean if a field has been set.

### SetCanDeleteAnyMessageGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanDeleteAnyMessageGroupNil(b bool)`

 SetCanDeleteAnyMessageGroupNil sets the value for CanDeleteAnyMessageGroup to be an explicit nil

### UnsetCanDeleteAnyMessageGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanDeleteAnyMessageGroup()`

UnsetCanDeleteAnyMessageGroup ensures that no value is present for CanDeleteAnyMessageGroup, not even an explicit nil
### GetCanDeleteOwnMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanDeleteOwnMessageGroup() interface{}`

GetCanDeleteOwnMessageGroup returns the CanDeleteOwnMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteOwnMessageGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanDeleteOwnMessageGroupOk() (*interface{}, bool)`

GetCanDeleteOwnMessageGroupOk returns a tuple with the CanDeleteOwnMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteOwnMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanDeleteOwnMessageGroup(v interface{})`

SetCanDeleteOwnMessageGroup sets CanDeleteOwnMessageGroup field to given value.

### HasCanDeleteOwnMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) HasCanDeleteOwnMessageGroup() bool`

HasCanDeleteOwnMessageGroup returns a boolean if a field has been set.

### SetCanDeleteOwnMessageGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanDeleteOwnMessageGroupNil(b bool)`

 SetCanDeleteOwnMessageGroupNil sets the value for CanDeleteOwnMessageGroup to be an explicit nil

### UnsetCanDeleteOwnMessageGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanDeleteOwnMessageGroup()`

UnsetCanDeleteOwnMessageGroup ensures that no value is present for CanDeleteOwnMessageGroup, not even an explicit nil
### GetCanMoveMessagesOutOfChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanMoveMessagesOutOfChannelGroup() interface{}`

GetCanMoveMessagesOutOfChannelGroup returns the CanMoveMessagesOutOfChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesOutOfChannelGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanMoveMessagesOutOfChannelGroupOk() (*interface{}, bool)`

GetCanMoveMessagesOutOfChannelGroupOk returns a tuple with the CanMoveMessagesOutOfChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesOutOfChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanMoveMessagesOutOfChannelGroup(v interface{})`

SetCanMoveMessagesOutOfChannelGroup sets CanMoveMessagesOutOfChannelGroup field to given value.

### HasCanMoveMessagesOutOfChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) HasCanMoveMessagesOutOfChannelGroup() bool`

HasCanMoveMessagesOutOfChannelGroup returns a boolean if a field has been set.

### SetCanMoveMessagesOutOfChannelGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanMoveMessagesOutOfChannelGroupNil(b bool)`

 SetCanMoveMessagesOutOfChannelGroupNil sets the value for CanMoveMessagesOutOfChannelGroup to be an explicit nil

### UnsetCanMoveMessagesOutOfChannelGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanMoveMessagesOutOfChannelGroup()`

UnsetCanMoveMessagesOutOfChannelGroup ensures that no value is present for CanMoveMessagesOutOfChannelGroup, not even an explicit nil
### GetCanMoveMessagesWithinChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanMoveMessagesWithinChannelGroup() interface{}`

GetCanMoveMessagesWithinChannelGroup returns the CanMoveMessagesWithinChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesWithinChannelGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanMoveMessagesWithinChannelGroupOk() (*interface{}, bool)`

GetCanMoveMessagesWithinChannelGroupOk returns a tuple with the CanMoveMessagesWithinChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesWithinChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanMoveMessagesWithinChannelGroup(v interface{})`

SetCanMoveMessagesWithinChannelGroup sets CanMoveMessagesWithinChannelGroup field to given value.

### HasCanMoveMessagesWithinChannelGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) HasCanMoveMessagesWithinChannelGroup() bool`

HasCanMoveMessagesWithinChannelGroup returns a boolean if a field has been set.

### SetCanMoveMessagesWithinChannelGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanMoveMessagesWithinChannelGroupNil(b bool)`

 SetCanMoveMessagesWithinChannelGroupNil sets the value for CanMoveMessagesWithinChannelGroup to be an explicit nil

### UnsetCanMoveMessagesWithinChannelGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanMoveMessagesWithinChannelGroup()`

UnsetCanMoveMessagesWithinChannelGroup ensures that no value is present for CanMoveMessagesWithinChannelGroup, not even an explicit nil
### GetCanSendMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanSendMessageGroup() interface{}`

GetCanSendMessageGroup returns the CanSendMessageGroup field if non-nil, zero value otherwise.

### GetCanSendMessageGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanSendMessageGroupOk() (*interface{}, bool)`

GetCanSendMessageGroupOk returns a tuple with the CanSendMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanSendMessageGroup(v interface{})`

SetCanSendMessageGroup sets CanSendMessageGroup field to given value.

### HasCanSendMessageGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) HasCanSendMessageGroup() bool`

HasCanSendMessageGroup returns a boolean if a field has been set.

### SetCanSendMessageGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanSendMessageGroupNil(b bool)`

 SetCanSendMessageGroupNil sets the value for CanSendMessageGroup to be an explicit nil

### UnsetCanSendMessageGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanSendMessageGroup()`

UnsetCanSendMessageGroup ensures that no value is present for CanSendMessageGroup, not even an explicit nil
### GetCanSubscribeGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanSubscribeGroup() interface{}`

GetCanSubscribeGroup returns the CanSubscribeGroup field if non-nil, zero value otherwise.

### GetCanSubscribeGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanSubscribeGroupOk() (*interface{}, bool)`

GetCanSubscribeGroupOk returns a tuple with the CanSubscribeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSubscribeGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanSubscribeGroup(v interface{})`

SetCanSubscribeGroup sets CanSubscribeGroup field to given value.


### SetCanSubscribeGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanSubscribeGroupNil(b bool)`

 SetCanSubscribeGroupNil sets the value for CanSubscribeGroup to be an explicit nil

### UnsetCanSubscribeGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanSubscribeGroup()`

UnsetCanSubscribeGroup ensures that no value is present for CanSubscribeGroup, not even an explicit nil
### GetCanResolveTopicsGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanResolveTopicsGroup() interface{}`

GetCanResolveTopicsGroup returns the CanResolveTopicsGroup field if non-nil, zero value otherwise.

### GetCanResolveTopicsGroupOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetCanResolveTopicsGroupOk() (*interface{}, bool)`

GetCanResolveTopicsGroupOk returns a tuple with the CanResolveTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResolveTopicsGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanResolveTopicsGroup(v interface{})`

SetCanResolveTopicsGroup sets CanResolveTopicsGroup field to given value.

### HasCanResolveTopicsGroup

`func (o *GetStreams200ResponseAllOfStreamsInner) HasCanResolveTopicsGroup() bool`

HasCanResolveTopicsGroup returns a boolean if a field has been set.

### SetCanResolveTopicsGroupNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetCanResolveTopicsGroupNil(b bool)`

 SetCanResolveTopicsGroupNil sets the value for CanResolveTopicsGroup to be an explicit nil

### UnsetCanResolveTopicsGroup
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetCanResolveTopicsGroup()`

UnsetCanResolveTopicsGroup ensures that no value is present for CanResolveTopicsGroup, not even an explicit nil
### GetSubscriberCount

`func (o *GetStreams200ResponseAllOfStreamsInner) GetSubscriberCount() interface{}`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetSubscriberCountOk() (*interface{}, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *GetStreams200ResponseAllOfStreamsInner) SetSubscriberCount(v interface{})`

SetSubscriberCount sets SubscriberCount field to given value.


### SetSubscriberCountNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetSubscriberCountNil(b bool)`

 SetSubscriberCountNil sets the value for SubscriberCount to be an explicit nil

### UnsetSubscriberCount
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetSubscriberCount()`

UnsetSubscriberCount ensures that no value is present for SubscriberCount, not even an explicit nil
### GetStreamWeeklyTraffic

`func (o *GetStreams200ResponseAllOfStreamsInner) GetStreamWeeklyTraffic() int32`

GetStreamWeeklyTraffic returns the StreamWeeklyTraffic field if non-nil, zero value otherwise.

### GetStreamWeeklyTrafficOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetStreamWeeklyTrafficOk() (*int32, bool)`

GetStreamWeeklyTrafficOk returns a tuple with the StreamWeeklyTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamWeeklyTraffic

`func (o *GetStreams200ResponseAllOfStreamsInner) SetStreamWeeklyTraffic(v int32)`

SetStreamWeeklyTraffic sets StreamWeeklyTraffic field to given value.


### SetStreamWeeklyTrafficNil

`func (o *GetStreams200ResponseAllOfStreamsInner) SetStreamWeeklyTrafficNil(b bool)`

 SetStreamWeeklyTrafficNil sets the value for StreamWeeklyTraffic to be an explicit nil

### UnsetStreamWeeklyTraffic
`func (o *GetStreams200ResponseAllOfStreamsInner) UnsetStreamWeeklyTraffic()`

UnsetStreamWeeklyTraffic ensures that no value is present for StreamWeeklyTraffic, not even an explicit nil
### GetIsDefault

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetStreams200ResponseAllOfStreamsInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetStreams200ResponseAllOfStreamsInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetStreams200ResponseAllOfStreamsInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


