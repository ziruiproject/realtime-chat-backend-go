package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/models"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []models.User {
	var SQL string = `SELECT id, name, email, profile_img, created_at, updated_at FROM users`
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.ErrorWithLog("Failed retriving users", err)
	defer helpers.ErrorCloseRowsDefer(rows)

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Profile, &user.CreatedAt, &user.UpdatedAt)
		helpers.ErrorWithLog("Failed scanning query", err)

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (models.User, error) {
	var SQL string = `SELECT id, name, email, profile_img, created_at, updated_at FROM users WHERE id = $1`
	rows, err := tx.QueryContext(ctx, SQL, id)
	helpers.ErrorWithLog("Failed retriving user", err)
	defer helpers.ErrorCloseRowsDefer(rows)

	var user = models.User{}
	if !rows.Next() {
		return user, errors.New("user not found")
	}

	err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Profile, &user.CreatedAt, &user.UpdatedAt)
	helpers.ErrorWithLog("Failed scanning query", err)
	return user, nil
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user models.User) models.User {
	var SQL string = "INSERT INTO users(id,name,email,password,created_at,updated_at) values($1, $2, $3, $4, $5, $6)"

//	Hash Password
	newPassword, err := helpers.HashPassword(user.Password)
	helpers.ErrorWithLog("Failed creating user", err)

	_, err = tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Email, newPassword, user.CreatedAt, user.UpdatedAt)
	helpers.ErrorWithLog("Failed creating user", err)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user models.User) models.User {
	var SQL string = `UPDATE users SET
			email = $1,
			name = $2,
			created_at = $3,
			updated_at = $4
			WHERE id::text = $5;`

	_, err := tx.QueryContext(ctx, SQL, user.Email, user.Name, user.CreatedAt, user.UpdatedAt, user.Id)
	helpers.ErrorWithLog("Failed updating user", err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) {
	var SQL string = `DELETE FROM users WHERE id = $1`
	_, err := tx.ExecContext(ctx, SQL, id)
	helpers.ErrorWithLog("Failed deleting user", err)
}

func (repository *UserRepositoryImpl) GetCredentials(ctx context.Context, tx *sql.Tx, email string) (models.User, error) {
	var SQL string = `SELECT id, password, created_at, updated_at FROM users WHERE email = $1`
	rows, err := tx.QueryContext(ctx, SQL, email)
	helpers.ErrorWithLog("Failed getting user", err)
	defer helpers.ErrorCloseRowsDefer(rows)

	var user = models.User{}
	if !rows.Next() {
		return user, errors.New("user not found")
	}

	err = rows.Scan(&user.Id, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	helpers.ErrorWithLog("Failed scanning query", err)
	return user, nil
}