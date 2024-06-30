package app

import (
	"database/sql"
	"online-shop-api/helper"
	"os"
	"time"
)

var dbName = []byte(os.Getenv("DB_NAME"))
var dbUser = []byte(os.Getenv("DB_USER"))
var dbPassword = []byte(os.Getenv("DB_PASS"))

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", string(dbUser)+":"+string(dbPassword)+"@tcp(localhost:3306)/"+string(dbName)+"?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
