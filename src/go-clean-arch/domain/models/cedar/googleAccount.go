package cedar

import "time"

type GoogleAccount struct {
	GoogleAccountId                int       `db:"google_account_id"`
	GoogleAccountUserId            string    `db:"google_account_user_id"`
	GoogleAccountUserNme           string    `db:"google_account_user_nme"`
	GoogleAccountScope             string    `db:"google_account_scope"`
	GoogleAccountTokenType         string    `db:"google_account_token_type"`
	GoogleAccountUserLocale        string    `db:"google_account_user_locale"`
	GoogleAccountAccessToken       string    `db:"google_account_access_token"`
	GoogleAccountRefreshToken      string    `db:"google_account_refresh_token"`
	GoogleAccountIdToken           string    `db:"google_account_id_token"`
	GoogleAccountUserEmail         string    `db:"google_account_user_email"`
	GoogleAccountExpiryDate        time.Time `db:"google_account_expiry_date"`
	GoogleAccountConnectedInd      bool      `db:"google_account_connected_ind"`
	OrganizationId                 int       `db:"organization_id"`
	GoogleAccountDefaultCustomerId string    `db:"google_account_default_customer_id"`

	GoogleAdwordsId               int    `db:"google_adwords_id"`
	GoogleAdwordsCustomerNme      string `db:"google_adwords_customer_nme"`
	GoogleAdwordsClientCustomerId int    `db:"google_adwords_client_customer_id"`
}
