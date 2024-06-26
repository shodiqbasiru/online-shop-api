package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"online-shop-api/helper"
	"online-shop-api/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	customer.Id = uuid.New().String()
	query := "INSERT INTO m_customer(id, name, address) VALUES (?,?,?)"
	_, err := tx.ExecContext(ctx, query, customer.Id, customer.CustomerName, customer.Address)
	helper.PanicIfError(err)

	return customer
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customerId string) (domain.Customer, error) {
	query := "SELECT id, name, address FROM m_customer WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}
	if rows.Next() {
		err := rows.Scan(
			&customer.Id,
			&customer.CustomerName,
			&customer.Address,
		)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return domain.Customer{}, errors.New("customer not found")
	}
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	query := "SELECT id, name, address FROM m_customer ORDER BY created_at DESC"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(
			&customer.Id,
			&customer.CustomerName,
			&customer.Address,
		)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}
	return customers
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	query := "UPDATE m_customer SET name = ?, address= ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, customer.CustomerName, customer.Address, customer.Id)
	helper.PanicIfError(err)
	return customer
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer) {
	query := "DELETE FROM m_customer WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, customer.Id)
	helper.PanicIfError(err)
}
