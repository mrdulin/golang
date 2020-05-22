package report

type CampaignStatus string

const (
	ENABLED CampaignStatus = "enabled"
	PAUSED  CampaignStatus = "paused"
)

type CampaignPerformanceReport struct {
	CampaignID     string
	CampaignNme    string
	CampaignStatus CampaignStatus
}
