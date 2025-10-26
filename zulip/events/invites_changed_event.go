package events

// InvitesChangedEvent A simple event sent when the set of invitations changes. This event is sent to organization administrators and the creator of the changed invitation; this tells clients they need to refetch data from `GET /invites` if they are displaying UI containing active invitations.
//
// **Changes**: Before Zulip 8.0 (feature level 209), this event was only sent to organization administrators.
type InvitesChangedEvent struct{ event }
