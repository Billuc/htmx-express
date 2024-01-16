package component

import (
	"html/template"
	"strings"
)

type Component struct {
	name     string
	template template.Template
	style    string
	script   string
}

func (c *Component) Render(props map[string]any) string {
	builder := new(strings.Builder)
	c.template.Execute(builder, props)
	return builder.String()
}
