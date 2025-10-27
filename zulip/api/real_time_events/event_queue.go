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
	client       APIRealTimeEvents
	logger       *slog.Logger
	lastEventID  atomic.Int64
	queueID      string
	bufferSize   int
	cancel       chan struct{}
	errorHandler EventQueueErrorHandler
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
