package zulip

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type NavigationViewsAPI interface {

	// AddNavigationView Add a navigation view
	//
	// Adds a new custom left sidebar navigation view configuration
	// for the current user.
	//
	// This can be used both to configure built-in navigation views,
	// or to add new navigation views.
	//
	// *Changes**: New in Zulip 11.0 (feature level 390).
	//
	AddNavigationView(ctx context.Context) AddNavigationViewRequest

	// AddNavigationViewExecute executes the request
	AddNavigationViewExecute(r AddNavigationViewRequest) (*Response, *http.Response, error)

	// EditNavigationView Update the navigation view
	//
	// Update the details of an existing configured navigation view,
	// such as its name or whether it's pinned.
	//
	// *Changes**: New in Zulip 11.0 (feature level 390).
	//
	EditNavigationView(ctx context.Context, fragment string) EditNavigationViewRequest

	// EditNavigationViewExecute executes the request
	EditNavigationViewExecute(r EditNavigationViewRequest) (*Response, *http.Response, error)

	// GetNavigationViews Get all navigation views
	//
	// Fetch all configured custom navigation views for the current user.
	//
	// *Changes**: New in Zulip 11.0 (feature level 390).
	//
	GetNavigationViews(ctx context.Context) GetNavigationViewsRequest

	// GetNavigationViewsExecute executes the request
	GetNavigationViewsExecute(r GetNavigationViewsRequest) (*GetNavigationViewsResponse, *http.Response, error)

	// RemoveNavigationView Remove a navigation view
	//
	// Remove a navigation view.
	//
	// *Changes**: New in Zulip 11.0 (feature level 390).
	//
	RemoveNavigationView(ctx context.Context, fragment string) RemoveNavigationViewRequest

	// RemoveNavigationViewExecute executes the request
	RemoveNavigationViewExecute(r RemoveNavigationViewRequest) (*Response, *http.Response, error)
}

type AddNavigationViewRequest struct {
	ctx        context.Context
	ApiService NavigationViewsAPI
	fragment   *string
	isPinned   *bool
	name       *string
}

// A unique identifier for the view, used to determine navigation behavior when clicked.  Clients should use this value to navigate to the corresponding URL hash.
func (r AddNavigationViewRequest) Fragment(fragment string) AddNavigationViewRequest {
	r.fragment = &fragment
	return r
}

// Determines whether the view appears directly in the sidebar or is hidden in the "More Views" menu.  - `true` - Pinned and visible in the sidebar. - `false` - Hidden and accessible via the "More Views" menu.
func (r AddNavigationViewRequest) IsPinned(isPinned bool) AddNavigationViewRequest {
	r.isPinned = &isPinned
	return r
}

// The user-facing name for custom navigation views. Omit this field for built-in views.
func (r AddNavigationViewRequest) Name(name string) AddNavigationViewRequest {
	r.name = &name
	return r
}

func (r AddNavigationViewRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.AddNavigationViewExecute(r)
}

// AddNavigationView Add a navigation view
//
// Adds a new custom left sidebar navigation view configuration
// for the current user.
//
// This can be used both to configure built-in navigation views,
// or to add new navigation views.
//
// *Changes**: New in Zulip 11.0 (feature level 390).
func (c *simpleClient) AddNavigationView(ctx context.Context) AddNavigationViewRequest {
	return AddNavigationViewRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) AddNavigationViewExecute(r AddNavigationViewRequest) (*Response, *http.Response, error) {
	var (
		httpMethod  = http.MethodPost
		postBody    interface{}
		headers     = make(map[string]string)
		queryParams = url.Values{}
		formParams  = url.Values{}
		response    = &Response{}
	)

	endpoint := "/navigation_views"

	if r.fragment == nil {
		return nil, nil, reportError("fragment is required and must be specified")
	}
	if r.isPinned == nil {
		return nil, nil, reportError("isPinned is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(formParams, "fragment", r.fragment, "", "")
	addParam(formParams, "is_pinned", r.isPinned, "", "")
	addOptionalParam(formParams, "name", r.name, "", "")
	req, err := c.prepareRequest(r.ctx, endpoint, httpMethod, postBody, headers, queryParams, formParams, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type EditNavigationViewRequest struct {
	ctx        context.Context
	ApiService NavigationViewsAPI
	fragment   string
	isPinned   *bool
	name       *string
}

// Determines whether the view is pinned (true) or hidden in the menu (false).
func (r EditNavigationViewRequest) IsPinned(isPinned bool) EditNavigationViewRequest {
	r.isPinned = &isPinned
	return r
}

// The user-facing name for custom navigation views. Omit this field for built-in views.
func (r EditNavigationViewRequest) Name(name string) EditNavigationViewRequest {
	r.name = &name
	return r
}

func (r EditNavigationViewRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.EditNavigationViewExecute(r)
}

// EditNavigationView Update the navigation view
//
// Update the details of an existing configured navigation view,
// such as its name or whether it's pinned.
//
// *Changes**: New in Zulip 11.0 (feature level 390).
func (c *simpleClient) EditNavigationView(ctx context.Context, fragment string) EditNavigationViewRequest {
	return EditNavigationViewRequest{
		ApiService: c,
		ctx:        ctx,
		fragment:   fragment,
	}
}

// Execute executes the request
func (c *simpleClient) EditNavigationViewExecute(r EditNavigationViewRequest) (*Response, *http.Response, error) {
	var (
		httpMethod  = http.MethodPatch
		postBody    interface{}
		headers     = make(map[string]string)
		queryParams = url.Values{}
		formParams  = url.Values{}
		response    = &Response{}
	)

	endpoint := "/navigation_views/{fragment}"
	endpoint = strings.Replace(endpoint, "{"+"fragment"+"}", url.PathEscape(r.fragment), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addOptionalParam(formParams, "is_pinned", r.isPinned, "", "")
	addOptionalParam(formParams, "name", r.name, "", "")
	req, err := c.prepareRequest(r.ctx, endpoint, httpMethod, postBody, headers, queryParams, formParams, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetNavigationViewsRequest struct {
	ctx        context.Context
	ApiService NavigationViewsAPI
}

func (r GetNavigationViewsRequest) Execute() (*GetNavigationViewsResponse, *http.Response, error) {
	return r.ApiService.GetNavigationViewsExecute(r)
}

// GetNavigationViews Get all navigation views
//
// Fetch all configured custom navigation views for the current user.
//
// *Changes**: New in Zulip 11.0 (feature level 390).
func (c *simpleClient) GetNavigationViews(ctx context.Context) GetNavigationViewsRequest {
	return GetNavigationViewsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetNavigationViewsExecute(r GetNavigationViewsRequest) (*GetNavigationViewsResponse, *http.Response, error) {
	var (
		httpMethod  = http.MethodGet
		postBody    interface{}
		headers     = make(map[string]string)
		queryParams = url.Values{}
		formParams  = url.Values{}
		response    = &GetNavigationViewsResponse{}
	)

	endpoint := "/navigation_views"

	// no Content-Type header

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, httpMethod, postBody, headers, queryParams, formParams, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RemoveNavigationViewRequest struct {
	ctx        context.Context
	ApiService NavigationViewsAPI
	fragment   string
}

func (r RemoveNavigationViewRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveNavigationViewExecute(r)
}

// RemoveNavigationView Remove a navigation view
//
// Remove a navigation view.
//
// *Changes**: New in Zulip 11.0 (feature level 390).
func (c *simpleClient) RemoveNavigationView(ctx context.Context, fragment string) RemoveNavigationViewRequest {
	return RemoveNavigationViewRequest{
		ApiService: c,
		ctx:        ctx,
		fragment:   fragment,
	}
}

// Execute executes the request
func (c *simpleClient) RemoveNavigationViewExecute(r RemoveNavigationViewRequest) (*Response, *http.Response, error) {
	var (
		httpMethod  = http.MethodDelete
		postBody    interface{}
		headers     = make(map[string]string)
		queryParams = url.Values{}
		formParams  = url.Values{}
		response    = &Response{}
	)

	endpoint := "/navigation_views/{fragment}"
	endpoint = strings.Replace(endpoint, "{"+"fragment"+"}", url.PathEscape(r.fragment), -1)

	// no Content-Type header

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, httpMethod, postBody, headers, queryParams, formParams, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}
