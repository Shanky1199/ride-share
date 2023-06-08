package main

import (
	"log"
	"ride_share/pkg"
	"ride_share/pkg/middlewares"
	"ride_share/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// Custom Middleware
	app.Use(middlewares.AuthMiddleware())
	pkg.InitDatabase()

	// Routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
