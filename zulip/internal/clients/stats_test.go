package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

// TestStatsConcurrent tests that statistics are correctly gathered
// when multiple concurrent requests are made with various error conditions
func TestStatsConcurrent(t *testing.T) {
	var requestCount atomic.Int32
	var successCount atomic.Int32
	var errorCount atomic.Int32
	var rateLimitCount atomic.Int32

	// Mock server that:
	// - Sleeps a bit to simulate real work
	// - Returns error every 5 requests
	// - Returns rate limit error every 13 requests
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := requestCount.Add(1)

		// Simulate some work
		time.Sleep(10 * time.Millisecond)

		// Rate limit every 13 requests (happens before error check)
		if count%13 == 0 {
			rateLimitCount.Add(1)
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
				RetryAfter: 50 * time.Millisecond,
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		// Error every 5 requests (but not when it's also a rate limit)
		if count%5 == 0 {
			errorCount.Add(1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			response := zulip.CodedError{
				Response: zulip.Response{
					Result: "error",
					Msg:    "simulated error",
				},
				Code: "BAD_REQUEST",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		// Success
		successCount.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := zulip.Response{
			Result: "success",
			Msg:    "",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client with retry support and stats gathering
	rc := &zulip.ZulipRC{
		Site:   server.URL,
		Email:  "test@example.com",
		APIKey: "test-key",
	}

	cfg, err := NewConfig(rc,
		WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
		WithGatherStats(true),
		WithMaxRetries(retryIndefinitely), // Retry rate limits forever
	)
	require.NoError(t, err)

	client := NewRetryClient(cfg)

	// Number of concurrent goroutines and requests per goroutine
	numWorkers := 10
	requestsPerWorker := 10
	totalRequests := numWorkers * requestsPerWorker

	// Track completion
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Launch concurrent workers
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			t.Logf("Worker %d starting", workerID)

			for j := 0; j < requestsPerWorker; j++ {
				endpoint := fmt.Sprintf("/api/v1/test/worker-%d", workerID)

				req, err := http.NewRequest("GET", server.URL+endpoint, nil)
				if err != nil {
					t.Errorf("Failed to create request: %v", err)
					continue
				}
				req.SetBasicAuth(rc.Email, rc.APIKey)

				var model zulip.Response
				_, err = client.CallAPI(context.Background(), endpoint, req, &model)
				// We expect some errors, so just continue
				if err != nil {
					// Verify it's one of our expected error types
					if _, ok := err.(*zulip.APIError); !ok {
						t.Errorf("Unexpected error type: %T", err)
					}
				}
				t.Logf("Worker %d completed request %d", workerID, j+1)
			}
			t.Logf("Worker %d done", workerID)
		}(i)
	}

	// Wait for all requests to complete
	wg.Wait()

	// Get statistics
	stats := client.Stats.GetStatistics()

	// Verify statistics
	t.Logf("Request count: %d", requestCount.Load())
	t.Logf("Success count: %d", successCount.Load())
	t.Logf("Error count: %d", errorCount.Load())
	t.Logf("Rate limit count: %d", rateLimitCount.Load())

	// Calculate expected values
	// Note: Rate limits cause retries, so actual request count to server will be higher
	actualServerRequests := int(requestCount.Load())
	actualErrorCount := int(errorCount.Load())
	actualRateLimitCount := int(rateLimitCount.Load())

	t.Logf("Actual server requests: %d", actualServerRequests)

	// The client should track statistics per endpoint
	totalCallCount := uint64(0)
	totalErrCount := uint64(0)
	totalRetryCount := uint64(0)
	var totalDuration time.Duration

	for endpoint, stat := range stats {
		t.Logf("Endpoint %s: Count=%d, ErrCount=%d, RetryCount=%d, TotalDuration=%v",
			endpoint, stat.Count, stat.ErrCount, stat.RetryCount, stat.TotalDuration)

		totalCallCount += stat.Count
		totalErrCount += stat.ErrCount
		totalRetryCount += stat.RetryCount
		totalDuration += stat.TotalDuration
	}

	// Check rate-limit-timeout stats
	if rateLimitStats, ok := stats["rate-limit-timeout"]; ok {
		t.Logf("Rate limit timeout: TotalDuration=%v", rateLimitStats.TotalDuration)
	}

	// Assertions
	// 1. Total calls (including retries) should equal server requests
	assert.Equal(t, uint64(actualServerRequests), totalCallCount,
		"Total call count should equal number of HTTP requests made (including retries)")

	// 2. Error count should match the errors returned (including rate limits on first attempt)
	expectedErrors := actualErrorCount + actualRateLimitCount
	assert.Equal(t, uint64(expectedErrors), totalErrCount,
		"Total error count should match server errors + rate limits")

	// 3. Retry count should match rate limit count (since we retry those)
	assert.Equal(t, uint64(actualRateLimitCount), totalRetryCount,
		"Total retry count should match rate limit count")

	// 4. Duration should be positive and reasonable
	assert.Greater(t, totalDuration, time.Duration(0),
		"Total duration should be positive")

	// Each request sleeps 10ms, so minimum expected duration
	minExpectedDuration := time.Duration(actualServerRequests) * 10 * time.Millisecond
	assert.GreaterOrEqual(t, totalDuration, minExpectedDuration,
		"Total duration should be at least the sum of all request sleep times")

	// 5. Verify we actually had some retries
	assert.Greater(t, actualRateLimitCount, 0,
		"Should have encountered at least one rate limit")

	// 6. Verify we had some errors
	assert.Greater(t, actualErrorCount, 0,
		"Should have encountered at least one error")

	t.Logf("\n=== Summary ===")
	t.Logf("Total client requests: %d", totalRequests)
	t.Logf("Total server requests: %d (includes retries)", actualServerRequests)
	t.Logf("Total tracked calls: %d", totalCallCount)
	t.Logf("Total tracked errors: %d", totalErrCount)
	t.Logf("Total tracked retries: %d", totalRetryCount)
	t.Logf("Total duration: %v", totalDuration)
}

// TestStatsIncrement specifically tests the increment function for race conditions
func TestStatsIncrement(t *testing.T) {
	var s stats

	numGoroutines := 100
	incrementsPerGoroutine := 1000
	expectedTotal := uint64(numGoroutines * incrementsPerGoroutine)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Concurrently increment the same key
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				s.Count("test-endpoint")
			}
		}()
	}

	wg.Wait()

	// Verify the count is exactly what we expect
	actual := s.count.get("test-endpoint")
	assert.Equal(t, expectedTotal, actual,
		"Count should be exactly the sum of all increments")
}

// TestStatsDurationIncrement tests that duration increments work correctly
func TestStatsDurationIncrement(t *testing.T) {
	var s stats

	numGoroutines := 50
	incrementsPerGoroutine := 100
	incrementAmount := 10 * time.Millisecond
	expectedTotal := time.Duration(numGoroutines*incrementsPerGoroutine) * incrementAmount

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Concurrently increment durations
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				s.Duration("test-endpoint", incrementAmount)
			}
		}()
	}

	wg.Wait()

	// Verify the duration is exactly what we expect
	actual := s.totalDuration.get("test-endpoint")
	assert.Equal(t, expectedTotal, actual,
		"Duration should be exactly the sum of all increments")
}

// WithLogger option for testing
func WithLogger(logger *slog.Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

// WithGatherStats option for testing
func WithGatherStats(gather bool) Option {
	return func(c *Config) {
		c.GatherStats = gather
	}
}

// WithMaxRetries option for testing
func WithMaxRetries(maxRetries int) Option {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}
