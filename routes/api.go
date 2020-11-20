package routes

import (
	"go-template/controllers"

	"github.com/gofiber/fiber/v2"
)

// RegisterAPIRoutes method to register api router
func RegisterAPIRoutes(app *fiber.App) {
	user := controllers.UserController{}
	app.Get("/api/user/:ID", user.GetByID)
	app.Post("/api/user/create", user.Create)
}
