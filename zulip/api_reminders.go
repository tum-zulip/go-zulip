package zulip

import (
	"bytes"
	"context"
	"io"
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
	ApiService                 RemindersAPI
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
//	**Changes**: New in Zulip 11.0 (feature level 415).
func (r CreateMessageReminderRequest) Note(note string) CreateMessageReminderRequest {
	r.note = &note
	return r
}

func (r CreateMessageReminderRequest) Execute() (*CreateMessageReminderResponse, *http.Response, error) {
	return r.ApiService.CreateMessageReminderExecute(r)
}

// CreateMessageReminder Create a message reminder
//
// Schedule a reminder to be sent to the current user at the specified time. The reminder will link the relevant message.
//
// *Changes**: New in Zulip 11.0 (feature level 381).
func (c *simpleClient) CreateMessageReminder(ctx context.Context) CreateMessageReminderRequest {
	return CreateMessageReminderRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateMessageReminderExecute(r CreateMessageReminderRequest) (*CreateMessageReminderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateMessageReminderResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reminders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.messageId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message_id", r.messageId, "", "")
	}
	if r.scheduledDeliveryTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "scheduled_delivery_timestamp", r.scheduledDeliveryTimestamp, "", "")
	}
	if r.note != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "note", r.note, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteReminderRequest struct {
	ctx        context.Context
	ApiService RemindersAPI
	reminderId int64
}

func (r DeleteReminderRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeleteReminderExecute(r)
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
		ApiService: c,
		ctx:        ctx,
		reminderId: reminderId,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteReminderExecute(r DeleteReminderRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reminders/{reminder_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"reminder_id"+"}", url.PathEscape(parameterValueToString(r.reminderId, "reminderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetRemindersRequest struct {
	ctx        context.Context
	ApiService RemindersAPI
}

func (r GetRemindersRequest) Execute() (*GetRemindersResponse, *http.Response, error) {
	return r.ApiService.GetRemindersExecute(r)
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
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetRemindersExecute(r GetRemindersRequest) (*GetRemindersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetRemindersResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reminders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}
