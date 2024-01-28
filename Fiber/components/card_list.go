package components

import (
	"fax/data/models"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func CardListComponent(cards []models.Card) g.Node {
	return Div(
		Class("cards"),

		CreateCardComponent(),

		Div(Class("card-list"),
			g.Group(g.Map(cards, CardComponent)),
		),
	)
}
