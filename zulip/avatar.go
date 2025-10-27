package zulip

type Avatar struct {
	// The URL of the new avatar for the user.
	AvatarURL string `json:"avatar_url,omitempty"`
	// The new avatar data source type for the user.
	// Value values are:
	//   - AvatarSourceGravatar = Gravatar avatar
	//   - AvatarSourceUploaded = Uploaded by user
	AvatarSource AvatarSource `json:"avatar_source,omitempty"`
	// The new medium-size avatar URL for user.
	AvatarURLMedium string `json:"avatar_url_medium,omitempty"`
}
