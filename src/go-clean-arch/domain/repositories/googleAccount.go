package repositories

import "go-clean-arch/domain/models"

type GoogleAccountRepository interface {
	FindByClientCustomerIds(ids []int) ([]models.GoogleAccount, error)
	FindByCampaignRanByZOWIForZELO() ([]models.GoogleAccount, error)
}