package services

import (
	"go-clean-arch/domain/repositories"
	"reflect"
	"testing"
)

func TestNewCampaignService(t *testing.T) {
	type args struct {
		campaignRepo repositories.CampaignRepository
	}
	tests := []struct {
		name string
		args args
		want ICampaignService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCampaignService(tt.args.campaignRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCampaignService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignService_FindValidGoogleCampaignIds(t *testing.T) {
	type fields struct {
		campaignRepo repositories.CampaignRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &CampaignService{
				campaignRepo: tt.fields.campaignRepo,
			}
			got, err := svc.FindValidGoogleCampaignIds()
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignService.FindValidGoogleCampaignIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CampaignService.FindValidGoogleCampaignIds() = %v, want %v", got, tt.want)
			}
		})
	}
}
