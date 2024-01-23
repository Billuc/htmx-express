package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./components", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"Name": "World"}, "layout")
	})

	app.Listen(":3000")
}
