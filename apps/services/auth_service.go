package services

import (
	"context"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/requests"
)

type AuthService interface {
	Login(ctx context.Context, request requests.UserLoginRequest) responses.AuthResponse
}
