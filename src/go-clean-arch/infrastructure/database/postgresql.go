package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPGDatabase() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 port=5431 user=sampleadmin password=samplepass dbname=nodejs-pg-knex-samples sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
