package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/controllers"
)

func SetupRoute(app *fiber.App, controller *controllers.Controllers) {
	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("Masuk Route")
		return controller.UserController.GetAll(c)
	})
}
