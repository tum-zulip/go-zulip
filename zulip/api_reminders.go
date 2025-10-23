package zulip

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RemindersAPI interface {

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
	DeleteReminder(ctx context.Context, reminderId int64) DeleteReminderRequest

	// DeleteReminderExecute executes the request
	DeleteReminderExecute(r DeleteReminderRequest) (*Response, *http.Response, error)

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

type CreateMessageReminderRequest struct {
	ctx                        context.Context
	apiService                 RemindersAPI
	messageId                  *int64
	scheduledDeliveryTimestamp *int64
	note                       *string
}

// The Id of the previously sent message to reference in the reminder message.
func (r CreateMessageReminderRequest) MessageId(messageId int64) CreateMessageReminderRequest {
	r.messageId = &messageId
	return r
}

// The UNIX timestamp for when the reminder will be sent, in UTC seconds.
func (r CreateMessageReminderRequest) ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp time.Time) CreateMessageReminderRequest {
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
func (c *simpleClient) CreateMessageReminder(ctx context.Context) CreateMessageReminderRequest {
	return CreateMessageReminderRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateMessageReminderExecute(r CreateMessageReminderRequest) (*CreateMessageReminderResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateMessageReminderResponse{}
		endpoint = "/reminders"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addOptionalParam(form, "message_id", r.messageId)
	addOptionalParam(form, "scheduled_delivery_timestamp", r.scheduledDeliveryTimestamp)
	addOptionalParam(form, "note", r.note)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type DeleteReminderRequest struct {
	ctx        context.Context
	apiService RemindersAPI
	reminderId int64
}

func (r DeleteReminderRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.DeleteReminderExecute(r)
}

// DeleteReminder Delete a reminder
//
// Delete, and therefore cancel sending, a previously [scheduled reminder].
//
// *Changes**: New in Zulip 11.0 (feature level 399).
//
// [scheduled reminder]: https://zulip.com/help/schedule-a-reminder
func (c *simpleClient) DeleteReminder(ctx context.Context, reminderId int64) DeleteReminderRequest {
	return DeleteReminderRequest{
		apiService: c,
		ctx:        ctx,
		reminderId: reminderId,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteReminderExecute(r DeleteReminderRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/reminders/{reminder_id}"
	)

	endpoint = strings.Replace(endpoint, "{reminder_id}", idToString(r.reminderId), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetRemindersRequest struct {
	ctx        context.Context
	apiService RemindersAPI
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
func (c *simpleClient) GetReminders(ctx context.Context) GetRemindersRequest {
	return GetRemindersRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetRemindersExecute(r GetRemindersRequest) (*GetRemindersResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetRemindersResponse{}
		endpoint = "/reminders"
	)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}
