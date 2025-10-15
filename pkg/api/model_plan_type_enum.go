package api

// The plan type of the organization.
// - 1 = Self-hosted organization (SELF_HOSTED)
// - 2 = Zulip Cloud free plan (LIMITED)
// - 3 = Zulip Cloud Standard plan (STANDARD)
// - 4 = Zulip Cloud Standard plan, sponsored for free (STANDARD_FREE)
type PlanType int

const (
	PlanTypeSelfHosted   PlanType = 1
	PlanTypeLimited      PlanType = 2
	PlanTypeStandard     PlanType = 3
	PlanTypeStandardFree PlanType = 4
)

var AllowedPlanTypeEnumValues = []PlanType{
	PlanTypeSelfHosted,
	PlanTypeLimited,
	PlanTypeStandard,
	PlanTypeStandardFree,
}

func NewPlanTypeFromValue(v int) (*PlanType, error) {
	ev := PlanType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedPlanTypeEnumValues,
			Value:   v,
			VarName: "PlanType",
		}
	}
}

func (v PlanType) IsValid() bool {
	for _, existing := range AllowedPlanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
