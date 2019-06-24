package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-clean-arch/infrastructure/database"
	"go-clean-arch/interfaces/repositories"
	"go-clean-arch/domain/services"
	"log"
)

func init() {
	viper.SetConfigName("config.dev")
	viper.AddConfigPath("./infrastructure/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	dbConf := database.PGDatabaseConfig{
		Host: viper.GetString("SQL_HOST"),
		Port: viper.GetString("SQL_PORT"),
		User: viper.GetString("SQL_USER"),
		Password: viper.GetString("SQL_PASSWORD"),
		Dbname: viper.GetString("SQL_DB"),
	}
	db, err := database.ConnectPGDatabase(&dbConf)
	if err != nil {
		log.Fatal(err)
	}

	defer func () {
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
	googleAccountService := services.NewGoogleAccountService(googleAccountRepo, locationRepo)
	googleAccounts, err := googleAccountService.FindGoogleAccountsForReport()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", googleAccounts)

}
