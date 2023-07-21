package helpers

import (
	"github.com/ziruiproject/realtime-chat-backend-go/apps/models"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"
)

func ToUserResponse(user models.User) responses.UserResponse {
	return responses.UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Profile: user.Profile,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
