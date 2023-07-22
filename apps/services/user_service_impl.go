package services

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/models"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/repositories"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/requests"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"
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

func (service *UserServiceImpl) GetAll(ctx context.Context) []responses.UserResponse {
	log.Println("Masuk Service")
	tx, err := service.DB.Begin()
	helpers.ErrorWithLog("Failed to make the transaction: ", err)
	defer helpers.CommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)

	var userResponse []responses.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, helpers.ToUserResponse(user))
	}

	return userResponse
}

func (service *UserServiceImpl) GetById(ctx context.Context, userId uuid.UUID) responses.UserResponse {
	tx, err := service.DB.Begin()
	helpers.ErrorWithLog("Failed to make transaction", err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.GetById(ctx, tx, userId)
	helpers.ErrorWithLog("Failed retrivign user", err)

	return helpers.ToUserResponse(user)
}

func (service *UserServiceImpl) Create(ctx context.Context, request requests.UserCreateRequset) responses.UserResponse {
	tx, err := service.DB.Begin()
	helpers.ErrorWithLog("Failed to make transaction", err)
	defer helpers.CommitOrRollback(tx)

	user := models.User{
		Id:        uuid.New(),
		Email:     request.Email,
		Name:      request.Name,
		Password:  request.Password,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helpers.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request requests.UserUpdateRequest, userId uuid.UUID) responses.UserResponse {
	tx, err := service.DB.Begin()
	helpers.ErrorWithLog("Failed to make transaction", err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.GetById(ctx, tx, userId)
	helpers.ErrorWithLog("Failed retriving user", err)

	user = models.User{
		Id:        user.Id,
		Name:      helpers.IsUpdateRequired(user.Name, request.Name).(string),
		Email:     helpers.IsUpdateRequired(user.Email, request.Email).(string),
		Profile:   helpers.IsUpdateRequired(user.Profile, request.Profile).(string),
		UpdatedAt: time.Now().Unix(),
	}

	user = service.UserRepository.Update(ctx, tx, user)

	return helpers.ToUserResponse(user)
}


