package scheduled_messages

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	. "github.com/tum-zulip/go-zulip/zulip/internal/apiutils"
	. "github.com/tum-zulip/go-zulip/zulip/models"
)

type APIScheduledMessages interface {

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

type scheduledMessagesService struct {
	client StructuredClient
}

var _ APIScheduledMessages = (*scheduledMessagesService)(nil)

type CreateScheduledMessageRequest struct {
	ctx                        context.Context
	apiService                 APIScheduledMessages
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
// **Changes**: In Zulip 9.0 (feature level 248), `RecipientTypeChannel` was added as an additional value for this parameter to indicate the type of a channel message.
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
// **Changes**: Before Zulip 10.0 (feature level 370), `"(no topic)"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
//
// [`POST /register`]: https://zulip.com/api/register-queue
// [topics are required]: https://zulip.com/help/require-topics
func (r CreateScheduledMessageRequest) Topic(topic string) CreateScheduledMessageRequest {
	r.topic = &topic
	return r
}

// Whether the message should be initially marked read by its sender. If unspecified, the server uses a heuristic based on the client name and the recipient.
//
// **Changes**: New in Zulip 8.0 (feature level 236).
func (r CreateScheduledMessageRequest) ReadBySender(readBySender bool) CreateScheduledMessageRequest {
	r.readBySender = &readBySender
	return r
}

func (r CreateScheduledMessageRequest) Execute() (*CreateScheduledMessageResponse, *http.Response, error) {
	return r.apiService.CreateScheduledMessageExecute(r)
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
func (s *scheduledMessagesService) CreateScheduledMessage(ctx context.Context) CreateScheduledMessageRequest {
	return CreateScheduledMessageRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *scheduledMessagesService) CreateScheduledMessageExecute(r CreateScheduledMessageRequest) (*CreateScheduledMessageResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateScheduledMessageResponse{}
		endpoint = "/scheduled_messages"
	)
	if r.recipientType == nil {
		return nil, nil, fmt.Errorf("recipientType is required and must be specified")
	}
	if r.to == nil {
		return nil, nil, fmt.Errorf("to is required and must be specified")
	}
	if r.content == nil {
		return nil, nil, fmt.Errorf("content is required and must be specified")
	}
	if r.scheduledDeliveryTimestamp == nil {
		return nil, nil, fmt.Errorf("scheduledDeliveryTimestamp is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "type", r.recipientType)
	AddParam(form, "to", r.to)
	AddParam(form, "content", r.content)
	AddOptionalParam(form, "topic", r.topic)
	AddParam(form, "scheduled_delivery_timestamp", r.scheduledDeliveryTimestamp)
	AddOptionalParam(form, "read_by_sender", r.readBySender)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type DeleteScheduledMessageRequest struct {
	ctx                context.Context
	apiService         APIScheduledMessages
	scheduledMessageId int64
}

func (r DeleteScheduledMessageRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.DeleteScheduledMessageExecute(r)
}

// DeleteScheduledMessage Delete a scheduled message
//
// Delete, and therefore cancel sending, a previously [scheduled message].
//
// *Changes**: New in Zulip 7.0 (feature level 173).
//
// [scheduled message]: https://zulip.com/help/schedule-a-message
func (s *scheduledMessagesService) DeleteScheduledMessage(ctx context.Context, scheduledMessageId int64) DeleteScheduledMessageRequest {
	return DeleteScheduledMessageRequest{
		apiService:         s,
		ctx:                ctx,
		scheduledMessageId: scheduledMessageId,
	}
}

// Execute executes the request
func (s *scheduledMessagesService) DeleteScheduledMessageExecute(r DeleteScheduledMessageRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/scheduled_messages/{scheduled_message_id}"
	)

	endpoint = strings.Replace(endpoint, "{scheduled_message_id}", IdToString(r.scheduledMessageId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetScheduledMessagesRequest struct {
	ctx        context.Context
	apiService APIScheduledMessages
}

func (r GetScheduledMessagesRequest) Execute() (*GetScheduledMessagesResponse, *http.Response, error) {
	return r.apiService.GetScheduledMessagesExecute(r)
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
func (s *scheduledMessagesService) GetScheduledMessages(ctx context.Context) GetScheduledMessagesRequest {
	return GetScheduledMessagesRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *scheduledMessagesService) GetScheduledMessagesExecute(r GetScheduledMessagesRequest) (*GetScheduledMessagesResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetScheduledMessagesResponse{}
		endpoint = "/scheduled_messages"
	)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateScheduledMessageRequest struct {
	ctx                        context.Context
	apiService                 APIScheduledMessages
	scheduledMessageId         int64
	recipientType              *string
	to                         *Recipient
	content                    *string
	topic                      *string
	scheduledDeliveryTimestamp *int64
}

// The type of scheduled message to be sent. `"direct"` for a direct message and `"stream"` or `"channel"` for a channel message.  When updating the type of the scheduled message, the `to` parameter is required. And, if updating the type of the scheduled message to `"stream"`/`"channel"`, then the `topic` parameter is also required.  Note that, while `"private"` is supported for scheduling direct messages, clients are encouraged to use to the modern convention of `"direct"` to indicate this message type, because support for `"private"` may eventually be removed.
//
// **Changes**: In Zulip 9.0 (feature level 248), `"channel"` was added as an additional value for this parameter to indicate the type of a channel message.
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
// **Changes**: Before Zulip 10.0 (feature level 370), `"(no topic)"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
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
	return r.apiService.UpdateScheduledMessageExecute(r)
}

// UpdateScheduledMessage Edit a scheduled message
//
// Edit an existing [scheduled message].
//
// *Changes**: New in Zulip 7.0 (feature level 184).
//
// [scheduled message]: https://zulip.com/help/schedule-a-message
func (s *scheduledMessagesService) UpdateScheduledMessage(ctx context.Context, scheduledMessageId int64) UpdateScheduledMessageRequest {
	return UpdateScheduledMessageRequest{
		apiService:         s,
		ctx:                ctx,
		scheduledMessageId: scheduledMessageId,
	}
}

// Execute executes the request
func (s *scheduledMessagesService) UpdateScheduledMessageExecute(r UpdateScheduledMessageRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/scheduled_messages/{scheduled_message_id}"
	)

	endpoint = strings.Replace(endpoint, "{scheduled_message_id}", IdToString(r.scheduledMessageId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "type", r.recipientType)
	if err := AddOptionalJSONParam(form, "to", r.to); err != nil {
		return nil, nil, err
	}
	AddOptionalParam(form, "content", r.content)
	AddOptionalParam(form, "topic", r.topic)
	AddOptionalParam(form, "scheduled_delivery_timestamp", r.scheduledDeliveryTimestamp)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, req, response)
	return response, httpResp, err
}
