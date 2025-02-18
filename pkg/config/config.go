package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	ServerPort string
	Database   struct {
		Dialect  string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	JWTSecret string
}

var Config Configuration

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
