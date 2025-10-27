package clients

import (
	"sync"
	"time"

	"github.com/tum-zulip/go-zulip/zulip/client/statistics"
)

type stats struct {
	count         stat[uint64]
	errCount      stat[uint64]
	retryCount    stat[uint64]
	totalDuration stat[time.Duration]
}

func (s *stats) Count(key string) {
	s.count.increment(key, 1)
}

func (s *stats) Error(key string) {
	s.errCount.increment(key, 1)
}

func (s *stats) Retry(key string) {
	s.retryCount.increment(key, 1)
}

func (s *stats) Duration(key string, duration time.Duration) {
	s.totalDuration.increment(key, duration)
}

type integer interface{ ~uint64 | ~int64 }

type stat[T integer] struct {
	sync.Map
}

func (m *stat[T]) increment(key string, value T) {
	c, _ := m.LoadOrStore(key, T(0))
	current := c.(T)

	for !m.CompareAndSwap(key, current, current+value) {
		c, _ := m.Load(key)
		current = c.(T)
	}
}

func (m *stat[T]) get(key string) T {
	v, ok := m.Load(key)
	if !ok {
		var t T
		return t
	}
	t, ok := v.(T)
	if !ok {
		panic("unexpected type in map")
	}
	return t
}

func (s *stats) GetStatistics() statistics.Statistics {
	stats := make(map[string]statistics.Statistic)
	if s == nil {
		return statistics.Statistics{Stats: stats}
	}

	s.totalDuration.Range(func(k any, v any) bool {
		endpoint := k.(string)
		duration := v.(time.Duration)

		stats[endpoint] = statistics.Statistic{
			TotalDuration: duration,
			Count:         s.count.get(endpoint),
			ErrCount:      s.errCount.get(endpoint),
			RetryCount:    s.retryCount.get(endpoint),
		}
		return true
	})
	return statistics.Statistics{Stats: stats}
}
