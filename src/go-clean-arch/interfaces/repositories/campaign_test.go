package repositories

import (
	"go-clean-arch/domain/models/cedar/campaign"
	"go-clean-arch/domain/repositories"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestNewCampaignRepository(t *testing.T) {
	type args struct {
		Db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want repositories.CampaignRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCampaignRepository(tt.args.Db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCampaignRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignRepository_FindById(t *testing.T) {
	type fields struct {
		Db *sqlx.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *campaign.Campaign
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &CampaignRepository{
				Db: tt.fields.Db,
			}
			got, err := cr.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignRepository.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CampaignRepository.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignRepository_FindValidGoogleCampaign(t *testing.T) {
	type fields struct {
		Db *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*campaign.Campaign
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &CampaignRepository{
				Db: tt.fields.Db,
			}
			got, err := cr.FindValidGoogleCampaign()
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignRepository.FindValidGoogleCampaign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CampaignRepository.FindValidGoogleCampaign() = %v, want %v", got, tt.want)
			}
		})
	}
}
