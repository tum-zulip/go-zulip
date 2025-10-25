package testutils

import (
    "encoding/json"
    "os"
    "path/filepath"
    "testing"
    "time"

    "github.com/stretchr/testify/require"
    "github.com/tum-zulip/go-zulip/zulip/client/statistics"
)

func TestWriteStatisticsToJSONAggregatesAcrossWrites(t *testing.T) {
    tempDir := t.TempDir()
    statsFile := filepath.Join(tempDir, "stats.json")
    t.Setenv("GO_ZULIP_STATS_FILE", statsFile)

    first := map[string]statistics.Statistic{
        "/endpoint": {
            Count:         1,
            ErrCount:      1,
            RetryCound:    0,
            TotalDuration: time.Second,
        },
    }

    second := map[string]statistics.Statistic{
        "/endpoint": {
            Count:         2,
            ErrCount:      0,
            RetryCound:    1,
            TotalDuration: 2 * time.Second,
        },
    }

    third := map[string]statistics.Statistic{
        "/other": {
            Count:         5,
            ErrCount:      2,
            RetryCound:    0,
            TotalDuration: 500 * time.Millisecond,
        },
    }

    writeStatisticsToJSON("alice@example.com", first)
    writeStatisticsToJSON("alice@example.com", second)
    writeStatisticsToJSON("bob@example.com", third)

    contents, err := os.ReadFile(statsFile)
    require.NoError(t, err)

    data := make(map[string]map[string]statistics.Statistic)
    require.NoError(t, json.Unmarshal(contents, &data))

    alice := data["alice@example.com"]["/endpoint"]
    require.EqualValues(t, 3, alice.Count)
    require.EqualValues(t, 1, alice.ErrCount)
    require.EqualValues(t, 1, alice.RetryCound)
    require.Equal(t, 3*time.Second, alice.TotalDuration)

    bob := data["bob@example.com"]["/other"]
    require.EqualValues(t, 5, bob.Count)
    require.EqualValues(t, 2, bob.ErrCount)
    require.EqualValues(t, 0, bob.RetryCound)
    require.Equal(t, 500*time.Millisecond, bob.TotalDuration)
}
