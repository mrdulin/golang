package main

import (
	"fmt"
	"go-clean-arch/domain/services"
	"go-clean-arch/infrastructure/config"
	"go-clean-arch/infrastructure/database"
	"go-clean-arch/interfaces/repositories"
	"log"
)

func init() {

}

func main() {
	appConfig, err := config.NewApplicationConfig()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%+v\n", err)
		}
	}()

	//userRepository := repositories.NewUserRepository(db)
	//users, _ := userRepository.FindAll()
	//fmt.Printf("%#v", &users)

	//campaignRepository := repositories.NewCampaignRepository(db)
	//campaign, err := campaignRepository.FindById("102")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%#v\n", campaign)
	//campaignJSON, err := json.Marshal(campaign)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("campaignJSON:", string(campaignJSON))

	googleAccountRepo := repositories.NewGoogleAccountRepository(db)
	locationRepo := repositories.NewLocationRepository(db)
	campaignRepo := repositories.NewCampaignRepository(db)

	googleAccountService := services.NewGoogleAccountService(googleAccountRepo, locationRepo)
	campaignService := services.NewCampaignService(campaignRepo)
	_, err = googleAccountService.FindGoogleAccountsForReport()
	checkError(err)
	//log.Printf("%#v", googleAccounts)

	googleCampaignIds, err := campaignService.FindValidGoogleCampaignIds()
	checkError(err)
	log.Printf("%#v", googleCampaignIds)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
