# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamId** | Pointer to **int32** | The unique ID of a channel.  | [optional] 
**Name** | Pointer to **string** | The name of a channel.  | [optional] 
**Description** | Pointer to **string** | The [description](/help/change-the-channel-description) of the channel in [Zulip-flavored Markdown](/help/format-your-message-using-markdown) format, intended to be used to prepopulate UI for editing a channel&#39;s description.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  See also &#x60;rendered_description&#x60;.  | [optional] 
**RenderedDescription** | Pointer to **string** | The [description](/help/change-the-channel-description) of the channel rendered as HTML, intended to be used when displaying the channel description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.  See also &#x60;description&#x60;.  | [optional] 
**DateCreated** | Pointer to **int32** | The UNIX timestamp for when the channel was created, in UTC seconds.  **Changes**: New in Zulip 4.0 (feature level 30).  | [optional] 
**CreatorId** | Pointer to **NullableInt32** | The ID of the user who created this channel.  A &#x60;null&#x60; value means the channel has no recorded creator, which is often because the channel is very old, or because it was created via a data import tool or [management command][management-commands].  **Changes**: New in Zulip 9.0 (feature level 254).  [management-commands]: https://zulip.readthedocs.io/en/latest/production/management-commands.html  | [optional] 
**InviteOnly** | Pointer to **bool** | Specifies whether the channel is private or not. Only people who have been invited can access a private channel.  | [optional] 
**Subscribers** | Pointer to **[]int32** | A list of user IDs of users who are also subscribed to a given channel. Included only if &#x60;include_subscribers&#x60; is &#x60;true&#x60;.  | [optional] 
**PartialSubscribers** | Pointer to **[]int32** | If [&#x60;include_subscribers&#x3D;\&quot;partial\&quot;&#x60;](/api/get-subscriptions#parameter-include_subscribers) was requested, the server may, at its discretion, send a &#x60;partial_subscribers&#x60; list rather than a &#x60;subscribers&#x60; list for channels with a large number of subscribers.  The &#x60;partial_subscribers&#x60; list contains an arbitrary subset of the channel&#39;s subscribers that is guaranteed to include all bot user subscribers as well as all users who have been active in the last 14 days, but otherwise can be chosen arbitrarily by the server.  **Changes**: New in Zulip 11.0 (feature level 412).  | [optional] 
**DesktopNotifications** | Pointer to **NullableBool** | A boolean specifying whether desktop notifications are enabled for the given channel.  A &#x60;null&#x60; value means the value of this setting should be inherited from the user-level default setting, &#x60;enable_stream_desktop_notifications&#x60;, for this channel.  | [optional] 
**EmailNotifications** | Pointer to **NullableBool** | A boolean specifying whether email notifications are enabled for the given channel.  A &#x60;null&#x60; value means the value of this setting should be inherited from the user-level default setting, &#x60;enable_stream_email_notifications&#x60;, for this channel.  | [optional] 
**WildcardMentionsNotify** | Pointer to **NullableBool** | A boolean specifying whether wildcard mentions trigger notifications as though they were personal mentions in this channel.  A &#x60;null&#x60; value means the value of this setting should be inherited from the user-level default setting, wildcard_mentions_notify, for this channel.  | [optional] 
**PushNotifications** | Pointer to **NullableBool** | A boolean specifying whether push notifications are enabled for the given channel.  A &#x60;null&#x60; value means the value of this setting should be inherited from the user-level default setting, &#x60;enable_stream_push_notifications&#x60;, for this channel.  | [optional] 
**AudibleNotifications** | Pointer to **NullableBool** | A boolean specifying whether audible notifications are enabled for the given channel.  A &#x60;null&#x60; value means the value of this setting should be inherited from the user-level default setting, &#x60;enable_stream_audible_notifications&#x60;, for this channel.  | [optional] 
**PinToTop** | Pointer to **bool** | A boolean specifying whether the given channel has been pinned to the top.  | [optional] 
**IsMuted** | Pointer to **bool** | Whether the user has muted the channel. Muted channels do not count towards your total unread count and do not show up in the &#x60;Combined feed&#x60; view (previously known as &#x60;All messages&#x60;).  **Changes**: Prior to Zulip 2.1.0, this feature was represented by the more confusingly named &#x60;in_home_view&#x60; (with the opposite value, &#x60;in_home_view&#x3D;!is_muted&#x60;).  | [optional] 
**InHomeView** | Pointer to **bool** | Legacy property for if the given channel is muted, with inverted meaning.  **Changes**: Deprecated in Zulip 2.1.0. Clients should use &#x60;is_muted&#x60; where available.  | [optional] 
**IsAnnouncementOnly** | Pointer to **bool** | Whether only organization administrators can post to the channel.  **Changes**: Deprecated in Zulip 3.0 (feature level 1). Clients should use &#x60;stream_post_policy&#x60; instead.  | [optional] 
**IsWebPublic** | Pointer to **bool** | Whether the channel has been configured to allow unauthenticated access to its message history from the web.  | [optional] 
**Color** | Pointer to **string** | The user&#39;s personal color for the channel.  | [optional] 
**StreamPostPolicy** | Pointer to **int32** | A deprecated representation of a superset of the users who have permission to post messages to the channel available for backwards-compatibility. Clients should use &#x60;can_send_message_group&#x60; instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 1 &#x3D; Any user can post. - 2 &#x3D; Only administrators can post. - 3 &#x3D; Only [full members][calc-full-member] can post. - 4 &#x3D; Only moderators can post.  **Changes**: Deprecated in Zulip 10.0 (feature level 333) and replaced by &#x60;can_send_message_group&#x60;, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  New in Zulip 3.0 (feature level 1), replacing the previous &#x60;is_announcement_only&#x60; boolean.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**MessageRetentionDays** | Pointer to **NullableInt32** | Number of days that messages sent to this channel will be stored before being automatically deleted by the [message retention policy](/help/message-retention-policy). There are two special values:  - &#x60;null&#x60;, the default, means the channel will inherit the organization   level setting. - &#x60;-1&#x60; encodes retaining messages in this channel forever.  **Changes**: New in Zulip 3.0 (feature level 17).  | [optional] 
**HistoryPublicToSubscribers** | Pointer to **bool** | Whether the history of the channel is public to its subscribers.  Currently always true for public channels (i.e. &#x60;\&quot;invite_only\&quot;: false&#x60; implies &#x60;\&quot;history_public_to_subscribers\&quot;: true&#x60;), but clients should not make that assumption, as we may change that behavior in the future.  | [optional] 
**FirstMessageId** | Pointer to **NullableInt32** | The ID of the first message in the channel.  Intended to help clients determine whether they need to display UI like the \&quot;show all topics\&quot; widget that would suggest the channel has older history that can be accessed.  Is &#x60;null&#x60; for channels with no message history.  | [optional] 
**FolderId** | Pointer to **NullableInt32** | The ID of the folder to which the channel belongs.  Is &#x60;null&#x60; if channel does not belong to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).  | [optional] 
**TopicsPolicy** | Pointer to [**TopicsPolicy**](TopicsPolicy.md) |  | [optional] 
**IsRecentlyActive** | Pointer to **bool** | Whether the channel has recent message activity. Clients should use this to implement [hiding inactive channels](/help/manage-inactive-channels).  **Changes**: New in Zulip 10.0 (feature level 323). Previously, clients implemented the demote_inactive_streams from local message history, resulting in a choppy loading experience.  | [optional] 
**StreamWeeklyTraffic** | Pointer to **NullableInt32** | The average number of messages sent to the channel per week, as estimated based on recent weeks, rounded to the nearest integer.  If &#x60;null&#x60;, the channel was recently created and there is insufficient data to estimate the average traffic.  | [optional] 
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
**IsArchived** | Pointer to **bool** | A boolean indicating whether the channel is [archived](/help/archive-a-channel).  **Changes**: New in Zulip 10.0 (feature level 315). Previously, subscriptions only included active channels. Note that some endpoints will never return archived channels unless the client declares explicit support for them via the &#x60;archived_channels&#x60; client capability.  | [optional] 
**SubscriberCount** | Pointer to **float32** | The total number of non-deactivated users (including bots) who are subscribed to the channel. Clients are responsible for updating this value using &#x60;peer_add&#x60; and &#x60;peer_remove&#x60; events.  The server&#39;s internals cannot guarantee this value is correctly synced with &#x60;peer_add&#x60; and &#x60;peer_remove&#x60; events for the channel. As a result, if a (rare) race occurs between a change in the channel&#39;s subscribers and fetching this value, it is possible for a client that is correctly following the events protocol to end up with a permanently off-by-one error in the channel&#39;s subscriber count.  Clients are recommended to fetch full subscriber data for a channel in contexts where it is important to avoid this risk. The official web application, for example, uses this field primarily while waiting to fetch a given channel&#39;s full subscriber list from the server.  **Changes**: New in Zulip 11.0 (feature level 394).  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamId

`func (o *Subscription) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *Subscription) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *Subscription) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *Subscription) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetName

`func (o *Subscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subscription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subscription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Subscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subscription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Subscription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRenderedDescription

`func (o *Subscription) GetRenderedDescription() string`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *Subscription) GetRenderedDescriptionOk() (*string, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *Subscription) SetRenderedDescription(v string)`

SetRenderedDescription sets RenderedDescription field to given value.

### HasRenderedDescription

`func (o *Subscription) HasRenderedDescription() bool`

HasRenderedDescription returns a boolean if a field has been set.

### GetDateCreated

`func (o *Subscription) GetDateCreated() int32`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Subscription) GetDateCreatedOk() (*int32, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Subscription) SetDateCreated(v int32)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Subscription) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetCreatorId

`func (o *Subscription) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Subscription) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Subscription) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Subscription) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *Subscription) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *Subscription) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetInviteOnly

`func (o *Subscription) GetInviteOnly() bool`

GetInviteOnly returns the InviteOnly field if non-nil, zero value otherwise.

### GetInviteOnlyOk

`func (o *Subscription) GetInviteOnlyOk() (*bool, bool)`

GetInviteOnlyOk returns a tuple with the InviteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteOnly

`func (o *Subscription) SetInviteOnly(v bool)`

SetInviteOnly sets InviteOnly field to given value.

### HasInviteOnly

`func (o *Subscription) HasInviteOnly() bool`

HasInviteOnly returns a boolean if a field has been set.

### GetSubscribers

`func (o *Subscription) GetSubscribers() []int32`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *Subscription) GetSubscribersOk() (*[]int32, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *Subscription) SetSubscribers(v []int32)`

SetSubscribers sets Subscribers field to given value.

### HasSubscribers

`func (o *Subscription) HasSubscribers() bool`

HasSubscribers returns a boolean if a field has been set.

### GetPartialSubscribers

`func (o *Subscription) GetPartialSubscribers() []int32`

GetPartialSubscribers returns the PartialSubscribers field if non-nil, zero value otherwise.

### GetPartialSubscribersOk

`func (o *Subscription) GetPartialSubscribersOk() (*[]int32, bool)`

GetPartialSubscribersOk returns a tuple with the PartialSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialSubscribers

`func (o *Subscription) SetPartialSubscribers(v []int32)`

SetPartialSubscribers sets PartialSubscribers field to given value.

### HasPartialSubscribers

`func (o *Subscription) HasPartialSubscribers() bool`

HasPartialSubscribers returns a boolean if a field has been set.

### GetDesktopNotifications

`func (o *Subscription) GetDesktopNotifications() bool`

GetDesktopNotifications returns the DesktopNotifications field if non-nil, zero value otherwise.

### GetDesktopNotificationsOk

`func (o *Subscription) GetDesktopNotificationsOk() (*bool, bool)`

GetDesktopNotificationsOk returns a tuple with the DesktopNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopNotifications

`func (o *Subscription) SetDesktopNotifications(v bool)`

SetDesktopNotifications sets DesktopNotifications field to given value.

### HasDesktopNotifications

`func (o *Subscription) HasDesktopNotifications() bool`

HasDesktopNotifications returns a boolean if a field has been set.

### SetDesktopNotificationsNil

`func (o *Subscription) SetDesktopNotificationsNil(b bool)`

 SetDesktopNotificationsNil sets the value for DesktopNotifications to be an explicit nil

### UnsetDesktopNotifications
`func (o *Subscription) UnsetDesktopNotifications()`

UnsetDesktopNotifications ensures that no value is present for DesktopNotifications, not even an explicit nil
### GetEmailNotifications

`func (o *Subscription) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *Subscription) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *Subscription) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *Subscription) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### SetEmailNotificationsNil

`func (o *Subscription) SetEmailNotificationsNil(b bool)`

 SetEmailNotificationsNil sets the value for EmailNotifications to be an explicit nil

### UnsetEmailNotifications
`func (o *Subscription) UnsetEmailNotifications()`

UnsetEmailNotifications ensures that no value is present for EmailNotifications, not even an explicit nil
### GetWildcardMentionsNotify

`func (o *Subscription) GetWildcardMentionsNotify() bool`

GetWildcardMentionsNotify returns the WildcardMentionsNotify field if non-nil, zero value otherwise.

### GetWildcardMentionsNotifyOk

`func (o *Subscription) GetWildcardMentionsNotifyOk() (*bool, bool)`

GetWildcardMentionsNotifyOk returns a tuple with the WildcardMentionsNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardMentionsNotify

`func (o *Subscription) SetWildcardMentionsNotify(v bool)`

SetWildcardMentionsNotify sets WildcardMentionsNotify field to given value.

### HasWildcardMentionsNotify

`func (o *Subscription) HasWildcardMentionsNotify() bool`

HasWildcardMentionsNotify returns a boolean if a field has been set.

### SetWildcardMentionsNotifyNil

`func (o *Subscription) SetWildcardMentionsNotifyNil(b bool)`

 SetWildcardMentionsNotifyNil sets the value for WildcardMentionsNotify to be an explicit nil

### UnsetWildcardMentionsNotify
`func (o *Subscription) UnsetWildcardMentionsNotify()`

UnsetWildcardMentionsNotify ensures that no value is present for WildcardMentionsNotify, not even an explicit nil
### GetPushNotifications

`func (o *Subscription) GetPushNotifications() bool`

GetPushNotifications returns the PushNotifications field if non-nil, zero value otherwise.

### GetPushNotificationsOk

`func (o *Subscription) GetPushNotificationsOk() (*bool, bool)`

GetPushNotificationsOk returns a tuple with the PushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotifications

`func (o *Subscription) SetPushNotifications(v bool)`

SetPushNotifications sets PushNotifications field to given value.

### HasPushNotifications

`func (o *Subscription) HasPushNotifications() bool`

HasPushNotifications returns a boolean if a field has been set.

### SetPushNotificationsNil

`func (o *Subscription) SetPushNotificationsNil(b bool)`

 SetPushNotificationsNil sets the value for PushNotifications to be an explicit nil

### UnsetPushNotifications
`func (o *Subscription) UnsetPushNotifications()`

UnsetPushNotifications ensures that no value is present for PushNotifications, not even an explicit nil
### GetAudibleNotifications

`func (o *Subscription) GetAudibleNotifications() bool`

GetAudibleNotifications returns the AudibleNotifications field if non-nil, zero value otherwise.

### GetAudibleNotificationsOk

`func (o *Subscription) GetAudibleNotificationsOk() (*bool, bool)`

GetAudibleNotificationsOk returns a tuple with the AudibleNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudibleNotifications

`func (o *Subscription) SetAudibleNotifications(v bool)`

SetAudibleNotifications sets AudibleNotifications field to given value.

### HasAudibleNotifications

`func (o *Subscription) HasAudibleNotifications() bool`

HasAudibleNotifications returns a boolean if a field has been set.

### SetAudibleNotificationsNil

`func (o *Subscription) SetAudibleNotificationsNil(b bool)`

 SetAudibleNotificationsNil sets the value for AudibleNotifications to be an explicit nil

### UnsetAudibleNotifications
`func (o *Subscription) UnsetAudibleNotifications()`

UnsetAudibleNotifications ensures that no value is present for AudibleNotifications, not even an explicit nil
### GetPinToTop

`func (o *Subscription) GetPinToTop() bool`

GetPinToTop returns the PinToTop field if non-nil, zero value otherwise.

### GetPinToTopOk

`func (o *Subscription) GetPinToTopOk() (*bool, bool)`

GetPinToTopOk returns a tuple with the PinToTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinToTop

`func (o *Subscription) SetPinToTop(v bool)`

SetPinToTop sets PinToTop field to given value.

### HasPinToTop

`func (o *Subscription) HasPinToTop() bool`

HasPinToTop returns a boolean if a field has been set.

### GetIsMuted

`func (o *Subscription) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *Subscription) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *Subscription) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *Subscription) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetInHomeView

`func (o *Subscription) GetInHomeView() bool`

GetInHomeView returns the InHomeView field if non-nil, zero value otherwise.

### GetInHomeViewOk

`func (o *Subscription) GetInHomeViewOk() (*bool, bool)`

GetInHomeViewOk returns a tuple with the InHomeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInHomeView

`func (o *Subscription) SetInHomeView(v bool)`

SetInHomeView sets InHomeView field to given value.

### HasInHomeView

`func (o *Subscription) HasInHomeView() bool`

HasInHomeView returns a boolean if a field has been set.

### GetIsAnnouncementOnly

`func (o *Subscription) GetIsAnnouncementOnly() bool`

GetIsAnnouncementOnly returns the IsAnnouncementOnly field if non-nil, zero value otherwise.

### GetIsAnnouncementOnlyOk

`func (o *Subscription) GetIsAnnouncementOnlyOk() (*bool, bool)`

GetIsAnnouncementOnlyOk returns a tuple with the IsAnnouncementOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnouncementOnly

`func (o *Subscription) SetIsAnnouncementOnly(v bool)`

SetIsAnnouncementOnly sets IsAnnouncementOnly field to given value.

### HasIsAnnouncementOnly

`func (o *Subscription) HasIsAnnouncementOnly() bool`

HasIsAnnouncementOnly returns a boolean if a field has been set.

### GetIsWebPublic

`func (o *Subscription) GetIsWebPublic() bool`

GetIsWebPublic returns the IsWebPublic field if non-nil, zero value otherwise.

### GetIsWebPublicOk

`func (o *Subscription) GetIsWebPublicOk() (*bool, bool)`

GetIsWebPublicOk returns a tuple with the IsWebPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebPublic

`func (o *Subscription) SetIsWebPublic(v bool)`

SetIsWebPublic sets IsWebPublic field to given value.

### HasIsWebPublic

`func (o *Subscription) HasIsWebPublic() bool`

HasIsWebPublic returns a boolean if a field has been set.

### GetColor

`func (o *Subscription) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Subscription) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Subscription) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Subscription) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetStreamPostPolicy

`func (o *Subscription) GetStreamPostPolicy() int32`

GetStreamPostPolicy returns the StreamPostPolicy field if non-nil, zero value otherwise.

### GetStreamPostPolicyOk

`func (o *Subscription) GetStreamPostPolicyOk() (*int32, bool)`

GetStreamPostPolicyOk returns a tuple with the StreamPostPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamPostPolicy

`func (o *Subscription) SetStreamPostPolicy(v int32)`

SetStreamPostPolicy sets StreamPostPolicy field to given value.

### HasStreamPostPolicy

`func (o *Subscription) HasStreamPostPolicy() bool`

HasStreamPostPolicy returns a boolean if a field has been set.

### GetMessageRetentionDays

`func (o *Subscription) GetMessageRetentionDays() int32`

GetMessageRetentionDays returns the MessageRetentionDays field if non-nil, zero value otherwise.

### GetMessageRetentionDaysOk

`func (o *Subscription) GetMessageRetentionDaysOk() (*int32, bool)`

GetMessageRetentionDaysOk returns a tuple with the MessageRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRetentionDays

`func (o *Subscription) SetMessageRetentionDays(v int32)`

SetMessageRetentionDays sets MessageRetentionDays field to given value.

### HasMessageRetentionDays

`func (o *Subscription) HasMessageRetentionDays() bool`

HasMessageRetentionDays returns a boolean if a field has been set.

### SetMessageRetentionDaysNil

`func (o *Subscription) SetMessageRetentionDaysNil(b bool)`

 SetMessageRetentionDaysNil sets the value for MessageRetentionDays to be an explicit nil

### UnsetMessageRetentionDays
`func (o *Subscription) UnsetMessageRetentionDays()`

UnsetMessageRetentionDays ensures that no value is present for MessageRetentionDays, not even an explicit nil
### GetHistoryPublicToSubscribers

`func (o *Subscription) GetHistoryPublicToSubscribers() bool`

GetHistoryPublicToSubscribers returns the HistoryPublicToSubscribers field if non-nil, zero value otherwise.

### GetHistoryPublicToSubscribersOk

`func (o *Subscription) GetHistoryPublicToSubscribersOk() (*bool, bool)`

GetHistoryPublicToSubscribersOk returns a tuple with the HistoryPublicToSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPublicToSubscribers

`func (o *Subscription) SetHistoryPublicToSubscribers(v bool)`

SetHistoryPublicToSubscribers sets HistoryPublicToSubscribers field to given value.

### HasHistoryPublicToSubscribers

`func (o *Subscription) HasHistoryPublicToSubscribers() bool`

HasHistoryPublicToSubscribers returns a boolean if a field has been set.

### GetFirstMessageId

`func (o *Subscription) GetFirstMessageId() int32`

GetFirstMessageId returns the FirstMessageId field if non-nil, zero value otherwise.

### GetFirstMessageIdOk

`func (o *Subscription) GetFirstMessageIdOk() (*int32, bool)`

GetFirstMessageIdOk returns a tuple with the FirstMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessageId

`func (o *Subscription) SetFirstMessageId(v int32)`

SetFirstMessageId sets FirstMessageId field to given value.

### HasFirstMessageId

`func (o *Subscription) HasFirstMessageId() bool`

HasFirstMessageId returns a boolean if a field has been set.

### SetFirstMessageIdNil

`func (o *Subscription) SetFirstMessageIdNil(b bool)`

 SetFirstMessageIdNil sets the value for FirstMessageId to be an explicit nil

### UnsetFirstMessageId
`func (o *Subscription) UnsetFirstMessageId()`

UnsetFirstMessageId ensures that no value is present for FirstMessageId, not even an explicit nil
### GetFolderId

`func (o *Subscription) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Subscription) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Subscription) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Subscription) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *Subscription) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *Subscription) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetTopicsPolicy

`func (o *Subscription) GetTopicsPolicy() TopicsPolicy`

GetTopicsPolicy returns the TopicsPolicy field if non-nil, zero value otherwise.

### GetTopicsPolicyOk

`func (o *Subscription) GetTopicsPolicyOk() (*TopicsPolicy, bool)`

GetTopicsPolicyOk returns a tuple with the TopicsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsPolicy

`func (o *Subscription) SetTopicsPolicy(v TopicsPolicy)`

SetTopicsPolicy sets TopicsPolicy field to given value.

### HasTopicsPolicy

`func (o *Subscription) HasTopicsPolicy() bool`

HasTopicsPolicy returns a boolean if a field has been set.

### GetIsRecentlyActive

`func (o *Subscription) GetIsRecentlyActive() bool`

GetIsRecentlyActive returns the IsRecentlyActive field if non-nil, zero value otherwise.

### GetIsRecentlyActiveOk

`func (o *Subscription) GetIsRecentlyActiveOk() (*bool, bool)`

GetIsRecentlyActiveOk returns a tuple with the IsRecentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecentlyActive

`func (o *Subscription) SetIsRecentlyActive(v bool)`

SetIsRecentlyActive sets IsRecentlyActive field to given value.

### HasIsRecentlyActive

`func (o *Subscription) HasIsRecentlyActive() bool`

HasIsRecentlyActive returns a boolean if a field has been set.

### GetStreamWeeklyTraffic

`func (o *Subscription) GetStreamWeeklyTraffic() int32`

GetStreamWeeklyTraffic returns the StreamWeeklyTraffic field if non-nil, zero value otherwise.

### GetStreamWeeklyTrafficOk

`func (o *Subscription) GetStreamWeeklyTrafficOk() (*int32, bool)`

GetStreamWeeklyTrafficOk returns a tuple with the StreamWeeklyTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamWeeklyTraffic

`func (o *Subscription) SetStreamWeeklyTraffic(v int32)`

SetStreamWeeklyTraffic sets StreamWeeklyTraffic field to given value.

### HasStreamWeeklyTraffic

`func (o *Subscription) HasStreamWeeklyTraffic() bool`

HasStreamWeeklyTraffic returns a boolean if a field has been set.

### SetStreamWeeklyTrafficNil

`func (o *Subscription) SetStreamWeeklyTrafficNil(b bool)`

 SetStreamWeeklyTrafficNil sets the value for StreamWeeklyTraffic to be an explicit nil

### UnsetStreamWeeklyTraffic
`func (o *Subscription) UnsetStreamWeeklyTraffic()`

UnsetStreamWeeklyTraffic ensures that no value is present for StreamWeeklyTraffic, not even an explicit nil
### GetCanAddSubscribersGroup

`func (o *Subscription) GetCanAddSubscribersGroup() ChannelCanAddSubscribersGroup`

GetCanAddSubscribersGroup returns the CanAddSubscribersGroup field if non-nil, zero value otherwise.

### GetCanAddSubscribersGroupOk

`func (o *Subscription) GetCanAddSubscribersGroupOk() (*ChannelCanAddSubscribersGroup, bool)`

GetCanAddSubscribersGroupOk returns a tuple with the CanAddSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddSubscribersGroup

`func (o *Subscription) SetCanAddSubscribersGroup(v ChannelCanAddSubscribersGroup)`

SetCanAddSubscribersGroup sets CanAddSubscribersGroup field to given value.

### HasCanAddSubscribersGroup

`func (o *Subscription) HasCanAddSubscribersGroup() bool`

HasCanAddSubscribersGroup returns a boolean if a field has been set.

### GetCanRemoveSubscribersGroup

`func (o *Subscription) GetCanRemoveSubscribersGroup() CanRemoveSubscribersGroup`

GetCanRemoveSubscribersGroup returns the CanRemoveSubscribersGroup field if non-nil, zero value otherwise.

### GetCanRemoveSubscribersGroupOk

`func (o *Subscription) GetCanRemoveSubscribersGroupOk() (*CanRemoveSubscribersGroup, bool)`

GetCanRemoveSubscribersGroupOk returns a tuple with the CanRemoveSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemoveSubscribersGroup

`func (o *Subscription) SetCanRemoveSubscribersGroup(v CanRemoveSubscribersGroup)`

SetCanRemoveSubscribersGroup sets CanRemoveSubscribersGroup field to given value.

### HasCanRemoveSubscribersGroup

`func (o *Subscription) HasCanRemoveSubscribersGroup() bool`

HasCanRemoveSubscribersGroup returns a boolean if a field has been set.

### GetCanAdministerChannelGroup

`func (o *Subscription) GetCanAdministerChannelGroup() CanAdministerChannelGroup`

GetCanAdministerChannelGroup returns the CanAdministerChannelGroup field if non-nil, zero value otherwise.

### GetCanAdministerChannelGroupOk

`func (o *Subscription) GetCanAdministerChannelGroupOk() (*CanAdministerChannelGroup, bool)`

GetCanAdministerChannelGroupOk returns a tuple with the CanAdministerChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdministerChannelGroup

`func (o *Subscription) SetCanAdministerChannelGroup(v CanAdministerChannelGroup)`

SetCanAdministerChannelGroup sets CanAdministerChannelGroup field to given value.

### HasCanAdministerChannelGroup

`func (o *Subscription) HasCanAdministerChannelGroup() bool`

HasCanAdministerChannelGroup returns a boolean if a field has been set.

### GetCanDeleteAnyMessageGroup

`func (o *Subscription) GetCanDeleteAnyMessageGroup() CanDeleteAnyMessageGroup`

GetCanDeleteAnyMessageGroup returns the CanDeleteAnyMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteAnyMessageGroupOk

`func (o *Subscription) GetCanDeleteAnyMessageGroupOk() (*CanDeleteAnyMessageGroup, bool)`

GetCanDeleteAnyMessageGroupOk returns a tuple with the CanDeleteAnyMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteAnyMessageGroup

`func (o *Subscription) SetCanDeleteAnyMessageGroup(v CanDeleteAnyMessageGroup)`

SetCanDeleteAnyMessageGroup sets CanDeleteAnyMessageGroup field to given value.

### HasCanDeleteAnyMessageGroup

`func (o *Subscription) HasCanDeleteAnyMessageGroup() bool`

HasCanDeleteAnyMessageGroup returns a boolean if a field has been set.

### GetCanDeleteOwnMessageGroup

`func (o *Subscription) GetCanDeleteOwnMessageGroup() CanDeleteOwnMessageGroup`

GetCanDeleteOwnMessageGroup returns the CanDeleteOwnMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteOwnMessageGroupOk

`func (o *Subscription) GetCanDeleteOwnMessageGroupOk() (*CanDeleteOwnMessageGroup, bool)`

GetCanDeleteOwnMessageGroupOk returns a tuple with the CanDeleteOwnMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteOwnMessageGroup

`func (o *Subscription) SetCanDeleteOwnMessageGroup(v CanDeleteOwnMessageGroup)`

SetCanDeleteOwnMessageGroup sets CanDeleteOwnMessageGroup field to given value.

### HasCanDeleteOwnMessageGroup

`func (o *Subscription) HasCanDeleteOwnMessageGroup() bool`

HasCanDeleteOwnMessageGroup returns a boolean if a field has been set.

### GetCanMoveMessagesOutOfChannelGroup

`func (o *Subscription) GetCanMoveMessagesOutOfChannelGroup() CanMoveMessagesOutOfChannelGroup`

GetCanMoveMessagesOutOfChannelGroup returns the CanMoveMessagesOutOfChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesOutOfChannelGroupOk

`func (o *Subscription) GetCanMoveMessagesOutOfChannelGroupOk() (*CanMoveMessagesOutOfChannelGroup, bool)`

GetCanMoveMessagesOutOfChannelGroupOk returns a tuple with the CanMoveMessagesOutOfChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesOutOfChannelGroup

`func (o *Subscription) SetCanMoveMessagesOutOfChannelGroup(v CanMoveMessagesOutOfChannelGroup)`

SetCanMoveMessagesOutOfChannelGroup sets CanMoveMessagesOutOfChannelGroup field to given value.

### HasCanMoveMessagesOutOfChannelGroup

`func (o *Subscription) HasCanMoveMessagesOutOfChannelGroup() bool`

HasCanMoveMessagesOutOfChannelGroup returns a boolean if a field has been set.

### GetCanMoveMessagesWithinChannelGroup

`func (o *Subscription) GetCanMoveMessagesWithinChannelGroup() CanMoveMessagesWithinChannelGroup`

GetCanMoveMessagesWithinChannelGroup returns the CanMoveMessagesWithinChannelGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesWithinChannelGroupOk

`func (o *Subscription) GetCanMoveMessagesWithinChannelGroupOk() (*CanMoveMessagesWithinChannelGroup, bool)`

GetCanMoveMessagesWithinChannelGroupOk returns a tuple with the CanMoveMessagesWithinChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesWithinChannelGroup

`func (o *Subscription) SetCanMoveMessagesWithinChannelGroup(v CanMoveMessagesWithinChannelGroup)`

SetCanMoveMessagesWithinChannelGroup sets CanMoveMessagesWithinChannelGroup field to given value.

### HasCanMoveMessagesWithinChannelGroup

`func (o *Subscription) HasCanMoveMessagesWithinChannelGroup() bool`

HasCanMoveMessagesWithinChannelGroup returns a boolean if a field has been set.

### GetCanSendMessageGroup

`func (o *Subscription) GetCanSendMessageGroup() CanSendMessageGroup`

GetCanSendMessageGroup returns the CanSendMessageGroup field if non-nil, zero value otherwise.

### GetCanSendMessageGroupOk

`func (o *Subscription) GetCanSendMessageGroupOk() (*CanSendMessageGroup, bool)`

GetCanSendMessageGroupOk returns a tuple with the CanSendMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSendMessageGroup

`func (o *Subscription) SetCanSendMessageGroup(v CanSendMessageGroup)`

SetCanSendMessageGroup sets CanSendMessageGroup field to given value.

### HasCanSendMessageGroup

`func (o *Subscription) HasCanSendMessageGroup() bool`

HasCanSendMessageGroup returns a boolean if a field has been set.

### GetCanSubscribeGroup

`func (o *Subscription) GetCanSubscribeGroup() CanSubscribeGroup`

GetCanSubscribeGroup returns the CanSubscribeGroup field if non-nil, zero value otherwise.

### GetCanSubscribeGroupOk

`func (o *Subscription) GetCanSubscribeGroupOk() (*CanSubscribeGroup, bool)`

GetCanSubscribeGroupOk returns a tuple with the CanSubscribeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSubscribeGroup

`func (o *Subscription) SetCanSubscribeGroup(v CanSubscribeGroup)`

SetCanSubscribeGroup sets CanSubscribeGroup field to given value.

### HasCanSubscribeGroup

`func (o *Subscription) HasCanSubscribeGroup() bool`

HasCanSubscribeGroup returns a boolean if a field has been set.

### GetCanResolveTopicsGroup

`func (o *Subscription) GetCanResolveTopicsGroup() CanResolveTopicsGroup`

GetCanResolveTopicsGroup returns the CanResolveTopicsGroup field if non-nil, zero value otherwise.

### GetCanResolveTopicsGroupOk

`func (o *Subscription) GetCanResolveTopicsGroupOk() (*CanResolveTopicsGroup, bool)`

GetCanResolveTopicsGroupOk returns a tuple with the CanResolveTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResolveTopicsGroup

`func (o *Subscription) SetCanResolveTopicsGroup(v CanResolveTopicsGroup)`

SetCanResolveTopicsGroup sets CanResolveTopicsGroup field to given value.

### HasCanResolveTopicsGroup

`func (o *Subscription) HasCanResolveTopicsGroup() bool`

HasCanResolveTopicsGroup returns a boolean if a field has been set.

### GetIsArchived

`func (o *Subscription) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Subscription) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Subscription) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *Subscription) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetSubscriberCount

`func (o *Subscription) GetSubscriberCount() float32`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *Subscription) GetSubscriberCountOk() (*float32, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *Subscription) SetSubscriberCount(v float32)`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *Subscription) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


