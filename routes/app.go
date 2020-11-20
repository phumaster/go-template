package routes

import "github.com/gofiber/fiber/v2"

// RegisterAppRoutes method to register App router
func RegisterAppRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!!!")
	})
}
