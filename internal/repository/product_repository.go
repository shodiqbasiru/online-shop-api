package repository

import (
	"context"
	"database/sql"
	"online-shop-api/internal/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
}
