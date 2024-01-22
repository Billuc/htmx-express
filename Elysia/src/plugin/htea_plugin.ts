import { Elysia } from "elysia";
import { serveResourcesPlugin } from "./serve_resources";
import { serveRoutesPlugin } from "./serve_routes";

export const hteaPlugin = new Elysia()
  .use(serveResourcesPlugin)
  .use(serveRoutesPlugin);
