package repositories

import (
	"go-clean-arch/domain/models/cedar"
)

type GoogleAccountRepository interface {
	FindByClientCustomerIds(ids []int) ([]cedar.GoogleAccount, error)
	FindByCampaignRanByZOWIForZELO() ([]cedar.GoogleAccount, error)
}
