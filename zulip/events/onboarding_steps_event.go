package events

// OnboardingStepsEvent Event sent when the set of onboarding steps to show for the current user has changed (e.g. because the user dismissed one).  Clients that feature a similar tutorial experience to the Zulip web app may want to handle these events.
//
// **Changes**: Before Zulip 8.0 (feature level 233), this event was named `hotspots`. Prior to this feature level, one-time notice onboarding steps were not supported.
type OnboardingStepsEvent struct {
	event

	// An array of dictionaries where each dictionary contains details about a single onboarding step.
	//
	// **Changes**: Before Zulip 8.0 (feature level 233), this array was named `hotspots`. Prior to this feature level, one-time notice onboarding steps were not supported, and the `type` field in these objects did not exist as all onboarding steps were implicitly hotspots.
	OnboardingSteps []OnboardingStep `json:"onboarding_steps,omitempty"`
}

// OnboardingStep Dictionary containing details of a single onboarding step.
type OnboardingStep struct {
	// The type of the onboarding step. Valid value is `"one_time_notice"`.
	//
	// **Changes**: Removed type `"hotspot"` in Zulip 9.0 (feature level 259).  New in Zulip 8.0 (feature level 233).
	Type string `json:"type,omitempty"`
	// The name of the onboarding step.
	Name string `json:"name,omitempty"`
}
