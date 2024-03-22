package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_AUTO_MIGRATE bool

	DB_HOST     string
	DB_PORT     int
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string

	JWT_SECRET_KEY string
}

func LoadConfig() (config Config) {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return
}
