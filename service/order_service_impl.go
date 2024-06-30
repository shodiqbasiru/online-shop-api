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
		CustomerId:   customer.Id,
		OrderDetails: detailsResponses,
	}
}

func (service *OrderServiceImpl) CreateOrderDetails(ctx context.Context, tx *sql.Tx, orderDetail []domain.OrderDetail) []domain.OrderDetail {
	return service.OrderRepository.SaveOrderDetails(ctx, tx, orderDetail)
}

func (service *OrderServiceImpl) GetById(ctx context.Context, orderId string) response.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindOrderId(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var orderDetails []response.OrderDetailResponse
	for _, detail := range order.OrderDetails {
		orderDetails = append(orderDetails, response.OrderDetailResponse{
			Id:        detail.Id,
			Qty:       detail.Qty,
			Price:     detail.Price,
			OrderId:   detail.OrderId,
			ProductId: detail.ProductId,
		})
	}

	return response.OrderResponse{
		Id:           order.Id,
		TransDate:    order.TransDate,
		Status:       order.Status,
		CustomerId:   order.CustomerId,
		OrderDetails: orderDetails,
	}
}

func (service *OrderServiceImpl) GetAll(ctx context.Context) []response.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := service.OrderRepository.FindAll(ctx, tx)

	var orderResponses []response.OrderResponse
	for _, order := range orders {
		var orderDetails []response.OrderDetailResponse
		for _, detail := range order.OrderDetails {
			orderDetails = append(orderDetails, response.OrderDetailResponse{
				Id:        detail.Id,
				Qty:       detail.Qty,
				Price:     detail.Price,
				OrderId:   detail.OrderId,
				ProductId: detail.ProductId,
			})
		}
		orderResponses = append(orderResponses, response.OrderResponse{
			Id:           order.Id,
			TransDate:    order.TransDate,
			Status:       order.Status,
			CustomerId:   order.CustomerId,
			OrderDetails: orderDetails,
		})
	}

	return orderResponses
}

func (service *OrderServiceImpl) UpdateStatusOrder(ctx context.Context, orderId string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	currentOrder, err := service.OrderRepository.FindOrderId(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	currentOrder.Status = domain.STATUS_SUCCESS
	service.OrderRepository.UpdateStatus(ctx, tx, currentOrder)

	return "Updated Status Order Successfully"
}

func (service *OrderServiceImpl) TaskCancelOrder(ctx context.Context) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders, err := service.OrderRepository.FindByStatus(ctx, tx, domain.STATUS_PENDING)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	for _, order := range orders {
		order.Status = domain.STATUS_CANCEL
		service.OrderRepository.UpdateStatus(ctx, tx, order)
	}

	return nil
}
