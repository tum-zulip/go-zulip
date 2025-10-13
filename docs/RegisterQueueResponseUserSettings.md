# RegisterQueueResponseUserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TwentyFourHourTime** | Pointer to **NullableBool** | Whether time should be [displayed in 24-hour notation](/help/change-the-time-format).  A &#x60;null&#x60; value indicates that the client should use the default time format for the user&#39;s locale.  **Changes**: Prior to Zulip 11.0 (feature level 408), &#x60;null&#x60; was not a valid value for this setting. Note that it was not possible to actually set the time format to &#x60;null&#x60; at this feature level.  | [optional] 
**WebMarkReadOnScrollPolicy** | Pointer to **int32** | Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \&quot;Always\&quot; setting when marking messages as read.  | [optional] 
**WebChannelDefaultView** | Pointer to **int32** | Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \&quot;Top unread topic in channel\&quot; is new in Zulip 11.0 (feature level 401).  The \&quot;List of topics\&quot; option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \&quot;Channel feed\&quot; behavior.  | [optional] 
**StarredMessageCounts** | Pointer to **bool** | Whether clients should display the [number of starred messages](/help/star-a-message#display-the-number-of-starred-messages).  | [optional] 
**ReceivesTypingNotifications** | Pointer to **bool** | Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.  | [optional] 
**WebSuggestUpdateTimezone** | Pointer to **bool** | Whether the user should be shown an alert, offering to update their [profile time zone](/help/change-your-timezone), when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.  **Changes**: New in Zulip 10.0 (feature level 329).  | [optional] 
**FluidLayoutWidth** | Pointer to **bool** | Whether to use the [maximum available screen width](/help/enable-full-width-display) for the web app&#39;s center panel (message feed, recent conversations) on wide screens.  | [optional] 
**HighContrastMode** | Pointer to **bool** | This setting is reserved for use to control variations in Zulip&#39;s design to help visually impaired users.  | [optional] 
**WebFontSizePx** | Pointer to **int32** | User-configured primary &#x60;font-size&#x60; for the web application, in pixels.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.  | [optional] 
**WebLineHeightPercent** | Pointer to **int32** | User-configured primary &#x60;line-height&#x60; for the web application, in percent, so a value of 120 represents a &#x60;line-height&#x60; of 1.2.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.  | [optional] 
**ColorScheme** | Pointer to **int32** | Controls which [color theme](/help/dark-theme) to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard &#x60;prefers-color-scheme&#x60; media query.  | [optional] 
**TranslateEmoticons** | Pointer to **bool** | Whether to [translate emoticons to emoji](/help/configure-emoticon-translations) in messages the user sends.  | [optional] 
**DisplayEmojiReactionUsers** | Pointer to **bool** | Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.  **Changes**: New in Zulip 6.0 (feature level 125).  | [optional] 
**DefaultLanguage** | Pointer to **string** | What [default language](/help/change-your-language) to use for the account.  This controls both the Zulip UI as well as email notifications sent to the user.  The value needs to be a standard language code that the Zulip server has translation data for; for example, &#x60;\&quot;en\&quot;&#x60; for English or &#x60;\&quot;de\&quot;&#x60; for German.  | [optional] 
**WebHomeView** | Pointer to **string** | The [home view](/help/configure-home-view) used when opening a new Zulip web app window or hitting the &#x60;Esc&#x60; keyboard shortcut repeatedly.  - \&quot;recent_topics\&quot; - Recent conversations view - \&quot;inbox\&quot; - Inbox view - \&quot;all_messages\&quot; - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;default_view&#x60;, which was new in Zulip 4.0 (feature level 42).  | [optional] 
**WebEscapeNavigatesToHomeView** | Pointer to **bool** | Whether the escape key navigates to the [configured home view](/help/configure-home-view).  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;escape_navigates_to_default_view&#x60;, which was new in Zulip 5.0 (feature level 107).  | [optional] 
**LeftSideUserlist** | Pointer to **bool** | Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.  | [optional] 
**Emojiset** | Pointer to **string** | The user&#39;s configured [emoji set](/help/emoji-and-emoticons#use-emoticons), used to display emoji to the user everywhere they appear in the UI.  - \&quot;google\&quot; - Google modern - \&quot;google-blob\&quot; - Google classic - \&quot;twitter\&quot; - Twitter - \&quot;text\&quot; - Plain text  | [optional] 
**DemoteInactiveStreams** | Pointer to **int32** | Whether to [hide inactive channels](/help/manage-inactive-channels) in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never  | [optional] 
**UserListStyle** | Pointer to **int32** | The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).  | [optional] 
**WebAnimateImagePreviews** | Pointer to **string** | Controls how animated images should be played in the message feed in the web/desktop application.  - \&quot;always\&quot; - Always play the animated images in the message feed. - \&quot;on_hover\&quot; - Play the animated images on hover over them in the message feed. - \&quot;never\&quot; - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275).  | [optional] 
**WebStreamUnreadsCountDisplayPolicy** | Pointer to **int32** | Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).  | [optional] 
**HideAiFeatures** | Pointer to **bool** | Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).  | [optional] 
**WebLeftSidebarShowChannelFolders** | Pointer to **bool** | Determines whether the web/desktop application&#39;s left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).  | [optional] 
**WebLeftSidebarUnreadsCountSummary** | Pointer to **bool** | Determines whether the web/desktop application&#39;s left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).  | [optional] 
**Timezone** | Pointer to **string** | The IANA identifier of the user&#39;s [profile time zone](/help/change-your-timezone), which is used primarily to display the user&#39;s local time to other users.  | [optional] 
**EnterSends** | Pointer to **bool** | Whether the user setting for [sending on pressing Enter](/help/configure-send-message-keys) in the compose box is enabled.  | [optional] 
**EnableDraftsSynchronization** | Pointer to **bool** | A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.  | [optional] 
**EnableStreamDesktopNotifications** | Pointer to **bool** | Enable visual desktop notifications for channel messages.  | [optional] 
**EnableStreamEmailNotifications** | Pointer to **bool** | Enable email notifications for channel messages.  | [optional] 
**EnableStreamPushNotifications** | Pointer to **bool** | Enable mobile notifications for channel messages.  | [optional] 
**EnableStreamAudibleNotifications** | Pointer to **bool** | Enable audible desktop notifications for channel messages.  | [optional] 
**NotificationSound** | Pointer to **string** | Notification sound name.  | [optional] 
**EnableDesktopNotifications** | Pointer to **bool** | Enable visual desktop notifications for direct messages and @-mentions.  | [optional] 
**EnableSounds** | Pointer to **bool** | Enable audible desktop notifications for direct messages and @-mentions.  | [optional] 
**EnableFollowedTopicDesktopNotifications** | Pointer to **bool** | Enable visual desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | [optional] 
**EnableFollowedTopicEmailNotifications** | Pointer to **bool** | Enable email notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | [optional] 
**EnableFollowedTopicPushNotifications** | Pointer to **bool** | Enable push notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | [optional] 
**EnableFollowedTopicAudibleNotifications** | Pointer to **bool** | Enable audible desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | [optional] 
**EmailNotificationsBatchingPeriodSeconds** | Pointer to **int32** | The duration (in seconds) for which the server should wait to batch email notifications before sending them.  | [optional] 
**EnableOfflineEmailNotifications** | Pointer to **bool** | Enable email notifications for direct messages and @-mentions received when the user is offline.  | [optional] 
**EnableOfflinePushNotifications** | Pointer to **bool** | Enable mobile notification for direct messages and @-mentions received when the user is offline.  | [optional] 
**EnableOnlinePushNotifications** | Pointer to **bool** | Enable mobile notification for direct messages and @-mentions received when the user is online.  | [optional] 
**EnableDigestEmails** | Pointer to **bool** | Enable digest emails when the user is away.  | [optional] 
**EnableMarketingEmails** | Pointer to **bool** | Enable marketing emails. Has no function outside Zulip Cloud.  | [optional] 
**EnableLoginEmails** | Pointer to **bool** | Enable email notifications for new logins to account.  | [optional] 
**MessageContentInEmailNotifications** | Pointer to **bool** | Include the message&#39;s content in email notifications for new messages.  | [optional] 
**PmContentInDesktopNotifications** | Pointer to **bool** | Include content of direct messages in desktop notifications.  | [optional] 
**WildcardMentionsNotify** | Pointer to **bool** | Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.  | [optional] 
**EnableFollowedTopicWildcardMentionsNotify** | Pointer to **bool** | Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.  **Changes**: New in Zulip 8.0 (feature level 189).  | [optional] 
**DesktopIconCountDisplay** | Pointer to **int32** | Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added &#x60;DMs, mentions, and followed topics&#x60; option, renumbering the options to insert it in order.  | [optional] 
**RealmNameInEmailNotificationsPolicy** | Pointer to **int32** | Whether to [include organization name in subject of message notification emails](/help/email-notifications#include-organization-name-in-subject-line).  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous &#x60;realm_name_in_notifications&#x60; boolean; &#x60;true&#x60; corresponded to &#x60;Always&#x60;, and &#x60;false&#x60; to &#x60;Never&#x60;.  | [optional] 
**AutomaticallyFollowTopicsPolicy** | Pointer to **int32** | Which [topics to follow automatically](/help/mute-a-topic).  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  | [optional] 
**AutomaticallyUnmuteTopicsInMutedStreamsPolicy** | Pointer to **int32** | Which [topics to unmute automatically in muted channels](/help/mute-a-topic).  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  | [optional] 
**AutomaticallyFollowTopicsWhereMentioned** | Pointer to **bool** | Whether the server will automatically mark the user as following topics where the user is mentioned.  **Changes**: New in Zulip 8.0 (feature level 235).  | [optional] 
**ResolvedTopicNoticeAutoReadPolicy** | Pointer to **string** | Controls whether the resolved-topic notices are marked as read.  - \&quot;always\&quot; - Always mark resolved-topic notices as read. - \&quot;except_followed\&quot; - Mark resolved-topic notices as read in topics not followed by the user. - \&quot;never\&quot; - Never mark resolved-topic notices as read.  **Changes**: New in Zulip 11.0 (feature level 385).  | [optional] 
**PresenceEnabled** | Pointer to **bool** | Display the presence status to other users when online.  | [optional] 
**AvailableNotificationSounds** | Pointer to **[]string** | Array containing the names of the notification sound options supported by this Zulip server. Only relevant to support UI for configuring notification sounds.  | [optional] 
**EmojisetChoices** | Pointer to [**[]RegisterQueueResponseUserSettingsEmojisetChoicesInner**](RegisterQueueResponseUserSettingsEmojisetChoicesInner.md) | Array of dictionaries where each dictionary describes an emoji set supported by this version of the Zulip server.  Only relevant to clients with configuration UI for choosing an emoji set; the currently selected emoji set is available in the &#x60;emojiset&#x60; key.  See [PATCH /settings](/api/update-settings) for details on the meaning of this setting.  | [optional] 
**SendPrivateTypingNotifications** | Pointer to **bool** | Whether the user has chosen to send [typing notifications](/help/typing-notifications) when composing direct messages. The client should send typing notifications for direct messages if and only if this setting is enabled.  **Changes**: New in Zulip 5.0 (feature level 105).  | [optional] 
**SendStreamTypingNotifications** | Pointer to **bool** | Whether the user has chosen to send [typing notifications](/help/typing-notifications) when composing channel messages. The client should send typing notifications for channel messages if and only if this setting is enabled.  **Changes**: New in Zulip 5.0 (feature level 105).  | [optional] 
**SendReadReceipts** | Pointer to **bool** | Whether other users are allowed to see whether you&#39;ve read messages.  **Changes**: New in Zulip 5.0 (feature level 105).  | [optional] 
**AllowPrivateDataExport** | Pointer to **bool** | Whether organization administrators are allowed to export your private data.  **Changes**: New in Zulip 10.0 (feature level 293).  | [optional] 
**EmailAddressVisibility** | Pointer to **int32** | The [policy][permission-level] for [which other users][help-email-visibility] in this organization can see the user&#39;s real email address.  - 1 &#x3D; Everyone - 2 &#x3D; Members only - 3 &#x3D; Administrators only - 4 &#x3D; Nobody - 5 &#x3D; Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.  [permission-level]: /api/roles-and-permissions#permission-levels [help-email-visibility]: /help/configure-email-visibility  | [optional] 
**WebNavigateToSentMessage** | Pointer to **bool** | Web/desktop app setting for whether the user&#39;s view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.  | [optional] 

## Methods

### NewRegisterQueueResponseUserSettings

`func NewRegisterQueueResponseUserSettings() *RegisterQueueResponseUserSettings`

NewRegisterQueueResponseUserSettings instantiates a new RegisterQueueResponseUserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseUserSettingsWithDefaults

`func NewRegisterQueueResponseUserSettingsWithDefaults() *RegisterQueueResponseUserSettings`

NewRegisterQueueResponseUserSettingsWithDefaults instantiates a new RegisterQueueResponseUserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwentyFourHourTime

`func (o *RegisterQueueResponseUserSettings) GetTwentyFourHourTime() bool`

GetTwentyFourHourTime returns the TwentyFourHourTime field if non-nil, zero value otherwise.

### GetTwentyFourHourTimeOk

`func (o *RegisterQueueResponseUserSettings) GetTwentyFourHourTimeOk() (*bool, bool)`

GetTwentyFourHourTimeOk returns a tuple with the TwentyFourHourTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwentyFourHourTime

`func (o *RegisterQueueResponseUserSettings) SetTwentyFourHourTime(v bool)`

SetTwentyFourHourTime sets TwentyFourHourTime field to given value.

### HasTwentyFourHourTime

`func (o *RegisterQueueResponseUserSettings) HasTwentyFourHourTime() bool`

HasTwentyFourHourTime returns a boolean if a field has been set.

### SetTwentyFourHourTimeNil

`func (o *RegisterQueueResponseUserSettings) SetTwentyFourHourTimeNil(b bool)`

 SetTwentyFourHourTimeNil sets the value for TwentyFourHourTime to be an explicit nil

### UnsetTwentyFourHourTime
`func (o *RegisterQueueResponseUserSettings) UnsetTwentyFourHourTime()`

UnsetTwentyFourHourTime ensures that no value is present for TwentyFourHourTime, not even an explicit nil
### GetWebMarkReadOnScrollPolicy

`func (o *RegisterQueueResponseUserSettings) GetWebMarkReadOnScrollPolicy() int32`

GetWebMarkReadOnScrollPolicy returns the WebMarkReadOnScrollPolicy field if non-nil, zero value otherwise.

### GetWebMarkReadOnScrollPolicyOk

`func (o *RegisterQueueResponseUserSettings) GetWebMarkReadOnScrollPolicyOk() (*int32, bool)`

GetWebMarkReadOnScrollPolicyOk returns a tuple with the WebMarkReadOnScrollPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMarkReadOnScrollPolicy

`func (o *RegisterQueueResponseUserSettings) SetWebMarkReadOnScrollPolicy(v int32)`

SetWebMarkReadOnScrollPolicy sets WebMarkReadOnScrollPolicy field to given value.

### HasWebMarkReadOnScrollPolicy

`func (o *RegisterQueueResponseUserSettings) HasWebMarkReadOnScrollPolicy() bool`

HasWebMarkReadOnScrollPolicy returns a boolean if a field has been set.

### GetWebChannelDefaultView

`func (o *RegisterQueueResponseUserSettings) GetWebChannelDefaultView() int32`

GetWebChannelDefaultView returns the WebChannelDefaultView field if non-nil, zero value otherwise.

### GetWebChannelDefaultViewOk

`func (o *RegisterQueueResponseUserSettings) GetWebChannelDefaultViewOk() (*int32, bool)`

GetWebChannelDefaultViewOk returns a tuple with the WebChannelDefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebChannelDefaultView

`func (o *RegisterQueueResponseUserSettings) SetWebChannelDefaultView(v int32)`

SetWebChannelDefaultView sets WebChannelDefaultView field to given value.

### HasWebChannelDefaultView

`func (o *RegisterQueueResponseUserSettings) HasWebChannelDefaultView() bool`

HasWebChannelDefaultView returns a boolean if a field has been set.

### GetStarredMessageCounts

`func (o *RegisterQueueResponseUserSettings) GetStarredMessageCounts() bool`

GetStarredMessageCounts returns the StarredMessageCounts field if non-nil, zero value otherwise.

### GetStarredMessageCountsOk

`func (o *RegisterQueueResponseUserSettings) GetStarredMessageCountsOk() (*bool, bool)`

GetStarredMessageCountsOk returns a tuple with the StarredMessageCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredMessageCounts

`func (o *RegisterQueueResponseUserSettings) SetStarredMessageCounts(v bool)`

SetStarredMessageCounts sets StarredMessageCounts field to given value.

### HasStarredMessageCounts

`func (o *RegisterQueueResponseUserSettings) HasStarredMessageCounts() bool`

HasStarredMessageCounts returns a boolean if a field has been set.

### GetReceivesTypingNotifications

`func (o *RegisterQueueResponseUserSettings) GetReceivesTypingNotifications() bool`

GetReceivesTypingNotifications returns the ReceivesTypingNotifications field if non-nil, zero value otherwise.

### GetReceivesTypingNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetReceivesTypingNotificationsOk() (*bool, bool)`

GetReceivesTypingNotificationsOk returns a tuple with the ReceivesTypingNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivesTypingNotifications

`func (o *RegisterQueueResponseUserSettings) SetReceivesTypingNotifications(v bool)`

SetReceivesTypingNotifications sets ReceivesTypingNotifications field to given value.

### HasReceivesTypingNotifications

`func (o *RegisterQueueResponseUserSettings) HasReceivesTypingNotifications() bool`

HasReceivesTypingNotifications returns a boolean if a field has been set.

### GetWebSuggestUpdateTimezone

`func (o *RegisterQueueResponseUserSettings) GetWebSuggestUpdateTimezone() bool`

GetWebSuggestUpdateTimezone returns the WebSuggestUpdateTimezone field if non-nil, zero value otherwise.

### GetWebSuggestUpdateTimezoneOk

`func (o *RegisterQueueResponseUserSettings) GetWebSuggestUpdateTimezoneOk() (*bool, bool)`

GetWebSuggestUpdateTimezoneOk returns a tuple with the WebSuggestUpdateTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSuggestUpdateTimezone

`func (o *RegisterQueueResponseUserSettings) SetWebSuggestUpdateTimezone(v bool)`

SetWebSuggestUpdateTimezone sets WebSuggestUpdateTimezone field to given value.

### HasWebSuggestUpdateTimezone

`func (o *RegisterQueueResponseUserSettings) HasWebSuggestUpdateTimezone() bool`

HasWebSuggestUpdateTimezone returns a boolean if a field has been set.

### GetFluidLayoutWidth

`func (o *RegisterQueueResponseUserSettings) GetFluidLayoutWidth() bool`

GetFluidLayoutWidth returns the FluidLayoutWidth field if non-nil, zero value otherwise.

### GetFluidLayoutWidthOk

`func (o *RegisterQueueResponseUserSettings) GetFluidLayoutWidthOk() (*bool, bool)`

GetFluidLayoutWidthOk returns a tuple with the FluidLayoutWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFluidLayoutWidth

`func (o *RegisterQueueResponseUserSettings) SetFluidLayoutWidth(v bool)`

SetFluidLayoutWidth sets FluidLayoutWidth field to given value.

### HasFluidLayoutWidth

`func (o *RegisterQueueResponseUserSettings) HasFluidLayoutWidth() bool`

HasFluidLayoutWidth returns a boolean if a field has been set.

### GetHighContrastMode

`func (o *RegisterQueueResponseUserSettings) GetHighContrastMode() bool`

GetHighContrastMode returns the HighContrastMode field if non-nil, zero value otherwise.

### GetHighContrastModeOk

`func (o *RegisterQueueResponseUserSettings) GetHighContrastModeOk() (*bool, bool)`

GetHighContrastModeOk returns a tuple with the HighContrastMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighContrastMode

`func (o *RegisterQueueResponseUserSettings) SetHighContrastMode(v bool)`

SetHighContrastMode sets HighContrastMode field to given value.

### HasHighContrastMode

`func (o *RegisterQueueResponseUserSettings) HasHighContrastMode() bool`

HasHighContrastMode returns a boolean if a field has been set.

### GetWebFontSizePx

`func (o *RegisterQueueResponseUserSettings) GetWebFontSizePx() int32`

GetWebFontSizePx returns the WebFontSizePx field if non-nil, zero value otherwise.

### GetWebFontSizePxOk

`func (o *RegisterQueueResponseUserSettings) GetWebFontSizePxOk() (*int32, bool)`

GetWebFontSizePxOk returns a tuple with the WebFontSizePx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebFontSizePx

`func (o *RegisterQueueResponseUserSettings) SetWebFontSizePx(v int32)`

SetWebFontSizePx sets WebFontSizePx field to given value.

### HasWebFontSizePx

`func (o *RegisterQueueResponseUserSettings) HasWebFontSizePx() bool`

HasWebFontSizePx returns a boolean if a field has been set.

### GetWebLineHeightPercent

`func (o *RegisterQueueResponseUserSettings) GetWebLineHeightPercent() int32`

GetWebLineHeightPercent returns the WebLineHeightPercent field if non-nil, zero value otherwise.

### GetWebLineHeightPercentOk

`func (o *RegisterQueueResponseUserSettings) GetWebLineHeightPercentOk() (*int32, bool)`

GetWebLineHeightPercentOk returns a tuple with the WebLineHeightPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLineHeightPercent

`func (o *RegisterQueueResponseUserSettings) SetWebLineHeightPercent(v int32)`

SetWebLineHeightPercent sets WebLineHeightPercent field to given value.

### HasWebLineHeightPercent

`func (o *RegisterQueueResponseUserSettings) HasWebLineHeightPercent() bool`

HasWebLineHeightPercent returns a boolean if a field has been set.

### GetColorScheme

`func (o *RegisterQueueResponseUserSettings) GetColorScheme() int32`

GetColorScheme returns the ColorScheme field if non-nil, zero value otherwise.

### GetColorSchemeOk

`func (o *RegisterQueueResponseUserSettings) GetColorSchemeOk() (*int32, bool)`

GetColorSchemeOk returns a tuple with the ColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorScheme

`func (o *RegisterQueueResponseUserSettings) SetColorScheme(v int32)`

SetColorScheme sets ColorScheme field to given value.

### HasColorScheme

`func (o *RegisterQueueResponseUserSettings) HasColorScheme() bool`

HasColorScheme returns a boolean if a field has been set.

### GetTranslateEmoticons

`func (o *RegisterQueueResponseUserSettings) GetTranslateEmoticons() bool`

GetTranslateEmoticons returns the TranslateEmoticons field if non-nil, zero value otherwise.

### GetTranslateEmoticonsOk

`func (o *RegisterQueueResponseUserSettings) GetTranslateEmoticonsOk() (*bool, bool)`

GetTranslateEmoticonsOk returns a tuple with the TranslateEmoticons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslateEmoticons

`func (o *RegisterQueueResponseUserSettings) SetTranslateEmoticons(v bool)`

SetTranslateEmoticons sets TranslateEmoticons field to given value.

### HasTranslateEmoticons

`func (o *RegisterQueueResponseUserSettings) HasTranslateEmoticons() bool`

HasTranslateEmoticons returns a boolean if a field has been set.

### GetDisplayEmojiReactionUsers

`func (o *RegisterQueueResponseUserSettings) GetDisplayEmojiReactionUsers() bool`

GetDisplayEmojiReactionUsers returns the DisplayEmojiReactionUsers field if non-nil, zero value otherwise.

### GetDisplayEmojiReactionUsersOk

`func (o *RegisterQueueResponseUserSettings) GetDisplayEmojiReactionUsersOk() (*bool, bool)`

GetDisplayEmojiReactionUsersOk returns a tuple with the DisplayEmojiReactionUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayEmojiReactionUsers

`func (o *RegisterQueueResponseUserSettings) SetDisplayEmojiReactionUsers(v bool)`

SetDisplayEmojiReactionUsers sets DisplayEmojiReactionUsers field to given value.

### HasDisplayEmojiReactionUsers

`func (o *RegisterQueueResponseUserSettings) HasDisplayEmojiReactionUsers() bool`

HasDisplayEmojiReactionUsers returns a boolean if a field has been set.

### GetDefaultLanguage

`func (o *RegisterQueueResponseUserSettings) GetDefaultLanguage() string`

GetDefaultLanguage returns the DefaultLanguage field if non-nil, zero value otherwise.

### GetDefaultLanguageOk

`func (o *RegisterQueueResponseUserSettings) GetDefaultLanguageOk() (*string, bool)`

GetDefaultLanguageOk returns a tuple with the DefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguage

`func (o *RegisterQueueResponseUserSettings) SetDefaultLanguage(v string)`

SetDefaultLanguage sets DefaultLanguage field to given value.

### HasDefaultLanguage

`func (o *RegisterQueueResponseUserSettings) HasDefaultLanguage() bool`

HasDefaultLanguage returns a boolean if a field has been set.

### GetWebHomeView

`func (o *RegisterQueueResponseUserSettings) GetWebHomeView() string`

GetWebHomeView returns the WebHomeView field if non-nil, zero value otherwise.

### GetWebHomeViewOk

`func (o *RegisterQueueResponseUserSettings) GetWebHomeViewOk() (*string, bool)`

GetWebHomeViewOk returns a tuple with the WebHomeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHomeView

`func (o *RegisterQueueResponseUserSettings) SetWebHomeView(v string)`

SetWebHomeView sets WebHomeView field to given value.

### HasWebHomeView

`func (o *RegisterQueueResponseUserSettings) HasWebHomeView() bool`

HasWebHomeView returns a boolean if a field has been set.

### GetWebEscapeNavigatesToHomeView

`func (o *RegisterQueueResponseUserSettings) GetWebEscapeNavigatesToHomeView() bool`

GetWebEscapeNavigatesToHomeView returns the WebEscapeNavigatesToHomeView field if non-nil, zero value otherwise.

### GetWebEscapeNavigatesToHomeViewOk

`func (o *RegisterQueueResponseUserSettings) GetWebEscapeNavigatesToHomeViewOk() (*bool, bool)`

GetWebEscapeNavigatesToHomeViewOk returns a tuple with the WebEscapeNavigatesToHomeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebEscapeNavigatesToHomeView

`func (o *RegisterQueueResponseUserSettings) SetWebEscapeNavigatesToHomeView(v bool)`

SetWebEscapeNavigatesToHomeView sets WebEscapeNavigatesToHomeView field to given value.

### HasWebEscapeNavigatesToHomeView

`func (o *RegisterQueueResponseUserSettings) HasWebEscapeNavigatesToHomeView() bool`

HasWebEscapeNavigatesToHomeView returns a boolean if a field has been set.

### GetLeftSideUserlist

`func (o *RegisterQueueResponseUserSettings) GetLeftSideUserlist() bool`

GetLeftSideUserlist returns the LeftSideUserlist field if non-nil, zero value otherwise.

### GetLeftSideUserlistOk

`func (o *RegisterQueueResponseUserSettings) GetLeftSideUserlistOk() (*bool, bool)`

GetLeftSideUserlistOk returns a tuple with the LeftSideUserlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftSideUserlist

`func (o *RegisterQueueResponseUserSettings) SetLeftSideUserlist(v bool)`

SetLeftSideUserlist sets LeftSideUserlist field to given value.

### HasLeftSideUserlist

`func (o *RegisterQueueResponseUserSettings) HasLeftSideUserlist() bool`

HasLeftSideUserlist returns a boolean if a field has been set.

### GetEmojiset

`func (o *RegisterQueueResponseUserSettings) GetEmojiset() string`

GetEmojiset returns the Emojiset field if non-nil, zero value otherwise.

### GetEmojisetOk

`func (o *RegisterQueueResponseUserSettings) GetEmojisetOk() (*string, bool)`

GetEmojisetOk returns a tuple with the Emojiset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiset

`func (o *RegisterQueueResponseUserSettings) SetEmojiset(v string)`

SetEmojiset sets Emojiset field to given value.

### HasEmojiset

`func (o *RegisterQueueResponseUserSettings) HasEmojiset() bool`

HasEmojiset returns a boolean if a field has been set.

### GetDemoteInactiveStreams

`func (o *RegisterQueueResponseUserSettings) GetDemoteInactiveStreams() int32`

GetDemoteInactiveStreams returns the DemoteInactiveStreams field if non-nil, zero value otherwise.

### GetDemoteInactiveStreamsOk

`func (o *RegisterQueueResponseUserSettings) GetDemoteInactiveStreamsOk() (*int32, bool)`

GetDemoteInactiveStreamsOk returns a tuple with the DemoteInactiveStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemoteInactiveStreams

`func (o *RegisterQueueResponseUserSettings) SetDemoteInactiveStreams(v int32)`

SetDemoteInactiveStreams sets DemoteInactiveStreams field to given value.

### HasDemoteInactiveStreams

`func (o *RegisterQueueResponseUserSettings) HasDemoteInactiveStreams() bool`

HasDemoteInactiveStreams returns a boolean if a field has been set.

### GetUserListStyle

`func (o *RegisterQueueResponseUserSettings) GetUserListStyle() int32`

GetUserListStyle returns the UserListStyle field if non-nil, zero value otherwise.

### GetUserListStyleOk

`func (o *RegisterQueueResponseUserSettings) GetUserListStyleOk() (*int32, bool)`

GetUserListStyleOk returns a tuple with the UserListStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListStyle

`func (o *RegisterQueueResponseUserSettings) SetUserListStyle(v int32)`

SetUserListStyle sets UserListStyle field to given value.

### HasUserListStyle

`func (o *RegisterQueueResponseUserSettings) HasUserListStyle() bool`

HasUserListStyle returns a boolean if a field has been set.

### GetWebAnimateImagePreviews

`func (o *RegisterQueueResponseUserSettings) GetWebAnimateImagePreviews() string`

GetWebAnimateImagePreviews returns the WebAnimateImagePreviews field if non-nil, zero value otherwise.

### GetWebAnimateImagePreviewsOk

`func (o *RegisterQueueResponseUserSettings) GetWebAnimateImagePreviewsOk() (*string, bool)`

GetWebAnimateImagePreviewsOk returns a tuple with the WebAnimateImagePreviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAnimateImagePreviews

`func (o *RegisterQueueResponseUserSettings) SetWebAnimateImagePreviews(v string)`

SetWebAnimateImagePreviews sets WebAnimateImagePreviews field to given value.

### HasWebAnimateImagePreviews

`func (o *RegisterQueueResponseUserSettings) HasWebAnimateImagePreviews() bool`

HasWebAnimateImagePreviews returns a boolean if a field has been set.

### GetWebStreamUnreadsCountDisplayPolicy

`func (o *RegisterQueueResponseUserSettings) GetWebStreamUnreadsCountDisplayPolicy() int32`

GetWebStreamUnreadsCountDisplayPolicy returns the WebStreamUnreadsCountDisplayPolicy field if non-nil, zero value otherwise.

### GetWebStreamUnreadsCountDisplayPolicyOk

`func (o *RegisterQueueResponseUserSettings) GetWebStreamUnreadsCountDisplayPolicyOk() (*int32, bool)`

GetWebStreamUnreadsCountDisplayPolicyOk returns a tuple with the WebStreamUnreadsCountDisplayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebStreamUnreadsCountDisplayPolicy

`func (o *RegisterQueueResponseUserSettings) SetWebStreamUnreadsCountDisplayPolicy(v int32)`

SetWebStreamUnreadsCountDisplayPolicy sets WebStreamUnreadsCountDisplayPolicy field to given value.

### HasWebStreamUnreadsCountDisplayPolicy

`func (o *RegisterQueueResponseUserSettings) HasWebStreamUnreadsCountDisplayPolicy() bool`

HasWebStreamUnreadsCountDisplayPolicy returns a boolean if a field has been set.

### GetHideAiFeatures

`func (o *RegisterQueueResponseUserSettings) GetHideAiFeatures() bool`

GetHideAiFeatures returns the HideAiFeatures field if non-nil, zero value otherwise.

### GetHideAiFeaturesOk

`func (o *RegisterQueueResponseUserSettings) GetHideAiFeaturesOk() (*bool, bool)`

GetHideAiFeaturesOk returns a tuple with the HideAiFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAiFeatures

`func (o *RegisterQueueResponseUserSettings) SetHideAiFeatures(v bool)`

SetHideAiFeatures sets HideAiFeatures field to given value.

### HasHideAiFeatures

`func (o *RegisterQueueResponseUserSettings) HasHideAiFeatures() bool`

HasHideAiFeatures returns a boolean if a field has been set.

### GetWebLeftSidebarShowChannelFolders

`func (o *RegisterQueueResponseUserSettings) GetWebLeftSidebarShowChannelFolders() bool`

GetWebLeftSidebarShowChannelFolders returns the WebLeftSidebarShowChannelFolders field if non-nil, zero value otherwise.

### GetWebLeftSidebarShowChannelFoldersOk

`func (o *RegisterQueueResponseUserSettings) GetWebLeftSidebarShowChannelFoldersOk() (*bool, bool)`

GetWebLeftSidebarShowChannelFoldersOk returns a tuple with the WebLeftSidebarShowChannelFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLeftSidebarShowChannelFolders

`func (o *RegisterQueueResponseUserSettings) SetWebLeftSidebarShowChannelFolders(v bool)`

SetWebLeftSidebarShowChannelFolders sets WebLeftSidebarShowChannelFolders field to given value.

### HasWebLeftSidebarShowChannelFolders

`func (o *RegisterQueueResponseUserSettings) HasWebLeftSidebarShowChannelFolders() bool`

HasWebLeftSidebarShowChannelFolders returns a boolean if a field has been set.

### GetWebLeftSidebarUnreadsCountSummary

`func (o *RegisterQueueResponseUserSettings) GetWebLeftSidebarUnreadsCountSummary() bool`

GetWebLeftSidebarUnreadsCountSummary returns the WebLeftSidebarUnreadsCountSummary field if non-nil, zero value otherwise.

### GetWebLeftSidebarUnreadsCountSummaryOk

`func (o *RegisterQueueResponseUserSettings) GetWebLeftSidebarUnreadsCountSummaryOk() (*bool, bool)`

GetWebLeftSidebarUnreadsCountSummaryOk returns a tuple with the WebLeftSidebarUnreadsCountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLeftSidebarUnreadsCountSummary

`func (o *RegisterQueueResponseUserSettings) SetWebLeftSidebarUnreadsCountSummary(v bool)`

SetWebLeftSidebarUnreadsCountSummary sets WebLeftSidebarUnreadsCountSummary field to given value.

### HasWebLeftSidebarUnreadsCountSummary

`func (o *RegisterQueueResponseUserSettings) HasWebLeftSidebarUnreadsCountSummary() bool`

HasWebLeftSidebarUnreadsCountSummary returns a boolean if a field has been set.

### GetTimezone

`func (o *RegisterQueueResponseUserSettings) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RegisterQueueResponseUserSettings) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RegisterQueueResponseUserSettings) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *RegisterQueueResponseUserSettings) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetEnterSends

`func (o *RegisterQueueResponseUserSettings) GetEnterSends() bool`

GetEnterSends returns the EnterSends field if non-nil, zero value otherwise.

### GetEnterSendsOk

`func (o *RegisterQueueResponseUserSettings) GetEnterSendsOk() (*bool, bool)`

GetEnterSendsOk returns a tuple with the EnterSends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterSends

`func (o *RegisterQueueResponseUserSettings) SetEnterSends(v bool)`

SetEnterSends sets EnterSends field to given value.

### HasEnterSends

`func (o *RegisterQueueResponseUserSettings) HasEnterSends() bool`

HasEnterSends returns a boolean if a field has been set.

### GetEnableDraftsSynchronization

`func (o *RegisterQueueResponseUserSettings) GetEnableDraftsSynchronization() bool`

GetEnableDraftsSynchronization returns the EnableDraftsSynchronization field if non-nil, zero value otherwise.

### GetEnableDraftsSynchronizationOk

`func (o *RegisterQueueResponseUserSettings) GetEnableDraftsSynchronizationOk() (*bool, bool)`

GetEnableDraftsSynchronizationOk returns a tuple with the EnableDraftsSynchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDraftsSynchronization

`func (o *RegisterQueueResponseUserSettings) SetEnableDraftsSynchronization(v bool)`

SetEnableDraftsSynchronization sets EnableDraftsSynchronization field to given value.

### HasEnableDraftsSynchronization

`func (o *RegisterQueueResponseUserSettings) HasEnableDraftsSynchronization() bool`

HasEnableDraftsSynchronization returns a boolean if a field has been set.

### GetEnableStreamDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableStreamDesktopNotifications() bool`

GetEnableStreamDesktopNotifications returns the EnableStreamDesktopNotifications field if non-nil, zero value otherwise.

### GetEnableStreamDesktopNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableStreamDesktopNotificationsOk() (*bool, bool)`

GetEnableStreamDesktopNotificationsOk returns a tuple with the EnableStreamDesktopNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableStreamDesktopNotifications(v bool)`

SetEnableStreamDesktopNotifications sets EnableStreamDesktopNotifications field to given value.

### HasEnableStreamDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableStreamDesktopNotifications() bool`

HasEnableStreamDesktopNotifications returns a boolean if a field has been set.

### GetEnableStreamEmailNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableStreamEmailNotifications() bool`

GetEnableStreamEmailNotifications returns the EnableStreamEmailNotifications field if non-nil, zero value otherwise.

### GetEnableStreamEmailNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableStreamEmailNotificationsOk() (*bool, bool)`

GetEnableStreamEmailNotificationsOk returns a tuple with the EnableStreamEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamEmailNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableStreamEmailNotifications(v bool)`

SetEnableStreamEmailNotifications sets EnableStreamEmailNotifications field to given value.

### HasEnableStreamEmailNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableStreamEmailNotifications() bool`

HasEnableStreamEmailNotifications returns a boolean if a field has been set.

### GetEnableStreamPushNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableStreamPushNotifications() bool`

GetEnableStreamPushNotifications returns the EnableStreamPushNotifications field if non-nil, zero value otherwise.

### GetEnableStreamPushNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableStreamPushNotificationsOk() (*bool, bool)`

GetEnableStreamPushNotificationsOk returns a tuple with the EnableStreamPushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamPushNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableStreamPushNotifications(v bool)`

SetEnableStreamPushNotifications sets EnableStreamPushNotifications field to given value.

### HasEnableStreamPushNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableStreamPushNotifications() bool`

HasEnableStreamPushNotifications returns a boolean if a field has been set.

### GetEnableStreamAudibleNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableStreamAudibleNotifications() bool`

GetEnableStreamAudibleNotifications returns the EnableStreamAudibleNotifications field if non-nil, zero value otherwise.

### GetEnableStreamAudibleNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableStreamAudibleNotificationsOk() (*bool, bool)`

GetEnableStreamAudibleNotificationsOk returns a tuple with the EnableStreamAudibleNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamAudibleNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableStreamAudibleNotifications(v bool)`

SetEnableStreamAudibleNotifications sets EnableStreamAudibleNotifications field to given value.

### HasEnableStreamAudibleNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableStreamAudibleNotifications() bool`

HasEnableStreamAudibleNotifications returns a boolean if a field has been set.

### GetNotificationSound

`func (o *RegisterQueueResponseUserSettings) GetNotificationSound() string`

GetNotificationSound returns the NotificationSound field if non-nil, zero value otherwise.

### GetNotificationSoundOk

`func (o *RegisterQueueResponseUserSettings) GetNotificationSoundOk() (*string, bool)`

GetNotificationSoundOk returns a tuple with the NotificationSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSound

`func (o *RegisterQueueResponseUserSettings) SetNotificationSound(v string)`

SetNotificationSound sets NotificationSound field to given value.

### HasNotificationSound

`func (o *RegisterQueueResponseUserSettings) HasNotificationSound() bool`

HasNotificationSound returns a boolean if a field has been set.

### GetEnableDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableDesktopNotifications() bool`

GetEnableDesktopNotifications returns the EnableDesktopNotifications field if non-nil, zero value otherwise.

### GetEnableDesktopNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableDesktopNotificationsOk() (*bool, bool)`

GetEnableDesktopNotificationsOk returns a tuple with the EnableDesktopNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableDesktopNotifications(v bool)`

SetEnableDesktopNotifications sets EnableDesktopNotifications field to given value.

### HasEnableDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableDesktopNotifications() bool`

HasEnableDesktopNotifications returns a boolean if a field has been set.

### GetEnableSounds

`func (o *RegisterQueueResponseUserSettings) GetEnableSounds() bool`

GetEnableSounds returns the EnableSounds field if non-nil, zero value otherwise.

### GetEnableSoundsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableSoundsOk() (*bool, bool)`

GetEnableSoundsOk returns a tuple with the EnableSounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSounds

`func (o *RegisterQueueResponseUserSettings) SetEnableSounds(v bool)`

SetEnableSounds sets EnableSounds field to given value.

### HasEnableSounds

`func (o *RegisterQueueResponseUserSettings) HasEnableSounds() bool`

HasEnableSounds returns a boolean if a field has been set.

### GetEnableFollowedTopicDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicDesktopNotifications() bool`

GetEnableFollowedTopicDesktopNotifications returns the EnableFollowedTopicDesktopNotifications field if non-nil, zero value otherwise.

### GetEnableFollowedTopicDesktopNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicDesktopNotificationsOk() (*bool, bool)`

GetEnableFollowedTopicDesktopNotificationsOk returns a tuple with the EnableFollowedTopicDesktopNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFollowedTopicDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableFollowedTopicDesktopNotifications(v bool)`

SetEnableFollowedTopicDesktopNotifications sets EnableFollowedTopicDesktopNotifications field to given value.

### HasEnableFollowedTopicDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableFollowedTopicDesktopNotifications() bool`

HasEnableFollowedTopicDesktopNotifications returns a boolean if a field has been set.

### GetEnableFollowedTopicEmailNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicEmailNotifications() bool`

GetEnableFollowedTopicEmailNotifications returns the EnableFollowedTopicEmailNotifications field if non-nil, zero value otherwise.

### GetEnableFollowedTopicEmailNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicEmailNotificationsOk() (*bool, bool)`

GetEnableFollowedTopicEmailNotificationsOk returns a tuple with the EnableFollowedTopicEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFollowedTopicEmailNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableFollowedTopicEmailNotifications(v bool)`

SetEnableFollowedTopicEmailNotifications sets EnableFollowedTopicEmailNotifications field to given value.

### HasEnableFollowedTopicEmailNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableFollowedTopicEmailNotifications() bool`

HasEnableFollowedTopicEmailNotifications returns a boolean if a field has been set.

### GetEnableFollowedTopicPushNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicPushNotifications() bool`

GetEnableFollowedTopicPushNotifications returns the EnableFollowedTopicPushNotifications field if non-nil, zero value otherwise.

### GetEnableFollowedTopicPushNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicPushNotificationsOk() (*bool, bool)`

GetEnableFollowedTopicPushNotificationsOk returns a tuple with the EnableFollowedTopicPushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFollowedTopicPushNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableFollowedTopicPushNotifications(v bool)`

SetEnableFollowedTopicPushNotifications sets EnableFollowedTopicPushNotifications field to given value.

### HasEnableFollowedTopicPushNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableFollowedTopicPushNotifications() bool`

HasEnableFollowedTopicPushNotifications returns a boolean if a field has been set.

### GetEnableFollowedTopicAudibleNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicAudibleNotifications() bool`

GetEnableFollowedTopicAudibleNotifications returns the EnableFollowedTopicAudibleNotifications field if non-nil, zero value otherwise.

### GetEnableFollowedTopicAudibleNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicAudibleNotificationsOk() (*bool, bool)`

GetEnableFollowedTopicAudibleNotificationsOk returns a tuple with the EnableFollowedTopicAudibleNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFollowedTopicAudibleNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableFollowedTopicAudibleNotifications(v bool)`

SetEnableFollowedTopicAudibleNotifications sets EnableFollowedTopicAudibleNotifications field to given value.

### HasEnableFollowedTopicAudibleNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableFollowedTopicAudibleNotifications() bool`

HasEnableFollowedTopicAudibleNotifications returns a boolean if a field has been set.

### GetEmailNotificationsBatchingPeriodSeconds

`func (o *RegisterQueueResponseUserSettings) GetEmailNotificationsBatchingPeriodSeconds() int32`

GetEmailNotificationsBatchingPeriodSeconds returns the EmailNotificationsBatchingPeriodSeconds field if non-nil, zero value otherwise.

### GetEmailNotificationsBatchingPeriodSecondsOk

`func (o *RegisterQueueResponseUserSettings) GetEmailNotificationsBatchingPeriodSecondsOk() (*int32, bool)`

GetEmailNotificationsBatchingPeriodSecondsOk returns a tuple with the EmailNotificationsBatchingPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotificationsBatchingPeriodSeconds

`func (o *RegisterQueueResponseUserSettings) SetEmailNotificationsBatchingPeriodSeconds(v int32)`

SetEmailNotificationsBatchingPeriodSeconds sets EmailNotificationsBatchingPeriodSeconds field to given value.

### HasEmailNotificationsBatchingPeriodSeconds

`func (o *RegisterQueueResponseUserSettings) HasEmailNotificationsBatchingPeriodSeconds() bool`

HasEmailNotificationsBatchingPeriodSeconds returns a boolean if a field has been set.

### GetEnableOfflineEmailNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableOfflineEmailNotifications() bool`

GetEnableOfflineEmailNotifications returns the EnableOfflineEmailNotifications field if non-nil, zero value otherwise.

### GetEnableOfflineEmailNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableOfflineEmailNotificationsOk() (*bool, bool)`

GetEnableOfflineEmailNotificationsOk returns a tuple with the EnableOfflineEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflineEmailNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableOfflineEmailNotifications(v bool)`

SetEnableOfflineEmailNotifications sets EnableOfflineEmailNotifications field to given value.

### HasEnableOfflineEmailNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableOfflineEmailNotifications() bool`

HasEnableOfflineEmailNotifications returns a boolean if a field has been set.

### GetEnableOfflinePushNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableOfflinePushNotifications() bool`

GetEnableOfflinePushNotifications returns the EnableOfflinePushNotifications field if non-nil, zero value otherwise.

### GetEnableOfflinePushNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableOfflinePushNotificationsOk() (*bool, bool)`

GetEnableOfflinePushNotificationsOk returns a tuple with the EnableOfflinePushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflinePushNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableOfflinePushNotifications(v bool)`

SetEnableOfflinePushNotifications sets EnableOfflinePushNotifications field to given value.

### HasEnableOfflinePushNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableOfflinePushNotifications() bool`

HasEnableOfflinePushNotifications returns a boolean if a field has been set.

### GetEnableOnlinePushNotifications

`func (o *RegisterQueueResponseUserSettings) GetEnableOnlinePushNotifications() bool`

GetEnableOnlinePushNotifications returns the EnableOnlinePushNotifications field if non-nil, zero value otherwise.

### GetEnableOnlinePushNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableOnlinePushNotificationsOk() (*bool, bool)`

GetEnableOnlinePushNotificationsOk returns a tuple with the EnableOnlinePushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOnlinePushNotifications

`func (o *RegisterQueueResponseUserSettings) SetEnableOnlinePushNotifications(v bool)`

SetEnableOnlinePushNotifications sets EnableOnlinePushNotifications field to given value.

### HasEnableOnlinePushNotifications

`func (o *RegisterQueueResponseUserSettings) HasEnableOnlinePushNotifications() bool`

HasEnableOnlinePushNotifications returns a boolean if a field has been set.

### GetEnableDigestEmails

`func (o *RegisterQueueResponseUserSettings) GetEnableDigestEmails() bool`

GetEnableDigestEmails returns the EnableDigestEmails field if non-nil, zero value otherwise.

### GetEnableDigestEmailsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableDigestEmailsOk() (*bool, bool)`

GetEnableDigestEmailsOk returns a tuple with the EnableDigestEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDigestEmails

`func (o *RegisterQueueResponseUserSettings) SetEnableDigestEmails(v bool)`

SetEnableDigestEmails sets EnableDigestEmails field to given value.

### HasEnableDigestEmails

`func (o *RegisterQueueResponseUserSettings) HasEnableDigestEmails() bool`

HasEnableDigestEmails returns a boolean if a field has been set.

### GetEnableMarketingEmails

`func (o *RegisterQueueResponseUserSettings) GetEnableMarketingEmails() bool`

GetEnableMarketingEmails returns the EnableMarketingEmails field if non-nil, zero value otherwise.

### GetEnableMarketingEmailsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableMarketingEmailsOk() (*bool, bool)`

GetEnableMarketingEmailsOk returns a tuple with the EnableMarketingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMarketingEmails

`func (o *RegisterQueueResponseUserSettings) SetEnableMarketingEmails(v bool)`

SetEnableMarketingEmails sets EnableMarketingEmails field to given value.

### HasEnableMarketingEmails

`func (o *RegisterQueueResponseUserSettings) HasEnableMarketingEmails() bool`

HasEnableMarketingEmails returns a boolean if a field has been set.

### GetEnableLoginEmails

`func (o *RegisterQueueResponseUserSettings) GetEnableLoginEmails() bool`

GetEnableLoginEmails returns the EnableLoginEmails field if non-nil, zero value otherwise.

### GetEnableLoginEmailsOk

`func (o *RegisterQueueResponseUserSettings) GetEnableLoginEmailsOk() (*bool, bool)`

GetEnableLoginEmailsOk returns a tuple with the EnableLoginEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLoginEmails

`func (o *RegisterQueueResponseUserSettings) SetEnableLoginEmails(v bool)`

SetEnableLoginEmails sets EnableLoginEmails field to given value.

### HasEnableLoginEmails

`func (o *RegisterQueueResponseUserSettings) HasEnableLoginEmails() bool`

HasEnableLoginEmails returns a boolean if a field has been set.

### GetMessageContentInEmailNotifications

`func (o *RegisterQueueResponseUserSettings) GetMessageContentInEmailNotifications() bool`

GetMessageContentInEmailNotifications returns the MessageContentInEmailNotifications field if non-nil, zero value otherwise.

### GetMessageContentInEmailNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetMessageContentInEmailNotificationsOk() (*bool, bool)`

GetMessageContentInEmailNotificationsOk returns a tuple with the MessageContentInEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContentInEmailNotifications

`func (o *RegisterQueueResponseUserSettings) SetMessageContentInEmailNotifications(v bool)`

SetMessageContentInEmailNotifications sets MessageContentInEmailNotifications field to given value.

### HasMessageContentInEmailNotifications

`func (o *RegisterQueueResponseUserSettings) HasMessageContentInEmailNotifications() bool`

HasMessageContentInEmailNotifications returns a boolean if a field has been set.

### GetPmContentInDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) GetPmContentInDesktopNotifications() bool`

GetPmContentInDesktopNotifications returns the PmContentInDesktopNotifications field if non-nil, zero value otherwise.

### GetPmContentInDesktopNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetPmContentInDesktopNotificationsOk() (*bool, bool)`

GetPmContentInDesktopNotificationsOk returns a tuple with the PmContentInDesktopNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmContentInDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) SetPmContentInDesktopNotifications(v bool)`

SetPmContentInDesktopNotifications sets PmContentInDesktopNotifications field to given value.

### HasPmContentInDesktopNotifications

`func (o *RegisterQueueResponseUserSettings) HasPmContentInDesktopNotifications() bool`

HasPmContentInDesktopNotifications returns a boolean if a field has been set.

### GetWildcardMentionsNotify

`func (o *RegisterQueueResponseUserSettings) GetWildcardMentionsNotify() bool`

GetWildcardMentionsNotify returns the WildcardMentionsNotify field if non-nil, zero value otherwise.

### GetWildcardMentionsNotifyOk

`func (o *RegisterQueueResponseUserSettings) GetWildcardMentionsNotifyOk() (*bool, bool)`

GetWildcardMentionsNotifyOk returns a tuple with the WildcardMentionsNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardMentionsNotify

`func (o *RegisterQueueResponseUserSettings) SetWildcardMentionsNotify(v bool)`

SetWildcardMentionsNotify sets WildcardMentionsNotify field to given value.

### HasWildcardMentionsNotify

`func (o *RegisterQueueResponseUserSettings) HasWildcardMentionsNotify() bool`

HasWildcardMentionsNotify returns a boolean if a field has been set.

### GetEnableFollowedTopicWildcardMentionsNotify

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicWildcardMentionsNotify() bool`

GetEnableFollowedTopicWildcardMentionsNotify returns the EnableFollowedTopicWildcardMentionsNotify field if non-nil, zero value otherwise.

### GetEnableFollowedTopicWildcardMentionsNotifyOk

`func (o *RegisterQueueResponseUserSettings) GetEnableFollowedTopicWildcardMentionsNotifyOk() (*bool, bool)`

GetEnableFollowedTopicWildcardMentionsNotifyOk returns a tuple with the EnableFollowedTopicWildcardMentionsNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFollowedTopicWildcardMentionsNotify

`func (o *RegisterQueueResponseUserSettings) SetEnableFollowedTopicWildcardMentionsNotify(v bool)`

SetEnableFollowedTopicWildcardMentionsNotify sets EnableFollowedTopicWildcardMentionsNotify field to given value.

### HasEnableFollowedTopicWildcardMentionsNotify

`func (o *RegisterQueueResponseUserSettings) HasEnableFollowedTopicWildcardMentionsNotify() bool`

HasEnableFollowedTopicWildcardMentionsNotify returns a boolean if a field has been set.

### GetDesktopIconCountDisplay

`func (o *RegisterQueueResponseUserSettings) GetDesktopIconCountDisplay() int32`

GetDesktopIconCountDisplay returns the DesktopIconCountDisplay field if non-nil, zero value otherwise.

### GetDesktopIconCountDisplayOk

`func (o *RegisterQueueResponseUserSettings) GetDesktopIconCountDisplayOk() (*int32, bool)`

GetDesktopIconCountDisplayOk returns a tuple with the DesktopIconCountDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopIconCountDisplay

`func (o *RegisterQueueResponseUserSettings) SetDesktopIconCountDisplay(v int32)`

SetDesktopIconCountDisplay sets DesktopIconCountDisplay field to given value.

### HasDesktopIconCountDisplay

`func (o *RegisterQueueResponseUserSettings) HasDesktopIconCountDisplay() bool`

HasDesktopIconCountDisplay returns a boolean if a field has been set.

### GetRealmNameInEmailNotificationsPolicy

`func (o *RegisterQueueResponseUserSettings) GetRealmNameInEmailNotificationsPolicy() int32`

GetRealmNameInEmailNotificationsPolicy returns the RealmNameInEmailNotificationsPolicy field if non-nil, zero value otherwise.

### GetRealmNameInEmailNotificationsPolicyOk

`func (o *RegisterQueueResponseUserSettings) GetRealmNameInEmailNotificationsPolicyOk() (*int32, bool)`

GetRealmNameInEmailNotificationsPolicyOk returns a tuple with the RealmNameInEmailNotificationsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmNameInEmailNotificationsPolicy

`func (o *RegisterQueueResponseUserSettings) SetRealmNameInEmailNotificationsPolicy(v int32)`

SetRealmNameInEmailNotificationsPolicy sets RealmNameInEmailNotificationsPolicy field to given value.

### HasRealmNameInEmailNotificationsPolicy

`func (o *RegisterQueueResponseUserSettings) HasRealmNameInEmailNotificationsPolicy() bool`

HasRealmNameInEmailNotificationsPolicy returns a boolean if a field has been set.

### GetAutomaticallyFollowTopicsPolicy

`func (o *RegisterQueueResponseUserSettings) GetAutomaticallyFollowTopicsPolicy() int32`

GetAutomaticallyFollowTopicsPolicy returns the AutomaticallyFollowTopicsPolicy field if non-nil, zero value otherwise.

### GetAutomaticallyFollowTopicsPolicyOk

`func (o *RegisterQueueResponseUserSettings) GetAutomaticallyFollowTopicsPolicyOk() (*int32, bool)`

GetAutomaticallyFollowTopicsPolicyOk returns a tuple with the AutomaticallyFollowTopicsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyFollowTopicsPolicy

`func (o *RegisterQueueResponseUserSettings) SetAutomaticallyFollowTopicsPolicy(v int32)`

SetAutomaticallyFollowTopicsPolicy sets AutomaticallyFollowTopicsPolicy field to given value.

### HasAutomaticallyFollowTopicsPolicy

`func (o *RegisterQueueResponseUserSettings) HasAutomaticallyFollowTopicsPolicy() bool`

HasAutomaticallyFollowTopicsPolicy returns a boolean if a field has been set.

### GetAutomaticallyUnmuteTopicsInMutedStreamsPolicy

`func (o *RegisterQueueResponseUserSettings) GetAutomaticallyUnmuteTopicsInMutedStreamsPolicy() int32`

GetAutomaticallyUnmuteTopicsInMutedStreamsPolicy returns the AutomaticallyUnmuteTopicsInMutedStreamsPolicy field if non-nil, zero value otherwise.

### GetAutomaticallyUnmuteTopicsInMutedStreamsPolicyOk

`func (o *RegisterQueueResponseUserSettings) GetAutomaticallyUnmuteTopicsInMutedStreamsPolicyOk() (*int32, bool)`

GetAutomaticallyUnmuteTopicsInMutedStreamsPolicyOk returns a tuple with the AutomaticallyUnmuteTopicsInMutedStreamsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyUnmuteTopicsInMutedStreamsPolicy

`func (o *RegisterQueueResponseUserSettings) SetAutomaticallyUnmuteTopicsInMutedStreamsPolicy(v int32)`

SetAutomaticallyUnmuteTopicsInMutedStreamsPolicy sets AutomaticallyUnmuteTopicsInMutedStreamsPolicy field to given value.

### HasAutomaticallyUnmuteTopicsInMutedStreamsPolicy

`func (o *RegisterQueueResponseUserSettings) HasAutomaticallyUnmuteTopicsInMutedStreamsPolicy() bool`

HasAutomaticallyUnmuteTopicsInMutedStreamsPolicy returns a boolean if a field has been set.

### GetAutomaticallyFollowTopicsWhereMentioned

`func (o *RegisterQueueResponseUserSettings) GetAutomaticallyFollowTopicsWhereMentioned() bool`

GetAutomaticallyFollowTopicsWhereMentioned returns the AutomaticallyFollowTopicsWhereMentioned field if non-nil, zero value otherwise.

### GetAutomaticallyFollowTopicsWhereMentionedOk

`func (o *RegisterQueueResponseUserSettings) GetAutomaticallyFollowTopicsWhereMentionedOk() (*bool, bool)`

GetAutomaticallyFollowTopicsWhereMentionedOk returns a tuple with the AutomaticallyFollowTopicsWhereMentioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyFollowTopicsWhereMentioned

`func (o *RegisterQueueResponseUserSettings) SetAutomaticallyFollowTopicsWhereMentioned(v bool)`

SetAutomaticallyFollowTopicsWhereMentioned sets AutomaticallyFollowTopicsWhereMentioned field to given value.

### HasAutomaticallyFollowTopicsWhereMentioned

`func (o *RegisterQueueResponseUserSettings) HasAutomaticallyFollowTopicsWhereMentioned() bool`

HasAutomaticallyFollowTopicsWhereMentioned returns a boolean if a field has been set.

### GetResolvedTopicNoticeAutoReadPolicy

`func (o *RegisterQueueResponseUserSettings) GetResolvedTopicNoticeAutoReadPolicy() string`

GetResolvedTopicNoticeAutoReadPolicy returns the ResolvedTopicNoticeAutoReadPolicy field if non-nil, zero value otherwise.

### GetResolvedTopicNoticeAutoReadPolicyOk

`func (o *RegisterQueueResponseUserSettings) GetResolvedTopicNoticeAutoReadPolicyOk() (*string, bool)`

GetResolvedTopicNoticeAutoReadPolicyOk returns a tuple with the ResolvedTopicNoticeAutoReadPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedTopicNoticeAutoReadPolicy

`func (o *RegisterQueueResponseUserSettings) SetResolvedTopicNoticeAutoReadPolicy(v string)`

SetResolvedTopicNoticeAutoReadPolicy sets ResolvedTopicNoticeAutoReadPolicy field to given value.

### HasResolvedTopicNoticeAutoReadPolicy

`func (o *RegisterQueueResponseUserSettings) HasResolvedTopicNoticeAutoReadPolicy() bool`

HasResolvedTopicNoticeAutoReadPolicy returns a boolean if a field has been set.

### GetPresenceEnabled

`func (o *RegisterQueueResponseUserSettings) GetPresenceEnabled() bool`

GetPresenceEnabled returns the PresenceEnabled field if non-nil, zero value otherwise.

### GetPresenceEnabledOk

`func (o *RegisterQueueResponseUserSettings) GetPresenceEnabledOk() (*bool, bool)`

GetPresenceEnabledOk returns a tuple with the PresenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceEnabled

`func (o *RegisterQueueResponseUserSettings) SetPresenceEnabled(v bool)`

SetPresenceEnabled sets PresenceEnabled field to given value.

### HasPresenceEnabled

`func (o *RegisterQueueResponseUserSettings) HasPresenceEnabled() bool`

HasPresenceEnabled returns a boolean if a field has been set.

### GetAvailableNotificationSounds

`func (o *RegisterQueueResponseUserSettings) GetAvailableNotificationSounds() []string`

GetAvailableNotificationSounds returns the AvailableNotificationSounds field if non-nil, zero value otherwise.

### GetAvailableNotificationSoundsOk

`func (o *RegisterQueueResponseUserSettings) GetAvailableNotificationSoundsOk() (*[]string, bool)`

GetAvailableNotificationSoundsOk returns a tuple with the AvailableNotificationSounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNotificationSounds

`func (o *RegisterQueueResponseUserSettings) SetAvailableNotificationSounds(v []string)`

SetAvailableNotificationSounds sets AvailableNotificationSounds field to given value.

### HasAvailableNotificationSounds

`func (o *RegisterQueueResponseUserSettings) HasAvailableNotificationSounds() bool`

HasAvailableNotificationSounds returns a boolean if a field has been set.

### GetEmojisetChoices

`func (o *RegisterQueueResponseUserSettings) GetEmojisetChoices() []RegisterQueueResponseUserSettingsEmojisetChoicesInner`

GetEmojisetChoices returns the EmojisetChoices field if non-nil, zero value otherwise.

### GetEmojisetChoicesOk

`func (o *RegisterQueueResponseUserSettings) GetEmojisetChoicesOk() (*[]RegisterQueueResponseUserSettingsEmojisetChoicesInner, bool)`

GetEmojisetChoicesOk returns a tuple with the EmojisetChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojisetChoices

`func (o *RegisterQueueResponseUserSettings) SetEmojisetChoices(v []RegisterQueueResponseUserSettingsEmojisetChoicesInner)`

SetEmojisetChoices sets EmojisetChoices field to given value.

### HasEmojisetChoices

`func (o *RegisterQueueResponseUserSettings) HasEmojisetChoices() bool`

HasEmojisetChoices returns a boolean if a field has been set.

### GetSendPrivateTypingNotifications

`func (o *RegisterQueueResponseUserSettings) GetSendPrivateTypingNotifications() bool`

GetSendPrivateTypingNotifications returns the SendPrivateTypingNotifications field if non-nil, zero value otherwise.

### GetSendPrivateTypingNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetSendPrivateTypingNotificationsOk() (*bool, bool)`

GetSendPrivateTypingNotificationsOk returns a tuple with the SendPrivateTypingNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPrivateTypingNotifications

`func (o *RegisterQueueResponseUserSettings) SetSendPrivateTypingNotifications(v bool)`

SetSendPrivateTypingNotifications sets SendPrivateTypingNotifications field to given value.

### HasSendPrivateTypingNotifications

`func (o *RegisterQueueResponseUserSettings) HasSendPrivateTypingNotifications() bool`

HasSendPrivateTypingNotifications returns a boolean if a field has been set.

### GetSendStreamTypingNotifications

`func (o *RegisterQueueResponseUserSettings) GetSendStreamTypingNotifications() bool`

GetSendStreamTypingNotifications returns the SendStreamTypingNotifications field if non-nil, zero value otherwise.

### GetSendStreamTypingNotificationsOk

`func (o *RegisterQueueResponseUserSettings) GetSendStreamTypingNotificationsOk() (*bool, bool)`

GetSendStreamTypingNotificationsOk returns a tuple with the SendStreamTypingNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStreamTypingNotifications

`func (o *RegisterQueueResponseUserSettings) SetSendStreamTypingNotifications(v bool)`

SetSendStreamTypingNotifications sets SendStreamTypingNotifications field to given value.

### HasSendStreamTypingNotifications

`func (o *RegisterQueueResponseUserSettings) HasSendStreamTypingNotifications() bool`

HasSendStreamTypingNotifications returns a boolean if a field has been set.

### GetSendReadReceipts

`func (o *RegisterQueueResponseUserSettings) GetSendReadReceipts() bool`

GetSendReadReceipts returns the SendReadReceipts field if non-nil, zero value otherwise.

### GetSendReadReceiptsOk

`func (o *RegisterQueueResponseUserSettings) GetSendReadReceiptsOk() (*bool, bool)`

GetSendReadReceiptsOk returns a tuple with the SendReadReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReadReceipts

`func (o *RegisterQueueResponseUserSettings) SetSendReadReceipts(v bool)`

SetSendReadReceipts sets SendReadReceipts field to given value.

### HasSendReadReceipts

`func (o *RegisterQueueResponseUserSettings) HasSendReadReceipts() bool`

HasSendReadReceipts returns a boolean if a field has been set.

### GetAllowPrivateDataExport

`func (o *RegisterQueueResponseUserSettings) GetAllowPrivateDataExport() bool`

GetAllowPrivateDataExport returns the AllowPrivateDataExport field if non-nil, zero value otherwise.

### GetAllowPrivateDataExportOk

`func (o *RegisterQueueResponseUserSettings) GetAllowPrivateDataExportOk() (*bool, bool)`

GetAllowPrivateDataExportOk returns a tuple with the AllowPrivateDataExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrivateDataExport

`func (o *RegisterQueueResponseUserSettings) SetAllowPrivateDataExport(v bool)`

SetAllowPrivateDataExport sets AllowPrivateDataExport field to given value.

### HasAllowPrivateDataExport

`func (o *RegisterQueueResponseUserSettings) HasAllowPrivateDataExport() bool`

HasAllowPrivateDataExport returns a boolean if a field has been set.

### GetEmailAddressVisibility

`func (o *RegisterQueueResponseUserSettings) GetEmailAddressVisibility() int32`

GetEmailAddressVisibility returns the EmailAddressVisibility field if non-nil, zero value otherwise.

### GetEmailAddressVisibilityOk

`func (o *RegisterQueueResponseUserSettings) GetEmailAddressVisibilityOk() (*int32, bool)`

GetEmailAddressVisibilityOk returns a tuple with the EmailAddressVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressVisibility

`func (o *RegisterQueueResponseUserSettings) SetEmailAddressVisibility(v int32)`

SetEmailAddressVisibility sets EmailAddressVisibility field to given value.

### HasEmailAddressVisibility

`func (o *RegisterQueueResponseUserSettings) HasEmailAddressVisibility() bool`

HasEmailAddressVisibility returns a boolean if a field has been set.

### GetWebNavigateToSentMessage

`func (o *RegisterQueueResponseUserSettings) GetWebNavigateToSentMessage() bool`

GetWebNavigateToSentMessage returns the WebNavigateToSentMessage field if non-nil, zero value otherwise.

### GetWebNavigateToSentMessageOk

`func (o *RegisterQueueResponseUserSettings) GetWebNavigateToSentMessageOk() (*bool, bool)`

GetWebNavigateToSentMessageOk returns a tuple with the WebNavigateToSentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebNavigateToSentMessage

`func (o *RegisterQueueResponseUserSettings) SetWebNavigateToSentMessage(v bool)`

SetWebNavigateToSentMessage sets WebNavigateToSentMessage field to given value.

### HasWebNavigateToSentMessage

`func (o *RegisterQueueResponseUserSettings) HasWebNavigateToSentMessage() bool`

HasWebNavigateToSentMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


