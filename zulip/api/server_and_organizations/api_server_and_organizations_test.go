package server_and_organizations_test

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_GetServerSettings(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetServerSettings(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.NotEmpty(t, resp.RealmName)
		assert.NotEmpty(t, resp.RealmUrl)
	})
}

func Test_CodePlaygrounds(t *testing.T) {
	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		name := fmt.Sprintf("Playground %s", UniqueName("code"))
		resp, httpResp, err := apiClient.AddCodePlayground(ctx).
			Name(name).
			PygmentsLanguage("python").
			UrlTemplate("https://play.z.example/run?code={code}").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)

		playgroundID := resp.ID
		removed := false
		defer func() {
			if removed {
				return
			}
			_, _, cleanupErr := apiClient.RemoveCodePlayground(context.Background(), playgroundID).Execute()
			if cleanupErr != nil {
				t.Logf("cleanup remove code playground %d: %v", playgroundID, cleanupErr)
			}
		}()

		removeResp, removeHTTP, err := apiClient.RemoveCodePlayground(ctx, playgroundID).Execute()
		require.NoError(t, err)
		require.NotNil(t, removeResp)
		RequireStatusOK(t, removeHTTP)
		assert.Equal(t, "success", removeResp.Result)
		removed = true
	})
}

func Test_Linkifiers(t *testing.T) {
	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		pattern := fmt.Sprintf("test-%s-(?P<id>[0-9]+)", UniqueName("linkifier"))
		urlTemplate := "https://z.example/issues/{id}"

		addResp, httpResp, err := apiClient.AddLinkifier(ctx).
			Pattern(pattern).
			UrlTemplate(urlTemplate).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, addResp)
		RequireStatusOK(t, httpResp)

		linkifierID := addResp.ID
		removed := false
		defer func() {
			if removed {
				return
			}
			_, _, cleanupErr := apiClient.RemoveLinkifier(context.Background(), linkifierID).Execute()
			if cleanupErr != nil {
				t.Logf("cleanup remove linkifier %d: %v", linkifierID, cleanupErr)
			}
		}()

		listResp, httpResp, err := apiClient.GetLinkifiers(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, listResp)
		RequireStatusOK(t, httpResp)

		linkifiers := listResp.Linkifiers
		require.NotEmpty(t, linkifiers)
		require.True(
			t,
			linkifierExists(linkifiers, linkifierID, pattern, urlTemplate),
			"created linkifier missing from listing",
		)

		updatedTemplate := urlTemplate + "?source=api-test"
		updateResp, httpResp, err := apiClient.UpdateLinkifier(ctx, linkifierID).
			Pattern(pattern).
			UrlTemplate(updatedTemplate).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, updateResp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, "success", updateResp.Result)

		updatedListResp, httpResp, err := apiClient.GetLinkifiers(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, updatedListResp)
		RequireStatusOK(t, httpResp)
		assert.True(
			t,
			linkifierExists(updatedListResp.Linkifiers, linkifierID, pattern, updatedTemplate),
			"updated linkifier missing or incorrect",
		)

		originalOrder := extractlinkifierIDs(updatedListResp.Linkifiers)
		movedOrder := moveIdToFront(originalOrder, linkifierID)
		if !int64SlicesEqual(originalOrder, movedOrder) {
			reorderResp, httpResp, err := apiClient.ReorderLinkifiers(ctx).
				OrderedLinkifierIDs(movedOrder).
				Execute()
			require.NoError(t, err)
			require.NotNil(t, reorderResp)
			RequireStatusOK(t, httpResp)
			assert.Equal(t, "success", reorderResp.Result)

			restoreResp, httpResp, err := apiClient.ReorderLinkifiers(ctx).
				OrderedLinkifierIDs(originalOrder).
				Execute()
			require.NoError(t, err)
			require.NotNil(t, restoreResp)
			RequireStatusOK(t, httpResp)
			assert.Equal(t, "success", restoreResp.Result)
		}

		removeResp, removeHTTP, err := apiClient.RemoveLinkifier(ctx, linkifierID).Execute()
		require.NoError(t, err)
		require.NotNil(t, removeResp)
		RequireStatusOK(t, removeHTTP)
		assert.Equal(t, "success", removeResp.Result)
		removed = true
	})
}

func Test_CustomProfileFields(t *testing.T) {
	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		const fieldName = "API Test Profile Field"

		listResp, httpResp, err := apiClient.GetCustomProfileFields(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, listResp)
		RequireStatusOK(t, httpResp)

		fields := listResp.CustomFields
		var fieldID int64
		for _, field := range fields {
			if field.Name == fieldName {
				fieldID = field.ID
				break
			}
		}

		if fieldID == 0 {
			createResp, createHTTP, err := apiClient.CreateCustomProfileField(ctx).
				FieldType(1).
				Name(fieldName).
				Hint("Created by go-z.automated tests").
				EditableByUser(true).
				Required(false).
				Execute()
			require.NoError(t, err)
			require.NotNil(t, createResp)
			RequireStatusOK(t, createHTTP)
			fieldID = createResp.ID

			listResp, httpResp, err = apiClient.GetCustomProfileFields(ctx).Execute()
			require.NoError(t, err)
			require.NotNil(t, listResp)
			RequireStatusOK(t, httpResp)
			fields = listResp.CustomFields
		}

		require.NotZero(t, fieldID, "profile field ID must be set")
		require.True(t, profileFieldExists(fields, fieldID), "profile field missing from listing")

		originalOrder := extractProfileFieldIDs(fields)
		movedOrder := moveIdToFront(originalOrder, fieldID)
		if !int64SlicesEqual(originalOrder, movedOrder) {
			reorderResp, reorderHTTP, err := apiClient.ReorderCustomProfileFields(ctx).
				Order(movedOrder).
				Execute()
			require.NoError(t, err)
			require.NotNil(t, reorderResp)
			RequireStatusOK(t, reorderHTTP)
			assert.Equal(t, "success", reorderResp.Result)

			restoreResp, restoreHTTP, err := apiClient.ReorderCustomProfileFields(ctx).
				Order(originalOrder).
				Execute()
			require.NoError(t, err)
			require.NotNil(t, restoreResp)
			RequireStatusOK(t, restoreHTTP)
			assert.Equal(t, "success", restoreResp.Result)
		}
	})
}

func Test_Presence(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.GetPresence(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)

		presences := resp.Presences
		assert.NotEmpty(t, presences)
		for email, details := range presences {
			assert.NotEmpty(t, details, "presence data missing for %s", email)
		}
	})
}

func Test_RealmExports(t *testing.T) {
	RequireFeatureLevel(t, 295)

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		consentsResp, httpResp, err := apiClient.GetRealmExportConsents(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, consentsResp)
		RequireStatusOK(t, httpResp)

		exportsResp, httpResp, err := apiClient.GetRealmExports(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, exportsResp)
		RequireStatusOK(t, httpResp)
		assert.NotNil(t, exportsResp.Exports)

		exportResp, httpResp, err := apiClient.ExportRealm(ctx).Execute()
		if err != nil {
			if handledRateLimit(t, err) {
				return
			}
		}
		require.NoError(t, err)
		require.NotNil(t, exportResp)
		RequireStatusOK(t, httpResp)
	})
}

func Test_WelcomeBotPreview(t *testing.T) {
	RequireFeatureLevel(t, 416)

	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.TestWelcomeBotCustomMessage(ctx).
			WelcomeMessageCustomText(fmt.Sprintf("Welcome preview from automated test %s", UniqueName("welcome"))).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Positive(t, resp.MessageID)
	})
}

func Test_RealmUserSettingsDefaults(t *testing.T) {
	RunForAdminAndOwnerClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.UpdateRealmUserSettingsDefaults(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		assert.Equal(t, "success", resp.Result)
	})
}

func Test_CustomEmojiLifecycle(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		emojiName := strings.ToLower(UniqueName("emoji"))
		emojiFile := newEmojiPNG(t)

		uploadResp, httpResp, err := apiClient.UploadCustomEmoji(ctx, emojiName).
			Filename(emojiFile).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, uploadResp)
		RequireStatusOK(t, httpResp)

		listResp, httpResp, err := apiClient.GetCustomEmoji(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, listResp)
		RequireStatusOK(t, httpResp)

		currentEmoji := findEmojiByName(listResp.Emoji, emojiName)
		require.NotNil(t, currentEmoji, "uploaded emoji not returned by GET /realm/emoji")
		assert.False(t, currentEmoji.Deactivated)

		deactivateResp, deactivateHTTP, err := apiClient.DeactivateCustomEmoji(ctx, emojiName).Execute()
		require.NoError(t, err)
		require.NotNil(t, deactivateResp)
		RequireStatusOK(t, deactivateHTTP)
		assert.Equal(t, "success", deactivateResp.Result)

		afterResp, httpResp, err := apiClient.GetCustomEmoji(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, afterResp)
		RequireStatusOK(t, httpResp)

		deactivated := findEmojiByName(afterResp.Emoji, emojiName)
		require.NotNil(t, deactivated, "deactivated emoji missing from listing")
		assert.True(t, deactivated.Deactivated)
	})
}

func linkifierExists(linkifiers []z.RealmLinkifiers, id int64, pattern, urlTemplate string) bool {
	for _, lf := range linkifiers {
		if lf.ID == id {
			return lf.Pattern == pattern && lf.UrlTemplate == urlTemplate
		}
	}
	return false
}

func extractlinkifierIDs(linkifiers []z.RealmLinkifiers) []int64 {
	ids := make([]int64, 0, len(linkifiers))
	for _, lf := range linkifiers {
		ids = append(ids, lf.ID)
	}
	return ids
}

func extractProfileFieldIDs(fields []z.CustomProfileField) []int64 {
	ids := make([]int64, 0, len(fields))
	for _, field := range fields {
		ids = append(ids, field.ID)
	}
	return ids
}

func moveIdToFront(ids []int64, id int64) []int64 {
	result := make([]int64, 0, len(ids))
	found := false
	for _, value := range ids {
		if value == id {
			found = true
			continue
		}
		result = append(result, value)
	}
	if !found {
		return append([]int64(nil), ids...)
	}
	return append([]int64{id}, result...)
}

func int64SlicesEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func profileFieldExists(fields []z.CustomProfileField, id int64) bool {
	for _, field := range fields {
		if field.ID == id {
			return true
		}
	}
	return false
}

func handledRateLimit(t *testing.T, err error) bool {
	var apiErr *z.APIError
	if !errors.As(err, &apiErr) {
		return false
	}

	body := string(apiErr.Body())
	if strings.Contains(body, "Exceeded rate limit") {
		t.Logf("realm export skipped due to rate limit: %s", strings.TrimSpace(body))
		return true
	}

	return false
}

func newEmojiPNG(t *testing.T) *os.File {
	t.Helper()

	tmp, err := os.CreateTemp("", "z.emoji-*.png")
	require.NoError(t, err)

	img := image.NewRGBA(image.Rect(0, 0, 16, 16))
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			img.Set(x, y, color.RGBA{R: 0xff, G: 0xa5, B: 0, A: 0xff})
		}
	}

	err = png.Encode(tmp, img)
	require.NoError(t, err)
	_, err = tmp.Seek(0, 0)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, tmp.Close())
		require.NoError(t, os.Remove(tmp.Name()))
	})

	return tmp
}

func findEmojiByName(emojis map[string]z.RealmEmoji, name string) *z.RealmEmoji {
	if len(emojis) == 0 {
		return nil
	}

	for _, emoji := range emojis {
		if strings.EqualFold(emoji.Name, name) {
			e := emoji
			return &e
		}
	}

	return nil
}
