package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/controllers"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"
	"github.com/ziruiproject/realtime-chat-backend-go/databases"
	"github.com/ziruiproject/realtime-chat-backend-go/middlewares"
	"github.com/ziruiproject/realtime-chat-backend-go/routes"
)

func main() {
	app := fiber.New()

	helpers.LoadEnv()

	db, err := databases.Connection()
	helpers.ErrorWithLog("Error when connecting to db: ", err)

	loadController := controllers.SetupController(db)

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	app.Use("/users/:id", middlewares.AuthMiddleware)

	routes.SetupRoute(app, loadController)
	app.Listen(":3000")
}
