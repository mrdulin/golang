package transaction

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
)
import (
	"github.com/icrowley/fake"
)

type User struct {
	UserID        int           `db:"user_id"`
	UserNme       string        `db:"user_nme"`
	UserEmail     string        `db:"user_email"`
	UserAddressId sql.NullInt64 `db:"user_address_id"`
}

type ITransactionSamples interface {
	CreateUserTransaction() (*User, error)
}

type TransactionSamples struct {
	Db *sqlx.DB
}

func NewTransactionSamples(Db *sqlx.DB) ITransactionSamples {
	return &TransactionSamples{Db}
}

func (ts *TransactionSamples) CreateUserTransaction() (*User, error) {
	tx := ts.Db.MustBegin()
	var lastInsertId int
	err := tx.QueryRowx(`INSERT INTO addresses (address_id, address_city, address_country, address_state) VALUES ($1, $2, $3, $4) RETURNING address_id`, 3, fake.City(), fake.Country(), fake.State()).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "insert address error")
	}
	log.Println("lastInsertId: ", lastInsertId)

	var user User
	err = tx.QueryRowx(`INSERT INTO users (user_id, user_nme, user_email, user_address_id) VALUES ($1, $2, $3, $4) RETURNING *;`, 6, fake.UserName(), fake.EmailAddress(), lastInsertId).StructScan(&user)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "insert user error")
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "tx.Commit()")
	}
	return &user, nil
}
