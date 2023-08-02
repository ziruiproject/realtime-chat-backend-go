package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/requests"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"
)

type UserService interface {
	GetAll(ctx context.Context) []responses.UserResponse
	GetById(ctx context.Context, userId uuid.UUID) responses.UserResponse
	Create(ctx context.Context, request requests.UserCreateRequest) responses.UserResponse
	Update(ctx context.Context, request requests.UserUpdateRequest, userId uuid.UUID) responses.UserResponse
	Delete(ctx context.Context, userId uuid.UUID)
}
