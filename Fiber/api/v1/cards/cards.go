package cards

import (
	"fax/data/repos"

	"github.com/gofiber/fiber/v2"
)

func InitTodos(group fiber.Router) {
	group.Get("/cards", getCards)
	// group.Post("/cards", createCard)
	// group.Get("/cards/:id", getCard)
	// group.Put("/cards/:id", updateCard)
	// group.Delete("/cards/:id", deleteCard)
}

func getCards(c *fiber.Ctx) error {
	repo, err := repos.GetCardRepo()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	cards, err := repo.GetManyCards()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Render("cards/card-list", fiber.Map{"Cards": cards})
}
