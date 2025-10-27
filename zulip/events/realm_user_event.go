package events

import "github.com/tum-zulip/go-zulip/zulip"

// RealmUserAddEvent Event sent to all users in a Zulip organization when a new user joins or when a guest user gains access to a user. Processing this event is important to being able to display basic details on other users given only their Id.  If the current user is a guest whose access to a newly created user is limited by a `can_access_all_users_group` policy, and the event queue was registered with the `user_list_incomplete` client capability, then the event queue will not receive an event for such a new user. If a newly created user is inaccessible to the current user via such a policy, but the client lacks `user_list_incomplete` client capability, then this event will be delivered to the queue, with an "Unknown user" object with the usual format but placeholder data whose only variable content is the user Id.
//
// **Changes**: Before Zulip 8.0 (feature level 232), the `user_list_incomplete` client capability did not exist, and so all clients whose access to a new user was prevented by `can_access_all_users_group` policy would receive a fake "Unknown user" event for such a user.  Starting with Zulip 8.0 (feature level 228), this event is also sent when a guest user gains access to a user.
type RealmUserAddEvent struct {
	event

	Person zulip.User `json:"person,omitempty"`
}

// RealmUserRemoveEvent Event sent to guest users when they lose access to a user.
//
// **Changes**: As of Zulip 8.0 (feature level 228), this event is no longer deprecated.  In Zulip 8.0 (feature level 222), this event was deprecated and no longer sent to clients. Prior to this feature level, it was sent to all users in a Zulip organization when a user was deactivated.
type RealmUserRemoveEvent struct {
	event

	Person UserInfo `json:"person,omitempty"`
}

// UserInfo Object containing details of the deactivated user.
type UserInfo struct {
	// The Id of the deactivated user.
	UserID int64 `json:"user_id,omitempty"`
	// The full name of the user.  **Deprecated**: We expect to remove this field in the future.
	// Deprecated
	FullName *string `json:"full_name,omitempty"`
}

// RealmUserUpdateEvent Event sent generally to all users who can access the modified user for changes in the set of users or those users metadata.
//
// **Changes**: Prior to Zulip 8.0 (feature level 228), this event was sent to all users in the organization.
type RealmUserUpdateEvent struct {
	event

	Person UserUpdate `json:"person,omitempty"`
}

// UserUpdate - Object containing the changed details of the user. It has multiple forms depending on the value changed.
//
// **Changes**: Removed `is_billing_admin` field in Zulip 10.0 (feature level 363), as it was replaced by the `can_manage_billing_group` realm setting.
type UserUpdate struct {
	// The Id of the user who is affected by this change.
	UserID int64 `json:"user_id,omitempty"`

	UserUpdateEventFullName      *UserUpdateEventFullName
	UserUpdateEventAvatar        *UserUpdateEventAvatar
	UserUpdateEventTimezone      *UserUpdateEventTimezone
	UserUpdateEventBotOwner      *UserUpdateEventBotOwner
	UserUpdateEventRole          *UserUpdateEventRole
	UserUpdateEventDeliveryEmail *UserUpdateEventDeliveryEmail
	UserUpdateEventCustomField   *UserUpdateEventCustomField
	UserUpdateEventEmail         *UserUpdateEventEmail
	UserUpdateEventActivation    *UserUpdateEventActivation
}

// UserUpdateEventAvatar When a user changes their avatar.
type UserUpdateEventAvatar struct {
	zulip.Avatar

	// The version number for the user's avatar. This is useful for cache-busting.
	AvatarVersion int64 `json:"avatar_version,omitempty"`
}

// UserUpdateEventFullName When a user changes their full name.
type UserUpdateEventFullName struct {
	// The new full name for the user.
	FullName string `json:"full_name,omitempty"`
}

// UserUpdateEventTimezone When a user changes their [profile time zone].
//
// [profile time zone]: https://zulip.com/help/change-your-timezone
type UserUpdateEventTimezone struct {
	// The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the `user_id`.
	// Deprecated
	Email string `json:"email,omitempty"`
	// The IANA identifier of the new profile time zone for the user.
	Timezone string `json:"timezone,omitempty"`
}

// UserUpdateEventBotOwner When the owner of a bot changes.
type UserUpdateEventBotOwner struct {
	// The user Id of the new bot owner.
	BotOwnerId int64 `json:"bot_owner_id,omitempty"`
}

// UserUpdateEventRole When the [role] of a user changes.
//
// [role]: https://zulip.com/help/user-roles
type UserUpdateEventRole struct {
	// The new [role] of the user.
	//
	// [role]: https://zulip.com/api/roles-and-permissions
	Role zulip.Role `json:"role,omitempty"`
}

// UserUpdateEventDeliveryEmail When the value of a user's delivery email as visible to you changes, either due to the email address changing or your access to the user's email changing via an update to their `email_address_visibility` setting.
//
// **Changes**: Prior to Zulip 7.0 (feature level 163), this event was sent only to the affected user, and this event would only be triggered by changing the affected user's delivery email.
type UserUpdateEventDeliveryEmail struct {
	// The new delivery email of the user.  This value can be `null` if the affected user changed their `email_address_visibility` setting such that you cannot access their real email.
	//
	// **Changes**: Before Zulip 7.0 (feature level 163), `null` was not a possible value for this event as it was only sent to the affected user when their email address was changed.
	DeliveryEmail *string `json:"delivery_email,omitempty"`
}

// UserUpdateEventCustomField When the user updates one of their custom profile fields.
type UserUpdateEventCustomField struct {
	CustomProfileField zulip.ProfileDataValue `json:"custom_profile_field,omitempty"`
}

// UserUpdateEventEmail When the Zulip API email address of a user changes, either due to the user's email address changing, or due to changes in the user's [email address visibility].
//
// [email address visibility]: https://zulip.com/help/configure-email-visibility
type UserUpdateEventEmail struct {
	// The new value of `email` for the user. The client should update any data structures associated with this user to use this new value as the user's Zulip API email address.
	NewEmail string `json:"new_email,omitempty"`
}

// UserUpdateEventActivation When a user is deactivated or reactivated. Only users who can access the modified user under the organization's `can_access_all_users_group` policy will receive this event.  Clients receiving a deactivation event should remove the user from all user groups in their data structures, because deactivated users cannot be members of groups.
//
// **Changes**: Prior to Zulip 10.0 (feature level 303), reactivation events were sent to users who could not access the reactivated user due to a `can_access_all_users_group` policy. Also, previously, Clients were not required to update group membership records during user deactivation.  New in Zulip 8.0 (feature level 222). Previously the server sent a `realm_user` event with `op` field set to `remove` when deactivating a user and a `realm_user` event with `op` field set to `add` when reactivating a user.
type UserUpdateEventActivation struct {
	// A boolean describing whether the user account has been deactivated.
	IsActive bool `json:"is_active,omitempty"`
}
