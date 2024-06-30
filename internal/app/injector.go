//go:build wireinject
// +build wireinject

package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"online-shop-api/internal/config"
	"online-shop-api/internal/controller"
	"online-shop-api/internal/database"
	"online-shop-api/internal/middleware"
	"online-shop-api/internal/repository"
	"online-shop-api/internal/service"
	"online-shop-api/scheduler"
	"online-shop-api/utils"
)

func ProvideValidatorOptions() []validator.Option {
	return []validator.Option{}
}

func InitializedServer() *http.Server {
	wire.Build(
		config.NewConfig,
		database.NewDB,
		ProvideValidatorOptions,
		validator.New,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		repository.NewProductRepository,
		service.NewProductService,
		controller.NewProductController,
		repository.NewUserRepository,
		service.NewAuthService,
		controller.NewAuthController,
		repository.NewCustomerRepository,
		service.NewCustomerService,
		controller.NewCustomerController,
		repository.NewOrderRepository,
		service.NewOrderService,
		controller.NewOrderController,
		scheduler.NewScheduler,
		NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
		utils.NewJWT,
	)

	return nil
}
