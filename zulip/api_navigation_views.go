package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type NavigationViewsAPI interface {

	/*
			AddNavigationView Add a navigation view

			Adds a new custom left sidebar navigation view configuration
		for the current user.

		This can be used both to configure built-in navigation views,
		or to add new navigation views.

		**Changes**: New in Zulip 11.0 (feature level 390).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return AddNavigationViewRequest
	*/
	AddNavigationView(ctx context.Context) AddNavigationViewRequest

	// AddNavigationViewExecute executes the request
	//  @return Response
	AddNavigationViewExecute(r AddNavigationViewRequest) (*Response, *http.Response, error)

	/*
			EditNavigationView Update the navigation view

			Update the details of an existing configured navigation view,
		such as its name or whether it's pinned.

		**Changes**: New in Zulip 11.0 (feature level 390).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param fragment The unique URL hash of the navigation view to be updated.  This also serves as the identifier for the navigation view.
			@return EditNavigationViewRequest
	*/
	EditNavigationView(ctx context.Context, fragment string) EditNavigationViewRequest

	// EditNavigationViewExecute executes the request
	//  @return Response
	EditNavigationViewExecute(r EditNavigationViewRequest) (*Response, *http.Response, error)

	/*
			GetNavigationViews Get all navigation views

			Fetch all configured custom navigation views for the current user.

		**Changes**: New in Zulip 11.0 (feature level 390).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GetNavigationViewsRequest
	*/
	GetNavigationViews(ctx context.Context) GetNavigationViewsRequest

	// GetNavigationViewsExecute executes the request
	//  @return GetNavigationViewsResponse
	GetNavigationViewsExecute(r GetNavigationViewsRequest) (*GetNavigationViewsResponse, *http.Response, error)

	/*
			RemoveNavigationView Remove a navigation view

			Remove a navigation view.

		**Changes**: New in Zulip 11.0 (feature level 390).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param fragment The unique URL hash of the navigation view to be removed.  This also serves as the identifier for the navigation view.
			@return RemoveNavigationViewRequest
	*/
	RemoveNavigationView(ctx context.Context, fragment string) RemoveNavigationViewRequest

	// RemoveNavigationViewExecute executes the request
	//  @return Response
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

// Determines whether the view appears directly in the sidebar or is hidden in the \\\&quot;More Views\\\&quot; menu.  - &#x60;true&#x60; - Pinned and visible in the sidebar. - &#x60;false&#x60; - Hidden and accessible via the \\\&quot;More Views\\\&quot; menu.
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

/*
AddNavigationView Add a navigation view

Adds a new custom left sidebar navigation view configuration
for the current user.

This can be used both to configure built-in navigation views,
or to add new navigation views.

**Changes**: New in Zulip 11.0 (feature level 390).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AddNavigationViewRequest
*/
func (c *Client) AddNavigationView(ctx context.Context) AddNavigationViewRequest {
	return AddNavigationViewRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *Client) AddNavigationViewExecute(r AddNavigationViewRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/navigation_views"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fragment == nil {
		return localVarReturnValue, nil, reportError("fragment is required and must be specified")
	}
	if r.isPinned == nil {
		return localVarReturnValue, nil, reportError("isPinned is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "fragment", r.fragment, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "is_pinned", r.isPinned, "", "")
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

/*
EditNavigationView Update the navigation view

Update the details of an existing configured navigation view,
such as its name or whether it's pinned.

**Changes**: New in Zulip 11.0 (feature level 390).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fragment The unique URL hash of the navigation view to be updated.  This also serves as the identifier for the navigation view.
	@return EditNavigationViewRequest
*/
func (c *Client) EditNavigationView(ctx context.Context, fragment string) EditNavigationViewRequest {
	return EditNavigationViewRequest{
		ApiService: c,
		ctx:        ctx,
		fragment:   fragment,
	}
}

// Execute executes the request
//
//	@return Response
func (c *Client) EditNavigationViewExecute(r EditNavigationViewRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/navigation_views/{fragment}"
	localVarPath = strings.Replace(localVarPath, "{"+"fragment"+"}", url.PathEscape(parameterValueToString(r.fragment, "fragment")), -1)

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
	if r.isPinned != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_pinned", r.isPinned, "", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type GetNavigationViewsRequest struct {
	ctx        context.Context
	ApiService NavigationViewsAPI
}

func (r GetNavigationViewsRequest) Execute() (*GetNavigationViewsResponse, *http.Response, error) {
	return r.ApiService.GetNavigationViewsExecute(r)
}

/*
GetNavigationViews Get all navigation views

Fetch all configured custom navigation views for the current user.

**Changes**: New in Zulip 11.0 (feature level 390).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetNavigationViewsRequest
*/
func (c *Client) GetNavigationViews(ctx context.Context) GetNavigationViewsRequest {
	return GetNavigationViewsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetNavigationViewsResponse
func (c *Client) GetNavigationViewsExecute(r GetNavigationViewsRequest) (*GetNavigationViewsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetNavigationViewsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/navigation_views"

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

type RemoveNavigationViewRequest struct {
	ctx        context.Context
	ApiService NavigationViewsAPI
	fragment   string
}

func (r RemoveNavigationViewRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveNavigationViewExecute(r)
}

/*
RemoveNavigationView Remove a navigation view

Remove a navigation view.

**Changes**: New in Zulip 11.0 (feature level 390).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fragment The unique URL hash of the navigation view to be removed.  This also serves as the identifier for the navigation view.
	@return RemoveNavigationViewRequest
*/
func (c *Client) RemoveNavigationView(ctx context.Context, fragment string) RemoveNavigationViewRequest {
	return RemoveNavigationViewRequest{
		ApiService: c,
		ctx:        ctx,
		fragment:   fragment,
	}
}

// Execute executes the request
//
//	@return Response
func (c *Client) RemoveNavigationViewExecute(r RemoveNavigationViewRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/navigation_views/{fragment}"
	localVarPath = strings.Replace(localVarPath, "{"+"fragment"+"}", url.PathEscape(parameterValueToString(r.fragment, "fragment")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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
