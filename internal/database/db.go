package database

import (
	"database/sql"
	"fmt"
	"online-shop-api/internal/config"
	"online-shop-api/internal/helper"
	"time"
)

func NewDB(config *config.Config) *sql.DB {
	dbHost, dbPort, dbName, dbUser, dbPassword := config.GetDatabaseConfig()

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", datasource)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
