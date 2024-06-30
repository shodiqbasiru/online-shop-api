package config

import (
	"os"
)

func GetDatabaseConfig() (string, string, string, string, string) {
	viper := NewViper()

	envName := viper.GetString("DATABASE_NAME")
	envUser := viper.GetString("DATABASE_USER")
	envPassword := viper.GetString("DATABASE_PASSWORD")
	dbHost := viper.GetString("DATABASE_HOST")
	dbPort := viper.GetString("DATABASE_PORT")
	dbName := os.Getenv(envName)
	dbUser := os.Getenv(envUser)
	dbPassword := os.Getenv(envPassword)

	return dbHost, dbPort, dbName, dbUser, dbPassword
}
