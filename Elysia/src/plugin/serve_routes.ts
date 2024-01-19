import * as cheerio from "cheerio";
import Elysia from "elysia";
import { html } from "@elysiajs/html";

const serve_route = async (path: string) => {
  const sanitizedPath = path.replace(/\/$/, "");
  const filepath = `src/routes${sanitizedPath}/index.html`;
  console.log(`Serving route from ${filepath}`);

  const filedata = await Bun.file(filepath).text();

  const $ = cheerio.load(filedata);
  $("head").append('<script src="https://unpkg.com/htmx.org@1.9.10"></script>');
  $("head").append(
    '<script src="https://unpkg.com/htmx.org/dist/ext/head-support.js"></script>'
  );

  $("body").attr("hx-ext", "head-support");

  return $.html();
};

export const serveRoutePlugin = new Elysia()
  .use(html())
  .get("/*", async ({ path, set }) => {
    const html = await serve_route(path);
    return html;
  });
