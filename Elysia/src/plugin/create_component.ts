import { Component } from "./component";
import { ComponentRenderer } from "./component_renderer";
import { ComponentStore } from "./component_store";

export default function createComponent<T>(
  name: string,
  componentDefinition: {
    style?: string;
    render: (props: T, renderer: ComponentRenderer) => string | Promise<string>;
    script?: string;
  }
): Component<T> {
  const component = new Component<T>(name, componentDefinition);

  const store = ComponentStore.instance;
  store.addComponent(component);

  return component;
}
