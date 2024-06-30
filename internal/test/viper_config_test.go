package test

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"online-shop-api/internal/helper"
	"testing"
)

func TestConfigENVViper(t *testing.T) {
	config := viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath("../../")

	config.AutomaticEnv()

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	assert.Equal(t, "online-shop-api", config.GetString("APP_NAME"))
	assert.Equal(t, "MSFB", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
}
