package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/tum-zulip/go-zulip/pkg/api"
)

const (
	// usernames
	testSite              = "http://localhost:9991"
	testOwnerUsername     = "desdemona@zulip.com"
	testAdminUsername     = "iago@zulip.com"
	testModeratorUsername = "shiva@zulip.com"
	testGuestUsername     = "polonius@zulip.com"
	testNormalUsername    = "AARON@zulip.com"
	otherNormalUsername   = "hamlet@zulip.com"
)

var (
	ownerClient     = namedClient{name: "owner", factory: GetOwnerClient}
	adminClient     = namedClient{name: "admin", factory: GetAdminClient}
	moderatorClient = namedClient{name: "moderator", factory: GetModeratorClient}
	normalClient    = namedClient{name: "normal", factory: GetNormalClient}
	otherClient     = namedClient{name: "other", factory: GetOtherNormalClient}
	guestClient     = namedClient{name: "guest", factory: GetGuestClient}
	botClient       = namedClient{name: "bot", factory: GetBotClient}
	allClients      = []namedClient{ownerClient, adminClient, moderatorClient, normalClient, botClient}
)

type namedClient struct {
	name    string
	factory func(*testing.T) *api.ZulipClient
}

func Test_getTestClient(t *testing.T) {
	runForClients(t, allClients, func(t *testing.T, client *api.ZulipClient) {
		if client == nil {
			t.Fatal("getTestClient returned nil client")
		}
	})
}

func GetOwnerClient(t *testing.T) *api.ZulipClient {
	t.Helper()
	return getTestClient(t, testOwnerUsername)
}

func GetAdminClient(t *testing.T) *api.ZulipClient {
	t.Helper()
	return getTestClient(t, testAdminUsername)
}

func GetModeratorClient(t *testing.T) *api.ZulipClient {
	t.Helper()
	return getTestClient(t, testModeratorUsername)
}

func GetGuestClient(t *testing.T) *api.ZulipClient {
	t.Helper()
	return getTestClient(t, testGuestUsername)
}

func GetBotClient(t *testing.T) *api.ZulipClient {
	t.Helper()
	t.Skip("Bot user tests are not implemented yet")
	return nil
}

func GetNormalClient(t *testing.T) *api.ZulipClient {
	t.Helper()
	return getTestClient(t, testNormalUsername)
}

func GetOtherNormalClient(t *testing.T) *api.ZulipClient {
	t.Helper()
	return getTestClient(t, otherNormalUsername)
}

func runForClients(t *testing.T, clients []namedClient, fn func(*testing.T, *api.ZulipClient)) {
	for _, client := range clients {
		t.Run(client.name, func(t *testing.T) {
			cl := client.factory(t)
			fn(t, cl)
		})
	}
}

func getTestClient(t *testing.T, username string) *api.ZulipClient {
	t.Helper()

	for attempt := 0; attempt < 2; attempt++ {
		url := fmt.Sprintf("%s/api/v1/dev_fetch_api_key?username=%s", testSite, username)
		resp, err := http.DefaultClient.Post(url, "application/json", nil)
		if err != nil {
			t.Fatalf("Failed to fetch API key for user %s: %v", username, err)
		}

		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			t.Fatalf("Failed to read API key response for user %s: %v", username, readErr)
		}

		if resp.StatusCode == http.StatusOK {
			return buildClientFromResponse(t, username, body)
		}

		if resp.StatusCode == http.StatusUnauthorized && attempt == 0 && tryReactivateUser(t, username, body) {
			continue
		}

		t.Fatalf("Failed to fetch API key for user %s: status code %d, response: %s", username, resp.StatusCode, string(body))
	}

	t.Fatalf("Failed to fetch API key for user %s after reactivation attempt", username)
	return nil
}

func buildClientFromResponse(t *testing.T, username string, body []byte) *api.ZulipClient {
	t.Helper()

	var result struct {
		APIKey string `json:"api_key"`
		EMail  string `json:"email"`
		UserId int    `json:"user_id"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("Failed to decode API key response for user %s: %v", username, err)
	}
	if result.APIKey == "" {
		t.Fatalf("Empty API key received for user %s", username)
	}
	if result.EMail != username {
		t.Fatalf("Unexpected email in API key response: got %s, want %s", result.EMail, username)
	}

	insecure := true
	rc := &api.ZulipRC{
		Site:     testSite,
		Email:    username,
		APIKey:   result.APIKey,
		Insecure: &insecure,
	}

	handler := slog.NewTextHandler(log.Default().Writer(), &slog.HandlerOptions{Level: slog.LevelDebug})

	client, err := api.NewZulipClient(rc, api.WithLogger(slog.New(handler)))
	if err != nil {
		t.Fatalf("Failed to create Zulip client: %v", err)
	}

	return client
}

func tryReactivateUser(t *testing.T, username string, body []byte) bool {
	t.Helper()

	var apiErr struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	}

	if err := json.Unmarshal(body, &apiErr); err != nil {
		return false
	}

	if apiErr.Code != "USER_DEACTIVATED" {
		return false
	}

	t.Logf("User %s is deactivated; attempting to reactivate via owner account", username)

	if err := ensureUserActive(t, username); err != nil {
		t.Fatalf("Failed to reactivate user %s: %v", username, err)
	}

	return true
}

func ensureUserActive(t *testing.T, username string) error {
	t.Helper()

	owner := GetOwnerClient(t)

	resp, _, err := owner.GetUserByEmail(context.Background(), username).Execute()
	if err != nil {
		return fmt.Errorf("fetching user metadata: %w", err)
	}

	if resp.User.IsActive {
		return nil
	}

	_, _, err = owner.ReactivateUser(context.Background(), resp.User.UserId).Execute()
	if err != nil {
		return fmt.Errorf("reactivating user: %w", err)
	}

	return nil
}

func uniqueName(prefix string) string {
	return fmt.Sprintf("%s-%d-%d", prefix, time.Now().UnixNano(), rand.Intn(100000))
}
