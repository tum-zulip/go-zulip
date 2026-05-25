package realtimeevents

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"
	"syscall"

	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/users"
	"github.com/tum-zulip/go-zulip/zulip/events"
)

// EventQueueErrorHandler receives asynchronous errors produced while polling events.
// Implementations can log, buffer, or otherwise handle the failure before the
// next polling attempt is made.
type EventQueueErrorHandler interface {
	Handle(ctx context.Context, resp *http.Response, err error)
}

// EventQueue represents a long-lived consumer of a Zulip event queue that has
// already been registered with the server.
type EventQueue interface {
	// Connect starts the event queue polling loop in a background goroutine and
	// returns a channel that receives events as they arrive. The goroutine runs
	// until the provided context is cancelled or Close is invoked.
	//
	// queueID must reference an already-registered Zulip event queue. The helper
	// does not register queues automatically. lastEventID is the highest event ID
	// the caller has processed so far; pass -1 to start with the server default.
	Connect(ctx context.Context, queueID string, lastEventID int64) (<-chan events.Event, error)
	// Close stops the polling goroutine. It does not delete the queue from the
	// server. Calling Close before Connect returns an error.
	Close() error
	// QueueID reports the identifier of the queue currently being consumed.
	QueueID() string
	// LastEventID reports the highest event ID delivered through this queue. If
	// no events have been received, it returns the value supplied to Connect.
	LastEventID() int64
	// UpdateRealmSettings loads realm settings from a RegisterQueueResponse.
	// Should be called after RegisterQueue to initialize settings from server.
	UpdateRealmSettings(resp *RegisterQueueResponse)
	// GetRealmSettings returns the current realm settings. Thread-safe.
	GetRealmSettings() *RealmSettings
	// StartTyping begins sending typing notifications for the given recipient.
	// Returns a TypingNotifier that must be closed when the user stops typing.
	// The typingSender must implement SetTypingStatus (typically client.Client).
	StartTyping(ctx context.Context, typingSender TypingSender, recipient zulip.Recipient) (*TypingNotifier, error)
}

// TypingSender is the interface required to send typing status notifications.
// This is satisfied by client.Client and users.APIUsers.
type TypingSender interface {
	SetTypingStatus(ctx context.Context) users.SetTypingStatusRequest
}

// RealmSettings stores typing-related realm configuration.
// Instances are immutable; updates create new instances.
type RealmSettings struct {
	ServerTypingStartedWaitPeriodMilliseconds   int64
	ServerTypingStoppedWaitPeriodMilliseconds   int64
	ServerTypingStartedExpiryPeriodMilliseconds int64
}

const (
	defaultTypingStartedWaitPeriodMs   = 10000
	defaultTypingStoppedWaitPeriodMs   = 5000
	defaultTypingStartedExpiryPeriodMs = 15000
)

func defaultRealmSettings() *RealmSettings {
	return &RealmSettings{
		ServerTypingStartedWaitPeriodMilliseconds:   defaultTypingStartedWaitPeriodMs,
		ServerTypingStoppedWaitPeriodMilliseconds:   defaultTypingStoppedWaitPeriodMs,
		ServerTypingStartedExpiryPeriodMilliseconds: defaultTypingStartedExpiryPeriodMs,
	}
}

// WithLogger returns an EventQueueErrorHandler that logs
// polling errors using slog at the error level.
func WithLogger(logger *slog.Logger) EventQueueOption {
	if logger == nil {
		panic("logger cannot be nil")
	}
	return func(q *eventQueue) {
		q.logger = logger
	}
}

type channelErrorHandler struct {
	logger *slog.Logger
	errs   chan error
}

func (h *channelErrorHandler) Handle(ctx context.Context, _ *http.Response, err error) {
	select {
	case <-ctx.Done():
		h.logger.ErrorContext(ctx, "event queue context cancelled while handling error")
		return
	case h.errs <- err:
		// Error sent
	}
}

// NewEventQueueChannelErrorHandler returns an EventQueueErrorHandler that
// forwards polling errors to the supplied channel until the context is cancelled.
func WithEventQueueChannelErrorHandler(logger *slog.Logger, errs chan error) EventQueueOption {
	if logger == nil {
		panic("logger cannot be nil")
	}
	if errs == nil {
		panic("errs channel cannot be nil")
	}
	return func(q *eventQueue) {
		q.errorHandler = &channelErrorHandler{logger: logger, errs: errs}
	}
}

// WithEventQueueChannelBufferSize sets the size of the internal event channel.
func WithEventQueueChannelBufferSize(size int) EventQueueOption {
	return func(q *eventQueue) {
		q.bufferSize = size
	}
}

type EventQueueOption func(*eventQueue)

type eventQueue struct {
	client        APIRealTimeEvents
	logger        *slog.Logger
	lastEventID   atomic.Int64
	queueID       string
	bufferSize    int
	cancel        chan struct{}
	errorHandler  EventQueueErrorHandler
	realmSettings atomic.Value // stores *RealmSettings
}

var _ EventQueue = (*eventQueue)(nil)

// NewEventQueue returns an EventQueue implementation that polls the Zulip
// RealTime Events API using the provided client. The handler is optional; when
// nil, polling errors are ignored.
func NewEventQueue(client APIRealTimeEvents, opts ...EventQueueOption) EventQueue {
	q := &eventQueue{
		client:      client,
		logger:      slog.Default(),
		lastEventID: atomic.Int64{},
	}

	q.realmSettings.Store(defaultRealmSettings())

	for _, opt := range opts {
		if opt != nil {
			opt(q)
		}
	}

	return q
}

func (q *eventQueue) QueueID() string {
	return q.queueID
}

func (q *eventQueue) LastEventID() int64 {
	return q.lastEventID.Load()
}

func (q *eventQueue) GetRealmSettings() *RealmSettings {
	settings, _ := q.realmSettings.Load().(*RealmSettings)
	return settings
}

func (q *eventQueue) UpdateRealmSettings(resp *RegisterQueueResponse) {
	settings := &RealmSettings{
		ServerTypingStartedWaitPeriodMilliseconds:   resp.ServerTypingStartedWaitPeriodMilliseconds,
		ServerTypingStoppedWaitPeriodMilliseconds:   resp.ServerTypingStoppedWaitPeriodMilliseconds,
		ServerTypingStartedExpiryPeriodMilliseconds: resp.ServerTypingStartedExpiryPeriodMilliseconds,
	}

	defaults := defaultRealmSettings()
	if settings.ServerTypingStartedWaitPeriodMilliseconds == 0 {
		settings.ServerTypingStartedWaitPeriodMilliseconds = defaults.ServerTypingStartedWaitPeriodMilliseconds
	}
	if settings.ServerTypingStoppedWaitPeriodMilliseconds == 0 {
		settings.ServerTypingStoppedWaitPeriodMilliseconds = defaults.ServerTypingStoppedWaitPeriodMilliseconds
	}
	if settings.ServerTypingStartedExpiryPeriodMilliseconds == 0 {
		settings.ServerTypingStartedExpiryPeriodMilliseconds = defaults.ServerTypingStartedExpiryPeriodMilliseconds
	}

	q.realmSettings.Store(settings)
}

func (q *eventQueue) StartTyping(
	ctx context.Context,
	typingSender TypingSender,
	recipient zulip.Recipient,
) (*TypingNotifier, error) {
	typingCtx, cancel := context.WithCancel(ctx)

	notifier := &TypingNotifier{
		sender:        typingSender,
		queue:         q,
		recipient:     recipient,
		ctx:           typingCtx,
		cancel:        cancel,
		active:        true,
		keepAliveDone: make(chan struct{}),
	}

	_, _, err := typingSender.SetTypingStatus(ctx).
		To(recipient).
		Op(zulip.TypingStatusOpStart).
		Execute()
	if err != nil {
		cancel()
		return nil, err
	}

	go notifier.runKeepAlive()

	return notifier, nil
}

func (q *eventQueue) Connect(ctx context.Context, queueID string, lastEventID int64) (<-chan events.Event, error) {
	if queueID == "" {
		return nil, errors.New("queueID cannot be empty")
	}

	q.queueID = queueID
	q.lastEventID.Store(lastEventID)

	// Verify that the queue exists using the caller's context before
	// installing signal handling. This avoids creating cancellers that
	// would be leaked if verification fails.
	if err := q.checkQueueExists(ctx); err != nil {
		return nil, err
	}

	events := make(chan events.Event, q.bufferSize)
	q.cancel = make(chan struct{})

	// Start polling in background goroutine using the wrapped context
	go q.pollEvents(ctx, events)

	return events, nil
}

func (q *eventQueue) Close() error {
	if q.cancel == nil {
		return errors.New("event queue not connected")
	}
	close(q.cancel)
	return nil
}

func (q *eventQueue) pollEvents(ctx context.Context, events chan<- events.Event) {
	defer close(events)

	running := true

	for running {
		select {
		case <-ctx.Done():
			q.logger.DebugContext(ctx, "event queue context cancelled, stopping poll loop")
			return
		default:
		}

		resp, httpResp, err := q.client.GetEvents(ctx).
			QueueID(q.QueueID()).
			LastEventID(q.LastEventID()).
			Execute()
		if err != nil {
			q.handleHTTPError(ctx, httpResp, err)
			if shouldStopPolling(err) {
				q.logger.WarnContext(ctx, "event queue stopping due to polling error", "error", err)
				return
			}
			continue
		}

		if resp == nil {
			q.logger.WarnContext(ctx, "no events received from server, but also no error")
			continue
		}
		running = q.processEvents(ctx, resp, events)
	}
}

func (q *eventQueue) handleHTTPError(ctx context.Context, httpResp *http.Response, err error) {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		q.logger.DebugContext(ctx, "event queue stopping after context cancellation", "error", err)
	} else {
		if q.errorHandler != nil {
			q.errorHandler.Handle(ctx, httpResp, err)
		} else if q.errorHandler == nil {
			q.logger.ErrorContext(ctx, "event queue polling error", "error", err)
		}
	}
}

func (q *eventQueue) processEvents(ctx context.Context, resp *GetEventsResponse, events chan<- events.Event) bool {
	// Process events from Response
	lastID, valid := int64(0), false

	defer func() {
		// Update last event ID if we processed any events
		if valid {
			q.lastEventID.Store(lastID)
		}

		q.logger.DebugContext(ctx, "polled events", "count", len(resp.Events), "last_event_id", lastID)
	}()

	for _, event := range resp.Events {
		if event == nil {
			q.logger.WarnContext(ctx, "received nil event from server")
			continue
		}

		q.updateRealmSettingsFromEvent(event)

		select {
		case <-ctx.Done():
			q.logger.DebugContext(ctx, "event queue context cancelled while processing events")
			return false
		case events <- event:
			lastID = event.GetID()
			valid = true
		}
	}
	return true
}

func (q *eventQueue) checkQueueExists(ctx context.Context) error {
	_, _, err := q.client.GetEvents(ctx).
		QueueID(q.QueueID()).
		LastEventID(q.LastEventID()).
		DontBlock(true).
		Execute()
	if err != nil {
		return err
	}
	return nil
}

func (q *eventQueue) updateRealmSettingsFromEvent(event events.Event) {
	if e, ok := event.(events.RealmUpdateEvent); ok {
		q.handleRealmUpdateEvent(e)
	}
}

func (q *eventQueue) handleRealmUpdateEvent(e events.RealmUpdateEvent) {
	current := q.GetRealmSettings()
	updated := &RealmSettings{
		ServerTypingStartedWaitPeriodMilliseconds:   current.ServerTypingStartedWaitPeriodMilliseconds,
		ServerTypingStoppedWaitPeriodMilliseconds:   current.ServerTypingStoppedWaitPeriodMilliseconds,
		ServerTypingStartedExpiryPeriodMilliseconds: current.ServerTypingStartedExpiryPeriodMilliseconds,
	}

	switch e.Property {
	case "server_typing_started_wait_period_milliseconds":
		if e.Value.Int64 != nil {
			updated.ServerTypingStartedWaitPeriodMilliseconds = *e.Value.Int64
		}
	case "server_typing_stopped_wait_period_milliseconds":
		if e.Value.Int64 != nil {
			updated.ServerTypingStoppedWaitPeriodMilliseconds = *e.Value.Int64
		}
	case "server_typing_started_expiry_period_milliseconds":
		if e.Value.Int64 != nil {
			updated.ServerTypingStartedExpiryPeriodMilliseconds = *e.Value.Int64
		}
	default:
		return
	}

	q.realmSettings.Store(updated)
	q.logger.Debug("updated realm settings from event", "property", e.Property)
}

func shouldStopPolling(err error) bool {
	if err == nil {
		return false
	}

	switch {
	case
		errors.Is(err, io.EOF),
		errors.Is(err, io.ErrUnexpectedEOF),
		errors.Is(err, net.ErrClosed):
		return true
	}

	switch {
	case errors.Is(err, syscall.ECONNRESET),
		errors.Is(err, syscall.ECONNABORTED),
		errors.Is(err, syscall.EPIPE),
		errors.Is(err, syscall.ECONNREFUSED):
		return true
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		return !netErr.Timeout()
	}

	return false
}
