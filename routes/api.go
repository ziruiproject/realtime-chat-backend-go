package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/controllers"
)

func SetupRoute(app *fiber.App, controller *controllers.Controllers) {
	app.Get("/users", func(c *fiber.Ctx) error {
		return controller.UserController.GetAll(c)
	})
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		return controller.UserController.GetById(c)
	})
	app.Post("/users", func(c *fiber.Ctx) error {
		return controller.UserController.Create(c)
	})
	app.Patch("/users/:id", func(c *fiber.Ctx) error {
		return controller.UserController.Update(c)
	})
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		return controller.UserController.Delete(c)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return controller.AuthController.Login(c)
	})
}
