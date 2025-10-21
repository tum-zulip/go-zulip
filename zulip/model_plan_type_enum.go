package zulip

// The plan type of the organization.
//   - PlanTypeSelfHosted = Self-hosted organization (SELF_HOSTED)
//   - PlanTypeLimited = Zulip Cloud free plan (LIMITED)
//   - PlanTypeStandard = Zulip Cloud Standard plan (STANDARD)
//   - PlanTypeStandardFree = Zulip Cloud Standard plan, sponsored for free (STANDARD_FREE)
type PlanType int

const (
	PlanTypeSelfHosted   PlanType = 1
	PlanTypeLimited      PlanType = 2
	PlanTypeStandard     PlanType = 3
	PlanTypeStandardFree PlanType = 4
)

var allowedPlanTypeEnumValues = []PlanType{
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
			Enum:    allowedPlanTypeEnumValues,
			Value:   v,
			VarName: "PlanType",
		}
	}
}

func (v PlanType) IsValid() bool {
	for _, existing := range allowedPlanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
