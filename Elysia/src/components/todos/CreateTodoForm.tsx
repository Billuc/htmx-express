import { ComponentRenderer } from "../../plugin/component_renderer";
import createComponent from "../../plugin/create_component";

export const CreateTodoForm = createComponent("create-todo-form", {
  style: /* css */ `
        .create-todo-form {
          text-align: center;
        }
    `,
  render(props: {}, renderer: ComponentRenderer) {
    return (
      <form
        hx-post="/api/todos"
        hx-target="previous .todo-list"
        hx-swap="beforeend"
        class="create-todo-form"
      >
        <input name="title" />
        <input type="submit" value="New todo" />
      </form>
    );
  },
});
