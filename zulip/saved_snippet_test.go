package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	z "github.com/tum-zulip/go-zulip/zulip"
)

func TestSavedSnippetMarshalJSON_EncodesUnixSeconds(t *testing.T) {
	created := time.Unix(1700000000, 500*int64(time.Millisecond)).UTC()
	snippet := z.SavedSnippet{
		ID:          11,
		Title:       "Snippet",
		Content:     "**bold**",
		DateCreated: created,
	}

	data, err := json.Marshal(snippet)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	value, ok := payload["date_created"]
	require.True(t, ok)
	require.IsType(t, float64(0), value)
	assert.InEpsilon(t, float64(created.Unix()), value, 0.001)
}

func TestSavedSnippetUnmarshalJSON_DecodesUnixSeconds(t *testing.T) {
	raw := []byte(`{"id":11,"title":"Snippet","content":"**bold**","date_created":1700000000}`)

	var snippet z.SavedSnippet
	require.NoError(t, json.Unmarshal(raw, &snippet))

	assert.Equal(t, int64(11), snippet.ID)
	assert.Equal(t, "Snippet", snippet.Title)
	assert.Equal(t, "**bold**", snippet.Content)
	assert.Equal(t, int64(1700000000), snippet.DateCreated.Unix())
	assert.Equal(t, time.UTC, snippet.DateCreated.Location())
}
