package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
	"online-shop-api/internal/helper"
	"online-shop-api/internal/model/domain"
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

func (repository *OrderRepositoryImpl) FindOrderId(ctx context.Context, tx *sql.Tx, orderId string) (domain.Order, error) {
	query := `
		SELECT 
			o.id, o.customer_id, o.status, o.created_at, 
			od.id, od.order_id, od.product_id, od.qty, od.price
		FROM 
			t_order o 
		JOIN 
			t_order_detail od ON o.id = od.order_id 
		WHERE 
			o.id = ?`

	rows, err := tx.QueryContext(ctx, query, orderId)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return domain.Order{}, err
	}
	defer rows.Close()

	var order domain.Order
	order.OrderDetails = make([]domain.OrderDetail, 0)

	for rows.Next() {
		var detail domain.OrderDetail
		err := rows.Scan(
			&order.Id,
			&order.CustomerId,
			&order.Status,
			&order.TransDate,
			&detail.Id,
			&detail.OrderId,
			&detail.ProductId,
			&detail.Qty,
			&detail.Price,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return domain.Order{}, err
		}
		order.OrderDetails = append(order.OrderDetails, detail)
	}

	if len(order.OrderDetails) == 0 {
		return domain.Order{}, errors.New("order not found")
	}

	return order, nil
}
func (repository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Order {
	query := `
		SELECT 
			o.id, o.customer_id, o.status, o.created_at, 
			od.id, od.order_id, od.product_id, od.qty, od.price
		FROM 
			t_order o 
		LEFT JOIN 
			t_order_detail od ON o.id = od.order_id`

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	orderMap := make(map[string]*domain.Order)

	for rows.Next() {
		order := domain.Order{}
		orderDetail := domain.OrderDetail{}

		err := rows.Scan(
			&order.Id,
			&order.CustomerId,
			&order.Status,
			&order.TransDate,
			&orderDetail.Id,
			&orderDetail.OrderId,
			&orderDetail.ProductId,
			&orderDetail.Qty,
			&orderDetail.Price,
		)
		helper.PanicIfError(err)

		if existingOrder, ok := orderMap[order.Id]; ok {
			existingOrder.OrderDetails = append(existingOrder.OrderDetails, orderDetail)
		} else {
			order.OrderDetails = append(order.OrderDetails, orderDetail)
			orderMap[order.Id] = &order
		}
	}

	var orders []domain.Order
	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders
}

func (repository *OrderRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	query := "UPDATE t_order SET status=? WHERE id=?"
	_, err := tx.ExecContext(ctx, query, order.Status, order.Id)
	helper.PanicIfError(err)
	return order
}

func (repository *OrderRepositoryImpl) FindByStatus(ctx context.Context, tx *sql.Tx, status string) ([]domain.Order, error) {
	query := "SELECT id, customer_id, status, created_at FROM t_order WHERE status = ?"
	rows, err := tx.QueryContext(ctx, query, status)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		err := rows.Scan(&order.Id, &order.CustomerId, &order.Status, &order.TransDate)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
