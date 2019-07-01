package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	cedarModels "go-clean-arch/domain/models/cedar/campaign"
	"go-clean-arch/domain/repositories"
)

type CampaignRepository struct {
	Db *sqlx.DB
}

func NewCampaignRepository(Db *sqlx.DB) repositories.CampaignRepository {
	return &CampaignRepository{Db}
}

func (cr *CampaignRepository) FindById(id string) (*cedarModels.Campaign, error) {
	campaign := cedarModels.Campaign{}
	query := `
		SELECT
			campaign_id,
			campaign_nme,
			organization_id
		FROM "CAMPAIGN" WHERE campaign_id = $1
	`
	err := cr.Db.Get(&campaign, query, id)
	if err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (cr *CampaignRepository) FindValidGoogleCampaign() ([]*cedarModels.Campaign, error) {
	var campaigns []*cedarModels.Campaign

	query := `
		select 
				c.campaign_id,
				gc.google_channel_campaign_id
		from "CAMPAIGN" as c
		inner join "LOCATION" as loc on c.location_id = loc.location_id
		inner join "GOOGLE_CHANNEL" as gc on c.campaign_channel_id = gc.campaign_channel_id
		where 
			c.campaign_channel_nme = 'google' and 
			c.campaign_duration_end_dte > CURRENT_DATE - 2 and
			loc.google_adwords_client_customer_id <> '' and 
			gc.google_channel_campaign_id <> ''
	`

	rows, err := cr.Db.Queryx(query)
	if err != nil {
		return nil, errors.Wrap(err, "cr.Db.Queryx")
	}
	for rows.Next() {
		var campaign cedarModels.Campaign
		err := rows.StructScan(&campaign)
		if err != nil {
			return nil, errors.Wrap(err, "rows.StructScan")
		}

		campaigns = append(campaigns, &campaign)
	}

	return campaigns, nil
}
