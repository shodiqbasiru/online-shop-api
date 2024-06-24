package repository

import (
	"context"
	"database/sql"
	"online-shop-api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId string) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
}
