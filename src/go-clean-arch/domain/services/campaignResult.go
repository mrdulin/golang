package services

import (
	models "go-clean-arch/domain/models/adChannel"
	"go-clean-arch/domain/repositories"
)

type ICampaignResultService interface {
	UpdateStatusTransaction(rows []models.AdPerformanceReportRow) error
}

type CampaignResultService struct {
	campaignResultRepo repositories.CampaignResultRepository
}

func NewCampaignResultService(campaignResultRepo repositories.CampaignResultRepository) ICampaignResultService {
	return &CampaignResultService{campaignResultRepo}
}

func (svc *CampaignResultService) UpdateStatusTransaction(rows []models.AdPerformanceReportRow) error {
	for _, row := range rows {
		return svc.campaignResultRepo.UpdateStatusTransaction(row)
	}
	return nil
}
