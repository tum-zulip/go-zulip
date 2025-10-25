package users

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/tum-zulip/go-zulip/zulip"
	. "github.com/tum-zulip/go-zulip/zulip/internal/apiutils"
)

type APIUsers interface {

	// AddAlertWords Add alert words
	//
	// Add words (or phrases) to the user's set of configured [alert words].
	//
	// [alert words]: https://zulip.com/help/dm-mention-alert-notifications#alert-words
	//
	AddAlertWords(ctx context.Context) AddAlertWordsRequest

	// AddAlertWordsExecute executes the request
	AddAlertWordsExecute(r AddAlertWordsRequest) (*AlertWordsResponse, *http.Response, error)

	// AddApnsToken Add an APNs device token
	//
	// This endpoint adds an APNs device token to register for iOS push notifications.
	//
	AddApnsToken(ctx context.Context) AddApnsTokenRequest

	// AddApnsTokenExecute executes the request
	AddApnsTokenExecute(r AddApnsTokenRequest) (*zulip.Response, *http.Response, error)

	// AddFcmToken Add an FCM registration token
	//
	// This endpoint adds an FCM registration token for push notifications.
	//
	AddFcmToken(ctx context.Context) AddFcmTokenRequest

	// AddFcmTokenExecute executes the request
	AddFcmTokenExecute(r AddFcmTokenRequest) (*zulip.Response, *http.Response, error)

	// CreateUser Create a user
	//
	// Create a new user account via the API.
	//
	// !!! warn ""
	//
	// *Note**: On Zulip Cloud, this feature is available only for
	// organizations on a [Zulip Cloud Standard]
	// or [Zulip Cloud Plus] plan. Administrators
	// can request the required `can_create_users` permission for a bot or
	// user by contacting [Zulip Cloud support] with an
	// explanation for why it is needed. Self-hosted installations can
	// toggle `can_create_users` on an account using the `manage.py
	// change_user_role` [management command].
	//
	// *Changes**: Before Zulip 4.0 (feature level 36), this endpoint was
	// available to all organization administrators.
	//
	// [Zulip Cloud support]: https://zulip.com/help/contact-support
	// [management command]: https://zulip.readthedocs.io/en/latest/production/management-commands.html
	//
	// [Zulip Cloud Standard]: https://zulip.com/plans/
	CreateUser(ctx context.Context) CreateUserRequest

	// CreateUserExecute executes the request
	CreateUserExecute(r CreateUserRequest) (*CreateUserResponse, *http.Response, error)

	// CreateUserGroup Create a user group
	//
	// Create a new [user group].
	//
	// [user group]: https://zulip.com/help/user-groups
	CreateUserGroup(ctx context.Context) CreateUserGroupRequest

	// CreateUserGroupExecute executes the request
	CreateUserGroupExecute(r CreateUserGroupRequest) (*CreateUserGroupResponse, *http.Response, error)

	// DeactivateOwnUser Deactivate own user
	//
	// Deactivates the current user's account. See also the administrative endpoint for
	// [deactivating another user].
	//
	// This endpoint is primarily useful to Zulip clients providing a user settings UI.
	//
	// [deactivating another user]: https://zulip.com/api/deactivate-user
	DeactivateOwnUser(ctx context.Context) DeactivateOwnUserRequest

	// DeactivateOwnUserExecute executes the request
	DeactivateOwnUserExecute(r DeactivateOwnUserRequest) (*zulip.Response, *http.Response, error)

	// DeactivateUser Deactivate a user
	//
	// [Deactivates a user]
	// given their user Id.
	//
	// [Deactivates a user]: https://zulip.com/help/deactivate-or-reactivate-a-user
	DeactivateUser(ctx context.Context, userId int64) DeactivateUserRequest

	// DeactivateUserExecute executes the request
	DeactivateUserExecute(r DeactivateUserRequest) (*zulip.Response, *http.Response, error)

	// DeactivateUserGroup Deactivate a user group
	//
	// Deactivate a user group. Deactivated user groups cannot be
	// used for mentions, permissions, or any other purpose, but can
	// be reactivated or renamed.
	//
	// Deactivating user groups is preferable to deleting them from
	// the database, since the deactivation model allows audit logs
	// of changes to sensitive group-valued permissions to be
	// maintained.
	//
	// *Changes**: New in Zulip 10.0 (feature level 290).
	//
	DeactivateUserGroup(ctx context.Context, userGroupId int64) DeactivateUserGroupRequest

	// DeactivateUserGroupExecute executes the request
	DeactivateUserGroupExecute(r DeactivateUserGroupRequest) (*zulip.Response, *http.Response, error)

	// GetAlertWords Get all alert words
	//
	// Get all of the user's configured [alert words].
	//
	// [alert words]: https://zulip.com/help/dm-mention-alert-notifications#alert-words
	//
	GetAlertWords(ctx context.Context) GetAlertWordsRequest

	// GetAlertWordsExecute executes the request
	GetAlertWordsExecute(r GetAlertWordsRequest) (*AlertWordsResponse, *http.Response, error)

	// GetAttachments Get attachments
	//
	// Fetch metadata on files uploaded by the requesting user.
	//
	GetAttachments(ctx context.Context) GetAttachmentsRequest

	// GetAttachmentsExecute executes the request
	GetAttachmentsExecute(r GetAttachmentsRequest) (*GetAttachmentsResponse, *http.Response, error)

	// GetIsUserGroupMember Get user group membership status
	//
	// Check whether a user is member of user group.
	//
	// *Changes**: Prior to Zulip 10.0 (feature level 303),
	// this would return true when passed a deactivated user
	// who was a member of the user group before being deactivated.
	//
	// New in Zulip 6.0 (feature level 127).
	//
	GetIsUserGroupMember(ctx context.Context, userGroupId int64, userId int64) GetIsUserGroupMemberRequest

	// GetIsUserGroupMemberExecute executes the request
	GetIsUserGroupMemberExecute(r GetIsUserGroupMemberRequest) (*GetIsUserGroupMemberResponse, *http.Response, error)

	// GetOwnUser Get own user
	//
	// Get basic data about the user/bot that requests this endpoint.
	//
	// *Changes**: Removed `is_billing_admin` field in Zulip 10.0 (feature level 363), as it was
	// replaced by the `can_manage_billing_group` realm setting.
	//
	GetOwnUser(ctx context.Context) GetOwnUserRequest

	// GetOwnUserExecute executes the request
	GetOwnUserExecute(r GetOwnUserRequest) (*GetOwnUserResponse, *http.Response, error)

	// GetUser Get a user
	//
	// Fetch details for a single user in the organization.
	//
	// You can also fetch details on [all users in the organization]
	// or [by a user's Zulip API email].
	//
	// *Changes**: New in Zulip 3.0 (feature level 1).
	//
	// [all users in the organization]: https://zulip.com/api/get-users
	// [by a user's Zulip API email]: https://zulip.com/api/get-user-by-email
	GetUser(ctx context.Context, userId int64) GetUserRequest

	// GetUserExecute executes the request
	GetUserExecute(r GetUserRequest) (*GetUserResponse, *http.Response, error)

	// GetUserByEmail Get a user by email
	//
	// Fetch details for a single user in the organization given a Zulip
	// API email address.
	//
	// You can also fetch details on [all users in the organization]
	// or [by user Id].
	//
	// Fetching by user Id is generally recommended when possible,
	// as a user might [change their email address]
	// or change their [email address visibility],
	// either of which could change the client's ability to look them up by that
	// email address.
	//
	// *Changes**: Starting with Zulip 10.0 (feature level 302), the real email
	// address can be used in the `email` parameter and will fetch the target user's
	// data if and only if the target's email visibility setting permits the requester
	// to see the email address.
	// The dummy email addresses of the form `user{id}@{realm.host}` still work, and
	// will now work for **all** users, via identifying them by the embedded user Id.
	//
	// New in Zulip Server 4.0 (feature level 39).
	//
	// [all users in the organization]: https://zulip.com/api/get-users
	// [by user Id]: https://zulip.com/api/get-user
	// [change their email address]: https://zulip.com/help/change-your-email-address
	// [email address visibility]: https://zulip.com/help/configure-email-visibility
	GetUserByEmail(ctx context.Context, email string) GetUserByEmailRequest

	// GetUserByEmailExecute executes the request
	GetUserByEmailExecute(r GetUserByEmailRequest) (*GetUserResponse, *http.Response, error)

	// GetUserGroupMembers Get user group members
	//
	// Get the members of a [user group].
	//
	// *Changes**: New in Zulip 6.0 (feature level 127).
	//
	// [user group]: https://zulip.com/help/user-groups
	GetUserGroupMembers(ctx context.Context, userGroupId int64) GetUserGroupMembersRequest

	// GetUserGroupMembersExecute executes the request
	GetUserGroupMembersExecute(r GetUserGroupMembersRequest) (*GetUserGroupMembersResponse, *http.Response, error)

	// GetUserGroupSubgroups Get subgroups of a user group
	//
	// Get the subgroups of a [user group].
	//
	// *Changes**: New in Zulip 6.0 (feature level 127).
	//
	// [user group]: https://zulip.com/help/user-groups
	GetUserGroupSubgroups(ctx context.Context, userGroupId int64) GetUserGroupSubgroupsRequest

	// GetUserGroupSubgroupsExecute executes the request
	GetUserGroupSubgroupsExecute(r GetUserGroupSubgroupsRequest) (*GetUserGroupSubgroupsResponse, *http.Response, error)

	// GetUserGroups Get user groups
	//
	// Fetches all of the user groups in the organization.
	//
	// !!! warn ""
	//
	// *Note**: This endpoint is only available to [members and administrators]; bots and guests
	// cannot use this endpoint.
	//
	// [members and administrators]: https://zulip.com/help/user-roles
	GetUserGroups(ctx context.Context) GetUserGroupsRequest

	// GetUserGroupsExecute executes the request
	GetUserGroupsExecute(r GetUserGroupsRequest) (*GetUserGroupsResponse, *http.Response, error)

	// GetUserPresence Get a user's presence
	//
	// Get the presence status for a specific user.
	//
	// This endpoint is most useful for embedding data about a user's
	// presence status in other sites (e.g. an employee directory). Full
	// Zulip clients like mobile/desktop apps will want to use the [main presence endpoint], which returns data for all
	// active users in the organization, instead.
	//
	// [main presence endpoint]: https://zulip.com/api/get-presence
	GetUserPresence(ctx context.Context, userIdOrEmail string) GetUserPresenceRequest

	// GetUserPresenceExecute executes the request
	GetUserPresenceExecute(r GetUserPresenceRequest) (*GetUserPresenceResponse, *http.Response, error)

	// GetUserStatus Get a user's status
	//
	// Get the [status] currently set by a
	// user in the organization.
	//
	// *Changes**: New in Zulip 9.0 (feature level 262). Previously,
	// user statuses could only be fetched via the [`POST /register`] endpoint.
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	// [status]: https://zulip.com/help/status-and-availability
	GetUserStatus(ctx context.Context, userId int64) GetUserStatusRequest

	// GetUserStatusExecute executes the request
	GetUserStatusExecute(r GetUserStatusRequest) (*GetUserStatusResponse, *http.Response, error)

	// GetUsers Get users
	//
	// Retrieve details on users in the organization.
	//
	// By default, returns all accessible users in the organization.
	// The `user_ids` query parameter can be used to limit the
	// results to a specific set of user Ids.
	//
	// Optionally includes values of [custom profile fields].
	//
	// You can also [fetch details on a single user].
	//
	// *Changes**: This endpoint did not support unauthenticated
	// access in organizations using the [public access option] prior to Zulip 11.0
	// (feature level 387).
	//
	// [custom profile fields]: https://zulip.com/help/custom-profile-fields
	// [fetch details on a single user]: https://zulip.com/api/get-user
	// [public access option]: https://zulip.com/help/public-access-option
	GetUsers(ctx context.Context) GetUsersRequest

	// GetUsersExecute executes the request
	GetUsersExecute(r GetUsersRequest) (*GetUsersResponse, *http.Response, error)

	// MuteUser Mute a user
	//
	// [Mute a user] from the perspective of the requesting
	// user. Messages sent by muted users will be automatically marked as read
	// and hidden for the user who muted them.
	//
	// Muted users should be implemented by clients as follows:
	//
	//   - The server will immediately mark all messages sent by the muted
	// user as read. This will automatically clear any existing mobile
	// push notifications related to the muted user.
	//   - The server will mark any new messages sent by the muted user as read
	// for the requesting user's account, which prevents all email and mobile
	// push notifications.
	//   - Clients should exclude muted users from presence lists or other UI
	// for viewing or composing one-on-one direct messages. One-on-one direct
	// messages sent by muted users should be hidden everywhere in the Zulip UI.
	//   - Channel messages and group direct messages sent by the muted
	// user should avoid displaying the content and name/avatar,
	// but should display that N messages by a muted user were
	// hidden (so that it is possible to interpret the messages by
	// other users who are talking with the muted user).
	//   - Group direct message conversations including the muted user
	// should display muted users as "Muted user", rather than
	// showing their name, in lists of such conversations, along with using
	// a blank grey avatar where avatars are displayed.
	//   - Administrative/settings UI elements for showing "All users that exist
	// on this channel or realm", e.g. for organization
	// administration or showing channel subscribers, should display
	// the user's name as normal.
	//
	// *Changes**: New in Zulip 4.0 (feature level 48).
	//
	// [Mute a user]: https://zulip.com/help/mute-a-user
	MuteUser(ctx context.Context, mutedUserId int64) MuteUserRequest

	// MuteUserExecute executes the request
	MuteUserExecute(r MuteUserRequest) (*zulip.Response, *http.Response, error)

	// ReactivateUser Reactivate a user
	//
	// [Reactivates a user]
	// given their user Id.
	//
	// [Reactivates a user]: https://zulip.com/help/deactivate-or-reactivate-a-user
	ReactivateUser(ctx context.Context, userId int64) ReactivateUserRequest

	// ReactivateUserExecute executes the request
	ReactivateUserExecute(r ReactivateUserRequest) (*zulip.Response, *http.Response, error)

	// RemoveAlertWords Remove alert words
	//
	// Remove words (or phrases) from the user's set of configured [alert words].
	//
	// Alert words are case insensitive.
	//
	// [alert words]: https://zulip.com/help/dm-mention-alert-notifications#alert-words
	//
	RemoveAlertWords(ctx context.Context) RemoveAlertWordsRequest

	// RemoveAlertWordsExecute executes the request
	RemoveAlertWordsExecute(r RemoveAlertWordsRequest) (*AlertWordsResponse, *http.Response, error)

	// RemoveApnsToken Remove an APNs device token
	//
	// This endpoint removes an APNs device token for iOS push notifications.
	//
	RemoveApnsToken(ctx context.Context) RemoveApnsTokenRequest

	// RemoveApnsTokenExecute executes the request
	RemoveApnsTokenExecute(r RemoveApnsTokenRequest) (*zulip.Response, *http.Response, error)

	// RemoveAttachment Delete an attachment
	//
	// Delete an uploaded file given its attachment Id.
	//
	// Note that uploaded files that have been referenced in at least
	// one message are automatically deleted once the last message
	// containing a link to them is deleted (whether directly or via
	// a [message retention policy].
	//
	// Uploaded files that are never used in a message are
	// automatically deleted a few weeks after being uploaded.
	//
	// Attachment Ids can be contained from [GET /attachments].
	//
	// [message retention policy]: https://zulip.com/help/message-retention-policy
	// [GET /attachments]: https://zulip.com/api/get-attachments
	RemoveAttachment(ctx context.Context, attachmentId int64) RemoveAttachmentRequest

	// RemoveAttachmentExecute executes the request
	RemoveAttachmentExecute(r RemoveAttachmentRequest) (*zulip.Response, *http.Response, error)

	// RemoveFcmToken Remove an FCM registration token
	//
	// This endpoint removes an FCM registration token for push notifications.
	//
	RemoveFcmToken(ctx context.Context) RemoveFcmTokenRequest

	// RemoveFcmTokenExecute executes the request
	RemoveFcmTokenExecute(r RemoveFcmTokenRequest) (*zulip.Response, *http.Response, error)

	// SetTypingStatus Set "typing" status
	//
	// Notify other users whether the current user is
	// [typing a message].
	//
	// Clients implementing Zulip's typing notifications
	// protocol should work as follows:
	//
	//   - Send a request to this endpoint with `"op": "start"` when a user
	// starts composing a message.
	//   - While the user continues to actively type or otherwise interact with
	// the compose UI (e.g. interacting with the compose box emoji picker),
	// send regular `"op": "start"` requests to this endpoint, using
	// `server_typing_started_wait_period_milliseconds` in the
	// [`POST /register`] response as the time interval
	// between each request.
	//   - Send a request to this endpoint with `"op": "stop"` when a user
	// has stopped using the compose UI for the time period indicated by
	// `server_typing_stopped_wait_period_milliseconds` in the
	// [`POST /register`] response or when a user
	// cancels the compose action (if it had previously sent a "start"
	// notification for that compose action).
	//   - Start displaying a visual typing indicator for a given conversation
	// when a [`typing op:start`] event is received
	// from the server.
	//   - Continue displaying a visual typing indicator for the conversation
	// until a [`typing op:stop`] event is received
	// from the server or the time period indicated by
	// `server_typing_started_expiry_period_milliseconds` in the
	// [`POST /register`] response has passed without
	// a new `typing "op": "start"` event for the conversation.
	//
	// This protocol is designed to allow the server-side typing notifications
	// implementation to be stateless while being resilient as network failures
	// will not result in a user being incorrectly displayed as perpetually
	// typing.
	//
	// See the subsystems documentation on [typing indicators]
	// for additional design details on Zulip's typing notifications protocol.
	//
	// *Changes**: Clients shouldn't care about the APIs prior to Zulip 8.0 (feature level 215)
	// for channel typing notifications, as no client actually implemented
	// the previous API for those.
	//
	// Support for displaying channel typing notifications was new
	// in Zulip 4.0 (feature level 58). Clients should indicate they support
	// processing channel typing notifications via the `stream_typing_notifications`
	// value in the [`client_capabilities`] parameter of the
	// `POST /register` endpoint.
	//
	// [typing a message]: https://zulip.com/help/typing-notifications
	// [`POST /register`]: https://zulip.com/api/register-queue
	// [`typing op:start`]: https://zulip.com/api/get-events#typing-start
	// [`typing op:stop`]: https://zulip.com/api/get-events#typing-stop
	// [`client_capabilities`]: https://zulip.com/api/register-queue#parameter-client_capabilities
	// [typing indicators]: https://zulip.readthedocs.io/en/latest/subsystems/typing-indicators.html
	//
	SetTypingStatus(ctx context.Context) SetTypingStatusRequest

	// SetTypingStatusExecute executes the request
	SetTypingStatusExecute(r SetTypingStatusRequest) (*zulip.Response, *http.Response, error)

	// SetTypingStatusForMessageEdit Set "typing" status for message editing
	//
	// Notify other users whether the current user is editing a message.
	//
	// Typing notifications for editing messages follow the same protocol as
	// [set-typing-status], see that endpoint for
	// details.
	//
	// *Changes**: Before Zulip 10.0 (feature level 361), the endpoint was
	// named `/message_edit_typing` with `message_id` a required parameter in
	// the request body. Clients are recommended to start using sending these
	// typing notifications starting from this feature level.
	//
	// New in Zulip 10.0 (feature level 351). Previously, typing notifications were
	// not available when editing messages.
	//
	// [set-typing-status]: https://zulip.com/api/set-typing-status
	SetTypingStatusForMessageEdit(ctx context.Context, messageId int64) SetTypingStatusForMessageEditRequest

	// SetTypingStatusForMessageEditExecute executes the request
	SetTypingStatusForMessageEditExecute(r SetTypingStatusForMessageEditRequest) (*zulip.Response, *http.Response, error)

	// UnmuteUser Unmute a user
	//
	// [Unmute a user]
	// from the perspective of the requesting user.
	//
	// *Changes**: New in Zulip 4.0 (feature level 48).
	//
	// [Unmute a user]: https://zulip.com/help/mute-a-user#see-your-list-of-muted-users
	UnmuteUser(ctx context.Context, mutedUserId int64) UnmuteUserRequest

	// UnmuteUserExecute executes the request
	UnmuteUserExecute(r UnmuteUserRequest) (*zulip.Response, *http.Response, error)

	// UpdatePresence Update your presence
	//
	// Update the current user's [presence] and fetch presence data
	// of other users in the organization.
	//
	// This endpoint is meant to be used by clients for both:
	//
	//   - Reporting the current user's presence status (`"active"` or `"idle"`) to the server.
	//   - Obtaining the presence data of all other users in the organization via regular polling.
	//
	// Accurate user presence is one of the most expensive parts of any
	// chat application (in terms of bandwidth and other resources). Therefore,
	// it is important that clients implementing Zulip's user presence system
	// use the modern [`last_update_id`] protocol to
	// minimize fetching duplicate user presence data.
	//
	// Client apps implementing presence are recommended to also consume [`presence` events], in order to learn about newly online users
	// immediately.
	//
	// The Zulip server is responsible for implementing [invisible mode],
	// which disables sharing a user's presence data. Nonetheless, clients
	// should check the `presence_enabled` field in user objects in order to
	// display the current user as online or offline based on whether they are
	// sharing their presence information.
	//
	// *Changes**: As of Zulip 8.0 (feature level 228), if the
	// `CAN_ACCESS_ALL_USERS_GROUP_LIMITS_PRESENCE` server-level setting is
	// `true`, then user presence data in the response is [limited to users the current user can see/access].
	//
	// [`presence` events]: https://zulip.com/api/get-events#presence
	// [limited to users the current user can see/access]: https://zulip.com/help/guest-users#configure-whether-guests-can-see-all-other-users
	// [invisible mode]: https://zulip.com/help/status-and-availability#invisible-mode
	// [presence]: https://zulip.com/help/status-and-availability#availability
	//
	// [`last_update_id`]: https://zulip.com/api/update-presence#parameter-last_update_id
	UpdatePresence(ctx context.Context) UpdatePresenceRequest

	// UpdatePresenceExecute executes the request
	UpdatePresenceExecute(r UpdatePresenceRequest) (*UpdatePresenceResponse, *http.Response, error)

	// UpdateSettings Update settings
	//
	// This endpoint is used to edit the current user's settings.
	//
	// *Changes**: Removed `dense_mode` setting in Zulip 10.0
	// (feature level 364) as we now have `web_font_size_px` and
	// `web_line_height_percent` settings for more control.
	//
	// Prior to Zulip 5.0 (feature level 80), this endpoint only
	// supported the `full_name`, `email`, `old_password`, and
	// `new_password` parameters. Notification settings were
	// managed by `PATCH /settings/notifications`, and all other
	// settings by `PATCH /settings/display`.
	//
	// The feature level 80 migration to merge these endpoints did not
	// change how request parameters are encoded. However, it did change
	// the handling of any invalid parameters present in a request
	// (see feature level 78 change below).
	//
	// As of feature level 80, the `PATCH /settings/display` and
	// `PATCH /settings/notifications` endpoints are deprecated aliases
	// for this endpoint for backwards-compatibility, and will be removed
	// once clients have migrated to use this endpoint.
	//
	// Prior to Zulip 5.0 (feature level 78), this endpoint indicated
	// which parameters it had processed by including in the response
	// object `"key": value` entries for values successfully changed by
	// the request. That was replaced by the more ergonomic
	// [`ignored_parameters_unsupported`] array.
	//
	// The `PATCH /settings/notifications` and `PATCH /settings/display`
	// endpoints also had this behavior of indicating processed parameters
	// before they became aliases of this endpoint in Zulip 5.0 (see
	// feature level 80 change above).
	//
	// Before feature level 78, request parameters that were not supported
	// (or were unchanged) were silently ignored.
	//
	// [`ignored_parameters_unsupported`]: https://zulip.com/api/rest-error-handling#ignored-parameters
	//
	UpdateSettings(ctx context.Context) UpdateSettingsRequest

	// UpdateSettingsExecute executes the request
	UpdateSettingsExecute(r UpdateSettingsRequest) (*zulip.Response, *http.Response, error)

	// UpdateStatus Update your status
	//
	// Change your [status].
	//
	// A request to this endpoint will only change the parameters passed.
	// For example, passing just `status_text` requests a change in the status
	// text, but will leave the status emoji unchanged.
	//
	// Clients that wish to set the user's status to a specific value should
	// pass all supported parameters.
	//
	// *Changes**: In Zulip 5.0 (feature level 86), added support for
	// `emoji_name`, `emoji_code`, and `reaction_type` parameters.
	//
	// [status]: https://zulip.com/help/status-and-availability
	UpdateStatus(ctx context.Context) UpdateStatusRequest

	// UpdateStatusExecute executes the request
	UpdateStatusExecute(r UpdateStatusRequest) (*zulip.Response, *http.Response, error)

	// UpdateStatusForUser Update user status
	//
	// Administrator endpoint for changing the [status] of
	// another user.
	//
	// *Changes**: New in Zulip 11.0 (feature level 407).
	//
	// [status]: https://zulip.com/help/status-and-availability
	UpdateStatusForUser(ctx context.Context, userId int64) UpdateStatusForUserRequest

	// UpdateStatusForUserExecute executes the request
	UpdateStatusForUserExecute(r UpdateStatusForUserRequest) (*zulip.Response, *http.Response, error)

	// UpdateUser Update a user
	//
	// Administrative endpoint to update the details of another user in the organization.
	//
	// Supports everything an administrator can do to edit details of another
	// user's account, including editing full name,
	// [role], and [custom profile fields].
	//
	// [role]: https://zulip.com/help/user-roles
	// [custom profile fields]: https://zulip.com/help/custom-profile-fields
	UpdateUser(ctx context.Context, userId int64) UpdateUserRequest

	// UpdateUserExecute executes the request
	UpdateUserExecute(r UpdateUserRequest) (*zulip.Response, *http.Response, error)

	// UpdateUserByEmail Update a user by email
	//
	// Administrative endpoint to update the details of another user in the organization by their email address.
	// Works the same way as [`PATCH /users/{user_id}`] but fetching the target user by their
	// real email address.
	//
	// The requester needs to have permission to view the target user's real email address, subject to the
	// user's email address visibility setting. Otherwise, the dummy address of the format
	// `user{id}@{realm.host}` needs be used. This follows the same rules as `GET /users/{email}`.
	//
	// *Changes**: New in Zulip 10.0 (feature level 313).
	//
	// [`PATCH /users/{user_id}`]: https://zulip.com/api/update-user
	UpdateUserByEmail(ctx context.Context, email string) UpdateUserByEmailRequest

	// UpdateUserByEmailExecute executes the request
	UpdateUserByEmailExecute(r UpdateUserByEmailRequest) (*zulip.Response, *http.Response, error)

	// UpdateUserGroup Update a user group
	//
	// Update the name, description or any of the permission settings
	// of a [user group].
	//
	// This endpoint is also used to reactivate a user group.
	//
	// Note that while permissions settings of deactivated groups can
	// be edited by this API endpoint, and those permissions settings
	// do affect the ability to modify the deactivated group and its
	// membership, the deactivated group itself cannot be mentioned
	// or used in the value of any permission without first being reactivated.
	//
	// *Changes**: Starting with Zulip 11.0 (feature level 386), this
	// endpoint can be used to reactivate a user group.
	//
	// Prior to Zulip 10.0 (feature level 340), only the name field
	// of deactivated groups could be modified.
	//
	// [user group]: https://zulip.com/help/user-groups
	UpdateUserGroup(ctx context.Context, userGroupId int64) UpdateUserGroupRequest

	// UpdateUserGroupExecute executes the request
	UpdateUserGroupExecute(r UpdateUserGroupRequest) (*zulip.Response, *http.Response, error)

	// UpdateUserGroupMembers Update user group members
	//
	// Update the members of a [user group]. The
	// user Ids must correspond to non-deactivated users.
	//
	// *Changes**: Prior to Zulip 11.0 (feature level 391), members
	// could not be added or removed from a deactivated group.
	//
	// *Changes**: Prior to Zulip 10.0 (feature level 303), group memberships of
	// deactivated users were visible to the API and could be edited via this endpoint.
	//
	// [user group]: https://zulip.com/help/user-groups
	UpdateUserGroupMembers(ctx context.Context, userGroupId int64) UpdateUserGroupMembersRequest

	// UpdateUserGroupMembersExecute executes the request
	UpdateUserGroupMembersExecute(r UpdateUserGroupMembersRequest) (*zulip.Response, *http.Response, error)

	// UpdateUserGroupSubgroups Update subgroups of a user group
	//
	// Update the subgroups of a [user group].
	//
	// *Changes**: Prior to Zulip 11.0 (feature level 391), subgroups
	// could not be added or removed from a deactivated group.
	//
	// *Changes**: New in Zulip 6.0 (feature level 127).
	//
	// [user group]: https://zulip.com/help/user-groups
	UpdateUserGroupSubgroups(ctx context.Context, userGroupId int64) UpdateUserGroupSubgroupsRequest

	// UpdateUserGroupSubgroupsExecute executes the request
	UpdateUserGroupSubgroupsExecute(r UpdateUserGroupSubgroupsRequest) (*zulip.Response, *http.Response, error)
}

type usersService struct {
	client StructuredClient
}

func NewUsersService(client StructuredClient) *usersService {
	return &usersService{client: client}
}

var _ APIUsers = (*usersService)(nil)

type AddAlertWordsRequest struct {
	ctx        context.Context
	apiService APIUsers
	alertWords *[]string
}

// An array of strings to be added to the user's set of configured alert words. Strings already present in the user's set of alert words already are ignored.  Alert words are case insensitive.
func (r AddAlertWordsRequest) AlertWords(alertWords []string) AddAlertWordsRequest {
	r.alertWords = &alertWords
	return r
}

func (r AddAlertWordsRequest) Execute() (*AlertWordsResponse, *http.Response, error) {
	return r.apiService.AddAlertWordsExecute(r)
}

// AddAlertWords Add alert words
//
// Add words (or phrases) to the user's set of configured [alert words].
//
// [alert words]: https://zulip.com/help/dm-mention-alert-notifications#alert-words
func (s *usersService) AddAlertWords(ctx context.Context) AddAlertWordsRequest {
	return AddAlertWordsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) AddAlertWordsExecute(r AddAlertWordsRequest) (*AlertWordsResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &AlertWordsResponse{}
		endpoint = "/users/me/alert_words"
	)
	if r.alertWords == nil {
		return nil, nil, fmt.Errorf("alertWords is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	if err := AddOptionalJSONParam(form, "alert_words", r.alertWords); err != nil {
		return nil, nil, err
	}
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type AddApnsTokenRequest struct {
	ctx        context.Context
	apiService APIUsers
	token      *string
	appid      *string
}

// The token provided by the device.
func (r AddApnsTokenRequest) Token(token string) AddApnsTokenRequest {
	r.token = &token
	return r
}

// The Id of the Zulip app that is making the request.
//
// **Changes**: In Zulip 8.0 (feature level 223), this parameter was made required. Previously, if it was unspecified, the server would use a default value (based on the `ZULIP_IOS_APP_Id` server setting, which defaulted to `"org.zulip.Zulip"`).
func (r AddApnsTokenRequest) Appid(appid string) AddApnsTokenRequest {
	r.appid = &appid
	return r
}

func (r AddApnsTokenRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.AddApnsTokenExecute(r)
}

// AddApnsToken Add an APNs device token
//
// This endpoint adds an APNs device token to register for iOS push notifications.
func (s *usersService) AddApnsToken(ctx context.Context) AddApnsTokenRequest {
	return AddApnsTokenRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) AddApnsTokenExecute(r AddApnsTokenRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/apns_device_token"
	)
	if r.token == nil {
		return nil, nil, fmt.Errorf("token is required and must be specified")
	}
	if r.appid == nil {
		return nil, nil, fmt.Errorf("appid is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "token", r.token)
	AddParam(form, "appid", r.appid)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type AddFcmTokenRequest struct {
	ctx        context.Context
	apiService APIUsers
	token      *string
}

// The token provided by the device.
func (r AddFcmTokenRequest) Token(token string) AddFcmTokenRequest {
	r.token = &token
	return r
}

func (r AddFcmTokenRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.AddFcmTokenExecute(r)
}

// AddFcmToken Add an FCM registration token
//
// This endpoint adds an FCM registration token for push notifications.
func (s *usersService) AddFcmToken(ctx context.Context) AddFcmTokenRequest {
	return AddFcmTokenRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) AddFcmTokenExecute(r AddFcmTokenRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/android_gcm_reg_id"
	)
	if r.token == nil {
		return nil, nil, fmt.Errorf("token is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "token", r.token)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type CreateUserRequest struct {
	ctx        context.Context
	apiService APIUsers
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
	return r.apiService.CreateUserExecute(r)
}

// CreateUser Create a user
//
// Create a new user account via the API.
//
// !!! warn ""
//
// *Note**: On Zulip Cloud, this feature is available only for
// organizations on a [Zulip Cloud Standard]
// or [Zulip Cloud Plus] plan. Administrators
// can request the required `can_create_users` permission for a bot or
// user by contacting [Zulip Cloud support] with an
// explanation for why it is needed. Self-hosted installations can
// toggle `can_create_users` on an account using the `manage.py
// change_user_role` [management command].
//
// *Changes**: Before Zulip 4.0 (feature level 36), this endpoint was
// available to all organization administrators.
//
// [Zulip Cloud support]: https://zulip.com/help/contact-support
// [management command]: https://zulip.readthedocs.io/en/latest/production/management-commands.html
// [Zulip Cloud Standard]: https://zulip.com/plans/
func (s *usersService) CreateUser(ctx context.Context) CreateUserRequest {
	return CreateUserRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) CreateUserExecute(r CreateUserRequest) (*CreateUserResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateUserResponse{}
		endpoint = "/users"
	)
	if r.email == nil {
		return nil, nil, fmt.Errorf("email is required and must be specified")
	}
	if r.password == nil {
		return nil, nil, fmt.Errorf("password is required and must be specified")
	}
	if r.fullName == nil {
		return nil, nil, fmt.Errorf("fullName is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "email", r.email)
	AddParam(form, "password", r.password)
	AddParam(form, "full_name", r.fullName)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type CreateUserGroupRequest struct {
	ctx                   context.Context
	apiService            APIUsers
	name                  *string
	description           *string
	members               *[]int64
	subgroups             *[]int64
	canAddMembersGroup    *zulip.GroupSettingValue
	canJoinGroup          *zulip.GroupSettingValue
	canLeaveGroup         *zulip.GroupSettingValue
	canManageGroup        *zulip.GroupSettingValue
	canMentionGroup       *zulip.GroupSettingValue
	canRemoveMembersGroup *zulip.GroupSettingValue
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

// An array containing the Ids of the initial subgroups for the new user group.  User can add subgroups to the new group irrespective of other permissions for the new group.
//
// **Changes**: New in Zulip 10.0 (feature level 311).
func (r CreateUserGroupRequest) Subgroups(subgroups []int64) CreateUserGroupRequest {
	r.subgroups = &subgroups
	return r
}

// A [group-setting value] defining the set of users who have permission to add members to this user group.
//
// **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the `can_manage_group` setting.
//
// [group-setting value]: https://zulip.com/api/group-setting-values
//
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
func (r CreateUserGroupRequest) CanAddMembersGroup(canAddMembersGroup zulip.GroupSettingValue) CreateUserGroupRequest {
	r.canAddMembersGroup = &canAddMembersGroup
	return r
}

// A [group-setting value] defining the set of users who have permission to join this user group.
//
// **Changes**: New in Zulip 10.0 (feature level 301).
//
// [group-setting value]: https://zulip.com/api/group-setting-values
//
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
func (r CreateUserGroupRequest) CanJoinGroup(canJoinGroup zulip.GroupSettingValue) CreateUserGroupRequest {
	r.canJoinGroup = &canJoinGroup
	return r
}

// A [group-setting value] defining the set of users who have permission to leave this user group.
//
// **Changes**: New in Zulip 10.0 (feature level 308).
//
// [group-setting value]: https://zulip.com/api/group-setting-values
//
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
func (r CreateUserGroupRequest) CanLeaveGroup(canLeaveGroup zulip.GroupSettingValue) CreateUserGroupRequest {
	r.canLeaveGroup = &canLeaveGroup
	return r
}

// A [group-setting value] defining the set of users who have permission to [manage this user group].  This setting cannot be set to `"role:internet"` and `"role:everyone"` [system groups].
//
// **Changes**: New in Zulip 10.0 (feature level 283).
//
// [group-setting value]: https://zulip.com/api/group-setting-values
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
// [manage this user group]: https://zulip.com/help/manage-user-groups
func (r CreateUserGroupRequest) CanManageGroup(canManageGroup zulip.GroupSettingValue) CreateUserGroupRequest {
	r.canManageGroup = &canManageGroup
	return r
}

// A [group-setting value] defining the set of users who have permission to [mention this user group].  This setting cannot be set to `"role:internet"` and `"role:owners"` [system groups].  Before Zulip 9.0 (feature level 258), this parameter could only be the integer form of a [group-setting value].  Before Zulip 8.0 (feature level 198), this parameter was named `can_mention_group_id`.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups].
//
// [group-setting value]: https://zulip.com/api/group-setting-values
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
// [mention this user group]: https://zulip.com/help/mention-a-user-or-group
func (r CreateUserGroupRequest) CanMentionGroup(canMentionGroup zulip.GroupSettingValue) CreateUserGroupRequest {
	r.canMentionGroup = &canMentionGroup
	return r
}

// A [group-setting value] defining the set of users who have permission to remove members from this user group.
//
// **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the `can_manage_group` setting.
//
// [group-setting value]: https://zulip.com/api/group-setting-values
//
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
func (r CreateUserGroupRequest) CanRemoveMembersGroup(canRemoveMembersGroup zulip.GroupSettingValue) CreateUserGroupRequest {
	r.canRemoveMembersGroup = &canRemoveMembersGroup
	return r
}

func (r CreateUserGroupRequest) Execute() (*CreateUserGroupResponse, *http.Response, error) {
	return r.apiService.CreateUserGroupExecute(r)
}

// CreateUserGroup Create a user group
//
// Create a new [user group].
//
// [user group]: https://zulip.com/help/user-groups
func (s *usersService) CreateUserGroup(ctx context.Context) CreateUserGroupRequest {
	return CreateUserGroupRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) CreateUserGroupExecute(r CreateUserGroupRequest) (*CreateUserGroupResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateUserGroupResponse{}
		endpoint = "/user_groups/create"
	)
	if r.name == nil {
		return nil, nil, fmt.Errorf("name is required and must be specified")
	}
	if r.description == nil {
		return nil, nil, fmt.Errorf("description is required and must be specified")
	}
	if r.members == nil {
		return nil, nil, fmt.Errorf("members is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "name", r.name)
	AddParam(form, "description", r.description)
	AddParam(form, "members", r.members)
	AddOptionalParam(form, "subgroups", r.subgroups)
	if err := AddOptionalJSONParam(form, "can_add_members_group", r.canAddMembersGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_join_group", r.canJoinGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_leave_group", r.canLeaveGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_manage_group", r.canManageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_mention_group", r.canMentionGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_remove_members_group", r.canRemoveMembersGroup); err != nil {
		return nil, nil, err
	}
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type DeactivateOwnUserRequest struct {
	ctx        context.Context
	apiService APIUsers
}

func (r DeactivateOwnUserRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.DeactivateOwnUserExecute(r)
}

// DeactivateOwnUser Deactivate own user
//
// Deactivates the current user's account. See also the administrative endpoint for
// [deactivating another user].
//
// This endpoint is primarily useful to Zulip clients providing a user settings UI.
//
// [deactivating another user]: https://zulip.com/api/deactivate-user
func (s *usersService) DeactivateOwnUser(ctx context.Context) DeactivateOwnUserRequest {
	return DeactivateOwnUserRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) DeactivateOwnUserExecute(r DeactivateOwnUserRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me"
	)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type DeactivateUserRequest struct {
	ctx                             context.Context
	apiService                      APIUsers
	userId                          int64
	deactivationNotificationComment *string
}

// If not `null`, requests that the deactivated user receive a notification email about their account deactivation.  If not `""`, encodes custom text written by the administrator to be included in the notification email.
//
// **Changes**: New in Zulip 5.0 (feature level 135).
func (r DeactivateUserRequest) DeactivationNotificationComment(deactivationNotificationComment string) DeactivateUserRequest {
	r.deactivationNotificationComment = &deactivationNotificationComment
	return r
}

func (r DeactivateUserRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.DeactivateUserExecute(r)
}

// DeactivateUser Deactivate a user
//
// [Deactivates a user]
// given their user Id.
//
// [Deactivates a user]: https://zulip.com/help/deactivate-or-reactivate-a-user
func (s *usersService) DeactivateUser(ctx context.Context, userId int64) DeactivateUserRequest {
	return DeactivateUserRequest{
		apiService: s,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
func (s *usersService) DeactivateUserExecute(r DeactivateUserRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/{user_id}"
	)

	endpoint = strings.Replace(endpoint, "{user_id}", IdToString(r.userId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "deactivation_notification_comment", r.deactivationNotificationComment)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type DeactivateUserGroupRequest struct {
	ctx         context.Context
	apiService  APIUsers
	userGroupId int64
}

func (r DeactivateUserGroupRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.DeactivateUserGroupExecute(r)
}

// DeactivateUserGroup Deactivate a user group
//
// Deactivate a user group. Deactivated user groups cannot be
// used for mentions, permissions, or any other purpose, but can
// be reactivated or renamed.
//
// Deactivating user groups is preferable to deleting them from
// the database, since the deactivation model allows audit logs
// of changes to sensitive group-valued permissions to be
// maintained.
//
// *Changes**: New in Zulip 10.0 (feature level 290).
func (s *usersService) DeactivateUserGroup(ctx context.Context, userGroupId int64) DeactivateUserGroupRequest {
	return DeactivateUserGroupRequest{
		apiService:  s,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
func (s *usersService) DeactivateUserGroupExecute(r DeactivateUserGroupRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/user_groups/{user_group_id}/deactivate"
	)

	endpoint = strings.Replace(endpoint, "{user_group_id}", IdToString(r.userGroupId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetAlertWordsRequest struct {
	ctx        context.Context
	apiService APIUsers
}

func (r GetAlertWordsRequest) Execute() (*AlertWordsResponse, *http.Response, error) {
	return r.apiService.GetAlertWordsExecute(r)
}

// GetAlertWords Get all alert words
//
// Get all of the user's configured [alert words].
//
// [alert words]: https://zulip.com/help/dm-mention-alert-notifications#alert-words
func (s *usersService) GetAlertWords(ctx context.Context) GetAlertWordsRequest {
	return GetAlertWordsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) GetAlertWordsExecute(r GetAlertWordsRequest) (*AlertWordsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &AlertWordsResponse{}
		endpoint = "/users/me/alert_words"
	)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetAttachmentsRequest struct {
	ctx        context.Context
	apiService APIUsers
}

func (r GetAttachmentsRequest) Execute() (*GetAttachmentsResponse, *http.Response, error) {
	return r.apiService.GetAttachmentsExecute(r)
}

// GetAttachments Get attachments
//
// Fetch metadata on files uploaded by the requesting user.
func (s *usersService) GetAttachments(ctx context.Context) GetAttachmentsRequest {
	return GetAttachmentsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) GetAttachmentsExecute(r GetAttachmentsRequest) (*GetAttachmentsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetAttachmentsResponse{}
		endpoint = "/attachments"
	)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetIsUserGroupMemberRequest struct {
	ctx              context.Context
	apiService       APIUsers
	userGroupId      int64
	userId           int64
	directMemberOnly *bool
}

// Whether to consider only the direct members of user group and not members of its subgroups. Default is `false`.
func (r GetIsUserGroupMemberRequest) DirectMemberOnly(directMemberOnly bool) GetIsUserGroupMemberRequest {
	r.directMemberOnly = &directMemberOnly
	return r
}

func (r GetIsUserGroupMemberRequest) Execute() (*GetIsUserGroupMemberResponse, *http.Response, error) {
	return r.apiService.GetIsUserGroupMemberExecute(r)
}

// GetIsUserGroupMember Get user group membership status
//
// Check whether a user is member of user group.
//
// *Changes**: Prior to Zulip 10.0 (feature level 303),
// this would return true when passed a deactivated user
// who was a member of the user group before being deactivated.
//
// New in Zulip 6.0 (feature level 127).
func (s *usersService) GetIsUserGroupMember(ctx context.Context, userGroupId int64, userId int64) GetIsUserGroupMemberRequest {
	return GetIsUserGroupMemberRequest{
		apiService:  s,
		ctx:         ctx,
		userGroupId: userGroupId,
		userId:      userId,
	}
}

// Execute executes the request
func (s *usersService) GetIsUserGroupMemberExecute(r GetIsUserGroupMemberRequest) (*GetIsUserGroupMemberResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetIsUserGroupMemberResponse{}
		endpoint = "/user_groups/{user_group_id}/members/{user_id}"
	)

	endpoint = strings.Replace(endpoint, "{user_group_id}", IdToString(r.userGroupId), -1)
	endpoint = strings.Replace(endpoint, "{user_id}", IdToString(r.userId), -1)

	AddOptionalParam(query, "direct_member_only", r.directMemberOnly)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetOwnUserRequest struct {
	ctx        context.Context
	apiService APIUsers
}

func (r GetOwnUserRequest) Execute() (*GetOwnUserResponse, *http.Response, error) {
	return r.apiService.GetOwnUserExecute(r)
}

// GetOwnUser Get own user
//
// Get basic data about the user/bot that requests this endpoint.
//
// *Changes**: Removed `is_billing_admin` field in Zulip 10.0 (feature level 363), as it was
// replaced by the `can_manage_billing_group` realm setting.
func (s *usersService) GetOwnUser(ctx context.Context) GetOwnUserRequest {
	return GetOwnUserRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) GetOwnUserExecute(r GetOwnUserRequest) (*GetOwnUserResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetOwnUserResponse{}
		endpoint = "/users/me"
	)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetUserRequest struct {
	ctx                        context.Context
	apiService                 APIUsers
	userId                     int64
	clientGravatar             *bool
	includeCustomProfileFields *bool
}

// Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.
//
// **Changes**: The default value of this parameter was `false` prior to Zulip 5.0 (feature level 92).
func (r GetUserRequest) ClientGravatar(clientGravatar bool) GetUserRequest {
	r.clientGravatar = &clientGravatar
	return r
}

// Whether the client wants [custom profile field] data to be included in the response.
//
// **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.
//
// [custom profile field]: https://zulip.com/help/custom-profile-fields
func (r GetUserRequest) IncludeCustomProfileFields(includeCustomProfileFields bool) GetUserRequest {
	r.includeCustomProfileFields = &includeCustomProfileFields
	return r
}

func (r GetUserRequest) Execute() (*GetUserResponse, *http.Response, error) {
	return r.apiService.GetUserExecute(r)
}

// GetUser Get a user
//
// Fetch details for a single user in the organization.
//
// You can also fetch details on [all users in the organization]
// or [by a user's Zulip API email].
//
// *Changes**: New in Zulip 3.0 (feature level 1).
//
// [all users in the organization]: https://zulip.com/api/get-users
// [by a user's Zulip API email]: https://zulip.com/api/get-user-by-email
func (s *usersService) GetUser(ctx context.Context, userId int64) GetUserRequest {
	return GetUserRequest{
		apiService: s,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
func (s *usersService) GetUserExecute(r GetUserRequest) (*GetUserResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetUserResponse{}
		endpoint = "/users/{user_id}"
	)

	endpoint = strings.Replace(endpoint, "{user_id}", IdToString(r.userId), -1)

	AddOptionalParam(query, "client_gravatar", r.clientGravatar)
	AddOptionalParam(query, "include_custom_profile_fields", r.includeCustomProfileFields)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetUserByEmailRequest struct {
	ctx                        context.Context
	apiService                 APIUsers
	email                      string
	clientGravatar             *bool
	includeCustomProfileFields *bool
}

// Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.
//
// **Changes**: The default value of this parameter was `false` prior to Zulip 5.0 (feature level 92).
func (r GetUserByEmailRequest) ClientGravatar(clientGravatar bool) GetUserByEmailRequest {
	r.clientGravatar = &clientGravatar
	return r
}

// Whether the client wants [custom profile field] data to be included in the response.
//
// **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.
//
// [custom profile field]: https://zulip.com/help/custom-profile-fields
func (r GetUserByEmailRequest) IncludeCustomProfileFields(includeCustomProfileFields bool) GetUserByEmailRequest {
	r.includeCustomProfileFields = &includeCustomProfileFields
	return r
}

func (r GetUserByEmailRequest) Execute() (*GetUserResponse, *http.Response, error) {
	return r.apiService.GetUserByEmailExecute(r)
}

// GetUserByEmail Get a user by email
//
// Fetch details for a single user in the organization given a Zulip
// API email address.
//
// You can also fetch details on [all users in the organization]
// or [by user Id].
//
// Fetching by user Id is generally recommended when possible,
// as a user might [change their email address]
// or change their [email address visibility],
// either of which could change the client's ability to look them up by that
// email address.
//
// *Changes**: Starting with Zulip 10.0 (feature level 302), the real email
// address can be used in the `email` parameter and will fetch the target user's
// data if and only if the target's email visibility setting permits the requester
// to see the email address.
// The dummy email addresses of the form `user{id}@{realm.host}` still work, and
// will now work for **all** users, via identifying them by the embedded user Id.
//
// New in Zulip Server 4.0 (feature level 39).
//
// [all users in the organization]: https://zulip.com/api/get-users
// [by user Id]: https://zulip.com/api/get-user
// [change their email address]: https://zulip.com/help/change-your-email-address
// [email address visibility]: https://zulip.com/help/configure-email-visibility
func (s *usersService) GetUserByEmail(ctx context.Context, email string) GetUserByEmailRequest {
	return GetUserByEmailRequest{
		apiService: s,
		ctx:        ctx,
		email:      email,
	}
}

// Execute executes the request
func (s *usersService) GetUserByEmailExecute(r GetUserByEmailRequest) (*GetUserResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetUserResponse{}
		endpoint = "/users/{email}"
	)

	endpoint = strings.Replace(endpoint, "{email}", url.PathEscape(r.email), -1)

	AddOptionalParam(query, "client_gravatar", r.clientGravatar)
	AddOptionalParam(query, "include_custom_profile_fields", r.includeCustomProfileFields)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetUserGroupMembersRequest struct {
	ctx              context.Context
	apiService       APIUsers
	userGroupId      int64
	directMemberOnly *bool
}

// Whether to consider only the direct members of user group and not members of its subgroups. Default is `false`.
func (r GetUserGroupMembersRequest) DirectMemberOnly(directMemberOnly bool) GetUserGroupMembersRequest {
	r.directMemberOnly = &directMemberOnly
	return r
}

func (r GetUserGroupMembersRequest) Execute() (*GetUserGroupMembersResponse, *http.Response, error) {
	return r.apiService.GetUserGroupMembersExecute(r)
}

// GetUserGroupMembers Get user group members
//
// Get the members of a [user group].
//
// *Changes**: New in Zulip 6.0 (feature level 127).
//
// [user group]: https://zulip.com/help/user-groups
func (s *usersService) GetUserGroupMembers(ctx context.Context, userGroupId int64) GetUserGroupMembersRequest {
	return GetUserGroupMembersRequest{
		apiService:  s,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
func (s *usersService) GetUserGroupMembersExecute(r GetUserGroupMembersRequest) (*GetUserGroupMembersResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetUserGroupMembersResponse{}
		endpoint = "/user_groups/{user_group_id}/members"
	)
	endpoint = strings.Replace(endpoint, "{user_group_id}", IdToString(r.userGroupId), -1)

	AddOptionalParam(query, "direct_member_only", r.directMemberOnly)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetUserGroupSubgroupsRequest struct {
	ctx                context.Context
	apiService         APIUsers
	userGroupId        int64
	directSubgroupOnly *bool
}

// Whether to consider only direct subgroups of the user group or subgroups of subgroups also.
func (r GetUserGroupSubgroupsRequest) DirectSubgroupOnly(directSubgroupOnly bool) GetUserGroupSubgroupsRequest {
	r.directSubgroupOnly = &directSubgroupOnly
	return r
}

func (r GetUserGroupSubgroupsRequest) Execute() (*GetUserGroupSubgroupsResponse, *http.Response, error) {
	return r.apiService.GetUserGroupSubgroupsExecute(r)
}

// GetUserGroupSubgroups Get subgroups of a user group
//
// Get the subgroups of a [user group].
//
// *Changes**: New in Zulip 6.0 (feature level 127).
//
// [user group]: https://zulip.com/help/user-groups
func (s *usersService) GetUserGroupSubgroups(ctx context.Context, userGroupId int64) GetUserGroupSubgroupsRequest {
	return GetUserGroupSubgroupsRequest{
		apiService:  s,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
func (s *usersService) GetUserGroupSubgroupsExecute(r GetUserGroupSubgroupsRequest) (*GetUserGroupSubgroupsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetUserGroupSubgroupsResponse{}
		endpoint = "/user_groups/{user_group_id}/subgroups"
	)

	endpoint = strings.Replace(endpoint, "{user_group_id}", IdToString(r.userGroupId), -1)

	AddOptionalParam(query, "direct_subgroup_only", r.directSubgroupOnly)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetUserGroupsRequest struct {
	ctx                      context.Context
	apiService               APIUsers
	includeDeactivatedGroups *bool
}

// Whether to include deactivated user groups in the response.
//
// **Changes**: In Zulip 10.0 (feature level 294), renamed `allow_deactivated` to `include_deactivated_groups`.  New in Zulip 10.0 (feature level 290). Previously, deactivated user groups did not exist and thus would never be included in the response.
func (r GetUserGroupsRequest) IncludeDeactivatedGroups(includeDeactivatedGroups bool) GetUserGroupsRequest {
	r.includeDeactivatedGroups = &includeDeactivatedGroups
	return r
}

func (r GetUserGroupsRequest) Execute() (*GetUserGroupsResponse, *http.Response, error) {
	return r.apiService.GetUserGroupsExecute(r)
}

// GetUserGroups Get user groups
//
// Fetches all of the user groups in the organization.
//
// !!! warn ""
//
// *Note**: This endpoint is only available to [members and administrators]; bots and guests
// cannot use this endpoint.
//
// [members and administrators]: https://zulip.com/help/user-roles
func (s *usersService) GetUserGroups(ctx context.Context) GetUserGroupsRequest {
	return GetUserGroupsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) GetUserGroupsExecute(r GetUserGroupsRequest) (*GetUserGroupsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetUserGroupsResponse{}
		endpoint = "/user_groups"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "include_deactivated_groups", r.includeDeactivatedGroups)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetUserPresenceRequest struct {
	ctx           context.Context
	apiService    APIUsers
	userIdOrEmail string
}

func (r GetUserPresenceRequest) Execute() (*GetUserPresenceResponse, *http.Response, error) {
	return r.apiService.GetUserPresenceExecute(r)
}

// GetUserPresence Get a user's presence
//
// Get the presence status for a specific user.
//
// This endpoint is most useful for embedding data about a user's
// presence status in other sites (e.g. an employee directory). Full
// Zulip clients like mobile/desktop apps will want to use the [main presence endpoint], which returns data for all
// active users in the organization, instead.
//
// [main presence endpoint]: https://zulip.com/api/get-presence
func (s *usersService) GetUserPresence(ctx context.Context, userIdOrEmail string) GetUserPresenceRequest {
	return GetUserPresenceRequest{
		apiService:    s,
		ctx:           ctx,
		userIdOrEmail: userIdOrEmail,
	}
}

// Execute executes the request
func (s *usersService) GetUserPresenceExecute(r GetUserPresenceRequest) (*GetUserPresenceResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetUserPresenceResponse{}
		endpoint = "/users/{user_id_or_email}/presence"
	)

	endpoint = strings.Replace(endpoint, "{user_id_or_email}", url.PathEscape(r.userIdOrEmail), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetUserStatusRequest struct {
	ctx        context.Context
	apiService APIUsers
	userId     int64
}

func (r GetUserStatusRequest) Execute() (*GetUserStatusResponse, *http.Response, error) {
	return r.apiService.GetUserStatusExecute(r)
}

// GetUserStatus Get a user's status
//
// Get the [status] currently set by a
// user in the organization.
//
// *Changes**: New in Zulip 9.0 (feature level 262). Previously,
// user statuses could only be fetched via the [`POST /register`] endpoint.
// [status]: https://zulip.com/help/status-and-availability
// [`POST /register`]: https://zulip.com/api/register-queue
func (s *usersService) GetUserStatus(ctx context.Context, userId int64) GetUserStatusRequest {
	return GetUserStatusRequest{
		apiService: s,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
func (s *usersService) GetUserStatusExecute(r GetUserStatusRequest) (*GetUserStatusResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetUserStatusResponse{}
		endpoint = "/users/{user_id}/status"
	)

	endpoint = strings.Replace(endpoint, "{user_id}", IdToString(r.userId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetUsersRequest struct {
	ctx                        context.Context
	apiService                 APIUsers
	clientGravatar             *bool
	includeCustomProfileFields *bool
	userIds                    *[]int64
}

// Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.
//
// **Changes**: The default value of this parameter was `false` prior to Zulip 5.0 (feature level 92).
func (r GetUsersRequest) ClientGravatar(clientGravatar bool) GetUsersRequest {
	r.clientGravatar = &clientGravatar
	return r
}

// Whether the client wants [custom profile field] data to be included in the response.
//
// **Changes**: New in Zulip 2.1.0. Previous versions do not offer these data via the API.
//
// [custom profile field]: https://zulip.com/help/custom-profile-fields
func (r GetUsersRequest) IncludeCustomProfileFields(includeCustomProfileFields bool) GetUsersRequest {
	r.includeCustomProfileFields = &includeCustomProfileFields
	return r
}

// Limits the results to the specified user Ids. If not provided, the server will return all accessible users in the organization.
//
// **Changes**: New in Zulip 11.0 (feature level 384).
func (r GetUsersRequest) UserIds(userIds []int64) GetUsersRequest {
	r.userIds = &userIds
	return r
}

func (r GetUsersRequest) Execute() (*GetUsersResponse, *http.Response, error) {
	return r.apiService.GetUsersExecute(r)
}

// GetUsers Get users
//
// Retrieve details on users in the organization.
//
// By default, returns all accessible users in the organization.
// The `user_ids` query parameter can be used to limit the
// results to a specific set of user Ids.
//
// Optionally includes values of [custom profile fields].
//
// You can also [fetch details on a single user].
//
// *Changes**: This endpoint did not support unauthenticated
// access in organizations using the [public access option] prior to Zulip 11.0
// (feature level 387).
//
// [custom profile fields]: https://zulip.com/help/custom-profile-fields
// [fetch details on a single user]: https://zulip.com/api/get-user
// [public access option]: https://zulip.com/help/public-access-option
func (s *usersService) GetUsers(ctx context.Context) GetUsersRequest {
	return GetUsersRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) GetUsersExecute(r GetUsersRequest) (*GetUsersResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetUsersResponse{}
		endpoint = "/users"
	)
	AddOptionalParam(query, "client_gravatar", r.clientGravatar)
	AddOptionalParam(query, "include_custom_profile_fields", r.includeCustomProfileFields)
	AddOptionalCSVParam(query, "user_ids", r.userIds)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type MuteUserRequest struct {
	ctx         context.Context
	apiService  APIUsers
	mutedUserId int64
}

func (r MuteUserRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.MuteUserExecute(r)
}

// MuteUser Mute a user
//
// [Mute a user] from the perspective of the requesting
// user. Messages sent by muted users will be automatically marked as read
// and hidden for the user who muted them.
//
// Muted users should be implemented by clients as follows:
//
//   - The server will immediately mark all messages sent by the muted user as read. This will automatically clear any existing mobile push notifications related to the muted user.
//   - The server will mark any new messages sent by the muted user as read for the requesting user's account, which prevents all email and mobile push notifications.
//   - Clients should exclude muted users from presence lists or other UI for viewing or composing one-on-one direct messages. One-on-one direct messages sent by muted users should be hidden everywhere in the Zulip UI.
//   - Channel messages and group direct messages sent by the muted user should avoid displaying the content and name/avatar, but should display that N messages by a muted user were hidden (so that it is possible to interpret the messages by other users who are talking with the muted user).
//   - Group direct message conversations including the muted user should display muted users as "Muted user", rather than showing their name, in lists of such conversations, along with using a blank grey avatar where avatars are displayed.
//   - Administrative/settings UI elements for showing "All users that exist on this channel or realm", e.g. for organization administration or showing channel subscribers, should display the user's name as normal.
//
// *Changes**: New in Zulip 4.0 (feature level 48).
//
// [Mute a user]: https://zulip.com/help/mute-a-user
func (s *usersService) MuteUser(ctx context.Context, mutedUserId int64) MuteUserRequest {
	return MuteUserRequest{
		apiService:  s,
		ctx:         ctx,
		mutedUserId: mutedUserId,
	}
}

// Execute executes the request
func (s *usersService) MuteUserExecute(r MuteUserRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/muted_users/{muted_user_id}"
	)

	endpoint = strings.Replace(endpoint, "{muted_user_id}", IdToString(r.mutedUserId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type ReactivateUserRequest struct {
	ctx        context.Context
	apiService APIUsers
	userId     int64
}

func (r ReactivateUserRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.ReactivateUserExecute(r)
}

// ReactivateUser Reactivate a user
//
// [Reactivates a user]
// given their user Id.
//
// [Reactivates a user]: https://zulip.com/help/deactivate-or-reactivate-a-user
func (s *usersService) ReactivateUser(ctx context.Context, userId int64) ReactivateUserRequest {
	return ReactivateUserRequest{
		apiService: s,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
func (s *usersService) ReactivateUserExecute(r ReactivateUserRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/{user_id}/reactivate"
	)

	endpoint = strings.Replace(endpoint, "{user_id}", IdToString(r.userId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RemoveAlertWordsRequest struct {
	ctx        context.Context
	apiService APIUsers
	alertWords *[]string
}

// An array of strings to be removed from the user's set of configured alert words. Strings that are not in the user's set of alert words are ignored.
func (r RemoveAlertWordsRequest) AlertWords(alertWords []string) RemoveAlertWordsRequest {
	r.alertWords = &alertWords
	return r
}

func (r RemoveAlertWordsRequest) Execute() (*AlertWordsResponse, *http.Response, error) {
	return r.apiService.RemoveAlertWordsExecute(r)
}

// RemoveAlertWords Remove alert words
//
// Remove words (or phrases) from the user's set of configured [alert words].
//
// Alert words are case insensitive.
//
// [alert words]: https://zulip.com/help/dm-mention-alert-notifications#alert-words
func (s *usersService) RemoveAlertWords(ctx context.Context) RemoveAlertWordsRequest {
	return RemoveAlertWordsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) RemoveAlertWordsExecute(r RemoveAlertWordsRequest) (*AlertWordsResponse, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &AlertWordsResponse{}
		endpoint = "/users/me/alert_words"
	)
	if r.alertWords == nil {
		return nil, nil, fmt.Errorf("alertWords is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "alert_words", r.alertWords)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RemoveApnsTokenRequest struct {
	ctx        context.Context
	apiService APIUsers
	token      *string
}

// The token provided by the device.
func (r RemoveApnsTokenRequest) Token(token string) RemoveApnsTokenRequest {
	r.token = &token
	return r
}

func (r RemoveApnsTokenRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.RemoveApnsTokenExecute(r)
}

// RemoveApnsToken Remove an APNs device token
//
// This endpoint removes an APNs device token for iOS push notifications.
func (s *usersService) RemoveApnsToken(ctx context.Context) RemoveApnsTokenRequest {
	return RemoveApnsTokenRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) RemoveApnsTokenExecute(r RemoveApnsTokenRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/apns_device_token"
	)
	if r.token == nil {
		return nil, nil, fmt.Errorf("token is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "token", r.token)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RemoveAttachmentRequest struct {
	ctx          context.Context
	apiService   APIUsers
	attachmentId int64
}

func (r RemoveAttachmentRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.RemoveAttachmentExecute(r)
}

// RemoveAttachment Delete an attachment
//
// Delete an uploaded file given its attachment Id.
//
// Note that uploaded files that have been referenced in at least
// one message are automatically deleted once the last message
// containing a link to them is deleted (whether directly or via
// a [message retention policy].
//
// Uploaded files that are never used in a message are
// automatically deleted a few weeks after being uploaded.
//
// Attachment Ids can be contained from [GET /attachments].
//
// [message retention policy]: https://zulip.com/help/message-retention-policy
// [GET /attachments]: https://zulip.com/api/get-attachments
func (s *usersService) RemoveAttachment(ctx context.Context, attachmentId int64) RemoveAttachmentRequest {
	return RemoveAttachmentRequest{
		apiService:   s,
		ctx:          ctx,
		attachmentId: attachmentId,
	}
}

// Execute executes the request
func (s *usersService) RemoveAttachmentExecute(r RemoveAttachmentRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/attachments/{attachment_id}"
	)

	endpoint = strings.Replace(endpoint, "{attachment_id}", IdToString(r.attachmentId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RemoveFcmTokenRequest struct {
	ctx        context.Context
	apiService APIUsers
	token      *string
}

// The token provided by the device.
func (r RemoveFcmTokenRequest) Token(token string) RemoveFcmTokenRequest {
	r.token = &token
	return r
}

func (r RemoveFcmTokenRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.RemoveFcmTokenExecute(r)
}

// RemoveFcmToken Remove an FCM registration token
//
// This endpoint removes an FCM registration token for push notifications.
func (s *usersService) RemoveFcmToken(ctx context.Context) RemoveFcmTokenRequest {
	return RemoveFcmTokenRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) RemoveFcmTokenExecute(r RemoveFcmTokenRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/android_gcm_reg_id"
	)
	if r.token == nil {
		return nil, nil, fmt.Errorf("token is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "token", r.token)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type SetTypingStatusRequest struct {
	ctx           context.Context
	apiService    APIUsers
	op            *zulip.TypingStatusOp
	recipientType *zulip.RecipientType
	to            *zulip.Recipient
	channelId     *int64
	topic         *string
}

// Whether the user has started (`TypingStatusOpStart`) or stopped (`TypingStatusOpStop`) typing.
func (r SetTypingStatusRequest) Op(op zulip.TypingStatusOp) SetTypingStatusRequest {
	r.op = &op
	return r
}

// Type of the message being composed.
//
// **Changes**: In Zulip 9.0 (feature level 248), `RecipientTypeChannel` was added as an additional value for this parameter to indicate a channel message is being composed.  In Zulip 8.0 (feature level 215), stopped supporting `RecipientTypePrivate` as a valid value for this parameter.  In Zulip 7.0 (feature level 174), `RecipientTypeDirect` was added as the preferred way to indicate a direct message is being composed, becoming the default value for this parameter and deprecating the original `RecipientTypePrivate`.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.
func (r SetTypingStatusRequest) RecipientType(recipientType zulip.RecipientType) SetTypingStatusRequest {
	r.recipientType = &recipientType
	return r
}

// User Ids of the recipients of the message being typed. Required for the `"direct"` type. Ignored in the case of `"stream"` or `"channel"` type.  Clients should send a JSON-encoded list of user Ids, even if there is only one recipient.
//
// **Changes**: In Zulip 8.0 (feature level 215), stopped using this parameter for the `"stream"` type. Previously, in the case of the `"stream"` type, it accepted a single-element list containing the Id of the channel. A new parameter, `stream_id`, is now used for this. Note that the `"channel"` type did not exist at this feature level.  Support for typing notifications for channel' messages is new in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  Before Zulip 2.0.0, this parameter accepted only a JSON-encoded list of email addresses. Support for the email address-based format was removed in Zulip 3.0 (feature level 11).
func (r SetTypingStatusRequest) To(to zulip.Recipient) SetTypingStatusRequest {
	r.to = &to
	if r.recipientType == nil {
		r.recipientType = to.RecipientType()
	}
	return r
}

// Id of the channel in which the message is being typed. Required for the `"stream"` or `"channel"` type. Ignored in the case of `"direct"` type.
//
// **Changes**: New in Zulip 8.0 (feature level 215). Previously, a single-element list containing the Id of the channel was passed in `to` parameter.
func (r SetTypingStatusRequest) ChannelId(channelId int64) SetTypingStatusRequest {
	r.channelId = &channelId
	return r
}

// Topic to which message is being typed. Required for the `"stream"` or `"channel"` type. Ignored in the case of `"direct"` type.  Note: When `"(no topic)"` or the value of `realm_empty_topic_display_name` found in the [POST /register] response is used for this parameter, it is interpreted as an empty string.
//
// **Changes**: Before Zulip 10.0 (feature level 372), `"(no topic)"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.
//
// [POST /register]: https://zulip.com/api/register-queue
func (r SetTypingStatusRequest) Topic(topic string) SetTypingStatusRequest {
	r.topic = &topic
	return r
}

func (r SetTypingStatusRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.SetTypingStatusExecute(r)
}

// SetTypingStatus Set "typing" status
//
// Notify other users whether the current user is
// [typing a message].
//
// Clients implementing Zulip's typing notifications
// protocol should work as follows:
//
//   - Send a request to this endpoint with `"op": "start"` when a user starts composing a message.
//   - While the user continues to actively type or otherwise interact with the compose UI (e.g. interacting with the compose box emoji picker), send regular `"op": "start"` requests to this endpoint, using `server_typing_started_wait_period_milliseconds` in the [`POST /register`] response as the time interval between each request.
//   - Send a request to this endpoint with `"op": "stop"` when a user has stopped using the compose UI for the time period indicated by `server_typing_stopped_wait_period_milliseconds` in the [`POST /register`] response or when a user cancels the compose action (if it had previously sent a "start" notification for that compose action).
//   - Start displaying a visual typing indicator for a given conversation when a [`typing op:start`] event is received from the server.
//   - Continue displaying a visual typing indicator for the conversation until a [`typing op:stop`] event is received from the server or the time period indicated by `server_typing_started_expiry_period_milliseconds` in the [`POST /register`] response has passed without a new `typing "op": "start"` event for the conversation.
//
// This protocol is designed to allow the server-side typing notifications
// implementation to be stateless while being resilient as network failures
// will not result in a user being incorrectly displayed as perpetually
// typing.
//
// See the subsystems documentation on [typing indicators]
// for additional design details on Zulip's typing notifications protocol.
//
// *Changes**: Clients shouldn't care about the APIs prior to Zulip 8.0 (feature level 215)
// for channel typing notifications, as no client actually implemented
// the previous API for those.
//
// Support for displaying channel typing notifications was new
// in Zulip 4.0 (feature level 58). Clients should indicate they support
// processing channel typing notifications via the `stream_typing_notifications`
// value in the [`client_capabilities`] parameter of the
// `POST /register` endpoint.
//
// [typing a message]: https://zulip.com/help/typing-notifications
// [`typing op:start`]: https://zulip.com/api/get-events#typing-start
// [`typing op:stop`]: https://zulip.com/api/get-events#typing-stop
// [`client_capabilities`]: https://zulip.com/api/register-queue#parameter-client_capabilities
// [typing indicators]: https://zulip.readthedocs.io/en/latest/subsystems/typing-indicators.html
// [`POST /register`]: https://zulip.com/api/register-queue
func (s *usersService) SetTypingStatus(ctx context.Context) SetTypingStatusRequest {
	return SetTypingStatusRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) SetTypingStatusExecute(r SetTypingStatusRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/typing"
	)
	if r.op == nil {
		return nil, nil, fmt.Errorf("op is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "type", r.recipientType)
	AddParam(form, "op", r.op)
	if err := AddOptionalJSONParam(form, "to", r.to); err != nil {
		return nil, nil, err
	}
	AddOptionalParam(form, "stream_id", r.channelId)
	AddOptionalParam(form, "topic", r.topic)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type SetTypingStatusForMessageEditRequest struct {
	ctx        context.Context
	apiService APIUsers
	messageId  int64
	op         *zulip.TypingStatusOp
}

// Whether the user has started (`TypingStatusOpStart`) or stopped (`TypingStatusOpStop`) editing.
func (r SetTypingStatusForMessageEditRequest) Op(op zulip.TypingStatusOp) SetTypingStatusForMessageEditRequest {
	r.op = &op
	return r
}

func (r SetTypingStatusForMessageEditRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.SetTypingStatusForMessageEditExecute(r)
}

// SetTypingStatusForMessageEdit Set "typing" status for message editing
//
// Notify other users whether the current user is editing a message.
//
// Typing notifications for editing messages follow the same protocol as
// [set-typing-status], see that endpoint for
// details.
//
// *Changes**: Before Zulip 10.0 (feature level 361), the endpoint was
// named `/message_edit_typing` with `message_id` a required parameter in
// the request body. Clients are recommended to start using sending these
// typing notifications starting from this feature level.
//
// New in Zulip 10.0 (feature level 351). Previously, typing notifications were
// not available when editing messages.
//
// [set-typing-status]: https://zulip.com/api/set-typing-status
func (s *usersService) SetTypingStatusForMessageEdit(ctx context.Context, messageId int64) SetTypingStatusForMessageEditRequest {
	return SetTypingStatusForMessageEditRequest{
		apiService: s,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (s *usersService) SetTypingStatusForMessageEditExecute(r SetTypingStatusForMessageEditRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/messages/{message_id}/typing"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", IdToString(r.messageId), -1)

	if r.op == nil {
		return nil, nil, fmt.Errorf("op is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "op", r.op)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UnmuteUserRequest struct {
	ctx         context.Context
	apiService  APIUsers
	mutedUserId int64
}

func (r UnmuteUserRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UnmuteUserExecute(r)
}

// UnmuteUser Unmute a user
//
// [Unmute a user]
// from the perspective of the requesting user.
//
// *Changes**: New in Zulip 4.0 (feature level 48).
//
// [Unmute a user]: https://zulip.com/help/mute-a-user#see-your-list-of-muted-users
func (s *usersService) UnmuteUser(ctx context.Context, mutedUserId int64) UnmuteUserRequest {
	return UnmuteUserRequest{
		apiService:  s,
		ctx:         ctx,
		mutedUserId: mutedUserId,
	}
}

// Execute executes the request
func (s *usersService) UnmuteUserExecute(r UnmuteUserRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/muted_users/{muted_user_id}"
	)

	endpoint = strings.Replace(endpoint, "{muted_user_id}", IdToString(r.mutedUserId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdatePresenceRequest struct {
	ctx              context.Context
	apiService       APIUsers
	status           *zulip.PresenceStatus
	lastUpdateId     *int64
	historyLimitDays *int32
	newUserInput     *bool
	pingOnly         *bool
	slimPresence     *bool
}

// The status of the user on this client.  Clients should report the user as `PresenceStatusActive` on this device if the client knows that the user is presently using the device (and thus would potentially see a notification immediately), even if the user has not directly interacted with the Zulip client.  Otherwise, it should report the user as `PresenceStatusIdle`.
// See the related [`new_user_input`] parameter for how a client should report whether the user is actively using the Zulip client.
//
// [`new_user_input`]: https://zulip.com/api/update-presence#parameter-new_user_input
func (r UpdatePresenceRequest) Status(status zulip.PresenceStatus) UpdatePresenceRequest {
	r.status = &status
	return r
}

// The identifier that specifies what presence data the client already has received, which allows the server to only return more recent user presence data.  This should be set to `-1` during initialization of the client in order to fetch all user presence data, unless the client is obtaining initial user presence metadata from the [`POST /register`] endpoint.  In subsequent queries to this endpoint, this value should be set to the most recent value of `presence_last_update_id` returned by the server in this endpoint's response, which implements incremental fetching of user presence data.  When this parameter is passed, the user presence data in the response will always be in the modern format.
//
// **Changes**: New in Zulip 9.0 (feature level 263). Previously, the server sent user presence data for all users who had been active in the last two weeks unconditionally.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r UpdatePresenceRequest) LastUpdateId(lastUpdateId int64) UpdatePresenceRequest {
	r.lastUpdateId = &lastUpdateId
	return r
}

// Limits how far back in time to fetch user presence data. If not specified, defaults to 14 days. A value of N means that the oldest presence data fetched will be from at most N days ago.  Note that this is only useful during the initial user presence data fetch, as subsequent fetches should use the `last_update_id` parameter, which will act as the limit on how much presence data is returned. `history_limit_days` is ignored if `last_update_id` is passed with a value greater than `0`, indicating that the client already has some presence data.
//
// **Changes**: New in Zulip 10.0 (feature level 288).
func (r UpdatePresenceRequest) HistoryLimitDays(historyLimitDays int32) UpdatePresenceRequest {
	r.historyLimitDays = &historyLimitDays
	return r
}

// Whether the user has interacted with the client (e.g. moved the mouse, used the keyboard, etc.) since the previous presence request from this client.  The server uses data from this parameter to implement certain [usage statistics].  User interface clients that might run in the background, without the user ever interacting with them, should be careful to only pass `true` if the user has actually interacted with the client in order to avoid corrupting usage statistics graphs.
//
// [usage statistics]: https://zulip.com/help/analytics
func (r UpdatePresenceRequest) NewUserInput(newUserInput bool) UpdatePresenceRequest {
	r.newUserInput = &newUserInput
	return r
}

// Whether the client is sending a ping-only request, meaning it only wants to update the user's presence `status` on the server.  Otherwise, also requests the server return user presence data for all users in the organization, which is further specified by the [`last_update_id`] parameter.
//
// [`last_update_id`]: https://zulip.com/api/update-presence#parameter-last_update_id
func (r UpdatePresenceRequest) PingOnly(pingOnly bool) UpdatePresenceRequest {
	r.pingOnly = &pingOnly
	return r
}

// Legacy parameter for configuring the format (modern or legacy) in which the server will return user presence data for the organization.  Modern clients should use [`last_update_id`], which guarantees that user presence data will be returned in the modern format, and should not pass this parameter as `true` unless interacting with an older server.  Legacy clients that do not yet support `last_update_id` may use the value of `true` to request the modern format for user presence data.  **Note**: The legacy format for user presence data will be removed entirely in a future release.
//
// **Changes**:**Deprecated** in Zulip 9.0 (feature level 263). Using the modern `last_update_id` parameter is the recommended way to request the modern format for user presence data.  New in Zulip 3.0 (no feature level as it was an unstable API at that point).
//
// [`last_update_id`]: https://zulip.com/api/update-presence#parameter-last_update_id
func (r UpdatePresenceRequest) SlimPresence(slimPresence bool) UpdatePresenceRequest {
	r.slimPresence = &slimPresence
	return r
}

func (r UpdatePresenceRequest) Execute() (*UpdatePresenceResponse, *http.Response, error) {
	return r.apiService.UpdatePresenceExecute(r)
}

// UpdatePresence Update your presence
//
// Update the current user's [presence] and fetch presence data
// of other users in the organization.
//
// This endpoint is meant to be used by clients for both:
//   - Reporting the current user's presence status (`"active"` or `"idle"`) to the server.
//   - Obtaining the presence data of all other users in the organization via regular polling.
//
// Accurate user presence is one of the most expensive parts of any
// chat application (in terms of bandwidth and other resources). Therefore,
// it is important that clients implementing Zulip's user presence system
// use the modern [`last_update_id`] protocol to
// minimize fetching duplicate user presence data.
//
// Client apps implementing presence are recommended to also consume [`presence` events], in order to learn about newly online users
// immediately.
//
// The Zulip server is responsible for implementing [invisible mode],
// which disables sharing a user's presence data. Nonetheless, clients
// should check the `presence_enabled` field in user objects in order to
// display the current user as online or offline based on whether they are
// sharing their presence information.
//
// *Changes**: As of Zulip 8.0 (feature level 228), if the
// `CAN_ACCESS_ALL_USERS_GROUP_LIMITS_PRESENCE` server-level setting is
// `true`, then user presence data in the response is [limited to users the current user can see/access].
//
// [limited to users the current user can see/access]: https://zulip.com/help/guest-users#configure-whether-guests-can-see-all-other-users
// [invisible mode]: https://zulip.com/help/status-and-availability#invisible-mode
// [presence]: https://zulip.com/help/status-and-availability#availability
// [`presence` events]: https://zulip.com/api/get-events#presence
// [`last_update_id`]: https://zulip.com/api/update-presence#parameter-last_update_id
func (s *usersService) UpdatePresence(ctx context.Context) UpdatePresenceRequest {
	return UpdatePresenceRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) UpdatePresenceExecute(r UpdatePresenceRequest) (*UpdatePresenceResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &UpdatePresenceResponse{}
		endpoint = "/users/me/presence"
	)
	if r.status == nil {
		return nil, nil, fmt.Errorf("status is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "last_update_id", r.lastUpdateId)
	AddOptionalParam(form, "history_limit_days", r.historyLimitDays)
	AddOptionalParam(form, "new_user_input", r.newUserInput)
	AddOptionalParam(form, "ping_only", r.pingOnly)
	AddOptionalParam(form, "slim_presence", r.slimPresence)
	AddParam(form, "status", r.status)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateSettingsRequest struct {
	ctx                                            context.Context
	apiService                                     APIUsers
	fullName                                       *string
	email                                          *string
	oldPassword                                    *string
	newPassword                                    *string
	twentyFourHourTime                             *bool
	webMarkReadOnScrollPolicy                      *int32
	webChannelDefaultView                          *zulip.ChannelDefaultView
	starredMessageCounts                           *bool
	receivesTypingNotifications                    *bool
	webSuggestUpdateTimezone                       *bool
	fluidLayoutWidth                               *bool
	highContrastMode                               *bool
	webFontSizePx                                  *int32
	webLineHeightPercent                           *int32
	colorScheme                                    *zulip.ColorScheme
	enableDraftsSynchronization                    *bool
	translateEmoticons                             *bool
	displayEmojiReactionUsers                      *bool
	defaultLanguage                                *string
	webHomeView                                    *zulip.HomeView
	webEscapeNavigatesToHomeView                   *bool
	leftSideUserlist                               *bool
	emojiset                                       *zulip.Emojiset
	demoteInactiveChannels                         *zulip.DemoteInactiveChannels
	userListStyle                                  *zulip.UserListStyle
	webAnimateImagePreviews                        *zulip.WebAnimateImagePreviews
	webChannelUnreadsCountDisplayPolicy            *zulip.UnreadsCountDisplay
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
	desktopIconCountDisplay                        *zulip.BadgeCount
	realmNameInEmailNotificationsPolicy            *int32
	automaticallyFollowTopicsPolicy                *zulip.TopicInteraction
	automaticallyUnmuteTopicsInMutedChannelsPolicy *zulip.TopicInteraction
	automaticallyFollowTopicsWhereMentioned        *bool
	resolvedTopicNoticeAutoReadPolicy              *zulip.ResolvedTopicNoticeAutoReadPolicy
	presenceEnabled                                *bool
	enterSends                                     *bool
	sendPrivateTypingNotifications                 *bool
	sendChannelTypingNotifications                 *bool
	sendReadReceipts                               *bool
	allowPrivateDataExport                         *bool
	emailAddressVisibility                         *zulip.EmailVisibility
	webNavigateToSentMessage                       *bool
}

// A new display name for the user.
func (r UpdateSettingsRequest) FullName(fullName string) UpdateSettingsRequest {
	r.fullName = &fullName
	return r
}

// Asks the server to initiate a confirmation sequence to change the user's email address to the indicated value. The user will need to demonstrate control of the new email address by clicking a confirmation link sent to that address.
func (r UpdateSettingsRequest) Email(email string) UpdateSettingsRequest {
	r.email = &email
	return r
}

// The user's old Zulip password (or LDAP password, if LDAP authentication is in use).  Required only when sending the `new_password` parameter.
func (r UpdateSettingsRequest) OldPassword(oldPassword string) UpdateSettingsRequest {
	r.oldPassword = &oldPassword
	return r
}

// The user's new Zulip password (or LDAP password, if LDAP authentication is in use).  The `old_password` parameter must be included in the request.
func (r UpdateSettingsRequest) NewPassword(newPassword string) UpdateSettingsRequest {
	r.newPassword = &newPassword
	return r
}

// Whether time should be [displayed in 24-hour notation].
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.
//
// [displayed in 24-hour notation]: https://zulip.com/help/change-the-time-format
func (r UpdateSettingsRequest) TwentyFourHourTime(twentyFourHourTime bool) UpdateSettingsRequest {
	r.twentyFourHourTime = &twentyFourHourTime
	return r
}

// Whether or not to mark messages as read when the user scrolls through their feed.
//   - 1 = Always
//   - 2 = Only in conversation views
//   - 3 = Never
//
// **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the "Always" setting when marking messages as read.
func (r UpdateSettingsRequest) WebMarkReadOnScrollPolicy(webMarkReadOnScrollPolicy int32) UpdateSettingsRequest {
	r.webMarkReadOnScrollPolicy = &webMarkReadOnScrollPolicy
	return r
}

// Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.
//   - ChannelDefaultViewTopTopicInChannel
//   - ChannelDefaultViewChannelFeed
//   - ChannelDefaultViewListOfTopics
//   - ChannelDefaultViewTopUnreadTopicInChannel
//
// **Changes**: The "Top unread topic in channel" is new in Zulip 11.0 (feature level 401).  The "List of topics" option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the "Channel feed" behavior.
func (r UpdateSettingsRequest) WebChannelDefaultView(webChannelDefaultView zulip.ChannelDefaultView) UpdateSettingsRequest {
	r.webChannelDefaultView = &webChannelDefaultView
	return r
}

// Whether clients should display the [number of starred messages].
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.
//
// [number of starred messages]: https://zulip.com/help/star-a-message#display-the-number-of-starred-messages
func (r UpdateSettingsRequest) StarredMessageCounts(starredMessageCounts bool) UpdateSettingsRequest {
	r.starredMessageCounts = &starredMessageCounts
	return r
}

// Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  By default, this is set to true, enabling user to receive typing notifications from other users.
//
// **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.
func (r UpdateSettingsRequest) ReceivesTypingNotifications(receivesTypingNotifications bool) UpdateSettingsRequest {
	r.receivesTypingNotifications = &receivesTypingNotifications
	return r
}

// Whether the user should be shown an alert, offering to update their [profile time zone], when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.
//
// **Changes**: New in Zulip 10.0 (feature level 329).
//
// [profile time zone]: https://zulip.com/help/change-your-timezone
func (r UpdateSettingsRequest) WebSuggestUpdateTimezone(webSuggestUpdateTimezone bool) UpdateSettingsRequest {
	r.webSuggestUpdateTimezone = &webSuggestUpdateTimezone
	return r
}

// Whether to use the [maximum available screen width] for the web app's center panel (message feed, recent conversations) on wide screens.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.
//
// [maximum available screen width]: https://zulip.com/help/enable-full-width-display
func (r UpdateSettingsRequest) FluidLayoutWidth(fluidLayoutWidth bool) UpdateSettingsRequest {
	r.fluidLayoutWidth = &fluidLayoutWidth
	return r
}

// This setting is reserved for use to control variations in Zulip's design to help visually impaired users.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.
func (r UpdateSettingsRequest) HighContrastMode(highContrastMode bool) UpdateSettingsRequest {
	r.highContrastMode = &highContrastMode
	return r
}

// User-configured primary `font-size` for the web application, in pixels.
//
// **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.
func (r UpdateSettingsRequest) WebFontSizePx(webFontSizePx int32) UpdateSettingsRequest {
	r.webFontSizePx = &webFontSizePx
	return r
}

// User-configured primary `line-height` for the web application, in percent, so a value of 120 represents a `line-height` of 1.2.
//
// **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.
func (r UpdateSettingsRequest) WebLineHeightPercent(webLineHeightPercent int32) UpdateSettingsRequest {
	r.webLineHeightPercent = &webLineHeightPercent
	return r
}

// Controls which [color theme] to use.
//
//   - ColorSchemeAutomatic
//   - ColorSchemeDark
//   - ColorSchemeLight
//
// Automatic detection is implementing using the standard `prefers-color-scheme` media query.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.
//
// [color theme]: https://zulip.com/help/dark-theme
func (r UpdateSettingsRequest) ColorScheme(colorScheme zulip.ColorScheme) UpdateSettingsRequest {
	r.colorScheme = &colorScheme
	return r
}

// A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.
//
// **Changes**: New in Zulip 5.0 (feature level 87).
func (r UpdateSettingsRequest) EnableDraftsSynchronization(enableDraftsSynchronization bool) UpdateSettingsRequest {
	r.enableDraftsSynchronization = &enableDraftsSynchronization
	return r
}

// Whether to [translate emoticons to emoji] in messages the user sends.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.
//
// [translate emoticons to emoji]: https://zulip.com/help/configure-emoticon-translations
func (r UpdateSettingsRequest) TranslateEmoticons(translateEmoticons bool) UpdateSettingsRequest {
	r.translateEmoticons = &translateEmoticons
	return r
}

// Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.
//
// **Changes**: New in Zulip 6.0 (feature level 125).
func (r UpdateSettingsRequest) DisplayEmojiReactionUsers(displayEmojiReactionUsers bool) UpdateSettingsRequest {
	r.displayEmojiReactionUsers = &displayEmojiReactionUsers
	return r
}

// What [default language] to use for the account.  This controls both the Zulip UI as well as email notifications sent to the user.  The value needs to be a standard language code that the Zulip server has translation data for; for example, `"en"` for English or `"de"` for German.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 63).
//
// [default language]: https://zulip.com/help/change-your-language
func (r UpdateSettingsRequest) DefaultLanguage(defaultLanguage string) UpdateSettingsRequest {
	r.defaultLanguage = &defaultLanguage
	return r
}

// The [home view] used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.
//   - HomeViewRecentTopics
//   - HomeViewInbox
//   - HomeViewAllMessages
//
// **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).  Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).
//
// [home view]: https://zulip.com/help/configure-home-view
func (r UpdateSettingsRequest) WebHomeView(webHomeView zulip.HomeView) UpdateSettingsRequest {
	r.webHomeView = &webHomeView
	return r
}

// Whether the escape key navigates to the [configured home view].
//
// **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `escape_navigates_to_default_view`, which was new in Zulip 5.0 (feature level 107).
//
// [configured home view]: https://zulip.com/help/configure-home-view
func (r UpdateSettingsRequest) WebEscapeNavigatesToHomeView(webEscapeNavigatesToHomeView bool) UpdateSettingsRequest {
	r.webEscapeNavigatesToHomeView = &webEscapeNavigatesToHomeView
	return r
}

// Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.
func (r UpdateSettingsRequest) LeftSideUserlist(leftSideUserlist bool) UpdateSettingsRequest {
	r.leftSideUserlist = &leftSideUserlist
	return r
}

// The user's configured [emoji set], used to display emoji to the user everywhere they appear in the UI.
//   - EmojisetGoogle = Google modern
//   - EmojisetGoogleBlob = Google classic
//   - EmojisetTwitter = Twitter
//   - EmojisetText = Plain text
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).
//
// [emoji set]: https://zulip.com/help/emoji-and-emoticons#use-emoticons
func (r UpdateSettingsRequest) Emojiset(emojiset zulip.Emojiset) UpdateSettingsRequest {
	r.emojiset = &emojiset
	return r
}

// Whether to [hide inactive channels] in the left sidebar.
//   - DemoteInactiveChannelsAutomatic
//   - DemoteInactiveChannelsAlways
//   - DemoteInactiveChannelsNever
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.
//
// [hide inactive channels]: https://zulip.com/help/manage-inactive-channels
func (r UpdateSettingsRequest) DemoteInactiveChannels(demoteInactiveChannels zulip.DemoteInactiveChannels) UpdateSettingsRequest {
	r.demoteInactiveChannels = &demoteInactiveChannels
	return r
}

// The style selected by the user for the right sidebar user list.
//   - UserListStyleCompact
//   - UserListStyleWithStatus
//   - UserListStyleWithAvatarAndStatus
//
// **Changes**: New in Zulip 6.0 (feature level 141).
func (r UpdateSettingsRequest) UserListStyle(userListStyle zulip.UserListStyle) UpdateSettingsRequest {
	r.userListStyle = &userListStyle
	return r
}

// Controls how animated images should be played in the message feed in the web/desktop application.
//   - WebAnimateImagePreviewsAlways
//   - WebAnimateImagePreviewsOnHover
//   - WebAnimateImagePreviewsNever
//
// **Changes**: New in Zulip 9.0 (feature level 275).
func (r UpdateSettingsRequest) WebAnimateImagePreviews(webAnimateImagePreviews zulip.WebAnimateImagePreviews) UpdateSettingsRequest {
	r.webAnimateImagePreviews = &webAnimateImagePreviews
	return r
}

// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.
//   - UnreadsCountDisplayAllChannels
//   - UnreadsCountDisplayUnmutedChannelsAndTopics
//   - UnreadsCountDisplayNoChannels
//
// **Changes**: New in Zulip 8.0 (feature level 210).
func (r UpdateSettingsRequest) WebChannelUnreadsCountDisplayPolicy(webChannelUnreadsCountDisplayPolicy zulip.UnreadsCountDisplay) UpdateSettingsRequest {
	r.webChannelUnreadsCountDisplayPolicy = &webChannelUnreadsCountDisplayPolicy
	return r
}

// Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.
//
// **Changes**: New in Zulip 10.0 (feature level 350).
func (r UpdateSettingsRequest) HideAiFeatures(hideAiFeatures bool) UpdateSettingsRequest {
	r.hideAiFeatures = &hideAiFeatures
	return r
}

// Determines whether the web/desktop application's left sidebar displays any channel folders configured by the organization.
//
// **Changes**: New in Zulip 11.0 (feature level 411).
func (r UpdateSettingsRequest) WebLeftSidebarShowChannelFolders(webLeftSidebarShowChannelFolders bool) UpdateSettingsRequest {
	r.webLeftSidebarShowChannelFolders = &webLeftSidebarShowChannelFolders
	return r
}

// Determines whether the web/desktop application's left sidebar displays the unread message count summary.
//
// **Changes**: New in Zulip 11.0 (feature level 398).
func (r UpdateSettingsRequest) WebLeftSidebarUnreadsCountSummary(webLeftSidebarUnreadsCountSummary bool) UpdateSettingsRequest {
	r.webLeftSidebarUnreadsCountSummary = &webLeftSidebarUnreadsCountSummary
	return r
}

// The IANA identifier of the user's [profile time zone], which is used primarily to display the user's local time to other users.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/display` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 64).
//
// [profile time zone]: https://zulip.com/help/change-your-timezone
func (r UpdateSettingsRequest) Timezone(timezone string) UpdateSettingsRequest {
	r.timezone = &timezone
	return r
}

// Enable visual desktop notifications for channel messages.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableChannelDesktopNotifications(enableChannelDesktopNotifications bool) UpdateSettingsRequest {
	r.enableChannelDesktopNotifications = &enableChannelDesktopNotifications
	return r
}

// Enable email notifications for channel messages.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableChannelEmailNotifications(enableChannelEmailNotifications bool) UpdateSettingsRequest {
	r.enableChannelEmailNotifications = &enableChannelEmailNotifications
	return r
}

// Enable mobile notifications for channel messages.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableChannelPushNotifications(enableChannelPushNotifications bool) UpdateSettingsRequest {
	r.enableChannelPushNotifications = &enableChannelPushNotifications
	return r
}

// Enable audible desktop notifications for channel messages.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableChannelAudibleNotifications(enableChannelAudibleNotifications bool) UpdateSettingsRequest {
	r.enableChannelAudibleNotifications = &enableChannelAudibleNotifications
	return r
}

// Notification sound name.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.  Unnecessary JSON-encoding of this parameter was removed in Zulip 4.0 (feature level 63).
func (r UpdateSettingsRequest) NotificationSound(notificationSound string) UpdateSettingsRequest {
	r.notificationSound = &notificationSound
	return r
}

// Enable visual desktop notifications for direct messages and @-mentions.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableDesktopNotifications(enableDesktopNotifications bool) UpdateSettingsRequest {
	r.enableDesktopNotifications = &enableDesktopNotifications
	return r
}

// Enable audible desktop notifications for direct messages and @-mentions.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableSounds(enableSounds bool) UpdateSettingsRequest {
	r.enableSounds = &enableSounds
	return r
}

// The duration (in seconds) for which the server should wait to batch email notifications before sending them.
//
// **Changes**: New in Zulip 5.0 (feature level 82)
func (r UpdateSettingsRequest) EmailNotificationsBatchingPeriodSeconds(emailNotificationsBatchingPeriodSeconds int32) UpdateSettingsRequest {
	r.emailNotificationsBatchingPeriodSeconds = &emailNotificationsBatchingPeriodSeconds
	return r
}

// Enable email notifications for direct messages and @-mentions received when the user is offline.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableOfflineEmailNotifications(enableOfflineEmailNotifications bool) UpdateSettingsRequest {
	r.enableOfflineEmailNotifications = &enableOfflineEmailNotifications
	return r
}

// Enable mobile notification for direct messages and @-mentions received when the user is offline.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableOfflinePushNotifications(enableOfflinePushNotifications bool) UpdateSettingsRequest {
	r.enableOfflinePushNotifications = &enableOfflinePushNotifications
	return r
}

// Enable mobile notification for direct messages and @-mentions received when the user is online.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableOnlinePushNotifications(enableOnlinePushNotifications bool) UpdateSettingsRequest {
	r.enableOnlinePushNotifications = &enableOnlinePushNotifications
	return r
}

// Enable visual desktop notifications for messages sent to followed topics.
//
// **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicDesktopNotifications(enableFollowedTopicDesktopNotifications bool) UpdateSettingsRequest {
	r.enableFollowedTopicDesktopNotifications = &enableFollowedTopicDesktopNotifications
	return r
}

// Enable email notifications for messages sent to followed topics.
//
// **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicEmailNotifications(enableFollowedTopicEmailNotifications bool) UpdateSettingsRequest {
	r.enableFollowedTopicEmailNotifications = &enableFollowedTopicEmailNotifications
	return r
}

// Enable push notifications for messages sent to followed topics.
//
// **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicPushNotifications(enableFollowedTopicPushNotifications bool) UpdateSettingsRequest {
	r.enableFollowedTopicPushNotifications = &enableFollowedTopicPushNotifications
	return r
}

// Enable audible desktop notifications for messages sent to followed topics.
//
// **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicAudibleNotifications(enableFollowedTopicAudibleNotifications bool) UpdateSettingsRequest {
	r.enableFollowedTopicAudibleNotifications = &enableFollowedTopicAudibleNotifications
	return r
}

// Enable digest emails when the user is away.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableDigestEmails(enableDigestEmails bool) UpdateSettingsRequest {
	r.enableDigestEmails = &enableDigestEmails
	return r
}

// Enable marketing emails. Has no function outside Zulip Cloud.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableMarketingEmails(enableMarketingEmails bool) UpdateSettingsRequest {
	r.enableMarketingEmails = &enableMarketingEmails
	return r
}

// Enable email notifications for new logins to account.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) EnableLoginEmails(enableLoginEmails bool) UpdateSettingsRequest {
	r.enableLoginEmails = &enableLoginEmails
	return r
}

// Include the message's content in email notifications for new messages.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) MessageContentInEmailNotifications(messageContentInEmailNotifications bool) UpdateSettingsRequest {
	r.messageContentInEmailNotifications = &messageContentInEmailNotifications
	return r
}

// Include content of direct messages in desktop notifications.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) PmContentInDesktopNotifications(pmContentInDesktopNotifications bool) UpdateSettingsRequest {
	r.pmContentInDesktopNotifications = &pmContentInDesktopNotifications
	return r
}

// Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) WildcardMentionsNotify(wildcardMentionsNotify bool) UpdateSettingsRequest {
	r.wildcardMentionsNotify = &wildcardMentionsNotify
	return r
}

// Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.
//
// **Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateSettingsRequest) EnableFollowedTopicWildcardMentionsNotify(enableFollowedTopicWildcardMentionsNotify bool) UpdateSettingsRequest {
	r.enableFollowedTopicWildcardMentionsNotify = &enableFollowedTopicWildcardMentionsNotify
	return r
}

// Unread count badge (appears in desktop sidebar and browser tab)
//   - BadgeCountAllUnreadMessages
//   - BadgeCountDMsMentionsAndFollowedTopics
//   - BadgeCountDMsAndMentions
//   - BadgeCountNone
//
// **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.  Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) DesktopIconCountDisplay(desktopIconCountDisplay zulip.BadgeCount) UpdateSettingsRequest {
	r.desktopIconCountDisplay = &desktopIconCountDisplay
	return r
}

// Whether to [include organization name in subject of message notification emails].
//   - 1 = Automatic
//   - 2 = Always
//   - 3 = Never
//
// **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.  Before Zulip 5.0 (feature level 80), the previous `realm_name_in_notifications` setting was managed by the `PATCH /settings/notifications` endpoint.
//
// [include organization name in subject of message notification emails]: https://zulip.com/help/email-notifications#include-organization-name-in-subject-line
func (r UpdateSettingsRequest) RealmNameInEmailNotificationsPolicy(realmNameInEmailNotificationsPolicy int32) UpdateSettingsRequest {
	r.realmNameInEmailNotificationsPolicy = &realmNameInEmailNotificationsPolicy
	return r
}

// Which [topics to follow automatically].
//   - TopicInteractionTopicsTheUserParticipatesIn
//   - TopicInteractionTopicsTheUserSendsAMessageTo
//   - TopicInteractionTopicsTheUserStarts
//   - TopicInteractionNever
//
// **Changes**: New in Zulip 8.0 (feature level 214).
//
// [topics to follow automatically]: https://zulip.com/help/mute-a-topic
func (r UpdateSettingsRequest) AutomaticallyFollowTopicsPolicy(automaticallyFollowTopicsPolicy zulip.TopicInteraction) UpdateSettingsRequest {
	r.automaticallyFollowTopicsPolicy = &automaticallyFollowTopicsPolicy
	return r
}

// Which [topics to unmute automatically in muted channels].
//   - TopicInteractionTopicsTheUserParticipatesIn
//   - TopicInteractionTopicsTheUserSendsAMessageTo
//   - TopicInteractionTopicsTheUserStarts
//   - TopicInteractionNever
//
// **Changes**: New in Zulip 8.0 (feature level 214).
//
// [topics to unmute automatically in muted channels]: https://zulip.com/help/mute-a-topic
func (r UpdateSettingsRequest) AutomaticallyUnmuteTopicsInMutedChannelsPolicy(automaticallyUnmuteTopicsInMutedChannelsPolicy zulip.TopicInteraction) UpdateSettingsRequest {
	r.automaticallyUnmuteTopicsInMutedChannelsPolicy = &automaticallyUnmuteTopicsInMutedChannelsPolicy
	return r
}

// Whether the server will automatically mark the user as following topics where the user is mentioned.
//
// **Changes**: New in Zulip 8.0 (feature level 235).
func (r UpdateSettingsRequest) AutomaticallyFollowTopicsWhereMentioned(automaticallyFollowTopicsWhereMentioned bool) UpdateSettingsRequest {
	r.automaticallyFollowTopicsWhereMentioned = &automaticallyFollowTopicsWhereMentioned
	return r
}

// Controls whether the resolved-topic notices are marked as read.
//   - ResolvedTopicNoticeAutoReadPolicyAlways
//   - ResolvedTopicNoticeAutoReadPolicyExceptFollowed
//   - ResolvedTopicNoticeAutoReadPolicyNever
//
// **Changes**: New in Zulip 11.0 (feature level 385).
func (r UpdateSettingsRequest) ResolvedTopicNoticeAutoReadPolicy(resolvedTopicNoticeAutoReadPolicy zulip.ResolvedTopicNoticeAutoReadPolicy) UpdateSettingsRequest {
	r.resolvedTopicNoticeAutoReadPolicy = &resolvedTopicNoticeAutoReadPolicy
	return r
}

// Display the presence status to other users when online.
//
// **Changes**: Before Zulip 5.0 (feature level 80), this setting was managed by the `PATCH /settings/notifications` endpoint.
func (r UpdateSettingsRequest) PresenceEnabled(presenceEnabled bool) UpdateSettingsRequest {
	r.presenceEnabled = &presenceEnabled
	return r
}

// Whether pressing Enter in the compose box sends a message (or saves a message edit).
//
// **Changes**: Before Zulip 5.0 (feature level 81), this setting was managed by the `POST /users/me/enter-sends` endpoint, with the same parameter format.
func (r UpdateSettingsRequest) EnterSends(enterSends bool) UpdateSettingsRequest {
	r.enterSends = &enterSends
	return r
}

// Whether [typing notifications] be sent when composing direct messages.
//
// **Changes**: New in Zulip 5.0 (feature level 105).
//
// [typing notifications]: https://zulip.com/help/typing-notifications
func (r UpdateSettingsRequest) SendPrivateTypingNotifications(sendPrivateTypingNotifications bool) UpdateSettingsRequest {
	r.sendPrivateTypingNotifications = &sendPrivateTypingNotifications
	return r
}

// Whether [typing notifications] be sent when composing channel messages.
//
// **Changes**: New in Zulip 5.0 (feature level 105).
//
// [typing notifications]: https://zulip.com/help/typing-notifications
func (r UpdateSettingsRequest) SendChannelTypingNotifications(sendChannelTypingNotifications bool) UpdateSettingsRequest {
	r.sendChannelTypingNotifications = &sendChannelTypingNotifications
	return r
}

// Whether other users are allowed to see whether you've read messages.
//
// **Changes**: New in Zulip 5.0 (feature level 105).
func (r UpdateSettingsRequest) SendReadReceipts(sendReadReceipts bool) UpdateSettingsRequest {
	r.sendReadReceipts = &sendReadReceipts
	return r
}

// Whether organization administrators are allowed to export your private data.
//
// **Changes**: New in Zulip 10.0 (feature level 293).
func (r UpdateSettingsRequest) AllowPrivateDataExport(allowPrivateDataExport bool) UpdateSettingsRequest {
	r.allowPrivateDataExport = &allowPrivateDataExport
	return r
}

// The [policy] this user has selected for [which other users] in this organization can see their real email address.
//   - EmailVisibilityEveryone
//   - EmailVisibilityMembersOnly
//   - EmailVisibilityAdministratorsOnly
//   - EmailVisibilityNobody
//   - EmailVisibilityModeratorsOnly
//
// **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.
//
// [policy]: https://zulip.com/api/roles-and-permissions#permission-levels
// [which other users]: https://zulip.com/help/configure-email-visibility
func (r UpdateSettingsRequest) EmailAddressVisibility(emailAddressVisibility zulip.EmailVisibility) UpdateSettingsRequest {
	r.emailAddressVisibility = &emailAddressVisibility
	return r
}

// Web/desktop app setting for whether the user's view should automatically go to the conversation where they sent a message.
//
// **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.
func (r UpdateSettingsRequest) WebNavigateToSentMessage(webNavigateToSentMessage bool) UpdateSettingsRequest {
	r.webNavigateToSentMessage = &webNavigateToSentMessage
	return r
}

func (r UpdateSettingsRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateSettingsExecute(r)
}

// UpdateSettings Update settings
//
// This endpoint is used to edit the current user's settings.
//
// *Changes**: Removed `dense_mode` setting in Zulip 10.0
// (feature level 364) as we now have `web_font_size_px` and
// `web_line_height_percent` settings for more control.
//
// Prior to Zulip 5.0 (feature level 80), this endpoint only
// supported the `full_name`, `email`, `old_password`, and
// `new_password` parameters. Notification settings were
// managed by `PATCH /settings/notifications`, and all other
// settings by `PATCH /settings/display`.
//
// The feature level 80 migration to merge these endpoints did not
// change how request parameters are encoded. However, it did change
// the handling of any invalid parameters present in a request
// (see feature level 78 change below).
//
// As of feature level 80, the `PATCH /settings/display` and
// `PATCH /settings/notifications` endpoints are deprecated aliases
// for this endpoint for backwards-compatibility, and will be removed
// once clients have migrated to use this endpoint.
//
// Prior to Zulip 5.0 (feature level 78), this endpoint indicated
// which parameters it had processed by including in the response
// object `"key": value` entries for values successfully changed by
// the request. That was replaced by the more ergonomic
// [`ignored_parameters_unsupported`] array.
//
// The `PATCH /settings/notifications` and `PATCH /settings/display`
// endpoints also had this behavior of indicating processed parameters
// before they became aliases of this endpoint in Zulip 5.0 (see
// feature level 80 change above).
//
// Before feature level 78, request parameters that were not supported
// (or were unchanged) were silently ignored.
//
// [`ignored_parameters_unsupported`]: https://zulip.com/api/rest-error-handling#ignored-parameters
func (s *usersService) UpdateSettings(ctx context.Context) UpdateSettingsRequest {
	return UpdateSettingsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) UpdateSettingsExecute(r UpdateSettingsRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/settings"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "full_name", r.fullName)
	AddOptionalParam(form, "email", r.email)
	AddOptionalParam(form, "old_password", r.oldPassword)
	AddOptionalParam(form, "new_password", r.newPassword)
	AddOptionalParam(form, "twenty_four_hour_time", r.twentyFourHourTime)
	AddOptionalParam(form, "web_mark_read_on_scroll_policy", r.webMarkReadOnScrollPolicy)
	AddOptionalParam(form, "web_channel_default_view", r.webChannelDefaultView)
	AddOptionalParam(form, "starred_message_counts", r.starredMessageCounts)
	AddOptionalParam(form, "receives_typing_notifications", r.receivesTypingNotifications)
	AddOptionalParam(form, "web_suggest_update_timezone", r.webSuggestUpdateTimezone)
	AddOptionalParam(form, "fluid_layout_width", r.fluidLayoutWidth)
	AddOptionalParam(form, "high_contrast_mode", r.highContrastMode)
	AddOptionalParam(form, "web_font_size_px", r.webFontSizePx)
	AddOptionalParam(form, "web_line_height_percent", r.webLineHeightPercent)
	AddOptionalParam(form, "color_scheme", r.colorScheme)
	AddOptionalParam(form, "enable_drafts_synchronization", r.enableDraftsSynchronization)
	AddOptionalParam(form, "translate_emoticons", r.translateEmoticons)
	AddOptionalParam(form, "display_emoji_reaction_users", r.displayEmojiReactionUsers)
	AddOptionalParam(form, "default_language", r.defaultLanguage)
	AddOptionalParam(form, "web_home_view", r.webHomeView)
	AddOptionalParam(form, "web_escape_navigates_to_home_view", r.webEscapeNavigatesToHomeView)
	AddOptionalParam(form, "left_side_userlist", r.leftSideUserlist)
	AddOptionalParam(form, "emojiset", r.emojiset)
	AddOptionalParam(form, "demote_inactive_streams", r.demoteInactiveChannels)
	AddOptionalParam(form, "user_list_style", r.userListStyle)
	AddOptionalParam(form, "web_animate_image_previews", r.webAnimateImagePreviews)
	AddOptionalParam(form, "web_stream_unreads_count_display_policy", r.webChannelUnreadsCountDisplayPolicy)
	AddOptionalParam(form, "hide_ai_features", r.hideAiFeatures)
	AddOptionalParam(form, "web_left_sidebar_show_channel_folders", r.webLeftSidebarShowChannelFolders)
	AddOptionalParam(form, "web_left_sidebar_unreads_count_summary", r.webLeftSidebarUnreadsCountSummary)
	AddOptionalParam(form, "timezone", r.timezone)
	AddOptionalParam(form, "enable_stream_desktop_notifications", r.enableChannelDesktopNotifications)
	AddOptionalParam(form, "enable_stream_email_notifications", r.enableChannelEmailNotifications)
	AddOptionalParam(form, "enable_stream_push_notifications", r.enableChannelPushNotifications)
	AddOptionalParam(form, "enable_stream_audible_notifications", r.enableChannelAudibleNotifications)
	AddOptionalParam(form, "notification_sound", r.notificationSound)
	AddOptionalParam(form, "enable_desktop_notifications", r.enableDesktopNotifications)
	AddOptionalParam(form, "enable_sounds", r.enableSounds)
	AddOptionalParam(form, "email_notifications_batching_period_seconds", r.emailNotificationsBatchingPeriodSeconds)
	AddOptionalParam(form, "enable_offline_email_notifications", r.enableOfflineEmailNotifications)
	AddOptionalParam(form, "enable_offline_push_notifications", r.enableOfflinePushNotifications)
	AddOptionalParam(form, "enable_online_push_notifications", r.enableOnlinePushNotifications)
	AddOptionalParam(form, "enable_followed_topic_desktop_notifications", r.enableFollowedTopicDesktopNotifications)
	AddOptionalParam(form, "enable_followed_topic_email_notifications", r.enableFollowedTopicEmailNotifications)
	AddOptionalParam(form, "enable_followed_topic_push_notifications", r.enableFollowedTopicPushNotifications)
	AddOptionalParam(form, "enable_followed_topic_audible_notifications", r.enableFollowedTopicAudibleNotifications)
	AddOptionalParam(form, "enable_digest_emails", r.enableDigestEmails)
	AddOptionalParam(form, "enable_marketing_emails", r.enableMarketingEmails)
	AddOptionalParam(form, "enable_login_emails", r.enableLoginEmails)
	AddOptionalParam(form, "message_content_in_email_notifications", r.messageContentInEmailNotifications)
	AddOptionalParam(form, "pm_content_in_desktop_notifications", r.pmContentInDesktopNotifications)
	AddOptionalParam(form, "wildcard_mentions_notify", r.wildcardMentionsNotify)
	AddOptionalParam(form, "enable_followed_topic_wildcard_mentions_notify", r.enableFollowedTopicWildcardMentionsNotify)
	AddOptionalParam(form, "desktop_icon_count_display", r.desktopIconCountDisplay)
	AddOptionalParam(form, "realm_name_in_email_notifications_policy", r.realmNameInEmailNotificationsPolicy)
	AddOptionalParam(form, "automatically_follow_topics_policy", r.automaticallyFollowTopicsPolicy)
	AddOptionalParam(form, "automatically_unmute_topics_in_muted_streams_policy", r.automaticallyUnmuteTopicsInMutedChannelsPolicy)
	AddOptionalParam(form, "automatically_follow_topics_where_mentioned", r.automaticallyFollowTopicsWhereMentioned)
	AddOptionalParam(form, "resolved_topic_notice_auto_read_policy", r.resolvedTopicNoticeAutoReadPolicy)
	AddOptionalParam(form, "presence_enabled", r.presenceEnabled)
	AddOptionalParam(form, "enter_sends", r.enterSends)
	AddOptionalParam(form, "send_private_typing_notifications", r.sendPrivateTypingNotifications)
	AddOptionalParam(form, "send_stream_typing_notifications", r.sendChannelTypingNotifications)
	AddOptionalParam(form, "send_read_receipts", r.sendReadReceipts)
	AddOptionalParam(form, "allow_private_data_export", r.allowPrivateDataExport)
	AddOptionalParam(form, "email_address_visibility", r.emailAddressVisibility)
	AddOptionalParam(form, "web_navigate_to_sent_message", r.webNavigateToSentMessage)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateStatusRequest struct {
	ctx          context.Context
	apiService   APIUsers
	statusText   *string
	away         *bool
	emojiName    *string
	emojiCode    *string
	reactionType *zulip.ReactionType
}

// The text content of the status message. Sending the empty string will clear the user's status.  **Note**: The limit on the size of the message is 60 characters.
func (r UpdateStatusRequest) StatusText(statusText string) UpdateStatusRequest {
	r.statusText = &statusText
	return r
}

// Whether the user should be marked as "away".
//
// **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, `away` is a legacy way to access the user's `presence_enabled` setting, with `away - !presence_enabled`. To be removed in a future release.
func (r UpdateStatusRequest) Away(away bool) UpdateStatusRequest {
	r.away = &away
	return r
}

// The name for the emoji to associate with this status.
//
// **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusRequest) EmojiName(emojiName string) UpdateStatusRequest {
	r.emojiName = &emojiName
	return r
}

// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.
//
// **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusRequest) EmojiCode(emojiCode string) UpdateStatusRequest {
	r.emojiCode = &emojiCode
	return r
}

// A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.
//
// Must be one of the following values:
//   - ReactionTypeRealmEmoji
//   - ReactionTypeUnicodeEmoji
//   - ReactionTypeZulipExtraEmoji
//   - ReactionTypeEmpty
//
// **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusRequest) ReactionType(reactionType zulip.ReactionType) UpdateStatusRequest {
	r.reactionType = &reactionType
	return r
}

func (r UpdateStatusRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateStatusExecute(r)
}

// UpdateStatus Update your status
//
// Change your [status].
//
// A request to this endpoint will only change the parameters passed.
// For example, passing just `status_text` requests a change in the status
// text, but will leave the status emoji unchanged.
//
// Clients that wish to set the user's status to a specific value should
// pass all supported parameters.
//
// *Changes**: In Zulip 5.0 (feature level 86), added support for
// `emoji_name`, `emoji_code`, and `reaction_type` parameters.
//
// [status]: https://zulip.com/help/status-and-availability
func (s *usersService) UpdateStatus(ctx context.Context) UpdateStatusRequest {
	return UpdateStatusRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *usersService) UpdateStatusExecute(r UpdateStatusRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/status"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "status_text", r.statusText)
	AddOptionalParam(form, "away", r.away)
	AddOptionalParam(form, "emoji_name", r.emojiName)
	AddOptionalParam(form, "emoji_code", r.emojiCode)
	AddOptionalParam(form, "reaction_type", r.reactionType)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateStatusForUserRequest struct {
	ctx          context.Context
	apiService   APIUsers
	userId       int64
	statusText   *string
	emojiName    *string
	emojiCode    *string
	reactionType *string
}

// The text content of the status message. Sending the empty string will clear the user's status.  **Note**: The limit on the size of the message is 60 characters.
func (r UpdateStatusForUserRequest) StatusText(statusText string) UpdateStatusForUserRequest {
	r.statusText = &statusText
	return r
}

// The name for the emoji to associate with this status.
//
// **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusForUserRequest) EmojiName(emojiName string) UpdateStatusForUserRequest {
	r.emojiName = &emojiName
	return r
}

// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.
//
// **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusForUserRequest) EmojiCode(emojiCode string) UpdateStatusForUserRequest {
	r.emojiCode = &emojiCode
	return r
}

// A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:
//   - ReactionTypeRealmEmoji
//   - ReactionTypeUnicodeEmoji
//   - ReactionTypeZulipExtraEmoji
//   - ReactionTypeEmpty
//
// **Changes**: New in Zulip 5.0 (feature level 86).
func (r UpdateStatusForUserRequest) ReactionType(reactionType string) UpdateStatusForUserRequest {
	r.reactionType = &reactionType
	return r
}

func (r UpdateStatusForUserRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateStatusForUserExecute(r)
}

// UpdateStatusForUser Update user status
//
// Administrator endpoint for changing the [status] of
// another user.
//
// *Changes**: New in Zulip 11.0 (feature level 407).
//
// [status]: https://zulip.com/help/status-and-availability
func (s *usersService) UpdateStatusForUser(ctx context.Context, userId int64) UpdateStatusForUserRequest {
	return UpdateStatusForUserRequest{
		apiService: s,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
func (s *usersService) UpdateStatusForUserExecute(r UpdateStatusForUserRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/{user_id}/status"
	)

	endpoint = strings.Replace(endpoint, "{user_id}", IdToString(r.userId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "status_text", r.statusText)
	AddOptionalParam(form, "emoji_name", r.emojiName)
	AddOptionalParam(form, "emoji_code", r.emojiCode)
	AddOptionalParam(form, "reaction_type", r.reactionType)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateUserRequest struct {
	ctx         context.Context
	apiService  APIUsers
	userId      int64
	fullName    *string
	role        *zulip.Role
	profileData *[]map[string]interface{}
	newEmail    *string
}

// The user's full name.
//
// **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 5.0 (feature level 106).
func (r UpdateUserRequest) FullName(fullName string) UpdateUserRequest {
	r.fullName = &fullName
	return r
}

// New [role] for the user. Roles are encoded as:  - Organization owner: 100 - Organization administrator: 200 - Organization moderator: 300 - Member: 400 - Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.
//
// **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of `is_admin` and `is_guest` boolean parameters. Organization moderator role added in Zulip 4.0 (feature level 60).
//
// [role]: https://zulip.com/api/roles-and-permissions
func (r UpdateUserRequest) Role(role zulip.Role) UpdateUserRequest {
	r.role = &role
	return r
}

// A dictionary containing the updated custom profile field data for the user.
func (r UpdateUserRequest) ProfileData(profileData []map[string]interface{}) UpdateUserRequest {
	r.profileData = &profileData
	return r
}

// New email address for the user. Requires the user making the request to be an organization owner and additionally have the `.can_change_user_emails` special permission.
//
// **Changes**: New in Zulip 10.0 (feature level 285).
func (r UpdateUserRequest) NewEmail(newEmail string) UpdateUserRequest {
	r.newEmail = &newEmail
	return r
}

func (r UpdateUserRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateUserExecute(r)
}

// UpdateUser Update a user
//
// Administrative endpoint to update the details of another user in the organization.
//
// Supports everything an administrator can do to edit details of another
// user's account, including editing full name,
// [role], and [custom profile fields].
//
// [role]: https://zulip.com/help/user-roles
// [custom profile fields]: https://zulip.com/help/custom-profile-fields
func (s *usersService) UpdateUser(ctx context.Context, userId int64) UpdateUserRequest {
	return UpdateUserRequest{
		apiService: s,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
func (s *usersService) UpdateUserExecute(r UpdateUserRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/{user_id}"
	)

	endpoint = strings.Replace(endpoint, "{user_id}", IdToString(r.userId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "full_name", r.fullName)
	AddOptionalParam(form, "role", r.role)
	AddOptionalParam(form, "profile_data", r.profileData)
	AddOptionalParam(form, "new_email", r.newEmail)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateUserByEmailRequest struct {
	ctx         context.Context
	apiService  APIUsers
	email       string
	fullName    *string
	role        *int32
	profileData *[]map[string]interface{}
	newEmail    *string
}

// The user's full name.
//
// **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 5.0 (feature level 106).
func (r UpdateUserByEmailRequest) FullName(fullName string) UpdateUserByEmailRequest {
	r.fullName = &fullName
	return r
}

// New [role] for the user. Roles are encoded as:  - Organization owner: 100 - Organization administrator: 200 - Organization moderator: 300 - Member: 400 - Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.
//
// **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of `is_admin` and `is_guest` boolean parameters. Organization moderator role added in Zulip 4.0 (feature level 60).
//
// [role]: https://zulip.com/api/roles-and-permissions
func (r UpdateUserByEmailRequest) Role(role int32) UpdateUserByEmailRequest {
	r.role = &role
	return r
}

// A dictionary containing the updated custom profile field data for the user.
func (r UpdateUserByEmailRequest) ProfileData(profileData []map[string]interface{}) UpdateUserByEmailRequest {
	r.profileData = &profileData
	return r
}

// New email address for the user. Requires the user making the request to be an organization owner and additionally have the `.can_change_user_emails` special permission.
//
// **Changes**: New in Zulip 10.0 (feature level 285).
func (r UpdateUserByEmailRequest) NewEmail(newEmail string) UpdateUserByEmailRequest {
	r.newEmail = &newEmail
	return r
}

func (r UpdateUserByEmailRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateUserByEmailExecute(r)
}

// UpdateUserByEmail Update a user by email
//
// Administrative endpoint to update the details of another user in the organization by their email address.
// Works the same way as [`PATCH /users/{user_id}`] but fetching the target user by their
// real email address.
//
// The requester needs to have permission to view the target user's real email address, subject to the
// user's email address visibility setting. Otherwise, the dummy address of the format
// `user{id}@{realm.host}` needs be used. This follows the same rules as `GET /users/{email}`.
//
// *Changes**: New in Zulip 10.0 (feature level 313).
//
// [`PATCH /users/{user_id}`]: https://zulip.com/api/update-user
func (s *usersService) UpdateUserByEmail(ctx context.Context, email string) UpdateUserByEmailRequest {
	return UpdateUserByEmailRequest{
		apiService: s,
		ctx:        ctx,
		email:      email,
	}
}

// Execute executes the request
func (s *usersService) UpdateUserByEmailExecute(r UpdateUserByEmailRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/{email}"
	)

	endpoint = strings.Replace(endpoint, "{email}", url.PathEscape(r.email), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "full_name", r.fullName)
	AddOptionalParam(form, "role", r.role)
	AddOptionalParam(form, "profile_data", r.profileData)
	AddOptionalParam(form, "new_email", r.newEmail)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateUserGroupRequest struct {
	ctx                   context.Context
	apiService            APIUsers
	userGroupId           int64
	name                  *string
	description           *string
	canAddMembersGroup    *zulip.GroupSettingValueUpdate
	canJoinGroup          *zulip.GroupSettingValueUpdate
	canLeaveGroup         *zulip.GroupSettingValueUpdate
	canManageGroup        *zulip.GroupSettingValueUpdate
	canMentionGroup       *zulip.GroupSettingValueUpdate
	canRemoveMembersGroup *zulip.GroupSettingValueUpdate
	deactivated           *bool
}

// The new name of the group.
//
// **Changes**: Before Zulip 7.0 (feature level 165), this was a required field.
func (r UpdateUserGroupRequest) Name(name string) UpdateUserGroupRequest {
	r.name = &name
	return r
}

// The new description of the group.
//
// **Changes**: Before Zulip 7.0 (feature level 165), this was a required field.
func (r UpdateUserGroupRequest) Description(description string) UpdateUserGroupRequest {
	r.description = &description
	return r
}

// The set of users who have permission to add members to this user group expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the `can_manage_group` setting.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values
//
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
func (r UpdateUserGroupRequest) CanAddMembersGroup(canAddMembersGroup zulip.GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canAddMembersGroup = &canAddMembersGroup
	return r
}

// The set of users who have permission to join this user group expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 10.0 (feature level 301).
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values
//
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
func (r UpdateUserGroupRequest) CanJoinGroup(canJoinGroup zulip.GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canJoinGroup = &canJoinGroup
	return r
}

// The set of users who have permission to leave this user group expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 10.0 (feature level 308).
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values
//
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
func (r UpdateUserGroupRequest) CanLeaveGroup(canLeaveGroup zulip.GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canLeaveGroup = &canLeaveGroup
	return r
}

// The set of users who have permission to [manage this user group] expressed as an [update to a group-setting value].  This setting cannot be set to `"role:internet"` and `"role:everyone"` [system groups].
//
// **Changes**: New in Zulip 10.0 (feature level 283).
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
// [manage this user group]: https://zulip.com/help/manage-user-groups
func (r UpdateUserGroupRequest) CanManageGroup(canManageGroup zulip.GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canManageGroup = &canManageGroup
	return r
}

// The set of users who have permission to [mention this group], expressed as an [update to a group-setting value].  This setting cannot be set to `"role:internet"` and `"role:owners"` [system groups].
//
// **Changes**: In Zulip 9.0 (feature level 260), this parameter was updated to only accept an object with the `old` and `new` fields described below. Prior to this feature level, this parameter could be either of the two forms of a [group-setting value].  Before Zulip 9.0 (feature level 258), this parameter could only be the integer form of a [group-setting value].  Before Zulip 8.0 (feature level 198), this parameter was named `can_mention_group_id`.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups].
//
// [mention this group]: https://zulip.com/help/mention-a-user-or-group
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
// [group-setting value]: https://zulip.com/api/group-setting-values
func (r UpdateUserGroupRequest) CanMentionGroup(canMentionGroup zulip.GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canMentionGroup = &canMentionGroup
	return r
}

// The set of users who have permission to remove members from this user group expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the `can_manage_group` setting.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values
//
// [system groups]: https://zulip.com/api/group-setting-values#system-groups
func (r UpdateUserGroupRequest) CanRemoveMembersGroup(canRemoveMembersGroup zulip.GroupSettingValueUpdate) UpdateUserGroupRequest {
	r.canRemoveMembersGroup = &canRemoveMembersGroup
	return r
}

// A deactivated user group can be reactivated by passing this parameter as `false`.  Passing `true` does nothing as user group is deactivated using [`POST /user_groups/{user_group_id}/deactivate`] endpoint.
//
// **Changes**: New in Zulip 11.0 (feature level 386).
//
// [`POST /user_groups/{user_group_id}/deactivate`]: deactivate-user-group
func (r UpdateUserGroupRequest) Deactivated(deactivated bool) UpdateUserGroupRequest {
	r.deactivated = &deactivated
	return r
}

func (r UpdateUserGroupRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateUserGroupExecute(r)
}

// UpdateUserGroup Update a user group
//
// Update the name, description or any of the permission settings
// of a [user group].
//
// This endpoint is also used to reactivate a user group.
//
// Note that while permissions settings of deactivated groups can
// be edited by this API endpoint, and those permissions settings
// do affect the ability to modify the deactivated group and its
// membership, the deactivated group itself cannot be mentioned
// or used in the value of any permission without first being reactivated.
//
// *Changes**: Starting with Zulip 11.0 (feature level 386), this
// endpoint can be used to reactivate a user group.
//
// Prior to Zulip 10.0 (feature level 340), only the name field
// of deactivated groups could be modified.
//
// [user group]: https://zulip.com/help/user-groups
func (s *usersService) UpdateUserGroup(ctx context.Context, userGroupId int64) UpdateUserGroupRequest {
	return UpdateUserGroupRequest{
		apiService:  s,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
func (s *usersService) UpdateUserGroupExecute(r UpdateUserGroupRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/user_groups/{user_group_id}"
	)

	endpoint = strings.Replace(endpoint, "{user_group_id}", IdToString(r.userGroupId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "name", r.name)
	AddOptionalParam(form, "description", r.description)
	if err := AddOptionalJSONParam(form, "can_add_members_group", r.canAddMembersGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_join_group", r.canJoinGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_leave_group", r.canLeaveGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_manage_group", r.canManageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_mention_group", r.canMentionGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_remove_members_group", r.canRemoveMembersGroup); err != nil {
		return nil, nil, err
	}
	AddOptionalParam(form, "deactivated", r.deactivated)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateUserGroupMembersRequest struct {
	ctx             context.Context
	apiService      APIUsers
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
func (r UpdateUserGroupMembersRequest) Add(add ...int64) UpdateUserGroupMembersRequest {
	r.add = &add
	return r
}

// The list of user group Ids to be removed from the user group.
//
// **Changes**: New in Zulip 10.0 (feature level 311).
func (r UpdateUserGroupMembersRequest) DeleteSubgroups(deleteSubgroups []int64) UpdateUserGroupMembersRequest {
	r.deleteSubgroups = &deleteSubgroups
	return r
}

// The list of user group Ids to be added to the user group.
//
// **Changes**: New in Zulip 10.0 (feature level 311).
func (r UpdateUserGroupMembersRequest) AddSubgroups(addSubgroups []int64) UpdateUserGroupMembersRequest {
	r.addSubgroups = &addSubgroups
	return r
}

func (r UpdateUserGroupMembersRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateUserGroupMembersExecute(r)
}

// UpdateUserGroupMembers Update user group members
//
// Update the members of a [user group]. The
// user Ids must correspond to non-deactivated users.
//
// *Changes**: Prior to Zulip 11.0 (feature level 391), members
// could not be added or removed from a deactivated group.
//
// *Changes**: Prior to Zulip 10.0 (feature level 303), group memberships of
// deactivated users were visible to the API and could be edited via this endpoint.
//
// [user group]: https://zulip.com/help/user-groups
func (s *usersService) UpdateUserGroupMembers(ctx context.Context, userGroupId int64) UpdateUserGroupMembersRequest {
	return UpdateUserGroupMembersRequest{
		apiService:  s,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
func (s *usersService) UpdateUserGroupMembersExecute(r UpdateUserGroupMembersRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/user_groups/{user_group_id}/members"
	)

	endpoint = strings.Replace(endpoint, "{user_group_id}", IdToString(r.userGroupId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "delete", r.delete)
	AddOptionalParam(form, "add", r.add)
	AddOptionalParam(form, "delete_subgroups", r.deleteSubgroups)
	AddOptionalParam(form, "add_subgroups", r.addSubgroups)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateUserGroupSubgroupsRequest struct {
	ctx         context.Context
	apiService  APIUsers
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

func (r UpdateUserGroupSubgroupsRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateUserGroupSubgroupsExecute(r)
}

// UpdateUserGroupSubgroups Update subgroups of a user group
//
// Update the subgroups of a [user group].
//
// *Changes**: Prior to Zulip 11.0 (feature level 391), subgroups
// could not be added or removed from a deactivated group.
//
// *Changes**: New in Zulip 6.0 (feature level 127).
//
// [user group]: https://zulip.com/help/user-groups
func (s *usersService) UpdateUserGroupSubgroups(ctx context.Context, userGroupId int64) UpdateUserGroupSubgroupsRequest {
	return UpdateUserGroupSubgroupsRequest{
		apiService:  s,
		ctx:         ctx,
		userGroupId: userGroupId,
	}
}

// Execute executes the request
func (s *usersService) UpdateUserGroupSubgroupsExecute(r UpdateUserGroupSubgroupsRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/user_groups/{user_group_id}/subgroups"
	)

	endpoint = strings.Replace(endpoint, "{user_group_id}", IdToString(r.userGroupId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "delete", r.delete)
	AddOptionalParam(form, "add", r.add)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}
