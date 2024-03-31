package main

import (
	"log"
	"main/database"
	"main/routes"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)
	app.Post("/tasks", routes.CreateTask)
}

func main() {
	// Connect to the database
	database.Connect()

	// Create a new Fiber instance
	app := fiber.New()

	// Define the routes
	setupRoutes(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
