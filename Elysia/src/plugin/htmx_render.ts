import { Component } from "./component";
import { ComponentRenderer } from "./component_renderer";

export default function htmx_render<TProps>(
  component: Component<TProps>,
  props: TProps
) {
  const renderer = new ComponentRenderer();

  const html = renderer.render(component, props);
  const accumulatedStyles = [...renderer.accumulatedStyles.values()];

  const styleHead = `
    <head hx-head="merge">
        ${accumulatedStyles
          .map(
            (n) =>
              `<link rel="stylesheet" href="/resources/${n}.css" hx-preserve="true"/>`
          )
          .join("\n")}
    </head>
  `;

  return styleHead + html;
}
