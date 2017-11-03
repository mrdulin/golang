package campaign

type CampaignState string

const (
	ENABLED CampaignState = "enabled"
	PAUSED  CampaignState = "paused"
	REMOVED CampaignState = "removed"
)
