package controllers

import (
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/requests"

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
	var apiResponse responses.ApiResponse
	userResponse := controller.UserService.GetAll(c.Context())

	if len(userResponse) == 0 {
		apiResponse = responses.ApiResponse{
			Code: 404,
			Status: "Users Not Found",
			Data: "",
		}

		return c.Status(404).JSON(apiResponse)
	}

	apiResponse = responses.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.Status(200).JSON(apiResponse)
}

func (controller *UserControllerImpl) GetById(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("id"))

	var apiResponse responses.ApiResponse

//	Check if id requested valid
	if err != nil {
		apiResponse = responses.ApiResponse{
			Code: 400,
			Status: "Bad Request",
			Data: "",
		}
		return c.Status(400).JSON(apiResponse)
	}

	userResponse := controller.UserService.GetById(c.Context(), userId)

//	Check if user empty
	if userResponse == (responses.UserResponse{}) {
		apiResponse = responses.ApiResponse{
			Code: 404,
			Status: "User Not Found",
			Data: "",
		}
		return c.Status(404).JSON(apiResponse)
	}

//	If all valid, set user data
	apiResponse = responses.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.Status(200).JSON(apiResponse)
}

func (controller *UserControllerImpl) Create(c *fiber.Ctx) error {
	var apiResponse responses.ApiResponse
	var requestBody = requests.UserCreateRequest{}
	err := c.BodyParser(&requestBody)

	if err != nil {
		apiResponse = responses.ApiResponse{
			Code: 500,
			Status: "Internal Server Error",
			Data: "",
		}
		return c.Status(500).JSON(apiResponse)
	}

	userResponse := controller.UserService.Create(c.Context(), requestBody)

	apiResponse = responses.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.Status(200).JSON(apiResponse)
}

func (controller *UserControllerImpl) Update(c *fiber.Ctx) error {
	var apiResponse responses.ApiResponse
	var requestBody = requests.UserUpdateRequest{}

	if err := c.BodyParser(&requestBody); err != nil {
		apiResponse = responses.ApiResponse{
			Code: 500,
			Status: "Internal Server Error",
			Data: "",
		}
		return c.Status(500).JSON(apiResponse)
	}

	userId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		apiResponse = responses.ApiResponse{
			Code: 400,
			Status: "Bad Request",
			Data: "",
		}
		return c.Status(400).JSON(apiResponse)
	}

	userResponse := controller.UserService.Update(c.Context(), requestBody, userId)

	apiResponse = responses.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.Status(200).JSON(apiResponse)
}

func (controller *UserControllerImpl) Delete(c *fiber.Ctx) error {
	var apiResponse responses.ApiResponse;
	userId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		apiResponse = responses.ApiResponse{
			Code: 400,
			Status: "Bad Request",
			Data: "",
		}
		return c.Status(400).JSON(apiResponse)
	}

	controller.UserService.Delete(c.Context(), userId)

	apiResponse = responses.ApiResponse{
		Code:   200,
		Status: "OK",
	}

	return c.Status(200).JSON(apiResponse)
}
