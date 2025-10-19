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

type ScheduledMessagesAPI interface {

	// CreateScheduledMessage Create a scheduled message
	//
	// Create a new [scheduled message].
	//
	// *Changes**: In Zulip 7.0 (feature level 184), moved support for
	// [editing a scheduled message] to a
	// separate API endpoint, which removed the `scheduled_message_id`
	// parameter from this endpoint.
	//
	// New in Zulip 7.0 (feature level 179).
	//
	// [scheduled message]: https://zulip.com/help/schedule-a-message
	// [editing a scheduled message]: https://zulip.com/api/update-scheduled-message
	CreateScheduledMessage(ctx context.Context) CreateScheduledMessageRequest

	// CreateScheduledMessageExecute executes the request
	CreateScheduledMessageExecute(r CreateScheduledMessageRequest) (*CreateScheduledMessageResponse, *http.Response, error)

	// DeleteScheduledMessage Delete a scheduled message
	//
	// Delete, and therefore cancel sending, a previously [scheduled message].
	//
	// *Changes**: New in Zulip 7.0 (feature level 173).
	//
	// [scheduled message]: https://zulip.com/help/schedule-a-message
	DeleteScheduledMessage(ctx context.Context, scheduledMessageId int64) DeleteScheduledMessageRequest

	// DeleteScheduledMessageExecute executes the request
	DeleteScheduledMessageExecute(r DeleteScheduledMessageRequest) (*Response, *http.Response, error)

	// GetScheduledMessages Get scheduled messages
	//
	// Fetch all [scheduled messages] for
	// the current user.
	//
	// Scheduled messages are messages the user has scheduled to be
	// sent in the future via the send later feature.
	//
	// *Changes**: New in Zulip 7.0 (feature level 173).
	//
	// [scheduled messages]: https://zulip.com/help/schedule-a-message
	GetScheduledMessages(ctx context.Context) GetScheduledMessagesRequest

	// GetScheduledMessagesExecute executes the request
	GetScheduledMessagesExecute(r GetScheduledMessagesRequest) (*GetScheduledMessagesResponse, *http.Response, error)

	// UpdateScheduledMessage Edit a scheduled message
	//
	// Edit an existing [scheduled message].
	//
	// *Changes**: New in Zulip 7.0 (feature level 184).
	//
	// [scheduled message]: https://zulip.com/help/schedule-a-message
	UpdateScheduledMessage(ctx context.Context, scheduledMessageId int64) UpdateScheduledMessageRequest

	// UpdateScheduledMessageExecute executes the request
	UpdateScheduledMessageExecute(r UpdateScheduledMessageRequest) (*Response, *http.Response, error)
}

type CreateScheduledMessageRequest struct {
	ctx                        context.Context
	ApiService                 ScheduledMessagesAPI
	recipientType              *RecipientType
	to                         *Recipient
	content                    *string
	scheduledDeliveryTimestamp *int64
	topic                      *string
	readBySender               *bool
}

// The type of scheduled message to be sent. `RecipientTypeDirect` for a direct message and `RecipientTypeStream` or `RecipientTypeChannel` for a channel message.
//
// Note that, while `RecipientTypePrivate` is supported for scheduling direct messages, clients are encouraged to use to the modern convention of `RecipientTypeDirect` to indicate this message type, because support for `RecipientTypePrivate` may eventually be removed.
//
//	**Changes**: In Zulip 9.0 (feature level 248), `RecipientTypeChannel` was added as an additional value for this parameter to indicate the type of a channel message.
func (r CreateScheduledMessageRequest) RecipientType(recipientType RecipientType) CreateScheduledMessageRequest {
	r.recipientType = &recipientType
	return r
}

func (r CreateScheduledMessageRequest) To(to Recipient) CreateScheduledMessageRequest {
	r.to = &to
	return r
}

// The content of the message.  Clients should use the `max_message_length` returned by the [`POST /register`] endpoint to determine the maximum message size.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r CreateScheduledMessageRequest) Content(content string) CreateScheduledMessageRequest {
	r.content = &content
	return r
}

// The UNIX timestamp for when the message will be sent, in UTC seconds.
func (r CreateScheduledMessageRequest) ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp time.Time) CreateScheduledMessageRequest {
	timestamp := scheduledDeliveryTimestamp.Unix()
	r.scheduledDeliveryTimestamp = &timestamp
	return r
}

// The topic of the message. Only required for channel messages (`"type": "stream"` or `"type": "channel"`), ignored otherwise.  Clients should use the `max_topic_length` returned by the [`POST /register`] endpoint to determine the maximum topic length.  Note: When `"(no topic)"` or the value of `realm_empty_topic_display_name` found in the [POST /register] response is used for this parameter, it is interpreted as an empty string.  When [topics are required], this parameter can't be `"(no topic)"`, an empty string, or the value of `realm_empty_topic_display_name`.
//
//	**Changes**: Before Zulip 10.0 (feature level 370), `"(no topic)"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
//
// [`POST /register`]: https://zulip.com/api/register-queue
// [topics are required]: https://zulip.com/help/require-topics
func (r CreateScheduledMessageRequest) Topic(topic string) CreateScheduledMessageRequest {
	r.topic = &topic
	return r
}

// Whether the message should be initially marked read by its sender. If unspecified, the server uses a heuristic based on the client name and the recipient.
//
//	**Changes**: New in Zulip 8.0 (feature level 236).
func (r CreateScheduledMessageRequest) ReadBySender(readBySender bool) CreateScheduledMessageRequest {
	r.readBySender = &readBySender
	return r
}

func (r CreateScheduledMessageRequest) Execute() (*CreateScheduledMessageResponse, *http.Response, error) {
	return r.ApiService.CreateScheduledMessageExecute(r)
}

// CreateScheduledMessage Create a scheduled message
//
// Create a new [scheduled message].
//
// *Changes**: In Zulip 7.0 (feature level 184), moved support for
// [editing a scheduled message] to a
// separate API endpoint, which removed the `scheduled_message_id`
// parameter from this endpoint.
//
// New in Zulip 7.0 (feature level 179).
//
// [scheduled message]: https://zulip.com/help/schedule-a-message
// [editing a scheduled message]: https://zulip.com/api/update-scheduled-message
func (c *simpleClient) CreateScheduledMessage(ctx context.Context) CreateScheduledMessageRequest {
	return CreateScheduledMessageRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateScheduledMessageExecute(r CreateScheduledMessageRequest) (*CreateScheduledMessageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateScheduledMessageResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scheduled_messages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recipientType == nil {
		return localVarReturnValue, nil, reportError("recipientType is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.content == nil {
		return localVarReturnValue, nil, reportError("content is required and must be specified")
	}
	if r.scheduledDeliveryTimestamp == nil {
		return localVarReturnValue, nil, reportError("scheduledDeliveryTimestamp is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "type", r.recipientType, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "to", r.to, "form", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "content", r.content, "", "")
	if r.topic != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "topic", r.topic, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "scheduled_delivery_timestamp", r.scheduledDeliveryTimestamp, "form", "")
	if r.readBySender != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "read_by_sender", r.readBySender, "form", "")
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

type DeleteScheduledMessageRequest struct {
	ctx                context.Context
	ApiService         ScheduledMessagesAPI
	scheduledMessageId int64
}

func (r DeleteScheduledMessageRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeleteScheduledMessageExecute(r)
}

// DeleteScheduledMessage Delete a scheduled message
//
// Delete, and therefore cancel sending, a previously [scheduled message].
//
// *Changes**: New in Zulip 7.0 (feature level 173).
//
// [scheduled message]: https://zulip.com/help/schedule-a-message
func (c *simpleClient) DeleteScheduledMessage(ctx context.Context, scheduledMessageId int64) DeleteScheduledMessageRequest {
	return DeleteScheduledMessageRequest{
		ApiService:         c,
		ctx:                ctx,
		scheduledMessageId: scheduledMessageId,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteScheduledMessageExecute(r DeleteScheduledMessageRequest) (*Response, *http.Response, error) {
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

	localVarPath := localBasePath + "/scheduled_messages/{scheduled_message_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"scheduled_message_id"+"}", url.PathEscape(parameterValueToString(r.scheduledMessageId, "scheduledMessageId")), -1)

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

type GetScheduledMessagesRequest struct {
	ctx        context.Context
	ApiService ScheduledMessagesAPI
}

func (r GetScheduledMessagesRequest) Execute() (*GetScheduledMessagesResponse, *http.Response, error) {
	return r.ApiService.GetScheduledMessagesExecute(r)
}

// GetScheduledMessages Get scheduled messages
//
// Fetch all [scheduled messages] for
// the current user.
//
// Scheduled messages are messages the user has scheduled to be
// sent in the future via the send later feature.
//
// *Changes**: New in Zulip 7.0 (feature level 173).
//
// [scheduled messages]: https://zulip.com/help/schedule-a-message
func (c *simpleClient) GetScheduledMessages(ctx context.Context) GetScheduledMessagesRequest {
	return GetScheduledMessagesRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetScheduledMessagesExecute(r GetScheduledMessagesRequest) (*GetScheduledMessagesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetScheduledMessagesResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scheduled_messages"

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

type UpdateScheduledMessageRequest struct {
	ctx                        context.Context
	ApiService                 ScheduledMessagesAPI
	scheduledMessageId         int64
	recipientType              *string
	to                         *Recipient
	content                    *string
	topic                      *string
	scheduledDeliveryTimestamp *int64
}

// The type of scheduled message to be sent. `"direct"` for a direct message and `"stream"` or `"channel"` for a channel message.  When updating the type of the scheduled message, the `to` parameter is required. And, if updating the type of the scheduled message to `"stream"`/`"channel"`, then the `topic` parameter is also required.  Note that, while `"private"` is supported for scheduling direct messages, clients are encouraged to use to the modern convention of `"direct"` to indicate this message type, because support for `"private"` may eventually be removed.
//
//	**Changes**: In Zulip 9.0 (feature level 248), `"channel"` was added as an additional value for this parameter to indicate the type of a channel message.
func (r UpdateScheduledMessageRequest) RecipientType(recipientType string) UpdateScheduledMessageRequest {
	r.recipientType = &recipientType
	return r
}

func (r UpdateScheduledMessageRequest) To(to Recipient) UpdateScheduledMessageRequest {
	r.to = &to
	return r
}

// The updated content of the scheduled message.  Clients should use the `max_message_length` returned by the [`POST /register`] endpoint to determine the maximum message size.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r UpdateScheduledMessageRequest) Content(content string) UpdateScheduledMessageRequest {
	r.content = &content
	return r
}

// The updated topic of the scheduled message.  Required when updating the `type` of the scheduled message to `"stream"` or `"channel"`. Ignored when the existing or updated `type` of the scheduled message is `"direct"` (or `"private"`).  Clients should use the `max_topic_length` returned by the [`POST /register`] endpoint to determine the maximum topic length.  Note: When `"(no topic)"` or the value of `realm_empty_topic_display_name` found in the [POST /register] response is used for this parameter, it is interpreted as an empty string.  When [topics are required], this parameter can't be `"(no topic)"`, an empty string, or the value of `realm_empty_topic_display_name`.
//
//	**Changes**: Before Zulip 10.0 (feature level 370), `"(no topic)"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
//
// [`POST /register`]: https://zulip.com/api/register-queue
// [topics are required]: https://zulip.com/help/require-topics
func (r UpdateScheduledMessageRequest) Topic(topic string) UpdateScheduledMessageRequest {
	r.topic = &topic
	return r
}

// The UNIX timestamp for when the message will be sent, in UTC seconds.  Required when updating a scheduled message that the server has already tried and failed to send. This state is indicated with `"failed": true` in `scheduled_messages` objects; see response description at [`GET /scheduled_messages`].
//
// [`GET /scheduled_messages`]: https://zulip.com/api/get-scheduled-messages#response
func (r UpdateScheduledMessageRequest) ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp time.Time) UpdateScheduledMessageRequest {
	timestamp := scheduledDeliveryTimestamp.Unix()
	r.scheduledDeliveryTimestamp = &timestamp
	return r
}

func (r UpdateScheduledMessageRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateScheduledMessageExecute(r)
}

// UpdateScheduledMessage Edit a scheduled message
//
// Edit an existing [scheduled message].
//
// *Changes**: New in Zulip 7.0 (feature level 184).
//
// [scheduled message]: https://zulip.com/help/schedule-a-message
func (c *simpleClient) UpdateScheduledMessage(ctx context.Context, scheduledMessageId int64) UpdateScheduledMessageRequest {
	return UpdateScheduledMessageRequest{
		ApiService:         c,
		ctx:                ctx,
		scheduledMessageId: scheduledMessageId,
	}
}

// Execute executes the request
func (c *simpleClient) UpdateScheduledMessageExecute(r UpdateScheduledMessageRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scheduled_messages/{scheduled_message_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"scheduled_message_id"+"}", url.PathEscape(parameterValueToString(r.scheduledMessageId, "scheduledMessageId")), -1)

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
	if r.recipientType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "type", r.recipientType, "", "")
	}
	if r.to != nil {
		paramJson, err := parameterToJson(*r.to)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("to", paramJson)
	}
	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "content", r.content, "", "")
	}
	if r.topic != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "topic", r.topic, "", "")
	}
	if r.scheduledDeliveryTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "scheduled_delivery_timestamp", r.scheduledDeliveryTimestamp, "form", "")
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
