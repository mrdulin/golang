package models

import "database/sql"

// User Model
type User struct {
	UserID        int           `db:"user_id"`
	UserNme       string        `db:"user_nme"`
	UserEmail     string        `db:"user_email"`
	UserAddressID sql.NullInt64 `db:"user_address_id"`
}
