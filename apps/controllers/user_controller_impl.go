package controllers

import (
	"log"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/requests"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	"github.com/google/uuid"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"

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

func (controller *UserControllerImpl) GetById(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("id"))
	helpers.ErrorWithLog("Failed parsing id", err)

	userResponse := controller.UserService.GetById(c.Context(), userId)
	apiReesponse := responses.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
		}

		return c.JSON(apiReesponse)
}
