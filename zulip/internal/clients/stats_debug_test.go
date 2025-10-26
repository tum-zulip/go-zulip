package clients

import (
	"testing"
	"time"
)

// TestDurationIncrementDebug tests duration increment in isolation
func TestDurationIncrementDebug(t *testing.T) {
	var s stats

	// Add a few durations
	s.Duration("test", 100*time.Millisecond)
	s.Duration("test", 50*time.Millisecond)
	s.Duration("test", 25*time.Millisecond)

	result := s.totalDuration.get("test")
	expected := 175 * time.Millisecond

	t.Logf("Expected: %v, Got: %v", expected, result)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
