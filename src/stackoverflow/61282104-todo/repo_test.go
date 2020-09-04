package repo_test

import (
	"testing"

	repo "github.com/mrdulin/golang/src/stackoverflow/61282104-todo"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserRepo_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(`select name from users;`).WillReturnRows(sqlmock.NewRows(nil))
	userRepo := repo.UserRepo{DB: db}
	rows, err := userRepo.GetUser()
	t.Logf("rows: %+v\n", rows)
	t.Logf("err: %+v\n", err)
	if rows != nil {
		t.Fatalf("rows should be nil, got: %+v", rows)
	}
	if err != nil {
		t.Fatalf("err should be nil, got: %+v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
