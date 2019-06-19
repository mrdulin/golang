package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-clean-arch/domain/models"
	"go-clean-arch/domain/repositories"
)

type CampaignRepository struct {
	Db *sqlx.DB
}

func NewCampaignRepository(Db *sqlx.DB) repositories.CampaignRepository {
	return &CampaignRepository{Db}
}

func (cr *CampaignRepository)FindById(id string) (models.Campaign, error) {
	campaign := models.Campaign{}
	query := `
		SELECT
			campaign_id,
			campaign_nme,
			organization_id
		FROM "CAMPAIGN" WHERE campaign_id = $1
	`
	err := cr.Db.Get(&campaign, query, id)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}