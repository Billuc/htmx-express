package components

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Layout(children ...g.Node) g.Node {
	return c.HTML5(
		c.HTML5Props{
			Title: "Component Index",
			Head: []g.Node{
				Script(Src("https://cdn.tailwindcss.com?plugins=typography")),
				Script(Src("https://unpkg.com/htmx.org@1.9.10")),
				Script(Src("https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"), Defer()),
			},
			Body: children,
		},
	)
}
