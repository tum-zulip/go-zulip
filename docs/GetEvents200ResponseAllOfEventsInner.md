# GetEvents200ResponseAllOfEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | 
**Type** | **string** |  | 
**AlertWords** | Pointer to **[]string** | An array of strings, where each string is an alert word (or phrase) configured by the user.  | [optional] 
**User** | Pointer to **interface{}** |  | [optional] 
**SettingName** | Pointer to **string** | Name of the changed display setting.  | [optional] 
**Setting** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf1Setting**](GetEvents200ResponseAllOfEventsInnerOneOf1Setting.md) |  | [optional] 
**LanguageName** | Pointer to **string** | Present only if the setting to be changed is &#x60;default_language&#x60;. Contains the name of the new default language in English.  | [optional] 
**NotificationName** | Pointer to **string** | Name of the changed notification setting.  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **string** | The name of the property that was changed.  | [optional] 
**Value** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf68Value**](GetEvents200ResponseAllOfEventsInnerOneOf68Value.md) |  | [optional] 
**Person** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf14Person**](GetEvents200ResponseAllOfEventsInnerOneOf14Person.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]GetEvents200ResponseAllOfEventsInnerOneOf6SubscriptionsInner**](GetEvents200ResponseAllOfEventsInnerOneOf6SubscriptionsInner.md) | A list of dictionaries, where each dictionary contains information about one of the newly unsubscribed channels.  | [optional] 
**StreamId** | Pointer to **int32** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  The unique ID of the channel to which message is being typed.  **Changes**: New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  | [optional] 
**StreamIds** | Pointer to **[]int32** | Array containing the IDs of the channels that were deleted.  **Changes**: New in Zulip 10.0 (feature level 343). Previously, these IDs were available only via the legacy &#x60;streams&#x60; array.  | [optional] 
**UserIds** | Pointer to **[]int32** | Array containing the IDs of the users who have been removed from the user group.  | [optional] 
**Message** | Pointer to [**MessagesEvent**](MessagesEvent.md) |  | [optional] 
**Flags** | **[]string** | The user&#39;s personal [message flags][message-flags] for the message with ID &#x60;message_id&#x60; following the edit.  A client application should compare these to the original flags to identify cases where a mention or alert word was added by the edit.  **Changes**: In Zulip 8.0 (feature level 224), the &#x60;wildcard_mentioned&#x60; flag was deprecated in favor of the &#x60;stream_wildcard_mentioned&#x60; and &#x60;topic_wildcard_mentioned&#x60; flags. The &#x60;wildcard_mentioned&#x60; flag exists for backwards compatibility with older clients and equals &#x60;stream_wildcard_mentioned || topic_wildcard_mentioned&#x60;. Clients supporting older server versions should treat this field as a previous name for the &#x60;stream_wildcard_mentioned&#x60; flag as topic wildcard mentions were not available prior to this feature level.  [message-flags]: /api/update-message-flags#available-flags  | 
**Presences** | Pointer to [**map[string]ModernPresenceFormat**](ModernPresenceFormat.md) | Only present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  A dictionary mapping user IDs to the presence data (modern format) for the modified user(s). Clients should support updating multiple users in a single event.  **Changes**: New in Zulip 11.0 (feature level 419).  | [optional] 
**UserId** | **int32** | The ID of the user whose setting was changed.  | 
**Email** | Pointer to **string** | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the &#x60;user_id&#x60;.  | [optional] 
**ServerTimestamp** | Pointer to **float32** | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  The timestamp of when the Zulip server received the user&#39;s presence as a UNIX timestamp.  | [optional] 
**Presence** | Pointer to [**map[string]GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue**](GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue.md) | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  Object containing the presence data (legacy format) of of the modified user.  | [optional] 
**Streams** | Pointer to [**[]GetEvents200ResponseAllOfEventsInnerOneOf17StreamsInner**](GetEvents200ResponseAllOfEventsInnerOneOf17StreamsInner.md) | Array of objects, each containing ID of the channel that was deleted.  **Changes**: **Deprecated** in Zulip 10.0 (feature level 343) and will be removed in a future release. Previously, these objects additionally contained all the standard fields for a channel object.  | [optional] 
**Name** | Pointer to **string** | The name of the channel whose details have changed.  | [optional] 
**RenderedDescription** | Pointer to **string** | Note: Only present if the changed property was &#x60;description&#x60;.  The short description of the channel rendered as HTML, intended to be used when displaying the channel description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**HistoryPublicToSubscribers** | Pointer to **bool** | Note: Only present if the changed property was &#x60;invite_only&#x60;.  Whether the history of the channel is public to its subscribers.  Currently always true for public channels (i.e. &#x60;\&quot;invite_only\&quot;: false&#x60; implies &#x60;\&quot;history_public_to_subscribers\&quot;: true&#x60;), but clients should not make that assumption, as we may change that behavior in the future.  | [optional] 
**IsWebPublic** | Pointer to **bool** | Note: Only present if the changed property was &#x60;invite_only&#x60;.  Whether the channel&#39;s history is now readable by web-public spectators.  **Changes**: New in Zulip 5.0 (feature level 71).  | [optional] 
**EmojiName** | Pointer to **string** | The [emoji name](/api/update-status#parameter-emoji_name) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**EmojiCode** | Pointer to **string** | The [emoji code](/api/update-status#parameter-emoji_code) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**ReactionType** | Pointer to **string** | The [emoji type](/api/update-status#parameter-reaction_type) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**MessageId** | **int32** | Indicates the message id of the message that is being edited.  | 
**Attachment** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf23Attachment**](GetEvents200ResponseAllOfEventsInnerOneOf23Attachment.md) |  | [optional] 
**UploadSpaceUsed** | Pointer to **int32** | The total size of all files uploaded by in the organization, in bytes.  | [optional] 
**PushAccountId** | Pointer to **string** | The push account ID for this client registration.  See [&#x60;POST /mobile_push/register&#x60;](/api/register-push-device) for details on push account IDs.  | [optional] 
**Status** | Pointer to **string** | The updated registration status. Will be &#x60;\&quot;active\&quot;&#x60;, &#x60;\&quot;failed\&quot;&#x60;, or &#x60;\&quot;pending\&quot;&#x60;.  | [optional] 
**ErrorCode** | Pointer to **NullableString** | If the status is &#x60;\&quot;failed\&quot;&#x60;, a [Zulip API error code](/api/rest-error-handling) indicating the type of failure that occurred.  The following error codes have recommended client behavior:  - &#x60;\&quot;INVALID_BOUNCER_PUBLIC_KEY\&quot;&#x60; - Inform the user to update app. - &#x60;\&quot;REQUEST_EXPIRED&#x60; - Retry with a fresh payload.   If the status is \&quot;failed\&quot;, an error code explaining the failure.  | [optional] 
**MsgType** | Pointer to **string** | The type of the message.  | [optional] 
**Content** | Pointer to **string** | Only present if this event changed the message content or updated the message content for an [inline URL preview][inline-url-previews].  The new content of the message with ID &#x60;message_id&#x60;, in the original [Zulip-flavored Markdown](/help/format-your-message-using-markdown) format.  | [optional] 
**SenderId** | Pointer to **int32** | The ID of the user who sent the message.  | [optional] 
**SubmessageId** | Pointer to **int32** | The ID of the submessage.  | [optional] 
**Away** | Pointer to **bool** | Whether the user has marked themself \&quot;away\&quot; with this status.  **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, &#x60;away&#x60; is a legacy way to access the user&#39;s &#x60;presence_enabled&#x60; setting, with &#x60;away &#x3D; !presence_enabled&#x60;. To be removed in a future release.  | [optional] 
**StatusText** | Pointer to **string** | The text content of the status message.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting or writing a message.  | [optional] 
**Fields** | Pointer to [**[]CustomProfileField**](CustomProfileField.md) | An array of dictionaries where each dictionary contains details of a single new custom profile field for the Zulip organization.  | [optional] 
**DefaultStreamGroups** | Pointer to [**[]DefaultChannelGroup**](DefaultChannelGroup.md) | An array of dictionaries where each dictionary contains details about a single default channel group.  | [optional] 
**DefaultStreams** | Pointer to **[]int32** | An array of IDs of all the [default channels](/help/set-default-streams-for-new-users) in the organization.  **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single default stream for the Zulip organization.  | [optional] 
**MessageIds** | **[]int32** | A sorted list of IDs of messages to which any channel or topic changes encoded in this event should be applied.  This list always includes &#x60;message_id&#x60;, even when there are no channel or topic changes to apply.  These messages are guaranteed to have all been previously sent to channel &#x60;stream_id&#x60; with topic &#x60;orig_subject&#x60;, and have been moved to &#x60;new_stream_id&#x60; with topic &#x60;subject&#x60; (if those fields are present in the event).  Clients processing these events should update all cached message history associated with the moved messages (including adjusting &#x60;unread_msgs&#x60; data structures, where the client may not have the message itself in its history) to reflect the new channel and topic.  Content changes should be applied only to the single message indicated by &#x60;message_id&#x60;.  **Changes**: Before Zulip 11.0 (feature level 393), this list was not guaranteed to be sorted.  | 
**MessageType** | Pointer to **string** | Type of message being composed. Must be &#x60;\&quot;stream\&quot;&#x60; or &#x60;\&quot;direct\&quot;&#x60;.  **Changes**: In Zulip 8.0 (feature level 215), replaced the value &#x60;\&quot;private\&quot;&#x60; with &#x60;\&quot;direct\&quot;&#x60;.  New in Zulip 4.0 (feature level 58). Previously all typing notifications were implicitly direct messages.  | [optional] 
**Topic** | Pointer to **string** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  Topic within the channel where the message is being typed.  **Changes**: New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  | [optional] 
**MutedTopics** | Pointer to [**[][]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner**]([]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner.md) | Array of tuples, where each tuple describes a muted topic. The first element of the tuple is the channel name in which the topic has to be muted, the second element is the topic name to be muted and the third element is an integer UNIX timestamp representing when the topic was muted.  **Changes**: Deprecated in Zulip 6.0 (feature level 134). Starting with this version, clients that explicitly requested the replacement &#x60;user_topic&#x60; event type when registering their event queue will not receive this legacy event type.  Before Zulip 3.0 (feature level 1), the &#x60;muted_topics&#x60; array objects were 2-item tuples and did not include the timestamp information for when the topic was muted.  | [optional] 
**TopicName** | Pointer to **string** | The name of the topic.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**LastUpdated** | Pointer to **int32** | An integer UNIX timestamp representing when the user-topic relationship was last changed.  | [optional] 
**VisibilityPolicy** | Pointer to **int32** | An integer indicating the user&#39;s visibility preferences for the topic, such as whether the topic is muted.  - 0 &#x3D; None. Used to indicate that the user no   longer has a special visibility policy for this topic. - 1 &#x3D; Muted. Used to record [muted topics](/help/mute-a-topic). - 2 &#x3D; Unmuted. Used to record unmuted topics. - 3 &#x3D; Followed. Used to record [followed topics](/help/follow-a-topic).  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.  In Zulip 7.0 (feature level 170), added unmuted as a visibility policy option.  | [optional] 
**MutedUsers** | Pointer to [**[]GetEvents200ResponseAllOfEventsInnerOneOf33MutedUsersInner**](GetEvents200ResponseAllOfEventsInnerOneOf33MutedUsersInner.md) | A list of dictionaries where each dictionary describes a muted user.  | [optional] 
**OnboardingSteps** | Pointer to [**[]OnboardingStep**](OnboardingStep.md) | An array of dictionaries where each dictionary contains details about a single onboarding step.  **Changes**: Before Zulip 8.0 (feature level 233), this array was named &#x60;hotspots&#x60;. Prior to this feature level, one-time notice onboarding steps were not supported, and the &#x60;type&#x60; field in these objects did not exist as all onboarding steps were implicitly hotspots.  | [optional] 
**RenderingOnly** | **bool** | Whether the event only updates the rendered content of the message.  This field should be used by clients to determine if the event only provides a rendering update to the message content, such as for an [inline URL preview][inline-url-previews]. When &#x60;true&#x60;, the event does not reflect a user-generated edit and does not modify the message history.  **Changes**: New in Zulip 5.0 (feature level 114). Clients can correctly identify these rendering update event with earlier Zulip versions by checking whether the &#x60;user_id&#x60; field was omitted.  | 
**EditTimestamp** | **int32** | The time when this message edit operation was processed by the server.  **Changes**: As of Zulip 5.0 (feature level 114), this field is present for all &#x60;update_message&#x60; events. Previously, this field was omitted for [inline URL preview][inline-url-previews] updates.  | 
**StreamName** | Pointer to **string** | Only present if the message was edited and originally sent to a channel.  The name of the channel that the message was sent to. Clients are recommended to use the &#x60;stream_id&#x60; field instead.  | [optional] 
**NewStreamId** | Pointer to **int32** | Only present if message(s) were moved to a different channel.  The post-edit channel for all of the messages with IDs in &#x60;message_ids&#x60;.  | [optional] 
**PropagateMode** | Pointer to **string** | Only present if this event moved messages to a different topic and/or channel.  The choice the editing user made about which messages should be affected by a channel/topic edit:  - &#x60;\&quot;change_one\&quot;&#x60;: Just change the one indicated in &#x60;message_id&#x60;. - &#x60;\&quot;change_later\&quot;&#x60;: Change messages in the same topic that had   been sent after this one. - &#x60;\&quot;change_all\&quot;&#x60;: Change all messages in that topic.  This parameter should be used to decide whether to change navigation and compose box state in response to the edit. For example, if the user was previously in topic narrow, and the topic was edited with &#x60;\&quot;change_later\&quot;&#x60; or &#x60;\&quot;change_all\&quot;&#x60;, the Zulip web app will automatically navigate to the new topic narrow. Similarly, a message being composed to the old topic should have its recipient changed to the new topic.  This navigation makes it much more convenient to move content between topics without disruption or messages continuing to be sent to the pre-edit topic by accident.  | [optional] 
**OrigSubject** | Pointer to **string** | Only present if this event moved messages to a different topic and/or channel.  The pre-edit topic for all of the messages with IDs in &#x60;message_ids&#x60;.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual pre-edit topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**Subject** | Pointer to **string** | Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  The post-edit topic for all of the messages with IDs in &#x60;message_ids&#x60;.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual post-edit topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**TopicLinks** | Pointer to [**[]GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner**](GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner.md) | Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  Data on any links to be included in the &#x60;topic&#x60; line (these are generated by [custom linkification filter](/help/add-a-custom-linkifier) that match content in the message&#39;s topic.), corresponding to the post-edit topic.  **Changes**: This field contained a list of urls before Zulip 4.0 (feature level 46).  New in Zulip 3.0 (feature level 1). Previously, this field was called &#x60;subject_links&#x60;; clients are recommended to rename &#x60;subject_links&#x60; to &#x60;topic_links&#x60; if present for compatibility with older Zulip servers.  | [optional] 
**OrigContent** | Pointer to **string** | Only present if this event changed the message content.  The original content of the message with ID &#x60;message_id&#x60; immediately prior to this edit, in the original [Zulip-flavored Markdown](/help/format-your-message-using-markdown) format.  | [optional] 
**OrigRenderedContent** | Pointer to **string** | Only present if this event changed the message content.  The original content of the message with ID &#x60;message_id&#x60; immediately prior to this edit, rendered as HTML.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**RenderedContent** | Pointer to **string** | Only present if this event changed the message content or updated the message content for an [inline URL preview][inline-url-previews].  The new content of the message with ID &#x60;message_id&#x60;, rendered in HTML.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**IsMeMessage** | Pointer to **bool** | Only present if this event changed the message content.  Whether the message with ID &#x60;message_id&#x60; is now a [/me status message][status-messages].  [status-messages]: /help/format-your-message-using-markdown#status-messages  | [optional] 
**Sender** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf38Sender**](GetEvents200ResponseAllOfEventsInnerOneOf38Sender.md) |  | [optional] 
**Recipients** | Pointer to [**[]GetEvents200ResponseAllOfEventsInnerOneOf38RecipientsInner**](GetEvents200ResponseAllOfEventsInnerOneOf38RecipientsInner.md) | Only present if &#x60;message_type&#x60; is &#x60;\&quot;direct\&quot;&#x60;.  Array of dictionaries describing the set of users who would be recipients of the message that was previously being typed. Each dictionary contains details about one of the recipients. The sending user is guaranteed to appear among the recipients.  | [optional] 
**Recipient** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf39Recipient**](GetEvents200ResponseAllOfEventsInnerOneOf39Recipient.md) |  | [optional] 
**Operation** | Pointer to **string** | Old name for the &#x60;op&#x60; field in this event type.  **Deprecated** in Zulip 4.0 (feature level 32), and replaced by the &#x60;op&#x60; field.  | [optional] 
**Flag** | Pointer to **string** | The [flag][message-flags] to be removed.  | [optional] 
**Messages** | Pointer to **[]int32** | Array containing the IDs of the messages from which the flag was removed.  | [optional] 
**All** | Pointer to **bool** | Will be &#x60;false&#x60; for all specified flags.  **Deprecated** and will be removed in a future release.  | [optional] 
**MessageDetails** | Pointer to [**map[string]GetEvents200ResponseAllOfEventsInnerOneOf42MessageDetailsValue**](GetEvents200ResponseAllOfEventsInnerOneOf42MessageDetailsValue.md) | Only present if the specified &#x60;flag&#x60; is &#x60;\&quot;read\&quot;&#x60;.  A set of data structures describing the messages that are being marked as unread with additional details to allow clients to update the &#x60;unread_msgs&#x60; data structure for these messages (which may not be otherwise known to the client).  **Changes**: New in Zulip 5.0 (feature level 121). Previously, marking already read messages as unread was not supported by the Zulip API.  | [optional] 
**Group** | Pointer to [**UserGroup**](UserGroup.md) |  | [optional] 
**GroupId** | Pointer to **int32** | The ID of the group which has been deleted.  | [optional] 
**Data** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf84Data**](GetEvents200ResponseAllOfEventsInnerOneOf84Data.md) |  | [optional] 
**DirectSubgroupIds** | Pointer to **[]int32** | Array containing the IDs of the subgroups that have been removed from the user group.  **Changes**: New in Zulip 6.0 (feature level 131). Previously, this was called &#x60;subgroup_ids&#x60;, but clients can ignore older events as this feature level predates subgroups being fully implemented.  | [optional] 
**RealmLinkifiers** | Pointer to [**[]GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner**](GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner.md) | An ordered array of dictionaries where each dictionary contains details about a single linkifier.  Clients should always process linkifiers in the order given; this is important if the realm has linkifiers with overlapping patterns. The order can be modified using [&#x60;PATCH /realm/linkifiers&#x60;](/api/reorder-linkifiers).  | [optional] 
**RealmFilters** | Pointer to [**[][]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner**]([]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner.md) | An array of tuples, where each tuple described a linkifier. The first element of the tuple was a string regex pattern which represented the pattern to be linkified on matching, for example &#x60;\&quot;#(?P&lt;id&gt;[123])\&quot;&#x60;. The second element was the URL format string that the pattern should be linkified with. A URL format string for the above example would be &#x60;\&quot;https://realm.com/my_realm_filter/%(id)s\&quot;&#x60;. And the third element was the ID of the realm filter.  | [optional] 
**RealmPlaygrounds** | Pointer to [**[]RealmPlayground**](RealmPlayground.md) | An array of dictionaries where each dictionary contains data about a single playground entry.  | [optional] 
**RealmEmoji** | Pointer to [**map[string]RealmEmoji**](RealmEmoji.md) | An object in which each key describes a realm emoji.  | [optional] 
**RealmDomain** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain**](GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain.md) |  | [optional] 
**Domain** | Pointer to **string** | The domain to be removed.  | [optional] 
**Exports** | Pointer to [**[]RealmExport**](RealmExport.md) | An array of dictionaries where each dictionary contains details about a data export of the organization.  **Changes**: Prior to Zulip 10.0 (feature level 304), &#x60;export_type&#x60; parameter was not present as only public data export was supported via API.  | [optional] 
**Consented** | Pointer to **bool** | Whether the user has consented for their private data export.  | [optional] 
**Bot** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf62Bot**](GetEvents200ResponseAllOfEventsInnerOneOf62Bot.md) |  | [optional] 
**RealmId** | Pointer to **int32** | The ID of the deactivated realm.  | [optional] 
**ZulipVersion** | Pointer to **string** | The Zulip version number, in the format where this appears in the [server_settings](/api/get-server-settings) and [register](/api/register-queue) responses.  **Changes**: New in Zulip 4.0 (feature level 59).  | [optional] 
**ZulipMergeBase** | Pointer to **string** | The Zulip merge base number, in the format where this appears in the [server_settings](/api/get-server-settings) and [register](/api/register-queue) responses.  **Changes**: New in Zulip 5.0 (feature level 88).  | [optional] 
**ZulipFeatureLevel** | Pointer to **int32** | The [Zulip feature level](/api/changelog) of the server after the restart.  Clients should use this to update their tracking of the server&#39;s capabilities, and may choose to refetch their state and create a new event queue when the API feature level has changed in a way that the client finds significant. Clients choosing to do so must implement a random delay strategy to spread such restarts over 5 or more minutes to avoid creating a synchronized thundering herd effect.  **Changes**: New in Zulip 4.0 (feature level 59).  | [optional] 
**ServerGeneration** | Pointer to **int32** | The timestamp at which the server started.  | [optional] 
**Immediate** | Pointer to **bool** | Whether the client should fetch a new event queue immediately, rather than using a backoff strategy to avoid thundering herds. A Zulip development server uses this parameter to reload clients immediately.  | [optional] 
**Drafts** | Pointer to [**[]Draft**](Draft.md) | An array containing objects for the newly created drafts.  | [optional] 
**Draft** | Pointer to [**Draft**](Draft.md) |  | [optional] 
**DraftId** | Pointer to **int32** | The ID of the draft that was just deleted.  | [optional] 
**NavigationView** | Pointer to [**NavigationView**](NavigationView.md) |  | [optional] 
**Fragment** | Pointer to **string** | The unique URL hash of the navigation view that was just deleted.  | [optional] 
**SavedSnippet** | Pointer to [**SavedSnippet**](SavedSnippet.md) |  | [optional] 
**SavedSnippetId** | Pointer to **int32** | The ID of the saved snippet that was just deleted.  **Changes**: New in Zulip 10.0 (feature level 297).  | [optional] 
**Reminders** | Pointer to [**[]Reminder**](Reminder.md) | An array of objects containing details of the newly created reminders.  | [optional] 
**ReminderId** | Pointer to **int32** | The ID of the reminder that was deleted.  | [optional] 
**ScheduledMessages** | Pointer to [**[]ScheduledMessage**](ScheduledMessage.md) | An array of objects containing details of the newly created scheduled messages.  | [optional] 
**ScheduledMessage** | Pointer to [**ScheduledMessage**](ScheduledMessage.md) |  | [optional] 
**ScheduledMessageId** | Pointer to **int32** | The ID of the scheduled message that was deleted.  | [optional] 
**ChannelFolder** | Pointer to [**ChannelFolder**](ChannelFolder.md) |  | [optional] 
**ChannelFolderId** | Pointer to **float32** | ID of the updated channel folder.  | [optional] 
**Order** | Pointer to **[]int32** | A list of channel folder IDs representing the new order.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInner

`func NewGetEvents200ResponseAllOfEventsInner(id int32, type_ string, flags []string, userId int32, messageId int32, messageIds []int32, renderingOnly bool, editTimestamp int32, ) *GetEvents200ResponseAllOfEventsInner`

NewGetEvents200ResponseAllOfEventsInner instantiates a new GetEvents200ResponseAllOfEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerWithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerWithDefaults() *GetEvents200ResponseAllOfEventsInner`

NewGetEvents200ResponseAllOfEventsInnerWithDefaults instantiates a new GetEvents200ResponseAllOfEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEvents200ResponseAllOfEventsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvents200ResponseAllOfEventsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *GetEvents200ResponseAllOfEventsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEvents200ResponseAllOfEventsInner) SetType(v string)`

SetType sets Type field to given value.


### GetAlertWords

`func (o *GetEvents200ResponseAllOfEventsInner) GetAlertWords() []string`

GetAlertWords returns the AlertWords field if non-nil, zero value otherwise.

### GetAlertWordsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetAlertWordsOk() (*[]string, bool)`

GetAlertWordsOk returns a tuple with the AlertWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertWords

`func (o *GetEvents200ResponseAllOfEventsInner) SetAlertWords(v []string)`

SetAlertWords sets AlertWords field to given value.

### HasAlertWords

`func (o *GetEvents200ResponseAllOfEventsInner) HasAlertWords() bool`

HasAlertWords returns a boolean if a field has been set.

### GetUser

`func (o *GetEvents200ResponseAllOfEventsInner) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetEvents200ResponseAllOfEventsInner) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *GetEvents200ResponseAllOfEventsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GetEvents200ResponseAllOfEventsInner) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetEvents200ResponseAllOfEventsInner) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetSettingName

`func (o *GetEvents200ResponseAllOfEventsInner) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *GetEvents200ResponseAllOfEventsInner) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *GetEvents200ResponseAllOfEventsInner) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetSetting

`func (o *GetEvents200ResponseAllOfEventsInner) GetSetting() GetEvents200ResponseAllOfEventsInnerOneOf1Setting`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSettingOk() (*GetEvents200ResponseAllOfEventsInnerOneOf1Setting, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *GetEvents200ResponseAllOfEventsInner) SetSetting(v GetEvents200ResponseAllOfEventsInnerOneOf1Setting)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *GetEvents200ResponseAllOfEventsInner) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetLanguageName

`func (o *GetEvents200ResponseAllOfEventsInner) GetLanguageName() string`

GetLanguageName returns the LanguageName field if non-nil, zero value otherwise.

### GetLanguageNameOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetLanguageNameOk() (*string, bool)`

GetLanguageNameOk returns a tuple with the LanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageName

`func (o *GetEvents200ResponseAllOfEventsInner) SetLanguageName(v string)`

SetLanguageName sets LanguageName field to given value.

### HasLanguageName

`func (o *GetEvents200ResponseAllOfEventsInner) HasLanguageName() bool`

HasLanguageName returns a boolean if a field has been set.

### GetNotificationName

`func (o *GetEvents200ResponseAllOfEventsInner) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *GetEvents200ResponseAllOfEventsInner) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *GetEvents200ResponseAllOfEventsInner) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetOp

`func (o *GetEvents200ResponseAllOfEventsInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GetEvents200ResponseAllOfEventsInner) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *GetEvents200ResponseAllOfEventsInner) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetProperty

`func (o *GetEvents200ResponseAllOfEventsInner) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *GetEvents200ResponseAllOfEventsInner) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *GetEvents200ResponseAllOfEventsInner) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *GetEvents200ResponseAllOfEventsInner) GetValue() GetEvents200ResponseAllOfEventsInnerOneOf68Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetValueOk() (*GetEvents200ResponseAllOfEventsInnerOneOf68Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetEvents200ResponseAllOfEventsInner) SetValue(v GetEvents200ResponseAllOfEventsInnerOneOf68Value)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetEvents200ResponseAllOfEventsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPerson

`func (o *GetEvents200ResponseAllOfEventsInner) GetPerson() GetEvents200ResponseAllOfEventsInnerOneOf14Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetPersonOk() (*GetEvents200ResponseAllOfEventsInnerOneOf14Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *GetEvents200ResponseAllOfEventsInner) SetPerson(v GetEvents200ResponseAllOfEventsInnerOneOf14Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *GetEvents200ResponseAllOfEventsInner) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetSubscriptions

`func (o *GetEvents200ResponseAllOfEventsInner) GetSubscriptions() []GetEvents200ResponseAllOfEventsInnerOneOf6SubscriptionsInner`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSubscriptionsOk() (*[]GetEvents200ResponseAllOfEventsInnerOneOf6SubscriptionsInner, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *GetEvents200ResponseAllOfEventsInner) SetSubscriptions(v []GetEvents200ResponseAllOfEventsInnerOneOf6SubscriptionsInner)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *GetEvents200ResponseAllOfEventsInner) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetStreamId

`func (o *GetEvents200ResponseAllOfEventsInner) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *GetEvents200ResponseAllOfEventsInner) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *GetEvents200ResponseAllOfEventsInner) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetStreamIds

`func (o *GetEvents200ResponseAllOfEventsInner) GetStreamIds() []int32`

GetStreamIds returns the StreamIds field if non-nil, zero value otherwise.

### GetStreamIdsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetStreamIdsOk() (*[]int32, bool)`

GetStreamIdsOk returns a tuple with the StreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamIds

`func (o *GetEvents200ResponseAllOfEventsInner) SetStreamIds(v []int32)`

SetStreamIds sets StreamIds field to given value.

### HasStreamIds

`func (o *GetEvents200ResponseAllOfEventsInner) HasStreamIds() bool`

HasStreamIds returns a boolean if a field has been set.

### GetUserIds

`func (o *GetEvents200ResponseAllOfEventsInner) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *GetEvents200ResponseAllOfEventsInner) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *GetEvents200ResponseAllOfEventsInner) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetMessage

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessage() MessagesEvent`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageOk() (*MessagesEvent, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEvents200ResponseAllOfEventsInner) SetMessage(v MessagesEvent)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetEvents200ResponseAllOfEventsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFlags

`func (o *GetEvents200ResponseAllOfEventsInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetEvents200ResponseAllOfEventsInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.


### GetPresences

`func (o *GetEvents200ResponseAllOfEventsInner) GetPresences() map[string]ModernPresenceFormat`

GetPresences returns the Presences field if non-nil, zero value otherwise.

### GetPresencesOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetPresencesOk() (*map[string]ModernPresenceFormat, bool)`

GetPresencesOk returns a tuple with the Presences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresences

`func (o *GetEvents200ResponseAllOfEventsInner) SetPresences(v map[string]ModernPresenceFormat)`

SetPresences sets Presences field to given value.

### HasPresences

`func (o *GetEvents200ResponseAllOfEventsInner) HasPresences() bool`

HasPresences returns a boolean if a field has been set.

### GetUserId

`func (o *GetEvents200ResponseAllOfEventsInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetEvents200ResponseAllOfEventsInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *GetEvents200ResponseAllOfEventsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetEvents200ResponseAllOfEventsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetEvents200ResponseAllOfEventsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetServerTimestamp

`func (o *GetEvents200ResponseAllOfEventsInner) GetServerTimestamp() float32`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetServerTimestampOk() (*float32, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *GetEvents200ResponseAllOfEventsInner) SetServerTimestamp(v float32)`

SetServerTimestamp sets ServerTimestamp field to given value.

### HasServerTimestamp

`func (o *GetEvents200ResponseAllOfEventsInner) HasServerTimestamp() bool`

HasServerTimestamp returns a boolean if a field has been set.

### GetPresence

`func (o *GetEvents200ResponseAllOfEventsInner) GetPresence() map[string]GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetPresenceOk() (*map[string]GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *GetEvents200ResponseAllOfEventsInner) SetPresence(v map[string]GetEvents200ResponseAllOfEventsInnerOneOf15PresenceValue)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *GetEvents200ResponseAllOfEventsInner) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetStreams

`func (o *GetEvents200ResponseAllOfEventsInner) GetStreams() []GetEvents200ResponseAllOfEventsInnerOneOf17StreamsInner`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetStreamsOk() (*[]GetEvents200ResponseAllOfEventsInnerOneOf17StreamsInner, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *GetEvents200ResponseAllOfEventsInner) SetStreams(v []GetEvents200ResponseAllOfEventsInnerOneOf17StreamsInner)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *GetEvents200ResponseAllOfEventsInner) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetName

`func (o *GetEvents200ResponseAllOfEventsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEvents200ResponseAllOfEventsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetEvents200ResponseAllOfEventsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRenderedDescription

`func (o *GetEvents200ResponseAllOfEventsInner) GetRenderedDescription() string`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRenderedDescriptionOk() (*string, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *GetEvents200ResponseAllOfEventsInner) SetRenderedDescription(v string)`

SetRenderedDescription sets RenderedDescription field to given value.

### HasRenderedDescription

`func (o *GetEvents200ResponseAllOfEventsInner) HasRenderedDescription() bool`

HasRenderedDescription returns a boolean if a field has been set.

### GetHistoryPublicToSubscribers

`func (o *GetEvents200ResponseAllOfEventsInner) GetHistoryPublicToSubscribers() bool`

GetHistoryPublicToSubscribers returns the HistoryPublicToSubscribers field if non-nil, zero value otherwise.

### GetHistoryPublicToSubscribersOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetHistoryPublicToSubscribersOk() (*bool, bool)`

GetHistoryPublicToSubscribersOk returns a tuple with the HistoryPublicToSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPublicToSubscribers

`func (o *GetEvents200ResponseAllOfEventsInner) SetHistoryPublicToSubscribers(v bool)`

SetHistoryPublicToSubscribers sets HistoryPublicToSubscribers field to given value.

### HasHistoryPublicToSubscribers

`func (o *GetEvents200ResponseAllOfEventsInner) HasHistoryPublicToSubscribers() bool`

HasHistoryPublicToSubscribers returns a boolean if a field has been set.

### GetIsWebPublic

`func (o *GetEvents200ResponseAllOfEventsInner) GetIsWebPublic() bool`

GetIsWebPublic returns the IsWebPublic field if non-nil, zero value otherwise.

### GetIsWebPublicOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetIsWebPublicOk() (*bool, bool)`

GetIsWebPublicOk returns a tuple with the IsWebPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebPublic

`func (o *GetEvents200ResponseAllOfEventsInner) SetIsWebPublic(v bool)`

SetIsWebPublic sets IsWebPublic field to given value.

### HasIsWebPublic

`func (o *GetEvents200ResponseAllOfEventsInner) HasIsWebPublic() bool`

HasIsWebPublic returns a boolean if a field has been set.

### GetEmojiName

`func (o *GetEvents200ResponseAllOfEventsInner) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *GetEvents200ResponseAllOfEventsInner) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *GetEvents200ResponseAllOfEventsInner) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### GetEmojiCode

`func (o *GetEvents200ResponseAllOfEventsInner) GetEmojiCode() string`

GetEmojiCode returns the EmojiCode field if non-nil, zero value otherwise.

### GetEmojiCodeOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetEmojiCodeOk() (*string, bool)`

GetEmojiCodeOk returns a tuple with the EmojiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiCode

`func (o *GetEvents200ResponseAllOfEventsInner) SetEmojiCode(v string)`

SetEmojiCode sets EmojiCode field to given value.

### HasEmojiCode

`func (o *GetEvents200ResponseAllOfEventsInner) HasEmojiCode() bool`

HasEmojiCode returns a boolean if a field has been set.

### GetReactionType

`func (o *GetEvents200ResponseAllOfEventsInner) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *GetEvents200ResponseAllOfEventsInner) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *GetEvents200ResponseAllOfEventsInner) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### GetMessageId

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *GetEvents200ResponseAllOfEventsInner) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetAttachment

`func (o *GetEvents200ResponseAllOfEventsInner) GetAttachment() GetEvents200ResponseAllOfEventsInnerOneOf23Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetAttachmentOk() (*GetEvents200ResponseAllOfEventsInnerOneOf23Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *GetEvents200ResponseAllOfEventsInner) SetAttachment(v GetEvents200ResponseAllOfEventsInnerOneOf23Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *GetEvents200ResponseAllOfEventsInner) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetUploadSpaceUsed

`func (o *GetEvents200ResponseAllOfEventsInner) GetUploadSpaceUsed() int32`

GetUploadSpaceUsed returns the UploadSpaceUsed field if non-nil, zero value otherwise.

### GetUploadSpaceUsedOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetUploadSpaceUsedOk() (*int32, bool)`

GetUploadSpaceUsedOk returns a tuple with the UploadSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpaceUsed

`func (o *GetEvents200ResponseAllOfEventsInner) SetUploadSpaceUsed(v int32)`

SetUploadSpaceUsed sets UploadSpaceUsed field to given value.

### HasUploadSpaceUsed

`func (o *GetEvents200ResponseAllOfEventsInner) HasUploadSpaceUsed() bool`

HasUploadSpaceUsed returns a boolean if a field has been set.

### GetPushAccountId

`func (o *GetEvents200ResponseAllOfEventsInner) GetPushAccountId() string`

GetPushAccountId returns the PushAccountId field if non-nil, zero value otherwise.

### GetPushAccountIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetPushAccountIdOk() (*string, bool)`

GetPushAccountIdOk returns a tuple with the PushAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushAccountId

`func (o *GetEvents200ResponseAllOfEventsInner) SetPushAccountId(v string)`

SetPushAccountId sets PushAccountId field to given value.

### HasPushAccountId

`func (o *GetEvents200ResponseAllOfEventsInner) HasPushAccountId() bool`

HasPushAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *GetEvents200ResponseAllOfEventsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEvents200ResponseAllOfEventsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEvents200ResponseAllOfEventsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetEvents200ResponseAllOfEventsInner) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetEvents200ResponseAllOfEventsInner) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetEvents200ResponseAllOfEventsInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GetEvents200ResponseAllOfEventsInner) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GetEvents200ResponseAllOfEventsInner) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetMsgType

`func (o *GetEvents200ResponseAllOfEventsInner) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *GetEvents200ResponseAllOfEventsInner) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *GetEvents200ResponseAllOfEventsInner) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetContent

`func (o *GetEvents200ResponseAllOfEventsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetEvents200ResponseAllOfEventsInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetEvents200ResponseAllOfEventsInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSenderId

`func (o *GetEvents200ResponseAllOfEventsInner) GetSenderId() int32`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSenderIdOk() (*int32, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *GetEvents200ResponseAllOfEventsInner) SetSenderId(v int32)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *GetEvents200ResponseAllOfEventsInner) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetSubmessageId

`func (o *GetEvents200ResponseAllOfEventsInner) GetSubmessageId() int32`

GetSubmessageId returns the SubmessageId field if non-nil, zero value otherwise.

### GetSubmessageIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSubmessageIdOk() (*int32, bool)`

GetSubmessageIdOk returns a tuple with the SubmessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmessageId

`func (o *GetEvents200ResponseAllOfEventsInner) SetSubmessageId(v int32)`

SetSubmessageId sets SubmessageId field to given value.

### HasSubmessageId

`func (o *GetEvents200ResponseAllOfEventsInner) HasSubmessageId() bool`

HasSubmessageId returns a boolean if a field has been set.

### GetAway

`func (o *GetEvents200ResponseAllOfEventsInner) GetAway() bool`

GetAway returns the Away field if non-nil, zero value otherwise.

### GetAwayOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetAwayOk() (*bool, bool)`

GetAwayOk returns a tuple with the Away field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAway

`func (o *GetEvents200ResponseAllOfEventsInner) SetAway(v bool)`

SetAway sets Away field to given value.

### HasAway

`func (o *GetEvents200ResponseAllOfEventsInner) HasAway() bool`

HasAway returns a boolean if a field has been set.

### GetStatusText

`func (o *GetEvents200ResponseAllOfEventsInner) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *GetEvents200ResponseAllOfEventsInner) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *GetEvents200ResponseAllOfEventsInner) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetFields

`func (o *GetEvents200ResponseAllOfEventsInner) GetFields() []CustomProfileField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetFieldsOk() (*[]CustomProfileField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetEvents200ResponseAllOfEventsInner) SetFields(v []CustomProfileField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GetEvents200ResponseAllOfEventsInner) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetDefaultStreamGroups

`func (o *GetEvents200ResponseAllOfEventsInner) GetDefaultStreamGroups() []DefaultChannelGroup`

GetDefaultStreamGroups returns the DefaultStreamGroups field if non-nil, zero value otherwise.

### GetDefaultStreamGroupsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetDefaultStreamGroupsOk() (*[]DefaultChannelGroup, bool)`

GetDefaultStreamGroupsOk returns a tuple with the DefaultStreamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStreamGroups

`func (o *GetEvents200ResponseAllOfEventsInner) SetDefaultStreamGroups(v []DefaultChannelGroup)`

SetDefaultStreamGroups sets DefaultStreamGroups field to given value.

### HasDefaultStreamGroups

`func (o *GetEvents200ResponseAllOfEventsInner) HasDefaultStreamGroups() bool`

HasDefaultStreamGroups returns a boolean if a field has been set.

### GetDefaultStreams

`func (o *GetEvents200ResponseAllOfEventsInner) GetDefaultStreams() []int32`

GetDefaultStreams returns the DefaultStreams field if non-nil, zero value otherwise.

### GetDefaultStreamsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetDefaultStreamsOk() (*[]int32, bool)`

GetDefaultStreamsOk returns a tuple with the DefaultStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStreams

`func (o *GetEvents200ResponseAllOfEventsInner) SetDefaultStreams(v []int32)`

SetDefaultStreams sets DefaultStreams field to given value.

### HasDefaultStreams

`func (o *GetEvents200ResponseAllOfEventsInner) HasDefaultStreams() bool`

HasDefaultStreams returns a boolean if a field has been set.

### GetMessageIds

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *GetEvents200ResponseAllOfEventsInner) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.


### GetMessageType

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *GetEvents200ResponseAllOfEventsInner) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *GetEvents200ResponseAllOfEventsInner) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTopic

`func (o *GetEvents200ResponseAllOfEventsInner) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GetEvents200ResponseAllOfEventsInner) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GetEvents200ResponseAllOfEventsInner) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetMutedTopics

`func (o *GetEvents200ResponseAllOfEventsInner) GetMutedTopics() [][]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner`

GetMutedTopics returns the MutedTopics field if non-nil, zero value otherwise.

### GetMutedTopicsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMutedTopicsOk() (*[][]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner, bool)`

GetMutedTopicsOk returns a tuple with the MutedTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedTopics

`func (o *GetEvents200ResponseAllOfEventsInner) SetMutedTopics(v [][]GetEvents200ResponseAllOfEventsInnerOneOf31MutedTopicsInnerInner)`

SetMutedTopics sets MutedTopics field to given value.

### HasMutedTopics

`func (o *GetEvents200ResponseAllOfEventsInner) HasMutedTopics() bool`

HasMutedTopics returns a boolean if a field has been set.

### GetTopicName

`func (o *GetEvents200ResponseAllOfEventsInner) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *GetEvents200ResponseAllOfEventsInner) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *GetEvents200ResponseAllOfEventsInner) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetEvents200ResponseAllOfEventsInner) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetEvents200ResponseAllOfEventsInner) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetEvents200ResponseAllOfEventsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVisibilityPolicy

`func (o *GetEvents200ResponseAllOfEventsInner) GetVisibilityPolicy() int32`

GetVisibilityPolicy returns the VisibilityPolicy field if non-nil, zero value otherwise.

### GetVisibilityPolicyOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetVisibilityPolicyOk() (*int32, bool)`

GetVisibilityPolicyOk returns a tuple with the VisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityPolicy

`func (o *GetEvents200ResponseAllOfEventsInner) SetVisibilityPolicy(v int32)`

SetVisibilityPolicy sets VisibilityPolicy field to given value.

### HasVisibilityPolicy

`func (o *GetEvents200ResponseAllOfEventsInner) HasVisibilityPolicy() bool`

HasVisibilityPolicy returns a boolean if a field has been set.

### GetMutedUsers

`func (o *GetEvents200ResponseAllOfEventsInner) GetMutedUsers() []GetEvents200ResponseAllOfEventsInnerOneOf33MutedUsersInner`

GetMutedUsers returns the MutedUsers field if non-nil, zero value otherwise.

### GetMutedUsersOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMutedUsersOk() (*[]GetEvents200ResponseAllOfEventsInnerOneOf33MutedUsersInner, bool)`

GetMutedUsersOk returns a tuple with the MutedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedUsers

`func (o *GetEvents200ResponseAllOfEventsInner) SetMutedUsers(v []GetEvents200ResponseAllOfEventsInnerOneOf33MutedUsersInner)`

SetMutedUsers sets MutedUsers field to given value.

### HasMutedUsers

`func (o *GetEvents200ResponseAllOfEventsInner) HasMutedUsers() bool`

HasMutedUsers returns a boolean if a field has been set.

### GetOnboardingSteps

`func (o *GetEvents200ResponseAllOfEventsInner) GetOnboardingSteps() []OnboardingStep`

GetOnboardingSteps returns the OnboardingSteps field if non-nil, zero value otherwise.

### GetOnboardingStepsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetOnboardingStepsOk() (*[]OnboardingStep, bool)`

GetOnboardingStepsOk returns a tuple with the OnboardingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingSteps

`func (o *GetEvents200ResponseAllOfEventsInner) SetOnboardingSteps(v []OnboardingStep)`

SetOnboardingSteps sets OnboardingSteps field to given value.

### HasOnboardingSteps

`func (o *GetEvents200ResponseAllOfEventsInner) HasOnboardingSteps() bool`

HasOnboardingSteps returns a boolean if a field has been set.

### GetRenderingOnly

`func (o *GetEvents200ResponseAllOfEventsInner) GetRenderingOnly() bool`

GetRenderingOnly returns the RenderingOnly field if non-nil, zero value otherwise.

### GetRenderingOnlyOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRenderingOnlyOk() (*bool, bool)`

GetRenderingOnlyOk returns a tuple with the RenderingOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderingOnly

`func (o *GetEvents200ResponseAllOfEventsInner) SetRenderingOnly(v bool)`

SetRenderingOnly sets RenderingOnly field to given value.


### GetEditTimestamp

`func (o *GetEvents200ResponseAllOfEventsInner) GetEditTimestamp() int32`

GetEditTimestamp returns the EditTimestamp field if non-nil, zero value otherwise.

### GetEditTimestampOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetEditTimestampOk() (*int32, bool)`

GetEditTimestampOk returns a tuple with the EditTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditTimestamp

`func (o *GetEvents200ResponseAllOfEventsInner) SetEditTimestamp(v int32)`

SetEditTimestamp sets EditTimestamp field to given value.


### GetStreamName

`func (o *GetEvents200ResponseAllOfEventsInner) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *GetEvents200ResponseAllOfEventsInner) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *GetEvents200ResponseAllOfEventsInner) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### GetNewStreamId

`func (o *GetEvents200ResponseAllOfEventsInner) GetNewStreamId() int32`

GetNewStreamId returns the NewStreamId field if non-nil, zero value otherwise.

### GetNewStreamIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetNewStreamIdOk() (*int32, bool)`

GetNewStreamIdOk returns a tuple with the NewStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStreamId

`func (o *GetEvents200ResponseAllOfEventsInner) SetNewStreamId(v int32)`

SetNewStreamId sets NewStreamId field to given value.

### HasNewStreamId

`func (o *GetEvents200ResponseAllOfEventsInner) HasNewStreamId() bool`

HasNewStreamId returns a boolean if a field has been set.

### GetPropagateMode

`func (o *GetEvents200ResponseAllOfEventsInner) GetPropagateMode() string`

GetPropagateMode returns the PropagateMode field if non-nil, zero value otherwise.

### GetPropagateModeOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetPropagateModeOk() (*string, bool)`

GetPropagateModeOk returns a tuple with the PropagateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateMode

`func (o *GetEvents200ResponseAllOfEventsInner) SetPropagateMode(v string)`

SetPropagateMode sets PropagateMode field to given value.

### HasPropagateMode

`func (o *GetEvents200ResponseAllOfEventsInner) HasPropagateMode() bool`

HasPropagateMode returns a boolean if a field has been set.

### GetOrigSubject

`func (o *GetEvents200ResponseAllOfEventsInner) GetOrigSubject() string`

GetOrigSubject returns the OrigSubject field if non-nil, zero value otherwise.

### GetOrigSubjectOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetOrigSubjectOk() (*string, bool)`

GetOrigSubjectOk returns a tuple with the OrigSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigSubject

`func (o *GetEvents200ResponseAllOfEventsInner) SetOrigSubject(v string)`

SetOrigSubject sets OrigSubject field to given value.

### HasOrigSubject

`func (o *GetEvents200ResponseAllOfEventsInner) HasOrigSubject() bool`

HasOrigSubject returns a boolean if a field has been set.

### GetSubject

`func (o *GetEvents200ResponseAllOfEventsInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetEvents200ResponseAllOfEventsInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetEvents200ResponseAllOfEventsInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTopicLinks

`func (o *GetEvents200ResponseAllOfEventsInner) GetTopicLinks() []GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner`

GetTopicLinks returns the TopicLinks field if non-nil, zero value otherwise.

### GetTopicLinksOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetTopicLinksOk() (*[]GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner, bool)`

GetTopicLinksOk returns a tuple with the TopicLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicLinks

`func (o *GetEvents200ResponseAllOfEventsInner) SetTopicLinks(v []GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner)`

SetTopicLinks sets TopicLinks field to given value.

### HasTopicLinks

`func (o *GetEvents200ResponseAllOfEventsInner) HasTopicLinks() bool`

HasTopicLinks returns a boolean if a field has been set.

### GetOrigContent

`func (o *GetEvents200ResponseAllOfEventsInner) GetOrigContent() string`

GetOrigContent returns the OrigContent field if non-nil, zero value otherwise.

### GetOrigContentOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetOrigContentOk() (*string, bool)`

GetOrigContentOk returns a tuple with the OrigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigContent

`func (o *GetEvents200ResponseAllOfEventsInner) SetOrigContent(v string)`

SetOrigContent sets OrigContent field to given value.

### HasOrigContent

`func (o *GetEvents200ResponseAllOfEventsInner) HasOrigContent() bool`

HasOrigContent returns a boolean if a field has been set.

### GetOrigRenderedContent

`func (o *GetEvents200ResponseAllOfEventsInner) GetOrigRenderedContent() string`

GetOrigRenderedContent returns the OrigRenderedContent field if non-nil, zero value otherwise.

### GetOrigRenderedContentOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetOrigRenderedContentOk() (*string, bool)`

GetOrigRenderedContentOk returns a tuple with the OrigRenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigRenderedContent

`func (o *GetEvents200ResponseAllOfEventsInner) SetOrigRenderedContent(v string)`

SetOrigRenderedContent sets OrigRenderedContent field to given value.

### HasOrigRenderedContent

`func (o *GetEvents200ResponseAllOfEventsInner) HasOrigRenderedContent() bool`

HasOrigRenderedContent returns a boolean if a field has been set.

### GetRenderedContent

`func (o *GetEvents200ResponseAllOfEventsInner) GetRenderedContent() string`

GetRenderedContent returns the RenderedContent field if non-nil, zero value otherwise.

### GetRenderedContentOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRenderedContentOk() (*string, bool)`

GetRenderedContentOk returns a tuple with the RenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedContent

`func (o *GetEvents200ResponseAllOfEventsInner) SetRenderedContent(v string)`

SetRenderedContent sets RenderedContent field to given value.

### HasRenderedContent

`func (o *GetEvents200ResponseAllOfEventsInner) HasRenderedContent() bool`

HasRenderedContent returns a boolean if a field has been set.

### GetIsMeMessage

`func (o *GetEvents200ResponseAllOfEventsInner) GetIsMeMessage() bool`

GetIsMeMessage returns the IsMeMessage field if non-nil, zero value otherwise.

### GetIsMeMessageOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetIsMeMessageOk() (*bool, bool)`

GetIsMeMessageOk returns a tuple with the IsMeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeMessage

`func (o *GetEvents200ResponseAllOfEventsInner) SetIsMeMessage(v bool)`

SetIsMeMessage sets IsMeMessage field to given value.

### HasIsMeMessage

`func (o *GetEvents200ResponseAllOfEventsInner) HasIsMeMessage() bool`

HasIsMeMessage returns a boolean if a field has been set.

### GetSender

`func (o *GetEvents200ResponseAllOfEventsInner) GetSender() GetEvents200ResponseAllOfEventsInnerOneOf38Sender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSenderOk() (*GetEvents200ResponseAllOfEventsInnerOneOf38Sender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GetEvents200ResponseAllOfEventsInner) SetSender(v GetEvents200ResponseAllOfEventsInnerOneOf38Sender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *GetEvents200ResponseAllOfEventsInner) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipients

`func (o *GetEvents200ResponseAllOfEventsInner) GetRecipients() []GetEvents200ResponseAllOfEventsInnerOneOf38RecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRecipientsOk() (*[]GetEvents200ResponseAllOfEventsInnerOneOf38RecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetEvents200ResponseAllOfEventsInner) SetRecipients(v []GetEvents200ResponseAllOfEventsInnerOneOf38RecipientsInner)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *GetEvents200ResponseAllOfEventsInner) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetRecipient

`func (o *GetEvents200ResponseAllOfEventsInner) GetRecipient() GetEvents200ResponseAllOfEventsInnerOneOf39Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRecipientOk() (*GetEvents200ResponseAllOfEventsInnerOneOf39Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *GetEvents200ResponseAllOfEventsInner) SetRecipient(v GetEvents200ResponseAllOfEventsInnerOneOf39Recipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *GetEvents200ResponseAllOfEventsInner) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetOperation

`func (o *GetEvents200ResponseAllOfEventsInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GetEvents200ResponseAllOfEventsInner) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GetEvents200ResponseAllOfEventsInner) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetFlag

`func (o *GetEvents200ResponseAllOfEventsInner) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *GetEvents200ResponseAllOfEventsInner) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *GetEvents200ResponseAllOfEventsInner) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetMessages

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessages() []int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessagesOk() (*[]int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetEvents200ResponseAllOfEventsInner) SetMessages(v []int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *GetEvents200ResponseAllOfEventsInner) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetAll

`func (o *GetEvents200ResponseAllOfEventsInner) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetEvents200ResponseAllOfEventsInner) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetEvents200ResponseAllOfEventsInner) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetMessageDetails

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageDetails() map[string]GetEvents200ResponseAllOfEventsInnerOneOf42MessageDetailsValue`

GetMessageDetails returns the MessageDetails field if non-nil, zero value otherwise.

### GetMessageDetailsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetMessageDetailsOk() (*map[string]GetEvents200ResponseAllOfEventsInnerOneOf42MessageDetailsValue, bool)`

GetMessageDetailsOk returns a tuple with the MessageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetails

`func (o *GetEvents200ResponseAllOfEventsInner) SetMessageDetails(v map[string]GetEvents200ResponseAllOfEventsInnerOneOf42MessageDetailsValue)`

SetMessageDetails sets MessageDetails field to given value.

### HasMessageDetails

`func (o *GetEvents200ResponseAllOfEventsInner) HasMessageDetails() bool`

HasMessageDetails returns a boolean if a field has been set.

### GetGroup

`func (o *GetEvents200ResponseAllOfEventsInner) GetGroup() UserGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetGroupOk() (*UserGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetEvents200ResponseAllOfEventsInner) SetGroup(v UserGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetEvents200ResponseAllOfEventsInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupId

`func (o *GetEvents200ResponseAllOfEventsInner) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GetEvents200ResponseAllOfEventsInner) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GetEvents200ResponseAllOfEventsInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetData

`func (o *GetEvents200ResponseAllOfEventsInner) GetData() GetEvents200ResponseAllOfEventsInnerOneOf84Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetDataOk() (*GetEvents200ResponseAllOfEventsInnerOneOf84Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEvents200ResponseAllOfEventsInner) SetData(v GetEvents200ResponseAllOfEventsInnerOneOf84Data)`

SetData sets Data field to given value.

### HasData

`func (o *GetEvents200ResponseAllOfEventsInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDirectSubgroupIds

`func (o *GetEvents200ResponseAllOfEventsInner) GetDirectSubgroupIds() []int32`

GetDirectSubgroupIds returns the DirectSubgroupIds field if non-nil, zero value otherwise.

### GetDirectSubgroupIdsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetDirectSubgroupIdsOk() (*[]int32, bool)`

GetDirectSubgroupIdsOk returns a tuple with the DirectSubgroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSubgroupIds

`func (o *GetEvents200ResponseAllOfEventsInner) SetDirectSubgroupIds(v []int32)`

SetDirectSubgroupIds sets DirectSubgroupIds field to given value.

### HasDirectSubgroupIds

`func (o *GetEvents200ResponseAllOfEventsInner) HasDirectSubgroupIds() bool`

HasDirectSubgroupIds returns a boolean if a field has been set.

### GetRealmLinkifiers

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmLinkifiers() []GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner`

GetRealmLinkifiers returns the RealmLinkifiers field if non-nil, zero value otherwise.

### GetRealmLinkifiersOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmLinkifiersOk() (*[]GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner, bool)`

GetRealmLinkifiersOk returns a tuple with the RealmLinkifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmLinkifiers

`func (o *GetEvents200ResponseAllOfEventsInner) SetRealmLinkifiers(v []GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner)`

SetRealmLinkifiers sets RealmLinkifiers field to given value.

### HasRealmLinkifiers

`func (o *GetEvents200ResponseAllOfEventsInner) HasRealmLinkifiers() bool`

HasRealmLinkifiers returns a boolean if a field has been set.

### GetRealmFilters

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmFilters() [][]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner`

GetRealmFilters returns the RealmFilters field if non-nil, zero value otherwise.

### GetRealmFiltersOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmFiltersOk() (*[][]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner, bool)`

GetRealmFiltersOk returns a tuple with the RealmFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmFilters

`func (o *GetEvents200ResponseAllOfEventsInner) SetRealmFilters(v [][]GetEvents200ResponseAllOfEventsInnerOneOf51RealmFiltersInnerInner)`

SetRealmFilters sets RealmFilters field to given value.

### HasRealmFilters

`func (o *GetEvents200ResponseAllOfEventsInner) HasRealmFilters() bool`

HasRealmFilters returns a boolean if a field has been set.

### GetRealmPlaygrounds

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmPlaygrounds() []RealmPlayground`

GetRealmPlaygrounds returns the RealmPlaygrounds field if non-nil, zero value otherwise.

### GetRealmPlaygroundsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmPlaygroundsOk() (*[]RealmPlayground, bool)`

GetRealmPlaygroundsOk returns a tuple with the RealmPlaygrounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPlaygrounds

`func (o *GetEvents200ResponseAllOfEventsInner) SetRealmPlaygrounds(v []RealmPlayground)`

SetRealmPlaygrounds sets RealmPlaygrounds field to given value.

### HasRealmPlaygrounds

`func (o *GetEvents200ResponseAllOfEventsInner) HasRealmPlaygrounds() bool`

HasRealmPlaygrounds returns a boolean if a field has been set.

### GetRealmEmoji

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmEmoji() map[string]RealmEmoji`

GetRealmEmoji returns the RealmEmoji field if non-nil, zero value otherwise.

### GetRealmEmojiOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmEmojiOk() (*map[string]RealmEmoji, bool)`

GetRealmEmojiOk returns a tuple with the RealmEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmoji

`func (o *GetEvents200ResponseAllOfEventsInner) SetRealmEmoji(v map[string]RealmEmoji)`

SetRealmEmoji sets RealmEmoji field to given value.

### HasRealmEmoji

`func (o *GetEvents200ResponseAllOfEventsInner) HasRealmEmoji() bool`

HasRealmEmoji returns a boolean if a field has been set.

### GetRealmDomain

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmDomain() GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain`

GetRealmDomain returns the RealmDomain field if non-nil, zero value otherwise.

### GetRealmDomainOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmDomainOk() (*GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain, bool)`

GetRealmDomainOk returns a tuple with the RealmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDomain

`func (o *GetEvents200ResponseAllOfEventsInner) SetRealmDomain(v GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain)`

SetRealmDomain sets RealmDomain field to given value.

### HasRealmDomain

`func (o *GetEvents200ResponseAllOfEventsInner) HasRealmDomain() bool`

HasRealmDomain returns a boolean if a field has been set.

### GetDomain

`func (o *GetEvents200ResponseAllOfEventsInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetEvents200ResponseAllOfEventsInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetEvents200ResponseAllOfEventsInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExports

`func (o *GetEvents200ResponseAllOfEventsInner) GetExports() []RealmExport`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetExportsOk() (*[]RealmExport, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *GetEvents200ResponseAllOfEventsInner) SetExports(v []RealmExport)`

SetExports sets Exports field to given value.

### HasExports

`func (o *GetEvents200ResponseAllOfEventsInner) HasExports() bool`

HasExports returns a boolean if a field has been set.

### GetConsented

`func (o *GetEvents200ResponseAllOfEventsInner) GetConsented() bool`

GetConsented returns the Consented field if non-nil, zero value otherwise.

### GetConsentedOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetConsentedOk() (*bool, bool)`

GetConsentedOk returns a tuple with the Consented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsented

`func (o *GetEvents200ResponseAllOfEventsInner) SetConsented(v bool)`

SetConsented sets Consented field to given value.

### HasConsented

`func (o *GetEvents200ResponseAllOfEventsInner) HasConsented() bool`

HasConsented returns a boolean if a field has been set.

### GetBot

`func (o *GetEvents200ResponseAllOfEventsInner) GetBot() GetEvents200ResponseAllOfEventsInnerOneOf62Bot`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetBotOk() (*GetEvents200ResponseAllOfEventsInnerOneOf62Bot, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *GetEvents200ResponseAllOfEventsInner) SetBot(v GetEvents200ResponseAllOfEventsInnerOneOf62Bot)`

SetBot sets Bot field to given value.

### HasBot

`func (o *GetEvents200ResponseAllOfEventsInner) HasBot() bool`

HasBot returns a boolean if a field has been set.

### GetRealmId

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmId() int32`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRealmIdOk() (*int32, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *GetEvents200ResponseAllOfEventsInner) SetRealmId(v int32)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *GetEvents200ResponseAllOfEventsInner) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetZulipVersion

`func (o *GetEvents200ResponseAllOfEventsInner) GetZulipVersion() string`

GetZulipVersion returns the ZulipVersion field if non-nil, zero value otherwise.

### GetZulipVersionOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetZulipVersionOk() (*string, bool)`

GetZulipVersionOk returns a tuple with the ZulipVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipVersion

`func (o *GetEvents200ResponseAllOfEventsInner) SetZulipVersion(v string)`

SetZulipVersion sets ZulipVersion field to given value.

### HasZulipVersion

`func (o *GetEvents200ResponseAllOfEventsInner) HasZulipVersion() bool`

HasZulipVersion returns a boolean if a field has been set.

### GetZulipMergeBase

`func (o *GetEvents200ResponseAllOfEventsInner) GetZulipMergeBase() string`

GetZulipMergeBase returns the ZulipMergeBase field if non-nil, zero value otherwise.

### GetZulipMergeBaseOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetZulipMergeBaseOk() (*string, bool)`

GetZulipMergeBaseOk returns a tuple with the ZulipMergeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipMergeBase

`func (o *GetEvents200ResponseAllOfEventsInner) SetZulipMergeBase(v string)`

SetZulipMergeBase sets ZulipMergeBase field to given value.

### HasZulipMergeBase

`func (o *GetEvents200ResponseAllOfEventsInner) HasZulipMergeBase() bool`

HasZulipMergeBase returns a boolean if a field has been set.

### GetZulipFeatureLevel

`func (o *GetEvents200ResponseAllOfEventsInner) GetZulipFeatureLevel() int32`

GetZulipFeatureLevel returns the ZulipFeatureLevel field if non-nil, zero value otherwise.

### GetZulipFeatureLevelOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetZulipFeatureLevelOk() (*int32, bool)`

GetZulipFeatureLevelOk returns a tuple with the ZulipFeatureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipFeatureLevel

`func (o *GetEvents200ResponseAllOfEventsInner) SetZulipFeatureLevel(v int32)`

SetZulipFeatureLevel sets ZulipFeatureLevel field to given value.

### HasZulipFeatureLevel

`func (o *GetEvents200ResponseAllOfEventsInner) HasZulipFeatureLevel() bool`

HasZulipFeatureLevel returns a boolean if a field has been set.

### GetServerGeneration

`func (o *GetEvents200ResponseAllOfEventsInner) GetServerGeneration() int32`

GetServerGeneration returns the ServerGeneration field if non-nil, zero value otherwise.

### GetServerGenerationOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetServerGenerationOk() (*int32, bool)`

GetServerGenerationOk returns a tuple with the ServerGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGeneration

`func (o *GetEvents200ResponseAllOfEventsInner) SetServerGeneration(v int32)`

SetServerGeneration sets ServerGeneration field to given value.

### HasServerGeneration

`func (o *GetEvents200ResponseAllOfEventsInner) HasServerGeneration() bool`

HasServerGeneration returns a boolean if a field has been set.

### GetImmediate

`func (o *GetEvents200ResponseAllOfEventsInner) GetImmediate() bool`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetImmediateOk() (*bool, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *GetEvents200ResponseAllOfEventsInner) SetImmediate(v bool)`

SetImmediate sets Immediate field to given value.

### HasImmediate

`func (o *GetEvents200ResponseAllOfEventsInner) HasImmediate() bool`

HasImmediate returns a boolean if a field has been set.

### GetDrafts

`func (o *GetEvents200ResponseAllOfEventsInner) GetDrafts() []Draft`

GetDrafts returns the Drafts field if non-nil, zero value otherwise.

### GetDraftsOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetDraftsOk() (*[]Draft, bool)`

GetDraftsOk returns a tuple with the Drafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrafts

`func (o *GetEvents200ResponseAllOfEventsInner) SetDrafts(v []Draft)`

SetDrafts sets Drafts field to given value.

### HasDrafts

`func (o *GetEvents200ResponseAllOfEventsInner) HasDrafts() bool`

HasDrafts returns a boolean if a field has been set.

### GetDraft

`func (o *GetEvents200ResponseAllOfEventsInner) GetDraft() Draft`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetDraftOk() (*Draft, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *GetEvents200ResponseAllOfEventsInner) SetDraft(v Draft)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *GetEvents200ResponseAllOfEventsInner) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetDraftId

`func (o *GetEvents200ResponseAllOfEventsInner) GetDraftId() int32`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetDraftIdOk() (*int32, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *GetEvents200ResponseAllOfEventsInner) SetDraftId(v int32)`

SetDraftId sets DraftId field to given value.

### HasDraftId

`func (o *GetEvents200ResponseAllOfEventsInner) HasDraftId() bool`

HasDraftId returns a boolean if a field has been set.

### GetNavigationView

`func (o *GetEvents200ResponseAllOfEventsInner) GetNavigationView() NavigationView`

GetNavigationView returns the NavigationView field if non-nil, zero value otherwise.

### GetNavigationViewOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetNavigationViewOk() (*NavigationView, bool)`

GetNavigationViewOk returns a tuple with the NavigationView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationView

`func (o *GetEvents200ResponseAllOfEventsInner) SetNavigationView(v NavigationView)`

SetNavigationView sets NavigationView field to given value.

### HasNavigationView

`func (o *GetEvents200ResponseAllOfEventsInner) HasNavigationView() bool`

HasNavigationView returns a boolean if a field has been set.

### GetFragment

`func (o *GetEvents200ResponseAllOfEventsInner) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *GetEvents200ResponseAllOfEventsInner) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *GetEvents200ResponseAllOfEventsInner) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetSavedSnippet

`func (o *GetEvents200ResponseAllOfEventsInner) GetSavedSnippet() SavedSnippet`

GetSavedSnippet returns the SavedSnippet field if non-nil, zero value otherwise.

### GetSavedSnippetOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSavedSnippetOk() (*SavedSnippet, bool)`

GetSavedSnippetOk returns a tuple with the SavedSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippet

`func (o *GetEvents200ResponseAllOfEventsInner) SetSavedSnippet(v SavedSnippet)`

SetSavedSnippet sets SavedSnippet field to given value.

### HasSavedSnippet

`func (o *GetEvents200ResponseAllOfEventsInner) HasSavedSnippet() bool`

HasSavedSnippet returns a boolean if a field has been set.

### GetSavedSnippetId

`func (o *GetEvents200ResponseAllOfEventsInner) GetSavedSnippetId() int32`

GetSavedSnippetId returns the SavedSnippetId field if non-nil, zero value otherwise.

### GetSavedSnippetIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetSavedSnippetIdOk() (*int32, bool)`

GetSavedSnippetIdOk returns a tuple with the SavedSnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippetId

`func (o *GetEvents200ResponseAllOfEventsInner) SetSavedSnippetId(v int32)`

SetSavedSnippetId sets SavedSnippetId field to given value.

### HasSavedSnippetId

`func (o *GetEvents200ResponseAllOfEventsInner) HasSavedSnippetId() bool`

HasSavedSnippetId returns a boolean if a field has been set.

### GetReminders

`func (o *GetEvents200ResponseAllOfEventsInner) GetReminders() []Reminder`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetRemindersOk() (*[]Reminder, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *GetEvents200ResponseAllOfEventsInner) SetReminders(v []Reminder)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *GetEvents200ResponseAllOfEventsInner) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetReminderId

`func (o *GetEvents200ResponseAllOfEventsInner) GetReminderId() int32`

GetReminderId returns the ReminderId field if non-nil, zero value otherwise.

### GetReminderIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetReminderIdOk() (*int32, bool)`

GetReminderIdOk returns a tuple with the ReminderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderId

`func (o *GetEvents200ResponseAllOfEventsInner) SetReminderId(v int32)`

SetReminderId sets ReminderId field to given value.

### HasReminderId

`func (o *GetEvents200ResponseAllOfEventsInner) HasReminderId() bool`

HasReminderId returns a boolean if a field has been set.

### GetScheduledMessages

`func (o *GetEvents200ResponseAllOfEventsInner) GetScheduledMessages() []ScheduledMessage`

GetScheduledMessages returns the ScheduledMessages field if non-nil, zero value otherwise.

### GetScheduledMessagesOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetScheduledMessagesOk() (*[]ScheduledMessage, bool)`

GetScheduledMessagesOk returns a tuple with the ScheduledMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessages

`func (o *GetEvents200ResponseAllOfEventsInner) SetScheduledMessages(v []ScheduledMessage)`

SetScheduledMessages sets ScheduledMessages field to given value.

### HasScheduledMessages

`func (o *GetEvents200ResponseAllOfEventsInner) HasScheduledMessages() bool`

HasScheduledMessages returns a boolean if a field has been set.

### GetScheduledMessage

`func (o *GetEvents200ResponseAllOfEventsInner) GetScheduledMessage() ScheduledMessage`

GetScheduledMessage returns the ScheduledMessage field if non-nil, zero value otherwise.

### GetScheduledMessageOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetScheduledMessageOk() (*ScheduledMessage, bool)`

GetScheduledMessageOk returns a tuple with the ScheduledMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessage

`func (o *GetEvents200ResponseAllOfEventsInner) SetScheduledMessage(v ScheduledMessage)`

SetScheduledMessage sets ScheduledMessage field to given value.

### HasScheduledMessage

`func (o *GetEvents200ResponseAllOfEventsInner) HasScheduledMessage() bool`

HasScheduledMessage returns a boolean if a field has been set.

### GetScheduledMessageId

`func (o *GetEvents200ResponseAllOfEventsInner) GetScheduledMessageId() int32`

GetScheduledMessageId returns the ScheduledMessageId field if non-nil, zero value otherwise.

### GetScheduledMessageIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetScheduledMessageIdOk() (*int32, bool)`

GetScheduledMessageIdOk returns a tuple with the ScheduledMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessageId

`func (o *GetEvents200ResponseAllOfEventsInner) SetScheduledMessageId(v int32)`

SetScheduledMessageId sets ScheduledMessageId field to given value.

### HasScheduledMessageId

`func (o *GetEvents200ResponseAllOfEventsInner) HasScheduledMessageId() bool`

HasScheduledMessageId returns a boolean if a field has been set.

### GetChannelFolder

`func (o *GetEvents200ResponseAllOfEventsInner) GetChannelFolder() ChannelFolder`

GetChannelFolder returns the ChannelFolder field if non-nil, zero value otherwise.

### GetChannelFolderOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetChannelFolderOk() (*ChannelFolder, bool)`

GetChannelFolderOk returns a tuple with the ChannelFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFolder

`func (o *GetEvents200ResponseAllOfEventsInner) SetChannelFolder(v ChannelFolder)`

SetChannelFolder sets ChannelFolder field to given value.

### HasChannelFolder

`func (o *GetEvents200ResponseAllOfEventsInner) HasChannelFolder() bool`

HasChannelFolder returns a boolean if a field has been set.

### GetChannelFolderId

`func (o *GetEvents200ResponseAllOfEventsInner) GetChannelFolderId() float32`

GetChannelFolderId returns the ChannelFolderId field if non-nil, zero value otherwise.

### GetChannelFolderIdOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetChannelFolderIdOk() (*float32, bool)`

GetChannelFolderIdOk returns a tuple with the ChannelFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFolderId

`func (o *GetEvents200ResponseAllOfEventsInner) SetChannelFolderId(v float32)`

SetChannelFolderId sets ChannelFolderId field to given value.

### HasChannelFolderId

`func (o *GetEvents200ResponseAllOfEventsInner) HasChannelFolderId() bool`

HasChannelFolderId returns a boolean if a field has been set.

### GetOrder

`func (o *GetEvents200ResponseAllOfEventsInner) GetOrder() []int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetEvents200ResponseAllOfEventsInner) GetOrderOk() (*[]int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetEvents200ResponseAllOfEventsInner) SetOrder(v []int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetEvents200ResponseAllOfEventsInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


