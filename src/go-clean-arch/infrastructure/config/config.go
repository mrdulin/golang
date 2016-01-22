package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

type IApplicationConfig interface {
	Load() (*viper.Viper, error)
}

type ApplicationConfig struct{}

func NewApplicationConfig() IApplicationConfig {
	return &ApplicationConfig{}
}

func (appConf *ApplicationConfig) Load() (*viper.Viper, error) {
	v := viper.New()
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	v.SetConfigName("config.dev")
	v.AddConfigPath(basePath)
	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig()")
	}
	return v, nil
}
