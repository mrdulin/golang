package repositories

import (
	"go-clean-arch/domain/models/cedar"
	"go-clean-arch/domain/repositories"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestNewGoogleAccountRepository(t *testing.T) {
	type args struct {
		Db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want repositories.GoogleAccountRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGoogleAccountRepository(tt.args.Db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoogleAccountRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAccountRepository_FindByClientCustomerIds(t *testing.T) {
	type fields struct {
		Db *sqlx.DB
	}
	type args struct {
		ids []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []cedar.GoogleAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			googleAccountRepo := &GoogleAccountRepository{
				Db: tt.fields.Db,
			}
			got, err := googleAccountRepo.FindByClientCustomerIds(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleAccountRepository.FindByClientCustomerIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoogleAccountRepository.FindByClientCustomerIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAccountRepository_FindByCampaignRanByZOWIForZELO(t *testing.T) {
	type fields struct {
		Db *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []cedar.GoogleAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			googleAccountRepo := &GoogleAccountRepository{
				Db: tt.fields.Db,
			}
			got, err := googleAccountRepo.FindByCampaignRanByZOWIForZELO()
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleAccountRepository.FindByCampaignRanByZOWIForZELO() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoogleAccountRepository.FindByCampaignRanByZOWIForZELO() = %v, want %v", got, tt.want)
			}
		})
	}
}
