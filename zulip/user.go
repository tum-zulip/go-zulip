package zulip

// User struct for User
type User struct {
	// The unique ID of the user.
	UserID        int64   `json:"user_id,omitempty"`
	DeliveryEmail *string `json:"delivery_email,omitempty"`
	// The Zulip API email address of the user or bot.  If you do not have permission to view the email address of the target user, this will be a fake email address that is usable for the Zulip API but nothing else.
	Email string `json:"email,omitempty"`
	// Full name of the user or bot, used for all display purposes.
	FullName string `json:"full_name,omitempty"`
	// The time the user account was created.
	DateJoined string `json:"date_joined,omitempty"`
	// A boolean specifying whether the user account has been deactivated.
	IsActive bool `json:"is_active,omitempty"`
	// A boolean specifying whether the user is an organization owner. If true, `is_admin` will also be true.
	//
	// **Changes**: New in Zulip 3.0 (feature level 8).
	IsOwner bool `json:"is_owner,omitempty"`
	// A boolean specifying whether the user is an organization administrator.
	IsAdmin bool `json:"is_admin,omitempty"`
	// A boolean specifying whether the user is a guest user.
	IsGuest bool `json:"is_guest,omitempty"`
	// A boolean specifying whether the user is a bot or full account.
	IsBot bool `json:"is_bot,omitempty"`

	// [Organization-level role] of the user. Possible values are:
	//   - RoleOwner
	//   - RoleAdmin
	//   - RoleModerator
	//   - RoleMember
	//   - RoleGuest
	//
	// **Changes**: New in Zulip 4.0 (feature level 59).
	//
	// [Organization-level role]: https://zulip.com/api/roles-and-permissions
	Role Role `json:"role,omitempty"`
	// The IANA identifier of the user's [profile time zone], which is used primarily to display the user's local time to other users.
	//
	// [profile time zone]: https://zulip.com/help/change-your-timezone
	Timezone  string  `json:"timezone,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// Version for the user's avatar. Used for cache-busting requests for the user's avatar. Clients generally shouldn't need to use this; most avatar URLs sent by Zulip will already end with `?v={avatar_version}`.
	AvatarVersion int32 `json:"avatar_version,omitempty"`
	// Only present if `is_bot` is false; bots can't have custom profile fields.  A dictionary containing custom profile field data for the user. Each entry maps the integer ID of a custom profile field in the organization to a dictionary containing the user's data for that field. Generally the data includes just a single `value` key; for those custom profile fields supporting Markdown, a `rendered_value` key will also be present.
	ProfileData map[int64]ProfileDataValue `json:"profile_data,omitempty"`

	BotType    *BotType `json:"bot_type,omitempty"`
	BotOwnerID *int64   `json:"bot_owner_id,omitempty"`
	// Whether the user is a system bot. System bots are special bot user accounts that are managed by the system, rather than the organization's administrators.
	//
	// **Changes**: This field was called `is_cross_realm_bot` before Zulip 5.0 (feature level 83).
	IsSystemBot *bool `json:"is_system_bot,omitempty"`
}

// ProfileDataValue `{id}`: Object with data about what value the user filled in the custom profile field with that ID.
type ProfileDataValue struct {
	// The ID of the custom profile field which user updated.
	ID int64 `json:"id,omitempty"`
	// User's personal value for this custom profile field.
	Value string `json:"value,omitempty"`
	// The `value` rendered in HTML. Will only be present for custom profile field types that support Markdown rendering.  This user-generated HTML content should be rendered using the same CSS and client-side security protections as are used for message content.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	RenderedValue *string `json:"rendered_value,omitempty"`
}
