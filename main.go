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
	app.Get("/tasks", routes.GetTasks)
	app.Get("/tasks/:id", routes.GetTaskById)

	app.Put("/tasks/:id", routes.UpdateTasksById)

	app.Post("/tasks", routes.CreateTask)

	app.Delete("/tasks/:id", routes.DeleteTaskById)

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
