import Elysia from "elysia";
import { ComponentStore } from "./component_store";

const serve_resource = async (
  path: string,
  headers: Record<string, string>
) => {
  const sanitizedPath = path.replace(/\/resources\//, "");
  console.log(`Serving resource for ${sanitizedPath}`);

  const store = ComponentStore.instance;
  const [componentName, resourceType] = sanitizedPath.split(".");
  const component = store.getComponent(componentName);

  if (resourceType === "css") {
    headers["Content-Type"] = "text/css";
    return component.style;
  } else if (resourceType === "js") {
    return component.script;
  }
};

export const serveResourcesPlugin = new Elysia().get(
  "/resources/*",
  ({ path, set }) => {
    const resource = serve_resource(path, set.headers);
    return resource;
  }
);
