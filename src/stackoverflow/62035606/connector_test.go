package connector_test

import (
	"fmt"
	"testing"

	connector "github.com/mrdulin/golang/src/stackoverflow/62035606"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockedDB struct {
	mock.Mock
}

func (db *MockedDB) Ping() error {
	args := db.Called()
	return args.Error(0)
}

func TestConnector_Pool(t *testing.T) {
	t.Run("should verifies connection to the database is still alive", func(t *testing.T) {
		testDB := new(MockedDB)
		c := connector.Connector{DB: testDB}
		testDB.On("Ping").Return(nil)
		pool, err := c.Pool()
		require.Nil(t, err, nil)
		require.Equal(t, 1, pool)
	})

	t.Run("should return error if connection to the database is not alive", func(t *testing.T) {
		testDB := new(MockedDB)
		c := connector.Connector{DB: testDB}
		testDB.On("Ping").Return(fmt.Errorf("network"))
		pool, err := c.Pool()
		require.Error(t, err, "network")
		require.Nil(t, pool)
	})
}
