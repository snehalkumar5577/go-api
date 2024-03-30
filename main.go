package main

import (
	"log"
	"main/database"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	database.Connect()
	app := fiber.New()

	app.Get("/", welcome)

	log.Fatal(app.Listen(":3000"))
}
