package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func Test_RateLimitError(t *testing.T) {
	data := []byte(`{"result":"error","msg":"API usage exceeded rate limit","code":"RATE_LIMIT_HIT","retry-after":30.5}`)
	var zulipErr zulip.RateLimitedError

	err := json.Unmarshal(data, &zulipErr)
	require.NoError(t, err)

	require.Equal(t, zulipErr.RetryAfter, time.Duration(30*time.Second+500*time.Millisecond))

}
