# RegisterQueue200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**QueueId** | Pointer to **NullableString** | The ID of the queue that has been allocated for your client.  Will be &#x60;null&#x60; only for unauthenticated access in realms that have enabled the [public access option](/help/public-access-option).  | [optional] 
**LastEventId** | Pointer to **int32** | The initial value of &#x60;last_event_id&#x60; to pass to &#x60;GET /api/v1/events&#x60;.  | [optional] 
**ZulipFeatureLevel** | Pointer to **int32** | The server&#39;s current [Zulip feature level](/api/changelog).  **Changes**: As of Zulip 3.0 (feature level 3), this is always present in the endpoint&#39;s response. Previously, it was only present if &#x60;event_types&#x60; included &#x60;zulip_version&#x60;.  New in Zulip 3.0 (feature level 1).  | [optional] 
**ZulipVersion** | Pointer to **string** | The server&#39;s version number. This is often a release version number, like &#x60;2.1.7&#x60;. But for a server running a [version from Git][git-release], it will be a Git reference to the commit, like &#x60;5.0-dev-1650-gc3fd37755f&#x60;.  **Changes**: As of Zulip 3.0 (feature level 3), this is always present in the endpoint&#39;s response. Previously, it was only present if &#x60;event_types&#x60; included &#x60;zulip_version&#x60;.  [git-release]: https://zulip.readthedocs.io/en/latest/overview/release-lifecycle.html#git-versions  | [optional] 
**ZulipMergeBase** | Pointer to **string** | The &#x60;git merge-base&#x60; between &#x60;zulip_version&#x60; and official branches in the public [Zulip server and web app repository](https://github.com/zulip/zulip), in the same format as &#x60;zulip_version&#x60;. This will equal &#x60;zulip_version&#x60; if the server is not running a fork of the Zulip server.  This will be &#x60;\&quot;\&quot;&#x60; if the server does not know its &#x60;merge-base&#x60;.  **Changes**: New in Zulip 5.0 (feature level 88).  | [optional] 
**AlertWords** | Pointer to **[]string** | Present if &#x60;alert_words&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of strings, each an [alert word](/help/dm-mention-alert-notifications#alert-words) that the current user has configured.  | [optional] 
**CustomProfileFields** | Pointer to [**[]CustomProfileField**](CustomProfileField.md) | Present if &#x60;custom_profile_fields&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary contains the details of a single custom profile field that is available to users in this Zulip organization. This must be combined with the custom profile field values on individual user objects to display users&#39; profiles.  | [optional] 
**CustomProfileFieldTypes** | Pointer to [**map[string]RegisterQueueResponseCustomProfileFieldTypesEntry**](RegisterQueueResponseCustomProfileFieldTypesEntry.md) | Present if &#x60;custom_profile_fields&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of objects; each object describes a type of custom profile field that could be configured on this Zulip server. Each custom profile type has an ID and the &#x60;type&#x60; property of a custom profile field is equal to one of these IDs.  This attribute is only useful for clients containing UI for changing the set of configured custom profile fields in a Zulip organization.  | [optional] 
**RealmDateCreated** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The UNIX timestamp (UTC) for when the organization was created.  **Changes**: New in Zulip 8.0 (feature level 203).  | [optional] 
**DemoOrganizationScheduledDeletionDate** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;, and the realm is a demo organization.  The UNIX timestamp (UTC) when the demo organization will be automatically deleted. Clients should use this to display a prominent warning to the user that the organization will be deleted at the indicated time.  **Changes**: New in Zulip 5.0 (feature level 94).  | [optional] 
**Drafts** | Pointer to [**[]Draft**](Draft.md) | An array containing draft objects for the user. These drafts are being stored on the backend for the purpose of syncing across devices. This array will be empty if &#x60;enable_drafts_synchronization&#x60; is set to &#x60;false&#x60;.  | [optional] 
**OnboardingSteps** | Pointer to [**[]OnboardingStep**](OnboardingStep.md) | Present if &#x60;onboarding_steps&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries, where each dictionary contains details about a single onboarding step that should be shown to the user.  We expect that only official Zulip clients will interact with this data.  **Changes**: Before Zulip 8.0 (feature level 233), this array was named &#x60;hotspots&#x60;. Prior to this feature level, one-time notice onboarding steps were not supported, and the &#x60;type&#x60; field in these objects did not exist as all onboarding steps were implicitly hotspots.  | [optional] 
**NavigationTourVideoUrl** | Pointer to **NullableString** | Present if &#x60;onboarding_steps&#x60; is present in &#x60;fetch_event_types&#x60;.  URL of the navigation tour video to display to new users during onboarding. If &#x60;null&#x60;, the onboarding video experience is disabled.  **Changes**: New in Zulip 10.0 (feature level 369).  | [optional] 
**MaxMessageId** | Pointer to **int32** | Present if &#x60;message&#x60; is present in &#x60;fetch_event_types&#x60;.  The highest message ID among all messages the user has received as of the moment of this request.  **Deprecated**: This field may be removed in future versions as it no longer has a clear purpose. Clients wishing to fetch the latest messages should pass &#x60;\&quot;anchor\&quot;: \&quot;latest\&quot;&#x60; to &#x60;GET /messages&#x60;.  | [optional] 
**MaxReminderNoteLength** | Pointer to **int32** | The maximum allowed length for a reminder note.  **Changes**: New in Zulip 11.0 (feature level 415).  | [optional] 
**MaxStreamNameLength** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum allowed length for a channel name, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 4.0 (feature level 53). Previously, this required &#x60;stream&#x60; in &#x60;fetch_event_types&#x60;, was called &#x60;stream_name_max_length&#x60;, and always had a value of 60.  | [optional] 
**MaxStreamDescriptionLength** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum allowed length for a channel description, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 4.0 (feature level 53). Previously, this required &#x60;stream&#x60; in &#x60;fetch_event_types&#x60;, was called &#x60;stream_description_max_length&#x60;, and always had a value of 1024.  | [optional] 
**MaxChannelFolderNameLength** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum allowed length for a channel folder name, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 11.0 (feature level 410). Clients should use 60 as a fallback value on previous feature levels.  | [optional] 
**MaxChannelFolderDescriptionLength** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum allowed length for a channel folder description, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 11.0 (feature level 410). Clients should use 1024 as a fallback value on previous feature levels.  | [optional] 
**MaxTopicLength** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum allowed length for a topic, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 4.0 (feature level 53). Previously, this property always had a value of 60.  | [optional] 
**MaxMessageLength** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum allowed length for a message, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 4.0 (feature level 53). Previously, this property always had a value of 10000.  | [optional] 
**ServerMinDeactivatedRealmDeletionDays** | Pointer to **NullableInt32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The minimum permitted number of days before full data deletion (users, channels, messages, etc.) of a deactivated organization. If &#x60;null&#x60;, then a deactivated organization&#39;s data can be deleted immediately.  **Changes**: New in Zulip 10.0 (feature level 332)  | [optional] 
**ServerMaxDeactivatedRealmDeletionDays** | Pointer to **NullableInt32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum permitted number of days before full data deletion (users, channels, messages, etc.) of a deactivated organization. If &#x60;null&#x60;, then a deactivated organization&#39;s data can be retained indefinitely.  **Changes**: New in Zulip 10.0 (feature level 332).  | [optional] 
**ServerPresencePingIntervalSeconds** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  For clients implementing the [presence](/api/get-presence) system, the time interval the client should use for sending presence requests to the server (and thus receive presence updates from the server).  It is important for presence implementations to use both this and &#x60;server_presence_offline_threshold_seconds&#x60; correctly, so that a Zulip server can change these values to manage the trade-off between load and freshness of presence data.  **Changes**: New in Zulip 7.0 (feature level 164). Clients should use 60 for older Zulip servers, since that&#39;s the value that was hardcoded in the Zulip mobile apps prior to this parameter being introduced.  | [optional] 
**ServerPresenceOfflineThresholdSeconds** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  How old a presence timestamp for a given user can be before the user should be displayed as offline by clients displaying Zulip presence data. See the related &#x60;server_presence_ping_interval_seconds&#x60; for details.  **Changes**: New in Zulip 7.0 (feature level 164). Clients should use 140 for older Zulip servers, since that&#39;s the value that was hardcoded in the Zulip client apps prior to this parameter being introduced.  | [optional] 
**ServerTypingStartedExpiryPeriodMilliseconds** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  For clients implementing [typing notifications](/api/set-typing-status) protocol, the time interval in milliseconds that the client should wait for additional [typing start](/api/get-events#typing-start) events from the server before removing an active typing indicator.  **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 15000 for older Zulip servers, since that&#39;s the value that was hardcoded in the Zulip apps prior to this parameter being introduced.  | [optional] 
**ServerTypingStoppedWaitPeriodMilliseconds** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  For clients implementing [typing notifications](/api/set-typing-status) protocol, the time interval in milliseconds that the client should wait when a user stops interacting with the compose UI before sending a stop notification to the server.  **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 5000 for older Zulip servers, since that&#39;s the value that was hardcoded in the Zulip apps prior to this parameter being introduced.  | [optional] 
**ServerTypingStartedWaitPeriodMilliseconds** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  For clients implementing [typing notifications](/api/set-typing-status) protocol, the time interval in milliseconds that the client should use to send regular start notifications to the server to indicate that the user is still actively interacting with the compose UI.  **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 10000 for older Zulip servers, since that&#39;s the value that was hardcoded in the Zulip apps prior to this parameter being introduced.  | [optional] 
**ScheduledMessages** | Pointer to [**[]ScheduledMessage**](ScheduledMessage.md) | Present if &#x60;scheduled_messages&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of all undelivered scheduled messages by the user.  **Changes**: New in Zulip 7.0 (feature level 179).  | [optional] 
**Reminders** | Pointer to [**[]ScheduledMessage**](ScheduledMessage.md) | Present if &#x60;reminders&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of all undelivered reminders scheduled by the user.  **Changes**: New in Zulip 11.0 (feature level 399).  | [optional] 
**MutedTopics** | Pointer to [**[][]EventEnvelopeOneOf31MutedTopicsInnerInner**]([]EventEnvelopeOneOf31MutedTopicsInnerInner.md) | Present if &#x60;muted_topics&#x60; is present in &#x60;fetch_event_types&#x60;.  Array of tuples, where each tuple describes a muted topic. The first element of the tuple is the channel name in which the topic has to be muted, the second element is the topic name to be muted and the third element is an integer UNIX timestamp representing when the topic was muted.  **Changes**: Deprecated in Zulip 6.0 (feature level 134). Starting with this version, &#x60;muted_topics&#x60; will only be present in the response if the &#x60;user_topic&#x60; object, which generalizes and replaces this field, is not explicitly requested via &#x60;fetch_event_types&#x60;.  Before Zulip 3.0 (feature level 1), the &#x60;muted_topics&#x60; array objects were 2-item tuples and did not include the timestamp information for when the topic was muted.  | [optional] 
**MutedUsers** | Pointer to [**[]EventEnvelopeOneOf33MutedUsersInner**](EventEnvelopeOneOf33MutedUsersInner.md) | Present if &#x60;muted_users&#x60; is present in &#x60;fetch_event_types&#x60;.  A list of dictionaries where each dictionary describes a [muted user](/api/mute-user).  **Changes**: New in Zulip 4.0 (feature level 48).  | [optional] 
**Presences** | Pointer to [**map[string]RegisterQueueResponsePresencesEntry**](RegisterQueueResponsePresencesEntry.md) | Present if &#x60;presence&#x60; is present in &#x60;fetch_event_types&#x60;.  A dictionary where each entry describes the presence details of a user in the Zulip organization.  The format of the entry (modern or legacy) depends on the value of [&#x60;slim_presence&#x60;](#parameter-slim_presence).  Users who have been offline for multiple weeks may not appear in this object.  | [optional] 
**PresenceLastUpdateId** | Pointer to **int32** | Present if &#x60;presence&#x60; is present in &#x60;fetch_event_types&#x60;.  Provides the &#x60;last_update_id&#x60; value of the latest presence data fetched by the server and included in the response in &#x60;presences&#x60;. This can be used as the value of the &#x60;presence_last_update_id&#x60; parameter when polling for presence data at the [/users/me/presence](/api/update-presence) endpoint to tell the server to only fetch the relevant newer data in order to skip redundant already-known presence information.  **Changes**: New in Zulip 9.0 (feature level 263).  | [optional] 
**ServerTimestamp** | Pointer to **float32** | Present if &#x60;presence&#x60; is present in &#x60;fetch_event_types&#x60;.  The time when the server fetched the &#x60;presences&#x60; data included in the response. Matches the similar field in presence responses.  **Changes**: New in Zulip 5.0 (feature level 70).  | [optional] 
**RealmDomains** | Pointer to [**[]RealmDomain**](RealmDomain.md) | Present if &#x60;realm_domains&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary describes a domain within which users can join the organization without and invitation.  | [optional] 
**RealmEmoji** | Pointer to [**map[string]RealmEmoji**](RealmEmoji.md) | Present if &#x60;realm_emoji&#x60; is present in &#x60;fetch_event_types&#x60;.  A dictionary of objects where each object describes a custom emoji that has been uploaded in this Zulip organization.  | [optional] 
**RealmLinkifiers** | Pointer to [**[]RegisterQueueResponseRealmLinkifiersItem**](RegisterQueueResponseRealmLinkifiersItem.md) | Present if &#x60;realm_linkifiers&#x60; is present in &#x60;fetch_event_types&#x60;.  An ordered array of objects where each object describes a single [linkifier](/help/add-a-custom-linkifier).  The order of the array reflects the order that each linkifier should be processed when linkifying messages and topics. By default, new linkifiers are ordered last. This order can be modified with [&#x60;PATCH /realm/linkifiers&#x60;](/api/reorder-linkifiers).  Clients will receive an empty array unless the event queue is registered with the client capability &#x60;{\&quot;linkifier_url_template\&quot;: true}&#x60;. See [&#x60;client_capabilities&#x60;](/api/register-queue#parameter-client_capabilities) parameter for how this can be specified.  **Changes**: Before Zulip 7.0 (feature level 176), the &#x60;linkifier_url_template&#x60; client capability was not required. The requirement was added because linkifiers were updated to contain a URL template instead of a URL format string, which was a not backwards-compatible change.  New in Zulip 4.0 (feature level 54). Clients can access this data for servers on earlier feature levels via the legacy &#x60;realm_filters&#x60; property.  | [optional] 
**RealmFilters** | Pointer to [**[][]EventEnvelopeOneOf51RealmFiltersInnerInner**]([]EventEnvelopeOneOf51RealmFiltersInnerInner.md) | Legacy property for [linkifiers](/help/add-a-custom-linkifier). Present if &#x60;realm_filters&#x60; is present in &#x60;fetch_event_types&#x60;.  When present, this is always an empty array.  **Changes**: Prior to Zulip 7.0 (feature level 176), this was an array of tuples, where each tuple described a linkifier. The first element of the tuple was a string regex pattern which represented the pattern to be linkified on matching, for example &#x60;\&quot;#(?P&lt;id&gt;[123])\&quot;&#x60;. The second element was a URL format string that the pattern should be linkified with. A URL format string for the above example would be &#x60;\&quot;https://realm.com/my_realm_filter/%(id)s\&quot;&#x60;. And the third element was the ID of the realm filter.  **Deprecated** in Zulip 4.0 (feature level 54), replaced by the &#x60;realm_linkifiers&#x60; key.  | [optional] 
**RealmPlaygrounds** | Pointer to [**[]RealmPlayground**](RealmPlayground.md) | Present if &#x60;realm_playgrounds&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary describes a [code playground](/help/code-blocks#code-playgrounds) configured for this Zulip organization.  **Changes**: New in Zulip 4.0 (feature level 49).  | [optional] 
**RealmUserGroups** | Pointer to [**[]UserGroup**](UserGroup.md) | Present if &#x60;realm_user_groups&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary describes a [user group](/help/user-groups) in the Zulip organization.  Deactivated groups will only be included if &#x60;include_deactivated_groups&#x60; client capability is set to &#x60;true&#x60;.  **Changes**: Prior to Zulip 10.0 (feature level 294), deactivated groups were included for all the clients.  | [optional] 
**RealmBots** | Pointer to [**[]Bot**](Bot.md) | Present if &#x60;realm_bot&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary describes a bot that the current user can administer. If the current user is an organization administrator, this will include all bots in the organization. Otherwise, it will only include bots owned by the user (either because the user created the bot or an administrator transferred the bot&#39;s ownership to the user).  | [optional] 
**RealmEmbeddedBots** | Pointer to [**[]RegisterQueueResponseRealmEmbeddedBotsItem**](RegisterQueueResponseRealmEmbeddedBotsItem.md) | Present if &#x60;realm_embedded_bots&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary describes an type of embedded bot that is available to be configured on this Zulip server.  Clients only need these data if they contain UI for creating or administering bots.  | [optional] 
**RealmIncomingWebhookBots** | Pointer to [**[]RegisterQueueResponseRealmIncomingWebhookBotsItem**](RegisterQueueResponseRealmIncomingWebhookBotsItem.md) | Present if &#x60;realm_incoming_webhook_bots&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary describes a type of incoming webhook integration that is available to be configured on this Zulip server.  Clients only need these data if they contain UI for creating or administering bots.  | [optional] 
**RecentPrivateConversations** | Pointer to [**[]RegisterQueueResponseRecentPrivateConversationsItem**](RegisterQueueResponseRecentPrivateConversationsItem.md) | Present if &#x60;recent_private_conversations&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries containing data on all direct message and group direct message conversations that the user has received (or sent) messages in, organized by conversation. This data set is designed to support UI elements such as the \&quot;Direct messages\&quot; widget in the web application showing recent direct message conversations that the user has participated in.  \&quot;Recent\&quot; is defined as the server&#39;s discretion; the original implementation interpreted that as \&quot;the 1000 most recent direct messages the user received\&quot;.  | [optional] 
**NavigationViews** | Pointer to [**[]NavigationView**](NavigationView.md) | Present if &#x60;navigation_views&#x60; is present in &#x60;fetch_event_types&#x60;. An array of dictionaries containing data on all of the current user&#39;s navigation views.  **Changes**: New in Zulip 11.0 (feature level 390).  | [optional] 
**SavedSnippets** | Pointer to [**[]SavedSnippet**](SavedSnippet.md) | Present if &#x60;saved_snippets&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries containing data on all of the current user&#39;s saved snippets.  **Changes**: New in Zulip 10.0 (feature level 297).  | [optional] 
**Subscriptions** | Pointer to [**[]Subscription**](Subscription.md) | Present if &#x60;subscription&#x60; is present in &#x60;fetch_event_types&#x60;.  A array of dictionaries where each dictionary describes the properties of a channel the user is subscribed to (as well as that user&#39;s personal per-channel settings).  **Changes**: Removed &#x60;email_address&#x60; field from the dictionary in Zulip 8.0 (feature level 226).  Removed &#x60;role&#x60; field from the dictionary in Zulip 6.0 (feature level 133).  | [optional] 
**Unsubscribed** | Pointer to [**[]Subscription**](Subscription.md) | Present if &#x60;subscription&#x60; is present in &#x60;fetch_event_types&#x60;.  A array of dictionaries where each dictionary describes one of the channels the user has unsubscribed from but was previously subscribed to along with the subscription details.  Unlike &#x60;never_subscribed&#x60;, the user might have messages in their personal message history that were sent to these channels.  **Changes**: Prior to Zulip 10.0 (feature level 349), if a user was in &#x60;can_administer_channel_group&#x60; of a channel that they had unsubscribed from, but not an organization administrator, the channel in question would not be part of this array.  Removed &#x60;email_address&#x60; field from the dictionary in Zulip 8.0 (feature level 226).  Removed &#x60;role&#x60; field from the dictionary in Zulip 6.0 (feature level 133).  | [optional] 
**NeverSubscribed** | Pointer to [**[]RegisterQueueResponseNeverSubscribedItem**](RegisterQueueResponseNeverSubscribedItem.md) | Present if &#x60;subscription&#x60; is present in &#x60;fetch_event_types&#x60;.  A array of dictionaries where each dictionary describes one of the channels that is visible to the user and the user has never been subscribed to.  Important for clients containing UI where one can browse channels to subscribe to.  **Changes**: Before Zulip 10.0 (feature level 362), archived channels did not appear in this list, even if the &#x60;archived_channels&#x60; [client capability][client-capabilities] was declared by the client.  Prior to Zulip 10.0 (feature level 349), if a user was in &#x60;can_administer_channel_group&#x60; of a channel that they never subscribed to, but not an organization administrator, the channel in question would not be part of this array.  | [optional] 
**ChannelFolders** | Pointer to [**[]ChannelFolder**](ChannelFolder.md) | Present if &#x60;channel_folders&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary describes one of the channel folders in the organization.  Only channel folders with one or more public web channels are visible to spectators.  **Changes**: New in Zulip 11.0 (feature level 389).  | [optional] 
**UnreadMsgs** | Pointer to [**RegisterQueueResponseUnreadMsgs**](RegisterQueueResponseUnreadMsgs.md) |  | [optional] 
**StarredMessages** | Pointer to **[]int32** | Present if &#x60;starred_messages&#x60; is present in &#x60;fetch_event_types&#x60;.  Array containing the IDs of all messages which have been [starred](/help/star-a-message) by the user.  | [optional] 
**Streams** | Pointer to [**[]BasicChannel**](BasicChannel.md) | Present if &#x60;stream&#x60; is present in &#x60;fetch_event_types&#x60;.  Array of dictionaries where each dictionary contains details about a single channel in the organization that is visible to the user.  For organization administrators, this will include all private channels in the organization.  **Changes**: Before Zulip 11.0 (feature level 378), archived channels did not appear in this list, even if the &#x60;archived_channels&#x60; [client capability][client-capabilities] was declared by the client.  As of Zulip 8.0 (feature level 205), this will include all web-public channels in the organization as well.  | [optional] 
**RealmDefaultStreams** | Pointer to **[]int32** | Present if &#x60;default_streams&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of IDs of all the [default channels](/help/set-default-streams-for-new-users) in the organization.  **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single default stream for the Zulip organization.  | [optional] 
**RealmDefaultStreamGroups** | Pointer to [**[]DefaultChannelGroup**](DefaultChannelGroup.md) | Present if &#x60;default_stream_groups&#x60; is present in &#x60;fetch_event_types&#x60;.  An array of dictionaries where each dictionary contains details about a single default channel group configured for this Zulip organization.  Default channel groups are an experimental feature.  | [optional] 
**StopWords** | Pointer to **[]string** | Present if &#x60;stop_words&#x60; is present in &#x60;fetch_event_types&#x60;.  An array containing the stop words used by the Zulip server&#39;s full-text search implementation. Useful for showing helpful error messages when a search returns limited results because a stop word in the query was ignored.  | [optional] 
**UserStatus** | Pointer to [**map[string]RegisterQueueResponseUserStatus**](RegisterQueueResponseUserStatus.md) | Present if &#x60;user_status&#x60; is present in &#x60;fetch_event_types&#x60;.  A dictionary which contains the [status](/help/status-and-availability) of all users in the Zulip organization who have set a status.  **Changes**: The emoji parameters are new in Zulip 5.0 (feature level 86). Previously, Zulip did not support emoji associated with statuses.  | [optional] 
**UserSettings** | Pointer to [**RegisterQueueResponseUserSettings**](RegisterQueueResponseUserSettings.md) |  | [optional] 
**UserTopics** | Pointer to [**[]RegisterQueueResponseUserTopicsItem**](RegisterQueueResponseUserTopicsItem.md) | Present if &#x60;user_topic&#x60; is present in &#x60;fetch_event_types&#x60;.  **Changes**: New in Zulip 6.0 (feature level 134), deprecating and replacing the previous &#x60;muted_topics&#x60; structure.  | [optional] 
**HasZoomToken** | Pointer to **bool** | Present if &#x60;video_calls&#x60; is present in &#x60;fetch_event_types&#x60;.  A boolean which signifies whether the user has a Zoom token and has thus completed OAuth flow for the [Zoom integration](/help/configure-call-provider). Clients need to know whether initiating Zoom OAuth is required before creating a Zoom call.  | [optional] 
**GiphyApiKey** | Pointer to **string** | Present if &#x60;giphy&#x60; is present in &#x60;fetch_event_types&#x60;.  GIPHY&#39;s client-side SDKs needs this API key to use the GIPHY API. GIPHY API keys are not secret (their main purpose appears to be allowing GIPHY to block a problematic app). Please don&#39;t use our API key for an app unrelated to Zulip.  Developers of clients should also read the [GIPHY API TOS](https://support.giphy.com/hc/en-us/articles/360028134111-GIPHY-API-Terms-of-Service-) before using this API key.  **Changes**: Added in Zulip 4.0 (feature level 47).  | [optional] 
**PushDevices** | Pointer to [**map[string]RegisterQueueResponsePushDevicesEntry**](RegisterQueueResponsePushDevicesEntry.md) | Present if &#x60;push_device&#x60; is present in &#x60;fetch_event_types&#x60;.  Dictionary where each entry describes the user&#39;s push device&#39;s registration status and error code (if registration failed).  **Changes**: New in Zulip 11.0 (feature level 406).  | [optional] 
**EnableDesktopNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableDigestEmails** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableLoginEmails** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableMarketingEmails** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EmailNotificationsBatchingPeriodSeconds** | Pointer to **int32** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableOfflineEmailNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableOfflinePushNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableOnlinePushNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableSounds** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableStreamDesktopNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableStreamEmailNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableStreamPushNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableStreamAudibleNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**WildcardMentionsNotify** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**MessageContentInEmailNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**NotificationSound** | Pointer to **string** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**PmContentInDesktopNotifications** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**DesktopIconCountDisplay** | Pointer to **int32** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**RealmNameInEmailNotificationsPolicy** | Pointer to **int32** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: In Zulip 7.0 (feature level 168), replaced previous &#x60;realm_name_in_notifications&#x60; global notifications setting with &#x60;realm_name_in_email_notifications_policy&#x60;.  **Deprecated** since Zulip 5.0 (feature level 89); both &#x60;realm_name_in_notifications&#x60; and the newer &#x60;realm_name_in_email_notifications_policy&#x60; are deprecated. Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**PresenceEnabled** | Pointer to **bool** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The current value of this global notification setting for the user. See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**AvailableNotificationSounds** | Pointer to **[]string** | Present if &#x60;update_global_notifications&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Array containing the names of the notification sound options supported by this Zulip server. Only relevant to support UI for configuring notification sounds.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**ColorScheme** | Pointer to **int32** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The color scheme selected by the user.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**DefaultLanguage** | Pointer to **string** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The default language chosen by the user.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**DemoteInactiveStreams** | Pointer to **int32** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether the user has chosen to hide inactive channels.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**Emojiset** | Pointer to **string** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The name of the emoji set that the user has chosen.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**EnableDraftsSynchronization** | Pointer to **bool** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether drafts synchronization is enabled for the user. If disabled, clients will receive an error when trying to use the &#x60;drafts&#x60; endpoints.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  New in Zulip 5.0 (feature level 87).  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**FluidLayoutWidth** | Pointer to **bool** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether the user has chosen for the layout width to be fluid.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**WebHomeView** | Pointer to **string** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The [home view](/help/configure-home-view) in Zulip, represented as the URL suffix after &#x60;#&#x60; to be rendered when Zulip loads.  Currently supported values are &#x60;all_messages&#x60; and &#x60;recent_topics&#x60;.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;default_view&#x60;, which was new in Zulip 4.0 (feature level 42).  **Deprecated** in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**HighContrastMode** | Pointer to **bool** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether has switched on high contrast mode.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**LeftSideUserlist** | Pointer to **bool** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether the user has chosen for the userlist to be displayed on the left side of the screen (for desktop app and web app) in narrow windows.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**StarredMessageCounts** | Pointer to **bool** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether the user has chosen the number of starred messages to be displayed similar to unread counts.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**Timezone** | Pointer to **string** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  The user&#39;s [profile time zone](/help/change-your-timezone), which is used primarily to display the user&#39;s local time to other users.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**TranslateEmoticons** | Pointer to **bool** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether the user has chosen for emoticons to be translated into emoji in the Zulip compose box.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**TwentyFourHourTime** | Pointer to **bool** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether the user has chosen a twenty four hour time display (true) or a twelve hour one (false).  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**ReceivesTypingNotifications** | Pointer to **bool** | Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.  | [optional] 
**EnterSends** | Pointer to **bool** | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Whether the user setting for [sending on pressing Enter][set-enter-send] in the compose box is enabled.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and process the &#x60;user_settings&#x60; event type instead.  Prior to Zulip 5.0 (feature level 84), this field was present in response if &#x60;realm_user&#x60; was present in &#x60;fetch_event_types&#x60;, not &#x60;update_display_settings&#x60;.  [capabilities]: /api/register-queue#parameter-client_capabilities [set-enter-send]: /help/configure-send-message-keys  | [optional] 
**EmojisetChoices** | Pointer to [**[]RegisterQueueResponseUserSettingsEmojisetChoicesInner**](RegisterQueueResponseUserSettingsEmojisetChoicesInner.md) | Present if &#x60;update_display_settings&#x60; is present in &#x60;fetch_event_types&#x60; and only for clients that did not include &#x60;user_settings_object&#x60; in their [&#x60;client_capabilities&#x60;][capabilities] when registering the event queue.  Array of dictionaries where each dictionary describes an emoji set supported by this version of the Zulip server.  Only relevant to clients with configuration UI for choosing an emoji set; the currently selected emoji set is available in the &#x60;emojiset&#x60; key.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the &#x60;user_settings_object&#x60; client capability and access the &#x60;user_settings&#x60; object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**RealmMessageEditHistoryVisibilityPolicy** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  What typesof message edit history are accessible to users via [message edit history](/help/view-a-messages-edit-history).  - \&quot;all\&quot; &#x3D; All edit history is visible. - \&quot;moves\&quot; &#x3D; Only moves are visible. - \&quot;none\&quot; &#x3D; No edit history is visible.  **Changes**: New in Zulip 10.0 (feature level 358), replacing the previous &#x60;allow_edit_history&#x60; boolean setting; &#x60;true&#x60; corresponds to &#x60;all&#x60;, and &#x60;false&#x60; to &#x60;none&#x60;.  | [optional] 
**RealmAllowEditHistory** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether this organization is configured to allow users to access [message edit history](/help/view-a-messages-edit-history).  The value of &#x60;realm_allow_edit_history&#x60; is set as &#x60;false&#x60; if the &#x60;realm_message_edit_history_visibility_policy&#x60; is configured as \&quot;None\&quot; and &#x60;true&#x60; if it is configured as \&quot;Moves only\&quot; or \&quot;All\&quot;.  **Changes**: Deprecated in Zulip 10.0 (feature level 358) and will be removed in the future, as it is an inaccurate version &#x60;realm_message_edit_history_visibility_policy&#x60;, which replaces this field.  | [optional] 
**RealmCanAddCustomEmojiGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to add custom emoji in the organization.  **Changes**: New in Zulip 10.0 (feature level 307). Previously, this permission was controlled by the enum &#x60;add_custom_emoji_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators.  Before Zulip 5.0 (feature level 85), the &#x60;realm_add_emoji_by_admins_only&#x60; boolean setting controlled this permission; &#x60;true&#x60; corresponded to &#x60;Admins&#x60;, and &#x60;false&#x60; to &#x60;Everyone&#x60;.  | [optional] 
**RealmCanAddSubscribersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to add subscribers to channels in the organization.  **Changes**: New in Zulip 10.0 (feature level 341). Previously, this permission was controlled by the enum &#x60;invite_to_stream_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators.  | [optional] 
**RealmCanDeleteAnyMessageGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to delete any message in the organization.  **Changes**: New in Zulip 10.0 (feature level 281). Previously, this permission was limited to administrators only and was uneditable.  | [optional] 
**RealmCanDeleteOwnMessageGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to delete messages that they have sent in the organization.  **Changes**: New in Zulip 10.0 (feature level 291). Previously, this permission was controlled by the enum &#x60;delete_own_message_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators, 5&#x3D;Everyone.  Before Zulip 5.0 (feature level 101), the &#x60;allow_message_deleting&#x60; boolean setting controlled this permission; &#x60;true&#x60; corresponded to &#x60;Everyone&#x60;, and &#x60;false&#x60; to &#x60;Admins&#x60;.  | [optional] 
**RealmCanSetDeleteMessagePolicyGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to change per-channel &#x60;can_delete_any_message_group&#x60; and &#x60;can_delete_own_message_group&#x60; permission settings. Note that the user must be a member of both this group and the &#x60;can_administer_channel_group&#x60; of the channel whose message delete settings they want to change.  Organization administrators can always change these settings of every channel.  **Changes**: New in Zulip 11.0 (feature level 407).  | [optional] 
**RealmCanSetTopicsPolicyGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to change per-channel &#x60;topics_policy&#x60; setting. Note that the user must be a member of both this group and the &#x60;can_administer_channel_group&#x60; of the channel whose &#x60;topics_policy&#x60; they want to change.  Organization administrators can always change the &#x60;topics_policy&#x60; setting of every channel.  **Changes**: New in Zulip 11.0 (feature level 392).  | [optional] 
**RealmCanInviteUsersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to send email invitations for inviting other users to the organization.  **Changes**: New in Zulip 10.0 (feature level 321). Previously, this permission was controlled by the enum &#x60;invite_to_realm_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators, 6&#x3D;Nobody.  Before Zulip 4.0 (feature level 50), the &#x60;invite_by_admins_only&#x60; boolean setting controlled this permission; &#x60;true&#x60; corresponded to &#x60;Admins&#x60;, and &#x60;false&#x60; to &#x60;Members&#x60;.  | [optional] 
**RealmCanMentionManyUsersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to use wildcard mentions in large channels.  All users will receive a warning/reminder when using mentions in large channels, even when permitted to do so.  **Changes**: New in Zulip 10.0 (feature level 352). Previously, this permission was controlled by the enum &#x60;wildcard_mention_policy&#x60;.  | [optional] 
**RealmCanMoveMessagesBetweenChannelsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to move messages from one channel to another in the organization.  **Changes**: New in Zulip 10.0 (feature level 310). Previously, this permission was controlled by the enum &#x60;move_messages_between_streams_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators, 6&#x3D;Nobody.  In Zulip 7.0 (feature level 159), &#x60;Nobody&#x60; was added as an option to &#x60;move_messages_between_streams_policy&#x60; enum.  | [optional] 
**RealmCanMoveMessagesBetweenTopicsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to move messages from one topic to another within a channel in the organization.  **Changes**: New in Zulip 10.0 (feature level 316). Previously, this permission was controlled by the enum &#x60;edit_topic_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators, 5&#x3D;Everyone, 6&#x3D;Nobody.  In Zulip 7.0 (feature level 159), &#x60;Nobody&#x60; was added as an option to &#x60;edit_topic_policy&#x60; enum.  | [optional] 
**RealmCanCreateGroups** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to create user groups in this organization.  **Changes**: New in Zulip 10.0 (feature level 299). Previously &#x60;realm_user_group_edit_policy&#x60; field used to control the permission to create user groups.  | [optional] 
**RealmCanCreateBotsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to create all types of bot users in the organization. See also &#x60;can_create_write_only_bots_group&#x60;.  **Changes**: New in Zulip 10.0 (feature level 344). Previously, this permission was controlled by the enum &#x60;bot_creation_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Generic bots limited to administrators, 3&#x3D;Administrators.  | [optional] 
**RealmCanCreateWriteOnlyBotsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to create bot users that can only send messages in the organization, i.e. incoming webhooks, in addition to the users who are present in &#x60;can_create_bots_group&#x60;.  **Changes**: New in Zulip 10.0 (feature level 344). Previously, this permission was controlled by the enum &#x60;bot_creation_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Generic bots limited to administrators, 3&#x3D;Administrators.  | [optional] 
**RealmCanManageAllGroups** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to administer all existing groups in this organization.  **Changes**: Prior to Zulip 10.0 (feature level 305), only users who were a member of the group or had the moderator role or above could exercise the permission on a given group.  New in Zulip 10.0 (feature level 299). Previously the &#x60;user_group_edit_policy&#x60; field controlled the permission to manage user groups. Valid values were as follows:  - 1 &#x3D; All members can create and edit user groups - 2 &#x3D; Only organization administrators can create and edit   user groups - 3 &#x3D; Only [full members][calc-full-member] can create and   edit user groups. - 4 &#x3D; Only organization administrators and moderators can   create and edit user groups.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**RealmCanManageBillingGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to manage plans and billing in the organization.  **Changes**: New in Zulip 10.0 (feature level 363). Previously, only owners and users with &#x60;is_billing_admin&#x60; property set to &#x60;true&#x60; were allowed to manage plans and billing.  | [optional] 
**RealmCanCreatePublicChannelGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to create public channels in this organization.  **Changes**: New in Zulip 9.0 (feature level 264). Previously &#x60;realm_create_public_stream_policy&#x60; field used to control the permission to create public channels.  | [optional] 
**RealmCanCreatePrivateChannelGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to create private channels in this organization.  **Changes**: New in Zulip 9.0 (feature level 266). Previously &#x60;realm_create_private_stream_policy&#x60; field used to control the permission to create private channels.  | [optional] 
**RealmCanCreateWebPublicChannelGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to create web-public channels in this organization.  Has no effect and should not be displayed in settings UI unless the Zulip server has the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; server-level setting enabled and the organization has enabled the &#x60;enable_spectator_access&#x60; realm setting.  **Changes**: New in Zulip 10.0 (feature level 280). Previously &#x60;realm_create_web_public_stream_policy&#x60; field used to control the permission to create web-public channels.  | [optional] 
**RealmCanResolveTopicsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to [resolve topics](/help/resolve-a-topic) in the organization.  **Changes**: New in Zulip 10.0 (feature level 367). Previously, permission to resolve topics was controlled by the more general &#x60;can_move_messages_between_topics_group permission for moving messages&#x60;.  | [optional] 
**RealmCreatePublicStreamPolicy** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A deprecated representation of a superset of the users who have permission to create public channels in the organization, available for backwards-compatibility. Clients should use &#x60;can_create_public_channel_group&#x60; instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 1 &#x3D; Members only - 2 &#x3D; Admins only - 3 &#x3D; [Full members][calc-full-member] only - 4 &#x3D; Admins and moderators only  **Changes**: Deprecated in Zulip 9.0 (feature level 264) and replaced by &#x60;realm_can_create_public_channel_group&#x60;, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  Before Zulip 5.0 (feature level 102), permission to create channels was controlled by the &#x60;realm_create_stream_policy&#x60; setting.  [permission-level]: /api/roles-and-permissions#permission-levels [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**RealmCreatePrivateStreamPolicy** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A deprecated representation of a superset of the users who have permission to create private channels in the organization, available for backwards-compatibility. Clients should use &#x60;can_create_private_channel_group&#x60; instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 1 &#x3D; Members only - 2 &#x3D; Admins only - 3 &#x3D; [Full members][calc-full-member] only - 4 &#x3D; Admins and moderators only  **Changes**: Deprecated in Zulip 9.0 (feature level 266) and replaced by &#x60;realm_can_create_private_channel_group&#x60;, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  **Changes**: Before Zulip 5.0 (feature level 102), permission to create channels was controlled by the &#x60;realm_create_stream_policy&#x60; setting.  [permission-level]: /api/roles-and-permissions#permission-levels [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**RealmCreateWebPublicStreamPolicy** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A deprecated representation of a superset of the users who have permission to create web-public channels in the organization, available for backwards-compatibility. Clients should use &#x60;can_create_web_public_channel_group&#x60; instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 2 &#x3D; Admins only - 4 &#x3D; Admins and moderators only - 6 &#x3D; Nobody - 7 &#x3D; Owners only  **Changes**: Deprecated in Zulip 10.0 (feature level 280) and replaced by &#x60;realm_can_create_web_public_channel_group&#x60;, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  **Changes**: Added in Zulip 5.0 (feature level 103).  [permission-level]: /api/roles-and-permissions#permission-levels  | [optional] 
**RealmWildcardMentionPolicy** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A deprecated representation of a superset of the users who have permission to use wildcard mentions in large channels, available for backwards-compatibility. Clients should use &#x60;can_mention_many_users_group&#x60; instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 1 &#x3D; Any user can use wildcard mentions in large channels. - 2 &#x3D; Only members can use wildcard mentions in large channels. - 3 &#x3D; Only [full members][calc-full-member] can use wildcard mentions in large channels. - 5 &#x3D; Only organization administrators can use wildcard mentions in large channels. - 6 &#x3D; Nobody can use wildcard mentions in large channels. - 7 &#x3D; Only organization administrators and moderators can use wildcard mentions in large channels.  All users will receive a warning/reminder when using mentions in large channels, even when permitted to do so.  **Changes**: Deprecated in Zulip 10.0 (feature level 352) and replaced by &#x60;realm_can_mention_many_users_group&#x60;, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  Channel administrators option removed in Zulip 6.0 (feature level 133).  Moderators option added in Zulip 4.0 (feature level 62).  New in Zulip 4.0 (feature level 33).  [permission-level]: /api/roles-and-permissions#permission-levels [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**RealmDefaultLanguage** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The [organization language][org-lang] for automated messages and invitation emails.  [org-lang]: /help/configure-organization-language  | [optional] 
**RealmWelcomeMessageCustomText** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  This organization&#39;s configured custom message for Welcome Bot to send to new user accounts, in Zulip Markdown format.  **Changes**: New in Zulip 11.0 (feature level 416).  | [optional] 
**RealmDescription** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The description of the organization, used on login and registration pages.  | [optional] 
**RealmDigestEmailsEnabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the organization has enabled [weekly digest emails](/help/digest-emails).  | [optional] 
**RealmDisallowDisposableEmailAddresses** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the organization disallows disposable email addresses.  | [optional] 
**RealmEmailChangesDisabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether users are allowed to change their own email address in this organization. This is typically disabled for organizations that synchronize accounts from LDAP or a similar corporate database.  | [optional] 
**RealmInviteRequired** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether an invitation is required to join this organization.  | [optional] 
**RealmCreateMultiuseInviteGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who are allowed to create [reusable invitation links](/help/invite-new-users#create-a-reusable-invitation-link) to the organization.  **Changes**: Prior to Zulip 10.0 (feature level 314), this value used to be of type integer and did not accept anonymous user groups.  New in Zulip 8.0 (feature level 209).  | [optional] 
**RealmInlineImagePreview** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether this organization has been configured to enable [previews of linked images](/help/image-video-and-website-previews).  | [optional] 
**RealmInlineUrlEmbedPreview** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether this organization has been configured to enable [previews of linked websites](/help/image-video-and-website-previews).  | [optional] 
**RealmTopicsPolicy** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The organization&#39;s default policy for sending channel messages to the [empty \&quot;general chat\&quot; topic](/help/require-topics).  - &#x60;\&quot;allow_empty_topic\&quot;&#x60;: Channel messages can be sent to the empty topic. - &#x60;\&quot;disable_empty_topic\&quot;&#x60;: Channel messages cannot be sent to the empty topic.  **Changes**: New in Zulip 11.0 (feature level 392). Previously, this was controlled by the boolean &#x60;realm_mandatory_topics&#x60; setting, which is now deprecated.  | [optional] 
**RealmMandatoryTopics** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether [topics are required](/help/require-topics) for messages in this organization.  **Changes**: Deprecated in Zulip 11.0 (feature level 392). This is now controlled by the realm &#x60;topics_policy&#x60; setting.  | [optional] 
**RealmMessageRetentionDays** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The default [message retention policy](/help/message-retention-policy) for this organization. It can have one special value:  - &#x60;-1&#x60; denoting that the messages will be retained forever for this realm, by default.  **Changes**: Prior to Zulip 3.0 (feature level 22), no limit was encoded as &#x60;null&#x60; instead of &#x60;-1&#x60;. Clients can correctly handle all server versions by treating both &#x60;-1&#x60; and &#x60;null&#x60; as indicating unlimited message retention.  | [optional] 
**RealmName** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The name of the organization, used in login pages etc.  | [optional] 
**RealmRequireE2eePushNotifications** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether this realm is configured to disallow sending mobile push notifications with message content through the legacy mobile push notifications APIs. The new API uses end-to-end encryption to protect message content and metadata from being accessible to the push bouncer service, APNs, and FCM. Clients that support the new E2EE API will use it automatically regardless of this setting.  If &#x60;true&#x60;, mobile push notifications sent to clients that lack support for E2EE push notifications will always have \&quot;New message\&quot; as their content. Note that these legacy mobile notifications will still contain metadata, which may include the message&#39;s ID, the sender&#39;s name, email address, and avatar.  In a future release, once the official mobile apps have implemented fully validated their E2EE protocol support, this setting will become strict, and disable the legacy protocol entirely.  **Changes**: New in Zulip 11.0 (feature level 409). Previously, this behavior was available only via the &#x60;PUSH_NOTIFICATION_REDACT_CONTENT&#x60; global server setting.  | [optional] 
**RealmRequireUniqueNames** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Indicates whether the organization is configured to require users to have unique full names. If true, the server will reject attempts to create a new user, or change the name of an existing user, where doing so would lead to two users whose names are identical modulo case and unicode normalization.  **Changes**: New in Zulip 9.0 (feature level 246). Previously, the Zulip server could not be configured to enforce unique names.  | [optional] 
**RealmNameChangesDisabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Indicates whether users are [allowed to change](/help/restrict-name-and-email-changes) their name via the Zulip UI in this organization. Typically disabled in organizations syncing this type of account information from an external user database like LDAP.  | [optional] 
**RealmAvatarChangesDisabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Indicates whether users are [allowed to change](/help/restrict-name-and-email-changes) their avatar via the Zulip UI in this organization. Typically disabled in organizations syncing this type of account information from an external user database like LDAP.  | [optional] 
**RealmEmailsRestrictedToDomains** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether [new users joining](/help/restrict-account-creation#configuring-email-domain-restrictions) this organization are required to have an email address in one of the &#x60;realm_domains&#x60; configured for the organization.  | [optional] 
**RealmSendWelcomeEmails** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether or not this organization is configured to send the standard Zulip [welcome emails](/help/disable-welcome-emails) to new users joining the organization.  | [optional] 
**RealmMessageContentAllowedInEmailNotifications** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether notification emails in this organization are allowed to contain Zulip the message content, or simply indicate that a new message was sent.  | [optional] 
**RealmEnableSpectatorAccess** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether web-public channels and related anonymous access APIs/features are enabled in this organization.  Can only be enabled if the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; [server setting][server-settings] is enabled on the Zulip server. See also the &#x60;can_create_web_public_channel_group&#x60; realm setting.  **Changes**: New in Zulip 5.0 (feature level 109).  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  | [optional] 
**RealmWantAdvertiseInCommunitiesDirectory** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the organization has given permission to be advertised in the Zulip [communities directory](/help/communities-directory).  Useful only to clients supporting changing this setting for the organization.  Giving permission via this setting does not guarantee that an organization will be listed in the Zulip communities directory.  **Changes**: New in Zulip 6.0 (feature level 129).  | [optional] 
**RealmVideoChatProvider** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The configured [video call provider](/help/configure-call-provider) for the organization.  - 0 &#x3D; None - 1 &#x3D; Jitsi Meet - 3 &#x3D; Zoom (User OAuth integration) - 4 &#x3D; BigBlueButton - 5 &#x3D; Zoom (Server to Server OAuth integration)  Note that only one of the [Zoom integrations][zoom-video-calls] can be configured on a Zulip server.  **Changes**: In Zulip 10.0 (feature level 353), added the Zoom Server to Server OAuth option.  In Zulip 3.0 (feature level 1), added the None option to disable video call UI.  [zoom-video-calls]: https://zulip.readthedocs.io/en/latest/production/video-calls.html#zoom  | [optional] 
**RealmJitsiServerUrl** | Pointer to **NullableString** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL of the custom Jitsi Meet server configured in this organization&#39;s settings.  &#x60;null&#x60;, the default, means that the organization is using the should use the server-level configuration, &#x60;server_jitsi_server_url&#x60;. A correct client supporting only the modern API should use &#x60;realm_jitsi_server_url || server_jitsi_server_url&#x60; to create calls.  **Changes**: New in Zulip 8.0 (feature level 212). Previously, this was only available as a server-level configuration, which was available via the &#x60;jitsi_server_url&#x60; field.  | [optional] 
**RealmGiphyRating** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The configured GIPHY rating for the organization.  **Changes**: New in Zulip 4.0 (feature level 55).  | [optional] 
**RealmWaitingPeriodThreshold** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Members whose accounts have been created at least this many days ago will be treated as [full members][calc-full-member] for the purpose of settings that restrict access to new members.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**RealmDigestWeekday** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The day of the week when the organization will send its weekly digest email to inactive users.  | [optional] 
**RealmDirectMessageInitiatorGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to start a new direct message conversation involving other non-bot users. Users who are outside this group and attempt to send the first direct message to a given collection of recipient users will receive an error, unless all other recipients are bots or the sender.  **Changes**: New in Zulip 9.0 (feature level 270).  Previously, access to send direct messages was controlled by the &#x60;private_message_policy&#x60; realm setting, which supported values of 1 (enabled) and 2 (disabled).  | [optional] 
**RealmDirectMessagePermissionGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who have permission to fully use direct messages. Users outside this group can only send direct messages to conversations where all the recipients are in this group, are bots, or are the sender, ensuring that every direct message conversation will be visible to at least one user in this group.  **Changes**: New in Zulip 9.0 (feature level 270).  Previously, access to send direct messages was controlled by the &#x60;private_message_policy&#x60; realm setting, which supported values of 1 (enabled) and 2 (disabled).  | [optional] 
**RealmDefaultCodeBlockLanguage** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The default pygments language code to be used for code blocks in this organization. If an empty string, no default has been set.  **Changes**: Prior to Zulip 8.0 (feature level 195), a server bug meant that both &#x60;null&#x60; and an empty string could represent that no default was set for this realm setting. Clients supporting older server versions should treat either value (&#x60;null&#x60; or &#x60;\&quot;\&quot;&#x60;) as no default being set.  | [optional] 
**RealmMessageContentDeleteLimitSeconds** | Pointer to **NullableInt32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Messages sent more than this many seconds ago cannot be deleted with this organization&#39;s [message deletion policy](/help/restrict-message-editing-and-deletion).  Will not be 0. A &#x60;null&#x60; value means no limit: messages can be deleted regardless of how long ago they were sent.  **Changes**: No limit was represented using the special value &#x60;0&#x60; before Zulip 5.0 (feature level 100).  | [optional] 
**RealmAuthenticationMethods** | Pointer to [**map[string]RealmAuthenticationMethod**](RealmAuthenticationMethod.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Dictionary of authentication method keys mapped to dictionaries that describe the properties of the named authentication method for the organization - its enabled status and availability for use by the organization.  Clients should use this to implement server-settings UI to change which methods are enabled for the organization. For authentication UI itself, clients should use the pre-authentication metadata returned by [&#x60;GET /server_settings&#x60;](/api/get-server-settings).  **Changes**: In Zulip 9.0 (feature level 241), the values in this dictionary were changed. Previously, the values were a simple boolean indicating whether the backend is enabled or not.  | [optional] 
**RealmAllowMessageEditing** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether this organization&#39;s [message edit policy][config-message-editing] allows editing the content of messages.  See [&#x60;PATCH /messages/{message_id}&#x60;](/api/update-message) for details and history of how message editing permissions work.  [config-message-editing]: /help/restrict-message-editing-and-deletion  | [optional] 
**RealmMessageContentEditLimitSeconds** | Pointer to **NullableInt32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Messages sent more than this many seconds ago cannot be edited with this organization&#39;s [message edit policy](/help/restrict-message-editing-and-deletion).  Will not be &#x60;0&#x60;. A &#x60;null&#x60; value means no limit, so messages can be edited regardless of how long ago they were sent.  See [&#x60;PATCH /messages/{message_id}&#x60;](/api/update-message) for details and history of how message editing permissions work.  **Changes**: Before Zulip 6.0 (feature level 138), no limit was represented using the special value &#x60;0&#x60;.  | [optional] 
**RealmMoveMessagesWithinStreamLimitSeconds** | Pointer to **NullableInt32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Messages sent more than this many seconds ago cannot be moved within a channel to another topic by users who have permission to do so based on this organization&#39;s [topic edit policy](/help/restrict-moving-messages). This setting does not affect moderators and administrators.  Will not be &#x60;0&#x60;. A &#x60;null&#x60; value means no limit, so message topics can be edited regardless of how long ago they were sent.  See [&#x60;PATCH /messages/{message_id}&#x60;](/api/update-message) for details and history of how message editing permissions work.  **Changes**: New in Zulip 7.0 (feature level 162). Previously, this time limit was always 72 hours for users who were not administrators or moderators.  | [optional] 
**RealmMoveMessagesBetweenStreamsLimitSeconds** | Pointer to **NullableInt32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Messages sent more than this many seconds ago cannot be moved between channels by users who have permission to do so based on this organization&#39;s [message move policy](/help/restrict-moving-messages). This setting does not affect moderators and administrators.  Will not be &#x60;0&#x60;. A &#x60;null&#x60; value means no limit, so messages can be moved regardless of how long ago they were sent.  See [&#x60;PATCH /messages/{message_id}&#x60;](/api/update-message) for details and history of how message editing permissions work.  **Changes**: New in Zulip 7.0 (feature level 162). Previously, there was no time limit for moving messages between channels for users with permission to do so.  | [optional] 
**RealmEnableReadReceipts** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether read receipts is enabled in the organization or not.  If disabled, read receipt data will be unavailable to clients, regardless of individual users&#39; personal read receipt settings. See also the &#x60;send_read_receipts&#x60; setting within &#x60;realm_user_settings_defaults&#x60;.  **Changes**: New in Zulip 6.0 (feature level 137).  | [optional] 
**RealmIconUrl** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL of the organization&#39;s [profile icon](/help/create-your-organization-profile).  | [optional] 
**RealmIconSource** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  String indicating whether the organization&#39;s [profile icon](/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization&#39;s icon.  - \&quot;G\&quot; means generated by Gravatar (the default). - \&quot;U\&quot; means uploaded by an organization administrator.  | [optional] 
**MaxIconFileSizeMib** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum file size allowed for the organization&#39;s icon. Useful for UI allowing editing the organization&#39;s icon.  **Changes**: New in Zulip 5.0 (feature level 72). Previously, this was called &#x60;max_icon_file_size&#x60;.  | [optional] 
**RealmLogoUrl** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL of the organization&#39;s wide logo configured in the [organization profile](/help/create-your-organization-profile).  | [optional] 
**RealmLogoSource** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  String indicating whether the organization&#39;s [profile wide logo](/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization&#39;s wide logo.  - \&quot;D\&quot; means the logo is the default Zulip logo. - \&quot;U\&quot; means uploaded by an organization administrator.  | [optional] 
**RealmNightLogoUrl** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL of the organization&#39;s dark theme wide-format logo configured in the [organization profile](/help/create-your-organization-profile).  | [optional] 
**RealmNightLogoSource** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  String indicating whether the organization&#39;s dark theme [profile wide logo](/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization&#39;s wide logo.  - \&quot;D\&quot; means the logo is the default Zulip logo. - \&quot;U\&quot; means uploaded by an organization administrator.  | [optional] 
**MaxLogoFileSizeMib** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum file size allowed for the uploaded organization logos.  **Changes**: New in Zulip 5.0 (feature level 72). Previously, this was called &#x60;max_logo_file_size&#x60;.  | [optional] 
**RealmBotDomain** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The fake email domain that will be used for new bots created this organization. Useful for UI for creating bots.  | [optional] 
**RealmUri** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL for the organization. Alias of &#x60;realm_url&#x60;.  **Changes**: Deprecated in Zulip 9.0 (feature level 257). The term \&quot;URI\&quot; is deprecated in [web standards](https://url.spec.whatwg.org/#goals).  | [optional] 
**RealmUrl** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL for the organization.  **Changes**: New in Zulip 9.0 (feature level 257), replacing the deprecated &#x60;realm_uri&#x60;.  | [optional] 
**RealmAvailableVideoChatProviders** | Pointer to [**map[string]VideoChatProviderEntry**](VideoChatProviderEntry.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Dictionary where each entry describes a supported [video call provider](/help/configure-call-provider) that is configured on this server and could be selected by an organization administrator.  Useful for administrative settings UI that allows changing the realm setting &#x60;video_chat_provider&#x60;.  | [optional] 
**RealmPresenceDisabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether online presence of other users is shown in this organization.  | [optional] 
**SettingsSendDigestEmails** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether this Zulip server is configured to allow organizations to enable [digest emails](/help/digest-emails).  Relevant for administrative settings UI that can change the digest email settings.  | [optional] 
**RealmIsZephyrMirrorRealm** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the organization is a Zephyr mirror realm.  | [optional] 
**RealmEmailAuthEnabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the organization has enabled Zulip&#39;s default email and password authentication feature. Determines whether Zulip stores a password for the user and clients should offer any UI for changing the user&#39;s Zulip password.  | [optional] 
**RealmPasswordAuthEnabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the organization allows any sort of password-based authentication (whether via EmailAuthBackend or LDAP passwords).  Determines whether a client might ever need to display a password prompt (clients will primarily look at this attribute in [server_settings](/api/get-server-settings) before presenting a login page).  | [optional] 
**RealmPushNotificationsEnabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether push notifications are enabled for this organization. Typically &#x60;true&#x60; for Zulip Cloud and self-hosted realms that have a valid registration for the [Mobile push notifications service](https://zulip.readthedocs.io/en/latest/production/mobile-push-notifications.html), and &#x60;false&#x60; for self-hosted servers that do not.  **Changes**: Before Zulip 8.0 (feature level 231), this incorrectly was &#x60;true&#x60; for servers that were partly configured to use the Mobile Push Notifications Service but not properly registered.  | [optional] 
**RealmPushNotificationsEnabledEndTimestamp** | Pointer to **NullableInt32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  If the server expects the realm&#39;s push notifications access to end at a definite time in the future, the UNIX timestamp (UTC) at which this is expected to happen. Mobile clients should use this field to display warnings to users when the indicated timestamp is near.  **Changes**: New in Zulip 8.0 (feature level 231).  | [optional] 
**RealmUploadQuotaMib** | Pointer to **NullableInt32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The total quota for uploaded files in this organization.  Clients are not responsible for checking this quota; it is included in the API only for display purposes.  If &#x60;null&#x60;, there is no limit.  **Changes**: Before Zulip 9.0 (feature level 251), this field was incorrectly measured in bytes, not MiB.  New in Zulip 5.0 (feature level 72). Previously, this was called &#x60;realm_upload_quota&#x60;.  | [optional] 
**RealmOrgType** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The [organization type](/help/organization-type) for the realm. Useful only to clients supporting changing this setting for the organization, or clients implementing onboarding content or other features that varies with organization type.  - 0 &#x3D; Unspecified - 10 &#x3D; Business - 20 &#x3D; Open-source project - 30 &#x3D; Education (non-profit) - 35 &#x3D; Education (for-profit) - 40 &#x3D; Research - 50 &#x3D; Event or conference - 60 &#x3D; Non-profit (registered) - 70 &#x3D; Government - 80 &#x3D; Political group - 90 &#x3D; Community - 100 &#x3D; Personal - 1000 &#x3D; Other  **Changes**: New in Zulip 6.0 (feature level 128).  | [optional] 
**RealmPlanType** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The plan type of the organization.  - 1 &#x3D; Self-hosted organization (SELF_HOSTED) - 2 &#x3D; Zulip Cloud free plan (LIMITED) - 3 &#x3D; Zulip Cloud Standard plan (STANDARD) - 4 &#x3D; Zulip Cloud Standard plan, sponsored for free (STANDARD_FREE)  | [optional] 
**RealmEnableGuestUserDmWarning** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether clients should show a warning when a user is composing a DM to a guest user in this organization.  **Changes**: New in Zulip 10.0 (feature level 348).  | [optional] 
**RealmEnableGuestUserIndicator** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether clients should display \&quot;(guest)\&quot; after the names of guest users to prominently highlight their status.  **Changes**: New in Zulip 8.0 (feature level 216).  | [optional] 
**RealmCanAccessAllUsersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who are allowed to access all users in the organization.  **Changes**: Prior to Zulip 10.0 (feature level 314), this value used to be of type integer and did not accept anonymous user groups.  New in Zulip 8.0 (feature level 225).  | [optional] 
**RealmCanSummarizeTopicsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A [group-setting value](/api/group-setting-values) defining the set of users who are allowed to use AI summarization.  **Changes**: New in Zulip 10.0 (feature level 350).  | [optional] 
**ZulipPlanIsNotLimited** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the organization is using a limited (Zulip Cloud Free) plan.  | [optional] 
**UpgradeTextForWideOrganizationLogo** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Text to use when displaying UI for wide organization logos, a feature that is currently not available on the Zulip Cloud Free plan.  Useful only for clients supporting administrative UI for uploading a new wide organization logo to brand the organization.  | [optional] 
**RealmDefaultExternalAccounts** | Pointer to [**map[string]RegisterQueueResponseRealmDefaultExternalAccountsEntry**](RegisterQueueResponseRealmDefaultExternalAccountsEntry.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Dictionary where each entry describes a default external account type that can be configured with Zulip&#39;s [custom profile fields feature](/help/custom-profile-fields).  **Changes**: New in Zulip 2.1.0.  | [optional] 
**JitsiServerUrl** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The base URL to be used to create Jitsi video calls. Equals &#x60;realm_jitsi_server_url || server_jitsi_server_url&#x60;.  **Changes**: Deprecated in Zulip 8.0 (feature level 212) and will eventually be removed. Previously, the Jitsi server to use was not configurable on a per-realm basis, and this field contained the server&#39;s configured Jitsi server. (Which is now provided as &#x60;server_jitsi_server_url&#x60;). Clients supporting older versions should fall back to this field when creating calls: using &#x60;realm_jitsi_server_url || server_jitsi_server_url&#x60; with newer servers and using &#x60;jitsi_server_url&#x60; with servers below feature level 212.  | [optional] 
**DevelopmentEnvironment** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether this Zulip server is a development environment. Used to control certain features or UI (such as error popups) that should only apply when connected to a Zulip development environment.  | [optional] 
**ServerGeneration** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A timestamp indicating when the process hosting this event queue was started. Clients will likely only find this value useful for inclusion in detailed error reports.  | [optional] 
**PasswordMinLength** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  This Zulip server&#39;s configured minimum required length for passwords. Necessary for password change UI to show whether the password will be accepted.  | [optional] 
**PasswordMaxLength** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  This Zulip server&#39;s configured maximum length for passwords. Necessary for password change UI to show whether the password will be accepted.  **Changes**: New in Zulip 10.0 (feature level 338).  | [optional] 
**PasswordMinGuesses** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  This Zulip server&#39;s configured minimum &#x60;zxcvbn&#x60; minimum guesses. Necessary for password change UI to show whether the password will be accepted.  | [optional] 
**GiphyRatingOptions** | Pointer to [**map[string]RegisterQueueResponseGiphyRatingOptionsEntry**](RegisterQueueResponseGiphyRatingOptionsEntry.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Dictionary where each entry describes a valid rating that is configured on this server and could be selected by an organization administrator.  Useful for administrative settings UI that allows changing the allowed rating of GIFs.  | [optional] 
**MaxFileUploadSizeMib** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum file size that can be uploaded to this Zulip organization.  | [optional] 
**MaxAvatarFileSizeMib** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The maximum avatar size that can be uploaded to this Zulip server.  | [optional] 
**ServerInlineImagePreview** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the server is configured with support for inline image previews. Clients containing administrative UI for changing &#x60;realm_inline_image_preview&#x60; should consult this field before offering that feature.  | [optional] 
**ServerInlineUrlEmbedPreview** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the server is configured with support for inline URL previews. Clients containing administrative UI for changing &#x60;realm_inline_url_embed_preview&#x60; should consult this field before offering that feature.  | [optional] 
**ServerThumbnailFormats** | Pointer to [**[]RegisterQueueResponseServerThumbnailFormatsItem**](RegisterQueueResponseServerThumbnailFormatsItem.md) | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  A list describing the image formats that uploaded images will be thumbnailed into. Any image with a source starting with &#x60;/user_uploads/thumbnail/&#x60; can have its last path component replaced with any of the names contained in this list, to obtain the desired thumbnail size.  **Changes**: New in Zulip 9.0 (feature level 273).  | [optional] 
**ServerAvatarChangesDisabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the server allows avatar changes. Similar to &#x60;realm_avatar_changes_disabled&#x60; but based on the &#x60;AVATAR_CHANGES_DISABLED&#x60; Zulip server level setting.  | [optional] 
**ServerNameChangesDisabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the server allows name changes. Similar to &#x60;realm_name_changes_disabled&#x60; but based on the &#x60;NAME_CHANGES_DISABLED&#x60; Zulip server level setting.  | [optional] 
**ServerNeedsUpgrade** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the server is running an old version based on the Zulip [server release lifecycle](https://zulip.readthedocs.io/en/latest/overview/release-lifecycle.html#upgrade-nag), such that the web app will display to the current user a prominent warning.  **Changes**: New in Zulip 5.0 (feature level 74).  | [optional] 
**ServerWebPublicStreamsEnabled** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The value of the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; Zulip server level setting. A server that has disabled this setting intends to not offer [web public channels](/help/public-access-option) to realms it hosts. (Zulip Cloud defaults to &#x60;true&#x60;; self-hosted servers default to &#x60;false&#x60;).  Clients should use this to determine whether to offer UI for the realm-level setting for enabling web-public channels (&#x60;realm_enable_spectator_access&#x60;).  **Changes**: New in Zulip 5.0 (feature level 110).  | [optional] 
**ServerEmojiDataUrl** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL to a JSON file that describes which emoji names map to which emoji codes, for all Unicode emoji this Zulip server accepts.  The data at the given URL is a JSON object with one property, &#x60;code_to_names&#x60;. The value of that property is a JSON object where each key is an [emoji code](/api/add-reaction#parameter-emoji_code) for an available Unicode emoji, and each value is the corresponding [emoji names](/api/add-reaction#parameter-emoji_name) for this emoji, with the canonical name for the emoji always appearing first.  The HTTP response at that URL will have appropriate HTTP caching headers, such any HTTP implementation should get a cached version if emoji haven&#39;t changed since the last request.  **Changes**: New in Zulip 6.0 (feature level 140).  | [optional] 
**ServerJitsiServerUrl** | Pointer to **NullableString** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL of the Jitsi server that the Zulip server is configured to use by default; the organization-level setting &#x60;realm_jitsi_server_url&#x60; takes precedence over this setting when both are set.  **Changes**: New in Zulip 8.0 (feature level 212). Previously, this value was available as the now-deprecated &#x60;jitsi_server_url&#x60;.  | [optional] 
**ServerCanSummarizeTopics** | Pointer to **bool** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;  Whether topic summarization is enabled in the server or not depending upon whether &#x60;TOPIC_SUMMARIZATION_MODEL&#x60; is set or not.  **Changes**: New in Zulip 10.0 (feature level 350).  | [optional] 
**EventQueueLongpollTimeoutSeconds** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Recommended client-side HTTP request timeout for [&#x60;GET /events&#x60;](/api/get-events) calls. This is guaranteed to be somewhat greater than the heartbeat frequency. It is important that clients respect this parameter, so that increases in the heartbeat frequency do not break clients.  **Changes**: New in Zulip 5.0 (feature level 74). Previously, this was hardcoded to 90 seconds, and clients should use that as a fallback value when interacting with servers where this field is not present.  | [optional] 
**RealmBilling** | Pointer to [**RegisterQueueResponseRealmBilling**](RegisterQueueResponseRealmBilling.md) |  | [optional] 
**RealmModerationRequestChannelId** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The ID of the private channel to which messages flagged by users for moderation are sent. Moderators can use this channel to review and act on reported content.  Will be &#x60;-1&#x60; if moderation requests are disabled.  Clients should check whether moderation requests are disabled to determine whether to present a \&quot;report message\&quot; feature in their UI within a given organization.  **Changes**: New in Zulip 10.0 (feature level 331). Previously, no \&quot;report message\&quot; feature existed in Zulip.  | [optional] 
**RealmNewStreamAnnouncementsStreamId** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The ID of the channel to which automated messages announcing the [creation of new channels][new-channel-announce] are sent.  Will be &#x60;-1&#x60; if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  [new-channel-announce]: /help/configure-automated-notices#new-channel-announcements  **Changes**: In Zulip 9.0 (feature level 241), renamed &#39;realm_notifications_stream_id&#39; to &#x60;realm_new_stream_announcements_stream_id&#x60;.  | [optional] 
**RealmSignupAnnouncementsStreamId** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The ID of the channel to which automated messages announcing that [new users have joined the organization][new-user-announce] are sent.  Will be &#x60;-1&#x60; if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  [new-user-announce]: /help/configure-automated-notices#new-user-announcements  **Changes**: In Zulip 9.0 (feature level 241), renamed &#39;realm_signup_notifications_stream_id&#39; to &#x60;realm_signup_announcements_stream_id&#x60;.  | [optional] 
**RealmZulipUpdateAnnouncementsStreamId** | Pointer to **int32** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  The ID of the channel to which automated messages announcing new features or other end-user updates about the Zulip software are sent.  Will be &#x60;-1&#x60; if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  **Changes**: New in Zulip 9.0 (feature level 242).  | [optional] 
**RealmEmptyTopicDisplayName** | Pointer to **string** | Present if &#x60;realm&#x60; is present in &#x60;fetch_event_types&#x60;.  Clients declaring the &#x60;empty_topic_name&#x60; client capability should use the value of &#x60;realm_empty_topic_display_name&#x60; to determine how to display the empty string topic.  Clients not declaring the &#x60;empty_topic_name&#x60; client capability receive &#x60;realm_empty_topic_display_name&#x60; value as the topic name replacing empty string.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic name.  | [optional] 
**RealmUserSettingsDefaults** | Pointer to [**RegisterQueueResponseRealmUserSettingsDefaults**](RegisterQueueResponseRealmUserSettingsDefaults.md) |  | [optional] 
**RealmUsers** | Pointer to [**[]User**](User.md) | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  A array of dictionaries where each entry describes a user whose account has not been deactivated. Note that unlike the usual User dictionary, this does not contain the &#x60;is_active&#x60; key, as all the users present in this array have active accounts.  If the current user is a guest whose access to users is limited by a &#x60;can_access_all_users_group&#x60; policy, and the event queue was registered with the &#x60;user_list_incomplete&#x60; client capability, then users that the current user cannot access will not be included in this array. If the current user&#39;s access to a user is restricted but the client lacks this capability, then that inaccessible user will appear in the users array as an \&quot;Unknown user\&quot; object with the usual format but placeholder data whose only variable content is the user ID.  See also &#x60;cross_realm_bots&#x60; and &#x60;realm_non_active_users&#x60;.  **Changes**: Before Zulip 8.0 (feature level 232), the &#x60;user_list_incomplete&#x60; client capability did not exist, and so all clients whose access to a new user was prevented by &#x60;can_access_all_users_group&#x60; policy would receive a fake \&quot;Unknown user\&quot; event for such users.  | [optional] 
**RealmNonActiveUsers** | Pointer to [**[]User**](User.md) | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  A array of dictionaries where each entry describes a user whose account has been deactivated. Note that unlike the usual User dictionary this does not contain the &#x60;is_active&#x60; key as all the users present in this array have deactivated accounts.  | [optional] 
**AvatarSource** | Pointer to **string** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  The avatar data source type for the current user.  Value values are &#x60;G&#x60; (gravatar) and &#x60;U&#x60; (uploaded by user).  | [optional] 
**AvatarUrlMedium** | Pointer to **string** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  The avatar URL for the current user at 500x500 resolution, appropriate for use in settings UI showing the user&#39;s avatar.  | [optional] 
**AvatarUrl** | Pointer to **string** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  The URL of the avatar for the current user at 100x100 resolution. See also &#x60;avatar_url_medium&#x60;.  | [optional] 
**CanCreateStreams** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is allowed to create at least one type of channel with the organization&#39;s [channel creation policy](/help/configure-who-can-create-channels). Its value will always equal &#x60;can_create_public_streams || can_create_private_streams&#x60;.  **Changes**: Deprecated in Zulip 5.0 (feature level 102), when the new &#x60;create_private_stream_policy&#x60; and &#x60;create_public_stream_policy&#x60; properties introduced the possibility that a user could only create one type of channel.  This field will be removed in a future release.  | [optional] 
**CanCreatePublicStreams** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is allowed to create public channels with the organization&#39;s [channel creation policy](/help/configure-who-can-create-channels).  **Changes**: New in Zulip 5.0 (feature level 102). In older versions, the deprecated &#x60;can_create_streams&#x60; property should be used to determine whether the user can create public channels.  | [optional] 
**CanCreatePrivateStreams** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is allowed to create private channels with the organization&#39;s [channel creation policy](/help/configure-who-can-create-channels).  **Changes**: New in Zulip 5.0 (feature level 102). In older versions, the deprecated &#x60;can_create_streams&#x60; property should be used to determine whether the user can create private channels.  | [optional] 
**CanCreateWebPublicStreams** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is allowed to create public channels with the organization&#39;s [channel creation policy](/help/configure-who-can-create-channels).  Note that this will be false if the Zulip server does not have the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; setting enabled or if the organization has not enabled the &#x60;enable_spectator_access&#x60; realm setting.  **Changes**: New in Zulip 5.0 (feature level 103).  | [optional] 
**CanSubscribeOtherUsers** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is allowed to subscribe other users to channels with the organization&#39;s [channels policy](/help/configure-who-can-invite-to-channels).  | [optional] 
**CanInviteOthersToRealm** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user [is allowed to invite others][who-can-send-invitations] to the organization.  **Changes**: New in Zulip 4.0 (feature level 51).  [who-can-send-invitations]: /help/restrict-account-creation#change-who-can-send-invitations  | [optional] 
**IsAdmin** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is at least an [organization administrator](/api/roles-and-permissions).  | [optional] 
**IsOwner** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is an [organization owner](/api/roles-and-permissions).  **Changes**: New in Zulip 3.0 (feature level 11).  | [optional] 
**IsModerator** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is at least an [organization moderator](/api/roles-and-permissions).  **Changes**: Prior to Zulip 11.0 (feature level 380), this was only true for users whose role was exactly the moderator role.  New in Zulip 4.0 (feature level 60).  | [optional] 
**IsGuest** | Pointer to **bool** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Whether the current user is a [guest user](/api/roles-and-permissions).  | [optional] 
**UserId** | Pointer to **int32** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  The unique ID for the current user.  | [optional] 
**Email** | Pointer to **string** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  The Zulip API email address for the current user. See also &#x60;delivery_email&#x60;; these may be the same or different depending on the user&#39;s &#x60;email_address_visibility&#x60; policy.  | [optional] 
**DeliveryEmail** | Pointer to **string** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  The user&#39;s email address, appropriate for UI for changing the user&#39;s email address. See also &#x60;email&#x60;.  | [optional] 
**FullName** | Pointer to **string** | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  The full name of the current user.  | [optional] 
**CrossRealmBots** | Pointer to [**[]RegisterQueueResponseCrossRealmBotsItem**](RegisterQueueResponseCrossRealmBotsItem.md) | Present if &#x60;realm_user&#x60; is present in &#x60;fetch_event_types&#x60;.  Array of dictionaries where each dictionary contains details of a single cross realm bot. Cross-realm bots are special system bot accounts like Notification Bot.  Most clients will want to combine this with &#x60;realm_users&#x60; in many contexts.  | [optional] 
**ServerSupportedPermissionSettings** | Pointer to [**RegisterQueueResponseServerSupportedPermissionSettings**](RegisterQueueResponseServerSupportedPermissionSettings.md) |  | [optional] 
**MaxBulkNewSubscriptionMessages** | Pointer to **float32** | Maximum number of new subscribers for which the server will respect the &#x60;send_new_subscription_messages&#x60; parameter when [adding subscribers to a channel](/api/subscribe#parameter-send_new_subscription_messages).  **Changes**: New in Zulip 11.0 (feature level 397).  | [optional] 

## Methods

### NewRegisterQueue200Response

`func NewRegisterQueue200Response(result interface{}, msg interface{}, ) *RegisterQueue200Response`

NewRegisterQueue200Response instantiates a new RegisterQueue200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueue200ResponseWithDefaults

`func NewRegisterQueue200ResponseWithDefaults() *RegisterQueue200Response`

NewRegisterQueue200ResponseWithDefaults instantiates a new RegisterQueue200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *RegisterQueue200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RegisterQueue200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RegisterQueue200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *RegisterQueue200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *RegisterQueue200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *RegisterQueue200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *RegisterQueue200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *RegisterQueue200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *RegisterQueue200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *RegisterQueue200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *RegisterQueue200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *RegisterQueue200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *RegisterQueue200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *RegisterQueue200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *RegisterQueue200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *RegisterQueue200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetQueueId

`func (o *RegisterQueue200Response) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *RegisterQueue200Response) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *RegisterQueue200Response) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *RegisterQueue200Response) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### SetQueueIdNil

`func (o *RegisterQueue200Response) SetQueueIdNil(b bool)`

 SetQueueIdNil sets the value for QueueId to be an explicit nil

### UnsetQueueId
`func (o *RegisterQueue200Response) UnsetQueueId()`

UnsetQueueId ensures that no value is present for QueueId, not even an explicit nil
### GetLastEventId

`func (o *RegisterQueue200Response) GetLastEventId() int32`

GetLastEventId returns the LastEventId field if non-nil, zero value otherwise.

### GetLastEventIdOk

`func (o *RegisterQueue200Response) GetLastEventIdOk() (*int32, bool)`

GetLastEventIdOk returns a tuple with the LastEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventId

`func (o *RegisterQueue200Response) SetLastEventId(v int32)`

SetLastEventId sets LastEventId field to given value.

### HasLastEventId

`func (o *RegisterQueue200Response) HasLastEventId() bool`

HasLastEventId returns a boolean if a field has been set.

### GetZulipFeatureLevel

`func (o *RegisterQueue200Response) GetZulipFeatureLevel() int32`

GetZulipFeatureLevel returns the ZulipFeatureLevel field if non-nil, zero value otherwise.

### GetZulipFeatureLevelOk

`func (o *RegisterQueue200Response) GetZulipFeatureLevelOk() (*int32, bool)`

GetZulipFeatureLevelOk returns a tuple with the ZulipFeatureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipFeatureLevel

`func (o *RegisterQueue200Response) SetZulipFeatureLevel(v int32)`

SetZulipFeatureLevel sets ZulipFeatureLevel field to given value.

### HasZulipFeatureLevel

`func (o *RegisterQueue200Response) HasZulipFeatureLevel() bool`

HasZulipFeatureLevel returns a boolean if a field has been set.

### GetZulipVersion

`func (o *RegisterQueue200Response) GetZulipVersion() string`

GetZulipVersion returns the ZulipVersion field if non-nil, zero value otherwise.

### GetZulipVersionOk

`func (o *RegisterQueue200Response) GetZulipVersionOk() (*string, bool)`

GetZulipVersionOk returns a tuple with the ZulipVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipVersion

`func (o *RegisterQueue200Response) SetZulipVersion(v string)`

SetZulipVersion sets ZulipVersion field to given value.

### HasZulipVersion

`func (o *RegisterQueue200Response) HasZulipVersion() bool`

HasZulipVersion returns a boolean if a field has been set.

### GetZulipMergeBase

`func (o *RegisterQueue200Response) GetZulipMergeBase() string`

GetZulipMergeBase returns the ZulipMergeBase field if non-nil, zero value otherwise.

### GetZulipMergeBaseOk

`func (o *RegisterQueue200Response) GetZulipMergeBaseOk() (*string, bool)`

GetZulipMergeBaseOk returns a tuple with the ZulipMergeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipMergeBase

`func (o *RegisterQueue200Response) SetZulipMergeBase(v string)`

SetZulipMergeBase sets ZulipMergeBase field to given value.

### HasZulipMergeBase

`func (o *RegisterQueue200Response) HasZulipMergeBase() bool`

HasZulipMergeBase returns a boolean if a field has been set.

### GetAlertWords

`func (o *RegisterQueue200Response) GetAlertWords() []string`

GetAlertWords returns the AlertWords field if non-nil, zero value otherwise.

### GetAlertWordsOk

`func (o *RegisterQueue200Response) GetAlertWordsOk() (*[]string, bool)`

GetAlertWordsOk returns a tuple with the AlertWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertWords

`func (o *RegisterQueue200Response) SetAlertWords(v []string)`

SetAlertWords sets AlertWords field to given value.

### HasAlertWords

`func (o *RegisterQueue200Response) HasAlertWords() bool`

HasAlertWords returns a boolean if a field has been set.

### GetCustomProfileFields

`func (o *RegisterQueue200Response) GetCustomProfileFields() []CustomProfileField`

GetCustomProfileFields returns the CustomProfileFields field if non-nil, zero value otherwise.

### GetCustomProfileFieldsOk

`func (o *RegisterQueue200Response) GetCustomProfileFieldsOk() (*[]CustomProfileField, bool)`

GetCustomProfileFieldsOk returns a tuple with the CustomProfileFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProfileFields

`func (o *RegisterQueue200Response) SetCustomProfileFields(v []CustomProfileField)`

SetCustomProfileFields sets CustomProfileFields field to given value.

### HasCustomProfileFields

`func (o *RegisterQueue200Response) HasCustomProfileFields() bool`

HasCustomProfileFields returns a boolean if a field has been set.

### GetCustomProfileFieldTypes

`func (o *RegisterQueue200Response) GetCustomProfileFieldTypes() map[string]RegisterQueueResponseCustomProfileFieldTypesEntry`

GetCustomProfileFieldTypes returns the CustomProfileFieldTypes field if non-nil, zero value otherwise.

### GetCustomProfileFieldTypesOk

`func (o *RegisterQueue200Response) GetCustomProfileFieldTypesOk() (*map[string]RegisterQueueResponseCustomProfileFieldTypesEntry, bool)`

GetCustomProfileFieldTypesOk returns a tuple with the CustomProfileFieldTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProfileFieldTypes

`func (o *RegisterQueue200Response) SetCustomProfileFieldTypes(v map[string]RegisterQueueResponseCustomProfileFieldTypesEntry)`

SetCustomProfileFieldTypes sets CustomProfileFieldTypes field to given value.

### HasCustomProfileFieldTypes

`func (o *RegisterQueue200Response) HasCustomProfileFieldTypes() bool`

HasCustomProfileFieldTypes returns a boolean if a field has been set.

### GetRealmDateCreated

`func (o *RegisterQueue200Response) GetRealmDateCreated() int32`

GetRealmDateCreated returns the RealmDateCreated field if non-nil, zero value otherwise.

### GetRealmDateCreatedOk

`func (o *RegisterQueue200Response) GetRealmDateCreatedOk() (*int32, bool)`

GetRealmDateCreatedOk returns a tuple with the RealmDateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDateCreated

`func (o *RegisterQueue200Response) SetRealmDateCreated(v int32)`

SetRealmDateCreated sets RealmDateCreated field to given value.

### HasRealmDateCreated

`func (o *RegisterQueue200Response) HasRealmDateCreated() bool`

HasRealmDateCreated returns a boolean if a field has been set.

### GetDemoOrganizationScheduledDeletionDate

`func (o *RegisterQueue200Response) GetDemoOrganizationScheduledDeletionDate() int32`

GetDemoOrganizationScheduledDeletionDate returns the DemoOrganizationScheduledDeletionDate field if non-nil, zero value otherwise.

### GetDemoOrganizationScheduledDeletionDateOk

`func (o *RegisterQueue200Response) GetDemoOrganizationScheduledDeletionDateOk() (*int32, bool)`

GetDemoOrganizationScheduledDeletionDateOk returns a tuple with the DemoOrganizationScheduledDeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemoOrganizationScheduledDeletionDate

`func (o *RegisterQueue200Response) SetDemoOrganizationScheduledDeletionDate(v int32)`

SetDemoOrganizationScheduledDeletionDate sets DemoOrganizationScheduledDeletionDate field to given value.

### HasDemoOrganizationScheduledDeletionDate

`func (o *RegisterQueue200Response) HasDemoOrganizationScheduledDeletionDate() bool`

HasDemoOrganizationScheduledDeletionDate returns a boolean if a field has been set.

### GetDrafts

`func (o *RegisterQueue200Response) GetDrafts() []Draft`

GetDrafts returns the Drafts field if non-nil, zero value otherwise.

### GetDraftsOk

`func (o *RegisterQueue200Response) GetDraftsOk() (*[]Draft, bool)`

GetDraftsOk returns a tuple with the Drafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrafts

`func (o *RegisterQueue200Response) SetDrafts(v []Draft)`

SetDrafts sets Drafts field to given value.

### HasDrafts

`func (o *RegisterQueue200Response) HasDrafts() bool`

HasDrafts returns a boolean if a field has been set.

### GetOnboardingSteps

`func (o *RegisterQueue200Response) GetOnboardingSteps() []OnboardingStep`

GetOnboardingSteps returns the OnboardingSteps field if non-nil, zero value otherwise.

### GetOnboardingStepsOk

`func (o *RegisterQueue200Response) GetOnboardingStepsOk() (*[]OnboardingStep, bool)`

GetOnboardingStepsOk returns a tuple with the OnboardingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingSteps

`func (o *RegisterQueue200Response) SetOnboardingSteps(v []OnboardingStep)`

SetOnboardingSteps sets OnboardingSteps field to given value.

### HasOnboardingSteps

`func (o *RegisterQueue200Response) HasOnboardingSteps() bool`

HasOnboardingSteps returns a boolean if a field has been set.

### GetNavigationTourVideoUrl

`func (o *RegisterQueue200Response) GetNavigationTourVideoUrl() string`

GetNavigationTourVideoUrl returns the NavigationTourVideoUrl field if non-nil, zero value otherwise.

### GetNavigationTourVideoUrlOk

`func (o *RegisterQueue200Response) GetNavigationTourVideoUrlOk() (*string, bool)`

GetNavigationTourVideoUrlOk returns a tuple with the NavigationTourVideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationTourVideoUrl

`func (o *RegisterQueue200Response) SetNavigationTourVideoUrl(v string)`

SetNavigationTourVideoUrl sets NavigationTourVideoUrl field to given value.

### HasNavigationTourVideoUrl

`func (o *RegisterQueue200Response) HasNavigationTourVideoUrl() bool`

HasNavigationTourVideoUrl returns a boolean if a field has been set.

### SetNavigationTourVideoUrlNil

`func (o *RegisterQueue200Response) SetNavigationTourVideoUrlNil(b bool)`

 SetNavigationTourVideoUrlNil sets the value for NavigationTourVideoUrl to be an explicit nil

### UnsetNavigationTourVideoUrl
`func (o *RegisterQueue200Response) UnsetNavigationTourVideoUrl()`

UnsetNavigationTourVideoUrl ensures that no value is present for NavigationTourVideoUrl, not even an explicit nil
### GetMaxMessageId

`func (o *RegisterQueue200Response) GetMaxMessageId() int32`

GetMaxMessageId returns the MaxMessageId field if non-nil, zero value otherwise.

### GetMaxMessageIdOk

`func (o *RegisterQueue200Response) GetMaxMessageIdOk() (*int32, bool)`

GetMaxMessageIdOk returns a tuple with the MaxMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageId

`func (o *RegisterQueue200Response) SetMaxMessageId(v int32)`

SetMaxMessageId sets MaxMessageId field to given value.

### HasMaxMessageId

`func (o *RegisterQueue200Response) HasMaxMessageId() bool`

HasMaxMessageId returns a boolean if a field has been set.

### GetMaxReminderNoteLength

`func (o *RegisterQueue200Response) GetMaxReminderNoteLength() int32`

GetMaxReminderNoteLength returns the MaxReminderNoteLength field if non-nil, zero value otherwise.

### GetMaxReminderNoteLengthOk

`func (o *RegisterQueue200Response) GetMaxReminderNoteLengthOk() (*int32, bool)`

GetMaxReminderNoteLengthOk returns a tuple with the MaxReminderNoteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReminderNoteLength

`func (o *RegisterQueue200Response) SetMaxReminderNoteLength(v int32)`

SetMaxReminderNoteLength sets MaxReminderNoteLength field to given value.

### HasMaxReminderNoteLength

`func (o *RegisterQueue200Response) HasMaxReminderNoteLength() bool`

HasMaxReminderNoteLength returns a boolean if a field has been set.

### GetMaxStreamNameLength

`func (o *RegisterQueue200Response) GetMaxStreamNameLength() int32`

GetMaxStreamNameLength returns the MaxStreamNameLength field if non-nil, zero value otherwise.

### GetMaxStreamNameLengthOk

`func (o *RegisterQueue200Response) GetMaxStreamNameLengthOk() (*int32, bool)`

GetMaxStreamNameLengthOk returns a tuple with the MaxStreamNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamNameLength

`func (o *RegisterQueue200Response) SetMaxStreamNameLength(v int32)`

SetMaxStreamNameLength sets MaxStreamNameLength field to given value.

### HasMaxStreamNameLength

`func (o *RegisterQueue200Response) HasMaxStreamNameLength() bool`

HasMaxStreamNameLength returns a boolean if a field has been set.

### GetMaxStreamDescriptionLength

`func (o *RegisterQueue200Response) GetMaxStreamDescriptionLength() int32`

GetMaxStreamDescriptionLength returns the MaxStreamDescriptionLength field if non-nil, zero value otherwise.

### GetMaxStreamDescriptionLengthOk

`func (o *RegisterQueue200Response) GetMaxStreamDescriptionLengthOk() (*int32, bool)`

GetMaxStreamDescriptionLengthOk returns a tuple with the MaxStreamDescriptionLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamDescriptionLength

`func (o *RegisterQueue200Response) SetMaxStreamDescriptionLength(v int32)`

SetMaxStreamDescriptionLength sets MaxStreamDescriptionLength field to given value.

### HasMaxStreamDescriptionLength

`func (o *RegisterQueue200Response) HasMaxStreamDescriptionLength() bool`

HasMaxStreamDescriptionLength returns a boolean if a field has been set.

### GetMaxChannelFolderNameLength

`func (o *RegisterQueue200Response) GetMaxChannelFolderNameLength() int32`

GetMaxChannelFolderNameLength returns the MaxChannelFolderNameLength field if non-nil, zero value otherwise.

### GetMaxChannelFolderNameLengthOk

`func (o *RegisterQueue200Response) GetMaxChannelFolderNameLengthOk() (*int32, bool)`

GetMaxChannelFolderNameLengthOk returns a tuple with the MaxChannelFolderNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChannelFolderNameLength

`func (o *RegisterQueue200Response) SetMaxChannelFolderNameLength(v int32)`

SetMaxChannelFolderNameLength sets MaxChannelFolderNameLength field to given value.

### HasMaxChannelFolderNameLength

`func (o *RegisterQueue200Response) HasMaxChannelFolderNameLength() bool`

HasMaxChannelFolderNameLength returns a boolean if a field has been set.

### GetMaxChannelFolderDescriptionLength

`func (o *RegisterQueue200Response) GetMaxChannelFolderDescriptionLength() int32`

GetMaxChannelFolderDescriptionLength returns the MaxChannelFolderDescriptionLength field if non-nil, zero value otherwise.

### GetMaxChannelFolderDescriptionLengthOk

`func (o *RegisterQueue200Response) GetMaxChannelFolderDescriptionLengthOk() (*int32, bool)`

GetMaxChannelFolderDescriptionLengthOk returns a tuple with the MaxChannelFolderDescriptionLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChannelFolderDescriptionLength

`func (o *RegisterQueue200Response) SetMaxChannelFolderDescriptionLength(v int32)`

SetMaxChannelFolderDescriptionLength sets MaxChannelFolderDescriptionLength field to given value.

### HasMaxChannelFolderDescriptionLength

`func (o *RegisterQueue200Response) HasMaxChannelFolderDescriptionLength() bool`

HasMaxChannelFolderDescriptionLength returns a boolean if a field has been set.

### GetMaxTopicLength

`func (o *RegisterQueue200Response) GetMaxTopicLength() int32`

GetMaxTopicLength returns the MaxTopicLength field if non-nil, zero value otherwise.

### GetMaxTopicLengthOk

`func (o *RegisterQueue200Response) GetMaxTopicLengthOk() (*int32, bool)`

GetMaxTopicLengthOk returns a tuple with the MaxTopicLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopicLength

`func (o *RegisterQueue200Response) SetMaxTopicLength(v int32)`

SetMaxTopicLength sets MaxTopicLength field to given value.

### HasMaxTopicLength

`func (o *RegisterQueue200Response) HasMaxTopicLength() bool`

HasMaxTopicLength returns a boolean if a field has been set.

### GetMaxMessageLength

`func (o *RegisterQueue200Response) GetMaxMessageLength() int32`

GetMaxMessageLength returns the MaxMessageLength field if non-nil, zero value otherwise.

### GetMaxMessageLengthOk

`func (o *RegisterQueue200Response) GetMaxMessageLengthOk() (*int32, bool)`

GetMaxMessageLengthOk returns a tuple with the MaxMessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageLength

`func (o *RegisterQueue200Response) SetMaxMessageLength(v int32)`

SetMaxMessageLength sets MaxMessageLength field to given value.

### HasMaxMessageLength

`func (o *RegisterQueue200Response) HasMaxMessageLength() bool`

HasMaxMessageLength returns a boolean if a field has been set.

### GetServerMinDeactivatedRealmDeletionDays

`func (o *RegisterQueue200Response) GetServerMinDeactivatedRealmDeletionDays() int32`

GetServerMinDeactivatedRealmDeletionDays returns the ServerMinDeactivatedRealmDeletionDays field if non-nil, zero value otherwise.

### GetServerMinDeactivatedRealmDeletionDaysOk

`func (o *RegisterQueue200Response) GetServerMinDeactivatedRealmDeletionDaysOk() (*int32, bool)`

GetServerMinDeactivatedRealmDeletionDaysOk returns a tuple with the ServerMinDeactivatedRealmDeletionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMinDeactivatedRealmDeletionDays

`func (o *RegisterQueue200Response) SetServerMinDeactivatedRealmDeletionDays(v int32)`

SetServerMinDeactivatedRealmDeletionDays sets ServerMinDeactivatedRealmDeletionDays field to given value.

### HasServerMinDeactivatedRealmDeletionDays

`func (o *RegisterQueue200Response) HasServerMinDeactivatedRealmDeletionDays() bool`

HasServerMinDeactivatedRealmDeletionDays returns a boolean if a field has been set.

### SetServerMinDeactivatedRealmDeletionDaysNil

`func (o *RegisterQueue200Response) SetServerMinDeactivatedRealmDeletionDaysNil(b bool)`

 SetServerMinDeactivatedRealmDeletionDaysNil sets the value for ServerMinDeactivatedRealmDeletionDays to be an explicit nil

### UnsetServerMinDeactivatedRealmDeletionDays
`func (o *RegisterQueue200Response) UnsetServerMinDeactivatedRealmDeletionDays()`

UnsetServerMinDeactivatedRealmDeletionDays ensures that no value is present for ServerMinDeactivatedRealmDeletionDays, not even an explicit nil
### GetServerMaxDeactivatedRealmDeletionDays

`func (o *RegisterQueue200Response) GetServerMaxDeactivatedRealmDeletionDays() int32`

GetServerMaxDeactivatedRealmDeletionDays returns the ServerMaxDeactivatedRealmDeletionDays field if non-nil, zero value otherwise.

### GetServerMaxDeactivatedRealmDeletionDaysOk

`func (o *RegisterQueue200Response) GetServerMaxDeactivatedRealmDeletionDaysOk() (*int32, bool)`

GetServerMaxDeactivatedRealmDeletionDaysOk returns a tuple with the ServerMaxDeactivatedRealmDeletionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMaxDeactivatedRealmDeletionDays

`func (o *RegisterQueue200Response) SetServerMaxDeactivatedRealmDeletionDays(v int32)`

SetServerMaxDeactivatedRealmDeletionDays sets ServerMaxDeactivatedRealmDeletionDays field to given value.

### HasServerMaxDeactivatedRealmDeletionDays

`func (o *RegisterQueue200Response) HasServerMaxDeactivatedRealmDeletionDays() bool`

HasServerMaxDeactivatedRealmDeletionDays returns a boolean if a field has been set.

### SetServerMaxDeactivatedRealmDeletionDaysNil

`func (o *RegisterQueue200Response) SetServerMaxDeactivatedRealmDeletionDaysNil(b bool)`

 SetServerMaxDeactivatedRealmDeletionDaysNil sets the value for ServerMaxDeactivatedRealmDeletionDays to be an explicit nil

### UnsetServerMaxDeactivatedRealmDeletionDays
`func (o *RegisterQueue200Response) UnsetServerMaxDeactivatedRealmDeletionDays()`

UnsetServerMaxDeactivatedRealmDeletionDays ensures that no value is present for ServerMaxDeactivatedRealmDeletionDays, not even an explicit nil
### GetServerPresencePingIntervalSeconds

`func (o *RegisterQueue200Response) GetServerPresencePingIntervalSeconds() int32`

GetServerPresencePingIntervalSeconds returns the ServerPresencePingIntervalSeconds field if non-nil, zero value otherwise.

### GetServerPresencePingIntervalSecondsOk

`func (o *RegisterQueue200Response) GetServerPresencePingIntervalSecondsOk() (*int32, bool)`

GetServerPresencePingIntervalSecondsOk returns a tuple with the ServerPresencePingIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPresencePingIntervalSeconds

`func (o *RegisterQueue200Response) SetServerPresencePingIntervalSeconds(v int32)`

SetServerPresencePingIntervalSeconds sets ServerPresencePingIntervalSeconds field to given value.

### HasServerPresencePingIntervalSeconds

`func (o *RegisterQueue200Response) HasServerPresencePingIntervalSeconds() bool`

HasServerPresencePingIntervalSeconds returns a boolean if a field has been set.

### GetServerPresenceOfflineThresholdSeconds

`func (o *RegisterQueue200Response) GetServerPresenceOfflineThresholdSeconds() int32`

GetServerPresenceOfflineThresholdSeconds returns the ServerPresenceOfflineThresholdSeconds field if non-nil, zero value otherwise.

### GetServerPresenceOfflineThresholdSecondsOk

`func (o *RegisterQueue200Response) GetServerPresenceOfflineThresholdSecondsOk() (*int32, bool)`

GetServerPresenceOfflineThresholdSecondsOk returns a tuple with the ServerPresenceOfflineThresholdSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPresenceOfflineThresholdSeconds

`func (o *RegisterQueue200Response) SetServerPresenceOfflineThresholdSeconds(v int32)`

SetServerPresenceOfflineThresholdSeconds sets ServerPresenceOfflineThresholdSeconds field to given value.

### HasServerPresenceOfflineThresholdSeconds

`func (o *RegisterQueue200Response) HasServerPresenceOfflineThresholdSeconds() bool`

HasServerPresenceOfflineThresholdSeconds returns a boolean if a field has been set.

### GetServerTypingStartedExpiryPeriodMilliseconds

`func (o *RegisterQueue200Response) GetServerTypingStartedExpiryPeriodMilliseconds() int32`

GetServerTypingStartedExpiryPeriodMilliseconds returns the ServerTypingStartedExpiryPeriodMilliseconds field if non-nil, zero value otherwise.

### GetServerTypingStartedExpiryPeriodMillisecondsOk

`func (o *RegisterQueue200Response) GetServerTypingStartedExpiryPeriodMillisecondsOk() (*int32, bool)`

GetServerTypingStartedExpiryPeriodMillisecondsOk returns a tuple with the ServerTypingStartedExpiryPeriodMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypingStartedExpiryPeriodMilliseconds

`func (o *RegisterQueue200Response) SetServerTypingStartedExpiryPeriodMilliseconds(v int32)`

SetServerTypingStartedExpiryPeriodMilliseconds sets ServerTypingStartedExpiryPeriodMilliseconds field to given value.

### HasServerTypingStartedExpiryPeriodMilliseconds

`func (o *RegisterQueue200Response) HasServerTypingStartedExpiryPeriodMilliseconds() bool`

HasServerTypingStartedExpiryPeriodMilliseconds returns a boolean if a field has been set.

### GetServerTypingStoppedWaitPeriodMilliseconds

`func (o *RegisterQueue200Response) GetServerTypingStoppedWaitPeriodMilliseconds() int32`

GetServerTypingStoppedWaitPeriodMilliseconds returns the ServerTypingStoppedWaitPeriodMilliseconds field if non-nil, zero value otherwise.

### GetServerTypingStoppedWaitPeriodMillisecondsOk

`func (o *RegisterQueue200Response) GetServerTypingStoppedWaitPeriodMillisecondsOk() (*int32, bool)`

GetServerTypingStoppedWaitPeriodMillisecondsOk returns a tuple with the ServerTypingStoppedWaitPeriodMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypingStoppedWaitPeriodMilliseconds

`func (o *RegisterQueue200Response) SetServerTypingStoppedWaitPeriodMilliseconds(v int32)`

SetServerTypingStoppedWaitPeriodMilliseconds sets ServerTypingStoppedWaitPeriodMilliseconds field to given value.

### HasServerTypingStoppedWaitPeriodMilliseconds

`func (o *RegisterQueue200Response) HasServerTypingStoppedWaitPeriodMilliseconds() bool`

HasServerTypingStoppedWaitPeriodMilliseconds returns a boolean if a field has been set.

### GetServerTypingStartedWaitPeriodMilliseconds

`func (o *RegisterQueue200Response) GetServerTypingStartedWaitPeriodMilliseconds() int32`

GetServerTypingStartedWaitPeriodMilliseconds returns the ServerTypingStartedWaitPeriodMilliseconds field if non-nil, zero value otherwise.

### GetServerTypingStartedWaitPeriodMillisecondsOk

`func (o *RegisterQueue200Response) GetServerTypingStartedWaitPeriodMillisecondsOk() (*int32, bool)`

GetServerTypingStartedWaitPeriodMillisecondsOk returns a tuple with the ServerTypingStartedWaitPeriodMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypingStartedWaitPeriodMilliseconds

`func (o *RegisterQueue200Response) SetServerTypingStartedWaitPeriodMilliseconds(v int32)`

SetServerTypingStartedWaitPeriodMilliseconds sets ServerTypingStartedWaitPeriodMilliseconds field to given value.

### HasServerTypingStartedWaitPeriodMilliseconds

`func (o *RegisterQueue200Response) HasServerTypingStartedWaitPeriodMilliseconds() bool`

HasServerTypingStartedWaitPeriodMilliseconds returns a boolean if a field has been set.

### GetScheduledMessages

`func (o *RegisterQueue200Response) GetScheduledMessages() []ScheduledMessage`

GetScheduledMessages returns the ScheduledMessages field if non-nil, zero value otherwise.

### GetScheduledMessagesOk

`func (o *RegisterQueue200Response) GetScheduledMessagesOk() (*[]ScheduledMessage, bool)`

GetScheduledMessagesOk returns a tuple with the ScheduledMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessages

`func (o *RegisterQueue200Response) SetScheduledMessages(v []ScheduledMessage)`

SetScheduledMessages sets ScheduledMessages field to given value.

### HasScheduledMessages

`func (o *RegisterQueue200Response) HasScheduledMessages() bool`

HasScheduledMessages returns a boolean if a field has been set.

### GetReminders

`func (o *RegisterQueue200Response) GetReminders() []ScheduledMessage`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *RegisterQueue200Response) GetRemindersOk() (*[]ScheduledMessage, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *RegisterQueue200Response) SetReminders(v []ScheduledMessage)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *RegisterQueue200Response) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetMutedTopics

`func (o *RegisterQueue200Response) GetMutedTopics() [][]EventEnvelopeOneOf31MutedTopicsInnerInner`

GetMutedTopics returns the MutedTopics field if non-nil, zero value otherwise.

### GetMutedTopicsOk

`func (o *RegisterQueue200Response) GetMutedTopicsOk() (*[][]EventEnvelopeOneOf31MutedTopicsInnerInner, bool)`

GetMutedTopicsOk returns a tuple with the MutedTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedTopics

`func (o *RegisterQueue200Response) SetMutedTopics(v [][]EventEnvelopeOneOf31MutedTopicsInnerInner)`

SetMutedTopics sets MutedTopics field to given value.

### HasMutedTopics

`func (o *RegisterQueue200Response) HasMutedTopics() bool`

HasMutedTopics returns a boolean if a field has been set.

### GetMutedUsers

`func (o *RegisterQueue200Response) GetMutedUsers() []EventEnvelopeOneOf33MutedUsersInner`

GetMutedUsers returns the MutedUsers field if non-nil, zero value otherwise.

### GetMutedUsersOk

`func (o *RegisterQueue200Response) GetMutedUsersOk() (*[]EventEnvelopeOneOf33MutedUsersInner, bool)`

GetMutedUsersOk returns a tuple with the MutedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedUsers

`func (o *RegisterQueue200Response) SetMutedUsers(v []EventEnvelopeOneOf33MutedUsersInner)`

SetMutedUsers sets MutedUsers field to given value.

### HasMutedUsers

`func (o *RegisterQueue200Response) HasMutedUsers() bool`

HasMutedUsers returns a boolean if a field has been set.

### GetPresences

`func (o *RegisterQueue200Response) GetPresences() map[string]RegisterQueueResponsePresencesEntry`

GetPresences returns the Presences field if non-nil, zero value otherwise.

### GetPresencesOk

`func (o *RegisterQueue200Response) GetPresencesOk() (*map[string]RegisterQueueResponsePresencesEntry, bool)`

GetPresencesOk returns a tuple with the Presences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresences

`func (o *RegisterQueue200Response) SetPresences(v map[string]RegisterQueueResponsePresencesEntry)`

SetPresences sets Presences field to given value.

### HasPresences

`func (o *RegisterQueue200Response) HasPresences() bool`

HasPresences returns a boolean if a field has been set.

### GetPresenceLastUpdateId

`func (o *RegisterQueue200Response) GetPresenceLastUpdateId() int32`

GetPresenceLastUpdateId returns the PresenceLastUpdateId field if non-nil, zero value otherwise.

### GetPresenceLastUpdateIdOk

`func (o *RegisterQueue200Response) GetPresenceLastUpdateIdOk() (*int32, bool)`

GetPresenceLastUpdateIdOk returns a tuple with the PresenceLastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceLastUpdateId

`func (o *RegisterQueue200Response) SetPresenceLastUpdateId(v int32)`

SetPresenceLastUpdateId sets PresenceLastUpdateId field to given value.

### HasPresenceLastUpdateId

`func (o *RegisterQueue200Response) HasPresenceLastUpdateId() bool`

HasPresenceLastUpdateId returns a boolean if a field has been set.

### GetServerTimestamp

`func (o *RegisterQueue200Response) GetServerTimestamp() float32`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *RegisterQueue200Response) GetServerTimestampOk() (*float32, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *RegisterQueue200Response) SetServerTimestamp(v float32)`

SetServerTimestamp sets ServerTimestamp field to given value.

### HasServerTimestamp

`func (o *RegisterQueue200Response) HasServerTimestamp() bool`

HasServerTimestamp returns a boolean if a field has been set.

### GetRealmDomains

`func (o *RegisterQueue200Response) GetRealmDomains() []RealmDomain`

GetRealmDomains returns the RealmDomains field if non-nil, zero value otherwise.

### GetRealmDomainsOk

`func (o *RegisterQueue200Response) GetRealmDomainsOk() (*[]RealmDomain, bool)`

GetRealmDomainsOk returns a tuple with the RealmDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDomains

`func (o *RegisterQueue200Response) SetRealmDomains(v []RealmDomain)`

SetRealmDomains sets RealmDomains field to given value.

### HasRealmDomains

`func (o *RegisterQueue200Response) HasRealmDomains() bool`

HasRealmDomains returns a boolean if a field has been set.

### GetRealmEmoji

`func (o *RegisterQueue200Response) GetRealmEmoji() map[string]RealmEmoji`

GetRealmEmoji returns the RealmEmoji field if non-nil, zero value otherwise.

### GetRealmEmojiOk

`func (o *RegisterQueue200Response) GetRealmEmojiOk() (*map[string]RealmEmoji, bool)`

GetRealmEmojiOk returns a tuple with the RealmEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmoji

`func (o *RegisterQueue200Response) SetRealmEmoji(v map[string]RealmEmoji)`

SetRealmEmoji sets RealmEmoji field to given value.

### HasRealmEmoji

`func (o *RegisterQueue200Response) HasRealmEmoji() bool`

HasRealmEmoji returns a boolean if a field has been set.

### GetRealmLinkifiers

`func (o *RegisterQueue200Response) GetRealmLinkifiers() []RegisterQueueResponseRealmLinkifiersItem`

GetRealmLinkifiers returns the RealmLinkifiers field if non-nil, zero value otherwise.

### GetRealmLinkifiersOk

`func (o *RegisterQueue200Response) GetRealmLinkifiersOk() (*[]RegisterQueueResponseRealmLinkifiersItem, bool)`

GetRealmLinkifiersOk returns a tuple with the RealmLinkifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmLinkifiers

`func (o *RegisterQueue200Response) SetRealmLinkifiers(v []RegisterQueueResponseRealmLinkifiersItem)`

SetRealmLinkifiers sets RealmLinkifiers field to given value.

### HasRealmLinkifiers

`func (o *RegisterQueue200Response) HasRealmLinkifiers() bool`

HasRealmLinkifiers returns a boolean if a field has been set.

### GetRealmFilters

`func (o *RegisterQueue200Response) GetRealmFilters() [][]EventEnvelopeOneOf51RealmFiltersInnerInner`

GetRealmFilters returns the RealmFilters field if non-nil, zero value otherwise.

### GetRealmFiltersOk

`func (o *RegisterQueue200Response) GetRealmFiltersOk() (*[][]EventEnvelopeOneOf51RealmFiltersInnerInner, bool)`

GetRealmFiltersOk returns a tuple with the RealmFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmFilters

`func (o *RegisterQueue200Response) SetRealmFilters(v [][]EventEnvelopeOneOf51RealmFiltersInnerInner)`

SetRealmFilters sets RealmFilters field to given value.

### HasRealmFilters

`func (o *RegisterQueue200Response) HasRealmFilters() bool`

HasRealmFilters returns a boolean if a field has been set.

### GetRealmPlaygrounds

`func (o *RegisterQueue200Response) GetRealmPlaygrounds() []RealmPlayground`

GetRealmPlaygrounds returns the RealmPlaygrounds field if non-nil, zero value otherwise.

### GetRealmPlaygroundsOk

`func (o *RegisterQueue200Response) GetRealmPlaygroundsOk() (*[]RealmPlayground, bool)`

GetRealmPlaygroundsOk returns a tuple with the RealmPlaygrounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPlaygrounds

`func (o *RegisterQueue200Response) SetRealmPlaygrounds(v []RealmPlayground)`

SetRealmPlaygrounds sets RealmPlaygrounds field to given value.

### HasRealmPlaygrounds

`func (o *RegisterQueue200Response) HasRealmPlaygrounds() bool`

HasRealmPlaygrounds returns a boolean if a field has been set.

### GetRealmUserGroups

`func (o *RegisterQueue200Response) GetRealmUserGroups() []UserGroup`

GetRealmUserGroups returns the RealmUserGroups field if non-nil, zero value otherwise.

### GetRealmUserGroupsOk

`func (o *RegisterQueue200Response) GetRealmUserGroupsOk() (*[]UserGroup, bool)`

GetRealmUserGroupsOk returns a tuple with the RealmUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmUserGroups

`func (o *RegisterQueue200Response) SetRealmUserGroups(v []UserGroup)`

SetRealmUserGroups sets RealmUserGroups field to given value.

### HasRealmUserGroups

`func (o *RegisterQueue200Response) HasRealmUserGroups() bool`

HasRealmUserGroups returns a boolean if a field has been set.

### GetRealmBots

`func (o *RegisterQueue200Response) GetRealmBots() []Bot`

GetRealmBots returns the RealmBots field if non-nil, zero value otherwise.

### GetRealmBotsOk

`func (o *RegisterQueue200Response) GetRealmBotsOk() (*[]Bot, bool)`

GetRealmBotsOk returns a tuple with the RealmBots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmBots

`func (o *RegisterQueue200Response) SetRealmBots(v []Bot)`

SetRealmBots sets RealmBots field to given value.

### HasRealmBots

`func (o *RegisterQueue200Response) HasRealmBots() bool`

HasRealmBots returns a boolean if a field has been set.

### GetRealmEmbeddedBots

`func (o *RegisterQueue200Response) GetRealmEmbeddedBots() []RegisterQueueResponseRealmEmbeddedBotsItem`

GetRealmEmbeddedBots returns the RealmEmbeddedBots field if non-nil, zero value otherwise.

### GetRealmEmbeddedBotsOk

`func (o *RegisterQueue200Response) GetRealmEmbeddedBotsOk() (*[]RegisterQueueResponseRealmEmbeddedBotsItem, bool)`

GetRealmEmbeddedBotsOk returns a tuple with the RealmEmbeddedBots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmbeddedBots

`func (o *RegisterQueue200Response) SetRealmEmbeddedBots(v []RegisterQueueResponseRealmEmbeddedBotsItem)`

SetRealmEmbeddedBots sets RealmEmbeddedBots field to given value.

### HasRealmEmbeddedBots

`func (o *RegisterQueue200Response) HasRealmEmbeddedBots() bool`

HasRealmEmbeddedBots returns a boolean if a field has been set.

### GetRealmIncomingWebhookBots

`func (o *RegisterQueue200Response) GetRealmIncomingWebhookBots() []RegisterQueueResponseRealmIncomingWebhookBotsItem`

GetRealmIncomingWebhookBots returns the RealmIncomingWebhookBots field if non-nil, zero value otherwise.

### GetRealmIncomingWebhookBotsOk

`func (o *RegisterQueue200Response) GetRealmIncomingWebhookBotsOk() (*[]RegisterQueueResponseRealmIncomingWebhookBotsItem, bool)`

GetRealmIncomingWebhookBotsOk returns a tuple with the RealmIncomingWebhookBots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmIncomingWebhookBots

`func (o *RegisterQueue200Response) SetRealmIncomingWebhookBots(v []RegisterQueueResponseRealmIncomingWebhookBotsItem)`

SetRealmIncomingWebhookBots sets RealmIncomingWebhookBots field to given value.

### HasRealmIncomingWebhookBots

`func (o *RegisterQueue200Response) HasRealmIncomingWebhookBots() bool`

HasRealmIncomingWebhookBots returns a boolean if a field has been set.

### GetRecentPrivateConversations

`func (o *RegisterQueue200Response) GetRecentPrivateConversations() []RegisterQueueResponseRecentPrivateConversationsItem`

GetRecentPrivateConversations returns the RecentPrivateConversations field if non-nil, zero value otherwise.

### GetRecentPrivateConversationsOk

`func (o *RegisterQueue200Response) GetRecentPrivateConversationsOk() (*[]RegisterQueueResponseRecentPrivateConversationsItem, bool)`

GetRecentPrivateConversationsOk returns a tuple with the RecentPrivateConversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentPrivateConversations

`func (o *RegisterQueue200Response) SetRecentPrivateConversations(v []RegisterQueueResponseRecentPrivateConversationsItem)`

SetRecentPrivateConversations sets RecentPrivateConversations field to given value.

### HasRecentPrivateConversations

`func (o *RegisterQueue200Response) HasRecentPrivateConversations() bool`

HasRecentPrivateConversations returns a boolean if a field has been set.

### GetNavigationViews

`func (o *RegisterQueue200Response) GetNavigationViews() []NavigationView`

GetNavigationViews returns the NavigationViews field if non-nil, zero value otherwise.

### GetNavigationViewsOk

`func (o *RegisterQueue200Response) GetNavigationViewsOk() (*[]NavigationView, bool)`

GetNavigationViewsOk returns a tuple with the NavigationViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationViews

`func (o *RegisterQueue200Response) SetNavigationViews(v []NavigationView)`

SetNavigationViews sets NavigationViews field to given value.

### HasNavigationViews

`func (o *RegisterQueue200Response) HasNavigationViews() bool`

HasNavigationViews returns a boolean if a field has been set.

### GetSavedSnippets

`func (o *RegisterQueue200Response) GetSavedSnippets() []SavedSnippet`

GetSavedSnippets returns the SavedSnippets field if non-nil, zero value otherwise.

### GetSavedSnippetsOk

`func (o *RegisterQueue200Response) GetSavedSnippetsOk() (*[]SavedSnippet, bool)`

GetSavedSnippetsOk returns a tuple with the SavedSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippets

`func (o *RegisterQueue200Response) SetSavedSnippets(v []SavedSnippet)`

SetSavedSnippets sets SavedSnippets field to given value.

### HasSavedSnippets

`func (o *RegisterQueue200Response) HasSavedSnippets() bool`

HasSavedSnippets returns a boolean if a field has been set.

### GetSubscriptions

`func (o *RegisterQueue200Response) GetSubscriptions() []Subscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *RegisterQueue200Response) GetSubscriptionsOk() (*[]Subscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *RegisterQueue200Response) SetSubscriptions(v []Subscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *RegisterQueue200Response) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *RegisterQueue200Response) GetUnsubscribed() []Subscription`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *RegisterQueue200Response) GetUnsubscribedOk() (*[]Subscription, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *RegisterQueue200Response) SetUnsubscribed(v []Subscription)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *RegisterQueue200Response) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetNeverSubscribed

`func (o *RegisterQueue200Response) GetNeverSubscribed() []RegisterQueueResponseNeverSubscribedItem`

GetNeverSubscribed returns the NeverSubscribed field if non-nil, zero value otherwise.

### GetNeverSubscribedOk

`func (o *RegisterQueue200Response) GetNeverSubscribedOk() (*[]RegisterQueueResponseNeverSubscribedItem, bool)`

GetNeverSubscribedOk returns a tuple with the NeverSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverSubscribed

`func (o *RegisterQueue200Response) SetNeverSubscribed(v []RegisterQueueResponseNeverSubscribedItem)`

SetNeverSubscribed sets NeverSubscribed field to given value.

### HasNeverSubscribed

`func (o *RegisterQueue200Response) HasNeverSubscribed() bool`

HasNeverSubscribed returns a boolean if a field has been set.

### GetChannelFolders

`func (o *RegisterQueue200Response) GetChannelFolders() []ChannelFolder`

GetChannelFolders returns the ChannelFolders field if non-nil, zero value otherwise.

### GetChannelFoldersOk

`func (o *RegisterQueue200Response) GetChannelFoldersOk() (*[]ChannelFolder, bool)`

GetChannelFoldersOk returns a tuple with the ChannelFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFolders

`func (o *RegisterQueue200Response) SetChannelFolders(v []ChannelFolder)`

SetChannelFolders sets ChannelFolders field to given value.

### HasChannelFolders

`func (o *RegisterQueue200Response) HasChannelFolders() bool`

HasChannelFolders returns a boolean if a field has been set.

### GetUnreadMsgs

`func (o *RegisterQueue200Response) GetUnreadMsgs() RegisterQueueResponseUnreadMsgs`

GetUnreadMsgs returns the UnreadMsgs field if non-nil, zero value otherwise.

### GetUnreadMsgsOk

`func (o *RegisterQueue200Response) GetUnreadMsgsOk() (*RegisterQueueResponseUnreadMsgs, bool)`

GetUnreadMsgsOk returns a tuple with the UnreadMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMsgs

`func (o *RegisterQueue200Response) SetUnreadMsgs(v RegisterQueueResponseUnreadMsgs)`

SetUnreadMsgs sets UnreadMsgs field to given value.

### HasUnreadMsgs

`func (o *RegisterQueue200Response) HasUnreadMsgs() bool`

HasUnreadMsgs returns a boolean if a field has been set.

### GetStarredMessages

`func (o *RegisterQueue200Response) GetStarredMessages() []int32`

GetStarredMessages returns the StarredMessages field if non-nil, zero value otherwise.

### GetStarredMessagesOk

`func (o *RegisterQueue200Response) GetStarredMessagesOk() (*[]int32, bool)`

GetStarredMessagesOk returns a tuple with the StarredMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredMessages

`func (o *RegisterQueue200Response) SetStarredMessages(v []int32)`

SetStarredMessages sets StarredMessages field to given value.

### HasStarredMessages

`func (o *RegisterQueue200Response) HasStarredMessages() bool`

HasStarredMessages returns a boolean if a field has been set.

### GetStreams

`func (o *RegisterQueue200Response) GetStreams() []BasicChannel`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *RegisterQueue200Response) GetStreamsOk() (*[]BasicChannel, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *RegisterQueue200Response) SetStreams(v []BasicChannel)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *RegisterQueue200Response) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetRealmDefaultStreams

`func (o *RegisterQueue200Response) GetRealmDefaultStreams() []int32`

GetRealmDefaultStreams returns the RealmDefaultStreams field if non-nil, zero value otherwise.

### GetRealmDefaultStreamsOk

`func (o *RegisterQueue200Response) GetRealmDefaultStreamsOk() (*[]int32, bool)`

GetRealmDefaultStreamsOk returns a tuple with the RealmDefaultStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDefaultStreams

`func (o *RegisterQueue200Response) SetRealmDefaultStreams(v []int32)`

SetRealmDefaultStreams sets RealmDefaultStreams field to given value.

### HasRealmDefaultStreams

`func (o *RegisterQueue200Response) HasRealmDefaultStreams() bool`

HasRealmDefaultStreams returns a boolean if a field has been set.

### GetRealmDefaultStreamGroups

`func (o *RegisterQueue200Response) GetRealmDefaultStreamGroups() []DefaultChannelGroup`

GetRealmDefaultStreamGroups returns the RealmDefaultStreamGroups field if non-nil, zero value otherwise.

### GetRealmDefaultStreamGroupsOk

`func (o *RegisterQueue200Response) GetRealmDefaultStreamGroupsOk() (*[]DefaultChannelGroup, bool)`

GetRealmDefaultStreamGroupsOk returns a tuple with the RealmDefaultStreamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDefaultStreamGroups

`func (o *RegisterQueue200Response) SetRealmDefaultStreamGroups(v []DefaultChannelGroup)`

SetRealmDefaultStreamGroups sets RealmDefaultStreamGroups field to given value.

### HasRealmDefaultStreamGroups

`func (o *RegisterQueue200Response) HasRealmDefaultStreamGroups() bool`

HasRealmDefaultStreamGroups returns a boolean if a field has been set.

### GetStopWords

`func (o *RegisterQueue200Response) GetStopWords() []string`

GetStopWords returns the StopWords field if non-nil, zero value otherwise.

### GetStopWordsOk

`func (o *RegisterQueue200Response) GetStopWordsOk() (*[]string, bool)`

GetStopWordsOk returns a tuple with the StopWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopWords

`func (o *RegisterQueue200Response) SetStopWords(v []string)`

SetStopWords sets StopWords field to given value.

### HasStopWords

`func (o *RegisterQueue200Response) HasStopWords() bool`

HasStopWords returns a boolean if a field has been set.

### GetUserStatus

`func (o *RegisterQueue200Response) GetUserStatus() map[string]RegisterQueueResponseUserStatus`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *RegisterQueue200Response) GetUserStatusOk() (*map[string]RegisterQueueResponseUserStatus, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *RegisterQueue200Response) SetUserStatus(v map[string]RegisterQueueResponseUserStatus)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *RegisterQueue200Response) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetUserSettings

`func (o *RegisterQueue200Response) GetUserSettings() RegisterQueueResponseUserSettings`

GetUserSettings returns the UserSettings field if non-nil, zero value otherwise.

### GetUserSettingsOk

`func (o *RegisterQueue200Response) GetUserSettingsOk() (*RegisterQueueResponseUserSettings, bool)`

GetUserSettingsOk returns a tuple with the UserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSettings

`func (o *RegisterQueue200Response) SetUserSettings(v RegisterQueueResponseUserSettings)`

SetUserSettings sets UserSettings field to given value.

### HasUserSettings

`func (o *RegisterQueue200Response) HasUserSettings() bool`

HasUserSettings returns a boolean if a field has been set.

### GetUserTopics

`func (o *RegisterQueue200Response) GetUserTopics() []RegisterQueueResponseUserTopicsItem`

GetUserTopics returns the UserTopics field if non-nil, zero value otherwise.

### GetUserTopicsOk

`func (o *RegisterQueue200Response) GetUserTopicsOk() (*[]RegisterQueueResponseUserTopicsItem, bool)`

GetUserTopicsOk returns a tuple with the UserTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTopics

`func (o *RegisterQueue200Response) SetUserTopics(v []RegisterQueueResponseUserTopicsItem)`

SetUserTopics sets UserTopics field to given value.

### HasUserTopics

`func (o *RegisterQueue200Response) HasUserTopics() bool`

HasUserTopics returns a boolean if a field has been set.

### GetHasZoomToken

`func (o *RegisterQueue200Response) GetHasZoomToken() bool`

GetHasZoomToken returns the HasZoomToken field if non-nil, zero value otherwise.

### GetHasZoomTokenOk

`func (o *RegisterQueue200Response) GetHasZoomTokenOk() (*bool, bool)`

GetHasZoomTokenOk returns a tuple with the HasZoomToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasZoomToken

`func (o *RegisterQueue200Response) SetHasZoomToken(v bool)`

SetHasZoomToken sets HasZoomToken field to given value.

### HasHasZoomToken

`func (o *RegisterQueue200Response) HasHasZoomToken() bool`

HasHasZoomToken returns a boolean if a field has been set.

### GetGiphyApiKey

`func (o *RegisterQueue200Response) GetGiphyApiKey() string`

GetGiphyApiKey returns the GiphyApiKey field if non-nil, zero value otherwise.

### GetGiphyApiKeyOk

`func (o *RegisterQueue200Response) GetGiphyApiKeyOk() (*string, bool)`

GetGiphyApiKeyOk returns a tuple with the GiphyApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiphyApiKey

`func (o *RegisterQueue200Response) SetGiphyApiKey(v string)`

SetGiphyApiKey sets GiphyApiKey field to given value.

### HasGiphyApiKey

`func (o *RegisterQueue200Response) HasGiphyApiKey() bool`

HasGiphyApiKey returns a boolean if a field has been set.

### GetPushDevices

`func (o *RegisterQueue200Response) GetPushDevices() map[string]RegisterQueueResponsePushDevicesEntry`

GetPushDevices returns the PushDevices field if non-nil, zero value otherwise.

### GetPushDevicesOk

`func (o *RegisterQueue200Response) GetPushDevicesOk() (*map[string]RegisterQueueResponsePushDevicesEntry, bool)`

GetPushDevicesOk returns a tuple with the PushDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushDevices

`func (o *RegisterQueue200Response) SetPushDevices(v map[string]RegisterQueueResponsePushDevicesEntry)`

SetPushDevices sets PushDevices field to given value.

### HasPushDevices

`func (o *RegisterQueue200Response) HasPushDevices() bool`

HasPushDevices returns a boolean if a field has been set.

### GetEnableDesktopNotifications

`func (o *RegisterQueue200Response) GetEnableDesktopNotifications() bool`

GetEnableDesktopNotifications returns the EnableDesktopNotifications field if non-nil, zero value otherwise.

### GetEnableDesktopNotificationsOk

`func (o *RegisterQueue200Response) GetEnableDesktopNotificationsOk() (*bool, bool)`

GetEnableDesktopNotificationsOk returns a tuple with the EnableDesktopNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDesktopNotifications

`func (o *RegisterQueue200Response) SetEnableDesktopNotifications(v bool)`

SetEnableDesktopNotifications sets EnableDesktopNotifications field to given value.

### HasEnableDesktopNotifications

`func (o *RegisterQueue200Response) HasEnableDesktopNotifications() bool`

HasEnableDesktopNotifications returns a boolean if a field has been set.

### GetEnableDigestEmails

`func (o *RegisterQueue200Response) GetEnableDigestEmails() bool`

GetEnableDigestEmails returns the EnableDigestEmails field if non-nil, zero value otherwise.

### GetEnableDigestEmailsOk

`func (o *RegisterQueue200Response) GetEnableDigestEmailsOk() (*bool, bool)`

GetEnableDigestEmailsOk returns a tuple with the EnableDigestEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDigestEmails

`func (o *RegisterQueue200Response) SetEnableDigestEmails(v bool)`

SetEnableDigestEmails sets EnableDigestEmails field to given value.

### HasEnableDigestEmails

`func (o *RegisterQueue200Response) HasEnableDigestEmails() bool`

HasEnableDigestEmails returns a boolean if a field has been set.

### GetEnableLoginEmails

`func (o *RegisterQueue200Response) GetEnableLoginEmails() bool`

GetEnableLoginEmails returns the EnableLoginEmails field if non-nil, zero value otherwise.

### GetEnableLoginEmailsOk

`func (o *RegisterQueue200Response) GetEnableLoginEmailsOk() (*bool, bool)`

GetEnableLoginEmailsOk returns a tuple with the EnableLoginEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLoginEmails

`func (o *RegisterQueue200Response) SetEnableLoginEmails(v bool)`

SetEnableLoginEmails sets EnableLoginEmails field to given value.

### HasEnableLoginEmails

`func (o *RegisterQueue200Response) HasEnableLoginEmails() bool`

HasEnableLoginEmails returns a boolean if a field has been set.

### GetEnableMarketingEmails

`func (o *RegisterQueue200Response) GetEnableMarketingEmails() bool`

GetEnableMarketingEmails returns the EnableMarketingEmails field if non-nil, zero value otherwise.

### GetEnableMarketingEmailsOk

`func (o *RegisterQueue200Response) GetEnableMarketingEmailsOk() (*bool, bool)`

GetEnableMarketingEmailsOk returns a tuple with the EnableMarketingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMarketingEmails

`func (o *RegisterQueue200Response) SetEnableMarketingEmails(v bool)`

SetEnableMarketingEmails sets EnableMarketingEmails field to given value.

### HasEnableMarketingEmails

`func (o *RegisterQueue200Response) HasEnableMarketingEmails() bool`

HasEnableMarketingEmails returns a boolean if a field has been set.

### GetEmailNotificationsBatchingPeriodSeconds

`func (o *RegisterQueue200Response) GetEmailNotificationsBatchingPeriodSeconds() int32`

GetEmailNotificationsBatchingPeriodSeconds returns the EmailNotificationsBatchingPeriodSeconds field if non-nil, zero value otherwise.

### GetEmailNotificationsBatchingPeriodSecondsOk

`func (o *RegisterQueue200Response) GetEmailNotificationsBatchingPeriodSecondsOk() (*int32, bool)`

GetEmailNotificationsBatchingPeriodSecondsOk returns a tuple with the EmailNotificationsBatchingPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotificationsBatchingPeriodSeconds

`func (o *RegisterQueue200Response) SetEmailNotificationsBatchingPeriodSeconds(v int32)`

SetEmailNotificationsBatchingPeriodSeconds sets EmailNotificationsBatchingPeriodSeconds field to given value.

### HasEmailNotificationsBatchingPeriodSeconds

`func (o *RegisterQueue200Response) HasEmailNotificationsBatchingPeriodSeconds() bool`

HasEmailNotificationsBatchingPeriodSeconds returns a boolean if a field has been set.

### GetEnableOfflineEmailNotifications

`func (o *RegisterQueue200Response) GetEnableOfflineEmailNotifications() bool`

GetEnableOfflineEmailNotifications returns the EnableOfflineEmailNotifications field if non-nil, zero value otherwise.

### GetEnableOfflineEmailNotificationsOk

`func (o *RegisterQueue200Response) GetEnableOfflineEmailNotificationsOk() (*bool, bool)`

GetEnableOfflineEmailNotificationsOk returns a tuple with the EnableOfflineEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflineEmailNotifications

`func (o *RegisterQueue200Response) SetEnableOfflineEmailNotifications(v bool)`

SetEnableOfflineEmailNotifications sets EnableOfflineEmailNotifications field to given value.

### HasEnableOfflineEmailNotifications

`func (o *RegisterQueue200Response) HasEnableOfflineEmailNotifications() bool`

HasEnableOfflineEmailNotifications returns a boolean if a field has been set.

### GetEnableOfflinePushNotifications

`func (o *RegisterQueue200Response) GetEnableOfflinePushNotifications() bool`

GetEnableOfflinePushNotifications returns the EnableOfflinePushNotifications field if non-nil, zero value otherwise.

### GetEnableOfflinePushNotificationsOk

`func (o *RegisterQueue200Response) GetEnableOfflinePushNotificationsOk() (*bool, bool)`

GetEnableOfflinePushNotificationsOk returns a tuple with the EnableOfflinePushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflinePushNotifications

`func (o *RegisterQueue200Response) SetEnableOfflinePushNotifications(v bool)`

SetEnableOfflinePushNotifications sets EnableOfflinePushNotifications field to given value.

### HasEnableOfflinePushNotifications

`func (o *RegisterQueue200Response) HasEnableOfflinePushNotifications() bool`

HasEnableOfflinePushNotifications returns a boolean if a field has been set.

### GetEnableOnlinePushNotifications

`func (o *RegisterQueue200Response) GetEnableOnlinePushNotifications() bool`

GetEnableOnlinePushNotifications returns the EnableOnlinePushNotifications field if non-nil, zero value otherwise.

### GetEnableOnlinePushNotificationsOk

`func (o *RegisterQueue200Response) GetEnableOnlinePushNotificationsOk() (*bool, bool)`

GetEnableOnlinePushNotificationsOk returns a tuple with the EnableOnlinePushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOnlinePushNotifications

`func (o *RegisterQueue200Response) SetEnableOnlinePushNotifications(v bool)`

SetEnableOnlinePushNotifications sets EnableOnlinePushNotifications field to given value.

### HasEnableOnlinePushNotifications

`func (o *RegisterQueue200Response) HasEnableOnlinePushNotifications() bool`

HasEnableOnlinePushNotifications returns a boolean if a field has been set.

### GetEnableSounds

`func (o *RegisterQueue200Response) GetEnableSounds() bool`

GetEnableSounds returns the EnableSounds field if non-nil, zero value otherwise.

### GetEnableSoundsOk

`func (o *RegisterQueue200Response) GetEnableSoundsOk() (*bool, bool)`

GetEnableSoundsOk returns a tuple with the EnableSounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSounds

`func (o *RegisterQueue200Response) SetEnableSounds(v bool)`

SetEnableSounds sets EnableSounds field to given value.

### HasEnableSounds

`func (o *RegisterQueue200Response) HasEnableSounds() bool`

HasEnableSounds returns a boolean if a field has been set.

### GetEnableStreamDesktopNotifications

`func (o *RegisterQueue200Response) GetEnableStreamDesktopNotifications() bool`

GetEnableStreamDesktopNotifications returns the EnableStreamDesktopNotifications field if non-nil, zero value otherwise.

### GetEnableStreamDesktopNotificationsOk

`func (o *RegisterQueue200Response) GetEnableStreamDesktopNotificationsOk() (*bool, bool)`

GetEnableStreamDesktopNotificationsOk returns a tuple with the EnableStreamDesktopNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamDesktopNotifications

`func (o *RegisterQueue200Response) SetEnableStreamDesktopNotifications(v bool)`

SetEnableStreamDesktopNotifications sets EnableStreamDesktopNotifications field to given value.

### HasEnableStreamDesktopNotifications

`func (o *RegisterQueue200Response) HasEnableStreamDesktopNotifications() bool`

HasEnableStreamDesktopNotifications returns a boolean if a field has been set.

### GetEnableStreamEmailNotifications

`func (o *RegisterQueue200Response) GetEnableStreamEmailNotifications() bool`

GetEnableStreamEmailNotifications returns the EnableStreamEmailNotifications field if non-nil, zero value otherwise.

### GetEnableStreamEmailNotificationsOk

`func (o *RegisterQueue200Response) GetEnableStreamEmailNotificationsOk() (*bool, bool)`

GetEnableStreamEmailNotificationsOk returns a tuple with the EnableStreamEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamEmailNotifications

`func (o *RegisterQueue200Response) SetEnableStreamEmailNotifications(v bool)`

SetEnableStreamEmailNotifications sets EnableStreamEmailNotifications field to given value.

### HasEnableStreamEmailNotifications

`func (o *RegisterQueue200Response) HasEnableStreamEmailNotifications() bool`

HasEnableStreamEmailNotifications returns a boolean if a field has been set.

### GetEnableStreamPushNotifications

`func (o *RegisterQueue200Response) GetEnableStreamPushNotifications() bool`

GetEnableStreamPushNotifications returns the EnableStreamPushNotifications field if non-nil, zero value otherwise.

### GetEnableStreamPushNotificationsOk

`func (o *RegisterQueue200Response) GetEnableStreamPushNotificationsOk() (*bool, bool)`

GetEnableStreamPushNotificationsOk returns a tuple with the EnableStreamPushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamPushNotifications

`func (o *RegisterQueue200Response) SetEnableStreamPushNotifications(v bool)`

SetEnableStreamPushNotifications sets EnableStreamPushNotifications field to given value.

### HasEnableStreamPushNotifications

`func (o *RegisterQueue200Response) HasEnableStreamPushNotifications() bool`

HasEnableStreamPushNotifications returns a boolean if a field has been set.

### GetEnableStreamAudibleNotifications

`func (o *RegisterQueue200Response) GetEnableStreamAudibleNotifications() bool`

GetEnableStreamAudibleNotifications returns the EnableStreamAudibleNotifications field if non-nil, zero value otherwise.

### GetEnableStreamAudibleNotificationsOk

`func (o *RegisterQueue200Response) GetEnableStreamAudibleNotificationsOk() (*bool, bool)`

GetEnableStreamAudibleNotificationsOk returns a tuple with the EnableStreamAudibleNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamAudibleNotifications

`func (o *RegisterQueue200Response) SetEnableStreamAudibleNotifications(v bool)`

SetEnableStreamAudibleNotifications sets EnableStreamAudibleNotifications field to given value.

### HasEnableStreamAudibleNotifications

`func (o *RegisterQueue200Response) HasEnableStreamAudibleNotifications() bool`

HasEnableStreamAudibleNotifications returns a boolean if a field has been set.

### GetWildcardMentionsNotify

`func (o *RegisterQueue200Response) GetWildcardMentionsNotify() bool`

GetWildcardMentionsNotify returns the WildcardMentionsNotify field if non-nil, zero value otherwise.

### GetWildcardMentionsNotifyOk

`func (o *RegisterQueue200Response) GetWildcardMentionsNotifyOk() (*bool, bool)`

GetWildcardMentionsNotifyOk returns a tuple with the WildcardMentionsNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardMentionsNotify

`func (o *RegisterQueue200Response) SetWildcardMentionsNotify(v bool)`

SetWildcardMentionsNotify sets WildcardMentionsNotify field to given value.

### HasWildcardMentionsNotify

`func (o *RegisterQueue200Response) HasWildcardMentionsNotify() bool`

HasWildcardMentionsNotify returns a boolean if a field has been set.

### GetMessageContentInEmailNotifications

`func (o *RegisterQueue200Response) GetMessageContentInEmailNotifications() bool`

GetMessageContentInEmailNotifications returns the MessageContentInEmailNotifications field if non-nil, zero value otherwise.

### GetMessageContentInEmailNotificationsOk

`func (o *RegisterQueue200Response) GetMessageContentInEmailNotificationsOk() (*bool, bool)`

GetMessageContentInEmailNotificationsOk returns a tuple with the MessageContentInEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContentInEmailNotifications

`func (o *RegisterQueue200Response) SetMessageContentInEmailNotifications(v bool)`

SetMessageContentInEmailNotifications sets MessageContentInEmailNotifications field to given value.

### HasMessageContentInEmailNotifications

`func (o *RegisterQueue200Response) HasMessageContentInEmailNotifications() bool`

HasMessageContentInEmailNotifications returns a boolean if a field has been set.

### GetNotificationSound

`func (o *RegisterQueue200Response) GetNotificationSound() string`

GetNotificationSound returns the NotificationSound field if non-nil, zero value otherwise.

### GetNotificationSoundOk

`func (o *RegisterQueue200Response) GetNotificationSoundOk() (*string, bool)`

GetNotificationSoundOk returns a tuple with the NotificationSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSound

`func (o *RegisterQueue200Response) SetNotificationSound(v string)`

SetNotificationSound sets NotificationSound field to given value.

### HasNotificationSound

`func (o *RegisterQueue200Response) HasNotificationSound() bool`

HasNotificationSound returns a boolean if a field has been set.

### GetPmContentInDesktopNotifications

`func (o *RegisterQueue200Response) GetPmContentInDesktopNotifications() bool`

GetPmContentInDesktopNotifications returns the PmContentInDesktopNotifications field if non-nil, zero value otherwise.

### GetPmContentInDesktopNotificationsOk

`func (o *RegisterQueue200Response) GetPmContentInDesktopNotificationsOk() (*bool, bool)`

GetPmContentInDesktopNotificationsOk returns a tuple with the PmContentInDesktopNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmContentInDesktopNotifications

`func (o *RegisterQueue200Response) SetPmContentInDesktopNotifications(v bool)`

SetPmContentInDesktopNotifications sets PmContentInDesktopNotifications field to given value.

### HasPmContentInDesktopNotifications

`func (o *RegisterQueue200Response) HasPmContentInDesktopNotifications() bool`

HasPmContentInDesktopNotifications returns a boolean if a field has been set.

### GetDesktopIconCountDisplay

`func (o *RegisterQueue200Response) GetDesktopIconCountDisplay() int32`

GetDesktopIconCountDisplay returns the DesktopIconCountDisplay field if non-nil, zero value otherwise.

### GetDesktopIconCountDisplayOk

`func (o *RegisterQueue200Response) GetDesktopIconCountDisplayOk() (*int32, bool)`

GetDesktopIconCountDisplayOk returns a tuple with the DesktopIconCountDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopIconCountDisplay

`func (o *RegisterQueue200Response) SetDesktopIconCountDisplay(v int32)`

SetDesktopIconCountDisplay sets DesktopIconCountDisplay field to given value.

### HasDesktopIconCountDisplay

`func (o *RegisterQueue200Response) HasDesktopIconCountDisplay() bool`

HasDesktopIconCountDisplay returns a boolean if a field has been set.

### GetRealmNameInEmailNotificationsPolicy

`func (o *RegisterQueue200Response) GetRealmNameInEmailNotificationsPolicy() int32`

GetRealmNameInEmailNotificationsPolicy returns the RealmNameInEmailNotificationsPolicy field if non-nil, zero value otherwise.

### GetRealmNameInEmailNotificationsPolicyOk

`func (o *RegisterQueue200Response) GetRealmNameInEmailNotificationsPolicyOk() (*int32, bool)`

GetRealmNameInEmailNotificationsPolicyOk returns a tuple with the RealmNameInEmailNotificationsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmNameInEmailNotificationsPolicy

`func (o *RegisterQueue200Response) SetRealmNameInEmailNotificationsPolicy(v int32)`

SetRealmNameInEmailNotificationsPolicy sets RealmNameInEmailNotificationsPolicy field to given value.

### HasRealmNameInEmailNotificationsPolicy

`func (o *RegisterQueue200Response) HasRealmNameInEmailNotificationsPolicy() bool`

HasRealmNameInEmailNotificationsPolicy returns a boolean if a field has been set.

### GetPresenceEnabled

`func (o *RegisterQueue200Response) GetPresenceEnabled() bool`

GetPresenceEnabled returns the PresenceEnabled field if non-nil, zero value otherwise.

### GetPresenceEnabledOk

`func (o *RegisterQueue200Response) GetPresenceEnabledOk() (*bool, bool)`

GetPresenceEnabledOk returns a tuple with the PresenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceEnabled

`func (o *RegisterQueue200Response) SetPresenceEnabled(v bool)`

SetPresenceEnabled sets PresenceEnabled field to given value.

### HasPresenceEnabled

`func (o *RegisterQueue200Response) HasPresenceEnabled() bool`

HasPresenceEnabled returns a boolean if a field has been set.

### GetAvailableNotificationSounds

`func (o *RegisterQueue200Response) GetAvailableNotificationSounds() []string`

GetAvailableNotificationSounds returns the AvailableNotificationSounds field if non-nil, zero value otherwise.

### GetAvailableNotificationSoundsOk

`func (o *RegisterQueue200Response) GetAvailableNotificationSoundsOk() (*[]string, bool)`

GetAvailableNotificationSoundsOk returns a tuple with the AvailableNotificationSounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNotificationSounds

`func (o *RegisterQueue200Response) SetAvailableNotificationSounds(v []string)`

SetAvailableNotificationSounds sets AvailableNotificationSounds field to given value.

### HasAvailableNotificationSounds

`func (o *RegisterQueue200Response) HasAvailableNotificationSounds() bool`

HasAvailableNotificationSounds returns a boolean if a field has been set.

### GetColorScheme

`func (o *RegisterQueue200Response) GetColorScheme() int32`

GetColorScheme returns the ColorScheme field if non-nil, zero value otherwise.

### GetColorSchemeOk

`func (o *RegisterQueue200Response) GetColorSchemeOk() (*int32, bool)`

GetColorSchemeOk returns a tuple with the ColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorScheme

`func (o *RegisterQueue200Response) SetColorScheme(v int32)`

SetColorScheme sets ColorScheme field to given value.

### HasColorScheme

`func (o *RegisterQueue200Response) HasColorScheme() bool`

HasColorScheme returns a boolean if a field has been set.

### GetDefaultLanguage

`func (o *RegisterQueue200Response) GetDefaultLanguage() string`

GetDefaultLanguage returns the DefaultLanguage field if non-nil, zero value otherwise.

### GetDefaultLanguageOk

`func (o *RegisterQueue200Response) GetDefaultLanguageOk() (*string, bool)`

GetDefaultLanguageOk returns a tuple with the DefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguage

`func (o *RegisterQueue200Response) SetDefaultLanguage(v string)`

SetDefaultLanguage sets DefaultLanguage field to given value.

### HasDefaultLanguage

`func (o *RegisterQueue200Response) HasDefaultLanguage() bool`

HasDefaultLanguage returns a boolean if a field has been set.

### GetDemoteInactiveStreams

`func (o *RegisterQueue200Response) GetDemoteInactiveStreams() int32`

GetDemoteInactiveStreams returns the DemoteInactiveStreams field if non-nil, zero value otherwise.

### GetDemoteInactiveStreamsOk

`func (o *RegisterQueue200Response) GetDemoteInactiveStreamsOk() (*int32, bool)`

GetDemoteInactiveStreamsOk returns a tuple with the DemoteInactiveStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemoteInactiveStreams

`func (o *RegisterQueue200Response) SetDemoteInactiveStreams(v int32)`

SetDemoteInactiveStreams sets DemoteInactiveStreams field to given value.

### HasDemoteInactiveStreams

`func (o *RegisterQueue200Response) HasDemoteInactiveStreams() bool`

HasDemoteInactiveStreams returns a boolean if a field has been set.

### GetEmojiset

`func (o *RegisterQueue200Response) GetEmojiset() string`

GetEmojiset returns the Emojiset field if non-nil, zero value otherwise.

### GetEmojisetOk

`func (o *RegisterQueue200Response) GetEmojisetOk() (*string, bool)`

GetEmojisetOk returns a tuple with the Emojiset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiset

`func (o *RegisterQueue200Response) SetEmojiset(v string)`

SetEmojiset sets Emojiset field to given value.

### HasEmojiset

`func (o *RegisterQueue200Response) HasEmojiset() bool`

HasEmojiset returns a boolean if a field has been set.

### GetEnableDraftsSynchronization

`func (o *RegisterQueue200Response) GetEnableDraftsSynchronization() bool`

GetEnableDraftsSynchronization returns the EnableDraftsSynchronization field if non-nil, zero value otherwise.

### GetEnableDraftsSynchronizationOk

`func (o *RegisterQueue200Response) GetEnableDraftsSynchronizationOk() (*bool, bool)`

GetEnableDraftsSynchronizationOk returns a tuple with the EnableDraftsSynchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDraftsSynchronization

`func (o *RegisterQueue200Response) SetEnableDraftsSynchronization(v bool)`

SetEnableDraftsSynchronization sets EnableDraftsSynchronization field to given value.

### HasEnableDraftsSynchronization

`func (o *RegisterQueue200Response) HasEnableDraftsSynchronization() bool`

HasEnableDraftsSynchronization returns a boolean if a field has been set.

### GetFluidLayoutWidth

`func (o *RegisterQueue200Response) GetFluidLayoutWidth() bool`

GetFluidLayoutWidth returns the FluidLayoutWidth field if non-nil, zero value otherwise.

### GetFluidLayoutWidthOk

`func (o *RegisterQueue200Response) GetFluidLayoutWidthOk() (*bool, bool)`

GetFluidLayoutWidthOk returns a tuple with the FluidLayoutWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFluidLayoutWidth

`func (o *RegisterQueue200Response) SetFluidLayoutWidth(v bool)`

SetFluidLayoutWidth sets FluidLayoutWidth field to given value.

### HasFluidLayoutWidth

`func (o *RegisterQueue200Response) HasFluidLayoutWidth() bool`

HasFluidLayoutWidth returns a boolean if a field has been set.

### GetWebHomeView

`func (o *RegisterQueue200Response) GetWebHomeView() string`

GetWebHomeView returns the WebHomeView field if non-nil, zero value otherwise.

### GetWebHomeViewOk

`func (o *RegisterQueue200Response) GetWebHomeViewOk() (*string, bool)`

GetWebHomeViewOk returns a tuple with the WebHomeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHomeView

`func (o *RegisterQueue200Response) SetWebHomeView(v string)`

SetWebHomeView sets WebHomeView field to given value.

### HasWebHomeView

`func (o *RegisterQueue200Response) HasWebHomeView() bool`

HasWebHomeView returns a boolean if a field has been set.

### GetHighContrastMode

`func (o *RegisterQueue200Response) GetHighContrastMode() bool`

GetHighContrastMode returns the HighContrastMode field if non-nil, zero value otherwise.

### GetHighContrastModeOk

`func (o *RegisterQueue200Response) GetHighContrastModeOk() (*bool, bool)`

GetHighContrastModeOk returns a tuple with the HighContrastMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighContrastMode

`func (o *RegisterQueue200Response) SetHighContrastMode(v bool)`

SetHighContrastMode sets HighContrastMode field to given value.

### HasHighContrastMode

`func (o *RegisterQueue200Response) HasHighContrastMode() bool`

HasHighContrastMode returns a boolean if a field has been set.

### GetLeftSideUserlist

`func (o *RegisterQueue200Response) GetLeftSideUserlist() bool`

GetLeftSideUserlist returns the LeftSideUserlist field if non-nil, zero value otherwise.

### GetLeftSideUserlistOk

`func (o *RegisterQueue200Response) GetLeftSideUserlistOk() (*bool, bool)`

GetLeftSideUserlistOk returns a tuple with the LeftSideUserlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftSideUserlist

`func (o *RegisterQueue200Response) SetLeftSideUserlist(v bool)`

SetLeftSideUserlist sets LeftSideUserlist field to given value.

### HasLeftSideUserlist

`func (o *RegisterQueue200Response) HasLeftSideUserlist() bool`

HasLeftSideUserlist returns a boolean if a field has been set.

### GetStarredMessageCounts

`func (o *RegisterQueue200Response) GetStarredMessageCounts() bool`

GetStarredMessageCounts returns the StarredMessageCounts field if non-nil, zero value otherwise.

### GetStarredMessageCountsOk

`func (o *RegisterQueue200Response) GetStarredMessageCountsOk() (*bool, bool)`

GetStarredMessageCountsOk returns a tuple with the StarredMessageCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredMessageCounts

`func (o *RegisterQueue200Response) SetStarredMessageCounts(v bool)`

SetStarredMessageCounts sets StarredMessageCounts field to given value.

### HasStarredMessageCounts

`func (o *RegisterQueue200Response) HasStarredMessageCounts() bool`

HasStarredMessageCounts returns a boolean if a field has been set.

### GetTimezone

`func (o *RegisterQueue200Response) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RegisterQueue200Response) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RegisterQueue200Response) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *RegisterQueue200Response) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTranslateEmoticons

`func (o *RegisterQueue200Response) GetTranslateEmoticons() bool`

GetTranslateEmoticons returns the TranslateEmoticons field if non-nil, zero value otherwise.

### GetTranslateEmoticonsOk

`func (o *RegisterQueue200Response) GetTranslateEmoticonsOk() (*bool, bool)`

GetTranslateEmoticonsOk returns a tuple with the TranslateEmoticons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslateEmoticons

`func (o *RegisterQueue200Response) SetTranslateEmoticons(v bool)`

SetTranslateEmoticons sets TranslateEmoticons field to given value.

### HasTranslateEmoticons

`func (o *RegisterQueue200Response) HasTranslateEmoticons() bool`

HasTranslateEmoticons returns a boolean if a field has been set.

### GetTwentyFourHourTime

`func (o *RegisterQueue200Response) GetTwentyFourHourTime() bool`

GetTwentyFourHourTime returns the TwentyFourHourTime field if non-nil, zero value otherwise.

### GetTwentyFourHourTimeOk

`func (o *RegisterQueue200Response) GetTwentyFourHourTimeOk() (*bool, bool)`

GetTwentyFourHourTimeOk returns a tuple with the TwentyFourHourTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwentyFourHourTime

`func (o *RegisterQueue200Response) SetTwentyFourHourTime(v bool)`

SetTwentyFourHourTime sets TwentyFourHourTime field to given value.

### HasTwentyFourHourTime

`func (o *RegisterQueue200Response) HasTwentyFourHourTime() bool`

HasTwentyFourHourTime returns a boolean if a field has been set.

### GetReceivesTypingNotifications

`func (o *RegisterQueue200Response) GetReceivesTypingNotifications() bool`

GetReceivesTypingNotifications returns the ReceivesTypingNotifications field if non-nil, zero value otherwise.

### GetReceivesTypingNotificationsOk

`func (o *RegisterQueue200Response) GetReceivesTypingNotificationsOk() (*bool, bool)`

GetReceivesTypingNotificationsOk returns a tuple with the ReceivesTypingNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivesTypingNotifications

`func (o *RegisterQueue200Response) SetReceivesTypingNotifications(v bool)`

SetReceivesTypingNotifications sets ReceivesTypingNotifications field to given value.

### HasReceivesTypingNotifications

`func (o *RegisterQueue200Response) HasReceivesTypingNotifications() bool`

HasReceivesTypingNotifications returns a boolean if a field has been set.

### GetEnterSends

`func (o *RegisterQueue200Response) GetEnterSends() bool`

GetEnterSends returns the EnterSends field if non-nil, zero value otherwise.

### GetEnterSendsOk

`func (o *RegisterQueue200Response) GetEnterSendsOk() (*bool, bool)`

GetEnterSendsOk returns a tuple with the EnterSends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterSends

`func (o *RegisterQueue200Response) SetEnterSends(v bool)`

SetEnterSends sets EnterSends field to given value.

### HasEnterSends

`func (o *RegisterQueue200Response) HasEnterSends() bool`

HasEnterSends returns a boolean if a field has been set.

### GetEmojisetChoices

`func (o *RegisterQueue200Response) GetEmojisetChoices() []RegisterQueueResponseUserSettingsEmojisetChoicesInner`

GetEmojisetChoices returns the EmojisetChoices field if non-nil, zero value otherwise.

### GetEmojisetChoicesOk

`func (o *RegisterQueue200Response) GetEmojisetChoicesOk() (*[]RegisterQueueResponseUserSettingsEmojisetChoicesInner, bool)`

GetEmojisetChoicesOk returns a tuple with the EmojisetChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojisetChoices

`func (o *RegisterQueue200Response) SetEmojisetChoices(v []RegisterQueueResponseUserSettingsEmojisetChoicesInner)`

SetEmojisetChoices sets EmojisetChoices field to given value.

### HasEmojisetChoices

`func (o *RegisterQueue200Response) HasEmojisetChoices() bool`

HasEmojisetChoices returns a boolean if a field has been set.

### GetRealmMessageEditHistoryVisibilityPolicy

`func (o *RegisterQueue200Response) GetRealmMessageEditHistoryVisibilityPolicy() string`

GetRealmMessageEditHistoryVisibilityPolicy returns the RealmMessageEditHistoryVisibilityPolicy field if non-nil, zero value otherwise.

### GetRealmMessageEditHistoryVisibilityPolicyOk

`func (o *RegisterQueue200Response) GetRealmMessageEditHistoryVisibilityPolicyOk() (*string, bool)`

GetRealmMessageEditHistoryVisibilityPolicyOk returns a tuple with the RealmMessageEditHistoryVisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMessageEditHistoryVisibilityPolicy

`func (o *RegisterQueue200Response) SetRealmMessageEditHistoryVisibilityPolicy(v string)`

SetRealmMessageEditHistoryVisibilityPolicy sets RealmMessageEditHistoryVisibilityPolicy field to given value.

### HasRealmMessageEditHistoryVisibilityPolicy

`func (o *RegisterQueue200Response) HasRealmMessageEditHistoryVisibilityPolicy() bool`

HasRealmMessageEditHistoryVisibilityPolicy returns a boolean if a field has been set.

### GetRealmAllowEditHistory

`func (o *RegisterQueue200Response) GetRealmAllowEditHistory() bool`

GetRealmAllowEditHistory returns the RealmAllowEditHistory field if non-nil, zero value otherwise.

### GetRealmAllowEditHistoryOk

`func (o *RegisterQueue200Response) GetRealmAllowEditHistoryOk() (*bool, bool)`

GetRealmAllowEditHistoryOk returns a tuple with the RealmAllowEditHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmAllowEditHistory

`func (o *RegisterQueue200Response) SetRealmAllowEditHistory(v bool)`

SetRealmAllowEditHistory sets RealmAllowEditHistory field to given value.

### HasRealmAllowEditHistory

`func (o *RegisterQueue200Response) HasRealmAllowEditHistory() bool`

HasRealmAllowEditHistory returns a boolean if a field has been set.

### GetRealmCanAddCustomEmojiGroup

`func (o *RegisterQueue200Response) GetRealmCanAddCustomEmojiGroup() GroupSettingValue`

GetRealmCanAddCustomEmojiGroup returns the RealmCanAddCustomEmojiGroup field if non-nil, zero value otherwise.

### GetRealmCanAddCustomEmojiGroupOk

`func (o *RegisterQueue200Response) GetRealmCanAddCustomEmojiGroupOk() (*GroupSettingValue, bool)`

GetRealmCanAddCustomEmojiGroupOk returns a tuple with the RealmCanAddCustomEmojiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanAddCustomEmojiGroup

`func (o *RegisterQueue200Response) SetRealmCanAddCustomEmojiGroup(v GroupSettingValue)`

SetRealmCanAddCustomEmojiGroup sets RealmCanAddCustomEmojiGroup field to given value.

### HasRealmCanAddCustomEmojiGroup

`func (o *RegisterQueue200Response) HasRealmCanAddCustomEmojiGroup() bool`

HasRealmCanAddCustomEmojiGroup returns a boolean if a field has been set.

### GetRealmCanAddSubscribersGroup

`func (o *RegisterQueue200Response) GetRealmCanAddSubscribersGroup() GroupSettingValue`

GetRealmCanAddSubscribersGroup returns the RealmCanAddSubscribersGroup field if non-nil, zero value otherwise.

### GetRealmCanAddSubscribersGroupOk

`func (o *RegisterQueue200Response) GetRealmCanAddSubscribersGroupOk() (*GroupSettingValue, bool)`

GetRealmCanAddSubscribersGroupOk returns a tuple with the RealmCanAddSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanAddSubscribersGroup

`func (o *RegisterQueue200Response) SetRealmCanAddSubscribersGroup(v GroupSettingValue)`

SetRealmCanAddSubscribersGroup sets RealmCanAddSubscribersGroup field to given value.

### HasRealmCanAddSubscribersGroup

`func (o *RegisterQueue200Response) HasRealmCanAddSubscribersGroup() bool`

HasRealmCanAddSubscribersGroup returns a boolean if a field has been set.

### GetRealmCanDeleteAnyMessageGroup

`func (o *RegisterQueue200Response) GetRealmCanDeleteAnyMessageGroup() GroupSettingValue`

GetRealmCanDeleteAnyMessageGroup returns the RealmCanDeleteAnyMessageGroup field if non-nil, zero value otherwise.

### GetRealmCanDeleteAnyMessageGroupOk

`func (o *RegisterQueue200Response) GetRealmCanDeleteAnyMessageGroupOk() (*GroupSettingValue, bool)`

GetRealmCanDeleteAnyMessageGroupOk returns a tuple with the RealmCanDeleteAnyMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanDeleteAnyMessageGroup

`func (o *RegisterQueue200Response) SetRealmCanDeleteAnyMessageGroup(v GroupSettingValue)`

SetRealmCanDeleteAnyMessageGroup sets RealmCanDeleteAnyMessageGroup field to given value.

### HasRealmCanDeleteAnyMessageGroup

`func (o *RegisterQueue200Response) HasRealmCanDeleteAnyMessageGroup() bool`

HasRealmCanDeleteAnyMessageGroup returns a boolean if a field has been set.

### GetRealmCanDeleteOwnMessageGroup

`func (o *RegisterQueue200Response) GetRealmCanDeleteOwnMessageGroup() GroupSettingValue`

GetRealmCanDeleteOwnMessageGroup returns the RealmCanDeleteOwnMessageGroup field if non-nil, zero value otherwise.

### GetRealmCanDeleteOwnMessageGroupOk

`func (o *RegisterQueue200Response) GetRealmCanDeleteOwnMessageGroupOk() (*GroupSettingValue, bool)`

GetRealmCanDeleteOwnMessageGroupOk returns a tuple with the RealmCanDeleteOwnMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanDeleteOwnMessageGroup

`func (o *RegisterQueue200Response) SetRealmCanDeleteOwnMessageGroup(v GroupSettingValue)`

SetRealmCanDeleteOwnMessageGroup sets RealmCanDeleteOwnMessageGroup field to given value.

### HasRealmCanDeleteOwnMessageGroup

`func (o *RegisterQueue200Response) HasRealmCanDeleteOwnMessageGroup() bool`

HasRealmCanDeleteOwnMessageGroup returns a boolean if a field has been set.

### GetRealmCanSetDeleteMessagePolicyGroup

`func (o *RegisterQueue200Response) GetRealmCanSetDeleteMessagePolicyGroup() GroupSettingValue`

GetRealmCanSetDeleteMessagePolicyGroup returns the RealmCanSetDeleteMessagePolicyGroup field if non-nil, zero value otherwise.

### GetRealmCanSetDeleteMessagePolicyGroupOk

`func (o *RegisterQueue200Response) GetRealmCanSetDeleteMessagePolicyGroupOk() (*GroupSettingValue, bool)`

GetRealmCanSetDeleteMessagePolicyGroupOk returns a tuple with the RealmCanSetDeleteMessagePolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanSetDeleteMessagePolicyGroup

`func (o *RegisterQueue200Response) SetRealmCanSetDeleteMessagePolicyGroup(v GroupSettingValue)`

SetRealmCanSetDeleteMessagePolicyGroup sets RealmCanSetDeleteMessagePolicyGroup field to given value.

### HasRealmCanSetDeleteMessagePolicyGroup

`func (o *RegisterQueue200Response) HasRealmCanSetDeleteMessagePolicyGroup() bool`

HasRealmCanSetDeleteMessagePolicyGroup returns a boolean if a field has been set.

### GetRealmCanSetTopicsPolicyGroup

`func (o *RegisterQueue200Response) GetRealmCanSetTopicsPolicyGroup() GroupSettingValue`

GetRealmCanSetTopicsPolicyGroup returns the RealmCanSetTopicsPolicyGroup field if non-nil, zero value otherwise.

### GetRealmCanSetTopicsPolicyGroupOk

`func (o *RegisterQueue200Response) GetRealmCanSetTopicsPolicyGroupOk() (*GroupSettingValue, bool)`

GetRealmCanSetTopicsPolicyGroupOk returns a tuple with the RealmCanSetTopicsPolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanSetTopicsPolicyGroup

`func (o *RegisterQueue200Response) SetRealmCanSetTopicsPolicyGroup(v GroupSettingValue)`

SetRealmCanSetTopicsPolicyGroup sets RealmCanSetTopicsPolicyGroup field to given value.

### HasRealmCanSetTopicsPolicyGroup

`func (o *RegisterQueue200Response) HasRealmCanSetTopicsPolicyGroup() bool`

HasRealmCanSetTopicsPolicyGroup returns a boolean if a field has been set.

### GetRealmCanInviteUsersGroup

`func (o *RegisterQueue200Response) GetRealmCanInviteUsersGroup() GroupSettingValue`

GetRealmCanInviteUsersGroup returns the RealmCanInviteUsersGroup field if non-nil, zero value otherwise.

### GetRealmCanInviteUsersGroupOk

`func (o *RegisterQueue200Response) GetRealmCanInviteUsersGroupOk() (*GroupSettingValue, bool)`

GetRealmCanInviteUsersGroupOk returns a tuple with the RealmCanInviteUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanInviteUsersGroup

`func (o *RegisterQueue200Response) SetRealmCanInviteUsersGroup(v GroupSettingValue)`

SetRealmCanInviteUsersGroup sets RealmCanInviteUsersGroup field to given value.

### HasRealmCanInviteUsersGroup

`func (o *RegisterQueue200Response) HasRealmCanInviteUsersGroup() bool`

HasRealmCanInviteUsersGroup returns a boolean if a field has been set.

### GetRealmCanMentionManyUsersGroup

`func (o *RegisterQueue200Response) GetRealmCanMentionManyUsersGroup() GroupSettingValue`

GetRealmCanMentionManyUsersGroup returns the RealmCanMentionManyUsersGroup field if non-nil, zero value otherwise.

### GetRealmCanMentionManyUsersGroupOk

`func (o *RegisterQueue200Response) GetRealmCanMentionManyUsersGroupOk() (*GroupSettingValue, bool)`

GetRealmCanMentionManyUsersGroupOk returns a tuple with the RealmCanMentionManyUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanMentionManyUsersGroup

`func (o *RegisterQueue200Response) SetRealmCanMentionManyUsersGroup(v GroupSettingValue)`

SetRealmCanMentionManyUsersGroup sets RealmCanMentionManyUsersGroup field to given value.

### HasRealmCanMentionManyUsersGroup

`func (o *RegisterQueue200Response) HasRealmCanMentionManyUsersGroup() bool`

HasRealmCanMentionManyUsersGroup returns a boolean if a field has been set.

### GetRealmCanMoveMessagesBetweenChannelsGroup

`func (o *RegisterQueue200Response) GetRealmCanMoveMessagesBetweenChannelsGroup() GroupSettingValue`

GetRealmCanMoveMessagesBetweenChannelsGroup returns the RealmCanMoveMessagesBetweenChannelsGroup field if non-nil, zero value otherwise.

### GetRealmCanMoveMessagesBetweenChannelsGroupOk

`func (o *RegisterQueue200Response) GetRealmCanMoveMessagesBetweenChannelsGroupOk() (*GroupSettingValue, bool)`

GetRealmCanMoveMessagesBetweenChannelsGroupOk returns a tuple with the RealmCanMoveMessagesBetweenChannelsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanMoveMessagesBetweenChannelsGroup

`func (o *RegisterQueue200Response) SetRealmCanMoveMessagesBetweenChannelsGroup(v GroupSettingValue)`

SetRealmCanMoveMessagesBetweenChannelsGroup sets RealmCanMoveMessagesBetweenChannelsGroup field to given value.

### HasRealmCanMoveMessagesBetweenChannelsGroup

`func (o *RegisterQueue200Response) HasRealmCanMoveMessagesBetweenChannelsGroup() bool`

HasRealmCanMoveMessagesBetweenChannelsGroup returns a boolean if a field has been set.

### GetRealmCanMoveMessagesBetweenTopicsGroup

`func (o *RegisterQueue200Response) GetRealmCanMoveMessagesBetweenTopicsGroup() GroupSettingValue`

GetRealmCanMoveMessagesBetweenTopicsGroup returns the RealmCanMoveMessagesBetweenTopicsGroup field if non-nil, zero value otherwise.

### GetRealmCanMoveMessagesBetweenTopicsGroupOk

`func (o *RegisterQueue200Response) GetRealmCanMoveMessagesBetweenTopicsGroupOk() (*GroupSettingValue, bool)`

GetRealmCanMoveMessagesBetweenTopicsGroupOk returns a tuple with the RealmCanMoveMessagesBetweenTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanMoveMessagesBetweenTopicsGroup

`func (o *RegisterQueue200Response) SetRealmCanMoveMessagesBetweenTopicsGroup(v GroupSettingValue)`

SetRealmCanMoveMessagesBetweenTopicsGroup sets RealmCanMoveMessagesBetweenTopicsGroup field to given value.

### HasRealmCanMoveMessagesBetweenTopicsGroup

`func (o *RegisterQueue200Response) HasRealmCanMoveMessagesBetweenTopicsGroup() bool`

HasRealmCanMoveMessagesBetweenTopicsGroup returns a boolean if a field has been set.

### GetRealmCanCreateGroups

`func (o *RegisterQueue200Response) GetRealmCanCreateGroups() GroupSettingValue`

GetRealmCanCreateGroups returns the RealmCanCreateGroups field if non-nil, zero value otherwise.

### GetRealmCanCreateGroupsOk

`func (o *RegisterQueue200Response) GetRealmCanCreateGroupsOk() (*GroupSettingValue, bool)`

GetRealmCanCreateGroupsOk returns a tuple with the RealmCanCreateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanCreateGroups

`func (o *RegisterQueue200Response) SetRealmCanCreateGroups(v GroupSettingValue)`

SetRealmCanCreateGroups sets RealmCanCreateGroups field to given value.

### HasRealmCanCreateGroups

`func (o *RegisterQueue200Response) HasRealmCanCreateGroups() bool`

HasRealmCanCreateGroups returns a boolean if a field has been set.

### GetRealmCanCreateBotsGroup

`func (o *RegisterQueue200Response) GetRealmCanCreateBotsGroup() GroupSettingValue`

GetRealmCanCreateBotsGroup returns the RealmCanCreateBotsGroup field if non-nil, zero value otherwise.

### GetRealmCanCreateBotsGroupOk

`func (o *RegisterQueue200Response) GetRealmCanCreateBotsGroupOk() (*GroupSettingValue, bool)`

GetRealmCanCreateBotsGroupOk returns a tuple with the RealmCanCreateBotsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanCreateBotsGroup

`func (o *RegisterQueue200Response) SetRealmCanCreateBotsGroup(v GroupSettingValue)`

SetRealmCanCreateBotsGroup sets RealmCanCreateBotsGroup field to given value.

### HasRealmCanCreateBotsGroup

`func (o *RegisterQueue200Response) HasRealmCanCreateBotsGroup() bool`

HasRealmCanCreateBotsGroup returns a boolean if a field has been set.

### GetRealmCanCreateWriteOnlyBotsGroup

`func (o *RegisterQueue200Response) GetRealmCanCreateWriteOnlyBotsGroup() GroupSettingValue`

GetRealmCanCreateWriteOnlyBotsGroup returns the RealmCanCreateWriteOnlyBotsGroup field if non-nil, zero value otherwise.

### GetRealmCanCreateWriteOnlyBotsGroupOk

`func (o *RegisterQueue200Response) GetRealmCanCreateWriteOnlyBotsGroupOk() (*GroupSettingValue, bool)`

GetRealmCanCreateWriteOnlyBotsGroupOk returns a tuple with the RealmCanCreateWriteOnlyBotsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanCreateWriteOnlyBotsGroup

`func (o *RegisterQueue200Response) SetRealmCanCreateWriteOnlyBotsGroup(v GroupSettingValue)`

SetRealmCanCreateWriteOnlyBotsGroup sets RealmCanCreateWriteOnlyBotsGroup field to given value.

### HasRealmCanCreateWriteOnlyBotsGroup

`func (o *RegisterQueue200Response) HasRealmCanCreateWriteOnlyBotsGroup() bool`

HasRealmCanCreateWriteOnlyBotsGroup returns a boolean if a field has been set.

### GetRealmCanManageAllGroups

`func (o *RegisterQueue200Response) GetRealmCanManageAllGroups() GroupSettingValue`

GetRealmCanManageAllGroups returns the RealmCanManageAllGroups field if non-nil, zero value otherwise.

### GetRealmCanManageAllGroupsOk

`func (o *RegisterQueue200Response) GetRealmCanManageAllGroupsOk() (*GroupSettingValue, bool)`

GetRealmCanManageAllGroupsOk returns a tuple with the RealmCanManageAllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanManageAllGroups

`func (o *RegisterQueue200Response) SetRealmCanManageAllGroups(v GroupSettingValue)`

SetRealmCanManageAllGroups sets RealmCanManageAllGroups field to given value.

### HasRealmCanManageAllGroups

`func (o *RegisterQueue200Response) HasRealmCanManageAllGroups() bool`

HasRealmCanManageAllGroups returns a boolean if a field has been set.

### GetRealmCanManageBillingGroup

`func (o *RegisterQueue200Response) GetRealmCanManageBillingGroup() GroupSettingValue`

GetRealmCanManageBillingGroup returns the RealmCanManageBillingGroup field if non-nil, zero value otherwise.

### GetRealmCanManageBillingGroupOk

`func (o *RegisterQueue200Response) GetRealmCanManageBillingGroupOk() (*GroupSettingValue, bool)`

GetRealmCanManageBillingGroupOk returns a tuple with the RealmCanManageBillingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanManageBillingGroup

`func (o *RegisterQueue200Response) SetRealmCanManageBillingGroup(v GroupSettingValue)`

SetRealmCanManageBillingGroup sets RealmCanManageBillingGroup field to given value.

### HasRealmCanManageBillingGroup

`func (o *RegisterQueue200Response) HasRealmCanManageBillingGroup() bool`

HasRealmCanManageBillingGroup returns a boolean if a field has been set.

### GetRealmCanCreatePublicChannelGroup

`func (o *RegisterQueue200Response) GetRealmCanCreatePublicChannelGroup() GroupSettingValue`

GetRealmCanCreatePublicChannelGroup returns the RealmCanCreatePublicChannelGroup field if non-nil, zero value otherwise.

### GetRealmCanCreatePublicChannelGroupOk

`func (o *RegisterQueue200Response) GetRealmCanCreatePublicChannelGroupOk() (*GroupSettingValue, bool)`

GetRealmCanCreatePublicChannelGroupOk returns a tuple with the RealmCanCreatePublicChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanCreatePublicChannelGroup

`func (o *RegisterQueue200Response) SetRealmCanCreatePublicChannelGroup(v GroupSettingValue)`

SetRealmCanCreatePublicChannelGroup sets RealmCanCreatePublicChannelGroup field to given value.

### HasRealmCanCreatePublicChannelGroup

`func (o *RegisterQueue200Response) HasRealmCanCreatePublicChannelGroup() bool`

HasRealmCanCreatePublicChannelGroup returns a boolean if a field has been set.

### GetRealmCanCreatePrivateChannelGroup

`func (o *RegisterQueue200Response) GetRealmCanCreatePrivateChannelGroup() GroupSettingValue`

GetRealmCanCreatePrivateChannelGroup returns the RealmCanCreatePrivateChannelGroup field if non-nil, zero value otherwise.

### GetRealmCanCreatePrivateChannelGroupOk

`func (o *RegisterQueue200Response) GetRealmCanCreatePrivateChannelGroupOk() (*GroupSettingValue, bool)`

GetRealmCanCreatePrivateChannelGroupOk returns a tuple with the RealmCanCreatePrivateChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanCreatePrivateChannelGroup

`func (o *RegisterQueue200Response) SetRealmCanCreatePrivateChannelGroup(v GroupSettingValue)`

SetRealmCanCreatePrivateChannelGroup sets RealmCanCreatePrivateChannelGroup field to given value.

### HasRealmCanCreatePrivateChannelGroup

`func (o *RegisterQueue200Response) HasRealmCanCreatePrivateChannelGroup() bool`

HasRealmCanCreatePrivateChannelGroup returns a boolean if a field has been set.

### GetRealmCanCreateWebPublicChannelGroup

`func (o *RegisterQueue200Response) GetRealmCanCreateWebPublicChannelGroup() GroupSettingValue`

GetRealmCanCreateWebPublicChannelGroup returns the RealmCanCreateWebPublicChannelGroup field if non-nil, zero value otherwise.

### GetRealmCanCreateWebPublicChannelGroupOk

`func (o *RegisterQueue200Response) GetRealmCanCreateWebPublicChannelGroupOk() (*GroupSettingValue, bool)`

GetRealmCanCreateWebPublicChannelGroupOk returns a tuple with the RealmCanCreateWebPublicChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanCreateWebPublicChannelGroup

`func (o *RegisterQueue200Response) SetRealmCanCreateWebPublicChannelGroup(v GroupSettingValue)`

SetRealmCanCreateWebPublicChannelGroup sets RealmCanCreateWebPublicChannelGroup field to given value.

### HasRealmCanCreateWebPublicChannelGroup

`func (o *RegisterQueue200Response) HasRealmCanCreateWebPublicChannelGroup() bool`

HasRealmCanCreateWebPublicChannelGroup returns a boolean if a field has been set.

### GetRealmCanResolveTopicsGroup

`func (o *RegisterQueue200Response) GetRealmCanResolveTopicsGroup() GroupSettingValue`

GetRealmCanResolveTopicsGroup returns the RealmCanResolveTopicsGroup field if non-nil, zero value otherwise.

### GetRealmCanResolveTopicsGroupOk

`func (o *RegisterQueue200Response) GetRealmCanResolveTopicsGroupOk() (*GroupSettingValue, bool)`

GetRealmCanResolveTopicsGroupOk returns a tuple with the RealmCanResolveTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanResolveTopicsGroup

`func (o *RegisterQueue200Response) SetRealmCanResolveTopicsGroup(v GroupSettingValue)`

SetRealmCanResolveTopicsGroup sets RealmCanResolveTopicsGroup field to given value.

### HasRealmCanResolveTopicsGroup

`func (o *RegisterQueue200Response) HasRealmCanResolveTopicsGroup() bool`

HasRealmCanResolveTopicsGroup returns a boolean if a field has been set.

### GetRealmCreatePublicStreamPolicy

`func (o *RegisterQueue200Response) GetRealmCreatePublicStreamPolicy() int32`

GetRealmCreatePublicStreamPolicy returns the RealmCreatePublicStreamPolicy field if non-nil, zero value otherwise.

### GetRealmCreatePublicStreamPolicyOk

`func (o *RegisterQueue200Response) GetRealmCreatePublicStreamPolicyOk() (*int32, bool)`

GetRealmCreatePublicStreamPolicyOk returns a tuple with the RealmCreatePublicStreamPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCreatePublicStreamPolicy

`func (o *RegisterQueue200Response) SetRealmCreatePublicStreamPolicy(v int32)`

SetRealmCreatePublicStreamPolicy sets RealmCreatePublicStreamPolicy field to given value.

### HasRealmCreatePublicStreamPolicy

`func (o *RegisterQueue200Response) HasRealmCreatePublicStreamPolicy() bool`

HasRealmCreatePublicStreamPolicy returns a boolean if a field has been set.

### GetRealmCreatePrivateStreamPolicy

`func (o *RegisterQueue200Response) GetRealmCreatePrivateStreamPolicy() int32`

GetRealmCreatePrivateStreamPolicy returns the RealmCreatePrivateStreamPolicy field if non-nil, zero value otherwise.

### GetRealmCreatePrivateStreamPolicyOk

`func (o *RegisterQueue200Response) GetRealmCreatePrivateStreamPolicyOk() (*int32, bool)`

GetRealmCreatePrivateStreamPolicyOk returns a tuple with the RealmCreatePrivateStreamPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCreatePrivateStreamPolicy

`func (o *RegisterQueue200Response) SetRealmCreatePrivateStreamPolicy(v int32)`

SetRealmCreatePrivateStreamPolicy sets RealmCreatePrivateStreamPolicy field to given value.

### HasRealmCreatePrivateStreamPolicy

`func (o *RegisterQueue200Response) HasRealmCreatePrivateStreamPolicy() bool`

HasRealmCreatePrivateStreamPolicy returns a boolean if a field has been set.

### GetRealmCreateWebPublicStreamPolicy

`func (o *RegisterQueue200Response) GetRealmCreateWebPublicStreamPolicy() int32`

GetRealmCreateWebPublicStreamPolicy returns the RealmCreateWebPublicStreamPolicy field if non-nil, zero value otherwise.

### GetRealmCreateWebPublicStreamPolicyOk

`func (o *RegisterQueue200Response) GetRealmCreateWebPublicStreamPolicyOk() (*int32, bool)`

GetRealmCreateWebPublicStreamPolicyOk returns a tuple with the RealmCreateWebPublicStreamPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCreateWebPublicStreamPolicy

`func (o *RegisterQueue200Response) SetRealmCreateWebPublicStreamPolicy(v int32)`

SetRealmCreateWebPublicStreamPolicy sets RealmCreateWebPublicStreamPolicy field to given value.

### HasRealmCreateWebPublicStreamPolicy

`func (o *RegisterQueue200Response) HasRealmCreateWebPublicStreamPolicy() bool`

HasRealmCreateWebPublicStreamPolicy returns a boolean if a field has been set.

### GetRealmWildcardMentionPolicy

`func (o *RegisterQueue200Response) GetRealmWildcardMentionPolicy() int32`

GetRealmWildcardMentionPolicy returns the RealmWildcardMentionPolicy field if non-nil, zero value otherwise.

### GetRealmWildcardMentionPolicyOk

`func (o *RegisterQueue200Response) GetRealmWildcardMentionPolicyOk() (*int32, bool)`

GetRealmWildcardMentionPolicyOk returns a tuple with the RealmWildcardMentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmWildcardMentionPolicy

`func (o *RegisterQueue200Response) SetRealmWildcardMentionPolicy(v int32)`

SetRealmWildcardMentionPolicy sets RealmWildcardMentionPolicy field to given value.

### HasRealmWildcardMentionPolicy

`func (o *RegisterQueue200Response) HasRealmWildcardMentionPolicy() bool`

HasRealmWildcardMentionPolicy returns a boolean if a field has been set.

### GetRealmDefaultLanguage

`func (o *RegisterQueue200Response) GetRealmDefaultLanguage() string`

GetRealmDefaultLanguage returns the RealmDefaultLanguage field if non-nil, zero value otherwise.

### GetRealmDefaultLanguageOk

`func (o *RegisterQueue200Response) GetRealmDefaultLanguageOk() (*string, bool)`

GetRealmDefaultLanguageOk returns a tuple with the RealmDefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDefaultLanguage

`func (o *RegisterQueue200Response) SetRealmDefaultLanguage(v string)`

SetRealmDefaultLanguage sets RealmDefaultLanguage field to given value.

### HasRealmDefaultLanguage

`func (o *RegisterQueue200Response) HasRealmDefaultLanguage() bool`

HasRealmDefaultLanguage returns a boolean if a field has been set.

### GetRealmWelcomeMessageCustomText

`func (o *RegisterQueue200Response) GetRealmWelcomeMessageCustomText() string`

GetRealmWelcomeMessageCustomText returns the RealmWelcomeMessageCustomText field if non-nil, zero value otherwise.

### GetRealmWelcomeMessageCustomTextOk

`func (o *RegisterQueue200Response) GetRealmWelcomeMessageCustomTextOk() (*string, bool)`

GetRealmWelcomeMessageCustomTextOk returns a tuple with the RealmWelcomeMessageCustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmWelcomeMessageCustomText

`func (o *RegisterQueue200Response) SetRealmWelcomeMessageCustomText(v string)`

SetRealmWelcomeMessageCustomText sets RealmWelcomeMessageCustomText field to given value.

### HasRealmWelcomeMessageCustomText

`func (o *RegisterQueue200Response) HasRealmWelcomeMessageCustomText() bool`

HasRealmWelcomeMessageCustomText returns a boolean if a field has been set.

### GetRealmDescription

`func (o *RegisterQueue200Response) GetRealmDescription() string`

GetRealmDescription returns the RealmDescription field if non-nil, zero value otherwise.

### GetRealmDescriptionOk

`func (o *RegisterQueue200Response) GetRealmDescriptionOk() (*string, bool)`

GetRealmDescriptionOk returns a tuple with the RealmDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDescription

`func (o *RegisterQueue200Response) SetRealmDescription(v string)`

SetRealmDescription sets RealmDescription field to given value.

### HasRealmDescription

`func (o *RegisterQueue200Response) HasRealmDescription() bool`

HasRealmDescription returns a boolean if a field has been set.

### GetRealmDigestEmailsEnabled

`func (o *RegisterQueue200Response) GetRealmDigestEmailsEnabled() bool`

GetRealmDigestEmailsEnabled returns the RealmDigestEmailsEnabled field if non-nil, zero value otherwise.

### GetRealmDigestEmailsEnabledOk

`func (o *RegisterQueue200Response) GetRealmDigestEmailsEnabledOk() (*bool, bool)`

GetRealmDigestEmailsEnabledOk returns a tuple with the RealmDigestEmailsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDigestEmailsEnabled

`func (o *RegisterQueue200Response) SetRealmDigestEmailsEnabled(v bool)`

SetRealmDigestEmailsEnabled sets RealmDigestEmailsEnabled field to given value.

### HasRealmDigestEmailsEnabled

`func (o *RegisterQueue200Response) HasRealmDigestEmailsEnabled() bool`

HasRealmDigestEmailsEnabled returns a boolean if a field has been set.

### GetRealmDisallowDisposableEmailAddresses

`func (o *RegisterQueue200Response) GetRealmDisallowDisposableEmailAddresses() bool`

GetRealmDisallowDisposableEmailAddresses returns the RealmDisallowDisposableEmailAddresses field if non-nil, zero value otherwise.

### GetRealmDisallowDisposableEmailAddressesOk

`func (o *RegisterQueue200Response) GetRealmDisallowDisposableEmailAddressesOk() (*bool, bool)`

GetRealmDisallowDisposableEmailAddressesOk returns a tuple with the RealmDisallowDisposableEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDisallowDisposableEmailAddresses

`func (o *RegisterQueue200Response) SetRealmDisallowDisposableEmailAddresses(v bool)`

SetRealmDisallowDisposableEmailAddresses sets RealmDisallowDisposableEmailAddresses field to given value.

### HasRealmDisallowDisposableEmailAddresses

`func (o *RegisterQueue200Response) HasRealmDisallowDisposableEmailAddresses() bool`

HasRealmDisallowDisposableEmailAddresses returns a boolean if a field has been set.

### GetRealmEmailChangesDisabled

`func (o *RegisterQueue200Response) GetRealmEmailChangesDisabled() bool`

GetRealmEmailChangesDisabled returns the RealmEmailChangesDisabled field if non-nil, zero value otherwise.

### GetRealmEmailChangesDisabledOk

`func (o *RegisterQueue200Response) GetRealmEmailChangesDisabledOk() (*bool, bool)`

GetRealmEmailChangesDisabledOk returns a tuple with the RealmEmailChangesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmailChangesDisabled

`func (o *RegisterQueue200Response) SetRealmEmailChangesDisabled(v bool)`

SetRealmEmailChangesDisabled sets RealmEmailChangesDisabled field to given value.

### HasRealmEmailChangesDisabled

`func (o *RegisterQueue200Response) HasRealmEmailChangesDisabled() bool`

HasRealmEmailChangesDisabled returns a boolean if a field has been set.

### GetRealmInviteRequired

`func (o *RegisterQueue200Response) GetRealmInviteRequired() bool`

GetRealmInviteRequired returns the RealmInviteRequired field if non-nil, zero value otherwise.

### GetRealmInviteRequiredOk

`func (o *RegisterQueue200Response) GetRealmInviteRequiredOk() (*bool, bool)`

GetRealmInviteRequiredOk returns a tuple with the RealmInviteRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmInviteRequired

`func (o *RegisterQueue200Response) SetRealmInviteRequired(v bool)`

SetRealmInviteRequired sets RealmInviteRequired field to given value.

### HasRealmInviteRequired

`func (o *RegisterQueue200Response) HasRealmInviteRequired() bool`

HasRealmInviteRequired returns a boolean if a field has been set.

### GetRealmCreateMultiuseInviteGroup

`func (o *RegisterQueue200Response) GetRealmCreateMultiuseInviteGroup() GroupSettingValue`

GetRealmCreateMultiuseInviteGroup returns the RealmCreateMultiuseInviteGroup field if non-nil, zero value otherwise.

### GetRealmCreateMultiuseInviteGroupOk

`func (o *RegisterQueue200Response) GetRealmCreateMultiuseInviteGroupOk() (*GroupSettingValue, bool)`

GetRealmCreateMultiuseInviteGroupOk returns a tuple with the RealmCreateMultiuseInviteGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCreateMultiuseInviteGroup

`func (o *RegisterQueue200Response) SetRealmCreateMultiuseInviteGroup(v GroupSettingValue)`

SetRealmCreateMultiuseInviteGroup sets RealmCreateMultiuseInviteGroup field to given value.

### HasRealmCreateMultiuseInviteGroup

`func (o *RegisterQueue200Response) HasRealmCreateMultiuseInviteGroup() bool`

HasRealmCreateMultiuseInviteGroup returns a boolean if a field has been set.

### GetRealmInlineImagePreview

`func (o *RegisterQueue200Response) GetRealmInlineImagePreview() bool`

GetRealmInlineImagePreview returns the RealmInlineImagePreview field if non-nil, zero value otherwise.

### GetRealmInlineImagePreviewOk

`func (o *RegisterQueue200Response) GetRealmInlineImagePreviewOk() (*bool, bool)`

GetRealmInlineImagePreviewOk returns a tuple with the RealmInlineImagePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmInlineImagePreview

`func (o *RegisterQueue200Response) SetRealmInlineImagePreview(v bool)`

SetRealmInlineImagePreview sets RealmInlineImagePreview field to given value.

### HasRealmInlineImagePreview

`func (o *RegisterQueue200Response) HasRealmInlineImagePreview() bool`

HasRealmInlineImagePreview returns a boolean if a field has been set.

### GetRealmInlineUrlEmbedPreview

`func (o *RegisterQueue200Response) GetRealmInlineUrlEmbedPreview() bool`

GetRealmInlineUrlEmbedPreview returns the RealmInlineUrlEmbedPreview field if non-nil, zero value otherwise.

### GetRealmInlineUrlEmbedPreviewOk

`func (o *RegisterQueue200Response) GetRealmInlineUrlEmbedPreviewOk() (*bool, bool)`

GetRealmInlineUrlEmbedPreviewOk returns a tuple with the RealmInlineUrlEmbedPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmInlineUrlEmbedPreview

`func (o *RegisterQueue200Response) SetRealmInlineUrlEmbedPreview(v bool)`

SetRealmInlineUrlEmbedPreview sets RealmInlineUrlEmbedPreview field to given value.

### HasRealmInlineUrlEmbedPreview

`func (o *RegisterQueue200Response) HasRealmInlineUrlEmbedPreview() bool`

HasRealmInlineUrlEmbedPreview returns a boolean if a field has been set.

### GetRealmTopicsPolicy

`func (o *RegisterQueue200Response) GetRealmTopicsPolicy() string`

GetRealmTopicsPolicy returns the RealmTopicsPolicy field if non-nil, zero value otherwise.

### GetRealmTopicsPolicyOk

`func (o *RegisterQueue200Response) GetRealmTopicsPolicyOk() (*string, bool)`

GetRealmTopicsPolicyOk returns a tuple with the RealmTopicsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmTopicsPolicy

`func (o *RegisterQueue200Response) SetRealmTopicsPolicy(v string)`

SetRealmTopicsPolicy sets RealmTopicsPolicy field to given value.

### HasRealmTopicsPolicy

`func (o *RegisterQueue200Response) HasRealmTopicsPolicy() bool`

HasRealmTopicsPolicy returns a boolean if a field has been set.

### GetRealmMandatoryTopics

`func (o *RegisterQueue200Response) GetRealmMandatoryTopics() bool`

GetRealmMandatoryTopics returns the RealmMandatoryTopics field if non-nil, zero value otherwise.

### GetRealmMandatoryTopicsOk

`func (o *RegisterQueue200Response) GetRealmMandatoryTopicsOk() (*bool, bool)`

GetRealmMandatoryTopicsOk returns a tuple with the RealmMandatoryTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMandatoryTopics

`func (o *RegisterQueue200Response) SetRealmMandatoryTopics(v bool)`

SetRealmMandatoryTopics sets RealmMandatoryTopics field to given value.

### HasRealmMandatoryTopics

`func (o *RegisterQueue200Response) HasRealmMandatoryTopics() bool`

HasRealmMandatoryTopics returns a boolean if a field has been set.

### GetRealmMessageRetentionDays

`func (o *RegisterQueue200Response) GetRealmMessageRetentionDays() int32`

GetRealmMessageRetentionDays returns the RealmMessageRetentionDays field if non-nil, zero value otherwise.

### GetRealmMessageRetentionDaysOk

`func (o *RegisterQueue200Response) GetRealmMessageRetentionDaysOk() (*int32, bool)`

GetRealmMessageRetentionDaysOk returns a tuple with the RealmMessageRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMessageRetentionDays

`func (o *RegisterQueue200Response) SetRealmMessageRetentionDays(v int32)`

SetRealmMessageRetentionDays sets RealmMessageRetentionDays field to given value.

### HasRealmMessageRetentionDays

`func (o *RegisterQueue200Response) HasRealmMessageRetentionDays() bool`

HasRealmMessageRetentionDays returns a boolean if a field has been set.

### GetRealmName

`func (o *RegisterQueue200Response) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *RegisterQueue200Response) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *RegisterQueue200Response) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *RegisterQueue200Response) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetRealmRequireE2eePushNotifications

`func (o *RegisterQueue200Response) GetRealmRequireE2eePushNotifications() bool`

GetRealmRequireE2eePushNotifications returns the RealmRequireE2eePushNotifications field if non-nil, zero value otherwise.

### GetRealmRequireE2eePushNotificationsOk

`func (o *RegisterQueue200Response) GetRealmRequireE2eePushNotificationsOk() (*bool, bool)`

GetRealmRequireE2eePushNotificationsOk returns a tuple with the RealmRequireE2eePushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmRequireE2eePushNotifications

`func (o *RegisterQueue200Response) SetRealmRequireE2eePushNotifications(v bool)`

SetRealmRequireE2eePushNotifications sets RealmRequireE2eePushNotifications field to given value.

### HasRealmRequireE2eePushNotifications

`func (o *RegisterQueue200Response) HasRealmRequireE2eePushNotifications() bool`

HasRealmRequireE2eePushNotifications returns a boolean if a field has been set.

### GetRealmRequireUniqueNames

`func (o *RegisterQueue200Response) GetRealmRequireUniqueNames() bool`

GetRealmRequireUniqueNames returns the RealmRequireUniqueNames field if non-nil, zero value otherwise.

### GetRealmRequireUniqueNamesOk

`func (o *RegisterQueue200Response) GetRealmRequireUniqueNamesOk() (*bool, bool)`

GetRealmRequireUniqueNamesOk returns a tuple with the RealmRequireUniqueNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmRequireUniqueNames

`func (o *RegisterQueue200Response) SetRealmRequireUniqueNames(v bool)`

SetRealmRequireUniqueNames sets RealmRequireUniqueNames field to given value.

### HasRealmRequireUniqueNames

`func (o *RegisterQueue200Response) HasRealmRequireUniqueNames() bool`

HasRealmRequireUniqueNames returns a boolean if a field has been set.

### GetRealmNameChangesDisabled

`func (o *RegisterQueue200Response) GetRealmNameChangesDisabled() bool`

GetRealmNameChangesDisabled returns the RealmNameChangesDisabled field if non-nil, zero value otherwise.

### GetRealmNameChangesDisabledOk

`func (o *RegisterQueue200Response) GetRealmNameChangesDisabledOk() (*bool, bool)`

GetRealmNameChangesDisabledOk returns a tuple with the RealmNameChangesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmNameChangesDisabled

`func (o *RegisterQueue200Response) SetRealmNameChangesDisabled(v bool)`

SetRealmNameChangesDisabled sets RealmNameChangesDisabled field to given value.

### HasRealmNameChangesDisabled

`func (o *RegisterQueue200Response) HasRealmNameChangesDisabled() bool`

HasRealmNameChangesDisabled returns a boolean if a field has been set.

### GetRealmAvatarChangesDisabled

`func (o *RegisterQueue200Response) GetRealmAvatarChangesDisabled() bool`

GetRealmAvatarChangesDisabled returns the RealmAvatarChangesDisabled field if non-nil, zero value otherwise.

### GetRealmAvatarChangesDisabledOk

`func (o *RegisterQueue200Response) GetRealmAvatarChangesDisabledOk() (*bool, bool)`

GetRealmAvatarChangesDisabledOk returns a tuple with the RealmAvatarChangesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmAvatarChangesDisabled

`func (o *RegisterQueue200Response) SetRealmAvatarChangesDisabled(v bool)`

SetRealmAvatarChangesDisabled sets RealmAvatarChangesDisabled field to given value.

### HasRealmAvatarChangesDisabled

`func (o *RegisterQueue200Response) HasRealmAvatarChangesDisabled() bool`

HasRealmAvatarChangesDisabled returns a boolean if a field has been set.

### GetRealmEmailsRestrictedToDomains

`func (o *RegisterQueue200Response) GetRealmEmailsRestrictedToDomains() bool`

GetRealmEmailsRestrictedToDomains returns the RealmEmailsRestrictedToDomains field if non-nil, zero value otherwise.

### GetRealmEmailsRestrictedToDomainsOk

`func (o *RegisterQueue200Response) GetRealmEmailsRestrictedToDomainsOk() (*bool, bool)`

GetRealmEmailsRestrictedToDomainsOk returns a tuple with the RealmEmailsRestrictedToDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmailsRestrictedToDomains

`func (o *RegisterQueue200Response) SetRealmEmailsRestrictedToDomains(v bool)`

SetRealmEmailsRestrictedToDomains sets RealmEmailsRestrictedToDomains field to given value.

### HasRealmEmailsRestrictedToDomains

`func (o *RegisterQueue200Response) HasRealmEmailsRestrictedToDomains() bool`

HasRealmEmailsRestrictedToDomains returns a boolean if a field has been set.

### GetRealmSendWelcomeEmails

`func (o *RegisterQueue200Response) GetRealmSendWelcomeEmails() bool`

GetRealmSendWelcomeEmails returns the RealmSendWelcomeEmails field if non-nil, zero value otherwise.

### GetRealmSendWelcomeEmailsOk

`func (o *RegisterQueue200Response) GetRealmSendWelcomeEmailsOk() (*bool, bool)`

GetRealmSendWelcomeEmailsOk returns a tuple with the RealmSendWelcomeEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmSendWelcomeEmails

`func (o *RegisterQueue200Response) SetRealmSendWelcomeEmails(v bool)`

SetRealmSendWelcomeEmails sets RealmSendWelcomeEmails field to given value.

### HasRealmSendWelcomeEmails

`func (o *RegisterQueue200Response) HasRealmSendWelcomeEmails() bool`

HasRealmSendWelcomeEmails returns a boolean if a field has been set.

### GetRealmMessageContentAllowedInEmailNotifications

`func (o *RegisterQueue200Response) GetRealmMessageContentAllowedInEmailNotifications() bool`

GetRealmMessageContentAllowedInEmailNotifications returns the RealmMessageContentAllowedInEmailNotifications field if non-nil, zero value otherwise.

### GetRealmMessageContentAllowedInEmailNotificationsOk

`func (o *RegisterQueue200Response) GetRealmMessageContentAllowedInEmailNotificationsOk() (*bool, bool)`

GetRealmMessageContentAllowedInEmailNotificationsOk returns a tuple with the RealmMessageContentAllowedInEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMessageContentAllowedInEmailNotifications

`func (o *RegisterQueue200Response) SetRealmMessageContentAllowedInEmailNotifications(v bool)`

SetRealmMessageContentAllowedInEmailNotifications sets RealmMessageContentAllowedInEmailNotifications field to given value.

### HasRealmMessageContentAllowedInEmailNotifications

`func (o *RegisterQueue200Response) HasRealmMessageContentAllowedInEmailNotifications() bool`

HasRealmMessageContentAllowedInEmailNotifications returns a boolean if a field has been set.

### GetRealmEnableSpectatorAccess

`func (o *RegisterQueue200Response) GetRealmEnableSpectatorAccess() bool`

GetRealmEnableSpectatorAccess returns the RealmEnableSpectatorAccess field if non-nil, zero value otherwise.

### GetRealmEnableSpectatorAccessOk

`func (o *RegisterQueue200Response) GetRealmEnableSpectatorAccessOk() (*bool, bool)`

GetRealmEnableSpectatorAccessOk returns a tuple with the RealmEnableSpectatorAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEnableSpectatorAccess

`func (o *RegisterQueue200Response) SetRealmEnableSpectatorAccess(v bool)`

SetRealmEnableSpectatorAccess sets RealmEnableSpectatorAccess field to given value.

### HasRealmEnableSpectatorAccess

`func (o *RegisterQueue200Response) HasRealmEnableSpectatorAccess() bool`

HasRealmEnableSpectatorAccess returns a boolean if a field has been set.

### GetRealmWantAdvertiseInCommunitiesDirectory

`func (o *RegisterQueue200Response) GetRealmWantAdvertiseInCommunitiesDirectory() bool`

GetRealmWantAdvertiseInCommunitiesDirectory returns the RealmWantAdvertiseInCommunitiesDirectory field if non-nil, zero value otherwise.

### GetRealmWantAdvertiseInCommunitiesDirectoryOk

`func (o *RegisterQueue200Response) GetRealmWantAdvertiseInCommunitiesDirectoryOk() (*bool, bool)`

GetRealmWantAdvertiseInCommunitiesDirectoryOk returns a tuple with the RealmWantAdvertiseInCommunitiesDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmWantAdvertiseInCommunitiesDirectory

`func (o *RegisterQueue200Response) SetRealmWantAdvertiseInCommunitiesDirectory(v bool)`

SetRealmWantAdvertiseInCommunitiesDirectory sets RealmWantAdvertiseInCommunitiesDirectory field to given value.

### HasRealmWantAdvertiseInCommunitiesDirectory

`func (o *RegisterQueue200Response) HasRealmWantAdvertiseInCommunitiesDirectory() bool`

HasRealmWantAdvertiseInCommunitiesDirectory returns a boolean if a field has been set.

### GetRealmVideoChatProvider

`func (o *RegisterQueue200Response) GetRealmVideoChatProvider() int32`

GetRealmVideoChatProvider returns the RealmVideoChatProvider field if non-nil, zero value otherwise.

### GetRealmVideoChatProviderOk

`func (o *RegisterQueue200Response) GetRealmVideoChatProviderOk() (*int32, bool)`

GetRealmVideoChatProviderOk returns a tuple with the RealmVideoChatProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmVideoChatProvider

`func (o *RegisterQueue200Response) SetRealmVideoChatProvider(v int32)`

SetRealmVideoChatProvider sets RealmVideoChatProvider field to given value.

### HasRealmVideoChatProvider

`func (o *RegisterQueue200Response) HasRealmVideoChatProvider() bool`

HasRealmVideoChatProvider returns a boolean if a field has been set.

### GetRealmJitsiServerUrl

`func (o *RegisterQueue200Response) GetRealmJitsiServerUrl() string`

GetRealmJitsiServerUrl returns the RealmJitsiServerUrl field if non-nil, zero value otherwise.

### GetRealmJitsiServerUrlOk

`func (o *RegisterQueue200Response) GetRealmJitsiServerUrlOk() (*string, bool)`

GetRealmJitsiServerUrlOk returns a tuple with the RealmJitsiServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmJitsiServerUrl

`func (o *RegisterQueue200Response) SetRealmJitsiServerUrl(v string)`

SetRealmJitsiServerUrl sets RealmJitsiServerUrl field to given value.

### HasRealmJitsiServerUrl

`func (o *RegisterQueue200Response) HasRealmJitsiServerUrl() bool`

HasRealmJitsiServerUrl returns a boolean if a field has been set.

### SetRealmJitsiServerUrlNil

`func (o *RegisterQueue200Response) SetRealmJitsiServerUrlNil(b bool)`

 SetRealmJitsiServerUrlNil sets the value for RealmJitsiServerUrl to be an explicit nil

### UnsetRealmJitsiServerUrl
`func (o *RegisterQueue200Response) UnsetRealmJitsiServerUrl()`

UnsetRealmJitsiServerUrl ensures that no value is present for RealmJitsiServerUrl, not even an explicit nil
### GetRealmGiphyRating

`func (o *RegisterQueue200Response) GetRealmGiphyRating() int32`

GetRealmGiphyRating returns the RealmGiphyRating field if non-nil, zero value otherwise.

### GetRealmGiphyRatingOk

`func (o *RegisterQueue200Response) GetRealmGiphyRatingOk() (*int32, bool)`

GetRealmGiphyRatingOk returns a tuple with the RealmGiphyRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmGiphyRating

`func (o *RegisterQueue200Response) SetRealmGiphyRating(v int32)`

SetRealmGiphyRating sets RealmGiphyRating field to given value.

### HasRealmGiphyRating

`func (o *RegisterQueue200Response) HasRealmGiphyRating() bool`

HasRealmGiphyRating returns a boolean if a field has been set.

### GetRealmWaitingPeriodThreshold

`func (o *RegisterQueue200Response) GetRealmWaitingPeriodThreshold() int32`

GetRealmWaitingPeriodThreshold returns the RealmWaitingPeriodThreshold field if non-nil, zero value otherwise.

### GetRealmWaitingPeriodThresholdOk

`func (o *RegisterQueue200Response) GetRealmWaitingPeriodThresholdOk() (*int32, bool)`

GetRealmWaitingPeriodThresholdOk returns a tuple with the RealmWaitingPeriodThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmWaitingPeriodThreshold

`func (o *RegisterQueue200Response) SetRealmWaitingPeriodThreshold(v int32)`

SetRealmWaitingPeriodThreshold sets RealmWaitingPeriodThreshold field to given value.

### HasRealmWaitingPeriodThreshold

`func (o *RegisterQueue200Response) HasRealmWaitingPeriodThreshold() bool`

HasRealmWaitingPeriodThreshold returns a boolean if a field has been set.

### GetRealmDigestWeekday

`func (o *RegisterQueue200Response) GetRealmDigestWeekday() int32`

GetRealmDigestWeekday returns the RealmDigestWeekday field if non-nil, zero value otherwise.

### GetRealmDigestWeekdayOk

`func (o *RegisterQueue200Response) GetRealmDigestWeekdayOk() (*int32, bool)`

GetRealmDigestWeekdayOk returns a tuple with the RealmDigestWeekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDigestWeekday

`func (o *RegisterQueue200Response) SetRealmDigestWeekday(v int32)`

SetRealmDigestWeekday sets RealmDigestWeekday field to given value.

### HasRealmDigestWeekday

`func (o *RegisterQueue200Response) HasRealmDigestWeekday() bool`

HasRealmDigestWeekday returns a boolean if a field has been set.

### GetRealmDirectMessageInitiatorGroup

`func (o *RegisterQueue200Response) GetRealmDirectMessageInitiatorGroup() GroupSettingValue`

GetRealmDirectMessageInitiatorGroup returns the RealmDirectMessageInitiatorGroup field if non-nil, zero value otherwise.

### GetRealmDirectMessageInitiatorGroupOk

`func (o *RegisterQueue200Response) GetRealmDirectMessageInitiatorGroupOk() (*GroupSettingValue, bool)`

GetRealmDirectMessageInitiatorGroupOk returns a tuple with the RealmDirectMessageInitiatorGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDirectMessageInitiatorGroup

`func (o *RegisterQueue200Response) SetRealmDirectMessageInitiatorGroup(v GroupSettingValue)`

SetRealmDirectMessageInitiatorGroup sets RealmDirectMessageInitiatorGroup field to given value.

### HasRealmDirectMessageInitiatorGroup

`func (o *RegisterQueue200Response) HasRealmDirectMessageInitiatorGroup() bool`

HasRealmDirectMessageInitiatorGroup returns a boolean if a field has been set.

### GetRealmDirectMessagePermissionGroup

`func (o *RegisterQueue200Response) GetRealmDirectMessagePermissionGroup() GroupSettingValue`

GetRealmDirectMessagePermissionGroup returns the RealmDirectMessagePermissionGroup field if non-nil, zero value otherwise.

### GetRealmDirectMessagePermissionGroupOk

`func (o *RegisterQueue200Response) GetRealmDirectMessagePermissionGroupOk() (*GroupSettingValue, bool)`

GetRealmDirectMessagePermissionGroupOk returns a tuple with the RealmDirectMessagePermissionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDirectMessagePermissionGroup

`func (o *RegisterQueue200Response) SetRealmDirectMessagePermissionGroup(v GroupSettingValue)`

SetRealmDirectMessagePermissionGroup sets RealmDirectMessagePermissionGroup field to given value.

### HasRealmDirectMessagePermissionGroup

`func (o *RegisterQueue200Response) HasRealmDirectMessagePermissionGroup() bool`

HasRealmDirectMessagePermissionGroup returns a boolean if a field has been set.

### GetRealmDefaultCodeBlockLanguage

`func (o *RegisterQueue200Response) GetRealmDefaultCodeBlockLanguage() string`

GetRealmDefaultCodeBlockLanguage returns the RealmDefaultCodeBlockLanguage field if non-nil, zero value otherwise.

### GetRealmDefaultCodeBlockLanguageOk

`func (o *RegisterQueue200Response) GetRealmDefaultCodeBlockLanguageOk() (*string, bool)`

GetRealmDefaultCodeBlockLanguageOk returns a tuple with the RealmDefaultCodeBlockLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDefaultCodeBlockLanguage

`func (o *RegisterQueue200Response) SetRealmDefaultCodeBlockLanguage(v string)`

SetRealmDefaultCodeBlockLanguage sets RealmDefaultCodeBlockLanguage field to given value.

### HasRealmDefaultCodeBlockLanguage

`func (o *RegisterQueue200Response) HasRealmDefaultCodeBlockLanguage() bool`

HasRealmDefaultCodeBlockLanguage returns a boolean if a field has been set.

### GetRealmMessageContentDeleteLimitSeconds

`func (o *RegisterQueue200Response) GetRealmMessageContentDeleteLimitSeconds() int32`

GetRealmMessageContentDeleteLimitSeconds returns the RealmMessageContentDeleteLimitSeconds field if non-nil, zero value otherwise.

### GetRealmMessageContentDeleteLimitSecondsOk

`func (o *RegisterQueue200Response) GetRealmMessageContentDeleteLimitSecondsOk() (*int32, bool)`

GetRealmMessageContentDeleteLimitSecondsOk returns a tuple with the RealmMessageContentDeleteLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMessageContentDeleteLimitSeconds

`func (o *RegisterQueue200Response) SetRealmMessageContentDeleteLimitSeconds(v int32)`

SetRealmMessageContentDeleteLimitSeconds sets RealmMessageContentDeleteLimitSeconds field to given value.

### HasRealmMessageContentDeleteLimitSeconds

`func (o *RegisterQueue200Response) HasRealmMessageContentDeleteLimitSeconds() bool`

HasRealmMessageContentDeleteLimitSeconds returns a boolean if a field has been set.

### SetRealmMessageContentDeleteLimitSecondsNil

`func (o *RegisterQueue200Response) SetRealmMessageContentDeleteLimitSecondsNil(b bool)`

 SetRealmMessageContentDeleteLimitSecondsNil sets the value for RealmMessageContentDeleteLimitSeconds to be an explicit nil

### UnsetRealmMessageContentDeleteLimitSeconds
`func (o *RegisterQueue200Response) UnsetRealmMessageContentDeleteLimitSeconds()`

UnsetRealmMessageContentDeleteLimitSeconds ensures that no value is present for RealmMessageContentDeleteLimitSeconds, not even an explicit nil
### GetRealmAuthenticationMethods

`func (o *RegisterQueue200Response) GetRealmAuthenticationMethods() map[string]RealmAuthenticationMethod`

GetRealmAuthenticationMethods returns the RealmAuthenticationMethods field if non-nil, zero value otherwise.

### GetRealmAuthenticationMethodsOk

`func (o *RegisterQueue200Response) GetRealmAuthenticationMethodsOk() (*map[string]RealmAuthenticationMethod, bool)`

GetRealmAuthenticationMethodsOk returns a tuple with the RealmAuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmAuthenticationMethods

`func (o *RegisterQueue200Response) SetRealmAuthenticationMethods(v map[string]RealmAuthenticationMethod)`

SetRealmAuthenticationMethods sets RealmAuthenticationMethods field to given value.

### HasRealmAuthenticationMethods

`func (o *RegisterQueue200Response) HasRealmAuthenticationMethods() bool`

HasRealmAuthenticationMethods returns a boolean if a field has been set.

### GetRealmAllowMessageEditing

`func (o *RegisterQueue200Response) GetRealmAllowMessageEditing() bool`

GetRealmAllowMessageEditing returns the RealmAllowMessageEditing field if non-nil, zero value otherwise.

### GetRealmAllowMessageEditingOk

`func (o *RegisterQueue200Response) GetRealmAllowMessageEditingOk() (*bool, bool)`

GetRealmAllowMessageEditingOk returns a tuple with the RealmAllowMessageEditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmAllowMessageEditing

`func (o *RegisterQueue200Response) SetRealmAllowMessageEditing(v bool)`

SetRealmAllowMessageEditing sets RealmAllowMessageEditing field to given value.

### HasRealmAllowMessageEditing

`func (o *RegisterQueue200Response) HasRealmAllowMessageEditing() bool`

HasRealmAllowMessageEditing returns a boolean if a field has been set.

### GetRealmMessageContentEditLimitSeconds

`func (o *RegisterQueue200Response) GetRealmMessageContentEditLimitSeconds() int32`

GetRealmMessageContentEditLimitSeconds returns the RealmMessageContentEditLimitSeconds field if non-nil, zero value otherwise.

### GetRealmMessageContentEditLimitSecondsOk

`func (o *RegisterQueue200Response) GetRealmMessageContentEditLimitSecondsOk() (*int32, bool)`

GetRealmMessageContentEditLimitSecondsOk returns a tuple with the RealmMessageContentEditLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMessageContentEditLimitSeconds

`func (o *RegisterQueue200Response) SetRealmMessageContentEditLimitSeconds(v int32)`

SetRealmMessageContentEditLimitSeconds sets RealmMessageContentEditLimitSeconds field to given value.

### HasRealmMessageContentEditLimitSeconds

`func (o *RegisterQueue200Response) HasRealmMessageContentEditLimitSeconds() bool`

HasRealmMessageContentEditLimitSeconds returns a boolean if a field has been set.

### SetRealmMessageContentEditLimitSecondsNil

`func (o *RegisterQueue200Response) SetRealmMessageContentEditLimitSecondsNil(b bool)`

 SetRealmMessageContentEditLimitSecondsNil sets the value for RealmMessageContentEditLimitSeconds to be an explicit nil

### UnsetRealmMessageContentEditLimitSeconds
`func (o *RegisterQueue200Response) UnsetRealmMessageContentEditLimitSeconds()`

UnsetRealmMessageContentEditLimitSeconds ensures that no value is present for RealmMessageContentEditLimitSeconds, not even an explicit nil
### GetRealmMoveMessagesWithinStreamLimitSeconds

`func (o *RegisterQueue200Response) GetRealmMoveMessagesWithinStreamLimitSeconds() int32`

GetRealmMoveMessagesWithinStreamLimitSeconds returns the RealmMoveMessagesWithinStreamLimitSeconds field if non-nil, zero value otherwise.

### GetRealmMoveMessagesWithinStreamLimitSecondsOk

`func (o *RegisterQueue200Response) GetRealmMoveMessagesWithinStreamLimitSecondsOk() (*int32, bool)`

GetRealmMoveMessagesWithinStreamLimitSecondsOk returns a tuple with the RealmMoveMessagesWithinStreamLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMoveMessagesWithinStreamLimitSeconds

`func (o *RegisterQueue200Response) SetRealmMoveMessagesWithinStreamLimitSeconds(v int32)`

SetRealmMoveMessagesWithinStreamLimitSeconds sets RealmMoveMessagesWithinStreamLimitSeconds field to given value.

### HasRealmMoveMessagesWithinStreamLimitSeconds

`func (o *RegisterQueue200Response) HasRealmMoveMessagesWithinStreamLimitSeconds() bool`

HasRealmMoveMessagesWithinStreamLimitSeconds returns a boolean if a field has been set.

### SetRealmMoveMessagesWithinStreamLimitSecondsNil

`func (o *RegisterQueue200Response) SetRealmMoveMessagesWithinStreamLimitSecondsNil(b bool)`

 SetRealmMoveMessagesWithinStreamLimitSecondsNil sets the value for RealmMoveMessagesWithinStreamLimitSeconds to be an explicit nil

### UnsetRealmMoveMessagesWithinStreamLimitSeconds
`func (o *RegisterQueue200Response) UnsetRealmMoveMessagesWithinStreamLimitSeconds()`

UnsetRealmMoveMessagesWithinStreamLimitSeconds ensures that no value is present for RealmMoveMessagesWithinStreamLimitSeconds, not even an explicit nil
### GetRealmMoveMessagesBetweenStreamsLimitSeconds

`func (o *RegisterQueue200Response) GetRealmMoveMessagesBetweenStreamsLimitSeconds() int32`

GetRealmMoveMessagesBetweenStreamsLimitSeconds returns the RealmMoveMessagesBetweenStreamsLimitSeconds field if non-nil, zero value otherwise.

### GetRealmMoveMessagesBetweenStreamsLimitSecondsOk

`func (o *RegisterQueue200Response) GetRealmMoveMessagesBetweenStreamsLimitSecondsOk() (*int32, bool)`

GetRealmMoveMessagesBetweenStreamsLimitSecondsOk returns a tuple with the RealmMoveMessagesBetweenStreamsLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMoveMessagesBetweenStreamsLimitSeconds

`func (o *RegisterQueue200Response) SetRealmMoveMessagesBetweenStreamsLimitSeconds(v int32)`

SetRealmMoveMessagesBetweenStreamsLimitSeconds sets RealmMoveMessagesBetweenStreamsLimitSeconds field to given value.

### HasRealmMoveMessagesBetweenStreamsLimitSeconds

`func (o *RegisterQueue200Response) HasRealmMoveMessagesBetweenStreamsLimitSeconds() bool`

HasRealmMoveMessagesBetweenStreamsLimitSeconds returns a boolean if a field has been set.

### SetRealmMoveMessagesBetweenStreamsLimitSecondsNil

`func (o *RegisterQueue200Response) SetRealmMoveMessagesBetweenStreamsLimitSecondsNil(b bool)`

 SetRealmMoveMessagesBetweenStreamsLimitSecondsNil sets the value for RealmMoveMessagesBetweenStreamsLimitSeconds to be an explicit nil

### UnsetRealmMoveMessagesBetweenStreamsLimitSeconds
`func (o *RegisterQueue200Response) UnsetRealmMoveMessagesBetweenStreamsLimitSeconds()`

UnsetRealmMoveMessagesBetweenStreamsLimitSeconds ensures that no value is present for RealmMoveMessagesBetweenStreamsLimitSeconds, not even an explicit nil
### GetRealmEnableReadReceipts

`func (o *RegisterQueue200Response) GetRealmEnableReadReceipts() bool`

GetRealmEnableReadReceipts returns the RealmEnableReadReceipts field if non-nil, zero value otherwise.

### GetRealmEnableReadReceiptsOk

`func (o *RegisterQueue200Response) GetRealmEnableReadReceiptsOk() (*bool, bool)`

GetRealmEnableReadReceiptsOk returns a tuple with the RealmEnableReadReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEnableReadReceipts

`func (o *RegisterQueue200Response) SetRealmEnableReadReceipts(v bool)`

SetRealmEnableReadReceipts sets RealmEnableReadReceipts field to given value.

### HasRealmEnableReadReceipts

`func (o *RegisterQueue200Response) HasRealmEnableReadReceipts() bool`

HasRealmEnableReadReceipts returns a boolean if a field has been set.

### GetRealmIconUrl

`func (o *RegisterQueue200Response) GetRealmIconUrl() string`

GetRealmIconUrl returns the RealmIconUrl field if non-nil, zero value otherwise.

### GetRealmIconUrlOk

`func (o *RegisterQueue200Response) GetRealmIconUrlOk() (*string, bool)`

GetRealmIconUrlOk returns a tuple with the RealmIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmIconUrl

`func (o *RegisterQueue200Response) SetRealmIconUrl(v string)`

SetRealmIconUrl sets RealmIconUrl field to given value.

### HasRealmIconUrl

`func (o *RegisterQueue200Response) HasRealmIconUrl() bool`

HasRealmIconUrl returns a boolean if a field has been set.

### GetRealmIconSource

`func (o *RegisterQueue200Response) GetRealmIconSource() string`

GetRealmIconSource returns the RealmIconSource field if non-nil, zero value otherwise.

### GetRealmIconSourceOk

`func (o *RegisterQueue200Response) GetRealmIconSourceOk() (*string, bool)`

GetRealmIconSourceOk returns a tuple with the RealmIconSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmIconSource

`func (o *RegisterQueue200Response) SetRealmIconSource(v string)`

SetRealmIconSource sets RealmIconSource field to given value.

### HasRealmIconSource

`func (o *RegisterQueue200Response) HasRealmIconSource() bool`

HasRealmIconSource returns a boolean if a field has been set.

### GetMaxIconFileSizeMib

`func (o *RegisterQueue200Response) GetMaxIconFileSizeMib() int32`

GetMaxIconFileSizeMib returns the MaxIconFileSizeMib field if non-nil, zero value otherwise.

### GetMaxIconFileSizeMibOk

`func (o *RegisterQueue200Response) GetMaxIconFileSizeMibOk() (*int32, bool)`

GetMaxIconFileSizeMibOk returns a tuple with the MaxIconFileSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconFileSizeMib

`func (o *RegisterQueue200Response) SetMaxIconFileSizeMib(v int32)`

SetMaxIconFileSizeMib sets MaxIconFileSizeMib field to given value.

### HasMaxIconFileSizeMib

`func (o *RegisterQueue200Response) HasMaxIconFileSizeMib() bool`

HasMaxIconFileSizeMib returns a boolean if a field has been set.

### GetRealmLogoUrl

`func (o *RegisterQueue200Response) GetRealmLogoUrl() string`

GetRealmLogoUrl returns the RealmLogoUrl field if non-nil, zero value otherwise.

### GetRealmLogoUrlOk

`func (o *RegisterQueue200Response) GetRealmLogoUrlOk() (*string, bool)`

GetRealmLogoUrlOk returns a tuple with the RealmLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmLogoUrl

`func (o *RegisterQueue200Response) SetRealmLogoUrl(v string)`

SetRealmLogoUrl sets RealmLogoUrl field to given value.

### HasRealmLogoUrl

`func (o *RegisterQueue200Response) HasRealmLogoUrl() bool`

HasRealmLogoUrl returns a boolean if a field has been set.

### GetRealmLogoSource

`func (o *RegisterQueue200Response) GetRealmLogoSource() string`

GetRealmLogoSource returns the RealmLogoSource field if non-nil, zero value otherwise.

### GetRealmLogoSourceOk

`func (o *RegisterQueue200Response) GetRealmLogoSourceOk() (*string, bool)`

GetRealmLogoSourceOk returns a tuple with the RealmLogoSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmLogoSource

`func (o *RegisterQueue200Response) SetRealmLogoSource(v string)`

SetRealmLogoSource sets RealmLogoSource field to given value.

### HasRealmLogoSource

`func (o *RegisterQueue200Response) HasRealmLogoSource() bool`

HasRealmLogoSource returns a boolean if a field has been set.

### GetRealmNightLogoUrl

`func (o *RegisterQueue200Response) GetRealmNightLogoUrl() string`

GetRealmNightLogoUrl returns the RealmNightLogoUrl field if non-nil, zero value otherwise.

### GetRealmNightLogoUrlOk

`func (o *RegisterQueue200Response) GetRealmNightLogoUrlOk() (*string, bool)`

GetRealmNightLogoUrlOk returns a tuple with the RealmNightLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmNightLogoUrl

`func (o *RegisterQueue200Response) SetRealmNightLogoUrl(v string)`

SetRealmNightLogoUrl sets RealmNightLogoUrl field to given value.

### HasRealmNightLogoUrl

`func (o *RegisterQueue200Response) HasRealmNightLogoUrl() bool`

HasRealmNightLogoUrl returns a boolean if a field has been set.

### GetRealmNightLogoSource

`func (o *RegisterQueue200Response) GetRealmNightLogoSource() string`

GetRealmNightLogoSource returns the RealmNightLogoSource field if non-nil, zero value otherwise.

### GetRealmNightLogoSourceOk

`func (o *RegisterQueue200Response) GetRealmNightLogoSourceOk() (*string, bool)`

GetRealmNightLogoSourceOk returns a tuple with the RealmNightLogoSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmNightLogoSource

`func (o *RegisterQueue200Response) SetRealmNightLogoSource(v string)`

SetRealmNightLogoSource sets RealmNightLogoSource field to given value.

### HasRealmNightLogoSource

`func (o *RegisterQueue200Response) HasRealmNightLogoSource() bool`

HasRealmNightLogoSource returns a boolean if a field has been set.

### GetMaxLogoFileSizeMib

`func (o *RegisterQueue200Response) GetMaxLogoFileSizeMib() int32`

GetMaxLogoFileSizeMib returns the MaxLogoFileSizeMib field if non-nil, zero value otherwise.

### GetMaxLogoFileSizeMibOk

`func (o *RegisterQueue200Response) GetMaxLogoFileSizeMibOk() (*int32, bool)`

GetMaxLogoFileSizeMibOk returns a tuple with the MaxLogoFileSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLogoFileSizeMib

`func (o *RegisterQueue200Response) SetMaxLogoFileSizeMib(v int32)`

SetMaxLogoFileSizeMib sets MaxLogoFileSizeMib field to given value.

### HasMaxLogoFileSizeMib

`func (o *RegisterQueue200Response) HasMaxLogoFileSizeMib() bool`

HasMaxLogoFileSizeMib returns a boolean if a field has been set.

### GetRealmBotDomain

`func (o *RegisterQueue200Response) GetRealmBotDomain() string`

GetRealmBotDomain returns the RealmBotDomain field if non-nil, zero value otherwise.

### GetRealmBotDomainOk

`func (o *RegisterQueue200Response) GetRealmBotDomainOk() (*string, bool)`

GetRealmBotDomainOk returns a tuple with the RealmBotDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmBotDomain

`func (o *RegisterQueue200Response) SetRealmBotDomain(v string)`

SetRealmBotDomain sets RealmBotDomain field to given value.

### HasRealmBotDomain

`func (o *RegisterQueue200Response) HasRealmBotDomain() bool`

HasRealmBotDomain returns a boolean if a field has been set.

### GetRealmUri

`func (o *RegisterQueue200Response) GetRealmUri() string`

GetRealmUri returns the RealmUri field if non-nil, zero value otherwise.

### GetRealmUriOk

`func (o *RegisterQueue200Response) GetRealmUriOk() (*string, bool)`

GetRealmUriOk returns a tuple with the RealmUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmUri

`func (o *RegisterQueue200Response) SetRealmUri(v string)`

SetRealmUri sets RealmUri field to given value.

### HasRealmUri

`func (o *RegisterQueue200Response) HasRealmUri() bool`

HasRealmUri returns a boolean if a field has been set.

### GetRealmUrl

`func (o *RegisterQueue200Response) GetRealmUrl() string`

GetRealmUrl returns the RealmUrl field if non-nil, zero value otherwise.

### GetRealmUrlOk

`func (o *RegisterQueue200Response) GetRealmUrlOk() (*string, bool)`

GetRealmUrlOk returns a tuple with the RealmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmUrl

`func (o *RegisterQueue200Response) SetRealmUrl(v string)`

SetRealmUrl sets RealmUrl field to given value.

### HasRealmUrl

`func (o *RegisterQueue200Response) HasRealmUrl() bool`

HasRealmUrl returns a boolean if a field has been set.

### GetRealmAvailableVideoChatProviders

`func (o *RegisterQueue200Response) GetRealmAvailableVideoChatProviders() map[string]VideoChatProviderEntry`

GetRealmAvailableVideoChatProviders returns the RealmAvailableVideoChatProviders field if non-nil, zero value otherwise.

### GetRealmAvailableVideoChatProvidersOk

`func (o *RegisterQueue200Response) GetRealmAvailableVideoChatProvidersOk() (*map[string]VideoChatProviderEntry, bool)`

GetRealmAvailableVideoChatProvidersOk returns a tuple with the RealmAvailableVideoChatProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmAvailableVideoChatProviders

`func (o *RegisterQueue200Response) SetRealmAvailableVideoChatProviders(v map[string]VideoChatProviderEntry)`

SetRealmAvailableVideoChatProviders sets RealmAvailableVideoChatProviders field to given value.

### HasRealmAvailableVideoChatProviders

`func (o *RegisterQueue200Response) HasRealmAvailableVideoChatProviders() bool`

HasRealmAvailableVideoChatProviders returns a boolean if a field has been set.

### GetRealmPresenceDisabled

`func (o *RegisterQueue200Response) GetRealmPresenceDisabled() bool`

GetRealmPresenceDisabled returns the RealmPresenceDisabled field if non-nil, zero value otherwise.

### GetRealmPresenceDisabledOk

`func (o *RegisterQueue200Response) GetRealmPresenceDisabledOk() (*bool, bool)`

GetRealmPresenceDisabledOk returns a tuple with the RealmPresenceDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPresenceDisabled

`func (o *RegisterQueue200Response) SetRealmPresenceDisabled(v bool)`

SetRealmPresenceDisabled sets RealmPresenceDisabled field to given value.

### HasRealmPresenceDisabled

`func (o *RegisterQueue200Response) HasRealmPresenceDisabled() bool`

HasRealmPresenceDisabled returns a boolean if a field has been set.

### GetSettingsSendDigestEmails

`func (o *RegisterQueue200Response) GetSettingsSendDigestEmails() bool`

GetSettingsSendDigestEmails returns the SettingsSendDigestEmails field if non-nil, zero value otherwise.

### GetSettingsSendDigestEmailsOk

`func (o *RegisterQueue200Response) GetSettingsSendDigestEmailsOk() (*bool, bool)`

GetSettingsSendDigestEmailsOk returns a tuple with the SettingsSendDigestEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsSendDigestEmails

`func (o *RegisterQueue200Response) SetSettingsSendDigestEmails(v bool)`

SetSettingsSendDigestEmails sets SettingsSendDigestEmails field to given value.

### HasSettingsSendDigestEmails

`func (o *RegisterQueue200Response) HasSettingsSendDigestEmails() bool`

HasSettingsSendDigestEmails returns a boolean if a field has been set.

### GetRealmIsZephyrMirrorRealm

`func (o *RegisterQueue200Response) GetRealmIsZephyrMirrorRealm() bool`

GetRealmIsZephyrMirrorRealm returns the RealmIsZephyrMirrorRealm field if non-nil, zero value otherwise.

### GetRealmIsZephyrMirrorRealmOk

`func (o *RegisterQueue200Response) GetRealmIsZephyrMirrorRealmOk() (*bool, bool)`

GetRealmIsZephyrMirrorRealmOk returns a tuple with the RealmIsZephyrMirrorRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmIsZephyrMirrorRealm

`func (o *RegisterQueue200Response) SetRealmIsZephyrMirrorRealm(v bool)`

SetRealmIsZephyrMirrorRealm sets RealmIsZephyrMirrorRealm field to given value.

### HasRealmIsZephyrMirrorRealm

`func (o *RegisterQueue200Response) HasRealmIsZephyrMirrorRealm() bool`

HasRealmIsZephyrMirrorRealm returns a boolean if a field has been set.

### GetRealmEmailAuthEnabled

`func (o *RegisterQueue200Response) GetRealmEmailAuthEnabled() bool`

GetRealmEmailAuthEnabled returns the RealmEmailAuthEnabled field if non-nil, zero value otherwise.

### GetRealmEmailAuthEnabledOk

`func (o *RegisterQueue200Response) GetRealmEmailAuthEnabledOk() (*bool, bool)`

GetRealmEmailAuthEnabledOk returns a tuple with the RealmEmailAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmailAuthEnabled

`func (o *RegisterQueue200Response) SetRealmEmailAuthEnabled(v bool)`

SetRealmEmailAuthEnabled sets RealmEmailAuthEnabled field to given value.

### HasRealmEmailAuthEnabled

`func (o *RegisterQueue200Response) HasRealmEmailAuthEnabled() bool`

HasRealmEmailAuthEnabled returns a boolean if a field has been set.

### GetRealmPasswordAuthEnabled

`func (o *RegisterQueue200Response) GetRealmPasswordAuthEnabled() bool`

GetRealmPasswordAuthEnabled returns the RealmPasswordAuthEnabled field if non-nil, zero value otherwise.

### GetRealmPasswordAuthEnabledOk

`func (o *RegisterQueue200Response) GetRealmPasswordAuthEnabledOk() (*bool, bool)`

GetRealmPasswordAuthEnabledOk returns a tuple with the RealmPasswordAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPasswordAuthEnabled

`func (o *RegisterQueue200Response) SetRealmPasswordAuthEnabled(v bool)`

SetRealmPasswordAuthEnabled sets RealmPasswordAuthEnabled field to given value.

### HasRealmPasswordAuthEnabled

`func (o *RegisterQueue200Response) HasRealmPasswordAuthEnabled() bool`

HasRealmPasswordAuthEnabled returns a boolean if a field has been set.

### GetRealmPushNotificationsEnabled

`func (o *RegisterQueue200Response) GetRealmPushNotificationsEnabled() bool`

GetRealmPushNotificationsEnabled returns the RealmPushNotificationsEnabled field if non-nil, zero value otherwise.

### GetRealmPushNotificationsEnabledOk

`func (o *RegisterQueue200Response) GetRealmPushNotificationsEnabledOk() (*bool, bool)`

GetRealmPushNotificationsEnabledOk returns a tuple with the RealmPushNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPushNotificationsEnabled

`func (o *RegisterQueue200Response) SetRealmPushNotificationsEnabled(v bool)`

SetRealmPushNotificationsEnabled sets RealmPushNotificationsEnabled field to given value.

### HasRealmPushNotificationsEnabled

`func (o *RegisterQueue200Response) HasRealmPushNotificationsEnabled() bool`

HasRealmPushNotificationsEnabled returns a boolean if a field has been set.

### GetRealmPushNotificationsEnabledEndTimestamp

`func (o *RegisterQueue200Response) GetRealmPushNotificationsEnabledEndTimestamp() int32`

GetRealmPushNotificationsEnabledEndTimestamp returns the RealmPushNotificationsEnabledEndTimestamp field if non-nil, zero value otherwise.

### GetRealmPushNotificationsEnabledEndTimestampOk

`func (o *RegisterQueue200Response) GetRealmPushNotificationsEnabledEndTimestampOk() (*int32, bool)`

GetRealmPushNotificationsEnabledEndTimestampOk returns a tuple with the RealmPushNotificationsEnabledEndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPushNotificationsEnabledEndTimestamp

`func (o *RegisterQueue200Response) SetRealmPushNotificationsEnabledEndTimestamp(v int32)`

SetRealmPushNotificationsEnabledEndTimestamp sets RealmPushNotificationsEnabledEndTimestamp field to given value.

### HasRealmPushNotificationsEnabledEndTimestamp

`func (o *RegisterQueue200Response) HasRealmPushNotificationsEnabledEndTimestamp() bool`

HasRealmPushNotificationsEnabledEndTimestamp returns a boolean if a field has been set.

### SetRealmPushNotificationsEnabledEndTimestampNil

`func (o *RegisterQueue200Response) SetRealmPushNotificationsEnabledEndTimestampNil(b bool)`

 SetRealmPushNotificationsEnabledEndTimestampNil sets the value for RealmPushNotificationsEnabledEndTimestamp to be an explicit nil

### UnsetRealmPushNotificationsEnabledEndTimestamp
`func (o *RegisterQueue200Response) UnsetRealmPushNotificationsEnabledEndTimestamp()`

UnsetRealmPushNotificationsEnabledEndTimestamp ensures that no value is present for RealmPushNotificationsEnabledEndTimestamp, not even an explicit nil
### GetRealmUploadQuotaMib

`func (o *RegisterQueue200Response) GetRealmUploadQuotaMib() int32`

GetRealmUploadQuotaMib returns the RealmUploadQuotaMib field if non-nil, zero value otherwise.

### GetRealmUploadQuotaMibOk

`func (o *RegisterQueue200Response) GetRealmUploadQuotaMibOk() (*int32, bool)`

GetRealmUploadQuotaMibOk returns a tuple with the RealmUploadQuotaMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmUploadQuotaMib

`func (o *RegisterQueue200Response) SetRealmUploadQuotaMib(v int32)`

SetRealmUploadQuotaMib sets RealmUploadQuotaMib field to given value.

### HasRealmUploadQuotaMib

`func (o *RegisterQueue200Response) HasRealmUploadQuotaMib() bool`

HasRealmUploadQuotaMib returns a boolean if a field has been set.

### SetRealmUploadQuotaMibNil

`func (o *RegisterQueue200Response) SetRealmUploadQuotaMibNil(b bool)`

 SetRealmUploadQuotaMibNil sets the value for RealmUploadQuotaMib to be an explicit nil

### UnsetRealmUploadQuotaMib
`func (o *RegisterQueue200Response) UnsetRealmUploadQuotaMib()`

UnsetRealmUploadQuotaMib ensures that no value is present for RealmUploadQuotaMib, not even an explicit nil
### GetRealmOrgType

`func (o *RegisterQueue200Response) GetRealmOrgType() int32`

GetRealmOrgType returns the RealmOrgType field if non-nil, zero value otherwise.

### GetRealmOrgTypeOk

`func (o *RegisterQueue200Response) GetRealmOrgTypeOk() (*int32, bool)`

GetRealmOrgTypeOk returns a tuple with the RealmOrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmOrgType

`func (o *RegisterQueue200Response) SetRealmOrgType(v int32)`

SetRealmOrgType sets RealmOrgType field to given value.

### HasRealmOrgType

`func (o *RegisterQueue200Response) HasRealmOrgType() bool`

HasRealmOrgType returns a boolean if a field has been set.

### GetRealmPlanType

`func (o *RegisterQueue200Response) GetRealmPlanType() int32`

GetRealmPlanType returns the RealmPlanType field if non-nil, zero value otherwise.

### GetRealmPlanTypeOk

`func (o *RegisterQueue200Response) GetRealmPlanTypeOk() (*int32, bool)`

GetRealmPlanTypeOk returns a tuple with the RealmPlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPlanType

`func (o *RegisterQueue200Response) SetRealmPlanType(v int32)`

SetRealmPlanType sets RealmPlanType field to given value.

### HasRealmPlanType

`func (o *RegisterQueue200Response) HasRealmPlanType() bool`

HasRealmPlanType returns a boolean if a field has been set.

### GetRealmEnableGuestUserDmWarning

`func (o *RegisterQueue200Response) GetRealmEnableGuestUserDmWarning() bool`

GetRealmEnableGuestUserDmWarning returns the RealmEnableGuestUserDmWarning field if non-nil, zero value otherwise.

### GetRealmEnableGuestUserDmWarningOk

`func (o *RegisterQueue200Response) GetRealmEnableGuestUserDmWarningOk() (*bool, bool)`

GetRealmEnableGuestUserDmWarningOk returns a tuple with the RealmEnableGuestUserDmWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEnableGuestUserDmWarning

`func (o *RegisterQueue200Response) SetRealmEnableGuestUserDmWarning(v bool)`

SetRealmEnableGuestUserDmWarning sets RealmEnableGuestUserDmWarning field to given value.

### HasRealmEnableGuestUserDmWarning

`func (o *RegisterQueue200Response) HasRealmEnableGuestUserDmWarning() bool`

HasRealmEnableGuestUserDmWarning returns a boolean if a field has been set.

### GetRealmEnableGuestUserIndicator

`func (o *RegisterQueue200Response) GetRealmEnableGuestUserIndicator() bool`

GetRealmEnableGuestUserIndicator returns the RealmEnableGuestUserIndicator field if non-nil, zero value otherwise.

### GetRealmEnableGuestUserIndicatorOk

`func (o *RegisterQueue200Response) GetRealmEnableGuestUserIndicatorOk() (*bool, bool)`

GetRealmEnableGuestUserIndicatorOk returns a tuple with the RealmEnableGuestUserIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEnableGuestUserIndicator

`func (o *RegisterQueue200Response) SetRealmEnableGuestUserIndicator(v bool)`

SetRealmEnableGuestUserIndicator sets RealmEnableGuestUserIndicator field to given value.

### HasRealmEnableGuestUserIndicator

`func (o *RegisterQueue200Response) HasRealmEnableGuestUserIndicator() bool`

HasRealmEnableGuestUserIndicator returns a boolean if a field has been set.

### GetRealmCanAccessAllUsersGroup

`func (o *RegisterQueue200Response) GetRealmCanAccessAllUsersGroup() GroupSettingValue`

GetRealmCanAccessAllUsersGroup returns the RealmCanAccessAllUsersGroup field if non-nil, zero value otherwise.

### GetRealmCanAccessAllUsersGroupOk

`func (o *RegisterQueue200Response) GetRealmCanAccessAllUsersGroupOk() (*GroupSettingValue, bool)`

GetRealmCanAccessAllUsersGroupOk returns a tuple with the RealmCanAccessAllUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanAccessAllUsersGroup

`func (o *RegisterQueue200Response) SetRealmCanAccessAllUsersGroup(v GroupSettingValue)`

SetRealmCanAccessAllUsersGroup sets RealmCanAccessAllUsersGroup field to given value.

### HasRealmCanAccessAllUsersGroup

`func (o *RegisterQueue200Response) HasRealmCanAccessAllUsersGroup() bool`

HasRealmCanAccessAllUsersGroup returns a boolean if a field has been set.

### GetRealmCanSummarizeTopicsGroup

`func (o *RegisterQueue200Response) GetRealmCanSummarizeTopicsGroup() GroupSettingValue`

GetRealmCanSummarizeTopicsGroup returns the RealmCanSummarizeTopicsGroup field if non-nil, zero value otherwise.

### GetRealmCanSummarizeTopicsGroupOk

`func (o *RegisterQueue200Response) GetRealmCanSummarizeTopicsGroupOk() (*GroupSettingValue, bool)`

GetRealmCanSummarizeTopicsGroupOk returns a tuple with the RealmCanSummarizeTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCanSummarizeTopicsGroup

`func (o *RegisterQueue200Response) SetRealmCanSummarizeTopicsGroup(v GroupSettingValue)`

SetRealmCanSummarizeTopicsGroup sets RealmCanSummarizeTopicsGroup field to given value.

### HasRealmCanSummarizeTopicsGroup

`func (o *RegisterQueue200Response) HasRealmCanSummarizeTopicsGroup() bool`

HasRealmCanSummarizeTopicsGroup returns a boolean if a field has been set.

### GetZulipPlanIsNotLimited

`func (o *RegisterQueue200Response) GetZulipPlanIsNotLimited() bool`

GetZulipPlanIsNotLimited returns the ZulipPlanIsNotLimited field if non-nil, zero value otherwise.

### GetZulipPlanIsNotLimitedOk

`func (o *RegisterQueue200Response) GetZulipPlanIsNotLimitedOk() (*bool, bool)`

GetZulipPlanIsNotLimitedOk returns a tuple with the ZulipPlanIsNotLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipPlanIsNotLimited

`func (o *RegisterQueue200Response) SetZulipPlanIsNotLimited(v bool)`

SetZulipPlanIsNotLimited sets ZulipPlanIsNotLimited field to given value.

### HasZulipPlanIsNotLimited

`func (o *RegisterQueue200Response) HasZulipPlanIsNotLimited() bool`

HasZulipPlanIsNotLimited returns a boolean if a field has been set.

### GetUpgradeTextForWideOrganizationLogo

`func (o *RegisterQueue200Response) GetUpgradeTextForWideOrganizationLogo() string`

GetUpgradeTextForWideOrganizationLogo returns the UpgradeTextForWideOrganizationLogo field if non-nil, zero value otherwise.

### GetUpgradeTextForWideOrganizationLogoOk

`func (o *RegisterQueue200Response) GetUpgradeTextForWideOrganizationLogoOk() (*string, bool)`

GetUpgradeTextForWideOrganizationLogoOk returns a tuple with the UpgradeTextForWideOrganizationLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTextForWideOrganizationLogo

`func (o *RegisterQueue200Response) SetUpgradeTextForWideOrganizationLogo(v string)`

SetUpgradeTextForWideOrganizationLogo sets UpgradeTextForWideOrganizationLogo field to given value.

### HasUpgradeTextForWideOrganizationLogo

`func (o *RegisterQueue200Response) HasUpgradeTextForWideOrganizationLogo() bool`

HasUpgradeTextForWideOrganizationLogo returns a boolean if a field has been set.

### GetRealmDefaultExternalAccounts

`func (o *RegisterQueue200Response) GetRealmDefaultExternalAccounts() map[string]RegisterQueueResponseRealmDefaultExternalAccountsEntry`

GetRealmDefaultExternalAccounts returns the RealmDefaultExternalAccounts field if non-nil, zero value otherwise.

### GetRealmDefaultExternalAccountsOk

`func (o *RegisterQueue200Response) GetRealmDefaultExternalAccountsOk() (*map[string]RegisterQueueResponseRealmDefaultExternalAccountsEntry, bool)`

GetRealmDefaultExternalAccountsOk returns a tuple with the RealmDefaultExternalAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDefaultExternalAccounts

`func (o *RegisterQueue200Response) SetRealmDefaultExternalAccounts(v map[string]RegisterQueueResponseRealmDefaultExternalAccountsEntry)`

SetRealmDefaultExternalAccounts sets RealmDefaultExternalAccounts field to given value.

### HasRealmDefaultExternalAccounts

`func (o *RegisterQueue200Response) HasRealmDefaultExternalAccounts() bool`

HasRealmDefaultExternalAccounts returns a boolean if a field has been set.

### GetJitsiServerUrl

`func (o *RegisterQueue200Response) GetJitsiServerUrl() string`

GetJitsiServerUrl returns the JitsiServerUrl field if non-nil, zero value otherwise.

### GetJitsiServerUrlOk

`func (o *RegisterQueue200Response) GetJitsiServerUrlOk() (*string, bool)`

GetJitsiServerUrlOk returns a tuple with the JitsiServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitsiServerUrl

`func (o *RegisterQueue200Response) SetJitsiServerUrl(v string)`

SetJitsiServerUrl sets JitsiServerUrl field to given value.

### HasJitsiServerUrl

`func (o *RegisterQueue200Response) HasJitsiServerUrl() bool`

HasJitsiServerUrl returns a boolean if a field has been set.

### GetDevelopmentEnvironment

`func (o *RegisterQueue200Response) GetDevelopmentEnvironment() bool`

GetDevelopmentEnvironment returns the DevelopmentEnvironment field if non-nil, zero value otherwise.

### GetDevelopmentEnvironmentOk

`func (o *RegisterQueue200Response) GetDevelopmentEnvironmentOk() (*bool, bool)`

GetDevelopmentEnvironmentOk returns a tuple with the DevelopmentEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentEnvironment

`func (o *RegisterQueue200Response) SetDevelopmentEnvironment(v bool)`

SetDevelopmentEnvironment sets DevelopmentEnvironment field to given value.

### HasDevelopmentEnvironment

`func (o *RegisterQueue200Response) HasDevelopmentEnvironment() bool`

HasDevelopmentEnvironment returns a boolean if a field has been set.

### GetServerGeneration

`func (o *RegisterQueue200Response) GetServerGeneration() int32`

GetServerGeneration returns the ServerGeneration field if non-nil, zero value otherwise.

### GetServerGenerationOk

`func (o *RegisterQueue200Response) GetServerGenerationOk() (*int32, bool)`

GetServerGenerationOk returns a tuple with the ServerGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGeneration

`func (o *RegisterQueue200Response) SetServerGeneration(v int32)`

SetServerGeneration sets ServerGeneration field to given value.

### HasServerGeneration

`func (o *RegisterQueue200Response) HasServerGeneration() bool`

HasServerGeneration returns a boolean if a field has been set.

### GetPasswordMinLength

`func (o *RegisterQueue200Response) GetPasswordMinLength() int32`

GetPasswordMinLength returns the PasswordMinLength field if non-nil, zero value otherwise.

### GetPasswordMinLengthOk

`func (o *RegisterQueue200Response) GetPasswordMinLengthOk() (*int32, bool)`

GetPasswordMinLengthOk returns a tuple with the PasswordMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinLength

`func (o *RegisterQueue200Response) SetPasswordMinLength(v int32)`

SetPasswordMinLength sets PasswordMinLength field to given value.

### HasPasswordMinLength

`func (o *RegisterQueue200Response) HasPasswordMinLength() bool`

HasPasswordMinLength returns a boolean if a field has been set.

### GetPasswordMaxLength

`func (o *RegisterQueue200Response) GetPasswordMaxLength() int32`

GetPasswordMaxLength returns the PasswordMaxLength field if non-nil, zero value otherwise.

### GetPasswordMaxLengthOk

`func (o *RegisterQueue200Response) GetPasswordMaxLengthOk() (*int32, bool)`

GetPasswordMaxLengthOk returns a tuple with the PasswordMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMaxLength

`func (o *RegisterQueue200Response) SetPasswordMaxLength(v int32)`

SetPasswordMaxLength sets PasswordMaxLength field to given value.

### HasPasswordMaxLength

`func (o *RegisterQueue200Response) HasPasswordMaxLength() bool`

HasPasswordMaxLength returns a boolean if a field has been set.

### GetPasswordMinGuesses

`func (o *RegisterQueue200Response) GetPasswordMinGuesses() int32`

GetPasswordMinGuesses returns the PasswordMinGuesses field if non-nil, zero value otherwise.

### GetPasswordMinGuessesOk

`func (o *RegisterQueue200Response) GetPasswordMinGuessesOk() (*int32, bool)`

GetPasswordMinGuessesOk returns a tuple with the PasswordMinGuesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinGuesses

`func (o *RegisterQueue200Response) SetPasswordMinGuesses(v int32)`

SetPasswordMinGuesses sets PasswordMinGuesses field to given value.

### HasPasswordMinGuesses

`func (o *RegisterQueue200Response) HasPasswordMinGuesses() bool`

HasPasswordMinGuesses returns a boolean if a field has been set.

### GetGiphyRatingOptions

`func (o *RegisterQueue200Response) GetGiphyRatingOptions() map[string]RegisterQueueResponseGiphyRatingOptionsEntry`

GetGiphyRatingOptions returns the GiphyRatingOptions field if non-nil, zero value otherwise.

### GetGiphyRatingOptionsOk

`func (o *RegisterQueue200Response) GetGiphyRatingOptionsOk() (*map[string]RegisterQueueResponseGiphyRatingOptionsEntry, bool)`

GetGiphyRatingOptionsOk returns a tuple with the GiphyRatingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiphyRatingOptions

`func (o *RegisterQueue200Response) SetGiphyRatingOptions(v map[string]RegisterQueueResponseGiphyRatingOptionsEntry)`

SetGiphyRatingOptions sets GiphyRatingOptions field to given value.

### HasGiphyRatingOptions

`func (o *RegisterQueue200Response) HasGiphyRatingOptions() bool`

HasGiphyRatingOptions returns a boolean if a field has been set.

### GetMaxFileUploadSizeMib

`func (o *RegisterQueue200Response) GetMaxFileUploadSizeMib() int32`

GetMaxFileUploadSizeMib returns the MaxFileUploadSizeMib field if non-nil, zero value otherwise.

### GetMaxFileUploadSizeMibOk

`func (o *RegisterQueue200Response) GetMaxFileUploadSizeMibOk() (*int32, bool)`

GetMaxFileUploadSizeMibOk returns a tuple with the MaxFileUploadSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileUploadSizeMib

`func (o *RegisterQueue200Response) SetMaxFileUploadSizeMib(v int32)`

SetMaxFileUploadSizeMib sets MaxFileUploadSizeMib field to given value.

### HasMaxFileUploadSizeMib

`func (o *RegisterQueue200Response) HasMaxFileUploadSizeMib() bool`

HasMaxFileUploadSizeMib returns a boolean if a field has been set.

### GetMaxAvatarFileSizeMib

`func (o *RegisterQueue200Response) GetMaxAvatarFileSizeMib() int32`

GetMaxAvatarFileSizeMib returns the MaxAvatarFileSizeMib field if non-nil, zero value otherwise.

### GetMaxAvatarFileSizeMibOk

`func (o *RegisterQueue200Response) GetMaxAvatarFileSizeMibOk() (*int32, bool)`

GetMaxAvatarFileSizeMibOk returns a tuple with the MaxAvatarFileSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvatarFileSizeMib

`func (o *RegisterQueue200Response) SetMaxAvatarFileSizeMib(v int32)`

SetMaxAvatarFileSizeMib sets MaxAvatarFileSizeMib field to given value.

### HasMaxAvatarFileSizeMib

`func (o *RegisterQueue200Response) HasMaxAvatarFileSizeMib() bool`

HasMaxAvatarFileSizeMib returns a boolean if a field has been set.

### GetServerInlineImagePreview

`func (o *RegisterQueue200Response) GetServerInlineImagePreview() bool`

GetServerInlineImagePreview returns the ServerInlineImagePreview field if non-nil, zero value otherwise.

### GetServerInlineImagePreviewOk

`func (o *RegisterQueue200Response) GetServerInlineImagePreviewOk() (*bool, bool)`

GetServerInlineImagePreviewOk returns a tuple with the ServerInlineImagePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInlineImagePreview

`func (o *RegisterQueue200Response) SetServerInlineImagePreview(v bool)`

SetServerInlineImagePreview sets ServerInlineImagePreview field to given value.

### HasServerInlineImagePreview

`func (o *RegisterQueue200Response) HasServerInlineImagePreview() bool`

HasServerInlineImagePreview returns a boolean if a field has been set.

### GetServerInlineUrlEmbedPreview

`func (o *RegisterQueue200Response) GetServerInlineUrlEmbedPreview() bool`

GetServerInlineUrlEmbedPreview returns the ServerInlineUrlEmbedPreview field if non-nil, zero value otherwise.

### GetServerInlineUrlEmbedPreviewOk

`func (o *RegisterQueue200Response) GetServerInlineUrlEmbedPreviewOk() (*bool, bool)`

GetServerInlineUrlEmbedPreviewOk returns a tuple with the ServerInlineUrlEmbedPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInlineUrlEmbedPreview

`func (o *RegisterQueue200Response) SetServerInlineUrlEmbedPreview(v bool)`

SetServerInlineUrlEmbedPreview sets ServerInlineUrlEmbedPreview field to given value.

### HasServerInlineUrlEmbedPreview

`func (o *RegisterQueue200Response) HasServerInlineUrlEmbedPreview() bool`

HasServerInlineUrlEmbedPreview returns a boolean if a field has been set.

### GetServerThumbnailFormats

`func (o *RegisterQueue200Response) GetServerThumbnailFormats() []RegisterQueueResponseServerThumbnailFormatsItem`

GetServerThumbnailFormats returns the ServerThumbnailFormats field if non-nil, zero value otherwise.

### GetServerThumbnailFormatsOk

`func (o *RegisterQueue200Response) GetServerThumbnailFormatsOk() (*[]RegisterQueueResponseServerThumbnailFormatsItem, bool)`

GetServerThumbnailFormatsOk returns a tuple with the ServerThumbnailFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerThumbnailFormats

`func (o *RegisterQueue200Response) SetServerThumbnailFormats(v []RegisterQueueResponseServerThumbnailFormatsItem)`

SetServerThumbnailFormats sets ServerThumbnailFormats field to given value.

### HasServerThumbnailFormats

`func (o *RegisterQueue200Response) HasServerThumbnailFormats() bool`

HasServerThumbnailFormats returns a boolean if a field has been set.

### GetServerAvatarChangesDisabled

`func (o *RegisterQueue200Response) GetServerAvatarChangesDisabled() bool`

GetServerAvatarChangesDisabled returns the ServerAvatarChangesDisabled field if non-nil, zero value otherwise.

### GetServerAvatarChangesDisabledOk

`func (o *RegisterQueue200Response) GetServerAvatarChangesDisabledOk() (*bool, bool)`

GetServerAvatarChangesDisabledOk returns a tuple with the ServerAvatarChangesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAvatarChangesDisabled

`func (o *RegisterQueue200Response) SetServerAvatarChangesDisabled(v bool)`

SetServerAvatarChangesDisabled sets ServerAvatarChangesDisabled field to given value.

### HasServerAvatarChangesDisabled

`func (o *RegisterQueue200Response) HasServerAvatarChangesDisabled() bool`

HasServerAvatarChangesDisabled returns a boolean if a field has been set.

### GetServerNameChangesDisabled

`func (o *RegisterQueue200Response) GetServerNameChangesDisabled() bool`

GetServerNameChangesDisabled returns the ServerNameChangesDisabled field if non-nil, zero value otherwise.

### GetServerNameChangesDisabledOk

`func (o *RegisterQueue200Response) GetServerNameChangesDisabledOk() (*bool, bool)`

GetServerNameChangesDisabledOk returns a tuple with the ServerNameChangesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNameChangesDisabled

`func (o *RegisterQueue200Response) SetServerNameChangesDisabled(v bool)`

SetServerNameChangesDisabled sets ServerNameChangesDisabled field to given value.

### HasServerNameChangesDisabled

`func (o *RegisterQueue200Response) HasServerNameChangesDisabled() bool`

HasServerNameChangesDisabled returns a boolean if a field has been set.

### GetServerNeedsUpgrade

`func (o *RegisterQueue200Response) GetServerNeedsUpgrade() bool`

GetServerNeedsUpgrade returns the ServerNeedsUpgrade field if non-nil, zero value otherwise.

### GetServerNeedsUpgradeOk

`func (o *RegisterQueue200Response) GetServerNeedsUpgradeOk() (*bool, bool)`

GetServerNeedsUpgradeOk returns a tuple with the ServerNeedsUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNeedsUpgrade

`func (o *RegisterQueue200Response) SetServerNeedsUpgrade(v bool)`

SetServerNeedsUpgrade sets ServerNeedsUpgrade field to given value.

### HasServerNeedsUpgrade

`func (o *RegisterQueue200Response) HasServerNeedsUpgrade() bool`

HasServerNeedsUpgrade returns a boolean if a field has been set.

### GetServerWebPublicStreamsEnabled

`func (o *RegisterQueue200Response) GetServerWebPublicStreamsEnabled() bool`

GetServerWebPublicStreamsEnabled returns the ServerWebPublicStreamsEnabled field if non-nil, zero value otherwise.

### GetServerWebPublicStreamsEnabledOk

`func (o *RegisterQueue200Response) GetServerWebPublicStreamsEnabledOk() (*bool, bool)`

GetServerWebPublicStreamsEnabledOk returns a tuple with the ServerWebPublicStreamsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWebPublicStreamsEnabled

`func (o *RegisterQueue200Response) SetServerWebPublicStreamsEnabled(v bool)`

SetServerWebPublicStreamsEnabled sets ServerWebPublicStreamsEnabled field to given value.

### HasServerWebPublicStreamsEnabled

`func (o *RegisterQueue200Response) HasServerWebPublicStreamsEnabled() bool`

HasServerWebPublicStreamsEnabled returns a boolean if a field has been set.

### GetServerEmojiDataUrl

`func (o *RegisterQueue200Response) GetServerEmojiDataUrl() string`

GetServerEmojiDataUrl returns the ServerEmojiDataUrl field if non-nil, zero value otherwise.

### GetServerEmojiDataUrlOk

`func (o *RegisterQueue200Response) GetServerEmojiDataUrlOk() (*string, bool)`

GetServerEmojiDataUrlOk returns a tuple with the ServerEmojiDataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerEmojiDataUrl

`func (o *RegisterQueue200Response) SetServerEmojiDataUrl(v string)`

SetServerEmojiDataUrl sets ServerEmojiDataUrl field to given value.

### HasServerEmojiDataUrl

`func (o *RegisterQueue200Response) HasServerEmojiDataUrl() bool`

HasServerEmojiDataUrl returns a boolean if a field has been set.

### GetServerJitsiServerUrl

`func (o *RegisterQueue200Response) GetServerJitsiServerUrl() string`

GetServerJitsiServerUrl returns the ServerJitsiServerUrl field if non-nil, zero value otherwise.

### GetServerJitsiServerUrlOk

`func (o *RegisterQueue200Response) GetServerJitsiServerUrlOk() (*string, bool)`

GetServerJitsiServerUrlOk returns a tuple with the ServerJitsiServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerJitsiServerUrl

`func (o *RegisterQueue200Response) SetServerJitsiServerUrl(v string)`

SetServerJitsiServerUrl sets ServerJitsiServerUrl field to given value.

### HasServerJitsiServerUrl

`func (o *RegisterQueue200Response) HasServerJitsiServerUrl() bool`

HasServerJitsiServerUrl returns a boolean if a field has been set.

### SetServerJitsiServerUrlNil

`func (o *RegisterQueue200Response) SetServerJitsiServerUrlNil(b bool)`

 SetServerJitsiServerUrlNil sets the value for ServerJitsiServerUrl to be an explicit nil

### UnsetServerJitsiServerUrl
`func (o *RegisterQueue200Response) UnsetServerJitsiServerUrl()`

UnsetServerJitsiServerUrl ensures that no value is present for ServerJitsiServerUrl, not even an explicit nil
### GetServerCanSummarizeTopics

`func (o *RegisterQueue200Response) GetServerCanSummarizeTopics() bool`

GetServerCanSummarizeTopics returns the ServerCanSummarizeTopics field if non-nil, zero value otherwise.

### GetServerCanSummarizeTopicsOk

`func (o *RegisterQueue200Response) GetServerCanSummarizeTopicsOk() (*bool, bool)`

GetServerCanSummarizeTopicsOk returns a tuple with the ServerCanSummarizeTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCanSummarizeTopics

`func (o *RegisterQueue200Response) SetServerCanSummarizeTopics(v bool)`

SetServerCanSummarizeTopics sets ServerCanSummarizeTopics field to given value.

### HasServerCanSummarizeTopics

`func (o *RegisterQueue200Response) HasServerCanSummarizeTopics() bool`

HasServerCanSummarizeTopics returns a boolean if a field has been set.

### GetEventQueueLongpollTimeoutSeconds

`func (o *RegisterQueue200Response) GetEventQueueLongpollTimeoutSeconds() int32`

GetEventQueueLongpollTimeoutSeconds returns the EventQueueLongpollTimeoutSeconds field if non-nil, zero value otherwise.

### GetEventQueueLongpollTimeoutSecondsOk

`func (o *RegisterQueue200Response) GetEventQueueLongpollTimeoutSecondsOk() (*int32, bool)`

GetEventQueueLongpollTimeoutSecondsOk returns a tuple with the EventQueueLongpollTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQueueLongpollTimeoutSeconds

`func (o *RegisterQueue200Response) SetEventQueueLongpollTimeoutSeconds(v int32)`

SetEventQueueLongpollTimeoutSeconds sets EventQueueLongpollTimeoutSeconds field to given value.

### HasEventQueueLongpollTimeoutSeconds

`func (o *RegisterQueue200Response) HasEventQueueLongpollTimeoutSeconds() bool`

HasEventQueueLongpollTimeoutSeconds returns a boolean if a field has been set.

### GetRealmBilling

`func (o *RegisterQueue200Response) GetRealmBilling() RegisterQueueResponseRealmBilling`

GetRealmBilling returns the RealmBilling field if non-nil, zero value otherwise.

### GetRealmBillingOk

`func (o *RegisterQueue200Response) GetRealmBillingOk() (*RegisterQueueResponseRealmBilling, bool)`

GetRealmBillingOk returns a tuple with the RealmBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmBilling

`func (o *RegisterQueue200Response) SetRealmBilling(v RegisterQueueResponseRealmBilling)`

SetRealmBilling sets RealmBilling field to given value.

### HasRealmBilling

`func (o *RegisterQueue200Response) HasRealmBilling() bool`

HasRealmBilling returns a boolean if a field has been set.

### GetRealmModerationRequestChannelId

`func (o *RegisterQueue200Response) GetRealmModerationRequestChannelId() int32`

GetRealmModerationRequestChannelId returns the RealmModerationRequestChannelId field if non-nil, zero value otherwise.

### GetRealmModerationRequestChannelIdOk

`func (o *RegisterQueue200Response) GetRealmModerationRequestChannelIdOk() (*int32, bool)`

GetRealmModerationRequestChannelIdOk returns a tuple with the RealmModerationRequestChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmModerationRequestChannelId

`func (o *RegisterQueue200Response) SetRealmModerationRequestChannelId(v int32)`

SetRealmModerationRequestChannelId sets RealmModerationRequestChannelId field to given value.

### HasRealmModerationRequestChannelId

`func (o *RegisterQueue200Response) HasRealmModerationRequestChannelId() bool`

HasRealmModerationRequestChannelId returns a boolean if a field has been set.

### GetRealmNewStreamAnnouncementsStreamId

`func (o *RegisterQueue200Response) GetRealmNewStreamAnnouncementsStreamId() int32`

GetRealmNewStreamAnnouncementsStreamId returns the RealmNewStreamAnnouncementsStreamId field if non-nil, zero value otherwise.

### GetRealmNewStreamAnnouncementsStreamIdOk

`func (o *RegisterQueue200Response) GetRealmNewStreamAnnouncementsStreamIdOk() (*int32, bool)`

GetRealmNewStreamAnnouncementsStreamIdOk returns a tuple with the RealmNewStreamAnnouncementsStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmNewStreamAnnouncementsStreamId

`func (o *RegisterQueue200Response) SetRealmNewStreamAnnouncementsStreamId(v int32)`

SetRealmNewStreamAnnouncementsStreamId sets RealmNewStreamAnnouncementsStreamId field to given value.

### HasRealmNewStreamAnnouncementsStreamId

`func (o *RegisterQueue200Response) HasRealmNewStreamAnnouncementsStreamId() bool`

HasRealmNewStreamAnnouncementsStreamId returns a boolean if a field has been set.

### GetRealmSignupAnnouncementsStreamId

`func (o *RegisterQueue200Response) GetRealmSignupAnnouncementsStreamId() int32`

GetRealmSignupAnnouncementsStreamId returns the RealmSignupAnnouncementsStreamId field if non-nil, zero value otherwise.

### GetRealmSignupAnnouncementsStreamIdOk

`func (o *RegisterQueue200Response) GetRealmSignupAnnouncementsStreamIdOk() (*int32, bool)`

GetRealmSignupAnnouncementsStreamIdOk returns a tuple with the RealmSignupAnnouncementsStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmSignupAnnouncementsStreamId

`func (o *RegisterQueue200Response) SetRealmSignupAnnouncementsStreamId(v int32)`

SetRealmSignupAnnouncementsStreamId sets RealmSignupAnnouncementsStreamId field to given value.

### HasRealmSignupAnnouncementsStreamId

`func (o *RegisterQueue200Response) HasRealmSignupAnnouncementsStreamId() bool`

HasRealmSignupAnnouncementsStreamId returns a boolean if a field has been set.

### GetRealmZulipUpdateAnnouncementsStreamId

`func (o *RegisterQueue200Response) GetRealmZulipUpdateAnnouncementsStreamId() int32`

GetRealmZulipUpdateAnnouncementsStreamId returns the RealmZulipUpdateAnnouncementsStreamId field if non-nil, zero value otherwise.

### GetRealmZulipUpdateAnnouncementsStreamIdOk

`func (o *RegisterQueue200Response) GetRealmZulipUpdateAnnouncementsStreamIdOk() (*int32, bool)`

GetRealmZulipUpdateAnnouncementsStreamIdOk returns a tuple with the RealmZulipUpdateAnnouncementsStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmZulipUpdateAnnouncementsStreamId

`func (o *RegisterQueue200Response) SetRealmZulipUpdateAnnouncementsStreamId(v int32)`

SetRealmZulipUpdateAnnouncementsStreamId sets RealmZulipUpdateAnnouncementsStreamId field to given value.

### HasRealmZulipUpdateAnnouncementsStreamId

`func (o *RegisterQueue200Response) HasRealmZulipUpdateAnnouncementsStreamId() bool`

HasRealmZulipUpdateAnnouncementsStreamId returns a boolean if a field has been set.

### GetRealmEmptyTopicDisplayName

`func (o *RegisterQueue200Response) GetRealmEmptyTopicDisplayName() string`

GetRealmEmptyTopicDisplayName returns the RealmEmptyTopicDisplayName field if non-nil, zero value otherwise.

### GetRealmEmptyTopicDisplayNameOk

`func (o *RegisterQueue200Response) GetRealmEmptyTopicDisplayNameOk() (*string, bool)`

GetRealmEmptyTopicDisplayNameOk returns a tuple with the RealmEmptyTopicDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmEmptyTopicDisplayName

`func (o *RegisterQueue200Response) SetRealmEmptyTopicDisplayName(v string)`

SetRealmEmptyTopicDisplayName sets RealmEmptyTopicDisplayName field to given value.

### HasRealmEmptyTopicDisplayName

`func (o *RegisterQueue200Response) HasRealmEmptyTopicDisplayName() bool`

HasRealmEmptyTopicDisplayName returns a boolean if a field has been set.

### GetRealmUserSettingsDefaults

`func (o *RegisterQueue200Response) GetRealmUserSettingsDefaults() RegisterQueueResponseRealmUserSettingsDefaults`

GetRealmUserSettingsDefaults returns the RealmUserSettingsDefaults field if non-nil, zero value otherwise.

### GetRealmUserSettingsDefaultsOk

`func (o *RegisterQueue200Response) GetRealmUserSettingsDefaultsOk() (*RegisterQueueResponseRealmUserSettingsDefaults, bool)`

GetRealmUserSettingsDefaultsOk returns a tuple with the RealmUserSettingsDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmUserSettingsDefaults

`func (o *RegisterQueue200Response) SetRealmUserSettingsDefaults(v RegisterQueueResponseRealmUserSettingsDefaults)`

SetRealmUserSettingsDefaults sets RealmUserSettingsDefaults field to given value.

### HasRealmUserSettingsDefaults

`func (o *RegisterQueue200Response) HasRealmUserSettingsDefaults() bool`

HasRealmUserSettingsDefaults returns a boolean if a field has been set.

### GetRealmUsers

`func (o *RegisterQueue200Response) GetRealmUsers() []User`

GetRealmUsers returns the RealmUsers field if non-nil, zero value otherwise.

### GetRealmUsersOk

`func (o *RegisterQueue200Response) GetRealmUsersOk() (*[]User, bool)`

GetRealmUsersOk returns a tuple with the RealmUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmUsers

`func (o *RegisterQueue200Response) SetRealmUsers(v []User)`

SetRealmUsers sets RealmUsers field to given value.

### HasRealmUsers

`func (o *RegisterQueue200Response) HasRealmUsers() bool`

HasRealmUsers returns a boolean if a field has been set.

### GetRealmNonActiveUsers

`func (o *RegisterQueue200Response) GetRealmNonActiveUsers() []User`

GetRealmNonActiveUsers returns the RealmNonActiveUsers field if non-nil, zero value otherwise.

### GetRealmNonActiveUsersOk

`func (o *RegisterQueue200Response) GetRealmNonActiveUsersOk() (*[]User, bool)`

GetRealmNonActiveUsersOk returns a tuple with the RealmNonActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmNonActiveUsers

`func (o *RegisterQueue200Response) SetRealmNonActiveUsers(v []User)`

SetRealmNonActiveUsers sets RealmNonActiveUsers field to given value.

### HasRealmNonActiveUsers

`func (o *RegisterQueue200Response) HasRealmNonActiveUsers() bool`

HasRealmNonActiveUsers returns a boolean if a field has been set.

### GetAvatarSource

`func (o *RegisterQueue200Response) GetAvatarSource() string`

GetAvatarSource returns the AvatarSource field if non-nil, zero value otherwise.

### GetAvatarSourceOk

`func (o *RegisterQueue200Response) GetAvatarSourceOk() (*string, bool)`

GetAvatarSourceOk returns a tuple with the AvatarSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSource

`func (o *RegisterQueue200Response) SetAvatarSource(v string)`

SetAvatarSource sets AvatarSource field to given value.

### HasAvatarSource

`func (o *RegisterQueue200Response) HasAvatarSource() bool`

HasAvatarSource returns a boolean if a field has been set.

### GetAvatarUrlMedium

`func (o *RegisterQueue200Response) GetAvatarUrlMedium() string`

GetAvatarUrlMedium returns the AvatarUrlMedium field if non-nil, zero value otherwise.

### GetAvatarUrlMediumOk

`func (o *RegisterQueue200Response) GetAvatarUrlMediumOk() (*string, bool)`

GetAvatarUrlMediumOk returns a tuple with the AvatarUrlMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrlMedium

`func (o *RegisterQueue200Response) SetAvatarUrlMedium(v string)`

SetAvatarUrlMedium sets AvatarUrlMedium field to given value.

### HasAvatarUrlMedium

`func (o *RegisterQueue200Response) HasAvatarUrlMedium() bool`

HasAvatarUrlMedium returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *RegisterQueue200Response) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *RegisterQueue200Response) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *RegisterQueue200Response) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *RegisterQueue200Response) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCanCreateStreams

`func (o *RegisterQueue200Response) GetCanCreateStreams() bool`

GetCanCreateStreams returns the CanCreateStreams field if non-nil, zero value otherwise.

### GetCanCreateStreamsOk

`func (o *RegisterQueue200Response) GetCanCreateStreamsOk() (*bool, bool)`

GetCanCreateStreamsOk returns a tuple with the CanCreateStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateStreams

`func (o *RegisterQueue200Response) SetCanCreateStreams(v bool)`

SetCanCreateStreams sets CanCreateStreams field to given value.

### HasCanCreateStreams

`func (o *RegisterQueue200Response) HasCanCreateStreams() bool`

HasCanCreateStreams returns a boolean if a field has been set.

### GetCanCreatePublicStreams

`func (o *RegisterQueue200Response) GetCanCreatePublicStreams() bool`

GetCanCreatePublicStreams returns the CanCreatePublicStreams field if non-nil, zero value otherwise.

### GetCanCreatePublicStreamsOk

`func (o *RegisterQueue200Response) GetCanCreatePublicStreamsOk() (*bool, bool)`

GetCanCreatePublicStreamsOk returns a tuple with the CanCreatePublicStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreatePublicStreams

`func (o *RegisterQueue200Response) SetCanCreatePublicStreams(v bool)`

SetCanCreatePublicStreams sets CanCreatePublicStreams field to given value.

### HasCanCreatePublicStreams

`func (o *RegisterQueue200Response) HasCanCreatePublicStreams() bool`

HasCanCreatePublicStreams returns a boolean if a field has been set.

### GetCanCreatePrivateStreams

`func (o *RegisterQueue200Response) GetCanCreatePrivateStreams() bool`

GetCanCreatePrivateStreams returns the CanCreatePrivateStreams field if non-nil, zero value otherwise.

### GetCanCreatePrivateStreamsOk

`func (o *RegisterQueue200Response) GetCanCreatePrivateStreamsOk() (*bool, bool)`

GetCanCreatePrivateStreamsOk returns a tuple with the CanCreatePrivateStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreatePrivateStreams

`func (o *RegisterQueue200Response) SetCanCreatePrivateStreams(v bool)`

SetCanCreatePrivateStreams sets CanCreatePrivateStreams field to given value.

### HasCanCreatePrivateStreams

`func (o *RegisterQueue200Response) HasCanCreatePrivateStreams() bool`

HasCanCreatePrivateStreams returns a boolean if a field has been set.

### GetCanCreateWebPublicStreams

`func (o *RegisterQueue200Response) GetCanCreateWebPublicStreams() bool`

GetCanCreateWebPublicStreams returns the CanCreateWebPublicStreams field if non-nil, zero value otherwise.

### GetCanCreateWebPublicStreamsOk

`func (o *RegisterQueue200Response) GetCanCreateWebPublicStreamsOk() (*bool, bool)`

GetCanCreateWebPublicStreamsOk returns a tuple with the CanCreateWebPublicStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateWebPublicStreams

`func (o *RegisterQueue200Response) SetCanCreateWebPublicStreams(v bool)`

SetCanCreateWebPublicStreams sets CanCreateWebPublicStreams field to given value.

### HasCanCreateWebPublicStreams

`func (o *RegisterQueue200Response) HasCanCreateWebPublicStreams() bool`

HasCanCreateWebPublicStreams returns a boolean if a field has been set.

### GetCanSubscribeOtherUsers

`func (o *RegisterQueue200Response) GetCanSubscribeOtherUsers() bool`

GetCanSubscribeOtherUsers returns the CanSubscribeOtherUsers field if non-nil, zero value otherwise.

### GetCanSubscribeOtherUsersOk

`func (o *RegisterQueue200Response) GetCanSubscribeOtherUsersOk() (*bool, bool)`

GetCanSubscribeOtherUsersOk returns a tuple with the CanSubscribeOtherUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSubscribeOtherUsers

`func (o *RegisterQueue200Response) SetCanSubscribeOtherUsers(v bool)`

SetCanSubscribeOtherUsers sets CanSubscribeOtherUsers field to given value.

### HasCanSubscribeOtherUsers

`func (o *RegisterQueue200Response) HasCanSubscribeOtherUsers() bool`

HasCanSubscribeOtherUsers returns a boolean if a field has been set.

### GetCanInviteOthersToRealm

`func (o *RegisterQueue200Response) GetCanInviteOthersToRealm() bool`

GetCanInviteOthersToRealm returns the CanInviteOthersToRealm field if non-nil, zero value otherwise.

### GetCanInviteOthersToRealmOk

`func (o *RegisterQueue200Response) GetCanInviteOthersToRealmOk() (*bool, bool)`

GetCanInviteOthersToRealmOk returns a tuple with the CanInviteOthersToRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInviteOthersToRealm

`func (o *RegisterQueue200Response) SetCanInviteOthersToRealm(v bool)`

SetCanInviteOthersToRealm sets CanInviteOthersToRealm field to given value.

### HasCanInviteOthersToRealm

`func (o *RegisterQueue200Response) HasCanInviteOthersToRealm() bool`

HasCanInviteOthersToRealm returns a boolean if a field has been set.

### GetIsAdmin

`func (o *RegisterQueue200Response) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *RegisterQueue200Response) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *RegisterQueue200Response) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *RegisterQueue200Response) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsOwner

`func (o *RegisterQueue200Response) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *RegisterQueue200Response) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *RegisterQueue200Response) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *RegisterQueue200Response) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsModerator

`func (o *RegisterQueue200Response) GetIsModerator() bool`

GetIsModerator returns the IsModerator field if non-nil, zero value otherwise.

### GetIsModeratorOk

`func (o *RegisterQueue200Response) GetIsModeratorOk() (*bool, bool)`

GetIsModeratorOk returns a tuple with the IsModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModerator

`func (o *RegisterQueue200Response) SetIsModerator(v bool)`

SetIsModerator sets IsModerator field to given value.

### HasIsModerator

`func (o *RegisterQueue200Response) HasIsModerator() bool`

HasIsModerator returns a boolean if a field has been set.

### GetIsGuest

`func (o *RegisterQueue200Response) GetIsGuest() bool`

GetIsGuest returns the IsGuest field if non-nil, zero value otherwise.

### GetIsGuestOk

`func (o *RegisterQueue200Response) GetIsGuestOk() (*bool, bool)`

GetIsGuestOk returns a tuple with the IsGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuest

`func (o *RegisterQueue200Response) SetIsGuest(v bool)`

SetIsGuest sets IsGuest field to given value.

### HasIsGuest

`func (o *RegisterQueue200Response) HasIsGuest() bool`

HasIsGuest returns a boolean if a field has been set.

### GetUserId

`func (o *RegisterQueue200Response) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RegisterQueue200Response) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RegisterQueue200Response) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RegisterQueue200Response) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmail

`func (o *RegisterQueue200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterQueue200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterQueue200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RegisterQueue200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDeliveryEmail

`func (o *RegisterQueue200Response) GetDeliveryEmail() string`

GetDeliveryEmail returns the DeliveryEmail field if non-nil, zero value otherwise.

### GetDeliveryEmailOk

`func (o *RegisterQueue200Response) GetDeliveryEmailOk() (*string, bool)`

GetDeliveryEmailOk returns a tuple with the DeliveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEmail

`func (o *RegisterQueue200Response) SetDeliveryEmail(v string)`

SetDeliveryEmail sets DeliveryEmail field to given value.

### HasDeliveryEmail

`func (o *RegisterQueue200Response) HasDeliveryEmail() bool`

HasDeliveryEmail returns a boolean if a field has been set.

### GetFullName

`func (o *RegisterQueue200Response) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *RegisterQueue200Response) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *RegisterQueue200Response) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *RegisterQueue200Response) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetCrossRealmBots

`func (o *RegisterQueue200Response) GetCrossRealmBots() []RegisterQueueResponseCrossRealmBotsItem`

GetCrossRealmBots returns the CrossRealmBots field if non-nil, zero value otherwise.

### GetCrossRealmBotsOk

`func (o *RegisterQueue200Response) GetCrossRealmBotsOk() (*[]RegisterQueueResponseCrossRealmBotsItem, bool)`

GetCrossRealmBotsOk returns a tuple with the CrossRealmBots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossRealmBots

`func (o *RegisterQueue200Response) SetCrossRealmBots(v []RegisterQueueResponseCrossRealmBotsItem)`

SetCrossRealmBots sets CrossRealmBots field to given value.

### HasCrossRealmBots

`func (o *RegisterQueue200Response) HasCrossRealmBots() bool`

HasCrossRealmBots returns a boolean if a field has been set.

### GetServerSupportedPermissionSettings

`func (o *RegisterQueue200Response) GetServerSupportedPermissionSettings() RegisterQueueResponseServerSupportedPermissionSettings`

GetServerSupportedPermissionSettings returns the ServerSupportedPermissionSettings field if non-nil, zero value otherwise.

### GetServerSupportedPermissionSettingsOk

`func (o *RegisterQueue200Response) GetServerSupportedPermissionSettingsOk() (*RegisterQueueResponseServerSupportedPermissionSettings, bool)`

GetServerSupportedPermissionSettingsOk returns a tuple with the ServerSupportedPermissionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSupportedPermissionSettings

`func (o *RegisterQueue200Response) SetServerSupportedPermissionSettings(v RegisterQueueResponseServerSupportedPermissionSettings)`

SetServerSupportedPermissionSettings sets ServerSupportedPermissionSettings field to given value.

### HasServerSupportedPermissionSettings

`func (o *RegisterQueue200Response) HasServerSupportedPermissionSettings() bool`

HasServerSupportedPermissionSettings returns a boolean if a field has been set.

### GetMaxBulkNewSubscriptionMessages

`func (o *RegisterQueue200Response) GetMaxBulkNewSubscriptionMessages() float32`

GetMaxBulkNewSubscriptionMessages returns the MaxBulkNewSubscriptionMessages field if non-nil, zero value otherwise.

### GetMaxBulkNewSubscriptionMessagesOk

`func (o *RegisterQueue200Response) GetMaxBulkNewSubscriptionMessagesOk() (*float32, bool)`

GetMaxBulkNewSubscriptionMessagesOk returns a tuple with the MaxBulkNewSubscriptionMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBulkNewSubscriptionMessages

`func (o *RegisterQueue200Response) SetMaxBulkNewSubscriptionMessages(v float32)`

SetMaxBulkNewSubscriptionMessages sets MaxBulkNewSubscriptionMessages field to given value.

### HasMaxBulkNewSubscriptionMessages

`func (o *RegisterQueue200Response) HasMaxBulkNewSubscriptionMessages() bool`

HasMaxBulkNewSubscriptionMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


