import { ComponentRenderer } from "../../plugin/component_renderer";
import createComponent from "../../plugin/create_component";

export const Todo = createComponent("todo", {
  style: /* css */ `
    .todo {
      padding: 4px 8px;
      background: #eee;
      border-radius: 0.5rem;
    }
  `,
  render(props: { text: string }, renderer: ComponentRenderer) {
    return <div class="todo">{props.text}</div>;
  },
});
