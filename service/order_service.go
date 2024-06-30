package service

import (
	"context"
	"database/sql"
	"online-shop-api/model/domain"
	"online-shop-api/model/dto/request"
	"online-shop-api/model/dto/response"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request request.OrderRequest) response.OrderResponse
	CreateOrderDetails(ctx context.Context, tx *sql.Tx, orderDetail []domain.OrderDetail) []domain.OrderDetail
	GetById(ctx context.Context, orderId string) response.OrderResponse
	UpdateStatusOrder(ctx context.Context, orderId string) string
	TaskCancelOrder(ctx context.Context) error
}
