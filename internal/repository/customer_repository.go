package repository

import (
	"context"
	"database/sql"
	"online-shop-api/internal/model/domain"
)

type CustomerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	FindById(ctx context.Context, tx *sql.Tx, customerId string) (domain.Customer, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
	Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer)
}
