package api

import (
	v1 "fax/api/v1"

	"github.com/gofiber/fiber/v2"
)

func InitApi(app *fiber.App) {
	api := app.Group("/api")

	v1.InitV1(api)
}
