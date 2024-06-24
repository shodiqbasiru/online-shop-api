package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"online-shop-api/app"
	"online-shop-api/controller"
	"online-shop-api/helper"
	"online-shop-api/repository"
	"online-shop-api/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(*categoryController)
	server := app.NewServer(router)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
