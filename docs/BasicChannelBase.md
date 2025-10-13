# BasicChannelBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamId** | Pointer to **int32** | The unique ID of the channel.  | [optional] 
**Name** | Pointer to **string** | The name of the channel.  | [optional] 
**IsArchived** | Pointer to **bool** | A boolean indicating whether the channel is [archived](/help/archive-a-channel).  **Changes**: New in Zulip 10.0 (feature level 315). Previously, this endpoint never returned archived channels.  | [optional] 
**Description** | Pointer to **string** | The short description of the channel in [Zulip-flavored Markdown](/help/format-your-message-using-markdown) format, intended to be used to prepopulate UI for editing a channel&#39;s description.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**DateCreated** | Pointer to **int32** | The UNIX timestamp for when the channel was created, in UTC seconds.  **Changes**: New in Zulip 4.0 (feature level 30).  | [optional] 
**CreatorId** | Pointer to **NullableInt32** | The ID of the user who created this channel.  A &#x60;null&#x60; value means the channel has no recorded creator, which is often because the channel is very old, or because it was created via a data import tool or [management command][management-commands].  **Changes**: New in Zulip 9.0 (feature level 254).  [management-commands]: https://zulip.readthedocs.io/en/latest/production/management-commands.html  | [optional] 
**InviteOnly** | Pointer to **bool** | Specifies whether the channel is private or not. Only people who have been invited can access a private channel.  | [optional] 
**RenderedDescription** | Pointer to **string** | The short description of the channel rendered as HTML, intended to be used when displaying the channel description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.  | [optional] 
**IsWebPublic** | Pointer to **bool** | Whether the channel has been configured to allow unauthenticated access to its message history from the web.  **Changes**: New in Zulip 2.1.0.  | [optional] 
**StreamPostPolicy** | Pointer to **int32** | A deprecated representation of a superset of the users who have permission to post messages to the channel available for backwards-compatibility. Clients should use &#x60;can_send_message_group&#x60; instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 1 &#x3D; Any user can post. - 2 &#x3D; Only administrators can post. - 3 &#x3D; Only [full members][calc-full-member] can post. - 4 &#x3D; Only moderators can post.  **Changes**: Deprecated in Zulip 10.0 (feature level 333) and replaced by &#x60;can_send_message_group&#x60;, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  New in Zulip 3.0 (feature level 1), replacing the previous &#x60;is_announcement_only&#x60; boolean.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**MessageRetentionDays** | Pointer to **NullableInt32** | Number of days that messages sent to this channel will be stored before being automatically deleted by the [message retention policy](/help/message-retention-policy). There are two special values:  - &#x60;null&#x60;, the default, means the channel will inherit the organization   level setting. - &#x60;-1&#x60; encodes retaining messages in this channel forever.  **Changes**: New in Zulip 3.0 (feature level 17).  | [optional] 
**HistoryPublicToSubscribers** | Pointer to **bool** | Whether the history of the channel is public to its subscribers.  Currently always true for public channels (i.e. &#x60;\&quot;invite_only\&quot;: false&#x60; implies &#x60;\&quot;history_public_to_subscribers\&quot;: true&#x60;), but clients should not make that assumption, as we may change that behavior in the future.  | [optional] 
**TopicsPolicy** | Pointer to [**TopicsPolicy**](TopicsPolicy.md) |  | [optional] 
**FirstMessageId** | Pointer to **NullableInt32** | The ID of the first message in the channel.  Intended to help clients determine whether they need to display UI like the \&quot;show all topics\&quot; widget that would suggest the channel has older history that can be accessed.  Is &#x60;null&#x60; for channels with no message history.  **Changes**: New in Zulip 2.1.0.  | [optional] 
**FolderId** | Pointer to **NullableInt32** | The ID of the folder to which the channel belongs.  Is &#x60;null&#x60; if channel does not belong to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).  | [optional] 
**IsRecentlyActive** | Pointer to **bool** | Whether the channel has recent message activity. Clients should use this to implement [hide inactive channels](/help/manage-inactive-channels) if &#x60;demote_inactive_streams&#x60; is enabled.  **Changes**: New in Zulip 10.0 (feature level 323). Previously, clients implemented the demote_inactive_streams from local message history, resulting in a choppy loading experience.  | [optional] 
**IsAnnouncementOnly** | Pointer to **bool** | Whether the given channel is announcement only or not.  **Changes**: Deprecated in Zulip 3.0 (feature level 1). Clients should use &#x60;stream_post_policy&#x60; instead.  | [optional] 
**CanAddSubscribersGroup** | Pointer to [**ChannelCanAddSubscribersGroup**](ChannelCanAddSubscribersGroup.md) |  | [optional] 
**CanRemoveSubscribersGroup** | Pointer to [**CanRemoveSubscribersGroup**](CanRemoveSubscribersGroup.md) |  | [optional] 
**CanAdministerChannelGroup** | Pointer to [**CanAdministerChannelGroup**](CanAdministerChannelGroup.md) |  | [optional] 
**CanDeleteAnyMessageGroup** | Pointer to [**CanDeleteAnyMessageGroup**](CanDeleteAnyMessageGroup.md) |  | [optional] 
**CanDeleteOwnMessageGroup** | Pointer to [**CanDeleteOwnMessageGroup**](CanDeleteOwnMessageGroup.md) |  | [optional] 
**CanMoveMessagesOutOfChannelGroup** | Pointer to [**CanMoveMessagesOutOfChannelGroup**](CanMoveMessagesOutOfChannelGroup.md) |  | [optional] 
**CanMoveMessagesWithinChannelGroup** | Pointer to [**CanMoveMessagesWithinChannelGroup**](CanMoveMessagesWithinChannelGroup.md) |  | [optional] 
**CanSendMessageGroup** | Pointer to [**CanSendMessageGroup**](CanSendMessageGroup.md) |  | [optional] 
**CanSubscribeGroup** | Pointer to [**CanSubscribeGroup**](CanSubscribeGroup.md) |  | [optional] 
**CanResolveTopicsGroup** | Pointer to [**CanResolveTopicsGroup**](CanResolveTopicsGroup.md) |  | [optional] 
**SubscriberCount** | Pointer to **float32** | The total number of non-deactivated users (including bots) who are subscribed to the channel. Clients are responsible for updating this value using &#x60;peer_add&#x60; and &#x60;peer_remove&#x60; events.  The server&#39;s internals cannot guarantee this value is correctly synced with &#x60;peer_add&#x60; and &#x60;peer_remove&#x60; events for the channel. As a result, if a (rare) race occurs between a change in the channel&#39;s subscribers and fetching this value, it is possible for a client that is correctly following the events protocol to end up with a permanently off-by-one error in the channel&#39;s subscriber count.  Clients are recommended to fetch full subscriber data for a channel in contexts where it is important to avoid this risk. The official web application, for example, uses this field primarily while waiting to fetch a given channel&#39;s full subscriber list from the server.  **Changes**: New in Zulip 11.0 (feature level 394).  | [optional] 

## Methods

### NewBasicChannelBase

`func NewBasicChannelBase() *BasicChannelBase`

NewBasicChannelBase instantiates a new BasicChannelBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicChannelBaseWithDefaults

`func NewBasicChannelBaseWithDefaults() *BasicChannelBase`

NewBasicChannelBaseWithDefaults instantiates a new BasicChannelBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamId

`func (o *BasicChannelBase) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *BasicChannelBase) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *BasicChannelBase) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *BasicChannelBase) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetName

`func (o *BasicChannelBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicChannelBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicChannelBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicChannelBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsArchived

`func (o *BasicChannelBase) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *BasicChannelBase) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *BasicChannelBase) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *BasicChannelBase) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetDescription

`func (o *BasicChannelBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BasicChannelBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BasicChannelBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BasicChannelBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDateCreated

`func (o *BasicChannelBase) GetDateCreated() int32`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BasicChannelBase) GetDateCreatedOk() (*int32, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BasicChannelBase) SetDateCreated(v int32)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *BasicChannelBase) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetCreatorId

`func (o *BasicChannelBase) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *BasicChannelBase) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *BasicChannelBase) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *BasicChannelBase) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *BasicChannelBase) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *BasicChannelBase) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetInviteOnly

`func (o *BasicChannelBase) GetInviteOnly() bool`

GetInviteOnly returns the InviteOnly field if non-nil, zero value otherwise.

### GetInviteOnlyOk

`func (o *BasicChannelBase) GetInviteOnlyOk() (*bool, bool)`

GetInviteOnlyOk returns a tuple with the InviteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteOnly

`func (o *BasicChannelBase) SetInviteOnly(v bool)`

SetInviteOnly sets InviteOnly field to given value.

### HasInviteOnly

`func (o *BasicChannelBase) HasInviteOnly() bool`

HasInviteOnly returns a boolean if a field has been set.

### GetRenderedDescription

`func (o *BasicChannelBase) GetRenderedDescription() string`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *BasicChannelBase) GetRenderedDescriptionOk() (*string, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *BasicChannelBase) SetRenderedDescription(v string)`

SetRenderedDescription sets RenderedDescription field to given value.

### HasRenderedDescription

`func (o *BasicChannelBase) HasRenderedDescription() bool`

HasRenderedDescription returns a boolean if a field has been set.

### GetIsWebPublic

`func (o *BasicChannelBase) GetIsWebPublic() bool`

GetIsWebPublic returns the IsWebPublic field if non-nil, zero value otherwise.

### GetIsWebPublicOk

`func (o *BasicChannelBase) GetIsWebPublicOk() (*bool, bool)`

GetIsWebPublicOk returns a tuple with the IsWebPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebPublic

`func (o *BasicChannelBase) SetIsWebPublic(v bool)`

SetIsWebPublic sets IsWebPublic field to given value.

### HasIsWebPublic

`func (o *BasicChannelBase) HasIsWebPublic() bool`

HasIsWebPublic returns a boolean if a field has been set.

### GetStreamPostPolicy

`func (o *BasicChannelBase) GetStreamPostPolicy() int32`

GetStreamPostPolicy returns the StreamPostPolicy field if non-nil, zero value otherwise.

### GetStreamPostPolicyOk

`func (o *BasicChannelBase) GetStreamPostPolicyOk() (*int32, bool)`

GetStreamPostPolicyOk returns a tuple with the StreamPostPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamPostPolicy

`func (o *BasicChannelBase) SetStreamPostPolicy(v int32)`

SetStreamPostPolicy sets StreamPostPolicy field to given value.

### HasStreamPostPolicy

`func (o *BasicChannelBase) HasStreamPostPolicy() bool`

HasStreamPostPolicy returns a boolean if a field has been set.

### GetMessageRetentionDays

`func (o *BasicChannelBase) GetMessageRetentionDays() int32`

GetMessageRetentionDays returns the MessageRetentionDays field if non-nil, zero value otherwise.

### GetMessageRetentionDaysOk

`func (o *BasicChannelBase) GetMessageRetentionDaysOk() (*int32, bool)`

GetMessageRetentionDaysOk returns a tuple with the MessageRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRetentionDays

`func (o *BasicChannelBase) SetMessageRetentionDays(v int32)`

SetMessageRetentionDays sets MessageRetentionDays field to given value.

### HasMessageRetentionDays

`func (o *BasicChannelBase) HasMessageRetentionDays() bool`

HasMessageRetentionDays returns a boolean if a field has been set.

### SetMessageRetentionDaysNil

`func (o *BasicChannelBase) SetMessageRetentionDaysNil(b bool)`

 SetMessageRetentionDaysNil sets the value for MessageRetentionDays to be an explicit nil

### UnsetMessageRetentionDays
`func (o *BasicChannelBase) UnsetMessageRetentionDays()`

UnsetMessageRetentionDays ensures that no value is present for MessageRetentionDays, not even an explicit nil
### GetHistoryPublicToSubscribers

`func (o *BasicChannelBase) GetHistoryPublicToSubscribers() bool`

GetHistoryPublicToSubscribers returns the HistoryPublicToSubscribers field if non-nil, zero value otherwise.

### GetHistoryPublicToSubscribersOk

`func (o *BasicChannelBase) GetHistoryPublicToSubscribersOk() (*bool, bool)`

GetHistoryPublicToSubscribersOk returns a tuple with the HistoryPublicToSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPublicToSubscribers

`func (o *BasicChannelBase) SetHistoryPublicToSubscribers(v bool)`

SetHistoryPublicToSubscribers sets HistoryPublicToSubscribers field to given value.

### HasHistoryPublicToSubscribers

`func (o *BasicChannelBase) HasHistoryPublicToSubscribers() bool`

HasHistoryPublicToSubscribers returns a boolean if a field has been set.

### GetTopicsPolicy

`func (o *BasicChannelBase) GetTopicsPolicy() TopicsPolicy`

GetTopicsPolicy returns the TopicsPolicy field if non-nil, zero value otherwise.

### GetTopicsPolicyOk

`func (o *BasicChannelBase) GetTopicsPolicyOk() (*TopicsPolicy, bool)`

GetTopicsPolicyOk returns a tuple with the TopicsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsPolicy

`func (o *BasicChannelBase) SetTopicsPolicy(v TopicsPolicy)`

SetTopicsPolicy sets TopicsPolicy field to given value.

### HasTopicsPolicy

`func (o *BasicChannelBase) HasTopicsPolicy() bool`

HasTopicsPolicy returns a boolean if a field has been set.

### GetFirstMessageId

`func (o *BasicChannelBase) GetFirstMessageId() int32`

GetFirstMessageId returns the FirstMessageId field if non-nil, zero value otherwise.

### GetFirstMessageIdOk

`func (o *BasicChannelBase) GetFirstMessageIdOk() (*int32, bool)`

GetFirstMessageIdOk returns a tuple with the FirstMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessageId

`func (o *BasicChannelBase) SetFirstMessageId(v int32)`

SetFirstMessageId sets FirstMessageId field to given value.

### HasFirstMessageId

`func (o *BasicChannelBase) HasFirstMessageId() bool`

HasFirstMessageId returns a boolean if a field has been set.

### SetFirstMessageIdNil

`func (o *BasicChannelBase) SetFirstMessageIdNil(b bool)`

 SetFirstMessageIdNil sets the value for FirstMessageId to be an explicit nil

### UnsetFirstMessageId
`func (o *BasicChannelBase) UnsetFirstMessageId()`

UnsetFirstMessageId ensures that no value is present for FirstMessageId, not even an explicit nil
### GetFolderId

`func (o *BasicChannelBase) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BasicChannelBase) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BasicChannelBase) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *BasicChannelBase) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *BasicChannelBase) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *BasicChannelBase) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetIsRecentlyActive

`func (o *BasicChannelBase) GetIsRecentlyActive() bool`

GetIsRecentlyActive returns the IsRecentlyActive field if non-nil, zero value otherwise.

### GetIsRecentlyActiveOk

`func (o *BasicChannelBase) GetIsRecentlyActiveOk() (*bool, bool)`

GetIsRecentlyActiveOk returns a tuple with the IsRecentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecentlyActive

`func (o *BasicChannelBase) SetIsRecentlyActive(v bool)`

SetIsRecentlyActive sets IsRecentlyActive field to given value.

### HasIsRecentlyActive

`func (o *BasicChannelBase) HasIsRecentlyActive() bool`

HasIsRecentlyActive returns a boolean if a field has been set.

### GetIsAnnouncementOnly

`func (o *BasicChannelBase) GetIsAnnouncementOnly() bool`

GetIsAnnouncementOnly returns the IsAnnouncementOnly field if non-nil, zero value otherwise.

### GetIsAnnouncementOnlyOk

`func (o *BasicChannelBase) GetIsAnnouncementOnlyOk() (*bool, bool)`

GetIsAnnouncementOnlyOk returns a tuple with the IsAnnouncementOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnouncementOnly

`func (o *BasicChannelBase) SetIsAnnouncementOnly(v bool)`

SetIsAnnouncementOnly sets IsAnnouncementOnly field to given value.

### HasIsAnnouncementOnly

`func (o *BasicChannelBase) HasIsAnnouncementOnly() bool`

HasIsAnnouncementOnly returns a boolean if a field has been set.

### GetCanAddSubscribersGroup

`func (o *BasicChannelBase) GetCanAddSubscribersGroup() ChannelCanAddSubscribersGroup`

GetCanAddSubscribersGroup returns the CanAddSubscribersGroup field if non-nil, zero value otherwise.

### GetCanAddSubscribersGroupOk

`func (o *BasicChannelBase) GetCanAddSubscribersGroupOk() (*ChannelCanAddSubscribersGroup, bool)`

GetCanAddSubscribersGroupOk returns a tuple with the CanAddSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddSubscribersGroup

`func (o *BasicChannelBase) SetCanAddSubscribersGroup(v ChannelCanAddSubscribersGroup)`

SetCanAddSubscribersGroup sets CanAddSubscribersGroup field to given value.

### HasCanAddSubscribersGroup

`func (o *BasicChannelBase) HasCanAddSubscribersGroup() bool`

HasCanAddSubscribersGroup returns a boolean if a field has been set.

### GetCanRemoveSubscribersGroup

`func (o *BasicChannelBase) GetCanRemoveSubscribersGroup() CanRemoveSubscribersGroup`

GetCanRemoveSubscribersGroup returns the CanRemoveSubscribersGroup field if non-nil, zero value otherwise.

### GetCanRemoveSubscribersGroupOk

`func (o *BasicChannelBase) GetCanRemoveSubscribersGroupOk() (*CanRemoveSubscribersGroup, bool)`

GetCanRemoveSubscribersGroupOk returns a tuple with the CanRemoveSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemoveSubscribersGroup

`func (o *BasicChannelBase) SetCanRemoveSubscribersGroup(v CanRemoveSubscribersGroup)`

SetCanRemoveSubscribersGroup sets CanRemoveSubscribersGroup field to given value.

### HasCanRemoveSubscribersGroup

`func (o *BasicChannelBase) HasCanRemoveSubscribersGroup() bool`

HasCanRemoveSubscribersGroup returns a boolean if a field has been set.

### GetCanAdministerChannelGroup

`func (o *BasicChannelBase) GetCanAdministerChannelGroup() CanAdministerChannelGroup`

GetCanAdministerChannelGroup returns the CanAdministerChannelGroup field if non-nil, zero value otherwise.

### GetCanAdministerChannelGroupOk

`func (o *BasicChannelBase) GetCanAdministerChannelGroupOk() (*CanAdministerChannelGroup, bool)`

GetCanAdministerChannelGroupOk returns a tuple with the CanAdministerChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdministerChannelGroup

`func (o *BasicChannelBase) SetCanAdministerChannelGroup(v CanAdministerChannelGroup)`

SetCanAdministerChannelGroup sets CanAdministerChannelGroup field to given value.

### HasCanAdministerChannelGroup

`func (o *BasicChannelBase) HasCanAdministerChannelGroup() bool`

HasCanAdministerChannelGroup returns a boolean if a field has been set.

### GetCanDeleteAnyMessageGroup

`func (o *BasicChannelBase) GetCanDeleteAnyMessageGroup() CanDeleteAnyMessageGroup`

GetCanDeleteAnyMessageGroup returns the CanDeleteAnyMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteAnyMessageGroupOk

`func (o *BasicChannelBase) GetCanDeleteAnyMessageGroupOk() (*CanDeleteAnyMessageGroup, bool)`

GetCanDeleteAnyMessageGroupOk returns a tuple with the CanDeleteAnyMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteAnyMessageGroup

`func (o *BasicChannelBase) SetCanDeleteAnyMessageGroup(v CanDeleteAnyMessageGroup)`

SetCanDeleteAnyMessageGroup sets CanDeleteAnyMessageGroup field to given value.

### HasCanDeleteAnyMessageGroup

`func (o *BasicChannelBase) HasCanDeleteAnyMessageGroup() bool`

HasCanDeleteAnyMessageGroup returns a boolean if a field has been set.

### GetCanDeleteOwnMessageGroup

`func (o *BasicChannelBase) GetCanDeleteOwnMessageGroup() CanDeleteOwnMessageGroup`

GetCanDeleteOwnMessageGroup returns the CanDeleteOwnMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteOwnMessageGroupOk

`func (o *BasicChannelBase) GetCanDeleteOwnMessageGroupOk() (*CanDeleteOwnMessageGroup, bool)`

GetCanDeleteOwnMessageGroupOk returns a tuple with the CanDeleteOwnMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteOwnMessageGroup

`func (o *BasicChannelBase) SetCanDeleteOwnMessageGroup(v CanDeleteOwnMessageGroup)`

SetCanDeleteOwnMessageGroup sets CanDeleteOwnMessageGroup field to given value.

### HasCanDeleteOwnMessageGroup

`func (o *BasicChannelBase) HasCanDeleteOwnMessageGroup() bool`

HasCanDeleteOwnMessageGroup returns a boolean if a field has been set.

### GetCanMoveMessagesOutOfChannelGroup

`func (o *BasicChannelBase) GetCanMoveMessagesOutOfChannelGroup() CanMoveMessagesOutOfChannelGroup`

GetCanMoveMessagesOutOfChannelGroup returns the CanMoveMessagesOutOfChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesOutOfChannelGroupOk

`func (o *BasicChannelBase) GetCanMoveMessagesOutOfChannelGroupOk() (*CanMoveMessagesOutOfChannelGroup, bool)`

GetCanMoveMessagesOutOfChannelGroupOk returns a tuple with the CanMoveMessagesOutOfChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesOutOfChannelGroup

`func (o *BasicChannelBase) SetCanMoveMessagesOutOfChannelGroup(v CanMoveMessagesOutOfChannelGroup)`

SetCanMoveMessagesOutOfChannelGroup sets CanMoveMessagesOutOfChannelGroup field to given value.

### HasCanMoveMessagesOutOfChannelGroup

`func (o *BasicChannelBase) HasCanMoveMessagesOutOfChannelGroup() bool`

HasCanMoveMessagesOutOfChannelGroup returns a boolean if a field has been set.

### GetCanMoveMessagesWithinChannelGroup

`func (o *BasicChannelBase) GetCanMoveMessagesWithinChannelGroup() CanMoveMessagesWithinChannelGroup`

GetCanMoveMessagesWithinChannelGroup returns the CanMoveMessagesWithinChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesWithinChannelGroupOk

`func (o *BasicChannelBase) GetCanMoveMessagesWithinChannelGroupOk() (*CanMoveMessagesWithinChannelGroup, bool)`

GetCanMoveMessagesWithinChannelGroupOk returns a tuple with the CanMoveMessagesWithinChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesWithinChannelGroup

`func (o *BasicChannelBase) SetCanMoveMessagesWithinChannelGroup(v CanMoveMessagesWithinChannelGroup)`

SetCanMoveMessagesWithinChannelGroup sets CanMoveMessagesWithinChannelGroup field to given value.

### HasCanMoveMessagesWithinChannelGroup

`func (o *BasicChannelBase) HasCanMoveMessagesWithinChannelGroup() bool`

HasCanMoveMessagesWithinChannelGroup returns a boolean if a field has been set.

### GetCanSendMessageGroup

`func (o *BasicChannelBase) GetCanSendMessageGroup() CanSendMessageGroup`

GetCanSendMessageGroup returns the CanSendMessageGroup field if non-nil, zero value otherwise.

### GetCanSendMessageGroupOk

`func (o *BasicChannelBase) GetCanSendMessageGroupOk() (*CanSendMessageGroup, bool)`

GetCanSendMessageGroupOk returns a tuple with the CanSendMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendMessageGroup

`func (o *BasicChannelBase) SetCanSendMessageGroup(v CanSendMessageGroup)`

SetCanSendMessageGroup sets CanSendMessageGroup field to given value.

### HasCanSendMessageGroup

`func (o *BasicChannelBase) HasCanSendMessageGroup() bool`

HasCanSendMessageGroup returns a boolean if a field has been set.

### GetCanSubscribeGroup

`func (o *BasicChannelBase) GetCanSubscribeGroup() CanSubscribeGroup`

GetCanSubscribeGroup returns the CanSubscribeGroup field if non-nil, zero value otherwise.

### GetCanSubscribeGroupOk

`func (o *BasicChannelBase) GetCanSubscribeGroupOk() (*CanSubscribeGroup, bool)`

GetCanSubscribeGroupOk returns a tuple with the CanSubscribeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSubscribeGroup

`func (o *BasicChannelBase) SetCanSubscribeGroup(v CanSubscribeGroup)`

SetCanSubscribeGroup sets CanSubscribeGroup field to given value.

### HasCanSubscribeGroup

`func (o *BasicChannelBase) HasCanSubscribeGroup() bool`

HasCanSubscribeGroup returns a boolean if a field has been set.

### GetCanResolveTopicsGroup

`func (o *BasicChannelBase) GetCanResolveTopicsGroup() CanResolveTopicsGroup`

GetCanResolveTopicsGroup returns the CanResolveTopicsGroup field if non-nil, zero value otherwise.

### GetCanResolveTopicsGroupOk

`func (o *BasicChannelBase) GetCanResolveTopicsGroupOk() (*CanResolveTopicsGroup, bool)`

GetCanResolveTopicsGroupOk returns a tuple with the CanResolveTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResolveTopicsGroup

`func (o *BasicChannelBase) SetCanResolveTopicsGroup(v CanResolveTopicsGroup)`

SetCanResolveTopicsGroup sets CanResolveTopicsGroup field to given value.

### HasCanResolveTopicsGroup

`func (o *BasicChannelBase) HasCanResolveTopicsGroup() bool`

HasCanResolveTopicsGroup returns a boolean if a field has been set.

### GetSubscriberCount

`func (o *BasicChannelBase) GetSubscriberCount() float32`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *BasicChannelBase) GetSubscriberCountOk() (*float32, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *BasicChannelBase) SetSubscriberCount(v float32)`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *BasicChannelBase) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


