package application

import (
	"go-clean-arch/domain/services"
	"go-clean-arch/infrastructure/config"
	"go-clean-arch/infrastructure/database"
	"go-clean-arch/interfaces/repositories"
	"log"
)

type CompositionRoot struct {
	AppConfig *config.ApplicationConfig

	CampaignService       services.ICampaignService
	CampaignResultService services.ICampaignResultService
	googleAccountService  services.IGoogleAccountService
}

func NewCompositionRoot() *CompositionRoot {
	appConfig, err := config.NewApplicationConfig()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	dbConf := database.PGDatabaseConfig{
		Host:     appConfig.SqlHost,
		Port:     appConfig.SqlPort,
		User:     appConfig.SqlUser,
		Password: appConfig.SqlPassword,
		Dbname:   appConfig.SqlDb,
	}
	db, err := database.ConnectPGDatabase(&dbConf)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// repositories
	campaignResultRepo := repositories.NewCampaignResultRepository(db)
	campaignRepo := repositories.NewCampaignRepository(db)
	googleAccountRepo := repositories.NewGoogleAccountRepository(db)
	locationRepo := repositories.NewLocationRepository(db)

	// services
	campaignService := services.NewCampaignService(campaignRepo)
	campaignResultService := services.NewCampaignResultService(campaignResultRepo)
	googleAccountService := services.NewGoogleAccountService(googleAccountRepo, locationRepo)

	return &CompositionRoot{
		appConfig,
		campaignService,
		campaignResultService,
		googleAccountService,
	}
}
