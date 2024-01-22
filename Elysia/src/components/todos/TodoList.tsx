import { ComponentRenderer } from "../../plugin/component_renderer";
import createComponent from "../../plugin/create_component";
import { Todo } from "./Todo";
import { CreateTodoForm } from "./CreateTodoForm";

export const TodoList = createComponent("todo-list", {
  style: /* css */ `
        .todo-list {
            display: flex;
            flex-direction: column;
            padding: 4px;
            gap: 4px;
        }
    `,
  render(props: { todos: { text: string }[] }, renderer: ComponentRenderer) {
    const todoItems = props.todos.map((todo) => renderer.render(Todo, todo));

    return (
      <div>
        <div class="todo-list">{todoItems}</div>
        {renderer.render(CreateTodoForm, {})}
      </div>
    );
  },
});
