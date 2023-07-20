package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/models"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []models.User {
	log.Println("Masuk Repository")
	var SQL string = "SELECT id, name, email, profile_img FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.ErrorWithLog("Failed executing query", err)
	defer helpers.ErrorCloseRowsDefer(rows)

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Profile)
		helpers.ErrorWithLog("Failed scanning query", err)

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id string) (models.User, error) {
	var SQL string = "SELECT id, name, email, profile_img FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helpers.ErrorWithLog("Failed executing query", err)
	defer helpers.ErrorCloseRowsDefer(rows)

	user := models.User{}

	if !rows.Next() {
		return user, errors.New("user not found")
	}

	err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Profile)
	helpers.ErrorWithLog("Failed scanning query", err)
	return user, nil
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user models.User) models.User {
	var SQL string = "INISERT INTO users(id, name, email, password, profile_img) values (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Email, user.Password, user.Profile)
	helpers.ErrorWithLog("Failed executing query", err)

	return user
}