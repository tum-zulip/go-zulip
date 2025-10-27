// Package navigationviews provides API methods for managing Zulip navigation views,
// including home view settings and navigation state management.
package navigationviews

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/internal/apiutils"
	"github.com/tum-zulip/go-zulip/zulip/internal/clients"
)

type APINavigationViews interface {
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
	AddNavigationViewExecute(r AddNavigationViewRequest) (*zulip.Response, *http.Response, error)

	// EditNavigationView Update the navigation view
	//
	// Update the details of an existing configured navigation view,
	// such as its name or whether it's pinned.
	//
	// *Changes**: New in Zulip 11.0 (feature level 390).
	//
	EditNavigationView(ctx context.Context, fragment string) EditNavigationViewRequest

	// EditNavigationViewExecute executes the request
	EditNavigationViewExecute(r EditNavigationViewRequest) (*zulip.Response, *http.Response, error)

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
	RemoveNavigationViewExecute(r RemoveNavigationViewRequest) (*zulip.Response, *http.Response, error)
}

type navigationViewsService struct {
	client clients.Client
}

func NewNavigationViewsService(client clients.Client) APINavigationViews {
	return &navigationViewsService{client: client}
}

var _ APINavigationViews = (*navigationViewsService)(nil)

type AddNavigationViewRequest struct {
	ctx        context.Context
	apiService APINavigationViews
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

func (r AddNavigationViewRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.AddNavigationViewExecute(r)
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
func (s *navigationViewsService) AddNavigationView(ctx context.Context) AddNavigationViewRequest {
	return AddNavigationViewRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request.
func (s *navigationViewsService) AddNavigationViewExecute(
	r AddNavigationViewRequest,
) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/navigation_views"
	)
	if r.fragment == nil {
		return nil, nil, errors.New("fragment is required and must be specified")
	}
	if r.isPinned == nil {
		return nil, nil, errors.New("isPinned is required and must be specified")
	}

	headers["Content-Type"] = apiutils.ContentTypeFormURLEncoded
	headers["Accept"] = apiutils.ContentTypeJSON

	apiutils.AddParam(form, "fragment", r.fragment)
	apiutils.AddParam(form, "is_pinned", r.isPinned)
	apiutils.AddOptParam(form, "name", r.name)
	req, err := apiutils.PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type EditNavigationViewRequest struct {
	ctx        context.Context
	apiService APINavigationViews
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

func (r EditNavigationViewRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.EditNavigationViewExecute(r)
}

// EditNavigationView Update the navigation view
//
// Update the details of an existing configured navigation view,
// such as its name or whether it's pinned.
//
// *Changes**: New in Zulip 11.0 (feature level 390).
func (s *navigationViewsService) EditNavigationView(ctx context.Context, fragment string) EditNavigationViewRequest {
	return EditNavigationViewRequest{
		apiService: s,
		ctx:        ctx,
		fragment:   fragment,
	}
}

// Execute executes the request.
func (s *navigationViewsService) EditNavigationViewExecute(
	r EditNavigationViewRequest,
) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/navigation_views/{fragment}"
	)

	path := strings.ReplaceAll(endpoint, "{fragment}", url.PathEscape(r.fragment))

	headers["Content-Type"] = apiutils.ContentTypeFormURLEncoded
	headers["Accept"] = apiutils.ContentTypeJSON

	apiutils.AddOptParam(form, "is_pinned", r.isPinned)
	apiutils.AddOptParam(form, "name", r.name)
	req, err := apiutils.PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetNavigationViewsRequest struct {
	ctx        context.Context
	apiService APINavigationViews
}

func (r GetNavigationViewsRequest) Execute() (*GetNavigationViewsResponse, *http.Response, error) {
	return r.apiService.GetNavigationViewsExecute(r)
}

// GetNavigationViews Get all navigation views
//
// Fetch all configured custom navigation views for the current user.
//
// *Changes**: New in Zulip 11.0 (feature level 390).
func (s *navigationViewsService) GetNavigationViews(ctx context.Context) GetNavigationViewsRequest {
	return GetNavigationViewsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request.
func (s *navigationViewsService) GetNavigationViewsExecute(
	r GetNavigationViewsRequest,
) (*GetNavigationViewsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetNavigationViewsResponse{}
		endpoint = "/navigation_views"
	)

	headers["Accept"] = apiutils.ContentTypeJSON
	req, err := apiutils.PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type RemoveNavigationViewRequest struct {
	ctx        context.Context
	apiService APINavigationViews
	fragment   string
}

func (r RemoveNavigationViewRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.RemoveNavigationViewExecute(r)
}

// RemoveNavigationView Remove a navigation view
//
// Remove a navigation view.
//
// *Changes**: New in Zulip 11.0 (feature level 390).
func (s *navigationViewsService) RemoveNavigationView(
	ctx context.Context,
	fragment string,
) RemoveNavigationViewRequest {
	return RemoveNavigationViewRequest{
		apiService: s,
		ctx:        ctx,
		fragment:   fragment,
	}
}

// Execute executes the request.
func (s *navigationViewsService) RemoveNavigationViewExecute(
	r RemoveNavigationViewRequest,
) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/navigation_views/{fragment}"
	)

	path := strings.ReplaceAll(endpoint, "{fragment}", url.PathEscape(r.fragment))

	headers["Accept"] = apiutils.ContentTypeJSON
	req, err := apiutils.PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}
