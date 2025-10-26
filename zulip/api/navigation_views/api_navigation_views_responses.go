package navigation_views

import (
	"github.com/tum-zulip/go-zulip/zulip"
)

// GetNavigationViewsResponse struct for GetNavigationViewsResponse
type GetNavigationViewsResponse struct {
	zulip.Response

	// An array of dictionaries containing the user's navigation views.
	NavigationViews []zulip.NavigationView `json:"navigation_views,omitempty"`
}
