package zulip

// The configured [video call provider] for the organization.
//   - 0 = None
//   - 1 = Jitsi Meet
//   - 3 = Zoom (User OAuth integration)
//   - 4 = BigBlueButton
//   - 5 = Zoom (Server to Server OAuth integration)
//
// Note that only one of the [Zoom integrations] can be configured on a Zulip server.
// **Changes**: In Zulip 10.0 (feature level 353), added the Zoom Server to Server OAuth option.  In Zulip 3.0 (feature level 1), added the None option to disable video call UI.
//
// [Zoom integrations]: https://zulip.readthedocs.io/en/latest/production/video-calls.html#zoom
// [video call provider]: https://zulip.com/help/configure-call-provider
type VideoChatProvider int

const (
	VideoChatProviderNone                     VideoChatProvider = 0
	VideoChatProviderJitsiMeet                VideoChatProvider = 1
	VideoChatProviderZoomUserOAuthIntegration VideoChatProvider = 3
	VideoChatProviderBigBlueButton            VideoChatProvider = 4
	VideoChatProviderZoomServerToServerOAuth  VideoChatProvider = 5
)

var AllowedVideoChatProviderEnumValues = []VideoChatProvider{
	VideoChatProviderNone,
	VideoChatProviderJitsiMeet,
	VideoChatProviderZoomUserOAuthIntegration,
	VideoChatProviderBigBlueButton,
	VideoChatProviderZoomServerToServerOAuth,
}

func NewVideoChatProviderFromValue(v int) (*VideoChatProvider, error) {
	ev := VideoChatProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedVideoChatProviderEnumValues,
			Value:   v,
			VarName: "VideoChatProvider",
		}
	}
}
func (v VideoChatProvider) IsValid() bool {
	for _, existing := range AllowedVideoChatProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
