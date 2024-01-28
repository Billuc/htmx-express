package api

import (
	components "fax/components"
	"fax/data/repos"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InitCards(group fiber.Router) {
	group.Get("/cards", getCards)
	group.Post("/cards", createCard)
	// group.Get("/cards/:id", getCard)
	// group.Put("/cards/:id", updateCard)
	// group.Delete("/cards/:id", deleteCard)
}

func getCards(c *fiber.Ctx) error {
	repo, err := repos.GetCardRepo()

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	cards, err := repo.GetManyCards()

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	components.CardListComponent(cards).Render(c.Response().BodyWriter())
	return c.SendStatus(200)
}

func createCard(c *fiber.Ctx) error {
	repo, err := repos.GetCardRepo()

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	card, err := repo.CreateCard(c.FormValue("title"))

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	components.CardComponent(*card).Render(c.Response().BodyWriter())
	return c.Type("html").SendStatus(200)
}
