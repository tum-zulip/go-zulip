package zulip

// Whether the data export is a public or a standard data export.
//   - ExportTypePublicData = Public data export.
//   - ExportTypeStandardData = Standard data export.
//
// **Changes**: New in Zulip 10.0 (feature level 304).
// Previously, the export type was not included in these objects because only public data exports could be created or listed via the API or UI.
type ExportType int

const (
	ExportTypePublicData   ExportType = 1
	ExportTypeStandardData ExportType = 2
)

var AllowedExportTypes = []ExportType{
	ExportTypePublicData,
	ExportTypeStandardData,
}

func ExportTypeFromValue(v int32) (*ExportType, error) {
	ev := ExportType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			VarName: "ExportType",
			Enum:    AllowedExportTypes,
			Value:   v,
		}
	}
}

func (v ExportType) IsValid() bool {
	for _, existing := range AllowedExportTypes {
		if existing == v {
			return true
		}
	}
	return false
}
