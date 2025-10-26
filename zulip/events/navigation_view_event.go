package events

import "github.com/tum-zulip/go-zulip/zulip"

// NavigationViewAddEvent Event containing details of a newly configured navigation view.
//
// **Changes**: New in Zulip 11.0 (feature level 390).
type NavigationViewAddEvent struct {
	event

	NavigationView zulip.NavigationView `json:"navigation_view,omitempty"`
}

// NavigationViewUpdateEvent Event containing details of an update to an existing navigation view.
//
// **Changes**: New in Zulip 11.0 (feature level 390).
type NavigationViewUpdateEvent struct {
	event

	// The unique URL hash of the navigation view being updated.
	Fragment string                   `json:"fragment,omitempty"`
	Data     NavigationViewUpdateData `json:"data,omitempty"`
}

// NavigationViewUpdateData A dictionary containing the updated properties of the navigation view.
type NavigationViewUpdateData struct {
	// The user-facing name for custom navigation views. Omit this field for built-in views.
	Name *string `json:"name,omitempty"`
	// Determines whether the view is pinned (true) or hidden in the menu (false).
	IsPinned *bool `json:"is_pinned,omitempty"`
}

// NavigationViewRemoveEvent Event containing the fragment of a deleted navigation view.
//
// **Changes**: New in Zulip 11.0 (feature level 390).
type NavigationViewRemoveEvent struct {
	event

	// The unique URL hash of the navigation view that was just deleted.
	Fragment string `json:"fragment,omitempty"`
}
