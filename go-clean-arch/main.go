package main

import (
	"log"
	"fmt"
	models "github.com/mrdulin/golang/go-clean-arch/domain/models/user"
	"github.com/mrdulin/golang/go-clean-arch/infrastructure/database"

)

func main() {
	db := database.ConnectPGDatabase()

	users := []models.User{}
	query := "select * from users"
	err := db.Select(&users, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", users)
}
