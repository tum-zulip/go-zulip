package zulip

import (
	"encoding/json"
	"time"
)

// SavedSnippet Object containing the details of the saved snippet.
type SavedSnippet struct {
	// The unique Id of the saved snippet.
	Id int64 `json:"id,omitempty"`
	// The title of the saved snippet.
	Title string `json:"title,omitempty"`
	// The content of the saved snippet in [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown) format.  Clients should insert this content into a message when using a saved snippet.
	Content string `json:"content,omitempty"`
	// The UNIX timestamp for when the saved snippet was created, in UTC seconds.
	DateCreated time.Time `json:"date_created,omitempty"`
}

type savedSnippetJSON struct {
	Id          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	DateCreated int64  `json:"date_created,omitempty"`
}

func (o SavedSnippet) MarshalJSON() ([]byte, error) {
	aux := savedSnippetJSON{
		Id:          o.Id,
		Title:       o.Title,
		Content:     o.Content,
		DateCreated: o.DateCreated.Unix(),
	}

	return json.Marshal(aux)
}

func (o *SavedSnippet) UnmarshalJSON(data []byte) error {
	var aux savedSnippetJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.Id = aux.Id
	o.Title = aux.Title
	o.Content = aux.Content
	o.DateCreated = time.Unix(aux.DateCreated, 0).UTC()

	return nil
}
