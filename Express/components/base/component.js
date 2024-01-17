import * as cheerio from "cheerio";

export default function Component(name, htmlFn, style, script) {
  function render(props) {
    const $ = cheerio.load(htmlFn(props));

    if (style) {
      $.root().append(/* html */ `<div 
            hx-trigger="revealed once" 
            hx-get="/api/style/${name}.css" 
            hx-swap="beforeend" 
            hx-target="style[hx-style]"
            ></div>`);
    }

    return $.html();
  }

  return {
    render,
    style,
    name,
  };
}
