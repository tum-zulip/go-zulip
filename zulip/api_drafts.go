package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type DraftsAPI interface {

	// CreateDrafts Create drafts
	//
	// Create one or more drafts on the server. These drafts will be automatically
	// synchronized to other clients via `drafts` events.
	//
	CreateDrafts(ctx context.Context) CreateDraftsRequest

	// CreateDraftsExecute executes the request
	CreateDraftsExecute(r CreateDraftsRequest) (*CreateDraftsResponse, *http.Response, error)

	// CreateSavedSnippet Create a saved snippet
	//
	// Create a new saved snippet for the current user.
	//
	// *Changes**: New in Zulip 10.0 (feature level 297).
	//
	CreateSavedSnippet(ctx context.Context) CreateSavedSnippetRequest

	// CreateSavedSnippetExecute executes the request
	CreateSavedSnippetExecute(r CreateSavedSnippetRequest) (*CreateSavedSnippetResponse, *http.Response, error)

	// DeleteDraft Delete a draft
	//
	// Delete a single draft from the server. The deletion will be automatically
	// synchronized to other clients via a `drafts` event.
	//
	DeleteDraft(ctx context.Context, draftId int64) DeleteDraftRequest

	// DeleteDraftExecute executes the request
	DeleteDraftExecute(r DeleteDraftRequest) (*Response, *http.Response, error)

	// DeleteSavedSnippet Delete a saved snippet
	//
	// Delete a saved snippet.
	//
	// *Changes**: New in Zulip 10.0 (feature level 297).
	//
	DeleteSavedSnippet(ctx context.Context, savedSnippetId int64) DeleteSavedSnippetRequest

	// DeleteSavedSnippetExecute executes the request
	DeleteSavedSnippetExecute(r DeleteSavedSnippetRequest) (*Response, *http.Response, error)

	// EditDraft Edit a draft
	//
	// Edit a draft on the server. The edit will be automatically
	// synchronized to other clients via `drafts` events.
	//
	EditDraft(ctx context.Context, draftId int64) EditDraftRequest

	// EditDraftExecute executes the request
	EditDraftExecute(r EditDraftRequest) (*Response, *http.Response, error)

	// EditSavedSnippet Edit a saved snippet
	//
	// Edit a saved snippet for the current user.
	//
	// *Changes**: New in Zulip 10.0 (feature level 368).
	//
	EditSavedSnippet(ctx context.Context, savedSnippetId int64) EditSavedSnippetRequest

	// EditSavedSnippetExecute executes the request
	EditSavedSnippetExecute(r EditSavedSnippetRequest) (*Response, *http.Response, error)

	// GetDrafts Get drafts
	//
	// Fetch all drafts for the current user.
	//
	GetDrafts(ctx context.Context) GetDraftsRequest

	// GetDraftsExecute executes the request
	GetDraftsExecute(r GetDraftsRequest) (*GetDraftsResponse, *http.Response, error)

	// GetSavedSnippets Get all saved snippets
	//
	// Fetch all the saved snippets for the current user.
	//
	// *Changes**: New in Zulip 10.0 (feature level 297).
	//
	GetSavedSnippets(ctx context.Context) GetSavedSnippetsRequest

	// GetSavedSnippetsExecute executes the request
	GetSavedSnippetsExecute(r GetSavedSnippetsRequest) (*GetSavedSnippetsResponse, *http.Response, error)
}

type CreateDraftsRequest struct {
	ctx        context.Context
	ApiService DraftsAPI
	drafts     *[]Draft
}

// A JSON-encoded list of containing new draft objects.
func (r CreateDraftsRequest) Drafts(drafts []Draft) CreateDraftsRequest {
	r.drafts = &drafts
	return r
}

func (r CreateDraftsRequest) Execute() (*CreateDraftsResponse, *http.Response, error) {
	return r.ApiService.CreateDraftsExecute(r)
}

// CreateDrafts Create drafts
//
// Create one or more drafts on the server. These drafts will be automatically
// synchronized to other clients via `drafts` events.
func (c *simpleClient) CreateDrafts(ctx context.Context) CreateDraftsRequest {
	return CreateDraftsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateDraftsExecute(r CreateDraftsRequest) (*CreateDraftsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateDraftsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drafts"

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
	if r.drafts != nil {
		for i := range *r.drafts {
			(*r.drafts)[i].Type = (*r.drafts)[i].Type.ToLegacy()
		}
		parameterAddToHeaderOrQuery(localVarFormParams, "drafts", r.drafts, "form", "multi")
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

type CreateSavedSnippetRequest struct {
	ctx        context.Context
	ApiService DraftsAPI
	title      *string
	content    *string
}

// The title of the saved snippet.
func (r CreateSavedSnippetRequest) Title(title string) CreateSavedSnippetRequest {
	r.title = &title
	return r
}

// The content of the saved snippet in [Zulip-flavored Markdown](https://zulip.com/help/format-your-message-using-markdown) format.  Clients should insert this content into a message when using a saved snippet.
func (r CreateSavedSnippetRequest) Content(content string) CreateSavedSnippetRequest {
	r.content = &content
	return r
}

func (r CreateSavedSnippetRequest) Execute() (*CreateSavedSnippetResponse, *http.Response, error) {
	return r.ApiService.CreateSavedSnippetExecute(r)
}

// CreateSavedSnippet Create a saved snippet
//
// Create a new saved snippet for the current user.
//
// *Changes**: New in Zulip 10.0 (feature level 297).
func (c *simpleClient) CreateSavedSnippet(ctx context.Context) CreateSavedSnippetRequest {
	return CreateSavedSnippetRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateSavedSnippetExecute(r CreateSavedSnippetRequest) (*CreateSavedSnippetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSavedSnippetResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/saved_snippets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.title == nil {
		return localVarReturnValue, nil, reportError("title is required and must be specified")
	}
	if r.content == nil {
		return localVarReturnValue, nil, reportError("content is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "title", r.title, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "content", r.content, "", "")
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

type DeleteDraftRequest struct {
	ctx        context.Context
	ApiService DraftsAPI
	draftId    int64
}

func (r DeleteDraftRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeleteDraftExecute(r)
}

// DeleteDraft Delete a draft
//
// Delete a single draft from the server. The deletion will be automatically
// synchronized to other clients via a `drafts` event.
func (c *simpleClient) DeleteDraft(ctx context.Context, draftId int64) DeleteDraftRequest {
	return DeleteDraftRequest{
		ApiService: c,
		ctx:        ctx,
		draftId:    draftId,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteDraftExecute(r DeleteDraftRequest) (*Response, *http.Response, error) {
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

	localVarPath := localBasePath + "/drafts/{draft_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"draft_id"+"}", url.PathEscape(parameterValueToString(r.draftId, "draftId")), -1)

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

type DeleteSavedSnippetRequest struct {
	ctx            context.Context
	ApiService     DraftsAPI
	savedSnippetId int64
}

func (r DeleteSavedSnippetRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeleteSavedSnippetExecute(r)
}

// DeleteSavedSnippet Delete a saved snippet
//
// Delete a saved snippet.
//
// *Changes**: New in Zulip 10.0 (feature level 297).
func (c *simpleClient) DeleteSavedSnippet(ctx context.Context, savedSnippetId int64) DeleteSavedSnippetRequest {
	return DeleteSavedSnippetRequest{
		ApiService:     c,
		ctx:            ctx,
		savedSnippetId: savedSnippetId,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteSavedSnippetExecute(r DeleteSavedSnippetRequest) (*Response, *http.Response, error) {
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

	localVarPath := localBasePath + "/saved_snippets/{saved_snippet_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"saved_snippet_id"+"}", url.PathEscape(parameterValueToString(r.savedSnippetId, "savedSnippetId")), -1)

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

type EditDraftRequest struct {
	ctx        context.Context
	ApiService DraftsAPI
	draftId    int64
	draft      *Draft
}

// A JSON-encoded object containing a replacement draft object for this Id.
func (r EditDraftRequest) Draft(draft Draft) EditDraftRequest {
	r.draft = &draft
	return r
}

func (r EditDraftRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.EditDraftExecute(r)
}

// EditDraft Edit a draft
//
// Edit a draft on the server. The edit will be automatically
// synchronized to other clients via `drafts` events.
func (c *simpleClient) EditDraft(ctx context.Context, draftId int64) EditDraftRequest {
	return EditDraftRequest{
		ApiService: c,
		ctx:        ctx,
		draftId:    draftId,
	}
}

// Execute executes the request
func (c *simpleClient) EditDraftExecute(r EditDraftRequest) (*Response, *http.Response, error) {
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

	localVarPath := localBasePath + "/drafts/{draft_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"draft_id"+"}", url.PathEscape(parameterValueToString(r.draftId, "draftId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.draft == nil {
		return localVarReturnValue, nil, reportError("draft is required and must be specified")
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
	r.draft.Type = r.draft.Type.ToLegacy()
	parameterAddToHeaderOrQuery(localVarFormParams, "draft", r.draft, "form", "")
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

type EditSavedSnippetRequest struct {
	ctx            context.Context
	ApiService     DraftsAPI
	savedSnippetId int64
	title          *string
	content        *string
}

// The title of the saved snippet.
func (r EditSavedSnippetRequest) Title(title string) EditSavedSnippetRequest {
	r.title = &title
	return r
}

// The content of the saved snippet in the original [Zulip-flavored Markdown](https://zulip.com/help/format-your-message-using-markdown) format.  Clients should insert this content into a message when using a saved snippet.
func (r EditSavedSnippetRequest) Content(content string) EditSavedSnippetRequest {
	r.content = &content
	return r
}

func (r EditSavedSnippetRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.EditSavedSnippetExecute(r)
}

// EditSavedSnippet Edit a saved snippet
//
// Edit a saved snippet for the current user.
//
// *Changes**: New in Zulip 10.0 (feature level 368).
func (c *simpleClient) EditSavedSnippet(ctx context.Context, savedSnippetId int64) EditSavedSnippetRequest {
	return EditSavedSnippetRequest{
		ApiService:     c,
		ctx:            ctx,
		savedSnippetId: savedSnippetId,
	}
}

// Execute executes the request
func (c *simpleClient) EditSavedSnippetExecute(r EditSavedSnippetRequest) (*Response, *http.Response, error) {
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

	localVarPath := localBasePath + "/saved_snippets/{saved_snippet_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"saved_snippet_id"+"}", url.PathEscape(parameterValueToString(r.savedSnippetId, "savedSnippetId")), -1)

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
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "title", r.title, "", "")
	}
	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "content", r.content, "", "")
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

type GetDraftsRequest struct {
	ctx        context.Context
	ApiService DraftsAPI
}

func (r GetDraftsRequest) Execute() (*GetDraftsResponse, *http.Response, error) {
	return r.ApiService.GetDraftsExecute(r)
}

// GetDrafts Get drafts
//
// Fetch all drafts for the current user.
func (c *simpleClient) GetDrafts(ctx context.Context) GetDraftsRequest {
	return GetDraftsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetDraftsExecute(r GetDraftsRequest) (*GetDraftsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDraftsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drafts"

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

type GetSavedSnippetsRequest struct {
	ctx        context.Context
	ApiService DraftsAPI
}

func (r GetSavedSnippetsRequest) Execute() (*GetSavedSnippetsResponse, *http.Response, error) {
	return r.ApiService.GetSavedSnippetsExecute(r)
}

// GetSavedSnippets Get all saved snippets
//
// Fetch all the saved snippets for the current user.
//
// *Changes**: New in Zulip 10.0 (feature level 297).
func (c *simpleClient) GetSavedSnippets(ctx context.Context) GetSavedSnippetsRequest {
	return GetSavedSnippetsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetSavedSnippetsExecute(r GetSavedSnippetsRequest) (*GetSavedSnippetsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSavedSnippetsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/saved_snippets"

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
