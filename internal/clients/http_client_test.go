package clients

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tum-zulip/go-zulip/zulip"
)

type recordingRoundTripper struct {
	req *http.Request
}

func (t *recordingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	t.req = req
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(http.NoBody),
		Header:     make(http.Header),
	}, nil
}

func TestBuildHTTPClient_PreservesCustomRoundTripper(t *testing.T) {
	customTransport := &recordingRoundTripper{}

	cfg, err := NewConfig(&zulip.RC{
		Site:   "https://zulip.example.com",
		Email:  "test@example.com",
		APIKey: "test-api-key",
	}, func(c *Config) {
		c.HTTPClient = &http.Client{
			Transport: customTransport,
			Timeout:   5 * time.Second,
		}
	})
	require.NoError(t, err)

	require.Equal(t, 5*time.Second, cfg.HTTPClient.Timeout)
	authTransport, ok := cfg.HTTPClient.Transport.(*basicAuthRoundTripper)
	require.True(t, ok)
	require.Same(t, customTransport, authTransport.base)

	req, err := http.NewRequest(http.MethodGet, "https://zulip.example.com/api/v1/users/me", nil)
	require.NoError(t, err)
	resp, err := cfg.HTTPClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.NotNil(t, customTransport.req)
	email, apiKey, ok := customTransport.req.BasicAuth()
	require.True(t, ok)
	assert.Equal(t, "test@example.com", email)
	assert.Equal(t, "test-api-key", apiKey)
}

func TestBuildHTTPClient_ClonesCustomHTTPTransport(t *testing.T) {
	customTransport := &http.Transport{
		MaxIdleConns:          42,
		MaxIdleConnsPerHost:   7,
		ResponseHeaderTimeout: 3 * time.Second,
	}

	cfg, err := NewConfig(&zulip.RC{
		Site:   "https://zulip.example.com",
		Email:  "test@example.com",
		APIKey: "test-api-key",
	}, func(c *Config) {
		c.HTTPClient = &http.Client{
			Transport: customTransport,
		}
	})
	require.NoError(t, err)

	authTransport, ok := cfg.HTTPClient.Transport.(*basicAuthRoundTripper)
	require.True(t, ok)

	clonedTransport, ok := authTransport.base.(*http.Transport)
	require.True(t, ok)
	require.NotSame(t, customTransport, clonedTransport)
	assert.Equal(t, customTransport.MaxIdleConns, clonedTransport.MaxIdleConns)
	assert.Equal(t, customTransport.MaxIdleConnsPerHost, clonedTransport.MaxIdleConnsPerHost)
	assert.Equal(t, customTransport.ResponseHeaderTimeout, clonedTransport.ResponseHeaderTimeout)
}
