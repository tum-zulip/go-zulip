package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type WebhooksAPI interface {

	/*
			ZulipOutgoingWebhooks Outgoing webhooks

			Outgoing webhooks allow you to build or set up Zulip integrations which are
		notified when certain types of messages are sent in Zulip.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ZulipOutgoingWebhooksRequest
	*/
	ZulipOutgoingWebhooks(ctx context.Context) ZulipOutgoingWebhooksRequest

	// ZulipOutgoingWebhooksExecute executes the request
	//  @return ZulipOutgoingWebhooksResponse
	ZulipOutgoingWebhooksExecute(r ZulipOutgoingWebhooksRequest) (*ZulipOutgoingWebhooksResponse, *http.Response, error)
}

type ZulipOutgoingWebhooksRequest struct {
	ctx        context.Context
	ApiService WebhooksAPI
}

func (r ZulipOutgoingWebhooksRequest) Execute() (*ZulipOutgoingWebhooksResponse, *http.Response, error) {
	return r.ApiService.ZulipOutgoingWebhooksExecute(r)
}

/*
ZulipOutgoingWebhooks Outgoing webhooks

Outgoing webhooks allow you to build or set up Zulip integrations which are
notified when certain types of messages are sent in Zulip.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ZulipOutgoingWebhooksRequest
*/
func (c *simpleClient) ZulipOutgoingWebhooks(ctx context.Context) ZulipOutgoingWebhooksRequest {
	return ZulipOutgoingWebhooksRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ZulipOutgoingWebhooksResponse
func (c *simpleClient) ZulipOutgoingWebhooksExecute(r ZulipOutgoingWebhooksRequest) (*ZulipOutgoingWebhooksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ZulipOutgoingWebhooksResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/zulip-outgoing-webhook"

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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
