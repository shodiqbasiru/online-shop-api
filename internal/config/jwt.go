package config

import "os"

func GetJwtConfig() []byte {
	viper := NewViper()
	envJwtSecret := viper.GetString("JWT_SECRET")
	secretKey := []byte(os.Getenv(envJwtSecret))
	return secretKey
}
