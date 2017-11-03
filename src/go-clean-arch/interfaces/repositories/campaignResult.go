package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	models "go-clean-arch/domain/models/adChannel"
	cedarCampaign "go-clean-arch/domain/models/cedar/campaign"
	"go-clean-arch/domain/repositories"
	"log"
)

type CampaignResultRepository struct {
	Db *sqlx.DB
}

func NewCampaignResultRepository(Db *sqlx.DB) repositories.CampaignResultRepository {
	return &CampaignResultRepository{Db}
}

func (repo *CampaignResultRepository) UpdateStatusTransaction(row models.AdPerformanceReportRow, campaignChannelStatus cedarCampaign.CampaignChannelStatus) error {
	log.Printf("%+v, %+v", row, campaignChannelStatus)
	tx, err := repo.Db.Begin()
	if err != nil {
		return errors.Wrap(err, "repo.Db.Begin()")
	}

	updateGoogleChannelQuery := `
		update "GOOGLE_CHANNEL" 
		set 
			google_channel_campaign_status = $1,
			google_channel_adgroup_status = $2,
			google_channel_ad_status = $3,
			google_channel_approval_status = $4
		where 
			google_channel_campaign_id = $5;
	`
	if _, err := tx.Exec(updateGoogleChannelQuery, row.CampaignState, row.AdGroupState, row.AdState, row.ApprovalStatus, row.CampaignID); err != nil {
		return errors.Wrap(err, "tx.Exec(updateGoogleChannelQuery)")
	}

	updateCampaignQuery := `
		update "CAMPAIGN" as c
		set
			campaign_channel_status = $1
		from 
			"GOOGLE_CHANNEL" gc
		where
			gc.campaign_channel_id = c.campaign_channel_id and
			gc.google_channel_campaign_id = $2;
	`

	if _, err := tx.Exec(updateCampaignQuery, campaignChannelStatus, row.CampaignID); err != nil {
		return errors.Wrap(err, "tx.Exec(updateCampaignQuery)")
	}
	return tx.Commit()
}
