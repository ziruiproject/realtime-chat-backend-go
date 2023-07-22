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

func (controller *UserControllerImpl) Create(c *fiber.Ctx) error {
	var requestBody = requests.UserCreateRequset{}
	log.Println("Masuk controller")
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Request Body:", string(c.Body()))
		log.Println(err)
		return err
	}
	log.Println("ini body", requestBody)

	userResponse := controller.UserService.Create(c.Context(), requestBody)

	apiResponse := responses.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
		}

		return c.JSON(apiResponse)
}

func (controller *UserControllerImpl) Update(c *fiber.Ctx) error {
	var requestBody = requests.UserUpdateRequest{}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	userId, err := uuid.Parse(c.Params("id"))
	helpers.ErrorWithLog("Failed parsing id", err)

	userResponse := controller.UserService.Update(c.Context(), requestBody, userId)

	apiResponse := responses.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
		}

		return c.JSON(apiResponse)
}

func (controller *UserControllerImpl) Delete(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("id"))
	helpers.ErrorWithLog("Failed parsing id", err)

	controller.UserService.Delete(c.Context(), userId)

	apiResponse := responses.ApiResponse{
		Code:   200,
		Status: "OK",
		}

		return c.JSON(apiResponse)
}
