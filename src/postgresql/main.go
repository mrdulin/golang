package main

import (
	"database/sql"
	"log"
	"math/rand"

	"github.com/bxcodec/faker"
	_ "github.com/lib/pq"
)

type userPO struct {
	user_id         int
	user_nme        string
	user_email      string
	user_address_id int
}

func main() {
	driverName := "postgres"
	dataSourceName := "user=sampleadmin password=samplepass port=5431 dbname=nodejs-pg-knex-samples sslmode=disable"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	users, err := findUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", users)

	//userPO, err := insertUser(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("insert user %+v\n", userPO)
}

// Query
func findUsers(db *sql.DB) ([]userPO, error) {
	var userPOs []userPO

	sql := `select user_nme from users`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u userPO
		err := rows.Scan(&u.user_nme)
		if err != nil {
			return nil, err
		}
		userPOs = append(userPOs, u)
	}

	return userPOs, nil
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
