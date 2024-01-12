import Component from "./base/component.js";

const Element = Component(
  "element",
  function (props) {
    return /* html */ `
        <div class="element">
            <span>${props.item.name}</span>
            <span>${props.item.age}</span>
        </div>
    `;
  },
  /* css */ `
    .element {
        display: flex;
        justify-content: space-between;
        background-color: #f5f5f5;
    }
    `
);

export default Element.render;
export const name = Element.name;
export const style = Element.style;
