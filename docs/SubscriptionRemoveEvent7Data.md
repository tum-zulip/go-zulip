# SubscriptionRemoveEvent7Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMessageEditing** | Pointer to **bool** | Whether this organization&#39;s [message edit policy][config-message-editing] allows editing the content of messages.  See [&#x60;PATCH /messages/{message_id}&#x60;](zulip.com/api/update-message for details and history of how message editing permissions work.  [config-message-editing]: /help/restrict-message-editing-and-deletion  | [optional] 
**AuthenticationMethods** | Pointer to [**map[string]RealmAuthenticationMethod**](RealmAuthenticationMethod.md) | Dictionary of authentication method keys mapped to dictionaries that describe the properties of the named authentication method for the organization - its enabled status and availability for use by the organization.  Clients should use this to implement server-settings UI to change which methods are enabled for the organization. For authentication UI itself, clients should use the pre-authentication metadata returned by [&#x60;GET /server_settings&#x60;](zulip.com/api/get-server-settings.  **Changes**: In Zulip 9.0 (feature level 243), the values in this dictionary were changed. Previously, the values were a simple boolean indicating whether the backend is enabled or not.  | [optional] 
**CanAccessAllUsersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who are allowed to access all users in the organization.  **Changes**: Prior to Zulip 10.0 (feature level 314), this value used to be of type integer and did not accept anonymous user groups.  New in Zulip 8.0 (feature level 225).  | [optional] 
**CanCreateGroups** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to create user groups in this organization.  **Changes**: New in Zulip 10.0 (feature level 299). Previously &#x60;user_group_edit_policy&#x60; field used to control the permission to create user groups.  | [optional] 
**CanCreateBotsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to create all types of bot users in the organization. See also &#x60;can_create_write_only_bots_group&#x60;.  **Changes**: New in Zulip 10.0 (feature level 344). Previously, this permission was controlled by the enum &#x60;bot_creation_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Generic bots limited to administrators, 3&#x3D;Administrators.  | [optional] 
**CanCreateWriteOnlyBotsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to create bot users that can only send messages in the organization, i.e. incoming webhooks, in addition to the users who are present in &#x60;can_create_bots_group&#x60;.  **Changes**: New in Zulip 10.0 (feature level 344). Previously, this permission was controlled by the enum &#x60;bot_creation_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Generic bots limited to administrators, 3&#x3D;Administrators.  | [optional] 
**CanCreatePublicChannelGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to create public channels in this organization.  **Changes**: New in Zulip 9.0 (feature level 264). Previously &#x60;realm_create_public_stream_policy&#x60; field used to control the permission to create public channels.  | [optional] 
**CanCreatePrivateChannelGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to create private channels in this organization.  **Changes**: New in Zulip 9.0 (feature level 266). Previously &#x60;realm_create_private_stream_policy&#x60; field used to control the permission to create private channels.  | [optional] 
**CanCreateWebPublicChannelGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to create web-public channels in this organization.  **Changes**: New in Zulip 10.0 (feature level 280). Previously &#x60;realm_create_web_public_stream_policy&#x60; field used to control the permission to create web-public channels.  | [optional] 
**CanAddCustomEmojiGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to add custom emoji in the organization.  **Changes**: New in Zulip 10.0 (feature level 307). Previously, this permission was controlled by the enum &#x60;add_custom_emoji_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators.  Before Zulip 5.0 (feature level 85), the &#x60;realm_add_emoji_by_admins_only&#x60; boolean setting controlled this permission; &#x60;true&#x60; corresponded to &#x60;Admins&#x60;, and &#x60;false&#x60; to &#x60;Everyone&#x60;.  | [optional] 
**CanAddSubscribersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to add subscribers to channels in the organization.  **Changes**: New in Zulip 10.0 (feature level 341). Previously, this permission was controlled by the enum &#x60;invite_to_stream_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators.  | [optional] 
**CanDeleteAnyMessageGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to delete any message in the organization.  **Changes**: New in Zulip 10.0 (feature level 281). Previously, this permission was limited to administrators only and was uneditable.  | [optional] 
**CanDeleteOwnMessageGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to delete messages that they have sent in the organization.  **Changes**: New in Zulip 10.0 (feature level 291). Previously, this permission was controlled by the enum &#x60;delete_own_message_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators, 5&#x3D;Everyone.  Before Zulip 5.0 (feature level 101), the &#x60;allow_message_deleting&#x60; boolean setting controlled this permission; &#x60;true&#x60; corresponded to &#x60;Everyone&#x60;, and &#x60;false&#x60; to &#x60;Admins&#x60;.  | [optional] 
**CanSetDeleteMessagePolicyGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to change per-channel &#x60;can_delete_any_message_group&#x60; and &#x60;can_delete_own_message_group&#x60; permission settings. Note that the user must be a member of both this group and the &#x60;can_administer_channel_group&#x60; of the channel whose message delete settings they want to change.  Organization administrators can always change these settings of every channel.  **Changes**: New in Zulip 11.0 (feature level 407).  | [optional] 
**CanSetTopicsPolicyGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to change per-channel &#x60;topics_policy&#x60; setting. Note that the user must be a member of both this group and the &#x60;can_administer_channel_group&#x60; of the channel whose &#x60;topics_policy&#x60; they want to change.  Organization administrators can always change the &#x60;topics_policy&#x60; setting of every channel.  **Changes**: New in Zulip 11.0 (feature level 392).  | [optional] 
**CanInviteUsersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to send email invitations for inviting other users to the organization.  **Changes**: New in Zulip 10.0 (feature level 321). Previously, this permission was controlled by the enum &#x60;invite_to_realm_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators, 6&#x3D;Nobody.  Before Zulip 4.0 (feature level 50), the &#x60;invite_by_admins_only&#x60; boolean setting controlled this permission; &#x60;true&#x60; corresponded to &#x60;Admins&#x60;, and &#x60;false&#x60; to &#x60;Members&#x60;.  | [optional] 
**CanMentionManyUsersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to use wildcard mentions in large channels.  All users will receive a warning/reminder when using mentions in large channels, even when permitted to do so.  **Changes**: New in Zulip 10.0 (feature level 352). Previously, this permission was controlled by the enum &#x60;wildcard_mention_policy&#x60;.  | [optional] 
**CanMoveMessagesBetweenChannelsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to move messages from one channel to another in the organization.  **Changes**: New in Zulip 10.0 (feature level 310). Previously, this permission was controlled by the enum &#x60;move_messages_between_streams_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators, 6&#x3D;Nobody.  In Zulip 7.0 (feature level 159), &#x60;Nobody&#x60; was added as an option to &#x60;move_messages_between_streams_policy&#x60; enum.  | [optional] 
**CanMoveMessagesBetweenTopicsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to move messages from one topic to another within a channel in the organization.  **Changes**: New in Zulip 10.0 (feature level 316). Previously, this permission was controlled by the enum &#x60;edit_topic_policy&#x60;. Values were 1&#x3D;Members, 2&#x3D;Admins, 3&#x3D;Full members, 4&#x3D;Moderators, 5&#x3D;Everyone, 6&#x3D;Nobody.  In Zulip 7.0 (feature level 159), &#x60;Nobody&#x60; was added as an option to &#x60;edit_topic_policy&#x60; enum.  | [optional] 
**CanResolveTopicsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to [resolve topics](zulip.com/help/resolve-a-topic in the organization.  **Changes**: New in Zulip 10.0 (feature level 367). Previously, permission to resolve topics was controlled by the more general can_move_messages_between_topics_group permission for moving messages.  | [optional] 
**CanManageAllGroups** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to administer all existing groups in this organization.  **Changes**: Prior to Zulip 10.0 (feature level 305), only users who were a member of the group or had the moderator role or above could exercise the permission on a given group.  New in Zulip 10.0 (feature level 299). Previously the &#x60;user_group_edit_policy&#x60; field controlled the permission to manage user groups. Valid values were as follows:  - 1 &#x3D; All members can create and edit user groups - 2 &#x3D; Only organization administrators can create and edit   user groups - 3 &#x3D; Only [full members][calc-full-member] can create and   edit user groups. - 4 &#x3D; Only organization administrators and moderators can   create and edit user groups.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**CanManageBillingGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to manage plans and billing in the organization.  **Changes**: New in Zulip 10.0 (feature level 363). Previously, only owners and users with &#x60;is_billing_admin&#x60; property set to &#x60;true&#x60; were allowed to manage plans and billing.  | [optional] 
**CanSummarizeTopicsGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who are allowed to use AI summarization.  **Changes**: New in Zulip 10.0 (feature level 350).  | [optional] 
**CreateMultiuseInviteGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who are allowed to create [reusable invitation links](zulip.com/help/invite-new-users#create-a-reusable-invitation-link to the organization.  **Changes**: Prior to Zulip 10.0 (feature level 314), this value used to be of type integer and did not accept anonymous user groups.  New in Zulip 8.0 (feature level 209).  | [optional] 
**DefaultCodeBlockLanguage** | Pointer to **string** | The default pygments language code to be used for code blocks in this organization. If an empty string, no default has been set.  **Changes**: Prior to Zulip 8.0 (feature level 195), a server bug meant that both &#x60;null&#x60; and an empty string could represent that no default was set for this realm setting in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue response. The documentation for both that endpoint and this event incorrectly stated that the only representation for no default language was &#x60;null&#x60;. This event in fact uses the empty string to indicate that no default has been set in all server versions.  | [optional] 
**DefaultLanguage** | Pointer to **string** | The default language for the organization.  | [optional] 
**Description** | Pointer to **string** | The description of the organization, used on login and registration pages.  | [optional] 
**DigestEmailsEnabled** | Pointer to **bool** | Whether the organization has enabled [weekly digest emails](zulip.com/help/digest-emails.  | [optional] 
**DigestWeekday** | Pointer to **int32** | The day of the week when the organization will send its weekly digest email to inactive users.  | [optional] 
**DirectMessageInitiatorGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to start a new direct message conversation involving other non-bot users. Users who are outside this group and attempt to send the first direct message to a given collection of recipient users will receive an error, unless all other recipients are bots or the sender.  **Changes**: New in Zulip 9.0 (feature level 270).  Previously, access to send direct messages was controlled by the &#x60;private_message_policy&#x60; realm setting, which supported values of 1 (enabled) and 2 (disabled).  | [optional] 
**DirectMessagePermissionGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value](zulip.com/api/group-setting-values defining the set of users who have permission to fully use direct messages. Users outside this group can only send direct messages to conversations where all the recipients are in this group, are bots, or are the sender, ensuring that every direct message conversation will be visible to at least one user in this group.  **Changes**: New in Zulip 9.0 (feature level 270).  Previously, access to send direct messages was controlled by the &#x60;private_message_policy&#x60; realm setting, which supported values of 1 (enabled) and 2 (disabled).  | [optional] 
**DisallowDisposableEmailAddresses** | Pointer to **bool** | Whether the organization disallows disposable email addresses.  | [optional] 
**EmailChangesDisabled** | Pointer to **bool** | Whether users are allowed to change their own email address in this organization. This is typically disabled for organizations that synchronize accounts from LDAP or a similar corporate database.  | [optional] 
**EnableReadReceipts** | Pointer to **bool** | Whether read receipts is enabled in the organization or not.  If disabled, read receipt data will be unavailable to clients, regardless of individual users&#39; personal read receipt settings. See also the &#x60;send_read_receipts&#x60; setting within &#x60;realm_user_settings_defaults&#x60;.  **Changes**: New in Zulip 6.0 (feature level 137).  | [optional] 
**EmailsRestrictedToDomains** | Pointer to **bool** | Whether [new users joining](zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions this organization are required to have an email address in one of the &#x60;realm_domains&#x60; configured for the organization.  | [optional] 
**EnableGuestUserDmWarning** | Pointer to **bool** | Whether clients should show a warning when a user is composing a DM to a guest user in this organization.  **Changes**: New in Zulip 10.0 (feature level 348).  | [optional] 
**EnableGuestUserIndicator** | Pointer to **bool** | Whether clients should display \&quot;(guest)\&quot; after the names of guest users to prominently highlight their status.  **Changes**: New in Zulip 8.0 (feature level 216).  | [optional] 
**EnableSpectatorAccess** | Pointer to **bool** | Whether web-public channels are enabled in this organization.  Can only be enabled if the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; [server setting][server-settings] is enabled on the Zulip server. See also the &#x60;can_create_web_public_channel_group&#x60; realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  **Changes**: New in Zulip 5.0 (feature level 109).  | [optional] 
**GiphyRating** | Pointer to **int32** | Maximum rating of the GIFs that will be retrieved from GIPHY.  **Changes**: New in Zulip 4.0 (feature level 55).  | [optional] 
**IconSource** | Pointer to **string** | String indicating whether the organization&#39;s [profile icon](zulip.com/help/create-your-organization-profile was uploaded by a user or is the default. Useful for UI allowing editing the organization&#39;s icon.  - \&quot;G\&quot; means generated by Gravatar (the default). - \&quot;U\&quot; means uploaded by an organization administrator.  | [optional] 
**IconUrl** | Pointer to **string** | The URL of the organization&#39;s [profile icon](zulip.com/help/create-your-organization-profile.  | [optional] 
**InlineImagePreview** | Pointer to **bool** | Whether this organization has been configured to enable [previews of linked images](zulip.com/help/image-video-and-website-previews.  | [optional] 
**InlineUrlEmbedPreview** | Pointer to **bool** | Whether this organization has been configured to enable [previews of linked websites](zulip.com/help/image-video-and-website-previews.  | [optional] 
**InviteRequired** | Pointer to **bool** | Whether an invitation is required to join this organization.  | [optional] 
**JitsiServerUrl** | Pointer to **NullableString** | The URL of the custom Jitsi Meet server configured in this organization&#39;s settings.  &#x60;null&#x60;, the default, means that the organization is using the should use the server-level configuration, &#x60;server_jitsi_server_url&#x60;.  **Changes**: New in Zulip 8.0 (feature level 212). Previously, this was only available as a server-level configuration, and required a server restart to change.  | [optional] 
**LogoSource** | Pointer to **string** | String indicating whether the organization&#39;s [profile wide logo](zulip.com/help/create-your-organization-profile was uploaded by a user or is the default. Useful for UI allowing editing the organization&#39;s wide logo.  - \&quot;D\&quot; means the logo is the default Zulip logo. - \&quot;U\&quot; means uploaded by an organization administrator.  | [optional] 
**LogoUrl** | Pointer to **string** | The URL of the organization&#39;s wide logo configured in the [organization profile](zulip.com/help/create-your-organization-profile.  | [optional] 
**TopicsPolicy** | Pointer to **string** | The organization&#39;s default policy for sending channel messages to the [empty \&quot;general chat\&quot; topic](zulip.com/help/general-chat-topic.  - &#x60;\&quot;allow_empty_topic\&quot;&#x60;: Channel messages can be sent to the empty topic. - &#x60;\&quot;disable_empty_topic\&quot;&#x60;: Channel messages cannot be sent to the empty topic.  **Changes**: New in Zulip 11.0 (feature level 392). Previously, this was controlled by the boolean realm &#x60;mandatory_topics&#x60; setting, which is now deprecated.  | [optional] 
**MandatoryTopics** | Pointer to **bool** | Whether [topics are required](zulip.com/help/require-topics for messages in this organization.  **Changes**: Deprecated in Zulip 11.0 (feature level 392). This is now controlled by the realm &#x60;topics_policy&#x60; setting.  | [optional] 
**MaxFileUploadSizeMib** | Pointer to **int32** | The new maximum file size that can be uploaded to this Zulip organization.  **Changes**: New in Zulip 10.0 (feature level 306). Previously, this field of the core state did not support being updated via the events system, as it was typically hardcoded for a given Zulip installation.  | [optional] 
**MessageContentAllowedInEmailNotifications** | Pointer to **bool** | Whether notification emails in this organization are allowed to contain Zulip the message content, or simply indicate that a new message was sent.  | [optional] 
**MessageContentDeleteLimitSeconds** | Pointer to **NullableInt32** | Messages sent more than this many seconds ago cannot be deleted with this organization&#39;s [message deletion policy](zulip.com/help/restrict-message-editing-and-deletion.  Will not be 0. A &#x60;null&#x60; value means no limit: messages can be deleted regardless of how long ago they were sent.  **Changes**: No limit was represented using the special value &#x60;0&#x60; before Zulip 5.0 (feature level 100).  | [optional] 
**MessageContentEditLimitSeconds** | Pointer to **NullableInt32** | Messages sent more than this many seconds ago cannot be edited with this organization&#39;s [message edit policy](zulip.com/help/restrict-message-editing-and-deletion.  Will not be &#x60;0&#x60;. A &#x60;null&#x60; value means no limit, so messages can be edited regardless of how long ago they were sent.  See [&#x60;PATCH /messages/{message_id}&#x60;](zulip.com/api/update-message for details and history of how message editing permissions work.  **Changes**: Before Zulip 6.0 (feature level 138), no limit was represented using the special value &#x60;0&#x60;.  | [optional] 
**MessageEditHistoryVisibilityPolicy** | Pointer to **string** | Which type of message edit history is configured to allow users to access [message edit history](zulip.com/help/view-a-messages-edit-history.  - \&quot;all\&quot; &#x3D; All edit history is visible. - \&quot;moves\&quot; &#x3D; Only moves are visible. - \&quot;none\&quot; &#x3D; No edit history is visible.  **Changes**: New in Zulip 10.0 (feature level 358), replacing the previous &#x60;allow_edit_history&#x60; boolean setting; &#x60;true&#x60; corresponds to &#x60;all&#x60;, and &#x60;false&#x60; to &#x60;none&#x60;.  | [optional] 
**ModerationRequestChannelId** | Pointer to **int32** | The ID of the private channel to which messages flagged by users for moderation are sent. Moderators can use this channel to review and act on reported content.  Will be &#x60;-1&#x60; if moderation requests are disabled.  Clients should check whether moderation requests are disabled to determine whether to present a \&quot;report message\&quot; feature in their UI within a given organization.  **Changes**: New in Zulip 10.0 (feature level 331). Previously, no \&quot;report message\&quot; features existed in Zulip.  | [optional] 
**MoveMessagesWithinStreamLimitSeconds** | Pointer to **NullableInt32** | Messages sent more than this many seconds ago cannot be moved within a channel to another topic by users who have permission to do so based on this organization&#39;s [topic edit policy](zulip.com/help/restrict-moving-messages. This setting does not affect moderators and administrators.  Will not be &#x60;0&#x60;. A &#x60;null&#x60; value means no limit, so message topics can be edited regardless of how long ago they were sent.  See [&#x60;PATCH /messages/{message_id}&#x60;](zulip.com/api/update-message for details and history of how message editing permissions work.  **Changes**: New in Zulip 7.0 (feature level 162). Previously, this time limit was always 72 hours for users who were not administrators or moderators.  | [optional] 
**MoveMessagesBetweenStreamsLimitSeconds** | Pointer to **NullableInt32** | Messages sent more than this many seconds ago cannot be moved between channels by users who have permission to do so based on this organization&#39;s [message move policy](zulip.com/help/restrict-moving-messages. This setting does not affect moderators and administrators.  Will not be &#x60;0&#x60;. A &#x60;null&#x60; value means no limit, so messages can be moved regardless of how long ago they were sent.  See [&#x60;PATCH /messages/{message_id}&#x60;](zulip.com/api/update-message for details and history of how message editing permissions work.  **Changes**: New in Zulip 7.0 (feature level 162). Previously, there was no time limit for moving messages between channels for users with permission to do so.  | [optional] 
**Name** | Pointer to **string** | The name of the organization, used in login pages etc.  | [optional] 
**NameChangesDisabled** | Pointer to **bool** | Indicates whether users are [allowed to change](zulip.com/help/restrict-name-and-email-changes their name via the Zulip UI in this organization. Typically disabled in organizations syncing this type of account information from an external user database like LDAP.  | [optional] 
**NightLogoSource** | Pointer to **string** | String indicating whether the organization&#39;s dark theme [profile wide logo](zulip.com/help/create-your-organization-profile was uploaded by a user or is the default. Useful for UI allowing editing the organization&#39;s wide logo.  - \&quot;D\&quot; means the logo is the default Zulip logo. - \&quot;U\&quot; means uploaded by an organization administrator.  | [optional] 
**NightLogoUrl** | Pointer to **string** | The URL of the organization&#39;s dark theme wide-format logo configured in the [organization profile](zulip.com/help/create-your-organization-profile.  | [optional] 
**NewStreamAnnouncementsStreamId** | Pointer to **int32** | The ID of the channel to which automated messages announcing the [creation of new channels][new-channel-announce] are sent.  Will be &#x60;-1&#x60; if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  [new-channel-announce]: /help/configure-automated-notices#new-channel-announcements  **Changes**: In Zulip 9.0 (feature level 241), renamed &#x60;notifications_stream_id&#x60; to &#x60;new_stream_announcements_stream_id&#x60;.  | [optional] 
**OrgType** | Pointer to **int32** | The [organization type](zulip.com/help/organization-type for the realm.  - 0 &#x3D; Unspecified - 10 &#x3D; Business - 20 &#x3D; Open-source project - 30 &#x3D; Education (non-profit) - 35 &#x3D; Education (for-profit) - 40 &#x3D; Research - 50 &#x3D; Event or conference - 60 &#x3D; Non-profit (registered) - 70 &#x3D; Government - 80 &#x3D; Political group - 90 &#x3D; Community - 100 &#x3D; Personal - 1000 &#x3D; Other  **Changes**: New in Zulip 6.0 (feature level 128).  | [optional] 
**PlanType** | Pointer to **int32** | The plan type of the organization.  - 1 &#x3D; Self-hosted organization (SELF_HOSTED) - 2 &#x3D; Zulip Cloud free plan (LIMITED) - 3 &#x3D; Zulip Cloud Standard plan (STANDARD) - 4 &#x3D; Zulip Cloud Standard plan, sponsored for free (STANDARD_FREE)  | [optional] 
**PresenceDisabled** | Pointer to **bool** | Whether online presence of other users is shown in this organization.  | [optional] 
**PushNotificationsEnabled** | Pointer to **bool** | Whether push notifications are enabled for this organization. Typically &#x60;true&#x60; for Zulip Cloud and self-hosted realms that have a valid registration for the [Mobile push notifications service](https://zulip.readthedocs.io/en/latest/production/mobile-push-notifications.html), and &#x60;false&#x60; for self-hosted servers that do not.  **Changes**: New in Zulip 8.0 (feature level 231). Previously, this value was never updated via events.  | [optional] 
**PushNotificationsEnabledEndTimestamp** | Pointer to **NullableInt32** | If the server expects the realm&#39;s push notifications access to end at a definite time in the future, the time at which this is expected to happen. Mobile clients should use this field to display warnings to users when the indicated timestamp is near.  **Changes**: New in Zulip 8.0 (feature level 231).  | [optional] 
**RequireE2eePushNotifications** | Pointer to **bool** | Whether this realm is configured to disallow sending mobile push notifications with message content through the legacy mobile push notifications APIs. The new API uses end-to-end encryption to protect message content and metadata from being accessible to the push bouncer service, APNs, and FCM. Clients that support the new E2EE API will use it automatically regardless of this setting.  If &#x60;true&#x60;, mobile push notifications sent to clients that lack support for E2EE push notifications will always have \&quot;New message\&quot; as their content. Note that these legacy mobile notifications will still contain metadata, which may include the message&#39;s ID, the sender&#39;s name, email address, and avatar.  In a future release, once the official mobile apps have implemented fully validated their E2EE protocol support, this setting will become strict, and disable the legacy protocol entirely.  **Changes**: New in Zulip 11.0 (feature level 409). Previously, this behavior was available only via the &#x60;PUSH_NOTIFICATION_REDACT_CONTENT&#x60; global server setting.  | [optional] 
**RequireUniqueNames** | Pointer to **bool** | Indicates whether the organization is configured to require users to have unique full names. If true, the server will reject attempts to create a new user, or change the name of an existing user, where doing so would lead to two users whose names are identical modulo case and unicode normalization.  **Changes**: New in Zulip 9.0 (feature level 246). Previously, the Zulip server could not be configured to enforce unique names.  | [optional] 
**SendWelcomeEmails** | Pointer to **bool** | Whether or not this organization is configured to send the standard Zulip [welcome emails](zulip.com/help/disable-welcome-emails to new users joining the organization.  | [optional] 
**SignupAnnouncementsStreamId** | Pointer to **int32** | The ID of the channel to which automated messages announcing that [new users have joined the organization][new-user-announce] are sent.  Will be &#x60;-1&#x60; if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  [new-user-announce]: /help/configure-automated-notices#new-user-announcements  **Changes**: In Zulip 9.0 (feature level 241), renamed &#x60;signup_notifications_stream_id&#x60; to &#x60;signup_announcements_stream_id&#x60;.  | [optional] 
**UploadQuotaMib** | Pointer to **NullableInt32** | The new upload quota for the Zulip organization.  If &#x60;null&#x60;, there is no limit.  **Changes**: New in Zulip 10.0 (feature level 306). Previously, this was present changed via an &#x60;upload_quota&#x60; field in &#x60;extra_data&#x60; property of [realm/update](#realm-update) event format for &#x60;plan_type&#x60; events.  | [optional] 
**VideoChatProvider** | Pointer to **int32** | The configured [video call provider](zulip.com/help/configure-call-provider for the organization.  - 0 &#x3D; None - 1 &#x3D; Jitsi Meet - 3 &#x3D; Zoom (User OAuth integration) - 4 &#x3D; BigBlueButton - 5 &#x3D; Zoom (Server to Server OAuth integration)  Note that only one of the [Zoom integrations][zoom-video-calls] can be configured on a Zulip server.  **Changes**: In Zulip 10.0 (feature level 353), added the Zoom Server to Server OAuth option.  In Zulip 3.0 (feature level 1), added the None option to disable video call UI.  [zoom-video-calls]: https://zulip.readthedocs.io/en/latest/production/video-calls.html#zoom  | [optional] 
**WaitingPeriodThreshold** | Pointer to **int32** | Members whose accounts have been created at least this many days ago will be treated as [full members][calc-full-member] for the purpose of settings that restrict access to new members.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member  | [optional] 
**WantAdvertiseInCommunitiesDirectory** | Pointer to **bool** | Whether the organization has given permission to be advertised in the Zulip [communities directory](zulip.com/help/communities-directory.  **Changes**: New in Zulip 6.0 (feature level 129).  | [optional] 
**WelcomeMessageCustomText** | Pointer to **string** | This organization&#39;s configured custom message for Welcome Bot to send to new user accounts, in Zulip Markdown format.  Maximum length is 8000 characters.  **Changes**: New in Zulip 11.0 (feature level 416).  | [optional] 
**ZulipUpdateAnnouncementsStreamId** | Pointer to **int32** | The ID of the channel to which automated messages announcing new features or other end-user updates about the Zulip software are sent.  Will be &#x60;-1&#x60; if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  **Changes**: New in Zulip 9.0 (feature level 242).  | [optional] 

## Methods

### NewEventEnvelopeOneOf67Data

`func NewEventEnvelopeOneOf67Data() *EventEnvelopeOneOf67Data`

NewEventEnvelopeOneOf67Data instantiates a new EventEnvelopeOneOf67Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf67DataWithDefaults

`func NewEventEnvelopeOneOf67DataWithDefaults() *EventEnvelopeOneOf67Data`

NewEventEnvelopeOneOf67DataWithDefaults instantiates a new EventEnvelopeOneOf67Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMessageEditing

`func (o *EventEnvelopeOneOf67Data) GetAllowMessageEditing() bool`

GetAllowMessageEditing returns the AllowMessageEditing field if non-nil, zero value otherwise.

### GetAllowMessageEditingOk

`func (o *EventEnvelopeOneOf67Data) GetAllowMessageEditingOk() (*bool, bool)`

GetAllowMessageEditingOk returns a tuple with the AllowMessageEditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessageEditing

`func (o *EventEnvelopeOneOf67Data) SetAllowMessageEditing(v bool)`

SetAllowMessageEditing sets AllowMessageEditing field to given value.

### HasAllowMessageEditing

`func (o *EventEnvelopeOneOf67Data) HasAllowMessageEditing() bool`

HasAllowMessageEditing returns a boolean if a field has been set.

### GetAuthenticationMethods

`func (o *EventEnvelopeOneOf67Data) GetAuthenticationMethods() map[string]RealmAuthenticationMethod`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *EventEnvelopeOneOf67Data) GetAuthenticationMethodsOk() (*map[string]RealmAuthenticationMethod, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *EventEnvelopeOneOf67Data) SetAuthenticationMethods(v map[string]RealmAuthenticationMethod)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *EventEnvelopeOneOf67Data) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### GetCanAccessAllUsersGroup

`func (o *EventEnvelopeOneOf67Data) GetCanAccessAllUsersGroup() GroupSettingValue`

GetCanAccessAllUsersGroup returns the CanAccessAllUsersGroup field if non-nil, zero value otherwise.

### GetCanAccessAllUsersGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanAccessAllUsersGroupOk() (*GroupSettingValue, bool)`

GetCanAccessAllUsersGroupOk returns a tuple with the CanAccessAllUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessAllUsersGroup

`func (o *EventEnvelopeOneOf67Data) SetCanAccessAllUsersGroup(v GroupSettingValue)`

SetCanAccessAllUsersGroup sets CanAccessAllUsersGroup field to given value.

### HasCanAccessAllUsersGroup

`func (o *EventEnvelopeOneOf67Data) HasCanAccessAllUsersGroup() bool`

HasCanAccessAllUsersGroup returns a boolean if a field has been set.

### GetCanCreateGroups

`func (o *EventEnvelopeOneOf67Data) GetCanCreateGroups() GroupSettingValue`

GetCanCreateGroups returns the CanCreateGroups field if non-nil, zero value otherwise.

### GetCanCreateGroupsOk

`func (o *EventEnvelopeOneOf67Data) GetCanCreateGroupsOk() (*GroupSettingValue, bool)`

GetCanCreateGroupsOk returns a tuple with the CanCreateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateGroups

`func (o *EventEnvelopeOneOf67Data) SetCanCreateGroups(v GroupSettingValue)`

SetCanCreateGroups sets CanCreateGroups field to given value.

### HasCanCreateGroups

`func (o *EventEnvelopeOneOf67Data) HasCanCreateGroups() bool`

HasCanCreateGroups returns a boolean if a field has been set.

### GetCanCreateBotsGroup

`func (o *EventEnvelopeOneOf67Data) GetCanCreateBotsGroup() GroupSettingValue`

GetCanCreateBotsGroup returns the CanCreateBotsGroup field if non-nil, zero value otherwise.

### GetCanCreateBotsGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanCreateBotsGroupOk() (*GroupSettingValue, bool)`

GetCanCreateBotsGroupOk returns a tuple with the CanCreateBotsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateBotsGroup

`func (o *EventEnvelopeOneOf67Data) SetCanCreateBotsGroup(v GroupSettingValue)`

SetCanCreateBotsGroup sets CanCreateBotsGroup field to given value.

### HasCanCreateBotsGroup

`func (o *EventEnvelopeOneOf67Data) HasCanCreateBotsGroup() bool`

HasCanCreateBotsGroup returns a boolean if a field has been set.

### GetCanCreateWriteOnlyBotsGroup

`func (o *EventEnvelopeOneOf67Data) GetCanCreateWriteOnlyBotsGroup() GroupSettingValue`

GetCanCreateWriteOnlyBotsGroup returns the CanCreateWriteOnlyBotsGroup field if non-nil, zero value otherwise.

### GetCanCreateWriteOnlyBotsGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanCreateWriteOnlyBotsGroupOk() (*GroupSettingValue, bool)`

GetCanCreateWriteOnlyBotsGroupOk returns a tuple with the CanCreateWriteOnlyBotsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateWriteOnlyBotsGroup

`func (o *EventEnvelopeOneOf67Data) SetCanCreateWriteOnlyBotsGroup(v GroupSettingValue)`

SetCanCreateWriteOnlyBotsGroup sets CanCreateWriteOnlyBotsGroup field to given value.

### HasCanCreateWriteOnlyBotsGroup

`func (o *EventEnvelopeOneOf67Data) HasCanCreateWriteOnlyBotsGroup() bool`

HasCanCreateWriteOnlyBotsGroup returns a boolean if a field has been set.

### GetCanCreatePublicChannelGroup

`func (o *EventEnvelopeOneOf67Data) GetCanCreatePublicChannelGroup() GroupSettingValue`

GetCanCreatePublicChannelGroup returns the CanCreatePublicChannelGroup field if non-nil, zero value otherwise.

### GetCanCreatePublicChannelGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanCreatePublicChannelGroupOk() (*GroupSettingValue, bool)`

GetCanCreatePublicChannelGroupOk returns a tuple with the CanCreatePublicChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreatePublicChannelGroup

`func (o *EventEnvelopeOneOf67Data) SetCanCreatePublicChannelGroup(v GroupSettingValue)`

SetCanCreatePublicChannelGroup sets CanCreatePublicChannelGroup field to given value.

### HasCanCreatePublicChannelGroup

`func (o *EventEnvelopeOneOf67Data) HasCanCreatePublicChannelGroup() bool`

HasCanCreatePublicChannelGroup returns a boolean if a field has been set.

### GetCanCreatePrivateChannelGroup

`func (o *EventEnvelopeOneOf67Data) GetCanCreatePrivateChannelGroup() GroupSettingValue`

GetCanCreatePrivateChannelGroup returns the CanCreatePrivateChannelGroup field if non-nil, zero value otherwise.

### GetCanCreatePrivateChannelGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanCreatePrivateChannelGroupOk() (*GroupSettingValue, bool)`

GetCanCreatePrivateChannelGroupOk returns a tuple with the CanCreatePrivateChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreatePrivateChannelGroup

`func (o *EventEnvelopeOneOf67Data) SetCanCreatePrivateChannelGroup(v GroupSettingValue)`

SetCanCreatePrivateChannelGroup sets CanCreatePrivateChannelGroup field to given value.

### HasCanCreatePrivateChannelGroup

`func (o *EventEnvelopeOneOf67Data) HasCanCreatePrivateChannelGroup() bool`

HasCanCreatePrivateChannelGroup returns a boolean if a field has been set.

### GetCanCreateWebPublicChannelGroup

`func (o *EventEnvelopeOneOf67Data) GetCanCreateWebPublicChannelGroup() GroupSettingValue`

GetCanCreateWebPublicChannelGroup returns the CanCreateWebPublicChannelGroup field if non-nil, zero value otherwise.

### GetCanCreateWebPublicChannelGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanCreateWebPublicChannelGroupOk() (*GroupSettingValue, bool)`

GetCanCreateWebPublicChannelGroupOk returns a tuple with the CanCreateWebPublicChannelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateWebPublicChannelGroup

`func (o *EventEnvelopeOneOf67Data) SetCanCreateWebPublicChannelGroup(v GroupSettingValue)`

SetCanCreateWebPublicChannelGroup sets CanCreateWebPublicChannelGroup field to given value.

### HasCanCreateWebPublicChannelGroup

`func (o *EventEnvelopeOneOf67Data) HasCanCreateWebPublicChannelGroup() bool`

HasCanCreateWebPublicChannelGroup returns a boolean if a field has been set.

### GetCanAddCustomEmojiGroup

`func (o *EventEnvelopeOneOf67Data) GetCanAddCustomEmojiGroup() GroupSettingValue`

GetCanAddCustomEmojiGroup returns the CanAddCustomEmojiGroup field if non-nil, zero value otherwise.

### GetCanAddCustomEmojiGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanAddCustomEmojiGroupOk() (*GroupSettingValue, bool)`

GetCanAddCustomEmojiGroupOk returns a tuple with the CanAddCustomEmojiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddCustomEmojiGroup

`func (o *EventEnvelopeOneOf67Data) SetCanAddCustomEmojiGroup(v GroupSettingValue)`

SetCanAddCustomEmojiGroup sets CanAddCustomEmojiGroup field to given value.

### HasCanAddCustomEmojiGroup

`func (o *EventEnvelopeOneOf67Data) HasCanAddCustomEmojiGroup() bool`

HasCanAddCustomEmojiGroup returns a boolean if a field has been set.

### GetCanAddSubscribersGroup

`func (o *EventEnvelopeOneOf67Data) GetCanAddSubscribersGroup() GroupSettingValue`

GetCanAddSubscribersGroup returns the CanAddSubscribersGroup field if non-nil, zero value otherwise.

### GetCanAddSubscribersGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanAddSubscribersGroupOk() (*GroupSettingValue, bool)`

GetCanAddSubscribersGroupOk returns a tuple with the CanAddSubscribersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddSubscribersGroup

`func (o *EventEnvelopeOneOf67Data) SetCanAddSubscribersGroup(v GroupSettingValue)`

SetCanAddSubscribersGroup sets CanAddSubscribersGroup field to given value.

### HasCanAddSubscribersGroup

`func (o *EventEnvelopeOneOf67Data) HasCanAddSubscribersGroup() bool`

HasCanAddSubscribersGroup returns a boolean if a field has been set.

### GetCanDeleteAnyMessageGroup

`func (o *EventEnvelopeOneOf67Data) GetCanDeleteAnyMessageGroup() GroupSettingValue`

GetCanDeleteAnyMessageGroup returns the CanDeleteAnyMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteAnyMessageGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanDeleteAnyMessageGroupOk() (*GroupSettingValue, bool)`

GetCanDeleteAnyMessageGroupOk returns a tuple with the CanDeleteAnyMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteAnyMessageGroup

`func (o *EventEnvelopeOneOf67Data) SetCanDeleteAnyMessageGroup(v GroupSettingValue)`

SetCanDeleteAnyMessageGroup sets CanDeleteAnyMessageGroup field to given value.

### HasCanDeleteAnyMessageGroup

`func (o *EventEnvelopeOneOf67Data) HasCanDeleteAnyMessageGroup() bool`

HasCanDeleteAnyMessageGroup returns a boolean if a field has been set.

### GetCanDeleteOwnMessageGroup

`func (o *EventEnvelopeOneOf67Data) GetCanDeleteOwnMessageGroup() GroupSettingValue`

GetCanDeleteOwnMessageGroup returns the CanDeleteOwnMessageGroup field if non-nil, zero value otherwise.

### GetCanDeleteOwnMessageGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanDeleteOwnMessageGroupOk() (*GroupSettingValue, bool)`

GetCanDeleteOwnMessageGroupOk returns a tuple with the CanDeleteOwnMessageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteOwnMessageGroup

`func (o *EventEnvelopeOneOf67Data) SetCanDeleteOwnMessageGroup(v GroupSettingValue)`

SetCanDeleteOwnMessageGroup sets CanDeleteOwnMessageGroup field to given value.

### HasCanDeleteOwnMessageGroup

`func (o *EventEnvelopeOneOf67Data) HasCanDeleteOwnMessageGroup() bool`

HasCanDeleteOwnMessageGroup returns a boolean if a field has been set.

### GetCanSetDeleteMessagePolicyGroup

`func (o *EventEnvelopeOneOf67Data) GetCanSetDeleteMessagePolicyGroup() GroupSettingValue`

GetCanSetDeleteMessagePolicyGroup returns the CanSetDeleteMessagePolicyGroup field if non-nil, zero value otherwise.

### GetCanSetDeleteMessagePolicyGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanSetDeleteMessagePolicyGroupOk() (*GroupSettingValue, bool)`

GetCanSetDeleteMessagePolicyGroupOk returns a tuple with the CanSetDeleteMessagePolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSetDeleteMessagePolicyGroup

`func (o *EventEnvelopeOneOf67Data) SetCanSetDeleteMessagePolicyGroup(v GroupSettingValue)`

SetCanSetDeleteMessagePolicyGroup sets CanSetDeleteMessagePolicyGroup field to given value.

### HasCanSetDeleteMessagePolicyGroup

`func (o *EventEnvelopeOneOf67Data) HasCanSetDeleteMessagePolicyGroup() bool`

HasCanSetDeleteMessagePolicyGroup returns a boolean if a field has been set.

### GetCanSetTopicsPolicyGroup

`func (o *EventEnvelopeOneOf67Data) GetCanSetTopicsPolicyGroup() GroupSettingValue`

GetCanSetTopicsPolicyGroup returns the CanSetTopicsPolicyGroup field if non-nil, zero value otherwise.

### GetCanSetTopicsPolicyGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanSetTopicsPolicyGroupOk() (*GroupSettingValue, bool)`

GetCanSetTopicsPolicyGroupOk returns a tuple with the CanSetTopicsPolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSetTopicsPolicyGroup

`func (o *EventEnvelopeOneOf67Data) SetCanSetTopicsPolicyGroup(v GroupSettingValue)`

SetCanSetTopicsPolicyGroup sets CanSetTopicsPolicyGroup field to given value.

### HasCanSetTopicsPolicyGroup

`func (o *EventEnvelopeOneOf67Data) HasCanSetTopicsPolicyGroup() bool`

HasCanSetTopicsPolicyGroup returns a boolean if a field has been set.

### GetCanInviteUsersGroup

`func (o *EventEnvelopeOneOf67Data) GetCanInviteUsersGroup() GroupSettingValue`

GetCanInviteUsersGroup returns the CanInviteUsersGroup field if non-nil, zero value otherwise.

### GetCanInviteUsersGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanInviteUsersGroupOk() (*GroupSettingValue, bool)`

GetCanInviteUsersGroupOk returns a tuple with the CanInviteUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInviteUsersGroup

`func (o *EventEnvelopeOneOf67Data) SetCanInviteUsersGroup(v GroupSettingValue)`

SetCanInviteUsersGroup sets CanInviteUsersGroup field to given value.

### HasCanInviteUsersGroup

`func (o *EventEnvelopeOneOf67Data) HasCanInviteUsersGroup() bool`

HasCanInviteUsersGroup returns a boolean if a field has been set.

### GetCanMentionManyUsersGroup

`func (o *EventEnvelopeOneOf67Data) GetCanMentionManyUsersGroup() GroupSettingValue`

GetCanMentionManyUsersGroup returns the CanMentionManyUsersGroup field if non-nil, zero value otherwise.

### GetCanMentionManyUsersGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanMentionManyUsersGroupOk() (*GroupSettingValue, bool)`

GetCanMentionManyUsersGroupOk returns a tuple with the CanMentionManyUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMentionManyUsersGroup

`func (o *EventEnvelopeOneOf67Data) SetCanMentionManyUsersGroup(v GroupSettingValue)`

SetCanMentionManyUsersGroup sets CanMentionManyUsersGroup field to given value.

### HasCanMentionManyUsersGroup

`func (o *EventEnvelopeOneOf67Data) HasCanMentionManyUsersGroup() bool`

HasCanMentionManyUsersGroup returns a boolean if a field has been set.

### GetCanMoveMessagesBetweenChannelsGroup

`func (o *EventEnvelopeOneOf67Data) GetCanMoveMessagesBetweenChannelsGroup() GroupSettingValue`

GetCanMoveMessagesBetweenChannelsGroup returns the CanMoveMessagesBetweenChannelsGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesBetweenChannelsGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanMoveMessagesBetweenChannelsGroupOk() (*GroupSettingValue, bool)`

GetCanMoveMessagesBetweenChannelsGroupOk returns a tuple with the CanMoveMessagesBetweenChannelsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesBetweenChannelsGroup

`func (o *EventEnvelopeOneOf67Data) SetCanMoveMessagesBetweenChannelsGroup(v GroupSettingValue)`

SetCanMoveMessagesBetweenChannelsGroup sets CanMoveMessagesBetweenChannelsGroup field to given value.

### HasCanMoveMessagesBetweenChannelsGroup

`func (o *EventEnvelopeOneOf67Data) HasCanMoveMessagesBetweenChannelsGroup() bool`

HasCanMoveMessagesBetweenChannelsGroup returns a boolean if a field has been set.

### GetCanMoveMessagesBetweenTopicsGroup

`func (o *EventEnvelopeOneOf67Data) GetCanMoveMessagesBetweenTopicsGroup() GroupSettingValue`

GetCanMoveMessagesBetweenTopicsGroup returns the CanMoveMessagesBetweenTopicsGroup field if non-nil, zero value otherwise.

### GetCanMoveMessagesBetweenTopicsGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanMoveMessagesBetweenTopicsGroupOk() (*GroupSettingValue, bool)`

GetCanMoveMessagesBetweenTopicsGroupOk returns a tuple with the CanMoveMessagesBetweenTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMoveMessagesBetweenTopicsGroup

`func (o *EventEnvelopeOneOf67Data) SetCanMoveMessagesBetweenTopicsGroup(v GroupSettingValue)`

SetCanMoveMessagesBetweenTopicsGroup sets CanMoveMessagesBetweenTopicsGroup field to given value.

### HasCanMoveMessagesBetweenTopicsGroup

`func (o *EventEnvelopeOneOf67Data) HasCanMoveMessagesBetweenTopicsGroup() bool`

HasCanMoveMessagesBetweenTopicsGroup returns a boolean if a field has been set.

### GetCanResolveTopicsGroup

`func (o *EventEnvelopeOneOf67Data) GetCanResolveTopicsGroup() GroupSettingValue`

GetCanResolveTopicsGroup returns the CanResolveTopicsGroup field if non-nil, zero value otherwise.

### GetCanResolveTopicsGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanResolveTopicsGroupOk() (*GroupSettingValue, bool)`

GetCanResolveTopicsGroupOk returns a tuple with the CanResolveTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResolveTopicsGroup

`func (o *EventEnvelopeOneOf67Data) SetCanResolveTopicsGroup(v GroupSettingValue)`

SetCanResolveTopicsGroup sets CanResolveTopicsGroup field to given value.

### HasCanResolveTopicsGroup

`func (o *EventEnvelopeOneOf67Data) HasCanResolveTopicsGroup() bool`

HasCanResolveTopicsGroup returns a boolean if a field has been set.

### GetCanManageAllGroups

`func (o *EventEnvelopeOneOf67Data) GetCanManageAllGroups() GroupSettingValue`

GetCanManageAllGroups returns the CanManageAllGroups field if non-nil, zero value otherwise.

### GetCanManageAllGroupsOk

`func (o *EventEnvelopeOneOf67Data) GetCanManageAllGroupsOk() (*GroupSettingValue, bool)`

GetCanManageAllGroupsOk returns a tuple with the CanManageAllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageAllGroups

`func (o *EventEnvelopeOneOf67Data) SetCanManageAllGroups(v GroupSettingValue)`

SetCanManageAllGroups sets CanManageAllGroups field to given value.

### HasCanManageAllGroups

`func (o *EventEnvelopeOneOf67Data) HasCanManageAllGroups() bool`

HasCanManageAllGroups returns a boolean if a field has been set.

### GetCanManageBillingGroup

`func (o *EventEnvelopeOneOf67Data) GetCanManageBillingGroup() GroupSettingValue`

GetCanManageBillingGroup returns the CanManageBillingGroup field if non-nil, zero value otherwise.

### GetCanManageBillingGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanManageBillingGroupOk() (*GroupSettingValue, bool)`

GetCanManageBillingGroupOk returns a tuple with the CanManageBillingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageBillingGroup

`func (o *EventEnvelopeOneOf67Data) SetCanManageBillingGroup(v GroupSettingValue)`

SetCanManageBillingGroup sets CanManageBillingGroup field to given value.

### HasCanManageBillingGroup

`func (o *EventEnvelopeOneOf67Data) HasCanManageBillingGroup() bool`

HasCanManageBillingGroup returns a boolean if a field has been set.

### GetCanSummarizeTopicsGroup

`func (o *EventEnvelopeOneOf67Data) GetCanSummarizeTopicsGroup() GroupSettingValue`

GetCanSummarizeTopicsGroup returns the CanSummarizeTopicsGroup field if non-nil, zero value otherwise.

### GetCanSummarizeTopicsGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCanSummarizeTopicsGroupOk() (*GroupSettingValue, bool)`

GetCanSummarizeTopicsGroupOk returns a tuple with the CanSummarizeTopicsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSummarizeTopicsGroup

`func (o *EventEnvelopeOneOf67Data) SetCanSummarizeTopicsGroup(v GroupSettingValue)`

SetCanSummarizeTopicsGroup sets CanSummarizeTopicsGroup field to given value.

### HasCanSummarizeTopicsGroup

`func (o *EventEnvelopeOneOf67Data) HasCanSummarizeTopicsGroup() bool`

HasCanSummarizeTopicsGroup returns a boolean if a field has been set.

### GetCreateMultiuseInviteGroup

`func (o *EventEnvelopeOneOf67Data) GetCreateMultiuseInviteGroup() GroupSettingValue`

GetCreateMultiuseInviteGroup returns the CreateMultiuseInviteGroup field if non-nil, zero value otherwise.

### GetCreateMultiuseInviteGroupOk

`func (o *EventEnvelopeOneOf67Data) GetCreateMultiuseInviteGroupOk() (*GroupSettingValue, bool)`

GetCreateMultiuseInviteGroupOk returns a tuple with the CreateMultiuseInviteGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMultiuseInviteGroup

`func (o *EventEnvelopeOneOf67Data) SetCreateMultiuseInviteGroup(v GroupSettingValue)`

SetCreateMultiuseInviteGroup sets CreateMultiuseInviteGroup field to given value.

### HasCreateMultiuseInviteGroup

`func (o *EventEnvelopeOneOf67Data) HasCreateMultiuseInviteGroup() bool`

HasCreateMultiuseInviteGroup returns a boolean if a field has been set.

### GetDefaultCodeBlockLanguage

`func (o *EventEnvelopeOneOf67Data) GetDefaultCodeBlockLanguage() string`

GetDefaultCodeBlockLanguage returns the DefaultCodeBlockLanguage field if non-nil, zero value otherwise.

### GetDefaultCodeBlockLanguageOk

`func (o *EventEnvelopeOneOf67Data) GetDefaultCodeBlockLanguageOk() (*string, bool)`

GetDefaultCodeBlockLanguageOk returns a tuple with the DefaultCodeBlockLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCodeBlockLanguage

`func (o *EventEnvelopeOneOf67Data) SetDefaultCodeBlockLanguage(v string)`

SetDefaultCodeBlockLanguage sets DefaultCodeBlockLanguage field to given value.

### HasDefaultCodeBlockLanguage

`func (o *EventEnvelopeOneOf67Data) HasDefaultCodeBlockLanguage() bool`

HasDefaultCodeBlockLanguage returns a boolean if a field has been set.

### GetDefaultLanguage

`func (o *EventEnvelopeOneOf67Data) GetDefaultLanguage() string`

GetDefaultLanguage returns the DefaultLanguage field if non-nil, zero value otherwise.

### GetDefaultLanguageOk

`func (o *EventEnvelopeOneOf67Data) GetDefaultLanguageOk() (*string, bool)`

GetDefaultLanguageOk returns a tuple with the DefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguage

`func (o *EventEnvelopeOneOf67Data) SetDefaultLanguage(v string)`

SetDefaultLanguage sets DefaultLanguage field to given value.

### HasDefaultLanguage

`func (o *EventEnvelopeOneOf67Data) HasDefaultLanguage() bool`

HasDefaultLanguage returns a boolean if a field has been set.

### GetDescription

`func (o *EventEnvelopeOneOf67Data) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventEnvelopeOneOf67Data) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventEnvelopeOneOf67Data) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventEnvelopeOneOf67Data) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigestEmailsEnabled

`func (o *EventEnvelopeOneOf67Data) GetDigestEmailsEnabled() bool`

GetDigestEmailsEnabled returns the DigestEmailsEnabled field if non-nil, zero value otherwise.

### GetDigestEmailsEnabledOk

`func (o *EventEnvelopeOneOf67Data) GetDigestEmailsEnabledOk() (*bool, bool)`

GetDigestEmailsEnabledOk returns a tuple with the DigestEmailsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestEmailsEnabled

`func (o *EventEnvelopeOneOf67Data) SetDigestEmailsEnabled(v bool)`

SetDigestEmailsEnabled sets DigestEmailsEnabled field to given value.

### HasDigestEmailsEnabled

`func (o *EventEnvelopeOneOf67Data) HasDigestEmailsEnabled() bool`

HasDigestEmailsEnabled returns a boolean if a field has been set.

### GetDigestWeekday

`func (o *EventEnvelopeOneOf67Data) GetDigestWeekday() int32`

GetDigestWeekday returns the DigestWeekday field if non-nil, zero value otherwise.

### GetDigestWeekdayOk

`func (o *EventEnvelopeOneOf67Data) GetDigestWeekdayOk() (*int32, bool)`

GetDigestWeekdayOk returns a tuple with the DigestWeekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestWeekday

`func (o *EventEnvelopeOneOf67Data) SetDigestWeekday(v int32)`

SetDigestWeekday sets DigestWeekday field to given value.

### HasDigestWeekday

`func (o *EventEnvelopeOneOf67Data) HasDigestWeekday() bool`

HasDigestWeekday returns a boolean if a field has been set.

### GetDirectMessageInitiatorGroup

`func (o *EventEnvelopeOneOf67Data) GetDirectMessageInitiatorGroup() GroupSettingValue`

GetDirectMessageInitiatorGroup returns the DirectMessageInitiatorGroup field if non-nil, zero value otherwise.

### GetDirectMessageInitiatorGroupOk

`func (o *EventEnvelopeOneOf67Data) GetDirectMessageInitiatorGroupOk() (*GroupSettingValue, bool)`

GetDirectMessageInitiatorGroupOk returns a tuple with the DirectMessageInitiatorGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMessageInitiatorGroup

`func (o *EventEnvelopeOneOf67Data) SetDirectMessageInitiatorGroup(v GroupSettingValue)`

SetDirectMessageInitiatorGroup sets DirectMessageInitiatorGroup field to given value.

### HasDirectMessageInitiatorGroup

`func (o *EventEnvelopeOneOf67Data) HasDirectMessageInitiatorGroup() bool`

HasDirectMessageInitiatorGroup returns a boolean if a field has been set.

### GetDirectMessagePermissionGroup

`func (o *EventEnvelopeOneOf67Data) GetDirectMessagePermissionGroup() GroupSettingValue`

GetDirectMessagePermissionGroup returns the DirectMessagePermissionGroup field if non-nil, zero value otherwise.

### GetDirectMessagePermissionGroupOk

`func (o *EventEnvelopeOneOf67Data) GetDirectMessagePermissionGroupOk() (*GroupSettingValue, bool)`

GetDirectMessagePermissionGroupOk returns a tuple with the DirectMessagePermissionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMessagePermissionGroup

`func (o *EventEnvelopeOneOf67Data) SetDirectMessagePermissionGroup(v GroupSettingValue)`

SetDirectMessagePermissionGroup sets DirectMessagePermissionGroup field to given value.

### HasDirectMessagePermissionGroup

`func (o *EventEnvelopeOneOf67Data) HasDirectMessagePermissionGroup() bool`

HasDirectMessagePermissionGroup returns a boolean if a field has been set.

### GetDisallowDisposableEmailAddresses

`func (o *EventEnvelopeOneOf67Data) GetDisallowDisposableEmailAddresses() bool`

GetDisallowDisposableEmailAddresses returns the DisallowDisposableEmailAddresses field if non-nil, zero value otherwise.

### GetDisallowDisposableEmailAddressesOk

`func (o *EventEnvelopeOneOf67Data) GetDisallowDisposableEmailAddressesOk() (*bool, bool)`

GetDisallowDisposableEmailAddressesOk returns a tuple with the DisallowDisposableEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowDisposableEmailAddresses

`func (o *EventEnvelopeOneOf67Data) SetDisallowDisposableEmailAddresses(v bool)`

SetDisallowDisposableEmailAddresses sets DisallowDisposableEmailAddresses field to given value.

### HasDisallowDisposableEmailAddresses

`func (o *EventEnvelopeOneOf67Data) HasDisallowDisposableEmailAddresses() bool`

HasDisallowDisposableEmailAddresses returns a boolean if a field has been set.

### GetEmailChangesDisabled

`func (o *EventEnvelopeOneOf67Data) GetEmailChangesDisabled() bool`

GetEmailChangesDisabled returns the EmailChangesDisabled field if non-nil, zero value otherwise.

### GetEmailChangesDisabledOk

`func (o *EventEnvelopeOneOf67Data) GetEmailChangesDisabledOk() (*bool, bool)`

GetEmailChangesDisabledOk returns a tuple with the EmailChangesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailChangesDisabled

`func (o *EventEnvelopeOneOf67Data) SetEmailChangesDisabled(v bool)`

SetEmailChangesDisabled sets EmailChangesDisabled field to given value.

### HasEmailChangesDisabled

`func (o *EventEnvelopeOneOf67Data) HasEmailChangesDisabled() bool`

HasEmailChangesDisabled returns a boolean if a field has been set.

### GetEnableReadReceipts

`func (o *EventEnvelopeOneOf67Data) GetEnableReadReceipts() bool`

GetEnableReadReceipts returns the EnableReadReceipts field if non-nil, zero value otherwise.

### GetEnableReadReceiptsOk

`func (o *EventEnvelopeOneOf67Data) GetEnableReadReceiptsOk() (*bool, bool)`

GetEnableReadReceiptsOk returns a tuple with the EnableReadReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReadReceipts

`func (o *EventEnvelopeOneOf67Data) SetEnableReadReceipts(v bool)`

SetEnableReadReceipts sets EnableReadReceipts field to given value.

### HasEnableReadReceipts

`func (o *EventEnvelopeOneOf67Data) HasEnableReadReceipts() bool`

HasEnableReadReceipts returns a boolean if a field has been set.

### GetEmailsRestrictedToDomains

`func (o *EventEnvelopeOneOf67Data) GetEmailsRestrictedToDomains() bool`

GetEmailsRestrictedToDomains returns the EmailsRestrictedToDomains field if non-nil, zero value otherwise.

### GetEmailsRestrictedToDomainsOk

`func (o *EventEnvelopeOneOf67Data) GetEmailsRestrictedToDomainsOk() (*bool, bool)`

GetEmailsRestrictedToDomainsOk returns a tuple with the EmailsRestrictedToDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsRestrictedToDomains

`func (o *EventEnvelopeOneOf67Data) SetEmailsRestrictedToDomains(v bool)`

SetEmailsRestrictedToDomains sets EmailsRestrictedToDomains field to given value.

### HasEmailsRestrictedToDomains

`func (o *EventEnvelopeOneOf67Data) HasEmailsRestrictedToDomains() bool`

HasEmailsRestrictedToDomains returns a boolean if a field has been set.

### GetEnableGuestUserDmWarning

`func (o *EventEnvelopeOneOf67Data) GetEnableGuestUserDmWarning() bool`

GetEnableGuestUserDmWarning returns the EnableGuestUserDmWarning field if non-nil, zero value otherwise.

### GetEnableGuestUserDmWarningOk

`func (o *EventEnvelopeOneOf67Data) GetEnableGuestUserDmWarningOk() (*bool, bool)`

GetEnableGuestUserDmWarningOk returns a tuple with the EnableGuestUserDmWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGuestUserDmWarning

`func (o *EventEnvelopeOneOf67Data) SetEnableGuestUserDmWarning(v bool)`

SetEnableGuestUserDmWarning sets EnableGuestUserDmWarning field to given value.

### HasEnableGuestUserDmWarning

`func (o *EventEnvelopeOneOf67Data) HasEnableGuestUserDmWarning() bool`

HasEnableGuestUserDmWarning returns a boolean if a field has been set.

### GetEnableGuestUserIndicator

`func (o *EventEnvelopeOneOf67Data) GetEnableGuestUserIndicator() bool`

GetEnableGuestUserIndicator returns the EnableGuestUserIndicator field if non-nil, zero value otherwise.

### GetEnableGuestUserIndicatorOk

`func (o *EventEnvelopeOneOf67Data) GetEnableGuestUserIndicatorOk() (*bool, bool)`

GetEnableGuestUserIndicatorOk returns a tuple with the EnableGuestUserIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGuestUserIndicator

`func (o *EventEnvelopeOneOf67Data) SetEnableGuestUserIndicator(v bool)`

SetEnableGuestUserIndicator sets EnableGuestUserIndicator field to given value.

### HasEnableGuestUserIndicator

`func (o *EventEnvelopeOneOf67Data) HasEnableGuestUserIndicator() bool`

HasEnableGuestUserIndicator returns a boolean if a field has been set.

### GetEnableSpectatorAccess

`func (o *EventEnvelopeOneOf67Data) GetEnableSpectatorAccess() bool`

GetEnableSpectatorAccess returns the EnableSpectatorAccess field if non-nil, zero value otherwise.

### GetEnableSpectatorAccessOk

`func (o *EventEnvelopeOneOf67Data) GetEnableSpectatorAccessOk() (*bool, bool)`

GetEnableSpectatorAccessOk returns a tuple with the EnableSpectatorAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpectatorAccess

`func (o *EventEnvelopeOneOf67Data) SetEnableSpectatorAccess(v bool)`

SetEnableSpectatorAccess sets EnableSpectatorAccess field to given value.

### HasEnableSpectatorAccess

`func (o *EventEnvelopeOneOf67Data) HasEnableSpectatorAccess() bool`

HasEnableSpectatorAccess returns a boolean if a field has been set.

### GetGiphyRating

`func (o *EventEnvelopeOneOf67Data) GetGiphyRating() int32`

GetGiphyRating returns the GiphyRating field if non-nil, zero value otherwise.

### GetGiphyRatingOk

`func (o *EventEnvelopeOneOf67Data) GetGiphyRatingOk() (*int32, bool)`

GetGiphyRatingOk returns a tuple with the GiphyRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiphyRating

`func (o *EventEnvelopeOneOf67Data) SetGiphyRating(v int32)`

SetGiphyRating sets GiphyRating field to given value.

### HasGiphyRating

`func (o *EventEnvelopeOneOf67Data) HasGiphyRating() bool`

HasGiphyRating returns a boolean if a field has been set.

### GetIconSource

`func (o *EventEnvelopeOneOf67Data) GetIconSource() string`

GetIconSource returns the IconSource field if non-nil, zero value otherwise.

### GetIconSourceOk

`func (o *EventEnvelopeOneOf67Data) GetIconSourceOk() (*string, bool)`

GetIconSourceOk returns a tuple with the IconSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSource

`func (o *EventEnvelopeOneOf67Data) SetIconSource(v string)`

SetIconSource sets IconSource field to given value.

### HasIconSource

`func (o *EventEnvelopeOneOf67Data) HasIconSource() bool`

HasIconSource returns a boolean if a field has been set.

### GetIconUrl

`func (o *EventEnvelopeOneOf67Data) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *EventEnvelopeOneOf67Data) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *EventEnvelopeOneOf67Data) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *EventEnvelopeOneOf67Data) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetInlineImagePreview

`func (o *EventEnvelopeOneOf67Data) GetInlineImagePreview() bool`

GetInlineImagePreview returns the InlineImagePreview field if non-nil, zero value otherwise.

### GetInlineImagePreviewOk

`func (o *EventEnvelopeOneOf67Data) GetInlineImagePreviewOk() (*bool, bool)`

GetInlineImagePreviewOk returns a tuple with the InlineImagePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineImagePreview

`func (o *EventEnvelopeOneOf67Data) SetInlineImagePreview(v bool)`

SetInlineImagePreview sets InlineImagePreview field to given value.

### HasInlineImagePreview

`func (o *EventEnvelopeOneOf67Data) HasInlineImagePreview() bool`

HasInlineImagePreview returns a boolean if a field has been set.

### GetInlineUrlEmbedPreview

`func (o *EventEnvelopeOneOf67Data) GetInlineUrlEmbedPreview() bool`

GetInlineUrlEmbedPreview returns the InlineUrlEmbedPreview field if non-nil, zero value otherwise.

### GetInlineUrlEmbedPreviewOk

`func (o *EventEnvelopeOneOf67Data) GetInlineUrlEmbedPreviewOk() (*bool, bool)`

GetInlineUrlEmbedPreviewOk returns a tuple with the InlineUrlEmbedPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineUrlEmbedPreview

`func (o *EventEnvelopeOneOf67Data) SetInlineUrlEmbedPreview(v bool)`

SetInlineUrlEmbedPreview sets InlineUrlEmbedPreview field to given value.

### HasInlineUrlEmbedPreview

`func (o *EventEnvelopeOneOf67Data) HasInlineUrlEmbedPreview() bool`

HasInlineUrlEmbedPreview returns a boolean if a field has been set.

### GetInviteRequired

`func (o *EventEnvelopeOneOf67Data) GetInviteRequired() bool`

GetInviteRequired returns the InviteRequired field if non-nil, zero value otherwise.

### GetInviteRequiredOk

`func (o *EventEnvelopeOneOf67Data) GetInviteRequiredOk() (*bool, bool)`

GetInviteRequiredOk returns a tuple with the InviteRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteRequired

`func (o *EventEnvelopeOneOf67Data) SetInviteRequired(v bool)`

SetInviteRequired sets InviteRequired field to given value.

### HasInviteRequired

`func (o *EventEnvelopeOneOf67Data) HasInviteRequired() bool`

HasInviteRequired returns a boolean if a field has been set.

### GetJitsiServerUrl

`func (o *EventEnvelopeOneOf67Data) GetJitsiServerUrl() string`

GetJitsiServerUrl returns the JitsiServerUrl field if non-nil, zero value otherwise.

### GetJitsiServerUrlOk

`func (o *EventEnvelopeOneOf67Data) GetJitsiServerUrlOk() (*string, bool)`

GetJitsiServerUrlOk returns a tuple with the JitsiServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitsiServerUrl

`func (o *EventEnvelopeOneOf67Data) SetJitsiServerUrl(v string)`

SetJitsiServerUrl sets JitsiServerUrl field to given value.

### HasJitsiServerUrl

`func (o *EventEnvelopeOneOf67Data) HasJitsiServerUrl() bool`

HasJitsiServerUrl returns a boolean if a field has been set.

### SetJitsiServerUrlNil

`func (o *EventEnvelopeOneOf67Data) SetJitsiServerUrlNil(b bool)`

 SetJitsiServerUrlNil sets the value for JitsiServerUrl to be an explicit nil

### UnsetJitsiServerUrl
`func (o *EventEnvelopeOneOf67Data) UnsetJitsiServerUrl()`

UnsetJitsiServerUrl ensures that no value is present for JitsiServerUrl, not even an explicit nil
### GetLogoSource

`func (o *EventEnvelopeOneOf67Data) GetLogoSource() string`

GetLogoSource returns the LogoSource field if non-nil, zero value otherwise.

### GetLogoSourceOk

`func (o *EventEnvelopeOneOf67Data) GetLogoSourceOk() (*string, bool)`

GetLogoSourceOk returns a tuple with the LogoSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSource

`func (o *EventEnvelopeOneOf67Data) SetLogoSource(v string)`

SetLogoSource sets LogoSource field to given value.

### HasLogoSource

`func (o *EventEnvelopeOneOf67Data) HasLogoSource() bool`

HasLogoSource returns a boolean if a field has been set.

### GetLogoUrl

`func (o *EventEnvelopeOneOf67Data) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *EventEnvelopeOneOf67Data) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *EventEnvelopeOneOf67Data) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *EventEnvelopeOneOf67Data) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetTopicsPolicy

`func (o *EventEnvelopeOneOf67Data) GetTopicsPolicy() string`

GetTopicsPolicy returns the TopicsPolicy field if non-nil, zero value otherwise.

### GetTopicsPolicyOk

`func (o *EventEnvelopeOneOf67Data) GetTopicsPolicyOk() (*string, bool)`

GetTopicsPolicyOk returns a tuple with the TopicsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsPolicy

`func (o *EventEnvelopeOneOf67Data) SetTopicsPolicy(v string)`

SetTopicsPolicy sets TopicsPolicy field to given value.

### HasTopicsPolicy

`func (o *EventEnvelopeOneOf67Data) HasTopicsPolicy() bool`

HasTopicsPolicy returns a boolean if a field has been set.

### GetMandatoryTopics

`func (o *EventEnvelopeOneOf67Data) GetMandatoryTopics() bool`

GetMandatoryTopics returns the MandatoryTopics field if non-nil, zero value otherwise.

### GetMandatoryTopicsOk

`func (o *EventEnvelopeOneOf67Data) GetMandatoryTopicsOk() (*bool, bool)`

GetMandatoryTopicsOk returns a tuple with the MandatoryTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryTopics

`func (o *EventEnvelopeOneOf67Data) SetMandatoryTopics(v bool)`

SetMandatoryTopics sets MandatoryTopics field to given value.

### HasMandatoryTopics

`func (o *EventEnvelopeOneOf67Data) HasMandatoryTopics() bool`

HasMandatoryTopics returns a boolean if a field has been set.

### GetMaxFileUploadSizeMib

`func (o *EventEnvelopeOneOf67Data) GetMaxFileUploadSizeMib() int32`

GetMaxFileUploadSizeMib returns the MaxFileUploadSizeMib field if non-nil, zero value otherwise.

### GetMaxFileUploadSizeMibOk

`func (o *EventEnvelopeOneOf67Data) GetMaxFileUploadSizeMibOk() (*int32, bool)`

GetMaxFileUploadSizeMibOk returns a tuple with the MaxFileUploadSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileUploadSizeMib

`func (o *EventEnvelopeOneOf67Data) SetMaxFileUploadSizeMib(v int32)`

SetMaxFileUploadSizeMib sets MaxFileUploadSizeMib field to given value.

### HasMaxFileUploadSizeMib

`func (o *EventEnvelopeOneOf67Data) HasMaxFileUploadSizeMib() bool`

HasMaxFileUploadSizeMib returns a boolean if a field has been set.

### GetMessageContentAllowedInEmailNotifications

`func (o *EventEnvelopeOneOf67Data) GetMessageContentAllowedInEmailNotifications() bool`

GetMessageContentAllowedInEmailNotifications returns the MessageContentAllowedInEmailNotifications field if non-nil, zero value otherwise.

### GetMessageContentAllowedInEmailNotificationsOk

`func (o *EventEnvelopeOneOf67Data) GetMessageContentAllowedInEmailNotificationsOk() (*bool, bool)`

GetMessageContentAllowedInEmailNotificationsOk returns a tuple with the MessageContentAllowedInEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContentAllowedInEmailNotifications

`func (o *EventEnvelopeOneOf67Data) SetMessageContentAllowedInEmailNotifications(v bool)`

SetMessageContentAllowedInEmailNotifications sets MessageContentAllowedInEmailNotifications field to given value.

### HasMessageContentAllowedInEmailNotifications

`func (o *EventEnvelopeOneOf67Data) HasMessageContentAllowedInEmailNotifications() bool`

HasMessageContentAllowedInEmailNotifications returns a boolean if a field has been set.

### GetMessageContentDeleteLimitSeconds

`func (o *EventEnvelopeOneOf67Data) GetMessageContentDeleteLimitSeconds() int32`

GetMessageContentDeleteLimitSeconds returns the MessageContentDeleteLimitSeconds field if non-nil, zero value otherwise.

### GetMessageContentDeleteLimitSecondsOk

`func (o *EventEnvelopeOneOf67Data) GetMessageContentDeleteLimitSecondsOk() (*int32, bool)`

GetMessageContentDeleteLimitSecondsOk returns a tuple with the MessageContentDeleteLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContentDeleteLimitSeconds

`func (o *EventEnvelopeOneOf67Data) SetMessageContentDeleteLimitSeconds(v int32)`

SetMessageContentDeleteLimitSeconds sets MessageContentDeleteLimitSeconds field to given value.

### HasMessageContentDeleteLimitSeconds

`func (o *EventEnvelopeOneOf67Data) HasMessageContentDeleteLimitSeconds() bool`

HasMessageContentDeleteLimitSeconds returns a boolean if a field has been set.

### SetMessageContentDeleteLimitSecondsNil

`func (o *EventEnvelopeOneOf67Data) SetMessageContentDeleteLimitSecondsNil(b bool)`

 SetMessageContentDeleteLimitSecondsNil sets the value for MessageContentDeleteLimitSeconds to be an explicit nil

### UnsetMessageContentDeleteLimitSeconds
`func (o *EventEnvelopeOneOf67Data) UnsetMessageContentDeleteLimitSeconds()`

UnsetMessageContentDeleteLimitSeconds ensures that no value is present for MessageContentDeleteLimitSeconds, not even an explicit nil
### GetMessageContentEditLimitSeconds

`func (o *EventEnvelopeOneOf67Data) GetMessageContentEditLimitSeconds() int32`

GetMessageContentEditLimitSeconds returns the MessageContentEditLimitSeconds field if non-nil, zero value otherwise.

### GetMessageContentEditLimitSecondsOk

`func (o *EventEnvelopeOneOf67Data) GetMessageContentEditLimitSecondsOk() (*int32, bool)`

GetMessageContentEditLimitSecondsOk returns a tuple with the MessageContentEditLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContentEditLimitSeconds

`func (o *EventEnvelopeOneOf67Data) SetMessageContentEditLimitSeconds(v int32)`

SetMessageContentEditLimitSeconds sets MessageContentEditLimitSeconds field to given value.

### HasMessageContentEditLimitSeconds

`func (o *EventEnvelopeOneOf67Data) HasMessageContentEditLimitSeconds() bool`

HasMessageContentEditLimitSeconds returns a boolean if a field has been set.

### SetMessageContentEditLimitSecondsNil

`func (o *EventEnvelopeOneOf67Data) SetMessageContentEditLimitSecondsNil(b bool)`

 SetMessageContentEditLimitSecondsNil sets the value for MessageContentEditLimitSeconds to be an explicit nil

### UnsetMessageContentEditLimitSeconds
`func (o *EventEnvelopeOneOf67Data) UnsetMessageContentEditLimitSeconds()`

UnsetMessageContentEditLimitSeconds ensures that no value is present for MessageContentEditLimitSeconds, not even an explicit nil
### GetMessageEditHistoryVisibilityPolicy

`func (o *EventEnvelopeOneOf67Data) GetMessageEditHistoryVisibilityPolicy() string`

GetMessageEditHistoryVisibilityPolicy returns the MessageEditHistoryVisibilityPolicy field if non-nil, zero value otherwise.

### GetMessageEditHistoryVisibilityPolicyOk

`func (o *EventEnvelopeOneOf67Data) GetMessageEditHistoryVisibilityPolicyOk() (*string, bool)`

GetMessageEditHistoryVisibilityPolicyOk returns a tuple with the MessageEditHistoryVisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEditHistoryVisibilityPolicy

`func (o *EventEnvelopeOneOf67Data) SetMessageEditHistoryVisibilityPolicy(v string)`

SetMessageEditHistoryVisibilityPolicy sets MessageEditHistoryVisibilityPolicy field to given value.

### HasMessageEditHistoryVisibilityPolicy

`func (o *EventEnvelopeOneOf67Data) HasMessageEditHistoryVisibilityPolicy() bool`

HasMessageEditHistoryVisibilityPolicy returns a boolean if a field has been set.

### GetModerationRequestChannelId

`func (o *EventEnvelopeOneOf67Data) GetModerationRequestChannelId() int32`

GetModerationRequestChannelId returns the ModerationRequestChannelId field if non-nil, zero value otherwise.

### GetModerationRequestChannelIdOk

`func (o *EventEnvelopeOneOf67Data) GetModerationRequestChannelIdOk() (*int32, bool)`

GetModerationRequestChannelIdOk returns a tuple with the ModerationRequestChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationRequestChannelId

`func (o *EventEnvelopeOneOf67Data) SetModerationRequestChannelId(v int32)`

SetModerationRequestChannelId sets ModerationRequestChannelId field to given value.

### HasModerationRequestChannelId

`func (o *EventEnvelopeOneOf67Data) HasModerationRequestChannelId() bool`

HasModerationRequestChannelId returns a boolean if a field has been set.

### GetMoveMessagesWithinStreamLimitSeconds

`func (o *EventEnvelopeOneOf67Data) GetMoveMessagesWithinStreamLimitSeconds() int32`

GetMoveMessagesWithinStreamLimitSeconds returns the MoveMessagesWithinStreamLimitSeconds field if non-nil, zero value otherwise.

### GetMoveMessagesWithinStreamLimitSecondsOk

`func (o *EventEnvelopeOneOf67Data) GetMoveMessagesWithinStreamLimitSecondsOk() (*int32, bool)`

GetMoveMessagesWithinStreamLimitSecondsOk returns a tuple with the MoveMessagesWithinStreamLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveMessagesWithinStreamLimitSeconds

`func (o *EventEnvelopeOneOf67Data) SetMoveMessagesWithinStreamLimitSeconds(v int32)`

SetMoveMessagesWithinStreamLimitSeconds sets MoveMessagesWithinStreamLimitSeconds field to given value.

### HasMoveMessagesWithinStreamLimitSeconds

`func (o *EventEnvelopeOneOf67Data) HasMoveMessagesWithinStreamLimitSeconds() bool`

HasMoveMessagesWithinStreamLimitSeconds returns a boolean if a field has been set.

### SetMoveMessagesWithinStreamLimitSecondsNil

`func (o *EventEnvelopeOneOf67Data) SetMoveMessagesWithinStreamLimitSecondsNil(b bool)`

 SetMoveMessagesWithinStreamLimitSecondsNil sets the value for MoveMessagesWithinStreamLimitSeconds to be an explicit nil

### UnsetMoveMessagesWithinStreamLimitSeconds
`func (o *EventEnvelopeOneOf67Data) UnsetMoveMessagesWithinStreamLimitSeconds()`

UnsetMoveMessagesWithinStreamLimitSeconds ensures that no value is present for MoveMessagesWithinStreamLimitSeconds, not even an explicit nil
### GetMoveMessagesBetweenStreamsLimitSeconds

`func (o *EventEnvelopeOneOf67Data) GetMoveMessagesBetweenStreamsLimitSeconds() int32`

GetMoveMessagesBetweenStreamsLimitSeconds returns the MoveMessagesBetweenStreamsLimitSeconds field if non-nil, zero value otherwise.

### GetMoveMessagesBetweenStreamsLimitSecondsOk

`func (o *EventEnvelopeOneOf67Data) GetMoveMessagesBetweenStreamsLimitSecondsOk() (*int32, bool)`

GetMoveMessagesBetweenStreamsLimitSecondsOk returns a tuple with the MoveMessagesBetweenStreamsLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveMessagesBetweenStreamsLimitSeconds

`func (o *EventEnvelopeOneOf67Data) SetMoveMessagesBetweenStreamsLimitSeconds(v int32)`

SetMoveMessagesBetweenStreamsLimitSeconds sets MoveMessagesBetweenStreamsLimitSeconds field to given value.

### HasMoveMessagesBetweenStreamsLimitSeconds

`func (o *EventEnvelopeOneOf67Data) HasMoveMessagesBetweenStreamsLimitSeconds() bool`

HasMoveMessagesBetweenStreamsLimitSeconds returns a boolean if a field has been set.

### SetMoveMessagesBetweenStreamsLimitSecondsNil

`func (o *EventEnvelopeOneOf67Data) SetMoveMessagesBetweenStreamsLimitSecondsNil(b bool)`

 SetMoveMessagesBetweenStreamsLimitSecondsNil sets the value for MoveMessagesBetweenStreamsLimitSeconds to be an explicit nil

### UnsetMoveMessagesBetweenStreamsLimitSeconds
`func (o *EventEnvelopeOneOf67Data) UnsetMoveMessagesBetweenStreamsLimitSeconds()`

UnsetMoveMessagesBetweenStreamsLimitSeconds ensures that no value is present for MoveMessagesBetweenStreamsLimitSeconds, not even an explicit nil
### GetName

`func (o *EventEnvelopeOneOf67Data) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventEnvelopeOneOf67Data) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventEnvelopeOneOf67Data) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventEnvelopeOneOf67Data) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameChangesDisabled

`func (o *EventEnvelopeOneOf67Data) GetNameChangesDisabled() bool`

GetNameChangesDisabled returns the NameChangesDisabled field if non-nil, zero value otherwise.

### GetNameChangesDisabledOk

`func (o *EventEnvelopeOneOf67Data) GetNameChangesDisabledOk() (*bool, bool)`

GetNameChangesDisabledOk returns a tuple with the NameChangesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameChangesDisabled

`func (o *EventEnvelopeOneOf67Data) SetNameChangesDisabled(v bool)`

SetNameChangesDisabled sets NameChangesDisabled field to given value.

### HasNameChangesDisabled

`func (o *EventEnvelopeOneOf67Data) HasNameChangesDisabled() bool`

HasNameChangesDisabled returns a boolean if a field has been set.

### GetNightLogoSource

`func (o *EventEnvelopeOneOf67Data) GetNightLogoSource() string`

GetNightLogoSource returns the NightLogoSource field if non-nil, zero value otherwise.

### GetNightLogoSourceOk

`func (o *EventEnvelopeOneOf67Data) GetNightLogoSourceOk() (*string, bool)`

GetNightLogoSourceOk returns a tuple with the NightLogoSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNightLogoSource

`func (o *EventEnvelopeOneOf67Data) SetNightLogoSource(v string)`

SetNightLogoSource sets NightLogoSource field to given value.

### HasNightLogoSource

`func (o *EventEnvelopeOneOf67Data) HasNightLogoSource() bool`

HasNightLogoSource returns a boolean if a field has been set.

### GetNightLogoUrl

`func (o *EventEnvelopeOneOf67Data) GetNightLogoUrl() string`

GetNightLogoUrl returns the NightLogoUrl field if non-nil, zero value otherwise.

### GetNightLogoUrlOk

`func (o *EventEnvelopeOneOf67Data) GetNightLogoUrlOk() (*string, bool)`

GetNightLogoUrlOk returns a tuple with the NightLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNightLogoUrl

`func (o *EventEnvelopeOneOf67Data) SetNightLogoUrl(v string)`

SetNightLogoUrl sets NightLogoUrl field to given value.

### HasNightLogoUrl

`func (o *EventEnvelopeOneOf67Data) HasNightLogoUrl() bool`

HasNightLogoUrl returns a boolean if a field has been set.

### GetNewStreamAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) GetNewStreamAnnouncementsStreamId() int32`

GetNewStreamAnnouncementsStreamId returns the NewStreamAnnouncementsStreamId field if non-nil, zero value otherwise.

### GetNewStreamAnnouncementsStreamIdOk

`func (o *EventEnvelopeOneOf67Data) GetNewStreamAnnouncementsStreamIdOk() (*int32, bool)`

GetNewStreamAnnouncementsStreamIdOk returns a tuple with the NewStreamAnnouncementsStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStreamAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) SetNewStreamAnnouncementsStreamId(v int32)`

SetNewStreamAnnouncementsStreamId sets NewStreamAnnouncementsStreamId field to given value.

### HasNewStreamAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) HasNewStreamAnnouncementsStreamId() bool`

HasNewStreamAnnouncementsStreamId returns a boolean if a field has been set.

### GetOrgType

`func (o *EventEnvelopeOneOf67Data) GetOrgType() int32`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *EventEnvelopeOneOf67Data) GetOrgTypeOk() (*int32, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *EventEnvelopeOneOf67Data) SetOrgType(v int32)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *EventEnvelopeOneOf67Data) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### GetPlanType

`func (o *EventEnvelopeOneOf67Data) GetPlanType() int32`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *EventEnvelopeOneOf67Data) GetPlanTypeOk() (*int32, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *EventEnvelopeOneOf67Data) SetPlanType(v int32)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *EventEnvelopeOneOf67Data) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetPresenceDisabled

`func (o *EventEnvelopeOneOf67Data) GetPresenceDisabled() bool`

GetPresenceDisabled returns the PresenceDisabled field if non-nil, zero value otherwise.

### GetPresenceDisabledOk

`func (o *EventEnvelopeOneOf67Data) GetPresenceDisabledOk() (*bool, bool)`

GetPresenceDisabledOk returns a tuple with the PresenceDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceDisabled

`func (o *EventEnvelopeOneOf67Data) SetPresenceDisabled(v bool)`

SetPresenceDisabled sets PresenceDisabled field to given value.

### HasPresenceDisabled

`func (o *EventEnvelopeOneOf67Data) HasPresenceDisabled() bool`

HasPresenceDisabled returns a boolean if a field has been set.

### GetPushNotificationsEnabled

`func (o *EventEnvelopeOneOf67Data) GetPushNotificationsEnabled() bool`

GetPushNotificationsEnabled returns the PushNotificationsEnabled field if non-nil, zero value otherwise.

### GetPushNotificationsEnabledOk

`func (o *EventEnvelopeOneOf67Data) GetPushNotificationsEnabledOk() (*bool, bool)`

GetPushNotificationsEnabledOk returns a tuple with the PushNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotificationsEnabled

`func (o *EventEnvelopeOneOf67Data) SetPushNotificationsEnabled(v bool)`

SetPushNotificationsEnabled sets PushNotificationsEnabled field to given value.

### HasPushNotificationsEnabled

`func (o *EventEnvelopeOneOf67Data) HasPushNotificationsEnabled() bool`

HasPushNotificationsEnabled returns a boolean if a field has been set.

### GetPushNotificationsEnabledEndTimestamp

`func (o *EventEnvelopeOneOf67Data) GetPushNotificationsEnabledEndTimestamp() int32`

GetPushNotificationsEnabledEndTimestamp returns the PushNotificationsEnabledEndTimestamp field if non-nil, zero value otherwise.

### GetPushNotificationsEnabledEndTimestampOk

`func (o *EventEnvelopeOneOf67Data) GetPushNotificationsEnabledEndTimestampOk() (*int32, bool)`

GetPushNotificationsEnabledEndTimestampOk returns a tuple with the PushNotificationsEnabledEndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotificationsEnabledEndTimestamp

`func (o *EventEnvelopeOneOf67Data) SetPushNotificationsEnabledEndTimestamp(v int32)`

SetPushNotificationsEnabledEndTimestamp sets PushNotificationsEnabledEndTimestamp field to given value.

### HasPushNotificationsEnabledEndTimestamp

`func (o *EventEnvelopeOneOf67Data) HasPushNotificationsEnabledEndTimestamp() bool`

HasPushNotificationsEnabledEndTimestamp returns a boolean if a field has been set.

### SetPushNotificationsEnabledEndTimestampNil

`func (o *EventEnvelopeOneOf67Data) SetPushNotificationsEnabledEndTimestampNil(b bool)`

 SetPushNotificationsEnabledEndTimestampNil sets the value for PushNotificationsEnabledEndTimestamp to be an explicit nil

### UnsetPushNotificationsEnabledEndTimestamp
`func (o *EventEnvelopeOneOf67Data) UnsetPushNotificationsEnabledEndTimestamp()`

UnsetPushNotificationsEnabledEndTimestamp ensures that no value is present for PushNotificationsEnabledEndTimestamp, not even an explicit nil
### GetRequireE2eePushNotifications

`func (o *EventEnvelopeOneOf67Data) GetRequireE2eePushNotifications() bool`

GetRequireE2eePushNotifications returns the RequireE2eePushNotifications field if non-nil, zero value otherwise.

### GetRequireE2eePushNotificationsOk

`func (o *EventEnvelopeOneOf67Data) GetRequireE2eePushNotificationsOk() (*bool, bool)`

GetRequireE2eePushNotificationsOk returns a tuple with the RequireE2eePushNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireE2eePushNotifications

`func (o *EventEnvelopeOneOf67Data) SetRequireE2eePushNotifications(v bool)`

SetRequireE2eePushNotifications sets RequireE2eePushNotifications field to given value.

### HasRequireE2eePushNotifications

`func (o *EventEnvelopeOneOf67Data) HasRequireE2eePushNotifications() bool`

HasRequireE2eePushNotifications returns a boolean if a field has been set.

### GetRequireUniqueNames

`func (o *EventEnvelopeOneOf67Data) GetRequireUniqueNames() bool`

GetRequireUniqueNames returns the RequireUniqueNames field if non-nil, zero value otherwise.

### GetRequireUniqueNamesOk

`func (o *EventEnvelopeOneOf67Data) GetRequireUniqueNamesOk() (*bool, bool)`

GetRequireUniqueNamesOk returns a tuple with the RequireUniqueNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireUniqueNames

`func (o *EventEnvelopeOneOf67Data) SetRequireUniqueNames(v bool)`

SetRequireUniqueNames sets RequireUniqueNames field to given value.

### HasRequireUniqueNames

`func (o *EventEnvelopeOneOf67Data) HasRequireUniqueNames() bool`

HasRequireUniqueNames returns a boolean if a field has been set.

### GetSendWelcomeEmails

`func (o *EventEnvelopeOneOf67Data) GetSendWelcomeEmails() bool`

GetSendWelcomeEmails returns the SendWelcomeEmails field if non-nil, zero value otherwise.

### GetSendWelcomeEmailsOk

`func (o *EventEnvelopeOneOf67Data) GetSendWelcomeEmailsOk() (*bool, bool)`

GetSendWelcomeEmailsOk returns a tuple with the SendWelcomeEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWelcomeEmails

`func (o *EventEnvelopeOneOf67Data) SetSendWelcomeEmails(v bool)`

SetSendWelcomeEmails sets SendWelcomeEmails field to given value.

### HasSendWelcomeEmails

`func (o *EventEnvelopeOneOf67Data) HasSendWelcomeEmails() bool`

HasSendWelcomeEmails returns a boolean if a field has been set.

### GetSignupAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) GetSignupAnnouncementsStreamId() int32`

GetSignupAnnouncementsStreamId returns the SignupAnnouncementsStreamId field if non-nil, zero value otherwise.

### GetSignupAnnouncementsStreamIdOk

`func (o *EventEnvelopeOneOf67Data) GetSignupAnnouncementsStreamIdOk() (*int32, bool)`

GetSignupAnnouncementsStreamIdOk returns a tuple with the SignupAnnouncementsStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) SetSignupAnnouncementsStreamId(v int32)`

SetSignupAnnouncementsStreamId sets SignupAnnouncementsStreamId field to given value.

### HasSignupAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) HasSignupAnnouncementsStreamId() bool`

HasSignupAnnouncementsStreamId returns a boolean if a field has been set.

### GetUploadQuotaMib

`func (o *EventEnvelopeOneOf67Data) GetUploadQuotaMib() int32`

GetUploadQuotaMib returns the UploadQuotaMib field if non-nil, zero value otherwise.

### GetUploadQuotaMibOk

`func (o *EventEnvelopeOneOf67Data) GetUploadQuotaMibOk() (*int32, bool)`

GetUploadQuotaMibOk returns a tuple with the UploadQuotaMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadQuotaMib

`func (o *EventEnvelopeOneOf67Data) SetUploadQuotaMib(v int32)`

SetUploadQuotaMib sets UploadQuotaMib field to given value.

### HasUploadQuotaMib

`func (o *EventEnvelopeOneOf67Data) HasUploadQuotaMib() bool`

HasUploadQuotaMib returns a boolean if a field has been set.

### SetUploadQuotaMibNil

`func (o *EventEnvelopeOneOf67Data) SetUploadQuotaMibNil(b bool)`

 SetUploadQuotaMibNil sets the value for UploadQuotaMib to be an explicit nil

### UnsetUploadQuotaMib
`func (o *EventEnvelopeOneOf67Data) UnsetUploadQuotaMib()`

UnsetUploadQuotaMib ensures that no value is present for UploadQuotaMib, not even an explicit nil
### GetVideoChatProvider

`func (o *EventEnvelopeOneOf67Data) GetVideoChatProvider() int32`

GetVideoChatProvider returns the VideoChatProvider field if non-nil, zero value otherwise.

### GetVideoChatProviderOk

`func (o *EventEnvelopeOneOf67Data) GetVideoChatProviderOk() (*int32, bool)`

GetVideoChatProviderOk returns a tuple with the VideoChatProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChatProvider

`func (o *EventEnvelopeOneOf67Data) SetVideoChatProvider(v int32)`

SetVideoChatProvider sets VideoChatProvider field to given value.

### HasVideoChatProvider

`func (o *EventEnvelopeOneOf67Data) HasVideoChatProvider() bool`

HasVideoChatProvider returns a boolean if a field has been set.

### GetWaitingPeriodThreshold

`func (o *EventEnvelopeOneOf67Data) GetWaitingPeriodThreshold() int32`

GetWaitingPeriodThreshold returns the WaitingPeriodThreshold field if non-nil, zero value otherwise.

### GetWaitingPeriodThresholdOk

`func (o *EventEnvelopeOneOf67Data) GetWaitingPeriodThresholdOk() (*int32, bool)`

GetWaitingPeriodThresholdOk returns a tuple with the WaitingPeriodThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingPeriodThreshold

`func (o *EventEnvelopeOneOf67Data) SetWaitingPeriodThreshold(v int32)`

SetWaitingPeriodThreshold sets WaitingPeriodThreshold field to given value.

### HasWaitingPeriodThreshold

`func (o *EventEnvelopeOneOf67Data) HasWaitingPeriodThreshold() bool`

HasWaitingPeriodThreshold returns a boolean if a field has been set.

### GetWantAdvertiseInCommunitiesDirectory

`func (o *EventEnvelopeOneOf67Data) GetWantAdvertiseInCommunitiesDirectory() bool`

GetWantAdvertiseInCommunitiesDirectory returns the WantAdvertiseInCommunitiesDirectory field if non-nil, zero value otherwise.

### GetWantAdvertiseInCommunitiesDirectoryOk

`func (o *EventEnvelopeOneOf67Data) GetWantAdvertiseInCommunitiesDirectoryOk() (*bool, bool)`

GetWantAdvertiseInCommunitiesDirectoryOk returns a tuple with the WantAdvertiseInCommunitiesDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantAdvertiseInCommunitiesDirectory

`func (o *EventEnvelopeOneOf67Data) SetWantAdvertiseInCommunitiesDirectory(v bool)`

SetWantAdvertiseInCommunitiesDirectory sets WantAdvertiseInCommunitiesDirectory field to given value.

### HasWantAdvertiseInCommunitiesDirectory

`func (o *EventEnvelopeOneOf67Data) HasWantAdvertiseInCommunitiesDirectory() bool`

HasWantAdvertiseInCommunitiesDirectory returns a boolean if a field has been set.

### GetWelcomeMessageCustomText

`func (o *EventEnvelopeOneOf67Data) GetWelcomeMessageCustomText() string`

GetWelcomeMessageCustomText returns the WelcomeMessageCustomText field if non-nil, zero value otherwise.

### GetWelcomeMessageCustomTextOk

`func (o *EventEnvelopeOneOf67Data) GetWelcomeMessageCustomTextOk() (*string, bool)`

GetWelcomeMessageCustomTextOk returns a tuple with the WelcomeMessageCustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessageCustomText

`func (o *EventEnvelopeOneOf67Data) SetWelcomeMessageCustomText(v string)`

SetWelcomeMessageCustomText sets WelcomeMessageCustomText field to given value.

### HasWelcomeMessageCustomText

`func (o *EventEnvelopeOneOf67Data) HasWelcomeMessageCustomText() bool`

HasWelcomeMessageCustomText returns a boolean if a field has been set.

### GetZulipUpdateAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) GetZulipUpdateAnnouncementsStreamId() int32`

GetZulipUpdateAnnouncementsStreamId returns the ZulipUpdateAnnouncementsStreamId field if non-nil, zero value otherwise.

### GetZulipUpdateAnnouncementsStreamIdOk

`func (o *EventEnvelopeOneOf67Data) GetZulipUpdateAnnouncementsStreamIdOk() (*int32, bool)`

GetZulipUpdateAnnouncementsStreamIdOk returns a tuple with the ZulipUpdateAnnouncementsStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipUpdateAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) SetZulipUpdateAnnouncementsStreamId(v int32)`

SetZulipUpdateAnnouncementsStreamId sets ZulipUpdateAnnouncementsStreamId field to given value.

### HasZulipUpdateAnnouncementsStreamId

`func (o *EventEnvelopeOneOf67Data) HasZulipUpdateAnnouncementsStreamId() bool`

HasZulipUpdateAnnouncementsStreamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


