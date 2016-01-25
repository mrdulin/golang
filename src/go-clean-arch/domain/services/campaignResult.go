package services

import (
	"go-clean-arch/domain/repositories"
)

type ICampaignResultService interface{

}

type CampaignResultService struct {
	campaignResultRepo repositories.CampaignResultRepository
}

func NewCampaignResultService(campaignResultRepo repositories.CampaignResultRepository) ICampaignResultService {
	return &CampaignResultService{campaignResultRepo}
}

