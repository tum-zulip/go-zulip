package realtimeevents

import (
	"context"
	"sync"
	"time"

	"github.com/tum-zulip/go-zulip/zulip"
)

// TypingNotifier manages automatic typing notifications.
// It sends periodic "start" notifications at the server-specified interval
// and sends a "stop" notification when Close is called.
//
// Call Close when the user stops typing, cancels, or submits the message.
type TypingNotifier struct {
	sender    TypingSender
	queue     *eventQueue
	recipient zulip.Recipient
	ctx       context.Context
	cancel    context.CancelFunc

	mu            sync.Mutex
	active        bool
	keepAliveDone chan struct{}
}

// Close stops the typing indicator and sends a "stop" notification.
// Safe to call multiple times.
func (t *TypingNotifier) Close() error {
	t.mu.Lock()

	if !t.active {
		t.mu.Unlock()
		return nil
	}

	t.active = false
	t.cancel()
	keepAliveDone := t.keepAliveDone

	t.mu.Unlock()

	if keepAliveDone != nil {
		<-keepAliveDone
	}

	_, _, err := t.sender.SetTypingStatus(context.Background()).
		To(t.recipient).
		Op(zulip.TypingStatusOpStop).
		Execute()

	return err
}

// runKeepAlive sends periodic "start" notifications at server-specified interval.
func (t *TypingNotifier) runKeepAlive() {
	defer close(t.keepAliveDone)

	settings := t.queue.GetRealmSettings()
	ticker := time.NewTicker(
		time.Duration(settings.ServerTypingStartedWaitPeriodMilliseconds) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-t.ctx.Done():
			return
		case <-ticker.C:
			t.mu.Lock()
			active := t.active
			t.mu.Unlock()

			if !active {
				return
			}

			//nolint:errcheck // Network errors shouldn't kill typing; server will expire indicator
			t.sender.SetTypingStatus(t.ctx).
				To(t.recipient).
				Op(zulip.TypingStatusOpStart).
				Execute()
		}
	}
}
