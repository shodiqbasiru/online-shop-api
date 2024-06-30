package config

import (
	"github.com/spf13/viper"
	"online-shop-api/internal/helper"
	"os"
)

type Config struct {
	Viper *viper.Viper
}

func NewConfig() *Config {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath("../../")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	helper.PanicIfError(err)
	return &Config{Viper: config}
}

func (c *Config) GetDatabaseConfig() (dbHost, dbPort, dbName, dbUser, dbPassword string) {
	envName := c.Viper.GetString("DATABASE_NAME")
	envUser := c.Viper.GetString("DATABASE_USER")
	envPassword := c.Viper.GetString("DATABASE_PASSWORD")
	dbHost = c.Viper.GetString("DATABASE_HOST")
	dbPort = c.Viper.GetString("DATABASE_PORT")
	dbName = os.Getenv(envName)
	dbUser = os.Getenv(envUser)
	dbPassword = os.Getenv(envPassword)

	return dbHost, dbPort, dbName, dbUser, dbPassword
}

func (c *Config) GetJwtConfig() []byte {
	envJwtSecret := c.Viper.GetString("JWT_SECRET")
	secretKey := []byte(os.Getenv(envJwtSecret))
	return secretKey
}
