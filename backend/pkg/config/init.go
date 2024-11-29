package config

import (
	"github.com/donnie4w/go-logger/logger"
	"github.com/spf13/viper"
)

func Init() (cfg *Config) {
	cfg = new(Config)

	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("Error reading config file, %v", err)
		panic(err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		logger.Fatalf("Error unmarshaling config file, %s", err)
		panic(err)
	}
	return
}
