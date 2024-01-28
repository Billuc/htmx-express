package components

import (
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

func CardsIndex() g.Node {
	return Div(
		Div(hx.Get("/api/v1/cards"), hx.Trigger("load")),
	)
}
