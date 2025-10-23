package models

import (
	"encoding/json"
	"fmt"

	. "github.com/tum-zulip/go-zulip/zulip/internal/utils"
)

// Todo: make not exported
type EventEnvelope struct {
	Event Event `json:"event"`
}

type eventPeeker struct {
	Type EventType `json:"type"`
	Id   int64     `json:"id"`
	Op   *EventOp  `json:"op,omitempty"`
}

func decodeAndWrap[T Event](event *EventEnvelope, data []byte) error {
	var t T
	err := NewStrictDecoder(data).Decode(&t)
	if err != nil {
		return err
	}

	event.Event = t
	return nil
}

func (e *EventEnvelope) UnmarshalJSON(data []byte) error {
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
		err = decodeAndWrap[AlertWordsEvent](e, data)
	case EventTypeAttachment:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[AttachmentAddEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[AttachmentRemoveEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeChannelFolder:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[ChannelFolderAddEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[ChannelFolderUpdateEvent](e, data)
		case EventOpReorder:
			err = decodeAndWrap[ChannelFolderReorderEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeCustomProfileFields:
		err = decodeAndWrap[CustomProfileFieldsEvent](e, data)
	case EventTypeDefaultChannelGroups:
		err = decodeAndWrap[DefaultChannelGroupsEvent](e, data)
	case EventTypeDefaultChannels:
		err = decodeAndWrap[DefaultChannelsEvent](e, data)
	case EventTypeDeleteMessage:
		err = decodeAndWrap[DeleteMessageEvent](e, data)
	case EventTypeDrafts:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[DraftsAddEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[DraftsUpdateEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[DraftsRemoveEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeHasZoomToken:
		err = decodeAndWrap[HasZoomTokenEvent](e, data)
	case EventTypeHeartbeat:
		err = decodeAndWrap[HeartbeatEvent](e, data)
	case EventTypeInvitesChanged:
		err = decodeAndWrap[InvitesChangedEvent](e, data)
	case EventTypeMessage:
		err = decodeAndWrap[MessageEvent](e, data)
	case EventTypeMutedTopics:
		err = decodeAndWrap[MutedTopicsEvent](e, data)
	case EventTypeMutedUsers:
		err = decodeAndWrap[MutedUsersEvent](e, data)
	case EventTypeNavigationView:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[NavigationViewAddEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[NavigationViewRemoveEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[NavigationViewUpdateEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeOnboardingSteps:
		err = decodeAndWrap[OnboardingStepsEvent](e, data)
	case EventTypePresence:
		err = decodeAndWrap[PresenceEvent](e, data)
	case EventTypePushDevice:
		err = decodeAndWrap[PushDeviceEvent](e, data)
	case EventTypeReaction:
		err = decodeAndWrap[ReactionEvent](e, data)
	case EventTypeRealm:
		switch op {
		case EventOpUpdate:
			err = decodeAndWrap[RealmUpdateEvent](e, data)
		case EventOpDeactivated:
			err = decodeAndWrap[RealmDeactivatedEvent](e, data)
		default:
			goto unknownEventError

		}
	case EventTypeRealmBot:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[RealmBotAddEvent](e, data)
		// RealmBotRemoveEvent Event sent to all users when a bot has been deactivated.
		// **Changes**: Deprecated and no longer sent since Zulip 8.0 (feature level 222).
		// Previously, this event was sent to all users in a Zulip organization when a bot was deactivated.
		case EventOpRemove:
			err = decodeAndWrap[RealmBotDeleteEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[RealmBotUpdateEvent](e, data)
		case EventOpDelete:
			err = decodeAndWrap[RealmBotDeleteEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeRealmDomains:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[RealmDomainsAddEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[RealmDomainsRemoveEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[RealmDomainsChangeEvent](e, data)
		case EventOpPeerAdd:
			err = decodeAndWrap[RealmDomainsAddEvent](e, data)
		case EventOpPeerRemove:
			err = decodeAndWrap[RealmDomainsRemoveEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeRealmEmoji:
		switch op {
		case EventOpUpdate:
			err = decodeAndWrap[RealmEmojiUpdateEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeRealmExport:
		err = decodeAndWrap[RealmExportEvent](e, data)
	case EventTypeRealmExportConsent:
		err = decodeAndWrap[RealmExportConsentEvent](e, data)
	case EventTypeRealmFilters:
		err = decodeAndWrap[RealmFiltersEvent](e, data)
	case EventTypeRealmLinkifiers:
		err = decodeAndWrap[RealmLinkifiersEvent](e, data)
	case EventTypeRealmPlaygrounds:
		err = decodeAndWrap[RealmPlaygroundsEvent](e, data)
	case EventTypeRealmUser:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[RealmUserAddEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[RealmUserRemoveEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[RealmUserUpdateEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeRealmUserSettingsDefaults:
		switch op {
		case EventOpUpdate:
			err = decodeAndWrap[RealmUserSettingsDefaultsUpdateEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeReminders:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[RemindersAddEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[RemindersRemoveEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeRestart:
		err = decodeAndWrap[RestartEvent](e, data)
	case EventTypeSavedSnippets:
		err = decodeAndWrap[SavedSnippetsEvent](e, data)
	case EventTypeScheduledMessages:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[ScheduledMessagesAddEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[ScheduledMessagesUpdateEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[ScheduledMessagesRemoveEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeChannel:
		switch op {
		case EventOpCreate:
			err = decodeAndWrap[ChannelCreateEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[ChannelUpdateEvent](e, data)
		case EventOpDelete:
			err = decodeAndWrap[ChannelDeleteEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeSubmessage:
		err = decodeAndWrap[SubmessageEvent](e, data)
	case EventTypeSubscription:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[SubscriptionAddEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[SubscriptionRemoveEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[SubscriptionUpdateEvent](e, data)
		case EventOpPeerAdd:
			err = decodeAndWrap[SubscriptionAddEvent](e, data)
		case EventOpPeerRemove:
			err = decodeAndWrap[SubscriptionRemoveEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeTyping:
		err = decodeAndWrap[TypingEvent](e, data)
	case EventTypeTypingEditMessage:
		err = decodeAndWrap[TypingEditMessageEvent](e, data)
	case EventTypeUpdateDisplaySettings:
		err = decodeAndWrap[UpdateDisplaySettingsEvent](e, data)
	case EventTypeUpdateGlobalNotifications:
		err = decodeAndWrap[UpdateGlobalNotificationsEvent](e, data)
	case EventTypeUpdateMessage:
		err = decodeAndWrap[UpdateMessageEvent](e, data)
	case EventTypeUpdateMessageFlags:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[UpdateMessageFlagsAddEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[UpdateMessageFlagsRemoveEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeUserGroup:
		switch op {
		case EventOpAdd:
			err = decodeAndWrap[UserGroupAddEvent](e, data)
		case EventOpRemove:
			err = decodeAndWrap[UserGroupRemoveEvent](e, data)
		case EventOpUpdate:
			err = decodeAndWrap[UserGroupUpdateEvent](e, data)
		case EventOpAddMembers, EventOpRemoveMembers:
			err = decodeAndWrap[UserGroupMembersEvent](e, data)
		case EventOpAddSubgroups, EventOpRemoveSubgroups:
			err = decodeAndWrap[UserGroupSubgroupsEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeUserSettings:
		switch op {
		case EventOpUpdate:
			err = decodeAndWrap[UserSettingsUpdateEvent](e, data)
		default:
			goto unknownEventError
		}
	case EventTypeUserStatus:
		err = decodeAndWrap[UserStatusEvent](e, data)
	case EventTypeUserTopic:
		err = decodeAndWrap[UserTopicEvent](e, data)
	case EventTypeWebReloadClient:
		err = decodeAndWrap[WebReloadClientEvent](e, data)
	default:
		goto unknownEventError
	}

	if err != nil {
		e.Event = &EventUnmarshalingError{
			EventCommon: EventCommon{
				Type: peeker.Type,
				Id:   peeker.Id,
			},
			Data: data,
			Err:  err,
		}
	}
	return nil

unknownEventError:
	e.Event = &EventUnmarshalingError{
		EventCommon: EventCommon{
			Type: peeker.Type,
			Id:   peeker.Id,
		},
		Data: data,
		Err:  fmt.Errorf("unknown event type %s with op %v", peeker.Type, peeker.Op),
	}
	return nil
}
