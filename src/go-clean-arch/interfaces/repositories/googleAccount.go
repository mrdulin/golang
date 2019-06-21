package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-clean-arch/domain/models"
	"go-clean-arch/domain/repositories"
)

type GoogleAccountRepository struct {
	Db *sqlx.DB
}

func NewGoogleAccountRepository(Db *sqlx.DB) repositories.GoogleAccountRepository {
	return &GoogleAccountRepository{Db}
}

func (googleAccountRepo *GoogleAccountRepository) FindByClientCustomerIds(ids []int) ([]models.GoogleAccount, error) {
	query := `
		SELECT
			gac.google_account_refresh_token,
			gac.google_account_user_nme,
			gaw.google_adwords_id,
			gaw.google_adwords_customer_nme,
			gaw.google_adwords_client_customer_id
		FROM
			"GOOGLE_ADWORDS" as gaw
		INNER JOIN "GOOGLE_ACCOUNT" as gac ON gaw.google_account_id = gac.google_account_id
		WHERE 
			gaw.google_adwords_client_customer_id IN (${clientCustomerIds.map(() => '?').join(',')});
	`

	var googleAccounts []models.GoogleAccount
	err := googleAccountRepo.Db.Select(&googleAccounts, query, ids)
	if err != nil {
		return googleAccounts, err
	}
	return googleAccounts, nil
}
