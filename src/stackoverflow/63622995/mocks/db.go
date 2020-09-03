package mocks

import (
	"github.com/mrdulin/golang/src/stackoverflow/63622995/db"
	"github.com/stretchr/testify/mock"
)

type MockedOrmDB struct {
	mock.Mock
}

func (s *MockedOrmDB) Table(name string) db.OrmDB {
	args := s.Called(name)
	return args.Get(0).(db.OrmDB)
}

func (s *MockedOrmDB) Where(query interface{}, args ...interface{}) db.OrmDB {
	arguments := s.Called(query, args)
	return arguments.Get(0).(db.OrmDB)
}

func (s *MockedOrmDB) Scan(dest interface{}) *db.OrmDBWithError {
	args := s.Called(dest)
	return args.Get(0).(*db.OrmDBWithError)
}
