package application

import (
	"fmt"
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

	DataStoreService datastore.IService
}

func NewCompositionRoot() *CompositionRoot {
	projectId := os.Getenv("GCP_PROJECT")
	fmt.Printf("projectId: %#v\n", projectId)
	dataStoreOptions := datastore.Options{ProjectID: projectId}
	dataStoreService, err := datastore.New(&dataStoreOptions)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	appConfig, err := config.New("dataStore", dataStoreService)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Printf("env: %#v\n", os.Getenv("ENV"))
	fmt.Printf("application config: %#v\n", appConfig)

	dbConf := database.PGDatabaseConfig{
		Host:     fmt.Sprintf("/cloudsql/%s", appConfig.SqlInstanceConnectionName),
		User:     appConfig.SqlUser,
		Password: appConfig.SqlPassword,
		Dbname:   appConfig.SqlDb,
	}
	db, err := database.ConnectPGDatabase(&dbConf)
	if err != nil {
		log.Fatalf("%+v\n", err)
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
