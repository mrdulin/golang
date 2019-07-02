package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type PGDatabaseConfig struct {
	Host, Port, User, Password, Dbname string
}

func ConnectPGDatabase(conf *PGDatabaseConfig) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.User, conf.Password, conf.Dbname)
	driverName := "postgres"
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Connect(driverName, dataSourceName)")
	}
	return db, nil
}
