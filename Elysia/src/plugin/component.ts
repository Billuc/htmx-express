import { ComponentRenderer } from "./component_renderer";

export interface ComponentDefinition<TProps> {
  style?: string;
  render: (
    props: TProps,
    renderer: ComponentRenderer
  ) => string | Promise<string>;
  script?: string;
}

export class Component<TProps> {
  name: string;
  style?: string;
  render: (
    props: TProps,
    renderer: ComponentRenderer
  ) => string | Promise<string>;
  script?: string;

  constructor(name: string, definition: ComponentDefinition<TProps>) {
    this.name = name;
    this.style = definition.style;
    this.render = definition.render;
    this.script = definition.script;
  }
}
