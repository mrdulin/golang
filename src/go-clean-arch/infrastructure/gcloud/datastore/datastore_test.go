package datastore

import (
	"testing"
)

var options = ServiceOptions{}

func TestDataStoreService_GetEnvVars(t *testing.T) {
	dataStoreService, err := NewDataStoreService(&options)
	if err != nil {
		t.Logf("%#v", err)
	}
	entities, err := dataStoreService.GetEnvVars()
	if err != nil {
		t.Logf("%#v", err)
	}
	t.Logf("%#v", entities)
}
