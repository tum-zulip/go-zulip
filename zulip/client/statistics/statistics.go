// Package statistics defines structures for tracking API client statistics.
package statistics

import "time"

type Statistics struct {
	Stats map[string]Statistic `json:"stats,omitempty"`
}

type Statistic struct {
	Count         uint64        `json:"count,omitempty"`
	ErrCount      uint64        `json:"err_count,omitempty"`
	RetryCount    uint64        `json:"retry_count,omitempty"`
	TotalDuration time.Duration `json:"total_duration,omitempty"`
}
