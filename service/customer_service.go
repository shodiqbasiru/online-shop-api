package service

import (
	"context"
	"online-shop-api/model/domain"
	"online-shop-api/model/dto/request"
	"online-shop-api/model/dto/response"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	GetById(ctx context.Context, customerId string) response.CustomerResponse
	GetAll(ctx context.Context) []response.CustomerResponse
	UpdateCustomer(ctx context.Context, request request.CustomerUpdateRequest) response.CustomerResponse
	DeleteCustomer(ctx context.Context, customerId string)
}
