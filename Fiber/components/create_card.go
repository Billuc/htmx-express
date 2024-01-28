package components

import (
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

func CreateCardComponent() g.Node {
	return Div(
		g.Attr("x-data", "{ showModal: false }"),
		g.Attr("@keydown.window.escape", "showModal = false"),
		Class("modal"),

		Button(g.Attr("@click", "showModal = !showModal"), g.Text("New Card")),

		Div(
			g.Attr("x-cloak"),
			g.Attr("x-transition.opacity"),
			g.Attr("x-show", "showModal"),
			g.Attr("@click", "showModal = false"),
			Class("modal-backdrop"),
		),

		Div(
			g.Attr("x-cloak"),
			g.Attr("x-transition.opacity"),
			g.Attr("x-show", "showModal"),
			Class("modal-card"),

			Div(
				Class("modal-card-body"),

				FormEl(
					hx.Post("/api/v1/cards"),
					hx.Swap("beforeend"),
					hx.Target("next .card-list"),
					g.Attr("@submit", "showModal = false"),

					Div(
						Class("form-group"),

						Label(For("title"), g.Text("Title")),
						Input(
							Type("text"),
							Name("title"),
							ID("title"),
							Class("form-control"),
						),
					),

					Div(
						Class("form-actions"),

						Input(Type("submit"), Value("Create")),
					),
				),
			),
		),
	)
}
