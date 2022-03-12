package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host       string
	Port       string
	UserName   string
	Password   string
	Database   string
	AuthDB     string
	Collection string
}

func ViperConfig() (dbcnfg *DBConfig, err error) {

	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return dbcnfg, fmt.Errorf("fatal error config file: %w ", err)
	}

	dbcnfg = &DBConfig{
		Host:       viper.GetString("mongodb.host"),
		Port:       viper.GetString("mongodb.port"),
		UserName:   viper.GetString("mongodb.username"),
		Password:   viper.GetString("mongodb.password"),
		Database:   viper.GetString("mongodb.database"),
		AuthDB:     viper.GetString("mongodb.auth_db"),
		Collection: viper.GetString("mongodb.collection"),
	}
	return
}
