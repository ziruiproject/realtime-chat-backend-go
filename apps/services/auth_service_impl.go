package services

import (
	"context"
	"database/sql"
	configs "github.com/ziruiproject/realtime-chat-backend-go/apps/configuration"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/requests"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/repositories"
)

type AuthServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
}

func NewAuthService(UserRepository repositories.UserRepository, db *sql.DB) AuthService {
	return &AuthServiceImpl{
		UserRepository: UserRepository,
		DB: db,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request requests.UserLoginRequest) responses.AuthResponse {
	tx, err := service.DB.Begin()
	helpers.ErrorWithLog("Failed to make the transaction: ", err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.GetCredentials(ctx, tx, request.Email)
	helpers.ErrorWithLog("Failed retriving user", err)

	if !helpers.VerifyPassword(user.Password, request.Password) {
		return responses.AuthResponse{
			UserId:  nil,
			Token:   "",
			Success: false,
		}
	}

	// IMPLEMENT JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()


	tokenString, err := token.SignedString([]byte(configs.EnvConfigs.SecretToken))
	helpers.ErrorWithLog("Failed generating token", err)

	return responses.AuthResponse{
		UserId:  &user.Id,
		Token:   tokenString,
		Success: true,
	}
}
