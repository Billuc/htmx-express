import Element from "./element.js";
import each from "../src/each.js";
import Component from "./base/component.js";

const List = Component("list", function (props) {
  return /* html */ `
            <div>
                ${each(props.items, (item) => Element({ item }))}
            </div>
        `;
});

export default List.render;
