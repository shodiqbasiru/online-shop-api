package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"online-shop-api/exception"
	"online-shop-api/helper"
	"online-shop-api/model/domain"
	"online-shop-api/model/dto/request"
	"online-shop-api/model/dto/response"
	"online-shop-api/repository"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	AuthService
	DB       *sql.DB
	Validate *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{CustomerRepository: customerRepository, DB: DB, Validate: validate}
}

func (service *CustomerServiceImpl) CreateCustomer(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer = service.CustomerRepository.Save(ctx, tx, customer)
	return customer, nil
}

func (service *CustomerServiceImpl) GetById(ctx context.Context, customerId string) response.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) GetAll(ctx context.Context) []response.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers := service.CustomerRepository.FindAll(ctx, tx)

	var customerResponses []response.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, helper.ToCustomerResponse(customer))
	}
	return customerResponses
}

func (service *CustomerServiceImpl) UpdateCustomer(ctx context.Context, request request.CustomerUpdateRequest) response.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customer.CustomerName = request.CustomerName
	customer.Address = request.Address

	customer = service.CustomerRepository.Update(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) DeleteCustomer(ctx context.Context, customerId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CustomerRepository.Delete(ctx, tx, customer)
}
