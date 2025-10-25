package statistics

import "time"

type Statistic struct {
	Count         uint64
	ErrCount      uint64
	RetryCound    uint64
	TotalDuration time.Duration
}
