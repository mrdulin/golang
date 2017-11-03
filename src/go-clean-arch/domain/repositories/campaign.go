package repositories

import (
	"go-clean-arch/domain/models/cedar/campaign"
)

type CampaignRepository interface {
	FindById(id string) (*campaign.Campaign, error)
	FindValidGoogleCampaign() ([]*campaign.Campaign, error)
}
