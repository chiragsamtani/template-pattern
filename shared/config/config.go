package config

import (
	"os"

	"github.com/spf13/viper"
)

type ImmutableConfig interface {
	GetPort() string
	GetDatabaseUsername() string
	GetDatabasePassword() string
	GetDatabaseHost() string
	GetDatabaseName() string
	GetSecretKey() string
}

type RootConfig struct {
	Port             string `mapstructure:"PORT"`
	DatabaseUsername string `mapstructure:"MYSQL_DATABASE_USERNAME"`
	DatabasePassword string `mapstructure:"MYSQL_DATABASE_PASSWORD"`
	DatabaseHost     string `mapstructure:"MYSQL_DATABASE_HOST"`
	DatabaseName     string `mapstructure:"MYSQL_DATABASE"`
	JWTSecretKey     string `mapstructure:"JWT_SECRET_KEY"`
}

func (rc *RootConfig) GetPort() string {
	return rc.Port
}

func (rc *RootConfig) GetDatabaseUsername() string {
	return rc.DatabaseUsername
}

func (rc *RootConfig) GetDatabasePassword() string {
	return rc.DatabasePassword
}

func (rc *RootConfig) GetDatabaseHost() string {
	return rc.DatabaseHost
}

func (rc *RootConfig) GetDatabaseName() string {
	return rc.DatabaseName
}

func (rc *RootConfig) GetSecretKey() string {
	return rc.JWTSecretKey
}

func GetConfigFromEnv() (*RootConfig, error) {
	// use local config by default
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	appEnv, exists := os.LookupEnv("APP_ENV")
	if exists {
		if appEnv == "staging" {
			viper.SetConfigName("app.staging")
		} else {
			viper.SetConfigName("app.local")
		}
	} else {
		viper.SetConfigName("app.local")
	}
	var config RootConfig
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}
