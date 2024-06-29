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

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	CustomerService CustomerService
	ProductService  ProductService
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, customerService CustomerService, productService ProductService, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{OrderRepository: orderRepository, CustomerService: customerService, ProductService: productService, DB: DB, Validate: validate}
}

func (service *OrderServiceImpl) CreateOrder(ctx context.Context, request request.OrderRequest) response.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := service.CustomerService.GetById(ctx, request.CustomerId)

	order := domain.Order{
		Status:     domain.STATUS_PENDING,
		CustomerId: customer.Id,
	}
	order = service.OrderRepository.SaveOrder(ctx, tx, order)

	var orderDetails []domain.OrderDetail
	for _, detailRequest := range request.OrderDetails {
		product := service.ProductService.GetById(ctx, detailRequest.ProductId)

		if product.Stock-detailRequest.Qty < 0 {
			panic(exception.NewBadRequestError("Stock Product is not enough"))
		}
		product.Stock = product.Stock - detailRequest.Qty

		service.ProductService.UpdateStock(ctx, domain.Product(product))
		detail := domain.OrderDetail{
			Qty:       detailRequest.Qty,
			Price:     detailRequest.Qty * product.Price,
			OrderId:   order.Id,
			ProductId: product.Id,
		}
		orderDetails = append(orderDetails, detail)
	}
	orderDetails = service.CreateOrderDetails(ctx, tx, orderDetails)

	var detailsResponses []response.OrderDetailResponse
	for _, detailResponse := range orderDetails {
		detailsResponses = append(detailsResponses, response.OrderDetailResponse{
			Id:        detailResponse.Id,
			OrderId:   detailResponse.OrderId,
			ProductId: detailResponse.ProductId,
			Qty:       detailResponse.Qty,
			Price:     detailResponse.Price,
		})
	}

	return response.OrderResponse{
		Id:           order.Id,
		TransDate:    order.TransDate,
		Status:       order.Status,
		CustomerId:   order.Customer.Id,
		OrderDetails: detailsResponses,
	}
}

func (service *OrderServiceImpl) CreateOrderDetails(ctx context.Context, tx *sql.Tx, orderDetail []domain.OrderDetail) []domain.OrderDetail {
	return service.OrderRepository.SaveOrderDetails(ctx, tx, orderDetail)
}
