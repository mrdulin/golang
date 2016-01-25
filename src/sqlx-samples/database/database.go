package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type IDataBase interface {
	Connect() (*sqlx.DB, error)
}

type DataBase struct {}

func NewDataBase() IDataBase {
	return &DataBase{}
}

func (db *DataBase) Connect() (*sqlx.DB, error) {
	driverName := "postgres"
	dataSourceName := "host=127.0.0.1 port=5431 user=sampleadmin password=samplepass dbname=nodejs-pg-knex-samples sslmode=disable"
	sqlxDb, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return sqlxDb, nil
}
