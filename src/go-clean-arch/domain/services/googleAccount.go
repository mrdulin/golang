package services

import (
	"fmt"
	"go-clean-arch/domain/models/cedar"
	"go-clean-arch/domain/repositories"
)

type IGoogleAccountService interface {
	FindGoogleAccountsForReport() ([]cedar.GoogleAccount, error)
}

type GoogleAccountService struct {
	googleAccountRepo repositories.GoogleAccountRepository
	locationRepo      repositories.LocationRepository
}

func NewGoogleAccountService(googleAccountRepo repositories.GoogleAccountRepository, locationRepo repositories.LocationRepository) IGoogleAccountService {
	return &GoogleAccountService{googleAccountRepo, locationRepo}
}

func (svc *GoogleAccountService) FindGoogleAccountsForReport() ([]cedar.GoogleAccount, error) {
	googleAccounts := make([]cedar.GoogleAccount, 0)
	locations, err := svc.locationRepo.FindLocationsBoundGoogleClientCustomerId()
	if err != nil {
		return googleAccounts, err
	}
	if len(locations) == 0 {
		return googleAccounts, fmt.Errorf("no location binds google adwords client customer id")
	}

	googleAdWordsClientCustomerIds := make([]int, 0)
	for _, location := range locations {
		if location.GoogleAdwordsClientCustomerId != 0 {
			googleAdWordsClientCustomerIds = append(googleAdWordsClientCustomerIds, location.GoogleAdwordsClientCustomerId)
		}
	}

	if len(googleAdWordsClientCustomerIds) == 0 {
		return googleAccounts, fmt.Errorf("no google adwords client customer ids")
	}

	googleAccountsForZOWI, err := svc.googleAccountRepo.FindByClientCustomerIds(googleAdWordsClientCustomerIds)
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
		return googleAccounts, fmt.Errorf("no google accounts for getting report")
	}

	return googleAccounts, nil
}
