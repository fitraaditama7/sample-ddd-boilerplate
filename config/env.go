package config

import (
	"ddd-boilerplate/pkg/logger"

	"github.com/spf13/viper"
)

func NewConfig() *Config {
	LoadConfigFromFile()
	return &Config{
		App:               NewAppConfig(),
		PostgreSQLConfig:  NewPostgreSQLConfig(),
		HttpClientConfig:  NewHttpClientConfig(),
		HealthCheckConfig: NewHealthCheckConfig(),
	}
}

func LoadConfigFromFile() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}
