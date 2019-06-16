package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"fmt"
)

type userPO struct {
	user_id int
	user_nme string
	user_email string
	user_address_id sql.NullString
	//user_address_id int
}

func main() {
	dialect := "postgres"
	db, err := gorm.Open(dialect, "host=127.0.0.1 port=5431 user=sampleadmin password=samplepass dbname=nodejs-pg-knex-samples sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	users := findUsersQueryBuilder(db)
	fmt.Printf("%#v", users)
}

func findUsersQueryBuilder (db *gorm.DB) []userPO {
	rows, err := db.Table("users").Select("user_id, user_nme, user_email, user_address_id").Rows()
	panicIf(err)
	defer rows.Close()

	userPOs := []userPO{}
	for rows.Next() {
		var userPO = userPO{}
		err := rows.Scan(&userPO.user_id, &userPO.user_nme, &userPO.user_email, &userPO.user_address_id)
		panicIf(err)
		userPOs = append(userPOs, userPO)
	}

	return userPOs
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}