package components

import (
	"fax/data/models"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func CardComponent(card models.Card) g.Node {
	return Div(
		Class("card"),

		H1(g.Text(card.Title)),
		P(g.Text(card.Description)),
	)
}
