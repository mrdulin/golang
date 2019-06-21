package models

import "database/sql"

// Campaign Model
type Campaign struct {
	CampaignId int `db:"campaign_id" json:"campaignId"`
	CampaignNme string `db:"campaign_nme" json:"campaignNme"`
	OrganizationId sql.NullInt64 `db:"organization_id" json:"organizationId"`
	LocationId int `db:"location_id"`
}
