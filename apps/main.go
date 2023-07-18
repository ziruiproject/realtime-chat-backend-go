package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/controllers"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"
	"github.com/ziruiproject/realtime-chat-backend-go/databases"
	"github.com/ziruiproject/realtime-chat-backend-go/routes"
)

func main() {
	app := fiber.New()
	db, err := databases.Connection()
	helpers.ErrorWithLog("Error when connecting to db: ", err)
	
	loadController := controllers.SetupController(db)

	routes.SetupRoute(app, loadController)
	app.Listen(":3000")
}
