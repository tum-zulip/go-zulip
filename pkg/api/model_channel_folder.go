package api

import (
	"encoding/json"
	"time"
)

type ChannelFolder struct {
	// The name of the channel folder.
	Name string `json:"name,omitempty"`
	// The UNIX timestamp for when the channel folder was created, in UTC seconds.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The Id of the user who created this channel folder.
	CreatorId *int64 `json:"creator_id,omitempty"`
	// The description of the channel folder.  See [Markdown message formatting](zulip.com/api/message-formatting) for details on Zulip's HTML format.
	Description string `json:"description,omitempty"`
	// The description of the channel folder rendered as HTML, intended to be used when displaying the channel folder description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.
	RenderedDescription string `json:"rendered_description,omitempty"`
	// This value determines in which order the channel folders will be displayed in the UI. The value is 0 indexed, and the value with the lower order will be displayed first.  **Changes**: New in Zulip 11.0 (feature level 414).
	Order int32 `json:"order,omitempty"`
	// The Id of the channel folder.
	Id int64 `json:"id,omitempty"`
	// Whether the channel folder is archived or not.
	IsArchived bool `json:"is_archived,omitempty"`
}

type channelFolderJSON struct {
	Name                string `json:"name,omitempty"`
	DateCreated         *int64 `json:"date_created,omitempty"`
	CreatorId           *int64 `json:"creator_id,omitempty"`
	Description         string `json:"description,omitempty"`
	RenderedDescription string `json:"rendered_description,omitempty"`
	Order               int32  `json:"order,omitempty"`
	Id                  int64  `json:"id,omitempty"`
	IsArchived          bool   `json:"is_archived,omitempty"`
}

func (o ChannelFolder) MarshalJSON() ([]byte, error) {
	aux := channelFolderJSON{
		Name:                o.Name,
		CreatorId:           o.CreatorId,
		Description:         o.Description,
		RenderedDescription: o.RenderedDescription,
		Order:               o.Order,
		Id:                  o.Id,
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
	o.CreatorId = aux.CreatorId
	o.Description = aux.Description
	o.RenderedDescription = aux.RenderedDescription
	o.Order = aux.Order
	o.Id = aux.Id
	o.IsArchived = aux.IsArchived

	if aux.DateCreated != nil {
		t := time.Unix(*aux.DateCreated, 0).UTC()
		o.DateCreated = &t
	} else {
		o.DateCreated = nil
	}

	return nil
}
