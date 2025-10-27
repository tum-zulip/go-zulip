package testutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/channels"
	"github.com/tum-zulip/go-zulip/zulip/api/messages"
	"github.com/tum-zulip/go-zulip/zulip/api/users"
	"github.com/tum-zulip/go-zulip/zulip/client"
	"github.com/tum-zulip/go-zulip/zulip/client/statistics"
)

const (
	// usernames
	TestSite              = "http://localhost:9991"
	TestOwnerUsername     = "desdemona@zulip.com"
	TestAdminUsername     = "iago@zulip.com"
	TestModeratorUsername = "shiva@zulip.com"
	TestGuestUsername     = "polonius@zulip.com"
	TestNormalUsername    = "AARON@zulip.com"
	OtherNormalUsername   = "hamlet@zulip.com"
	DeactivateTestUser    = "cordelia@zulip.com"
)

// TODO: Add guest client to tests.

var (
	OwnerClient     = namedClient{name: "owner", factory: GetOwnerClient}
	AdminClient     = namedClient{name: "admin", factory: GetAdminClient}
	ModeratorClient = namedClient{name: "moderator", factory: GetModeratorClient}
	NormalClient    = namedClient{name: "normal", factory: GetNormalClient}
	BotClient       = namedClient{name: "bot", factory: GetBotClient}
	AllClients      = []namedClient{OwnerClient, AdminClient, ModeratorClient, NormalClient}
)

type namedClient struct {
	name    string
	factory func(*testing.T) client.Client
}

func GetOwnerClient(t *testing.T) client.Client {
	t.Helper()
	_, client := GetTestClient(t, TestOwnerUsername)
	return client
}

func GetAdminClient(t *testing.T) client.Client {
	t.Helper()
	_, client := GetTestClient(t, TestAdminUsername)
	return client
}

func GetModeratorClient(t *testing.T) client.Client {
	t.Helper()
	_, client := GetTestClient(t, TestModeratorUsername)
	return client
}

func GetGuestClient(t *testing.T) client.Client {
	t.Helper()
	_, client := GetTestClient(t, TestGuestUsername)
	return client
}

func GetBotClient(t *testing.T) client.Client {
	t.Helper()
	rc := createBot(t, "test-bot")
	client, err := client.NewClient(rc,
		client.SkipWarnOnInsecureTLS(),
		client.EnableStatistics())
	require.NoError(t, err)
	return client
}

func GetNormalClient(t *testing.T) client.Client {
	t.Helper()
	_, client := GetTestClient(t, TestNormalUsername)
	return client
}

func GetOtherNormalClient(t *testing.T) client.Client {
	t.Helper()
	_, client := GetTestClient(t, OtherNormalUsername)
	return client
}

func RunForClients(t *testing.T, clients []namedClient, fn func(*testing.T, client.Client)) {
	for _, client := range clients {
		t.Run(client.name, func(t *testing.T) {
			cl := client.factory(t)
			fn(t, cl)
		})
	}
}

func RunForAllClients(t *testing.T, fn func(*testing.T, client.Client)) {
	RunForClients(t, AllClients, fn)
}

func RunForAdminAndOwnerClients(t *testing.T, fn func(*testing.T, client.Client)) {
	RunForClients(t, []namedClient{OwnerClient, AdminClient}, fn)
}

var valueCache sync.Map

func GetTestClient(t *testing.T, username string) (*z.ZulipRC, client.Client) {
	type cachevalue struct {
		rc     *z.ZulipRC
		client client.Client
	}

	// Check if client already exists in cache
	if cachedClient, ok := valueCache.Load(username); ok {
		cv, cacheOk := cachedClient.(cachevalue)
		if !cacheOk {
			t.Fatalf("Failed to cast cached client for user %s", username)
		}
		return cv.rc, cv.client
	}

	// Create client if not in cache
	newRC, newClient := getTestClient(t, username)

	// Store in cache atomically
	actual, loaded := valueCache.LoadOrStore(username, cachevalue{rc: newRC, client: newClient})

	// Register cleanup for EVERY client instance, not just once per username
	// This ensures statistics from all parallel tests are saved
	t.Cleanup(func() {
		require.NoError(t, os.MkdirAll("/tmp/go-zulip-stats", 0o750))

		data, err := json.MarshalIndent(map[string]map[string]statistics.Statistic{
			username: newClient.GetStatistics(),
		}, "", "  ")
		if err != nil {
			t.Logf("Failed to marshal statistics for user %s: %v", username, err)
			return
		}
		f, err := os.CreateTemp("/tmp/go-zulip-stats", "fmt-go-zulip-stats-*.json")
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write(data)
		if err != nil {
			t.Logf("Failed to write statistics for user %s: %v", username, err)
			return
		}
	})

	if loaded {
		// Another goroutine created the client first, use that one
		cv, ok := actual.(cachevalue)
		if !ok {
			t.Fatalf("Failed to cast cached client for user %s", username)
		}
		return cv.rc, cv.client
	}

	return newRC, newClient
}

func getTestClient(t *testing.T, username string) (*z.ZulipRC, client.Client) {
	t.Helper()

	info := fetchUserInfo(t, username)

	insecure := true
	rc := &z.ZulipRC{
		Site:     TestSite,
		Email:    info.EMail,
		APIKey:   info.APIKey,
		Insecure: &insecure,
	}

	client, err := client.NewClient(rc,
		client.SkipWarnOnInsecureTLS(),
		client.EnableStatistics())
	if err != nil {
		t.Fatalf("Failed to create z.client: %v", err)
	}

	return rc, client
}

type UserInfo struct {
	APIKey string `json:"api_key"`
	EMail  string `json:"email"`
	UserID int    `json:"user_id"`
}

func fetchUserInfo(t *testing.T, username string) UserInfo {
	t.Helper()

	for attempt := 0; attempt < 2; attempt++ {
		url := fmt.Sprintf("%s/api/v1/dev_fetch_api_key?username=%s", TestSite, username)
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
			var result UserInfo

			if err := json.Unmarshal(body, &result); err != nil {
				t.Fatalf("Failed to decode API key response for user %s: %v", username, err)
			}
			if result.APIKey == "" {
				t.Fatalf("Empty API key received for user %s", username)
			}
			if result.EMail != username {
				t.Fatalf("Unexpected email in API key response: got %s, want %s", result.EMail, username)
			}
			return result

		}

		if resp.StatusCode == http.StatusUnauthorized && attempt == 0 && tryReactivateUser(t, username, body) {
			continue
		}

		t.Fatalf(
			"Failed to fetch API key for user %s: status code %d, response: %s",
			username,
			resp.StatusCode,
			string(body),
		)
	}
	t.Fatalf("Failed to fetch API key for user %s after reactivation attempt", username)
	return UserInfo{}
}

func CreateRandomChannel(t *testing.T, apiClient client.Client, subscribers ...int64) (string, int64) {
	t.Helper()

	name := UniqueName("test-channel")
	{

		subs := append([]int64(nil), subscribers...)
		if len(subs) == 0 {
			subs = []int64{GetUserID(t, apiClient)}
		}

		resp, httpResp, err := apiClient.Subscribe(context.Background()).
			Subscriptions([]channels.SubscriptionRequest{{Name: name}}).
			Principals(z.UserIdsAsPrincipals(subs...)).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
	}

	resp, httpResp, err := apiClient.GetChannelID(context.Background()).Channel(name).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return name, resp.ChannelID
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

	_, _, err = owner.ReactivateUser(context.Background(), resp.User.UserID).Execute()
	if err != nil {
		return fmt.Errorf("reactivating user: %w", err)
	}

	return nil
}

func UniqueName(prefix string) string {
	return fmt.Sprintf("%s%d%d", prefix, rand.Int(), time.Now().UnixNano())
}

func getOwnUser(t *testing.T, apiClient client.Client) *users.GetOwnUserResponse {
	resp, httpResp, err := apiClient.GetOwnUser(context.Background()).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpResp.StatusCode)
	return resp
}

func createBot(t *testing.T, botName string) *z.ZulipRC {
	t.Helper()

	rc, _ := GetTestClient(t, TestOwnerUsername)

	type botResponse struct {
		z.Response

		UserID int64  `json:"user_id"`
		APIKey string `json:"api_key"`
	}

	data := url.Values{}
	data.Set("full_name", botName)
	data.Set("bot_type", strconv.Itoa(int(z.BotTypeGeneric)))
	data.Set("short_name", botName)

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	for k, v := range data {
		err := w.WriteField(k, v[0])
		require.NoError(t, err)
	}
	require.NoError(t, w.Close())

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/bots", TestSite),
		body,
	)
	require.NoError(t, err)
	req.SetBasicAuth(rc.Email, rc.APIKey)
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode, "bot creation failed: %s", string(respBody))

	var botResp botResponse
	err = json.Unmarshal(respBody, &botResp)
	require.NoError(t, err)
	require.NotZero(t, botResp.UserID)
	require.NotEmpty(t, botResp.APIKey)

	insecure := true
	return &z.ZulipRC{
		Site:     TestSite,
		Email:    fmt.Sprintf("%s-bot@zulip.com", botName),
		APIKey:   botResp.APIKey,
		Insecure: &insecure,
	}
}

var idCache sync.Map

func GetUserID(t *testing.T, apiClient client.Client) int64 {
	t.Helper()

	if id, ok := idCache.Load(apiClient); ok {
		return id.(int64)
	}

	resp := getOwnUser(t, apiClient)

	require.NotNil(t, resp)

	idCache.Store(apiClient, resp.UserID)

	return resp.UserID
}

func GetUserEmail(t *testing.T, apiClient client.Client) string {
	t.Helper()

	resp := getOwnUser(t, apiClient)

	return resp.Email
}

func CreateRandomUserGroup(t *testing.T, apiClient client.Client, members ...int64) int64 {
	t.Helper()

	name := UniqueName("test-group-")

	resp, _, err := apiClient.CreateUserGroup(context.Background()).
		Name(name).
		Description("Test Group").
		Members(members).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)

	if resp.GroupID == 0 {
		// older Zulip versions did not return GroupID in response
		resp, httpResp, err := apiClient.GetUserGroups(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		for _, g := range resp.UserGroups {
			if g.Name == name {
				return g.ID
			}
		}
		t.Fatalf("Created user group not found in list")
	}

	return resp.GroupID
}

func RequireStatusOK(t *testing.T, httpResp *http.Response) {
	t.Helper()
	require.NotNil(t, httpResp)
	assert.Equal(t, 200, httpResp.StatusCode)
}

var channelCahe atomic.Value

func GetChannelWithAllClients(t *testing.T) (string, int64) {
	t.Helper()

	type channel struct {
		id   int64
		name string
	}

	if v := channelCahe.Load(); v != nil {
		c := v.(channel)
		return c.name, c.id
	}

	clientFactories := []func(*testing.T) client.Client{
		GetOwnerClient,
		GetAdminClient,
		GetModeratorClient,
		GetNormalClient,
		GetOtherNormalClient,
	}

	allClientIDs := make([]int64, 0, len(clientFactories))
	for _, factory := range clientFactories {
		apiClient := factory(t)
		userID := GetUserID(t, apiClient)
		allClientIDs = append(allClientIDs, userID)
	}

	name, id := CreateRandomChannel(t, GetAdminClient(t), allClientIDs...)

	if channelCahe.CompareAndSwap(nil, channel{id, name}) {
		return name, id
	}

	v := channelCahe.Load()
	c := v.(channel)
	return c.name, c.id
}

func SendChannelMessage(t *testing.T, apiClient client.Client, channelID int64, topic, content string) int64 {
	t.Helper()

	resp, httpResp, err := apiClient.SendMessage(context.Background()).
		To(z.ChannelAsRecipient(channelID)).
		Topic(topic).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp.ID
}

var serverFeatureLevel atomic.Int64

func GetFeatureLevel(t *testing.T) int {
	if level := serverFeatureLevel.Load(); level != 0 {
		return int(level)
	}

	client := GetOwnerClient(t)

	featureLevel, _, err := client.GetServerSettings(context.Background()).Execute()
	require.NoError(t, err)
	require.NotNil(t, featureLevel)
	serverFeatureLevel.Store(int64(featureLevel.ZulipFeatureLevel))

	return int(featureLevel.ZulipFeatureLevel)
}

func RequireFeatureLevel(t *testing.T, minLevel int) {
	t.Helper()

	currentLevel := GetFeatureLevel(t)
	if currentLevel < minLevel {
		t.Skipf("Skipping test: requires feature level %d, server has %d", minLevel, currentLevel)
	}
}

type ChannelMessage struct {
	ChannelID int64
	Topic     string
	MessageID int64
}

func CreateChannelMessage(t *testing.T, apiClient client.Client, channelID int64) ChannelMessage {
	t.Helper()

	topic := UniqueName("topic")
	content := fmt.Sprintf("automated test message %s", UniqueName("content"))
	messageID := SendChannelMessage(t, apiClient, channelID, topic, content)
	assert.Positive(t, messageID)

	return ChannelMessage{
		ChannelID: channelID,
		Topic:     topic,
		MessageID: messageID,
	}
}

func CreateDirectMessage(t *testing.T, apiClient client.Client, to int64) int64 {
	t.Helper()

	content := UniqueName("content")
	resp, httpResp, err := apiClient.SendMessage(context.Background()).
		To(z.UserAsRecipient(to)).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)
	require.Positive(t, resp.ID)

	return resp.ID
}

func UploadFileForTest(t *testing.T, ctx context.Context, apiClient client.Client) *messages.UploadFileResponse {
	t.Helper()

	tmp, err := os.CreateTemp("", "z.upload-*.txt")
	require.NoError(t, err)
	defer func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}()

	_, err = tmp.WriteString("uploaded from automated test")
	require.NoError(t, err)
	_, err = tmp.Seek(0, 0)
	require.NoError(t, err)

	resp, httpResp, err := apiClient.UploadFile(ctx).
		Filename(tmp).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp
}
