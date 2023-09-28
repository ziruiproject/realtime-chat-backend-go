package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/services"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/requests"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"
)

type AuthControllerImpl struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(c *fiber.Ctx) error {
	var apiResponse responses.ApiResponse
	var requestBody = requests.UserLoginRequest{}
	err := c.BodyParser(&requestBody)

	if err != nil {
		apiResponse = responses.ApiResponse{
			Code: 500,
			Status: "Internal Server Error",
			Data: "",
		}
		return c.Status(500).JSON(apiResponse)
	}

	userResponse := controller.AuthService.Login(c.Context(), requestBody)

	if userResponse.UserId == nil {
		apiResponse = responses.ApiResponse{
			Code: 401,
			Status: "Unauthorized",
			Data: "",
		}
		return c.Status(401).JSON(apiResponse)
	}

	apiResponse = responses.ApiResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}

	return c.Status(200).JSON(apiResponse)
}
