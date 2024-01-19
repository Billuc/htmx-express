import { Component } from "./component";

export class ComponentRenderer {
  accumulatedStyles: Set<string>;
  accumulatedScripts: Set<string>;

  constructor() {
    this.accumulatedStyles = new Set();
    this.accumulatedScripts = new Set();
  }

  render<T>(component: Component<T>, props: T) {
    if (component.style) {
      this.accumulatedStyles.add(component.name);
    }

    if (component.script) {
      this.accumulatedScripts.add(component.name);
    }

    return component.render(props, this);
  }
}
