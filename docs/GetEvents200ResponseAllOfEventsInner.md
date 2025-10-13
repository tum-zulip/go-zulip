# EventEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | 
**Type** | **string** |  | 
**AlertWords** | Pointer to **[]string** | An array of strings, where each string is an alert word (or phrase) configured by the user.  | [optional] 
**User** | Pointer to **interface{}** |  | [optional] 
**SettingName** | Pointer to **string** | Name of the changed display setting.  | [optional] 
**Setting** | Pointer to [**EventEnvelopeOneOf1Setting**](EventEnvelopeOneOf1Setting.md) |  | [optional] 
**LanguageName** | Pointer to **string** | Present only if the setting to be changed is &#x60;default_language&#x60;. Contains the name of the new default language in English.  | [optional] 
**NotificationName** | Pointer to **string** | Name of the changed notification setting.  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **string** | The name of the property that was changed.  | [optional] 
**Value** | Pointer to [**EventEnvelopeOneOf68Value**](EventEnvelopeOneOf68Value.md) |  | [optional] 
**Person** | Pointer to [**EventEnvelopeOneOf14Person**](EventEnvelopeOneOf14Person.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]EventEnvelopeOneOf6SubscriptionsInner**](EventEnvelopeOneOf6SubscriptionsInner.md) | A list of dictionaries, where each dictionary contains information about one of the newly unsubscribed channels.  | [optional] 
**StreamId** | Pointer to **int32** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  The unique ID of the channel to which message is being typed.  **Changes**: New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  | [optional] 
**StreamIds** | Pointer to **[]int32** | Array containing the IDs of the channels that were deleted.  **Changes**: New in Zulip 10.0 (feature level 343). Previously, these IDs were available only via the legacy &#x60;streams&#x60; array.  | [optional] 
**UserIds** | Pointer to **[]int32** | Array containing the IDs of the users who have been removed from the user group.  | [optional] 
**Message** | Pointer to [**MessagesEvent**](MessagesEvent.md) |  | [optional] 
**Flags** | **[]string** | The user&#39;s personal [message flags][message-flags] for the message with ID &#x60;message_id&#x60; following the edit.  A client application should compare these to the original flags to identify cases where a mention or alert word was added by the edit.  **Changes**: In Zulip 8.0 (feature level 224), the &#x60;wildcard_mentioned&#x60; flag was deprecated in favor of the &#x60;stream_wildcard_mentioned&#x60; and &#x60;topic_wildcard_mentioned&#x60; flags. The &#x60;wildcard_mentioned&#x60; flag exists for backwards compatibility with older clients and equals &#x60;stream_wildcard_mentioned || topic_wildcard_mentioned&#x60;. Clients supporting older server versions should treat this field as a previous name for the &#x60;stream_wildcard_mentioned&#x60; flag as topic wildcard mentions were not available prior to this feature level.  [message-flags]: /api/update-message-flags#available-flags  | 
**Presences** | Pointer to [**map[string]ModernPresenceFormat**](ModernPresenceFormat.md) | Only present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  A dictionary mapping user IDs to the presence data (modern format) for the modified user(s). Clients should support updating multiple users in a single event.  **Changes**: New in Zulip 11.0 (feature level 419).  | [optional] 
**UserId** | **int32** | The ID of the user whose setting was changed.  | 
**Email** | Pointer to **string** | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the &#x60;user_id&#x60;.  | [optional] 
**ServerTimestamp** | Pointer to **float32** | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  The timestamp of when the Zulip server received the user&#39;s presence as a UNIX timestamp.  | [optional] 
**Presence** | Pointer to [**map[string]EventEnvelopeOneOf15PresenceValue**](EventEnvelopeOneOf15PresenceValue.md) | Not present for clients that support the &#x60;simplified_presence_events&#x60; [client capability](/api/register-queue#parameter-client_capabilities).  Object containing the presence data (legacy format) of of the modified user.  | [optional] 
**Streams** | Pointer to [**[]EventEnvelopeOneOf17StreamsInner**](EventEnvelopeOneOf17StreamsInner.md) | Array of objects, each containing ID of the channel that was deleted.  **Changes**: **Deprecated** in Zulip 10.0 (feature level 343) and will be removed in a future release. Previously, these objects additionally contained all the standard fields for a channel object.  | [optional] 
**Name** | Pointer to **string** | The name of the channel whose details have changed.  | [optional] 
**RenderedDescription** | Pointer to **string** | Note: Only present if the changed property was &#x60;description&#x60;.  The short description of the channel rendered as HTML, intended to be used when displaying the channel description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**HistoryPublicToSubscribers** | Pointer to **bool** | Note: Only present if the changed property was &#x60;invite_only&#x60;.  Whether the history of the channel is public to its subscribers.  Currently always true for public channels (i.e. &#x60;\&quot;invite_only\&quot;: false&#x60; implies &#x60;\&quot;history_public_to_subscribers\&quot;: true&#x60;), but clients should not make that assumption, as we may change that behavior in the future.  | [optional] 
**IsWebPublic** | Pointer to **bool** | Note: Only present if the changed property was &#x60;invite_only&#x60;.  Whether the channel&#39;s history is now readable by web-public spectators.  **Changes**: New in Zulip 5.0 (feature level 71).  | [optional] 
**EmojiName** | Pointer to **string** | The [emoji name](/api/update-status#parameter-emoji_name) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**EmojiCode** | Pointer to **string** | The [emoji code](/api/update-status#parameter-emoji_code) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**ReactionType** | Pointer to **string** | The [emoji type](/api/update-status#parameter-reaction_type) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**MessageId** | **int32** | Indicates the message id of the message that is being edited.  | 
**Attachment** | Pointer to [**EventEnvelopeOneOf23Attachment**](EventEnvelopeOneOf23Attachment.md) |  | [optional] 
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
**MutedTopics** | Pointer to [**[][]EventEnvelopeOneOf31MutedTopicsInnerInner**]([]EventEnvelopeOneOf31MutedTopicsInnerInner.md) | Array of tuples, where each tuple describes a muted topic. The first element of the tuple is the channel name in which the topic has to be muted, the second element is the topic name to be muted and the third element is an integer UNIX timestamp representing when the topic was muted.  **Changes**: Deprecated in Zulip 6.0 (feature level 134). Starting with this version, clients that explicitly requested the replacement &#x60;user_topic&#x60; event type when registering their event queue will not receive this legacy event type.  Before Zulip 3.0 (feature level 1), the &#x60;muted_topics&#x60; array objects were 2-item tuples and did not include the timestamp information for when the topic was muted.  | [optional] 
**TopicName** | Pointer to **string** | The name of the topic.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**LastUpdated** | Pointer to **int32** | An integer UNIX timestamp representing when the user-topic relationship was last changed.  | [optional] 
**VisibilityPolicy** | Pointer to **int32** | An integer indicating the user&#39;s visibility preferences for the topic, such as whether the topic is muted.  - 0 &#x3D; None. Used to indicate that the user no   longer has a special visibility policy for this topic. - 1 &#x3D; Muted. Used to record [muted topics](/help/mute-a-topic). - 2 &#x3D; Unmuted. Used to record unmuted topics. - 3 &#x3D; Followed. Used to record [followed topics](/help/follow-a-topic).  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.  In Zulip 7.0 (feature level 170), added unmuted as a visibility policy option.  | [optional] 
**MutedUsers** | Pointer to [**[]EventEnvelopeOneOf33MutedUsersInner**](EventEnvelopeOneOf33MutedUsersInner.md) | A list of dictionaries where each dictionary describes a muted user.  | [optional] 
**OnboardingSteps** | Pointer to [**[]OnboardingStep**](OnboardingStep.md) | An array of dictionaries where each dictionary contains details about a single onboarding step.  **Changes**: Before Zulip 8.0 (feature level 233), this array was named &#x60;hotspots&#x60;. Prior to this feature level, one-time notice onboarding steps were not supported, and the &#x60;type&#x60; field in these objects did not exist as all onboarding steps were implicitly hotspots.  | [optional] 
**RenderingOnly** | **bool** | Whether the event only updates the rendered content of the message.  This field should be used by clients to determine if the event only provides a rendering update to the message content, such as for an [inline URL preview][inline-url-previews]. When &#x60;true&#x60;, the event does not reflect a user-generated edit and does not modify the message history.  **Changes**: New in Zulip 5.0 (feature level 114). Clients can correctly identify these rendering update event with earlier Zulip versions by checking whether the &#x60;user_id&#x60; field was omitted.  | 
**EditTimestamp** | **int32** | The time when this message edit operation was processed by the server.  **Changes**: As of Zulip 5.0 (feature level 114), this field is present for all &#x60;update_message&#x60; events. Previously, this field was omitted for [inline URL preview][inline-url-previews] updates.  | 
**StreamName** | Pointer to **string** | Only present if the message was edited and originally sent to a channel.  The name of the channel that the message was sent to. Clients are recommended to use the &#x60;stream_id&#x60; field instead.  | [optional] 
**NewStreamId** | Pointer to **int32** | Only present if message(s) were moved to a different channel.  The post-edit channel for all of the messages with IDs in &#x60;message_ids&#x60;.  | [optional] 
**PropagateMode** | Pointer to **string** | Only present if this event moved messages to a different topic and/or channel.  The choice the editing user made about which messages should be affected by a channel/topic edit:  - &#x60;\&quot;change_one\&quot;&#x60;: Just change the one indicated in &#x60;message_id&#x60;. - &#x60;\&quot;change_later\&quot;&#x60;: Change messages in the same topic that had   been sent after this one. - &#x60;\&quot;change_all\&quot;&#x60;: Change all messages in that topic.  This parameter should be used to decide whether to change navigation and compose box state in response to the edit. For example, if the user was previously in topic narrow, and the topic was edited with &#x60;\&quot;change_later\&quot;&#x60; or &#x60;\&quot;change_all\&quot;&#x60;, the Zulip web app will automatically navigate to the new topic narrow. Similarly, a message being composed to the old topic should have its recipient changed to the new topic.  This navigation makes it much more convenient to move content between topics without disruption or messages continuing to be sent to the pre-edit topic by accident.  | [optional] 
**OrigSubject** | Pointer to **string** | Only present if this event moved messages to a different topic and/or channel.  The pre-edit topic for all of the messages with IDs in &#x60;message_ids&#x60;.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual pre-edit topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**Subject** | Pointer to **string** | Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  The post-edit topic for all of the messages with IDs in &#x60;message_ids&#x60;.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual post-edit topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**TopicLinks** | Pointer to [**[]EventEnvelopeOneOf36TopicLinksInner**](EventEnvelopeOneOf36TopicLinksInner.md) | Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  Data on any links to be included in the &#x60;topic&#x60; line (these are generated by [custom linkification filter](/help/add-a-custom-linkifier) that match content in the message&#39;s topic.), corresponding to the post-edit topic.  **Changes**: This field contained a list of urls before Zulip 4.0 (feature level 46).  New in Zulip 3.0 (feature level 1). Previously, this field was called &#x60;subject_links&#x60;; clients are recommended to rename &#x60;subject_links&#x60; to &#x60;topic_links&#x60; if present for compatibility with older Zulip servers.  | [optional] 
**OrigContent** | Pointer to **string** | Only present if this event changed the message content.  The original content of the message with ID &#x60;message_id&#x60; immediately prior to this edit, in the original [Zulip-flavored Markdown](/help/format-your-message-using-markdown) format.  | [optional] 
**OrigRenderedContent** | Pointer to **string** | Only present if this event changed the message content.  The original content of the message with ID &#x60;message_id&#x60; immediately prior to this edit, rendered as HTML.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**RenderedContent** | Pointer to **string** | Only present if this event changed the message content or updated the message content for an [inline URL preview][inline-url-previews].  The new content of the message with ID &#x60;message_id&#x60;, rendered in HTML.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**IsMeMessage** | Pointer to **bool** | Only present if this event changed the message content.  Whether the message with ID &#x60;message_id&#x60; is now a [/me status message][status-messages].  [status-messages]: /help/format-your-message-using-markdown#status-messages  | [optional] 
**Sender** | Pointer to [**EventEnvelopeOneOf38Sender**](EventEnvelopeOneOf38Sender.md) |  | [optional] 
**Recipients** | Pointer to [**[]EventEnvelopeOneOf38RecipientsInner**](EventEnvelopeOneOf38RecipientsInner.md) | Only present if &#x60;message_type&#x60; is &#x60;\&quot;direct\&quot;&#x60;.  Array of dictionaries describing the set of users who would be recipients of the message that was previously being typed. Each dictionary contains details about one of the recipients. The sending user is guaranteed to appear among the recipients.  | [optional] 
**Recipient** | Pointer to [**EventEnvelopeOneOf39Recipient**](EventEnvelopeOneOf39Recipient.md) |  | [optional] 
**Operation** | Pointer to **string** | Old name for the &#x60;op&#x60; field in this event type.  **Deprecated** in Zulip 4.0 (feature level 32), and replaced by the &#x60;op&#x60; field.  | [optional] 
**Flag** | Pointer to **string** | The [flag][message-flags] to be removed.  | [optional] 
**Messages** | Pointer to **[]int32** | Array containing the IDs of the messages from which the flag was removed.  | [optional] 
**All** | Pointer to **bool** | Will be &#x60;false&#x60; for all specified flags.  **Deprecated** and will be removed in a future release.  | [optional] 
**MessageDetails** | Pointer to [**map[string]EventEnvelopeOneOf42MessageDetailsValue**](EventEnvelopeOneOf42MessageDetailsValue.md) | Only present if the specified &#x60;flag&#x60; is &#x60;\&quot;read\&quot;&#x60;.  A set of data structures describing the messages that are being marked as unread with additional details to allow clients to update the &#x60;unread_msgs&#x60; data structure for these messages (which may not be otherwise known to the client).  **Changes**: New in Zulip 5.0 (feature level 121). Previously, marking already read messages as unread was not supported by the Zulip API.  | [optional] 
**Group** | Pointer to [**UserGroup**](UserGroup.md) |  | [optional] 
**GroupId** | Pointer to **int32** | The ID of the group which has been deleted.  | [optional] 
**Data** | Pointer to [**EventEnvelopeOneOf84Data**](EventEnvelopeOneOf84Data.md) |  | [optional] 
**DirectSubgroupIds** | Pointer to **[]int32** | Array containing the IDs of the subgroups that have been removed from the user group.  **Changes**: New in Zulip 6.0 (feature level 131). Previously, this was called &#x60;subgroup_ids&#x60;, but clients can ignore older events as this feature level predates subgroups being fully implemented.  | [optional] 
**RealmLinkifiers** | Pointer to [**[]EventEnvelopeOneOf50RealmLinkifiersInner**](EventEnvelopeOneOf50RealmLinkifiersInner.md) | An ordered array of dictionaries where each dictionary contains details about a single linkifier.  Clients should always process linkifiers in the order given; this is important if the realm has linkifiers with overlapping patterns. The order can be modified using [&#x60;PATCH /realm/linkifiers&#x60;](/api/reorder-linkifiers).  | [optional] 
**RealmFilters** | Pointer to [**[][]EventEnvelopeOneOf51RealmFiltersInnerInner**]([]EventEnvelopeOneOf51RealmFiltersInnerInner.md) | An array of tuples, where each tuple described a linkifier. The first element of the tuple was a string regex pattern which represented the pattern to be linkified on matching, for example &#x60;\&quot;#(?P&lt;id&gt;[123])\&quot;&#x60;. The second element was the URL format string that the pattern should be linkified with. A URL format string for the above example would be &#x60;\&quot;https://realm.com/my_realm_filter/%(id)s\&quot;&#x60;. And the third element was the ID of the realm filter.  | [optional] 
**RealmPlaygrounds** | Pointer to [**[]RealmPlayground**](RealmPlayground.md) | An array of dictionaries where each dictionary contains data about a single playground entry.  | [optional] 
**RealmEmoji** | Pointer to [**map[string]RealmEmoji**](RealmEmoji.md) | An object in which each key describes a realm emoji.  | [optional] 
**RealmDomain** | Pointer to [**EventEnvelopeOneOf55RealmDomain**](EventEnvelopeOneOf55RealmDomain.md) |  | [optional] 
**Domain** | Pointer to **string** | The domain to be removed.  | [optional] 
**Exports** | Pointer to [**[]RealmExport**](RealmExport.md) | An array of dictionaries where each dictionary contains details about a data export of the organization.  **Changes**: Prior to Zulip 10.0 (feature level 304), &#x60;export_type&#x60; parameter was not present as only public data export was supported via API.  | [optional] 
**Consented** | Pointer to **bool** | Whether the user has consented for their private data export.  | [optional] 
**Bot** | Pointer to [**EventEnvelopeOneOf62Bot**](EventEnvelopeOneOf62Bot.md) |  | [optional] 
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

### NewEventEnvelope

`func NewEventEnvelope(id int32, type_ string, flags []string, userId int32, messageId int32, messageIds []int32, renderingOnly bool, editTimestamp int32, ) *EventEnvelope`

NewEventEnvelope instantiates a new EventEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeWithDefaults

`func NewEventEnvelopeWithDefaults() *EventEnvelope`

NewEventEnvelopeWithDefaults instantiates a new EventEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventEnvelope) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventEnvelope) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventEnvelope) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *EventEnvelope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventEnvelope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventEnvelope) SetType(v string)`

SetType sets Type field to given value.


### GetAlertWords

`func (o *EventEnvelope) GetAlertWords() []string`

GetAlertWords returns the AlertWords field if non-nil, zero value otherwise.

### GetAlertWordsOk

`func (o *EventEnvelope) GetAlertWordsOk() (*[]string, bool)`

GetAlertWordsOk returns a tuple with the AlertWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertWords

`func (o *EventEnvelope) SetAlertWords(v []string)`

SetAlertWords sets AlertWords field to given value.

### HasAlertWords

`func (o *EventEnvelope) HasAlertWords() bool`

HasAlertWords returns a boolean if a field has been set.

### GetUser

`func (o *EventEnvelope) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventEnvelope) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventEnvelope) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *EventEnvelope) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *EventEnvelope) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *EventEnvelope) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetSettingName

`func (o *EventEnvelope) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *EventEnvelope) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *EventEnvelope) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *EventEnvelope) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetSetting

`func (o *EventEnvelope) GetSetting() EventEnvelopeOneOf1Setting`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *EventEnvelope) GetSettingOk() (*EventEnvelopeOneOf1Setting, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *EventEnvelope) SetSetting(v EventEnvelopeOneOf1Setting)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *EventEnvelope) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetLanguageName

`func (o *EventEnvelope) GetLanguageName() string`

GetLanguageName returns the LanguageName field if non-nil, zero value otherwise.

### GetLanguageNameOk

`func (o *EventEnvelope) GetLanguageNameOk() (*string, bool)`

GetLanguageNameOk returns a tuple with the LanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageName

`func (o *EventEnvelope) SetLanguageName(v string)`

SetLanguageName sets LanguageName field to given value.

### HasLanguageName

`func (o *EventEnvelope) HasLanguageName() bool`

HasLanguageName returns a boolean if a field has been set.

### GetNotificationName

`func (o *EventEnvelope) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *EventEnvelope) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *EventEnvelope) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *EventEnvelope) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetOp

`func (o *EventEnvelope) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EventEnvelope) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EventEnvelope) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *EventEnvelope) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetProperty

`func (o *EventEnvelope) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *EventEnvelope) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *EventEnvelope) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *EventEnvelope) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *EventEnvelope) GetValue() EventEnvelopeOneOf68Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventEnvelope) GetValueOk() (*EventEnvelopeOneOf68Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventEnvelope) SetValue(v EventEnvelopeOneOf68Value)`

SetValue sets Value field to given value.

### HasValue

`func (o *EventEnvelope) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPerson

`func (o *EventEnvelope) GetPerson() EventEnvelopeOneOf14Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *EventEnvelope) GetPersonOk() (*EventEnvelopeOneOf14Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *EventEnvelope) SetPerson(v EventEnvelopeOneOf14Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *EventEnvelope) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetSubscriptions

`func (o *EventEnvelope) GetSubscriptions() []EventEnvelopeOneOf6SubscriptionsInner`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *EventEnvelope) GetSubscriptionsOk() (*[]EventEnvelopeOneOf6SubscriptionsInner, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *EventEnvelope) SetSubscriptions(v []EventEnvelopeOneOf6SubscriptionsInner)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *EventEnvelope) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetStreamId

`func (o *EventEnvelope) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *EventEnvelope) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *EventEnvelope) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *EventEnvelope) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetStreamIds

`func (o *EventEnvelope) GetStreamIds() []int32`

GetStreamIds returns the StreamIds field if non-nil, zero value otherwise.

### GetStreamIdsOk

`func (o *EventEnvelope) GetStreamIdsOk() (*[]int32, bool)`

GetStreamIdsOk returns a tuple with the StreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamIds

`func (o *EventEnvelope) SetStreamIds(v []int32)`

SetStreamIds sets StreamIds field to given value.

### HasStreamIds

`func (o *EventEnvelope) HasStreamIds() bool`

HasStreamIds returns a boolean if a field has been set.

### GetUserIds

`func (o *EventEnvelope) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *EventEnvelope) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *EventEnvelope) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *EventEnvelope) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetMessage

`func (o *EventEnvelope) GetMessage() MessagesEvent`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventEnvelope) GetMessageOk() (*MessagesEvent, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventEnvelope) SetMessage(v MessagesEvent)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventEnvelope) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFlags

`func (o *EventEnvelope) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *EventEnvelope) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *EventEnvelope) SetFlags(v []string)`

SetFlags sets Flags field to given value.


### GetPresences

`func (o *EventEnvelope) GetPresences() map[string]ModernPresenceFormat`

GetPresences returns the Presences field if non-nil, zero value otherwise.

### GetPresencesOk

`func (o *EventEnvelope) GetPresencesOk() (*map[string]ModernPresenceFormat, bool)`

GetPresencesOk returns a tuple with the Presences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresences

`func (o *EventEnvelope) SetPresences(v map[string]ModernPresenceFormat)`

SetPresences sets Presences field to given value.

### HasPresences

`func (o *EventEnvelope) HasPresences() bool`

HasPresences returns a boolean if a field has been set.

### GetUserId

`func (o *EventEnvelope) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventEnvelope) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventEnvelope) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *EventEnvelope) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EventEnvelope) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EventEnvelope) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EventEnvelope) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetServerTimestamp

`func (o *EventEnvelope) GetServerTimestamp() float32`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *EventEnvelope) GetServerTimestampOk() (*float32, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *EventEnvelope) SetServerTimestamp(v float32)`

SetServerTimestamp sets ServerTimestamp field to given value.

### HasServerTimestamp

`func (o *EventEnvelope) HasServerTimestamp() bool`

HasServerTimestamp returns a boolean if a field has been set.

### GetPresence

`func (o *EventEnvelope) GetPresence() map[string]EventEnvelopeOneOf15PresenceValue`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EventEnvelope) GetPresenceOk() (*map[string]EventEnvelopeOneOf15PresenceValue, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EventEnvelope) SetPresence(v map[string]EventEnvelopeOneOf15PresenceValue)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EventEnvelope) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetStreams

`func (o *EventEnvelope) GetStreams() []EventEnvelopeOneOf17StreamsInner`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *EventEnvelope) GetStreamsOk() (*[]EventEnvelopeOneOf17StreamsInner, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *EventEnvelope) SetStreams(v []EventEnvelopeOneOf17StreamsInner)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *EventEnvelope) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetName

`func (o *EventEnvelope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventEnvelope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventEnvelope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventEnvelope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRenderedDescription

`func (o *EventEnvelope) GetRenderedDescription() string`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *EventEnvelope) GetRenderedDescriptionOk() (*string, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *EventEnvelope) SetRenderedDescription(v string)`

SetRenderedDescription sets RenderedDescription field to given value.

### HasRenderedDescription

`func (o *EventEnvelope) HasRenderedDescription() bool`

HasRenderedDescription returns a boolean if a field has been set.

### GetHistoryPublicToSubscribers

`func (o *EventEnvelope) GetHistoryPublicToSubscribers() bool`

GetHistoryPublicToSubscribers returns the HistoryPublicToSubscribers field if non-nil, zero value otherwise.

### GetHistoryPublicToSubscribersOk

`func (o *EventEnvelope) GetHistoryPublicToSubscribersOk() (*bool, bool)`

GetHistoryPublicToSubscribersOk returns a tuple with the HistoryPublicToSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPublicToSubscribers

`func (o *EventEnvelope) SetHistoryPublicToSubscribers(v bool)`

SetHistoryPublicToSubscribers sets HistoryPublicToSubscribers field to given value.

### HasHistoryPublicToSubscribers

`func (o *EventEnvelope) HasHistoryPublicToSubscribers() bool`

HasHistoryPublicToSubscribers returns a boolean if a field has been set.

### GetIsWebPublic

`func (o *EventEnvelope) GetIsWebPublic() bool`

GetIsWebPublic returns the IsWebPublic field if non-nil, zero value otherwise.

### GetIsWebPublicOk

`func (o *EventEnvelope) GetIsWebPublicOk() (*bool, bool)`

GetIsWebPublicOk returns a tuple with the IsWebPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebPublic

`func (o *EventEnvelope) SetIsWebPublic(v bool)`

SetIsWebPublic sets IsWebPublic field to given value.

### HasIsWebPublic

`func (o *EventEnvelope) HasIsWebPublic() bool`

HasIsWebPublic returns a boolean if a field has been set.

### GetEmojiName

`func (o *EventEnvelope) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *EventEnvelope) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *EventEnvelope) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *EventEnvelope) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### GetEmojiCode

`func (o *EventEnvelope) GetEmojiCode() string`

GetEmojiCode returns the EmojiCode field if non-nil, zero value otherwise.

### GetEmojiCodeOk

`func (o *EventEnvelope) GetEmojiCodeOk() (*string, bool)`

GetEmojiCodeOk returns a tuple with the EmojiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiCode

`func (o *EventEnvelope) SetEmojiCode(v string)`

SetEmojiCode sets EmojiCode field to given value.

### HasEmojiCode

`func (o *EventEnvelope) HasEmojiCode() bool`

HasEmojiCode returns a boolean if a field has been set.

### GetReactionType

`func (o *EventEnvelope) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *EventEnvelope) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *EventEnvelope) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *EventEnvelope) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### GetMessageId

`func (o *EventEnvelope) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EventEnvelope) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EventEnvelope) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetAttachment

`func (o *EventEnvelope) GetAttachment() EventEnvelopeOneOf23Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *EventEnvelope) GetAttachmentOk() (*EventEnvelopeOneOf23Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *EventEnvelope) SetAttachment(v EventEnvelopeOneOf23Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *EventEnvelope) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetUploadSpaceUsed

`func (o *EventEnvelope) GetUploadSpaceUsed() int32`

GetUploadSpaceUsed returns the UploadSpaceUsed field if non-nil, zero value otherwise.

### GetUploadSpaceUsedOk

`func (o *EventEnvelope) GetUploadSpaceUsedOk() (*int32, bool)`

GetUploadSpaceUsedOk returns a tuple with the UploadSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpaceUsed

`func (o *EventEnvelope) SetUploadSpaceUsed(v int32)`

SetUploadSpaceUsed sets UploadSpaceUsed field to given value.

### HasUploadSpaceUsed

`func (o *EventEnvelope) HasUploadSpaceUsed() bool`

HasUploadSpaceUsed returns a boolean if a field has been set.

### GetPushAccountId

`func (o *EventEnvelope) GetPushAccountId() string`

GetPushAccountId returns the PushAccountId field if non-nil, zero value otherwise.

### GetPushAccountIdOk

`func (o *EventEnvelope) GetPushAccountIdOk() (*string, bool)`

GetPushAccountIdOk returns a tuple with the PushAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushAccountId

`func (o *EventEnvelope) SetPushAccountId(v string)`

SetPushAccountId sets PushAccountId field to given value.

### HasPushAccountId

`func (o *EventEnvelope) HasPushAccountId() bool`

HasPushAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *EventEnvelope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventEnvelope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventEnvelope) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventEnvelope) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *EventEnvelope) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *EventEnvelope) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *EventEnvelope) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *EventEnvelope) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *EventEnvelope) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *EventEnvelope) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetMsgType

`func (o *EventEnvelope) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *EventEnvelope) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *EventEnvelope) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *EventEnvelope) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetContent

`func (o *EventEnvelope) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EventEnvelope) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EventEnvelope) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *EventEnvelope) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSenderId

`func (o *EventEnvelope) GetSenderId() int32`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *EventEnvelope) GetSenderIdOk() (*int32, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *EventEnvelope) SetSenderId(v int32)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *EventEnvelope) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetSubmessageId

`func (o *EventEnvelope) GetSubmessageId() int32`

GetSubmessageId returns the SubmessageId field if non-nil, zero value otherwise.

### GetSubmessageIdOk

`func (o *EventEnvelope) GetSubmessageIdOk() (*int32, bool)`

GetSubmessageIdOk returns a tuple with the SubmessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmessageId

`func (o *EventEnvelope) SetSubmessageId(v int32)`

SetSubmessageId sets SubmessageId field to given value.

### HasSubmessageId

`func (o *EventEnvelope) HasSubmessageId() bool`

HasSubmessageId returns a boolean if a field has been set.

### GetAway

`func (o *EventEnvelope) GetAway() bool`

GetAway returns the Away field if non-nil, zero value otherwise.

### GetAwayOk

`func (o *EventEnvelope) GetAwayOk() (*bool, bool)`

GetAwayOk returns a tuple with the Away field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAway

`func (o *EventEnvelope) SetAway(v bool)`

SetAway sets Away field to given value.

### HasAway

`func (o *EventEnvelope) HasAway() bool`

HasAway returns a boolean if a field has been set.

### GetStatusText

`func (o *EventEnvelope) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *EventEnvelope) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *EventEnvelope) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *EventEnvelope) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetFields

`func (o *EventEnvelope) GetFields() []CustomProfileField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EventEnvelope) GetFieldsOk() (*[]CustomProfileField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EventEnvelope) SetFields(v []CustomProfileField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EventEnvelope) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetDefaultStreamGroups

`func (o *EventEnvelope) GetDefaultStreamGroups() []DefaultChannelGroup`

GetDefaultStreamGroups returns the DefaultStreamGroups field if non-nil, zero value otherwise.

### GetDefaultStreamGroupsOk

`func (o *EventEnvelope) GetDefaultStreamGroupsOk() (*[]DefaultChannelGroup, bool)`

GetDefaultStreamGroupsOk returns a tuple with the DefaultStreamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStreamGroups

`func (o *EventEnvelope) SetDefaultStreamGroups(v []DefaultChannelGroup)`

SetDefaultStreamGroups sets DefaultStreamGroups field to given value.

### HasDefaultStreamGroups

`func (o *EventEnvelope) HasDefaultStreamGroups() bool`

HasDefaultStreamGroups returns a boolean if a field has been set.

### GetDefaultStreams

`func (o *EventEnvelope) GetDefaultStreams() []int32`

GetDefaultStreams returns the DefaultStreams field if non-nil, zero value otherwise.

### GetDefaultStreamsOk

`func (o *EventEnvelope) GetDefaultStreamsOk() (*[]int32, bool)`

GetDefaultStreamsOk returns a tuple with the DefaultStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStreams

`func (o *EventEnvelope) SetDefaultStreams(v []int32)`

SetDefaultStreams sets DefaultStreams field to given value.

### HasDefaultStreams

`func (o *EventEnvelope) HasDefaultStreams() bool`

HasDefaultStreams returns a boolean if a field has been set.

### GetMessageIds

`func (o *EventEnvelope) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *EventEnvelope) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *EventEnvelope) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.


### GetMessageType

`func (o *EventEnvelope) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *EventEnvelope) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *EventEnvelope) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *EventEnvelope) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTopic

`func (o *EventEnvelope) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *EventEnvelope) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *EventEnvelope) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *EventEnvelope) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetMutedTopics

`func (o *EventEnvelope) GetMutedTopics() [][]EventEnvelopeOneOf31MutedTopicsInnerInner`

GetMutedTopics returns the MutedTopics field if non-nil, zero value otherwise.

### GetMutedTopicsOk

`func (o *EventEnvelope) GetMutedTopicsOk() (*[][]EventEnvelopeOneOf31MutedTopicsInnerInner, bool)`

GetMutedTopicsOk returns a tuple with the MutedTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedTopics

`func (o *EventEnvelope) SetMutedTopics(v [][]EventEnvelopeOneOf31MutedTopicsInnerInner)`

SetMutedTopics sets MutedTopics field to given value.

### HasMutedTopics

`func (o *EventEnvelope) HasMutedTopics() bool`

HasMutedTopics returns a boolean if a field has been set.

### GetTopicName

`func (o *EventEnvelope) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *EventEnvelope) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *EventEnvelope) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *EventEnvelope) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EventEnvelope) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EventEnvelope) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EventEnvelope) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EventEnvelope) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVisibilityPolicy

`func (o *EventEnvelope) GetVisibilityPolicy() int32`

GetVisibilityPolicy returns the VisibilityPolicy field if non-nil, zero value otherwise.

### GetVisibilityPolicyOk

`func (o *EventEnvelope) GetVisibilityPolicyOk() (*int32, bool)`

GetVisibilityPolicyOk returns a tuple with the VisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityPolicy

`func (o *EventEnvelope) SetVisibilityPolicy(v int32)`

SetVisibilityPolicy sets VisibilityPolicy field to given value.

### HasVisibilityPolicy

`func (o *EventEnvelope) HasVisibilityPolicy() bool`

HasVisibilityPolicy returns a boolean if a field has been set.

### GetMutedUsers

`func (o *EventEnvelope) GetMutedUsers() []EventEnvelopeOneOf33MutedUsersInner`

GetMutedUsers returns the MutedUsers field if non-nil, zero value otherwise.

### GetMutedUsersOk

`func (o *EventEnvelope) GetMutedUsersOk() (*[]EventEnvelopeOneOf33MutedUsersInner, bool)`

GetMutedUsersOk returns a tuple with the MutedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedUsers

`func (o *EventEnvelope) SetMutedUsers(v []EventEnvelopeOneOf33MutedUsersInner)`

SetMutedUsers sets MutedUsers field to given value.

### HasMutedUsers

`func (o *EventEnvelope) HasMutedUsers() bool`

HasMutedUsers returns a boolean if a field has been set.

### GetOnboardingSteps

`func (o *EventEnvelope) GetOnboardingSteps() []OnboardingStep`

GetOnboardingSteps returns the OnboardingSteps field if non-nil, zero value otherwise.

### GetOnboardingStepsOk

`func (o *EventEnvelope) GetOnboardingStepsOk() (*[]OnboardingStep, bool)`

GetOnboardingStepsOk returns a tuple with the OnboardingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingSteps

`func (o *EventEnvelope) SetOnboardingSteps(v []OnboardingStep)`

SetOnboardingSteps sets OnboardingSteps field to given value.

### HasOnboardingSteps

`func (o *EventEnvelope) HasOnboardingSteps() bool`

HasOnboardingSteps returns a boolean if a field has been set.

### GetRenderingOnly

`func (o *EventEnvelope) GetRenderingOnly() bool`

GetRenderingOnly returns the RenderingOnly field if non-nil, zero value otherwise.

### GetRenderingOnlyOk

`func (o *EventEnvelope) GetRenderingOnlyOk() (*bool, bool)`

GetRenderingOnlyOk returns a tuple with the RenderingOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderingOnly

`func (o *EventEnvelope) SetRenderingOnly(v bool)`

SetRenderingOnly sets RenderingOnly field to given value.


### GetEditTimestamp

`func (o *EventEnvelope) GetEditTimestamp() int32`

GetEditTimestamp returns the EditTimestamp field if non-nil, zero value otherwise.

### GetEditTimestampOk

`func (o *EventEnvelope) GetEditTimestampOk() (*int32, bool)`

GetEditTimestampOk returns a tuple with the EditTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditTimestamp

`func (o *EventEnvelope) SetEditTimestamp(v int32)`

SetEditTimestamp sets EditTimestamp field to given value.


### GetStreamName

`func (o *EventEnvelope) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *EventEnvelope) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *EventEnvelope) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *EventEnvelope) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### GetNewStreamId

`func (o *EventEnvelope) GetNewStreamId() int32`

GetNewStreamId returns the NewStreamId field if non-nil, zero value otherwise.

### GetNewStreamIdOk

`func (o *EventEnvelope) GetNewStreamIdOk() (*int32, bool)`

GetNewStreamIdOk returns a tuple with the NewStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStreamId

`func (o *EventEnvelope) SetNewStreamId(v int32)`

SetNewStreamId sets NewStreamId field to given value.

### HasNewStreamId

`func (o *EventEnvelope) HasNewStreamId() bool`

HasNewStreamId returns a boolean if a field has been set.

### GetPropagateMode

`func (o *EventEnvelope) GetPropagateMode() string`

GetPropagateMode returns the PropagateMode field if non-nil, zero value otherwise.

### GetPropagateModeOk

`func (o *EventEnvelope) GetPropagateModeOk() (*string, bool)`

GetPropagateModeOk returns a tuple with the PropagateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateMode

`func (o *EventEnvelope) SetPropagateMode(v string)`

SetPropagateMode sets PropagateMode field to given value.

### HasPropagateMode

`func (o *EventEnvelope) HasPropagateMode() bool`

HasPropagateMode returns a boolean if a field has been set.

### GetOrigSubject

`func (o *EventEnvelope) GetOrigSubject() string`

GetOrigSubject returns the OrigSubject field if non-nil, zero value otherwise.

### GetOrigSubjectOk

`func (o *EventEnvelope) GetOrigSubjectOk() (*string, bool)`

GetOrigSubjectOk returns a tuple with the OrigSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigSubject

`func (o *EventEnvelope) SetOrigSubject(v string)`

SetOrigSubject sets OrigSubject field to given value.

### HasOrigSubject

`func (o *EventEnvelope) HasOrigSubject() bool`

HasOrigSubject returns a boolean if a field has been set.

### GetSubject

`func (o *EventEnvelope) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EventEnvelope) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EventEnvelope) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EventEnvelope) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTopicLinks

`func (o *EventEnvelope) GetTopicLinks() []EventEnvelopeOneOf36TopicLinksInner`

GetTopicLinks returns the TopicLinks field if non-nil, zero value otherwise.

### GetTopicLinksOk

`func (o *EventEnvelope) GetTopicLinksOk() (*[]EventEnvelopeOneOf36TopicLinksInner, bool)`

GetTopicLinksOk returns a tuple with the TopicLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicLinks

`func (o *EventEnvelope) SetTopicLinks(v []EventEnvelopeOneOf36TopicLinksInner)`

SetTopicLinks sets TopicLinks field to given value.

### HasTopicLinks

`func (o *EventEnvelope) HasTopicLinks() bool`

HasTopicLinks returns a boolean if a field has been set.

### GetOrigContent

`func (o *EventEnvelope) GetOrigContent() string`

GetOrigContent returns the OrigContent field if non-nil, zero value otherwise.

### GetOrigContentOk

`func (o *EventEnvelope) GetOrigContentOk() (*string, bool)`

GetOrigContentOk returns a tuple with the OrigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigContent

`func (o *EventEnvelope) SetOrigContent(v string)`

SetOrigContent sets OrigContent field to given value.

### HasOrigContent

`func (o *EventEnvelope) HasOrigContent() bool`

HasOrigContent returns a boolean if a field has been set.

### GetOrigRenderedContent

`func (o *EventEnvelope) GetOrigRenderedContent() string`

GetOrigRenderedContent returns the OrigRenderedContent field if non-nil, zero value otherwise.

### GetOrigRenderedContentOk

`func (o *EventEnvelope) GetOrigRenderedContentOk() (*string, bool)`

GetOrigRenderedContentOk returns a tuple with the OrigRenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigRenderedContent

`func (o *EventEnvelope) SetOrigRenderedContent(v string)`

SetOrigRenderedContent sets OrigRenderedContent field to given value.

### HasOrigRenderedContent

`func (o *EventEnvelope) HasOrigRenderedContent() bool`

HasOrigRenderedContent returns a boolean if a field has been set.

### GetRenderedContent

`func (o *EventEnvelope) GetRenderedContent() string`

GetRenderedContent returns the RenderedContent field if non-nil, zero value otherwise.

### GetRenderedContentOk

`func (o *EventEnvelope) GetRenderedContentOk() (*string, bool)`

GetRenderedContentOk returns a tuple with the RenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedContent

`func (o *EventEnvelope) SetRenderedContent(v string)`

SetRenderedContent sets RenderedContent field to given value.

### HasRenderedContent

`func (o *EventEnvelope) HasRenderedContent() bool`

HasRenderedContent returns a boolean if a field has been set.

### GetIsMeMessage

`func (o *EventEnvelope) GetIsMeMessage() bool`

GetIsMeMessage returns the IsMeMessage field if non-nil, zero value otherwise.

### GetIsMeMessageOk

`func (o *EventEnvelope) GetIsMeMessageOk() (*bool, bool)`

GetIsMeMessageOk returns a tuple with the IsMeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeMessage

`func (o *EventEnvelope) SetIsMeMessage(v bool)`

SetIsMeMessage sets IsMeMessage field to given value.

### HasIsMeMessage

`func (o *EventEnvelope) HasIsMeMessage() bool`

HasIsMeMessage returns a boolean if a field has been set.

### GetSender

`func (o *EventEnvelope) GetSender() EventEnvelopeOneOf38Sender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *EventEnvelope) GetSenderOk() (*EventEnvelopeOneOf38Sender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *EventEnvelope) SetSender(v EventEnvelopeOneOf38Sender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *EventEnvelope) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipients

`func (o *EventEnvelope) GetRecipients() []EventEnvelopeOneOf38RecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EventEnvelope) GetRecipientsOk() (*[]EventEnvelopeOneOf38RecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EventEnvelope) SetRecipients(v []EventEnvelopeOneOf38RecipientsInner)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *EventEnvelope) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetRecipient

`func (o *EventEnvelope) GetRecipient() EventEnvelopeOneOf39Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *EventEnvelope) GetRecipientOk() (*EventEnvelopeOneOf39Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *EventEnvelope) SetRecipient(v EventEnvelopeOneOf39Recipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *EventEnvelope) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetOperation

`func (o *EventEnvelope) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *EventEnvelope) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *EventEnvelope) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *EventEnvelope) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetFlag

`func (o *EventEnvelope) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *EventEnvelope) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *EventEnvelope) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *EventEnvelope) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetMessages

`func (o *EventEnvelope) GetMessages() []int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *EventEnvelope) GetMessagesOk() (*[]int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *EventEnvelope) SetMessages(v []int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *EventEnvelope) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetAll

`func (o *EventEnvelope) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *EventEnvelope) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *EventEnvelope) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *EventEnvelope) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetMessageDetails

`func (o *EventEnvelope) GetMessageDetails() map[string]EventEnvelopeOneOf42MessageDetailsValue`

GetMessageDetails returns the MessageDetails field if non-nil, zero value otherwise.

### GetMessageDetailsOk

`func (o *EventEnvelope) GetMessageDetailsOk() (*map[string]EventEnvelopeOneOf42MessageDetailsValue, bool)`

GetMessageDetailsOk returns a tuple with the MessageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetails

`func (o *EventEnvelope) SetMessageDetails(v map[string]EventEnvelopeOneOf42MessageDetailsValue)`

SetMessageDetails sets MessageDetails field to given value.

### HasMessageDetails

`func (o *EventEnvelope) HasMessageDetails() bool`

HasMessageDetails returns a boolean if a field has been set.

### GetGroup

`func (o *EventEnvelope) GetGroup() UserGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *EventEnvelope) GetGroupOk() (*UserGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *EventEnvelope) SetGroup(v UserGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *EventEnvelope) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupId

`func (o *EventEnvelope) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *EventEnvelope) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *EventEnvelope) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *EventEnvelope) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetData

`func (o *EventEnvelope) GetData() EventEnvelopeOneOf84Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventEnvelope) GetDataOk() (*EventEnvelopeOneOf84Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventEnvelope) SetData(v EventEnvelopeOneOf84Data)`

SetData sets Data field to given value.

### HasData

`func (o *EventEnvelope) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDirectSubgroupIds

`func (o *EventEnvelope) GetDirectSubgroupIds() []int32`

GetDirectSubgroupIds returns the DirectSubgroupIds field if non-nil, zero value otherwise.

### GetDirectSubgroupIdsOk

`func (o *EventEnvelope) GetDirectSubgroupIdsOk() (*[]int32, bool)`

GetDirectSubgroupIdsOk returns a tuple with the DirectSubgroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSubgroupIds

`func (o *EventEnvelope) SetDirectSubgroupIds(v []int32)`

SetDirectSubgroupIds sets DirectSubgroupIds field to given value.

### HasDirectSubgroupIds

`func (o *EventEnvelope) HasDirectSubgroupIds() bool`

HasDirectSubgroupIds returns a boolean if a field has been set.

### GetRealmLinkifiers

`func (o *EventEnvelope) GetRealmLinkifiers() []EventEnvelopeOneOf50RealmLinkifiersInner`

GetRealmLinkifiers returns the RealmLinkifiers field if non-nil, zero value otherwise.

### GetRealmLinkifiersOk

`func (o *EventEnvelope) GetRealmLinkifiersOk() (*[]EventEnvelopeOneOf50RealmLinkifiersInner, bool)`

GetRealmLinkifiersOk returns a tuple with the RealmLinkifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmLinkifiers

`func (o *EventEnvelope) SetRealmLinkifiers(v []EventEnvelopeOneOf50RealmLinkifiersInner)`

SetRealmLinkifiers sets RealmLinkifiers field to given value.

### HasRealmLinkifiers

`func (o *EventEnvelope) HasRealmLinkifiers() bool`

HasRealmLinkifiers returns a boolean if a field has been set.

### GetRealmFilters

`func (o *EventEnvelope) GetRealmFilters() [][]EventEnvelopeOneOf51RealmFiltersInnerInner`

GetRealmFilters returns the RealmFilters field if non-nil, zero value otherwise.

### GetRealmFiltersOk

`func (o *EventEnvelope) GetRealmFiltersOk() (*[][]EventEnvelopeOneOf51RealmFiltersInnerInner, bool)`

GetRealmFiltersOk returns a tuple with the RealmFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmFilters

`func (o *EventEnvelope) SetRealmFilters(v [][]EventEnvelopeOneOf51RealmFiltersInnerInner)`

SetRealmFilters sets RealmFilters field to given value.

### HasRealmFilters

`func (o *EventEnvelope) HasRealmFilters() bool`

HasRealmFilters returns a boolean if a field has been set.

### GetRealmPlaygrounds

`func (o *EventEnvelope) GetRealmPlaygrounds() []RealmPlayground`

GetRealmPlaygrounds returns the RealmPlaygrounds field if non-nil, zero value otherwise.

### GetRealmPlaygroundsOk

`func (o *EventEnvelope) GetRealmPlaygroundsOk() (*[]RealmPlayground, bool)`

GetRealmPlaygroundsOk returns a tuple with the RealmPlaygrounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPlaygrounds

`func (o *EventEnvelope) SetRealmPlaygrounds(v []RealmPlayground)`

SetRealmPlaygrounds sets RealmPlaygrounds field to given value.

### HasRealmPlaygrounds

`func (o *EventEnvelope) HasRealmPlaygrounds() bool`

HasRealmPlaygrounds returns a boolean if a field has been set.

### GetRealmEmoji

`func (o *EventEnvelope) GetRealmEmoji() map[string]RealmEmoji`

GetRealmEmoji returns the RealmEmoji field if non-nil, zero value otherwise.

### GetRealmEmojiOk

`func (o *EventEnvelope) GetRealmEmojiOk() (*map[string]RealmEmoji, bool)`

GetRealmEmojiOk returns a tuple with the RealmEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmoji

`func (o *EventEnvelope) SetRealmEmoji(v map[string]RealmEmoji)`

SetRealmEmoji sets RealmEmoji field to given value.

### HasRealmEmoji

`func (o *EventEnvelope) HasRealmEmoji() bool`

HasRealmEmoji returns a boolean if a field has been set.

### GetRealmDomain

`func (o *EventEnvelope) GetRealmDomain() EventEnvelopeOneOf55RealmDomain`

GetRealmDomain returns the RealmDomain field if non-nil, zero value otherwise.

### GetRealmDomainOk

`func (o *EventEnvelope) GetRealmDomainOk() (*EventEnvelopeOneOf55RealmDomain, bool)`

GetRealmDomainOk returns a tuple with the RealmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDomain

`func (o *EventEnvelope) SetRealmDomain(v EventEnvelopeOneOf55RealmDomain)`

SetRealmDomain sets RealmDomain field to given value.

### HasRealmDomain

`func (o *EventEnvelope) HasRealmDomain() bool`

HasRealmDomain returns a boolean if a field has been set.

### GetDomain

`func (o *EventEnvelope) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EventEnvelope) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EventEnvelope) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EventEnvelope) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExports

`func (o *EventEnvelope) GetExports() []RealmExport`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *EventEnvelope) GetExportsOk() (*[]RealmExport, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *EventEnvelope) SetExports(v []RealmExport)`

SetExports sets Exports field to given value.

### HasExports

`func (o *EventEnvelope) HasExports() bool`

HasExports returns a boolean if a field has been set.

### GetConsented

`func (o *EventEnvelope) GetConsented() bool`

GetConsented returns the Consented field if non-nil, zero value otherwise.

### GetConsentedOk

`func (o *EventEnvelope) GetConsentedOk() (*bool, bool)`

GetConsentedOk returns a tuple with the Consented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsented

`func (o *EventEnvelope) SetConsented(v bool)`

SetConsented sets Consented field to given value.

### HasConsented

`func (o *EventEnvelope) HasConsented() bool`

HasConsented returns a boolean if a field has been set.

### GetBot

`func (o *EventEnvelope) GetBot() EventEnvelopeOneOf62Bot`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *EventEnvelope) GetBotOk() (*EventEnvelopeOneOf62Bot, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *EventEnvelope) SetBot(v EventEnvelopeOneOf62Bot)`

SetBot sets Bot field to given value.

### HasBot

`func (o *EventEnvelope) HasBot() bool`

HasBot returns a boolean if a field has been set.

### GetRealmId

`func (o *EventEnvelope) GetRealmId() int32`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *EventEnvelope) GetRealmIdOk() (*int32, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *EventEnvelope) SetRealmId(v int32)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *EventEnvelope) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetZulipVersion

`func (o *EventEnvelope) GetZulipVersion() string`

GetZulipVersion returns the ZulipVersion field if non-nil, zero value otherwise.

### GetZulipVersionOk

`func (o *EventEnvelope) GetZulipVersionOk() (*string, bool)`

GetZulipVersionOk returns a tuple with the ZulipVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipVersion

`func (o *EventEnvelope) SetZulipVersion(v string)`

SetZulipVersion sets ZulipVersion field to given value.

### HasZulipVersion

`func (o *EventEnvelope) HasZulipVersion() bool`

HasZulipVersion returns a boolean if a field has been set.

### GetZulipMergeBase

`func (o *EventEnvelope) GetZulipMergeBase() string`

GetZulipMergeBase returns the ZulipMergeBase field if non-nil, zero value otherwise.

### GetZulipMergeBaseOk

`func (o *EventEnvelope) GetZulipMergeBaseOk() (*string, bool)`

GetZulipMergeBaseOk returns a tuple with the ZulipMergeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipMergeBase

`func (o *EventEnvelope) SetZulipMergeBase(v string)`

SetZulipMergeBase sets ZulipMergeBase field to given value.

### HasZulipMergeBase

`func (o *EventEnvelope) HasZulipMergeBase() bool`

HasZulipMergeBase returns a boolean if a field has been set.

### GetZulipFeatureLevel

`func (o *EventEnvelope) GetZulipFeatureLevel() int32`

GetZulipFeatureLevel returns the ZulipFeatureLevel field if non-nil, zero value otherwise.

### GetZulipFeatureLevelOk

`func (o *EventEnvelope) GetZulipFeatureLevelOk() (*int32, bool)`

GetZulipFeatureLevelOk returns a tuple with the ZulipFeatureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipFeatureLevel

`func (o *EventEnvelope) SetZulipFeatureLevel(v int32)`

SetZulipFeatureLevel sets ZulipFeatureLevel field to given value.

### HasZulipFeatureLevel

`func (o *EventEnvelope) HasZulipFeatureLevel() bool`

HasZulipFeatureLevel returns a boolean if a field has been set.

### GetServerGeneration

`func (o *EventEnvelope) GetServerGeneration() int32`

GetServerGeneration returns the ServerGeneration field if non-nil, zero value otherwise.

### GetServerGenerationOk

`func (o *EventEnvelope) GetServerGenerationOk() (*int32, bool)`

GetServerGenerationOk returns a tuple with the ServerGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGeneration

`func (o *EventEnvelope) SetServerGeneration(v int32)`

SetServerGeneration sets ServerGeneration field to given value.

### HasServerGeneration

`func (o *EventEnvelope) HasServerGeneration() bool`

HasServerGeneration returns a boolean if a field has been set.

### GetImmediate

`func (o *EventEnvelope) GetImmediate() bool`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *EventEnvelope) GetImmediateOk() (*bool, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *EventEnvelope) SetImmediate(v bool)`

SetImmediate sets Immediate field to given value.

### HasImmediate

`func (o *EventEnvelope) HasImmediate() bool`

HasImmediate returns a boolean if a field has been set.

### GetDrafts

`func (o *EventEnvelope) GetDrafts() []Draft`

GetDrafts returns the Drafts field if non-nil, zero value otherwise.

### GetDraftsOk

`func (o *EventEnvelope) GetDraftsOk() (*[]Draft, bool)`

GetDraftsOk returns a tuple with the Drafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrafts

`func (o *EventEnvelope) SetDrafts(v []Draft)`

SetDrafts sets Drafts field to given value.

### HasDrafts

`func (o *EventEnvelope) HasDrafts() bool`

HasDrafts returns a boolean if a field has been set.

### GetDraft

`func (o *EventEnvelope) GetDraft() Draft`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *EventEnvelope) GetDraftOk() (*Draft, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *EventEnvelope) SetDraft(v Draft)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *EventEnvelope) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetDraftId

`func (o *EventEnvelope) GetDraftId() int32`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *EventEnvelope) GetDraftIdOk() (*int32, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *EventEnvelope) SetDraftId(v int32)`

SetDraftId sets DraftId field to given value.

### HasDraftId

`func (o *EventEnvelope) HasDraftId() bool`

HasDraftId returns a boolean if a field has been set.

### GetNavigationView

`func (o *EventEnvelope) GetNavigationView() NavigationView`

GetNavigationView returns the NavigationView field if non-nil, zero value otherwise.

### GetNavigationViewOk

`func (o *EventEnvelope) GetNavigationViewOk() (*NavigationView, bool)`

GetNavigationViewOk returns a tuple with the NavigationView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationView

`func (o *EventEnvelope) SetNavigationView(v NavigationView)`

SetNavigationView sets NavigationView field to given value.

### HasNavigationView

`func (o *EventEnvelope) HasNavigationView() bool`

HasNavigationView returns a boolean if a field has been set.

### GetFragment

`func (o *EventEnvelope) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *EventEnvelope) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *EventEnvelope) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *EventEnvelope) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetSavedSnippet

`func (o *EventEnvelope) GetSavedSnippet() SavedSnippet`

GetSavedSnippet returns the SavedSnippet field if non-nil, zero value otherwise.

### GetSavedSnippetOk

`func (o *EventEnvelope) GetSavedSnippetOk() (*SavedSnippet, bool)`

GetSavedSnippetOk returns a tuple with the SavedSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippet

`func (o *EventEnvelope) SetSavedSnippet(v SavedSnippet)`

SetSavedSnippet sets SavedSnippet field to given value.

### HasSavedSnippet

`func (o *EventEnvelope) HasSavedSnippet() bool`

HasSavedSnippet returns a boolean if a field has been set.

### GetSavedSnippetId

`func (o *EventEnvelope) GetSavedSnippetId() int32`

GetSavedSnippetId returns the SavedSnippetId field if non-nil, zero value otherwise.

### GetSavedSnippetIdOk

`func (o *EventEnvelope) GetSavedSnippetIdOk() (*int32, bool)`

GetSavedSnippetIdOk returns a tuple with the SavedSnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippetId

`func (o *EventEnvelope) SetSavedSnippetId(v int32)`

SetSavedSnippetId sets SavedSnippetId field to given value.

### HasSavedSnippetId

`func (o *EventEnvelope) HasSavedSnippetId() bool`

HasSavedSnippetId returns a boolean if a field has been set.

### GetReminders

`func (o *EventEnvelope) GetReminders() []Reminder`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *EventEnvelope) GetRemindersOk() (*[]Reminder, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *EventEnvelope) SetReminders(v []Reminder)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *EventEnvelope) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetReminderId

`func (o *EventEnvelope) GetReminderId() int32`

GetReminderId returns the ReminderId field if non-nil, zero value otherwise.

### GetReminderIdOk

`func (o *EventEnvelope) GetReminderIdOk() (*int32, bool)`

GetReminderIdOk returns a tuple with the ReminderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderId

`func (o *EventEnvelope) SetReminderId(v int32)`

SetReminderId sets ReminderId field to given value.

### HasReminderId

`func (o *EventEnvelope) HasReminderId() bool`

HasReminderId returns a boolean if a field has been set.

### GetScheduledMessages

`func (o *EventEnvelope) GetScheduledMessages() []ScheduledMessage`

GetScheduledMessages returns the ScheduledMessages field if non-nil, zero value otherwise.

### GetScheduledMessagesOk

`func (o *EventEnvelope) GetScheduledMessagesOk() (*[]ScheduledMessage, bool)`

GetScheduledMessagesOk returns a tuple with the ScheduledMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessages

`func (o *EventEnvelope) SetScheduledMessages(v []ScheduledMessage)`

SetScheduledMessages sets ScheduledMessages field to given value.

### HasScheduledMessages

`func (o *EventEnvelope) HasScheduledMessages() bool`

HasScheduledMessages returns a boolean if a field has been set.

### GetScheduledMessage

`func (o *EventEnvelope) GetScheduledMessage() ScheduledMessage`

GetScheduledMessage returns the ScheduledMessage field if non-nil, zero value otherwise.

### GetScheduledMessageOk

`func (o *EventEnvelope) GetScheduledMessageOk() (*ScheduledMessage, bool)`

GetScheduledMessageOk returns a tuple with the ScheduledMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessage

`func (o *EventEnvelope) SetScheduledMessage(v ScheduledMessage)`

SetScheduledMessage sets ScheduledMessage field to given value.

### HasScheduledMessage

`func (o *EventEnvelope) HasScheduledMessage() bool`

HasScheduledMessage returns a boolean if a field has been set.

### GetScheduledMessageId

`func (o *EventEnvelope) GetScheduledMessageId() int32`

GetScheduledMessageId returns the ScheduledMessageId field if non-nil, zero value otherwise.

### GetScheduledMessageIdOk

`func (o *EventEnvelope) GetScheduledMessageIdOk() (*int32, bool)`

GetScheduledMessageIdOk returns a tuple with the ScheduledMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessageId

`func (o *EventEnvelope) SetScheduledMessageId(v int32)`

SetScheduledMessageId sets ScheduledMessageId field to given value.

### HasScheduledMessageId

`func (o *EventEnvelope) HasScheduledMessageId() bool`

HasScheduledMessageId returns a boolean if a field has been set.

### GetChannelFolder

`func (o *EventEnvelope) GetChannelFolder() ChannelFolder`

GetChannelFolder returns the ChannelFolder field if non-nil, zero value otherwise.

### GetChannelFolderOk

`func (o *EventEnvelope) GetChannelFolderOk() (*ChannelFolder, bool)`

GetChannelFolderOk returns a tuple with the ChannelFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFolder

`func (o *EventEnvelope) SetChannelFolder(v ChannelFolder)`

SetChannelFolder sets ChannelFolder field to given value.

### HasChannelFolder

`func (o *EventEnvelope) HasChannelFolder() bool`

HasChannelFolder returns a boolean if a field has been set.

### GetChannelFolderId

`func (o *EventEnvelope) GetChannelFolderId() float32`

GetChannelFolderId returns the ChannelFolderId field if non-nil, zero value otherwise.

### GetChannelFolderIdOk

`func (o *EventEnvelope) GetChannelFolderIdOk() (*float32, bool)`

GetChannelFolderIdOk returns a tuple with the ChannelFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFolderId

`func (o *EventEnvelope) SetChannelFolderId(v float32)`

SetChannelFolderId sets ChannelFolderId field to given value.

### HasChannelFolderId

`func (o *EventEnvelope) HasChannelFolderId() bool`

HasChannelFolderId returns a boolean if a field has been set.

### GetOrder

`func (o *EventEnvelope) GetOrder() []int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *EventEnvelope) GetOrderOk() (*[]int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *EventEnvelope) SetOrder(v []int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *EventEnvelope) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


