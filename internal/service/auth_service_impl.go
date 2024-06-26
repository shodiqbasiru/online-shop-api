package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"online-shop-api/internal/exception"
	"online-shop-api/internal/helper"
	"online-shop-api/internal/model/domain"
	"online-shop-api/internal/model/dto/request"
	"online-shop-api/internal/model/dto/response"
	"online-shop-api/internal/repository"
	"online-shop-api/utils"
)

type AuthServiceImpl struct {
	repository.UserRepository
	CustomerService
	DB       *sql.DB
	validate *validator.Validate
	JWT      *utils.JWT
}

func NewAuthService(userRepository repository.UserRepository, customerService CustomerService, DB *sql.DB, validate *validator.Validate, JWT *utils.JWT) AuthService {
	return &AuthServiceImpl{UserRepository: userRepository, CustomerService: customerService, DB: DB, validate: validate, JWT: JWT}
}

func (service *AuthServiceImpl) RegisterUser(ctx context.Context, request request.RegisterRequest) response.RegisterResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := domain.RoleCustomer
	password, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	customer := domain.Customer{
		CustomerName: request.CustomerName,
	}

	customer, err = service.CustomerService.CreateCustomer(ctx, customer)
	helper.PanicIfError(err)

	user := domain.User{
		NoHp:       request.NoHp,
		Email:      request.Email,
		Password:   password,
		Role:       role,
		CustomerId: customer.Id,
	}

	user = service.UserRepository.Register(ctx, tx, user)

	return helper.ToRegisterResponse(user, customer)
}

func (service *AuthServiceImpl) RegisterAdmin(ctx context.Context, request request.RegisterRequest) response.RegisterResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := domain.RoleAdmin
	password, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	customer := domain.Customer{
		CustomerName: request.CustomerName,
	}

	customer, err = service.CustomerService.CreateCustomer(ctx, customer)
	helper.PanicIfError(err)

	user := domain.User{
		NoHp:       request.NoHp,
		Email:      request.Email,
		Password:   password,
		Role:       role,
		CustomerId: customer.Id,
	}

	user = service.UserRepository.Register(ctx, tx, user)

	return helper.ToRegisterResponse(user, customer)
}

func (service *AuthServiceImpl) LoginUser(ctx context.Context, request request.LoginRequest) response.LoginResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmailOrNoHp(ctx, tx, request.EmailOrNoHp)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	match := helper.CheckPasswordIsMatch(request.Password, user.Password)
	if !match {
		panic(exception.NewBadRequestError("email or password is wrong"))
	}

	token, err := service.JWT.GenerateJwtToken(user)
	helper.PanicIfError(err)

	return response.LoginResponse{
		Id:    user.Id,
		Role:  user.Role,
		Token: token,
	}
}
