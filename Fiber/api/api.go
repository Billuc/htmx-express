package api

import (
	"github.com/gofiber/fiber/v2"
)

func InitApi(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	InitCards(v1)
}
