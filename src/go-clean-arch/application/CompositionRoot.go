package application

import (
	"go-clean-arch/domain/services"
	"go-clean-arch/infrastructure/database"
	"go-clean-arch/interfaces/repositories"
	"log"
	"os"
)

type CompositionRoot struct {
	CampaignService       services.ICampaignService
	CampaignResultService services.ICampaignResultService
}

func NewCompositionRoot() *CompositionRoot {
	dbConf := database.PGDatabaseConfig{
		Host:     os.Getenv("SQL_HOST"),
		Port:     os.Getenv("SQL_PORT"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Dbname:   os.Getenv("SQL_DB"),
	}
	db, err := database.ConnectPGDatabase(&dbConf)
	if err != nil {
		log.Fatal(err)
	}
	// repositories
	campaignResultRepo := repositories.NewCampaignResultRepository(db)
	campaignRepo := repositories.NewCampaignRepository(db)

	// services
	campaignService := services.NewCampaignService(campaignRepo)
	campaignResultService := services.NewCampaignResultService(campaignResultRepo)
	return &CompositionRoot{
		campaignService,
		campaignResultService,
	}
}
