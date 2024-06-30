package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"online-shop-api/internal/helper"
	"online-shop-api/internal/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	product.Id = uuid.New().String()
	query := "INSERT INTO m_product(id, name, description, price, stock, category_id) VALUES (?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, query, product.Id, product.Name, product.Description, product.Price, product.Stock, product.CategoryId)
	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error) {
	query := "SELECT id, name, description, price, stock, category_id FROM m_product WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	var product domain.Product
	if rows.Next() {
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CategoryId,
		)
		helper.PanicIfError(err)

		return product, nil
	} else {
		return domain.Product{}, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	query := "SELECT id, name, description, price, stock, category_id FROM m_product ORDER BY created_at DESC"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CategoryId,
		)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return products
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	query := "UPDATE m_product SET name = ?, description = ?, price = ?, stock = ?, category_id = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Stock, product.CategoryId, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	query := "DELETE FROM m_product WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, product.Id)
	helper.PanicIfError(err)
}
