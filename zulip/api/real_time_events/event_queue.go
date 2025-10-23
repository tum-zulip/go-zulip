package real_time_events

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"
	"syscall"

	"github.com/tum-zulip/go-zulip/zulip"
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
	// queueId must reference an already-registered Zulip event queue. The helper
	// does not register queues automatically. lastEventId is the highest event ID
	// the caller has processed so far; pass -1 to start with the server default.
	Connect(ctx context.Context, queueId string, lastEventId int64) (<-chan zulip.Event, error)
	// Close stops the polling goroutine. It does not delete the queue from the
	// server. Calling Close before Connect returns an error.
	Close() error
	// QueueId reports the identifier of the queue currently being consumed.
	QueueId() string
	// LastEventId reports the highest event ID delivered through this queue. If
	// no events have been received, it returns the value supplied to Connect.
	LastEventId() int64
}

type loggingErrorHandler struct{}

func (h *loggingErrorHandler) Handle(ctx context.Context, resp *http.Response, err error) {
	slog.ErrorContext(ctx, "event queue error", "error", err)
}

// NewEventQueueLoggingErrorHandler returns an EventQueueErrorHandler that logs
// polling errors using slog at the error level.
func NewEventQueueLoggingErrorHandler() EventQueueErrorHandler {
	return &loggingErrorHandler{}
}

type channelErrorHandler struct {
	errs chan error
}

func (h *channelErrorHandler) Handle(ctx context.Context, resp *http.Response, err error) {
	select {
	case <-ctx.Done():
		slog.ErrorContext(ctx, "event queue context cancelled while handling error")
		return
	case h.errs <- err:
		// Error sent
	}
}

// NewEventQueueChannelErrorHandler returns an EventQueueErrorHandler that
// forwards polling errors to the supplied channel until the context is cancelled.
func NewEventQueueChannelErrorHandler(errs chan error) EventQueueErrorHandler {
	return &channelErrorHandler{errs: errs}
}

type eventQueue struct {
	client       APIRealTimeEvents
	lastEventId  atomic.Int64
	queueId      string
	cancel       context.CancelFunc
	errorHandler EventQueueErrorHandler
}

var _ EventQueue = (*eventQueue)(nil)

// NewEventQueue returns an EventQueue implementation that polls the Zulip
// RealTime Events API using the provided client. The handler is optional; when
// nil, polling errors are ignored.
func NewEventQueue(client APIRealTimeEvents, handler EventQueueErrorHandler) *eventQueue {
	return &eventQueue{
		client:       client,
		lastEventId:  atomic.Int64{},
		errorHandler: handler,
	}
}

func (q *eventQueue) QueueId() string {
	return q.queueId
}

func (q *eventQueue) LastEventId() int64 {
	return q.lastEventId.Load()
}

func (q *eventQueue) Connect(ctx context.Context, queueId string, lastEventId int64) (<-chan zulip.Event, error) {
	// Create cancellable context for the polling goroutine
	ctx, q.cancel = context.WithCancel(ctx)

	if queueId == "" {
		q.cancel = nil
		return nil, fmt.Errorf("queueId cannot be empty")
	}

	q.queueId = queueId
	q.lastEventId.Store(lastEventId)

	// Verify that the queue exists
	if err := q.checkQueueExists(ctx); err != nil {
		q.cancel = nil
		return nil, err
	}

	events := make(chan zulip.Event, 32) // default buffer size

	// Start polling in background goroutine
	go q.pollEvents(ctx, events)

	return events, nil
}

func (q *eventQueue) Close() error {
	if q.cancel == nil {
		return fmt.Errorf("event queue not connected")
	}
	q.cancel()
	return nil
}

func (q *eventQueue) pollEvents(ctx context.Context, events chan<- zulip.Event) {
	defer close(events)

	running := true

	for running {
		select {
		case <-ctx.Done():
			slog.DebugContext(ctx, "event queue context cancelled, stopping poll loop")
			running = false
			return
		default:
		}

		resp, httpResp, err := q.client.GetEvents(ctx).
			QueueId(q.QueueId()).
			LastEventId(q.LastEventId()).
			Execute()

		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				slog.DebugContext(ctx, "event queue stopping after context cancellation", "error", err)
			} else {
				if q.errorHandler != nil {
					q.errorHandler.Handle(ctx, httpResp, err)
				}

				if shouldStopPolling(err) {
					slog.WarnContext(ctx, "event queue stopping due to polling error", "error", err)
					return
				} else if q.errorHandler == nil {
					slog.ErrorContext(ctx, "event queue polling error", "error", err)
				}
			}
			continue
		}

		if resp == nil {
			slog.WarnContext(ctx, "no events received from server, but also no error")
			continue
		}
		running = q.processEvents(ctx, resp, events)
	}
}

func (q *eventQueue) processEvents(ctx context.Context, resp *GetEventsResponse, events chan<- zulip.Event) bool {
	// Process events from response
	lastId, valid := int64(0), false

	defer func() {
		// Update last event ID if we processed any events
		if valid {
			q.lastEventId.Store(lastId)
		}

		slog.DebugContext(ctx, "polled events", "count", len(resp.Events), "last_event_id", lastId)
	}()

	for _, event := range resp.Events {
		if event == nil {
			slog.WarnContext(ctx, "received nil event from server")
			continue
		}
		select {
		case <-ctx.Done():
			slog.DebugContext(ctx, "event queue context cancelled while processing events")
			return false
		case events <- event:
			lastId = event.GetId()
			valid = true
		}
	}
	return true
}

func (q *eventQueue) checkQueueExists(ctx context.Context) error {
	_, _, err := q.client.GetEvents(ctx).
		QueueId(q.QueueId()).
		LastEventId(q.LastEventId()).
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

	if netErr, ok := err.(net.Error); ok {
		return !netErr.Timeout()
	}

	return false
}
