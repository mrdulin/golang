package no_results

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	UserID        int    `db:"user_id"`
	UserNme       string `db:"user_nme"`
	UserEmail     string `db:"user_email"`
	UserAddressId string `db:"user_address_id"`
}

func findUsersByIds(db *sqlx.DB, ids []int) ([]User, error) {
	var users []User

	query := `select * from users where user_id in (?);`

	query, args, err := sqlx.In(query, ids)
	log.Printf("query = %s, args = %v", query, args)
	if err != nil {
		return users, err
	}
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	if err != nil {
		return users, err
	}

	return users, nil
}
