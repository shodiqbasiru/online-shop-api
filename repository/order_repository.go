package repository

import (
	"context"
	"database/sql"
	"online-shop-api/model/domain"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	SaveOrderDetails(ctx context.Context, tx *sql.Tx, orderDetails []domain.OrderDetail) []domain.OrderDetail
	FindOrderId(ctx context.Context, tx *sql.Tx, orderId string) (domain.Order, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Order
	UpdateStatus(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	FindByStatus(ctx context.Context, tx *sql.Tx, status string) ([]domain.Order, error)
}
