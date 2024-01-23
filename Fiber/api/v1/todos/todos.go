package todos

import (
	"github.com/gofiber/fiber/v2"
)

func InitTodos(group fiber.Router) {
	group.Get("/todos", getTodos)
	group.Post("/todos", createTodo)
	group.Get("/todos/:id", getTodo)
	group.Put("/todos/:id", updateTodo)
	group.Delete("/todos/:id", deleteTodo)
}
