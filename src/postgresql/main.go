package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/bxcodec/faker"
	_ "github.com/lib/pq"
)

func main() {
	driverName := "postgres"
	dataSourceName := "user=sampleadmin password=samplepass port=5431 dbname=nodejs-pg-knex-samples sslmode=disable"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	users := findUsers(db)
	log.Printf("%+v\n", users)

	userPO, err := insertUser(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("insert user %+v\n", userPO)
}

type userPO struct {
	user_id         int
	user_nme        string
	user_email      string
	user_address_id int
}

// Query
func findUsers(db *sql.DB) []userPO {
	sql := `
		select * from users
	`

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()

	userPOs := []userPO{}

	for rows.Next() {
		var u userPO
		err := rows.Scan(&u.user_id, &u.user_nme, &u.user_email, &u.user_address_id)
		if err != nil {
			log.Println(err.Error())
		}
		userPOs = append(userPOs, u)
	}

	return userPOs
}

func insertUser(db *sql.DB) (userPO, error) {
	sql := `
		insert into users(user_id, user_nme, user_email, user_address_id)
		values($1, $2, $3, $4)
		returning *;
	`

	var userPO userPO
	err := db.QueryRow(sql, rand.Intn(1000), faker.NAME, faker.Email, rand.Intn(1000)).Scan(&userPO.user_id, &userPO.user_nme, &userPO.user_email, &userPO.user_address_id)

	return userPO, err
}
