package service

import (
	"context"
	"online-shop-api/model/domain"
	"online-shop-api/model/dto/request"
	"online-shop-api/model/dto/response"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request request.OrderRequest) response.OrderResponse
	CreateOrderDetails(ctx context.Context, orderDetail []domain.OrderDetail) []domain.OrderDetail
}
