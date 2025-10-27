// Package reminders provides API methods for managing Zulip scheduled reminders,
// including creating, retrieving, and deleting reminder notifications.
package reminders

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tum-zulip/go-zulip/internal/apiutils"
	"github.com/tum-zulip/go-zulip/internal/clients"
	"github.com/tum-zulip/go-zulip/zulip"
)

type APIReminders interface {
	// CreateMessageReminder Create a message reminder
	//
	// Schedule a reminder to be sent to the current user at the specified time. The reminder will link the relevant message.
	//
	// *Changes**: New in Zulip 11.0 (feature level 381).
	//
	CreateMessageReminder(ctx context.Context) CreateMessageReminderRequest

	// CreateMessageReminderExecute executes the request
	CreateMessageReminderExecute(r CreateMessageReminderRequest) (*CreateMessageReminderResponse, *http.Response, error)

	// DeleteReminder Delete a reminder
	//
	// Delete, and therefore cancel sending, a previously [scheduled reminder].
	//
	// *Changes**: New in Zulip 11.0 (feature level 399).
	//
	// [scheduled reminder]: https://zulip.com/help/schedule-a-reminder
	DeleteReminder(ctx context.Context, reminderID int64) DeleteReminderRequest

	// DeleteReminderExecute executes the request
	DeleteReminderExecute(r DeleteReminderRequest) (*zulip.Response, *http.Response, error)

	// GetReminders Get reminders
	//
	// Fetch all [reminders] for the
	// current user.
	//
	// Reminders are messages the user has scheduled to be sent in the
	// future to themself.
	//
	// *Changes**: New in Zulip 11.0 (feature level 399).
	//
	// [reminders]: https://zulip.com/help/schedule-a-reminder
	GetReminders(ctx context.Context) GetRemindersRequest

	// GetRemindersExecute executes the request
	GetRemindersExecute(r GetRemindersRequest) (*GetRemindersResponse, *http.Response, error)
}

type remindersService struct {
	client clients.Client
}

func NewRemindersService(client clients.Client) APIReminders {
	return &remindersService{client: client}
}

var _ APIReminders = (*remindersService)(nil)

type CreateMessageReminderRequest struct {
	ctx                        context.Context
	apiService                 APIReminders
	messageID                  *int64
	scheduledDeliveryTimestamp *int64
	note                       *string
}

// The ID of the previously sent message to reference in the reminder message.
func (r CreateMessageReminderRequest) MessageID(messageID int64) CreateMessageReminderRequest {
	r.messageID = &messageID
	return r
}

// The UNIX timestamp for when the reminder will be sent, in UTC seconds.
func (r CreateMessageReminderRequest) ScheduledDeliveryTimestamp(
	scheduledDeliveryTimestamp time.Time,
) CreateMessageReminderRequest {
	timeStamp := scheduledDeliveryTimestamp.Unix()
	r.scheduledDeliveryTimestamp = &timeStamp
	return r
}

// A note associated with the reminder shown in the Notification Bot message.
//
// **Changes**: New in Zulip 11.0 (feature level 415).
func (r CreateMessageReminderRequest) Note(note string) CreateMessageReminderRequest {
	r.note = &note
	return r
}

func (r CreateMessageReminderRequest) Execute() (*CreateMessageReminderResponse, *http.Response, error) {
	return r.apiService.CreateMessageReminderExecute(r)
}

// CreateMessageReminder Create a message reminder
//
// Schedule a reminder to be sent to the current user at the specified time. The reminder will link the relevant message.
//
// *Changes**: New in Zulip 11.0 (feature level 381).
func (s *remindersService) CreateMessageReminder(ctx context.Context) CreateMessageReminderRequest {
	return CreateMessageReminderRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request.
func (s *remindersService) CreateMessageReminderExecute(
	r CreateMessageReminderRequest,
) (*CreateMessageReminderResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateMessageReminderResponse{}
		endpoint = "/reminders"
	)
	headers["Content-Type"] = apiutils.ContentTypeFormURLEncoded
	headers["Accept"] = apiutils.ContentTypeJSON

	apiutils.AddOptParam(form, "message_id", r.messageID)
	apiutils.AddOptParam(form, "scheduled_delivery_timestamp", r.scheduledDeliveryTimestamp)
	apiutils.AddOptParam(form, "note", r.note)
	req, err := apiutils.PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type DeleteReminderRequest struct {
	ctx        context.Context
	apiService APIReminders
	reminderID int64
}

func (r DeleteReminderRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.DeleteReminderExecute(r)
}

// DeleteReminder Delete a reminder
//
// Delete, and therefore cancel sending, a previously [scheduled reminder].
//
// *Changes**: New in Zulip 11.0 (feature level 399).
//
// [scheduled reminder]: https://zulip.com/help/schedule-a-reminder
func (s *remindersService) DeleteReminder(ctx context.Context, reminderID int64) DeleteReminderRequest {
	return DeleteReminderRequest{
		apiService: s,
		ctx:        ctx,
		reminderID: reminderID,
	}
}

// Execute executes the request.
func (s *remindersService) DeleteReminderExecute(r DeleteReminderRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/reminders/{reminder_id}"
	)

	path := strings.ReplaceAll(endpoint, "{reminder_id}", apiutils.IDToString(r.reminderID))

	headers["Accept"] = apiutils.ContentTypeJSON
	req, err := apiutils.PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetRemindersRequest struct {
	ctx        context.Context
	apiService APIReminders
}

func (r GetRemindersRequest) Execute() (*GetRemindersResponse, *http.Response, error) {
	return r.apiService.GetRemindersExecute(r)
}

// GetReminders Get reminders
//
// Fetch all [reminders] for the
// current user.
//
// Reminders are messages the user has scheduled to be sent in the
// future to themself.
//
// *Changes**: New in Zulip 11.0 (feature level 399).
//
// [reminders]: https://zulip.com/help/schedule-a-reminder
func (s *remindersService) GetReminders(ctx context.Context) GetRemindersRequest {
	return GetRemindersRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request.
func (s *remindersService) GetRemindersExecute(r GetRemindersRequest) (*GetRemindersResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetRemindersResponse{}
		endpoint = "/reminders"
	)

	headers["Accept"] = apiutils.ContentTypeJSON
	req, err := apiutils.PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}
