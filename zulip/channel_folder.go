package zulip

import (
	"encoding/json"
	"time"
)

type ChannelFolder struct {
	// The name of the channel folder.
	Name string `json:"name,omitempty"`
	// The UNIX timestamp for when the channel folder was created, in UTC seconds.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ID of the user who created this channel folder.
	CreatorID *int64 `json:"creator_id,omitempty"`
	// The description of the channel folder.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	Description string `json:"description,omitempty"`
	// The description of the channel folder rendered as HTML, intended to be used when displaying the channel folder description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.
	RenderedDescription string `json:"rendered_description,omitempty"`
	// This value determines in which order the channel folders will be displayed in the UI. The value is 0 indexed, and the value with the lower order will be displayed first.
	//
	// **Changes**: New in Zulip 11.0 (feature level 414).
	Order int32 `json:"order,omitempty"`
	// The ID of the channel folder.
	ID int64 `json:"id,omitempty"`
	// Whether the channel folder is archived or not.
	IsArchived bool `json:"is_archived,omitempty"`
}

type channelFolderJSON struct {
	Name                string `json:"name,omitempty"`
	DateCreated         *int64 `json:"date_created,omitempty"`
	CreatorID           *int64 `json:"creator_id,omitempty"`
	Description         string `json:"description,omitempty"`
	RenderedDescription string `json:"rendered_description,omitempty"`
	Order               int32  `json:"order,omitempty"`
	ID                  int64  `json:"id,omitempty"`
	IsArchived          bool   `json:"is_archived,omitempty"`
}

func (o ChannelFolder) MarshalJSON() ([]byte, error) {
	aux := channelFolderJSON{
		Name:                o.Name,
		CreatorID:           o.CreatorID,
		Description:         o.Description,
		RenderedDescription: o.RenderedDescription,
		Order:               o.Order,
		ID:                  o.ID,
		IsArchived:          o.IsArchived,
	}

	if o.DateCreated != nil {
		ts := o.DateCreated.Unix()
		aux.DateCreated = &ts
	}

	return json.Marshal(aux)
}

func (o *ChannelFolder) UnmarshalJSON(data []byte) error {
	var aux channelFolderJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.Name = aux.Name
	o.CreatorID = aux.CreatorID
	o.Description = aux.Description
	o.RenderedDescription = aux.RenderedDescription
	o.Order = aux.Order
	o.ID = aux.ID
	o.IsArchived = aux.IsArchived

	if aux.DateCreated != nil {
		t := time.Unix(*aux.DateCreated, 0).UTC()
		o.DateCreated = &t
	} else {
		o.DateCreated = nil
	}

	return nil
}
