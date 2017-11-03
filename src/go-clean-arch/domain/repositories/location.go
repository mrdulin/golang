package repositories

import (
	"go-clean-arch/domain/models/cedar"
)

type LocationRepository interface {
	FindLocationsBoundGoogleClientCustomerId() ([]cedar.Location, error)
}
