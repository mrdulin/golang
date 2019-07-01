package application

import (
	"go-clean-arch/domain/services"
	"go-clean-arch/infrastructure/config"
	"go-clean-arch/infrastructure/database"
	"go-clean-arch/infrastructure/gcloud/datastore"
	"go-clean-arch/interfaces/repositories"
	"log"
	"os"
)

type CompositionRoot struct {
	AppConfig *config.ApplicationConfig

	CampaignService       services.ICampaignService
	CampaignResultService services.ICampaignResultService
	GoogleAccountService  services.IGoogleAccountService

	DataStoreService datastore.IDataStoreService
}

func NewCompositionRoot() *CompositionRoot {
	dataStoreOptions := datastore.Options{ProjectID: os.Getenv("GCP_PROJECT")}
	dataStoreService, err := datastore.New(&dataStoreOptions)

	appConfig, err := config.New("dataStore", dataStoreService)
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
		dataStoreService,
	}
}
