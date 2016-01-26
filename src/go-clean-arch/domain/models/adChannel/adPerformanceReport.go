package models

type AdPerformanceReportRow struct {
	Text           string `xml:",chardata"`
	AdGroupID      string `xml:"adGroupID,attr"`
	AdGroup        string `xml:"adGroup,attr"`
	AdGroupState   string `xml:"adGroupState,attr"`
	AdType         string `xml:"adType,attr"`
	CampaignID     string `xml:"campaignID,attr"`
	Campaign       string `xml:"campaign,attr"`
	CampaignState  string `xml:"campaignState,attr"`
	ApprovalStatus string `xml:"approvalStatus,attr"`
	AdID           string `xml:"adID,attr"`
	AdState        string `xml:"adState,attr"`
}

type AdPerformanceReport struct {
	BaseReport
	Table struct {
		Text string `xml:",chardata"`
		Row  []AdPerformanceReportRow `xml:"row"`
	} `xml:"table"`
}

func (report *AdPerformanceReport) GetRows() []AdPerformanceReportRow {
	return report.Table.Row
}
