package testutils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/messages"
	"github.com/tum-zulip/go-zulip/zulip/api/users"
	"github.com/tum-zulip/go-zulip/zulip/client"
	"github.com/tum-zulip/go-zulip/zulip/client/statistics"
	"github.com/tum-zulip/go-zulip/zulip/zuliprc"
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

var (
	OwnerClient     = namedClient{name: "owner", factory: GetOwnerClient}
	AdminClient     = namedClient{name: "admin", factory: GetAdminClient}
	ModeratorClient = namedClient{name: "moderator", factory: GetModeratorClient}
	NormalClient    = namedClient{name: "normal", factory: GetNormalClient}
	BotClient       = namedClient{name: "bot", factory: GetBotClient}
	AllClients      = []namedClient{OwnerClient, AdminClient, ModeratorClient, NormalClient}
)

// TODO(janez): Add guest client to tests

var (
	cache       sync.Map // map[string]client.Client - keyed by username
	cleanupOnce sync.Map // map[string]*sync.Once - keyed by username
)

type namedClient struct {
	name    string
	factory func(*testing.T) client.Client
}

func GetOwnerClient(t *testing.T) client.Client {
	t.Helper()
	return GetTestClient(t, TestOwnerUsername)
}

func GetAdminClient(t *testing.T) client.Client {
	t.Helper()
	return GetTestClient(t, TestAdminUsername)
}

func GetModeratorClient(t *testing.T) client.Client {
	t.Helper()
	return GetTestClient(t, TestModeratorUsername)
}

func GetGuestClient(t *testing.T) client.Client {
	t.Helper()
	return GetTestClient(t, TestGuestUsername)
}

func GetBotClient(t *testing.T) client.Client {
	t.Helper()
	t.Skip("Bot user tests are not implemented yet")
	return nil
}

func GetNormalClient(t *testing.T) client.Client {
	t.Helper()
	return GetTestClient(t, TestNormalUsername)
}

func GetOtherNormalClient(t *testing.T) client.Client {
	t.Helper()
	return GetTestClient(t, OtherNormalUsername)
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

func GetTestClient(t *testing.T, username string) client.Client {
	// Check if client already exists in cache
	if cachedClient, ok := cache.Load(username); ok {
		return cachedClient.(client.Client)
	}

	// Create client if not in cache
	newClient := getTestClient(t, username)

	// Store in cache atomically
	actual, loaded := cache.LoadOrStore(username, newClient)

	// Register cleanup only once per username
	once, _ := cleanupOnce.LoadOrStore(username, &sync.Once{})
	once.(*sync.Once).Do(func() {
		t.Cleanup(func() {
			writeStatisticsToJSON(username, newClient.GetStatistics())
		})
	})

	if loaded {
		// Another goroutine created the client first, use that one
		return actual.(client.Client)
	}

	return newClient
}

func writeStatisticsToJSON(username string, stats map[string]statistics.Statistic) {
	if len(stats) == 0 {
		return
	}

	statsPath, err := statisticsFilePath()
	if err != nil {
		log.Printf("Failed to determine statistics file path: %v", err)
		return
	}

	if statsPath == "" {
		return
	}

	if err := os.MkdirAll(filepath.Dir(statsPath), 0o755); err != nil {
		log.Printf("Failed to create statistics directory: %v", err)
		return
	}

	lock, err := acquireFileLock(statsPath+".lock", 5*time.Second)
	if err != nil {
		log.Printf("Failed to acquire statistics lock: %v", err)
		return
	}
	defer func() {
		if err := lock.release(); err != nil {
			log.Printf("Failed to release statistics lock: %v", err)
		}
	}()

	aggregated, err := readAggregatedStatistics(statsPath)
	if err != nil {
		log.Printf("Failed to read existing statistics JSON: %v", err)
		aggregated = make(map[string]map[string]statistics.Statistic)
	}

	mergeStatistics(aggregated, username, stats)

	if err := writeAggregatedStatistics(statsPath, aggregated); err != nil {
		log.Printf("Failed to write statistics JSON: %v", err)
	}
}

func statisticsFilePath() (string, error) {
	if override := os.Getenv("GO_ZULIP_STATS_FILE"); override != "" {
		return filepath.Abs(override)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	root, err := findModuleRoot(cwd)
	if err != nil {
		// Fall back to current working directory when go.mod isn't found.
		return filepath.Join(cwd, "test_statistics.json"), nil
	}

	return filepath.Join(root, "test_statistics.json"), nil
}

func findModuleRoot(start string) (string, error) {
	dir := start
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found starting from %s", start)
		}
		dir = parent
	}
}

func readAggregatedStatistics(path string) (map[string]map[string]statistics.Statistic, error) {
	data := make(map[string]map[string]statistics.Statistic)

	contents, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return data, nil
		}
		return nil, err
	}

	if len(contents) == 0 {
		return data, nil
	}

	if err := json.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	if data == nil {
		data = make(map[string]map[string]statistics.Statistic)
	}

	return data, nil
}

func mergeStatistics(store map[string]map[string]statistics.Statistic, username string, stats map[string]statistics.Statistic) {
	if store == nil {
		return
	}

	if _, ok := store[username]; !ok {
		store[username] = make(map[string]statistics.Statistic)
	}

	for endpoint, metric := range stats {
		existing := store[username][endpoint]
		existing.Count += metric.Count
		existing.ErrCount += metric.ErrCount
		existing.RetryCound += metric.RetryCound
		existing.TotalDuration += metric.TotalDuration
		store[username][endpoint] = existing
	}
}

type fileLock struct {
	file *os.File
	path string
}

func acquireFileLock(lockPath string, timeout time.Duration) (*fileLock, error) {
	deadline := time.Now().Add(timeout)

	for {
		f, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0o600)
		if err == nil {
			return &fileLock{file: f, path: lockPath}, nil
		}

		if !errors.Is(err, os.ErrExist) {
			return nil, err
		}

		if time.Now().After(deadline) {
			return nil, fmt.Errorf("timeout acquiring lock %s", lockPath)
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (l *fileLock) release() error {
	if l == nil {
		return nil
	}

	if l.file != nil {
		if err := l.file.Close(); err != nil {
			return err
		}
	}

	if err := os.Remove(l.path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func writeAggregatedStatistics(path string, payload map[string]map[string]statistics.Statistic) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), "stats-*.json")
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(tmp)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(payload); err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return err
	}

	if err := tmp.Close(); err != nil {
		os.Remove(tmp.Name())
		return err
	}

	if err := os.Rename(tmp.Name(), path); err != nil {
		if removeErr := os.Remove(path); removeErr != nil && !errors.Is(removeErr, os.ErrNotExist) {
			os.Remove(tmp.Name())
			return fmt.Errorf("removing stale statistics file: %w", removeErr)
		}

		if err := os.Rename(tmp.Name(), path); err != nil {
			os.Remove(tmp.Name())
			return err
		}
	}

	return nil
}

func getTestClient(t *testing.T, username string) client.Client {
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

func buildClientFromResponse(t *testing.T, username string, body []byte) client.Client {
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
	rc := &zuliprc.ZulipRC{
		Site:     TestSite,
		Email:    username,
		APIKey:   result.APIKey,
		Insecure: &insecure,
	}

	handler := slog.NewTextHandler(log.Default().Writer(), &slog.HandlerOptions{Level: slog.LevelInfo})

	client, err := client.NewClient(rc,
		client.WithLogger(slog.New(handler)),
		client.SkipWarnOnInsecureTLS(),
		client.EnableStatistics())

	if err != nil {
		t.Fatalf("Failed to create z.client: %v", err)
	}

	return client
}

func CreateRandomChannel(t *testing.T, apiClient client.Client, subscribers ...int64) (string, int64) {
	t.Helper()

	subs := append([]int64(nil), subscribers...)
	if len(subs) == 0 {
		resp, httpResp, err := apiClient.GetOwnUser(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpResp)
		subs = []int64{resp.UserId}
	}

	name := UniqueName("test-channel")
	resp, httpResp, err := apiClient.CreateChannel(context.Background()).
		Name(name).
		Description("Created by channel API tests").
		Subscribers(subs...).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return name, resp.Id
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

func UniqueName(prefix string) string {
	return fmt.Sprintf("%s%d%d", prefix, rand.Intn(1000), time.Now().UnixNano())
}

func getOwnUser(t *testing.T, apiClient client.Client) *users.GetOwnUserResponse {
	resp, httpResp, err := apiClient.GetOwnUser(context.Background()).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpResp.StatusCode)
	return resp
}

var idCache sync.Map

func GetUserId(t *testing.T, apiClient client.Client) int64 {
	t.Helper()

	if id, ok := idCache.Load(apiClient); ok {
		return id.(int64)
	}

	resp := getOwnUser(t, apiClient)

	require.NotNil(t, resp)

	idCache.Store(apiClient, resp.UserId)

	return resp.UserId
}

func GetUserEmail(t *testing.T, apiClient client.Client) string {
	t.Helper()

	resp := getOwnUser(t, apiClient)

	return resp.Email
}

func CreateRandomUserGroup(t *testing.T, apiClient client.Client, members ...int64) int64 {
	t.Helper()

	groupId := rand.Intn(1000000)

	resp, httpResp, err := apiClient.CreateUserGroup(context.Background()).Name(fmt.Sprintf("test-group-%d", groupId)).Description("Test Group").Members(members).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpResp.StatusCode)

	return resp.GroupId
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

	allClientIds := make([]int64, 0, len(clientFactories))
	for _, factory := range clientFactories {
		apiClient := factory(t)
		userId := GetUserId(t, apiClient)
		allClientIds = append(allClientIds, userId)
	}

	name, id := CreateRandomChannel(t, GetAdminClient(t), allClientIds...)

	if channelCahe.CompareAndSwap(nil, channel{id, name}) {
		return name, id
	}

	v := channelCahe.Load()
	c := v.(channel)
	return c.name, c.id

}

func SendChannelMessage(t *testing.T, apiClient client.Client, channelId int64, topic, content string) int64 {
	t.Helper()

	resp, httpResp, err := apiClient.SendMessage(context.Background()).
		To(z.ChannelAsRecipient(channelId)).
		Topic(topic).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	RequireStatusOK(t, httpResp)

	return resp.Id
}

type ChannelMessage struct {
	ChannelId int64
	Topic     string
	MessageId int64
}

func CreateChannelMessage(t *testing.T, apiClient client.Client, channelId int64) ChannelMessage {
	t.Helper()

	topic := UniqueName("topic")
	content := fmt.Sprintf("automated test message %s", UniqueName("content"))
	messageId := SendChannelMessage(t, apiClient, channelId, topic, content)
	assert.Greater(t, messageId, int64(0))

	return ChannelMessage{
		ChannelId: channelId,
		Topic:     topic,
		MessageId: messageId,
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
	require.Greater(t, resp.Id, int64(0))

	return resp.Id
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
