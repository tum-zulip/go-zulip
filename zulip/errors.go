package zulip

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

// APIError Provides access to the body, error and model on returned errors.
type APIError struct {
	body  []byte
	error string

	// the model error returned by the API
	model interface{}
}

// CodedError struct for CodedError
type CodedError struct {
	Response

	// A string that identifies the error.
	Code string `json:"code"`
}

func (e CodedError) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.Result, e.Msg, e.Code)
}

// RateLimitedError struct for RateLimitedError
type RateLimitedError struct {
	CodedError

	// How many seconds the client must wait before making additional requests.
	RetryAfter time.Duration `json:"retry-after"`
}

type rateLimitErrorJSON struct {
	CodedError
	RetryAfter int64 `json:"retry-after"`
}

func (o *RateLimitedError) MarshalJSON() ([]byte, error) {
	secs := math.Ceil(o.RetryAfter.Seconds())
	if secs < 0 {
		return nil, fmt.Errorf("retry-after cannot be negative")
	}
	model := rateLimitErrorJSON{
		CodedError: o.CodedError,
		RetryAfter: int64(secs),
	}

	return json.Marshal(model)
}

func (o *RateLimitedError) UnmarshalJSON(data []byte) error {
	var model rateLimitErrorJSON

	err := json.Unmarshal(data, &model)
	if err != nil {
		return err
	}

	o.CodedError = model.CodedError
	o.RetryAfter = time.Duration(model.RetryAfter) * time.Second
	return nil
}

// BadEventQueueIdError struct for BadEventQueueIdError
type BadEventQueueIdError struct {
	CodedError

	// The string that identifies the invalid event queue.
	QueueId string `json:"queue_id"`
}

// NonExistingChannelIdError struct for NonExistingChannelIdError
type NonExistingChannelIdError struct {
	CodedError

	// The channel Id that could not be found.
	ChannelId int64 `json:"stream_id"`
}

// InvitationFailedError struct for InvitationFailedError
type InvitationFailedError struct {
	CodedError

	// An array of arrays of length 3, where each inner array consists of (a) an email address that was skipped while sending invitations, (b) the corresponding error message, and (c) a boolean which is `true` when the email address already uses Zulip and the corresponding user is deactivated in the organization.
	Errors [][]interface{} `json:"errors"`
	// A boolean specifying whether any invitations were sent.
	SentInvitations bool `json:"sent_invitations"`
	// A boolean specifying whether the limit on the number of invitations that can be sent in the organization in a day has been reached.
	DailyLimitReached bool `json:"daily_limit_reached"`
	// A boolean specifying whether the organization have enough unused Zulip licenses to invite specified number of users.
	LicenseLimitReached bool `json:"license_limit_reached"`
}

// NonExistingChannelNameError struct for NonExistingChannelNameError
type NonExistingChannelNameError struct {
	CodedError

	// The name of the channel that could not be found.
	Channel string `json:"stream"`
}

// IncompatibleParametersError struct for IncompatibleParametersError
type IncompatibleParametersError struct {
	CodedError

	// A string containing the parameters, separated by commas, that are incompatible.
	Parameters string `json:"parameters"`
}

// MissingArgumentError struct for MissingArgumentError
type MissingArgumentError struct {
	CodedError

	// It contains the information about the missing parameter.
	VarName string `json:"var_name"`
}

// DeactivateOwnUserError struct for DeactivateOwnUserError
type DeactivateOwnUserError struct {
	CodedError

	// An internationalized string that notes if the current user is the only organization owner or the only user in the organization.
	Entity string `json:"entity"`
	// Whether the current user is the only organization owner (meaning there are other active users in the organization) or the only active user in the organization.
	IsLastOwner bool `json:"is_last_owner"`
}

// Error returns non-empty string if there was an error.
func (e APIError) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.error, e.model, string(e.body))
}

// Body returns the raw bytes of the response
func (e APIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e APIError) Model() interface{} {
	return e.model
}

type ErrInvalidEnumValue struct {
	Value   interface{}
	Enum    interface{}
	VarName string
}

func (e *ErrInvalidEnumValue) Error() string {
	return fmt.Sprintf("invalid value '%v' for '%v', valid values are %v", e.Value, e.VarName, e.Enum)
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
