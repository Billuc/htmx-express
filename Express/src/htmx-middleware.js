import express from "express";
import * as fs from "node:fs/promises";
import * as cheerio from "cheerio";

/**
 *
 * @param {express.Request} req
 * @param {express.Response} res
 * @param {*} next
 */
export default async function htmxMiddleware(req, res, next) {
  if (!req.path.startsWith("/api")) {
    console.log(`Getting route from ${req.path}`);

    const fileName = fileNameFromPath(req.path);
    const data = await fs.readFile(fileName, { encoding: "utf-8" });

    const $ = cheerio.load(data);
    $("head").append(
      '<script src="https://unpkg.com/htmx.org@1.9.10"></script>'
    );
    $("head").append("<style hx-style></style>");

    res.send($.html());
  } else {
    next();
  }
}

function fileNameFromPath(path) {
  return `./routes${path.endsWith("/") ? path + "index" : path}.html`;
}
