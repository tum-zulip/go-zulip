package zulip

type UserSettings struct {
	// Whether time should be [displayed in 24-hour notation].  A `null` value indicates that the client should use the default time format for the user's locale.  **Changes**: Prior to Zulip 11.0 (feature level 408), `null` was not a valid value for this setting. Note that it was not possible to actually set the time format to `null` at this feature level.
	//
	// [displayed in 24-hour notation]: https://zulip.com/help/change-the-time-format
	TwentyFourHourTime *bool `json:"twenty_four_hour_time,omitempty"`
	// Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \"Always\" setting when marking messages as read.
	WebMarkReadOnScrollPolicy MarkReadOnScrollPolicy `json:"web_mark_read_on_scroll_policy,omitempty"`
	// Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \"Top unread topic in channel\" is new in Zulip 11.0 (feature level 401).  The \"List of topics\" option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \"Channel feed\" behavior.
	WebChannelDefaultView ChannelDefaultView `json:"web_channel_default_view,omitempty"`
	// Whether clients should display the [number of starred messages].
	//
	// [number of starred messages]: https://zulip.com/help/star-a-message#display-the-number-of-starred-messages
	StarredMessageCounts bool `json:"starred_message_counts,omitempty"`
	// Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.
	ReceivesTypingNotifications bool `json:"receives_typing_notifications,omitempty"`
	// Whether the user should be shown an alert, offering to update their [profile time zone], when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.  **Changes**: New in Zulip 10.0 (feature level 329).
	//
	// [profile time zone]: https://zulip.com/help/change-your-timezone
	WebSuggestUpdateTimezone bool `json:"web_suggest_update_timezone,omitempty"`
	// Whether to use the [maximum available screen width] for the web app's center panel (message feed, recent conversations) on wide screens.
	//
	// [maximum available screen width]: https://zulip.com/help/enable-full-width-display
	FluidLayoutWidth bool `json:"fluid_layout_width,omitempty"`
	// This setting is reserved for use to control variations in Zulip's design to help visually impaired users.
	HighContrastMode bool `json:"high_contrast_mode,omitempty"`
	// User-configured primary `font-size` for the web application, in pixels.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.
	WebFontSizePx int64 `json:"web_font_size_px,omitempty"`
	// User-configured primary `line-height` for the web application, in percent, so a value of 120 represents a `line-height` of 1.2.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.
	WebLineHeightPercent int `json:"web_line_height_percent,omitempty"`
	// Controls which [color theme] to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard `prefers-color-scheme` media query.
	//
	// [color theme]: https://zulip.com/help/dark-theme
	ColorScheme ColorScheme `json:"color_scheme,omitempty"`
	// Whether to [translate emoticons to emoji] in messages the user sends.
	//
	// [translate emoticons to emoji]: https://zulip.com/help/configure-emoticon-translations
	TranslateEmoticons bool `json:"translate_emoticons,omitempty"`
	// Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.  **Changes**: New in Zulip 6.0 (feature level 125).
	DisplayEmojiReactionUsers bool `json:"display_emoji_reaction_users,omitempty"`
	// What [default language] to use for the account.  This controls both the Zulip UI as well as email notifications sent to the user.  The value needs to be a standard language code that the Zulip server has translation data for; for example, `\"en\"` for English or `\"de\"` for German.
	//
	// [default language]: https://zulip.com/help/change-your-language
	DefaultLanguage string `json:"default_language,omitempty"`
	// The [home view] used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.  - \"recent_topics\" - Recent conversations view - \"inbox\" - Inbox view - \"all_messages\" - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).
	//
	// [home view]: https://zulip.com/help/configure-home-view
	WebHomeView WebHomeView `json:"web_home_view,omitempty"`
	// Whether the escape key navigates to the [configured home view].  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `escape_navigates_to_default_view`, which was new in Zulip 5.0 (feature level 107).
	//
	// [configured home view]: https://zulip.com/help/configure-home-view
	WebEscapeNavigatesToHomeView bool `json:"web_escape_navigates_to_home_view,omitempty"`
	// Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.
	LeftSideUserlist bool `json:"left_side_userlist,omitempty"`
	// The user's configured [emoji set], used to display emoji to the user everywhere they appear in the UI.  - \"google\" - Google modern - \"google-blob\" - Google classic - \"twitter\" - Twitter - \"text\" - Plain text
	//
	// [emoji set]: https://zulip.com/help/emoji-and-emoticons#use-emoticons
	Emojiset Emojiset `json:"emojiset,omitempty"`
	// Whether to [hide inactive channels] in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never
	//
	// [hide inactive channels]: https://zulip.com/help/manage-inactive-channels
	DemoteInactiveChannels DemoteInactiveChannels `json:"demote_inactive_streams,omitempty"`
	// The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).
	UserListStyle UserListStyle `json:"user_list_style,omitempty"`
	// Controls how animated images should be played in the message feed in the web/desktop application.  - \"always\" - Always play the animated images in the message feed. - \"on_hover\" - Play the animated images on hover over them in the message feed. - \"never\" - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275).
	WebAnimateImagePreviews WebAnimateImagePreviews `json:"web_animate_image_previews,omitempty"`
	// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).
	WebChannelUnreadsCountDisplayPolicy WebChannelUnreadsCountDisplayPolicy `json:"web_stream_unreads_count_display_policy,omitempty"`
	// Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).
	HideAiFeatures bool `json:"hide_ai_features,omitempty"`
	// Determines whether the web/desktop application's left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).
	WebLeftSidebarShowChannelFolders bool `json:"web_left_sidebar_show_channel_folders,omitempty"`
	// Determines whether the web/desktop application's left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).
	WebLeftSidebarUnreadsCountSummary bool `json:"web_left_sidebar_unreads_count_summary,omitempty"`
	// The IANA identifier of the user's [profile time zone], which is used primarily to display the user's local time to other users.
	//
	// [profile time zone]: https://zulip.com/help/change-your-timezone
	Timezone string `json:"timezone,omitempty"`
	// Whether the user setting for [sending on pressing Enter] in the compose box is enabled.
	//
	// [sending on pressing Enter]: https://zulip.com/help/configure-send-message-keys
	EnterSends bool `json:"enter_sends,omitempty"`
	// A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.
	EnableDraftsSynchronization bool `json:"enable_drafts_synchronization,omitempty"`
	// Enable visual desktop notifications for channel messages.
	EnableChannelDesktopNotifications bool `json:"enable_stream_desktop_notifications,omitempty"`
	// Enable email notifications for channel messages.
	EnableChannelEmailNotifications bool `json:"enable_stream_email_notifications,omitempty"`
	// Enable mobile notifications for channel messages.
	EnableChannelPushNotifications bool `json:"enable_stream_push_notifications,omitempty"`
	// Enable audible desktop notifications for channel messages.
	EnableChannelAudibleNotifications bool `json:"enable_stream_audible_notifications,omitempty"`
	// Notification sound name.
	NotificationSound string `json:"notification_sound,omitempty"`
	// Enable visual desktop notifications for direct messages and @-mentions.
	EnableDesktopNotifications bool `json:"enable_desktop_notifications,omitempty"`
	// Enable audible desktop notifications for direct messages and @-mentions.
	EnableSounds bool `json:"enable_sounds,omitempty"`
	// Enable visual desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicDesktopNotifications bool `json:"enable_followed_topic_desktop_notifications,omitempty"`
	// Enable email notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicEmailNotifications bool `json:"enable_followed_topic_email_notifications,omitempty"`
	// Enable push notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicPushNotifications bool `json:"enable_followed_topic_push_notifications,omitempty"`
	// Enable audible desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicAudibleNotifications bool `json:"enable_followed_topic_audible_notifications,omitempty"`
	// The duration (in seconds) for which the server should wait to batch email notifications before sending them.
	EmailNotificationsBatchingPeriodSeconds int `json:"email_notifications_batching_period_seconds,omitempty"`
	// Enable email notifications for direct messages and @-mentions received when the user is offline.
	EnableOfflineEmailNotifications bool `json:"enable_offline_email_notifications,omitempty"`
	// Enable mobile notification for direct messages and @-mentions received when the user is offline.
	EnableOfflinePushNotifications bool `json:"enable_offline_push_notifications,omitempty"`
	// Enable mobile notification for direct messages and @-mentions received when the user is online.
	EnableOnlinePushNotifications bool `json:"enable_online_push_notifications,omitempty"`
	// Enable digest emails when the user is away.
	EnableDigestEmails bool `json:"enable_digest_emails,omitempty"`
	// Enable marketing emails. Has no function outside Zulip Cloud.
	EnableMarketingEmails bool `json:"enable_marketing_emails,omitempty"`
	// Enable email notifications for new logins to account.
	EnableLoginEmails bool `json:"enable_login_emails,omitempty"`
	// Include the message's content in email notifications for new messages.
	MessageContentInEmailNotifications bool `json:"message_content_in_email_notifications,omitempty"`
	// Include content of direct messages in desktop notifications.
	PmContentInDesktopNotifications bool `json:"pm_content_in_desktop_notifications,omitempty"`
	// Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.
	WildcardMentionsNotify bool `json:"wildcard_mentions_notify,omitempty"`
	// Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicWildcardMentionsNotify bool `json:"enable_followed_topic_wildcard_mentions_notify,omitempty"`
	// Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.
	DesktopIconCountDisplay DesktopIconCountDisplay `json:"desktop_icon_count_display,omitempty"`
	// Whether to [include organization name in subject of message notification emails].  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.
	//
	// [include organization name in subject of message notification emails]: https://zulip.com/help/email-notifications#include-organization-name-in-subject-line
	RealmNameInEmailNotificationsPolicy NameInEmailNotificationsPolicy `json:"realm_name_in_email_notifications_policy,omitempty"`
	// Which [topics to follow automatically].  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).
	//
	// [topics to follow automatically]: https://zulip.com/help/mute-a-topic
	AutomaticallyFollowTopicsPolicy AutomaticallyFollowTopicsPolicy `json:"automatically_follow_topics_policy,omitempty"`
	// Which [topics to unmute automatically in muted channels].  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).
	//
	// [topics to unmute automatically in muted channels]: https://zulip.com/help/mute-a-topic
	AutomaticallyUnmuteTopicsInMutedChannelsPolicy AutomaticallyUnmuteTopicsInMutedChannelsPolicy `json:"automatically_unmute_topics_in_muted_streams_policy,omitempty"`
	// Whether the server will automatically mark the user as following topics where the user is mentioned.  **Changes**: New in Zulip 8.0 (feature level 235).
	AutomaticallyFollowTopicsWhereMentioned bool `json:"automatically_follow_topics_where_mentioned,omitempty"`
	// Controls whether the resolved-topic notices are marked as read.  - \"always\" - Always mark resolved-topic notices as read. - \"except_followed\" - Mark resolved-topic notices as read in topics not followed by the user. - \"never\" - Never mark resolved-topic notices as read.  **Changes**: New in Zulip 11.0 (feature level 385).
	ResolvedTopicNoticeAutoReadPolicy ResolvedTopicNoticeAutoReadPolicy `json:"resolved_topic_notice_auto_read_policy,omitempty"`
	// Display the presence status to other users when online.
	PresenceEnabled bool `json:"presence_enabled,omitempty"`
	// Array containing the names of the notification sound options supported by this Zulip server. Only relevant to support UI for configuring notification sounds.
	AvailableNotificationSounds []string `json:"available_notification_sounds,omitempty"`
	// Array of dictionaries where each dictionary describes an emoji set supported by this version of the Zulip server.  Only relevant to clients with configuration UI for choosing an emoji set; the currently selected emoji set is available in the `emojiset` key.  See [PATCH /settings] for details on the meaning of this setting.
	//
	// [PATCH /settings]: https://zulip.com/api/update-settings
	EmojisetChoices []UserSettingsEmojisetChoice `json:"emojiset_choices,omitempty"`
	// Whether the user has chosen to send [typing notifications] when composing direct messages. The client should send typing notifications for direct messages if and only if this setting is enabled.  **Changes**: New in Zulip 5.0 (feature level 105).
	//
	// [typing notifications]: https://zulip.com/help/typing-notifications
	SendPrivateTypingNotifications bool `json:"send_private_typing_notifications,omitempty"`
	// Whether the user has chosen to send [typing notifications] when composing channel messages. The client should send typing notifications for channel messages if and only if this setting is enabled.  **Changes**: New in Zulip 5.0 (feature level 105).
	//
	// [typing notifications]: https://zulip.com/help/typing-notifications
	SendChannelTypingNotifications bool `json:"send_stream_typing_notifications,omitempty"`
	// Whether other users are allowed to see whether you've read messages.  **Changes**: New in Zulip 5.0 (feature level 105).
	SendReadReceipts bool `json:"send_read_receipts,omitempty"`
	// Whether organization administrators are allowed to export your private data.  **Changes**: New in Zulip 10.0 (feature level 293).
	AllowPrivateDataExport bool `json:"allow_private_data_export,omitempty"`
	// The [policy] for [which other users] in this organization can see the user's real email address.
	//   - 1 = Everyone
	//   - 2 = Members only
	//   - 3 = Administrators only
	//   - 4 = Nobody
	//   - 5 = Moderators only
	//
	// **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.
	//
	// [policy]: https://zulip.com/api/roles-and-permissions#permission-levels
	// [which other users]: https://zulip.com/help/configure-email-visibility
	EmailAddressVisibility EmailAddressVisibility `json:"email_address_visibility,omitempty"`
	// Web/desktop app setting for whether the user's view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.
	WebNavigateToSentMessage bool `json:"web_navigate_to_sent_message,omitempty"`
}
