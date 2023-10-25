package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App *AppConfig

	//DB
	PostgreSQLConfig  *PostgreSQLConfig
	HttpClientConfig  *HttpClientConfig
	RedisConfig       *RedisConfig
	KafkaConfig       *KafkaConfig
	HealthCheckConfig *HealthCheckConfig
}

type AppConfig struct {
	Host string
	Port string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Host: viper.GetString("APP_HOST"),
		Port: viper.GetString("APP_PORT"),
	}
}

type PostgreSQLConfig struct {
	HostPrimary   string
	HostSecondary string
	Port          string
	Name          string
	Username      string
	Password      string
}

func NewPostgreSQLConfig() *PostgreSQLConfig {
	return &PostgreSQLConfig{
		HostPrimary:   viper.GetString("POSTGRESQL_PRIMARY_HOST"),
		HostSecondary: viper.GetString("POSTGRESQL_SECONDARY_HOST"),
		Port:          viper.GetString("POSTGRESQL_PORT"),
		Username:      viper.GetString("POSTGRESQL_USERNAME"),
		Password:      viper.GetString("POSTGRESQL_PASSWORD"),
		Name:          viper.GetString("POSTGRESQL_DATABASE_NAME"),
	}
}

type RedisConfig struct {
	Host string
	DB   int
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host: viper.GetString("REDIS_HOST"),
		DB:   viper.GetInt("REDIS_DB"),
	}
}

type KafkaConfig struct {
	BrokerURL string
	Partition int
}

type HttpClientConfig struct {
	Host                string
	MaxTimeout          int
	MaxIdleConns        int
	MaxConnsPerHost     int
	MaxIdleConnsPerHost int
}

func NewHttpClientConfig() *HttpClientConfig {
	return &HttpClientConfig{
		Host:                viper.GetString("HTTP_CLIENT_HOST"),
		MaxIdleConns:        viper.GetInt("HTTP_CLIENT_MAX_IDLE_CONNS"),
		MaxConnsPerHost:     viper.GetInt("HTTP_CLIENT_MAX_CONNS_PER_HOST"),
		MaxIdleConnsPerHost: viper.GetInt("MAX_IDLE_CONNS_PER_HOST"),
		MaxTimeout:          viper.GetInt("MAX_TIMEOUT"),
	}
}

type HealthCheckConfig struct {
	WaitStartupTime   int
	WaitLivenessTime  int
	WaitReadinessTime int
	HealthCheckTime   int
}

func NewHealthCheckConfig() *HealthCheckConfig {
	return &HealthCheckConfig{
		WaitStartupTime:   viper.GetInt("HEALTH_CHECK_WAIT_STARTUP_TIME"),
		WaitLivenessTime:  viper.GetInt("HEALTH_CHECK_WAIT_LIVENESS_TIME"),
		WaitReadinessTime: viper.GetInt("HEALTH_CHECK_WAIT_READINESS_TIME"),
		HealthCheckTime:   viper.GetInt("HEALTH_CHECK_TIME"),
	}
}
