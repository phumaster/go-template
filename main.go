package main

import (
	"go-template/connectors"
	"go-template/middleware"
	"go-template/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// @title Go template API
// @version 1.0
// @description This is a Go template using Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name phumaster
// @contact.email phumaster.dev@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()
	PORT := ":3000"
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	middleware.Init(app)
	routes.RegisterAppRoutes(app)
	routes.RegisterAPIRoutes(app)
	routes.RegisterSwaggerRoutes(app)
	connectors.InitDB()
	log.Fatal(app.Listen(PORT))
}
