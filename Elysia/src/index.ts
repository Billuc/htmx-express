import { Elysia } from "elysia";
import { serveRoutePlugin } from "./plugin/serve_routes";
import htmx_render from "./plugin/htmx_render";
import { serveResourcesPlugin } from "./plugin/serve_resources";
import { TodoList } from "./components/todos/TodoList";

const app = new Elysia()
  .get("/hello", () => "Hello Elysssia")
  .decorate("render", htmx_render)
  .get("/api/todos", ({ render }) =>
    render(TodoList, { todos: [{ text: "todo1" }, { text: "todo2" }] })
  )
  .use(serveResourcesPlugin)
  .use(serveRoutePlugin)
  .listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
);
