import { ComponentRenderer } from "../../plugin/component_renderer";
import createComponent from "../../plugin/create_component";
import { Todo } from "./Todo";

export const TodoList = createComponent("todo-list", {
  style: /* css */ `
        .todo-list {
            display: flex;
            flex-direction: column;
            background-color: red;
            padding: 20px;
        }
    `,
  render(props: { todos: { text: string }[] }, renderer: ComponentRenderer) {
    const todoItems = props.todos.map((todo) => renderer.render(Todo, todo));

    return <div class="todo-list">{todoItems}</div>;
  },
});
