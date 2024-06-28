package repository

import (
	"context"
	"database/sql"
	"online-shop-api/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, email string, password string) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindByEmailOrNoHp(ctx context.Context, tx *sql.Tx, emailOrNoHp string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
}
