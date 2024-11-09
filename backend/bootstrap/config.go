package bootstrap

import (
	"naive-backend/global"

	"github.com/donnie4w/go-logger/logger"
	"github.com/spf13/viper"
)

func ConfigInit() error {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error reading config file, %v", err)
		return err
	}

	if err := viper.Unmarshal(&global.Cfg); err != nil {
		logger.Errorf("Error unmarshaling config file, %s", err)
		return err
	}
	return nil
}
