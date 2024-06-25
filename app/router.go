package app

import (
	"github.com/julienschmidt/httprouter"
	"online-shop-api/controller"
)

func NewRouter(categoryController *controller.CategoryController, productController *controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	// Category End Point
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Product End Point
	router.POST("/api/products", productController.Create)
	router.GET("/api/products/:productId", productController.FindById)
	router.GET("/api/products", productController.FindAll)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	return router
}
