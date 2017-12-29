package services

import (
	"go-clean-arch/domain/models"
	"go-clean-arch/domain/repositories"
)

type ICampaignService interface {
	FindValidGoogleCampaignIds() ([]int, error)
}

type CampaignService struct {
	campaignRepo repositories.CampaignRepository
}

func NewCampaignService(campaignRepo repositories.CampaignRepository) ICampaignService {
	return &CampaignService{campaignRepo}
}

func (svc *CampaignService) FindValidGoogleCampaignIds() ([]int, error) {
	campaigns, err := svc.campaignRepo.FindValidGoogleCampaign()
	if err != nil {
		return nil, err
	}

	var googleChannelCampaignIds []int
	for _, campaign := range campaigns {
		googleChannelCampaignIds = append(googleChannelCampaignIds, campaign.GoogleChannelCampaignId)
	}

	if len(googleChannelCampaignIds) == 0 {
		return nil, &models.AppError{Err: "no valid campaign found", Code: models.BusinessErrorCode}
	}

	return googleChannelCampaignIds, nil
}
