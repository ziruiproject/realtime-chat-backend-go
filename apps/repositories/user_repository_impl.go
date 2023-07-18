package repositories

import (
	"context"
	"database/sql"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/models"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []models.User {
	var SQL string = "SELECT id, name, email, profile FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.ErrorWithLog("Failed executing query", err)
	defer helpers.ErrorCloseDefer(rows)

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(user.ID, user.Name, user.Email, user.Profile)
		helpers.ErrorWithLog("Failed scanning query", err)

		users = append(users, user)
	}

	return users
}

//
//func (repo *UserRepositoryImpl) GetUsers() (*models.User, error) {
//	query := `SELECT id, name, email, profile FROM users`
//	row := repo.Db.QueryRow(query)
//
//	user := &models.User{}
//	err := row.Scan(&user.ID, &user.Name, &user.Profile)
//
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, errors.New("user not found")
//		}
//		return nil, errors.Wrap(err, "failed to fetch user by ID")
//	}
//
//	return user, nil
//}
//
//}
