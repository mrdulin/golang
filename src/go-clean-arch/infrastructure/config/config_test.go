package config

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestNewApplicationConfig(t *testing.T) {
	tests := []struct {
		name string
		want IApplicationConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApplicationConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApplicationConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplicationConfig_Load(t *testing.T) {
	tests := []struct {
		name    string
		appConf *ApplicationConfig
		want    *viper.Viper
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appConf := &ApplicationConfig{}
			got, err := appConf.Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplicationConfig.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplicationConfig.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
