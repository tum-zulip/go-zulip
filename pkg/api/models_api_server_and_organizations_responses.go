package api

// AddCodePlaygroundResponse struct for AddCodePlaygroundResponse
type AddCodePlaygroundResponse struct {
	Response

	// The numeric Id assigned to this playground.
	Id int64 `json:"id,omitempty"`
}

// AddLinkifierResponse struct for AddLinkifierResponse
type AddLinkifierResponse struct {
	Response

	// The numeric Id assigned to this filter.
	Id int64 `json:"id,omitempty"`
}

// CreateCustomProfileFieldResponse struct for CreateCustomProfileFieldResponse
type CreateCustomProfileFieldResponse struct {
	Response

	// The Id for the custom profile field.
	Id int64 `json:"id,omitempty"`
}

// ExportRealmResponse struct for ExportRealmResponse
type ExportRealmResponse struct {
	Response

	// The Id of the data export created.  **Changes**: New in Zulip 7.0 (feature level 182).
	Id int64 `json:"id,omitempty"`
}

// GetCustomEmojiResponse struct for GetCustomEmojiResponse
type GetCustomEmojiResponse struct {
	Response

	// An object that contains `emoji` objects, each identified with their emoji Id as the key.
	Emoji map[string]RealmEmoji `json:"emoji,omitempty"`
}

// GetCustomProfileFieldsResponse struct for GetCustomProfileFieldsResponse
type GetCustomProfileFieldsResponse struct {
	Response

	// An array containing all the custom profile fields defined in this Zulip organization.
	CustomFields []CustomProfileField `json:"custom_fields,omitempty"`
}

// GetLinkifiersResponse struct for GetLinkifiersResponse
type GetLinkifiersResponse struct {
	Response

	// An ordered array of objects, where each object describes a linkifier.  Clients should always process linkifiers in the order given; this is important if the realm has linkifiers with overlapping patterns. The order can be modified using [`PATCH /realm/linkifiers`](zulip.com/api/reorder-linkifiers.
	Linkifiers []RealmLinkifiers `json:"linkifiers,omitempty"`
}

// GetPresenceResponse struct for GetPresenceResponse
type GetPresenceResponse struct {
	Response

	// The time when the server fetched the `presences` data included in the response.
	ServerTimestamp float32 `json:"server_timestamp,omitempty"`
	// A dictionary where each entry describes the presence details of a user in the Zulip organization.
	Presences map[string]map[string]LegacyPresenceFormat `json:"presences,omitempty"`
}

// GetRealmExportConsentsResponse struct for GetRealmExportConsentsResponse
type GetRealmExportConsentsResponse struct {
	Response

	// An array of objects where each object contains a user Id and whether the user has consented for their private data to be exported.
	ExportConsents []ExportConsent `json:"export_consents,omitempty"`
}

// ExportConsent struct for ExportConsent
type ExportConsent struct {
	// The user Id.
	UserId int64 `json:"user_id,omitempty"`
	// Whether the user has consented for their private data export.
	Consented bool `json:"consented,omitempty"`
}

// GetRealmExportsResponse struct for GetRealmExportsResponse
type GetRealmExportsResponse struct {
	Response

	// An array of dictionaries where each dictionary contains details about a data export of the organization.
	Exports []RealmExport `json:"exports,omitempty"`
}

// TestWelcomeBotCustomMessageResponse struct for TestWelcomeBotCustomMessageResponse
type TestWelcomeBotCustomMessageResponse struct {
	Response

	// The message_id of the test welcome bot custom message.
	MessageId int64 `json:"message_id,omitempty"`
}

// GetServerSettingsResponse struct for GetServerSettingsResponse
type GetServerSettingsResponse struct {
	Response

	// Deprecated
	AuthenticationMethods *AuthMethods `json:"authentication_methods,omitempty"`
	// A list of dictionaries describing the available external authentication methods (E.g. Google, GitHub, or SAML) enabled for this organization.  The list is sorted in the order in which these authentication methods should be displayed.  **Changes**: New in Zulip 2.1.0.
	ExternalAuthenticationMethods []ExternalAuthMethod `json:"external_authentication_methods,omitempty"`
	// An integer indicating what features are available on the server. The feature level increases monotonically; a value of N means the server supports all API features introduced before feature level N. This is designed to provide a simple way for client apps to decide whether the server supports a given feature or API change. See the [changelog](zulip.com/api/changelog) for details on what each feature level means.  **Changes**: New in Zulip 3.0 (feature level 1). We recommend using an implied value of 0 for Zulip servers that do not send this field.
	ZulipFeatureLevel int `json:"zulip_feature_level,omitempty"`
	// The server's version number. This is often a release version number, like `2.1.7`. But for a server running a [version from Git][git-release], it will be a Git reference to the commit, like `5.0-dev-1650-gc3fd37755f`.  [git-release]: https://zulip.readthedocs.io/en/latest/overview/release-lifecycle.html#git-versions
	ZulipVersion string `json:"zulip_version,omitempty"`
	// The `git merge-base` between `zulip_version` and official branches in the public [Zulip server and web app repository](https://github.com/zulip/zulip), in the same format as `zulip_version`. This will equal `zulip_version` if the server is not running a fork of the Zulip server.  This will be `\"\"` if unavailable.  **Changes**: New in Zulip 5.0 (feature level 88).
	ZulipMergeBase string `json:"zulip_merge_base,omitempty"`
	// Whether mobile/push notifications are configured.
	PushNotificationsEnabled bool `json:"push_notifications_enabled,omitempty"`
	// Whether the Zulip client that has sent a request to this endpoint is deemed incompatible with the server.
	IsIncompatible bool `json:"is_incompatible,omitempty"`
	// Setting for allowing users authenticate with an email-password combination.
	EmailAuthEnabled bool `json:"email_auth_enabled,omitempty"`
	// Whether all valid usernames for authentication to this organization will be email addresses. This is important for clients to know whether to do client side validation of email address format in a login prompt.  This value will be false if the server has [LDAP authentication][ldap-auth] enabled with a username and password combination.  [ldap-auth]: https://zulip.readthedocs.io/en/latest/production/authentication-methods.html#ldap-including-active-directory
	RequireEmailFormatUsernames bool `json:"require_email_format_usernames,omitempty"`
	// The organization's canonical URL. Alias of `realm_url`.  **Changes**: Deprecated in Zulip 9.0 (feature level 257). The term \"URI\" is deprecated in [web standards](https://url.spec.whatwg.org/#goals).
	// Deprecated
	RealmUri *string `json:"realm_uri,omitempty"`
	// The organization's canonical URL.  **Changes**: New in Zulip 9.0 (feature level 257), replacing the deprecated `realm_uri`.
	RealmUrl string `json:"realm_url,omitempty"`
	// The organization's name (for display purposes).
	RealmName string `json:"realm_name,omitempty"`
	// The URL for the organization's logo formatted as a square image, used for identifying the organization in small locations in the mobile and desktop apps.
	RealmIcon string `json:"realm_icon,omitempty"`
	// HTML description of the organization, as configured by the [organization profile](zulip.com/help/create-your-organization-profile.
	RealmDescription string `json:"realm_description,omitempty"`
	// Whether the organization has enabled the creation of [web-public channels](zulip.com/help/public-access-option) and at least one web-public channel on the server currently exists. Clients that support viewing content in web-public channels without an account can use this to determine whether to offer that feature on the login page for an organization.  **Changes**: New in Zulip 5.0 (feature level 116).
	RealmWebPublicAccessEnabled bool `json:"realm_web_public_access_enabled,omitempty"`
}

// AuthMethods Each key-value pair in the object indicates whether the authentication method is enabled on this server.  **Changes**: Deprecated in Zulip 2.1.0, in favor of the more expressive `external_authentication_methods`.
type AuthMethods struct {
	// Whether the user can authenticate using password.
	Password bool `json:"password,omitempty"`
	// Whether the user can authenticate using development API key.
	Dev bool `json:"dev,omitempty"`
	// Whether the user can authenticate using email.
	Email bool `json:"email,omitempty"`
	// Whether the user can authenticate using LDAP.
	Ldap bool `json:"ldap,omitempty"`
	// Whether the user can authenticate using REMOTE_USER.
	Remoteuser bool `json:"remoteuser,omitempty"`
	// Whether the user can authenticate using their GitHub account.
	Github bool `json:"github,omitempty"`
	// Whether the user can authenticate using their Microsoft Entra Id account.
	Azuread bool `json:"azuread,omitempty"`
	// Whether the user can authenticate using their GitLab account.  **Changes**: New in Zulip 3.0 (feature level 1).
	Gitlab bool `json:"gitlab,omitempty"`
	// Whether the user can authenticate using their Apple account.
	Apple bool `json:"apple,omitempty"`
	// Whether the user can authenticate using their Google account.
	Google bool `json:"google,omitempty"`
	// Whether the user can authenticate using SAML.
	Saml bool `json:"saml,omitempty"`
	// Whether the user can authenticate using OpenId Connect.
	OpenidConnect bool `json:"openid connect,omitempty"`
}

// ExternalAuthMethod struct for ExternalAuthMethod
type ExternalAuthMethod struct {
	// A unique, table, machine-readable name for the authentication method, intended to be used by clients with special behavior for specific authentication methods to correctly identify the method.
	Name string `json:"name,omitempty"`
	// Display name of the authentication method, to be used in all buttons for the authentication method.
	DisplayName string `json:"display_name,omitempty"`
	// URL for an image to be displayed as an icon in all buttons for the external authentication method.  When `null`, no icon should be displayed.
	DisplayIcon *string `json:"display_icon,omitempty"`
	// URL to be used to initiate authentication using this method.
	LoginUrl string `json:"login_url,omitempty"`
	// URL to be used to initiate account registration using this method.
	SignupUrl string `json:"signup_url,omitempty"`
}
