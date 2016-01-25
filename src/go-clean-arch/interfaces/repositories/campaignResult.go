package repositories

import (
	"database/sql"
	"go-clean-arch/domain/repositories"
)

type CampaignResultRepository struct {
	Db *sql.DB
}

func NewCampaignResultRepository(Db *sql.DB) repositories.CampaignResultRepository {
	return &CampaignResultRepository{Db}
}

func (repo *CampaignResultRepository) UpdateStatusTransaction() error {

	return nil
}
