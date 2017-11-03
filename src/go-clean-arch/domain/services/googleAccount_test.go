package services

import (
	"go-clean-arch/domain/models/cedar"
	"go-clean-arch/domain/repositories"
	"reflect"
	"testing"
)

func TestNewGoogleAccountService(t *testing.T) {
	type args struct {
		googleAccountRepo repositories.GoogleAccountRepository
		locationRepo      repositories.LocationRepository
	}
	tests := []struct {
		name string
		args args
		want IGoogleAccountService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGoogleAccountService(tt.args.googleAccountRepo, tt.args.locationRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoogleAccountService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAccountService_FindGoogleAccountsForReport(t *testing.T) {
	type fields struct {
		googleAccountRepo repositories.GoogleAccountRepository
		locationRepo      repositories.LocationRepository
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
			svc := &GoogleAccountService{
				googleAccountRepo: tt.fields.googleAccountRepo,
				locationRepo:      tt.fields.locationRepo,
			}
			got, err := svc.FindGoogleAccountsForReport()
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleAccountService.FindGoogleAccountsForReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoogleAccountService.FindGoogleAccountsForReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
