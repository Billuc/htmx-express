import { Elysia } from "elysia";
import { serveRoutePlugin } from "./plugin/serve_routes";

const app = new Elysia()
  .get("/hello", () => "Hello Elysssia")
  .use(serveRoutePlugin)
  .listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
);
