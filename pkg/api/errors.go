package api

import (
	"fmt"
	"reflect"
	"strings"
)

type ErrInvalidEnumValue struct {
	Value   interface{}
	Enum    interface{}
	VarName string
}

func (e *ErrInvalidEnumValue) Error() string {
	return fmt.Sprintf("invalid value '%v' for '%v', valid values are %v", e.Value, e.VarName, e.Enum)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// CodedError struct for CodedError
type CodedError struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
	// A string that identifies the error.
	Code string `json:"code"`
}

// RateLimitedError struct for RateLimitedError
type RateLimitedError struct {
	CodedError

	// How many seconds the client must wait before making additional requests.
	RetryAfter int `json:"retry-after,omitempty"`
}

// BadEventQueueIdError struct for BadEventQueueIdError
type BadEventQueueIdError struct {
	CodedError

	// The string that identifies the invalid event queue.
	QueueId string `json:"queue_id,omitempty"`
}

// NonExistingChannelIdError struct for NonExistingChannelIdError
type NonExistingChannelIdError struct {
	CodedError

	// The channel Id that could not be found.
	StreamId int64 `json:"stream_id,omitempty"`
}

// InvitationFailedError struct for InvitationFailedError
type InvitationFailedError struct {
	CodedError

	// An array of arrays of length 3, where each inner array consists of (a) an email address that was skipped while sending invitations, (b) the corresponding error message, and (c) a boolean which is `true` when the email address already uses Zulip and the corresponding user is deactivated in the organization.
	Errors [][]interface{} `json:"errors,omitempty"`
	// A boolean specifying whether any invitations were sent.
	SentInvitations bool `json:"sent_invitations,omitempty"`
	// A boolean specifying whether the limit on the number of invitations that can be sent in the organization in a day has been reached.
	DailyLimitReached bool `json:"daily_limit_reached,omitempty"`
	// A boolean specifying whether the organization have enough unused Zulip licenses to invite specified number of users.
	LicenseLimitReached bool `json:"license_limit_reached,omitempty"`
}

// NonExistingChannelNameError struct for NonExistingChannelNameError
type NonExistingChannelNameError struct {
	CodedError

	// The name of the channel that could not be found.
	Stream string `json:"stream,omitempty"`
}

// IncompatibleParametersError struct for IncompatibleParametersError
type IncompatibleParametersError struct {
	CodedError

	// A string containing the parameters, separated by commas, that are incompatible.
	Parameters string `json:"parameters,omitempty"`
}

// MissingArgumentError struct for MissingArgumentError
type MissingArgumentError struct {
	CodedError

	// It contains the information about the missing parameter.
	VarName string `json:"var_name,omitempty"`
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
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
