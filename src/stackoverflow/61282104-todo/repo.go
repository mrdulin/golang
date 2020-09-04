package repo

import (
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
}

func (r *UserRepo) GetUser() (*sql.Rows, error) {
	sqlQuery := "select name from users;"
	rows, err := r.DB.Query(sqlQuery)
	return rows, err
}
