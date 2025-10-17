package zulip

import (
	"encoding/json"
)

// EventEnvelope - struct for EventEnvelope
type EventEnvelope struct {
	Event Event
}

type eventPeeker struct {
	Type EventType `json:"type"`
	Op   *EventOp  `json:"op,omitempty"`
}

func decodeAndWrap[T Event](envelope *EventEnvelope, data []byte) error {
	var t T
	err := newStrictDecoder(data).Decode(&t)
	if err != nil {
		return err
	}
	envelope.Event = t
	return nil
}

func (src *EventEnvelope) UnmarshalJSON(data []byte) error {
	var peeker eventPeeker
	if err := json.Unmarshal(data, &peeker); err != nil {
		return err
	}
	var err error
	var op EventOp
	if peeker.Op != nil {
		op = *peeker.Op
	}
	switch peeker.Type {
	case EventTypeAlertWords:
		err = decodeAndWrap[AlertWordsEvent](src, data)
	case EventTypeAttachment:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[AttachmentAddEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[AttachmentRemoveEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeChannelFolder:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[ChannelFolderAddEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[ChannelFolderUpdateEvent](src, data)
		case EventOpReorder:
			err = decodeAndWrap[ChannelFolderReorderEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeCustomProfileFields:
		err = decodeAndWrap[CustomProfileFieldsEvent](src, data)
	case EventTypeDefaultChannelGroups:
		err = decodeAndWrap[DefaultChannelGroupsEvent](src, data)
	case EventTypeDefaultChannels:
		err = decodeAndWrap[DefaultChannelsEvent](src, data)
	case EventTypeDeleteMessage:
		err = decodeAndWrap[DeleteMessageEvent](src, data)
	case EventTypeDrafts:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[DraftsAddEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[DraftsUpdateEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[DraftsRemoveEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeHasZoomToken:
		err = decodeAndWrap[HasZoomTokenEvent](src, data)
	case EventTypeHeartbeat:
		err = decodeAndWrap[HeartbeatEvent](src, data)
	case EventTypeInvitesChanged:
		err = decodeAndWrap[InvitesChangedEvent](src, data)
	case EventTypeMessage:
		err = decodeAndWrap[MessageEvent](src, data)
	case EventTypeMutedTopics:
		err = decodeAndWrap[MutedTopicsEvent](src, data)
	case EventTypeMutedUsers:
		err = decodeAndWrap[MutedUsersEvent](src, data)
	case EventTypeNavigationView:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[NavigationViewAddEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[NavigationViewRemoveEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[NavigationViewUpdateEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeOnboardingSteps:
		err = decodeAndWrap[OnboardingStepsEvent](src, data)
	case EventTypePresence:
		err = decodeAndWrap[PresenceEvent](src, data)
	case EventTypePushDevice:
		err = decodeAndWrap[PushDeviceEvent](src, data)
	case EventTypeReaction:
		err = decodeAndWrap[ReactionEvent](src, data)
	case EventTypeRealm:
		switch op {
		case EventOpUpdate:
			err = decodeAndWrap[RealmUpdateEvent](src, data)
		case EventOpDeactivated:
			err = decodeAndWrap[RealmDeactivatedEvent](src, data)
		default:
			goto unknownEventError

		}
	case EventTypeRealmBot:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[RealmBotAddEvent](src, data)
		// RealmBotRemoveEvent Event sent to all users when a bot has been deactivated.
		// **Changes**: Deprecated and no longer sent since Zulip 8.0 (feature level 222).
		// Previously, this event was sent to all users in a Zulip organization when a bot was deactivated.
		case EventOpRemove:
			err = decodeAndWrap[RealmBotDeleteEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[RealmBotUpdateEvent](src, data)
		case EventOpDelete:
			err = decodeAndWrap[RealmBotDeleteEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeRealmDomains:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[RealmDomainsAddEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[RealmDomainsRemoveEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[RealmDomainsChangeEvent](src, data)
		case EventOpPeerAdd:
			err = decodeAndWrap[RealmDomainsAddEvent](src, data)
		case EventOpPeerRemove:
			err = decodeAndWrap[RealmDomainsRemoveEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeRealmEmoji:
		switch op {
		case EventOpUpdate:
			err = decodeAndWrap[RealmEmojiUpdateEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeRealmExport:
		err = decodeAndWrap[RealmExportEvent](src, data)
	case EventTypeRealmExportConsent:
		err = decodeAndWrap[RealmExportConsentEvent](src, data)
	case EventTypeRealmFilters:
		err = decodeAndWrap[RealmFiltersEvent](src, data)
	case EventTypeRealmLinkifiers:
		err = decodeAndWrap[RealmLinkifiersEvent](src, data)
	case EventTypeRealmPlaygrounds:
		err = decodeAndWrap[RealmPlaygroundsEvent](src, data)
	case EventTypeRealmUser:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[RealmUserAddEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[RealmUserRemoveEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[RealmUserUpdateEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeRealmUserSettingsDefaults:
		switch op {
		case EventOpUpdate:
			err = decodeAndWrap[RealmUserSettingsDefaultsUpdateEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeReminders:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[RemindersAddEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[RemindersRemoveEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeRestart:
		err = decodeAndWrap[RestartEvent](src, data)
	case EventTypeSavedSnippets:
		err = decodeAndWrap[SavedSnippetsEvent](src, data)
	case EventTypeScheduledMessages:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[ScheduledMessagesAddEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[ScheduledMessagesUpdateEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[ScheduledMessagesRemoveEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeChannel:
		switch op {
		case EventOpCreate:
			err = decodeAndWrap[ChannelCreateEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[ChannelUpdateEvent](src, data)
		case EventOpDelete:
			err = decodeAndWrap[ChannelDeleteEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeSubmessage:
		err = decodeAndWrap[SubmessageEvent](src, data)
	case EventTypeSubscription:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[SubscriptionAddEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[SubscriptionRemoveEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[SubscriptionUpdateEvent](src, data)
		case EventOpPeerAdd:
			err = decodeAndWrap[SubscriptionAddEvent](src, data)
		case EventOpPeerRemove:
			err = decodeAndWrap[SubscriptionRemoveEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeTyping:
		err = decodeAndWrap[TypingEvent](src, data)
	case EventTypeTypingEditMessage:
		err = decodeAndWrap[TypingEditMessageEvent](src, data)
	case EventTypeUpdateDisplaySettings:
		err = decodeAndWrap[UpdateDisplaySettingsEvent](src, data)
	case EventTypeUpdateGlobalNotifications:
		err = decodeAndWrap[UpdateGlobalNotificationsEvent](src, data)
	case EventTypeUpdateMessage:
		err = decodeAndWrap[UpdateMessageEvent](src, data)
	case EventTypeUpdateMessageFlags:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[UpdateMessageFlagsAddEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[UpdateMessageFlagsRemoveEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeUserGroup:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[UserGroupAddEvent](src, data)
		case EventOpRemove:
			err = decodeAndWrap[UserGroupRemoveEvent](src, data)
		case EventOpUpdate:
			err = decodeAndWrap[UserGroupUpdateEvent](src, data)
		case EventOpAddMembers, EventOpRemoveMembers:
			err = decodeAndWrap[UserGroupMembersEvent](src, data)
		case EventOpAddSubgroups, EventOpRemoveSubgroups:
			err = decodeAndWrap[UserGroupSubgroupsEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeUserSettings:
		switch op {
		case EventOpUpdate:
			err = decodeAndWrap[UserSettingsUpdateEvent](src, data)
		default:
			goto unknownEventError
		}
	case EventTypeUserStatus:
		err = decodeAndWrap[UserStatusEvent](src, data)
	case EventTypeUserTopic:
		err = decodeAndWrap[UserTopicEvent](src, data)
	case EventTypeWebReloadClient:
		err = decodeAndWrap[WebReloadClientEvent](src, data)
	default:
		err = &json.UnmarshalTypeError{
			Value:  "EventEnvelope",
			Struct: "EventEnvelope",
			Field:  "type",
		}
		goto unknownEventError
	}

	if err != nil {
		return err
	}
	return nil

unknownEventError:
	return &json.UnmarshalTypeError{
		Value:  "EventEnvelope",
		Struct: "EventEnvelope",
		Field:  "type",
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EventEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(src.Event)
}
