package clients

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tum-zulip/go-zulip/zulip"
)

func TestSimpleClient_DoHTTPCall_SanitizesCredentials(t *testing.T) {
	// Create a test server that returns a response with sensitive cookie
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Set-Cookie", "session=sensitive-cookie-123")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"success"}`))
	}))
	defer server.Close()

	// Create a buffer to capture logs
	var logBuffer bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logBuffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Create client config with test credentials
	rc := &zulip.RC{
		Site:   server.URL,
		Email:  "test@example.com",
		APIKey: "test-api-key-secret",
	}

	cfg, err := NewConfig(rc, func(c *Config) {
		c.Logger = logger
	})
	require.NoError(t, err)

	client := NewSimpleClient(cfg)

	// Make a request
	req, err := http.NewRequest(http.MethodGet, server.URL+"/api/v1/users/me", nil)
	require.NoError(t, err)

	_, err = client.doHTTPCall(context.Background(), "/users/me", req)
	require.NoError(t, err)

	logOutput := logBuffer.String()

	// Verify sensitive cookie value is NOT in logs
	assert.NotContains(t, logOutput, "session=sensitive-cookie-123", "Cookie values should not appear in logs")

	// Verify Set-Cookie header is sanitized
	assert.Contains(t, logOutput, "Set-Cookie:", "Should log Set-Cookie header name")
	// The cookie value should be redacted (not the literal string)
	if assert.Contains(t, logOutput, "Set-Cookie:", "Set-Cookie header should be present") {
		// Verify it's been sanitized - should not contain the actual value
		assert.NotContains(t, logOutput, "session=sensitive-cookie-123")
	}

	// Verify we still log debug info (non-sensitive parts)
	assert.Contains(t, logOutput, "HTTP Request", "Should log request marker")
	assert.Contains(t, logOutput, "HTTP Response", "Should log response marker")
}

func TestSimpleClient_DoHTTPCall_PreservesNonSensitiveHeaders(t *testing.T) {
	// Create a test server with non-sensitive headers
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Custom-Header", "custom-value")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"success"}`))
	}))
	defer server.Close()

	var logBuffer bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logBuffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	rc := &zulip.RC{
		Site:   server.URL,
		Email:  "test@example.com",
		APIKey: "test-api-key",
	}

	cfg, err := NewConfig(rc, func(c *Config) {
		c.Logger = logger
	})
	require.NoError(t, err)

	client := NewSimpleClient(cfg)

	req, err := http.NewRequest(http.MethodGet, server.URL+"/api/v1/users/me", nil)
	require.NoError(t, err)

	_, err = client.doHTTPCall(context.Background(), "/users/me", req)
	require.NoError(t, err)

	logOutput := logBuffer.String()

	// Non-sensitive headers should be preserved
	assert.Contains(t, logOutput, "Content-Type", "Should preserve Content-Type header")
	assert.Contains(t, logOutput, "application/json", "Should preserve Content-Type value")
	assert.Contains(t, logOutput, "X-Custom-Header", "Should preserve custom header")
	assert.Contains(t, logOutput, "custom-value", "Should preserve custom header value")
}

func TestSimpleClient_DoHTTPCall_NoLogsWhenDebugDisabled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"success"}`))
	}))
	defer server.Close()

	var logBuffer bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logBuffer, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Debug disabled
	}))

	rc := &zulip.RC{
		Site:   server.URL,
		Email:  "test@example.com",
		APIKey: "test-api-key-secret",
	}

	cfg, err := NewConfig(rc, func(c *Config) {
		c.Logger = logger
	})
	require.NoError(t, err)

	client := NewSimpleClient(cfg)

	req, err := http.NewRequest(http.MethodGet, server.URL+"/api/v1/users/me", nil)
	require.NoError(t, err)

	_, err = client.doHTTPCall(context.Background(), "/users/me", req)
	require.NoError(t, err)

	logOutput := logBuffer.String()

	// No HTTP dumps should be logged when debug is disabled
	assert.NotContains(t, logOutput, "HTTP Request", "Should not log request when debug disabled")
	assert.NotContains(t, logOutput, "HTTP Response", "Should not log response when debug disabled")
}

func TestSimpleClient_DoHTTPCall_WithRequestAuth(t *testing.T) {
	// This test verifies that if a request has an Authorization header
	// (e.g., manually added for testing), it gets sanitized in logs
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"success"}`))
	}))
	defer server.Close()

	var logBuffer bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logBuffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	rc := &zulip.RC{
		Site:   server.URL,
		Email:  "test@example.com",
		APIKey: "test-api-key-secret",
	}

	cfg, err := NewConfig(rc, func(c *Config) {
		c.Logger = logger
	})
	require.NoError(t, err)

	client := NewSimpleClient(cfg)

	// Create a request and manually add an Authorization header for testing
	req, err := http.NewRequest(http.MethodGet, server.URL+"/api/v1/users/me", nil)
	require.NoError(t, err)
	req.SetBasicAuth("test@example.com", "test-api-key-secret")

	_, err = client.doHTTPCall(context.Background(), "/users/me", req)
	require.NoError(t, err)

	logOutput := logBuffer.String()

	// Verify the credentials are not leaked in logs
	assert.NotContains(t, logOutput, "test-api-key-secret", "Should not log API key")
	assert.NotContains(t, logOutput, "test@example.com:test-api-key-secret", "Should not log email:apikey")

	// The Authorization header should be sanitized if present
	// Note: Due to how httputil.DumpRequestOut works with the transport layer,
	// the Authorization header might not always appear in the dump, but if it does,
	// it should be sanitized
	if assert.Contains(t, logOutput, "Authorization:", "If Authorization header is logged") {
		assert.Contains(t, logOutput, "[REDACTED]", "It should be sanitized")
	}
}
