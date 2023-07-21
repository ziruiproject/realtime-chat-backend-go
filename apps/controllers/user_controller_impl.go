package controllers

import (
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/services"
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
	apiResponse := responses.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.JSON(apiResponse)
}
