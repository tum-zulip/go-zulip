package api_client

import (
	"sync"
	"time"
)

type Statistics struct {
	sync.Map
}

func (m *Statistics) Increment(key string, value time.Duration) {
	c, _ := m.LoadOrStore(key, time.Duration(0))
	current := c.(time.Duration)

	for !m.CompareAndSwap(key, current, current+value) {
		value, _ := m.Load(key)
		current = value.(time.Duration)
	}
}

func (m *Statistics) GetStatistics() map[string]time.Duration {
	stats := make(map[string]time.Duration)
	if m == nil {
		return stats
	}

	m.Range(func(k any, v any) bool {
		endpoint := k.(string)
		duration := v.(time.Duration)
		stats[endpoint] = duration
		return true
	})
	return stats
}
