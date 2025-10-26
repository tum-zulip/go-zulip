package events

// Event is the interface that all Zulip events implement.
// Use type assertion or GetType() to determine the specific event type and handle accordingly.
type Event interface {
	// GetType returns the EventType of this event (e.g., EventTypeMessage, EventTypeChannel).
	GetType() EventType
	// GetId returns the unique ID of this event. Events appear in increasing order but may not be consecutive.
	GetId() int64
	// HasOp returns true if this event includes an operation type field (Op).
	HasOp() bool
	// GetOp returns the operation type for this event (e.g., EventOpAdd, EventOpRemove, EventOpUpdate).
	// Only meaningful if HasOp() returns true; otherwise returns empty string.
	GetOp() EventOp
	// GetOpOk returns the operation type and a boolean indicating if an operation exists.
	// Useful for safely checking if an operation is present.
	GetOpOk() (EventOp, bool)
}

type event struct {
	// The Id of the event. Events appear in increasing order but may not be consecutive.
	Id   int64     `json:"id,omitempty"`
	Type EventType `json:"type,omitempty"`
	Op   EventOp   `json:"op,omitempty"`
}

// GetType implements Event.GetType by returning the Type field.
func (e event) GetType() EventType {
	return e.Type
}

// GetId implements Event.GetId by returning the Id field.
func (e event) GetId() int64 {
	return e.Id
}

// HasOp implements Event.HasOp for event by returning false (no operation field).
func (e event) HasOp() bool {
	return false
}

// GetOp implements Event.GetOp for event by returning empty string (no operation).
func (e event) GetOp() EventOp {
	return e.Op
}

// GetOpOk implements Event.GetOpOk for event by returning (nil, false) (no operation).
func (e event) GetOpOk() (EventOp, bool) {
	return e.Op, e.Op != ""
}

type EventUnmarshalingError struct {
	event

	Type EventType
	Err  error
	Data []byte
}

func (e EventUnmarshalingError) Error() string {
	return e.Err.Error()
}
