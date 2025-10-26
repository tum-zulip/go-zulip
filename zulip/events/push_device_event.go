package events

// PushDeviceEvent Event sent to a user's clients when the metadata in the `push_devices` dictionary for the user changes.  Helps clients to live-update the `push_devices` dictionary returned in [`POST /register`] response.
//
// **Changes**: New in Zulip 11.0 (feature level 406).
//
// [`POST /register`]: https://zulip.com/api/register-queue
type PushDeviceEvent struct {
	event
	// The push account Id for this client registration.  See [`POST /mobile_push/register`] for details on push account Ids.
	//
	// [`POST /mobile_push/register`]: https://zulip.com/api/register-push-device
	PushAccountId string `json:"push_account_id,omitempty"`
	// The updated registration status. Will be `"active"`, `"failed"`, or `"pending"`.
	Status string `json:"status,omitempty"`
	// If the status is `"failed"`, a [Zulip API error code] indicating the type of failure that occurred.  The following error codes have recommended client behavior:  - `"INVALId_BOUNCER_PUBLIC_KEY"` - Inform the user to update app. - `"REQUEST_EXPIRED` - Retry with a fresh payload.   If the status is "failed", an error code explaining the failure.
	//
	// [Zulip API error code]: https://zulip.com/api/rest-error-handling
	ErrorCode *string `json:"error_code,omitempty"`
}
