package api

type EventOp string

const (
	EventOpAdd             EventOp = "add"
	EventOpRemove          EventOp = "remove"
	EventOpUpdate          EventOp = "update"
	EventOpOnboardingSteps EventOp = "onboarding_steps"
	EventOpDeactivated     EventOp = "deactivated"
	EventOpCreate          EventOp = "create"
	EventOpDelete          EventOp = "delete"
	EventOpChange          EventOp = "change"
	EventOpReorder         EventOp = "reorder"
	EventOpStart           EventOp = "start"
	EventOpStop            EventOp = "stop"
	EventOpUpdateDict      EventOp = "update_dict"
	EventOpPeerAdd         EventOp = "peer_add"
	EventOpPeerRemove      EventOp = "peer_remove"
	EventOpAddMembers      EventOp = "add_members"
	EventOpAddSubgroups    EventOp = "add_subgroups"
	EventOpRemoveMembers   EventOp = "remove_members"
	EventOpRemoveSubgroups EventOp = "remove_subgroups"
)
