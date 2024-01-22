import { Elysia } from "elysia";
import { html } from "@elysiajs/html";
import htmx_render from "./htmx_render";

export function createHteaApi() {
  return new Elysia().use(html()).decorate("render", htmx_render);
}
