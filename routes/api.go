package routes

import "github.com/gofiber/fiber/v2"

func SetupRoute(app *fiber.App) {
	app.Get("/", homeHandler)
}
func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
