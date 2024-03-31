package routes

import (
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func CreateResponseTask(taskModel models.Task) Task {
	return Task{
		ID:        int(taskModel.ID),
		Name:      taskModel.Name,
		CreatedAt: taskModel.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.Database.Db.Create(&task)
	responseTask := CreateResponseTask(task)

	return c.Status(201).JSON(responseTask)
}
