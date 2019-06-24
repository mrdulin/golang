package services

import (
	"fmt"
	"go-clean-arch/domain/models"
	"go-clean-arch/interfaces/repositories"
)

type IGoogleAccountService interface {
	FindGoogleAccountsForReport()([]models.GoogleAccount, error)
}

type GoogleAccountService struct {
	googleAccountRepo *repositories.GoogleAccountRepository
	locationRepo *repositories.LocationRepository
}

func NewGoogleAccountService(googleAccountRepo *repositories.GoogleAccountRepository, locationRepo *repositories.LocationRepository) IGoogleAccountService {
	return &GoogleAccountService{googleAccountRepo, locationRepo}
}

func (svc *GoogleAccountService) FindGoogleAccountsForReport() ([]models.GoogleAccount, error) {
	locations, err := svc.locationRepo.FindLocationsBoundGoogleClientCustomerId()
	if len(locations) == 0 {
		return nil, fmt.Errorf("No location binds google adwords client customer ID.")
	}

}