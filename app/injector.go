//go:build wireinject
// +build wireinject

package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"online-shop-api/controller"
	"online-shop-api/middleware"
	"online-shop-api/repository"
	"online-shop-api/service"
)

func ProvideValidatorOptions() []validator.Option {
	return []validator.Option{}
}

func InitializedServer() *http.Server {
	wire.Build(
		NewDB,
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
		NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
