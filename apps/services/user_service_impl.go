package services

import (
	"context"
	"database/sql"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
}

func (service *UserServiceImpl) GetAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.ErrorWithLog("Failed to make the transaction", err)
	defer helpers.CommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)

	var userResponse []web.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, helpers.ToUserResponse(user))
	}

	return userResponse
}
