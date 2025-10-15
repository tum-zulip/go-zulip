package zulip

// GetNavigationViewsResponse struct for GetNavigationViewsResponse
type GetNavigationViewsResponse struct {
	Response

	// An array of dictionaries containing the user's navigation views.
	NavigationViews []NavigationView `json:"navigation_views,omitempty"`
}

// NavigationView Represents a user's personal configuration for a specific navigation view (displayed most visibly at the top of the web application's left sidebar).  Navigation views can be either an override to the default behavior of a built-in view, or a custom view.  **Changes**: New in Zulip 11.0 (feature level 390).
type NavigationView struct {
	// A unique identifier for the view, used to determine navigation behavior when clicked.  Clients should use this value to navigate to the corresponding URL hash.
	Fragment string `json:"fragment"`
	// Determines whether the view appears directly in the sidebar or is hidden in the \"More Views\" menu.  - `true` - Pinned and visible in the sidebar. - `false` - Hidden and accessible via the \"More Views\" menu.
	IsPinned bool `json:"is_pinned"`
	// The user-facing name for custom navigation views. Omit this field for built-in views.
	Name *string `json:"name,omitempty"`
}
