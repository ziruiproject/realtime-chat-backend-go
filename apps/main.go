package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ziruiproject/realtime-chat-backend-go/routes"
)

func main() {
	app := fiber.New()
	routes.SetupRoute(app)
	app.Listen(":3000")
}
