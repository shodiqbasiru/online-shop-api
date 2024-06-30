package config

import (
	"github.com/spf13/viper"
	"online-shop-api/internal/helper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath("../../")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	return config
}
