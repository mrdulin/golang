package config_test

import (
	"reflect"
	"testing"

	config "github.com/mrdulin/golang/src/stackoverflow/62351902"
)

func TestGetConfigName(t *testing.T) {
	t.Run("should return test1", func(t *testing.T) {
		oGetAccountNumber := config.GetAccountNumber
		var toBeCallTimes int
		config.GetAccountNumber = func() *string {
			toBeCallTimes++
			n := "123"
			return &n
		}
		defer func() { config.GetAccountNumber = oGetAccountNumber }()
		got := config.GetConfigName()
		want := "test1"
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %+v, want: %+v", got, want)
		}
	})

	t.Run("should return test2", func(t *testing.T) {
		oGetAccountNumber := config.GetAccountNumber
		var toBeCallTimes int
		config.GetAccountNumber = func() *string {
			toBeCallTimes++
			n := "456"
			return &n
		}
		defer func() { config.GetAccountNumber = oGetAccountNumber }()
		got := config.GetConfigName()
		want := "test2"
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %+v, want: %+v", got, want)
		}
	})
}
