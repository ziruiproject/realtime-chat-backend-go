package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository repositories.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) GetAll(ctx context.Context) []web.UserResponse {
	log.Println(service.DB)
	log.Println("Masuk Service")
	tx, err := service.DB.Begin()
	helpers.ErrorWithLog("Failed to make the transaction: ", err)
	defer helpers.CommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)

	var userResponse []web.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, helpers.ToUserResponse(user))
	}

	return userResponse
}
