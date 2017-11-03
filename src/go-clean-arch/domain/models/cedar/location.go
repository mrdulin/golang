package cedar

type Location struct {
	LocationId                    int `db:"location_id"`
	GoogleAdwordsClientCustomerId int `db:"google_adwords_client_customer_id"`
	CampaignId                    int `db:"campaign_id"`
}
