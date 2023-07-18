package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/services"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) GetAll(c *fiber.Ctx) error {
	log.Println("Masuk Controller")
	userResponse := controller.UserService.GetAll(c.Context())
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.JSON(apiResponse)
}
