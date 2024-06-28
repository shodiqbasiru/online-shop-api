package main

import (
	_ "github.com/go-sql-driver/mysql"
	"online-shop-api/app"
)

func main() {
	app.RunServer()
}
