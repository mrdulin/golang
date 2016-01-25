package models

type DateRangeType string

const (
	TODAY        DateRangeType = "TODAY"
	YESTERDAY    DateRangeType = "YESTERDAY"
	LAST_7_DAYS  DateRangeType = "LAST_7_DAYS"
	LAST_30_DAYS DateRangeType = "LAST_30_DAYS"
	ALL_TIME     DateRangeType = "ALL_TIME"
)

type ReportDefinition struct {
	Selector      Selector      `json:"selector,omitempty"`
	ReportName    string        `json:"reportName,omitempty"`
	ReportType    string        `json:"reportType,omitempty"`
	DateRangeType DateRangeType `json:"reportType,omitempty"`
}
