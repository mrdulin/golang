package repositories

import "go-clean-arch/domain/models"

type LocationRepository interface {
	FindLocationsBoundGoogleClientCustomerId() ([]models.Location, error)
}
