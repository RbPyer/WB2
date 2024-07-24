package config

import (
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Host string
	Port string
}

func NewConfig() *ServerConfig {
	err := initConfig()
	if err != nil {
		log.Fatalf("init config err: %v", err)
	}

	return &ServerConfig{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}

}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
