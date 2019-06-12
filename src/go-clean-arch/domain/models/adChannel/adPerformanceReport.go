package models

import "encoding/xml"

type AdPerformanceReport struct {
	XMLName    xml.Name `xml:"report"`
	Text       string   `xml:",chardata"`
	ReportName struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"report-name"`
	DateRange struct {
		Text string `xml:",chardata"`
		Date string `xml:"date,attr"`
	} `xml:"date-range"`
	Table struct {
		Text string `xml:",chardata"`
		Row  []struct {
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
		} `xml:"row"`
	} `xml:"table"`
}

