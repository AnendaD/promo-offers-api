package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Level   string `mapstructure:"level"`
	Address string `mapstructure:"address"`
}

var AppConfig Config

func Load() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	viper.Unmarshal(&AppConfig)
}
