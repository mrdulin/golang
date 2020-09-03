package transaction

import (
	"log"
	"testing"

	"github.com/mrdulin/golang/src/sqlx-samples/database"
)

var (
	db                 database.IDataBase
	transactionSamples ITransactionSamples
)

func init() {
	db = database.NewDataBase()
	sqlxDb, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	transactionSamples = NewTransactionSamples(sqlxDb)
}

func TestCreateUserTransaction(t *testing.T) {
	user, err := transactionSamples.CreateUserTransaction()
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%#v", user)
}
