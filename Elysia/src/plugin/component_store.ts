class ComponentStore {
  static readonly _instance = new ComponentStore();
  components: Record<string, Component> = {};

  private constructor() {
    this.components = {};
  }

  static getStore() {
    return ComponentStore._instance;
  }
}

class ComponentStoreInstance {
  components: Record<string, Component>;
  accumulatedComponents: Set<string>;

  constructor(components: Record<string, Component>) {
    this.components = components;
    this.accumulatedComponents = new Set();
  }

  get(component: string) {
    return this.components[component];
  }
}

function createComponent<T>(
  name: string,
  componentDefinition: {
    style?: string;
    render: (props: T, store: ComponentStore) => string | Promise<string>;
    script?: string;
  }
): (props: T) => Promise<string> {
  const store = ComponentStore.getStore();
  store.components[name] = componentDefinition;

  // Add accumulated styles
  return async (props: T) => await componentDefinition.render(props, store);
}
