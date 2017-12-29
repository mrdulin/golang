package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func Connect() (*sqlx.DB, error) {
	conf := struct {
		SQL_INSTANCE_CONNECTION_NAME, SQL_USER, SQL_PASSWORD, SQL_DB string
	}{
		os.Getenv("SQL_INSTANCE_CONNECTION_NAME"),
		os.Getenv("SQL_USER"),
		os.Getenv("SQL_PASSWORD"),
		os.Getenv("SQL_DB"),
	}

	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s host=/cloudsql/%s sslmode=disable", conf.SQL_USER, conf.SQL_PASSWORD, conf.SQL_DB, conf.SQL_INSTANCE_CONNECTION_NAME)
	driverName := "postgres"
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	return db, nil
}
