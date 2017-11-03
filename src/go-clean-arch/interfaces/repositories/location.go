package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-clean-arch/domain/models/cedar"
	"go-clean-arch/domain/repositories"
)

type LocationRepository struct {
	Db *sqlx.DB
}

func NewLocationRepository(Db *sqlx.DB) repositories.LocationRepository {
	return &LocationRepository{Db}
}

func (locationRepo *LocationRepository) FindLocationsBoundGoogleClientCustomerId() ([]cedar.Location, error) {
	var locations []cedar.Location
	query := `
		SELECT 
			DISTINCT ON
			(loc.google_adwords_client_customer_id) google_adwords_client_customer_id,
			loc.location_id,
			c.campaign_id
		FROM "LOCATION" as loc
		INNER JOIN "CAMPAIGN" AS c ON c.location_id = loc.location_id
		WHERE google_adwords_client_customer_id <> '';
	`

	err := locationRepo.Db.Select(&locations, query)
	if err != nil {
		return locations, err
	}
	return locations, nil
}
