import express from "express";
import List from "./components/list.js";
import htmxMiddleware from "./src/htmx-middleware.js";
import registerHtmxComponent from "./src/register-htmx-component.js";
import * as element from "./components/element.js";

const app = express();
const port = 3002;

const items = [
  { name: "Toto", age: 8 },
  { name: "Dad", age: 42 },
];

app.use(htmxMiddleware);
app.use(registerHtmxComponent(element));

app.get("/api/elements", (req, res) => {
  res.send(List({ items }));
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
