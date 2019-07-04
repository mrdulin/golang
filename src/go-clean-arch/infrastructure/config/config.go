package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-clean-arch/infrastructure/gcloud/datastore"
	"os"
	"path/filepath"
	"runtime"
)

type Source string

const (
	DataStore Source = "dataStore"
	Os        Source = "os"
)

type IApplicationConfig interface {
	Load() (*ApplicationConfig, error)
}

type ApplicationConfig struct {
	SqlInstanceConnectionName, SqlHost, SqlPort, SqlUser, SqlPassword, SqlDb, AdChannelApi string
	RefreshToken                                                                           string
	ClientCustomerId                                                                       int
}

func New(configSource Source, dataStoreService datastore.IService) (*ApplicationConfig, error) {
	if os.Getenv("ENV") != "production" {
		fmt.Println("load env vars from local fs env file")
		v := viper.New()
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)

		v.SetConfigName("config.dev")
		v.AddConfigPath(basePath)
		err := v.ReadInConfig()
		if err != nil {
			return nil, errors.Wrap(err, "viper.ReadInConfig()")
		}
		return &ApplicationConfig{
			SqlHost:          v.GetString("SQL_HOST"),
			SqlPort:          v.GetString("SQL_PORT"),
			SqlUser:          v.GetString("SQL_USER"),
			SqlPassword:      v.GetString("SQL_PASSWORD"),
			SqlDb:            v.GetString("SQL_DB"),
			AdChannelApi:     v.GetString("AD_CHANNEL_API"),
			ClientCustomerId: v.GetInt("ClientCustomerId"),
			RefreshToken:     v.GetString("RefreshToken"),
		}, nil
	} else {
		switch configSource {
		case Os:
			fmt.Println("load env vars from OS")
			return &ApplicationConfig{
				SqlHost:      os.Getenv("SQL_HOST"),
				SqlPort:      os.Getenv("SQL_PORT"),
				SqlUser:      os.Getenv("SQL_USER"),
				SqlPassword:  os.Getenv("SQL_PASSWORD"),
				SqlDb:        os.Getenv("SQL_DB"),
				AdChannelApi: os.Getenv("AD_CHANNEL_API"),
			}, nil
		case DataStore:
			fmt.Println("load env vars from dataStore")
			envVars, err := dataStoreService.GetEnvVars()
			if err != nil {
				return nil, err
			}
			return &ApplicationConfig{
				SqlInstanceConnectionName: envVars.SQL_INSTANCE_CONNECTION_NAME,
				SqlDb:                     envVars.SQL_DATABASE,
				SqlUser:                   envVars.SQL_USER,
				SqlPassword:               envVars.SQL_PASSWORD,
				AdChannelApi:              envVars.AD_CHANNEL_API_BASE_URL,
			}, nil
		default:
			return &ApplicationConfig{}, nil
		}
	}
}
