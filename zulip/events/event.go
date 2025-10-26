package events

// Event is the interface that all Zulip events implement.
// Use type assertion or GetType() to determine the specific event type and handle accordingly.
type Event interface {
	// GetType returns the EventType of this event (e.g., EventTypeMessage, EventTypeChannel).
	GetType() EventType
	// GetId returns the unique ID of this event. Events appear in increasing order but may not be consecutive.
	GetId() int64
	// GetOp returns the operation type and a boolean indicating if an operation exists.
	GetOp() (EventOp, bool)
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

// GetOp implements Event.GetOp for event by returning ("", false) (no operation).
func (e event) GetOp() (EventOp, bool) {
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
