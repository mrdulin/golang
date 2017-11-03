package cedar

type GoogleChannel struct {
	GoogleChannelCampaignId string `db:"google_channel_campaign_id"`

	// TODO:
	GoogleChannelCampaignStatus string `db:"google_channel_campaign_status"`
	GoogleChannelAdGroupStatus  string `db:"google_channel_adgroup_status"`
	GoogleChannelAdStatus       string `db:"google_channel_ad_status"`
}
