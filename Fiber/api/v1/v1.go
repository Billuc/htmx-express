package v1

import (
	"github.com/gofiber/fiber/v2"
)

func InitV1(group fiber.Router) {
	group.Group("/v1")
}
