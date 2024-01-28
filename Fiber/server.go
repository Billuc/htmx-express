package main

import (
	"fax/api"
	"fax/components"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api.InitApi(app)

	app.Get("/cards", func(c *fiber.Ctx) error {
		components.Layout(components.CardsIndex()).Render(c.Response().BodyWriter())
		return c.Type("html").SendStatus(200)
	})

	app.Static("/assets", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Listen(":3000")
}
