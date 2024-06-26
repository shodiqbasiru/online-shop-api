package repository

import (
	"context"
	"database/sql"
	"online-shop-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
