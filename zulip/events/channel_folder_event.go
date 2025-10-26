package events

import "github.com/tum-zulip/go-zulip/zulip"

// ChannelFolderAddEvent Event sent to users in an organization when a channel folder is created.
//
// **Changes**: New in Zulip 11.0 (feature level 389).
type ChannelFolderAddEvent struct {
	event

	ChannelFolder zulip.ChannelFolder `json:"channel_folder,omitempty"`
}

// ChannelFolderUpdateEvent Event sent to users in an organization when a channel folder is updated.
//
// **Changes**: New in Zulip 11.0 (feature level 389).
type ChannelFolderUpdateEvent struct {
	event

	// Id of the updated channel folder.
	ChannelFolderId int64            `json:"channel_folder_id,omitempty"`
	Data            FolderUpdateData `json:"data,omitempty"`
}

// FolderUpdateData Dictionary containing the changed details of the channel folder.
type FolderUpdateData struct {
	// The new name of the channel folder. Only present if the channel folder's name changed.
	Name *string `json:"name,omitempty"`
	// The new description of the channel folder. Only present if the description changed.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	Description *string `json:"description,omitempty"`
	// The new rendered description of the channel folder. Only present if the description changed.
	RenderedDescription *string `json:"rendered_description,omitempty"`
	// Whether the channel folder is archived or not. Only present if the channel folder is archived or unarchived.
	IsArchived *bool `json:"is_archived,omitempty"`
}

// ChannelFolderReorderEvent Event sent to users in an organization when channel folders are reordered.
//
// **Changes**: New in Zulip 11.0 (feature level 418).
type ChannelFolderReorderEvent struct {
	event

	// A list of channel folder Ids representing the new order.
	Order []int64 `json:"order,omitempty"`
}
