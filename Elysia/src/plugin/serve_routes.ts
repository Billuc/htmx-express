import * as cheerio from "cheerio";
import Elysia from "elysia";

const serve_route = async (path: string) => {
  const sanitizedPath = path.replace(/\/$/, "");
  const filepath = `src/routes${sanitizedPath}/index.html`;
  console.log(`Serving route from ${filepath}`);

  const filedata = await Bun.file(filepath).text();

  const $ = cheerio.load(filedata);
  $("head").append('<script src="https://unpkg.com/htmx.org@1.9.10"></script>');
  $("head").append("<style hx-style></style>");

  return $.html();
};

export const serveRoutePlugin = new Elysia().get(
  "/*",
  async ({ path, set }) => {
    const html = await serve_route(path);
    set.headers["Content-Type"] = "text/html";
    return html;
  }
);
