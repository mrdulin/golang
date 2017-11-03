package repositories

import (
	"go-clean-arch/domain/models/cedar"
	"go-clean-arch/domain/repositories"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestNewLocationRepository(t *testing.T) {
	type args struct {
		Db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want repositories.LocationRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLocationRepository(tt.args.Db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLocationRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocationRepository_FindLocationsBoundGoogleClientCustomerId(t *testing.T) {
	type fields struct {
		Db *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []cedar.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			locationRepo := &LocationRepository{
				Db: tt.fields.Db,
			}
			got, err := locationRepo.FindLocationsBoundGoogleClientCustomerId()
			if (err != nil) != tt.wantErr {
				t.Errorf("LocationRepository.FindLocationsBoundGoogleClientCustomerId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocationRepository.FindLocationsBoundGoogleClientCustomerId() = %v, want %v", got, tt.want)
			}
		})
	}
}
