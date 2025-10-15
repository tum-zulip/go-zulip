package zulip

// RealmEmoji `{emoji_id}`: Object containing details about the emoji with the specified Id. It has the following properties:
type RealmEmoji struct {
	// The Id for this emoji, same as the object's key.
	Id string `json:"id,omitempty"`
	// The user-friendly name for this emoji. Users in the organization can use this emoji by writing this name between colons (`:name :`).
	Name string `json:"name,omitempty"`
	// The path relative to the organization's URL where the emoji's image can be found.
	SourceUrl string `json:"source_url,omitempty"`
	// Only non-null when the emoji's image is animated.  The path relative to the organization's URL where a still (not animated) version of the emoji can be found. (This is currently always the first frame of the animation).  This is useful for clients to display the emoji in contexts where continuously animating it would be a bad user experience (E.g. because it would be distracting).  **Changes**: New in Zulip 5.0 (added as optional field in feature level 97 and then made mandatory, but nullable, in feature level 113).
	StillUrl *string `json:"still_url,omitempty"`
	// Whether the emoji has been deactivated or not.
	Deactivated bool `json:"deactivated,omitempty"`
	// The user Id of the user who uploaded the custom emoji. Will be `null` if the uploader is unknown.  **Changes**: New in Zulip 3.0 (feature level 7). Previously was accessible via an `author` object with an `id` field.
	AuthorId *int64 `json:"author_id,omitempty"`
}
