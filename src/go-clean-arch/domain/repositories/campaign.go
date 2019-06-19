package repositories

import "go-clean-arch/domain/models"

type CampaignRepository interface {
	FindById(id string) (models.Campaign, error)
}