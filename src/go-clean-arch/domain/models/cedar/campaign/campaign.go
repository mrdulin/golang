package campaign

import "database/sql"

type CampaignChannelStatus string

const (
	SUCCESS     CampaignChannelStatus = "success"
	FAILED      CampaignChannelStatus = "failed"
	DISAPPROVED CampaignChannelStatus = "disapproved"
)

// Campaign Model
type Campaign struct {
	CampaignId            int            `db:"campaign_id" json:"campaignId"`
	CampaignNme           string         `db:"campaign_nme" json:"campaignNme"`
	OrganizationId        sql.NullInt64  `db:"organization_id" json:"organizationId"`
	LocationId            int            `db:"location_id"`
	CampaignChannelStatus sql.NullString `db:"campaign_channel_status"`

	GoogleChannelCampaignId int `db:"google_channel_campaign_id"`
}
