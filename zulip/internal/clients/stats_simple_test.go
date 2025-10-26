package clients

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/zuliprc"
)

// TestRateLimitRetrySimple tests a simple rate limit retry scenario
func TestRateLimitRetrySimple(t *testing.T) {
	var requestCount atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := requestCount.Add(1)
		t.Logf("Request #%d received", count)

		// First request: return rate limit
		// Second request: return success
		if count == 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			response := zulip.RateLimitedError{
				CodedError: zulip.CodedError{
					Response: zulip.Response{
						Result: "error",
						Msg:    "rate limited",
					},
					Code: "RATE_LIMIT_HIT",
				},
				RetryAfter: 100 * time.Millisecond,
			}
			jsonBytes, _ := json.Marshal(response)
			t.Logf("JSON being sent: %s", string(jsonBytes))
			w.Write(jsonBytes)
			t.Logf("Returned rate limit error with 100ms retry-after")
			return
		}

		// Success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := zulip.Response{
			Result: "success",
		}
		json.NewEncoder(w).Encode(response)
		t.Logf("Returned success")
	}))
	defer server.Close()

	rc := &zuliprc.ZulipRC{
		Site:   server.URL,
		Email:  "test@example.com",
		APIKey: "test-key",
	}

	cfg, err := NewConfig(rc,
		WithGatherStats(true),
		WithMaxRetries(retryIndefinitely),
	)
	assert.NoError(t, err)

	client := NewRetryClient(cfg)

	req, err := http.NewRequest("GET", server.URL+"/api/v1/test", nil)
	assert.NoError(t, err)
	req.SetBasicAuth(rc.Email, rc.APIKey)

	var model zulip.Response
	_, err = client.CallAPI(context.Background(), "/api/v1/test", req, &model)
	assert.NoError(t, err)

	// Check stats
	stats := client.Stats.GetStatistics()
	t.Logf("Stats: %+v", stats)

	// Should have made 2 requests total (1 rate limited, 1 successful)
	assert.Equal(t, int32(2), requestCount.Load())

	// Check rate-limit-timeout stat
	if rateLimitStats, ok := stats["rate-limit-timeout"]; ok {
		t.Logf("Rate limit timeout duration: %v", rateLimitStats.TotalDuration)
		assert.Equal(t, 100*time.Millisecond, rateLimitStats.TotalDuration,
			"Rate limit timeout should be exactly 100ms")
	} else {
		t.Error("Expected rate-limit-timeout stats to be present")
	}
}
