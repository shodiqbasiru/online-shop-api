package app

import (
	"github.com/julienschmidt/httprouter"
	"online-shop-api/internal/controller"
	"online-shop-api/internal/exception"
)

func NewRouter(
	categoryController *controller.CategoryController,
	productController *controller.ProductController,
	authController *controller.AuthController,
	customerController *controller.CustomerController,
	orderController *controller.OrderController,
) *httprouter.Router {
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

	// Auth Endpoint
	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/register-admin", authController.RegisterAdmin)
	router.POST("/api/auth/login", authController.LoginUser)

	// Customer EndPoint
	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/:customerId", customerController.FindById)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)

	// Order EndPoint
	router.POST("/api/orders", orderController.CreateOrder)
	router.GET("/api/orders/:orderId", orderController.FindById)
	router.GET("/api/orders/:orderId/status", orderController.UpdateStatusOrder)
	router.GET("/api/orders", orderController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
