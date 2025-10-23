package zulip

import (
	"context"
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
	apiService DraftsAPI
	drafts     *[]Draft
}

// A JSON-encoded list of containing new draft objects.
func (r CreateDraftsRequest) Drafts(drafts []Draft) CreateDraftsRequest {
	r.drafts = &drafts
	return r
}

func (r CreateDraftsRequest) Execute() (*CreateDraftsResponse, *http.Response, error) {
	return r.apiService.CreateDraftsExecute(r)
}

// CreateDrafts Create drafts
//
// Create one or more drafts on the server. These drafts will be automatically
// synchronized to other clients via `drafts` events.
func (c *simpleClient) CreateDrafts(ctx context.Context) CreateDraftsRequest {
	return CreateDraftsRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateDraftsExecute(r CreateDraftsRequest) (*CreateDraftsResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateDraftsResponse{}
		endpoint = "/drafts"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	if r.drafts != nil {
		for i := range *r.drafts {
			(*r.drafts)[i].Type = (*r.drafts)[i].Type.ToLegacy()
		}
		addParam(form, "drafts", r.drafts)
	}
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type CreateSavedSnippetRequest struct {
	ctx        context.Context
	apiService DraftsAPI
	title      *string
	content    *string
}

// The title of the saved snippet.
func (r CreateSavedSnippetRequest) Title(title string) CreateSavedSnippetRequest {
	r.title = &title
	return r
}

// The content of the saved snippet in [Zulip-flavored Markdown] format.  Clients should insert this content into a message when using a saved snippet.
//
// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
func (r CreateSavedSnippetRequest) Content(content string) CreateSavedSnippetRequest {
	r.content = &content
	return r
}

func (r CreateSavedSnippetRequest) Execute() (*CreateSavedSnippetResponse, *http.Response, error) {
	return r.apiService.CreateSavedSnippetExecute(r)
}

// CreateSavedSnippet Create a saved snippet
//
// Create a new saved snippet for the current user.
//
// *Changes**: New in Zulip 10.0 (feature level 297).
func (c *simpleClient) CreateSavedSnippet(ctx context.Context) CreateSavedSnippetRequest {
	return CreateSavedSnippetRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateSavedSnippetExecute(r CreateSavedSnippetRequest) (*CreateSavedSnippetResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateSavedSnippetResponse{}
		endpoint = "/saved_snippets"
	)
	if r.title == nil {
		return nil, nil, reportError("title is required and must be specified")
	}
	if r.content == nil {
		return nil, nil, reportError("content is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "title", r.title)
	addParam(form, "content", r.content)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type DeleteDraftRequest struct {
	ctx        context.Context
	apiService DraftsAPI
	draftId    int64
}

func (r DeleteDraftRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.DeleteDraftExecute(r)
}

// DeleteDraft Delete a draft
//
// Delete a single draft from the server. The deletion will be automatically
// synchronized to other clients via a `drafts` event.
func (c *simpleClient) DeleteDraft(ctx context.Context, draftId int64) DeleteDraftRequest {
	return DeleteDraftRequest{
		apiService: c,
		ctx:        ctx,
		draftId:    draftId,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteDraftExecute(r DeleteDraftRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/drafts/{draft_id}"
	)

	endpoint = strings.Replace(endpoint, "{draft_id}", idToString(r.draftId), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type DeleteSavedSnippetRequest struct {
	ctx            context.Context
	apiService     DraftsAPI
	savedSnippetId int64
}

func (r DeleteSavedSnippetRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.DeleteSavedSnippetExecute(r)
}

// DeleteSavedSnippet Delete a saved snippet
//
// Delete a saved snippet.
//
// *Changes**: New in Zulip 10.0 (feature level 297).
func (c *simpleClient) DeleteSavedSnippet(ctx context.Context, savedSnippetId int64) DeleteSavedSnippetRequest {
	return DeleteSavedSnippetRequest{
		apiService:     c,
		ctx:            ctx,
		savedSnippetId: savedSnippetId,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteSavedSnippetExecute(r DeleteSavedSnippetRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/saved_snippets/{saved_snippet_id}"
	)

	endpoint = strings.Replace(endpoint, "{saved_snippet_id}", idToString(r.savedSnippetId), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type EditDraftRequest struct {
	ctx        context.Context
	apiService DraftsAPI
	draftId    int64
	draft      *Draft
}

// A JSON-encoded object containing a replacement draft object for this Id.
func (r EditDraftRequest) Draft(draft Draft) EditDraftRequest {
	r.draft = &draft
	return r
}

func (r EditDraftRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.EditDraftExecute(r)
}

// EditDraft Edit a draft
//
// Edit a draft on the server. The edit will be automatically
// synchronized to other clients via `drafts` events.
func (c *simpleClient) EditDraft(ctx context.Context, draftId int64) EditDraftRequest {
	return EditDraftRequest{
		apiService: c,
		ctx:        ctx,
		draftId:    draftId,
	}
}

// Execute executes the request
func (c *simpleClient) EditDraftExecute(r EditDraftRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/drafts/{draft_id}"
	)

	endpoint = strings.Replace(endpoint, "{draft_id}", idToString(r.draftId), -1)

	if r.draft == nil {
		return nil, nil, reportError("draft is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	r.draft.Type = r.draft.Type.ToLegacy()
	addParam(form, "draft", r.draft)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type EditSavedSnippetRequest struct {
	ctx            context.Context
	apiService     DraftsAPI
	savedSnippetId int64
	title          *string
	content        *string
}

// The title of the saved snippet.
func (r EditSavedSnippetRequest) Title(title string) EditSavedSnippetRequest {
	r.title = &title
	return r
}

// The content of the saved snippet in the original [Zulip-flavored Markdown] format.  Clients should insert this content into a message when using a saved snippet.
//
// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
func (r EditSavedSnippetRequest) Content(content string) EditSavedSnippetRequest {
	r.content = &content
	return r
}

func (r EditSavedSnippetRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.EditSavedSnippetExecute(r)
}

// EditSavedSnippet Edit a saved snippet
//
// Edit a saved snippet for the current user.
//
// *Changes**: New in Zulip 10.0 (feature level 368).
func (c *simpleClient) EditSavedSnippet(ctx context.Context, savedSnippetId int64) EditSavedSnippetRequest {
	return EditSavedSnippetRequest{
		apiService:     c,
		ctx:            ctx,
		savedSnippetId: savedSnippetId,
	}
}

// Execute executes the request
func (c *simpleClient) EditSavedSnippetExecute(r EditSavedSnippetRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/saved_snippets/{saved_snippet_id}"
	)

	endpoint = strings.Replace(endpoint, "{saved_snippet_id}", idToString(r.savedSnippetId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addOptionalParam(form, "title", r.title)
	addOptionalParam(form, "content", r.content)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetDraftsRequest struct {
	ctx        context.Context
	apiService DraftsAPI
}

func (r GetDraftsRequest) Execute() (*GetDraftsResponse, *http.Response, error) {
	return r.apiService.GetDraftsExecute(r)
}

// GetDrafts Get drafts
//
// Fetch all drafts for the current user.
func (c *simpleClient) GetDrafts(ctx context.Context) GetDraftsRequest {
	return GetDraftsRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetDraftsExecute(r GetDraftsRequest) (*GetDraftsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetDraftsResponse{}
		endpoint = "/drafts"
	)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetSavedSnippetsRequest struct {
	ctx        context.Context
	apiService DraftsAPI
}

func (r GetSavedSnippetsRequest) Execute() (*GetSavedSnippetsResponse, *http.Response, error) {
	return r.apiService.GetSavedSnippetsExecute(r)
}

// GetSavedSnippets Get all saved snippets
//
// Fetch all the saved snippets for the current user.
//
// *Changes**: New in Zulip 10.0 (feature level 297).
func (c *simpleClient) GetSavedSnippets(ctx context.Context) GetSavedSnippetsRequest {
	return GetSavedSnippetsRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetSavedSnippetsExecute(r GetSavedSnippetsRequest) (*GetSavedSnippetsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetSavedSnippetsResponse{}
		endpoint = "/saved_snippets"
	)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}
