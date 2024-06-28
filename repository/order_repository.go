package repository

import (
	"context"
	"database/sql"
	"online-shop-api/model/domain"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	SaveOrderDetails(ctx context.Context, tx *sql.Tx, orderDetails []domain.OrderDetail) []domain.OrderDetail
}
