package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	configs "github.com/ziruiproject/realtime-chat-backend-go/apps/configuration"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/web/responses"
	"log"
)

func AuthMiddleware(c *fiber.Ctx) error {
//	GET AUTH HEADER
	authHeader := c.Get("X-AUTH-HEADER")
	

	if authHeader == "" {
		return c.Status(401).JSON(responses.ApiResponse{
			Code: 401,
			Status: "Unauthorized",
			Data: "",
		})
	}

//	PARSE TOKEN
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.EnvConfigs.SecretToken), nil
	})

//	CHECK IF ERROR PARSING TOKEN
	if err != nil {
		return c.Status(401).JSON(responses.ApiResponse{
			Code: 401,
			Status: "Unauthorized",
			Data: "",
		})
	}

//	CHECK IF THE TOKEN VALID
	if !token.Valid {
		return c.Status(401).JSON(responses.ApiResponse{
			Code: 401,
			Status: "Unauthorized",
			Data: "",
		})
	}

	return c.Next()
	
}
