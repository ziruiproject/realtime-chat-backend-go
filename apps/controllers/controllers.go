package controllers

import (
	"database/sql"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/repositories"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/services"
)

type Controllers struct {
	UserController UserController
	AuthController AuthController
}

func SetupController(db *sql.DB) *Controllers {
	return &Controllers{
		UserController: initUserController(db),
		AuthController: initAuthController(db),
	}
}

func initUserController(db *sql.DB) UserController {
	UserRepository := repositories.NewUserRepository()
	UserService := services.NewUserService(UserRepository, db)
	return NewUserController(UserService)
}

func initAuthController(db *sql.DB) AuthController {
	UserRepository := repositories.NewUserRepository()
	AuthService := services.NewAuthService(UserRepository, db)
	return NewAuthController(AuthService)
}
