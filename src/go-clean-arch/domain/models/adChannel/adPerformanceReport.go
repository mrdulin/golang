package models

import (
	"go-clean-arch/domain/models/adChannel/ad"
	"go-clean-arch/domain/models/adChannel/adGroup"
	"go-clean-arch/domain/models/adChannel/campaign"
)

type AdPerformanceReportRow struct {
	Text           string                 `xml:",chardata"`
	AdGroupID      string                 `xml:"adGroupID,attr"`
	AdGroup        string                 `xml:"adGroup,attr"`
	AdGroupState   adGroup.AdGroupState   `xml:"adGroupState,attr"`
	AdType         string                 `xml:"adType,attr"`
	CampaignID     string                 `xml:"campaignID,attr"`
	Campaign       string                 `xml:"campaign,attr"`
	CampaignState  campaign.CampaignState `xml:"campaignState,attr"`
	ApprovalStatus ad.ApprovalStatus      `xml:"approvalStatus,attr"`
	AdID           string                 `xml:"adID,attr"`
	AdState        ad.AdState             `xml:"adState,attr"`
}

type AdPerformanceReport struct {
	BaseReport
	Table struct {
		Text string                   `xml:",chardata"`
		Row  []AdPerformanceReportRow `xml:"row"`
	} `xml:"table"`
}

func (report *AdPerformanceReport) GetRows() []AdPerformanceReportRow {
	return report.Table.Row
}
