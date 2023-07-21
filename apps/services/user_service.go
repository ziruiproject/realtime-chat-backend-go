package services

import (
	"context"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/models"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web"
)

type UserService interface {
	GetAll(ctx context.Context) []web.UserResponse
	GetById(ctx context.Context, userId string) web.UserResponse
	Create(ctx context.Context, user models.User) web.UserResponse
	Update(ctx context.Context, user models.User) web.UserResponse
	Delete(ctx context.Context, userId string)
}
