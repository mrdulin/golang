package repositories

import (
	models "go-clean-arch/domain/models/adChannel"
	"go-clean-arch/domain/repositories"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestNewCampaignResultRepository(t *testing.T) {
	type args struct {
		Db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want repositories.CampaignResultRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCampaignResultRepository(tt.args.Db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCampaignResultRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignResultRepository_UpdateStatusTransaction(t *testing.T) {
	type fields struct {
		Db *sqlx.DB
	}
	type args struct {
		row *models.AdPerformanceReportRow
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CampaignResultRepository{
				Db: tt.fields.Db,
			}
			if err := repo.UpdateStatusTransaction(tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("CampaignResultRepository.UpdateStatusTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
