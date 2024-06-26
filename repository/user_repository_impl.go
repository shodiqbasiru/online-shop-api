package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"online-shop-api/helper"
	"online-shop-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	user.Id = uuid.New().String()
	query := "INSERT INTO m_user(id, name, no_hp, email, password, role) VALUES (?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, query, user.Id, user.Name, user.NoHp, user.Email, user.Password, user.Role)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, email string, password string) domain.User {
	query := "SELECT id FROM m_user WHERE email = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, query, email, password)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id)
		helper.PanicIfError(err)
		return user
	} else {
		return domain.User{}
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	query := "SELECT id, name, no_hp, email, password, role FROM m_user WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.NoHp,
			&user.Email,
			&user.Password,
			&user.Role,
		)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return domain.User{}, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	query := "SELECT id, name, no_hp, email, password, role FROM m_user WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.NoHp,
			&user.Email,
			&user.Password,
			&user.Role,
		)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return domain.User{}, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindByEmailOrNoHp(ctx context.Context, tx *sql.Tx, emailOrNoHp string) (domain.User, error) {
	query := "SELECT id,name,password FROM m_user WHERE email = ? OR no_hp = ?"
	rows, err := tx.QueryContext(ctx, query, emailOrNoHp, emailOrNoHp)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return domain.User{}, errors.New("email or no hp is not registered")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT id, name, no_hp, email, password, role FROM m_user ORDER BY created_at DESC"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.NoHp,
			&user.Email,
			&user.Password,
			&user.Role,
		)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "UPDATE m_user SET name = ?, no_hp = ?, email = ?, password = ?, role = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Name, user.NoHp, user.Email, user.Password, user.Role, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	query := "DELETE FROM m_user WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Id)
	helper.PanicIfError(err)
}