package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"online-shop-api/helper"
	"online-shop-api/model/domain"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) SaveOrder(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	order.Id = uuid.New().String()
	query := "INSERT INTO t_order(id, customer_id, status) VALUES (?,?,?)"
	_, err := tx.ExecContext(ctx, query, order.Id, order.CustomerId, order.Status)
	helper.PanicIfError(err)
	return order
}

func (repository *OrderRepositoryImpl) SaveOrderDetails(ctx context.Context, tx *sql.Tx, orderDetails []domain.OrderDetail) []domain.OrderDetail {
	query := "INSERT INTO t_order_detail(id, order_id, product_id, qty, price) VALUES (?,?,?,?,?)"
	for _, detail := range orderDetails {
		detail.Id = uuid.New().String()
		_, err := tx.ExecContext(ctx, query, detail.Id, detail.OrderId, detail.ProductId, detail.Qty, detail.Price)
		helper.PanicIfError(err)
	}
	return orderDetails
}
