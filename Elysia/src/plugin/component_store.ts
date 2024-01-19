import { Component } from "./component";

export class ComponentStore {
  private static readonly _instance: ComponentStore = new ComponentStore();
  private _components: Map<string, Component<any>>;

  private constructor() {
    this._components = new Map();
  }

  public static get instance(): ComponentStore {
    return ComponentStore._instance;
  }

  public getComponent(name: string): Component<any> {
    const component = this._components.get(name);

    if (!component) {
      throw new Error(`Component ${name} not found`);
    }

    return component;
  }

  public addComponent(component: Component<any>) {
    if (this._components.has(component.name)) {
      throw new Error(`Component ${component.name} already exists`);
    }

    this._components.set(component.name, component);
  }
}
