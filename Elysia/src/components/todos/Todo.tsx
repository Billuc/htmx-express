import { ComponentRenderer } from "../../plugin/component_renderer";
import createComponent from "../../plugin/create_component";

export const Todo = createComponent("todo", {
  render(props: { text: string }, renderer: ComponentRenderer) {
    return <div>{props.text}</div>;
  },
});
