package services

import (
	"context"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web"
)

type UserService interface {
	GetAll(ctx context.Context) []web.UserResponse
}
