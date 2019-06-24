package services

import (
	"fmt"
	"go-clean-arch/domain/models"
	"go-clean-arch/domain/repositories"
)

type IGoogleAccountService interface {
	FindGoogleAccountsForReport()([]models.GoogleAccount, error)
}

type GoogleAccountService struct {
	googleAccountRepo repositories.GoogleAccountRepository
	locationRepo repositories.LocationRepository
}

func NewGoogleAccountService(googleAccountRepo repositories.GoogleAccountRepository, locationRepo repositories.LocationRepository) IGoogleAccountService {
	return &GoogleAccountService{googleAccountRepo, locationRepo}
}

func (svc *GoogleAccountService) FindGoogleAccountsForReport() ([]models.GoogleAccount, error) {
	googleAccounts := make([]models.GoogleAccount, 0)
	locations, err := svc.locationRepo.FindLocationsBoundGoogleClientCustomerId()
	if err != nil {
		return googleAccounts, err
	}
	if len(locations) == 0 {
		return googleAccounts, fmt.Errorf("No location binds google adwords client customer ID.")
	}

	googleAdwordsClientCustomerIds := make([]int, 0)
	for _, location := range locations {
		if location.GoogleAdwordsClientCustomerId != 0 {
			googleAdwordsClientCustomerIds = append(googleAdwordsClientCustomerIds, location.GoogleAdwordsClientCustomerId)
		}
	}

	if len(googleAdwordsClientCustomerIds) == 0 {
		return googleAccounts, fmt.Errorf("No google adwords client customer ids")
	}

	googleAccountsForZOWI, err := svc.googleAccountRepo.FindByClientCustomerIds(googleAdwordsClientCustomerIds)
	if err != nil {
		return googleAccounts, err
	}
	googleAccounts = append(googleAccounts, googleAccountsForZOWI...)

	googleAccountsForZELO, err := svc.googleAccountRepo.FindByCampaignRanByZOWIForZELO()
	if err != nil {
		return googleAccounts, err
	}
	googleAccounts = append(googleAccounts, googleAccountsForZELO...)

	if len(googleAccounts) == 0 {
		return googleAccounts, fmt.Errorf("No google accounts for getting report")
	}

	return googleAccounts, nil
}