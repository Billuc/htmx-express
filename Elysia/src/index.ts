import { Elysia } from "elysia";
import { hteaPlugin } from "./plugin/htea_plugin";
import { todosApi } from "./api/todos";

const app = new Elysia().use(todosApi).use(hteaPlugin).listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
);
