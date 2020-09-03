package no_results

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/mrdulin/golang/src/sqlx-samples/database"
)

var (
	sqlxDb *sqlx.DB
	db     database.IDataBase
)

func init() {
	var err error
	db = database.NewDataBase()
	sqlxDb, err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

func TestFindUsersByIds_1(t *testing.T) {
	t.Skip()
	ids := []int{1, 2, 3}
	users, err := findUsersByIds(sqlxDb, ids)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("users = %+v", users)
}

func TestFindUsersByIds_2(t *testing.T) {
	ids := []int{10, 11, 100}
	users, err := findUsersByIds(sqlxDb, ids)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("users = %+v", users)
}
