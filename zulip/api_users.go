package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type UsersAPI interface {

	/*
			AddAlertWords Add alert words

			Add words (or phrases) to the user's set of configured [alert words][alert-words].

		[alert-words]: /help/dm-mention-alert-notifications#alert-words


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return AddAlertWordsRequest
	*/
	AddAlertWords(ctx context.Context) AddAlertWordsRequest

	// AddAlertWordsExecute executes the request
	//  @return AlertWordsResponse
	AddAlertWordsExecute(r AddAlertWordsRequest) (*AlertWordsResponse, *http.Response, error)

	/*
		AddApnsToken Add an APNs device token

		This endpoint adds an APNs device token to register for iOS push notifications.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return AddApnsTokenRequest
	*/
	AddApnsToken(ctx context.Context) AddApnsTokenRequest

	// AddApnsTokenExecute executes the request
	//  @return Response
	AddApnsTokenExecute(r AddApnsTokenRequest) (*Response, *http.Response, error)

	/*
		AddFcmToken Add an FCM registration token

		This endpoint adds an FCM registration token for push notifications.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return AddFcmTokenRequest
	*/
	AddFcmToken(ctx context.Context) AddFcmTokenRequest

	// AddFcmTokenExecute executes the request
	//  @return Response
	AddFcmTokenExecute(r AddFcmTokenRequest) (*Response, *http.Response, error)

	/*
			CreateUser Create a user

			Create a new user account via the API.

		!!! warn ""

		    **Note**: On Zulip Cloud, this feature is available only for
		    organizations on a [Zulip Cloud Standard](https://zulip.com/plans/)
		    or [Zulip Cloud Plus](https://zulip.com/plans/) plan. Administrators
		    can request the required `can_create_users` permission for a bot or
		    user by contacting [Zulip Cloud support][support] with an
		    explanation for why it is needed. Self-hosted installations can
		    toggle `can_create_users` on an account using the `manage.py
		    change_user_role` [management command][management-commands].

		**Changes**: Before Zulip 4.0 (feature level 36), this endpoint was
		available to all organization administrators.

		[support]: /help/contact-support
		[management-commands]: https://zulip.readthedocs.io/en/latest/production/management-commands.html


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return CreateUserRequest
	*/
	CreateUser(ctx context.Context) CreateUserRequest

	// CreateUserExecute executes the request
	//  @return CreateUserResponse
	CreateUserExecute(r CreateUserRequest) (*CreateUserResponse, *http.Response, error)

	/*
		CreateUserGroup Create a user group

		Create a new [user group](zulip.com/help/user-groups.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return CreateUserGroupRequest
	*/
	CreateUserGroup(ctx context.Context) CreateUserGroupRequest

	// CreateUserGroupExecute executes the request
	//  @return CreateUserGroupResponse
	CreateUserGroupExecute(r CreateUserGroupRequest) (*CreateUserGroupResponse, *http.Response, error)

	/*
			DeactivateOwnUser Deactivate own user

			Deactivates the current user's account. See also the administrative endpoint for
		[deactivating another user](zulip.com/api/deactivate-user.

		This endpoint is primarily useful to Zulip clients providing a user settings UI.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return DeactivateOwnUserRequest
	*/
	DeactivateOwnUser(ctx context.Context) DeactivateOwnUserRequest

	// DeactivateOwnUserExecute executes the request
	//  @return Response
	DeactivateOwnUserExecute(r DeactivateOwnUserRequest) (*Response, *http.Response, error)

	/*
			DeactivateUser Deactivate a user

			[Deactivates a
		user](https://zulip.com/help/deactivate-or-reactivate-a-user)
		given their user Id.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId The target user's Id.
			@return DeactivateUserRequest
	*/
	DeactivateUser(ctx context.Context, userId int64) DeactivateUserRequest

	// DeactivateUserExecute executes the request
	//  @return Response
	DeactivateUserExecute(r DeactivateUserRequest) (*Response, *http.Response, error)

	/*
			DeactivateUserGroup Deactivate a user group

			Deactivate a user group. Deactivated user groups cannot be
		used for mentions, permissions, or any other purpose, but can
		be reactivated or renamed.

		Deactivating user groups is preferable to deleting them from
		the database, since the deactivation model allows audit logs
		of changes to sensitive group-valued permissions to be
		maintained.

		**Changes**: New in Zulip 10.0 (feature level 290).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userGroupId The Id of the target user group.
			@return DeactivateUserGroupRequest
	*/
	DeactivateUserGroup(ctx context.Context, userGroupId int64) DeactivateUserGroupRequest

	// DeactivateUserGroupExecute executes the request
	//  @return Response
	DeactivateUserGroupExecute(r DeactivateUserGroupRequest) (*Response, *http.Response, error)

	/*
			GetAlertWords Get all alert words

			Get all of the user's configured [alert words][alert-words].

		[alert-words]: /help/dm-mention-alert-notifications#alert-words


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GetAlertWordsRequest
	*/
	GetAlertWords(ctx context.Context) GetAlertWordsRequest

	// GetAlertWordsExecute executes the request
	//  @return AlertWordsResponse
	GetAlertWordsExecute(r GetAlertWordsRequest) (*AlertWordsResponse, *http.Response, error)

	/*
		GetAttachments Get attachments

		Fetch metadata on files uploaded by the requesting user.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetAttachmentsRequest
	*/
	GetAttachments(ctx context.Context) GetAttachmentsRequest

	// GetAttachmentsExecute executes the request
	//  @return GetAttachmentsResponse
	GetAttachmentsExecute(r GetAttachmentsRequest) (*GetAttachmentsResponse, *http.Response, error)

	/*
			GetIsUserGroupMember Get user group membership status

			Check whether a user is member of user group.

		**Changes**: Prior to Zulip 10.0 (feature level 303),
		this would return true when passed a deactivated user
		who was a member of the user group before being deactivated.

		New in Zulip 6.0 (feature level 127).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userGroupId The Id of the target user group.
			@param userId The target user's Id.
			@return GetIsUserGroupMemberRequest
	*/
	GetIsUserGroupMember(ctx context.Context, userGroupId int64, userId int64) GetIsUserGroupMemberRequest

	// GetIsUserGroupMemberExecute executes the request
	//  @return GetIsUserGroupMemberResponse
	GetIsUserGroupMemberExecute(r GetIsUserGroupMemberRequest) (*GetIsUserGroupMemberResponse, *http.Response, error)

	/*
			GetOwnUser Get own user

			Get basic data about the user/bot that requests this endpoint.

		**Changes**: Removed `is_billing_admin` field in Zulip 10.0 (feature level 363), as it was
		replaced by the `can_manage_billing_group` realm setting.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GetOwnUserRequest
	*/
	GetOwnUser(ctx context.Context) GetOwnUserRequest

	// GetOwnUserExecute executes the request
	//  @return GetOwnUserResponse
	GetOwnUserExecute(r GetOwnUserRequest) (*GetOwnUserResponse, *http.Response, error)

	/*
			GetUser Get a user

			Fetch details for a single user in the organization.

		You can also fetch details on [all users in the organization](zulip.com/api/get-users
		or [by a user's Zulip API email](zulip.com/api/get-user-by-email.

		**Changes**: New in Zulip 3.0 (feature level 1).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId The target user's Id.
			@return GetUserRequest
	*/
	GetUser(ctx context.Context, userId int64) GetUserRequest

	// GetUserExecute executes the request
	//  @return GetUserResponse
	GetUserExecute(r GetUserRequest) (*GetUserResponse, *http.Response, error)

	/*
			GetUserByEmail Get a user by email

			Fetch details for a single user in the organization given a Zulip
		API email address.

		You can also fetch details on [all users in the organization](zulip.com/api/get-users
		or [by user Id](zulip.com/api/get-user.

		Fetching by user Id is generally recommended when possible,
		as a user might [change their email address](zulip.com/help/change-your-email-address
		or change their [email address visibility](zulip.com/help/configure-email-visibility,
		either of which could change the client's ability to look them up by that
		email address.

		**Changes**: Starting with Zulip 10.0 (feature level 302), the real email
		address can be used in the `email` parameter and will fetch the target user's
		data if and only if the target's email visibility setting permits the requester
		to see the email address.
		The dummy email addresses of the form `user{id}@{realm.host}` still work, and
		will now work for **all** users, via identifying them by the embedded user Id.

		New in Zulip Server 4.0 (feature level 39).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param email The email address of the user to fetch. Two forms are supported:  - The real email address of the user (`delivery_email`). The lookup will   succeed if and only if the user exists and their email address visibility   setting permits the client to see the email address.  - The dummy Zulip API email address of the form `user{user_id}@{realm_host}`. This   is identical to simply [getting user by Id](zulip.com/api/get-user. If the server or   realm change domains, the dummy email address used has to be adjustment to   match the new realm domain. This is legacy behavior for   backwards-compatibility, and will be removed in a future release.  **Changes**: Starting with Zulip 10.0 (feature level 302), lookups by real email address match the semantics of the target's email visibility setting and dummy email addresses work for all users, independently of their email visibility setting.  Previously, lookups were done only using the Zulip API email addresses.
			@return GetUserByEmailRequest
	*/
	GetUserByEmail(ctx context.Context, email string) GetUserByEmailRequest

	// GetUserByEmailExecute executes the request
	//  @return GetUserResponse
	GetUserByEmailExecute(r GetUserByEmailRequest) (*GetUserResponse, *http.Response, error)

	/*
			GetUserGroupMembers Get user group members

			Get the members of a [user group](zulip.com/help/user-groups.

		**Changes**: New in Zulip 6.0 (feature level 127).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userGroupId The Id of the target user group.
			@return GetUserGroupMembersRequest
	*/
	GetUserGroupMembers(ctx context.Context, userGroupId int64) GetUserGroupMembersRequest

	// GetUserGroupMembersExecute executes the request
	//  @return GetUserGroupMembersResponse
	GetUserGroupMembersExecute(r GetUserGroupMembersRequest) (*GetUserGroupMembersResponse, *http.Response, error)

	/*
			GetUserGroupSubgroups Get subgroups of a user group

			Get the subgroups of a [user group](zulip.com/help/user-groups.

		**Changes**: New in Zulip 6.0 (feature level 127).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userGroupId The Id of the target user group.
			@return GetUserGroupSubgroupsRequest
	*/
	GetUserGroupSubgroups(ctx context.Context, userGroupId int64) GetUserGroupSubgroupsRequest

	// GetUserGroupSubgroupsExecute executes the request
	//  @return GetUserGroupSubgroupsResponse
	GetUserGroupSubgroupsExecute(r GetUserGroupSubgroupsRequest) (*GetUserGroupSubgroupsResponse, *http.Response, error)

	/*
			GetUserGroups Get user groups

			Fetches all of the user groups in the organization.

		!!! warn ""

		    **Note**: This endpoint is only available to [members and
		    administrators](zulip.com/help/user-roles; bots and guests
		    cannot use this endpoint.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GetUserGroupsRequest
	*/
	GetUserGroups(ctx context.Context) GetUserGroupsRequest

	// GetUserGroupsExecute executes the request
	//  @return GetUserGroupsResponse
	GetUserGroupsExecute(r GetUserGroupsRequest) (*GetUserGroupsResponse, *http.Response, error)

	/*
			GetUserPresence Get a user's presence

			Get the presence status for a specific user.

		This endpoint is most useful for embedding data about a user's
		presence status in other sites (e.g. an employee directory). Full
		Zulip clients like mobile/desktop apps will want to use the [main
		presence endpoint](zulip.com/api/get-presence, which returns data for all
		active users in the organization, instead.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userIdOrEmail The Id or Zulip API email address of the user whose presence you want to fetch.  **Changes**: New in Zulip 4.0 (feature level 43). Previous versions only supported identifying the user by Zulip API email.
			@return GetUserPresenceRequest
	*/
	GetUserPresence(ctx context.Context, userIdOrEmail string) GetUserPresenceRequest

	// GetUserPresenceExecute executes the request
	//  @return GetUserPresenceResponse
	GetUserPresenceExecute(r GetUserPresenceRequest) (*GetUserPresenceResponse, *http.Response, error)

	/*
			GetUserStatus Get a user's status

			Get the [status](zulip.com/help/status-and-availability) currently set by a
		user in the organization.

		**Changes**: New in Zulip 9.0 (feature level 262). Previously,
		user statuses could only be fetched via the [`POST
		/register`](zulip.com/api/register-queue) endpoint.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId The target user's Id.
			@return GetUserStatusRequest
	*/
	GetUserStatus(ctx context.Context, userId int64) GetUserStatusRequest

	// GetUserStatusExecute executes the request
	//  @return GetUserStatusResponse
	GetUserStatusExecute(r GetUserStatusRequest) (*GetUserStatusResponse, *http.Response, error)

	/*
			GetUsers Get users

			Retrieve details on users in the organization.

		By default, returns all accessible users in the organization.
		The `user_ids` query parameter can be used to limit the
		results to a specific set of user Ids.

		Optionally includes values of [custom profile fields](zulip.com/help/custom-profile-fields.

		You can also [fetch details on a single user](zulip.com/api/get-user.

		**Changes**: This endpoint did not support unauthenticated
		access in organizations using the [public access
		option](zulip.com/help/public-access-option) prior to Zulip 11.0
		(feature level 387).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GetUsersRequest
	*/
	GetUsers(ctx context.Context) GetUsersRequest

	// GetUsersExecute executes the request
	//  @return GetUsersResponse
	GetUsersExecute(r GetUsersRequest) (*GetUsersResponse, *http.Response, error)

	/*
			MuteUser Mute a user

			[Mute a user](zulip.com/help/mute-a-user) from the perspective of the requesting
		user. Messages sent by muted users will be automatically marked as read
		and hidden for the user who muted them.

		Muted users should be implemented by clients as follows:

		- The server will immediately mark all messages sent by the muted
		  user as read. This will automatically clear any existing mobile
		  push notifications related to the muted user.
		- The server will mark any new messages sent by the muted user as read
		  for the requesting user's account, which prevents all email and mobile
		  push notifications.
		- Clients should exclude muted users from presence lists or other UI
		  for viewing or composing one-on-one direct messages. One-on-one direct
		  messages sent by muted users should be hidden everywhere in the Zulip UI.
		- Channel messages and group direct messages sent by the muted
		  user should avoid displaying the content and name/avatar,
		  but should display that N messages by a muted user were
		  hidden (so that it is possible to interpret the messages by
		  other users who are talking with the muted user).
		- Group direct message conversations including the muted user
		  should display muted users as "Muted user", rather than
		  showing their name, in lists of such conversations, along with using
		  a blank grey avatar where avatars are displayed.
		- Administrative/settings UI elements for showing "All users that exist
		  on this channel or realm", e.g. for organization
		  administration or showing channel subscribers, should display
		  the user's name as normal.

		**Changes**: New in Zulip 4.0 (feature level 48).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param mutedUserId The Id of the user to mute/unmute.  **Changes**: Before Zulip 8.0 (feature level 188), bot users could not be muted/unmuted, and specifying a bot user's Id returned an error response.
			@return MuteUserRequest
	*/
	MuteUser(ctx context.Context, mutedUserId int64) MuteUserRequest

	// MuteUserExecute executes the request
	//  @return Response
	MuteUserExecute(r MuteUserRequest) (*Response, *http.Response, error)

	/*
			ReactivateUser Reactivate a user

			[Reactivates a
		user](https://zulip.com/help/deactivate-or-reactivate-a-user)
		given their user Id.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId The target user's Id.
			@return ReactivateUserRequest
	*/
	ReactivateUser(ctx context.Context, userId int64) ReactivateUserRequest

	// ReactivateUserExecute executes the request
	//  @return Response
	ReactivateUserExecute(r ReactivateUserRequest) (*Response, *http.Response, error)

	/*
			RemoveAlertWords Remove alert words

			Remove words (or phrases) from the user's set of configured [alert words][alert-words].

		Alert words are case insensitive.

		[alert-words]: /help/dm-mention-alert-notifications#alert-words


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return RemoveAlertWordsRequest
	*/
	RemoveAlertWords(ctx context.Context) RemoveAlertWordsRequest

	// RemoveAlertWordsExecute executes the request
	//  @return AlertWordsResponse
	RemoveAlertWordsExecute(r RemoveAlertWordsRequest) (*AlertWordsResponse, *http.Response, error)

	/*
		RemoveApnsToken Remove an APNs device token

		This endpoint removes an APNs device token for iOS push notifications.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RemoveApnsTokenRequest
	*/
	RemoveApnsToken(ctx context.Context) RemoveApnsTokenRequest

	// RemoveApnsTokenExecute executes the request
	//  @return Response
	RemoveApnsTokenExecute(r RemoveApnsTokenRequest) (*Response, *http.Response, error)

	/*
			RemoveAttachment Delete an attachment

			Delete an uploaded file given its attachment Id.

		Note that uploaded files that have been referenced in at least
		one message are automatically deleted once the last message
		containing a link to them is deleted (whether directly or via
		a [message retention policy](zulip.com/help/message-retention-policy).

		Uploaded files that are never used in a message are
		automatically deleted a few weeks after being uploaded.

		Attachment Ids can be contained from [GET /attachments](zulip.com/api/get-attachments.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param attachmentId The Id of the attachment to be deleted.
			@return RemoveAttachmentRequest
	*/
	RemoveAttachment(ctx context.Context, attachmentId int64) RemoveAttachmentRequest

	// RemoveAttachmentExecute executes the request
	//  @return Response
	RemoveAttachmentExecute(r RemoveAttachmentRequest) (*Response, *http.Response, error)

	/*
		RemoveFcmToken Remove an FCM registration token

		This endpoint removes an FCM registration token for push notifications.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RemoveFcmTokenRequest
	*/
	RemoveFcmToken(ctx context.Context) RemoveFcmTokenRequest

	// RemoveFcmTokenExecute executes the request
	//  @return Response
	RemoveFcmTokenExecute(r RemoveFcmTokenRequest) (*Response, *http.Response, error)

	/*
			SetTypingStatus Set \"typing\" status

			Notify other users whether the current user is
		[typing a message][help-typing].

		Clients implementing Zulip's typing notifications
		protocol should work as follows:

		- Send a request to this endpoint with `"op": "start"` when a user
		  starts composing a message.
		- While the user continues to actively type or otherwise interact with
		  the compose UI (e.g. interacting with the compose box emoji picker),
		  send regular `"op": "start"` requests to this endpoint, using
		  `server_typing_started_wait_period_milliseconds` in the
		  [`POST /register`][api-register] response as the time interval
		  between each request.
		- Send a request to this endpoint with `"op": "stop"` when a user
		  has stopped using the compose UI for the time period indicated by
		  `server_typing_stopped_wait_period_milliseconds` in the
		  [`POST /register`][api-register] response or when a user
		  cancels the compose action (if it had previously sent a "start"
		  notification for that compose action).
		- Start displaying a visual typing indicator for a given conversation
		  when a [`typing op:start`][start-typing] event is received
		  from the server.
		- Continue displaying a visual typing indicator for the conversation
		  until a [`typing op:stop`][stop-typing] event is received
		  from the server or the time period indicated by
		  `server_typing_started_expiry_period_milliseconds` in the
		  [`POST /register`][api-register] response has passed without
		  a new `typing "op": "start"` event for the conversation.

		This protocol is designed to allow the server-side typing notifications
		implementation to be stateless while being resilient as network failures
		will not result in a user being incorrectly displayed as perpetually
		typing.

		See the subsystems documentation on [typing indicators][typing-protocol-docs]
		for additional design details on Zulip's typing notifications protocol.

		**Changes**: Clients shouldn't care about the APIs prior to Zulip 8.0 (feature level 215)
		for channel typing notifications, as no client actually implemented
		the previous API for those.

		Support for displaying channel typing notifications was new
		in Zulip 4.0 (feature level 58). Clients should indicate they support
		processing channel typing notifications via the `stream_typing_notifications`
		value in the `client_capabilities` parameter of the
		[`POST /register`][client-capabilities] endpoint.

		[help-typing]: /help/typing-notifications
		[api-register]: /api/register-queue
		[start-typing]: /api/get-events#typing-start
		[stop-typing]: /api/get-events#typing-stop
		[client-capabilities]: /api/register-queue#parameter-client_capabilities
		[typing-protocol-docs]: https://zulip.readthedocs.io/en/latest/subsystems/typing-indicators.html


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return SetTypingStatusRequest
	*/
	SetTypingStatus(ctx context.Context) SetTypingStatusRequest

	// SetTypingStatusExecute executes the request
	//  @return Response
	SetTypingStatusExecute(r SetTypingStatusRequest) (*Response, *http.Response, error)

	/*
			SetTypingStatusForMessageEdit Set \"typing\" status for message editing

			Notify other users whether the current user is editing a message.

		Typing notifications for editing messages follow the same protocol as
		[set-typing-status](zulip.com/api/set-typing-status, see that endpoint for
		details.

		**Changes**: Before Zulip 10.0 (feature level 361), the endpoint was
		named `/message_edit_typing` with `message_id` a required parameter in
		the request body. Clients are recommended to start using sending these
		typing notifications starting from this feature level.

		New in Zulip 10.0 (feature level 351). Previously, typing notifications were
		not available when editing messages.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param messageId The target message's Id.
			@return SetTypingStatusForMessageEditRequest
	*/
	SetTypingStatusForMessageEdit(ctx context.Context, messageId int64) SetTypingStatusForMessageEditRequest

	// SetTypingStatusForMessageEditExecute executes the request
	//  @return Response
	SetTypingStatusForMessageEditExecute(r SetTypingStatusForMessageEditRequest) (*Response, *http.Response, error)

	/*
			UnmuteUser Unmute a user

			[Unmute a user](zulip.com/help/mute-a-user#see-your-list-of-muted-users
		from the perspective of the requesting user.

		**Changes**: New in Zulip 4.0 (feature level 48).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param mutedUserId The Id of the user to mute/unmute.  **Changes**: Before Zulip 8.0 (feature level 188), bot users could not be muted/unmuted, and specifying a bot user's Id returned an error response.
			@return UnmuteUserRequest
	*/
	UnmuteUser(ctx context.Context, mutedUserId int64) UnmuteUserRequest

	// UnmuteUserExecute executes the request
	//  @return Response
	UnmuteUserExecute(r UnmuteUserRequest) (*Response, *http.Response, error)

	/*
			UpdatePresence Update your presence

			Update the current user's [presence][availability] and fetch presence data
		of other users in the organization.

		This endpoint is meant to be used by clients for both:

		- Reporting the current user's presence status (`"active"` or `"idle"`)
		  to the server.

		- Obtaining the presence data of all other users in the organization via
		  regular polling.

		Accurate user presence is one of the most expensive parts of any
		chat application (in terms of bandwidth and other resources). Therefore,
		it is important that clients implementing Zulip's user presence system
		use the modern [`last_update_id`](#parameter-last_update_id) protocol to
		minimize fetching duplicate user presence data.

		Client apps implementing presence are recommended to also consume [`presence`
		events](zulip.com/api/get-events#presence), in order to learn about newly online users
		immediately.

		The Zulip server is responsible for implementing [invisible mode][invisible],
		which disables sharing a user's presence data. Nonetheless, clients
		should check the `presence_enabled` field in user objects in order to
		display the current user as online or offline based on whether they are
		sharing their presence information.

		**Changes**: As of Zulip 8.0 (feature level 228), if the
		`CAN_ACCESS_ALL_USERS_GROUP_LIMITS_PRESENCE` server-level setting is
		`true`, then user presence data in the response is [limited to users
		the current user can see/access][limit-visibility].

		[limit-visibility]: /help/guest-users#configure-whether-guests-can-see-all-other-users
		[invisible]: /help/status-and-availability#invisible-mode
		[availability]: /help/status-and-availability#availability


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return UpdatePresenceRequest
	*/
	UpdatePresence(ctx context.Context) UpdatePresenceRequest

	// UpdatePresenceExecute executes the request
	//  @return UpdatePresenceResponse
	UpdatePresenceExecute(r UpdatePresenceRequest) (*UpdatePresenceResponse, *http.Response, error)

	/*
			UpdateSettings Update settings

			This endpoint is used to edit the current user's settings.

		**Changes**: Removed `dense_mode` setting in Zulip 10.0
		(feature level 364) as we now have `web_font_size_px` and
		`web_line_height_percent` settings for more control.

		Prior to Zulip 5.0 (feature level 80), this endpoint only
		supported the `full_name`, `email`, `old_password`, and
		`new_password` parameters. Notification settings were
		managed by `PATCH /settings/notifications`, and all other
		settings by `PATCH /settings/display`.

		The feature level 80 migration to merge these endpoints did not
		change how request parameters are encoded. However, it did change
		the handling of any invalid parameters present in a request
		(see feature level 78 change below).

		As of feature level 80, the `PATCH /settings/display` and
		`PATCH /settings/notifications` endpoints are deprecated aliases
		for this endpoint for backwards-compatibility, and will be removed
		once clients have migrated to use this endpoint.

		Prior to Zulip 5.0 (feature level 78), this endpoint indicated
		which parameters it had processed by including in the response
		object `"key": value` entries for values successfully changed by
		the request. That was replaced by the more ergonomic
		[`ignored_parameters_unsupported`][ignored-parameters] array.

		The `PATCH /settings/notifications` and `PATCH /settings/display`
		endpoints also had this behavior of indicating processed parameters
		before they became aliases of this endpoint in Zulip 5.0 (see
		feature level 80 change above).

		Before feature level 78, request parameters that were not supported
		(or were unchanged) were silently ignored.

		[ignored-parameters]: /api/rest-error-handling#ignored-parameters


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return UpdateSettingsRequest
	*/
	UpdateSettings(ctx context.Context) UpdateSettingsRequest

	// UpdateSettingsExecute executes the request
	//  @return Response
	UpdateSettingsExecute(r UpdateSettingsRequest) (*Response, *http.Response, error)

	/*
			UpdateStatus Update your status

			Change your [status](zulip.com/help/status-and-availability.

		A request to this endpoint will only change the parameters passed.
		For example, passing just `status_text` requests a change in the status
		text, but will leave the status emoji unchanged.

		Clients that wish to set the user's status to a specific value should
		pass all supported parameters.

		**Changes**: In Zulip 5.0 (feature level 86), added support for
		`emoji_name`, `emoji_code`, and `reaction_type` parameters.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return UpdateStatusRequest
	*/
	UpdateStatus(ctx context.Context) UpdateStatusRequest

	// UpdateStatusExecute executes the request
	//  @return Response
	UpdateStatusExecute(r UpdateStatusRequest) (*Response, *http.Response, error)

	/*
			UpdateStatusForUser Update user status

			Administrator endpoint for changing the [status](zulip.com/help/status-and-availability) of
		another user.

		**Changes**: New in Zulip 11.0 (feature level 407).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId The target user's Id.
			@return UpdateStatusForUserRequest
	*/
	UpdateStatusForUser(ctx context.Context, userId int64) UpdateStatusForUserRequest

	// UpdateStatusForUserExecute executes the request
	//  @return Response
	UpdateStatusForUserExecute(r UpdateStatusForUserRequest) (*Response, *http.Response, error)

	/*
			UpdateUser Update a user

			Administrative endpoint to update the details of another user in the organization.

		Supports everything an administrator can do to edit details of another
		user's account, including editing full name,
		[role](zulip.com/help/user-roles, and [custom profile
		fields](zulip.com/help/custom-profile-fields.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId The target user's Id.
			@return UpdateUserRequest
	*/
	UpdateUser(ctx context.Context, userId int64) UpdateUserRequest

	// UpdateUserExecute executes the request
	//  @return Response
	UpdateUserExecute(r UpdateUserRequest) (*Response, *http.Response, error)

	/*
			UpdateUserByEmail Update a user by email

			Administrative endpoint to update the details of another user in the organization by their email address.
		Works the same way as [`PATCH /users/{user_id}`](zulip.com/api/update-user) but fetching the target user by their
		real email address.

		The requester needs to have permission to view the target user's real email address, subject to the
		user's email address visibility setting. Otherwise, the dummy address of the format
		`user{id}@{realm.host}` needs be used. This follows the same rules as `GET /users/{email}`.

		**Changes**: New in Zulip 10.0 (feature level 313).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param email The email address of the user, specified following the same rules as [`GET /users/{email}`](zulip.com/api/get-user-by-email.
			@return UpdateUserByEmailRequest
	*/
	UpdateUserByEmail(ctx context.Context, email string) UpdateUserByEmailRequest

	// UpdateUserByEmailExecute executes the request
	//  @return Response
	UpdateUserByEmailExecute(r UpdateUserByEmailRequest) (*Response, *http.Response, error)

	/*
			UpdateUserGroup Update a user group

			Update the name, description or any of the permission settings
		of a [user group](zulip.com/help/user-groups.

		This endpoint is also used to reactivate a user group.

		Note that while permissions settings of deactivated groups can
		be edited by this API endpoint, and those permissions settings
		do affect the ability to modify the deactivated group and its
		membership, the deactivated group itself cannot be mentioned
		or used in the value of any permission without first being reactivated.

		**Changes**: Starting with Zulip 11.0 (feature level 386), this
		endpoint can be used to reactivate a user group.

		Prior to Zulip 10.0 (feature level 340), only the name field
		of deactivated groups could be modified.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userGroupId The Id of the target user group.
			@return UpdateUserGroupRequest
	*/
	UpdateUserGroup(ctx context.Context, userGroupId int64) UpdateUserGroupRequest

	// UpdateUserGroupExecute executes the request
	//  @return Response
	UpdateUserGroupExecute(r UpdateUserGroupRequest) (*Response, *http.Response, error)

	/*
			UpdateUserGroupMembers Update user group members

			Update the members of a [user group](zulip.com/help/user-groups. The
		user Ids must correspond to non-deactivated users.

		**Changes**: Prior to Zulip 11.0 (feature level 391), members
		could not be added or removed from a deactivated group.

		**Changes**: Prior to Zulip 10.0 (feature level 303), group memberships of
		deactivated users were visible to the API and could be edited via this endpoint.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userGroupId The Id of the target user group.
			@return UpdateUserGroupMembersRequest
	*/
	UpdateUserGroupMembers(ctx context.Context, userGroupId int64) UpdateUserGroupMembersRequest

	// UpdateUserGroupMembersExecute executes the request
	//  @return Response
	UpdateUserGroupMembersExecute(r UpdateUserGroupMembersRequest) (*Response, *http.Response, error)

	/*
			UpdateUserGroupSubgroups Update subgroups of a user group

			Update the subgroups of a [user group](zulip.com/help/user-groups.

		**Changes**: Prior to Zulip 11.0 (feature level 391), subgroups
		could not be added or removed from a deactivated group.

		**Changes**: New in Zulip 6.0 (feature level 127).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userGroupId The Id of the target user group.
			@return UpdateUserGroupSubgroupsRequest
	*/
	UpdateUserGroupSubgroups(ctx context.Context, userGroupId int64) UpdateUserGroupSubgroupsRequest

	// UpdateUserGroupSubgroupsExecute executes the request
	//  @return Response
	UpdateUserGroupSubgroupsExecute(r UpdateUserGroupSubgroupsRequest) (*Response, *http.Response, error)
}

type AddAlertWordsRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	alertWords *[]string
}

// An array of strings to be added to the user&#39;s set of configured alert words. Strings already present in the user&#39;s set of alert words already are ignored.  Alert words are case insensitive.
func (r AddAlertWordsRequest) AlertWords(alertWords []string) AddAlertWordsRequest {
	r.alertWords = &alertWords
	return r
}

func (r AddAlertWordsRequest) Execute() (*AlertWordsResponse, *http.Response, error) {
	return r.ApiService.AddAlertWordsExecute(r)
}

/*
AddAlertWords Add alert words

Add words (or phrases) to the user's set of configured [alert words][alert-words].

[alert-words]: /help/dm-mention-alert-notifications#alert-words

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AddAlertWordsRequest
*/
func (c *simpleClient) AddAlertWords(ctx context.Context) AddAlertWordsRequest {
	return AddAlertWordsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddAlertWordsResponse
func (c *simpleClient) AddAlertWordsExecute(r AddAlertWordsRequest) (*AlertWordsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlertWordsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/alert_words"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.alertWords == nil {
		return localVarReturnValue, nil, reportError("alertWords is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	paramJson, err := parameterToJson(r.alertWords)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarFormParams.Add("alert_words", paramJson)
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type AddApnsTokenRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	token      *string
	appid      *string
}

// The token provided by the device.
func (r AddApnsTokenRequest) Token(token string) AddApnsTokenRequest {
	r.token = &token
	return r
}

// The Id of the Zulip app that is making the request.  **Changes**: In Zulip 8.0 (feature level 223), this parameter was made required. Previously, if it was unspecified, the server would use a default value (based on the &#x60;ZULIP_IOS_APP_Id&#x60; server setting, which defaulted to &#x60;\\\&quot;org.zulip.Zulip\\\&quot;&#x60;).
func (r AddApnsTokenRequest) Appid(appid string) AddApnsTokenRequest {
	r.appid = &appid
	return r
}

func (r AddApnsTokenRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.AddApnsTokenExecute(r)
}

/*
AddApnsToken Add an APNs device token

This endpoint adds an APNs device token to register for iOS push notifications.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AddApnsTokenRequest
*/
func (c *simpleClient) AddApnsToken(ctx context.Context) AddApnsTokenRequest {
	return AddApnsTokenRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) AddApnsTokenExecute(r AddApnsTokenRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/apns_device_token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}
	if r.appid == nil {
		return localVarReturnValue, nil, reportError("appid is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "token", r.token, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "appid", r.appid, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type AddFcmTokenRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	token      *string
}

// The token provided by the device.
func (r AddFcmTokenRequest) Token(token string) AddFcmTokenRequest {
	r.token = &token
	return r
}

func (r AddFcmTokenRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.AddFcmTokenExecute(r)
}

/*
AddFcmToken Add an FCM registration token

This endpoint adds an FCM registration token for push notifications.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AddFcmTokenRequest
*/
func (c *simpleClient) AddFcmToken(ctx context.Context) AddFcmTokenRequest {
	return AddFcmTokenRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) AddFcmTokenExecute(r AddFcmTokenRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/android_gcm_reg_id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "token", r.token, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateUserRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	email      *string
	password   *string
	fullName   *string
}

// The email address of the new user.
func (r CreateUserRequest) Email(email string) CreateUserRequest {
	r.email = &email
	return r
}

// The password of the new user.
func (r CreateUserRequest) Password(password string) CreateUserRequest {
	r.password = &password
	return r
}

// The full name of the new user.
func (r CreateUserRequest) FullName(fullName string) CreateUserRequest {
	r.fullName = &fullName
	return r
}

func (r CreateUserRequest) Execute() (*CreateUserResponse, *http.Response, error) {
	return r.ApiService.CreateUserExecute(r)
}

/*
CreateUser Create a user

Create a new user account via the API.

!!! warn ""

	**Note**: On Zulip Cloud, this feature is available only for
	organizations on a [Zulip Cloud Standard](https://zulip.com/plans/)
	or [Zulip Cloud Plus](https://zulip.com/plans/) plan. Administrators
	can request the required `can_create_users` permission for a bot or
	user by contacting [Zulip Cloud support][support] with an
	explanation for why it is needed. Self-hosted installations can
	toggle `can_create_users` on an account using the `manage.py
	change_user_role` [management command][management-commands].

**Changes**: Before Zulip 4.0 (feature level 36), this endpoint was
available to all organization administrators.

[support]: /help/contact-support
[management-commands]: https://zulip.readthedocs.io/en/latest/production/management-commands.html

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateUserRequest
*/
func (c *simpleClient) CreateUser(ctx context.Context) CreateUserRequest {
	return CreateUserRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateUserResponse
func (c *simpleClient) CreateUserExecute(r CreateUserRequest) (*CreateUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateUserResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
	}
	if r.password == nil {
		return localVarReturnValue, nil, reportError("password is required and must be specified")
	}
	if r.fullName == nil {
		return localVarReturnValue, nil, reportError("fullName is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "email", r.email, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "full_name", r.fullName, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateUserGroupRequest struct {
	ctx                   context.Context
	ApiService            UsersAPI
	name                  *string
	description           *string
	members               *[]int64
	subgroups             *[]int64
	canAddMembersGroup    *GroupSettingValue
	canJoinGroup          *GroupSettingValue
	canLeaveGroup         *GroupSettingValue
	canManageGroup        *GroupSettingValue
	canMentionGroup       *GroupSettingValue
	canRemoveMembersGroup *GroupSettingValue
}

// The name of the user group.
func (r CreateUserGroupRequest) Name(name string) CreateUserGroupRequest {
	r.name = &name
	return r
}

// The description of the user group.
func (r CreateUserGroupRequest) Description(description string) CreateUserGroupRequest {
	r.description = &description
	return r
}

// An array containing the user Ids of the initial members for the new user group.
func (r CreateUserGroupRequest) Members(members []int64) CreateUserGroupRequest {
	r.members = &members
	return r
}

// An array containing the Ids of the initial subgroups for the new user group.  User can add subgroups to the new group irrespective of other permissions for the new group.  **Changes**: New in Zulip 10.0 (feature level 311).
func (r CreateUserGroupRequest) Subgroups(subgroups []int64) CreateUserGroupRequest {
	r.subgroups = &subgroups
	return r
}

// A [group-setting value][setting-values] defining the set of users who have permission to add members to this user group.  **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups
func (r CreateUserGroupRequest) CanAddMembersGroup(canAddMembersGroup GroupSettingValue) CreateUserGroupRequest {
	r.canAddMembersGroup = &canAddMembersGroup
	return r
}

// A [group-setting value][setting-values] defining the set of users who have permission to join this user group.  **Changes**: New in Zulip 10.0 (feature level 301).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups
func (r CreateUserGroupRequest) CanJoinGroup(canJoinGroup GroupSettingValue) CreateUserGroupRequest {
	r.canJoinGroup = &canJoinGroup
	return r
}

// A [group-setting value][setting-values] defining the set of users who have permission to leave this user group.  **Changes**: New in Zulip 10.0 (feature level 308).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups
func (r CreateUserGroupRequest) CanLeaveGroup(canLeaveGroup GroupSettingValue) CreateUserGroupRequest {
	r.canLeaveGroup = &canLeaveGroup
	return r
}

// A [group-setting value][setting-values] defining the set of users who have permission to [manage this user group][manage-user-groups].  This setting cannot be set to &#x60;\\\&quot;role:internet\\\&quot;&#x60; and &#x60;\\\&quot;role:everyone\\\&quot;&#x60; [system groups][system-groups].  **Changes**: New in Zulip 10.0 (feature level 283).  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups [manage-user-groups]: /help/manage-user-groups
func (r CreateUserGroupRequest) CanManageGroup(canManageGroup GroupSettingValue) CreateUserGroupRequest {
	r.canManageGroup = &canManageGroup
	return r
}

// A [group-setting value][setting-values] defining the set of users who have permission to [mention this user group][mentions].  This setting cannot be set to &#x60;\\\&quot;role:internet\\\&quot;&#x60; and &#x60;\\\&quot;role:owners\\\&quot;&#x60; [system groups][system-groups].  Before Zulip 9.0 (feature level 258), this parameter could only be the integer form of a [group-setting value][setting-values].  Before Zulip 8.0 (feature level 198), this parameter was named &#x60;can_mention_group_id&#x60;.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups][system-groups].  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups [mentions]: /help/mention-a-user-or-group
func (r CreateUserGroupRequest) CanMentionGroup(canMentionGroup GroupSettingValue) CreateUserGroupRequest {
	r.canMentionGroup = &canMentionGroup
	return r
}

// A [group-setting value][setting-values] defining the set of users who have permission to remove members from this user group.  **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups
func (r CreateUserGroupRequest) CanRemoveMembersGroup(canRemoveMembersGroup GroupSettingValue) CreateUserGroupRequest {
	r.canRemoveMembersGroup = &canRemoveMembersGroup
	return r
}

func (r CreateUserGroupRequest) Execute() (*CreateUserGroupResponse, *http.Response, error) {
	return r.ApiService.CreateUserGroupExecute(r)
}

/*
CreateUserGroup Create a user group

Create a new [user group](zulip.com/help/user-groups.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateUserGroupRequest
*/
func (c *simpleClient) CreateUserGroup(ctx context.Context) CreateUserGroupRequest {
	return CreateUserGroupRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateUserGroupResponse
func (c *simpleClient) CreateUserGroupExecute(r CreateUserGroupRequest) (*CreateUserGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateUserGroupResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if r.description == nil {
		return localVarReturnValue, nil, reportError("description is required and must be specified")
	}
	if r.members == nil {
		return localVarReturnValue, nil, reportError("members is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "members", r.members, "form", "multi")
	if r.subgroups != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "subgroups", r.subgroups, "form", "multi")
	}
	if r.canAddMembersGroup != nil {
		paramJson, err := parameterToJson(*r.canAddMembersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_add_members_group", paramJson)
	}
	if r.canJoinGroup != nil {
		paramJson, err := parameterToJson(*r.canJoinGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_join_group", paramJson)
	}
	if r.canLeaveGroup != nil {
		paramJson, err := parameterToJson(*r.canLeaveGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_leave_group", paramJson)
	}
	if r.canManageGroup != nil {
		paramJson, err := parameterToJson(*r.canManageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_manage_group", paramJson)
	}
	if r.canMentionGroup != nil {
		paramJson, err := parameterToJson(*r.canMentionGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_mention_group", paramJson)
	}
	if r.canRemoveMembersGroup != nil {
		paramJson, err := parameterToJson(*r.canRemoveMembersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_remove_members_group", paramJson)
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeactivateOwnUserRequest struct {
	ctx        context.Context
	ApiService UsersAPI
}

func (r DeactivateOwnUserRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeactivateOwnUserExecute(r)
}

/*
DeactivateOwnUser Deactivate own user

Deactivates the current user's account. See also the administrative endpoint for
[deactivating another user](zulip.com/api/deactivate-user.

This endpoint is primarily useful to Zulip clients providing a user settings UI.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DeactivateOwnUserRequest
*/
func (c *simpleClient) DeactivateOwnUser(ctx context.Context) DeactivateOwnUserRequest {
	return DeactivateOwnUserRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) DeactivateOwnUserExecute(r DeactivateOwnUserRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeactivateUserRequest struct {
	ctx                             context.Context
	ApiService                      UsersAPI
	userId                          int64
	deactivationNotificationComment *string
}

// If not &#x60;null&#x60;, requests that the deactivated user receive a notification email about their account deactivation.  If not &#x60;\\\&quot;\\\&quot;&#x60;, encodes custom text written by the administrator to be included in the notification email.  **Changes**: New in Zulip 5.0 (feature level 135).
func (r DeactivateUserRequest) DeactivationNotificationComment(deactivationNotificationComment string) DeactivateUserRequest {
	r.deactivationNotificationComment = &deactivationNotificationComment
	return r
}

func (r DeactivateUserRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeactivateUserExecute(r)
}

/*
DeactivateUser Deactivate a user

[Deactivates a
user](https://zulip.com/help/deactivate-or-reactivate-a-user)
given their user Id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The target user's Id.
	@return DeactivateUserRequest
*/
func (c *simpleClient) DeactivateUser(ctx context.Context, userId int64) DeactivateUserRequest {
	return DeactivateUserRequest{
		ApiService: c,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) DeactivateUserExecute(r DeactivateUserRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.deactivationNotificationComment != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "deactivation_notification_comment", r.deactivationNotificationComment, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeactivateUserGroupRequest struct {
	ctx         context.Context
	ApiService  UsersAPI
	userGroupId int64
}

func (r DeactivateUserGroupRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeactivateUserGroupExecute(r)
}

/*
DeactivateUserGroup Deactivate a user group

Deactivate a user group. Deactivated user groups cannot be
used for mentions, permissions, or any other purpose, but can
be reactivated or renamed.

Deactivating user groups is preferable to deleting them from
the database, since the deactivation model allows audit logs
of changes to sensitive group-valued permissions to be
maintained.

**Changes**: New in Zulip 10.0 (feature level 290).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userGroupId The Id of the target user group.
	@return DeactivateUserGroupRequest
*/
func (c *simpleClient) DeactivateUserGroup(ctx context.Context, userGroupId int64) DeactivateUserGroupRequest {
	return DeactivateUserGroupRequest{
		ApiService:  c,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) DeactivateUserGroupExecute(r DeactivateUserGroupRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups/{user_group_id}/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"user_group_id"+"}", url.PathEscape(parameterValueToString(r.userGroupId, "userGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetAlertWordsRequest struct {
	ctx        context.Context
	ApiService UsersAPI
}

func (r GetAlertWordsRequest) Execute() (*AlertWordsResponse, *http.Response, error) {
	return r.ApiService.GetAlertWordsExecute(r)
}

/*
GetAlertWords Get all alert words

Get all of the user's configured [alert words][alert-words].

[alert-words]: /help/dm-mention-alert-notifications#alert-words

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetAlertWordsRequest
*/
func (c *simpleClient) GetAlertWords(ctx context.Context) GetAlertWordsRequest {
	return GetAlertWordsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAlertWordsResponse
func (c *simpleClient) GetAlertWordsExecute(r GetAlertWordsRequest) (*AlertWordsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlertWordsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/alert_words"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetAttachmentsRequest struct {
	ctx        context.Context
	ApiService UsersAPI
}

func (r GetAttachmentsRequest) Execute() (*GetAttachmentsResponse, *http.Response, error) {
	return r.ApiService.GetAttachmentsExecute(r)
}

/*
GetAttachments Get attachments

Fetch metadata on files uploaded by the requesting user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetAttachmentsRequest
*/
func (c *simpleClient) GetAttachments(ctx context.Context) GetAttachmentsRequest {
	return GetAttachmentsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAttachmentsResponse
func (c *simpleClient) GetAttachmentsExecute(r GetAttachmentsRequest) (*GetAttachmentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAttachmentsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/attachments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetIsUserGroupMemberRequest struct {
	ctx              context.Context
	ApiService       UsersAPI
	userGroupId      int64
	userId           int64
	directMemberOnly *bool
}

// Whether to consider only the direct members of user group and not members of its subgroups. Default is &#x60;false&#x60;.
func (r GetIsUserGroupMemberRequest) DirectMemberOnly(directMemberOnly bool) GetIsUserGroupMemberRequest {
	r.directMemberOnly = &directMemberOnly
	return r
}

func (r GetIsUserGroupMemberRequest) Execute() (*GetIsUserGroupMemberResponse, *http.Response, error) {
	return r.ApiService.GetIsUserGroupMemberExecute(r)
}

/*
GetIsUserGroupMember Get user group membership status

Check whether a user is member of user group.

**Changes**: Prior to Zulip 10.0 (feature level 303),
this would return true when passed a deactivated user
who was a member of the user group before being deactivated.

New in Zulip 6.0 (feature level 127).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userGroupId The Id of the target user group.
	@param userId The target user's Id.
	@return GetIsUserGroupMemberRequest
*/
func (c *simpleClient) GetIsUserGroupMember(ctx context.Context, userGroupId int64, userId int64) GetIsUserGroupMemberRequest {
	return GetIsUserGroupMemberRequest{
		ApiService:  c,
		ctx:         ctx,
		userGroupId: userGroupId,
		userId:      userId,
	}
}

// Execute executes the request
//
//	@return GetIsUserGroupMemberResponse
func (c *simpleClient) GetIsUserGroupMemberExecute(r GetIsUserGroupMemberRequest) (*GetIsUserGroupMemberResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetIsUserGroupMemberResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups/{user_group_id}/members/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_group_id"+"}", url.PathEscape(parameterValueToString(r.userGroupId, "userGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.directMemberOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direct_member_only", r.directMemberOnly, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetOwnUserRequest struct {
	ctx        context.Context
	ApiService UsersAPI
}

func (r GetOwnUserRequest) Execute() (*GetOwnUserResponse, *http.Response, error) {
	return r.ApiService.GetOwnUserExecute(r)
}

/*
GetOwnUser Get own user

Get basic data about the user/bot that requests this endpoint.

**Changes**: Removed `is_billing_admin` field in Zulip 10.0 (feature level 363), as it was
replaced by the `can_manage_billing_group` realm setting.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetOwnUserRequest
*/
func (c *simpleClient) GetOwnUser(ctx context.Context) GetOwnUserRequest {
	return GetOwnUserRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOwnUserResponse
func (c *simpleClient) GetOwnUserExecute(r GetOwnUserRequest) (*GetOwnUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOwnUserResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetUserRequest struct {
	ctx                        context.Context
	ApiService                 UsersAPI
	userId                     int64
	clientGravatar             *bool
	includeCustomProfileFields *bool
}

// Whether the client supports computing gravatars URLs. If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  **Changes**: The default value of this parameter was &#x60;false&#x60; prior to Zulip 5.0 (feature level 92).
func (r GetUserRequest) ClientGravatar(clientGravatar bool) GetUserRequest {
	r.clientGravatar = &clientGravatar
	return r
}

// Whether the client wants [custom profile field](zulip.com/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.
func (r GetUserRequest) IncludeCustomProfileFields(includeCustomProfileFields bool) GetUserRequest {
	r.includeCustomProfileFields = &includeCustomProfileFields
	return r
}

func (r GetUserRequest) Execute() (*GetUserResponse, *http.Response, error) {
	return r.ApiService.GetUserExecute(r)
}

/*
GetUser Get a user

Fetch details for a single user in the organization.

You can also fetch details on [all users in the organization](zulip.com/api/get-users
or [by a user's Zulip API email](zulip.com/api/get-user-by-email.

**Changes**: New in Zulip 3.0 (feature level 1).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The target user's Id.
	@return GetUserRequest
*/
func (c *simpleClient) GetUser(ctx context.Context, userId int64) GetUserRequest {
	return GetUserRequest{
		ApiService: c,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return GetUserResponse
func (c *simpleClient) GetUserExecute(r GetUserRequest) (*GetUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUserResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clientGravatar != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "client_gravatar", r.clientGravatar, "form", "")
	} else {
		var defaultValue bool = true
		r.clientGravatar = &defaultValue
	}
	if r.includeCustomProfileFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_custom_profile_fields", r.includeCustomProfileFields, "form", "")
	} else {
		var defaultValue bool = false
		r.includeCustomProfileFields = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetUserByEmailRequest struct {
	ctx                        context.Context
	ApiService                 UsersAPI
	email                      string
	clientGravatar             *bool
	includeCustomProfileFields *bool
}

// Whether the client supports computing gravatars URLs. If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  **Changes**: The default value of this parameter was &#x60;false&#x60; prior to Zulip 5.0 (feature level 92).
func (r GetUserByEmailRequest) ClientGravatar(clientGravatar bool) GetUserByEmailRequest {
	r.clientGravatar = &clientGravatar
	return r
}

// Whether the client wants [custom profile field](zulip.com/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.
func (r GetUserByEmailRequest) IncludeCustomProfileFields(includeCustomProfileFields bool) GetUserByEmailRequest {
	r.includeCustomProfileFields = &includeCustomProfileFields
	return r
}

func (r GetUserByEmailRequest) Execute() (*GetUserResponse, *http.Response, error) {
	return r.ApiService.GetUserByEmailExecute(r)
}

/*
GetUserByEmail Get a user by email

Fetch details for a single user in the organization given a Zulip
API email address.

You can also fetch details on [all users in the organization](zulip.com/api/get-users
or [by user Id](zulip.com/api/get-user.

Fetching by user Id is generally recommended when possible,
as a user might [change their email address](zulip.com/help/change-your-email-address
or change their [email address visibility](zulip.com/help/configure-email-visibility,
either of which could change the client's ability to look them up by that
email address.

**Changes**: Starting with Zulip 10.0 (feature level 302), the real email
address can be used in the `email` parameter and will fetch the target user's
data if and only if the target's email visibility setting permits the requester
to see the email address.
The dummy email addresses of the form `user{id}@{realm.host}` still work, and
will now work for **all** users, via identifying them by the embedded user Id.

New in Zulip Server 4.0 (feature level 39).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param email The email address of the user to fetch. Two forms are supported:  - The real email address of the user (`delivery_email`). The lookup will   succeed if and only if the user exists and their email address visibility   setting permits the client to see the email address.  - The dummy Zulip API email address of the form `user{user_id}@{realm_host}`. This   is identical to simply [getting user by Id](zulip.com/api/get-user. If the server or   realm change domains, the dummy email address used has to be adjustment to   match the new realm domain. This is legacy behavior for   backwards-compatibility, and will be removed in a future release.  **Changes**: Starting with Zulip 10.0 (feature level 302), lookups by real email address match the semantics of the target's email visibility setting and dummy email addresses work for all users, independently of their email visibility setting.  Previously, lookups were done only using the Zulip API email addresses.
	@return GetUserByEmailRequest
*/
func (c *simpleClient) GetUserByEmail(ctx context.Context, email string) GetUserByEmailRequest {
	return GetUserByEmailRequest{
		ApiService: c,
		ctx:        ctx,
		email:      email,
	}
}

// Execute executes the request
//
//	@return GetUserResponse
func (c *simpleClient) GetUserByEmailExecute(r GetUserByEmailRequest) (*GetUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUserResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clientGravatar != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "client_gravatar", r.clientGravatar, "form", "")
	} else {
		var defaultValue bool = true
		r.clientGravatar = &defaultValue
	}
	if r.includeCustomProfileFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_custom_profile_fields", r.includeCustomProfileFields, "form", "")
	} else {
		var defaultValue bool = false
		r.includeCustomProfileFields = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetUserGroupMembersRequest struct {
	ctx              context.Context
	ApiService       UsersAPI
	userGroupId      int64
	directMemberOnly *bool
}

// Whether to consider only the direct members of user group and not members of its subgroups. Default is &#x60;false&#x60;.
func (r GetUserGroupMembersRequest) DirectMemberOnly(directMemberOnly bool) GetUserGroupMembersRequest {
	r.directMemberOnly = &directMemberOnly
	return r
}

func (r GetUserGroupMembersRequest) Execute() (*GetUserGroupMembersResponse, *http.Response, error) {
	return r.ApiService.GetUserGroupMembersExecute(r)
}

/*
GetUserGroupMembers Get user group members

Get the members of a [user group](zulip.com/help/user-groups.

**Changes**: New in Zulip 6.0 (feature level 127).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userGroupId The Id of the target user group.
	@return GetUserGroupMembersRequest
*/
func (c *simpleClient) GetUserGroupMembers(ctx context.Context, userGroupId int64) GetUserGroupMembersRequest {
	return GetUserGroupMembersRequest{
		ApiService:  c,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
//
//	@return GetUserGroupMembersResponse
func (c *simpleClient) GetUserGroupMembersExecute(r GetUserGroupMembersRequest) (*GetUserGroupMembersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUserGroupMembersResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups/{user_group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"user_group_id"+"}", url.PathEscape(parameterValueToString(r.userGroupId, "userGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.directMemberOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direct_member_only", r.directMemberOnly, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetUserGroupSubgroupsRequest struct {
	ctx                context.Context
	ApiService         UsersAPI
	userGroupId        int64
	directSubgroupOnly *bool
}

// Whether to consider only direct subgroups of the user group or subgroups of subgroups also.
func (r GetUserGroupSubgroupsRequest) DirectSubgroupOnly(directSubgroupOnly bool) GetUserGroupSubgroupsRequest {
	r.directSubgroupOnly = &directSubgroupOnly
	return r
}

func (r GetUserGroupSubgroupsRequest) Execute() (*GetUserGroupSubgroupsResponse, *http.Response, error) {
	return r.ApiService.GetUserGroupSubgroupsExecute(r)
}

/*
GetUserGroupSubgroups Get subgroups of a user group

Get the subgroups of a [user group](zulip.com/help/user-groups.

**Changes**: New in Zulip 6.0 (feature level 127).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userGroupId The Id of the target user group.
	@return GetUserGroupSubgroupsRequest
*/
func (c *simpleClient) GetUserGroupSubgroups(ctx context.Context, userGroupId int64) GetUserGroupSubgroupsRequest {
	return GetUserGroupSubgroupsRequest{
		ApiService:  c,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
//
//	@return GetUserGroupSubgroupsResponse
func (c *simpleClient) GetUserGroupSubgroupsExecute(r GetUserGroupSubgroupsRequest) (*GetUserGroupSubgroupsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUserGroupSubgroupsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups/{user_group_id}/subgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"user_group_id"+"}", url.PathEscape(parameterValueToString(r.userGroupId, "userGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.directSubgroupOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direct_subgroup_only", r.directSubgroupOnly, "form", "")
	} else {
		var defaultValue bool = false
		r.directSubgroupOnly = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetUserGroupsRequest struct {
	ctx                      context.Context
	ApiService               UsersAPI
	includeDeactivatedGroups *bool
}

// Whether to include deactivated user groups in the response.  **Changes**: In Zulip 10.0 (feature level 294), renamed &#x60;allow_deactivated&#x60; to &#x60;include_deactivated_groups&#x60;.  New in Zulip 10.0 (feature level 290). Previously, deactivated user groups did not exist and thus would never be included in the response.
func (r GetUserGroupsRequest) IncludeDeactivatedGroups(includeDeactivatedGroups bool) GetUserGroupsRequest {
	r.includeDeactivatedGroups = &includeDeactivatedGroups
	return r
}

func (r GetUserGroupsRequest) Execute() (*GetUserGroupsResponse, *http.Response, error) {
	return r.ApiService.GetUserGroupsExecute(r)
}

/*
GetUserGroups Get user groups

Fetches all of the user groups in the organization.

!!! warn ""

	   **Note**: This endpoint is only available to [members and
	   administrators](zulip.com/help/user-roles; bots and guests
	   cannot use this endpoint.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetUserGroupsRequest
*/
func (c *simpleClient) GetUserGroups(ctx context.Context) GetUserGroupsRequest {
	return GetUserGroupsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUserGroupsResponse
func (c *simpleClient) GetUserGroupsExecute(r GetUserGroupsRequest) (*GetUserGroupsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUserGroupsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.includeDeactivatedGroups != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "include_deactivated_groups", r.includeDeactivatedGroups, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetUserPresenceRequest struct {
	ctx           context.Context
	ApiService    UsersAPI
	userIdOrEmail string
}

func (r GetUserPresenceRequest) Execute() (*GetUserPresenceResponse, *http.Response, error) {
	return r.ApiService.GetUserPresenceExecute(r)
}

/*
GetUserPresence Get a user's presence

Get the presence status for a specific user.

This endpoint is most useful for embedding data about a user's
presence status in other sites (e.g. an employee directory). Full
Zulip clients like mobile/desktop apps will want to use the [main
presence endpoint](zulip.com/api/get-presence, which returns data for all
active users in the organization, instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userIdOrEmail The Id or Zulip API email address of the user whose presence you want to fetch.  **Changes**: New in Zulip 4.0 (feature level 43). Previous versions only supported identifying the user by Zulip API email.
	@return GetUserPresenceRequest
*/
func (c *simpleClient) GetUserPresence(ctx context.Context, userIdOrEmail string) GetUserPresenceRequest {
	return GetUserPresenceRequest{
		ApiService:    c,
		ctx:           ctx,
		userIdOrEmail: userIdOrEmail,
	}
}

// Execute executes the request
//
//	@return GetUserPresenceResponse
func (c *simpleClient) GetUserPresenceExecute(r GetUserPresenceRequest) (*GetUserPresenceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUserPresenceResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id_or_email}/presence"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id_or_email"+"}", url.PathEscape(parameterValueToString(r.userIdOrEmail, "userIdOrEmail")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetUserStatusRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	userId     int64
}

func (r GetUserStatusRequest) Execute() (*GetUserStatusResponse, *http.Response, error) {
	return r.ApiService.GetUserStatusExecute(r)
}

/*
GetUserStatus Get a user's status

Get the [status](zulip.com/help/status-and-availability) currently set by a
user in the organization.

**Changes**: New in Zulip 9.0 (feature level 262). Previously,
user statuses could only be fetched via the [`POST
/register`](zulip.com/api/register-queue) endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The target user's Id.
	@return GetUserStatusRequest
*/
func (c *simpleClient) GetUserStatus(ctx context.Context, userId int64) GetUserStatusRequest {
	return GetUserStatusRequest{
		ApiService: c,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return GetUserStatusResponse
func (c *simpleClient) GetUserStatusExecute(r GetUserStatusRequest) (*GetUserStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUserStatusResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetUsersRequest struct {
	ctx                        context.Context
	ApiService                 UsersAPI
	clientGravatar             *bool
	includeCustomProfileFields *bool
	userIds                    *[]int64
}

// Whether the client supports computing gravatars URLs. If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  **Changes**: The default value of this parameter was &#x60;false&#x60; prior to Zulip 5.0 (feature level 92).
func (r GetUsersRequest) ClientGravatar(clientGravatar bool) GetUsersRequest {
	r.clientGravatar = &clientGravatar
	return r
}

// Whether the client wants [custom profile field](zulip.com/help/custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.
func (r GetUsersRequest) IncludeCustomProfileFields(includeCustomProfileFields bool) GetUsersRequest {
	r.includeCustomProfileFields = &includeCustomProfileFields
	return r
}

// Limits the results to the specified user Ids. If not provided, the server will return all accessible users in the organization.  **Changes**: New in Zulip 11.0 (feature level 384).
func (r GetUsersRequest) UserIds(userIds []int64) GetUsersRequest {
	r.userIds = &userIds
	return r
}

func (r GetUsersRequest) Execute() (*GetUsersResponse, *http.Response, error) {
	return r.ApiService.GetUsersExecute(r)
}

/*
GetUsers Get users

Retrieve details on users in the organization.

By default, returns all accessible users in the organization.
The `user_ids` query parameter can be used to limit the
results to a specific set of user Ids.

Optionally includes values of [custom profile fields](zulip.com/help/custom-profile-fields.

You can also [fetch details on a single user](zulip.com/api/get-user.

**Changes**: This endpoint did not support unauthenticated
access in organizations using the [public access
option](zulip.com/help/public-access-option) prior to Zulip 11.0
(feature level 387).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetUsersRequest
*/
func (c *simpleClient) GetUsers(ctx context.Context) GetUsersRequest {
	return GetUsersRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUsersResponse
func (c *simpleClient) GetUsersExecute(r GetUsersRequest) (*GetUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUsersResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clientGravatar != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "client_gravatar", r.clientGravatar, "form", "")
	} else {
		var defaultValue bool = true
		r.clientGravatar = &defaultValue
	}
	if r.includeCustomProfileFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_custom_profile_fields", r.includeCustomProfileFields, "form", "")
	} else {
		var defaultValue bool = false
		r.includeCustomProfileFields = &defaultValue
	}
	if r.userIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_ids", r.userIds, "", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MuteUserRequest struct {
	ctx         context.Context
	ApiService  UsersAPI
	mutedUserId int64
}

func (r MuteUserRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.MuteUserExecute(r)
}

/*
MuteUser Mute a user

[Mute a user](zulip.com/help/mute-a-user) from the perspective of the requesting
user. Messages sent by muted users will be automatically marked as read
and hidden for the user who muted them.

Muted users should be implemented by clients as follows:

  - The server will immediately mark all messages sent by the muted
    user as read. This will automatically clear any existing mobile
    push notifications related to the muted user.
  - The server will mark any new messages sent by the muted user as read
    for the requesting user's account, which prevents all email and mobile
    push notifications.
  - Clients should exclude muted users from presence lists or other UI
    for viewing or composing one-on-one direct messages. One-on-one direct
    messages sent by muted users should be hidden everywhere in the Zulip UI.
  - Channel messages and group direct messages sent by the muted
    user should avoid displaying the content and name/avatar,
    but should display that N messages by a muted user were
    hidden (so that it is possible to interpret the messages by
    other users who are talking with the muted user).
  - Group direct message conversations including the muted user
    should display muted users as "Muted user", rather than
    showing their name, in lists of such conversations, along with using
    a blank grey avatar where avatars are displayed.
  - Administrative/settings UI elements for showing "All users that exist
    on this channel or realm", e.g. for organization
    administration or showing channel subscribers, should display
    the user's name as normal.

**Changes**: New in Zulip 4.0 (feature level 48).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param mutedUserId The Id of the user to mute/unmute.  **Changes**: Before Zulip 8.0 (feature level 188), bot users could not be muted/unmuted, and specifying a bot user's Id returned an error response.
	@return MuteUserRequest
*/
func (c *simpleClient) MuteUser(ctx context.Context, mutedUserId int64) MuteUserRequest {
	return MuteUserRequest{
		ApiService:  c,
		ctx:         ctx,
		mutedUserId: mutedUserId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) MuteUserExecute(r MuteUserRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/muted_users/{muted_user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"muted_user_id"+"}", url.PathEscape(parameterValueToString(r.mutedUserId, "mutedUserId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ReactivateUserRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	userId     int64
}

func (r ReactivateUserRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.ReactivateUserExecute(r)
}

/*
ReactivateUser Reactivate a user

[Reactivates a
user](https://zulip.com/help/deactivate-or-reactivate-a-user)
given their user Id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The target user's Id.
	@return ReactivateUserRequest
*/
func (c *simpleClient) ReactivateUser(ctx context.Context, userId int64) ReactivateUserRequest {
	return ReactivateUserRequest{
		ApiService: c,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) ReactivateUserExecute(r ReactivateUserRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/reactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemoveAlertWordsRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	alertWords *[]string
}

// An array of strings to be removed from the user&#39;s set of configured alert words. Strings that are not in the user&#39;s set of alert words are ignored.
func (r RemoveAlertWordsRequest) AlertWords(alertWords []string) RemoveAlertWordsRequest {
	r.alertWords = &alertWords
	return r
}

func (r RemoveAlertWordsRequest) Execute() (*AlertWordsResponse, *http.Response, error) {
	return r.ApiService.RemoveAlertWordsExecute(r)
}

/*
RemoveAlertWords Remove alert words

Remove words (or phrases) from the user's set of configured [alert words][alert-words].

Alert words are case insensitive.

[alert-words]: /help/dm-mention-alert-notifications#alert-words

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RemoveAlertWordsRequest
*/
func (c *simpleClient) RemoveAlertWords(ctx context.Context) RemoveAlertWordsRequest {
	return RemoveAlertWordsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlertWordsResponse
func (c *simpleClient) RemoveAlertWordsExecute(r RemoveAlertWordsRequest) (*AlertWordsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlertWordsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/alert_words"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.alertWords == nil {
		return localVarReturnValue, nil, reportError("alertWords is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "alert_words", r.alertWords, "form", "multi")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemoveApnsTokenRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	token      *string
}

// The token provided by the device.
func (r RemoveApnsTokenRequest) Token(token string) RemoveApnsTokenRequest {
	r.token = &token
	return r
}

func (r RemoveApnsTokenRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveApnsTokenExecute(r)
}

/*
RemoveApnsToken Remove an APNs device token

This endpoint removes an APNs device token for iOS push notifications.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RemoveApnsTokenRequest
*/
func (c *simpleClient) RemoveApnsToken(ctx context.Context) RemoveApnsTokenRequest {
	return RemoveApnsTokenRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) RemoveApnsTokenExecute(r RemoveApnsTokenRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/apns_device_token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "token", r.token, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemoveAttachmentRequest struct {
	ctx          context.Context
	ApiService   UsersAPI
	attachmentId int64
}

func (r RemoveAttachmentRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveAttachmentExecute(r)
}

/*
RemoveAttachment Delete an attachment

Delete an uploaded file given its attachment Id.

Note that uploaded files that have been referenced in at least
one message are automatically deleted once the last message
containing a link to them is deleted (whether directly or via
a [message retention policy](zulip.com/help/message-retention-policy).

Uploaded files that are never used in a message are
automatically deleted a few weeks after being uploaded.

Attachment Ids can be contained from [GET /attachments](zulip.com/api/get-attachments.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param attachmentId The Id of the attachment to be deleted.
	@return RemoveAttachmentRequest
*/
func (c *simpleClient) RemoveAttachment(ctx context.Context, attachmentId int64) RemoveAttachmentRequest {
	return RemoveAttachmentRequest{
		ApiService:   c,
		ctx:          ctx,
		attachmentId: attachmentId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) RemoveAttachmentExecute(r RemoveAttachmentRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/attachments/{attachment_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"attachment_id"+"}", url.PathEscape(parameterValueToString(r.attachmentId, "attachmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemoveFcmTokenRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	token      *string
}

// The token provided by the device.
func (r RemoveFcmTokenRequest) Token(token string) RemoveFcmTokenRequest {
	r.token = &token
	return r
}

func (r RemoveFcmTokenRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveFcmTokenExecute(r)
}

/*
RemoveFcmToken Remove an FCM registration token

This endpoint removes an FCM registration token for push notifications.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RemoveFcmTokenRequest
*/
func (c *simpleClient) RemoveFcmToken(ctx context.Context) RemoveFcmTokenRequest {
	return RemoveFcmTokenRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) RemoveFcmTokenExecute(r RemoveFcmTokenRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/android_gcm_reg_id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "token", r.token, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SetTypingStatusRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	op         *TypingStatusOp
	type_      *RecipientType
	to         *[]int64
	channelId  *int64
	topic      *string
}

// Whether the user has started (&#x60;\\\&quot;start\\\&quot;&#x60;) or stopped (&#x60;\\\&quot;stop\\\&quot;&#x60;) typing.
func (r SetTypingStatusRequest) Op(op TypingStatusOp) SetTypingStatusRequest {
	r.op = &op
	return r
}

// Type of the message being composed.  **Changes**: In Zulip 9.0 (feature level 248), &#x60;\\\&quot;channel\\\&quot;&#x60; was added as an additional value for this parameter to indicate a channel message is being composed.  In Zulip 8.0 (feature level 215), stopped supporting &#x60;\\\&quot;private\\\&quot;&#x60; as a valid value for this parameter.  In Zulip 7.0 (feature level 174), &#x60;\\\&quot;direct\\\&quot;&#x60; was added as the preferred way to indicate a direct message is being composed, becoming the default value for this parameter and deprecating the original &#x60;\\\&quot;private\\\&quot;&#x60;.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.
func (r SetTypingStatusRequest) Type_(type_ RecipientType) SetTypingStatusRequest {
	r.type_ = &type_
	return r
}

// User Ids of the recipients of the message being typed. Required for the &#x60;\\\&quot;direct\\\&quot;&#x60; type. Ignored in the case of &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; type.  Clients should send a JSON-encoded list of user Ids, even if there is only one recipient.  **Changes**: In Zulip 8.0 (feature level 215), stopped using this parameter for the &#x60;\\\&quot;stream\\\&quot;&#x60; type. Previously, in the case of the &#x60;\\\&quot;stream\\\&quot;&#x60; type, it accepted a single-element list containing the Id of the channel. A new parameter, &#x60;stream_id&#x60;, is now used for this. Note that the &#x60;\\\&quot;channel\\\&quot;&#x60; type did not exist at this feature level.  Support for typing notifications for channel&#39; messages is new in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  Before Zulip 2.0.0, this parameter accepted only a JSON-encoded list of email addresses. Support for the email address-based format was removed in Zulip 3.0 (feature level 11).
func (r SetTypingStatusRequest) To(to []int64) SetTypingStatusRequest {
	r.to = &to
	return r
}

// Id of the channel in which the message is being typed. Required for the &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; type. Ignored in the case of &#x60;\\\&quot;direct\\\&quot;&#x60; type.  **Changes**: New in Zulip 8.0 (feature level 215). Previously, a single-element list containing the Id of the channel was passed in &#x60;to&#x60; parameter.
func (r SetTypingStatusRequest) ChannelId(channelId int64) SetTypingStatusRequest {
	r.channelId = &channelId
	return r
}

// Topic to which message is being typed. Required for the &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; type. Ignored in the case of &#x60;\\\&quot;direct\\\&quot;&#x60; type.  Note: When &#x60;\\\&quot;(no topic)\\\&quot;&#x60; or the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](zulip.com/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 372), &#x60;\\\&quot;(no topic)\\\&quot;&#x60; was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.
func (r SetTypingStatusRequest) Topic(topic string) SetTypingStatusRequest {
	r.topic = &topic
	return r
}

func (r SetTypingStatusRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.SetTypingStatusExecute(r)
}

/*
SetTypingStatus Set \"typing\" status

Notify other users whether the current user is
[typing a message][help-typing].

Clients implementing Zulip's typing notifications
protocol should work as follows:

  - Send a request to this endpoint with `"op": "start"` when a user
    starts composing a message.
  - While the user continues to actively type or otherwise interact with
    the compose UI (e.g. interacting with the compose box emoji picker),
    send regular `"op": "start"` requests to this endpoint, using
    `server_typing_started_wait_period_milliseconds` in the
    [`POST /register`][api-register] response as the time interval
    between each request.
  - Send a request to this endpoint with `"op": "stop"` when a user
    has stopped using the compose UI for the time period indicated by
    `server_typing_stopped_wait_period_milliseconds` in the
    [`POST /register`][api-register] response or when a user
    cancels the compose action (if it had previously sent a "start"
    notification for that compose action).
  - Start displaying a visual typing indicator for a given conversation
    when a [`typing op:start`][start-typing] event is received
    from the server.
  - Continue displaying a visual typing indicator for the conversation
    until a [`typing op:stop`][stop-typing] event is received
    from the server or the time period indicated by
    `server_typing_started_expiry_period_milliseconds` in the
    [`POST /register`][api-register] response has passed without
    a new `typing "op": "start"` event for the conversation.

This protocol is designed to allow the server-side typing notifications
implementation to be stateless while being resilient as network failures
will not result in a user being incorrectly displayed as perpetually
typing.

See the subsystems documentation on [typing indicators][typing-protocol-docs]
for additional design details on Zulip's typing notifications protocol.

**Changes**: Clients shouldn't care about the APIs prior to Zulip 8.0 (feature level 215)
for channel typing notifications, as no client actually implemented
the previous API for those.

Support for displaying channel typing notifications was new
in Zulip 4.0 (feature level 58). Clients should indicate they support
processing channel typing notifications via the `stream_typing_notifications`
value in the `client_capabilities` parameter of the
[`POST /register`][client-capabilities] endpoint.

[help-typing]: /help/typing-notifications
[api-register]: /api/register-queue
[start-typing]: /api/get-events#typing-start
[stop-typing]: /api/get-events#typing-stop
[client-capabilities]: /api/register-queue#parameter-client_capabilities
[typing-protocol-docs]: https://zulip.readthedocs.io/en/latest/subsystems/typing-indicators.html

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SetTypingStatusRequest
*/
func (c *simpleClient) SetTypingStatus(ctx context.Context) SetTypingStatusRequest {
	return SetTypingStatusRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) SetTypingStatusExecute(r SetTypingStatusRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/typing"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.op == nil {
		return localVarReturnValue, nil, reportError("op is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "type", r.type_, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "op", r.op, "", "")
	if r.to != nil {
		paramJson, err := parameterToJson(r.to)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("to", paramJson)
	}
	if r.channelId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "stream_id", r.channelId, "form", "")
	}
	if r.topic != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "topic", r.topic, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SetTypingStatusForMessageEditRequest struct {
	ctx        context.Context
	ApiService UsersAPI
	messageId  int64
	op         *TypingStatusOp
}

// Whether the user has started (&#x60;\\\&quot;start\\\&quot;&#x60;) or stopped (&#x60;\\\&quot;stop\\\&quot;&#x60;) editing.
func (r SetTypingStatusForMessageEditRequest) Op(op TypingStatusOp) SetTypingStatusForMessageEditRequest {
	r.op = &op
	return r
}

func (r SetTypingStatusForMessageEditRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.SetTypingStatusForMessageEditExecute(r)
}

/*
SetTypingStatusForMessageEdit Set \"typing\" status for message editing

Notify other users whether the current user is editing a message.

Typing notifications for editing messages follow the same protocol as
[set-typing-status](zulip.com/api/set-typing-status, see that endpoint for
details.

**Changes**: Before Zulip 10.0 (feature level 361), the endpoint was
named `/message_edit_typing` with `message_id` a required parameter in
the request body. Clients are recommended to start using sending these
typing notifications starting from this feature level.

New in Zulip 10.0 (feature level 351). Previously, typing notifications were
not available when editing messages.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param messageId The target message's Id.
	@return SetTypingStatusForMessageEditRequest
*/
func (c *simpleClient) SetTypingStatusForMessageEdit(ctx context.Context, messageId int64) SetTypingStatusForMessageEditRequest {
	return SetTypingStatusForMessageEditRequest{
		ApiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) SetTypingStatusForMessageEditExecute(r SetTypingStatusForMessageEditRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages/{message_id}/typing"
	localVarPath = strings.Replace(localVarPath, "{"+"message_id"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.op == nil {
		return localVarReturnValue, nil, reportError("op is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "op", r.op, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UnmuteUserRequest struct {
	ctx         context.Context
	ApiService  UsersAPI
	mutedUserId int64
}

func (r UnmuteUserRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UnmuteUserExecute(r)
}

/*
UnmuteUser Unmute a user

[Unmute a user](zulip.com/help/mute-a-user#see-your-list-of-muted-users
from the perspective of the requesting user.

**Changes**: New in Zulip 4.0 (feature level 48).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param mutedUserId The Id of the user to mute/unmute.  **Changes**: Before Zulip 8.0 (feature level 188), bot users could not be muted/unmuted, and specifying a bot user's Id returned an error response.
	@return UnmuteUserRequest
*/
func (c *simpleClient) UnmuteUser(ctx context.Context, mutedUserId int64) UnmuteUserRequest {
	return UnmuteUserRequest{
		ApiService:  c,
		ctx:         ctx,
		mutedUserId: mutedUserId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UnmuteUserExecute(r UnmuteUserRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/muted_users/{muted_user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"muted_user_id"+"}", url.PathEscape(parameterValueToString(r.mutedUserId, "mutedUserId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdatePresenceRequest struct {
	ctx              context.Context
	ApiService       UsersAPI
	status           *PresenceStatus
	lastUpdateId     *int64
	historyLimitDays *int32
	newUserInput     *bool
	pingOnly         *bool
	slimPresence     *bool
}

// The status of the user on this client.  Clients should report the user as &#x60;\\\&quot;active\\\&quot;&#x60; on this device if the client knows that the user is presently using the device (and thus would potentially see a notification immediately), even if the user has not directly interacted with the Zulip client.  Otherwise, it should report the user as &#x60;\\\&quot;idle\\\&quot;&#x60;.  See the related [&#x60;new_user_input&#x60;](#parameter-new_user_input) parameter for how a client should report whether the user is actively using the Zulip client.
func (r UpdatePresenceRequest) Status(status PresenceStatus) UpdatePresenceRequest {
	r.status = &status
	return r
}

// The identifier that specifies what presence data the client already has received, which allows the server to only return more recent user presence data.  This should be set to &#x60;-1&#x60; during initialization of the client in order to fetch all user presence data, unless the client is obtaining initial user presence metadata from the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint.  In subsequent queries to this endpoint, this value should be set to the most recent value of &#x60;presence_last_update_id&#x60; returned by the server in this endpoint&#39;s response, which implements incremental fetching of user presence data.  When this parameter is passed, the user presence data in the response will always be in the modern format.  **Changes**: New in Zulip 9.0 (feature level 263). Previously, the server sent user presence data for all users who had been active in the last two weeks unconditionally.
func (r UpdatePresenceRequest) LastUpdateId(lastUpdateId int64) UpdatePresenceRequest {
	r.lastUpdateId = &lastUpdateId
	return r
}

// Limits how far back in time to fetch user presence data. If not specified, defaults to 14 days. A value of N means that the oldest presence data fetched will be from at most N days ago.  Note that this is only useful during the initial user presence data fetch, as subsequent fetches should use the &#x60;last_update_id&#x60; parameter, which will act as the limit on how much presence data is returned. &#x60;history_limit_days&#x60; is ignored if &#x60;last_update_id&#x60; is passed with a value greater than &#x60;0&#x60;, indicating that the client already has some presence data.  **Changes**: New in Zulip 10.0 (feature level 288).
func (r UpdatePresenceRequest) HistoryLimitDays(historyLimitDays int32) UpdatePresenceRequest {
	r.historyLimitDays = &historyLimitDays
	return r
}

// Whether the user has interacted with the client (e.g. moved the mouse, used the keyboard, etc.) since the previous presence request from this client.  The server uses data from this parameter to implement certain [usage statistics](zulip.com/help/analytics.  User interface clients that might run in the background, without the user ever interacting with them, should be careful to only pass &#x60;true&#x60; if the user has actually interacted with the client in order to avoid corrupting usage statistics graphs.
func (r UpdatePresenceRequest) NewUserInput(newUserInput bool) UpdatePresenceRequest {
	r.newUserInput = &newUserInput
	return r
}

// Whether the client is sending a ping-only request, meaning it only wants to update the user&#39;s presence &#x60;status&#x60; on the server.  Otherwise, also requests the server return user presence data for all users in the organization, which is further specified by the [&#x60;last_update_id&#x60;](#parameter-last_update_id) parameter.
func (r UpdatePresenceRequest) PingOnly(pingOnly bool) UpdatePresenceRequest {
	r.pingOnly = &pingOnly
	return r
}

// Legacy parameter for configuring the format (modern or legacy) in which the server will return user presence data for the organization.  Modern clients should use [&#x60;last_update_id&#x60;](#parameter-last_update_id), which guarantees that user presence data will be returned in the modern format, and should not pass this parameter as &#x60;true&#x60; unless interacting with an older server.  Legacy clients that do not yet support &#x60;last_update_id&#x60; may use the value of &#x60;true&#x60; to request the modern format for user presence data.  **Note**: The legacy format for user presence data will be removed entirely in a future release.  **Changes**: **Deprecated** in Zulip 9.0 (feature level 263). Using the modern &#x60;last_update_id&#x60; parameter is the recommended way to request the modern format for user presence data.  New in Zulip 3.0 (no feature level as it was an unstable API at that point).
func (r UpdatePresenceRequest) SlimPresence(slimPresence bool) UpdatePresenceRequest {
	r.slimPresence = &slimPresence
	return r
}

func (r UpdatePresenceRequest) Execute() (*UpdatePresenceResponse, *http.Response, error) {
	return r.ApiService.UpdatePresenceExecute(r)
}

/*
UpdatePresence Update your presence

Update the current user's [presence][availability] and fetch presence data
of other users in the organization.

This endpoint is meant to be used by clients for both:

  - Reporting the current user's presence status (`"active"` or `"idle"`)
    to the server.

  - Obtaining the presence data of all other users in the organization via
    regular polling.

Accurate user presence is one of the most expensive parts of any
chat application (in terms of bandwidth and other resources). Therefore,
it is important that clients implementing Zulip's user presence system
use the modern [`last_update_id`](#parameter-last_update_id) protocol to
minimize fetching duplicate user presence data.

Client apps implementing presence are recommended to also consume [`presence`
events](zulip.com/api/get-events#presence), in order to learn about newly online users
immediately.

The Zulip server is responsible for implementing [invisible mode][invisible],
which disables sharing a user's presence data. Nonetheless, clients
should check the `presence_enabled` field in user objects in order to
display the current user as online or offline based on whether they are
sharing their presence information.

**Changes**: As of Zulip 8.0 (feature level 228), if the
`CAN_ACCESS_ALL_USERS_GROUP_LIMITS_PRESENCE` server-level setting is
`true`, then user presence data in the response is [limited to users
the current user can see/access][limit-visibility].

[limit-visibility]: /help/guest-users#configure-whether-guests-can-see-all-other-users
[invisible]: /help/status-and-availability#invisible-mode
[availability]: /help/status-and-availability#availability

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UpdatePresenceRequest
*/
func (c *simpleClient) UpdatePresence(ctx context.Context) UpdatePresenceRequest {
	return UpdatePresenceRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdatePresenceResponse
func (c *simpleClient) UpdatePresenceExecute(r UpdatePresenceRequest) (*UpdatePresenceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdatePresenceResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/presence"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.status == nil {
		return localVarReturnValue, nil, reportError("status is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.lastUpdateId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "last_update_id", r.lastUpdateId, "", "")
	}
	if r.historyLimitDays != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "history_limit_days", r.historyLimitDays, "", "")
	}
	if r.newUserInput != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "new_user_input", r.newUserInput, "", "")
	}
	if r.pingOnly != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ping_only", r.pingOnly, "", "")
	}
	if r.slimPresence != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "slim_presence", r.slimPresence, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "status", r.status, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateSettingsRequest struct {
	ctx                                            context.Context
	ApiService                                     UsersAPI
	fullName                                       *string
	email                                          *string
	oldPassword                                    *string
	newPassword                                    *string
	twentyFourHourTime                             *bool
	webMarkReadOnScrollPolicy                      *int32
	webChannelDefaultView                          *int32
	starredMessageCounts                           *bool
	receivesTypingNotifications                    *bool
	webSuggestUpdateTimezone                       *bool
	fluidLayoutWidth                               *bool
	highContrastMode                               *bool
	webFontSizePx                                  *int32
	webLineHeightPercent                           *int32
	colorScheme                                    *ColorScheme
	enableDraftsSynchronization                    *bool
	translateEmoticons                             *bool
	displayEmojiReactionUsers                      *bool
	defaultLanguage                                *string
	webHomeView                                    *string
	webEscapeNavigatesToHomeView                   *bool
	leftSideUserlist                               *bool
	emojiset                                       *string
	demoteInactiveChannels                         *int32
	userListStyle                                  *int32
	webAnimateImagePreviews                        *string
	webChannelUnreadsCountDisplayPolicy            *int32
	hideAiFeatures                                 *bool
	webLeftSidebarShowChannelFolders               *bool
	webLeftSidebarUnreadsCountSummary              *bool
	timezone                                       *string
	enableChannelDesktopNotifications              *bool
	enableChannelEmailNotifications                *bool
	enableChannelPushNotifications                 *bool
	enableChannelAudibleNotifications              *bool
	notificationSound                              *string
	enableDesktopNotifications                     *bool
	enableSounds                                   *bool
	emailNotificationsBatchingPeriodSeconds        *int32
	enableOfflineEmailNotifications                *bool
	enableOfflinePushNotifications                 *bool
	enableOnlinePushNotifications                  *bool
	enableFollowedTopicDesktopNotifications        *bool
	enableFollowedTopicEmailNotifications          *bool
	enableFollowedTopicPushNotifications           *bool
	enableFollowedTopicAudibleNotifications        *bool
	enableDigestEmails                             *bool
	enableMarketingEmails                          *bool
	enableLoginEmails                              *bool
	messageContentInEmailNotifications             *bool
	pmContentInDesktopNotifications                *bool
	wildcardMentionsNotify                         *bool
	enableFollowedTopicWildcardMentionsNotify      *bool
	desktopIconCountDisplay                        *int32
	realmNameInEmailNotificationsPolicy            *int32
	automaticallyFollowTopicsPolicy                *int32
	automaticallyUnmuteTopicsInMutedChannelsPolicy *int32
	automaticallyFollowTopicsWhereMentioned        *bool
	resolvedTopicNoticeAutoReadPolicy              *string
	presenceEnabled                                *bool
	enterSends                                     *bool
	sendPrivateTypingNotifications                 *bool
	sendChannelTypingNotifications                 *bool
	sendReadReceipts                               *bool
	allowPrivateDataExport                         *bool
	emailAddressVisibility                         *int32
	webNavigateToSentMessage                       *bool
}

// A new display name for the user.
func (r UpdateSettingsRequest) FullName(fullName string) UpdateSettingsRequest {
	r.fullName = &fullName
	return r
}

// Asks the server to initiate a confirmation sequence to change the user&#39;s email address to the indicated value. The user will need to demonstrate control of the new email address by clicking a confirmation link sent to that address.
func (r UpdateSettingsRequest) Email(email string) UpdateSettingsRequest {
	r.email = &email
	return r
}

// The user&#39;s old Zulip password (or LDAP password, if LDAP authentication is in use).  Required only when sending the &#x60;new_password&#x60; parameter.
func (r UpdateSettingsRequest) OldPassword(oldPassword string) UpdateSettingsRequest {
	r.oldPassword = &oldPassword
	return r
}

// The user&#39;s new Zulip password (or LDAP password, if LDAP authentication is in use).  The &#x60;old_password&#x60; parameter must be included in the request.
func (r UpdateSettingsRequest) NewPassword(newPassword string) UpdateSettingsRequest {
	r.newPassword = &newPassword
	return r
}

// Whether time should be [displayed in 24-hour notation](zulip.com/help/change-the-time-format.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.
func (r UpdateSettingsRequest) TwentyFourHourTime(twentyFourHourTime bool) UpdateSettingsRequest {
	r.twentyFourHourTime = &twentyFourHourTime
	return r
}

// Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \\\&quot;Always\\\&quot; setting when marking messages as read.
func (r UpdateSettingsRequest) WebMarkReadOnScrollPolicy(webMarkReadOnScrollPolicy int32) UpdateSettingsRequest {
	r.webMarkReadOnScrollPolicy = &webMarkReadOnScrollPolicy
	return r
}

// Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \\\&quot;Top unread topic in channel\\\&quot; is new in Zulip 11.0 (feature level 401).  The \\\&quot;List of topics\\\&quot; option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \\\&quot;Channel feed\\\&quot; behavior.
func (r UpdateSettingsRequest) WebChannelDefaultView(webChannelDefaultView int32) UpdateSettingsRequest {
	r.webChannelDefaultView = &webChannelDefaultView
	return r
}

// Whether clients should display the [number of starred messages](zulip.com/help/star-a-message#display-the-number-of-starred-messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.
func (r UpdateSettingsRequest) StarredMessageCounts(starredMessageCounts bool) UpdateSettingsRequest {
	r.starredMessageCounts = &starredMessageCounts
	return r
}

// Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  By default, this is set to true, enabling user to receive typing notifications from other users.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.
func (r UpdateSettingsRequest) ReceivesTypingNotifications(receivesTypingNotifications bool) UpdateSettingsRequest {
	r.receivesTypingNotifications = &receivesTypingNotifications
	return r
}

// Whether the user should be shown an alert, offering to update their [profile time zone](zulip.com/help/change-your-timezone, when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.  **Changes**: New in Zulip 10.0 (feature level 329).
func (r UpdateSettingsRequest) WebSuggestUpdateTimezone(webSuggestUpdateTimezone bool) UpdateSettingsRequest {
	r.webSuggestUpdateTimezone = &webSuggestUpdateTimezone
	return r
}

// Whether to use the [maximum available screen width](zulip.com/help/enable-full-width-display) for the web app&#39;s center panel (message feed, recent conversations) on wide screens.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.
func (r UpdateSettingsRequest) FluidLayoutWidth(fluidLayoutWidth bool) UpdateSettingsRequest {
	r.fluidLayoutWidth = &fluidLayoutWidth
	return r
}

// This setting is reserved for use to control variations in Zulip&#39;s design to help visually impaired users.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.
func (r UpdateSettingsRequest) HighContrastMode(highContrastMode bool) UpdateSettingsRequest {
	r.highContrastMode = &highContrastMode
	return r
}

// User-configured primary &#x60;font-size&#x60; for the web application, in pixels.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.
func (r UpdateSettingsRequest) WebFontSizePx(webFontSizePx int32) UpdateSettingsRequest {
	r.webFontSizePx = &webFontSizePx
	return r
}

// User-configured primary &#x60;line-height&#x60; for the web application, in percent, so a value of 120 represents a &#x60;line-height&#x60; of 1.2.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.
func (r UpdateSettingsRequest) WebLineHeightPercent(webLineHeightPercent int32) UpdateSettingsRequest {
	r.webLineHeightPercent = &webLineHeightPercent
	return r
}

// Controls which [color theme](zulip.com/help/dark-theme) to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard &#x60;prefers-color-scheme&#x60; media query.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.
func (r UpdateSettingsRequest) ColorScheme(colorScheme ColorScheme) UpdateSettingsRequest {
	r.colorScheme = &colorScheme
	return r
}

// A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.  **Changes**: New in Zulip 5.0 (feature level 87).
func (r UpdateSettingsRequest) EnableDraftsSynchronization(enableDraftsSynchronization bool) UpdateSettingsRequest {
	r.enableDraftsSynchronization = &enableDraftsSynchronization
	return r
}

// Whether to [translate emoticons to emoji](zulip.com/help/configure-emoticon-translations) in messages the user sends.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.
func (r UpdateSettingsRequest) TranslateEmoticons(translateEmoticons bool) UpdateSettingsRequest {
	r.translateEmoticons = &translateEmoticons
	return r
}

// Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.  **Changes**: New in Zulip 6.0 (feature level 125).
func (r UpdateSettingsRequest) DisplayEmojiReactionUsers(displayEmojiReactionUsers bool) UpdateSettingsRequest {
	r.displayEmojiReactionUsers = &displayEmojiReactionUsers
	return r
}

// What [default language](zulip.com/help/change-your-language) to use for the account.  This controls both the Zulip UI as well as email notifications sent to the user.  The value needs to be a standard language code that the Zulip server has translation data for; for example, &#x60;\\\&quot;en\\\&quot;&#x60; for English or &#x60;\\\&quot;de\\\&quot;&#x60; for German.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 63).
func (r UpdateSettingsRequest) DefaultLanguage(defaultLanguage string) UpdateSettingsRequest {
	r.defaultLanguage = &defaultLanguage
	return r
}

// The [home view](zulip.com/help/configure-home-view) used when opening a new Zulip web app window or hitting the &#x60;Esc&#x60; keyboard shortcut repeatedly.  - \\\&quot;recent_topics\\\&quot; - Recent conversations view - \\\&quot;inbox\\\&quot; - Inbox view - \\\&quot;all_messages\\\&quot; - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;default_view&#x60;, which was new in Zulip 4.0 (feature level 42).  Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).
func (r UpdateSettingsRequest) WebHomeView(webHomeView string) UpdateSettingsRequest {
	r.webHomeView = &webHomeView
	return r
}

// Whether the escape key navigates to the [configured home view](zulip.com/help/configure-home-view.  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;escape_navigates_to_default_view&#x60;, which was new in Zulip 5.0 (feature level 107).
func (r UpdateSettingsRequest) WebEscapeNavigatesToHomeView(webEscapeNavigatesToHomeView bool) UpdateSettingsRequest {
	r.webEscapeNavigatesToHomeView = &webEscapeNavigatesToHomeView
	return r
}

// Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.
func (r UpdateSettingsRequest) LeftSideUserlist(leftSideUserlist bool) UpdateSettingsRequest {
	r.leftSideUserlist = &leftSideUserlist
	return r
}

// The user&#39;s configured [emoji set](zulip.com/help/emoji-and-emoticons#use-emoticons, used to display emoji to the user everywhere they appear in the UI.  - \\\&quot;google\\\&quot; - Google modern - \\\&quot;google-blob\\\&quot; - Google classic - \\\&quot;twitter\\\&quot; - Twitter - \\\&quot;text\\\&quot; - Plain text  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).
func (r UpdateSettingsRequest) Emojiset(emojiset string) UpdateSettingsRequest {
	r.emojiset = &emojiset
	return r
}

// Whether to [hide inactive channels](zulip.com/help/manage-inactive-channels) in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.
func (r UpdateSettingsRequest) DemoteInactiveChannels(demoteInactiveChannels int32) UpdateSettingsRequest {
	r.demoteInactiveChannels = &demoteInactiveChannels
	return r
}

// The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).
func (r UpdateSettingsRequest) UserListStyle(userListStyle int32) UpdateSettingsRequest {
	r.userListStyle = &userListStyle
	return r
}

// Controls how animated images should be played in the message feed in the web/desktop application.  - \\\&quot;always\\\&quot; - Always play the animated images in the message feed. - \\\&quot;on_hover\\\&quot; - Play the animated images on hover over them in the message feed. - \\\&quot;never\\\&quot; - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275).
func (r UpdateSettingsRequest) WebAnimateImagePreviews(webAnimateImagePreviews string) UpdateSettingsRequest {
	r.webAnimateImagePreviews = &webAnimateImagePreviews
	return r
}

// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).
func (r UpdateSettingsRequest) WebChannelUnreadsCountDisplayPolicy(webChannelUnreadsCountDisplayPolicy int32) UpdateSettingsRequest {
	r.webChannelUnreadsCountDisplayPolicy = &webChannelUnreadsCountDisplayPolicy
	return r
}

// Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).
func (r UpdateSettingsRequest) HideAiFeatures(hideAiFeatures bool) UpdateSettingsRequest {
	r.hideAiFeatures = &hideAiFeatures
	return r
}

// Determines whether the web/desktop application&#39;s left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).
func (r UpdateSettingsRequest) WebLeftSidebarShowChannelFolders(webLeftSidebarShowChannelFolders bool) UpdateSettingsRequest {
	r.webLeftSidebarShowChannelFolders = &webLeftSidebarShowChannelFolders
	return r
}

// Determines whether the web/desktop application&#39;s left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).
func (r UpdateSettingsRequest) WebLeftSidebarUnreadsCountSummary(webLeftSidebarUnreadsCountSummary bool) UpdateSettingsRequest {
	r.webLeftSidebarUnreadsCountSummary = &webLeftSidebarUnreadsCountSummary
	return r
}

// The IANA identifier of the user&#39;s [profile time zone](zulip.com/help/change-your-timezone, which is used primarily to display the user&#39;s local time to other users.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/display&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).
func (r UpdateSettingsRequest) Timezone(timezone string) UpdateSettingsRequest {
	r.timezone = &timezone
	return r
}

// Enable visual desktop notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableChannelDesktopNotifications(enableChannelDesktopNotifications bool) UpdateSettingsRequest {
	r.enableChannelDesktopNotifications = &enableChannelDesktopNotifications
	return r
}

// Enable email notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableChannelEmailNotifications(enableChannelEmailNotifications bool) UpdateSettingsRequest {
	r.enableChannelEmailNotifications = &enableChannelEmailNotifications
	return r
}

// Enable mobile notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableChannelPushNotifications(enableChannelPushNotifications bool) UpdateSettingsRequest {
	r.enableChannelPushNotifications = &enableChannelPushNotifications
	return r
}

// Enable audible desktop notifications for channel messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableChannelAudibleNotifications(enableChannelAudibleNotifications bool) UpdateSettingsRequest {
	r.enableChannelAudibleNotifications = &enableChannelAudibleNotifications
	return r
}

// Notification sound name.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 63).
func (r UpdateSettingsRequest) NotificationSound(notificationSound string) UpdateSettingsRequest {
	r.notificationSound = &notificationSound
	return r
}

// Enable visual desktop notifications for direct messages and @-mentions.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableDesktopNotifications(enableDesktopNotifications bool) UpdateSettingsRequest {
	r.enableDesktopNotifications = &enableDesktopNotifications
	return r
}

// Enable audible desktop notifications for direct messages and @-mentions.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableSounds(enableSounds bool) UpdateSettingsRequest {
	r.enableSounds = &enableSounds
	return r
}

// The duration (in seconds) for which the server should wait to batch email notifications before sending them.  **Changes**: New in Zulip 5.0 (feature level 82)
func (r UpdateSettingsRequest) EmailNotificationsBatchingPeriodSeconds(emailNotificationsBatchingPeriodSeconds int32) UpdateSettingsRequest {
	r.emailNotificationsBatchingPeriodSeconds = &emailNotificationsBatchingPeriodSeconds
	return r
}

// Enable email notifications for direct messages and @-mentions received when the user is offline.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableOfflineEmailNotifications(enableOfflineEmailNotifications bool) UpdateSettingsRequest {
	r.enableOfflineEmailNotifications = &enableOfflineEmailNotifications
	return r
}

// Enable mobile notification for direct messages and @-mentions received when the user is offline.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableOfflinePushNotifications(enableOfflinePushNotifications bool) UpdateSettingsRequest {
	r.enableOfflinePushNotifications = &enableOfflinePushNotifications
	return r
}

// Enable mobile notification for direct messages and @-mentions received when the user is online.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableOnlinePushNotifications(enableOnlinePushNotifications bool) UpdateSettingsRequest {
	r.enableOnlinePushNotifications = &enableOnlinePushNotifications
	return r
}

// Enable visual desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicDesktopNotifications(enableFollowedTopicDesktopNotifications bool) UpdateSettingsRequest {
	r.enableFollowedTopicDesktopNotifications = &enableFollowedTopicDesktopNotifications
	return r
}

// Enable email notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicEmailNotifications(enableFollowedTopicEmailNotifications bool) UpdateSettingsRequest {
	r.enableFollowedTopicEmailNotifications = &enableFollowedTopicEmailNotifications
	return r
}

// Enable push notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicPushNotifications(enableFollowedTopicPushNotifications bool) UpdateSettingsRequest {
	r.enableFollowedTopicPushNotifications = &enableFollowedTopicPushNotifications
	return r
}

// Enable audible desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicAudibleNotifications(enableFollowedTopicAudibleNotifications bool) UpdateSettingsRequest {
	r.enableFollowedTopicAudibleNotifications = &enableFollowedTopicAudibleNotifications
	return r
}

// Enable digest emails when the user is away.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableDigestEmails(enableDigestEmails bool) UpdateSettingsRequest {
	r.enableDigestEmails = &enableDigestEmails
	return r
}

// Enable marketing emails. Has no function outside Zulip Cloud.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableMarketingEmails(enableMarketingEmails bool) UpdateSettingsRequest {
	r.enableMarketingEmails = &enableMarketingEmails
	return r
}

// Enable email notifications for new logins to account.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) EnableLoginEmails(enableLoginEmails bool) UpdateSettingsRequest {
	r.enableLoginEmails = &enableLoginEmails
	return r
}

// Include the message&#39;s content in email notifications for new messages.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) MessageContentInEmailNotifications(messageContentInEmailNotifications bool) UpdateSettingsRequest {
	r.messageContentInEmailNotifications = &messageContentInEmailNotifications
	return r
}

// Include content of direct messages in desktop notifications.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) PmContentInDesktopNotifications(pmContentInDesktopNotifications bool) UpdateSettingsRequest {
	r.pmContentInDesktopNotifications = &pmContentInDesktopNotifications
	return r
}

// Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) WildcardMentionsNotify(wildcardMentionsNotify bool) UpdateSettingsRequest {
	r.wildcardMentionsNotify = &wildcardMentionsNotify
	return r
}

// Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.  **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicWildcardMentionsNotify(enableFollowedTopicWildcardMentionsNotify bool) UpdateSettingsRequest {
	r.enableFollowedTopicWildcardMentionsNotify = &enableFollowedTopicWildcardMentionsNotify
	return r
}

// Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added &#x60;DMs, mentions, and followed topics&#x60; option, renumbering the options to insert it in order.  Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) DesktopIconCountDisplay(desktopIconCountDisplay int32) UpdateSettingsRequest {
	r.desktopIconCountDisplay = &desktopIconCountDisplay
	return r
}

// Whether to [include organization name in subject of message notification emails](zulip.com/help/email-notifications#include-organization-name-in-subject-line.  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous &#x60;realm_name_in_notifications&#x60; boolean; &#x60;true&#x60; corresponded to &#x60;Always&#x60;, and &#x60;false&#x60; to &#x60;Never&#x60;.  Before Zulip 5.0 (feature level 80), the previous &#x60;realm_name_in_notifications&#x60; setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) RealmNameInEmailNotificationsPolicy(realmNameInEmailNotificationsPolicy int32) UpdateSettingsRequest {
	r.realmNameInEmailNotificationsPolicy = &realmNameInEmailNotificationsPolicy
	return r
}

// Which [topics to follow automatically](zulip.com/help/mute-a-topic.  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).
func (r UpdateSettingsRequest) AutomaticallyFollowTopicsPolicy(automaticallyFollowTopicsPolicy int32) UpdateSettingsRequest {
	r.automaticallyFollowTopicsPolicy = &automaticallyFollowTopicsPolicy
	return r
}

// Which [topics to unmute automatically in muted channels](zulip.com/help/mute-a-topic.  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).
func (r UpdateSettingsRequest) AutomaticallyUnmuteTopicsInMutedChannelsPolicy(automaticallyUnmuteTopicsInMutedChannelsPolicy int32) UpdateSettingsRequest {
	r.automaticallyUnmuteTopicsInMutedChannelsPolicy = &automaticallyUnmuteTopicsInMutedChannelsPolicy
	return r
}

// Whether the server will automatically mark the user as following topics where the user is mentioned.  **Changes**: New in Zulip 8.0 (feature level 235).
func (r UpdateSettingsRequest) AutomaticallyFollowTopicsWhereMentioned(automaticallyFollowTopicsWhereMentioned bool) UpdateSettingsRequest {
	r.automaticallyFollowTopicsWhereMentioned = &automaticallyFollowTopicsWhereMentioned
	return r
}

// Controls whether the resolved-topic notices are marked as read.  - \\\&quot;always\\\&quot; - Always mark resolved-topic notices as read. - \\\&quot;except_followed\\\&quot; - Mark resolved-topic notices as read in topics not followed by the user. - \\\&quot;never\\\&quot; - Never mark resolved-topic notices as read.  **Changes**: New in Zulip 11.0 (feature level 385).
func (r UpdateSettingsRequest) ResolvedTopicNoticeAutoReadPolicy(resolvedTopicNoticeAutoReadPolicy string) UpdateSettingsRequest {
	r.resolvedTopicNoticeAutoReadPolicy = &resolvedTopicNoticeAutoReadPolicy
	return r
}

// Display the presence status to other users when online.  **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the &#x60;PATCH /settings/notifications&#x60; endpoint.
func (r UpdateSettingsRequest) PresenceEnabled(presenceEnabled bool) UpdateSettingsRequest {
	r.presenceEnabled = &presenceEnabled
	return r
}

// Whether pressing Enter in the compose box sends a message (or saves a message edit).  **Changes**: Before Zulip 5.0 (feature level 81), this setting was managed by the &#x60;POST /users/me/enter-sends&#x60; endpoint, with the same parameter format.
func (r UpdateSettingsRequest) EnterSends(enterSends bool) UpdateSettingsRequest {
	r.enterSends = &enterSends
	return r
}

// Whether [typing notifications](zulip.com/help/typing-notifications) be sent when composing direct messages.  **Changes**: New in Zulip 5.0 (feature level 105).
func (r UpdateSettingsRequest) SendPrivateTypingNotifications(sendPrivateTypingNotifications bool) UpdateSettingsRequest {
	r.sendPrivateTypingNotifications = &sendPrivateTypingNotifications
	return r
}

// Whether [typing notifications](zulip.com/help/typing-notifications) be sent when composing channel messages.  **Changes**: New in Zulip 5.0 (feature level 105).
func (r UpdateSettingsRequest) SendChannelTypingNotifications(sendChannelTypingNotifications bool) UpdateSettingsRequest {
	r.sendChannelTypingNotifications = &sendChannelTypingNotifications
	return r
}

// Whether other users are allowed to see whether you&#39;ve read messages.  **Changes**: New in Zulip 5.0 (feature level 105).
func (r UpdateSettingsRequest) SendReadReceipts(sendReadReceipts bool) UpdateSettingsRequest {
	r.sendReadReceipts = &sendReadReceipts
	return r
}

// Whether organization administrators are allowed to export your private data.  **Changes**: New in Zulip 10.0 (feature level 293).
func (r UpdateSettingsRequest) AllowPrivateDataExport(allowPrivateDataExport bool) UpdateSettingsRequest {
	r.allowPrivateDataExport = &allowPrivateDataExport
	return r
}

// The [policy][permission-level] this user has selected for [which other users][help-email-visibility] in this organization can see their real email address.  - 1 &#x3D; Everyone - 2 &#x3D; Members only - 3 &#x3D; Administrators only - 4 &#x3D; Nobody - 5 &#x3D; Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.  [permission-level]: /api/roles-and-permissions#permission-levels [help-email-visibility]: /help/configure-email-visibility
func (r UpdateSettingsRequest) EmailAddressVisibility(emailAddressVisibility int32) UpdateSettingsRequest {
	r.emailAddressVisibility = &emailAddressVisibility
	return r
}

// Web/desktop app setting for whether the user&#39;s view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.
func (r UpdateSettingsRequest) WebNavigateToSentMessage(webNavigateToSentMessage bool) UpdateSettingsRequest {
	r.webNavigateToSentMessage = &webNavigateToSentMessage
	return r
}

func (r UpdateSettingsRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateSettingsExecute(r)
}

/*
UpdateSettings Update settings

This endpoint is used to edit the current user's settings.

**Changes**: Removed `dense_mode` setting in Zulip 10.0
(feature level 364) as we now have `web_font_size_px` and
`web_line_height_percent` settings for more control.

Prior to Zulip 5.0 (feature level 80), this endpoint only
supported the `full_name`, `email`, `old_password`, and
`new_password` parameters. Notification settings were
managed by `PATCH /settings/notifications`, and all other
settings by `PATCH /settings/display`.

The feature level 80 migration to merge these endpoints did not
change how request parameters are encoded. However, it did change
the handling of any invalid parameters present in a request
(see feature level 78 change below).

As of feature level 80, the `PATCH /settings/display` and
`PATCH /settings/notifications` endpoints are deprecated aliases
for this endpoint for backwards-compatibility, and will be removed
once clients have migrated to use this endpoint.

Prior to Zulip 5.0 (feature level 78), this endpoint indicated
which parameters it had processed by including in the response
object `"key": value` entries for values successfully changed by
the request. That was replaced by the more ergonomic
[`ignored_parameters_unsupported`][ignored-parameters] array.

The `PATCH /settings/notifications` and `PATCH /settings/display`
endpoints also had this behavior of indicating processed parameters
before they became aliases of this endpoint in Zulip 5.0 (see
feature level 80 change above).

Before feature level 78, request parameters that were not supported
(or were unchanged) were silently ignored.

[ignored-parameters]: /api/rest-error-handling#ignored-parameters

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UpdateSettingsRequest
*/
func (c *simpleClient) UpdateSettings(ctx context.Context) UpdateSettingsRequest {
	return UpdateSettingsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateSettingsExecute(r UpdateSettingsRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.fullName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "full_name", r.fullName, "", "")
	}
	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "email", r.email, "", "")
	}
	if r.oldPassword != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "old_password", r.oldPassword, "", "")
	}
	if r.newPassword != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "new_password", r.newPassword, "", "")
	}
	if r.twentyFourHourTime != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "twenty_four_hour_time", r.twentyFourHourTime, "form", "")
	}
	if r.webMarkReadOnScrollPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_mark_read_on_scroll_policy", r.webMarkReadOnScrollPolicy, "form", "")
	}
	if r.webChannelDefaultView != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_channel_default_view", r.webChannelDefaultView, "form", "")
	}
	if r.starredMessageCounts != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "starred_message_counts", r.starredMessageCounts, "form", "")
	}
	if r.receivesTypingNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "receives_typing_notifications", r.receivesTypingNotifications, "form", "")
	}
	if r.webSuggestUpdateTimezone != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_suggest_update_timezone", r.webSuggestUpdateTimezone, "form", "")
	}
	if r.fluidLayoutWidth != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "fluid_layout_width", r.fluidLayoutWidth, "form", "")
	}
	if r.highContrastMode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "high_contrast_mode", r.highContrastMode, "form", "")
	}
	if r.webFontSizePx != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_font_size_px", r.webFontSizePx, "form", "")
	}
	if r.webLineHeightPercent != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_line_height_percent", r.webLineHeightPercent, "form", "")
	}
	if r.colorScheme != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color_scheme", r.colorScheme, "form", "")
	}
	if r.enableDraftsSynchronization != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_drafts_synchronization", r.enableDraftsSynchronization, "form", "")
	}
	if r.translateEmoticons != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "translate_emoticons", r.translateEmoticons, "form", "")
	}
	if r.displayEmojiReactionUsers != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "display_emoji_reaction_users", r.displayEmojiReactionUsers, "form", "")
	}
	if r.defaultLanguage != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "default_language", r.defaultLanguage, "", "")
	}
	if r.webHomeView != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_home_view", r.webHomeView, "", "")
	}
	if r.webEscapeNavigatesToHomeView != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_escape_navigates_to_home_view", r.webEscapeNavigatesToHomeView, "form", "")
	}
	if r.leftSideUserlist != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "left_side_userlist", r.leftSideUserlist, "form", "")
	}
	if r.emojiset != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "emojiset", r.emojiset, "", "")
	}
	if r.demoteInactiveChannels != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "demote_inactive_streams", r.demoteInactiveChannels, "form", "")
	}
	if r.userListStyle != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_list_style", r.userListStyle, "form", "")
	}
	if r.webAnimateImagePreviews != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_animate_image_previews", r.webAnimateImagePreviews, "", "")
	}
	if r.webChannelUnreadsCountDisplayPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_stream_unreads_count_display_policy", r.webChannelUnreadsCountDisplayPolicy, "form", "")
	}
	if r.hideAiFeatures != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hide_ai_features", r.hideAiFeatures, "form", "")
	}
	if r.webLeftSidebarShowChannelFolders != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_left_sidebar_show_channel_folders", r.webLeftSidebarShowChannelFolders, "form", "")
	}
	if r.webLeftSidebarUnreadsCountSummary != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_left_sidebar_unreads_count_summary", r.webLeftSidebarUnreadsCountSummary, "form", "")
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "timezone", r.timezone, "", "")
	}
	if r.enableChannelDesktopNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_stream_desktop_notifications", r.enableChannelDesktopNotifications, "form", "")
	}
	if r.enableChannelEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_stream_email_notifications", r.enableChannelEmailNotifications, "form", "")
	}
	if r.enableChannelPushNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_stream_push_notifications", r.enableChannelPushNotifications, "form", "")
	}
	if r.enableChannelAudibleNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_stream_audible_notifications", r.enableChannelAudibleNotifications, "form", "")
	}
	if r.notificationSound != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "notification_sound", r.notificationSound, "", "")
	}
	if r.enableDesktopNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_desktop_notifications", r.enableDesktopNotifications, "form", "")
	}
	if r.enableSounds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_sounds", r.enableSounds, "form", "")
	}
	if r.emailNotificationsBatchingPeriodSeconds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "email_notifications_batching_period_seconds", r.emailNotificationsBatchingPeriodSeconds, "form", "")
	}
	if r.enableOfflineEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_offline_email_notifications", r.enableOfflineEmailNotifications, "form", "")
	}
	if r.enableOfflinePushNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_offline_push_notifications", r.enableOfflinePushNotifications, "form", "")
	}
	if r.enableOnlinePushNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_online_push_notifications", r.enableOnlinePushNotifications, "form", "")
	}
	if r.enableFollowedTopicDesktopNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_desktop_notifications", r.enableFollowedTopicDesktopNotifications, "form", "")
	}
	if r.enableFollowedTopicEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_email_notifications", r.enableFollowedTopicEmailNotifications, "form", "")
	}
	if r.enableFollowedTopicPushNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_push_notifications", r.enableFollowedTopicPushNotifications, "form", "")
	}
	if r.enableFollowedTopicAudibleNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_audible_notifications", r.enableFollowedTopicAudibleNotifications, "form", "")
	}
	if r.enableDigestEmails != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_digest_emails", r.enableDigestEmails, "form", "")
	}
	if r.enableMarketingEmails != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_marketing_emails", r.enableMarketingEmails, "form", "")
	}
	if r.enableLoginEmails != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_login_emails", r.enableLoginEmails, "form", "")
	}
	if r.messageContentInEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message_content_in_email_notifications", r.messageContentInEmailNotifications, "form", "")
	}
	if r.pmContentInDesktopNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pm_content_in_desktop_notifications", r.pmContentInDesktopNotifications, "form", "")
	}
	if r.wildcardMentionsNotify != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "wildcard_mentions_notify", r.wildcardMentionsNotify, "form", "")
	}
	if r.enableFollowedTopicWildcardMentionsNotify != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_wildcard_mentions_notify", r.enableFollowedTopicWildcardMentionsNotify, "form", "")
	}
	if r.desktopIconCountDisplay != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "desktop_icon_count_display", r.desktopIconCountDisplay, "form", "")
	}
	if r.realmNameInEmailNotificationsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "realm_name_in_email_notifications_policy", r.realmNameInEmailNotificationsPolicy, "form", "")
	}
	if r.automaticallyFollowTopicsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "automatically_follow_topics_policy", r.automaticallyFollowTopicsPolicy, "form", "")
	}
	if r.automaticallyUnmuteTopicsInMutedChannelsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "automatically_unmute_topics_in_muted_streams_policy", r.automaticallyUnmuteTopicsInMutedChannelsPolicy, "form", "")
	}
	if r.automaticallyFollowTopicsWhereMentioned != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "automatically_follow_topics_where_mentioned", r.automaticallyFollowTopicsWhereMentioned, "form", "")
	}
	if r.resolvedTopicNoticeAutoReadPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "resolved_topic_notice_auto_read_policy", r.resolvedTopicNoticeAutoReadPolicy, "", "")
	}
	if r.presenceEnabled != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "presence_enabled", r.presenceEnabled, "form", "")
	}
	if r.enterSends != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enter_sends", r.enterSends, "form", "")
	}
	if r.sendPrivateTypingNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_private_typing_notifications", r.sendPrivateTypingNotifications, "form", "")
	}
	if r.sendChannelTypingNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_stream_typing_notifications", r.sendChannelTypingNotifications, "form", "")
	}
	if r.sendReadReceipts != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_read_receipts", r.sendReadReceipts, "form", "")
	}
	if r.allowPrivateDataExport != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_private_data_export", r.allowPrivateDataExport, "form", "")
	}
	if r.emailAddressVisibility != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "email_address_visibility", r.emailAddressVisibility, "form", "")
	}
	if r.webNavigateToSentMessage != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_navigate_to_sent_message", r.webNavigateToSentMessage, "form", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateStatusRequest struct {
	ctx          context.Context
	ApiService   UsersAPI
	statusText   *string
	away         *bool
	emojiName    *string
	emojiCode    *string
	reactionType *string
}

// The text content of the status message. Sending the empty string will clear the user&#39;s status.  **Note**: The limit on the size of the message is 60 characters.
func (r UpdateStatusRequest) StatusText(statusText string) UpdateStatusRequest {
	r.statusText = &statusText
	return r
}

// Whether the user should be marked as \\\&quot;away\\\&quot;.  **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, &#x60;away&#x60; is a legacy way to access the user&#39;s &#x60;presence_enabled&#x60; setting, with &#x60;away &#x3D; !presence_enabled&#x60;. To be removed in a future release.
func (r UpdateStatusRequest) Away(away bool) UpdateStatusRequest {
	r.away = &away
	return r
}

// The name for the emoji to associate with this status.  **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusRequest) EmojiName(emojiName string) UpdateStatusRequest {
	r.emojiName = &emojiName
	return r
}

// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the &#x60;reaction_type&#x60;.  **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusRequest) EmojiCode(emojiCode string) UpdateStatusRequest {
	r.emojiCode = &emojiCode
	return r
}

// A string indicating the type of emoji. Each emoji &#x60;reaction_type&#x60; has an independent namespace for values of &#x60;emoji_code&#x60;.  Must be one of the following values:  - &#x60;unicode_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - &#x60;realm_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be the Id of   the uploaded [custom emoji](zulip.com/help/custom-emoji.  - &#x60;zulip_extra_emoji&#x60; : These are special emoji included with Zulip.   In this namespace, &#x60;emoji_code&#x60; will be the name of the emoji (e.g.   \\\&quot;zulip\\\&quot;).  **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusRequest) ReactionType(reactionType string) UpdateStatusRequest {
	r.reactionType = &reactionType
	return r
}

func (r UpdateStatusRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateStatusExecute(r)
}

/*
UpdateStatus Update your status

Change your [status](zulip.com/help/status-and-availability.

A request to this endpoint will only change the parameters passed.
For example, passing just `status_text` requests a change in the status
text, but will leave the status emoji unchanged.

Clients that wish to set the user's status to a specific value should
pass all supported parameters.

**Changes**: In Zulip 5.0 (feature level 86), added support for
`emoji_name`, `emoji_code`, and `reaction_type` parameters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UpdateStatusRequest
*/
func (c *simpleClient) UpdateStatus(ctx context.Context) UpdateStatusRequest {
	return UpdateStatusRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateStatusExecute(r UpdateStatusRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.statusText != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "status_text", r.statusText, "", "")
	}
	if r.away != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "away", r.away, "form", "")
	}
	if r.emojiName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "emoji_name", r.emojiName, "", "")
	}
	if r.emojiCode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "emoji_code", r.emojiCode, "", "")
	}
	if r.reactionType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "reaction_type", r.reactionType, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateStatusForUserRequest struct {
	ctx          context.Context
	ApiService   UsersAPI
	userId       int64
	statusText   *string
	emojiName    *string
	emojiCode    *string
	reactionType *string
}

// The text content of the status message. Sending the empty string will clear the user&#39;s status.  **Note**: The limit on the size of the message is 60 characters.
func (r UpdateStatusForUserRequest) StatusText(statusText string) UpdateStatusForUserRequest {
	r.statusText = &statusText
	return r
}

// The name for the emoji to associate with this status.  **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusForUserRequest) EmojiName(emojiName string) UpdateStatusForUserRequest {
	r.emojiName = &emojiName
	return r
}

// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the &#x60;reaction_type&#x60;.  **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusForUserRequest) EmojiCode(emojiCode string) UpdateStatusForUserRequest {
	r.emojiCode = &emojiCode
	return r
}

// A string indicating the type of emoji. Each emoji &#x60;reaction_type&#x60; has an independent namespace for values of &#x60;emoji_code&#x60;.  Must be one of the following values:  - &#x60;unicode_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - &#x60;realm_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be the Id of   the uploaded [custom emoji](zulip.com/help/custom-emoji.  - &#x60;zulip_extra_emoji&#x60; : These are special emoji included with Zulip.   In this namespace, &#x60;emoji_code&#x60; will be the name of the emoji (e.g.   \\\&quot;zulip\\\&quot;).  **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusForUserRequest) ReactionType(reactionType string) UpdateStatusForUserRequest {
	r.reactionType = &reactionType
	return r
}

func (r UpdateStatusForUserRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateStatusForUserExecute(r)
}

/*
UpdateStatusForUser Update user status

Administrator endpoint for changing the [status](zulip.com/help/status-and-availability) of
another user.

**Changes**: New in Zulip 11.0 (feature level 407).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The target user's Id.
	@return UpdateStatusForUserRequest
*/
func (c *simpleClient) UpdateStatusForUser(ctx context.Context, userId int64) UpdateStatusForUserRequest {
	return UpdateStatusForUserRequest{
		ApiService: c,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateStatusForUserExecute(r UpdateStatusForUserRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.statusText != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "status_text", r.statusText, "", "")
	}
	if r.emojiName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "emoji_name", r.emojiName, "", "")
	}
	if r.emojiCode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "emoji_code", r.emojiCode, "", "")
	}
	if r.reactionType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "reaction_type", r.reactionType, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateUserRequest struct {
	ctx         context.Context
	ApiService  UsersAPI
	userId      int64
	fullName    *string
	role        *Role
	profileData *[]map[string]interface{}
	newEmail    *string
}

// The user&#39;s full name.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 5.0 (feature level 106).
func (r UpdateUserRequest) FullName(fullName string) UpdateUserRequest {
	r.fullName = &fullName
	return r
}

// New [role](zulip.com/api/roles-and-permissions) for the user. Roles are encoded as:  - Organization owner: 100 - Organization administrator: 200 - Organization moderator: 300 - Member: 400 - Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.  **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of &#x60;is_admin&#x60; and &#x60;is_guest&#x60; boolean parameters. Organization moderator role added in Zulip 4.0 (feature level 60).
func (r UpdateUserRequest) Role(role Role) UpdateUserRequest {
	r.role = &role
	return r
}

// A dictionary containing the updated custom profile field data for the user.
func (r UpdateUserRequest) ProfileData(profileData []map[string]interface{}) UpdateUserRequest {
	r.profileData = &profileData
	return r
}

// New email address for the user. Requires the user making the request to be an organization owner and additionally have the &#x60;.can_change_user_emails&#x60; special permission.  **Changes**: New in Zulip 10.0 (feature level 285).
func (r UpdateUserRequest) NewEmail(newEmail string) UpdateUserRequest {
	r.newEmail = &newEmail
	return r
}

func (r UpdateUserRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateUserExecute(r)
}

/*
UpdateUser Update a user

Administrative endpoint to update the details of another user in the organization.

Supports everything an administrator can do to edit details of another
user's account, including editing full name,
[role](zulip.com/help/user-roles, and [custom profile
fields](zulip.com/help/custom-profile-fields.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The target user's Id.
	@return UpdateUserRequest
*/
func (c *simpleClient) UpdateUser(ctx context.Context, userId int64) UpdateUserRequest {
	return UpdateUserRequest{
		ApiService: c,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateUserExecute(r UpdateUserRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.fullName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "full_name", r.fullName, "", "")
	}
	if r.role != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "role", r.role, "form", "")
	}
	if r.profileData != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "profile_data", r.profileData, "form", "multi")
	}
	if r.newEmail != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "new_email", r.newEmail, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateUserByEmailRequest struct {
	ctx         context.Context
	ApiService  UsersAPI
	email       string
	fullName    *string
	role        *int32
	profileData *[]map[string]interface{}
	newEmail    *string
}

// The user&#39;s full name.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 5.0 (feature level 106).
func (r UpdateUserByEmailRequest) FullName(fullName string) UpdateUserByEmailRequest {
	r.fullName = &fullName
	return r
}

// New [role](zulip.com/api/roles-and-permissions) for the user. Roles are encoded as:  - Organization owner: 100 - Organization administrator: 200 - Organization moderator: 300 - Member: 400 - Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.  **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of &#x60;is_admin&#x60; and &#x60;is_guest&#x60; boolean parameters. Organization moderator role added in Zulip 4.0 (feature level 60).
func (r UpdateUserByEmailRequest) Role(role int32) UpdateUserByEmailRequest {
	r.role = &role
	return r
}

// A dictionary containing the updated custom profile field data for the user.
func (r UpdateUserByEmailRequest) ProfileData(profileData []map[string]interface{}) UpdateUserByEmailRequest {
	r.profileData = &profileData
	return r
}

// New email address for the user. Requires the user making the request to be an organization owner and additionally have the &#x60;.can_change_user_emails&#x60; special permission.  **Changes**: New in Zulip 10.0 (feature level 285).
func (r UpdateUserByEmailRequest) NewEmail(newEmail string) UpdateUserByEmailRequest {
	r.newEmail = &newEmail
	return r
}

func (r UpdateUserByEmailRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateUserByEmailExecute(r)
}

/*
UpdateUserByEmail Update a user by email

Administrative endpoint to update the details of another user in the organization by their email address.
Works the same way as [`PATCH /users/{user_id}`](zulip.com/api/update-user) but fetching the target user by their
real email address.

The requester needs to have permission to view the target user's real email address, subject to the
user's email address visibility setting. Otherwise, the dummy address of the format
`user{id}@{realm.host}` needs be used. This follows the same rules as `GET /users/{email}`.

**Changes**: New in Zulip 10.0 (feature level 313).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param email The email address of the user, specified following the same rules as [`GET /users/{email}`](zulip.com/api/get-user-by-email.
	@return UpdateUserByEmailRequest
*/
func (c *simpleClient) UpdateUserByEmail(ctx context.Context, email string) UpdateUserByEmailRequest {
	return UpdateUserByEmailRequest{
		ApiService: c,
		ctx:        ctx,
		email:      email,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateUserByEmailExecute(r UpdateUserByEmailRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.fullName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "full_name", r.fullName, "", "")
	}
	if r.role != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "role", r.role, "form", "")
	}
	if r.profileData != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "profile_data", r.profileData, "form", "multi")
	}
	if r.newEmail != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "new_email", r.newEmail, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateUserGroupRequest struct {
	ctx                   context.Context
	ApiService            UsersAPI
	userGroupId           int64
	name                  *string
	description           *string
	canAddMembersGroup    *GroupSettingValueUpdate
	canJoinGroup          *GroupSettingValueUpdate
	canLeaveGroup         *GroupSettingValueUpdate
	canManageGroup        *GroupSettingValueUpdate
	canMentionGroup       *GroupSettingValueUpdate
	canRemoveMembersGroup *GroupSettingValueUpdate
	deactivated           *bool
}

// The new name of the group.  **Changes**: Before Zulip 7.0 (feature level 165), this was a required field.
func (r UpdateUserGroupRequest) Name(name string) UpdateUserGroupRequest {
	r.name = &name
	return r
}

// The new description of the group.  **Changes**: Before Zulip 7.0 (feature level 165), this was a required field.
func (r UpdateUserGroupRequest) Description(description string) UpdateUserGroupRequest {
	r.description = &description
	return r
}

// The set of users who have permission to add members to this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups
func (r UpdateUserGroupRequest) CanAddMembersGroup(canAddMembersGroup GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canAddMembersGroup = &canAddMembersGroup
	return r
}

// The set of users who have permission to join this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 301).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups
func (r UpdateUserGroupRequest) CanJoinGroup(canJoinGroup GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canJoinGroup = &canJoinGroup
	return r
}

// The set of users who have permission to leave this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 308).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups
func (r UpdateUserGroupRequest) CanLeaveGroup(canLeaveGroup GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canLeaveGroup = &canLeaveGroup
	return r
}

// The set of users who have permission to [manage this user group][manage-user-groups] expressed as an [update to a group-setting value][update-group-setting].  This setting cannot be set to &#x60;\\\&quot;role:internet\\\&quot;&#x60; and &#x60;\\\&quot;role:everyone\\\&quot;&#x60; [system groups][system-groups].  **Changes**: New in Zulip 10.0 (feature level 283).  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups [manage-user-groups]: /help/manage-user-groups
func (r UpdateUserGroupRequest) CanManageGroup(canManageGroup GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canManageGroup = &canManageGroup
	return r
}

// The set of users who have permission to [mention this group][mentions], expressed as an [update to a group-setting value][update-group-setting].  This setting cannot be set to &#x60;\\\&quot;role:internet\\\&quot;&#x60; and &#x60;\\\&quot;role:owners\\\&quot;&#x60; [system groups][system-groups].  **Changes**: In Zulip 9.0 (feature level 260), this parameter was updated to only accept an object with the &#x60;old&#x60; and &#x60;new&#x60; fields described below. Prior to this feature level, this parameter could be either of the two forms of a [group-setting value][setting-values].  Before Zulip 9.0 (feature level 258), this parameter could only be the integer form of a [group-setting value][setting-values].  Before Zulip 8.0 (feature level 198), this parameter was named &#x60;can_mention_group_id&#x60;.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups][system-groups].  [mentions]: /help/mention-a-user-or-group [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups [setting-values]: /api/group-setting-values
func (r UpdateUserGroupRequest) CanMentionGroup(canMentionGroup GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canMentionGroup = &canMentionGroup
	return r
}

// The set of users who have permission to remove members from this user group expressed as an [update to a group-setting value][update-group-setting].  **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  [update-group-setting]: /api/group-setting-values#updating-group-setting-values [system-groups]: /api/group-setting-values#system-groups
func (r UpdateUserGroupRequest) CanRemoveMembersGroup(canRemoveMembersGroup GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canRemoveMembersGroup = &canRemoveMembersGroup
	return r
}

// A deactivated user group can be reactivated by passing this parameter as &#x60;false&#x60;.  Passing &#x60;true&#x60; does nothing as user group is deactivated using [&#x60;POST /user_groups/{user_group_id}/deactivate&#x60;](deactivate-user-group) endpoint.  **Changes**: New in Zulip 11.0 (feature level 386).
func (r UpdateUserGroupRequest) Deactivated(deactivated bool) UpdateUserGroupRequest {
	r.deactivated = &deactivated
	return r
}

func (r UpdateUserGroupRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateUserGroupExecute(r)
}

/*
UpdateUserGroup Update a user group

Update the name, description or any of the permission settings
of a [user group](zulip.com/help/user-groups.

This endpoint is also used to reactivate a user group.

Note that while permissions settings of deactivated groups can
be edited by this API endpoint, and those permissions settings
do affect the ability to modify the deactivated group and its
membership, the deactivated group itself cannot be mentioned
or used in the value of any permission without first being reactivated.

**Changes**: Starting with Zulip 11.0 (feature level 386), this
endpoint can be used to reactivate a user group.

Prior to Zulip 10.0 (feature level 340), only the name field
of deactivated groups could be modified.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userGroupId The Id of the target user group.
	@return UpdateUserGroupRequest
*/
func (c *simpleClient) UpdateUserGroup(ctx context.Context, userGroupId int64) UpdateUserGroupRequest {
	return UpdateUserGroupRequest{
		ApiService:  c,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateUserGroupExecute(r UpdateUserGroupRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups/{user_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_group_id"+"}", url.PathEscape(parameterValueToString(r.userGroupId, "userGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.canAddMembersGroup != nil {
		paramJson, err := parameterToJson(*r.canAddMembersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_add_members_group", paramJson)
	}
	if r.canJoinGroup != nil {
		paramJson, err := parameterToJson(*r.canJoinGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_join_group", paramJson)
	}
	if r.canLeaveGroup != nil {
		paramJson, err := parameterToJson(*r.canLeaveGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_leave_group", paramJson)
	}
	if r.canManageGroup != nil {
		paramJson, err := parameterToJson(*r.canManageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_manage_group", paramJson)
	}
	if r.canMentionGroup != nil {
		paramJson, err := parameterToJson(*r.canMentionGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_mention_group", paramJson)
	}
	if r.canRemoveMembersGroup != nil {
		paramJson, err := parameterToJson(*r.canRemoveMembersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_remove_members_group", paramJson)
	}
	if r.deactivated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "deactivated", r.deactivated, "form", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateUserGroupMembersRequest struct {
	ctx             context.Context
	ApiService      UsersAPI
	userGroupId     int64
	delete          *[]int64
	add             *[]int64
	deleteSubgroups *[]int64
	addSubgroups    *[]int64
}

// The list of user Ids to be removed from the user group.
func (r UpdateUserGroupMembersRequest) Delete(delete []int64) UpdateUserGroupMembersRequest {
	r.delete = &delete
	return r
}

// The list of user Ids to be added to the user group.
func (r UpdateUserGroupMembersRequest) Add(add []int64) UpdateUserGroupMembersRequest {
	r.add = &add
	return r
}

// The list of user group Ids to be removed from the user group.  **Changes**: New in Zulip 10.0 (feature level 311).
func (r UpdateUserGroupMembersRequest) DeleteSubgroups(deleteSubgroups []int64) UpdateUserGroupMembersRequest {
	r.deleteSubgroups = &deleteSubgroups
	return r
}

// The list of user group Ids to be added to the user group.  **Changes**: New in Zulip 10.0 (feature level 311).
func (r UpdateUserGroupMembersRequest) AddSubgroups(addSubgroups []int64) UpdateUserGroupMembersRequest {
	r.addSubgroups = &addSubgroups
	return r
}

func (r UpdateUserGroupMembersRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateUserGroupMembersExecute(r)
}

/*
UpdateUserGroupMembers Update user group members

Update the members of a [user group](zulip.com/help/user-groups. The
user Ids must correspond to non-deactivated users.

**Changes**: Prior to Zulip 11.0 (feature level 391), members
could not be added or removed from a deactivated group.

**Changes**: Prior to Zulip 10.0 (feature level 303), group memberships of
deactivated users were visible to the API and could be edited via this endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userGroupId The Id of the target user group.
	@return UpdateUserGroupMembersRequest
*/
func (c *simpleClient) UpdateUserGroupMembers(ctx context.Context, userGroupId int64) UpdateUserGroupMembersRequest {
	return UpdateUserGroupMembersRequest{
		ApiService:  c,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateUserGroupMembersExecute(r UpdateUserGroupMembersRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups/{user_group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"user_group_id"+"}", url.PathEscape(parameterValueToString(r.userGroupId, "userGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.delete != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "delete", r.delete, "form", "multi")
	}
	if r.add != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "add", r.add, "form", "multi")
	}
	if r.deleteSubgroups != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "delete_subgroups", r.deleteSubgroups, "form", "multi")
	}
	if r.addSubgroups != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "add_subgroups", r.addSubgroups, "form", "multi")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateUserGroupSubgroupsRequest struct {
	ctx         context.Context
	ApiService  UsersAPI
	userGroupId int64
	delete      *[]int64
	add         *[]int64
}

// The list of user group Ids to be removed from the user group.
func (r UpdateUserGroupSubgroupsRequest) Delete(delete []int64) UpdateUserGroupSubgroupsRequest {
	r.delete = &delete
	return r
}

// The list of user group Ids to be added to the user group.
func (r UpdateUserGroupSubgroupsRequest) Add(add []int64) UpdateUserGroupSubgroupsRequest {
	r.add = &add
	return r
}

func (r UpdateUserGroupSubgroupsRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateUserGroupSubgroupsExecute(r)
}

/*
UpdateUserGroupSubgroups Update subgroups of a user group

Update the subgroups of a [user group](zulip.com/help/user-groups.

**Changes**: Prior to Zulip 11.0 (feature level 391), subgroups
could not be added or removed from a deactivated group.

**Changes**: New in Zulip 6.0 (feature level 127).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userGroupId The Id of the target user group.
	@return UpdateUserGroupSubgroupsRequest
*/
func (c *simpleClient) UpdateUserGroupSubgroups(ctx context.Context, userGroupId int64) UpdateUserGroupSubgroupsRequest {
	return UpdateUserGroupSubgroupsRequest{
		ApiService:  c,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateUserGroupSubgroupsExecute(r UpdateUserGroupSubgroupsRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_groups/{user_group_id}/subgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"user_group_id"+"}", url.PathEscape(parameterValueToString(r.userGroupId, "userGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.delete != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "delete", r.delete, "form", "multi")
	}
	if r.add != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "add", r.add, "form", "multi")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}
