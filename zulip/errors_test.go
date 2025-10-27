package zulip_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	z "github.com/tum-zulip/go-zulip/zulip"
)

func Test_RateLimitError(t *testing.T) {
	data := []byte(
		`{"result":"error","msg":"API usage exceeded rate limit","code":"RATE_LIMIT_HIT","retry-after":30.5}`,
	)
	var zulipErr z.RateLimitedError

	err := json.Unmarshal(data, &zulipErr)
	require.NoError(t, err)

	require.Equal(t, 30*time.Second+500*time.Millisecond, zulipErr.RetryAfter)
}
