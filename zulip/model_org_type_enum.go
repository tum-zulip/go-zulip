package zulip

// The [organization type] for the realm.
//   - 0 = Unspecified
//   - 10 = Business
//   - 20 = Open-source project
//   - 30 = Education (non-profit)
//   - 35 = Education (for-profit)
//   - 40 = Research
//   - 50 = Event or conference
//   - 60 = Non-profit (registered)
//   - 70 = Government
//   - 80 = Political group
//   - 90 = Community
//   - 100 = Personal
//   - 1000 = Other
//
// **Changes**: New in Zulip 6.0 (feature level 128).
//
// [organization type]: https://zulip.com/help/organization-type
type OrgType int

const (
	OrgTypeUnspecified         OrgType = 0
	OrgTypeBusiness            OrgType = 10
	OrgTypeOpenSource          OrgType = 20
	OrgTypeEducationNonProfit  OrgType = 30
	OrgTypeEducationForProfit  OrgType = 35
	OrgTypeResearch            OrgType = 40
	OrgTypeEventOrConference   OrgType = 50
	OrgTypeNonProfitRegistered OrgType = 60
	OrgTypeGovernment          OrgType = 70
	OrgTypePoliticalGroup      OrgType = 80
	OrgTypeCommunity           OrgType = 90
	OrgTypePersonal            OrgType = 100
	OrgTypeOther               OrgType = 1000
)

var AllowedOrgTypeEnumValues = []OrgType{
	OrgTypeUnspecified,
	OrgTypeBusiness,
	OrgTypeOpenSource,
	OrgTypeEducationNonProfit,
	OrgTypeEducationForProfit,
	OrgTypeResearch,
	OrgTypeEventOrConference,
	OrgTypeNonProfitRegistered,
	OrgTypeGovernment,
	OrgTypePoliticalGroup,
	OrgTypeCommunity,
	OrgTypePersonal,
	OrgTypeOther,
}

func NewOrgTypeFromValue(v int) (*OrgType, error) {
	ev := OrgType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedOrgTypeEnumValues,
			Value:   v,
			VarName: "OrgType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgType) IsValid() bool {
	for _, existing := range AllowedOrgTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
