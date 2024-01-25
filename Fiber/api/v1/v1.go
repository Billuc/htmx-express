package v1

import (
	"fax/api/v1/cards"

	"github.com/gofiber/fiber/v2"
)

func InitV1(group fiber.Router) {
	v1 := group.Group("/v1")

	cards.InitCards(v1)
}
