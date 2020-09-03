package main

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/mrdulin/golang/src/stackoverflow/63622995/db"
	"github.com/mrdulin/golang/src/stackoverflow/63622995/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderPersister_GetOrder(t *testing.T) {
	assert := assert.New(t)
	t.Run("should get order", func(t *testing.T) {
		testDb := new(mocks.MockedOrmDB)
		id := uuid.New()
		testDb.
			On("Table", "orders").
			Return(testDb).
			On("Where", "order_id = ?", mock.Anything).
			Return(testDb).
			On("Scan", mock.Anything).Run(func(args mock.Arguments) {
			ret := args.Get(0).(*Order)
			ret.order_id = "123"
		}).
			Return(&db.OrmDBWithError{Error: nil})
		op := OrderPersister{DB: testDb}
		got, err := op.GetOrder(id)
		testDb.AssertExpectations(t)
		assert.Nil(err)
		assert.Equal(Order{order_id: "123"}, *got)
	})

	t.Run("should return error", func(t *testing.T) {
		testDb := new(mocks.MockedOrmDB)
		id := uuid.New()
		testDb.
			On("Table", "orders").
			Return(testDb).
			On("Where", "order_id = ?", mock.Anything).
			Return(testDb).
			On("Scan", mock.Anything).
			Return(&db.OrmDBWithError{Error: errors.New("network")})
		op := OrderPersister{DB: testDb}
		got, err := op.GetOrder(id)
		testDb.AssertExpectations(t)
		assert.Equal(Order{}, *got)
		assert.Equal(err.Error(), "network")
	})
}
