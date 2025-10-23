package navigation_views

import . "github.com/tum-zulip/go-zulip/zulip/models"

// GetNavigationViewsResponse struct for GetNavigationViewsResponse
type GetNavigationViewsResponse struct {
	Response

	// An array of dictionaries containing the user's navigation views.
	NavigationViews []NavigationView `json:"navigation_views,omitempty"`
}
