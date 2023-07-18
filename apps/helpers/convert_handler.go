package helpers

import (
	"github.com/ziruiproject/realtime-chat-backend-go/apps/models"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web"
)

func ToUserResponse(user models.User) web.UserResponse{
	return web.UserResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Profile: user.Profile,
	}
}
